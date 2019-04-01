package gdb

type IDataLoader interface {
	Load(key string) (string, error)
}

type ITheDataLoader interface {
	Load() (string, error)
}

type FuncDataLoader func(key string) (string, error)

func (this FuncDataLoader) Load(key string) (string, error) {
	return this(key)
}

type FuncTheDataLoader func() (string, error)

func (this FuncTheDataLoader) Load() (string, error) {
	return this()
}

type TheDataLoaderChain struct {
	loaders []ITheDataLoader
}

func (this *TheDataLoaderChain) AddLoader(loader ITheDataLoader) {
	this.loaders = append(this.loaders, loader)
}

func (this *TheDataLoaderChain) Load() (string, error) {
	var firstError error
	for _, loader := range this.loaders {
		data, err := loader.Load()
		if err == nil {
			return data, nil
		} else if firstError == nil {
			firstError = err
		}
	}
	return "", firstError
}
