package errors

import (
	errors2 "errors"
	"fmt"
	"runtime"
)

func WithCtx(text string) error {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		text = fmt.Sprintf("%s\n  at %s:%d", text, file, line)
	}
	return errors2.New(text)
}

func WithCtxf(format string, a ...any) error {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		format = fmt.Sprintf("%s\n  at %s:%d", format, file, line)
	}
	return fmt.Errorf(format, a...)
}
