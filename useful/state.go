package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type state struct {
	Run func(a Any) (Any, Any)
}

func State(x, y Any) state {
	return state{
		Run: func(x Any) (Any, Any) {
			return x, y
		},
	}
}

func (x state) Of(y Any) Point {
	return state{func(z Any) (Any, Any) {
		return y, z
	}}
}

func (x state) Ap(v Applicative) Applicative {
	app := x.Chain(func(f Any) Monad {
		fun := v.(Functor)
		app := fun.Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		})
		return app.(Monad)
	})
	return app.(Applicative)
}

func (x state) Chain(f func(x Any) Monad) Monad {
	return state{func(s Any) (Any, Any) {
		a, b := x.Run(s)
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return res.(state).Run(b)
	}}
}

func (x state) Map(f func(x Any) Any) Functor {
	fun := x.Chain(func(y Any) Monad {
		return x.Of(f(y)).(Monad)
	})
	return fun.(Functor)
}

// Derived

func (x state) EvalState(y Any) Any {
	a, _ := x.Run(y)
	return a
}

func (x state) ExecState(y Any) Any {
	_, b := x.Run(y)
	return b
}

func (x state) Get() state {
	return state{func(z Any) (Any, Any) {
		return z, z
	}}
}

func (x state) Modify(f func(x Any) Any) state {
	return state{func(z Any) (Any, Any) {
		fun := NewFunction(f)
		res, _ := fun.Call(z)
		return nil, res
	}}
}

func (x state) Put(a Any, b Any) state {
	return state{func(z Any) (Any, Any) {
		return a, b
	}}
}

var (
	State_ = state_{}
)

type state_ struct{}

func (f state_) As(x Any) state {
	return x.(state)
}

func (f state_) Ref() state {
	return state{}
}

func (r state_) Of(x Any) Point {
	return state{}.Of(x)
}
