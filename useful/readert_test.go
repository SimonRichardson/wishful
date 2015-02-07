package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractReaderT(x Any) Any {
	reader := x.(ReaderT)
	return reader.Run(Empty{})
}

// Functor Laws

func Test_ReaderT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewReaderT(Id{}).Of(x).(Functor)
	}).Identity(extractReaderT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ReaderT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewReaderT(Id{}).Of(x).(Functor)
	}).Composition(extractReaderT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_ReaderT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewReaderT(Id{})).LeftIdentity(extractReaderT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ReaderT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewReaderT(Id{})).RightIdentity(extractReaderT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ReaderT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(NewReaderT(Id{})).Associativity(extractReaderT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
