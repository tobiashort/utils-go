package assert

import (
	"fmt"
)

var PanicFunc = func(v any) {
	panic(v)
}

func Nil(v any, msg string) {
	if v != nil {
		PanicFunc(fmt.Errorf("%s", msg))
	}
}

func Nilf(v any, format string, args ...any) {
	if v != nil {
		PanicFunc(fmt.Errorf(format, args...))
	}
}

func NotNil(v any, msg string) {
	if v == nil {
		PanicFunc(fmt.Errorf("%s", msg))
	}
}

func NotNilf(v any, format string, args ...any) {
	if v == nil {
		PanicFunc(fmt.Errorf(format, args...))
	}
}

func True(cond bool, msg string) {
	if !cond {
		PanicFunc(fmt.Errorf("%s", msg))
	}
}

func Truef(cond bool, format string, args ...any) {
	if !cond {
		PanicFunc(fmt.Errorf(format, args...))
	}
}

func False(cond bool, msg string) {
	if cond {
		PanicFunc(fmt.Errorf("%s", msg))
	}
}

func Falsef(cond bool, format string, args ...any) {
	if cond {
		PanicFunc(fmt.Errorf(format, args...))
	}
}
