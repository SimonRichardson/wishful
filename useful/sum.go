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

func (x Sum) Of(v Any) Point {
	p, _ := FromAnyToInt(v)
	return NewSum(p)
}

func (x Sum) Empty() Monoid {
	return NewSum(Int(0))
}

func (x Sum) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x Sum) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x Sum) Map(f func(v Any) Any) Functor {
	return x.Chain(func(x Any) Monad {
		p, _ := FromAnyToInt(f(x))
		return NewSum(p)
	}).(Functor)
}
