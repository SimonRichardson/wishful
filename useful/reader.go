package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Reader struct {
	Run func(Any) Any
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
		return Reader_.Of(f(x)).(Monad)
	}).(Functor)
}

var (
	Reader_ = reader{}
)

type reader struct{}

func (r reader) Of(x Any) Point {
	return Reader{}.Of(x)
}

func (r reader) Ask() Reader {
	return Reader{
		Run: func(a Any) Any {
			return a
		},
	}
}
