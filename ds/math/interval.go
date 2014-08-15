package math

import (
	"bytes"
	"fmt"
)

// Interval defines an integer interval, which can be closed, open or
// semi-closed/open.
type Interval struct {
	// Start and End define the starting and ending points of the interval.
	Start, End int
	// StartClosed and EndClosed define whether interval end is open or
	// closed.
	StartClosed, EndClosed bool
}

// String returns the string representaion of the interval.
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

// Valid returns true if interval is valid, i.e. start <= end.
// A lot of operation will work incorrectly on an invalid interval, but they are
// allowed.
func (i *Interval) Valid() bool {
	return i.Start <= i.End
}

// Fix makes invalid interval valid.
func (i *Interval) Fix() {
	if !i.Valid() {
		i.Start, i.End = i.End, i.Start
	}
}

// Empty returns true if interval is valid and contains at least one point.
func (i *Interval) Empty() bool {
	return !i.Valid() || i.Start == i.End && !i.StartClosed && !i.EndClosed
}

// NewClosed returns a closed interval with given start and end points.
func NewClosed(s, e int) *Interval {
	return &Interval{s, e, true, true}
}

// NewOpen returns an open interval with given start and end points.
func NewOpen(s, e int) *Interval {
	return &Interval{s, e, false, false}
}

// NewInterval returns a new interval with given end points and given end point
// inclusion.
func NewInterval(s, e int, sc, ec bool) *Interval {
	return &Interval{s, e, sc, ec}
}

// Intersects returns true if o intersects i.
func (i *Interval) Intersects(o *Interval) bool {
	// Real intersection.
	if i.Start > o.Start && i.Start < o.End || o.Start > i.Start && o.Start < i.End {
		return true
	}

	// Ends are equal and ends are closed on both sides.
	if i.StartClosed && o.EndClosed && i.Start == o.End || i.EndClosed && o.StartClosed && i.End == o.Start {
		return true
	}
	return false
}
