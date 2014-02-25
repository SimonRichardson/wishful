package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"reflect"
)

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
func (x Lens) SliceLens(index int) Lens {
	return NewLens(func(a AnyVal) Store {
		// Is there away to clone the array, without cloning
		// the whole thing?
		src := reflect.ValueOf(a)
		dst := reflect.New(src.Type()).Elem()
		for i := 0; i < src.Len(); i++ {
			dst = reflect.Append(dst, src.Index(i))
		}
		val := dst.Index(index)

		return NewStore(
			func(b AnyVal) AnyVal {
				val.Set(reflect.ValueOf(b))
				return dst.Interface()
			},
			func() AnyVal {
				return val.Interface()
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
