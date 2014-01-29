package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

// Applicative Laws

// Some

func Test_Option_Some_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// Some

func Test_Option_Some_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Some{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Some{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(None{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(None{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
