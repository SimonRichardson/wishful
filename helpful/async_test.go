package helpful

import (
	"testing"
	"testing/quick"

	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func success(x Any) Monad {
	return Promise_.As(Promise_.Ref().Of(x))
}

// Manual tests

func Test_Async(t *testing.T) {
	f := func(x string) string {
		get := Async(success)
		a := get(x)
		b := EitherT_.As(a).Fold(
			Identity,
			Identity,
		)
		return Promise_.As(b).Extract().(string)
	}
	g := func(x string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
