package wishful

type Applicative interface {
	Of(v AnyVal) Applicative
	Ap(v Applicative) Applicative
}

type Applicative struct{}

func (o Applicative) Identity() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(v)
	}
	return f, g
}

func (o Applicative) Composition() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Compose).Ap(o.Of(Identity)).Ap(o.Of(Identity)).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(Identity).Ap(o.Of(v)))
	}
	return f, g
}

func (o Applicative) Homomorphism() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(Identity(v))
	}
	return f, g
}

func (o Applicative) Interchange() {
	f := func(v AnyVal) AnyVal {
		return o.Of(Identity).Ap(o.Of(v))
	}
	g := func(v AnyVal) AnyVal {
		return o.Of(Thrush(v)).Ap(o.Of(Identity))
	}
	return f, g
}
