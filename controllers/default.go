package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"bulletin_board/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout.html"
	c.TplName = "default/index.html"

	o := orm.NewOrm()
	o.Using("default")

	var products []*models.Products
	num, err := o.QueryTable("products").All(&products)

	if err != orm.ErrNoRows && num > 0 {
		c.Data["Records"] = products
	}
}
