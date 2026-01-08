package main

import (
	"fmt"

	"github.com/tobiashort/utils-go/errors"
)

func foo() error {
	return errors.WithCtx("error occurred")
}

func bar() error {
	return foo()
}

func main() {
	err := bar()
	if err != nil {
		fmt.Println(err)
	}
}
