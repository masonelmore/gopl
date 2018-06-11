package main

import (
	"math/rand"
	"testing"
	"time"
)

var minStrLen, maxStrLen = 1, 10
var benchmarks = []struct {
	name string
	args []string
}{
	{"0", []string{}},
	{"10", buildRandomArgs(10, minStrLen, maxStrLen)},
	{"100", buildRandomArgs(100, minStrLen, maxStrLen)},
	{"1000", buildRandomArgs(1000, minStrLen, maxStrLen)},
	{"10000", buildRandomArgs(10000, minStrLen, maxStrLen)},
	{"100000", buildRandomArgs(100000, minStrLen, maxStrLen)},
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
var result string

func BenchmarkEchoConcat(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = echoConcat(bm.args)
			}
		})
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = echoJoin(bm.args)
			}
		})
	}
}

func _string(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func buildRandomArgs(count, minStrLen, maxStrLen int) []string {
	args := make([]string, count)
	for i := range args {
		n := rand.Intn(maxStrLen-minStrLen) + minStrLen
		args[i] = _string(n)
	}
	return args
}
