package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	concat = fromMonadToSemigroupConcat(func(a Semigroup, b Semigroup) Any {
		return Append(a, b)
	})
)

func fromMonadToApplicativeAp(x Monad, y Applicative) Applicative {
	res := x.Chain(func(f Any) Monad {
		fun := y.(Functor)
		res := fun.Map(func(g Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(g)
			return res
		})
		return res.(Monad)
	})
	return res.(Applicative)
}

func fromMonadToSemigroupConcat(f func(a Semigroup, b Semigroup) Any) func(x Monad, y Semigroup) Semigroup {
	return func(x Monad, y Semigroup) Semigroup {
		res := x.Chain(func(a Any) Monad {
			fun := y.(Functor)
			res := fun.Map(func(b Any) Any {
				sem0 := a.(Semigroup)
				sem1 := b.(Semigroup)
				return f(sem0, sem1)
			})
			return res.(Monad)
		})
		return res.(Semigroup)
	}
}

func concatAnyvals(x Any) func(y Any) Any {
	return func(y Any) Any {
		a := x.(Semigroup)
		b := y.(Semigroup)
		return Append(a, b)
	}
}
