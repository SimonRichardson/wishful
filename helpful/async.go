package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	EitherPromise Monad = NewEitherT(Promise_.Ref())
)

func Async(f func(x Any) Monad) func(x Any) Monad {
	return func(x Any) Monad {
		return EitherT_.As(EitherPromise).From(
			NewPromise(
				func(resolve func(x Any) Any) Any {
					fun := NewFunction(f)
					res, _ := fun.Call(x)
					return Promise_.As(res).Fork(func(x Any) Any {
						return resolve(NewRight(x))
					})
				},
			),
		)
	}
}
