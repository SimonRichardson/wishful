package wishful

// A combinator
func Apply(f Morphism) Morphism {
	return func(x Any) Any {
		return f(x)
	}
}

// B combinator
func Compose(f Morphism) func(Morphism) Morphism {
	return func(g Morphism) Morphism {
		return func(a Any) Any {
			return f(g(a))
		}
	}
}

// K combinator
func Constant(a Any) Morphism {
	return func(b Any) Any {
		return a
	}
}
func ConstantNoArgs(a Any) func() Any {
	return func() Any {
		return a
	}
}

// I combinator
func Identity(x Any) Any {
	return x
}

// T combinator
func Thrush(x Any) func(Morphism) Any {
	return func(f Morphism) Any {
		return f(x)
	}
}
