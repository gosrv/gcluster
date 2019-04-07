package gutil

import (
	"errors"
	"io"
)

type Buffer struct {
	buf        []byte
	readPos    int
	writePos   int
	totalRead  int
	totalWrite int
}

func (this *Buffer) Fill(reader io.Reader) (int, error) {
	this.Pack()
	if this.writePos >= len(this.buf) {
		return 0, errors.New("buffer error")
	}

	n, err := reader.Read(this.buf[this.writePos:])
	if n > 0 {
		this.writePos += n
		this.totalWrite += n
	}
	return n, err
}

func (this *Buffer) Len() int {
	return this.writePos - this.readPos
}

func (this *Buffer) Pack() {
	if this.readPos == 0 {
		return
	}
	if this.readPos != this.writePos {
		copy(this.buf, this.buf[this.readPos:this.writePos])
	}
	this.writePos -= this.readPos
	this.readPos = 0
}

func (this *Buffer) Peek(n int) []byte {
	if this.Len() < n {
		return nil
	}

	return this.buf[this.readPos : this.readPos+n]
}

func (this *Buffer) Read(n int) []byte {
	if this.Len() < n {
		return nil
	}
	rdata := this.buf[this.readPos : this.readPos+n]
	this.readPos += n
	this.totalRead += n
	return rdata
}

func (this *Buffer) IsFull() bool {
	return this.Len() == len(this.buf)
}

func NewBuffer(n int) *Buffer {
	return &Buffer{
		buf:      make([]byte, n, n),
		readPos:  0,
		writePos: 0,
	}
}
