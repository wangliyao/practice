package admin

import (
  "github.com/astaxie/beego/orm"
  "practice/models"
  "github.com/astaxie/beego/logs"
)

type PostController struct {
	baseController
}

const (
  publish = iota
  unpublish
)

func (this *PostController) Index(){
  var post []*models.Post
  orm := orm.NewOrm()
  qs := orm.QueryTable("post")
  num, err :=  qs.All(&post)
  this.Data["posts"] = post
  this.Data["num"] = num
  this.Data["error"] = err
  logs.Info("结构是", post)
  this.TplName = "admin/post/index.html"
}

func (this *PostController) New(){
  this.TplName = "admin/post/new.html"
}

func (this *PostController) Create(){
  var post models.Post
  var image models.Image
  title := this.GetString("title")
  content := this.GetString("content")
  tag := this.GetString("tag")
  status ,_ := this.GetInt8("status", 0)
  logs.Info("状态是", this.GetString("status"))
  user := this.FindUser()
  path := ""
  logs.Info("正文是", content)
  f, h, errs := this.GetFile("uploadfile")
  if errs == nil{
    path = "./upload/" + h.Filename
    f.Close()
  }
  this.SaveToFile("uploadfile", path)
  images := image.GetImage(this.GetString("CoverId"))

  _, err := post.Save(title, content, tag, path, status, user, images)
  if len(err.Errors) != 0 {
    for _, errs := range err.Errors {
      this.Data["errMsg"] = errs.Message
     }
    this.TplName = "admin/post/new.html"
    return
  }
  this.Redirect("/admin/posts", 302)
}

func (this *PostController) Delete(){
  var post models.Post
  id := this.Ctx.Input.Param(":id")
   o := orm.NewOrm()
   qs := o.QueryTable("post")
   qs.Filter("id", id).One(&post)
  if _, err := o.Delete(&post); err == nil {
    this.Data["errMsg"] = "删除成功"
  }else{
    this.Data["errMsg"] = "删除失败"
  }
  this.Redirect("/admin/posts", 302)
}

func (this *PostController) Update(){
  var image models.Image
  var post models.Post
  id := this.Ctx.Input.Param(":id")
  o := orm.NewOrm()
  qs := o.QueryTable("post")
  qs.Filter("id", id).One(&post)
  title := this.GetString("title")
  content := this.GetString("content")
  tag := this.GetString("tag")
  status ,_ := this.GetInt8("status", 0)
  images := image.GetImage(this.GetString("CoverId"))

  if o.Read(&post) == nil {
    post.Title = title
    post.Content = content
    post.Tags = tag
    post.Status = status
    post.Cover = images
    if _, err := o.Update(&post); err == nil {
      this.Redirect("/admin/posts", 302)
    }else {
      this.Data["error"] = err
      this.Data["post"] = &post
      this.TplName = "admin/post/edit.html"
      return
    }
  }
}

func (this *PostController) Edit(){
   id := this.Ctx.Input.Param(":id")
   var post models.Post
   orm := orm.NewOrm()
   qs := orm.QueryTable("post")
   err :=  qs.Filter("id", id).One(&post)
   this.Data["post"] = post
   logs.Info(id)
   this.Data["error"] = err

   result := post.FindImage()
   logs.Info("image结果", result)
   this.Data["CoverUrl"] = "/static/upload/" + result.FileName
   
   this.TplName = "admin/post/edit.html"
}