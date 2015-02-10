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

func (o MonoidLaws) LeftIdentity(run Morphism) (func(int) Any, func(int) Any) {
	f := func(v int) Any {
		a := o.x.(Monoid).Empty().(Semigroup)
		b := o.x.Of(v).(Semigroup)
		return run(a.Concat(b))
	}
	g := func(v int) Any {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o MonoidLaws) RightIdentity(run Morphism) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(v).(Semigroup)
		b := o.x.(Monoid).Empty().(Semigroup)
		return run(a.Concat(b))
	}
	g := func(v int) Any {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o MonoidLaws) Associativity(run Morphism) (func(Int, Int, Int) Any, func(Int, Int, Int) Any) {
	f := func(x Int, y Int, z Int) Any {
		a := o.x.Of(x).(Semigroup)
		b := o.x.Of(y).(Semigroup)
		c := o.x.Of(z).(Semigroup)

		return run(a.Concat(b).Concat(c))
	}
	g := func(x Int, y Int, z Int) Any {
		a := o.x.Of(x).(Semigroup)
		b := o.x.Of(y).(Semigroup)
		c := o.x.Of(z).(Semigroup)

		return run(a.Concat(b.Concat(c)))
	}
	return f, g
}
