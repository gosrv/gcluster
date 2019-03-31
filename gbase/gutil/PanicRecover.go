package gutil

import (
	"log"
	"runtime/debug"
)

func PanicRecover() {
	err := recover()
	if err != nil {
		log.Printf("panic with error:", err, string(debug.Stack()))
	}
}
