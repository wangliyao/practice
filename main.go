package main

import (
	_ "practice/routers"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/logs"
)


func initLogs() {
  var fileconfig = `{"filename":"development.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`
  var consoleconfig = `{"level":7}`
  logs.SetLogger(logs.AdapterConsole,consoleconfig)
  logs.SetLogger(logs.AdapterFile, fileconfig)
  logs.Async(1e3)
}

func main() {
  initLogs()
	beego.Run()
}

