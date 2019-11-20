package semverfmt

import (
	"bytes"
	"strconv"

	"github.com/blang/semver"
)

// Sprintf formats a semantic version according to the string
// this code was adapted from github.com/tomnomnom/unfurl
func Sprintf(v semver.Version, format string) string {
	out := &bytes.Buffer{}

	inFormat := false
	for _, r := range format {

		if r == '%' && !inFormat {
			inFormat = true
			continue
		}

		if !inFormat {
			out.WriteRune(r)
			continue
		}

		switch r {

		// a literal percent rune
		case '%':
			out.WriteRune('%')

		// the major version
		case 'M':
			out.WriteString(strconv.FormatUint(v.Major, 10))

		// the minor version
		case 'm':
			out.WriteString(strconv.FormatUint(v.Minor, 10))

		// the patch version
		case 'p':
			out.WriteString(strconv.FormatUint(v.Patch, 10))

		}

		inFormat = false
	}

	return out.String()
}
