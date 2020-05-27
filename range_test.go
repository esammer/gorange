package gorange

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRange(t *testing.T) {
	r1 := Range{Begin: IntValue(1), End: IntValue(10)}
	r2 := Range{Begin: IntValue(5), End: IntValue(15)}
	r3 := r1.Merge(r2)
	r4 := Range{Begin: IntValue(20), End: IntValue(21)}
	r5 := Range{Begin: IntValue(0), End: IntValue(20)}

	require.Equal(t, "{Begin: 1, End: 10}", r1.String())
	require.Equal(t, "{Begin: 5, End: 15}", r2.String())
	require.Equal(t, "{Begin: 1, End: 15}", r3.String())
	require.Equal(t, "{Begin: 20, End: 21}", r4.String())
	require.Equal(t, "{Begin: 0, End: 20}", r5.String())

	// Check that r3 has the expected values post-Merge().
	require.Equal(t, IntValue(1), r3.Begin)
	require.Equal(t, IntValue(15), r3.End)

	require.True(t, r1.LessThan(r2))
	require.False(t, r2.LessThan(r1))
	// If two Begin values are mutually !LessThan() they're equal.
	require.False(t, r1.LessThan(r3))
	require.False(t, r3.LessThan(r1))

	require.True(t, r1.Contains(IntValue(3)))
	require.True(t, r1.Contains(IntValue(1)))
	require.False(t, r1.Contains(IntValue(10)))

	// Overlapping ranges. [1b [2b 1e] 2e]
	require.True(t, r1.Intersects(r2))
	require.True(t, r2.Intersects(r1))

	// Fully contained ranges. [5b [1b 1e] 5e]
	// 1's begin and end are within 5.
	require.True(t, r1.Intersects(r5))
	require.True(t, r5.Intersects(r1))

	// Non-overlapping ranges. [1b 1e] [4b 4e]
	require.False(t, r4.Intersects(r1))
	require.False(t, r1.Intersects(r4))
}

func BenchmarkRange(b *testing.B) {
	r1 := Range{Begin: IntValue(0), End: IntValue(10)}
	r2 := Range{Begin: IntValue(0), End: IntValue(5)}

	b.Run("LessThan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r1.LessThan(r2)
		}
	})
	b.Run("Before", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r1.Before(r2)
		}
	})
	b.Run("After", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r1.After(r2)
		}
	})
	b.Run("Intersects", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r1.Intersects(r2)
		}
	})
	b.Run("Merge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r1.Merge(r2)
		}
	})
	b.Run("Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r1.Contains(r1.Begin)
		}
	})
}
