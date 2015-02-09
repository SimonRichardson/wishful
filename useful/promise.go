package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type promise struct {
	Fork func(resolve func(x Any) Any) Any
}

func Promise(f func(resolve func(x Any) Any) Any) promise {
	return promise{
		Fork: f,
	}
}

func (x promise) Of(v Any) Point {
	return promise{func(resolve func(x Any) Any) Any {
		return resolve(v)
	}}
}

func (x promise) Ap(v Applicative) Applicative {
	return promise{func(resolve func(x Any) Any) Any {
		return x.Fork(func(f Any) Any {
			fun := v.(Functor)
			pro := fun.Map(func(x Any) Any {
				fun := NewFunction(f)
				res, _ := fun.Call(x)
				return res
			})
			return pro.(promise).Fork(resolve)
		})
	}}
}

func (x promise) Chain(f func(v Any) Monad) Monad {
	return promise{func(resolve func(x Any) Any) Any {
		return x.Fork(func(a Any) Any {
			p := f(a).(promise)
			return p.Fork(resolve)
		})
	}}
}

func (x promise) Map(f func(v Any) Any) Functor {
	return promise{func(resolve func(v Any) Any) Any {
		return x.Fork(func(a Any) Any {
			return resolve(f(a))
		})
	}}
}

func (x promise) Extract() Any {
	return x.Fork(Identity)
}

func (x promise) Extend(f func(p Comonad) Any) Comonad {
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

func (f promise_) As(x Any) promise {
	return x.(promise)
}

func (f promise_) Ref() promise {
	return promise{}
}

func (f promise_) Of(x Any) Point {
	return promise{}.Of(x)
}
