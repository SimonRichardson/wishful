package wishful

type Monad interface {
	Chain(f func(v AnyVal) Monad) Monad
}
