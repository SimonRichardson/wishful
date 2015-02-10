package wishful

type Semigroup interface {
	Concat(x Semigroup) Semigroup
}

type SemigroupLaws struct {
	x Point
}

func NewSemigroupLaws(point Point) SemigroupLaws {
	return SemigroupLaws{
		x: point,
	}
}

func (o SemigroupLaws) Associativity(run Morphism) (func(Int, Int, Int) Any, func(Int, Int, Int) Any) {
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
