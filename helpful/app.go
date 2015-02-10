package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	ST  = NewStateT(Promise_.Ref())
	App = NewReaderT(ST)
)

var (
	App_ = app{}
)

type app struct{}

func (a app) LiftT(p Promise) Monad {
	return App.Lift(ST.Func(func(x Any) Point {
		return NewPromise(func(resolve Morphism) Any {
			return p.Fork(func(y Any) Any {
				return resolve(NewTuple2(x, y))
			})
		})
	}))
}

func (a app) Lift() Morphism {
	var (
		x = func(x Any) Any {
			return App.Lift(x.(Functor))
		}
		y = func(x Any) Any {
			return ST.Lift(x.(Functor))
		}
	)
	return Compose(x)(y)
}

func (a app) Map(f Morphism) Monad {
	var (
		x = func(x Any) Any {
			return App.Of(x)
		}
	)
	return App.Chain(func(y Any) Monad {
		return Compose(x)(f)(y).(Monad)
	})
}

func (a app) MapLift(lift Morphism) func(Morphism) Any {
	return func(f Morphism) Any {
		return App.Chain(func(y Any) Monad {
			return Compose(lift)(f)(y).(Monad)
		})
	}
}
