package holidays

import (
	"github.com/pkg/errors"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/de"
	"github.com/rickar/cal/v2/es"
	"github.com/rickar/cal/v2/fr"
	"github.com/rickar/cal/v2/gb"
	"github.com/rickar/cal/v2/it"
	"github.com/rickar/cal/v2/us"
	"time"
)

func IsHoliday(countryCode string, d time.Time) (actualHolidayInCalendar, observedNonWorkingDay bool, h *cal.Holiday, finalErr error) {
	hh := ByCountryCode(countryCode)
	if hh == nil {
		finalErr = errors.Errorf("country code '%+v' is not supported", countryCode)
		return
	}
	c := cal.NewBusinessCalendar()
	// add holidays that the business observes
	c.AddHoliday(hh...)
	actualHolidayInCalendar, observedNonWorkingDay, h = c.IsHoliday(d)
	return actualHolidayInCalendar, observedNonWorkingDay, h, nil
}

func ByCountryCode(countryCode string) []*cal.Holiday {
	switch countryCode {
	case "us","en-us":
		return us.Holidays
	case "fr":
		return fr.Holidays
	case "it":
		return it.Holidays
	case "de":
		return de.Holidays
	case "en", "gb", "uk":
		return gb.Holidays
	case "es":
		return es.Holidays
	default:
		return nil
	}
}