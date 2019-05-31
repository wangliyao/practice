package models

import (
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego"
  _ "github.com/go-sql-driver/mysql"
  "time"
)

type BaseModel struct {
  Id        int32
  DeletedAt *time.Time `orm:"type(datetime);null"`
}

func init() {
  dbuser := beego.AppConfig.String("dbuser")
  dbport := beego.AppConfig.String("dbport")
  dbpassword := beego.AppConfig.String("dbpassword")
  dbname := beego.AppConfig.String("dbname")
  dbhost := beego.AppConfig.String("dbhost")

  dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
  orm.RegisterDataBase("default", "mysql", dburl)

  orm.RegisterModel(new(User), new(Comments), new(Post), new(Image))
	orm.RunSyncdb("default", false, true)

  if beego.AppConfig.String("runmode") == "dev" {
    orm.Debug = true
  }

}