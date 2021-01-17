package common
//https://gitee.com/y2h/beeqor/blob/master/initial/init.go
import (
	"database/sql/driver"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/gorm"
	"reflect"
	"regexp"
	"smartapp/initial"
	"strconv"
	"strings"
	"time"
	"unicode"

	//"github.com/gogather/com"
	_ "github.com/go-sql-driver/mysql"
)
var (
	log                      = logs.NewLogger()

	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)
type Logger struct {
}
func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
var LogFormatter = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			source          = fmt.Sprintf("(%v)", values[1])
		)

		messages = []interface{}{source}

		if len(values) == 2 {
			//remove the line break
			//remove the brackets
			source = fmt.Sprintf(" %v ", values[1])

			messages = []interface{}{ source}
		}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf(" [%.2fms] ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						if t.IsZero() {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", "0000-00-00 00:00:00"))
						} else {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
						}
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
						default:
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength&&formattedValues!=nil {

						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf(" \n[%v] ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned "))
		} else {
			messages = append(messages, " ")
			messages = append(messages, values[2:]...)
			messages = append(messages, " ")
		}
	}

	return
}
func (logger *Logger) Print(values ...interface{}) {
	//if len(values) > 1{
   //
	//}
   //var (
   //  level           = values[0]
   //   source          = values[1]
   //)
   //if level == "sql" {
   //	 affected:=fmt.Sprintf(" \n[%v]", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned ")
   //  sql := values[3].(string)
   //   beego.Trace(sql, level, source)
   //} else {
	// beego.Trace(values)
   //}
    formatter:=LogFormatter(values...)
	log.Notice(generateFmtStr(len(formatter)),formatter...)

}
func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
func InitSqlLog(){

	err := log.SetLogger(logs.AdapterFile, `
   {
   "filename": "logs/sql/sql.log",
   "level": 7,
   "maxlines": 0,
   "maxsize": 0,
   "daily": true,
   "maxdays": 10,
   "color": true
	}`)
	if err!=nil{
		panic("sql日志出现错误:"+err.Error())
	}
	log.EnableFuncCallDepth(false) //输出文件名和行号
	if beego.AppConfig.String("runmode") == "dev" {
		_ = log.SetLogger(logs.AdapterConsole)
	}else{
		log.Async()
		//logs.SetLogFuncCallDepth(3)    //
	}
	log.SetLevel(logs.LevelInfo)
}
func InitSql()  {
	return
	InitSqlLog()
	user := beego.AppConfig.String("mysql::user")
	pwd := beego.AppConfig.String("mysql::pwd")
	host := beego.AppConfig.String("mysql::host")
	port, err := beego.AppConfig.Int("mysql::port")
	dbname := beego.AppConfig.String("mysql::dbname")
	if nil != err {
		port = 3306
	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	//if com.FileExist("install.lock") { //么有安装就暂时不指定数据库
	initial.DB,err = gorm.Open( "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd, host, port, dbname))
	//} else {
	//	err = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8", user, passwd, host, port))
	//}
	if err!=nil{
		panic("连接数据库服务失败:"+err.Error())
	}
	// 全局禁用表名复数
	initial.DB.SingularTable(true)
	initial.DB.SetLogger(&Logger{})
	// 启用Logger,显示详细日志

	if ok,_:=beego.AppConfig.Bool("mysql::sql_log");ok == true{
		initial.DB.LogMode(true)
	}else{
		beego.Debug("数据库日志记录功能关闭 来源于配置 sql_log \n",ok,ok)
	}



}
