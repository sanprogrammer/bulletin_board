package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"bulletin_board/models"
	"strconv"
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

func (c *MainController) Show() {
	o := orm.NewOrm()
	o.Using("default")

	// convert the string value to an int
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	// read one
	product := models.Products{Id: id}
	err := o.Read(&product)
	if err != orm.ErrNoRows {
		c.Data["Name"]    = product.Name
		c.Data["Price"]   = product.Price
		c.Data["Created"] = product.Created
	}

	c.Layout = "layout.html"
	c.TplName = "default/show.html"
}
