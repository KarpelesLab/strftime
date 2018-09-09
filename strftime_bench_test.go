// +build bench

package strftime_test

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"

	magicaltux "github.com/MagicalTux/strftime"
	cactus "github.com/cactus/gostrftime"
	fastly "github.com/fastly/go-utils/strftime"
	jehiah "github.com/jehiah/go-strftime"
	leekchan "github.com/leekchan/timeutil"
	lestrrat "github.com/lestrrat-go/strftime"
	tebeka "github.com/tebeka/strftime"
)

func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
}

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
	// Grr, uses byte slices, and does it faster, but with more allocs
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

func BenchmarkMagicalTux(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		magicaltux.EnFormat(benchfmt, t)
	}
}
