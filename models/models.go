package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Products struct {
	Id     int    `form:"-"`
	Name   string `form:"name,text,name:" orm:"size(50)" valid:"MinSize(1);MaxSize(50)"`
	Price  int    `form:"price,int,price:"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Products))
}
