package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	mlog, _ := New("debug", "./ylog", 0)
	Export(mlog)
	defer mlog.Close()

	Debug("Hello, World")
	Error("This is an error msg")
	Release("This is an release msg")
	Fatal("This is an fatal msg")
}
