package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

func Append(a Semigroup, b Semigroup) Semigroup {
	return a.Concat(b)
}

func LiftA2(f func(a Any, b Any) Any, a Applicative, b Applicative) Applicative {
	x := a.(Functor)
	y := x.Map(func(a Any) Any {
		return func(b Any) Any {
			return f(a, b)
		}
	})
	return b.Ap(y.(Applicative))
}

func LiftA3(f func(a Any, b Any, c Any) Any, a Applicative, b Applicative, c Applicative) Applicative {
	x := a.(Functor)
	y := x.Map(func(a Any) Any {
		return func(b Any) Any {
			return func(b Any) Any {
				return f(a, b, c)
			}
		}
	})
	return b.Ap(y.(Applicative))
}
