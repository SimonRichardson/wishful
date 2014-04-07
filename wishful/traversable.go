package wishful

type Traversable interface {
	Traverse(f func(v AnyVal) AnyVal) Traversable
}

type TraversableLaws struct {
	x Point
}

func NewTraversableLaws(point Point) TraversableLaws {
	return TraversableLaws{
		x: point,
	}
}

func (o TraversableLaws) Identity(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(id{}.Of(v)).(Traversable)
		return run(a.Traverse(Identity))
	}
	g := func(v int) AnyVal {
		return run(id{}.Of(o.x.Of(v)))
	}
	return f, g
}

func (o TraversableLaws) Composition(run func(v AnyVal) AnyVal) (func(v int) AnyVal, func(v int) AnyVal) {
	f := func(v int) AnyVal {
		a := o.x.Of(id{}.Of(v)).(Traversable)
		return run(a.Traverse(Compose(Identity)(Identity)))
	}
	g := func(v int) AnyVal {
		a := o.x.Of(id{}.Of(v)).(Traversable)
		return run(a.Traverse(Identity).Traverse(Identity))
	}
	return f, g
}

// Used for testing, as we can't import useful package because import cycle error.

type id struct {
	x AnyVal
}

func (x id) Of(y AnyVal) Point {
	return id{x: y}
}

func (x id) Map(f func(x AnyVal) AnyVal) Functor {
	return id{}.Of(f(x.x)).(Functor)
}

func (x id) Traverse(f func(x AnyVal) AnyVal) Traversable {
	return f(x.x).(Functor).Map(func(x AnyVal) AnyVal {
		return id{}.Of(x)
	}).(Traversable)
}
