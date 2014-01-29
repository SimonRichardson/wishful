package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func extractIdT(x AnyVal) AnyVal {
	io := x.(IdT)
	return io.Run
}

// Applicative Laws

func Test_IdT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Identity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Composition(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Homomorphism(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Interchange(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_IdT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(NewIdT(Id{})).Identity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(NewIdT(Id{})).Composition(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws
/*
func Test_IdT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(IdT{}).LeftIdentity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(IdT{}).RightIdentity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(IdT{}).Associativity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
*/
