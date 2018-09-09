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

var EnglishFormatter = &Formatter{englishLocale}

// Format is a shortcut to format a date in a given locale easily. Best performance is achieved by using language constants such as language.AmericanEnglish or language.French.
func Format(l language.Tag, f string, t time.Time) string {
	locale, ok := strftimeLocaleTable[l]
	if !ok {
		// need to match locale
		_, i, _ := strftimeLocaleMatcher.Match(l)
		locale = strftimeLocales[i]
	}
	b := &bytes.Buffer{}
	strftimeInternal(locale, b, f, t)

	return b.String()
}

// EnFormat formats time t using format f and English locale.
func EnFormat(f string, t time.Time) string {
	return EnglishFormatter.Format(f, t)
}

// EnFormatF formats time t using format f in English locale and outputs it to the provided io.Writer.
func EnFormatF(o io.Writer, f string, t time.Time) error {
	return EnglishFormatter.FormatF(o, f, t)
}

// New creates a new Formatter by matching given language tags against known tags.
//
// One sample use is as follows:
// t, q, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
// f := strftime.New(t...)
func New(l ...language.Tag) *Formatter {
	if len(l) == 1 {
		if locale, ok := strftimeLocaleTable[l[0]]; ok {
			return &Formatter{locale}
		}
	}
	// need to match locale
	_, i, _ := strftimeLocaleMatcher.Match(l...)
	locale := strftimeLocales[i]

	return &Formatter{locale}
}

// Format formats time using provided format, and returns a string.
func (obj *Formatter) Format(f string, t time.Time) string {
	b := &bytes.Buffer{}
	strftimeInternal(obj.l, b, f, t)

	return b.String()
}

// FormatF formats time using provided format, and outputs it to the provided io.Writer.
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
