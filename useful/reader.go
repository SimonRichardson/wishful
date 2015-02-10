package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Reader struct {
	Run func(Any) Any
}

func NewReader(f func(Any) Any) Reader {
	return Reader{
		Run: f,
	}
}

func (r Reader) Of(x Any) Point {
	return Reader{
		Run: func(y Any) Any {
			return x
		},
	}
}

func (r Reader) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(r, v)
}

func (r Reader) Chain(f func(Any) Monad) Monad {
	return Reader{
		Run: func(x Any) Any {
			y := f(r.Run(x)).(Reader)
			return y.Run(x)
		},
	}
}

func (r Reader) Map(f func(Any) Any) Functor {
	return r.Chain(func(x Any) Monad {
		return Reader{}.Of(f(x)).(Monad)
	}).(Functor)
}

var (
	Reader_ = reader_{}
)

type reader_ struct{}

func (f reader_) As(x Any) Reader {
	return x.(Reader)
}

func (f reader_) Ref() Reader {
	return Reader{}
}

func (r reader_) Of(x Any) Point {
	return Reader{}.Of(x)
}

func (r reader_) Ask() Reader {
	return Reader{
		Run: func(a Any) Any {
			return a
		},
	}
}
