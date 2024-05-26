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
		{time: lo.Must(time.Parse(time.RFC3339, "2024-05-26T10:10:10Z")), percent: 57.6},
		{time: lo.Must(time.Parse(time.RFC3339, "2024-06-27T10:10:10Z")), percent: 70.0},
	}

	for _, tc := range testCases {
		assert.InDelta(t, tc.percent, getStatusForWeek(tc.time), 0.1)
	}
}
