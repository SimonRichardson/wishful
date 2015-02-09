package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractReader(x Any) Any {
	return Reader_.As(x).Run(Identity)
}

// Applicative Laws

func Test_Reader_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(reader{}).Identity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(reader{}).Composition(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(reader{}).Homomorphism(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(reader{}).Interchange(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Reader_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return reader{}.Of(x).(Functor)
	}).Identity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return reader{}.Of(x).(Functor)
	}).Composition(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Reader_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(reader{}).LeftIdentity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(reader{}).RightIdentity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(reader{}).Associativity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
