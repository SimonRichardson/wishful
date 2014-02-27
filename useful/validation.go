package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Validation interface {
	Fold(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) AnyVal
	Bimap(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) Monad
}

type Failure struct {
	x AnyVal
}

type Success struct {
	x AnyVal
}

func NewFailure(x AnyVal) Failure {
	return Failure{
		x: x,
	}
}

func NewSuccess(x AnyVal) Success {
	return Success{
		x: x,
	}
}

func (x Failure) Of(v AnyVal) Point {
	return NewSuccess(v)
}

func (x Success) Of(v AnyVal) Point {
	return NewSuccess(v)
}

func (x Failure) Ap(v Applicative) Applicative {
	return v.(Validation).Fold(
		func(y AnyVal) AnyVal {
			return NewFailure(concatAnyvals(x.x)(y))
		},
		func(y AnyVal) AnyVal {
			return NewFailure(x.x)
		},
	).(Applicative)
}

func (x Success) Ap(v Applicative) Applicative {
	return v.(Functor).Map(func(g AnyVal) AnyVal {
		fun := NewFunction(x.x)
		res, _ := fun.Call(g)
		return res
	}).(Applicative)
}

func (x Failure) Chain(f func(v AnyVal) Monad) Monad {
	return x
}

func (x Success) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x Failure) Map(f func(v AnyVal) AnyVal) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x Success) Map(f func(v AnyVal) AnyVal) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x Failure) Concat(y Semigroup) Semigroup {
	a := y.(Validation)
	b := a.Bimap(
		concatAnyvals(x.x),
		Identity,
	)
	return b.(Semigroup)
}

func (x Success) Concat(y Semigroup) Semigroup {
	a := y.(Functor)
	b := a.Map(concatAnyvals(x.x))
	return b.(Semigroup)
}

// Derived

func (x Failure) Fold(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) AnyVal {
	return f(x.x)
}

func (x Success) Fold(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) AnyVal {
	return g(x.x)
}

func (x Failure) Bimap(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) Monad {
	return NewFailure(f(x.x))
}

func (x Success) Bimap(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) Monad {
	return NewSuccess(g(x.x))
}
