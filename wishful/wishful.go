package wishful

// Identity

type Id struct {
	x AnyVal
}

func NewId(x AnyVal) Id {
	return Id{
		x: x,
	}
}

func (x Id) Fold(f func(x AnyVal) AnyVal) AnyVal {
	return f(x.x)
}

// IdentityT

type IdT struct {
	m   Applicative
	Run AnyVal
}

func NewIdT(m Applicative) IdT {
	return IdT{
		m:   m,
		Run: Empty{},
	}
}

func (x IdT) Lift(m Applicative) IdT {
	return NewIdT(m)
}

// IO

type IO struct {
	UnsafePerform func() AnyVal
}

func NewIO(unsafe func() AnyVal) IO {
	return IO{
		UnsafePerform: unsafe,
	}
}

// Option

type Option interface {
}

type Some struct {
	x AnyVal
}

type None struct {
}

func NewSome(x AnyVal) Some {
	return Some{
		x: x,
	}
}

func NewNone() None {
	return None{}
}

func (x Some) Fold(f func(x AnyVal) AnyVal, g func() AnyVal) AnyVal {
	return f(x.x)
}

func (x None) Fold(f func(x AnyVal) AnyVal, g func() AnyVal) AnyVal {
	return g()
}

// Promise

type Promise struct {
	Fork func(resolve func(x AnyVal) AnyVal) AnyVal
}
