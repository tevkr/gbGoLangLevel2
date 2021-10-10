package MutexSet_test

import (
	"github.com/tevkr/gbGoLangLevel2/PR5/part_3/MutexSet"
	"testing"
)

var testCount = 1000

func calcPercent(percent int) int {
	if percent <= 0 {
		return 0
	}
	return int(float32(testCount) * (float32(percent) / 100))
}

// 10% Запись
func BenchmarkMutexSetAdd10Percent(b *testing.B) {
	var set = MutexSet.NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(10))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

// 90% чтение
func BenchmarkMutexSetHas90Percent(b *testing.B) {
	var set = MutexSet.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(90))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

// 50% Запись
func BenchmarkMutexSetAdd50Percent(b *testing.B) {
	var set = MutexSet.NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(50))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

// 50% чтение
func BenchmarkMutexSetHas50Percent(b *testing.B) {
	var set = MutexSet.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(50))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

// 90% Запись
func BenchmarkMutexSetAdd90Percent(b *testing.B) {
	var set = MutexSet.NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(90))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

// 10% чтение
func BenchmarkMutexSetHas10Percent(b *testing.B) {
	var set = MutexSet.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(10))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}