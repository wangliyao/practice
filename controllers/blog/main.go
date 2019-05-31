package blog

import(
  "github.com/astaxie/beego/orm"
  "practice/models"
  "github.com/astaxie/beego/utils/pagination"
  "strconv"
  "github.com/astaxie/beego/logs"
)

type MainController struct {
  baseController
}

func (this *MainController ) Go404() {
}

func (this *MainController ) Index() {
  orm := orm.NewOrm()
  var post []*models.Post
  countPosts, _ := strconv.ParseInt("1", 0, 64)
  postsPerPage := 5
  paginator := pagination.NewPaginator(this.Ctx.Request, postsPerPage, countPosts)
  orm.QueryTable("post").Offset(paginator.Offset()).Limit(postsPerPage).OrderBy("-id").RelatedSel().All(&post)
  logs.Info("offset", paginator.Offset())
  this.Data["posts"] = post
  logs.Info(paginator)
  this.Data["paginator"] = paginator
  this.display("index")
}



