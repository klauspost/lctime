package lctime

import (
	"bytes"
	"time"
)

// Strftime formats a time.Time. It's locale-aware, so make sure you call
// SetLocale if needed.
func Strftime(format string, t time.Time) string {
	return lc.Strftime(format, t)
}

// StrftimeLoc formats a time.Time. It's locale-aware, so make sure you call
// SetLocale if needed.
func StrftimeLoc(locale, format string, t time.Time) (string, error) {
	lc, err := loadLocale(locale)
	if err != nil {
		return "", err
	}
	return lc.Strftime(format, t), nil
}

func (lc *localeData) Strftime(format string, t time.Time) string {
	if len(format) < 1 {
		return format
	}

	buf := new(bytes.Buffer)
	end := len(format)

	for i := 0; i < end; i++ {
		if format[i] == '%' && i+2 <= end {
			buf.WriteString(lc.parseDirective(format[i:i+2], t))
			i++
			continue
		}

		buf.WriteByte(format[i])
	}

	return buf.String()
}

func (lc *localeData) parseDirective(direc string, t time.Time) string {
	if len(direc) < 2 {
		return direc
	}

	switch direc[:2] {
	case "%a":
		return lc.pera(t)
	case "%A":
		return lc.perA(t)
	case "%b":
		return lc.perb(t)
	case "%B":
		return lc.perB(t)
	case "%c":
		return lc.perc(t)
	case "%C":
		return lc.perC(t)
	case "%d":
		return lc.perd(t)
	case "%D":
		return lc.perD(t)
	case "%e":
		return lc.pere(t)
	case "%F":
		return lc.perF(t)
	case "%g":
		return lc.perg(t)
	case "%G":
		return lc.perG(t)
	case "%H":
		return lc.perH(t)
	case "%I":
		return lc.perI(t)
	case "%j":
		return lc.perj(t)
	case "%m":
		return lc.perm(t)
	case "%M":
		return lc.perM(t)
	case "%n":
		return lc.pern(t)
	case "%p":
		return lc.perp(t)
	case "%r":
		return lc.perr(t)
	case "%R":
		return lc.perR(t)
	case "%S":
		return lc.perS(t)
	case "%t":
		return lc.pert(t)
	case "%T":
		return lc.perT(t)
	case "%u":
		return lc.peru(t)
	case "%U":
		return lc.perU(t)
	case "%V":
		return lc.perV(t)
	case "%w":
		return lc.perw(t)
	case "%W":
		return lc.perW(t)
	case "%x":
		return lc.perx(t)
	case "%X":
		return lc.perX(t)
	case "%y":
		return lc.pery(t)
	case "%Y":
		return lc.perY(t)
	case "%z":
		return lc.perz(t)
	case "%Z":
		return lc.perZ(t)
	case "%%":
		return lc.perper(t)
	}

	return direc
}
