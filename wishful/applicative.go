package wishful

type Applicative interface {
	Of(v AnyVal) Applicative
	Ap(v Applicative) Applicative
}

func (o Laws) Identity() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(v)
	}
	return f, g
}

func (o Laws) Composition() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Compose).Ap(o.Of(Identity)).Ap(o.Of(Identity)).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(Identity).Ap(o.Of(v)))
	}
	return f, g
}

func (o Laws) Homomorphism() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(Identity(v))
	}
	return f, g
}

func (o Laws) Interchange() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(Thrush(v)).Ap(o.Of(Identity))
	}
	return f, g
}
