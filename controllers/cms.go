package controllers

import (
	"web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CmsController struct {
	beego.Controller
}

//获取主页数据
func (this *CmsController) Get() {

	//有Orm对象
	// o := orm.NewOrm()
	// //有一个要插入数据的结构体对像
	// user := models.User{}
	// //对结构体赋值
	// user.Name = "11111"
	// //插入
	// id, err := o.Insert(&user)
	// if err != nil {
	// 	beego.Info("插入失败", err)
	// 	return
	// }
	// beego.Info("插入成功di为：：", id)

	//查询
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "bob"
	// err := o.Read(&user, "Name")
	// if err != nil {
	// 	beego.Info("查询失败", err)
	// 	return
	// }
	// beego.Info("查询成功", user)

	//更新
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Id = 1
	// err := o.Read(&user)
	// if err == nil{
	// 	user.Name="cool"
	// 	id,err := o.Update(&user)
	// 	if err != nil{
	// 		beego.Info("更新失败",err)
	// 	}
	// 	beego.Info("更新成功",id)
	// }

	//删除
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 2
	id, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除错误", err)
	}
	beego.Info("删除成功", id)

}
