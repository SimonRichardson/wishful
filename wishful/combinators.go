package wishful

// A combinator
func Apply(f func(x Any) Any) func(x Any) Any {
	return func(x Any) Any {
		return f(x)
	}
}

// B combinator
func Compose(f func(x Any) Any) func(g func(y Any) Any) func(z Any) Any {
	return func(g func(y Any) Any) func(z Any) Any {
		return func(a Any) Any {
			return f(g(a))
		}
	}
}

// K combinator
func Constant(a Any) func(x Any) Any {
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
func Thrush(x Any) func(f func(v Any) Any) Any {
	return func(f func(v Any) Any) Any {
		return f(x)
	}
}
