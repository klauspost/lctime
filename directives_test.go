package lctime

import (
	"testing"
	"time"
)

func TestPera(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1997, 5, 3, 1, 42, 58, 856796, time.UTC),
			"en_US", "Sat"},
		{time.Date(2037, 5, 26, 7, 51, 33, 975892, time.UTC),
			"es_MX", "mar"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := pera(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerA(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1988, 3, 8, 10, 2, 14, 280313, time.UTC),
			"en_US", "Tuesday"},
		{time.Date(1980, 3, 28, 17, 49, 9, 914495, time.UTC),
			"es_MX", "viernes"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perA(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerb(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1962, 9, 25, 18, 45, 19, 195633, time.UTC),
			"en_US", "Sep"},
		{time.Date(1833, 7, 17, 11, 38, 3, 809041, time.UTC),
			"es_MX", "jul"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perb(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerB(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(2051, 8, 26, 6, 47, 53, 788607, time.UTC),
			"en_US", "August"},
		{time.Date(2069, 7, 11, 9, 39, 12, 27157, time.UTC),
			"es_MX", "julio"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perB(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerc(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(2074, 10, 3, 20, 9, 2, 797808, time.UTC),
			"en_US", "Wed 03 Oct 2074 08:09:02 PM UTC"},
		{time.Date(1859, 10, 11, 21, 57, 51, 269055, time.UTC),
			"es_MX", "mar 11 oct 1859 21:57:51 UTC"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perc(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerC(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1858, 9, 15, 21, 13, 21, 930564, time.UTC), "18"},
	}

	for i, test := range tests {
		if got := perC(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerd(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2018, 4, 30, 17, 58, 33, 652858, time.UTC), "30"},
	}

	for i, test := range tests {
		if got := perd(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerD(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2013, 6, 1, 7, 39, 27, 218412, time.UTC), "06/01/13"},
	}

	for i, test := range tests {
		if got := perD(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPere(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1846, 9, 4, 22, 16, 5, 108059, time.UTC), " 4"},
		{time.Date(2045, 11, 23, 8, 14, 34, 971351, time.UTC), "23"},
	}

	for i, test := range tests {
		if got := pere(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerF(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1873, 9, 20, 1, 56, 2, 554999, time.UTC), "1873-09-20"},
	}

	for i, test := range tests {
		if got := perF(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerg(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1805, 11, 25, 6, 54, 12, 611641, time.UTC), "05"},
	}

	for i, test := range tests {
		if got := perg(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerG(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2052, 4, 23, 3, 48, 19, 290262, time.UTC), "2052"},
	}

	for i, test := range tests {
		if got := perG(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerH(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2098, 5, 20, 20, 38, 16, 790937, time.UTC), "20"},
	}

	for i, test := range tests {
		if got := perH(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerI(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1847, 9, 20, 0, 42, 43, 49440, time.UTC), "12"},
	}

	for i, test := range tests {
		if got := perI(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerj(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2000, 11, 19, 3, 37, 54, 446382, time.UTC), "324"},
	}

	for i, test := range tests {
		if got := perj(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerm(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1894, 8, 24, 21, 37, 21, 802850, time.UTC), "08"},
	}

	for i, test := range tests {
		if got := perm(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerM(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1979, 6, 20, 10, 8, 33, 349343, time.UTC), "08"},
	}

	for i, test := range tests {
		if got := perM(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPern(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1867, 3, 11, 5, 33, 29, 668689, time.UTC), "\n"},
	}

	for i, test := range tests {
		if got := pern(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerp(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1899, 2, 25, 20, 48, 58, 389229, time.UTC),
			"en_US", "PM"},
		{time.Date(1866, 6, 24, 7, 46, 55, 436140, time.UTC),
			"es_MX", ""},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perp(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerr(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1956, 6, 16, 3, 20, 31, 254473, time.UTC),
			"en_US", "03:20:31 AM"},
		{time.Date(2096, 11, 25, 22, 53, 17, 105946, time.UTC),
			"es_MX", ""},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perr(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerR(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1866, 7, 10, 18, 2, 16, 179906, time.UTC), "18:02"},
	}

	for i, test := range tests {
		if got := perR(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerS(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1901, 7, 18, 0, 46, 6, 900202, time.UTC), "06"},
	}

	for i, test := range tests {
		if got := perS(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPert(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1983, 1, 17, 4, 43, 51, 827019, time.UTC), "\t"},
	}

	for i, test := range tests {
		if got := pert(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerT(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2009, 7, 2, 1, 0, 49, 252443, time.UTC), "01:00:49"},
	}

	for i, test := range tests {
		if got := perT(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPeru(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2025, 5, 5, 8, 1, 29, 632563, time.UTC), "1"},
		{time.Date(1960, 9, 4, 3, 6, 47, 547101, time.UTC), "7"},
	}

	for i, test := range tests {
		if got := peru(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerU(t *testing.T) {
	t.SkipNow()

	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1949, 8, 16, 21, 54, 30, 113731, time.UTC),
			"en_US", "33"},
		{time.Date(1954, 11, 29, 0, 33, 29, 550957, time.UTC),
			"es_MX", "48"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perU(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerV(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1969, 10, 24, 1, 2, 57, 31564, time.UTC), "43"},
	}

	for i, test := range tests {
		if got := perV(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerw(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1920, 1, 1, 9, 33, 48, 124153, time.UTC), "4"},
	}

	for i, test := range tests {
		if got := perw(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerW(t *testing.T) {
	t.SkipNow()

	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(2081, 9, 29, 0, 22, 12, 982588, time.UTC),
			"en_US", "39"},
		{time.Date(1824, 1, 8, 4, 6, 57, 252617, time.UTC),
			"es_MX", "01"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perW(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerx(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(2004, 7, 1, 15, 14, 56, 299181, time.UTC),
			"en_US", "07/01/2004"},
		{time.Date(1860, 4, 25, 8, 51, 49, 613045, time.UTC),
			"es_MX", "25/04/60"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perx(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerX(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(1889, 6, 28, 17, 52, 27, 460104, time.UTC),
			"en_US", "05:52:27 PM"},
		{time.Date(1977, 7, 18, 18, 44, 4, 794184, time.UTC),
			"es_MX", "18:44:04"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perX(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPery(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1841, 9, 9, 3, 10, 9, 885689, time.UTC), "41"},
	}

	for i, test := range tests {
		if got := pery(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerY(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2047, 8, 23, 15, 23, 11, 369518, time.UTC), "2047"},
	}

	for i, test := range tests {
		if got := perY(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerz(t *testing.T) {
	tests := []struct {
		input  time.Time
		locale string
		want   string
	}{
		{time.Date(2065, 6, 19, 5, 29, 39, 858124, time.UTC),
			"en_US", "+0000"},
		{time.Date(1827, 2, 10, 1, 53, 54, 444622, time.UTC),
			"es_MX", "+0000"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := perz(test.input); got != test.want {
			t.Error("locale:", test.locale)
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerZ(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1961, 3, 10, 11, 18, 58, 211914, time.UTC), "UTC"},
	}

	for i, test := range tests {
		if got := perZ(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerper(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1835, 4, 5, 11, 9, 18, 795572, time.UTC), "%"},
	}

	for i, test := range tests {
		if got := perper(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}
