package wishful

// Identity

type Id struct {
	x AnyVal
}

// Option

type Option interface {
}
type Some struct {
	x AnyVal
}
type None struct {
}
