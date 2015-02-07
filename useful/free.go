package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Free interface {
	Of(Any) Point
	Chain(f func(Any) Monad) Monad
	Map(f func(Any) Any) Functor

	Run() Any
}

type Return struct {
	val Any
}

func NewReturn(x Any) Return {
	return Return{
		val: x,
	}
}

func (r Return) Of(x Any) Point {
	return NewReturn(x)
}

func (r Return) Chain(f func(Any) Monad) Monad {
	return f(r.val)
}

func (r Return) Map(f func(Any) Any) Functor {
	return r.Chain(func(x Any) Monad {
		return Free_.Of(f(x)).(Monad)
	}).(Functor)
}

func (r Return) Run() Any {
	return r.val
}

type Suspend struct {
	functor Functor
}

func NewSuspend(f Functor) Suspend {
	return Suspend{
		functor: f,
	}
}

func (s Suspend) Of(x Any) Point {
	return NewReturn(x)
}

func (s Suspend) Chain(f func(Any) Monad) Monad {
	return Suspend{
		functor: s.functor.Map(func(x Any) Any {
			return x.(Monad).Chain(f)
		}),
	}
}

func (s Suspend) Map(f func(Any) Any) Functor {
	return s.Chain(func(x Any) Monad {
		return Free_.Of(f(x)).(Monad)
	}).(Functor)
}

func (s Suspend) Run() Any {
	// Mutated state
	var x Free = s
	for {
		if _, ok := x.(Return); ok {
			break
		} else if s, ok := x.(Suspend); ok {
			var (
				a = s.functor.(Free)
				b = a.Run().(Free)
			)
			x = b
		}
	}
	return x.(Return).Run()
}

var (
	Free_ = free{}
)

type free struct{}

func (f free) Of(x Any) Point {
	return NewReturn(x)
}

func (f free) Lift(x Functor) Free {
	return NewSuspend(x.Map(func(y Any) Any {
		return NewReturn(y)
	}))
}
