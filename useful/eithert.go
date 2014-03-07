package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type EitherT struct {
	m   Point
	Run AnyVal
}

func NewEitherT(m Point) EitherT {
	return EitherT{
		m:   m,
		Run: Empty{},
	}
}

func (x EitherT) Of(v AnyVal) Point {
	return EitherT{
		m:   x.m,
		Run: x.m.Of(NewRight(v)),
	}
}

func (x EitherT) From(v AnyVal) EitherT {
	return EitherT{
		m:   x.m,
		Run: v,
	}
}

func (x EitherT) Fold(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) AnyVal {
	return x.Run.(Monad).Chain(func(o AnyVal) Monad {
		return x.m.Of(o.(Foldable).Fold(f, g)).(Monad)
	})
}

func (x EitherT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f AnyVal) Monad {
		return v.(Functor).Map(func(x AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		}).(Monad)
	})
	return mon.(Applicative)
}

func (x EitherT) Chain(f func(v AnyVal) Monad) Monad {
	mon := x.Run.(Monad)
	tra := EitherT{
		m: x.m,
		Run: mon.Chain(func(y AnyVal) Monad {
			return y.(Foldable).Fold(
				func(v AnyVal) AnyVal {
					return x.m.Of(NewLeft(v))
				},
				func(v AnyVal) AnyVal {
					return f(v).(EitherT).Run
				},
			).(Monad)
		}),
	}
	return tra
}

func (x EitherT) Map(f func(v AnyVal) AnyVal) Functor {
	mon := x.Chain(func(y AnyVal) Monad {
		app := NewEitherT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}

// Derived

func (x EitherT) Swap() Monad {
	return x.Fold(
		func(v AnyVal) AnyVal {
			return NewRight(v)
		},
		func(v AnyVal) AnyVal {
			return NewLeft(v)
		},
	).(Monad)
}

func (x EitherT) Bimap(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) Monad {
	return x.Fold(
		func(v AnyVal) AnyVal {
			return NewLeft(f(v))
		},
		func(v AnyVal) AnyVal {
			return NewRight(g(v))
		},
	).(Monad)
}
