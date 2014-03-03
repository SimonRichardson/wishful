package wishful

type Comonad interface {
	Extend(f func(a Comonad) AnyVal) Comonad
	Extract() AnyVal
}

type ComonadLaws struct {
	x Point
}

func NewComonadLaws(point Point) ComonadLaws {
	return ComonadLaws{
		x: point,
	}
}

func (o ComonadLaws) Identity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Comonad)
		return run(a.Extend(func(x Comonad) AnyVal {
			return a.Extract()
		}))
	}
	g := func(v int) AnyVal {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o ComonadLaws) Composition(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	extract := func(y Comonad) AnyVal {
		return y.Extract()
	}
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Comonad)
		b := Compose(func(x AnyVal) AnyVal {
			return a.Extract()
		})(func(x AnyVal) AnyVal {
			return a.Extend(extract)
		})

		return run(b(a))
	}
	g := func(v int) AnyVal {
		a := o.x.Of(v).(Comonad)
		return run(extract(a))
	}
	return f, g
}

func (o ComonadLaws) Associativity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	extract := func(y Comonad) AnyVal {
		return y.Extract()
	}
	duplicate := func(x Comonad) func(y AnyVal) AnyVal {
		return func(y AnyVal) AnyVal {
			return x.Extend(extract)
		}
	}
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Comonad)
		b := Compose(duplicate(a))(duplicate(a))
		return run(b(a))
	}
	g := func(v int) AnyVal {
		a := o.x.Of(v).(Functor)
		c := a.(Comonad)
		return run(a.Map(Compose(duplicate(c))(duplicate(c))).(Comonad).Extract())
	}
	return f, g
}
