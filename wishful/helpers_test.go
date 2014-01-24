package wishful

import (
	"testing"
	"testing/quick"
)

// IncInt

func Test_IncInt(t *testing.T) {
	f := func(v int) int {
		return v + 1
	}
	g := func(v int) int {
		return IncInt(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Inc

func Test_Inc_Int(t *testing.T) {
	f := func(v int) int {
		return v + 1
	}
	g := func(v int) int {
		return Inc(v).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Inc_Float32(t *testing.T) {
	f := func(v float32) float32 {
		return v + 1
	}
	g := func(v float32) float32 {
		return Inc(v).(float32)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Inc_Float64(t *testing.T) {
	f := func(v float64) float64 {
		return v + 1
	}
	g := func(v float64) float64 {
		return Inc(v).(float64)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Inc_string(t *testing.T) {
	f := func(v string) string {
		return v
	}
	g := func(v string) string {
		return Inc(v).(string)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// DecInt

func Test_DecInt(t *testing.T) {
	f := func(v int) int {
		return v - 1
	}
	g := func(v int) int {
		return DecInt(v)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// Dec

func Test_Dec_Int(t *testing.T) {
	f := func(v int) int {
		return v - 1
	}
	g := func(v int) int {
		return Dec(v).(int)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dec_Float32(t *testing.T) {
	f := func(v float32) float32 {
		return v - 1
	}
	g := func(v float32) float32 {
		return Dec(v).(float32)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dec_Float64(t *testing.T) {
	f := func(v float64) float64 {
		return v - 1
	}
	g := func(v float64) float64 {
		return Dec(v).(float64)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dec_string(t *testing.T) {
	f := func(v string) string {
		return v
	}
	g := func(v string) string {
		return Dec(v).(string)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
