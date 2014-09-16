package wishful

func IncInt(a int) int {
	return a + 1
}

func Inc(a Any) Any {
	if obj, ok := a.(int); ok {
		return obj + 1
	}
	if obj, ok := a.(float32); ok {
		return obj + 1.0
	}
	if obj, ok := a.(float64); ok {
		return obj + 1.0
	}
	return a
}

func DecInt(a int) int {
	return a - 1
}

func Dec(a Any) Any {
	if obj, ok := a.(int); ok {
		return obj - 1
	}
	if obj, ok := a.(float32); ok {
		return obj - 1.0
	}
	if obj, ok := a.(float64); ok {
		return obj - 1.0
	}
	return a
}
