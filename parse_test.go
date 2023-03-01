package httpprd_test

import (
	"github.com/reiver/go-httpprd"

	"github.com/reiver/go-utf8"

	"io"
	"strings"

	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedName string
		ExpectedVersion string
	}{
		{
			Value:        "HTTP",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP ",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP,",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP/",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP{",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP\r\n",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP,gopher,finger",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},



		{
			Value:        "HTTP/0.9",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},
		{
			Value:        "HTTP/0.9 ",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},
		{
			Value:        "HTTP/0.9,",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},
		{
			Value:        "HTTP/0.9/",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},
		{
			Value:        "HTTP/0.9{",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},
		{
			Value:        "HTTP/0.9\r\n",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},
		{
			Value:        "HTTP/0.9,gopher,finger",
			ExpectedName: "HTTP",
			ExpectedVersion:   "0.9",
		},



		{
			Value:        "HTTP/1.0",
			ExpectedName: "HTTP",
			ExpectedVersion:   "1.0",
		},
		{
			Value:        "HTTP/1.1",
			ExpectedName: "HTTP",
			ExpectedVersion:   "1.1",
		},



		{
			Value:        "finger",
			ExpectedName: "finger",
			ExpectedVersion:     "",
		},
		{
			Value:        "finger/1.0",
			ExpectedName: "finger",
			ExpectedVersion:     "1.0",
		},
		{
			Value:        "finger/123.45",
			ExpectedName: "finger",
			ExpectedVersion:     "123.45",
		},
		{
			Value:        "finger/😈",
			ExpectedName: "finger",
			ExpectedVersion:     "😈",
		},
		{
			Value:        "finger/😈😈",
			ExpectedName: "finger",
			ExpectedVersion:     "😈😈",
		},



		{
			Value:        "❤️",
			ExpectedName: "❤️",
			ExpectedVersion: "",
		},
		{
			Value:        "❤️/123.45",
			ExpectedName: "❤️",
			ExpectedVersion: "123.45",
		},



		{
			Value:        "❤️❤️/2.0",
			ExpectedName: "❤️❤️",
			ExpectedVersion:   "2.0",
		},



		{
			Value:        "HTTP,1.1",
			ExpectedName: "HTTP",
			ExpectedVersion:   "",
		},
		{
			Value:        "HTTP/1.1.1",
			ExpectedName: "HTTP",
			ExpectedVersion:   "1.1.1",
		},
		{
			Value:        "HTTP/1.1.1 ",
			ExpectedName: "HTTP",
			ExpectedVersion:   "1.1.1",
		},
		{
			Value:        "apple/banana/cherry ",
			ExpectedName: "apple",
			ExpectedVersion:   "banana",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		rs := utf8.RuneScannerWrap(reader)
		var runescanner io.RuneScanner = &rs

		actualName, actualVersion, err := httpprd.Parse(runescanner)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %v", err, err)
			continue
		}

		{
			var expected string = test.ExpectedName
			var actual   string = actualName

			if expected != actual {
				t.Errorf("For test #%d, the actual value for 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED NAME: %q", expected)
				t.Logf("ACTUAL   NAME: %q", actual)
				continue
			}
		}

		{
			var expected string = test.ExpectedVersion
			var actual   string = actualVersion

			if expected != actual {
				t.Errorf("For test #%d, the actual value for 'version' is not what was expected.", testNumber)
				t.Logf("EXPECTED VERSION: %q", expected)
				t.Logf("ACTUAL   VERSION: %q", actual)
				continue
			}
		}

	}
}
