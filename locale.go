package strftime

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

type strftimeLocaleInfo struct {
	DTfmt  string // %c
	Dfmt   string // %x
	Tfmt   string // %X
	Tfmt12 string //  with am/pm

	// era-related stuff
	DTfmtEra string // alternative %c
	DfmtEra  string
	TfmtEra  string

	AmPm [2]string

	AbDay [7]string
	Day   [7]string

	// functions for extended %O(x) and %E(x)
	Oprint func(int) string
	Eyear  func(time.Time, rune) string // rune can be 'C', 'y' or 'Y'

	AbMonth [12]string
	Month   [12]string
}

var usLocale = &strftimeLocaleInfo{
	DTfmt:  "%a %b %e %H:%M:%S %Y",
	Dfmt:   "%m/%d/%Y",
	Tfmt:   "%H:%M:%S",
	Tfmt12: "%I:%M:%S %p",
	AmPm:   [2]string{"AM", "PM"},

	AbDay:   [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
	Day:     [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
	Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
}

var strftimeLocaleTable = map[language.Tag]*strftimeLocaleInfo{
	language.AmericanEnglish: usLocale,
	language.German: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d.%m.%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
		Day:     [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
		AbMonth: [12]string{"Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"},
		Month:   [12]string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
	},
	language.French: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d/%m/%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		Day:     [7]string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		AbMonth: [12]string{"janv.", "févr.", "mars", "avril", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		Month:   [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	language.Korean: &strftimeLocaleInfo{
		DTfmt:  "%x (%a) %r",
		Dfmt:   "%Y년 %m월 %d일",
		Tfmt:   "%H시 %M분 %S초",
		Tfmt12: "%p %I시 %M분 %S초",
		AmPm:   [2]string{"오전", "오후"},

		AbDay:   [7]string{"일", "월", "화", "수", "목", "금", "토"},
		Day:     [7]string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"},
		AbMonth: [12]string{" 1월", " 2월", " 3월", " 4월", " 5월", " 6월", " 7월", " 8월", " 9월", "10월", "11월", "12월"},
		Month:   [12]string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
	},
	language.Japanese: &strftimeLocaleInfo{
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
	},
}

// TODO need something more generic for era
func strftimeJapaneseEra(t time.Time, r rune) string {
	era := "西暦"
	y, m, d := t.Date()

	if (y > 1989) || ((y == 1989) && (m > 1)) || ((y == 1989) && (m == 1) && (d >= 8)) {
		era = "平成"
		y -= 1988
	} else if (y > 1926) || ((y == 1926) && (m > 12)) || ((y == 1926) && (m == 12) && (d >= 25)) {
		era = "昭和"
		y -= 1925
	} else if (y > 1912) || ((y == 1912) && (m > 7)) || ((y == 1912) && (m == 7) && (d >= 30)) {
		era = "大正"
		y -= 1911
	} else if (y > 1868) || ((y == 1868) && (m > 10)) || ((y == 1868) && (m == 10) && (d >= 23)) {
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
			return era + "元年"
		} else {
			return fmt.Sprintf("%s%d年", era, y)
		}
	}
	return fmt.Sprintf("%%E%c", r)
}
