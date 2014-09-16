package useful

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Value struct {
	Val int
}

func (v Value) Concat(a Semigroup) Semigroup {
	return Value{v.Val + a.(Value).Val}
}

// Manual tests

func Test_Validation_Failure_New(t *testing.T) {
	f := func(x int) Failure {
		return NewFailure(x)
	}
	g := func(x int) Failure {
		return Failure{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Of(t *testing.T) {
	f := func(x int) Validation {
		return Success{}.Of(x).(Validation)
	}
	g := func(x int) Validation {
		return Failure{}.Of(x).(Validation)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Ap(t *testing.T) {
	f := func(x int) Validation {
		return Failure{Value{x}}.Ap(Failure{Value{1}}).(Validation)
	}
	g := func(x int) Validation {
		return Failure{Value{x + 1}}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApWithSuccess(t *testing.T) {
	f := func(x int) Validation {
		return Failure{Value{x}}.Ap(Success{Value{1}}).(Validation)
	}
	g := func(x int) Validation {
		return Failure{Value{x}}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Chain(t *testing.T) {
	f := func(x int) Validation {
		return Failure{}.Chain(func(v Any) Monad {
			return Failure{}
		}).(Validation)
	}
	g := func(x int) Validation {
		return Failure{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Map(t *testing.T) {
	f := func(x int) Validation {
		return Failure{}.Map(Identity).(Validation)
	}
	g := func(x int) Validation {
		return Failure{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Concat(t *testing.T) {
	f := func(x int) Validation {
		return Failure{Value{x}}.Concat(Failure{Value{1}}).(Validation)
	}
	g := func(x int) Validation {
		return Failure{Value{x + 1}}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_Bimap(t *testing.T) {
	f := func(x int) Validation {
		return Failure{x}.Bimap(func(v Any) Any {
			return v.(int) + 1
		}, Identity).(Validation)
	}
	g := func(x int) Validation {
		return Failure{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_Bimap(t *testing.T) {
	f := func(x int) Validation {
		return Success{x}.Bimap(Identity, func(v Any) Any {
			return v.(int) + 1
		}).(Validation)
	}
	g := func(x int) Validation {
		return Success{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

// Failure

func Test_Validation_Failure_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Failure{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Failure{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Failure{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Failure{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Success{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Success{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Success{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Success{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// Failure

func Test_Validation_Failure_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Failure{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Failure{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Success{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Success{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// Failure

func Test_Validation_Failure_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Failure{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Failure{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Failure_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Failure{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Success{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Success{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Validation_Success_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Success{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

// Failure

func Test_Validation_Failure_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Failure{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Success

func Test_Validation_Success_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Success{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
