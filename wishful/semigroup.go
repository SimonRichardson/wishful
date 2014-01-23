package wishful

type Semigroup interface {
	Concat(x Semigroup) Semigroup
}
