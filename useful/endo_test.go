package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractEndo(x Any) Any {
	endo := x.(Endo)
	return endo.Fork(Identity)
}

// Manual

func Test_Endo_NewEndo(t *testing.T) {
	f := func(x int) Any {
		return extractEndo(NewEndo(Constant(x)))
	}
	g := func(x int) Any {
		return extractEndo(Endo{}.Of(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Endo_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Endo{}.Of(x).(Functor)
	}).Identity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Endo_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Endo{}.Of(x).(Functor)
	}).Composition(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monoid Laws

func Test_Endo_MonoidLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonoidLaws(Endo{}).LeftIdentity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Endo_MonoidLaws_RightIdentity(t *testing.T) {
	f, g := NewMonoidLaws(Endo{}).RightIdentity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Endo_MonoidLaws_Associativity(t *testing.T) {
	f, g := NewMonoidLaws(Endo{}).Associativity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

func Test_Endo_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Endo{}).Associativity(extractEndo)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
