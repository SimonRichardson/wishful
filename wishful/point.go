package wishful

type Point interface {
	Of(v AnyVal) Point
}
