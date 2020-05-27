package gorange

import (
	"fmt"
)

type Range struct {
	Begin RangeValue
	End   RangeValue
}

type RangeValue interface {
	LessThan(value RangeValue) bool
}

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

func (this *Range) Contains(v RangeValue) bool {
	return (this.Begin.LessThan(v) || (!this.Begin.LessThan(v) && !v.LessThan(this.Begin))) &&
		v.LessThan(this.End)
}

func (this *Range) Intersects(r Range) bool {
	return !this.Before(r) && !this.After(r)
}

func (this *Range) Before(r Range) bool {
	return this.End.LessThan(r.Begin)
}

func (this *Range) After(r Range) bool {
	return r.End.LessThan(this.Begin)
}

func (this *Range) LessThan(r Range) bool {
	return this.Begin.LessThan(r.Begin)
}

func (this *Range) String() string {
	return fmt.Sprintf("{Begin: %v, End: %v}", this.Begin, this.End)
}
