// Package strftime implements C-like strftime functionality with locale support.
package strftime

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

// japaneseLocale defines the Japanese locale information for formatting dates and times.
// It includes specialized formatting for Japanese era years and Japanese numerals.
var japaneseLocale = &strftimeLocaleInfo{
	tag:      language.Japanese,
	DTfmt:    "%Y年%m月%d日 %H時%M分%S秒",   // Date and time format
	Dfmt:     "%Y年%m月%d日",             // Date format
	Tfmt:     "%H時%M分%S秒",             // Time format
	Tfmt12:   "%p%I時%M分%S秒",           // 12-hour time format
	DTfmtEra: "%EY%m月%d日 %H時%M分%S秒",   // Date and time format with era
	DfmtEra:  "%EY%m月%d日",             // Date format with era
	AmPm:     [...]string{"午前", "午後"}, // AM/PM indicators
	Eyear:    strftimeJapaneseEra,     // Era year formatting function
	Oprint:   strftimeJapaneseDigit,   // Japanese numeral formatting function

	// Day names in Japanese
	AbDay: [...]string{"日", "月", "火", "水", "木", "金", "土"},
	Day:   [...]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},

	// Month names in Japanese
	AbMonth: [...]string{" 1月", " 2月", " 3月", " 4月", " 5月", " 6月", " 7月", " 8月", " 9月", "10月", "11月", "12月"},
	Month:   [...]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
}

// strftimeJapaneseEra formats years according to the Japanese era calendar system.
// The Japanese calendar uses era names (Reiwa, Heisei, Showa, etc.) and years are counted
// from the beginning of each era.
//
// Parameters:
//   - t: Time value to format
//   - r: Format request byte ('C' for era name, 'y' for year within era, 'Y' for full era + year)
//
// Returns: Formatted string representation of the year according to Japanese era system
func strftimeJapaneseEra(t time.Time, r byte) string {
	era := "西暦" // Default to Western calendar (Seireki)
	y, m, d := t.Date()

	switch {
	case (y > 2019) || ((y == 2019) && (m > 5)) || ((y == 2019) && (m == 5) && (d >= 1)):
		// Reiwa era (令和) - from May 1, 2019
		era = "令和"
		y -= 2018
	case (y > 1989) || ((y == 1989) && (m > 1)) || ((y == 1989) && (m == 1) && (d >= 8)):
		// Heisei era (平成) - from January 8, 1989 to April 30, 2019
		era = "平成"
		y -= 1988
	case (y > 1926) || ((y == 1926) && (m > 12)) || ((y == 1926) && (m == 12) && (d >= 25)):
		// Showa era (昭和) - from December 25, 1926 to January 7, 1989
		era = "昭和"
		y -= 1925
	case (y > 1912) || ((y == 1912) && (m > 7)) || ((y == 1912) && (m == 7) && (d >= 30)):
		// Taisho era (大正) - from July 30, 1912 to December 24, 1926
		era = "大正"
		y -= 1911
	case (y > 1868) || ((y == 1868) && (m > 10)) || ((y == 1868) && (m == 10) && (d >= 23)):
		// Meiji era (明治) - from October 23, 1868 to July 29, 1912
		era = "明治"
		y -= 1867
	}

	switch r {
	case 'C':
		return era // Return just the era name
	case 'y':
		return strconv.FormatInt(int64(y), 10) // Return just the year within era
	case 'Y':
		if y == 1 {
			// First year of a given era is called "Gannen" (元年)
			return era + "元年"
		} else {
			return fmt.Sprintf("%s%d年", era, y)
		}
	}
	return fmt.Sprintf("%%E%c", r) // Return unhandled format
}

// Japanese numeral characters for digits 0-9
var jpDigits = [...]string{"〇", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

// Japanese numeral unit markers and their values
var jpUnits = [...]struct {
	U string // Unit character
	V uint   // Unit value
}{
	//{"兆", 1000000000000}, // Trillion (uncommented as it's rarely needed)
	{"億", 100000000}, // Hundred million
	{"万", 10000},     // Ten thousand
	{"千", 1000},      // Thousand
	{"百", 100},       // Hundred
	{"十", 10},        // Ten
}

// strftimeJapaneseDigit converts a numeric value into traditional Japanese numeral representation.
// For example, 42 becomes "四十二" (four-ten-two).
//
// This function handles conversion to the traditional Japanese counting system, which groups
// digits differently than the Western system (by 10,000 rather than 1,000).
//
// Parameters:
//   - b: Byte slice to append the formatted result to
//   - v: Integer value to convert to Japanese numerals
//
// Returns: The byte slice with Japanese numerals appended
func strftimeJapaneseDigit(b []byte, v int) []byte {
	if v < 0 {
		// Generally shouldn't happen in date formatting
		v = -v
		b = append(b, '-')
	}
	u := uint(v)

	appd := false // Track if anything has been appended

	// Process each unit (億, 万, 千, 百, 十)
	for _, unit := range jpUnits {
		if u >= unit.V {
			n := u / unit.V
			if n >= 10 {
				// Recursively format numbers ≥ 10 for this unit
				b = strftimeJapaneseDigit(b, int(n))
			} else if n > 1 {
				// For values > 1, add the digit
				b = append(b, []byte(jpDigits[n])...)
			}
			// Add the unit marker (for n=1, only the marker is added without a digit)
			b = append(b, []byte(unit.U)...)
			u = u - (n * unit.V)
			appd = true
		}
	}

	// Add remaining digit or zero if nothing was appended
	if u > 0 || !appd {
		b = append(b, []byte(jpDigits[u])...)
	}

	return b
}
