package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Tuple2 struct {
	_1 AnyVal
	_2 AnyVal
}

func NewTuple2(a AnyVal, b AnyVal) Tuple2 {
	return Tuple2{_1: a, _2: b}
}

func (t Tuple2) Get1() AnyVal {
	return t._1
}

func (t Tuple2) Get2() AnyVal {
	return t._2
}

type Tuple3 struct {
	Tuple2
	_3 AnyVal
}

func NewTuple3(a AnyVal, b AnyVal, c AnyVal) Tuple3 {
	return Tuple3{
		Tuple2: Tuple2{_1: a, _2: b},
		_3:     c,
	}
}

func (t Tuple3) Get3() AnyVal {
	return t._3
}

type Tuple4 struct {
	Tuple3
	_4 AnyVal
}

func NewTuple4(a AnyVal, b AnyVal, c AnyVal, d AnyVal) Tuple4 {
	return Tuple4{
		Tuple3: Tuple3{
			Tuple2: Tuple2{_1: a, _2: b},
			_3:     c,
		},
		_4: d,
	}
}

func (t Tuple4) Get4() AnyVal {
	return t._4
}
