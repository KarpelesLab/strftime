package strftime

import "golang.org/x/text/language"

// note: only character for "hour" changes between Simplified Chinese(时) and Traditional Chinese(時)

var (
	simplifiedChineseLocale = &strftimeLocaleInfo{
		tag:    language.SimplifiedChinese,
		DTfmt:  "%Y年%m月%d日 %A %H时%M分%S秒",
		Dfmt:   "%Y年%m月%d日",
		Tfmt:   "%H时%M分%S秒",
		Tfmt12: "%p %I时%M分%S秒",
		AmPm:   [2]string{"上午", "下午"},

		AbDay:   [7]string{"日", "一", "二", "三", "四", "五", "六"},
		Day:     [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		AbMonth: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		Month:   [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	}
	traditionalChineseLocale = &strftimeLocaleInfo{
		tag:    language.TraditionalChinese,
		DTfmt:  "%Y年%m月%d日 (%A) %H時%M分%S秒",
		Dfmt:   "%Y年%m月%d日",
		Tfmt:   "%H時%M分%S秒",
		Tfmt12: "%p %I時%M分%S秒",
		AmPm:   [2]string{"上午", "下午"},

		AbDay:   [7]string{"日", "一", "二", "三", "四", "五", "六"},
		Day:     [7]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		AbMonth: [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		Month:   [12]string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
	}
)
