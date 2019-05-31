package admin

type MainController struct {
	baseController
}


func (this *MainController ) Index() {
	this.Data["src"] = "/admin/content"
	this.display("index")
  }


func (this *MainController) Content() {
  this.TplName = "admin/content.html"
}