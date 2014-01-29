package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type IdT struct {
	m   Point
	Run AnyVal
}

func NewIdT(m Point) IdT {
	return IdT{
		m:   m,
		Run: Empty{},
	}
}

func (x IdT) Of(v AnyVal) Point {
	return IdT{
		m:   x.m,
		Run: x.m.Of(v),
	}
}

func (x IdT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f AnyVal) Monad {
		return v.(Functor).Map(func(x AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		}).(Monad)
	})
	return mon.(Applicative)
}

func (x IdT) Chain(f func(v AnyVal) Monad) Monad {
	mon := x.Run.(Monad)
	tra := IdT{
		m: x.m,
		Run: mon.Chain(func(y AnyVal) Monad {
			idt := f(y).(IdT)
			return idt.Run.(Monad)
		}),
	}
	return tra
}

func (x IdT) Map(f func(v AnyVal) AnyVal) Functor {
	mon := x.Chain(func(y AnyVal) Monad {
		app := NewIdT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}
