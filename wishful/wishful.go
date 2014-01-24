package wishful

// Identity

type Id struct {
	x AnyVal
}

// Option

type Option interface {
}
type Some struct {
	x AnyVal
}
type None struct {
}

// Promise

type Promise struct {
	fork func(resolve func(x AnyVal) AnyVal) AnyVal
}

func (x Promise) Fork(f func(x AnyVal) AnyVal) AnyVal {
	return x.fork(f)
}
