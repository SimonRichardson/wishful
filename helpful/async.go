package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	EitherPromise EitherT = NewEitherT(Promise{})
)

func Async(f func(x Any) Promise) func(x Any) EitherT {
	return func(x Any) EitherT {
		return EitherPromise.From(
			NewPromise(
				func(resolve func(x Any) Any) Any {
					fun := NewFunction(f)
					res, _ := fun.Call(x)
					return res.(Promise).Fork(func(x Any) Any {
						return resolve(NewRight(x))
					})
				},
			),
		)
	}
}
