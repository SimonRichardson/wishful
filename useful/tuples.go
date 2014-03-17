package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Tuple2 struct {
	_1 AnyVal
	_2 AnyVal
}

type Tuple3 struct {
	Tuple2
	_3 AnyVal
}

type Tuple4 struct {
	Tuple3
	_4 AnyVal
}
