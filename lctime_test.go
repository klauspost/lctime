package lctime

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	gotWant    = "got '%v', want '%v'"
	gotWantKey = "%s: got '%v', want '%v'"
	gotWantIdx = "%d: got '%v', want '%v'"
)

func TestSetLocale(t *testing.T) {
	fp := filepath.Join(localeDir, "bad.json")
	fd, err := os.Create(fp)
	if err != nil {
		t.Error(err)
	}
	fd.Close()
	defer os.Remove(fp)

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
		{"bad", ErrCorruptLocale},
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
