// The following directives are based on The Open Group Base Specifications
// Issue 7 strftime man page.
// http://pubs.opengroup.org/onlinepubs/9699919799/functions/strftime.html

package lctime

import (
	"fmt"
	"time"
)

// pera returns the locale's abbreviated weekday name.
func (lc *localeData) pera(t time.Time) string {
	return lc.ShortDays[int(t.Weekday())]
}

// perA returns the locale's full weekday name.
func (lc *localeData) perA(t time.Time) string {
	return lc.Days[int(t.Weekday())]
}

// perb returns the locale's abbreviated month name.
func (lc *localeData) perb(t time.Time) string {
	return lc.ShortMonths[int(t.Month())-1]
}

// perB returns the locale's full month name.
func (lc *localeData) perB(t time.Time) string {
	return lc.Months[int(t.Month())-1]
}

// perc returns the locale's appropriate date and time representation.
func (lc *localeData) perc(t time.Time) string {
	return lc.Strftime(lc.DateTime, t)
}

// perC returns the year divided by 100 and truncated to an integer, as a
// decimal number.
func (lc *localeData) perC(t time.Time) string {
	return fmt.Sprint(t.Year() / 100)
}

// perd returns the day of the month as a decimal number [01,31].
func (lc *localeData) perd(t time.Time) string {
	return fmt.Sprintf("%02d", t.Day())
}

// perD returns the date formatted as %m/%d/%y.
func (lc *localeData) perD(t time.Time) string {
	return lc.Strftime("%m/%d/%y", t)
}

// pere returns the day of the month as a decimal number [1,31]; a single digit
// is preceded by a space.
func (lc *localeData) pere(t time.Time) string {
	d := t.Day()
	if d < 10 {
		return fmt.Sprintf(" %d", d)
	}
	return fmt.Sprintf("%d", d)
}

// perF returns the date formatted as %Y-%m-%d.
func (lc *localeData) perF(t time.Time) string {
	return lc.Strftime("%Y-%m-%d", t)
}

// perg returns the last 2 digits of the week-based year as a decimal number
// [00,99].
func (lc *localeData) perg(t time.Time) string {
	y, _ := t.ISOWeek()
	return fmt.Sprintf("%02d", y%100)
}

// perG returns the week-based year as a decimal number (for example, 1977).
func (lc *localeData) perG(t time.Time) string {
	y, _ := t.ISOWeek()
	return fmt.Sprintf("%d", y)
}

// perH returns the hour (24-hour clock) as a decimal number [00,23].
func (lc *localeData) perH(t time.Time) string {
	return fmt.Sprintf("%02d", t.Hour())
}

// perI returns the hour (12-hour clock) as a decimal number [01,12].
func (lc *localeData) perI(t time.Time) string {
	hr := t.Hour() % 12
	if hr == 0 {
		hr = 12
	}

	return fmt.Sprintf("%02d", hr)
}

// perj returns the day of the year as a decimal number [001,366].
func (lc *localeData) perj(t time.Time) string {
	return fmt.Sprintf("%03d", t.YearDay())
}

// perm returns the month as a decimal number [01,12].
func (lc *localeData) perm(t time.Time) string {
	return fmt.Sprintf("%02d", t.Month())
}

// perM returns the minute as a decimal number [00,59].
func (lc *localeData) perM(t time.Time) string {
	return fmt.Sprintf("%02d", t.Minute())
}

// pern returns a newline.
func (lc *localeData) pern(t time.Time) string {
	return "\n"
}

// perp returns the locale's equivalent of either a.m. or p.m.
func (lc *localeData) perp(t time.Time) string {
	if t.Hour() < 12 {
		return lc.AMPM[0]
	}
	return lc.AMPM[1]
}

// perr returns the time in a.m. and p.m. notation.
func (lc *localeData) perr(t time.Time) string {
	return lc.Strftime(lc.TimeAMPM, t)
}

// perR returns the time formatted as %H:%M.
func (lc *localeData) perR(t time.Time) string {
	return lc.Strftime("%H:%M", t)
}

// perS returns the second as a decimal number [00,60].
func (lc *localeData) perS(t time.Time) string {
	return fmt.Sprintf("%02d", t.Second())
}

// pert returns a tab.
func (lc *localeData) pert(t time.Time) string {
	return "\t"
}

// perT returns the time formatted as %H:%M:%S
func (lc *localeData) perT(t time.Time) string {
	return lc.Strftime("%H:%M:%S", t)
}

// peru returns the weekday as a decimal number [1,7], with 1 representing
// Monday.
func (lc *localeData) peru(t time.Time) string {
	d := int(t.Weekday())
	if d == 0 {
		return "7"
	}

	return fmt.Sprint(d)
}

// perU returns the week number of the year as a decimal number [00,53]. The
// first Sunday of January is the first day of week 1; days in the new year
// before this are in week 0.
func (lc *localeData) perU(t time.Time) string {
	_, wn := t.ISOWeek()
	return fmt.Sprintf("%02d", wn)
}

// perV returns the week number of the year (Monday as the first day of the
// week) as a decimal number [01,53]. If the week containing 1 January has four
// or more days in the new year, then it is considered week 1. Otherwise, it is
// the last week of the previous year, and the next week is week 1. Both January
// 4th and the first Thursday of January are always in week 1.
func (lc *localeData) perV(t time.Time) string {
	_, wn := t.ISOWeek()
	return fmt.Sprintf("%02d", wn)
}

// perw returns the weekday as a decimal number [0,6], with 0 representing
// Sunday.
func (lc *localeData) perw(t time.Time) string {
	return fmt.Sprintf("%d", t.Weekday())
}

// perW returns the week number of the year as a decimal number [00,53]. The
// first Monday of January is the first day of week 1; days in the new year
// before this are in week 0.
func (lc *localeData) perW(t time.Time) string {
	_, wn := t.ISOWeek()
	return fmt.Sprintf("%02d", wn-1)
}

// perx returns the locale's appropriate date representation.
func (lc *localeData) perx(t time.Time) string {
	return lc.Strftime(lc.Date, t)
}

// perX returns the locale's appropriate time representation.
func (lc *localeData) perX(t time.Time) string {
	return lc.Strftime(lc.Time, t)
}

// pery returns the last two digits of the year as a decimal number [00,99].
func (lc *localeData) pery(t time.Time) string {
	return fmt.Sprintf("%02d", t.Year()%100)
}

// perY returns the year as a decimal number (for example, 1997).
func (lc *localeData) perY(t time.Time) string {
	return fmt.Sprint(t.Year())
}

// perz returns the offset from UTC in the ISO 8601:2000 standard format ( +hhmm
// or -hhmm ), or by no characters if no timezone is determinable. For example,
// "-0430" means 4 hours 30 minutes behind UTC (west of Greenwich). If tm_isdst
// is zero, the standard time offset is used. If tm_isdst is greater than zero,
// the daylight savings time offset is used. If tm_isdst is negative, no
// characters are returned.
func (lc *localeData) perz(t time.Time) string {
	_, off := t.Zone()
	return fmt.Sprintf("%+05d", off)
}

// perZ returns the timezone name or abbreviation, or by no bytes if no timezone
// information exists.
func (lc *localeData) perZ(t time.Time) string {
	tz, _ := t.Zone()
	return tz
}

// perper returns a %.
func (lc *localeData) perper(t time.Time) string {
	return "%"
}
