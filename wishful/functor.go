package wishful

type Functor interface {
	Map(f func(v AnyVal) AnyVal) Functor
}

// Identity

func (x Id) Map(f func(v AnyVal) AnyVal) Functor {
	return NewId(f(x.X))
}

// IdentityT

func (x IdT) Map(f func(v AnyVal) AnyVal) Functor {
	mon := x.Chain(func(y AnyVal) Monad {
		app := NewIdT(x.m).Of(f(y))
		return app.(Monad)
	})
	return mon.(Functor)
}

// IO

func (x IO) Map(f func(x AnyVal) AnyVal) Functor {
	res := x.Chain(func(x AnyVal) Monad {
		return IO{func() AnyVal {
			return f(x)
		}}
	})
	return res.(Functor)
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
