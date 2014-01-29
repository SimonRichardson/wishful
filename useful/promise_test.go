package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func extractPromise(x AnyVal) AnyVal {
	promise := x.(Promise)
	return promise.Fork(Identity)
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
