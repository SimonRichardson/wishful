package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractPromise(x Any) Any {
	return Promise_.As(x).Extract()
}

// Applicative Laws

func Test_promise_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(promise{}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(promise{}).Composition(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(promise{}).Homomorphism(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(promise{}).Interchange(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_promise_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return promise{}.Of(x).(Functor)
	}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return promise{}.Of(x).(Functor)
	}).Composition(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_promise_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(promise{}).LeftIdentity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(promise{}).RightIdentity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(promise{}).Associativity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Comonad Laws

func Test_promise_ComonadLaws_Identity(t *testing.T) {
	f, g := NewComonadLaws(promise{}).Identity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_ComonadLaws_Composition(t *testing.T) {
	f, g := NewComonadLaws(promise{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_promise_ComonadLaws_Associativity(t *testing.T) {
	f, g := NewComonadLaws(promise{}).Associativity(extractPromise)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
