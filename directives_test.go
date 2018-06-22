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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "fre"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "Чт."},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "Mi"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "মঙ্গল"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "五"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "Qua"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "ven."},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "ج"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.pera(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "fredag"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "Четверг"},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "Mittwoch"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "মঙ্গলবার"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "星期五"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "quarta"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "vendredi"},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "الجمعة"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perA(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "jan"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "дек."},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "Okt"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "অক্টোবর"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "11月"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "Set"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "sept."},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "نوف"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perb(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "januar"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "Декабрь"},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "Oktober"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "অক্টোবর"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "十一月"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "setembro"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "septembre"},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "نوفمبر"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perB(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "fre 07 jan 1994 09:51:27 UTC"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "Чт. 01 дек. 1988 17:51:22"},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "Mi 13 Okt 1965 08:31:35 UTC"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "মঙ্গলবার 16 অক্টোবর 1827 12:40:34 পূর্বাহ্ণ UTC"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "1849年11月23日 星期五 15时11分46秒"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "Qua 07 Set 1938 04:06:42 UTC"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "ven. 20 sept. 2024 23:33:13 UTC"},
		//{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
		//	"ar_EG", "01 نوف, 1805  11:12:14 ص UTC"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perc(test.input); got != test.want {
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
		if got := lc.perC(test.input); got != test.want {
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
		if got := lc.perd(test.input); got != test.want {
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
		if got := lc.perD(test.input); got != test.want {
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
		if got := lc.pere(test.input); got != test.want {
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
		if got := lc.perF(test.input); got != test.want {
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
		{time.Date(1800, 11, 3, 21, 30, 57, 427721, time.UTC), "00"},
	}

	for i, test := range tests {
		if got := lc.perg(test.input); got != test.want {
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
		{time.Date(1907, 4, 15, 21, 47, 4, 910983, time.UTC), "1907"},
	}

	for i, test := range tests {
		if got := lc.perG(test.input); got != test.want {
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
		if got := lc.perH(test.input); got != test.want {
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
		{time.Date(1941, 3, 11, 19, 6, 28, 673085, time.UTC), "07"},
	}

	for i, test := range tests {
		if got := lc.perI(test.input); got != test.want {
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
		if got := lc.perj(test.input); got != test.want {
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
		if got := lc.perm(test.input); got != test.want {
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
		if got := lc.perM(test.input); got != test.want {
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
		if got := lc.pern(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", ""},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", ""},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", ""},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "পূর্বাহ্ণ"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "下午"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", ""},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", ""},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "ص"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perp(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", ""},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", ""},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", ""},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "12:40:34 পূর্বাহ্ণ UTC"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "下午 03时11分46秒"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", ""},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", ""},
		//{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
		//	"ar_EG", "11:12:14 ص"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perr(test.input); got != test.want {
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
		{time.Date(1981, 9, 17, 3, 2, 12, 378798, time.UTC), "03:02"},
	}

	for i, test := range tests {
		if got := lc.perR(test.input); got != test.want {
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
		if got := lc.perS(test.input); got != test.want {
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
		if got := lc.pert(test.input); got != test.want {
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
		if got := lc.perT(test.input); got != test.want {
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
		{time.Date(1882, 6, 16, 11, 36, 31, 317944, time.UTC), "5"},
		{time.Date(2081, 6, 11, 0, 40, 2, 298881, time.UTC), "3"},
	}

	for i, test := range tests {
		if got := lc.peru(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerU(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(1949, 8, 16, 21, 54, 30, 113731, time.UTC), "33"},
		{time.Date(1956, 11, 29, 0, 33, 29, 550957, time.UTC), "48"},
		{time.Date(1954, 1, 22, 9, 33, 5, 730246, time.UTC), "03"},
	}

	for i, test := range tests {
		if got := lc.perU(test.input); got != test.want {
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
		{time.Date(2052, 11, 6, 23, 58, 31, 538378, time.UTC), "45"},
		{time.Date(2089, 7, 12, 3, 38, 23, 593912, time.UTC), "28"},
	}

	for i, test := range tests {
		if got := lc.perV(test.input); got != test.want {
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
		if got := lc.perw(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerW(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2081, 9, 29, 0, 22, 12, 982588, time.UTC), "39"},
		{time.Date(1824, 1, 8, 4, 6, 57, 252617, time.UTC), "01"},
		{time.Date(1817, 10, 24, 3, 54, 45, 627187, time.UTC), "42"},
		{time.Date(1834, 8, 20, 13, 5, 25, 433542, time.UTC), "33"},
	}

	for i, test := range tests {
		if got := lc.perW(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "07-01-1994"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "01.12.1988"},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "13.10.1965"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "মঙ্গলবার 16 অক্টোবর 1827"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "1849年11月23日"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "07-09-1938"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "20/09/2024"},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "01 نوف, 1805"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perx(test.input); got != test.want {
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
		{time.Date(1994, 1, 7, 9, 51, 27, 222686, time.UTC),
			"da_DK", "09:51:27"},
		{time.Date(1988, 12, 1, 17, 51, 22, 853401, time.UTC),
			"ru_RU", "17:51:22"},
		{time.Date(1965, 10, 13, 8, 31, 35, 101666, time.UTC),
			"de_DE", "08:31:35"},
		{time.Date(1827, 10, 16, 0, 40, 34, 153567, time.UTC),
			"bn_IN", "12:40:34 UTC"},
		{time.Date(1849, 11, 23, 15, 11, 46, 889860, time.UTC),
			"zh_CN", "15时11分46秒"},
		{time.Date(1938, 9, 7, 4, 6, 42, 79287, time.UTC),
			"pt_BR", "04:06:42"},
		{time.Date(2024, 9, 20, 23, 33, 13, 961479, time.UTC),
			"fr_FR", "23:33:13"},
		{time.Date(1805, 11, 1, 11, 12, 14, 651205, time.UTC),
			"ar_EG", "UTC 11:12:14"},
	}

	for i, test := range tests {
		SetLocale(test.locale)

		if got := lc.perX(test.input); got != test.want {
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
		{time.Date(2000, 9, 2, 4, 6, 36, 41497, time.UTC), "00"},
	}

	for i, test := range tests {
		if got := lc.pery(test.input); got != test.want {
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
		if got := lc.perY(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestPerz(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2065, 6, 19, 5, 29, 39, 858124, time.UTC), "+0000"},
	}

	for i, test := range tests {
		if got := lc.perz(test.input); got != test.want {
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
		if got := lc.perZ(test.input); got != test.want {
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
		if got := lc.perper(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}
