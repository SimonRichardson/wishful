package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Free interface {
	Of(Any) Point
	Chain(f func(Any) Monad) Monad
	Map(f func(Any) Any) Functor

	Run() Any
}
