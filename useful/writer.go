package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Writer struct {
	Run func() Tuple2
}

func NewWriter(x Any, y []Any) Writer {
	return Writer{
		Run: func() Tuple2 {
			return NewTuple2(x, y)
		},
	}
}

func (w Writer) Of(x Any) Point {
	return Writer{
		Run: func() Tuple2 {
			return NewTuple2(x, []Any{})
		},
	}
}

func (w Writer) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(w, v)
}

func (w Writer) Chain(f func(Any) Monad) Monad {
	return Writer{
		Run: func() Tuple2 {
			var (
				exe0 = w.Run()
				a    = exe0.Fst()
				b    = exe0.Snd().([]Any)

				exe1 = f(a).(Writer).Run()
				x    = exe1.Fst()
				y    = exe1.Snd().([]Any)
			)
			return NewTuple2(x, append(b, y...))
		},
	}
}

func (w Writer) Map(f func(Any) Any) Functor {
	return w.Chain(func(x Any) Monad {
		return Writer{
			Run: func() Tuple2 {
				return NewTuple2(f(x), []Any{})
			},
		}
	}).(Functor)
}

func (w Writer) Tell(x Any) Writer {
	return Writer{
		Run: func() Tuple2 {
			b := w.Run().Snd().([]Any)
			return NewTuple2(Empty{}, append(b, x))
		},
	}
}

// Static methods

var (
	Writer_ = writer{}
)

type writer struct{}

func (w writer) Of(x Any) Point {
	return Writer{}.Of(x)
}
