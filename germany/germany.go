package germany

import (
	"github.com/hochfrequenz/go-local-days/local_days"
)

// NewGermanLocalDaysCalculator returns a converter that works for Germany. Internally it's based on the local_days.NewTimeZoneBasedLocalTimeConverter and tzdata for "Europe/Berlin".
func NewGermanLocalDaysCalculator() local_days.LocalDaysCalculator {
	const zoneName = "Europe/Berlin"
	return local_days.NewTimeZoneBasedLocalTimeConverter(zoneName)
}
