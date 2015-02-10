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

type some struct {
	x Any
}

type none struct {
}

func NewSome(x Any) some {
	return some{
		x: x,
	}
}

func NewNone() none {
	return none{}
}

func (x some) Of(v Any) Point {
	return NewSome(v)
}

func (x none) Of(v Any) Point {
	return NewSome(v)
}

func (x some) Empty() Monoid {
	return NewNone()
}

func (x none) Empty() Monoid {
	return NewNone()
}

func (x some) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x none) Ap(v Applicative) Applicative {
	return x
}

func (x some) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x none) Chain(f func(v Any) Monad) Monad {
	return x
}

func (x some) Fold(f func(v Any) Any, g func() Any) Any {
	return f(x.x)
}

func (x none) Fold(f func(v Any) Any, g func() Any) Any {
	return g()
}

func (x some) Map(f func(v Any) Any) Functor {
	res := x.Chain(func(v Any) Monad {
		return NewSome(f(v))
	})
	return res.(Functor)
}

func (x none) Map(f func(v Any) Any) Functor {
	return x
}

func (x some) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x none) Concat(y Semigroup) Semigroup {
	return x
}

// Derived

func (x some) GetOrElse(f func() Any) Any {
	return x.x
}

func (x none) GetOrElse(f func() Any) Any {
	return f()
}

func (x some) OrElse(y Option) Option {
	return some{}.Of(x.x).(Option)
}

func (x none) OrElse(y Option) Option {
	return y
}

//

var (
	Option_ = option_{}
)

type option_ struct{}

func (o option_) As(x Any) Option {
	return x.(Option)
}

func (o option_) Ref() Option {
	return some{}
}

func (o option_) Of(x Any) Point {
	return NewSome(x)
}

func (o option_) Empty() Monoid {
	return NewNone()
}
