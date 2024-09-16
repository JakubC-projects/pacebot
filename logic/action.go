package logic

import (
	"time"

	"github.com/samber/lo"
)

type ActionMilestone struct {
	Date      time.Time
	Milestone float64
}

var milestones = []ActionMilestone{
	{Date: lo.Must(time.Parse(time.RFC3339, "2024-08-01T00:00:00Z")), Milestone: 70},
	{Date: lo.Must(time.Parse(time.RFC3339, "2024-09-01T00:00:00Z")), Milestone: 76},
	{Date: lo.Must(time.Parse(time.RFC3339, "2024-10-01T00:00:00Z")), Milestone: 82},
	{Date: lo.Must(time.Parse(time.RFC3339, "2024-11-01T00:00:00Z")), Milestone: 88},
	{Date: lo.Must(time.Parse(time.RFC3339, "2024-12-01T00:00:00Z")), Milestone: 94},
	{Date: lo.Must(time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")), Milestone: 100},
}

func getStatusForNextMilestone(now time.Time) float64 {
	for _, m := range milestones {
		if m.Date.After(now) {
			return m.Milestone
		}
	}
	return 100
}
