package strftime_test

import (
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
		{`%Oy`, `六`, ref},
		{`%OH`, `二十二`, ref},
		{`%OI`, `十`, ref},
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
