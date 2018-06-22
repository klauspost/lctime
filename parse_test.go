package lctime

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDirective(t *testing.T) {
	lc, err := loadLocale("en_US")
	if err != nil {
		t.Fatal(err)
	}
	dt := time.Date(2024, 2, 22, 0, 36, 21, 795187684, time.UTC)

	tests := []struct {
		input string
		want  string
	}{
		{"%a", lc.pera(dt)},
		{"%A", lc.perA(dt)},
		{"%b", lc.perb(dt)},
		{"%B", lc.perB(dt)},
		{"%c", lc.perc(dt)},
		{"%C", lc.perC(dt)},
		{"%d", lc.perd(dt)},
		{"%D", lc.perD(dt)},
		{"%e", lc.pere(dt)},
		{"%F", lc.perF(dt)},
		{"%g", lc.perg(dt)},
		{"%G", lc.perG(dt)},
		{"%H", lc.perH(dt)},
		{"%I", lc.perI(dt)},
		{"%j", lc.perj(dt)},
		{"%m", lc.perm(dt)},
		{"%M", lc.perM(dt)},
		{"%n", lc.pern(dt)},
		{"%p", lc.perp(dt)},
		{"%r", lc.perr(dt)},
		{"%R", lc.perR(dt)},
		{"%S", lc.perS(dt)},
		{"%t", lc.pert(dt)},
		{"%T", lc.perT(dt)},
		{"%u", lc.peru(dt)},
		{"%U", lc.perU(dt)},
		{"%V", lc.perV(dt)},
		{"%w", lc.perw(dt)},
		{"%W", lc.perW(dt)},
		{"%x", lc.perx(dt)},
		{"%X", lc.perX(dt)},
		{"%y", lc.pery(dt)},
		{"%Y", lc.perY(dt)},
		{"%z", lc.perz(dt)},
		{"%Z", lc.perZ(dt)},
		{"%%", lc.perper(dt)},
		{"--", "--"},
		{"", ""},
	}

	for i, test := range tests {
		if got := lc.parseDirective(test.input, dt); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestStrftime(t *testing.T) {
	SetLocale("en_US")
	dt := time.Date(2444, 3, 8, 3, 8, 59, 284117260, time.UTC)

	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"%a", "Tue"},
		{"%a %b", "Tue Mar"},
		{"%A %B %Y", "Tuesday March 2444"},
		{"abc", "abc"},
		{"%", "%"},
		{"%%", "%"},
		{"%FT%T%z", "2444-03-08T03:08:59+0000"},
	}

	for i, test := range tests {
		if got := Strftime(test.input, dt); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func ExampleStrftime() {
	// Initial locale based on env vars. If not set, then POSIX is used.
	t := time.Date(2000, 1, 2, 3, 4, 5, 6, time.UTC)
	fmt.Println(Strftime("%c", t))
	// Output: Sun 02 Jan 2000 03:04:05 AM UTC
}

func ExampleStrftimeLoc() {
	t := time.Date(2015, 12, 25, 3, 2, 1, 0, time.UTC)
	txt, err := StrftimeLoc("es_MX", "%A, %d de %B de %Y", t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txt)
	// Output: viernes, 25 de diciembre de 2015
}

func ExampleLocalizer() {
	t := time.Date(2015, 12, 25, 3, 2, 1, 0, time.UTC)
	l, err := NewLocalizer("da_DK")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l.Strftime("%A den %d. %B %Y", t))
	// Output: fredag den 25. december 2015
}
