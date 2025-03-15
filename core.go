// Package strftime implements C-like strftime functionality with locale support.
package strftime

import (
	"bytes"
	"strings"
	"time"
)

// appendStrftime formats a time according to the specified format string and locale information,
// appending the result to a byte slice.
//
// This is the core formatting function that processes the format string and applies the appropriate
// formatting based on the conversion specifiers (% directives).
//
// Parameters:
//   - l: Pointer to locale information for language-specific formatting
//   - b: Byte slice to append formatted output to
//   - f: Format string as bytes with strftime-style directives
//   - t: Time value to format
//
// Returns: The extended byte slice with formatted time appended
//
// Supported format directives include:
//   - %a - Abbreviated weekday name
//   - %A - Full weekday name
//   - %b - Abbreviated month name
//   - %B - Full month name
//   - %c - Preferred date and time representation
//   - %C - Century (year/100)
//   - %d - Day of month as decimal (01-31)
//   - %D - Equivalent to %m/%d/%y
//   - %e - Day of month as decimal with leading space (1-31)
//   - %f - Microseconds
//   - %F - Equivalent to %Y-%m-%d (ISO 8601)
//   - %g - Last two digits of ISO week-based year
//   - %G - ISO week-based year
//   - %h - Same as %b
//   - %H - Hour (00-23)
//   - %I - Hour (01-12)
//   - %j - Day of year (001-366)
//   - %k - Hour with leading space (0-23)
//   - %l - Hour with leading space (1-12)
//   - %m - Month as decimal (01-12)
//   - %M - Minute (00-59)
//   - %n - Newline character
//   - %p - AM or PM
//   - %P - am or pm
//   - %r - Time in 12-hour format with AM/PM
//   - %R - Time in 24-hour format (%H:%M)
//   - %s - Seconds since Unix epoch
//   - %S - Second (00-60)
//   - %t - Tab character
//   - %T - Time in 24-hour format (%H:%M:%S)
//   - %u - Weekday as decimal (1-7, Monday=1)
//   - %U - Week number (00-53, Sunday as first day)
//   - %v - Date in form of %e-%b-%Y
//   - %V - ISO week number (01-53)
//   - %w - Weekday as decimal (0-6, Sunday=0)
//   - %W - Week number (00-53, Monday as first day)
//   - %x - Preferred date representation
//   - %X - Preferred time representation
//   - %y - Year without century (00-99)
//   - %Y - Year with century
//   - %z - Time zone offset
//   - %Z - Time zone name
//   - %% - Percent sign
//
// Extended modifiers supported (before specifier):
//   - %E - Alternative format (for date/time) - depends on locale, mainly used for era-based dates
//   - %O - Alternative numeral format - depends on locale, mainly used for non-latin numerals
//   - %- - No zero padding
func appendStrftime(l *strftimeLocaleInfo, b []byte, f []byte, t time.Time) []byte {
	var skip, i int

	for len(f) > 0 {
		i = bytes.IndexByte(f, '%')
		if i > 0 {
			b = append(b, f[:i]...)
			f = f[i:]
		} else if i == -1 {
			// end of string
			return append(b, f...)
		}
		// at this point, f always starts with a % symbol

		if len(f) < 2 {
			// can't have anything anymore
			return append(b, f...)
		}

		skip = 2 // number of bytes to skip

		switch f[1] {
		case 'E':
			if len(f) < 3 {
				// not enough data to process
				skip = 0
				break
			}
			skip = 3
			// Era modifier
			switch f[2] {
			case 'c':
				if l.DTfmtEra != "" {
					b = appendStrftime(l, b, []byte(l.DTfmtEra), t)
				} else {
					b = appendStrftime(l, b, []byte(l.DTfmt), t)
				}
			case 'C':
				if l.Eyear != nil {
					b = append(b, []byte(l.Eyear(t, 'C'))...)
				} else {
					b = appendInt(b, t.Year()/100, 1)
				}
			case 'x':
				if l.DfmtEra != "" {
					b = appendStrftime(l, b, []byte(l.DfmtEra), t)
				} else {
					b = appendStrftime(l, b, []byte(l.Dfmt), t)
				}
			case 'X':
				if l.TfmtEra != "" {
					b = appendStrftime(l, b, []byte(l.TfmtEra), t)
				} else {
					b = appendStrftime(l, b, []byte(l.Tfmt), t)
				}
			case 'y':
				if l.Eyear != nil {
					b = append(b, []byte(l.Eyear(t, 'y'))...)
				} else {
					b = appendInt(b, t.Year()%100, 2)
				}
			case 'Y':
				if l.Eyear != nil {
					b = append(b, []byte(l.Eyear(t, 'Y'))...)
				} else {
					b = appendInt(b, t.Year(), 1)
				}
			default:
				skip = 0
			}
		case 'O':
			if len(f) < 3 {
				// not enough data to process
				skip = 0
				break
			}
			skip = 3
			// alternative digits output (japanese, etc)
			var v uint8
			switch f[2] {
			case 'd', 'e': // day (two decimals)
				v = uint8(t.Day())
			case 'H':
				v = uint8(t.Hour())
			case 'I':
				// Noon is 12PM, midnight is 12AM.
				v = uint8(t.Hour() % 12)
				if v == 0 {
					v = 12
				}
			case 'm':
				v = uint8(t.Month())
			case 'M':
				v = uint8(t.Minute())
			case 'S':
				v = uint8(t.Second())
			case 'U':
				v = uint8(((t.YearDay() - 1) - int(t.Weekday()) + 7) / 7)
			case 'V':
				_, w := t.ISOWeek()
				v = uint8(w)
			case 'w':
				v = uint8(t.Weekday())
			case 'W': // same as %U, but with monday
				wday := int(t.Weekday()+6) % 7 // weekday but Monday = 0
				v = uint8(((t.YearDay() - 1) - wday + 7) / 7)
			case 'y':
				v = uint8(t.Year() % 100)
			default:
				skip = 0
			}
			if skip != 0 {
				if l.Oprint != nil {
					b = l.Oprint(b, int(v))
				} else {
					switch f[2] {
					case 'e':
						b = appendUint8Sp(b, v, 2)
					case 'w', 'W':
						b = appendUint8(b, v, 1)
					default:
						b = appendUint8(b, v, 2)
					}
				}
			}
		case '-':
			if len(f) < 3 {
				// not enough data to process
				skip = 0
				break
			}
			skip = 3
			// no zero padding modified
			switch f[2] {
			case 'd': // day (two decimals)
				b = appendUint8(b, uint8(t.Day()), 1)
			case 'H':
				b = appendUint8(b, uint8(t.Hour()), 1)
			case 'I':
				// Noon is 12PM, midnight is 12AM.
				h := t.Hour() % 12
				if h == 0 {
					h = 12
				}
				b = appendUint8(b, uint8(h), 1)
			case 'j':
				b = appendInt(b, t.YearDay(), 1)
			case 'm':
				b = appendUint8(b, uint8(t.Month()), 1)
			case 'M':
				b = appendUint8(b, uint8(t.Minute()), 1)
			case 'S':
				b = appendUint8(b, uint8(t.Second()), 1)
			default:
				skip = 0
			}
		case 'a': // day (abbreviated)
			b = append(b, []byte(l.AbDay[t.Weekday()])...)
		case 'A': // day
			b = append(b, []byte(l.Day[t.Weekday()])...)
		case 'b', 'h': // month (abbreviated)
			b = append(b, []byte(l.AbMonth[int(t.Month())-1])...)
		case 'B': // month
			b = append(b, []byte(l.Month[int(t.Month())-1])...)
		case 'c': // date & time format
			b = appendStrftime(l, b, []byte(l.DTfmt), t)
		case 'C': // century part of year
			b = appendInt(b, t.Year()/100, 1)
		case 'd': // day (two decimals)
			b = appendUint8(b, uint8(t.Day()), 2)
		case 'D': // date (month/day/year format)
			b = appendStrftime(l, b, []byte("%m/%d/%y"), t)
		case 'e': // day
			b = appendUint8Sp(b, uint8(t.Day()), 2)
		case 'f':
			b = appendInt(b, t.Nanosecond()/1000, 6)
		case 'F':
			b = appendStrftime(l, b, []byte("%Y-%m-%d"), t)
		case 'g':
			y, _ := t.ISOWeek()
			b = appendInt(b, y%100, 2)
		case 'G':
			y, _ := t.ISOWeek()
			b = appendInt(b, y, 1)
		case 'H':
			b = appendUint8(b, uint8(t.Hour()), 2)
		case 'I':
			// Noon is 12PM, midnight is 12AM.
			h := t.Hour() % 12
			if h == 0 {
				h = 12
			}
			b = appendUint8(b, uint8(h), 2)
		case 'j':
			b = appendInt(b, t.YearDay(), 3)
		case 'k':
			b = appendUint8Sp(b, uint8(t.Hour()), 2)
		case 'l':
			// Noon is 12PM, midnight is 12AM.
			h := t.Hour() % 12
			if h == 0 {
				h = 12
			}
			b = appendUint8Sp(b, uint8(h), 2)
		case 'm':
			b = appendUint8(b, uint8(t.Month()), 2)
		case 'M':
			b = appendUint8(b, uint8(t.Minute()), 2)
		case 'n':
			b = append(b, '\n')
		case 'p':
			if t.Hour() >= 12 {
				b = append(b, []byte(strings.ToUpper(l.AmPm[1]))...)
			} else {
				b = append(b, []byte(strings.ToUpper(l.AmPm[0]))...)
			}
		case 'P':
			if t.Hour() >= 12 {
				b = append(b, []byte(strings.ToLower(l.AmPm[1]))...)
			} else {
				b = append(b, []byte(strings.ToLower(l.AmPm[0]))...)
			}
		case 'r':
			b = appendStrftime(l, b, []byte("%I:%M:%S %p"), t)
		case 'R':
			b = appendStrftime(l, b, []byte("%H:%M"), t)
		case 's':
			b = appendInt64(b, t.Unix(), 1)
		case 'S':
			b = appendUint8(b, uint8(t.Second()), 2)
		case 't':
			b = append(b, '\t')
		case 'T':
			b = appendStrftime(l, b, []byte("%H:%M:%S"), t)
		case 'u':
			wday := (int(t.Weekday()+6) % 7) + 1 // weekday but Monday = 1
			b = appendUint8(b, uint8(wday), 1)
		case 'U':
			b = appendUint8(b, uint8(((t.YearDay()-1)-int(t.Weekday())+7)/7), 2)
		case 'v': // non-standard extension found in https://github.com/lestrrat-go/strftime
			b = appendStrftime(l, b, []byte("%e-%b-%Y"), t)
		case 'V':
			_, w := t.ISOWeek()
			b = appendUint8(b, uint8(w), 2)
		case 'w':
			b = appendUint8(b, uint8(t.Weekday()), 1)
		case 'W': // same as %U, but with monday
			wday := int(t.Weekday()+6) % 7 // weekday but Monday = 0
			b = appendUint8(b, uint8(((t.YearDay()-1)-wday+7)/7), 2)
		case 'x':
			b = appendStrftime(l, b, []byte(l.Dfmt), t)
		case 'X':
			b = appendStrftime(l, b, []byte(l.Tfmt), t)
		case 'y':
			b = appendInt(b, t.Year()%100, 2)
		case 'Y':
			b = appendInt(b, t.Year(), 1)
		case 'z':
			_, z := t.Zone()
			z = z / 60 // convert seconds → minutes
			if z < 0 {
				b = append(b, '-')
				z = -z
			} else {
				b = append(b, '+')
			}
			b = appendUint8(b, uint8(z/60), 2)
			b = appendUint8(b, uint8(z%60), 2)
		case 'Z':
			n, _ := t.Zone()
			b = append(b, []byte(n)...)
		case '%':
			b = append(b, '%')
		default:
			skip = 0
		}

		// move f pointer
		if skip == 0 {
			b = append(b, f[0])
			f = f[1:]
		} else {
			f = f[skip:]
		}
	}
	return b
}
