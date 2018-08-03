package gotest

import (
	"testing"
)


/// run : go test   -test.bench=".*"

func BenchMark_Division(b *testing.B)  {

	for i := 0; i < b.N; i ++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B)  {
	b.StopTimer()
	b.StartTimer()
	for i:=0; i <b.N; i++ {
		Division(4,5)
	}


}
