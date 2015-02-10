package useful

import . "github.com/SimonRichardson/wishful/wishful"

type WriterT struct {
	m   Point
	Run func() Tuple2
}

func NewWriterT(m Point) WriterT {
	return WriterT{
		m: m,
		Run: func() Tuple2 {
			return NewTuple2(Empty{}, []Any{})
		},
	}
}

func (w WriterT) Of(a Any) Point {
	return WriterT{
		m: w.m,
		Run: func() Tuple2 {
			return NewTuple2(w.m.Of(a), []Any{})
		},
	}
}

func (w WriterT) Chain(f func(Any) Monad) Monad {
	return WriterT{
		m: w.m,
		Run: func() Tuple2 {
			var (
				a = w.Run()
				b = a.Fst().(Monad)
				c = b.Chain(f).(WriterT).Run()
			)
			return NewTuple2(c.Fst(), a.Snd())
		},
	}
}

func (w WriterT) Map(f Morphism) Functor {
	return w.Chain(func(a Any) Monad {
		return w.Of(f(a)).(Monad)
	}).(Functor)
}

func (w WriterT) Tell(x Any) WriterT {
	return WriterT{
		m: w.m,
		Run: func() Tuple2 {
			var (
				a = w.Run()
				b = a.Snd().([]Any)
			)
			return NewTuple2(a.Fst(), append(b, x))
		},
	}
}

var (
	WriterT_ = writerT_{}
)

type writerT_ struct{}

func (f writerT_) As(x Any) WriterT {
	return x.(WriterT)
}

func (f writerT_) Ref() WriterT {
	return WriterT{}
}
