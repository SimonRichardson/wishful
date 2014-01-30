package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Endo struct {
	Fork func(v AnyVal) AnyVal
}

func NewEndo(x func(v AnyVal) AnyVal) Endo {
	return Endo{
		Fork: x,
	}
}

func (x Endo) Of(v AnyVal) Point {
	return NewEndo(func(x AnyVal) AnyVal {
		return v
	})
}

func (x Endo) Empty() Monoid {
	return NewEndo(func(v AnyVal) AnyVal {
		return v
	})
}

func (x Endo) Concat(y Semigroup) Semigroup {
	return NewEndo(func(v AnyVal) AnyVal {
		a := y.(Endo)
		return x.Fork(a.Fork(v))
	})
}
