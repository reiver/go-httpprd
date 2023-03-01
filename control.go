package httpprd

import (
	"unicode"
)

// isControl returns true if the rune is a control-character, and returns false otherwise.
//
// This functin handles Unicode control-characters and not just the control characters suggested by IETF RFC-2616.
func isControl(r rune) bool {
	return unicode.IsControl(r)
}
