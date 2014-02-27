package useful

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

// Manual tests

func Test_Either_Left_New(t *testing.T) {
	f := func(x int) Left {
		return NewLeft(x)
	}
	g := func(x int) Left {
		return Left{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Of(t *testing.T) {
	f := func(x int) Either {
		return Right{}.Of(x).(Either)
	}
	g := func(x int) Either {
		return Left{}.Of(x).(Either)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Ap(t *testing.T) {
	f := func(x int) Either {
		return Left{}.Ap(Left{}).(Either)
	}
	g := func(x int) Either {
		return Left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Chain(t *testing.T) {
	f := func(x int) Either {
		return Left{}.Chain(func(v AnyVal) Monad {
			return Left{}
		}).(Either)
	}
	g := func(x int) Either {
		return Left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Map(t *testing.T) {
	f := func(x int) Either {
		return Left{}.Map(Identity).(Either)
	}
	g := func(x int) Either {
		return Left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Concat(t *testing.T) {
	f := func(x int) Either {
		return Left{}.Concat(Left{}).(Either)
	}
	g := func(x int) Either {
		return Left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Swap(t *testing.T) {
	f := func(x int) Either {
		return Left{x}.Swap().(Either)
	}
	g := func(x int) Either {
		return Right{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_Swap(t *testing.T) {
	f := func(x int) Either {
		return Right{x}.Swap().(Either)
	}
	g := func(x int) Either {
		return Left{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Bimap(t *testing.T) {
	f := func(x int) Either {
		return Left{x}.Bimap(func(v AnyVal) AnyVal {
			return v.(int) + 1
		}, Identity).(Either)
	}
	g := func(x int) Either {
		return Left{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_Bimap(t *testing.T) {
	f := func(x int) Either {
		return Right{x}.Bimap(Identity, func(v AnyVal) AnyVal {
			return v.(int) + 1
		}).(Either)
	}
	g := func(x int) Either {
		return Right{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

// Left

func Test_Either_Left_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Left{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Left{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Left{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Left{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Right

func Test_Either_Right_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(Right{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(Right{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(Right{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(Right{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// Left

func Test_Either_Left_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Left{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Left{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Right

func Test_Either_Right_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(Right{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(Right{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// Left

func Test_Either_Left_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Left{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Left{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Left{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Right

func Test_Either_Right_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(Right{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(Right{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(Right{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

// Left

func Test_Either_Left_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Left{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Right

func Test_Either_Right_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(Right{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
