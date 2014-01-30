package useful

import (
	"errors"
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
	if obj, ok := v.(int); ok {
		return NewSum(Int(obj))
	} else if obj, ok := v.(Int); ok {
		return NewSum(obj)
	} else {
		panic(errors.New("Invalid type for Sum"))
	}
}

func (x Sum) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Sum) Concat(y Semigroup) Semigroup {
	return fromMonadToSemigroupConcat(x, y)
}

func (x Sum) Map(f func(v AnyVal) AnyVal) Functor {
	return x.Chain(func(x AnyVal) Monad {
		return NewSum(f(x).(Int))
	}).(Functor)
}
