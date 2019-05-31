package blog

import(
  "github.com/astaxie/beego/logs"
  "github.com/astaxie/beego"
  "practice/models"
  "strings"
)

type AccountConterller struct {
  baseController
}

func (this *AccountConterller) Login(){
  this.display("accounts", "login")
}

func (this *AccountConterller) LoginCreate(){
  var user models.User
  email := this.GetString("email")
  password := this.GetString("password")
  password = user.EncodeMd5(strings.ToLower(password))
  user.Email = email
  user.Password = password
  if err := user.Read("email", "password"); err != nil {
    logs.Info(err)
    this.Data["errMsg"] = "账号或密码错误"
    this.display("accounts", "login")
    return
  }
  this.Ctx.SetSecureCookie(beego.AppConfig.String("sercet") ,"authgo", user.Email, 7200*3600)
  this.Redirect("/",302)
  return 
}

func (this *AccountConterller) New() {
  this.display("accounts", "new")
}

func (this *AccountConterller) SignOut(){
  this.Ctx.SetSecureCookie(beego.AppConfig.String("sercet") ,"authgo", "")
  this.Redirect("/",302)
  return
}

func (this *AccountConterller) Create() {
  var user models.User
  email := this.GetString("email")
  logs.Info("email 是", email)
  username := this.GetString("username")
  password := this.GetString("password")
  u, err := user.Register(email, password, username)
  logs.Info(len(err.Errors))
  logs.Info(u)
  if len(err.Errors) != 0 {
    for _, errs := range err.Errors {
      logs.Info(errs.Message)
      this.Data["errMsg"] = errs.Message
     }
    this.display("accounts", "new")
    return
  }
  this.Ctx.SetSecureCookie(beego.AppConfig.String("sercet") ,"authgo", u.Email, 7200*3600)
  logs.Info("结果是")
  this.Redirect("/",302)
  return 
}