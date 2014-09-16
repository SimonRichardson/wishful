package wishful

type Foldable interface {
	Fold(f func(v Any) Any, g func(v Any) Any) Any
}
