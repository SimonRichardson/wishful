package wishful

type Applicative interface {
	Of(v AnyVal) Applicative
	Ap(v Applicative) Applicative
}

// Identity

func (x Id) Of(v AnyVal) Applicative {
	return Id{v}
}

func (x Id) Ap(v Applicative) Applicative {
	return fromMonadToApplicativeAp(x, v)
}

// Option

func (x Some) Of(v AnyVal) Applicative {
	return Some{v}
}

func (x None) Of(v AnyVal) Applicative {
	return Some{v}
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
