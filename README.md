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
It provides a specific implementation for Germany which is used to unit test the general implementation. 

<!-- todo: add go playground example here -->

### Conventions
All times returned by the packages function in `LocalDaysCalculator` are in UTC because the purpose of the package is to spare you from dealing with any non-UTC times.

## Implicit Requirements

The package requires your relevant timezone data to be present on the system on which you're using it.
It does _not_ include timezone data itself and will panic if the local timezone data is not found.
Please import the [`time/tzdata`](https://pkg.go.dev/time/tzdata) package from the std library, if necessary.

The package does not include any workarounds to actual timezone data (e.g. in the case of Germany calculating the last Sunday in March or October.)
You can do it but you probably shouldn't.
