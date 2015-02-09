package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractIdT(x Any) Any {
	return IdT_.As(x).Run
}

// Applicative Laws

func Test_IdT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(IdT(id{})).Identity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(IdT(id{})).Composition(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(IdT(id{})).Homomorphism(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(IdT(id{})).Interchange(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_IdT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return IdT(id{}).Of(x).(Functor)
	}).Identity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return IdT(id{}).Of(x).(Functor)
	}).Composition(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_IdT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(IdT(id{})).LeftIdentity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(IdT(id{})).RightIdentity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(IdT(id{})).Associativity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
