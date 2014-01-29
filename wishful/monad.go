package wishful

type Monad interface {
	Chain(f func(v AnyVal) Monad) Monad
}

type MonadLaws struct {
	x Point
}

func NewMonadLaws(point Point) MonadLaws {
	return MonadLaws{
		x: point,
	}
}

func (o MonadLaws) LeftIdentity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x AnyVal) Monad {
			return Apply(func(x AnyVal) AnyVal {
				return o.x.Of(x)
			})(x).(Monad)
		}))
	}
	g := func(v int) AnyVal {
		return run(Apply(func(x AnyVal) AnyVal {
			return o.x.Of(x)
		})(v))
	}
	return f, g
}

func (o MonadLaws) RightIdentity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x AnyVal) Monad {
			return o.x.Of(x).(Monad)
		}))
	}
	g := func(v int) AnyVal {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o MonadLaws) Associativity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x AnyVal) Monad {
			return o.x.Of(x).(Monad)
		}).Chain(func(x AnyVal) Monad {
			return o.x.Of(x).(Monad)
		}))
	}
	g := func(v int) AnyVal {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x AnyVal) Monad {
			b := o.x.Of(x).(Monad)
			return b.Chain(func(x AnyVal) Monad {
				return o.x.Of(x).(Monad)
			})
		}))
	}
	return f, g
}
