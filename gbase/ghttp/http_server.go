package ghttp

import (
	"bytes"
	"github.com/globalsign/mgo/bson"
	"github.com/gosrv/gcluster/gbase/controller"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gutil"
	"github.com/gosrv/goioc"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"sync"
	"time"
)

type HttpServer struct {
	gioc.IBeanCondition
	gioc.IConfigBase
	log                  *logrus.Logger `log:"engine"`
	httpHost             string         `cfg.d:"http.host"`
	serverMux            *http.ServeMux
	controlPointGroupMgr controller.IControlPointGroupMgr `bean`
	ctxs                 *sync.Map
	filter               ControlPointFilter
}

type ControlPointFilter func(group, path string, point *controller.ControlPoint) *controller.ControlPoint

func AllControlPointPassFilter(group, path string, point *controller.ControlPoint) *controller.ControlPoint {
	return point
}

func NewHttpServer(cfgBase string, filter ControlPointFilter) *HttpServer {
	if filter == nil {
		filter = AllControlPointPassFilter
	}

	return &HttpServer{
		IBeanCondition: gioc.NewConditionOnValue(cfgBase, true),
		IConfigBase:    gioc.NewConfigBase(cfgBase),
		filter:         filter,
		ctxs:           new(sync.Map),
	}
}

type ExpireSession struct {
	ctx        gnet.ISessionCtx
	activeTime time.Time
}

func NewExpireSession(ctx gnet.ISessionCtx, activeTime time.Time) *ExpireSession {
	return &ExpireSession{ctx: ctx, activeTime: activeTime}
}

func (this *HttpServer) initRequestCtx(writer http.ResponseWriter, request *http.Request) gnet.ISessionCtx {
	cokieid, err := request.Cookie("sessionid")
	var ectx *ExpireSession
	if err == nil {
		ctxins, _ := this.ctxs.Load(cokieid.Value)
		if ctxins != nil {
			ectx = ctxins.(*ExpireSession)
		}
	}
	if ectx == nil {
		cokieid = &http.Cookie{Name: "sessionid", Value: bson.NewObjectId().Hex(), Expires: time.Now().Add(time.Hour)}
		http.SetCookie(writer, cokieid)
		ectx = NewExpireSession(gnet.NewSessionCtx(), time.Now())
		ectx.ctx.SetAttribute(gnet.ScopeSession, gnet.ISessionCtxType, ectx.ctx)
		this.ctxs.Store(cokieid.Value, ectx)
	}
	ectx.activeTime = time.Now()
	ctx := ectx.ctx

	ctx.Clear(gnet.ScopeRequest)
	params := NewHttpParam(request.URL.Query())
	ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(params), params)
	headers := NewHttpHeader(request.Header)
	ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(headers), headers)
	cookies := NewHttpCookie(request.Cookies())
	ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(cookies), cookies)
	form := NewHttpForm(request.Form)
	ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(form), form)
	bytes.NewBuffer(nil)

	ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(writer), writer)
	ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(request), request)
	return ctx
}

func (this *HttpServer) BeanStart() {
	this.serverMux = http.NewServeMux()
	allGroup := this.controlPointGroupMgr.GetAllControlGroup()
	for name, group := range allGroup {
		for route, point := range group.GetAllControlPoints() {
			if !reflect.TypeOf(point.Controller).AssignableTo(IHttpControllerType) {
				continue
			}
			cpoint := this.filter(name, route.(string), point)
			if cpoint == nil {
				continue
			}
			lastRoute := name + route.(string)
			this.log.WithField("route", lastRoute).Debug("connect http route")
			this.serverMux.HandleFunc(lastRoute, func(writer http.ResponseWriter, request *http.Request) {
				ctx := this.initRequestCtx(writer, request)
				rep := cpoint.Controller.Trigger().Trigger(cpoint, ctx)
				if rep != nil {
					data, err := cpoint.Controller.(IHttpController).ViewRender().RendView(rep)
					if err != nil {
						this.log.Warnf("http view rend error %v", err)
					} else {
						_, err = writer.Write(data)
						if err != nil {
							this.log.Warnf("http response write error %v", err)
						}
					}
				}
			})
		}
	}

	gutil.RecoverGo(func() {
		http.ListenAndServe(this.httpHost, this.serverMux)
	})
	this.log.WithField("host", this.httpHost).Info("http server start")

	gutil.RecoverGo(func() {
		for {
			now := time.Now()
			this.ctxs.Range(func(key, value interface{}) bool {
				ectx := value.(*ExpireSession)
				if now.Sub(ectx.activeTime) > time.Minute*10 {
					this.ctxs.Delete(key)
				}
				return true
			})
			time.Sleep(time.Minute)
		}
	})
}

func (this *HttpServer) BeanStop() {

}
