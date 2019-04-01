package gdb

type IDataSaver interface {
	Save(key string, value string) error
}

type ITheDataSaver interface {
	Save(value string) error
}

type FuncTheDataSaver = func(value string) error
type FuncTheDataSaverWraper struct {
	funcSaver FuncTheDataSaver
}

func NewFuncTheDataSaverWraper(funcSaver FuncTheDataSaver) *FuncTheDataSaverWraper {
	return &FuncTheDataSaverWraper{funcSaver: funcSaver}
}

func (this *FuncTheDataSaverWraper) Save(value string) error {
	return this.funcSaver(value)
}

type TheDataSaverChain struct {
	savers []ITheDataSaver
}

func (this *TheDataSaverChain) Save(value string) error {
	var firstError error = nil
	for _, saver := range this.savers {
		err := saver.Save(value)
		if err != nil && firstError == nil {
			firstError = err
		}
	}
	return firstError
}

func (this *TheDataSaverChain) AddSaver(saver ITheDataSaver) {
	this.savers = append(this.savers, saver)
}

func (this *TheDataSaverChain) SaveDepth(value string, depth int) {
	for i := 0; (depth <= 0 || i < depth) && i < len(this.savers); i++ {
		this.savers[i].Save(value)
	}
}
