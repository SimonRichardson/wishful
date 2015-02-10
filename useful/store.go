package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Store struct {
	Set func(x Any) Any
	Get func() Any
}

func NewStore(set Morphism, get func() Any) Store {
	return Store{
		set,
		get,
	}
}

func (x Store) Map(f Morphism) Functor {
	return x.Extend(func(x Store) Any {
		return f(x.Extract())
	})
}

// Derived

func (x Store) Extend(f func(Store) Any) Store {
	return Store{
		func(y Any) Any {
			fun := NewFunction(f)
			res, _ := fun.Call(Store{
				x.Set,
				func() Any {
					return y
				},
			})
			return res
		},
		x.Get,
	}
}

func (x Store) Extract() Any {
	return x.Set(x.Get())
}
