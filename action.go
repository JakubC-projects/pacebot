package peacefulroad

import (
	"time"

	"github.com/samber/lo"
)

var ActionStartTime = lo.Must(time.Parse(time.RFC3339, "2024-03-06T00:00:00Z"))
var ActionEndTime = lo.Must(time.Parse(time.RFC3339, "2024-07-11T00:00:00Z"))
var ActionStartPercentage = 35.0
var ActionEndPercentage = 70.0
