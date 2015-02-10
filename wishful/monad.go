package wishful

type Monad interface {
	Chain(f func(Any) Monad) Monad
}

type MonadLaws struct {
	x Point
}

func NewMonadLaws(point Point) MonadLaws {
	return MonadLaws{
		x: point,
	}
}

func (o MonadLaws) LeftIdentity(run Morphism) (func(int) Any, func(int) Any) {
	f := func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			return Apply(func(x Any) Any {
				return o.x.Of(x)
			})(x).(Monad)
		}))
	}
	g := func(v int) Any {
		return run(Apply(func(x Any) Any {
			return o.x.Of(x)
		})(v))
	}
	return f, g
}

func (o MonadLaws) RightIdentity(run Morphism) (func(int) Any, func(int) Any) {
	f := func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			return o.x.Of(x).(Monad)
		}))
	}
	g := func(v int) Any {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o MonadLaws) Associativity(run Morphism) (func(int) Any, func(int) Any) {
	f := func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			return o.x.Of(x).(Monad)
		}).Chain(func(x Any) Monad {
			return o.x.Of(x).(Monad)
		}))
	}
	g := func(v int) Any {
		a := o.x.Of(v).(Monad)
		return run(a.Chain(func(x Any) Monad {
			b := o.x.Of(x).(Monad)
			return b.Chain(func(x Any) Monad {
				return o.x.Of(x).(Monad)
			})
		}))
	}
	return f, g
}
