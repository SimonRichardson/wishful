package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type StateT struct {
	m   Point
	Run func(x AnyVal) Point
}

func NewStateT(m Point) StateT {
	return StateT{
		m: m,
		Run: func(x AnyVal) Point {
			return nil
		},
	}
}

func (x StateT) Lift(m Functor) StateT {
	return StateT{
		m: x.m,
		Run: func(b AnyVal) Point {
			return m.Map(func(c AnyVal) AnyVal {
				return Tuple2{_1: c, _2: b}
			}).(Point)
		},
	}
}

func (x StateT) Of(a AnyVal) Point {
	return StateT{
		m: x.m,
		Run: func(b AnyVal) Point {
			return x.m.Of(Tuple2{_1: a, _2: b})
		},
	}
}

func (x StateT) Chain(f func(a AnyVal) Monad) Monad {
	return StateT{
		m: x.m,
		Run: func(b AnyVal) Point {
			result := x.Run(b)
			return result.(Monad).Chain(func(t AnyVal) Monad {
				tup := t.(Tuple2)
				fun := NewFunction(f)
				res, _ := fun.Call(tup._1)
				return res.(StateT).Run(tup._2).(Monad)
			}).(Point)
		},
	}
}

func (x StateT) Get() StateT {
	return StateT{
		m: x.m,
		Run: func(b AnyVal) Point {
			return x.m.Of(Tuple2{_1: b, _2: b})
		},
	}
}

func (x StateT) Modify(f func(b AnyVal) AnyVal) StateT {
	return StateT{
		m: x.m,
		Run: func(b AnyVal) Point {
			fun := NewFunction(f)
			res, _ := fun.Call(b)
			return x.m.Of(Tuple2{_1: Empty{}, _2: res})
		},
	}
}

func (x StateT) Put(v AnyVal) StateT {
	return x.Modify(func(a AnyVal) AnyVal {
		return v
	})
}

func (x StateT) EvalState(s AnyVal) AnyVal {
	return x.Run(s).(Functor).Map(func(t AnyVal) AnyVal {
		return t.(Tuple2)._1
	})
}

func (x StateT) ExecState(s AnyVal) AnyVal {
	return x.Run(s).(Functor).Map(func(t AnyVal) AnyVal {
		return t.(Tuple2)._2
	})
}

func (x StateT) Map(f func(x AnyVal) AnyVal) Functor {
	return x.Chain(func(a AnyVal) Monad {
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return x.Of(res).(Monad)
	}).(Functor)
}

func (x StateT) Ap(a Applicative) Applicative {
	return x.Chain(func(f AnyVal) Monad {
		return a.(Functor).Map(func(b AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(b)
			return res
		}).(Monad)
	}).(Applicative)
}
