package ghttp

import "net/http"

type HttpHeader struct {
	Headers map[string][]string
}

func NewHttpHeader(headers map[string][]string) *HttpHeader {
	return &HttpHeader{Headers: headers}
}

type HttpParam struct {
	Params      map[string][]string
	ParamSingle map[string]string
}

func (this *HttpParam) buildSingle() {
	this.ParamSingle = make(map[string]string)
	for k, v := range this.Params {
		if len(v) >= 1 {
			this.ParamSingle[k] = v[0]
		}
	}
}

func NewHttpParam(params map[string][]string) *HttpParam {
	param := &HttpParam{Params: params}
	param.buildSingle()
	return param
}

type HttpForm struct {
	Params      map[string][]string
	ParamSingle map[string]string
}

func (this *HttpForm) buildSingle() {
	this.ParamSingle = make(map[string]string)
	for k, v := range this.Params {
		if len(v) >= 1 {
			this.ParamSingle[k] = v[0]
		}
	}
}

func NewHttpForm(params map[string][]string) *HttpForm {
	form := &HttpForm{Params: params}
	form.buildSingle()
	return form
}

type HttpBody struct {
	data []byte
}

func NewHttpBody(data []byte) *HttpBody {
	return &HttpBody{data: data}
}

type HttpCookie struct {
	Cookies     []*http.Cookie
	ParamSingle map[string]string
}

func (this *HttpCookie) buildSingle() {
	this.ParamSingle = make(map[string]string)
	for _, v := range this.Cookies {
		this.ParamSingle[v.Name] = v.Value
	}
}

func NewHttpCookie(cookies []*http.Cookie) *HttpCookie {
	cookie := &HttpCookie{Cookies: cookies}
	cookie.buildSingle()
	return cookie
}
