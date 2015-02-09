package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type readerT struct {
	m   Point
	Run func(x Any) Point
}

func ReaderT(m Point) readerT {
	return readerT{
		m: m,
		Run: func(x Any) Point {
			return nil
		},
	}
}

func (x readerT) Lift(m Functor) readerT {
	return readerT{
		m: x.m,
		Run: func(b Any) Point {
			return m.Map(func(c Any) Any {
				return Tuple2{_1: c, _2: b}
			}).(Point)
		},
	}
}

func (x readerT) Of(a Any) Point {
	return readerT{
		m: x.m,
		Run: func(b Any) Point {
			return x.m.Of(a)
		},
	}
}

func (x readerT) Chain(f func(a Any) Monad) Monad {
	return readerT{
		m: x.m,
		Run: func(b Any) Point {
			result := x.Run(b)
			return result.(Monad).Chain(func(t Any) Monad {
				return f(t).(readerT).Run(b).(Monad)
			}).(Point)
		},
	}
}

func (x readerT) Map(f func(x Any) Any) Functor {
	return x.Chain(func(a Any) Monad {
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return x.Of(res).(Monad)
	}).(Functor)
}

var (
	ReaderT_ = readerT_{}
)

type readerT_ struct{}

func (f readerT_) As(x Any) readerT {
	return x.(readerT)
}

func (f readerT_) Ref() readerT {
	return readerT{}
}
