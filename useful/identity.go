package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

// Identity

func (x Id) Fold(f func(x AnyVal) AnyVal) AnyVal {
	return f(x.X)
}

// IdentityT

func (x IdT) Lift(m Applicative) IdT {
	return NewIdT(m)
}
