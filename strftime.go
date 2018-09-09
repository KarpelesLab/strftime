package strftime

import (
	"bufio"
	"bytes"
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
		_, i, _ := strftimeLocaleMatcher.Match(l)
		locale = strftimeLocaleTable[strftimeLocaleTags[i]]
	}
	b := &bytes.Buffer{}
	strftimeInternal(locale, b, f, t)

	return b.String()
}

// FormatUS formats time t using format f and English locale.
func FormatEnglish(f string, t time.Time) string {
	b := &bytes.Buffer{}
	strftimeInternal(englishLocale, b, f, t)

	return b.String()
}

func New(l language.Tag) *Formatter {
	locale, ok := strftimeLocaleTable[l]
	if !ok {
		// need to match locale
		_, i, _ := strftimeLocaleMatcher.Match(l)
		locale = strftimeLocaleTable[strftimeLocaleTags[i]]
	}

	return &Formatter{locale}
}

func (obj *Formatter) Format(f string, t time.Time) string {
	b := &bytes.Buffer{}
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
