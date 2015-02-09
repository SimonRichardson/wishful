package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

// Functor Laws

// Return

func Test_Free_Return_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return ret{}.Of(x).(Functor)
	}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Free_Return_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return ret{}.Of(x).(Functor)
	}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// suspend

func Test_Free_suspend_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return suspend{}.Of(x).(Functor)
	}).Identity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Free_suspend_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return suspend{}.Of(x).(Functor)
	}).Composition(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

// Return

func Test_Free_Return_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(ret{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Free_Return_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(ret{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Free_Return_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(ret{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// suspend

func Test_Free_suspend_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(suspend{}).LeftIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Free_suspend_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(suspend{}).RightIdentity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Free_suspend_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(suspend{}).Associativity(Identity)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
