package strftime

import (
	"bufio"
	"io"
	"time"

	"golang.org/x/text/language"
)

type Formatter struct {
	l *strftimeLocaleInfo
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
	b := makeWriterBuf()
	strftimeInternal(locale, b, f, t)

	return b.String()
}

// FormatUS formats time t using format f and US locale.
func FormatUS(f string, t time.Time) string {
	b := makeWriterBuf()
	strftimeInternal(usLocale, b, f, t)

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
	b := makeWriterBuf()
	strftimeInternal(obj.l, b, f, t)

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
