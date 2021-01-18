package date_utils

import "time"

const(
	apiDateAndTimeLayout = "2006-01-02T15:04:05Z"
)

// Set GetNow for unit operations on current time
// now.Before, now.Format("2006") only the year
func GetNow() time.Time {
	return time.Now().UTC()
}

func GetDateNowInStringFormat() string {
	now := GetNow()
	return now.Format(apiDateAndTimeLayout)
}
