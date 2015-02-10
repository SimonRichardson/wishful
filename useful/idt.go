package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type IdT struct {
	m   Point
	Run Any
}

func NewIdT(m Point) IdT {
	return IdT{
		m:   m,
		Run: Empty{},
	}
}

func (x IdT) Of(v Any) Point {
	return IdT{
		m:   x.m,
		Run: x.m.Of(v),
	}
}

func (x IdT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f Any) Monad {
		return v.(Functor).Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		}).(Monad)
	})
	return mon.(Applicative)
}

func (x IdT) Chain(f func(v Any) Monad) Monad {
	mon := x.Run.(Monad)
	tra := IdT{
		m: x.m,
		Run: mon.Chain(func(y Any) Monad {
			idt := f(y).(IdT)
			return idt.Run.(Monad)
		}),
	}
	return tra
}

func (x IdT) Map(f func(v Any) Any) Functor {
	mon := x.Chain(func(y Any) Monad {
		app := NewIdT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}

var (
	IdT_ = idT_{}
)

type idT_ struct{}

func (f idT_) As(x Any) IdT {
	return x.(IdT)
}

func (f idT_) Ref() IdT {
	return IdT{}
}
