package wishful

type Foldable interface {
	Fold(f Morphism, g Morphism) Any
}
