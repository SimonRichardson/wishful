package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

// Manual tests

func Test_Either_Left_New(t *testing.T) {
	f := func(x int) left {
		return NewLeft(x)
	}
	g := func(x int) left {
		return left{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Of(t *testing.T) {
	f := func(x int) Either {
		return right{}.Of(x).(Either)
	}
	g := func(x int) Either {
		return left{}.Of(x).(Either)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Ap(t *testing.T) {
	f := func(x int) Either {
		return left{}.Ap(left{}).(Either)
	}
	g := func(x int) Either {
		return left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Chain(t *testing.T) {
	f := func(x int) Either {
		return left{}.Chain(func(v Any) Monad {
			return left{}
		}).(Either)
	}
	g := func(x int) Either {
		return left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Map(t *testing.T) {
	f := func(x int) Either {
		return left{}.Map(Identity).(Either)
	}
	g := func(x int) Either {
		return left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Concat(t *testing.T) {
	f := func(x int) Either {
		return left{}.Concat(left{}).(Either)
	}
	g := func(x int) Either {
		return left{}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Swap(t *testing.T) {
	f := func(x int) Either {
		return left{x}.Swap().(Either)
	}
	g := func(x int) Either {
		return right{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_Swap(t *testing.T) {
	f := func(x int) Either {
		return right{x}.Swap().(Either)
	}
	g := func(x int) Either {
		return left{x}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_Bimap(t *testing.T) {
	f := func(x int) Either {
		return left{x}.Bimap(func(v Any) Any {
			return v.(int) + 1
		}, Identity).(Either)
	}
	g := func(x int) Either {
		return left{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_Bimap(t *testing.T) {
	f := func(x int) Either {
		return right{x}.Bimap(Identity, func(v Any) Any {
			return v.(int) + 1
		}).(Either)
	}
	g := func(x int) Either {
		return right{x + 1}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

// left

func Test_Either_Left_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(left{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(left{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(left{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(left{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// right

func Test_Either_Right_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(right{}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(right{}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(right{}).Homomorphism(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(right{}).Interchange(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

// left

func Test_Either_Left_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return left{}.Of(x).(Functor)
	}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return left{}.Of(x).(Functor)
	}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// right

func Test_Either_Right_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return right{}.Of(x).(Functor)
	}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return right{}.Of(x).(Functor)
	}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// left

func Test_Either_Left_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(left{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(left{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Left_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(left{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// right

func Test_Either_Right_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(right{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(right{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Either_Right_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(right{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Semigroup Laws

// left

func Test_Either_Left_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(left{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// right

func Test_Either_Right_SemigroupLaws_Associativity(t *testing.T) {
	f, g := NewSemigroupLaws(right{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
