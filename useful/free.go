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

type ret struct {
	val Any
}

func NewReturn(x Any) ret {
	return ret{
		val: x,
	}
}

func (r ret) Of(x Any) Point {
	return NewReturn(x)
}

func (r ret) Chain(f func(Any) Monad) Monad {
	return f(r.val)
}

func (r ret) Map(f func(Any) Any) Functor {
	return r.Chain(func(x Any) Monad {
		return Free_.Of(f(x)).(Monad)
	}).(Functor)
}

func (r ret) Run() Any {
	return r.val
}

type suspend struct {
	functor Functor
}

func NewSuspend(f Functor) suspend {
	return suspend{
		functor: f,
	}
}

func (s suspend) Of(x Any) Point {
	return NewReturn(x)
}

func (s suspend) Chain(f func(Any) Monad) Monad {
	return suspend{
		functor: s.functor.Map(func(x Any) Any {
			return x.(Monad).Chain(f)
		}),
	}
}

func (s suspend) Map(f func(Any) Any) Functor {
	return s.Chain(func(x Any) Monad {
		return Free_.Of(f(x)).(Monad)
	}).(Functor)
}

func (s suspend) Run() Any {
	// Mutated state
	var x Free = s
	for {
		if _, ok := x.(ret); ok {
			break
		} else if s, ok := x.(suspend); ok {
			var (
				a = s.functor.(Free)
				b = a.Run().(Free)
			)
			x = b
		}
	}
	return x.(ret).Run()
}

var (
	Free_ = free_{}
)

type free_ struct{}

func (f free_) As(x Any) Free {
	return x.(Free)
}

func (f free_) Ref() Free {
	return ret{}
}

func (f free_) Of(x Any) Point {
	return NewReturn(x)
}

func (f free_) Lift(x Functor) Free {
	return NewSuspend(x.Map(func(y Any) Any {
		return NewReturn(y)
	}))
}
