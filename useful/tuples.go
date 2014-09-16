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

func (t Tuple2) Get1() Any {
	return t._1
}

func (t Tuple2) Get2() Any {
	return t._2
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

func (t Tuple3) Get3() Any {
	return t._3
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

func (t Tuple4) Get4() Any {
	return t._4
}
