package wishful

type Traversable interface {
	Traverse(g func(Any) Functor) Any
}
