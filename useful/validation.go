package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Validation interface {
	Of(v Any) Point
	Ap(v Applicative) Applicative
	Chain(f func(v Any) Monad) Monad
	Concat(y Semigroup) Semigroup
	Map(f func(v Any) Any) Functor
	Fold(f func(v Any) Any, g func(v Any) Any) Any
	Bimap(f func(v Any) Any, g func(v Any) Any) Monad
}

type Failure struct {
	x Any
}

type Success struct {
	x Any
}

func NewFailure(x Any) Failure {
	return Failure{
		x: x,
	}
}

func NewSuccess(x Any) Success {
	return Success{
		x: x,
	}
}

func (x Failure) Of(v Any) Point {
	return NewSuccess(v)
}

func (x Success) Of(v Any) Point {
	return NewSuccess(v)
}

func (x Failure) Ap(v Applicative) Applicative {
	return v.(Validation).Fold(
		func(y Any) Any {
			return NewFailure(concatAnyvals(x.x)(y))
		},
		func(y Any) Any {
			return NewFailure(x.x)
		},
	).(Applicative)
}

func (x Success) Ap(v Applicative) Applicative {
	return v.(Functor).Map(func(g Any) Any {
		fun := NewFunction(x.x)
		res, _ := fun.Call(g)
		return res
	}).(Applicative)
}

func (x Failure) Chain(f func(v Any) Monad) Monad {
	return x
}

func (x Success) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x Failure) Map(f func(v Any) Any) Functor {
	return x.Bimap(Identity, f).(Functor)
}

func (x Success) Map(f func(v Any) Any) Functor {
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

func (x Failure) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Success) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Failure) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return NewFailure(f(x.x))
}

func (x Success) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return NewSuccess(g(x.x))
}
