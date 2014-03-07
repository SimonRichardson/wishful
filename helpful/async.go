package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	EitherPromise EitherT = NewEitherT(Promise{})
)

func Async(f func(x AnyVal, y func(a AnyVal, b AnyVal) AnyVal) AnyVal) func(x AnyVal) EitherT {
	return func(x AnyVal) EitherT {
		return EitherPromise.From(
			NewPromise(
				func(resolve func(x AnyVal) AnyVal) AnyVal {
					fun := NewFunction(f)
					res, _ := fun.Call(x, func(l AnyVal, r AnyVal) AnyVal {
						if l != nil {
							return resolve(NewLeft(l))
						} else {
							return resolve(NewRight(r))
						}
					})
					return res
				},
			),
		)
	}
}
