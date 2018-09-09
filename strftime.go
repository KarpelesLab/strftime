package strftime

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
)

type Formatter struct {
	l *strftimeLocaleInfo
}

type strftimeWriter interface {
	io.Writer
	WriteByte(byte) error
	WriteRune(rune) (int, error)
	WriteString(s string) (n int, err error)
}

// Format will format the time t following strftime format specified in f using language l.
func Format(l language.Tag, f string, t time.Time) string {
	locale, ok := strftimeLocaleTable[l]
	if !ok {
		// need to match locale
		avTag := make([]language.Tag, len(strftimeLocaleTable))
		n := int(0)
		for tag, _ := range strftimeLocaleTable {
			avTag[n] = tag
			n += 1
		}
		m := language.NewMatcher(avTag)
		_, i, _ := m.Match(l)
		locale = strftimeLocaleTable[avTag[i]]
	}
	b := strings.Builder{}
	strftimeInternal(locale, &b, f, t)

	return b.String()
}

// FormatUS formats time t using format f and US locale.
func FormatUS(f string, t time.Time) string {
	b := strings.Builder{}
	strftimeInternal(usLocale, &b, f, t)

	return b.String()
}

func New(l language.Tag) *Formatter {
	locale, ok := strftimeLocaleTable[l]
	if !ok {
		// need to match locale
		avTag := make([]language.Tag, len(strftimeLocaleTable))
		n := int(0)
		for tag, _ := range strftimeLocaleTable {
			avTag[n] = tag
			n += 1
		}
		m := language.NewMatcher(avTag)
		_, i, _ := m.Match(l)
		locale = strftimeLocaleTable[avTag[i]]
	}

	return &Formatter{locale}
}

func (obj *Formatter) Format(f string, t time.Time) string {
	b := strings.Builder{}
	strftimeInternal(obj.l, &b, f, t)

	return b.String()
}

func (obj *Formatter) FormatF(o io.Writer, f string, t time.Time) error {
	if b, ok := o.(strftimeWriter); ok {
		// output implements the necessary methods to write runes & strings
		strftimeInternal(obj.l, b, f, t)
		return nil
	} else {
		w := bufio.NewWriter(o)
		strftimeInternal(obj.l, w, f, t)
		return w.Flush()
	}
}

func writeInt(w io.Writer, x int, width int) error {
	u := uint(x)
	var buf [20]byte
	i := len(buf)

	// Assemble decimal in reverse order.
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		i--
		buf[i] = '0'
	}

	if x < 0 {
		i--
		buf[i] = '-'
	}

	_, err := w.Write(buf[i:])
	return err
}

// optimized writeInt for unsigned values below 256
func writeUint8(w io.Writer, u uint8, width int) error {
	var buf [3]byte
	i := len(buf)

	// Assemble decimal in reverse order.
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		i--
		buf[i] = '0'
	}

	_, err := w.Write(buf[i:])
	return err
}

// version with space padding
func writeUint8Sp(w io.Writer, u uint8, width int) error {
	var buf [3]byte
	i := len(buf)

	// Assemble decimal in reverse order.
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		i--
		buf[i] = ' '
	}

	_, err := w.Write(buf[i:])
	return err
}

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
			if thisPercent == 'E' && l.DTfmtEra != "" {
				strftimeInternal(l, b, l.DTfmtEra, t)
			} else {
				strftimeInternal(l, b, l.DTfmt, t)
			}
		case 'C': // century part of year
			if thisPercent == 'E' && l.Eyear != nil {
				b.WriteString(l.Eyear(t, r))
			} else {
				b.WriteString(strconv.FormatInt(int64(t.Year()/100), 10))
			}
		case 'd': // day (two decimals)
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Day()))
			} else {
				writeUint8(b, uint8(t.Day()), 2)
			}
		case 'D': // date (month/day/year format)
			strftimeInternal(l, b, "%m/%d/%y", t)
		case 'e': // day
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Day()))
			} else {
				writeUint8Sp(b, uint8(t.Day()), 2)
			}
		case 'E':
			prevPercent = 'E'
		case 'f':
			writeInt(b, t.Nanosecond()/1000, 6)
		case 'F':
			strftimeInternal(l, b, "%Y-%m-%d", t)
		case 'g':
			y, _ := t.ISOWeek()
			writeInt(b, y%100, 1)
		case 'G':
			y, _ := t.ISOWeek()
			writeInt(b, y, 1)
		case 'H':
			if thisPercent == 'O' && l.Oprint != nil {
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
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(h))
			} else {
				writeUint8(b, uint8(h), 2)
			}
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
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(int(t.Month())))
			} else {
				writeUint8(b, uint8(t.Month()), 2)
			}
		case 'M':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Minute()))
			} else {
				writeUint8(b, uint8(t.Minute()), 2)
			}
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
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Second()))
			} else {
				writeUint8(b, uint8(t.Second()), 2)
			}
		case 't':
			b.WriteByte('\t')
		case 'T':
			strftimeInternal(l, b, "%H:%M:%S", t)
		case 'u':
			wday := (int(t.Weekday()+6) % 7) + 1 // weekday but Monday = 1
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(wday))
			} else {
				b.WriteString(strconv.FormatInt(int64(wday), 10))
			}
		case 'U':
			// TODO test me
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint((((t.YearDay() - 1) - int(t.Weekday()) + 7) / 7)))
			} else {
				writeUint8(b, uint8(((t.YearDay()-1)-int(t.Weekday())+7)/7), 2)
			}
		case 'v': // non-standard extension found in https://github.com/lestrrat-go/strftime
			strftimeInternal(l, b, "%e-%b-%Y", t)
		case 'V':
			_, w := t.ISOWeek()
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(w))
			} else {
				writeUint8(b, uint8(w), 2)
			}
		case 'w':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(int(t.Weekday())))
			} else {
				b.WriteString(strconv.FormatInt(int64(t.Weekday()), 10))
			}
		case 'W': // same as %U, but with monday
			// TODO test me
			wday := int(t.Weekday()+6) % 7 // weekday but Monday = 0
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(((t.YearDay() - 1) - wday + 7) / 7))
			} else {
				writeUint8(b, uint8(((t.YearDay()-1)-wday+7)/7), 2)
			}
		case 'x':
			if thisPercent == 'E' && l.DfmtEra != "" {
				strftimeInternal(l, b, l.DfmtEra, t)
			} else {
				strftimeInternal(l, b, l.Dfmt, t)
			}
		case 'X':
			if thisPercent == 'E' && l.TfmtEra != "" {
				strftimeInternal(l, b, l.TfmtEra, t)
			} else {
				strftimeInternal(l, b, l.Tfmt, t)
			}
		case 'y':
			if thisPercent == 'E' && l.Eyear != nil {
				b.WriteString(l.Eyear(t, r))
			} else if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Year() % 100))
			} else {
				writeInt(b, t.Year()%100, 2)
			}
		case 'Y':
			if thisPercent == 'E' && l.Eyear != nil {
				b.WriteString(l.Eyear(t, r))
			} else {
				b.WriteString(strconv.FormatInt(int64(t.Year()), 10))
			}
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
		case '%':
			b.WriteByte('%')
		default:
			b.WriteRune('%')
			b.WriteRune(r)
		}
	}
}
