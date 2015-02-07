package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractReader(x Any) Any {
	reader := x.(Reader)
	return reader.Run(Identity)
}

// Applicative Laws

func Test_Reader_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Reader{}).Identity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Reader{}).Composition(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Reader{}).Homomorphism(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Reader{}).Interchange(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Reader_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Reader{}.Of(x).(Functor)
	}).Identity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Reader{}.Of(x).(Functor)
	}).Composition(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Reader_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Reader{}).LeftIdentity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Reader{}).RightIdentity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Reader_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Reader{}).Associativity(extractReader)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
