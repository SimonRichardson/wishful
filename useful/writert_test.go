package useful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/wishful"
)

func extractWriterT(x Any) Any {
	writer := x.(WriterT)
	return writer.Run()
}

// Functor Laws

func Test_WriterT_FunctorLaws_Identity(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewWriterT(Id{}).Of(x).(Functor)
	}).Identity(extractWriterT)

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_WriterT_FunctorLaws_Composition(t *testing.T) {
	f, g := NewFunctorLaws(func(x Any) Functor {
		return NewWriterT(Id{}).Of(x).(Functor)
	}).Composition(extractWriterT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Monad Laws

func Test_WriterT_MonadLaws_LeftIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewWriterT(Id{})).LeftIdentity(extractWriterT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_WriterT_MonadLaws_RightIdentity(t *testing.T) {
	f, g := NewMonadLaws(NewWriterT(Id{})).RightIdentity(extractWriterT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_WriterT_MonadLaws_Associativity(t *testing.T) {
	f, g := NewMonadLaws(NewWriterT(Id{})).Associativity(extractWriterT)
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
