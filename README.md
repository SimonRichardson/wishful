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

### General

An algebra is a set of values, a set of operators that it is closed under and some laws it must obey.

Each algebra is a separate specification. An algebra may have dependencies on other algebras which must be implemented. An algebra may also state other algebra methods which do not need to be implemented and how they can be derived from new methods.

### Support

Current supported versions of go.

- 1.2
- Release
- Tip

Support for go version 1.1 has been dropped, until I get time to work out why
casting of int to Int is resolved. See the following issue at [Travis](https://travis-ci.org/SimonRichardson/wishful/jobs/18049902)

### Fantasy Land Compatible

`wishful` is fantasy-land compatible

[
  ![](https://raw.github.com/fantasyland/fantasy-land/master/logo.png)
](https://github.com/fantasyland/fantasy-land)
