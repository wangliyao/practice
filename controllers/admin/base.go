package admin

import (
  "github.com/astaxie/beego/orm"
  "practice/models"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/logs"
)

type baseController struct {
  beego.Controller
}

func (this *baseController) display(tpl ...string) {
  var tplName string
  layout := [3]string{"admin/layout.html", "admin/header.html", "admin/footer.html"}
  if len(tpl) == 1 {
	tplName = "admin/" + tpl[0] + ".html"
	logs.Info(tplName)
  }else{
    tplName ="admin/" + tpl[0] + "/" + tpl[1] + ".html"
    layout[1] = ""
    layout[2] = ""
  }
  this.Layout = layout[0]
  this.TplName = tplName
  this.LayoutSections = make(map[string]string)
  this.LayoutSections["header"] = layout[1]
  this.LayoutSections["footer"] = layout[2]
  this.LayoutSections["SideBar"] = ""
}

func (this *baseController) FindUser() *models.User {
  var user models.User
  email := this.prependcurrentUser()
  user.Email = email
  orm.NewOrm().Read(&user, "email")
  logs.Info("结果是", user.Id)
  return &user
}


func (this *baseController)Prepare(){
  this.Data["prependcurrentUser"] = this.prependcurrentUser()
}

func (this *baseController)prependcurrentUser() string {
  email, _:= this.Ctx.GetSecureCookie(beego.AppConfig.String("sercet") ,"authgo")
  return email
}