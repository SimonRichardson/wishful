package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type cofree struct {
	value   Any
	functor Functor
}

func Cofree(x Any, y Functor) cofree {
	return cofree{
		value:   x,
		functor: y,
	}
}

func (c cofree) Map(f func(Any) Any) Functor {
	return cofree{
		value: f(c.value),
		functor: c.functor.Map(func(a Any) Any {
			return a.(cofree).Map(f)
		}),
	}
}

func (c cofree) Extract() Any {
	return c.value
}

func (c cofree) Extend(f func(cofree) Any) cofree {
	return cofree{
		value: f(c),
		functor: c.functor.Map(func(a Any) Any {
			return a.(cofree).Extend(f)
		}),
	}
}

func (c cofree) Traverse(g func(Any) Functor) Functor {
	var do func(Any) Functor
	do = func(h Any) Functor {
		var (
			a = h.(cofree)
			b = a.functor.(Traversable).Traverse(do)
		)
		return g(a.value).Map(func(x Any) Any {
			return func(i Functor) cofree {
				return cofree{
					value:   x,
					functor: i,
				}
			}
		}).(Applicative).Ap(b.(Applicative)).(Functor)
	}
	return do(c)
}

var (
	Cofree_ = cofree_{}
)

type cofree_ struct{}

func (e cofree_) As(x Any) cofree {
	return x.(cofree)
}

func (e cofree_) Ref() cofree {
	return cofree{}
}
