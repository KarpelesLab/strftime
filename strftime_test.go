package strftime_test

import (
	"bytes"
	"testing"
	"time"

	"golang.org/x/text/language"

	"github.com/KarpelesLab/strftime"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	var ref = time.Unix(1136239445, 456841962).UTC()
	f := strftime.New(language.English)

	cmp := []struct{ A, B string }{
		{`%A`, `Monday`},
		{`%a`, `Mon`},
		{`%b`, `Jan`}, // same as %h
		{`%B`, `January`},
		{`%C`, `20`},
		{`%c`, `Mon Jan  2 22:04:05 2006`},
		{`%D`, `01/02/06`},
		{`%d`, `02`},
		{`%e`, ` 2`},
		{`%f`, `456841`},
		{`%F`, `2006-01-02`},
		{`%g`, `06`},
		{`%G`, `2006`},
		{`%H`, `22`},
		{`%h`, `Jan`},
		{`%I`, `10`},
		{`%j`, `002`},
		{`%k`, `22`},
		{`%l`, `10`},
		{`%M`, `04`},
		{`%m`, `01`},
		{`%n`, "\n"},
		{`%p`, `PM`},
		{`%P`, `pm`},
		{`%R`, `22:04`},
		{`%r`, `10:04:05 PM`},
		{`%S`, `05`},
		{`%s`, `1136239445`},
		{`%T`, `22:04:05`},
		{`%t`, "\t"},
		{`%U`, `01`},
		{`%u`, `1`},
		{`%V`, `01`},
		{`%v`, ` 2-Jan-2006`},
		{`%W`, `01`},
		{`%w`, `1`},
		{`%X`, `22:04:05`},
		{`%x`, `01/02/06`},
		{`%Y`, `2006`},
		{`%y`, `06`},
		{`%Z`, `UTC`},
		{`%z`, `+0000`},
		{`%%`, `%`},

		{`%-d`, `2`},
		{`%-m`, `1`},
		{`%-H`, `22`},
		{`%-I`, `10`},
		{`%-M`, `4`},
		{`%-S`, `5`},
		{`%-j`, `2`},

		{`%E %`, `%E %`},
		{`%E`, `%E`},
		{`%-`, `%-`},
		{`Test %O`, `Test %O`},

		// full test from https://github.com/lestrrat-go/strftime
		{`%A %a %B %b %C %c %D %d %e %F %H %h %I %j %k %l %M %m %n %p %R %r %S %T %t %U %u %V %v %W %w %X %x %Y %y %Z %z`, "Monday Mon January Jan 20 Mon Jan  2 22:04:05 2006 01/02/06 02  2 2006-01-02 22 Jan 10 002 22 10 04 01 \n PM 22:04 10:04:05 PM 05 22:04:05 \t 01 1 01  2-Jan-2006 01 1 22:04:05 01/02/06 2006 06 UTC +0000"},

		{`%Ec`, `Mon Jan  2 22:04:05 2006`},
		{`%EC`, `20`},
		{`%Ex`, `01/02/06`},
		{`%EX`, `22:04:05`},
		{`%Ey`, `06`},
		{`%EY`, `2006`},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, ref), `matching for `+x.A)
	}
}

