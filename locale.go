package strftime

import (
	"time"

	"golang.org/x/text/language"
)

type strftimeLocaleInfo struct {
	tag language.Tag

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
	Oprint func([]byte, int) []byte
	Eyear  func(time.Time, byte) string // byte can be 'C', 'y' or 'Y'

	AbMonth [12]string
	Month   [12]string
}

var (
	strftimeLocaleMatcher language.Matcher
	strftimeLocaleTable   map[language.Tag]*strftimeLocaleInfo
)

func init() {
	// initialize and fill the variables
	strftimeLocaleTable = make(map[language.Tag]*strftimeLocaleInfo)
	matcherTable := make([]language.Tag, len(strftimeLocales))

	for i, loc := range strftimeLocales {
		strftimeLocaleTable[loc.tag] = loc
		matcherTable[i] = loc.tag
	}

	strftimeLocaleMatcher = language.NewMatcher(matcherTable)
}

var strftimeLocales = [...]*strftimeLocaleInfo{
	englishLocale,
	americanEnglishLocale,
	britishEnglishLocale,
	&strftimeLocaleInfo{
		tag:   language.Spanish,
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d/%m/%y",
		Tfmt:  "%T",

		AbDay:   [7]string{"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
		Day:     [7]string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		AbMonth: [12]string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"},
		Month:   [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
	},
	&strftimeLocaleInfo{
		tag:   language.German,
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d.%m.%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
		Day:     [7]string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
		AbMonth: [12]string{"Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"},
		Month:   [12]string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
	},
	&strftimeLocaleInfo{
		tag:   language.French,
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d/%m/%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		Day:     [7]string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		AbMonth: [12]string{"janv.", "févr.", "mars", "avril", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		Month:   [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
	},
	&strftimeLocaleInfo{
		tag:   language.Italian,
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d/%m/%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"dom", "lun", "mar", "mer", "gio", "ven", "sab"},
		Day:     [7]string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"},
		AbMonth: [12]string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
		Month:   [12]string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
	},
	&strftimeLocaleInfo{
		tag:   language.Dutch,
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d-%m-%y",
		Tfmt:  "%T",

		AbDay:   [7]string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		Day:     [7]string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
		AbMonth: [12]string{"jan", "feb", "mrt", "apr", "mei", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		Month:   [12]string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
	},
	&strftimeLocaleInfo{
		tag:   language.Polish,
		DTfmt: "%a, %-d %b %Y, %T",
		Dfmt:  "%d.%m.%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"nie", "pon", "wto", "śro", "czw", "pią", "sob"},
		Day:     [7]string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
		AbMonth: [12]string{"sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
		Month:   [12]string{"styczeń", "luty", "marzec", "kwiecień", "maj", "czerwiec", "lipiec", "sierpień", "wrzesień", "październik", "listopad", "grudzień"},
	},
	&strftimeLocaleInfo{
		tag:   language.Portuguese,
		DTfmt: "%a %d %b %Y %T %Z",
		Dfmt:  "%d-%m-%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"},
		Day:     [7]string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"},
		AbMonth: [12]string{"Jan", "Fev", "Mar", "Abr", "Mai", "Jun", "Jul", "Ago", "Set", "Out", "Nov", "Dez"},
		Month:   [12]string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"},
	},
	&strftimeLocaleInfo{
		tag:   language.Russian,
		DTfmt: "%a %d %b %Y %T",
		Dfmt:  "%d.%m.%Y",
		Tfmt:  "%T",

		AbDay:   [7]string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
		Day:     [7]string{"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"},
		AbMonth: [12]string{"янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
		Month:   [12]string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"},
	},
	&strftimeLocaleInfo{
		tag:      language.Thai,
		DTfmt:    "%a %e %b %Ey, %H:%M:%S",
		Dfmt:     "%d/%m/%Ey",
		Tfmt:     "%H:%M:%S",
		Tfmt12:   "%I:%M:%S %p",
		DTfmtEra: "วัน%Aที่ %e %B %EC %Ey, %H.%M.%S น.",
		DfmtEra:  "%e %b %Ey",
		TfmtEra:  "%H.%M.%S น.",
		AmPm:     [2]string{"AM", "PM"},
		// TODO: thai era handling

		AbDay:   [7]string{"อา.", "จ.", "อ.", "พ.", "พฤ.", "ศ.", "ส."},
		Day:     [7]string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"},
		AbMonth: [12]string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
		Month:   [12]string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"},
	},
	&strftimeLocaleInfo{
		tag:    language.Korean,
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
	japaneseLocale,
	simplifiedChineseLocale,
	traditionalChineseLocale,
}
