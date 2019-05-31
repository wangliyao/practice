package models

import (
  "time"
)

type Comments struct {
  Post *Post `orm:"rel(fk);index"`
  User *User `orm:"rel(fk);index"`
  Comment string `orm:"type(text)"`
  Submittime time.Time `orn:"auto_now_add;type(datetime);index"`
  BaseModel
}