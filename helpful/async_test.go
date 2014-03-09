package helpful

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
	"testing"
	"testing/quick"
)

func success(x AnyVal) Promise {
	return Promise{}.Of(x).(Promise)
}

// Manual tests

func Test_Async(t *testing.T) {
	f := func(x string) string {
		get := Async(success)
		a := get(x)
		b := a.Fold(
			Identity,
			Identity,
		)
		return b.(Promise).Extract().(string)
	}
	g := func(x string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
