package wishful

type Monad interface {
	Chain(f func(v AnyVal) Monad) Monad
}

// Identity

func (x Id) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

// IO

func (x IO) Chain(f func(x AnyVal) Monad) Monad {
	return NewIO(func() AnyVal {
		io := f(x.UnsafePerform()).(IO)
		return io.UnsafePerform()
	})
}

// Option

func (x Some) Chain(f func(v AnyVal) Monad) Monad {
	return f(x.x)
}

func (x None) Chain(f func(v AnyVal) Monad) Monad {
	return x
}

// Promises

func (x Promise) Chain(f func(v AnyVal) Monad) Monad {
	return Promise{func(resolve func(x AnyVal) AnyVal) AnyVal {
		return x.Fork(func(a AnyVal) AnyVal {
			p := f(a).(Promise)
			return p.Fork(resolve)
		})
	}}
}
