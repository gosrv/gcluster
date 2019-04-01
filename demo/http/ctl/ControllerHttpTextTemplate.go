package ctl

import (
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gbase/gnet"
	"net/http"
)

/**
逻辑消息控制器
*/
type controllerHttpTemplate struct {
	// 控制器标记，必须集成这个接口才会被http模块判定为控制器
	ghttp.IHttpController
}

func NewControllerHttpTemplate() *controllerHttpTemplate {
	viewRender := ghttp.NewTextTemplateViewRender("tmol", "demo/conf/tmpl/", ".html")
	return &controllerHttpTemplate{
		// 起始路径，所有本控制器的路径都会在以起始路径开始
		IHttpController: ghttp.NewHttpController("/tmpl", viewRender),
	}
}

//	/demo/hello/raw
func (this *controllerHttpTemplate) HelloRaw(ctx gnet.ISessionCtx, writer http.ResponseWriter, request *http.Request) ghttp.ModAndView {
	return *ghttp.NewModAndView("hello", "/demo/hello/raw")
}

//  /demo/hello?Account=abc&Password=123
func (this *controllerHttpTemplate) Hello(ctx gnet.ISessionCtx, params *struct{ Account, Password string }) *ghttp.ModAndView {
	return ghttp.NewModAndView("hello", params)
}

//	/demo/hello/param?p1=1&p2=2
func (this *controllerHttpTemplate) HelloParam(ctx gnet.ISessionCtx, params *ghttp.HttpParam) *ghttp.ModAndView {
	return ghttp.NewModAndView("hello", params)
}

//	/demo/hello/form
func (this *controllerHttpTemplate) HelloForm(ctx gnet.ISessionCtx, form *ghttp.HttpForm) *ghttp.ModAndView {
	return ghttp.NewModAndView("hello", form.ParamSingle)
}

//	/demo/hello/header
func (this *controllerHttpTemplate) HelloHeader(ctx gnet.ISessionCtx, header *ghttp.HttpHeader) *ghttp.ModAndView {
	return ghttp.NewModAndView("hello", header)
}

//	/demo/hello/cookie
func (this *controllerHttpTemplate) HelloCookie(ctx gnet.ISessionCtx, cookie *ghttp.HttpCookie) *ghttp.ModAndView {
	return ghttp.NewModAndView("hello", cookie)
}