func TestValues(t *testing.T) {
	f := strftime.New(language.English)

	cmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%c %w %W %g %G %U`, `Mon Jan  2 22:04:05 2006 1 01 06 2006 01`, time.Unix(1136239445, 0)},
		{`%c %w %W %g %G %U`, `Sat Jan  1 04:05:06 2005 6 00 04 2004 00`, time.Unix(1104552306, 0)},
		{`%c %w %W %g %G %U`, `Tue Dec 30 04:05:06 2008 2 52 09 2009 52`, time.Unix(1230609906, 0)},
		{`%c %w %W %g %G %U`, `Tue Nov 15 08:17:19 1994 2 46 94 1994 46`, time.Unix(784887439, 0)},
		{`%c %w %W %g %G %U`, `Tue Aug  7 22:12:35 1984 2 32 84 1984 32`, time.Unix(460764755, 0)},
		{`%c %w %W %g %G %U`, `Fri Mar 20 03:38:53 2082 5 11 82 2082 11`, time.Unix(3541203533, 0)},
		{`%c %w %W %g %G %U`, `Sun Feb  7 06:28:15 2106 0 05 06 2106 06`, time.Unix(0xffffffff, 0)},
		{`%c %w %W %g %G %U`, `Mon Feb 20 00:36:15 36812 1 08 12 36812 08`, time.Unix(0xffffffffff, 0)},
		{`%c %w %W %g %G %U`, `Fri Dec  7 10:44:15 8921556 5 49 56 8921556 49`, time.Unix(0xffffffffffff, 0)},
		{`%c %w %W %g %G %U`, `Sun Nov 24 17:31:45 1833 0 46 33 1833 47`, time.Unix(-0xffffffff, 0)},
		{`%c %w %W %g %G %U`, `Mon Jan  1 00:00:00 1 1 01 01 1 00`, time.Unix(-62135596800, 0)},
		{`%c %w %W %g %G %U`, `Fri Jan  7 00:00:00 0 5 01 00 0 01`, time.Unix(-62166700800, 0)},
		{`%c %w %W %g %G %U`, `Sat Jan  1 00:00:00 0 6 00 -01 -1 00`, time.Unix(-62167219200, 0)},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, x.T.UTC()), `matching for `+x.A)
	}
}

func TestJapanese(t *testing.T) {
	var ref = time.Unix(1136239445, 456841962).UTC()
	f := strftime.New(language.Japanese)

	cmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%Ec`, `平成18年01月02日 22時04分05秒`, ref},
		{`%Ex`, `平成18年01月02日`, ref},
		{`%Ex`, `昭和64年01月07日`, time.Unix(600134400, 0)},
		{`%Ex`, `平成元年01月08日`, time.Unix(600220800, 0)},
		{`%Ex`, `昭和20年01月01日`, time.Unix(-788918400, 0)},
		{`%Ex`, `明治34年01月01日`, time.Unix(-2177452800, 0)},
		{`%Ex`, `西暦1801年01月01日`, time.Unix(-5333126400, 0)},
		{`%Ex`, `明治7年01月01日`, time.Unix(-3029443200, 0)},
		{`%Ex`, `大正4年01月01日`, time.Unix(-1735689600, 0)},
		{`%Ex`, `令和元年10月17日`, time.Unix(1571307445, 0)},
		{`%Oy`, `六`, ref},
		{`%OH`, `二十二`, ref},
		{`%OI`, `十`, ref},
		{`%Od %Om %OH:%OM:%OS %OV %OW %Ow`, `二 一 二十二:四:五 一 一 一`, ref},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, x.T.UTC()), `matching for `+x.A)
	}
}

func TestFrench(t *testing.T) {
	ref := time.Unix(1136239445, 456841962).UTC()
	f := strftime.New(language.French)

	cmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%p %P`, ` `, ref},
		{`%A %d %B %Y`, `lundi 02 janvier 2006`, ref},
		{`%a %d %b %Y`, `lun. 02 janv. 2006`, ref},
		{`%x`, `02/01/2006`, ref},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, x.T.UTC()), `matching for `+x.A)
	}
}

// TestGerman tests the German locale formatting
func TestGerman(t *testing.T) {
	ref := time.Unix(1136239445, 456841962).UTC()
	f := strftime.New(language.German)

	cmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%A`, `Montag`, ref},
		{`%a`, `Mo`, ref},
		{`%B`, `Januar`, ref},
		{`%b`, `Jan`, ref},
		{`%x`, `02.01.2006`, ref},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, x.T.UTC()), `matching for `+x.A)
	}
}

