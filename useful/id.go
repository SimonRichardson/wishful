package useful

import (
	. "github.com/SimonRichardson/wishful"
)

type Id struct {
	x AnyVal
}

func NewId(x AnyVal) Id {
	return Id{
		x: x,
	}
}

func (x Id) Of(v AnyVal) Applicative {
	return NewId(v)
}

func (x Id) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x Id) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Id) Concat(y Semigroup) Semigroup {
	return fromMonadToSemigroupConcat(x, y)
}

func (x Id) Map(f func(v AnyVal) AnyVal) Functor {
	return NewId(f(x.x))
}
