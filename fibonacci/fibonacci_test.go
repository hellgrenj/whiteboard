package fibonacci_test

import (
	"log"
	"os"
	"testing"

	"github.com/hellgrenj/whiteboard/fibonacci"
)

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}
func BenchmarkSlowFib(b *testing.B) {
	defer quiet()()
	for i := 0; i < b.N; i++ {
		fibonacci.SlowFib(30)
	}
}
func BenchmarkFasterFib(b *testing.B) {
	defer quiet()()
	for i := 0; i < b.N; i++ {
		fibonacci.FasterFib(30)
	}
}
