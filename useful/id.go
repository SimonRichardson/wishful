package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Id struct {
	x Any
}

func NewId(x Any) Id {
	return Id{
		x: x,
	}
}

func (x Id) Of(v Any) Point {
	return NewId(v)
}

func (x Id) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x Id) Chain(f func(Any) Monad) Monad {
	return f(x.x)
}

func (x Id) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x Id) Map(f Morphism) Functor {
	return x.Chain(func(x Any) Monad {
		return NewId(f(x))
	}).(Functor)
}

func (x Id) Extract() Any {
	return x.x
}

func (x Id) Extend(f func(Comonad) Any) Comonad {
	return x.Map(func(y Any) Any {
		fun := NewFunction(f)
		res, _ := fun.Call(x.Of(y))
		return res
	}).(Comonad)
}

var (
	Id_ = id_{}
)

type id_ struct{}

func (f id_) As(x Any) Id {
	return x.(Id)
}

func (f id_) Ref() Id {
	return Id{}
}

func (f id_) Of(x Any) Point {
	return Id{}.Of(x)
}
