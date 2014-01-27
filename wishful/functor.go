package wishful

type Functor interface {
	Map(f func(v AnyVal) AnyVal) Functor
}

// Identity

func (x Id) Map(f func(v AnyVal) AnyVal) Functor {
	return NewId(f(x.x))
}

// Option

func (x Some) Map(f func(v AnyVal) AnyVal) Functor {
	res := x.Chain(func(v AnyVal) Monad {
		return NewSome(f(v))
	})
	return res.(Functor)
}

func (x None) Map(f func(v AnyVal) AnyVal) Functor {
	return x
}

// Promise

func (x Promise) Map(f func(v AnyVal) AnyVal) Functor {
	return Promise{func(resolve func(v AnyVal) AnyVal) AnyVal {
		return x.Fork(func(a AnyVal) AnyVal {
			return resolve(f(a))
		})
	}}
}
