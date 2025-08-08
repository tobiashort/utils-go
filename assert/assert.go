package assert

import (
	"fmt"
)

func Nil(v any, format string, args ...any) {
	if v != nil {
		panic(fmt.Sprintf(format, args...))
	}
}

func NotNil(v any, format string, args ...any) {
	if v == nil {
		panic(fmt.Sprintf(format, args...))
	}
}

func True(cond bool, format string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(format, args...))
	}
}

func False(cond bool, format string, args ...any) {
	if cond {
		panic(fmt.Sprintf(format, args...))
	}
}
