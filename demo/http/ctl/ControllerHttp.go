package ctl

import (
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gbase/gnet"
	"net/http"
)

/**
逻辑消息控制器
*/
type controllerHttp struct {
	// 控制器标记
	ghttp.IHttpController
}

func NewControllerHttp() *controllerHttp {
	return &controllerHttp{
		IHttpController: ghttp.NewHttpRestController("/demo"),
	}
}

//	/demo/hello/raw
func (this *controllerHttp) HelloRaw(ctx gnet.ISessionCtx, writer http.ResponseWriter, request *http.Request) interface{} {
	return "/demo/hello/raw"
}

//  /demo/hello?Account=abc&Password=123
func (this *controllerHttp) Hello(ctx gnet.ISessionCtx, params *struct{ Account, Password string }) interface{} {
	return params
}

//	/demo/hello/param?p1=1&p2=2
func (this *controllerHttp) HelloParam(ctx gnet.ISessionCtx, params *ghttp.HttpParam) interface{} {
	return params
}

//	/demo/hello/form
func (this *controllerHttp) HelloForm(ctx gnet.ISessionCtx, form *ghttp.HttpForm) interface{} {
	return form
}

//	/demo/hello/header
func (this *controllerHttp) HelloHeader(ctx gnet.ISessionCtx, header *ghttp.HttpHeader) interface{} {
	return header
}

//	/demo/hello/cookie
func (this *controllerHttp) HelloCookie(ctx gnet.ISessionCtx, cookie *ghttp.HttpCookie) interface{} {
	return cookie
}
