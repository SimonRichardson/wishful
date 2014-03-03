package helpful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

func Append(a Semigroup, b Semigroup) Semigroup {
	return a.Concat(b)
}

func Join(a Monad) Monad {
	return a.Chain(func(a AnyVal) Monad {
		return a.(Monad)
	})
}

func LiftA2(f func(a AnyVal, b AnyVal) AnyVal, a Applicative, b Applicative) Applicative {
	x := a.(Functor)
	y := x.Map(func(a AnyVal) AnyVal {
		return func(b AnyVal) AnyVal {
			return f(a, b)
		}
	})
	return b.Ap(y.(Applicative))
}

func LiftA3(f func(a AnyVal, b AnyVal, c AnyVal) AnyVal, a Applicative, b Applicative, c Applicative) Applicative {
	x := a.(Functor)
	y := x.Map(func(a AnyVal) AnyVal {
		return func(b AnyVal) AnyVal {
			return func(b AnyVal) AnyVal {
				return f(a, b, c)
			}
		}
	})
	return b.Ap(y.(Applicative))
}
