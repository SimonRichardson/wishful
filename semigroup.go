package wishful

type Semigroup interface {
    concat(x Semigroup) Semigroup
}