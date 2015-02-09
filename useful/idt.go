package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type idT struct {
	m   Point
	Run Any
}

func IdT(m Point) idT {
	return idT{
		m:   m,
		Run: Empty{},
	}
}

func (x idT) Of(v Any) Point {
	return idT{
		m:   x.m,
		Run: x.m.Of(v),
	}
}

func (x idT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f Any) Monad {
		return v.(Functor).Map(func(x Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(x)
			return res
		}).(Monad)
	})
	return mon.(Applicative)
}

func (x idT) Chain(f func(v Any) Monad) Monad {
	mon := x.Run.(Monad)
	tra := idT{
		m: x.m,
		Run: mon.Chain(func(y Any) Monad {
			idt := f(y).(idT)
			return idt.Run.(Monad)
		}),
	}
	return tra
}

func (x idT) Map(f func(v Any) Any) Functor {
	mon := x.Chain(func(y Any) Monad {
		app := IdT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}

var (
	IdT_ = idT_{}
)

type idT_ struct{}

func (f idT_) As(x Any) idT {
	return x.(idT)
}

func (f idT_) Ref() idT {
	return idT{}
}
