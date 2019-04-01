package ghttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"reflect"
	"sync"
)

type IViewRender interface {
	RendView(obj interface{}) ([]byte, error)
}

type restViewRender struct {
}

func NewRestViewRender() *restViewRender {
	return &restViewRender{}
}

func (this *restViewRender) RendView(obj interface{}) ([]byte, error) {
	switch obj.(type) {
	case string:
		return []byte(obj.(string)), nil
	case []byte:
		return obj.([]byte), nil
	default:
		repType := reflect.TypeOf(obj)
		repValue := reflect.ValueOf(obj)
		for repType.Kind() == reflect.Ptr {
			repType = repType.Elem()
			repValue = repValue.Elem()
		}
		if repType.Kind() == reflect.Struct {
			return json.Marshal(repValue.Interface())
		} else {
			return nil, errors.New(fmt.Sprintf("unsupport response type %v", reflect.TypeOf(repType)))
		}
	}
}

type ModAndView struct {
	view string
	mod  interface{} //struct, map
}

func NewModAndView(view string, mod interface{}) *ModAndView {
	return &ModAndView{view: view, mod: mod}
}

type textTemplateViewRender struct {
	name          string
	prefix        string
	suffix        string
	cacheTemplate sync.Map //string,*template.Template
}

func NewTextTemplateViewRender(name string, prefix string, suffix string) *textTemplateViewRender {
	return &textTemplateViewRender{name: name, prefix: prefix, suffix: suffix}
}

func (this *textTemplateViewRender) obtainTemplate(view string) (*template.Template, error) {
	tmpl, ok := this.cacheTemplate.Load(view)
	if ok {
		return tmpl.(*template.Template), nil
	}
	tmpFileName := this.prefix + view + this.suffix
	fdata, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		return nil, err
	}
	tmpl, err = template.New(view).Parse(string(fdata))
	if err != nil {
		return nil, err
	}
	this.cacheTemplate.Store(view, tmpl)
	return tmpl.(*template.Template), nil
}

func (this *textTemplateViewRender) RendView(obj interface{}) ([]byte, error) {
	rval := reflect.ValueOf(obj)
	for rval.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}
	obj = rval.Interface()

	switch obj.(type) {
	case string:
		tmpl, err := this.obtainTemplate(obj.(string))
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		err = tmpl.Execute(writer, nil)
		return writer.Bytes(), err
	case ModAndView:
		mav := obj.(ModAndView)
		tmpl, err := this.obtainTemplate(mav.view)
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		err = tmpl.Execute(writer, mav.mod)
		return writer.Bytes(), err
	default:
		return nil, errors.New(fmt.Sprintf("not support response type %v", reflect.TypeOf(obj)))
	}
}
