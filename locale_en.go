// Package strftime implements C-like strftime functionality with locale support.
package strftime

import "golang.org/x/text/language"

var (
	// englishLocale defines the standard English locale information for formatting
	// dates and times. Used as the default locale when no specific locale is requested.
	englishLocale = &strftimeLocaleInfo{
		tag:    language.English,
		DTfmt:  "%a %b %e %H:%M:%S %Y", // Example: "Mon Jan  2 22:04:05 2006"
		Dfmt:   "%m/%d/%y",             // Example: "01/02/06"
		Tfmt:   "%H:%M:%S",             // Example: "22:04:05"
		Tfmt12: "%I:%M:%S %p",          // Example: "10:04:05 PM"
		AmPm:   [2]string{"AM", "PM"},  // AM/PM indicators

		// Day names in English
		AbDay: [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},                              // Abbreviated
		Day:   [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}, // Full

		// Month names in English
		AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},                                       // Abbreviated
		Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}, // Full
	}

	// americanEnglishLocale defines the American English locale information.
	// The main difference from standard English is the date format, which uses full year (%Y vs %y).
	americanEnglishLocale = &strftimeLocaleInfo{
		tag:    language.AmericanEnglish,
		DTfmt:  "%a %b %e %H:%M:%S %Y", // Example: "Mon Jan  2 22:04:05 2006"
		Dfmt:   "%m/%d/%Y",             // Example: "01/02/2006" (note the 4-digit year)
		Tfmt:   "%H:%M:%S",             // Example: "22:04:05"
		Tfmt12: "%I:%M:%S %p",          // Example: "10:04:05 PM"
		AmPm:   [2]string{"AM", "PM"},  // AM/PM indicators

		// Day names (same as standard English)
		AbDay: [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:   [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},

		// Month names (same as standard English)
		AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	}

	// britishEnglishLocale defines the British English locale information.
	// Notable differences include the date/time format and lowercase am/pm indicators.
	britishEnglishLocale = &strftimeLocaleInfo{
		tag:    language.BritishEnglish,
		DTfmt:  "%a %d %b %Y %T %Z",   // Example: "Mon 02 Jan 2006 22:04:05 UTC"
		Dfmt:   "%m/%d/%y",            // Example: "01/02/06"
		Tfmt:   "%T",                  // Example: "22:04:05" (using %T shorthand)
		Tfmt12: "%l:%M:%S %P %Z",      // Example: "10:04:05 pm UTC" (note lowercase pm)
		AmPm:   [2]string{"am", "pm"}, // Lowercase AM/PM indicators

		// Day names (same as standard English)
		AbDay: [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:   [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},

		// Month names (same as standard English)
		AbMonth: [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Month:   [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	}
)
