package useful

import (
	. "github.com/SimonRichardson/wishful"
)

type IdT struct {
	m   Applicative
	Run AnyVal
}

func NewIdT(m Applicative) IdT {
	return IdT{
		m:   m,
		Run: Empty{},
	}
}

func (x IdT) Of(v AnyVal) Applicative {
	return IdT{
		m:   x.m,
		Run: x.m.Of(v),
	}
}

func (x IdT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f AnyVal) Monad {
		fun := f.(func(f AnyVal) AnyVal)
		return v.(Functor).Map(fun).(Monad)
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
