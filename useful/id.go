package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type id struct {
	x Any
}

func Id(x Any) id {
	return id{
		x: x,
	}
}

func (x id) Of(v Any) Point {
	return Id(v)
}

func (x id) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x id) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x id) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x id) Map(f func(v Any) Any) Functor {
	return x.Chain(func(x Any) Monad {
		return Id(f(x))
	}).(Functor)
}

func (x id) Extract() Any {
	return x.x
}

func (x id) Extend(f func(p Comonad) Any) Comonad {
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

func (f id_) As(x Any) id {
	return x.(id)
}

func (f id_) Ref() id {
	return id{}
}

func (f id_) Of(x Any) Point {
	return id{}.Of(x)
}
