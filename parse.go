package lctime

import (
	"bytes"
	"time"
)

// Strftime formats a time.Time. It's locale-aware, so make sure you call
// SetLocale if needed.
func Strftime(format string, t time.Time) string {
	if len(format) < 1 {
		return format
	}

	buf := new(bytes.Buffer)
	end := len(format)

	for i := 0; i < end; i++ {
		if format[i] == '%' && i+2 <= end {
			buf.WriteString(parseDirective(format[i:i+2], t))
			i++
			continue
		}

		buf.WriteByte(format[i])
	}

	return buf.String()
}

func parseDirective(direc string, t time.Time) string {
	if len(direc) < 2 {
		return direc
	}

	switch direc[:2] {
	case "%a":
		return pera(t)
	case "%A":
		return perA(t)
	case "%b":
		return perb(t)
	case "%B":
		return perB(t)
	case "%c":
		return perc(t)
	case "%C":
		return perC(t)
	case "%d":
		return perd(t)
	case "%D":
		return perD(t)
	case "%e":
		return pere(t)
	case "%F":
		return perF(t)
	case "%g":
		return perg(t)
	case "%G":
		return perG(t)
	case "%H":
		return perH(t)
	case "%I":
		return perI(t)
	case "%j":
		return perj(t)
	case "%m":
		return perm(t)
	case "%M":
		return perM(t)
	case "%n":
		return pern(t)
	case "%p":
		return perp(t)
	case "%r":
		return perr(t)
	case "%R":
		return perR(t)
	case "%S":
		return perS(t)
	case "%t":
		return pert(t)
	case "%T":
		return perT(t)
	case "%u":
		return peru(t)
	case "%U":
		return perU(t)
	case "%V":
		return perV(t)
	case "%w":
		return perw(t)
	case "%W":
		return perW(t)
	case "%x":
		return perx(t)
	case "%X":
		return perX(t)
	case "%y":
		return pery(t)
	case "%Y":
		return perY(t)
	case "%z":
		return perz(t)
	case "%Z":
		return perZ(t)
	case "%%":
		return perper(t)
	}

	return direc
}
