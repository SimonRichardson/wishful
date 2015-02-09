package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Done(x Any) Free {
	return Return(x)
}

func Cont(f func() Free) Free {
	return Suspend(LiftFunc(f))
}

func Trampoline(x Free) Any {
	return x.Run()
}
