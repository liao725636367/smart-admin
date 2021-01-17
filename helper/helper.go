package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//func Query(query string, args ...interface{})orm.RawSeter{
//	//o :=orm.NewOrm()
//	//r := o.Raw(query,...args)
//}
//var goSrcRegexp = regexp.MustCompile(`jinzhu/gorm(@.*)?/.*.go`)
//var goTestRegexp = regexp.MustCompile(`jinzhu/gorm(@.*)?/.*test.go`)
/**
获取编译时路径 自动根据文件名称映射模板
*/
//根据编译时的调用栈回溯去找到对应的 文件代码路径
func GetControllerStackFile(skip int) string {
	_,fileSkip,_,_:=runtime.Caller(skip)
	for i := skip+1; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		//if ok && (!goSrcRegexp.MatchString(file) || goTestRegexp.MatchString(file)) {
		if ok &&  file != fileSkip&&strings.Contains(file,"controllers") { //判断在controllers目录下的控制器才能自动定位模板
			beego.Alert("回溯文件对比",file,fileSkip)
			return file
		}
	}
	return ""
}

// aes加密函数
//text  加密字符串
// 密钥 key只能为16位、24位或32位，分别对应AES-128, AES-192和 AES-256
func Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}
// aes解密函数
//text  加密字符串
// 密钥 key只能为16位、24位或32位，分别对应AES-128, AES-192和 AES-256
func Decrypt(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv= key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
//密码加密算法
func EncodePwd(pwd string,salt string )string {
	pwd =pwd +"_"+salt
	hashPwd,_:=bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.DefaultCost)
	return string(hashPwd)
}
//密码检查
func CheckPwd(pwd string,salt string,inputPwd string)bool{
	inputPwd =inputPwd +"_"+salt
	password,_:=bcrypt.GenerateFromPassword([]byte(inputPwd),bcrypt.DefaultCost)
	beego.Error("加密的密码:",string(password))

	err :=bcrypt.CompareHashAndPassword([]byte(pwd),[]byte(inputPwd))
	if err==nil{ //err为nil代表匹配
		return true
	}else{
		return false
	}
}

//切片去重转map
//切片转 map
func SliceUniqueMap(slice []string) map[string]string {
	maps := make(map[string]string)
	for index,item:=range slice{
		maps[item]=strconv.Itoa(index)
	}
	return maps
}
func TimeStr(values ... interface{} )string {
	return time.Now().Format("2006-01-02 15:04:05")
}
