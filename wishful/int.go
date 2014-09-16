package wishful

import (
	"errors"
	"math"
)

type Int int

func (o Int) Of(x int) Int {
	return Int(x)
}

func (o Int) Concat(x Semigroup) Semigroup {
	y, _ := FromAnyToInt(x)
	return o.Of(int(o) + int(y))
}

func FromAnyToInt(v Any) (Int, error) {
	if obj, ok := v.(int); ok {
		return Int(obj), nil
	} else if obj, ok := v.(Int); ok {
		return obj, nil
	} else {
		return Int(int(math.NaN())), errors.New("Type error, invalid Int")
	}
}
