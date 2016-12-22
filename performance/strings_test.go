package performance

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	strs = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}
)

func StringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < len(strs); j++ {
			strings.Join([]string{s, strs[j]}, "")
		}
	}
}

func StringsPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < len(strs); j++ {
			s += strs[j]
		}
	}
}

func BytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < len(strs); j++ {
			buf.WriteString(strs[j])
		}
	}
}

func StringsFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < len(strs); j++ {
			s += (strs[j])
		}
	}
}

func TestMain(t *testing.T) {

	fmt.Printf("%-20s:%s\n", "strings.Join", testing.Benchmark(StringsJoin))
	fmt.Printf("%-20s:%s\n", "bytes.Buffer", testing.Benchmark(BytesBuffer))
	fmt.Printf("%-20s:%s\n", "StringsPlus", testing.Benchmark(StringsPlus))
	fmt.Printf("%-20s:%s\n", "StringsFormat", testing.Benchmark(StringsFormat))
}
