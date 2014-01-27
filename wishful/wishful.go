package wishful

// Identity

type Id struct {
	X AnyVal
}

func NewId(x AnyVal) Id {
	return Id{
		X: x,
	}
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
	X AnyVal
}
type None struct {
}

func NewSome(x AnyVal) Some {
	return Some{
		X: x,
	}
}
func NewNone() None {
	return None{}
}

// Promise

type Promise struct {
	Fork func(resolve func(x AnyVal) AnyVal) AnyVal
}
