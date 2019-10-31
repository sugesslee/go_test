package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	//"labix.org/v2/mgo/bson"
)

const (
	MONGODB_URL = "112.74.36.53:37017"
)

func main() {
	//创建连接
	session, err := mgo.Dial(MONGODB_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// db := session.DB("xtest")   //数据库名称
	// collection := db.C("xtest") // 集合名称
	c := session.DB("test").C("test")
	// //插入数据
	err = c.Insert(&Person{"Tommy", "123456"}, &Person{"Hanleilei", "98765"},
		&Person{"喜洋洋", "98765"}, &Person{"灰太狼", "46577"},
	)
	if err != nil {
		panic(err)
	}
	//查询并赋值 Find().One()
	result := Person{}
	err = c.Find(bson.M{"name": "Tommy"}).One(&result)
	if err != nil {
		panic(err)
	}
	//输出
	fmt.Println("Phone ", result.Phone)
	// //集合中元素数量 Count()
	// countNum, err := c.Count()
	// fmt.Println("obj numbers ", countNum)
	// //查询多条数据 Find().Iter()
	// var onep = Person{}
	// iter := c.Find(nil).Iter()
	// for iter.Next(&onep) {
	// 	fmt.Println("姓名 ", onep.Name)
	// }
	// //查询多条数据 Find().All()
	// var personAll []Person
	// err = c.Find(nil).All(&personAll)
	// for i := 0; i < len(personAll); i++ {
	// 	fmt.Println("Person ", personAll[i].Name, personAll[i].Phone)
	// }
	// //更新数据 Update()
	// abc := Person{}
	// err = c.Find(bson.M{"name": "Tommy"}).One(&abc)
	// fmt.Println("Tommy phone is ", abc.Phone)
	// err = c.Update(bson.M{"name": "Tommy"}, bson.M{"$set": bson.M{"phone": "10086"}})
	// err = c.Find(bson.M{"name": "Tommy"}).One(&abc)
	// fmt.Println("Tommy phone is ", abc.Phone)
	// //删除数据 Remove()
	// fmt.Println(c.Count())
	// err = c.Remove(bson.M{"phone": "46577"})
	// fmt.Println(c.Count())
	fmt.Println("end")
}

type Person struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}
