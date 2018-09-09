package strftime_test

import (
	"testing"
	"time"

	"golang.org/x/text/language"

	"github.com/MagicalTux/strftime"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	var ref = time.Unix(1136239445, 456841962).UTC()
	f := strftime.New(language.English)

	cmp := []struct{ A, B string }{
		{`%A`, `Monday`},
		{`%a`, `Mon`},
		{`%B`, `January`},
		{`%C`, `20`},
		{`%c`, `Mon Jan  2 22:04:05 2006`},
		{`%D`, `01/02/06`},
		{`%d`, `02`},
		{`%e`, ` 2`},
		{`%f`, `456841`},
		{`%F`, `2006-01-02`},
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
		{`%R`, `22:04`},
		{`%r`, `10:04:05 PM`},
		{`%S`, `05`},
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

		// full test from https://github.com/lestrrat-go/strftime
		{`%A %a %B %b %C %c %D %d %e %F %H %h %I %j %k %l %M %m %n %p %R %r %S %T %t %U %u %V %v %W %w %X %x %Y %y %Z %z`, "Monday Mon January Jan 20 Mon Jan  2 22:04:05 2006 01/02/06 02  2 2006-01-02 22 Jan 10 002 22 10 04 01 \n PM 22:04 10:04:05 PM 05 22:04:05 \t 01 1 01  2-Jan-2006 01 1 22:04:05 01/02/06 2006 06 UTC +0000"},
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, ref), `matching for `+x.A)
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
	}

	for _, x := range cmp {
		assert.Equal(t, x.B, f.Format(x.A, x.T), `matching for `+x.A)
	}
}
