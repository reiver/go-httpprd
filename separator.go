package httpprd

// isSeparator return true if the rune is a separator-character,
// as defined by IETF RFC-2616, and return false otherwise.
//
// I.e.,:
//       separators     = "(" | ")" | "<" | ">" | "@"
//                      | "," | ";" | ":" | "\" | <">
//                      | "/" | "[" | "]" | "?" | "="
//                      | "{" | "}" | SP | HT
func isSeparator(r rune) bool {

	switch r {
	case
		'(' , ')' , '<' , '>' , '@' ,
		',' , ';' , ':' , '\\', '"' ,
		'/' , '[' , ']' , '?' , '=' ,
		'{' , '}' , ' ' , '\t':
		return true
	default:
		return false
	}
}
