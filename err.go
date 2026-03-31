package jserr

import (
	"runtime"
	"syscall/js"

	"github.com/tinywasm/fmt"
)

type Error struct {
	js.Value
}

func New(val ...any) *Error {
	return &Error{Value: js.Global().Get("Error").New(val...)}
}

func Wrap(err error) *Error {
	return New(err.Error())
}

func (e *Error) Log() {
	defer Recover()
	Log(e.Value)
}

func (e *Error) Error() string {
	return e.Call("toString").String()
}

func Recover() {
	if r := recover(); r != nil {
		pc, f, l, _ := runtime.Caller(3)
		err := New(fmt.Sprintf(errMsg, f, l, runtime.FuncForPC(pc).Name(), r))
		err.Log()
	}
}

const errMsg = `
file: %v
  line: %d
	func: %s
	err: %w
`

func Log(vals ...any) {
	defer Recover()
	js.Global().Get("console").Call("log", vals...)
}
