package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type io struct {
	UnsafePerform func() Any
}

func IO(unsafe func() Any) io {
	return io{
		UnsafePerform: unsafe,
	}
}

func (x io) Of(v Any) Point {
	return IO(func() Any {
		return v
	})
}

func (x io) Ap(v Applicative) Applicative {
	res := x.Chain(func(f Any) Monad {
		fun := v.(Functor)
		res := fun.Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		})
		return res.(Monad)
	})
	return res.(Applicative)
}

func (x io) Chain(f func(x Any) Monad) Monad {
	return IO(func() Any {
		io := f(x.UnsafePerform()).(io)
		return io.UnsafePerform()
	})
}

func (x io) Map(f func(x Any) Any) Functor {
	res := x.Chain(func(x Any) Monad {
		return io{func() Any {
			return f(x)
		}}
	})
	return res.(Functor)
}

var (
	IO_ = io_{}
)

type io_ struct{}

func (f io_) As(x Any) io {
	return x.(io)
}

func (f io_) Ref() io {
	return io{}
}

func (f io_) Of(x Any) Point {
	return io{}.Of(x)
}
