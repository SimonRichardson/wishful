package useful

import (
	. "github.com/SimonRichardson/wishful"
)

type Sum struct {
	x AnyVal
}

func NewSum(x AnyVal) Sum {
	return Sum{x}
}

func (o Sum) Of(x AnyVal) Applicative {
	return NewSum(x)
}

func (o Sum) Empty() Monoid {
	return NewSum(0)
}

func (o Sum) Chain(f func(v AnyVal) Monad) Monad {
	return f(o.x)
}

func (o Sum) Concat(x Semigroup) Semigroup {
	return o.Chain(func(a AnyVal) AnyVal {
		return x.Map(func(b AnyVal) AnyVal {
			return a.(int) + b.(int)
		})
	})
}

func (o Sum) Map(f func(v AnyVal) AnyVal) Functor {
	return NewSum(f(o.x))
}
