package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func extractIdT(x AnyVal) AnyVal {
	io := x.(IdT)
	return io.Run
}

// Applicative Laws

func Test_IdT_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Identity(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Composition(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Homomorphism(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_IdT_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(NewIdT(Id{})).Interchange(extractIdT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
