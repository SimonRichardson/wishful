package wishful

// A combinator
func Apply(f func(x AnyVal) AnyVal) func(x AnyVal) AnyVal {
	return func(x AnyVal) AnyVal {
		return f(x)
	}
}

// B combinator
func Compose(f func(x AnyVal) AnyVal) func(g func(y AnyVal) AnyVal) func(z AnyVal) AnyVal {
	return func(g func(y AnyVal) AnyVal) func(z AnyVal) AnyVal {
		return func(a AnyVal) AnyVal {
			return f(g(a))
		}
	}
}

// K combinator
func Constant(a AnyVal) func(x AnyVal) AnyVal {
	return func(b AnyVal) AnyVal {
		return a
	}
}

// I combinator
func Identity(x AnyVal) AnyVal {
	return x
}

// T combinator
func Thrush(x AnyVal) func(f func(v AnyVal) AnyVal) AnyVal {
	return func(f func(v AnyVal) AnyVal) AnyVal {
		return f(x)
	}
}
