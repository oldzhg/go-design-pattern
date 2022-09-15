package singleton_test

import (
	"github.com/stretchr/testify/assert"
	"go-design-pattern/singleton"
	"testing"
)

func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, singleton.GetLazySingleton(), singleton.GetLazySingleton())
}
func BenchmarkGetLazySingletonParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazySingleton() != singleton.GetLazySingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
