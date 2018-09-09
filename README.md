# strftime

Fast strftime in Go with [BCP 47 language tags](https://golang.org/x/text/language).

[![Build Status](https://travis-ci.org/MagicalTux/strftime.png?branch=master)](https://travis-ci.org/MagicalTux/strftime)

[![GoDoc](https://godoc.org/github.com/MagicalTux/strftime?status.svg)](https://godoc.org/github.com/MagicalTux/strftime)

# Usage

It is either possible to instanciate an object for a given locale (via `strftime.New`) or directly call strftime.Format().

For example (in French):

```go
f := strftime.New(language.French)
out.WriteString(f.Format(`pattern`, time.Now()));
```

# Description

This version of strftime for Go has multiple goals in mind:

* Support for [BCP 47 language tags](https://golang.org/x/text/language)
* Easy to use
* Ability to just return string or write to a io.Writer
* Be as complete as possible in terms of conversion specifications

## Pattern support

| pattern | description |
|:--------|:------------|
| %A      | national representation of the full weekday name |
| %a      | national representation of the abbreviated weekday |
| %B      | national representation of the full month name |
| %b      | national representation of the abbreviated month name |
| %C      | (year / 100) as decimal number; single digits are preceded by a zero |
| %c      | national representation of time and date |
| %D      | equivalent to %m/%d/%y |
| %d      | day of the month as a decimal number (01-31) |
| %e      | the day of the month as a decimal number (1-31); single digits are preceded by a blank |
| %F      | equivalent to %Y-%m-%d |
| %H      | the hour (24-hour clock) as a decimal number (00-23) |
| %h      | same as %b |
| %I      | the hour (12-hour clock) as a decimal number (01-12) |
| %j      | the day of the year as a decimal number (001-366) |
| %k      | the hour (24-hour clock) as a decimal number (0-23); single digits are preceded by a blank |
| %l      | the hour (12-hour clock) as a decimal number (1-12); single digits are preceded by a blank |
| %M      | the minute as a decimal number (00-59) |
| %m      | the month as a decimal number (01-12) |
| %n      | a newline |
| %p      | national representation of either "ante meridiem" (a.m.)  or "post meridiem" (p.m.)  as appropriate. |
| %R      | equivalent to %H:%M |
| %r      | equivalent to %I:%M:%S %p |
| %S      | the second as a decimal number (00-60) |
| %T      | equivalent to %H:%M:%S |
| %t      | a tab |
| %U      | the week number of the year (Sunday as the first day of the week) as a decimal number (00-53) |
| %u      | the weekday (Monday as the first day of the week) as a decimal number (1-7) |
| %V      | the week number of the year (Monday as the first day of the week) as a decimal number (01-53) |
| %v      | equivalent to %e-%b-%Y |
| %W      | the week number of the year (Monday as the first day of the week) as a decimal number (00-53) |
| %w      | the weekday (Sunday as the first day of the week) as a decimal number (0-6) |
| %X      | national representation of the time |
| %x      | national representation of the date |
| %Y      | the year with century as a decimal number |
| %y      | the year without century as a decimal number (00-99) |
| %Z      | the time zone name |
| %z      | the time zone offset from UTC |
| %%      | a '%' |

## Performances / other libraries

```
// On Linux Gentoo 4.14.13-gentoo
// go version go1.11 linux/amd64
$ go test -tags bench -benchmem -bench .
<snip>
BenchmarkTebeka-12        	  500000	      3780 ns/op	     288 B/op	      21 allocs/op
BenchmarkJehiah-12        	 1000000	      1549 ns/op	     256 B/op	      17 allocs/op
BenchmarkFastly-12        	  500000	      3642 ns/op	     192 B/op	      11 allocs/op
BenchmarkLestrrat-12      	 1000000	      1335 ns/op	     240 B/op	       3 allocs/op
BenchmarkMagicalTux-12    	 2000000	       804 ns/op	     227 B/op	       9 allocs/op
PASS
ok  	github.com/MagicalTux/strftime	9.141s
```

This library is much faster than other libraries for common cases. In case of format pattern re-use, [Lestrrat's implementation](github.com/lestrrat-go/strftime) is still faster (but has no locale awareness).

| Import Path                         | Score      | Note                            |
|:------------------------------------|-----------:|:--------------------------------|
| github.com/MagicalTux/strftime      | 804 ns/op  |                                 |
| github.com/lestrrat-go/strftime     | 1335 ns/op | Using `Format()` (NOT cached)   |
| github.com/jehiah/go-strftime       | 1549 ns/op |                                 |
| github.com/fastly/go-utils/strftime | 3642 ns/op | cgo version on Linux            |
| github.com/tebeka/strftime          | 3780 ns/op |                                 |

Please note that this benchmark only uses the subset of conversion specifications that are supported by *ALL* of the libraries compared.

## Internals

This strftime implementation works by writing to a io.Writer by default. If using Format instead of FormatF, it will write internally to a buffer, which is then returned as string.

Go provides two methods to accomplish this starting go 1.10: bytes.Buffer (available in earlier versions) or the new strings.Builder. In go 1.10 strings.Builder is still a bit slow, bug go 1.11 benchmark gets much closer to bytes.Buffer, with slightly less memory usage (208 B/op with strings.Builder vs 227 B/op with bytes.Buffer).

Because of this, it has been decided to stay with bytes.Buffer for now, but tests will have to be run again with new Go versions.