// TestRussian tests the Russian locale formatting
func TestRussian(t *testing.T) {
	ref := time.Unix(1136239445, 456841962).UTC()
	f := strftime.New(language.Russian)

	cmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%A`, `Понедельник`, ref},
		{`%a`, `Пн`, ref},
		{`%B`, `Январь`, ref},
		{`%b`, `янв`, ref},
		{`%x`, `02.01.2006`, ref},
		{`%c`, `Пн 02 янв 2006 22:04:05`, ref},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, x.T.UTC()), `matching for `+x.A)
	}
}

func TestLocale(t *testing.T) {
	var ref = time.Unix(1136239445, 456841962).UTC()

	cmp := []struct{ A, B string }{
		{`nn;q=0.3, en-us;q=0.8, en,`, `Mon Jan  2 22:04:05 2006`},
		{`gsw, en;q=0.7, en-US;q=0.8`, `Mon Jan  2 22:04:05 2006`},
		{`gsw, nl, da`, `ma 02 jan 2006 22:04:05 UTC`},
		{`fr`, `lun. 02 janv. 2006 22:04:05 UTC`},
		{`invalid`, `Mon Jan  2 22:04:05 2006`},
	}

	for _, x := range cmp {
		tags, _, _ := language.ParseAcceptLanguage(x.A)
		f := strftime.New(tags...)
		assert.Equal(t, x.B, f.Format(`%c`, ref), `language detect for `+x.A)
	}
}

// TestChinese tests both Simplified and Traditional Chinese locale formatting
func TestChinese(t *testing.T) {
	ref := time.Unix(1136239445, 456841962).UTC()

	// Test Simplified Chinese
	scf := strftime.New(language.SimplifiedChinese)
	cmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%A`, `星期一`, ref},
		{`%a`, `一`, ref},
		{`%B`, `一月`, ref},
		{`%b`, `1月`, ref},
		{`%Y年%m月%d日`, `2006年01月02日`, ref},
		{`%H时%M分%S秒`, `22时04分05秒`, ref}, // Note the simplified character for hour: 时
		{`%p`, `下午`, ref},               // PM in Chinese
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, scf.Format(x.A, x.T.UTC()), `Simplified Chinese matching for `+x.A)
	}

	// Test Traditional Chinese
	tcf := strftime.New(language.TraditionalChinese)
	tcmp := []struct {
		A, B string
		T    time.Time
	}{
		{`%A`, `星期一`, ref},
		{`%a`, `一`, ref},
		{`%B`, `一月`, ref},
		{`%b`, `1月`, ref},
		{`%Y年%m月%d日`, `2006年01月02日`, ref},
		{`%H時%M分%S秒`, `22時04分05秒`, ref}, // Note the traditional character for hour: 時
		{`%p`, `下午`, ref},               // PM in Chinese
	}

	for _, x := range tcmp {
		assert.Equal(t, x.B, tcf.Format(x.A, x.T.UTC()), `Traditional Chinese matching for `+x.A)
	}
}

// TestEdgeCases tests various edge cases and uncommon format combinations
func TestEdgeCases(t *testing.T) {
	f := strftime.New(language.English)

	// Test leap year Feb 29
	leapDay := time.Date(2020, 2, 29, 12, 30, 45, 0, time.UTC)
	assert.Equal(t, "02/29/20", f.Format("%D", leapDay), "Leap day formatting")
	assert.Equal(t, "060", f.Format("%j", leapDay), "Day of year for leap day")

	// Test December 31 (last day of year)
	lastDay := time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC)
	assert.Equal(t, "366", f.Format("%j", lastDay), "Day of year for Dec 31 in leap year")
	assert.Equal(t, "52", f.Format("%U", lastDay), "Week number for last day of year")

	// Test midnight formatting
	midnight := time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, "00:00:00", f.Format("%T", midnight), "Midnight in 24h format")
	assert.Equal(t, "12:00:00 AM", f.Format("%r", midnight), "Midnight in 12h format")

	// Test noon formatting
	noon := time.Date(2020, 1, 15, 12, 0, 0, 0, time.UTC)
	assert.Equal(t, "12:00:00", f.Format("%T", noon), "Noon in 24h format")
	assert.Equal(t, "12:00:00 PM", f.Format("%r", noon), "Noon in 12h format")

	// Test milliseconds/microseconds
	withMicro := time.Date(2020, 1, 15, 12, 0, 0, 123456789, time.UTC)
	assert.Equal(t, "123456", f.Format("%f", withMicro), "Microseconds formatting")
}

