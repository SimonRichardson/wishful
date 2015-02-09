package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type endo struct {
	Fork func(v Any) Any
}

func Endo(x func(v Any) Any) endo {
	return endo{
		Fork: x,
	}
}

func (x endo) Of(v Any) Point {
	return Endo(func(x Any) Any {
		return v
	})
}

func (x endo) Empty() Monoid {
	return Endo(func(v Any) Any {
		return v
	})
}

func (x endo) Concat(y Semigroup) Semigroup {
	return Endo(func(v Any) Any {
		a := y.(endo)
		return x.Fork(a.Fork(v))
	})
}

func (x endo) Map(f func(v Any) Any) Functor {
	return Endo(func(v Any) Any {
		return f(x.Fork(v))
	})
}

var (
	Endo_ = endo_{}
)

type endo_ struct{}

func (e endo_) As(x Any) endo {
	return x.(endo)
}

func (e endo_) Ref() endo {
	return endo{}
}

func (f endo_) Of(x Any) Point {
	return endo{}.Of(x)
}
