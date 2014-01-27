package wishful

type Monoid interface {
	Empty() Monoid
}

// Option

func (x Some) Empty() Monoid {
	return NewNone()
}

func (x None) Empty() Monoid {
	return NewNone()
}
