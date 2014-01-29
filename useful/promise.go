package useful

type Promise struct {
	Fork func(resolve func(x AnyVal) AnyVal) AnyVal
}

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

func (x Promise) Chain(f func(v AnyVal) Monad) Monad {
	return Promise{func(resolve func(x AnyVal) AnyVal) AnyVal {
		return x.Fork(func(a AnyVal) AnyVal {
			p := f(a).(Promise)
			return p.Fork(resolve)
		})
	}}
}

func (x Promise) Map(f func(v AnyVal) AnyVal) Functor {
	return Promise{func(resolve func(v AnyVal) AnyVal) AnyVal {
		return x.Fork(func(a AnyVal) AnyVal {
			return resolve(f(a))
		})
	}}
}
