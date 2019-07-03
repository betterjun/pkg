package main

import (
	"fmt"

	"github.com/betterjun/pkg/mongo"
	"gopkg.in/mgo.v2/bson"
)

// type Person struct {
// 	Id    bson.ObjectId `bson:"_id"`
// 	Name  string        `bson:"tname"` //bson:"name" 表示mongodb数据库中对应的字段名称
// 	Phone string        `bson:"tphone"`
// }

type Person struct {
	Id    bson.ObjectId `bson:"_id"`
	Name  string
	Phone string
}

/**
 * 添加person对象
 */
func AddPerson3(p Person) string {
	session := mongodb.GetClonedSession()
	defer session.Close()
	c := session.Collection("person")

	p.Id = bson.NewObjectId()
	err := c.Insert(p)
	if err != nil {
		fmt.Println(err)
		return "false"
	}
	return p.Id.Hex()
}

func main() {
	var p Person

	// // 登录方式跟数据库配置有关系
	// // 用airuser登录，鉴权数据库必须是air才可以成功
	// //err := mongodb.InitDefault("47.99.46.207", "airuser", "airuser2019", 10, 10)
	// // 用admin登录，所有数据库都可以操作
	// err := mongodb.InitDefault("47.99.46.207", "admin", "nhpw2019", 10, 10)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 在哪个数据库登录，所有操作都在该数据库
	//err := mongodb.InitDefaultDBSession("47.99.46.207", "admin", "admin", "nhpw2019", 10, 10)
	err := mongodb.InitDefaultDBSession("47.99.46.207", "air", "airuser", "airuser2019", 10, 10)
	//err := mongodb.InitDefaultDBSession("47.99.46.207", "air", "admin", "nhpw2019", 10, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Name = "3"
	AddPerson3(p)
}
