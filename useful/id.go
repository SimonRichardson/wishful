package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Id struct {
	x AnyVal
}

func NewId(x AnyVal) Id {
	return Id{
		x: x,
	}
}

func (x Id) Of(v AnyVal) Point {
	return NewId(v)
}

func (x Id) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x Id) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Id) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x Id) Map(f func(v AnyVal) AnyVal) Functor {
	return x.Chain(func(x AnyVal) Monad {
		return NewId(f(x))
	}).(Functor)
}

func (x Id) Extract() AnyVal {
	return x.x
}

func (x Id) Extend(f func(p Comonad) AnyVal) Comonad {
	return x.Map(func(y AnyVal) AnyVal {
		fun := NewFunction(f)
		res, _ := fun.Call(x.Of(y))
		return res
	}).(Comonad)
}

func (x Id) Traverse(f func(x AnyVal) AnyVal) Traversable {
	return f(x.x).(Functor).Map(func(x AnyVal) AnyVal {
		return NewId(x)
	}).(Traversable)
}

func (x Id) Sequence() AnyVal {
	return x.Traverse(Identity)
}
