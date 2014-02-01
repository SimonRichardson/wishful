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

func (x State) Chain(f func(x AnyVal) Monad) Monad {
	return State{func(s AnyVal) (AnyVal, AnyVal) {
		a, b := x.Run(s)
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return res.(State).Run(b)
	}}
}

func (x State) Map(f func(x AnyVal) AnyVal) Functor {
	fun := x.Chain(func(y AnyVal) Monad {
		return x.Of(f(y)).(Monad)
	})
	return fun.(Functor)
}

// Derived

func (x State) EvalState(y AnyVal) AnyVal {
	a, _ := x.Run(y)
	return a
}

func (x State) ExecState(y AnyVal) AnyVal {
	_, b := x.Run(y)
	return b
}

func (x State) Get() State {
	return State{func(z AnyVal) (AnyVal, AnyVal) {
		return z, z
	}}
}

func (x State) Modify(f func(x AnyVal) AnyVal) State {
	return State{func(z AnyVal) (AnyVal, AnyVal) {
		fun := NewFunction(f)
		res, _ := fun.Call(z)
		return nil, res
	}}
}

func (x State) Put(a AnyVal, b AnyVal) State {
	return State{func(z AnyVal) (AnyVal, AnyVal) {
		return a, b
	}}
}
