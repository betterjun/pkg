/*
	包cfg，配置文件管理模块，定义配置文件数据读取接口。
*/
package cfg

import (
	"github.com/betterjun/go-toml"
)

// 配置数据
var tree *toml.TomlTree

// 载入配置文件。
func LoadConfig(file string) (err error) {
	tree, err = toml.LoadFile(file)
	if err != nil {
		return err
	}

	return nil
}

// 获取字符串配置数据。
func GetString(key string, def ...string) (ret string) {
	return tree.GetString(key, def...)
}

// 获取64位整数配置数据。
func GetInt64(key string, def ...int64) (ret int64) {
	return tree.GetInt64(key, def...)
}

// 获取32位整数配置数据。
func GetInt32(key string, def ...int64) (ret int32) {
	return int32(tree.GetInt64(key, def...))
}

// 获取整数配置数据。
func GetInt(key string, def ...int64) (ret int) {
	return int(tree.GetInt64(key, def...))
}

// 获取64位浮点数配置数据。
func GetFloat64(key string, def ...float64) (ret float64) {
	return tree.GetFloat64(key, def...)
}

// 获取通用配置项，需要再次转换。
func Get(key string) interface{} {
	return tree.Get(key)
}
