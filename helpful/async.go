package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	EitherPromise EitherT = NewEitherT(Promise{})
)

func Async(f func(x AnyVal) Promise) func(x AnyVal) EitherT {
	return func(x AnyVal) EitherT {
		return EitherPromise.From(
			NewPromise(
				func(resolve func(x AnyVal) AnyVal) AnyVal {
					fun := NewFunction(f)
					res, _ := fun.Call(x)
					return res.(Promise).Fork(func(x AnyVal) AnyVal {
						return resolve(NewRight(x))
					})
				},
			),
		)
	}
}
