package wishful

// Identity

type Id struct {
	x AnyVal
}

func NewId(x AnyVal) Id {
	return Id{
		x: x,
	}
}

// IO

type IO struct {
	UnsafePerform func() AnyVal
}

func NewIO(unsafe func() AnyVal) IO {
	return IO{
		UnsafePerform: unsafe,
	}
}

// Option

type Option interface {
}
type Some struct {
	x AnyVal
}
type None struct {
}

func NewSome(x AnyVal) Some {
	return Some{
		x: x,
	}
}
func NewNone() None {
	return None{}
}

// Promise

type Promise struct {
	fork func(resolve func(x AnyVal) AnyVal) AnyVal
}

func (x Promise) Fork(f func(x AnyVal) AnyVal) AnyVal {
	return x.fork(f)
}
