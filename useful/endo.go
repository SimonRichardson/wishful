package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Endo struct {
	Fork func(v Any) Any
}

func NewEndo(x func(v Any) Any) Endo {
	return Endo{
		Fork: x,
	}
}

func (x Endo) Of(v Any) Point {
	return NewEndo(func(x Any) Any {
		return v
	})
}

func (x Endo) Empty() Monoid {
	return NewEndo(func(v Any) Any {
		return v
	})
}

func (x Endo) Concat(y Semigroup) Semigroup {
	return NewEndo(func(v Any) Any {
		a := y.(Endo)
		return x.Fork(a.Fork(v))
	})
}

func (x Endo) Map(f func(v Any) Any) Functor {
	return NewEndo(func(v Any) Any {
		return f(x.Fork(v))
	})
}
