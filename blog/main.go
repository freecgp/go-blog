package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/beego/bee/logger"
	_"blog/models"
	"strings"
)

func initTempalte(){
	err :=beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x,"/")
		y1 := strings.Trim(y,"/")
		return strings.Compare(x1,y1) == 0
	})
	if err != nil{
		beeLogger.Log.Error("Init tempaltes failed!")
	}else {
		beeLogger.Log.Info("Init templates successfully!")
	}
}

func main() {
	initTempalte()
	beego.Run()
}

