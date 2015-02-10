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

func (x Sum) Chain(f func(Any) Monad) Monad {
	return f(x.x)
}

func (x Sum) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x Sum) Map(f Morphism) Functor {
	return x.Chain(func(x Any) Monad {
		p, _ := FromAnyToInt(f(x))
		return NewSum(p)
	}).(Functor)
}

var (
	Sum_ = sum_{}
)

type sum_ struct{}

func (f sum_) As(x Any) Sum {
	return x.(Sum)
}

func (f sum_) Ref() Sum {
	return Sum{}
}

func (f sum_) Of(x Any) Point {
	return Sum{}.Of(x)
}