// TestCommonCombinations tests format string combinations commonly used in applications
func TestCommonCombinations(t *testing.T) {
	ref := time.Unix(1136239445, 456841962).UTC() // 2006-01-02 22:04:05.456841962 UTC
	f := strftime.New(language.English)

	// ISO 8601 date and variations
	assert.Equal(t, "2006-01-02", f.Format("%Y-%m-%d", ref), "ISO 8601 date format")
	assert.Equal(t, "2006-01-02T22:04:05", f.Format("%Y-%m-%dT%H:%M:%S", ref), "ISO 8601 datetime without timezone")
	assert.Equal(t, "2006-01-02T22:04:05+0000", f.Format("%Y-%m-%dT%H:%M:%S%z", ref), "ISO 8601 datetime with timezone")

	// Common log formats
	assert.Equal(t, "02/Jan/2006:22:04:05 +0000", f.Format("%d/%b/%Y:%H:%M:%S %z", ref), "Common Log Format date")
	assert.Equal(t, "Mon, 02 Jan 2006 22:04:05 +0000", f.Format("%a, %d %b %Y %H:%M:%S %z", ref), "RFC 7231 format (HTTP Date)")

	// Common display formats
	assert.Equal(t, "Jan  2, 2006", f.Format("%b %e, %Y", ref), "Common US date display format")
	assert.Equal(t, "Monday, January  2, 2006", f.Format("%A, %B %e, %Y", ref), "Long US date display format")
	assert.Equal(t, "10:04 PM", f.Format("%I:%M %p", ref), "12-hour time display")
}

// TestBoundaryConditions tests boundary time conditions
func TestBoundaryConditions(t *testing.T) {
	f := strftime.New(language.English)

	// Test Unix epoch
	epoch := time.Unix(0, 0).UTC() // 1970-01-01 00:00:00 UTC
	assert.Equal(t, "1970-01-01", f.Format("%Y-%m-%d", epoch), "Unix epoch date")
	assert.Equal(t, "00:00:00", f.Format("%H:%M:%S", epoch), "Unix epoch time")
	assert.Equal(t, "Thu Jan  1 00:00:00 1970", f.Format("%c", epoch), "Unix epoch full datetime")

	// Test year boundary transition
	yearEnd := time.Date(2020, 12, 31, 23, 59, 59, 999999999, time.UTC)
	yearStart := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, "2020-12-31 23:59:59", f.Format("%Y-%m-%d %H:%M:%S", yearEnd), "Year end")
	assert.Equal(t, "2021-01-01 00:00:00", f.Format("%Y-%m-%d %H:%M:%S", yearStart), "Year start")

	// Test very distant past and future dates
	distantPast := time.Date(-100, 1, 1, 0, 0, 0, 0, time.UTC)
	distantFuture := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)

	assert.Equal(t, "-100", f.Format("%Y", distantPast), "Distant past year")
	assert.Equal(t, "9999", f.Format("%Y", distantFuture), "Distant future year")
}

