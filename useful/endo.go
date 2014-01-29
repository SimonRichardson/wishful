package useful

import (
	. "github.com/SimonRichardson/wishful"
)

type Endo struct {
	x func(v AnyVal) AnyVal
}

func NewEndo(x func(v AnyVal) AnyVal) Endo {
	return Endo{x}
}

func (o Endo) Of(x AnyVal) Applicative {
	return NewEndo(func(v AnyVal) AnyVal {
		return x
	})
}

func (o Endo) Empty() Monoid {
	return NewEndo(Identity)
}

func (o Endo) Chain(f func(v AnyVal) Monad) Monad {
	return NewEndo(func(v AnyVal) AnyVal {
		return f(o.x(v)).x
	})
}

func (o Endo) Concat(x Semigroup) Semigroup {
	return o.Chain(func(a AnyVal) AnyVal {
		return x.Map(func(b AnyVal) AnyVal {
			return a.(int) * b.(int)
		})
	})
}

func (o Endo) Map(f func(v AnyVal) AnyVal) Functor {
	return o.Chain(func(v AnyVal) AnyVal {
		return NewEndo(f(v))
	})
}
