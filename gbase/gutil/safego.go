package gutil

func RecoverGo(gofunc func()) {
	go func() {
		defer PanicRecover()
		gofunc()
	}()
}
