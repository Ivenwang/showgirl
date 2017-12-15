package mysql

import (
	"github.com/astaxie/beego/orm"
)

func NewOrm() orm.Ormer {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("ShowGirl")

	return o
}

func NewShowgirlOrm() orm.Ormer {
	orm.Debug = true
	o := orm.NewOrm()
	return o
}

func NewShowgirlSlaveOrm() orm.Ormer {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("ShowGirl")
	return o
}
