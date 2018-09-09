package strftime

import (
	"strconv"
	"strings"
	"time"
)

func strftimeInternal(l *strftimeLocaleInfo, b strftimeWriter, f string, t time.Time) {
	prevPercent := rune(0)

	for _, r := range f {
		if prevPercent == 0 {
			if r == '%' {
				prevPercent = r
			} else {
				b.WriteRune(r)
			}
			continue
		}
		thisPercent := prevPercent
		prevPercent = 0

		switch thisPercent {
		case 'E':
			// Era modifier
			switch r {
			case 'c':
				if l.DTfmtEra != "" {
					strftimeInternal(l, b, l.DTfmtEra, t)
				} else {
					strftimeInternal(l, b, l.DTfmt, t)
				}
			case 'C':
				if l.Eyear != nil {
					b.WriteString(l.Eyear(t, r))
				} else {
					b.WriteString(strconv.FormatInt(int64(t.Year()/100), 10))
				}
			case 'x':
				if l.DfmtEra != "" {
					strftimeInternal(l, b, l.DfmtEra, t)
				} else {
					strftimeInternal(l, b, l.Dfmt, t)
				}
			case 'X':
				if l.TfmtEra != "" {
					strftimeInternal(l, b, l.TfmtEra, t)
				} else {
					strftimeInternal(l, b, l.Tfmt, t)
				}
			case 'y':
				if l.Eyear != nil {
					b.WriteString(l.Eyear(t, r))
				} else {
					writeInt(b, t.Year()%100, 2)
				}
			case 'Y':
				if l.Eyear != nil {
					b.WriteString(l.Eyear(t, r))
				} else {
					b.WriteString(strconv.FormatInt(int64(t.Year()), 10))
				}
			default:
				b.WriteRune('%')
				b.WriteRune('E')
				b.WriteRune(r)
			}
		case 'O':
			switch r {
			case 'd': // day (two decimals)
				if l.Oprint != nil {
					b.WriteString(l.Oprint(t.Day()))
				} else {
					writeUint8(b, uint8(t.Day()), 2)
				}
			case 'e': // day
				if l.Oprint != nil {
					b.WriteString(l.Oprint(t.Day()))
				} else {
					writeUint8Sp(b, uint8(t.Day()), 2)
				}
			case 'H':
				if l.Oprint != nil {
					b.WriteString(l.Oprint(t.Hour()))
				} else {
					writeUint8(b, uint8(t.Hour()), 2)
				}
			case 'I':
				// Noon is 12PM, midnight is 12AM.
				h := t.Hour() % 12
				if h == 0 {
					h = 12
				}
				if l.Oprint != nil {
					b.WriteString(l.Oprint(h))
				} else {
					writeUint8(b, uint8(h), 2)
				}
			case 'm':
				if l.Oprint != nil {
					b.WriteString(l.Oprint(int(t.Month())))
				} else {
					writeUint8(b, uint8(t.Month()), 2)
				}
			case 'M':
				if l.Oprint != nil {
					b.WriteString(l.Oprint(t.Minute()))
				} else {
					writeUint8(b, uint8(t.Minute()), 2)
				}
			case 'S':
				if l.Oprint != nil {
					b.WriteString(l.Oprint(t.Second()))
				} else {
					writeUint8(b, uint8(t.Second()), 2)
				}
			case 'U':
				if l.Oprint != nil {
					b.WriteString(l.Oprint((((t.YearDay() - 1) - int(t.Weekday()) + 7) / 7)))
				} else {
					writeUint8(b, uint8(((t.YearDay()-1)-int(t.Weekday())+7)/7), 2)
				}
			case 'V':
				_, w := t.ISOWeek()
				if l.Oprint != nil {
					b.WriteString(l.Oprint(w))
				} else {
					writeUint8(b, uint8(w), 2)
				}
			case 'w':
				if l.Oprint != nil {
					b.WriteString(l.Oprint(int(t.Weekday())))
				} else {
					b.WriteString(strconv.FormatInt(int64(t.Weekday()), 10))
				}
			case 'W': // same as %U, but with monday
				wday := int(t.Weekday()+6) % 7 // weekday but Monday = 0
				if l.Oprint != nil {
					b.WriteString(l.Oprint(((t.YearDay() - 1) - wday + 7) / 7))
				} else {
					writeUint8(b, uint8(((t.YearDay()-1)-wday+7)/7), 2)
				}
			case 'y':
				if l.Oprint != nil {
					b.WriteString(l.Oprint(t.Year() % 100))
				} else {
					writeInt(b, t.Year()%100, 2)
				}
			default:
				b.WriteRune('%')
				b.WriteRune('O')
				b.WriteRune(r)
			}
		case '-':
			// no zero padding modified
			switch r {
			case 'd': // day (two decimals)
				writeUint8(b, uint8(t.Day()), 1)
			case 'H':
				writeUint8(b, uint8(t.Hour()), 1)
			case 'I':
				// Noon is 12PM, midnight is 12AM.
				h := t.Hour() % 12
				if h == 0 {
					h = 12
				}
				writeUint8(b, uint8(h), 1)
			case 'j':
				writeInt(b, t.YearDay(), 1)
			case 'm':
				writeUint8(b, uint8(t.Month()), 1)
			case 'M':
				writeUint8(b, uint8(t.Minute()), 1)
			case 'S':
				writeUint8(b, uint8(t.Second()), 1)
			default:
				b.WriteRune('%')
				b.WriteRune('-')
				b.WriteRune(r)
			}
		case '%':
			switch r {
			case 'a': // day (abbreviated)
				b.WriteString(l.AbDay[t.Weekday()])
			case 'A': // day
				b.WriteString(l.Day[t.Weekday()])
			case 'b', 'h': // month (abbreviated)
				b.WriteString(l.AbMonth[int(t.Month())-1])
			case 'B': // month
				b.WriteString(l.Month[int(t.Month())-1])
			case 'c': // date & time format
				strftimeInternal(l, b, l.DTfmt, t)
			case 'C': // century part of year
				b.WriteString(strconv.FormatInt(int64(t.Year()/100), 10))
			case 'd': // day (two decimals)
				writeUint8(b, uint8(t.Day()), 2)
			case 'D': // date (month/day/year format)
				strftimeInternal(l, b, "%m/%d/%y", t)
			case 'e': // day
				writeUint8Sp(b, uint8(t.Day()), 2)
			case 'E':
				prevPercent = 'E'
			case 'f':
				writeInt(b, t.Nanosecond()/1000, 6)
			case 'F':
				strftimeInternal(l, b, "%Y-%m-%d", t)
			case 'g':
				y, _ := t.ISOWeek()
				writeInt(b, y%100, 2)
			case 'G':
				y, _ := t.ISOWeek()
				writeInt(b, y, 1)
			case 'H':
				writeUint8(b, uint8(t.Hour()), 2)
			case 'I':
				// Noon is 12PM, midnight is 12AM.
				h := t.Hour() % 12
				if h == 0 {
					h = 12
				}
				writeUint8(b, uint8(h), 2)
			case 'j':
				writeInt(b, t.YearDay(), 3)
			case 'k':
				writeUint8Sp(b, uint8(t.Hour()), 2)
			case 'l':
				// Noon is 12PM, midnight is 12AM.
				h := t.Hour() % 12
				if h == 0 {
					h = 12
				}
				writeUint8Sp(b, uint8(h), 2)
			case 'm':
				writeUint8(b, uint8(t.Month()), 2)
			case 'M':
				writeUint8(b, uint8(t.Minute()), 2)
			case 'n':
				b.WriteByte('\n')
			case 'O':
				prevPercent = 'O'
			case 'p':
				if t.Hour() >= 12 {
					b.WriteString(strings.ToUpper(l.AmPm[1]))
				} else {
					b.WriteString(strings.ToUpper(l.AmPm[0]))
				}
			case 'P':
				if t.Hour() >= 12 {
					b.WriteString(strings.ToLower(l.AmPm[1]))
				} else {
					b.WriteString(strings.ToLower(l.AmPm[0]))
				}
			case 'r':
				strftimeInternal(l, b, "%I:%M:%S %p", t)
			case 'R':
				strftimeInternal(l, b, "%H:%M", t)
			case 's':
				b.WriteString(strconv.FormatInt(t.Unix(), 10))
			case 'S':
				writeUint8(b, uint8(t.Second()), 2)
			case 't':
				b.WriteByte('\t')
			case 'T':
				strftimeInternal(l, b, "%H:%M:%S", t)
			case 'u':
				wday := (int(t.Weekday()+6) % 7) + 1 // weekday but Monday = 1
				b.WriteString(strconv.FormatInt(int64(wday), 10))
			case 'U':
				writeUint8(b, uint8(((t.YearDay()-1)-int(t.Weekday())+7)/7), 2)
			case 'v': // non-standard extension found in https://github.com/lestrrat-go/strftime
				strftimeInternal(l, b, "%e-%b-%Y", t)
			case 'V':
				_, w := t.ISOWeek()
				writeUint8(b, uint8(w), 2)
			case 'w':
				b.WriteString(strconv.FormatInt(int64(t.Weekday()), 10))
			case 'W': // same as %U, but with monday
				wday := int(t.Weekday()+6) % 7 // weekday but Monday = 0
				writeUint8(b, uint8(((t.YearDay()-1)-wday+7)/7), 2)
			case 'x':
				strftimeInternal(l, b, l.Dfmt, t)
			case 'X':
				strftimeInternal(l, b, l.Tfmt, t)
			case 'y':
				writeInt(b, t.Year()%100, 2)
			case 'Y':
				b.WriteString(strconv.FormatInt(int64(t.Year()), 10))
			case 'z':
				_, z := t.Zone()
				z = z / 60 // convert seconds → minutes
				if z < 0 {
					b.WriteByte('-')
					z = -z
				} else {
					b.WriteByte('+')
				}
				writeUint8(b, uint8(z/60), 2)
				writeUint8(b, uint8(z%60), 2)
			case 'Z':
				n, _ := t.Zone()
				b.WriteString(n)
			case '-':
				prevPercent = '-'
			case '%':
				b.WriteByte('%')
			default:
				b.WriteRune('%')
				b.WriteRune(r)
			}
		}
	}
}
