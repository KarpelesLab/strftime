package benchmarks

import (
	"fmt"
	"testing"
	"time"

	klbstrftime "github.com/KarpelesLab/strftime"
	cactus "github.com/cactus/gostrftime"
	fastly "github.com/fastly/go-utils/strftime"
	jehiah "github.com/jehiah/go-strftime"
	leekchan "github.com/leekchan/timeutil"
	lestrrat "github.com/lestrrat-go/strftime"
	tebeka "github.com/tebeka/strftime"
	"golang.org/x/text/language"
)

const benchfmt = `%A %a %B %b %d %H %I %M %m %p %S %Y %y %Z`

func BenchmarkCactus(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		cactus.Format(benchfmt, t)
	}
}

func BenchmarkLeekchan(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		leekchan.Strftime(&t, benchfmt)
	}
}

func BenchmarkTebeka(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		tebeka.Format(benchfmt, t)
	}
}

func BenchmarkJehiah(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		jehiah.Format(benchfmt, t)
	}
}

func BenchmarkFastly(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		fastly.Strftime(benchfmt, t)
	}
}

func BenchmarkLestrrat(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		lestrrat.Format(benchfmt, t)
	}
}

func BenchmarkKarpelesLab(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		klbstrftime.EnFormat(benchfmt, t)
	}
}

// BenchmarkFormatters runs benchmarks on various formatters and format strings
func BenchmarkFormatters(b *testing.B) {
	formats := map[string]string{
		"Simple":        "%Y-%m-%d",
		"Complex":       "%A, %B %d, %Y at %H:%M:%S %Z",
		"WithModifiers": "Year: %Y, ISO Week: %V, Day: %j, Time: %H:%M:%S",
		"WithSpecial":   "%Ec %EC %Ex %EX %Ey %EY",
		"WithNumerals":  "%Od %Om %OH:%OM:%OS %OV %OW %Ow",
		"WithEscapes":   "100%% complete with %%% and more %",
	}

	times := map[string]time.Time{
		"Current": time.Now().UTC(),
		"Epoch":   time.Unix(0, 0).UTC(),
		"Y2K":     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		"Future":  time.Date(2050, 12, 31, 23, 59, 59, 0, time.UTC),
	}

	locales := map[string]*klbstrftime.Formatter{
		"English":  klbstrftime.New(language.English),
		"Japanese": klbstrftime.New(language.Japanese),
		"Chinese":  klbstrftime.New(language.SimplifiedChinese),
		"German":   klbstrftime.New(language.German),
		"Russian":  klbstrftime.New(language.Russian),
		"French":   klbstrftime.New(language.French),
	}

	for localeName, formatter := range locales {
		for formatName, format := range formats {
			for timeName, timeValue := range times {
				name := fmt.Sprintf("Format/%s/%s/%s", localeName, formatName, timeName)
				b.Run(name, func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						formatter.Format(format, timeValue)
					}
				})
			}
		}
	}
}
