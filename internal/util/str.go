package util

import "strings"

// ParseFile parses a filename into its name and extension.
// It returns an error if the filename is empty or if it does not contain an extension.
// @param filename: the filename to parse
// @return string: the filename combine with string "_unlockeds" and extension
func SetUnlockFilename(filename *string) error {
	fn, ext, err := ParseFile(*filename)
	if err != nil {
		return err
	}

	*filename = strings.Join([]string{fn, "_unlockeds.", ext}, "")

	return nil
}
