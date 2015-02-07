package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractWriter(x Any) Any {
	writer := x.(Writer)
	return writer.Run()
}

// Applicative Laws

func Test_Writer_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Writer{}).Identity(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Writer_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Writer{}).Composition(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Writer_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Writer{}).Homomorphism(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Writer_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Writer{}).Interchange(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_Writer_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Writer{}.Of(x).(Functor)
	}).Identity(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Writer_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return Writer{}.Of(x).(Functor)
	}).Composition(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_Writer_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Writer{}).LeftIdentity(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Writer_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Writer{}).RightIdentity(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Writer_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Writer{}).Associativity(extractWriter)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
