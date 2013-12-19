package wishful

// A combinator
func apply(f func(x AnyVal) AnyVal) func(x AnyVal) AnyVal {
    return func(x AnyVal) AnyVal {
        return f(x)
    }
}

// B combinator
func compose(f func(x AnyVal) AnyVal) func(g func(y AnyVal) AnyVal) func(z AnyVal) AnyVal  {
    return func(g func(y AnyVal) AnyVal) func(z AnyVal) AnyVal {
        return func(a AnyVal) AnyVal {
            return f(g(a))
        }
    }
}

// K combinator
func constant(a AnyVal) func(x AnyVal) AnyVal {
    return func(b AnyVal) AnyVal {
        return a
    }
}

// I combinator
func identity(x AnyVal) AnyVal {
    return x
}