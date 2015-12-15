// The following directives are based on The Open Group Base Specifications
// Issue 7 strftime man page.
// http://pubs.opengroup.org/onlinepubs/9699919799/functions/strftime.html

package lctime

import (
	"fmt"
	"time"
)

// pera returns the locale's abbreviated weekday name.
func pera(t time.Time) string {
	return localeData["abday"][int(t.Weekday())]
}

// perA returns the locale's full weekday name.
func perA(t time.Time) string {
	return localeData["day"][int(t.Weekday())]
}

// perb returns the locale's abbreviated month name.
func perb(t time.Time) string {
	return localeData["abmon"][int(t.Month())-1]
}

// perB returns the locale's full month name.
func perB(t time.Time) string {
	return localeData["mon"][int(t.Month())-1]
}

// perc returns the locale's appropriate date and time representation.
func perc(t time.Time) string {
	return Strftime(localeData["d_t_fmt"][0], t)
}

// perC returns the year divided by 100 and truncated to an integer, as a
// decimal number.
//
// If a minimum field width is not specified, the number of characters placed
// into the array pointed to by s will be the number of digits in the year
// divided by 100 or two, whichever is greater. If a minimum field width is
// specified, the number of characters placed into the array pointed to by s
// will be the number of digits in the year divided by 100 or the minimum field
// width, whichever is greater.
func perC(t time.Time) string {
	return fmt.Sprint(t.Year() / 100)
}

// perd returns the day of the month as a decimal number [01,31].
func perd(t time.Time) string {
	return fmt.Sprintf("%02d", t.Day())
}

// perD returns %m/%d/%y.
func perD(t time.Time) string {
	return Strftime("%m/%d/%y", t)
}

// pere returns the day of the month as a decimal number [1,31]; a single digit
// is preceded by a space.
func pere(t time.Time) string {
	d := t.Day()
	if d < 10 {
		return fmt.Sprintf(" %d", d)
	}
	return fmt.Sprintf("%d", d)
}

// perF returns %+4Y-%m-%d if no flag and no minimum field width are specified.
//
// If a minimum field width of x is specified, the year shall be output as if by
// the Y specifier with whatever flag was given and a minimum field width of
// x-6. If x is less than 6, the behavior shall be as if x equalled 6.
//
// If the minimum field width is specified to be 10, and the year is four digits
// long, then the output string produced will match the ISO 8601:2000 standard
// subclause 4.1.2.2 complete representation, extended format date
// representation of a specific day. If a + flag is specified, a minimum field
// width of x is specified, and x-7 bytes are sufficient to hold the digits of
// the year (not including any needed sign character), then the output will
// match the ISO 8601:2000 standard subclause 4.1.2.4 complete representation,
// expanded format date representation of a specific day.
func perF(t time.Time) string {
	// return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	// BUG(jaime): %F not implemented
	return "-1"
}

// perg returns the last 2 digits of the week-based year as a decimal number
// [00,99].
func perg(t time.Time) string {
	y, _ := t.ISOWeek()
	return fmt.Sprintf("%02d", y%100)
}

// perG returns the week-based year as a decimal number (for example, 1977).
func perG(t time.Time) string {
	y, _ := t.ISOWeek()
	return fmt.Sprintf("%d", y)
}

// perH returns the hour (24-hour clock) as a decimal number [00,23].
func perH(t time.Time) string {
	return fmt.Sprintf("%02d", t.Hour())
}

// perI returns the hour (12-hour clock) as a decimal number [01,12].
func perI(t time.Time) string {
	hr := t.Hour() % 12
	if hr == 0 {
		hr = 12
	}

	return fmt.Sprintf("%02d", hr)
}

// perj returns the day of the year as a decimal number [001,366].
func perj(t time.Time) string {
	return fmt.Sprintf("%03d", t.YearDay())
}

// perm returns the month as a decimal number [01,12].
func perm(t time.Time) string {
	return fmt.Sprintf("%02d", t.Month())
}

