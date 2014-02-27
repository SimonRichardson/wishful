# Fantasy Land Specification

(aka "Algebraic GO Specification")

This project specifies interoperability of common algebraic
structures:

* Semigroup
* Monoid
* Functor
* Applicative
* Chain
* Monad

## General

An algebra is a set of values, a set of operators that it is closed
under and some laws it must obey.

Each Fantasy Land algebra is a separate specification. An algebra may
have dependencies on other algebras which must be implemented. An
algebra may also state other algebra methods which do not need to be
implemented and how they can be derived from new methods.

## Terminology

1. "value" is any GO value, including any which have the
   structures defined below.
2. "equivalent" is an appropriate definition of equivalence for the given value.
    The definition should ensure that the two values can be safely swapped out in a program that respects abstractions. For example:
    - Two lists are equivalent if they are equivalent at all indices.
    - Two plain old GO objects, interpreted as dictionaries, are equivalent when they are equivalent for all keys.
    - Two promises are equivalent when they yield equivalent values.
    - Two functions are equivalent if they yield equivalent outputs for equivalent inputs.

## Algebras

### Semigroup

1. `a.Concat(b).Concat(c)` is equivalent to `a.Concat(b.Concat(c))` (associativity)

#### `Concat` method

A value which has a Semigroup must provide a `concat` method. The
`concat` method takes one argument:

    s.Concat(b)

1. `b` must be a value of the same Semigroup

    1. If `b` is not the same semigroup, behaviour of `Concat` is
       unspecified.

2. `Concat` must return a value of the same Semigroup.

### Monoid

A value that implements the Monoid specification must also implement
the Semigroup specficiation.

1. `m.Concat(m.Empty())` is equivalent to `m` (right identity)
2. `m.Empty().Concat(m)` is equivalent to `m` (left identity)

#### `empty` method

A value which has a Monoid must provide an `empty` method on itself. 
The `empty` method takes no arguments:

    m.Empty()

1. `Empty` must return a value of the same Monoid

### Functor

1. `u.Map(func(a AnyVal) AnyVal { return a; }))` is equivalent to `u` (identity)
2. `u.Map(func(x AnyVal) AnyVal { return f(g(x)); })` is equivalent to `u.Map(g).Map(f)` (composition)

#### `Map` method

A value which has a Functor must provide a `Map` method. The `Map`
method takes one argument:

    u.Map(f)

1. `f` must be a function,

    1. If `f` is not a function, the behaviour of `Map` is
       unspecified.
    2. `f` can return any value.

2. `Map` must return a value of the same Functor

### Applicative

A value that implements the Applicative specification must also
implement the Functor specification.

A value which satisfies the specification of a Applicative does not
need to implement:

* Functor's `Map`; derivable as `func(f func(x AnyVal) AnyVal) AnyVal { return x.Of(f).Ap(x); })}`

1. `a.Of(func(a AnyVal) AnyVal { return a; }).Ap(v)` is equivalent to `v` (identity)
2. `a.Of(func(f AnyVal) AnyVal { return func(g) { return func(x) { return f(g(x))}; }; }).Ap(u).Ap(v).Ap(w)` is equivalent to `u.Ap(v.Ap(w))` (composition)
3. `a.Of(f).Ap(a.Of(x))` is equivalent to `a.Of(f(x))` (homomorphism)
4. `u.Ap(a.Of(y))` is equivalent to `a.Of(func(f AnyVal) AnyVal { return f(y); }).Ap(u)` (interchange)

#### `Ap` method

A value which has an Applicative must provide an `Ap` method. The `Ap`
method takes one argument:

    a.Ap(b)

1. `a` must be an Applicative of a function,

    1. If `a` does not represent a function, the behaviour of `Ap` is
       unspecified.

2. `b` must be an Applicative of any value

3. `ap` must apply the function in Applicative `a` to the value in
   Applicative `b`

#### `Of` method

A value which has an Applicative must provide an `Of` method on itself. 
The `Of` method takes one argument:

    a.Of(b)

1. `Of` must provide a value of the same Applicative

    1. No parts of `b` should be checked

### Chain

1. `m.Chain(f).Chain(g)` is equivalent to `m.Chain(func(x AnyVal) AnyVal { return f(x).Chain(g); })` (associativity)

#### `Chain` method

A value which has a Chain must provide a `Chain` method. The `Chain`
method takes one argument:

    m.Chain(f)

1. `f` must be a function which returns a value

    1. If `f` is not a function, the behaviour of `Chain` is
       unspecified.
    2. `f` must return a value of the same Chain

2. `Chain` must return a value of the same Chain

### Monad

A value that implements the Monad specification must also implement
the Applicative and Chain specifications.

A value which satisfies the specification of a Monad does not need to
implement:

* Applicative's `Ap`; derivable as `funct(m AnyVal) Anyval { return this.chain(funct(f AnyVal) Monad { return m.Map(f); }); }`
* Functor's `map`; derivable as `function(f) { return x.Chain(func(a AnyVal) Monad { return x.Of(f(a)); })}`

1. `m.Of(a).Chain(f)` is equivalent to `f(a)` (left identity)
2. `m.Chain(m.Of)` is equivalent to `m` (right identity)






## Notes

1. If there's more than a single way to implement the methods and
   laws, the implementation should choose one and provide wrappers for
   other uses.
2. It's discouraged to overload the specified methods. It can easily
   result in broken and buggy behaviour.
3. It is recommended to throw an exception on unspecified behaviour.
