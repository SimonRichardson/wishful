package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractPromise(x Any) Any {
	promise := x.(Promise)
	return promise.Extract()
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
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Promise{}.Of(x).(Functor)
	}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Promise{}.Of(x).(Functor)
	}).Composition(extractPromise)
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

// Comonad Laws

func Test_Promise_ComonadLaws_Identity(t *testing.T) {
	f, g := NewComonadLaws(Promise{}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_ComonadLaws_Composition(t *testing.T) {
	f, g := NewComonadLaws(Promise{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Promise_ComonadLaws_Associativity(t *testing.T) {
	f, g := NewComonadLaws(Promise{}).Associativity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
