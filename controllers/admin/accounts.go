package admin

import(
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"practice/models"
)

type AccountsController struct {
	baseController
}

func (this *AccountsController) Index(){
	user := this.FindUser()
	logs.Info(user)
	result := user.FindImage()
	logs.Info("image结果", result)
	this.Data["avatarUrl"] = "/static/upload/" + result.FileName
	this.Data["user"] = user
  this.TplName = "admin/accounts/index.html"
}


func (this *AccountsController) Create(){
	var account models.User
	var image models.Image
	user := this.FindUser()
	o := orm.NewOrm()
  qs := o.QueryTable("user")
  qs.Filter("id", user.Id).One(&account)
  nickname := this.GetString("username")
  email := this.GetString("email")
	password := this.GetString("password")
	images := image.GetImage(this.GetString("avatarId"))
	logs.Info("头像图片位", images)
	account.Email = email
	account.Nickname = nickname
	account.Avatar = images

	// account.Avatar = this.GetString("avator")
	if password != "" {
		account.Password = 	account.EncodeMd5(password)
	}
	_, err := o.Update(&account)
	if err != nil {
		this.Data[err] = "更新失败"
	}
	result := account.FindImage()
	logs.Info("image结果", result)
	this.Data["avatarUrl"] = "/static/upload/" + result.FileName
	this.Data["user"] = account
	
  this.TplName = "admin/accounts/index.html"
}