package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type IO struct {
	UnsafePerform func() AnyVal
}

func NewIO(unsafe func() AnyVal) IO {
	return IO{
		UnsafePerform: unsafe,
	}
}

func (x IO) Of(v AnyVal) Point {
	return NewIO(func() AnyVal {
		return v
	})
}

func (x IO) Ap(v Applicative) Applicative {
	res := x.Chain(func(f AnyVal) Monad {
		fun := v.(Functor)
		res := fun.Map(func(x AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		})
		return res.(Monad)
	})
	return res.(Applicative)
}

func (x IO) Chain(f func(x AnyVal) Monad) Monad {
	return NewIO(func() AnyVal {
		io := f(x.UnsafePerform()).(IO)
		return io.UnsafePerform()
	})
}

func (x IO) Map(f func(x AnyVal) AnyVal) Functor {
	res := x.Chain(func(x AnyVal) Monad {
		return IO{func() AnyVal {
			return f(x)
		}}
	})
	return res.(Functor)
}
