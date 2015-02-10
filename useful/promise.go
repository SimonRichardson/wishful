package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Promise struct {
	Fork func(func(x Any) Any) Any
}

func NewPromise(f func(func(x Any) Any) Any) Promise {
	return Promise{
		Fork: f,
	}
}

func (x Promise) Of(v Any) Point {
	return Promise{func(resolve func(x Any) Any) Any {
		return resolve(v)
	}}
}

func (x Promise) Ap(v Applicative) Applicative {
	return Promise{func(resolve func(x Any) Any) Any {
		return x.Fork(func(f Any) Any {
			fun := v.(Functor)
			pro := fun.Map(func(x Any) Any {
				fun := NewFunction(f)
				res, _ := fun.Call(x)
				return res
			})
			return pro.(Promise).Fork(resolve)
		})
	}}
}

func (x Promise) Chain(f func(v Any) Monad) Monad {
	return Promise{func(resolve func(x Any) Any) Any {
		return x.Fork(func(a Any) Any {
			p := f(a).(Promise)
			return p.Fork(resolve)
		})
	}}
}

func (x Promise) Map(f func(v Any) Any) Functor {
	return Promise{func(resolve func(v Any) Any) Any {
		return x.Fork(func(a Any) Any {
			return resolve(f(a))
		})
	}}
}

func (x Promise) Extract() Any {
	return x.Fork(Identity)
}

func (x Promise) Extend(f func(p Comonad) Any) Comonad {
	return x.Map(func(y Any) Any {
		fun := NewFunction(f)
		res, _ := fun.Call(x.Of(y))
		return res
	}).(Comonad)
}

var (
	Promise_ = promise_{}
)

type promise_ struct{}

func (f promise_) As(x Any) Promise {
	return x.(Promise)
}

func (f promise_) Ref() Promise {
	return Promise{}
}

func (f promise_) Of(x Any) Point {
	return Promise{}.Of(x)
}
