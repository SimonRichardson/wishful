package wishful

type Functor interface {
	Map(f func(v AnyVal) AnyVal) Functor
}

type FunctorLaws struct {
	x Point
}

func NewFunctorLaws(point Point) FunctorLaws {
	return FunctorLaws{
		x: point,
	}
}

func (o FunctorLaws) Identity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Identity))
	}
	g := func(v int) AnyVal {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o FunctorLaws) Composition(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Compose(Identity)(Identity)))
	}
	g := func(v int) AnyVal {
		a := o.x.Of(v).(Functor)
		return run(a.Map(Identity).Map(Identity))
	}
	return f, g
}
