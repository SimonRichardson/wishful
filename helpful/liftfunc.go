package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func LiftFunc(x func() Free) Functor {
	return funcF{
		val: func() Any {
			return x()
		},
	}
}

type funcF struct {
	val func() Any
}

func (x funcF) Map(f Morphism) Functor {
	return funcF{
		val: func() Any {
			return f(x.val())
		},
	}
}

func (x funcF) Run() Free {
	return x.val().(Free)
}
