package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type StateT struct {
	m   Point
	Run func(x Any) Point
}

func NewStateT(m Point) StateT {
	return StateT{
		m: m,
		Run: func(x Any) Point {
			return nil
		},
	}
}

func (x StateT) Lift(m Functor) StateT {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			return m.Map(func(c Any) Any {
				return Tuple2{_1: c, _2: b}
			}).(Point)
		},
	}
}

func (x StateT) Of(a Any) Point {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			return x.m.Of(Tuple2{_1: a, _2: b})
		},
	}
}

func (x StateT) Chain(f func(a Any) Monad) Monad {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			result := x.Run(b)
			return result.(Monad).Chain(func(t Any) Monad {
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
		Run: func(b Any) Point {
			return x.m.Of(Tuple2{_1: b, _2: b})
		},
	}
}

func (x StateT) Modify(f func(b Any) Any) StateT {
	return StateT{
		m: x.m,
		Run: func(b Any) Point {
			fun := NewFunction(f)
			res, _ := fun.Call(b)
			return x.m.Of(Tuple2{_1: Empty{}, _2: res})
		},
	}
}

func (x StateT) Put(v Any) StateT {
	return x.Modify(func(a Any) Any {
		return v
	})
}

func (x StateT) EvalState(s Any) Any {
	return x.Run(s).(Functor).Map(func(t Any) Any {
		return t.(Tuple2)._1
	})
}

func (x StateT) ExecState(s Any) Any {
	return x.Run(s).(Functor).Map(func(t Any) Any {
		return t.(Tuple2)._2
	})
}

func (x StateT) Map(f func(x Any) Any) Functor {
	return x.Chain(func(a Any) Monad {
		fun := NewFunction(f)
		res, _ := fun.Call(a)
		return x.Of(res).(Monad)
	}).(Functor)
}

func (x StateT) Ap(a Applicative) Applicative {
	return x.Chain(func(f Any) Monad {
		return a.(Functor).Map(func(b Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(b)
			return res
		}).(Monad)
	}).(Applicative)
}
