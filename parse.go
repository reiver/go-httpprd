package httpprd

import (
	"github.com/reiver/go-fck"
	"github.com/reiver/go-ord/en"
	"github.com/reiver/go-utf8"

	"io"
)

const (
	errNilRuneScanner = fck.Error("nil io.RuneScanner")
	errRuneError      = fck.Error("rune error")
)

// Parse parses an HTTP-product, as defined by IETF RFC-2616,
// and return the HTTP-product 'name' and HTTP-product 'version'.
//
// For example, this:
//
//	"HTTP/1.1"
//
// Would result in an HTTP-product name of:
//
//	"HTTP"
//
// And (would result in) an HTTP-product version of:
//
//	"1.1"
//
// From the IETF RFC-2616:
//
//	product         = token ["/" product-version]
//	product-version = token
//
//	token          = 1*<any CHAR except CTLs or separators>
//	separators     = "(" | ")" | "<" | ">" | "@"
//	               | "," | ";" | ":" | "\" | <">
//	               | "/" | "[" | "]" | "?" | "="
//	               | "{" | "}" | SP | HT
func Parse(rs io.RuneScanner) (name string, version string, err error) {

	if nil == rs {
/////////////// RETURN
		return "", "", errNilRuneScanner
	}

	{
		var count int64 = 0

		{
			var buffer []rune

			for {
				r, size, err := rs.ReadRune()
				count++
				if utf8.RuneError == r && nil == err {
					err = errRuneError
				}
				if io.EOF != err && nil != err {
/////////////////////////////////////// RETURN
					return "", "", fck.Errorf("problem reading the %s character: %v", orden.FormatInt64(count), err)
				}

				var iscontrol bool = isControl(r)
				var isseparator bool = isSeparator(r)

				{
					if 0 < size && utf8.RuneError != r && !isseparator && !iscontrol {
						buffer = append(buffer, r)
					}

					if isseparator || iscontrol || io.EOF == err {
						name = string(buffer)
					}
				}

				{
					if io.EOF == err {
/////////////////////////////////////////////// RETURN
						return name, version, nil
					}
					if '/' == r {
			/////////////////////// BREAK
						break
					}
					if isseparator || iscontrol {
						err := rs.UnreadRune()
						if nil != err {
/////////////////////////////////////////////////////// RETURN
							return name, version, fck.Errorf("problem unreading %s character (i.e., 0x%X): %v", orden.FormatInt64(count), r, err)
						}

/////////////////////////////////////////////// RETURN
						return name, version, nil
					}
				}
			}
		}

		{
			var buffer []rune

			for {
				r, size, err := rs.ReadRune()
				count++
				if utf8.RuneError == r && nil == err {
					err = errRuneError
				}
				if io.EOF != err && nil != err {
/////////////////////////////////////// RETURN
					return "", "", fck.Errorf("problem reading the %s character: %v", orden.FormatInt64(count), err)
				}

				var iscontrol bool = isControl(r)
				var isseparator bool = isSeparator(r)

				{
					if 0 < size && utf8.RuneError != r && !isseparator && !iscontrol {
						buffer = append(buffer, r)
					}

					if isseparator || iscontrol || io.EOF == err {
						version = string(buffer)
					}
				}

				{
					if io.EOF == err {
/////////////////////////////////////////////// RETURN
						return name, version, nil
					}
					if '/' == r {
			/////////////////////// BREAK
						break
					}
					if isseparator || iscontrol {
						err := rs.UnreadRune()
						if nil != err {
/////////////////////////////////////////////////////// RETURN
							return name, version, fck.Errorf("problem unreading %s character (i.e., 0x%X): %v", orden.FormatInt64(count), r, err)
						}

/////////////////////////////////////////////// RETURN
						return name, version, nil
					}
				}
			}
		}
	}

	return name, version, nil
}
