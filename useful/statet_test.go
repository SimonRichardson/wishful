package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractStateT(x Any) Any {
	io := x.(StateT)
	return io.Run(1)
}

// Applicative Laws

func Test_StateT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(id{})).Identity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(id{})).Composition(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(id{})).Homomorphism(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(NewStateT(id{})).Interchange(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_StateT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewStateT(id{}).Of(x).(Functor)
	}).Identity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewStateT(id{}).Of(x).(Functor)
	}).Composition(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_StateT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewStateT(id{})).LeftIdentity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewStateT(id{})).RightIdentity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_StateT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(NewStateT(id{})).Associativity(extractStateT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
