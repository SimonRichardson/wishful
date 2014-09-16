package wishful

type Applicative interface {
	Ap(v Applicative) Applicative
}

type ApplicativeLaws struct {
	x Point
}

func NewApplicativeLaws(point Point) ApplicativeLaws {
	return ApplicativeLaws{
		x: point,
	}
}

func (o ApplicativeLaws) Identity(run func(v Any) Any) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(Identity).(Applicative)
		b := o.x.Of(v).(Applicative)
		return run(a.Ap(b))
	}
	g := func(v int) Any {
		return run(o.x.Of(v))
	}
	return f, g
}

func (o ApplicativeLaws) Composition(run func(v Any) Any) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(Compose).(Applicative)
		b := o.x.Of(Identity).(Applicative)
		c := o.x.Of(Identity).(Applicative)
		d := o.x.Of(v).(Applicative)
		return run(a.Ap(b).Ap(c).Ap(d))
	}
	g := func(v int) Any {
		a := o.x.Of(Identity).(Applicative)
		b := o.x.Of(Identity).(Applicative)
		c := o.x.Of(v).(Applicative)
		return run(a.Ap(b.Ap(c)))
	}
	return f, g
}

func (o ApplicativeLaws) Homomorphism(run func(v Any) Any) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(Identity).(Applicative)
		b := o.x.Of(v).(Applicative)
		return run(a.Ap(b).(Applicative))
	}
	g := func(v int) Any {
		return run(o.x.Of(Identity(v)))
	}
	return f, g
}

func (o ApplicativeLaws) Interchange(run func(v Any) Any) (func(v int) Any, func(v int) Any) {
	f := func(v int) Any {
		a := o.x.Of(Identity).(Applicative)
		b := o.x.Of(v).(Applicative)
		return run(a.Ap(b).(Applicative))
	}
	g := func(v int) Any {
		a := o.x.Of(Thrush(v)).(Applicative)
		b := o.x.Of(Identity).(Applicative)
		return run(a.Ap(b))
	}
	return f, g
}
