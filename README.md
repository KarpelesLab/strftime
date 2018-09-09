# strftime

Fast strftime in Go with [BCP 47 language tags](https://golang.org/x/text/language).

[![Build Status](https://travis-ci.org/MagicalTux/strftime.png?branch=master)](https://travis-ci.org/MagicalTux/strftime)

[![GoDoc](https://godoc.org/github.com/MagicalTux/strftime?status.svg)](https://godoc.org/github.com/MagicalTux/strftime)

# Usage

It is either possible to instanciate an object for a given locale (via `strftime.New`) or directly call strftime.Format().

## Examples

To simply get a timestamp in English:

```go
fmt.Printf("%s: something happened", strftime.EnFormat(`%c`, time.Now()));
```

Or get a quick result in French:

```go
fmt.Printf("%s: Quelque chose est arrivé", strftime.Format(language.French, `%c`, time.Now()));
```

Or display a result in the appropriate language for a web user:

```go
tags, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
f := strftime.New(tags...)
f.FormatF(w, `%c`, time.Now());
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
| %f      | Microseconds (6 digits) |
| %G      | Year matching the week going by ISO-8601:1988 standards |
| %g      | Two digits representation of %G |
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
| %P      | lower-case version of %p |
| %R      | equivalent to %H:%M |
| %r      | equivalent to %I:%M:%S %p |
| %S      | the second as a decimal number (00-60) |
| %s      | Unix Epoch Time timestamp (seconds since January 1st 1970) |
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

Era modifiers are available. For locales in which there is no era, normal values (without era modifier) are returned.

| pattern | description |
|:--------|:------------|
| %Ec     | national representation of time and date |
| %EC     | name of era the date is in |
| %EX     | national representation of the time |
| %Ex     | national representation of the date |
| %EY     | full era name and year represented in locale |
| %Ey     | year as decimal number in era (if any) or same as %y |

## Performances / other libraries

```
// On Linux Gentoo 4.14.13-gentoo
// go version go1.11 linux/amd64
$ go test -tags bench -benchmem -bench .
<snip>
BenchmarkCactus-12        	 1000000	      1709 ns/op	     216 B/op	       7 allocs/op
BenchmarkLeekchan-12      	  500000	      2728 ns/op	    1376 B/op	      22 allocs/op
BenchmarkTebeka-12        	  300000	      3836 ns/op	     288 B/op	      21 allocs/op
BenchmarkJehiah-12        	 1000000	      1541 ns/op	     256 B/op	      17 allocs/op
BenchmarkFastly-12        	  500000	      3735 ns/op	     192 B/op	      11 allocs/op
BenchmarkLestrrat-12      	 1000000	      1319 ns/op	     240 B/op	       3 allocs/op
BenchmarkMagicalTux-12    	 2000000	       810 ns/op	     227 B/op	       9 allocs/op
PASS
ok  	github.com/MagicalTux/strftime	11.566s
```

This library is much faster than other libraries for common cases. In case of format pattern re-use, [Lestrrat's implementation](https://github.com/lestrrat-go/strftime) is still faster (but has no locale awareness).

| Import Path                         | Score      | Note                            |
|:------------------------------------|-----------:|:--------------------------------|
| github.com/MagicalTux/strftime      | 810 ns/op  |                                 |
| github.com/lestrrat-go/strftime     | 1319 ns/op | Using `Format()` (NOT cached)   |
| github.com/jehiah/go-strftime       | 1541 ns/op |                                 |
| github.com/cactus/gostrftime        | 1709 ns/op |                                 |
| github.com/leekchan/timeutil        | 2728 ns/op |                                 |
| github.com/fastly/go-utils/strftime | 3735 ns/op | cgo version on Linux            |
| github.com/tebeka/strftime          | 3836 ns/op |                                 |

Please note that this benchmark only uses the subset of conversion specifications that are supported by *ALL* of the libraries compared.

## Internals

This strftime implementation works by writing to a io.Writer by default. If using Format instead of FormatF, it will write internally to a buffer, which is then returned as string.

Go provides two methods to accomplish this starting go 1.10: bytes.Buffer (available in earlier versions) or the new strings.Builder. In go 1.10 strings.Builder is still a bit slow, bug go 1.11 benchmark gets much closer to bytes.Buffer, with slightly less memory usage (208 B/op with strings.Builder vs 227 B/op with bytes.Buffer).

Because of this, it has been decided to stay with bytes.Buffer for now, but tests will have to be run again with new Go versions.
