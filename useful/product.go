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

type product struct {
	x Int
}

func Product(x Int) product {
	return product{
		x: x,
	}
}

func (x product) Of(v Any) Point {
	p, _ := FromAnyToInt(v)
	return Product(p)
}

func (x product) Empty() Monoid {
	return Product(Int(1))
}

func (x product) Chain(f func(Any) Monad) Monad {
	return f(x.x)
}

func (x product) Concat(y Semigroup) Semigroup {
	return productConcat(x, y)
}

func (x product) Map(f func(Any) Any) Functor {
	return x.Chain(func(x Any) Monad {
		p, _ := FromAnyToInt(f(x))
		return Product(p)
	}).(Functor)
}

var (
	Product_ = product_{}
)

type product_ struct{}

func (f product_) As(x Any) product {
	return x.(product)
}

func (f product_) Ref() product {
	return product{}
}

func (f product_) Of(x Any) Point {
	return product{}.Of(x)
}
