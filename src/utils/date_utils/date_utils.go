package date_utils

import "time"

const(
	apiDateAndTimeLayout = "2006-01-02T15:04:05Z"
	apiDBLayout			 = "2006-01-02 15:04:05"
)

// Set GetNow for unit operations on current time
// now.Before, now.Format("2006") only the year
func GetNow() time.Time {
	return time.Now().UTC()
}

func GetDateNowInStringFormat() string {
	return GetNow().Format(apiDateAndTimeLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}