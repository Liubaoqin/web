package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户表
type User struct {
	Id   int
	Name string
}

//内容表
type Content struct {
	Id      int
	Title   string
	Content string
	Pic     string
	Addtime int
	Zan     int
}

//评价表
type Evaluate struct {
	Id      int
	Uid     int
	Content string
	Addtime int
	Zan     int
}

func init() {
	driverName := beego.AppConfig.String("driverName")

	// 注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	// 数据库连接
	user := beego.AppConfig.String("mysqlUser")
	pass := beego.AppConfig.String("mysqlPass")
	host := beego.AppConfig.String("host")
	dbname := beego.AppConfig.String("dbName")
	port := beego.AppConfig.String("port")

	dbconn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	fmt.Println(dbconn)
	err := orm.RegisterDataBase("default", driverName, dbconn)
	orm.RegisterModel(new(User), new(Content), new(Evaluate))
	orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("mysql连接失败")
	}
	fmt.Println("mysql连接成功")
}
