// nolint:scopelint
package semverfmt

import (
	"strconv"
	"testing"

	"github.com/blang/semver"
)

func TestSprintf(t *testing.T) {
	cases := []struct {
		version  string
		format   string
		expected string
	}{
		{"v3.5.7", "%M", "3"},
		{"v3.5.7", "v%M", "v3"},
		{"v3.5.7", "%M.%m", "3.5"},
		{"v3.5.7", "v%M.%m", "v3.5"},
		{"1.2.3", "%M.%m.%p-blah", "1.2.3-blah"},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v, _ := semver.ParseTolerant(c.version)

			actual := Sprintf(v, c.format)

			if actual != c.expected {
				t.Errorf("want %s for version(%s) and format(%s); have %s", c.expected, c.version, c.format, actual)
			}
		})
	}
}
