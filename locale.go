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
	language.BritishEnglish: &strftimeLocaleInfo{
		DTfmt:  "%a %d %b %Y %T %Z",
		Dfmt:   "%m/%d/%y",
		Tfmt:   "%T",
		Tfmt12: "%l:%M:%S %P %Z",
		AmPm:   [2]string{"am", "pm"},

		AbDay:   [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:     [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	},
	language.Spanish: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d/%m/%y",
		Tfmt:  "%T",

		AbDay:   [7]string{"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
		Day:     [7]string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		AbMonth: [12]string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"},
		Month:   [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
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
	language.Italian: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d/%m/%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"dom", "lun", "mar", "mer", "gio", "ven", "sab"},
		Day:     [7]string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"},
		AbMonth: [12]string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
		Month:   [12]string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
	},
	language.Dutch: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d-%m-%y",
		Tfmt:  "%T",

		AbDay:   [7]string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		Day:     [7]string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
		AbMonth: [12]string{"jan", "feb", "mrt", "apr", "mei", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		Month:   [12]string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
	},
	language.Polish: &strftimeLocaleInfo{
		DTfmt: "%a, %-d %b %Y, %T",
		Dfmt:  "%d.%m.%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"nie", "pon", "wto", "śro", "czw", "pią", "sob"},
		Day:     [7]string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
		AbMonth: [12]string{"sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
		Month:   [12]string{"styczeń", "luty", "marzec", "kwiecień", "maj", "czerwiec", "lipiec", "sierpień", "wrzesień", "październik", "listopad", "grudzień"},
	},
	language.Portuguese: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d-%m-%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"},
		Day:     [7]string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"},
		AbMonth: [12]string{"Jan", "Fev", "Mar", "Abr", "Mai", "Jun", "Jul", "Ago", "Set", "Out", "Nov", "Dez"},
		Month:   [12]string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"},
	},
	language.Russian: &strftimeLocaleInfo{
		DTfmt: "%a %d %b %Y %T",
		Dfmt:  "%d.%m.%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
		Day:     [7]string{"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"},
		AbMonth: [12]string{"янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
		Month:   [12]string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"},
	},
	language.Thai: &strftimeLocaleInfo{
		DTfmt:    "%a %e %b %Ey, %H:%M:%S",
		Dfmt:     "%d/%m/%Ey",
		Tfmt:     "%H:%M:%S",
		Tfmt12:   "%I:%M:%S %p",
		DTfmtEra: "วัน%Aที่ %e %B %EC %Ey, %H.%M.%S น.",
		DfmtEra:  "%e %b %Ey",
		TfmtEra:  "%H.%M.%S น.",
		AmPm:     [2]string{"AM", "PM"},
		// TODO: that era handling

		AbDay:   [7]string{"อา.", "จ.", "อ.", "พ.", "พฤ.", "ศ.", "ส."},
		Day:     [7]string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"},
		AbMonth: [12]string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
		Month:   [12]string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"},
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
	// only character for "hour" changes between Simplified Chinese(时) and Traditional Chinese(時)
	language.SimplifiedChinese: &strftimeLocaleInfo{
		DTfmt:  "%Y年%m月%d日 %A %H时%M分%S秒",
		Dfmt:   "%Y年%m月%d日",
		Tfmt:   "%H时%M分%S秒",
		Tfmt12: "%p %I时%M分%S秒",
		AmPm:   [2]string{"上午", "下午"},

		AbDay:   [7]string{"日", "一", "二", "三", "四", "五", "六"},
		Day:     [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		AbMonth: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		Month:   [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	},
	language.TraditionalChinese: &strftimeLocaleInfo{
		DTfmt:  "%Y年%m月%d日 (%A) %H時%M分%S秒",
		Dfmt:   "%Y年%m月%d日",
		Tfmt:   "%H時%M分%S秒",
		Tfmt12: "%p %I時%M分%S秒",
		AmPm:   [2]string{"上午", "下午"},

		AbDay:   [7]string{"日", "一", "二", "三", "四", "五", "六"},
		Day:     [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		AbMonth: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		Month:   [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	},
}

// TODO need something more generic for era
func strftimeJapaneseEra(t time.Time, r rune) string {
	era := "西暦"
	y, m, d := t.Date()

	switch {
	case (y > 1989) || ((y == 1989) && (m > 1)) || ((y == 1989) && (m == 1) && (d >= 8)):
		era = "平成"
		y -= 1988
	case (y > 1926) || ((y == 1926) && (m > 12)) || ((y == 1926) && (m == 12) && (d >= 25)):
		era = "昭和"
		y -= 1925
	case (y > 1912) || ((y == 1912) && (m > 7)) || ((y == 1912) && (m == 7) && (d >= 30)):
		era = "大正"
		y -= 1911
	case (y > 1868) || ((y == 1868) && (m > 10)) || ((y == 1868) && (m == 10) && (d >= 23)):
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
