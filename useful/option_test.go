package useful

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/wishful"
)

// Manual tests

func Test_Option_None_Empty(t *testing.T) {
	f := func(x int) Option {
		return None{}.Empty().(Option)
	}
	g := func(x int) Option {
		return Some{}.Empty().(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Of(t *testing.T) {
	f := func(x int) Option {
		return None{}.Of(x).(Option)
	}
	g := func(x int) Option {
		return Some{}.Of(x).(Option)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Ap(t *testing.T) {
	f := func(x int) Option {
		return None{}.Ap(None{}).(Option)
	}
	g := func(x int) Option {
		return None{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Chain(t *testing.T) {
	f := func(x int) Option {
		return None{}.Chain(func(v AnyVal) Monad {
			return None{}
		}).(Option)
	}
	g := func(x int) Option {
		return None{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Map(t *testing.T) {
	f := func(x int) Option {
		return None{}.Map(Identity).(Option)
	}
	g := func(x int) Option {
		return None{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_Concat(t *testing.T) {
	f := func(x int) Option {
		return None{}.Concat(None{}).(Option)
	}
	g := func(x int) Option {
		return None{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_GetOrElse(t *testing.T) {
	f := func(x int, y int) int {
		return x
	}
	g := func(x int, y int) int {
		return Some{}.Of(x).(Option).GetOrElse(ConstantNoArgs(y)).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_GetOrElse(t *testing.T) {
	f := func(x int, y int) int {
		return y
	}
	g := func(x int, y int) int {
		return None{}.GetOrElse(ConstantNoArgs(y)).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_OrElse(t *testing.T) {
	f := func(x int, y int) Option {
		return Some{x}
	}
	g := func(x int, y int) Option {
		return Some{x}.OrElse(Some{y})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_OrElse(t *testing.T) {
	f := func(x int, y int) Option {
		return Some{y}
	}
	g := func(x int, y int) Option {
		return None{}.OrElse(Some{y})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

// Some

func Test_Option_Some_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Some{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(None{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// Some

func Test_Option_Some_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Some{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Some{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(None{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(None{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// Some

func Test_Option_Some_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Some{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Some{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_Some_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Some{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(None{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(None{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Option_None_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(None{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

// Some

func Test_Option_Some_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Some{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// None

func Test_Option_None_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(None{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
