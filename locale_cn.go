// Package strftime implements C-like strftime functionality with locale support.
package strftime

import "golang.org/x/text/language"

// Note: The main difference between Simplified Chinese (时) and Traditional Chinese (時)
// is the character used for "hour" in time formatting.

var (
	// simplifiedChineseLocale defines the Simplified Chinese (Mandarin) locale information
	// for formatting dates and times according to Chinese conventions.
	simplifiedChineseLocale = &strftimeLocaleInfo{
		tag:    language.SimplifiedChinese,
		DTfmt:  "%Y年%m月%d日 %A %H时%M分%S秒", // Date and time format (note the 时 character for hour)
		Dfmt:   "%Y年%m月%d日",              // Date format
		Tfmt:   "%H时%M分%S秒",              // Time format
		Tfmt12: "%p %I时%M分%S秒",           // 12-hour time format
		AmPm:   [2]string{"上午", "下午"},    // AM/PM indicators (morning/afternoon)

		// Weekday names - abbreviated versions are just the day numbers in Chinese
		AbDay: [7]string{"日", "一", "二", "三", "四", "五", "六"},               // Sun, Mon, Tue, etc.
		Day:   [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}, // Sunday, Monday, etc.

		// Month names - abbreviated versions use Arabic numerals with 月(month)
		AbMonth: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		// Full month names use Chinese numerals
		Month: [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	}

	// traditionalChineseLocale defines the Traditional Chinese locale information for formatting
	// dates and times (used primarily in Taiwan, Hong Kong, and Macau).
	traditionalChineseLocale = &strftimeLocaleInfo{
		tag:    language.TraditionalChinese,
		DTfmt:  "%Y年%m月%d日 (%A) %H時%M分%S秒", // Date and time format (note the 時 character for hour)
		Dfmt:   "%Y年%m月%d日",                // Date format
		Tfmt:   "%H時%M分%S秒",                // Time format
		Tfmt12: "%p %I時%M分%S秒",             // 12-hour time format
		AmPm:   [2]string{"上午", "下午"},      // AM/PM indicators (morning/afternoon)

		// Weekday names - abbreviated versions are just the day numbers in Chinese
		AbDay: [7]string{"日", "一", "二", "三", "四", "五", "六"},               // Sun, Mon, Tue, etc.
		Day:   [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}, // Sunday, Monday, etc.

		// Month names - abbreviated versions use Arabic numerals with 月(month)
		AbMonth: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		// Full month names use Chinese numerals
		Month: [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	}
)
