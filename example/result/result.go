package main

import "fmt"

type Result[T any] struct {
	Val T
	Err error
}

func Ok[T any](v T) Result[T] {
	return Result[T]{
		Val: v,
		Err: nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

func foo() Result[string] {
	// return Ok("baz")
	// return Ok("bar")
	return Err[string](fmt.Errorf("error"))
}

func main() {
	r := foo()
	switch {
	case r.Val == "baz":
		fmt.Println("baz")
	case r.Val == "bar":
		fmt.Println("bar")
	case r.Err != nil:
		fmt.Println("error")
	default:
		panic("unreachable")
	}
}
