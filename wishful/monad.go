package wishful

type Monad interface {
	Chain(f func(v AnyVal) Monad) Monad
}

// Identity

func (x Id) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

// Option

func (x Some) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x None) Chain(f func(v AnyVal) Monad) Monad {
	return x
}
