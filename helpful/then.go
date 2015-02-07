package helpful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

func Then(x Monad, y Monad) Monad {
	return x.Chain(func(x Any) Monad {
		return y
	})
}
