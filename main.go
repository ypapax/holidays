package holidays

import (
	"github.com/pkg/errors"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
	"time"
)

func IsHoliday(countryCode string, d time.Time) (actualHolidayInCalendar, observedNonWorkingDay bool, h *cal.Holiday, finalErr error) {
	var holidays []*cal.Holiday
	switch countryCode {
	case "us","en-us":
		holidays = us.Holidays
	default:
		finalErr = errors.Errorf("country code '%+v' is not supported", countryCode)
		return
	}
	c := cal.NewBusinessCalendar()
	// add holidays that the business observes
	c.AddHoliday(holidays...)
	actualHolidayInCalendar, observedNonWorkingDay, h = c.IsHoliday(d)
	return actualHolidayInCalendar, observedNonWorkingDay, h, nil
}