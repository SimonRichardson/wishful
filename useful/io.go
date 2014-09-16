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

func (x IO) Chain(f func(x Any) Monad) Monad {
	return NewIO(func() Any {
		io := f(x.UnsafePerform()).(IO)
		return io.UnsafePerform()
	})
}

func (x IO) Map(f func(x Any) Any) Functor {
	res := x.Chain(func(x Any) Monad {
		return IO{func() Any {
			return f(x)
		}}
	})
	return res.(Functor)
}
