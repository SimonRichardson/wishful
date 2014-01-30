package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Sum struct {
	x Int
}

func NewSum(x Int) Sum {
	return Sum{
		x: x,
	}
}

func (x Sum) Of(v AnyVal) Point {
	p, _ := fromAnyValToInt(v)
	return NewSum(p)
}

func (x Sum) Empty() Monoid {
	return NewSum(Int(0))
}

func (x Sum) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Sum) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x Sum) Map(f func(v AnyVal) AnyVal) Functor {
	return x.Chain(func(x AnyVal) Monad {
		p, _ := fromAnyValToInt(f(x))
		return NewSum(p)
	}).(Functor)
}
