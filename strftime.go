// go implementation of strftime
package strftime

import (
	"strings"
	"time"
)

// taken from time/format.go
var conversion = map[string]string {
	/*stdLongMonth      */ "B":"January",
	/*stdMonth          */ "b": "Jan",
	// stdNumMonth       */ "m": "1",
	/*stdZeroMonth      */ "m": "01",
	/*stdLongWeekDay    */ "A": "Monday",
	/*stdWeekDay        */ "a": "Mon",
	// stdDay            */ "d": "2",
	// stdUnderDay       */ "d": "_2",
	/*stdZeroDay        */ "d": "02",
	/*stdHour           */ "H": "15",
	// stdHour12         */ "I": "3",
	/*stdZeroHour12     */ "I": "03",
	// stdMinute         */ "M": "4",
	/*stdZeroMinute     */ "M": "04",
	// stdSecond         */ "S": "5",
	/*stdZeroSecond     */ "S": "05",
	/*stdLongYear       */ "Y": "2006",
	/*stdYear           */ "y": "06",
	/*stdPM             */ "p": "PM",
	// stdpm             */ "p": "pm",
	/*stdTZ             */ "Z": "MST",
	// stdISO8601TZ      */ "z": "Z0700",  // prints Z for UTC
	// stdISO8601ColonTZ */ "z": "Z07:00", // prints Z for UTC
	/*stdNumTZ          */ "z": "-0700",  // always numeric
	// stdNumShortTZ     */ "b": "-07",    // always numeric
	// stdNumColonTZ     */ "b": "-07:00", // always numeric
}

// This is an alternative to time.Format because no one knows 
// what date 040305 is supposed to create when used as a 'layout' string
// this takes standard strftime format options. For a complete list
// of format options see http://strftime.org/
func Format(format string, t time.Time) string {
	formatChunks := strings.Split(format, "%")
	var layout []string
	for _, chunk := range formatChunks {
		if len(chunk) == 0 {
			continue
		}
		if layoutCmd, ok := conversion[chunk[0:1]]; ok {
			layout = append(layout, layoutCmd)
			if len(chunk) > 1 {
				layout = append(layout, chunk[1:])
			}
		} else {
			layout = append(layout, "%", chunk)
		}
	}
	return t.Format(strings.Join(layout, ""))
}
