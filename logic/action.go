package logic

import (
	"time"

	"github.com/samber/lo"
)

var actionStartTime = lo.Must(time.Parse(time.RFC3339, "2024-03-06T00:00:00Z"))
var actionEndTime = lo.Must(time.Parse(time.RFC3339, "2024-07-11T00:00:00Z"))
var actionStartPercentage = 35.0
var actionEndPercentage = 70.0

func getStatusForWeek(now time.Time) float64 {
	today := now.Truncate(time.Hour * 24)
	daysToNextMonday := int(7-today.Weekday())%7 + 1

	nextMonday := today.AddDate(0, 0, daysToNextMonday)

	actionDuration := float64(actionEndTime.Sub(actionStartTime))
	elapsedDuration := float64(nextMonday.Sub(actionStartTime))

	return elapsedDuration/actionDuration*(actionEndPercentage-actionStartPercentage) + actionStartPercentage
}
