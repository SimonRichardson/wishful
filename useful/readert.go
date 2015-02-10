package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type ReaderT struct {
	m   Point
	Run func(x Any) Point
}

func NewReaderT(m Point) ReaderT {
	return ReaderT{
		m: m,
		Run: func(x Any) Point {
			return nil
		},
	}
}

func (x ReaderT) Lift(m Functor) ReaderT {
	return ReaderT{
		m: x.m,
		Run: func(b Any) Point {
			return m.Map(func(c Any) Any {
				return Tuple2{_1: c, _2: b}
			}).(Point)
		},
	}
}

func (x ReaderT) Of(a Any) Point {
	return ReaderT{
		m: x.m,
		Run: func(b Any) Point {
			return x.m.Of(a)
		},
	}
}

func (x ReaderT) Chain(f func(a Any) Monad) Monad {
	return ReaderT{
		m: x.m,
		Run: func(b Any) Point {
			result := x.Run(b)
			return result.(Monad).Chain(func(t Any) Monad {
				return f(t).(ReaderT).Run(b).(Monad)
			}).(Point)
		},
	}
}

func (x ReaderT) Map(f Morphism) Functor {
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

func (f readerT_) As(x Any) ReaderT {
	return x.(ReaderT)
}

func (f readerT_) Ref() ReaderT {
	return ReaderT{}
}
