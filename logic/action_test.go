package logic

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestStatusForWeek(t *testing.T) {
	testCases := []struct {
		time    time.Time
		percent float64
	}{
		{time: lo.Must(time.Parse(time.RFC3339, "2024-04-22T10:10:10Z")), percent: 49.9},
		{time: lo.Must(time.Parse(time.RFC3339, "2024-04-25T10:10:10Z")), percent: 49.9},
		{time: lo.Must(time.Parse(time.RFC3339, "2024-04-20T10:10:10Z")), percent: 48.0},
	}

	for _, tc := range testCases {
		assert.InDelta(t, tc.percent, getStatusForWeek(tc.time), 0.1)
	}
}
