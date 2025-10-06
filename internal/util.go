package internal

import (
	"fmt"
	"strconv"

	"github.com/elbombardi/yt_captions/models"
)

func newTimeStamp(ms int64) models.TimeStamp {
	return models.TimeStamp{
		Milliseconds: ms,
		HMS:          convertMillisecondsToHMS(ms),
	}
}

// convertMillisecondsToHMS converts milliseconds to a string in the format H:M:S
func convertMillisecondsToHMS(ms int64) string {
	hours := ms / 3600000
	minutes := (ms % 3600000) / 60000
	seconds := (ms % 60000) / 1000
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// toMilliseconds converts a string containing seconds a float  to milliseconds.
func toMilliseconds(st string) (int64, error) {
	sec, err := strconv.ParseFloat(st, 64)
	if err != nil {
		return 0, err
	}
	return int64(sec * 1000), nil
}
