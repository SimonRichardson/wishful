package wishful

type Foldable interface {
	Fold(f func(v AnyVal) AnyVal, g func(v AnyVal) AnyVal) AnyVal
}
