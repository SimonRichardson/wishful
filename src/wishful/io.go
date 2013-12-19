package wishful

type IO struct {
    unsafePerform func() AnyVal
}

func (o IO) of(x AnyVal) IO {
    return IO{func () AnyVal {
        return x
    }}
}

// Methods
func (o IO) chain(f func(x AnyVal) IO) IO {
    return IO{func () AnyVal {
        return f(o.unsafePerform()).unsafePerform()
    }}
}

// Derived
func (o IO) ap(x IO) IO {
    return o.chain(func(f AnyVal) IO {
        return x.fmap(f.(func(f AnyVal) AnyVal))
    })
}
func (o IO) fmap(f func(x AnyVal) AnyVal) IO {
    return o.chain(func(x AnyVal) IO {
        return IO{func() AnyVal {
            return f(x)
        }}
    })
}