// Package local_days encapsulate the logic to convert to local date times behind an interface (LocalDaysCalculator) and provides a straight forward implementation NewTimeZoneBasedLocalTimeConverter.
package local_days

import (
	"fmt"
	"log"
	"time"
)

// NewTimeZoneBasedLocalTimeConverter returns a ToLocalTimeConverter that internally uses the timezone data from the timezone with the given zoneName (e.g. "Europe/Berlin"). It requires the tzdata to be available on the system and will panic if this is not the case.
func NewTimeZoneBasedLocalTimeConverter(zoneName string) LocalDaysCalculator {
	location, err := time.LoadLocation(zoneName)
	if err != nil {
		errorMsg := fmt.Errorf("The timezone data for '%s' could not be found. Import \"time/tzdata\" anywhere in your project or build with `-tags timetzdata`: https://pkg.go.dev/time/tzdata", zoneName)
		log.Panic(errorMsg)
	}
	return locationBasedLocalTimeConverter{location: location}
}

type locationBasedLocalTimeConverter struct {
	location *time.Location
}

// ToLocalTimeConverter contains a method to convert a time into a local time. This will, in most cases, happen on the basis of timezone data, but you are free to write your own conversion, although you're probably missing out on details at one point.
type ToLocalTimeConverter interface {
	// toLocalTime converts a timestamp to a "local" time by adjusting date, time and UTC-offset. The actual point in time in UTC or Unix does _not_ change.
	toLocalTime(timestamp time.Time) time.Time
}

func (l locationBasedLocalTimeConverter) toLocalTime(timestamp time.Time) time.Time {
	return timestamp.In(l.location)
}

// LocalDaysCalculator is an interface that encapsulates common date time operations that involve local date times.
type LocalDaysCalculator interface {
	// AddLocalDays converts timestamp to local time, then adds 1 day and returns UTC. This will effectively add 24h on 363 out of 365 cases. But on the days on which the calendar switches from Daylight saving time (DST) to "normal" time or vice versa it might add 25 or 23 hours.
	AddLocalDays(timestamp time.Time, number int) time.Time
	// StartOfLocalDay converts timestamp to local time, then sets hour, minute and seconds to 0 and returns as UTC. The return value is always <= the given timestamp.
	StartOfLocalDay(timestamp time.Time) time.Time
	// StartOfNextLocalDay converts timestamp to local time, then returns the next start of local day (midnight, 00:00am local time) as UTC. The return value is always > the given timestamp.
	StartOfNextLocalDay(timestamp time.Time) time.Time
	// StartOfLocalMonth converts timestamp to local time, then returns the start of the local month (day, hours, minutes, seconds=0) as UTC. The return value is always <= the given timestamp.
	StartOfLocalMonth(timestamp time.Time) time.Time
	// StartOfNextLocalMonth converts timestamp to local time, then returns the start of the next local month (day, hours, minutes, seconds=0) as UTC. The return value is always > the given timestamp.
	StartOfNextLocalMonth(timestamp time.Time) time.Time
	// GetLocalWeekday returns the weekday of the given timestamp in local timezone.
	GetLocalWeekday(timestamp time.Time) time.Weekday
	// NextLocalWeekday returns the start of the next local weekday (as specified) in UTC. The result always > than the given timestamp. It might be up to 7 days later than the given timestamp. If e.g. providing a tuesday and requesting the next tuesday, the result will be the timestamp + 7 Local days
	NextLocalWeekday(timestamp time.Time, weekday time.Weekday) time.Time
	// IsLocalMidnight returns true if and only if timestamp is midnight in local time
	IsLocalMidnight(timestamp time.Time) bool
}

// the following implementations are tested by the package "germany"

func (l locationBasedLocalTimeConverter) AddLocalDays(timestamp time.Time, number int) time.Time {
	return l.toLocalTime(timestamp).AddDate(0, 0, number).UTC()
}

func (l locationBasedLocalTimeConverter) StartOfLocalDay(timestamp time.Time) time.Time {
	localTime := l.toLocalTime(timestamp)
	localMidnight := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 0, 0, 0, 0, l.location)
	return localMidnight.UTC()
}

func (l locationBasedLocalTimeConverter) StartOfNextLocalDay(timestamp time.Time) time.Time {
	plusOneDayInLocally := l.toLocalTime(timestamp).AddDate(0, 0, 1)
	return l.StartOfLocalDay(plusOneDayInLocally.UTC())
}

func (l locationBasedLocalTimeConverter) StartOfLocalMonth(timestamp time.Time) time.Time {
	localTime := l.toLocalTime(timestamp)
	startOfLocalMonth := time.Date(localTime.Year(), localTime.Month(), 1, 0, 0, 0, 0, l.location)
	return startOfLocalMonth.UTC()
}

func (l locationBasedLocalTimeConverter) StartOfNextLocalMonth(timestamp time.Time) time.Time {
	localTime := l.toLocalTime(timestamp)
	startOfNextLocalMonth := time.Date(localTime.Year(), localTime.Month()+1, 1, 0, 0, 0, 0, l.location)
	return startOfNextLocalMonth.UTC()
}

func (l locationBasedLocalTimeConverter) GetLocalWeekday(timestamp time.Time) time.Weekday {
	return l.toLocalTime(timestamp).Weekday()
}

func (l locationBasedLocalTimeConverter) NextLocalWeekday(timestamp time.Time, weekday time.Weekday) time.Time {
	localTime := l.StartOfNextLocalDay(timestamp).In(l.location)
	for {
		if localTime.Weekday() == weekday {
			return localTime.UTC()
		}
		localTime = localTime.AddDate(0, 0, 1)
	}
}

func (l locationBasedLocalTimeConverter) IsLocalMidnight(timestamp time.Time) bool {
	localTime := l.toLocalTime(timestamp)
	return localTime.UTC() == l.StartOfLocalDay(timestamp)
}
