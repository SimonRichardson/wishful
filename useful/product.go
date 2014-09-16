package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	productConcat = fromMonadToSemigroupConcat(func(a Semigroup, b Semigroup) Any {
		// This is a bit horrid
		x, _ := FromAnyToInt(a)
		y, _ := FromAnyToInt(b)
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

func (x Product) Of(v Any) Point {
	p, _ := FromAnyToInt(v)
	return NewProduct(p)
}

func (x Product) Empty() Monoid {
	return NewProduct(Int(1))
}

func (x Product) Chain(f func(Any) Monad) Monad {
	return f(x.x)
}

func (x Product) Concat(y Semigroup) Semigroup {
	return productConcat(x, y)
}

func (x Product) Map(f func(Any) Any) Functor {
	return x.Chain(func(x Any) Monad {
		p, _ := FromAnyToInt(f(x))
		return NewProduct(p)
	}).(Functor)
}
