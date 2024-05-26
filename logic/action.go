package logic

import (
	"time"

	"github.com/samber/lo"
)

var actionStartTime = lo.Must(time.Parse(time.RFC3339, "2024-05-20T00:00:00Z"))
var actionEndTime = lo.Must(time.Parse(time.RFC3339, "2024-07-01T00:00:00Z"))
var actionStartPercentage = 55.12
var actionEndPercentage = 70.0

func getStatusForWeek(now time.Time) float64 {
	today := now.Truncate(time.Hour * 24)
	daysToNextMonday := int(7-today.Weekday())%7 + 1

	nextMonday := today.AddDate(0, 0, daysToNextMonday)

	actionDuration := float64(actionEndTime.Sub(actionStartTime))
	elapsedDuration := float64(nextMonday.Sub(actionStartTime))

	return elapsedDuration/actionDuration*(actionEndPercentage-actionStartPercentage) + actionStartPercentage
}
