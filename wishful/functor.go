package wishful

type Functor interface {
	Map(f func(v AnyVal) AnyVal) Functor
}
