package strftime

import (
	"bufio"
	"fmt"
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
		case 'a':
			b.WriteString(l.AbDay[t.Weekday()])
		case 'A':
			b.WriteString(l.Day[t.Weekday()])
		case 'b', 'h':
			b.WriteString(l.AbMonth[int(t.Month())-1])
		case 'B':
			b.WriteString(l.Month[int(t.Month())-1])
		case 'c':
			if thisPercent == 'E' && l.DTfmtEra != "" {
				strftimeInternal(l, b, l.DTfmtEra, t)
			} else {
				strftimeInternal(l, b, l.DTfmt, t)
			}
		case 'C':
			if thisPercent == 'E' && l.Eyear != nil {
				b.WriteString(l.Eyear(t, r))
			} else {
				b.WriteString(strconv.FormatInt(int64(t.Year()/100), 10))
			}
		case 'd':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Day()))
			} else {
				fmt.Fprintf(b, "%02d", t.Day())
			}
		case 'D':
			strftimeInternal(l, b, "%m/%d/%y", t)
		case 'e':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Day()))
			} else {
				b.WriteString(fmt.Sprintf("%2d", t.Day()))
			}
		case 'E':
			prevPercent = 'E'
		case 'f':
			b.WriteString(fmt.Sprintf("%06d", t.Nanosecond()/1000))
		case 'F':
			strftimeInternal(l, b, "%Y-%m-%d", t)
		case 'g':
			y, _ := t.ISOWeek()
			b.WriteString(fmt.Sprintf("%d", y%100))
		case 'G':
			y, _ := t.ISOWeek()
			b.WriteString(fmt.Sprintf("%d", y))
		case 'H':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Hour()))
			} else {
				b.WriteString(fmt.Sprintf("%02d", t.Hour()))
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
				b.WriteString(fmt.Sprintf("%02d", h))
			}
		case 'j':
			b.WriteString(fmt.Sprintf("%03d", t.YearDay()))
		case 'k':
			b.WriteString(fmt.Sprintf("%2d", t.Hour()))
		case 'l':
			// Noon is 12PM, midnight is 12AM.
			h := t.Hour() % 12
			if h == 0 {
				h = 12
			}
			b.WriteString(fmt.Sprintf("%2d", h))
		case 'm':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(int(t.Month())))
			} else {
				b.WriteString(fmt.Sprintf("%02d", t.Month()))
			}
		case 'M':
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(t.Minute()))
			} else {
				b.WriteString(fmt.Sprintf("%02d", t.Minute()))
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
				b.WriteString(fmt.Sprintf("%02d", t.Second()))
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
				fmt.Fprintf(b, "%02d", ((t.YearDay()-1)-int(t.Weekday())+7)/7)
			}
		case 'v': // non-standard extension found in https://github.com/lestrrat-go/strftime
			strftimeInternal(l, b, "%e-%b-%Y", t)
		case 'V':
			_, w := t.ISOWeek()
			if thisPercent == 'O' && l.Oprint != nil {
				b.WriteString(l.Oprint(w))
			} else {
				fmt.Fprintf(b, "%02d", w)
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
				fmt.Fprintf(b, "%02d", ((t.YearDay()-1)-wday+7)/7)
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
				b.WriteString(fmt.Sprintf("%02d", t.Year()%100))
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
			b.WriteString(fmt.Sprintf("%02d%02d", z/60, z%60))
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
