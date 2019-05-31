package models

import (
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/logs"
  "github.com/astaxie/beego/validation"
)

type Post struct {
  User *User `orm:"rel(fk);index"`
  Title string `orm:"size(100);index"`
  Content string `orm:"type(text)"`
  Tags string `orm:"size(100);index;null"`
  Status int8 `orm:"index"`
  Cover *Image `orm:"rel(one);null"`
  Comments []*Comments `orm:"reverse(many);null"`
  BaseModel
}

func (post Post)Save(title, content, tag, path string, status int8, user *User, image *Image) (p *Post,e validation.Validation) {
   logs.Info(content)
   o := orm.NewOrm();
   post.Title = title
   post.Content = content
   post.Tags = tag
  //  post.Cover = path
   post.Status = status
   post.User = user
   logs.Info("正确的图片是", image)
   post.Cover = image
   logs.Info("model用户是", user)

   valid := validation.Validation{}
   valid.Required(post.Title, "title").Message("标题不存在")
   valid.Required(post.Content, "content").Message("正文不存在")
   logs.Info(post)
   if valid.HasErrors(){
    return &post, valid
   }

   if _,err := o.Insert(&post); err != nil {
    logs.Info(err)
    valid.SetError("post", "保存失败")
  }
  return &post, valid
}

func (post *Post)FindImage() *Image {
  o := orm.NewOrm()
  if post.Cover != nil {
    o.Read(post.Cover)
  }
  logs.Info("结果为", post.Cover)
  return post.Cover
}