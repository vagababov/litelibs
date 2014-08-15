package math

import (
	"testing"
)

func TestIntersects(t *testing.T) {
	tests := []struct {
		f, s *Interval
		want bool
	}{
		{
			NewClosed(1, 2),
			NewClosed(2, 3),
			true,
		},
		{
			NewOpen(1, 2),
			NewOpen(2, 3),
			false,
		},
		{
			NewClosed(1, 2),
			NewClosed(4, 5),
			false,
		},
		{
			NewInterval(1, 2, true, false),
			NewInterval(2, 3, true, false),
			false,
		},
		{
			NewInterval(1, 2, false, true),
			NewInterval(2, 3, true, false),
			true,
		},
		{
			NewInterval(2, 5, true, false),
			NewInterval(3, 4, false, false),
			true,
		},
	}
	for _, test := range tests {
		if got, want := test.f.Intersects(test.s), test.want; got != want {
			t.Errorf("%v ∩ %v: got: %v want: %v", test.f, test.s, got, want)
		}
		if got, want := test.s.Intersects(test.f), test.want; got != want {
			t.Errorf("%v ∩ %v: got: %v want: %v", test.s, test.f, got, want)
		}
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		i    *Interval
		want bool
	}{
		{
			NewClosed(2, 1),
			false,
		},
		{
			NewOpen(1, 2),
			true,
		},
		{
			NewOpen(1, 1),
			true,
		},
		{
			NewClosed(1, 1),
			true,
		},
	}
	for _, test := range tests {
		if got, want := test.i.Valid(), test.want; got != want {
			t.Errorf("%v.Valid: got: %v want: %v", test.i, got, want)
		}
		test.i.Fix()
		if got, want := test.i.Valid(), true; got != want {
			t.Errorf("%v.Valid: got: %v want: %v", test.i, got, want)
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		i    *Interval
		want bool
	}{
		{
			NewOpen(1, 1),
			true,
		},
		{
			NewClosed(1, 1),
			false,
		},
		{
			NewInterval(1, 1, true, false),
			false,
		},
		{
			NewInterval(1, 1, false, true),
			false,
		},
		{
			NewOpen(1, 2),
			false,
		},
	}
	for _, test := range tests {
		if got, want := test.i.Empty(), test.want; got != want {
			t.Errorf("%v.Empty: got: %v want: %v", test.i, got, want)
		}
	}
}
