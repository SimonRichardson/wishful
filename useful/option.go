package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Option interface {
	Of(v Any) Point
	Empty() Monoid
	Ap(v Applicative) Applicative
	Chain(f func(v Any) Monad) Monad
	Concat(y Semigroup) Semigroup
	Fold(f func(v Any) Any, g func() Any) Any
	Map(f func(v Any) Any) Functor
	GetOrElse(f func() Any) Any
	OrElse(y Option) Option
}

type Some struct {
	x Any
}

type None struct {
}

func NewSome(x Any) Some {
	return Some{
		x: x,
	}
}

func NewNone() None {
	return None{}
}

func (x Some) Of(v Any) Point {
	return NewSome(v)
}

func (x None) Of(v Any) Point {
	return NewSome(v)
}

func (x Some) Empty() Monoid {
	return NewNone()
}

func (x None) Empty() Monoid {
	return NewNone()
}

func (x Some) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x None) Ap(v Applicative) Applicative {
	return x
}

func (x Some) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x None) Chain(f func(v Any) Monad) Monad {
	return x
}

func (x Some) Fold(f func(v Any) Any, g func() Any) Any {
	return f(x.x)
}

func (x None) Fold(f func(v Any) Any, g func() Any) Any {
	return g()
}

func (x Some) Map(f func(v Any) Any) Functor {
	res := x.Chain(func(v Any) Monad {
		return NewSome(f(v))
	})
	return res.(Functor)
}

func (x None) Map(f func(v Any) Any) Functor {
	return x
}

func (x Some) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x None) Concat(y Semigroup) Semigroup {
	return x
}

// Derived

func (x Some) GetOrElse(f func() Any) Any {
	return x.x
}

func (x None) GetOrElse(f func() Any) Any {
	return f()
}

func (x Some) OrElse(y Option) Option {
	return Some{}.Of(x.x).(Option)
}

func (x None) OrElse(y Option) Option {
	return y
}
