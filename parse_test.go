package lctime

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDirective(t *testing.T) {
	SetLocale("en_US")
	dt := time.Date(2024, 2, 22, 0, 36, 21, 795187684, time.UTC)

	tests := []struct {
		input string
		want  string
	}{
		{"%a", pera(dt)},
		{"%A", perA(dt)},
		{"%b", perb(dt)},
		{"%B", perB(dt)},
		{"%c", perc(dt)},
		{"%C", perC(dt)},
		{"%d", perd(dt)},
		{"%D", perD(dt)},
		{"%e", pere(dt)},
		{"%F", perF(dt)},
		{"%g", perg(dt)},
		{"%G", perG(dt)},
		{"%H", perH(dt)},
		{"%I", perI(dt)},
		{"%j", perj(dt)},
		{"%m", perm(dt)},
		{"%M", perM(dt)},
		{"%n", pern(dt)},
		{"%p", perp(dt)},
		{"%r", perr(dt)},
		{"%R", perR(dt)},
		{"%S", perS(dt)},
		{"%t", pert(dt)},
		{"%T", perT(dt)},
		{"%u", peru(dt)},
		{"%U", perU(dt)},
		{"%V", perV(dt)},
		{"%w", perw(dt)},
		{"%W", perW(dt)},
		{"%x", perx(dt)},
		{"%X", perX(dt)},
		{"%y", pery(dt)},
		{"%Y", perY(dt)},
		{"%z", perz(dt)},
		{"%Z", perZ(dt)},
		{"%%", perper(dt)},
		{"--", "--"},
		{"", ""},
	}

	for i, test := range tests {
		if got := parseDirective(test.input, dt); got != test.want {
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
	t := time.Date(2000, 1, 2, 3, 4, 5, 6, time.UTC)
	fmt.Println(Strftime("%Y-%m-%d %H:%M:%S", t))
	// Output: 2000-01-02 03:04:05
}

func ExampleStrftime_withLocale() {
	SetLocale("es_MX")
	t := time.Date(2000, 1, 2, 3, 4, 5, 6, time.UTC)
	fmt.Println(Strftime("%d de %B de %Y", t))
	// Output: 02 de enero de 2000
}
