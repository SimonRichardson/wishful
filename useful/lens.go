package useful

import (
	"reflect"

	. "github.com/SimonRichardson/wishful/wishful"
)

type Accessor interface {
	Get(Any) Any
	Set(Any, Any) Any
}

type SliceIndex struct {
	Index int
}

func (s SliceIndex) extract(val Any) (reflect.Value, reflect.Value) {
	src := reflect.ValueOf(val)
	dst := reflect.New(src.Type()).Elem()
	dst.Set(src)

	return dst, dst.Index(s.Index)
}

func (s SliceIndex) Get(x Any) Any {
	_, b := s.extract(x)
	return b.Interface()
}

func (s SliceIndex) Set(x Any, y Any) Any {
	a, b := s.extract(x)
	b.Set(reflect.ValueOf(y))
	return a.Interface()
}

type Lens struct {
	Run func(a Any) Store
}

func NewLens(run func(a Any) Store) Lens {
	return Lens{
		Run: run,
	}
}

func (x Lens) Id() Lens {
	return NewLens(func(a Any) Store {
		return NewStore(
			Identity,
			ConstantNoArgs(a),
		)
	})
}

func (x Lens) AccessorLens(accessor Accessor) Lens {
	return NewLens(func(a Any) Store {
		return NewStore(
			func(b Any) Any {
				return accessor.Set(a, b)
			},
			func() Any {
				return accessor.Get(a)
			},
		)
	})
}

func (x Lens) ObjectLens(property string) Lens {
	return NewLens(func(a Any) Store {
		src := reflect.ValueOf(a)
		dst := reflect.New(src.Type()).Elem()
		dst.Set(src)
		val := dst.FieldByName(property)

		return NewStore(
			func(b Any) Any {
				val.Set(reflect.ValueOf(b))
				return dst.Interface()
			},
			func() Any {
				return val.Interface()
			},
		)
	})
}

func (x Lens) SliceLens(index int) Lens {
	return NewLens(func(a Any) Store {
		src := reflect.ValueOf(a)
		dst := reflect.New(src.Type()).Elem()
		dst.Set(src)
		val := dst.Index(index)

		return NewStore(
			func(b Any) Any {
				val.Set(reflect.ValueOf(b))
				return dst.Interface()
			},
			func() Any {
				return val.Interface()
			},
		)
	})
}

func (x Lens) Compose(y Lens) Lens {
	return NewLens(func(target Any) Store {
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
