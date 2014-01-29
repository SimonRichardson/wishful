package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func extractIO(x AnyVal) AnyVal {
	io := x.(IO)
	return io.UnsafePerform()
}

// Applicative Laws

func Test_IO_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(IO{}).Identity(extractIO)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IO_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(IO{}).Composition(extractIO)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IO_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(IO{}).Homomorphism(extractIO)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IO_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(IO{}).Interchange(extractIO)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_IO_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(IO{}).Identity(extractIO)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IO_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(IO{}).Composition(extractIO)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
