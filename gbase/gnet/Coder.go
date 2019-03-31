package gnet

import "github.com/gosrv/gcluster/gbase/gutil"

type NetMsgEncoder interface {
	Encode(interface{}) ([]byte, error)
}

type NetMsgDecoder interface {
	Decode(reader *gutil.Buffer) (interface{}, error)
}
