/*
Package lctime provides a way to format dates and times using strftime
directives. More importantly, it does so in a locale-aware fashion. This
allows developers to format time based on a user's locale.

An initial locale is loaded at import time. It's determined by the first,
non-empty locale identifier from these environment variables.

    1. LC_TIME
    2. LC_ALL
    3. LANG

If all of the previous variables are empty, then POSIX will be the initial
locale used. All locales use UTF-8 character encoding.

The formats used are loosely based on glibc locale files.

These are the supported strftime directives. They're loosely based on The Open
Group Base Specifications Issue 7,
http://pubs.opengroup.org/onlinepubs/9699919799/functions/strftime.html.

   %a  locale's abbreviated weekday name
   %A  locale's full weekday name
   %b  locale's abbreviated month name
   %B  locale's full month name
   %c  locale's appropriate date and time representation
   %C  year divided by 100 and truncated to an integer
   %d  day of the month as a decimal number [01,31]
   %D  %m/%d/%y
   %e  day of the month as a decimal number [1,31]
   %F  %Y/%m/%d
   %g  last 2 digits of the week-based year as a decimal number
   %G  week-based year as a decimal number (for example, 1977)
   %H  hour (24-hour clock) as a decimal number [00,23]
   %I  hour (12-hour clock) as a decimal number [01,12]
   %j  day of the year as a decimal number [001,366]
   %m  month as a decimal number [01,12]
   %M  minute as a decimal number [00,59]
   %n  returns a newline
   %p  locale's equivalent of either a.m. or p.m.
   %r  time in a.m. and p.m. notation.
   %R  time in 24-hour notation %H:%M
   %S  second as a decimal number [00,60]
   %t  returns a tab
   %T  %H:%M:%S
   %u  weekday as a decimal number [1,7]
   %U  week number of the year as a decimal number [00,53]
   %V  week number of the year
   %w  weekday as a decimal number [0,6]
   %W  week number of the year as a decimal number [00,53]
   %x  locale's appropriate date representation
   %X  locale's appropriate time representation
   %y  last two digits of the year as a decimal number [00,99]
   %Y  year as a decimal number (for example, 1997)
   %z  offset from UTC in the ISO 8601:2000 standard format
   %Z  timezone name or abbreviation
   %%  %
*/
package lctime

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/klauspost/lctime/internal/locales"
)

// Localizer provides translation to a locale.
type Localizer interface {
	Strftime(format string, t time.Time) string
}

type localeData struct {
	ID string

	Days        []string
	ShortDays   []string
	Months      []string
	ShortMonths []string
	AMPM        []string

	Date     string
	DateTime string
	Time     string
	TimeAMPM string
}

var (
	// ErrNoLocale is returned when a given locale was not found.
	ErrNoLocale = errors.New("Locale not found")
	// ErrCorruptLocale is returned if a given locale file is corrupted.
	ErrCorruptLocale = errors.New("Corrupted locale")

	lc     localeData
	loaded sync.Map
)

func init() {
	vars := []string{
		os.Getenv("LC_TIME"),
		os.Getenv("LC_ALL"),
		os.Getenv("LANG"),
		"POSIX",
	}

	for _, v := range vars {
		if v != "" {
			SetLocale(v)
			break
		}
	}
}

// SetLocale activates the given locale.
func SetLocale(id string) error {
	l, err := loadLocale(id)
	if err != nil {
		lc = localeData{}
		return err
	}
	lc = *l
	return nil
}

// NewLocalizer provides a localizer to a specific locale.
func NewLocalizer(id string) (Localizer, error) {
	return loadLocale(id)
}

// loadLocale will load the locale or fetch it from cache.
func loadLocale(id string) (*localeData, error) {
	id = removeCodeset(id)
	if l, ok := loaded.Load(id); ok {
		lc, ok := l.(*localeData)
		if ok {
			return lc, nil
		}
	}

	bys, err := locale.Asset(id + ".json")
	if err != nil {
		return nil, ErrNoLocale
	}

	// All locales use UTF-8.
	var lc localeData
	if err := json.Unmarshal(bys, &lc); err != nil {
		return nil, ErrCorruptLocale
	}
	
	loaded.Store(id, &lc)
	return &lc, nil
}

// GetLocale returns the currently active locale.
func GetLocale() string {
	return lc.ID
}

// GetLocales returns a slice of available locales.
func GetLocales() []string {
	lcs := locale.AssetNames()

	names := make([]string, 0, len(lcs))
	for _, l := range lcs {
		names = append(names, strings.TrimSuffix(l, filepath.Ext(l)))
	}

	return names
}

// cleanID takes in a valid locale ID and returns the same ID, except without
// the codeset. A valid locale ID is defined like this:
// [language[_territory][.codeset][@modifier]]
func removeCodeset(id string) string {
	dot, at := -1, -1

	for i := 0; i < len(id); i++ {
		if id[i] == '.' {
			dot = i
		} else if id[i] == '@' {
			at = i
		}
	}

	if dot == at || dot == -1 {
		return id
	}

	if at == -1 {
		return id[:dot]
	}

	return id[:dot] + id[at:]
}
