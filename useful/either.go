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

// right

type left struct {
	x Any
}

func Left(x Any) left {
	return left{
		x: x,
	}
}

func (x left) Of(v Any) Point {
	return Right(v)
}

func (x left) Ap(v Applicative) Applicative {
	return x
}

func (x left) Chain(f func(v Any) Monad) Monad {
	return x
}

func (x left) Map(f func(v Any) Any) Functor {
	return x
}

func (x left) Concat(y Semigroup) Semigroup {
	return x
}

func (x left) Swap() Monad {
	return Right(x.x)
}

func (x left) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return Left(f(x.x))
}

func (x left) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x left) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x left) Traverse(f func(Any) Any, p Point) Functor {
	return p.Of(Left(x.x)).(Functor)
}

// right

type right struct {
	x Any
}

func Right(x Any) right {
	return right{
		x: x,
	}
}

func (x right) Of(v Any) Point {
	return Right(v)
}

func (x right) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x right) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x right) Map(f func(v Any) Any) Functor {
	res := x.Chain(func(v Any) Monad {
		return Right(f(v))
	})
	return res.(Functor)
}

func (x right) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x right) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x right) Swap() Monad {
	return Left(x.x)
}

func (x right) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return Right(g(x.x))
}

func (x right) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x right) Traverse(f func(Any) Any, p Point) Functor {
	return f(x.x).(Functor).Map(func(a Any) Any {
		return Either_.Of(a)
	})
}

// Either_

var (
	Either_ = either_{}
)

type either_ struct{}

func (e either_) As(x Any) Either {
	return x.(Either)
}

func (e either_) Ref() Either {
	return right{}
}

func (e either_) Of(x Any) Point {
	return Right(x)
}
