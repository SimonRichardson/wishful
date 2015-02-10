package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type State struct {
	Run func(a Any) (Any, Any)
}

func NewState(x, y Any) State {
	return State{
		Run: func(x Any) (Any, Any) {
			return x, y
		},
	}
}

func (x State) Of(y Any) Point {
	return State{func(z Any) (Any, Any) {
		return y, z
	}}
}

func (x State) Ap(v Applicative) Applicative {
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

func (x State) Chain(f func(Any) Monad) Monad {
	return State{func(s Any) (Any, Any) {
		a, b := x.Run(s)
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return res.(State).Run(b)
	}}
}

func (x State) Map(f Morphism) Functor {
	fun := x.Chain(func(y Any) Monad {
		return x.Of(f(y)).(Monad)
	})
	return fun.(Functor)
}

// Derived

func (x State) EvalState(y Any) Any {
	a, _ := x.Run(y)
	return a
}

func (x State) ExecState(y Any) Any {
	_, b := x.Run(y)
	return b
}

func (x State) Get() State {
	return State{func(z Any) (Any, Any) {
		return z, z
	}}
}

func (x State) Modify(f Morphism) State {
	return State{func(z Any) (Any, Any) {
		fun := NewFunction(f)
		res, _ := fun.Call(z)
		return nil, res
	}}
}

func (x State) Put(a Any, b Any) State {
	return State{func(z Any) (Any, Any) {
		return a, b
	}}
}

var (
	State_ = state_{}
)

type state_ struct{}

func (f state_) As(x Any) State {
	return x.(State)
}

func (f state_) Ref() State {
	return State{}
}

func (r state_) Of(x Any) Point {
	return State{}.Of(x)
}
