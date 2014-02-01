package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type State struct {
	Run func(a AnyVal) (AnyVal, AnyVal)
}

func (x State) Of(y AnyVal) Point {
	return State{func(z AnyVal) (AnyVal, AnyVal) {
		return y, z
	}}
}

func (x State) Ap(v Applicative) Applicative {
	app := x.Chain(func(f AnyVal) Monad {
		fun := v.(Functor)
		app := fun.Map(func(x AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		})
		return app.(Monad)
	})
	return app.(Applicative)
}

func (o State) Chain(f func(x AnyVal) Monad) Monad {
	return State{func(s AnyVal) (AnyVal, AnyVal) {
		a, b := o.Run(s)
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return res.(State).Run(b)
	}}
}

func (o State) Map(f func(x AnyVal) AnyVal) Functor {
	fun := o.Chain(func(x AnyVal) Monad {
		return State{}.Of(f(x)).(Monad)
	})
	return fun.(Functor)
}

func (o State) EvalState(x AnyVal) AnyVal {
	a, _ := o.Run(x)
	return a
}

func (o State) ExecState(x AnyVal) AnyVal {
	_, b := o.Run(x)
	return b
}
