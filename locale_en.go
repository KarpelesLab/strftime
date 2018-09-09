package strftime

import "golang.org/x/text/language"

var (
	englishLocale = &strftimeLocaleInfo{
		tag:    language.English,
		DTfmt:  "%a %b %e %H:%M:%S %Y",
		Dfmt:   "%m/%d/%y",
		Tfmt:   "%H:%M:%S",
		Tfmt12: "%I:%M:%S %p",
		AmPm:   [2]string{"AM", "PM"},

		AbDay:   [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:     [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	}

	americanEnglishLocale = &strftimeLocaleInfo{
		tag:    language.AmericanEnglish,
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

	britishEnglishLocale = &strftimeLocaleInfo{
		tag:    language.BritishEnglish,
		DTfmt:  "%a %d %b %Y %T %Z",
		Dfmt:   "%m/%d/%y",
		Tfmt:   "%T",
		Tfmt12: "%l:%M:%S %P %Z",
		AmPm:   [2]string{"am", "pm"},

		AbDay:   [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:     [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	}
)
