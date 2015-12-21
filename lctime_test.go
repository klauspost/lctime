package lctime

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

const (
	gotWant    = "got '%v', want '%v'"
	gotWantKey = "%s: got '%v', want '%v'"
	gotWantIdx = "%d: got '%v', want '%v'"
)

func TestSetLocale(t *testing.T) {
	tests := []struct {
		input string
		want  error
	}{
		{"ab_CD", ErrNoLocale},
		{"fake", ErrNoLocale},
		{"POSIX", nil},
		{"en_US", nil},
		{"en_US.UTF-8", nil},
		{"gez_ET@abegede", nil},
		{"es_MX", nil},
		{"", ErrNoLocale},
		{"eu_FR@euro", nil},
	}

	for i, test := range tests {
		if got := SetLocale(test.input); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestGetLocale(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"fake", ""},
		{"POSIX", "POSIX"},
		{"gez_ET@abegede", "gez_ET@abegede"},
		{"ab_CD", ""},
		{"", ""},
		{"eu_FR@euro", "eu_FR@euro"},
		{"en_US", "en_US"},
		{"es_MX", "es_MX"},
	}

	for i, test := range tests {
		SetLocale(test.input)

		if got := GetLocale(); got != test.want {
			t.Errorf(gotWantIdx, i, got, test.want)
		}
	}
}

func TestGetLocales(t *testing.T) {
	want := make([]string, 0, 8)
	err := filepath.Walk("internal/locales", func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() || fi.Name()[0] == '.' {
			return nil
		}

		fiName := fi.Name()
		if filepath.Ext(fiName) != ".json" {
			return nil
		}

		want = append(want, strings.TrimSuffix(fiName, filepath.Ext(fiName)))
		return nil
	})
	if err != nil {
		t.Fatal("Locale walk:", err)
	}

	got := GetLocales()

	sort.Strings(got)
	sort.Strings(want)

	if len(got) != len(want) {
		t.Error("Unexpected locale length")
		t.Fatalf(gotWant, len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf(gotWantIdx, i, got[i], want[i])
		}
	}
}
