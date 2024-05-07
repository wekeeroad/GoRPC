package errcode

import "fmt"

type Error struct {
	code int
	msg  string
}

var _codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("err code %d is exist, please change that", code))
	}
	_codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (err *Error) Code() int {
	return err.code
}

func (err *Error) Msg() string {
	return err.msg
}
