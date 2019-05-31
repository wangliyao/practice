package models

import(
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"fmt"

)

type Image struct {
	BaseModel
	File string `orm:"size(70)"`
	FileName string `orm:"size(70)"`
	User     *User   `orm:"reverse(one)"`
	Post     *Post  `orm:"reverse(one)"`
}



func (image Image)Save(file string, filename string) (string, error){
	o := orm.NewOrm();
	image.File = file
	image.FileName = filename
	imageId ,err := o.Insert(&image)
	logs.Info(err)
	return fmt.Sprint(imageId), err
}

func (image Image)GetImage(id string) *Image {
	o := orm.NewOrm();
	logs.Info("图片id为", id)
	err := o.QueryTable("image").Filter("id", id).One(&image)
	logs.Info("错误为", err)
	return &image
}