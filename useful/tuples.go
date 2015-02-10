package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Tuple2 struct {
	_1 Any
	_2 Any
}

func NewTuple2(a Any, b Any) Tuple2 {
	return Tuple2{_1: a, _2: b}
}

func (t Tuple2) Fst() Any {
	return t._1
}

func (t Tuple2) Snd() Any {
	return t._2
}

func (t Tuple2) MapFst(f Morphism) Tuple2 {
	return NewTuple2(f(t._1), t._2)
}

func (t Tuple2) MapSnd(f Morphism) Tuple2 {
	return NewTuple2(t._1, f(t._2))
}

func (t Tuple2) Bimap(f Morphism, g Morphism) Tuple2 {
	return NewTuple2(f(t._1), g(t._2))
}

func (t Tuple2) Append(x Any) Tuple3 {
	return NewTuple3(t._1, t._2, x)
}

func (t Tuple2) Slice() []Any {
	return []Any{t._1, t._2}
}

type Tuple3 struct {
	Tuple2
	_3 Any
}

func NewTuple3(a Any, b Any, c Any) Tuple3 {
	return Tuple3{
		Tuple2: Tuple2{_1: a, _2: b},
		_3:     c,
	}
}

func (t Tuple3) Trd() Any {
	return t._3
}

func (t Tuple3) MapTrd(f Morphism) Tuple3 {
	return NewTuple3(t._1, t._2, f(t._3))
}

func (t Tuple3) Append(x Any) Tuple4 {
	return NewTuple4(t._1, t._2, t._3, x)
}

func (t Tuple3) Slice() []Any {
	return []Any{t._1, t._2, t._3}
}

type Tuple4 struct {
	Tuple3
	_4 Any
}

func NewTuple4(a Any, b Any, c Any, d Any) Tuple4 {
	return Tuple4{
		Tuple3: Tuple3{
			Tuple2: Tuple2{_1: a, _2: b},
			_3:     c,
		},
		_4: d,
	}
}

func (t Tuple4) Fth() Any {
	return t._4
}

func (t Tuple4) MapFth(f Morphism) Tuple4 {
	return NewTuple4(t._1, t._2, t._3, f(t._4))
}
