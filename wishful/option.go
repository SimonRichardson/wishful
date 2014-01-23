package wishful

type Option interface {
	GetOrElse(x AnyVal) AnyVal
	OrElse(x Option) Option
}

type Some struct {
	x AnyVal
}

type None struct{}

func ToOption(x AnyVal) Option {
	if x == nil {
		return Some{x}
	}
	return None{}
}

func (x Some) Of(y AnyVal) Applicative {
	return Some{y}
}
func (x None) Of(y AnyVal) Applicative {
	return Some{y}
}

func (x Some) Empty() Monoid {
	return None{}
}
func (x None) Empty() Monoid {
	return None{}
}

// Methods
func (x Some) Chain(f func(y AnyVal) Monad) Monad {
	return f(x.x)
}
func (x None) Chain(f func(y AnyVal) Monad) Monad {
	return x
}

func (x Some) Concat(y Semigroup) Semigroup {
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
func (x None) Concat(y Semigroup) Semigroup {
	return x
}

func (x Some) Ap(y Applicative) Applicative {
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
func (x None) Ap(y Applicative) Applicative {
	return x
}

func (x Some) Map(f func(y AnyVal) AnyVal) Functor {
	res := x.Chain(func(y AnyVal) Monad {
		return Some{f(y)}
	})
	return res.(Functor)
}
func (x None) Map(f func(y AnyVal) AnyVal) Functor {
	return x
}

// Derived
func (x Some) OrElse(y Option) Option {
	return x
}
func (x None) OrElse(y Option) Option {
	return y
}

func (x Some) GetOrElse(y AnyVal) AnyVal {
	return x.x
}
func (x None) GetOrElse(y AnyVal) AnyVal {
	return y
}
