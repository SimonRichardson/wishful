package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
)

type Store struct {
	Set func(x AnyVal) AnyVal
	Get func() AnyVal
}

func NewStore(set func(x AnyVal) AnyVal, get func() AnyVal) Store {
	return Store{
		set,
		get,
	}
}

func (x Store) Map(f func(x AnyVal) AnyVal) Functor {
	return x.Extend(func(x Store) AnyVal {
		return f(x.Extract())
	})
}

// Derived

func (x Store) Extend(f func(x Store) AnyVal) Store {
	return Store{
		func(y AnyVal) AnyVal {
			fun := NewFunction(f)
			res, _ := fun.Call(Store{
				x.Set,
				func() AnyVal {
					return y
				},
			})
			return res
		},
		x.Get,
	}
}

func (x Store) Extract() AnyVal {
	return x.Set(x.Get())
}
