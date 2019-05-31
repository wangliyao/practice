package admin

import(
  "github.com/astaxie/beego/logs"
  "practice/models"
)

type ImageController struct {
	baseController
}

func (this *ImageController ) Create()  {
  var image models.Image
  f, h, _ := this.GetFile("files")
  logs.Info(f)
  fileName := h.Filename
  path := "./static/upload/" + fileName
  logs.Info("文件名",fileName)
  
  err := this.SaveToFile("files", path)
  logs.Info("错误是", err)
  id := ""
  if err == nil {
	id, err = image.Save(path, fileName)
  }
  this.Data["json"] = map[string]interface{}{"fileName": fileName, "id": id,"err": err}

  f.Close()
  this.ServeJSON()
  return
}