package controller

const (
	TypeController = "TypeController"
)

/**
控制器,网络消息处理器
*/
type IController interface {
	Group() string
	Type() string
	Finder() IControlPointFinder
	Trigger() IControlPointTrigger
}

type controller struct {
	group   string
	ctype   string
	finder  IControlPointFinder
	trigger IControlPointTrigger
}

func (this *controller) Type() string {
	return this.ctype
}

func (this *controller) Group() string {
	return this.group
}

func (this *controller) Finder() IControlPointFinder {
	return this.finder
}

func (this *controller) Trigger() IControlPointTrigger {
	return this.trigger
}

func NewController(group string, finder IControlPointFinder, trigger IControlPointTrigger) *controller {
	return &controller{group: group, finder: finder, trigger: trigger}
}

// 类型路由器
func NewTypeController(group string) *controller {
	return &controller{
		group:   group,
		ctype:   TypeController,
		finder:  FuncTypeControlPointFinder(typeControlPointFinder),
		trigger: FuncControlPointTrigger(typeControlPointTrigger),
	}
}
