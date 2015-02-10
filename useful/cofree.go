package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Cofree struct {
	value   Any
	functor Functor
}

func NewCofree(x Any, y Functor) Cofree {
	return Cofree{
		value:   x,
		functor: y,
	}
}

func (c Cofree) Map(f func(Any) Any) Functor {
	return Cofree{
		value: f(c.value),
		functor: c.functor.Map(func(a Any) Any {
			return a.(Cofree).Map(f)
		}),
	}
}

func (c Cofree) Extract() Any {
	return c.value
}

func (c Cofree) Extend(f func(Cofree) Any) Cofree {
	return Cofree{
		value: f(c),
		functor: c.functor.Map(func(a Any) Any {
			return a.(Cofree).Extend(f)
		}),
	}
}

func (c Cofree) Traverse(g func(Any) Functor) Functor {
	var do func(Any) Functor
	do = func(h Any) Functor {
		var (
			a = h.(Cofree)
			b = a.functor.(Traversable).Traverse(do)
		)
		return g(a.value).Map(func(x Any) Any {
			return func(i Functor) Cofree {
				return Cofree{
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

func (e cofree_) As(x Any) Cofree {
	return x.(Cofree)
}

func (e cofree_) Ref() Cofree {
	return Cofree{}
}
