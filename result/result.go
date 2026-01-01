package result

type Result[T any] struct {
	Val T
	Err error
}

func Ok[T any](val T) Result[T] {
	return Result[T]{
		Val: val,
		Err: nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}
