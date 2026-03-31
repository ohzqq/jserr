package jserr

import (
	"errors"
	"runtime"
	"syscall/js"

	"github.com/tinywasm/fmt"
)

var NotEnoughArgsErr = errors.New("not enough args")

type Error struct {
	js.Value
}

func New(val string) *Error {
	return &Error{Value: js.Global().Get("Error").New(val)}
}

func WrapError(err error) *Error {
	return New(err.Error())
}

func (e *Error) Log() {
	js.Global().Get("console").Call("log", e.Value)
}

func Log(vals ...js.Value) {
	js.Global().Get("console").Call("log", vals...)
}

func (e *Error) Error() string {
	return e.Call("toString").String()
}

func CheckArgs(args []js.Value) error {
	if len(args) < 1 {
		return NotEnoughArgsErr
	}
	return nil
}

func Recover() {
	if r := recover(); r != nil {
		pc, f, l, _ := runtime.Caller(3)
		fmt.Printf(errMsg, f, l, runtime.FuncForPC(pc).Name(), r)
	}
}

const errMsg = `
file: %v
  line: %d
	func: %s
	err: %w
`
