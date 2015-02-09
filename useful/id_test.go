package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

// Applicative Laws

func Test_Id_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(id{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(id{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Id_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return id{}.Of(x).(Functor)
	}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return id{}.Of(x).(Functor)
	}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Id_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(id{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(id{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(id{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

func Test_Id_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(id{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Comonad Laws

func Test_Id_ComonadLaws_Identity(t *testing.T) {
	f, g := NewComonadLaws(id{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ComonadLaws_Composition(t *testing.T) {
	f, g := NewComonadLaws(id{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Id_ComonadLaws_Associativity(t *testing.T) {
	f, g := NewComonadLaws(id{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
