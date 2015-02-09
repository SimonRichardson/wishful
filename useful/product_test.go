package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

// Manual tests

func Test_Product_Invalid(t *testing.T) {
	product{}.Of("x")
}

// Functor Laws

func Test_Product_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return product{}.Of(x).(Functor)
	}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Product_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return product{}.Of(x).(Functor)
	}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Product_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(product{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Product_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(product{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Product_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(product{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monoid Laws

func Test_Product_MonoidLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonoidLaws(product{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Product_MonoidLaws_RightIdentity(t *testing.T) {
	f, g := NewMonoidLaws(product{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Product_MonoidLaws_Associativity(t *testing.T) {
	f, g := NewMonoidLaws(product{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

func Test_Product_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(product{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
