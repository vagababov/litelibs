package math

import (
	"testing"
)

func Test(t *testing.T) {
	tests := struct {
		f, s *Interval
		want bool }{
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
			New(1,2, true, false),
			New(2, 3, true, false),
			false,
		},
		{
			New(1,2, false, true),
			New(2, 3, true, false),
			true,
		},
		{
			New(2, 5, true, false),
			New(3, 4, false, false),
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
