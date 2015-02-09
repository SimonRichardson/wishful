package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type reader struct {
	Run func(Any) Any
}

func (r reader) Of(x Any) Point {
	return reader{
		Run: func(y Any) Any {
			return x
		},
	}
}

func (r reader) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(r, v)
}

func (r reader) Chain(f func(Any) Monad) Monad {
	return reader{
		Run: func(x Any) Any {
			y := f(r.Run(x)).(reader)
			return y.Run(x)
		},
	}
}

func (r reader) Map(f func(Any) Any) Functor {
	return r.Chain(func(x Any) Monad {
		return reader{}.Of(f(x)).(Monad)
	}).(Functor)
}

var (
	Reader_ = reader_{}
)

type reader_ struct{}

func (f reader_) As(x Any) reader {
	return x.(reader)
}

func (f reader_) Ref() reader {
	return reader{}
}

func (r reader_) Of(x Any) Point {
	return reader{}.Of(x)
}

func (r reader_) Ask() reader {
	return reader{
		Run: func(a Any) Any {
			return a
		},
	}
}
