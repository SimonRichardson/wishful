package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractEitherT(x Any) Any {
	io := x.(EitherT)
	return io.Run
}

// Applicative Laws

func Test_EitherT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(NewEitherT(Id{})).Identity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(NewEitherT(Id{})).Composition(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(NewEitherT(Id{})).Homomorphism(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(NewEitherT(Id{})).Interchange(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_EitherT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewEitherT(Id{}).Of(x).(Functor)
	}).Identity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewEitherT(Id{}).Of(x).(Functor)
	}).Composition(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_EitherT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewEitherT(Id{})).LeftIdentity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewEitherT(Id{})).RightIdentity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_EitherT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(NewEitherT(Id{})).Associativity(extractEitherT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
