package gorange

import (
	"fmt"
)

// A continuous range of values.
//
// A range is made up of a begin and end value. These range values are implementations of the RangeValue interface. For
// the purposes of methods like Contains(), Range is inclusive of its Begin value and exclusive of its End. Like
// time.Time, callers should generally use values rather than pointers when working with Ranges since they only contain
// interface values (which are 2 * architecture word size per https://golang.org/src/go/types/sizes.go, or 16B on
// amd64). Ranges must be finite, containing both a Begin and End value. Infinite (unbound) ranges are not currently
// supported and will yield undefined behavior.
//
// Unless otherwise stated, all methods support overlapping ranges.
type Range struct {
	Begin RangeValue
	End   RangeValue
}

// A value in a range.
//
// A range value is any type that can determine its order relative to another range value of the same type.
// Implementations should return true if and only if the receiver is less than its argument. Specifically, if two values
// are equal, false should be returned. Range relies on the property that two values that are mutually less than each
// other are considered to be equal.
//
// This library includes implementations for most Go builtin types including string, all forms of int and uint, float,
// bool, and time.Time.
type RangeValue interface {
	LessThan(value RangeValue) bool
}

// Merge two ranges into a single continuous range.
//
// The resulting range's Begin is the lesser of the two Begins, and its End is the greater of the two Ends. Note that
// merging two non-continuous ranges results in a range containing both which may not be the desired result. Consider
// the following example:
//
// 	r1 := Range{Begin: IntValue(0), End: IntValue(5)}
// 	r2 := Range{Begin: IntValue(10), End: IntValue(20)}
// 	r3 := r1.Merge(r2) // r3: Range{Begin: IntValue(0), End: IntValue(20)}
//
// In this case, [0, 5).Merge([10, 20)) yields a larger range of [0, 20).
func (this *Range) Merge(r Range) Range {
	merged := Range{Begin: this.Begin, End: this.End}

	if !this.Begin.LessThan(r.Begin) {
		merged.Begin = r.Begin
	}

	if this.End.LessThan(r.End) {
		merged.End = r.End
	}

	return merged
}

// Determine whether a given value is within the range.
//
// Returns true if Begin <= v < End.
func (this *Range) Contains(v RangeValue) bool {
	return (this.Begin.LessThan(v) || (!this.Begin.LessThan(v) && !v.LessThan(this.Begin))) &&
		v.LessThan(this.End)
}

// Determine if a range overlaps or intersects with r.
func (this *Range) Intersects(r Range) bool {
	return !this.Before(r) && !this.After(r)
}

// Determine if a range ends before a given range begins.
//
// Returns true if the receivers End is less than r's Begin.
func (this *Range) Before(r Range) bool {
	return this.End.LessThan(r.Begin)
}

// Determine if a range begins after a given range ends.
//
// Returns true if r's End is less than the receiver's Begin.
func (this *Range) After(r Range) bool {
	return r.End.LessThan(this.Begin)
}

// Determine if one range is less than another.
//
// Returns true if the receiver's Begin is less than r's Begin, false in any other case.
func (this *Range) LessThan(r Range) bool {
	return this.Begin.LessThan(r.Begin)
}

func (this *Range) String() string {
	return fmt.Sprintf("{Begin: %v, End: %v}", this.Begin, this.End)
}
