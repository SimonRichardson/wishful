package wishful

type Value interface {}

type Option interface {
    // Methods
    Chain(func(v Value) Option) Option
}

type Some struct {
    x Value
}

type None struct {}

// Methods
func (o Some) Chain(f func(x Value) Option) Option {
    return f(o.x)
}

func (o None) Chain(f func(x Value) Option) Option {
    return o
}
