package math

import (
	"bytes"
	"fmt"
)

type Interval struct {
	Start, End             int
	StartClosed, EndClosed bool
}

func (i *Interval) String() string {
	var b bytes.Buffer
	if i.StartClosed {
		b.WriteRune('[')
	} else {
		b.WriteRune('(')
	}
	b.WriteString(fmt.Sprintf("%d, %d)", i.Start, i.End))
	if i.EndClosed {
		b.WriteRune(']')
	} else {
		b.WriteRune(')')
	}
	return b.String()
}

func NewClosed(s, e int) *Interval {
	return &Interval{s, e, true, true}
}

func NewOpen(s, e int) *Interval {
	return &Interval{s, e, false, false}
}

func NewInterval(s, e int, sc, ec bool) *Interval {
	return &Interval{s, e, sc, ec}
}

func (i *Interval) Intersects(o *Interval) bool {
	// Real intersection.
	if i.Start > o.Start && i.Start < i.End || o.Start > i.Start && o.Start < i.End {
		return true
	}
	// Ends are equal and ends are closed on both sides.
	if i.StartClosed && o.StartClosed && i.Start == o.End || i.EndClosed && o.StartClosed && i.End == o.Start {
		return true
	}
	return false
}