// perM returns the minute as a decimal number [00,59].
func perM(t time.Time) string {
	return fmt.Sprintf("%02d", t.Minute())
}

// pern returns a newline.
func pern(t time.Time) string {
	return "\n"
}

// perp returns the locale's equivalent of either a.m. or p.m.
func perp(t time.Time) string {
	if t.Hour() < 12 {
		return localeData["am_pm"][0]
	}
	return localeData["am_pm"][1]
}

// perr returns the time in a.m. and p.m. notation. In POSIX locale this shall
// be equivalent to %I : %M : %S %p.
func perr(t time.Time) string {
	return Strftime(localeData["t_fmt_ampm"][0], t)
}

// perR returns the time in 24-hour notation ( %H : %M ).
func perR(t time.Time) string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

// perS returns the second as a decimal number [00,60].
func perS(t time.Time) string {
	return fmt.Sprintf("%02d", t.Second())
}

// pert returns a tab.
func pert(t time.Time) string {
	return "\t"
}

// perT returns the time ( %H : %M : %S ).
func perT(t time.Time) string {
	return Strftime("%H:%M:%S", t)
}

// peru returns the weekday as a decimal number [1,7], with 1 representing
// Monday.
func peru(t time.Time) string {
	d := int(t.Weekday())
	if d == 0 {
		return "7"
	}

	return fmt.Sprint(d)
}

// perU returns the week number of the year as a decimal number [00,53]. The
// first Sunday of January is the first day of week 1; days in the new year
// before this are in week 0.
func perU(t time.Time) string {
	// BUG(jaime): %U not implemented
	return "-1"
}

// perV returns the week number of the year (Monday as the first day of the
// week) as a decimal number [01,53]. If the week containing 1 January has four
// or more days in the new year, then it is considered week 1. Otherwise, it is
// the last week of the previous year, and the next week is week 1. Both January
// 4th and the first Thursday of January are always in week 1.
func perV(t time.Time) string {
	_, wn := t.ISOWeek()
	return fmt.Sprintf("%02d", wn)
}

// perw returns the weekday as a decimal number [0,6], with 0 representing
// Sunday.
func perw(t time.Time) string {
	return fmt.Sprintf("%d", t.Weekday())
}

// perW returns the week number of the year as a decimal number [00,53]. The
// first Monday of January is the first day of week 1; days in the new year
// before this are in week 0.
func perW(t time.Time) string {
	// BUG(jaime): %W not implemented
	return "-1"
}

// perx returns the locale's appropriate date representation.
func perx(t time.Time) string {
	return Strftime(localeData["d_fmt"][0], t)
}

// perX returns the locale's appropriate time representation.
func perX(t time.Time) string {
	return Strftime(localeData["t_fmt"][0], t)
}

// pery returns the last two digits of the year as a decimal number [00,99].
func pery(t time.Time) string {
	return fmt.Sprintf("%02d", t.Year()%100)
}

// perY returns the year as a decimal number (for example, 1997).
func perY(t time.Time) string {
	return fmt.Sprint(t.Year())
}

// perz returns the offset from UTC in the ISO 8601:2000 standard format ( +hhmm
// or -hhmm ), or by no characters if no timezone is determinable. For example,
// "-0430" means 4 hours 30 minutes behind UTC (west of Greenwich). If tm_isdst
// is zero, the standard time offset is used. If tm_isdst is greater than zero,
// the daylight savings time offset is used. If tm_isdst is negative, no
// characters are returned.
func perz(t time.Time) string {
	_, off := t.Zone()
	return fmt.Sprintf("%+05d", off)
}

// perZ returns the timezone name or abbreviation, or by no bytes if no timezone
// information exists.
func perZ(t time.Time) string {
	tz, _ := t.Zone()
	return tz
}

// perper returns a %.
func perper(t time.Time) string {
	return "%"
}
