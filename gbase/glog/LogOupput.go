package glog

import (
	"io"
	"os"
)

type ILogOutput interface {
	LogOutputName() string
	ConfigLogOutput(writer *LogOutputWriter, cfg map[string]string)
}

type LogOutputWriter struct {
	writers []io.Writer
}

func NewLogOutputWriter() *LogOutputWriter {
	return &LogOutputWriter{}
}

func (this *LogOutputWriter) AddWriter(writer io.Writer) {
	this.writers = append(this.writers, writer)
}

func (this *LogOutputWriter) Write(p []byte) (n int, err error) {
	for _, w := range this.writers {
		wn, werr := w.Write(p)
		if n == 0 {
			n = wn
		}
		if werr != nil && err == nil {
			err = werr
		}
	}
	return n, err
}

type LogOutputConsole struct {
}

func NewLogOutputConsole() *LogOutputConsole {
	return &LogOutputConsole{}
}

func (this *LogOutputConsole) LogOutputName() string {
	return "console"
}

func (this *LogOutputConsole) ConfigLogOutput(writer *LogOutputWriter, cfg map[string]string) {
	writer.AddWriter(os.Stdout)
}

type LogOutputFile struct {
}

func NewLogOutputFile() *LogOutputFile {
	return &LogOutputFile{}
}

func (this *LogOutputFile) LogOutputName() string {
	return "file"
}

func (this *LogOutputFile) ConfigLogOutput(writer *LogOutputWriter, cfg map[string]string) {
	file, err := os.OpenFile(cfg["name"], os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		Panic("config file log output failed %v %v", cfg, err.Error())
	}
	writer.AddWriter(file)
}

var LogOutputConfigs = make(map[string]ILogOutput)

func init() {
	consolecfg := NewLogOutputConsole()
	LogOutputConfigs[consolecfg.LogOutputName()] = consolecfg
	filecfg := NewLogOutputFile()
	LogOutputConfigs[filecfg.LogOutputName()] = filecfg
}
