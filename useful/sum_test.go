package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

// Functor Laws

func Test_Sum_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Sum{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Sum_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Sum{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Sum_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Sum{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Sum_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Sum{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Sum_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Sum{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monoid Laws

func Test_Sum_MonoidLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonoidLaws(Sum{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Sum_MonoidLaws_RightIdentity(t *testing.T) {
	f, g := NewMonoidLaws(Sum{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Sum_MonoidLaws_Associativity(t *testing.T) {
	f, g := NewMonoidLaws(Sum{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

func Test_Sum_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Sum{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
