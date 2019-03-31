package gproto

type IEntity interface {
	Push(data interface{}) error
	Request(data interface{}) (interface{}, error)
}
