package german_time

import (
	"fmt"
	"log"
	"time"
)

const zoneName = "Europe/Berlin"

// returns the location that is necessary to convert to/from German/Berlin local time
func getBerlinLocation() *time.Location {
	berlin, err := time.LoadLocation(zoneName)
	if err != nil {
		errorMsg := fmt.Errorf("The timezone data for '%s' could not be found. Import \"time/tzdata\" anywhere in your project or build with `-tags timetzdata`: https://pkg.go.dev/time/tzdata", zoneName)
		log.Panic(errorMsg)
	}
	return berlin
}

// returns the German time of the given timestamp. This method is not exported/public because users of the library should only get UTC objects returned.
func toGermanTime(timestamp time.Time) time.Time {
	berlin := getBerlinLocation()
	return timestamp.In(berlin)
}

// AddGermanDays converts timestamp to German local time, then adds 1 day and returns UTC. This will effectively add 24h on 363 out of 365 cases. But on the days on which germany switches from Daylight saving time (DST) to "normal" time or vice versa it might add 25 or 23 hours.
func AddGermanDays(timestamp time.Time, number int) time.Time {
	return toGermanTime(timestamp).AddDate(0, 0, number).UTC()
}

// StartOfGermanDay converts timestamp to German local time, then sets hour, minute and seconds to 0 and returns as UTC. The return value is always <= the given timestamp.
func StartOfGermanDay(timestamp time.Time) time.Time {
	germanTime := toGermanTime(timestamp)
	midnightInGermany := time.Date(germanTime.Year(), germanTime.Month(), germanTime.Day(), 0, 0, 0, 0, getBerlinLocation())
	return midnightInGermany.UTC()
}

// StartOfNextGermanDay converts timestamp to German local time, then returns the next start of German day (midnight, 00:00am German local time) as UTC. The return value is always > the given timestamp.
func StartOfNextGermanDay(timestamp time.Time) time.Time {
	plusOneDayInGermany := toGermanTime(timestamp).AddDate(0, 0, 1)
	return StartOfGermanDay(plusOneDayInGermany.UTC())
}

// StartOfGermanMonth converts timestamp to German local time, then returns the start of the German month (day, hours, minutes, seconds=0) as UTC. The return value is always <= the given timestamp.
func StartOfGermanMonth(timestamp time.Time) time.Time {
	germanTime := toGermanTime(timestamp)
	startOfMonthInGermany := time.Date(germanTime.Year(), germanTime.Month(), 1, 0, 0, 0, 0, getBerlinLocation())
	return startOfMonthInGermany.UTC()
}

// StartOfNextGermanMonth converts timestamp to German local time, then returns the start of the next German month (day, hours, minutes, seconds=0) as UTC. The return value is always > the given timestamp.
func StartOfNextGermanMonth(timestamp time.Time) time.Time {
	germanTime := toGermanTime(timestamp)
	startOfNextMonthInGermany := time.Date(germanTime.Year(), germanTime.Month()+1, 1, 0, 0, 0, 0, getBerlinLocation())
	return startOfNextMonthInGermany.UTC()
}

// GetGermanWeekday returns the weekday of the given timestamp in Germany.
func GetGermanWeekday(timestamp time.Time) time.Weekday {
	return toGermanTime(timestamp).Weekday()
}

// NextGermanWeekday returns the start of the next German weekday (as specified) in UTC. The result always > than the given timestamp. It might be up to 7 days later than the given timestamp. If e.g. providing a tuesday and requesting the next tuesday, the result will be the timestamp + 7 German days
func NextGermanWeekday(timestamp time.Time, weekday time.Weekday) time.Time {
	germanTime := StartOfNextGermanDay(timestamp).In(getBerlinLocation())
	for {
		if germanTime.Weekday() == weekday {
			return germanTime.UTC()
		}
		germanTime = germanTime.AddDate(0, 0, 1)
	}
}

// IsGermanMidnight returns true if and only if timestamp is midnight in German local time (namely 22:00:00 UTC in CEST or 23:00:00 UTC in CET).
func IsGermanMidnight(timestamp time.Time) bool {
	germanTime := toGermanTime(timestamp)
	return germanTime.UTC() == StartOfGermanDay(timestamp)
}
