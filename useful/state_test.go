package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func extractState(x AnyVal) AnyVal {
	state := x.(State)
	return state.EvalState(1)
}

// Applicative Laws

func Test_State_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(State{}).Identity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(State{}).Composition(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(State{}).Homomorphism(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(State{}).Interchange(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_State_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(State{}).Identity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(State{}).Composition(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_State_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(State{}).LeftIdentity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(State{}).RightIdentity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(State{}).Associativity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
