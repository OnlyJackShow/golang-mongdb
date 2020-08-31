package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//"gopkg.in/mgo.v2/bson"

	"fmt"
)

var (
	Session *mgo.Session
	SE=Sinit()
)

type City struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Id   int           `bson:"id"`
	Name string        `bson:"name"`
}

func Sinit() *mgo.Session {
	link := "mongodb://10.1.1.69"
	Session, err := mgo.Dial(link)
	if err != nil {
		fmt.Println("link err:", err.Error())
	}
	var lo mgo.Credential
	lo.Username = "admin"
	lo.Password = "123456"
	err = Session.Login(&lo) //如果有账号和密码的需要验证
	if err != nil {
		fmt.Println("身份验证失败：", err.Error())
	}
	return Session
}

func main() {
	/*c := SE.DB("admin").C("hello") //选择ChatRoom库的account表
	var rest []City
	c.Find(nil).All(&rest)
	fmt.Println(rest)*/
	//Insert()
	Search()
}

//mongdb的插入
func Insert()  {
	for i:=2;i<10000000000;i++{
		var rest City
		rest.Id=i
		rest.Name=fmt.Sprintf("%s%d","name",i)
		SE.DB("admin").C("hello").Insert(&rest) //选择数据库和表
	}
}

//mongdb的查询
func Search()  {
	var result City
	SE.DB("admin").C("hello").Find(bson.M{"id":16451885}).One(&result)
	//SE.DB("admin").C("hello").Find(nil).All(&result)  //查询所有
	fmt.Println(result)
}