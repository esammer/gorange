package gorange

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStringValue_LessThan(t *testing.T) {
	v := StringValue("abc")
	require.Equal(t, true, v.LessThan(StringValue("def")))
	require.Equal(t, false, v.LessThan(StringValue("123")))
	require.Equal(t, false, v.LessThan(v))
}

func TestIntValue_LessThan(t *testing.T) {
	v := IntValue(1)
	require.Equal(t, true, v.LessThan(IntValue(2)))
	require.Equal(t, false, v.LessThan(IntValue(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestInt8Value_LessThan(t *testing.T) {
	v := Int8Value(1)
	require.Equal(t, true, v.LessThan(Int8Value(2)))
	require.Equal(t, false, v.LessThan(Int8Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestInt16Value_LessThan(t *testing.T) {
	v := Int16Value(1)
	require.Equal(t, true, v.LessThan(Int16Value(2)))
	require.Equal(t, false, v.LessThan(Int16Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestInt32Value_LessThan(t *testing.T) {
	v := Int32Value(1)
	require.Equal(t, true, v.LessThan(Int32Value(2)))
	require.Equal(t, false, v.LessThan(Int32Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestInt64Value_LessThan(t *testing.T) {
	v := Int64Value(1)
	require.Equal(t, true, v.LessThan(Int64Value(2)))
	require.Equal(t, false, v.LessThan(Int64Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestUIntValue_LessThan(t *testing.T) {
	v := UIntValue(1)
	require.Equal(t, true, v.LessThan(UIntValue(2)))
	require.Equal(t, false, v.LessThan(UIntValue(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestUInt8Value_LessThan(t *testing.T) {
	v := UInt8Value(1)
	require.Equal(t, true, v.LessThan(UInt8Value(2)))
	require.Equal(t, false, v.LessThan(UInt8Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestUInt16Value_LessThan(t *testing.T) {
	v := UInt16Value(1)
	require.Equal(t, true, v.LessThan(UInt16Value(2)))
	require.Equal(t, false, v.LessThan(UInt16Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestUInt32Value_LessThan(t *testing.T) {
	v := UInt32Value(1)
	require.Equal(t, true, v.LessThan(UInt32Value(2)))
	require.Equal(t, false, v.LessThan(UInt32Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestUInt64Value_LessThan(t *testing.T) {
	v := UInt64Value(1)
	require.Equal(t, true, v.LessThan(UInt64Value(2)))
	require.Equal(t, false, v.LessThan(UInt64Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestFloat32Value_LessThan(t *testing.T) {
	v := Float32Value(1)
	require.Equal(t, true, v.LessThan(Float32Value(2)))
	require.Equal(t, false, v.LessThan(Float32Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestFloat64Value_LessThan(t *testing.T) {
	v := Float64Value(1)
	require.Equal(t, true, v.LessThan(Float64Value(2)))
	require.Equal(t, false, v.LessThan(Float64Value(0)))
	require.Equal(t, false, v.LessThan(v))
}

func TestBoolValue_LessThan(t *testing.T) {
	v := BoolValue(false)
	require.Equal(t, true, v.LessThan(BoolValue(true)))
	require.Equal(t, false, BoolValue(true).LessThan(v))
	require.Equal(t, false, v.LessThan(v))
}

func TestBytesValue_LessThan(t *testing.T) {
	v := BytesValue("abc")
	require.Equal(t, true, v.LessThan(BytesValue("def")))
	require.Equal(t, false, v.LessThan(BytesValue("123")))
	require.Equal(t, false, v.LessThan(v))
}

func TestTimeValue_LessThan(t *testing.T) {
	t0 := time.Time{}
	t1 := time.Now()

	v := TimeValue(t1)
	// FIXME: Non-deterministic, but unlikely that instructions between t1 and here take zero nanos.
	require.Equal(t, true, v.LessThan(TimeValue(time.Now())))
	require.Equal(t, false, v.LessThan(TimeValue(t0)))
	require.Equal(t, false, v.LessThan(v))
}
