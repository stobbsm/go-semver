// Ⓒ 2017 Matthew Stobbs <matthew@sproutingcommunications.com>
//
// Licensed under the MIT license. See the LICENSE file for details

// Package semver helps create semantic Version numbers according to the
// semver specification
package version

import (
	"bytes"
	"fmt"
	"unicode"
)

// Version holds the structure of the semantic version information
type Version struct {
	major   uint
	minor   uint
	patch   uint
	release string
	build   string
}

var version Version

// Set creates and returns a pointer to a semver Version number
func Set(major, minor, patch uint, release string, build string) {
	version = Version{
		major:   major,
		minor:   minor,
		patch:   patch,
		release: release,
		build:   build,
	}
}

// Get returns the String() of the internal version
func Get() string {
	return version.String()
}

// String returns a string of the Version
func (v *Version) String() string {
	var Version = fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)

	if v.release != "" {
		v.release = checkString(v.release)
		Version = fmt.Sprintf("%s-%s", Version, v.release)
	}

	if v.build != "" {
		Version = fmt.Sprintf("%s+%s", Version, v.build)
	}

	return Version
}

// TODO: Fix to make static build numbers possible
func (v *Version) setbuild() {
	if v.build == "" {
		v.build = "nil"
	}
}

// Checks a provided string to make sure it meets the requirements as a character in the Version number.
func checkString(s string) string {
	var result bytes.Buffer
	for _, r := range s {
		if unicode.IsLetter(r) {
			result.WriteRune(r)
		}
		if unicode.IsDigit(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
