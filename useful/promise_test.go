package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func extractPromise(x AnyVal) AnyVal {
	promise := x.(Promise)
	return promise.Extract()
}

// Manual tests

func Test_Promise_Extend(t *testing.T) {
	f := func(x int) int {
		return x
	}
	g := func(x int) int {
		a := Promise{}.Of(x)
		b := a.(Promise).Extend(func(p Promise) AnyVal {
			return p.Extract()
		})
		return b.Extract().(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

func Test_Promise_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Promise{}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Promise{}).Composition(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Promise{}).Homomorphism(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Promise{}).Interchange(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Promise_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Promise{}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Promise{}).Composition(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Promise_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Promise{}).LeftIdentity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Promise{}).RightIdentity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Promise{}).Associativity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
