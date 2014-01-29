package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

// Applicative Laws

func Test_Id_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Id{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Id_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
