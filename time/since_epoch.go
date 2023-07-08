package time

import (
	"time"
)

func Since_Epoch() int64 {
	epoch_time := time.Date(2023, time.January, 6, 0, 0, 0, 0, time.UTC)
	time_now := time.Now()

	return time_now.Sub(epoch_time).Milliseconds()
}
