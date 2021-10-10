package RWMutexSet_test

import (
	"github.com/tevkr/gbGoLangLevel2/PR5/part_3/RWMutexSet"
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
func BenchmarkRWMutexSetAdd10Percent(b *testing.B) {
	var set = RWMutexSet.NewSet()

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
func BenchmarkRWMutexSetHas90Percent(b *testing.B) {
	var set = RWMutexSet.NewSet()
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
func BenchmarkRWMutexSetAdd50Percent(b *testing.B) {
	var set = RWMutexSet.NewSet()

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
func BenchmarkRWMutexSetHas50Percent(b *testing.B) {
	var set = RWMutexSet.NewSet()
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
func BenchmarkRWMutexSetAdd90Percent(b *testing.B) {
	var set = RWMutexSet.NewSet()

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
func BenchmarkRWMutexSetHas10Percent(b *testing.B) {
	var set = RWMutexSet.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(calcPercent(10))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}