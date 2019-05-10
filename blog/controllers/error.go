package controllers

import (
	"blog/syserror"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	if c.IsAjax(){
		c.Ctx.Output.Status = 200
		c.Data["json"] = map[string]interface{}{
			"code": 1002,
			"msg":"illegal request!",
		}
		c.ServeJSON()
	}
	c.Data["content"] = "page not found"
	c.TplName = "404.html"
}

func (c *ErrorController) Error500() {
	err, ok:=c.Data["error"].(error)
	if !ok{
		err = syserror.New("unkown error", nil)
	}
	serr, ok := err.(syserror.Error)
	if !ok{
		serr = syserror.New(err.Error(),nil)
	}
	if serr.ErrorReson() != nil{
		logs.Info(serr.Error(),serr.ErrorReson())
	}
	if c.IsAjax(){
		c.jsonerr(serr)
	}else {
		c.Data["content"] = serr.Error()
		c.TplName = "500.html"
	}
}


func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.html"
}

func (c *ErrorController) jsonerr( serr syserror.Error)  {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code": serr.Code(),
		"msg":serr.Code(),
	}
	c.ServeJSON()
}