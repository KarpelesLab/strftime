package strftime

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

var japaneseLocale = &strftimeLocaleInfo{
	tag:      language.Japanese,
	DTfmt:    "%Y年%m月%d日 %H時%M分%S秒",
	Dfmt:     "%Y年%m月%d日",
	Tfmt:     "%H時%M分%S秒",
	Tfmt12:   "%p%I時%M分%S秒",
	DTfmtEra: "%EY%m月%d日 %H時%M分%S秒",
	DfmtEra:  "%EY%m月%d日",
	AmPm:     [2]string{"午前", "午後"},
	Eyear:    strftimeJapaneseEra,

	AbDay:   [7]string{"日", "月", "火", "水", "木", "金", "土"},
	Day:     [7]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},
	AbMonth: [12]string{" 1月", " 2月", " 3月", " 4月", " 5月", " 6月", " 7月", " 8月", " 9月", "10月", "11月", "12月"},
	Month:   [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
}

// TODO need something more generic for era
func strftimeJapaneseEra(t time.Time, r byte) string {
	era := "西暦"
	y, m, d := t.Date()

	switch {
	case (y > 2019) || ((y == 2019) && (m > 5)) || ((y == 2019) && (m == 5) && (d >= 1)):
		// ??? era
	case (y > 1989) || ((y == 1989) && (m > 1)) || ((y == 1989) && (m == 1) && (d >= 8)):
		// Heisei era
		era = "平成"
		y -= 1988
	case (y > 1926) || ((y == 1926) && (m > 12)) || ((y == 1926) && (m == 12) && (d >= 25)):
		// Showa era
		era = "昭和"
		y -= 1925
	case (y > 1912) || ((y == 1912) && (m > 7)) || ((y == 1912) && (m == 7) && (d >= 30)):
		// Taisho era
		era = "大正"
		y -= 1911
	case (y > 1868) || ((y == 1868) && (m > 10)) || ((y == 1868) && (m == 10) && (d >= 23)):
		// Meiji era
		era = "明治"
		y -= 1867
	}

	switch r {
	case 'C':
		return era
	case 'y':
		return strconv.FormatInt(int64(y), 10)
	case 'Y':
		if y == 1 {
			// first year of a given era is called "Gannen"
			return era + "元年"
		} else {
			return fmt.Sprintf("%s%d年", era, y)
		}
	}
	return fmt.Sprintf("%%E%c", r)
}
