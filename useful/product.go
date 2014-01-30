package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	productConcat = fromMonadToSemigroupConcat(func(a Semigroup, b Semigroup) AnyVal {
		// This is a bit horrid
		x, _ := fromAnyValToInt(a)
		y, _ := fromAnyValToInt(b)
		return int(x) * int(y)
	})
)

type Product struct {
	x Int
}

func NewProduct(x Int) Product {
	return Product{
		x: x,
	}
}

func (x Product) Of(v AnyVal) Point {
	p, _ := fromAnyValToInt(v)
	return NewProduct(p)
}

func (x Product) Empty() Monoid {
	return NewProduct(Int(1))
}

func (x Product) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Product) Concat(y Semigroup) Semigroup {
	return productConcat(x, y)
}

func (x Product) Map(f func(v AnyVal) AnyVal) Functor {
	return x.Chain(func(x AnyVal) Monad {
		p, _ := fromAnyValToInt(f(x))
		return NewProduct(p)
	}).(Functor)
}
