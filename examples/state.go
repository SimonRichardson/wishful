package main

import (
	"fmt"
	"time"

	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func main() {
	M := NewStateT(PureOption)

	program := M.Lift(a("Hello")).Chain(func(x AnyVal) Monad {
		p := x.(Promise).Map(func(x AnyVal) AnyVal {
			time.Sleep(time.Second)
			return fmt.Sprintf("%s %s", x, "World!")
		})
		return M.Modify(Constant(p))
	})

	state := program.(StateT).ExecState(a(""))
	state.(Option).Map(func(x AnyVal) AnyVal {
		x.(Promise).Fork(func(x AnyVal) AnyVal {
			fmt.Println(x)
			return x
		})
		return x
	})
}

func a(v string) Option {
	return NewSome(PurePromise.Of(v))
}
