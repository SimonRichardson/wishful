package wishful

type Applicative interface {
	Of(v AnyVal) Applicative
	Ap(v Applicative) Applicative
}

// Identity

func (x Id) Of(v AnyVal) Applicative {
	return NewId(v)
}

func (x Id) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

// IdentityT

func (x IdT) Of(v AnyVal) Applicative {
	return IdT{
		m:   x.m,
		Run: x.m.Of(v),
	}
}

func (x IdT) Ap(v Applicative) Applicative {
	mon := x.Chain(func(f AnyVal) Monad {
		fun := f.(func(f AnyVal) AnyVal)
		return v.(Functor).Map(fun).(Monad)
	})
	return mon.(Applicative)
}

// IO

func (x IO) Of(v AnyVal) Applicative {
	return NewIO(func() AnyVal {
		return v
	})
}

func (x IO) Ap(v Applicative) Applicative {
	res := x.Chain(func(f AnyVal) Monad {
		fun := v.(Functor)
		res := fun.Map(f.(func(f AnyVal) AnyVal))
		return res.(Monad)
	})
	return res.(Applicative)
}

// Option

func (x Some) Of(v AnyVal) Applicative {
	return NewSome(v)
}

func (x None) Of(v AnyVal) Applicative {
	return NewSome(v)
}

func (x Some) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

func (x None) Ap(v Applicative) Applicative {
	return x
}

// Promise

func (x Promise) Of(v AnyVal) Applicative {
	return Promise{func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(v)
	}}
}

func (x Promise) Ap(v Applicative) Applicative {
	return Promise{func(resolve func(x AnyVal) AnyVal) AnyVal {
		return x.Fork(func(f AnyVal) AnyVal {
			fun := v.(Functor)
			pro := fun.Map(f.(func(x AnyVal) AnyVal)).(Promise)
			return pro.Fork(resolve)
		})
	}}
}

// Common

func fromMonadToApplicativeAp(x Monad, y Applicative) Applicative {
	res := x.Chain(func(f AnyVal) Monad {
		fun := y.(Functor)
		res := fun.Map(func(g AnyVal) AnyVal {
			app := f.(func(AnyVal) AnyVal)
			return app(g)
		})
		return res.(Monad)
	})
	return res.(Applicative)
}
