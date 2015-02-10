package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type IO struct {
	UnsafePerform func() Any
}

func NewIO(unsafe func() Any) IO {
	return IO{
		UnsafePerform: unsafe,
	}
}

func (x IO) Of(v Any) Point {
	return NewIO(func() Any {
		return v
	})
}

func (x IO) Ap(v Applicative) Applicative {
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

func (x IO) Chain(f func(Any) Monad) Monad {
	return NewIO(func() Any {
		IO := f(x.UnsafePerform()).(IO)
		return IO.UnsafePerform()
	})
}

func (x IO) Map(f Morphism) Functor {
	res := x.Chain(func(x Any) Monad {
		return IO{func() Any {
			return f(x)
		}}
	})
	return res.(Functor)
}

var (
	IO_ = io_{}
)

type io_ struct{}

func (f io_) As(x Any) IO {
	return x.(IO)
}

func (f io_) Ref() IO {
	return IO{}
}

func (f io_) Of(x Any) Point {
	return IO{}.Of(x)
}
