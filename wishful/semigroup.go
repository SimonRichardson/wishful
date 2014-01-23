package wishful

type Semigroup interface {
	Concat(x Semigroup) Semigroup
}

// Identity

func (x Id) Concat(y Semigroup) Semigroup {
	return fromMonadToSemigroupConcat(x, y)
}

// Option

func (x Some) Concat(y Semigroup) Semigroup {
	return fromMonadToSemigroupConcat(x, y)
}

func (x None) Concat(y Semigroup) Semigroup {
	return x
}

// Common

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
