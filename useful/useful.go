package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

func fromMonadToApplicativeAp(x Monad, y Applicative) Applicative {
	res := x.Chain(func(f AnyVal) Monad {
		fun := y.(Functor)
		res := fun.Map(func(g AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(g)
			return res
		})
		return res.(Monad)
	})
	return res.(Applicative)
}

func fromMonadToSemigroupConcat(x Monad, y Semigroup) Semigroup {
	res := x.Chain(func(a AnyVal) Monad {
		fun := y.(Functor)
		res := fun.Map(func(b AnyVal) AnyVal {
			sem0 := a.(Semigroup)
			sem1 := b.(Semigroup)
			return sem0.Concat(sem1)
		})
		return res.(Monad)
	})
	return res.(Semigroup)
}
