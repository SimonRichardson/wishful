package wishful

type Monoid interface {
	Empty() Monoid
}

// Option

func (x Some) Empty() Monoid {
	return None{}
}

func (x None) Empty() Monoid {
	return None{}
}
