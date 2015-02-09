package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractEndo(x Any) Any {
	return Endo_.As(x).Fork(Identity)
}

// Manual

func Test_endo_endo(t *testing.T) {
	f := func(x int) Any {
		return extractEndo(Endo(Constant(x)))
	}
	g := func(x int) Any {
		return extractEndo(endo{}.Of(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_endo_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return endo{}.Of(x).(Functor)
	}).Identity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_endo_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return endo{}.Of(x).(Functor)
	}).Composition(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monoid Laws

func Test_endo_MonoidLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonoidLaws(endo{}).LeftIdentity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_endo_MonoidLaws_RightIdentity(t *testing.T) {
	f, g := NewMonoidLaws(endo{}).RightIdentity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_endo_MonoidLaws_Associativity(t *testing.T) {
	f, g := NewMonoidLaws(endo{}).Associativity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

func Test_endo_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(endo{}).Associativity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
