package wishful

import (
	"errors"
	"reflect"
)

type Function struct {
	callable reflect.Value
}

func NewFunction(f AnyVal) Function {
	return Function{
		callable: reflect.ValueOf(f),
	}
}

func (f Function) Call(args ...AnyVal) (AnyVal, error) {
	return f.Apply(args)
}

func (f Function) Apply(args []AnyVal) (AnyVal, error) {
	vargs := make([]reflect.Value, len(args))
	for i, v := range args {
		vargs[i] = reflect.ValueOf(v)
	}

	result := f.callable.Call(vargs)
	if len(result) != 1 {
		return nil, errors.New("Expected 1 value to be returned.")
	}

	return result[0].Interface(), nil
}
