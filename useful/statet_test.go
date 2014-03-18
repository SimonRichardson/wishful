package useful

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/wishful"
)

func extractStateT(x AnyVal) AnyVal {
	io := x.(StateT)
	return io.Run(1)
}

// Applicative Laws

func Test_StateT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(Id{})).Identity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(Id{})).Composition(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(Id{})).Homomorphism(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(Id{})).Interchange(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_StateT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(NewStateT(Id{})).Identity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(NewStateT(Id{})).Composition(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_StateT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewStateT(Id{})).LeftIdentity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewStateT(Id{})).RightIdentity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(NewStateT(Id{})).Associativity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
