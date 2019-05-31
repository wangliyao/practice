package blog

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/logs"
  "practice/models"
)

type baseController struct {
  beego.Controller
}

func (this *baseController) display(tpl ...string) {
  var tplName string
  layout := [3]string{"layout.html", "header.html", "footer.html"}
  if len(tpl) == 1 {
    tplName = tpl[0] + ".html"
  }else{
    tplName = tpl[0] + "/" + tpl[1] + ".html"
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


func (this *baseController)Prepare(){
  this.Data["prependcurrentUser"] = this.prependcurrentUser()
  this.Data["IsLogin"] = this.IsLogin()

}

func (this *baseController)prependcurrentUser() string {
  email, _:= this.Ctx.GetSecureCookie(beego.AppConfig.String("sercet") ,"authgo")
  return email
}

func (this *baseController ) IsLogin() bool {
  email, ok:= this.Ctx.GetSecureCookie(beego.AppConfig.String("sercet") ,"authgo")
  logs.Info("email是", email)
  if ok {
    var user models.User
    user.Email = email
    if err := user.Read("email"); err == nil {
      return true
    }
    return false
  } else {
    return false
  }
}