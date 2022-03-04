package transfer

import (
	`strings`
)

func path(base string, separator string, filename string) string {
	paths := strings.Split(base, separator)
	paths = append(paths, filename)

	return strings.Join(paths, separator)
}
