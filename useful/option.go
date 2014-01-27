package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

// Option

func (x Some) Fold(f func(x AnyVal) AnyVal, g func() AnyVal) AnyVal {
	return f(x.X)
}

func (x None) Fold(f func(x AnyVal) AnyVal, g func() AnyVal) AnyVal {
	return g()
}
