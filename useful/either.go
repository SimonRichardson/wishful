package useful

import . "github.com/SimonRichardson/wishful/wishful"

type Either interface {
	Of(Any) Point
	Ap(Applicative) Applicative
	Chain(func(Any) Monad) Monad
	Concat(Semigroup) Semigroup
	Map(Morphism) Functor
	Bimap(Morphism, Morphism) Monad
	Fold(Morphism, Morphism) Any
	Swap() Monad
	Sequence(Point) Any
	Traverse(Morphism, Point) Functor
}

// right

type left struct {
	x Any
}

func NewLeft(x Any) left {
	return left{
		x: x,
	}
}

func (x left) Of(v Any) Point {
	return NewRight(v)
}

func (x left) Ap(v Applicative) Applicative {
	return x
}

func (x left) Chain(f func(Any) Monad) Monad {
	return x
}

func (x left) Map(f Morphism) Functor {
	return x
}

func (x left) Concat(y Semigroup) Semigroup {
	return x
}

func (x left) Swap() Monad {
	return NewRight(x.x)
}

func (x left) Bimap(f Morphism, g Morphism) Monad {
	return NewLeft(f(x.x))
}

func (x left) Fold(f Morphism, g Morphism) Any {
	return f(x.x)
}

func (x left) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x left) Traverse(f Morphism, p Point) Functor {
	return p.Of(NewLeft(x.x)).(Functor)
}

// right

type right struct {
	x Any
}

func NewRight(x Any) right {
	return right{
		x: x,
	}
}

func (x right) Of(v Any) Point {
	return NewRight(v)
}

func (x right) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x right) Chain(f func(v Any) Monad) Monad {
	return f(x.x)
}

func (x right) Map(f Morphism) Functor {
	res := x.Chain(func(v Any) Monad {
		return NewRight(f(v))
	})
	return res.(Functor)
}

func (x right) Concat(y Semigroup) Semigroup {
	return concat(x, y)
}

func (x right) Fold(f Morphism, g Morphism) Any {
	return g(x.x)
}

func (x right) Swap() Monad {
	return NewLeft(x.x)
}

func (x right) Bimap(f Morphism, g Morphism) Monad {
	return NewRight(g(x.x))
}

func (x right) Sequence(p Point) Any {
	return x.Traverse(Identity, p)
}

func (x right) Traverse(f Morphism, p Point) Functor {
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
	return NewRight(x)
}
