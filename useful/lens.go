package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"reflect"
)

type Accessor interface {
	Get(x AnyVal) AnyVal
	Set(x AnyVal, y AnyVal) AnyVal
}

type SliceIndex struct {
	Index int
}

func (s SliceIndex) extract(val AnyVal) (reflect.Value, reflect.Value) {
	src := reflect.ValueOf(val)
	num := src.Len()

	dst := reflect.MakeSlice(src.Type(), num, num)
	for i := 0; i < num; i++ {
		dst.Index(i).Set(src.Index(i))
	}

	return dst, dst.Index(s.Index)
}

func (s SliceIndex) Get(x AnyVal) AnyVal {
	_, b := s.extract(x)
	return b.Interface()
}

func (s SliceIndex) Set(x AnyVal, y AnyVal) AnyVal {
	a, b := s.extract(x)
	b.Set(reflect.ValueOf(y))
	return a.Interface()
}

type Lens struct {
	Run func(a AnyVal) Store
}

func NewLens(run func(a AnyVal) Store) Lens {
	return Lens{
		Run: run,
	}
}

func (x Lens) Id() Lens {
	return NewLens(func(a AnyVal) Store {
		return NewStore(
			Identity,
			ConstantNoArgs(a),
		)
	})
}

func (x Lens) SliceLens(accessor Accessor) Lens {
	return NewLens(func(a AnyVal) Store {
		return NewStore(
			func(b AnyVal) AnyVal {
				return accessor.Set(a, b)
			},
			func() AnyVal {
				return accessor.Get(a)
			},
		)
	})
}

func (x Lens) Compose(y Lens) Lens {
	return NewLens(func(target AnyVal) Store {
		a := y.Run(target)
		b := x.Run(a.Get())
		return NewStore(
			Compose(a.Set)(b.Set),
			b.Get,
		)
	})
}

func (x Lens) AndThen(y Lens) Lens {
	return y.Compose(x)
}
