package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractState(x Any) Any {
	return State_.As(x).EvalState(1)
}

// Manual tests

func Test_State_ExecState(t *testing.T) {
	f := func(x int, y int) int {
		return y
	}
	g := func(x int, y int) int {
		a := state{}.Of(x).(state)
		return a.ExecState(y).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_Get(t *testing.T) {
	f := func(x int) (Any, Any) {
		return x, x
	}
	g := func(x int) (Any, Any) {
		a := state{}.Get()
		return a.Run(x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_Modify(t *testing.T) {
	f := func(x int) (Any, Any) {
		return nil, x
	}
	g := func(x int) (Any, Any) {
		a := state{}.Modify(Identity)
		return a.Run(x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_Put(t *testing.T) {
	f := func(x int, y int) (Any, Any) {
		return x, y
	}
	g := func(x int, y int) (Any, Any) {
		a := state{}.Put(x, y)
		return a.Run(x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Applicative Laws

func Test_State_ApplicativeLaws_Identity(t *testing.T) {
	f, g := NewApplicativeLaws(state{}).Identity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_ApplicativeLaws_Composition(t *testing.T) {
	f, g := NewApplicativeLaws(state{}).Composition(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_ApplicativeLaws_Homomorphism(t *testing.T) {
	f, g := NewApplicativeLaws(state{}).Homomorphism(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_ApplicativeLaws_Interchange(t *testing.T) {
	f, g := NewApplicativeLaws(state{}).Interchange(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Functor Laws

func Test_State_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return state{}.Of(x).(Functor)
	}).Identity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return state{}.Of(x).(Functor)
	}).Composition(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_State_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(state{}).LeftIdentity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(state{}).RightIdentity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_State_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(state{}).Associativity(extractState)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
