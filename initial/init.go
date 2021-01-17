package initial

import (
	"github.com/astaxie/beego/cache"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"time"
)

var DB                       *gorm.DB //db全局对象
var Cache	cache.Cache
var CacheTime time.Duration
var Json jsoniter.API
type Map = map[string]interface{}

type MenusMap   []map[string]interface{}
//自定义排序
func (c MenusMap) Len() int {
	return len(c)
}
func (c MenusMap) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//元素排序判断
func (c MenusMap) Less(i, j int) bool {
	si, ok := c[i]["sort"].(int)
	if !ok {
		si = 0
	}
	sj, okj := c[j]["sort"].(int)
	if okj {
		sj = 0
	}
	return si < sj
}

func SliceFirst(slice []interface{})interface{}{
	return slice[0]
}
