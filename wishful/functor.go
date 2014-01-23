package wishful

type Functor interface {
	Map(f func(v AnyVal) AnyVal) Functor
}

// Identity

func (x Id) Map(f func(v AnyVal) AnyVal) Functor {
	return Id{f(x.x)}
}

// Option

func (x Some) Map(f func(v AnyVal) AnyVal) Functor {
	res := x.Chain(func(v AnyVal) Monad {
		return Some{f(v)}
	})
	return res.(Functor)
}

func (x None) Map(f func(v AnyVal) AnyVal) Functor {
	return x
}
