package proto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDate32_Time(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		v := time.Date(2011, 10, 10, 14, 59, 31, 401235, time.UTC)
		d := TimeToDate32(v)
		assert.Equal(t, Date32(31693), d)
		assert.Equal(t, NewDate32(2011, 10, 10), d)
		assert.Equal(t, d.String(), "2011-10-10")
		assert.Equal(t, d, TimeToDate32(d.Time()))
	})
	t.Run("Range", func(t *testing.T) {
		var (
			start = time.Date(1925, 1, 1, 0, 0, 0, 0, time.UTC)
			end   = time.Date(2283, 11, 11, 0, 0, 0, 0, time.UTC)
		)
		for v := start; v.Before(end); v = v.AddDate(0, 0, 1) {
			date := TimeToDate32(v)
			newTime := date.Time()
			require.True(t, newTime.Equal(v))
			newDate32 := NewDate32(newTime.Year(), newTime.Month(), newTime.Day())
			require.Equal(t, date, newDate32)
			require.Equal(t, v.Format("2006-01-02"), date.String())
		}
	})
}

func BenchmarkDate32_Time(b *testing.B) {
	b.ReportAllocs()

	v := Date32(100)
	var t time.Time
	for i := 0; i < b.N; i++ {
		t = v.Time()
	}
	_ = t.IsZero()
}
