package wishful

import (
	"testing"
	"testing/quick"
)

// Identity

func Test_IdentityOf(t *testing.T) {
	f := func(v int) Option {
		return Id{v}
	}
	g := func(v int) Option {
		return Id{}.Of(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ApWithIdentity(t *testing.T) {
	f := func(v int) Option {
		return Id{v}
	}
	g := func(v int) Option {
		return Id{Identity}.Ap(Id{v})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Option

func Test_OptionSomeOf(t *testing.T) {
	f := func(v int) Option {
		return Some{v}
	}
	g := func(v int) Option {
		return Some{}.Of(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_OptionNoneOf(t *testing.T) {
	f := func(v int) Option {
		return Some{v}
	}
	g := func(v int) Option {
		return None{}.Of(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ApWithOptionSome(t *testing.T) {
	f := func(v int) Option {
		return Some{v}
	}
	g := func(v int) Option {
		return Some{Identity}.Ap(Some{v})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ApWithOptionNoneForApMethod(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return Some{Identity}.Ap(None{})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ApWithOptionNoneForApConstructor(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Ap(None{})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ApWithOptionNoneForApConstructorWithSome(t *testing.T) {
	f := func(v int) Option {
		return None{}
	}
	g := func(v int) Option {
		return None{}.Ap(Some{v})
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Promise

func Test_PromiseOf(t *testing.T) {
	f := func(v int) int {
		return Promise{func(resolve func(x AnyVal) AnyVal) AnyVal {
			return resolve(v)
		}}.Fork(func(x AnyVal) AnyVal {
			return x
		}).(int)
	}
	g := func(v int) int {
		return Promise{}.Of(v).(Promise).Fork(func(x AnyVal) AnyVal {
			return x
		}).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_PromiseAp(t *testing.T) {
	f := func(v int) int {
		return IncInt(v)
	}
	g := func(v int) int {
		app := Promise{}.Of(Inc).Ap(Promise{}.Of(v))
		pro := app.(Promise)
		return pro.Fork(func(x AnyVal) AnyVal {
			return x
		}).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
