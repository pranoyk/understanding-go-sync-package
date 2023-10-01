package synconce

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkSyncOnceFunc(b *testing.B) {
	b.ResetTimer()

	b.Run("sync.Once", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var once sync.Once
			bk := NewBookKeeper("John", "123456789", "1234567890")
			once.Do(bk.helloWorld)
		}
		b.ReportAllocs()
	})

	b.Run("sync.OnceFunc", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bk := NewBookKeeper("John", "123456789", "1234567890")
			printHelloWorld := sync.OnceFunc(bk.helloWorld)
			printHelloWorld()
		}
		b.ReportAllocs()
	})
}

func BenchmarkSyncOnceValue(b *testing.B) {
	b.ResetTimer()

	b.Run("sync.Once", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var once sync.Once
			bk := NewBookKeeper("John", "123456789", "1234567890")
			var val string
			once.Do(func() { val = bk.getName() })
			fmt.Println(val)
			b.ReportAllocs()
		}
	})

	b.Run("sync.OnceValue", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bk := NewBookKeeper("John", "123456789", "1234567890")
			onceValue := sync.OnceValue[string](bk.getName)
			fmt.Println(onceValue())
			b.ReportAllocs()
		}
	})
}

func BenchmarkSyncOnceValues(b *testing.B) {
	b.ResetTimer()

	b.Run("sync.Once", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var once sync.Once
			bk := NewBookKeeper("John", "123456789", "1234567890")
			var name, phoneNumber string
			once.Do(func() { name, phoneNumber = bk.getNameAndPhoneNumber() })
			fmt.Println(name, phoneNumber)
			b.ReportAllocs()
		}
	})

	b.Run("sync.OnceValues", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			bk := NewBookKeeper("John", "123456789", "1234567890")
			onceValues := sync.OnceValues[string, string](bk.getNameAndPhoneNumber)
			fmt.Println(onceValues())
			b.ReportAllocs()
		}
	})
}