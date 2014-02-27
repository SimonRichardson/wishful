package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Either interface {
	Swap() Monad
	Bimap(f func(v AnyVal) AnyVal) Monad
}

type Left struct {
	x AnyVal
}

type Right struct {
	x AnyVal
}

func NewLeft(x AnyVal) Left {
	return Left{
		x: x,
	}
}

func NewRight(x AnyVal) Right {
	return Right{
		x: x,
	}
}

func (x Left) Of(v AnyVal) Point {
	return NewRight(v)
}

func (x Right) Of(v AnyVal) Point {
	return NewRight(v)
}

func (x Left) Ap(v Applicative) Applicative {
	return x
}

func (x Right) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x Left) Chain(f func(v AnyVal) Monad) Monad {
	return x
}

func (x Right) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Left) Map(f func(v AnyVal) AnyVal) Functor {
	return x
}

func (x Right) Map(f func(v AnyVal) AnyVal) Functor {
	res := x.Chain(func(v AnyVal) Monad {
		return NewRight(f(v))
	})
	return res.(Functor)
}

func (x Left) Concat(y Semigroup) Semigroup {
	return x
}

func (x Right) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

// Derived
func (x Left) Swap() Monad {
	return NewRight(x.x)
}

func (x Right) Swap() Monad {
	return NewLeft(x.x)
}

func (x Left) Bimap(f func(v AnyVal) AnyVal) Monad {
	return NewLeft(f(x.x))
}

func (x Right) Bimap(f func(v AnyVal) AnyVal) Monad {
	return NewRight(f(x.x))
}
