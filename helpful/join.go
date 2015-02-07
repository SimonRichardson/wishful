package helpful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

func Join(x Monad) Monad {
	return x.Chain(func(y Any) Monad {
		return y.(Monad)
	})
}
