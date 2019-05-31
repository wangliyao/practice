package routers

import (
	"github.com/astaxie/beego"
  "practice/controllers/blog"
  "practice/controllers/admin"
  "github.com/astaxie/beego/context"
  "github.com/astaxie/beego/logs"
)

func init() {

  var FilterUser = func(ctx *context.Context) {
    val, _ := ctx.GetSecureCookie(beego.AppConfig.String("sercet") ,"authgo")
    logs.Info("cookies", val)
    if val == "" {
        ctx.Redirect(302, "/signin")
    }
  }

    beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)
    beego.Router("/", &blog.MainController{}, "*:Index")
    beego.Router("/404.html", &blog.MainController{}, "*:Go404")
    beego.Router("/signin", &blog.AccountConterller{}, "get:Login;post:LoginCreate")
    beego.Router("/signup", &blog.AccountConterller{}, "*:New")
    beego.Router("/register", &blog.AccountConterller{}, "post:Create")
    beego.Router("/signout", &blog.AccountConterller{}, "*:SignOut")


    ns := beego.NewNamespace("/admin",
      beego.NSRouter("/", &admin.MainController{}, "*:Index"),
      beego.NSRouter("/content", &admin.MainController{}, "*:Content"),
      beego.NSRouter("/posts", &admin.PostController{}, "*:Index;post:Create"),
      beego.NSRouter("/posts/new", &admin.PostController{}, "*:New"),
      beego.NSRouter("/posts/:id", &admin.PostController{}, "get:Edit"),
      beego.NSRouter("/posts/:id/delete", &admin.PostController{}, "*:Delete"),
      beego.NSRouter("/posts/:id/update", &admin.PostController{}, "*:Update"),
      beego.NSRouter("/accounts", &admin.AccountsController{}, "*:Index;post:Create"),
      beego.NSRouter("/images",  &admin.ImageController{}, "post:Create"),
    )

    beego.AddNamespace(ns)
}
