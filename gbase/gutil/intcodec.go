package gutil

import "errors"

type intCodec struct{}

var IntCodec = &intCodec{}

func (this *intCodec) PackFixInt32BE(val int32, buf []byte, offset int) (int, error) {
	if offset+4 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	buf[offset+3] = (byte)(val & 0xFF)
	buf[offset+2] = (byte)((val >> 8) & 0xFF)
	buf[offset+1] = (byte)((val >> 16) & 0xFF)
	buf[offset+0] = (byte)((val >> 24) & 0xFF)
	return offset + 4, nil
}

func (this *intCodec) PackFixInt16BE(val int16, buf []byte, offset int) (int, error) {
	if offset+2 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	buf[offset+1] = (byte)(val & 0xFF)
	buf[offset+0] = (byte)((val >> 8) & 0xFF)
	return offset + 2, nil
}

func (this *intCodec) PackFixInt32LE(val int32, buf []byte, offset int) (int, error) {
	if offset+4 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	buf[offset] = (byte)(val & 0xFF)
	buf[offset+1] = (byte)((val >> 8) & 0xFF)
	buf[offset+2] = (byte)((val >> 16) & 0xFF)
	buf[offset+3] = (byte)((val >> 24) & 0xFF)
	return offset + 4, nil
}

func (this *intCodec) PackFixInt16LE(val int16, buf []byte, offset int) (int, error) {
	if offset+2 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	buf[offset] = (byte)(val & 0xFF)
	buf[offset+1] = (byte)((val >> 8) & 0xFF)
	return offset + 2, nil
}

func (this *intCodec) UnpackFixInt32BE(buf []byte, offset int) (int32, error) {
	if offset+4 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	return int32(buf[offset+3]&0xFF) |
		(int32(buf[offset+2]&0xFF) << 8) |
		(int32(buf[offset+1]&0xFF) << 16) |
		(int32(buf[offset+0]&0xFF) << 24), nil
}

func (this *intCodec) UnpackFixInt16BE(buf []byte, offset int) (int16, error) {
	if offset+2 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	return int16(buf[offset+1]&0xFF) |
		(int16(buf[offset+0]&0xFF) << 8), nil
}

func (this *intCodec) UnpackFixInt32LE(buf []byte, offset int) (int32, error) {
	if offset+4 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	return int32(buf[offset]&0xFF) |
		(int32(buf[offset+1]&0xFF) << 8) |
		(int32(buf[offset+2]&0xFF) << 16) |
		(int32(buf[offset+3]&0xFF) << 24), nil
}

func (this *intCodec) UnpackFixInt16LE(buf []byte, offset int) (int16, error) {
	if offset+2 > len(buf) {
		return -1, errors.New("index out of bound")
	}
	return int16(buf[offset]&0xFF) |
		(int16(buf[offset+1]&0xFF) << 8), nil
}
