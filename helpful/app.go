package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	ST  = NewStateT(Promise_.Ref())
	App = ReaderT(ST)
)

var (
	App_ = app{}
)

type app struct{}

func (a app) LiftT(t Monad) Monad {
	p := Promise_.As(t)
	return App.Lift(ST.Func(func(x Any) Point {
		return Promise(func(resolve func(Any) Any) Any {
			return p.Fork(func(y Any) Any {
				return resolve(NewTuple2(x, y))
			})
		})
	}))
}

func (a app) Lift() func(Any) Any {
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

func (a app) Map(f func(Any) Any) Monad {
	var (
		x = func(x Any) Any {
			return App.Of(x)
		}
	)
	return App.Chain(func(y Any) Monad {
		return Compose(x)(f)(y).(Monad)
	})
}

func (a app) MapLift(lift func(Any) Any) func(func(Any) Any) Any {
	return func(f func(Any) Any) Any {
		return App.Chain(func(y Any) Monad {
			return Compose(lift)(f)(y).(Monad)
		})
	}
}
