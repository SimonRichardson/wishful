wishful
=======

Monads in Go lang

[![Build Status](https://api.travis-ci.org/SimonRichardson/wishful.png?branch=develop)](https://travis-ci.org/SimonRichardson/wishful)

### Wishful

(aka "Algebraic GO Specification")

This project specifies interoperability of common algebraic structures:

Semigroup
Monoid
Functor
Applicative
Chain
Monad

### Useful

The useful lib provides actual implementations of Wishful spec.

### General

An algebra is a set of values, a set of operators that it is closed under and 
some laws it must obey.

Each algebra is a separate specification. An algebra may have dependencies on 
other algebras which must be implemented. An algebra may also state other 
algebra methods which do not need to be implemented and how they can be 
derived from new methods.

### Support

Current supported versions of go.

- 1.3
- Release
- Tip

### Fantasy Land Compatible

`wishful` is fantasy-land compatible

[
  ![](https://raw.github.com/fantasyland/fantasy-land/master/logo.png)
](https://github.com/fantasyland/fantasy-land)
