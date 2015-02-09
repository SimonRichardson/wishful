package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractEitherT(x Any) Any {
	return EitherT_.As(x).Run
}

// Applicative Laws

func Test_EitherT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(EitherT(id{})).Identity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(EitherT(id{})).Composition(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(EitherT(id{})).Homomorphism(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(EitherT(id{})).Interchange(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_EitherT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return EitherT(id{}).Of(x).(Functor)
	}).Identity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return EitherT(id{}).Of(x).(Functor)
	}).Composition(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_EitherT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(EitherT(id{})).LeftIdentity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(EitherT(id{})).RightIdentity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(EitherT(id{})).Associativity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