// TestLocaleDetection tests more complex locale detection scenarios
func TestLocaleDetection(t *testing.T) {
	ref := time.Unix(1136239445, 456841962).UTC()

	// Test various locale detection patterns
	localeTests := []struct {
		name          string
		acceptLang    string
		expectedTag   language.Tag
		expectedMonth string // %B for January
		expectedDay   string // %A for Monday
	}{
		{"Simple English", "en", language.English, "January", "Monday"},
		{"American English", "en-US", language.AmericanEnglish, "January", "Monday"},
		{"British English", "en-GB", language.BritishEnglish, "January", "Monday"},
		{"Simple German", "de", language.German, "Januar", "Montag"},
		{"Quality Values", "fr;q=0.8, en;q=0.7", language.French, "janvier", "lundi"},
		{"Multiple Locales with Quality", "es;q=0.5, it;q=0.9", language.Italian, "gennaio", "lunedì"},
		// Note: actual behavior depends on the implementation of strftimeLocaleMatcher
		// With the current implementation, it selects Spanish for the Brazilian Portuguese tag
		{"Complex Chain", "pt-BR, es;q=0.8, en-US;q=0.6, en;q=0.4", language.Spanish, "enero", "lunes"},
	}

	for _, tc := range localeTests {
		t.Run(tc.name, func(t *testing.T) {
			tags, _, _ := language.ParseAcceptLanguage(tc.acceptLang)

			// Create formatter with the parsed tags
			f := strftime.New(tags...)

			// Get actual formatted output
			gotMonth := f.Format("%B", ref)
			gotDay := f.Format("%A", ref)

			// Compare against expected values
			assert.Equal(t, tc.expectedMonth, gotMonth, "Month name should match expected locale")
			assert.Equal(t, tc.expectedDay, gotDay, "Day name should match expected locale")
		})
	}
}

// TestFormatErrors tests improper format strings and edge cases
func TestFormatErrors(t *testing.T) {
	f := strftime.New(language.English)
	ref := time.Unix(1136239445, 456841962).UTC()

	// Test incomplete format specifiers
	assert.Equal(t, "%", f.Format("%", ref), "Single % at end should remain as %")
	assert.Equal(t, "Test % string", f.Format("Test % string", ref), "Single % in middle should remain as %")

	// Test unknown format specifiers
	assert.Equal(t, "Test %Q string", f.Format("Test %Q string", ref), "Unknown specifier %Q should remain as %Q")

	// Test incomplete modifiers
	assert.Equal(t, "Test %E string", f.Format("Test %E string", ref), "Incomplete %E should remain as %E")
	assert.Equal(t, "Test %O string", f.Format("Test %O string", ref), "Incomplete %O should remain as %O")

	// Test multiple consecutive % signs
	assert.Equal(t, "Test % % % string", f.Format("Test % % % string", ref), "Multiple consecutive % should remain as is")
	assert.Equal(t, "Test % string", f.Format("Test %% string", ref), "%% should produce a single %")
	assert.Equal(t, "Test %% string", f.Format("Test %%% string", ref), "%%% should produce %% in the output")
}

func TestApi(t *testing.T) {
	// testing through different methods
	f := strftime.New(language.English)
	ref := time.Unix(1136239445, 456841962).UTC()
	good := `Mon Jan  2 22:04:05 2006`

	assert.Equal(t, good, strftime.Format(language.English, `%c`, ref), `testing strftime.Format`)
	assert.Equal(t, good, strftime.EnFormat(`%c`, ref), `testing for strftime.EnFormat`)

	buf := &bytes.Buffer{}
	strftime.EnFormatF(buf, `%c`, ref)
	assert.Equal(t, good, buf.String(), `testing for strftime.EnFormatF`)

	out := []byte("Test: ")
	out = f.AppendFormat(out, `%c`, ref)
	assert.Equal(t, append([]byte("Test: "), []byte(good)...), out, `testing for Formatter.AppendFormat`)

	buf = &bytes.Buffer{}
	f.FormatF(buf, `%c`, ref)
	assert.Equal(t, good, buf.String(), `testing for Formatter.FormatF`)
}
