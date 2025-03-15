// Package strftime implements C-like strftime format functionality with multiple locale support.
// It provides formatting of time.Time instances into strings using format specifiers similar to
// those used in C's strftime function. Supports various locales including English, Japanese,
// Chinese, and many European languages.
package strftime

import (
	"io"
	"time"

	"golang.org/x/text/language"
)

// Formatter represents a time formatter with specific locale settings.
// It handles the formatting of time values according to the specified locale.
type Formatter struct {
	l *strftimeLocaleInfo
}

// EnglishFormatter is a pre-initialized English locale formatter.
// It can be used directly without calling New() for English locale formatting.
var EnglishFormatter = &Formatter{englishLocale}

// Format is a shortcut to format a date in a given locale easily.
// Best performance is achieved by using language constants such as
// language.AmericanEnglish or language.French.
//
// Parameters:
//   - l: Language tag to determine the locale for formatting
//   - f: Format string with strftime-compatible format specifiers
//   - t: Time value to format
//
// Returns: Formatted time string according to the specified locale and format
func Format(l language.Tag, f string, t time.Time) string {
	locale, ok := strftimeLocaleTable[l]
	if !ok {
		// need to match locale
		_, i, _ := strftimeLocaleMatcher.Match(l)
		locale = strftimeLocales[i]
	}
	b := appendStrftime(locale, make([]byte, 0, 32+len(f)*2), []byte(f), t)

	return string(b)
}

// EnFormat formats time t using format f with English locale.
// This is a convenience function that uses the pre-initialized EnglishFormatter.
//
// Parameters:
//   - f: Format string with strftime-compatible format specifiers
//   - t: Time value to format
//
// Returns: Formatted time string in English locale
func EnFormat(f string, t time.Time) string {
	return EnglishFormatter.Format(f, t)
}

// EnFormatF formats time t using format f in English locale and outputs it to the provided io.Writer.
// This is a more efficient alternative to EnFormat when the output is to be written directly.
//
// Parameters:
//   - o: io.Writer to write the formatted output to
//   - f: Format string with strftime-compatible format specifiers
//   - t: Time value to format
//
// Returns: Error if writing to the io.Writer fails
func EnFormatF(o io.Writer, f string, t time.Time) error {
	return EnglishFormatter.FormatF(o, f, t)
}

// New creates a new Formatter by matching given language tags against known tags.
// If multiple language tags are provided, the best matching locale will be selected.
//
// One sample use is as follows:
//
//	t, q, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
//	f := strftime.New(t...)
//
// Parameters:
//   - l: One or more language tags to match against known locales
//
// Returns: A new Formatter instance configured for the best matching locale
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
// Uses the locale associated with this Formatter.
//
// Parameters:
//   - f: Format string with strftime-compatible format specifiers
//   - t: Time value to format
//
// Returns: Formatted time string according to this Formatter's locale
func (obj *Formatter) Format(f string, t time.Time) string {
	b := appendStrftime(obj.l, make([]byte, 0, 32+len(f)*2), []byte(f), t)

	return string(b)
}

// AppendFormat is like Format but appends the textual representation to b and returns the extended buffer.
// This is more efficient when building strings as it avoids unnecessary allocations.
//
// Parameters:
//   - b: Byte slice to append the formatted time to
//   - f: Format string with strftime-compatible format specifiers
//   - t: Time value to format
//
// Returns: The extended byte slice containing the original content followed by the formatted time
func (obj *Formatter) AppendFormat(b []byte, f string, t time.Time) []byte {
	return appendStrftime(obj.l, b, []byte(f), t)
}

// FormatF formats time using provided format, and outputs it to the provided io.Writer.
// This is more efficient than Format when the output is to be written directly.
//
// Parameters:
//   - o: io.Writer to write the formatted output to
//   - f: Format string with strftime-compatible format specifiers
//   - t: Time value to format
//
// Returns: Error if writing to the io.Writer fails
func (obj *Formatter) FormatF(o io.Writer, f string, t time.Time) error {
	// output implements the necessary methods to write runes & strings
	b := appendStrftime(obj.l, make([]byte, 0, 32+len(f)*2), []byte(f), t)
	_, err := o.Write(b)
	return err
}
