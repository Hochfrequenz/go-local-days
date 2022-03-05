# go-german-time

Timezones are annoying and converting between UTC and local times is error-prone.

The package `go-german-time` provides easy to use helper functions that are useful if your business logic relies on the German calendar but your application internally uses UTC (which it hopefully does).
The package does not contain any rocket science; in fact all functions are shorter than 5 lines of code.
But the package still saves you from thinking about stuff you do not want to think about and that would bloat your code.

## Installation

```bash
go get github.com/Hochfrequenz/go-german-time
```

## Usage Examples
<!-- todo: add go playground example here -->

### Conventions
All times returned by the packages function are in UTC because the purpose of the package is to spare you from dealing with any non-UTC times.

## Implicit Requirements

The package `go-german-time` requires the timezone `Europe/Berlin` to be present in the system on which you're using it.
It does _not_ include timezone data itself and will panic if `Europe/Berlin` data is not found.
Please import the [`time/tzdata`](https://pkg.go.dev/time/tzdata) package from the std library, if necessary.

The package does not include any workarounds to actual timezone data like calculating the last Sunday in March or October.
