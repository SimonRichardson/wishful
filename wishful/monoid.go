package wishful

type Monoid interface {
	Empty() Monoid
}

type MonoidLaws struct {
	x Point
}

func NewMonoidLaws(point Point) MonoidLaws {
	return MonoidLaws{
		x: point,
	}
}

func (o MonoidLaws) LeftIdentity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.(Monoid).Empty().(Semigroup)
		b := o.x.Of(v).(Semigroup)
		return run(a.Concat(b))
	}
	g := func(v int) AnyVal {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o MonoidLaws) RightIdentity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(v).(Semigroup)
		b := o.x.(Monoid).Empty().(Semigroup)
		return run(a.Concat(b))
	}
	g := func(v int) AnyVal {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o MonoidLaws) Associativity(run func(v AnyVal) AnyVal) (func(x Int, y Int, z Int) AnyVal, func(x Int, y Int, z Int) AnyVal) {
	f := func(x Int, y Int, z Int) AnyVal {
		a := o.x.Of(x).(Semigroup)
		b := o.x.Of(y).(Semigroup)
		c := o.x.Of(z).(Semigroup)

		return run(a.Concat(b).Concat(c))
	}
	g := func(x Int, y Int, z Int) AnyVal {
		a := o.x.Of(x).(Semigroup)
		b := o.x.Of(y).(Semigroup)
		c := o.x.Of(z).(Semigroup)

		return run(a.Concat(b.Concat(c)))
	}
	return f, g
}
