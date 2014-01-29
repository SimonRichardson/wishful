package useful

import (
	. "github.com/SimonRichardson/wishful"
)

type Product struct {
	x AnyVal
}

func NewProduct(x AnyVal) Product {
	return Product{x}
}

func (o Product) Of(x AnyVal) Applicative {
	return NewProduct(x)
}

func (o Product) Empty() Monoid {
	return NewProduct(1)
}

func (o Product) Chain(f func(v AnyVal) Monad) Monad {
	return f(o.x)
}

func (o Product) Concat(x Semigroup) Semigroup {
	return o.Chain(func(a AnyVal) AnyVal {
		return x.Map(func(b AnyVal) AnyVal {
			return a.(int) * b.(int)
		})
	})
}

func (o Product) Map(f func(v AnyVal) AnyVal) Functor {
	return NewProduct(f(o.x))
}
