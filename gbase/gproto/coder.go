package gproto

type ICoder interface {
	Code(data interface{}) interface{}
}

type IEncoder interface {
	Encode(data interface{}) interface{}
}

type IDecoder interface {
	Decode(data interface{}) interface{}
}
