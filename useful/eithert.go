package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type eitherT struct {
	m   Point
	Run Any
}

func EitherT(m Point) eitherT {
	return eitherT{
		m:   m,
		Run: Empty{},
	}
}

func (x eitherT) Of(v Any) Point {
	return eitherT{
		m:   x.m,
		Run: x.m.Of(Right(v)),
	}
}

func (x eitherT) From(v Any) eitherT {
	return eitherT{
		m:   x.m,
		Run: v,
	}
}

func (x eitherT) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return x.Run.(Monad).Chain(func(o Any) Monad {
		return x.m.Of(o.(Foldable).Fold(f, g)).(Monad)
	})
}

func (x eitherT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f Any) Monad {
		return v.(Functor).Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		}).(Monad)
	})
	return mon.(Applicative)
}

func (x eitherT) Chain(f func(v Any) Monad) Monad {
	mon := x.Run.(Monad)
	tra := eitherT{
		m: x.m,
		Run: mon.Chain(func(y Any) Monad {
			return y.(Foldable).Fold(
				func(v Any) Any {
					return x.m.Of(Left(v))
				},
				func(v Any) Any {
					return f(v).(eitherT).Run
				},
			).(Monad)
		}),
	}
	return tra
}

func (x eitherT) Map(f func(v Any) Any) Functor {
	mon := x.Chain(func(y Any) Monad {
		app := EitherT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}

func (x eitherT) Swap() Monad {
	return x.Fold(
		func(v Any) Any {
			return Right(v)
		},
		func(v Any) Any {
			return Left(v)
		},
	).(Monad)
}

func (x eitherT) Bimap(f func(v Any) Any, g func(v Any) Any) Monad {
	return x.Fold(
		func(v Any) Any {
			return Left(f(v))
		},
		func(v Any) Any {
			return Right(g(v))
		},
	).(Monad)
}

var (
	EitherT_ = eitherT_{}
)

type eitherT_ struct{}

func (e eitherT_) As(x Any) eitherT {
	return x.(eitherT)
}

func (e eitherT_) Ref() eitherT {
	return eitherT{}
}
