# go-local-days

[![Go Reference](https://pkg.go.dev/badge/github.com/hochfrequenz/go-local-days.svg)](https://pkg.go.dev/github.com/hochfrequenz/go-local-days)
![Unittest status badge](https://github.com/hochfrequenz/go-local-days/workflows/Unittests/badge.svg)
![Linter status badge](https://github.com/hochfrequenz/go-local-days/workflows/golangci-lint/badge.svg)

Timezones are annoying and converting between UTC and local times is error-prone.

The package `go-local-time` provides easy to use helper functions that are useful if your business logic relies on local date times but your application internally uses UTC (which it absolutely should).

## Why a package for such trivial computations?
This package does not contain any rocket science; In fact all functions are shorter than 5 lines of code, and you could probably implement them all by yourself in less than one or two hours.
But the package still saves you from thinking about stuff you do not want to think about and that would bloat your code.

It saves you from re-inventing the wheel for a problem that is already solved.
The package also comes with ~30 unittests, half of them pretty easy and straight forward but the other half of them addressing edge cases around the transition to/from daylight saving time.

## Installation

```bash
go get github.com/hochfrequenz/go-local-days
```

## Usage Examples

The package contains a general implementation that can easily be adapted to your local time zone.

```
package main

import (
	"fmt"
	"github.com/hochfrequenz/go-local-days/local_days"
	"time"
)

func main() {
	const zoneName = "Europe/Berlin"
	berlin := local_days.NewTimeZoneBasedLocalTimeConverter(zoneName)
	dateInGermanDaylightSavingTime := time.Date(2022, 10, 29, 10, 0, 0, 0, time.UTC) // local time in Germany: UTC+2
	theNextDayAtTheSameTimeInGermany := berlin.AddLocalDays(dateInGermanDaylightSavingTime, 1)
	fmt.Println(theNextDayAtTheSameTimeInGermany) // 2022-10-30 11:00:00 +0000 UTC (because this is UTC+1)
	// note that the hour switched from 10 to 11 because the "local day" in Germany has 25 hours on that day
}
```
[Go Playground](https://play.golang.com/p/JPlItKzIpK7)

For more code snippets, see the extensive [tests with examples from Germany](germany/germany_test.go).

### Conventions

All times returned by the packages function in `LocalDaysCalculator` are in UTC because the purpose of the package is to spare you from dealing with any non-UTC times.

### Full List of Features

See the `LocalDaysCalculator` interface:

```go
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
```

## Implicit Requirements

The package requires your relevant timezone data to be present on the system on which you're using it.
It does _not_ include timezone data itself and will panic if the local timezone data is not found.
Please import the [`time/tzdata`](https://pkg.go.dev/time/tzdata) package from the std library, if necessary.

The package does not include any workarounds to actual timezone data (e.g. in the case of Germany calculating the last Sunday in March or October.)
You can do it but you probably shouldn't.
