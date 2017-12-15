package mysql

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// // 需要在init中注册定义的model
	// //orm.RegisterModel(new(AccountInfo), new(UserInfo))

	// orm.RegisterDriver("mysql", orm.DRMySQL)

	//连接池配置
	//最大空闲连接数
	maxIdle := beego.AppConfig.DefaultInt("Mysql::MaxIdle", 6)

	//最大数据库连接数
	maxConn := beego.AppConfig.DefaultInt("Mysql::MaxConnent", 30)

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("Mysql::MysqlConn"), maxIdle, maxConn)
	orm.RegisterDataBase("ShowGirlSlave", "mysql", beego.AppConfig.String("Mysql::MysqlConnSlave"), maxIdle, maxConn)

}
