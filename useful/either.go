package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Either interface {
	Of(Any) Point
	Ap(Applicative) Applicative
	Chain(func(Any) Monad) Monad
	Concat(Semigroup) Semigroup
	Map(func(Any) Any) Functor
	Bimap(func(Any) Any, func(Any) Any) Monad
	Fold(func(Any) Any, func(Any) Any) Any
	Swap() Monad
	Sequence(Point) Any
	Traverse(func(Any) Any, Point) Functor
}

type Left struct {
	x Any
}

type Right struct {
	x Any
}

func NewLeft(x Any) Left {
	return Left{
		x: x,
	}
}

func NewRight(x Any) Right {
	return Right{
		x: x,
	}
}

func (x Left) Of(v Any) Point {
	return NewRight(v)
}

func (x Right) Of(v Any) Point {
	return NewRight(v)
}

func (x Left) Ap(v Applicative) Applicative {
	return x
}

func (x Right) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x Left) Chain(f func(v Any) Monad) Monad {
	return x
}

func (x Right) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x Left) Map(f func(v Any) Any) Functor {
	return x
}

func (x Right) Map(f func(v Any) Any) Functor {
	res := x.Chain(func(v Any) Monad {
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

func (x Left) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return NewLeft(f(x.x))
}

func (x Right) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return NewRight(g(x.x))
}

func (x Left) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Right) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Left) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x Right) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x Left) Traverse(f func(Any) Any, p Point) Functor {
	return p.Of(NewLeft(x.x)).(Functor)
}

func (x Right) Traverse(f func(Any) Any, p Point) Functor {
	return f(x.x).(Functor).Map(func(a Any) Any {
		return NewRight(a)
	})
}
