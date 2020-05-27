package gorange

import (
	"bytes"
	"time"
)

// Specific impls to allow range comparisons.

type StringValue string

func (this StringValue) LessThan(v RangeValue) bool {
	return this < v.(StringValue)
}

type IntValue int

func (this IntValue) LessThan(v RangeValue) bool {
	return this < v.(IntValue)
}

type Int8Value int32

func (this Int8Value) LessThan(v RangeValue) bool {
	return this < v.(Int8Value)
}

type Int16Value int32

func (this Int16Value) LessThan(v RangeValue) bool {
	return this < v.(Int16Value)
}

type Int32Value int32

func (this Int32Value) LessThan(v RangeValue) bool {
	return this < v.(Int32Value)
}

type Int64Value int64

func (this Int64Value) LessThan(v RangeValue) bool {
	return this < v.(Int64Value)
}

type UIntValue uint

func (this UIntValue) LessThan(v RangeValue) bool {
	return this < v.(UIntValue)
}

type UInt8Value uint16

func (this UInt8Value) LessThan(v RangeValue) bool {
	return this < v.(UInt8Value)
}

type UInt16Value uint16

func (this UInt16Value) LessThan(v RangeValue) bool {
	return this < v.(UInt16Value)
}

type UInt32Value uint32

func (this UInt32Value) LessThan(v RangeValue) bool {
	return this < v.(UInt32Value)
}

type UInt64Value uint64

func (this UInt64Value) LessThan(v RangeValue) bool {
	return this < v.(UInt64Value)
}

type Float32Value float32

func (this Float32Value) LessThan(v RangeValue) bool {
	return this < v.(Float32Value)
}

type Float64Value float32

func (this Float64Value) LessThan(v RangeValue) bool {
	return this < v.(Float64Value)
}

type BoolValue bool

func (this BoolValue) LessThan(v RangeValue) bool {
	return this != v.(BoolValue) && this == false
}

type BytesValue []byte

func (this BytesValue) LessThan(v RangeValue) bool {
	switch bytes.Compare(this, v.(BytesValue)) {
	case 0, 1:
		return false
	case -1:
		return true
	default:
		panic("compare returned illegal value")
	}
}

type TimeValue time.Time

func (this TimeValue) LessThan(v RangeValue) bool {
	return time.Time(this).Before(time.Time(v.(TimeValue)))
}

func (this TimeValue) String() string {
	return time.Time(this).String()
}
