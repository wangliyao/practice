package models

import (
  "time"
  "github.com/astaxie/beego/logs"
  "crypto/md5"
  "encoding/hex"
  "strings"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"

)

type User struct {
  Username string `orm:"unique;size(15);index;"`
  Password string `orm:"size(32);"`
  Nickname string `orm:"size(15);index;null"`
  Email string `orm:"size(50);index;"`
  Lastlogin  time.Time `orm:"auto_now_add;type(datetime);index;null"`
  Permission string `orm:"size(100);index;null"`
  Avatar *Image `orm:"rel(one);null"`
  Post []*Post `orm:"reverse(many);null"`
  Comments []*Comments `orm:"reverse(many);null"`
  BaseModel
}


func (user *User) Read(fields ...string) error {
  if err := orm.NewOrm().Read(user, fields...); err != nil {
    return err 
  }
  return nil
}

func (u User)EncodeMd5(str string) (md5Digest string){
  password := []byte(str)
  result := md5.Sum(password)
  md5Digest = hex.EncodeToString(result[:])
  return
}
 

func (u User) Register(email string, password string, username string) (user *User, v validation.Validation) {
  logs.Info("开始注册！！！！")
  o := orm.NewOrm()
  u.Email = strings.ToLower(strings.TrimSpace(email))
  logs.Info("密码:", password)
  u.Password = u.EncodeMd5(strings.ToLower(password))
  logs.Info("encode密码", u.Password)
  u.Username = strings.ToLower(strings.TrimSpace(username))
  u.Nickname = "Uber"
  valid := validation.Validation{}
  valid.Required(u.Username, "username").Message("用户名不存在")
  valid.MaxSize(u.Username, 15,"nameMax").Message("用户名最大长度15位")
  valid.Required(u.Password, "password").Message("密码不存在")
  valid.Required(u.Email, "email").Message("邮箱不存在")
  valid.Email(u.Email, "emailVali").Message("邮箱格式不正确")
  if valid.HasErrors(){
    return &u, valid
  }

  logs.Info("对象是", u)
  logs.Info("验证是", valid)

  if _,err := o.Insert(&u); err != nil {
    logs.Info(err)
    valid.SetError("user", "保存失败")
  }
  
  return &u, valid

}

func (user *User)FindImage() *Image {
  o := orm.NewOrm()
  if user.Avatar != nil {
    o.Read(user.Avatar)
  }
  logs.Info("结果为", user.Avatar)
  return user.Avatar
}