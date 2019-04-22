package gmxdriver

import (
	"fmt"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gmx"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"reflect"
	"strings"
)

const (
	GmxTag = "gmx"
	GmxOpt = "gmx.opt"
)

type GMXDriver struct {
	ghttp.IHttpController
	gmx.IMXManager
	mgr   *gmx.MXManager
	insId int
}

var _ gioc.ITagProcessor = (*GMXDriver)(nil)

func NewGMXDriver(path string) *GMXDriver {
	driver := &GMXDriver{
		IHttpController: ghttp.NewHttpRestController(path),
		mgr:             gmx.NewMXManager(),
	}
	driver.IMXManager = driver.mgr
	return driver
}

func (this *GMXDriver) TagProcessorName() string {
	return "gmx"
}

func (this *GMXDriver) PrepareProcess() {

}

func (this *GMXDriver) TagProcess(bean interface{}, fType reflect.StructField, fValue reflect.Value, tags map[string]string) {
	gmxTag, gmxTagOk := tags[GmxTag]
	if !gmxTagOk {
		return
	}
	gmxOpt := tags[GmxOpt]
	canRead := true
	canWrite := true
	if len(gmxOpt) > 0 {
		canRead = strings.Contains(gmxOpt, "r")
		canWrite = strings.Contains(gmxOpt, "w")
	}

	for fValue.Kind() == reflect.Ptr {
		if canWrite && fValue.IsNil() {
			fValue.Set(reflect.New(fValue.Type().Elem()))
		}
		fValue = fValue.Elem()
	}
	fValue = util.Hack.ValuePatchWrite(fValue)
	// 如果没有指定名字，则创建一个
	gmxName := gmxTag
	if len(gmxName) == 0 {
		this.insId++
		beanType := reflect.TypeOf(bean)
		for beanType.Kind() == reflect.Ptr {
			beanType = beanType.Elem()
		}
		gmxName = fmt.Sprintf("%v.%v.%v", beanType.Name(), fType.Name, this.insId)
	}
	if fValue.CanAddr() && canWrite {
		err := this.mgr.AddItemInsRW(gmxName, fValue.Addr().Interface(), canRead, canWrite)
		util.VerifyNoError(err)
	} else {
		err := this.mgr.AddItemInsRW(gmxName, fValue.Interface(), canRead, canWrite)
		util.VerifyNoError(err)
	}
}

///keys?key=...
func (this *GMXDriver) Keys(ctx gnet.ISessionCtx, param *ghttp.HttpParam) []byte {
	rep, _ := this.mgr.HandleKeys()
	return rep
}

///get?key=...,...,...
func (this *GMXDriver) Get(ctx gnet.ISessionCtx, param *ghttp.HttpParam) []byte {
	keys := strings.Split(param.ParamSingle["key"], ",")
	rep, _ := this.mgr.HandleGet(keys)
	return rep
}

///set?key=...,...,...value=...,...,...
func (this *GMXDriver) Set(ctx gnet.ISessionCtx, param *ghttp.HttpParam) []byte {
	keys := strings.Split(param.ParamSingle["key"], ",")
	vals := strings.Split(param.ParamSingle["value"], ",")
	rep, _ := this.mgr.HandleSet(keys, vals)
	return rep
}

///call?key=...&params=...,...,...
func (this *GMXDriver) Call(ctx gnet.ISessionCtx, param *ghttp.HttpParam) []byte {
	key := param.ParamSingle["key"]
	params := strings.Split(param.ParamSingle["params"], ",")
	rep, _ := this.mgr.HandleCall(key, params)
	return rep
}
