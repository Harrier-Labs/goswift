package goswift

import (
	"testing"
)

func BenchmarkSet(b *testing.B) {
	key := "name"
	val := "leoantony"
	cache := NewCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Set(key, val, 0)
	}
}

func BenchmarkSetWithExpiry(b *testing.B) {
	key := "name"
	val := "leoantony"
	cache := NewCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Set(key, val, 10000)
	}
}

func BenchmarkGet(b *testing.B) {
	key := "name"
	val := "leoantony"
	cache := NewCache()
	cache.Set(key, val, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Get(key)
	}
}

func BenchmarkExists(b *testing.B) {
	key := "name"
	val := "leoantony"
	cache := NewCache()
	cache.Set(key, val, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Exists(key)
	}
}

func BenchmarkDel(b *testing.B) {
	key := "name"
	val := "leoantony"
	cache := NewCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Set(key, val, 0)
		cache.Del(key)
	}
}

func BenchmarkUpdate(b *testing.B) {
	key := "name"
	val := "leoantony"
	newVal := "johndoe"
	cache := NewCache()
	cache.Set(key, val, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Update(key, newVal)
	}
}

func BenchmarkHset(b *testing.B) {
	key := "user"
	field := "name"
	val := "leoantony"
	cache := NewCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.Hset(key, field, val, 0)
	}
}

func BenchmarkHGet(b *testing.B) {
	key := "user"
	field := "name"
	val := "leoantony"
	cache := NewCache()
	cache.Hset(key, field, val, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.HGet(key, field)
	}
}

func BenchmarkHGetAll(b *testing.B) {
	key := "user"
	cache := NewCache()
	cache.Hset(key, "name", "leoantony", 0)
	cache.Hset(key, "age", 25, 0)
	cache.Hset(key, "city", "New York", 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.HGetAll(key)
	}
}

type TestStruct struct {
	Name string
	Age  int
	City string
}

func BenchmarkHMset(b *testing.B) {
	key := "user"
	data := TestStruct{
		Name: "leoantony",
		Age:  25,
		City: "New York",
	}
	cache := NewCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.HMset(key, data, 0)
	}
}

func BenchmarkAllData(b *testing.B) {
	cache := NewCache()
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	cache.Set("key3", "value3", 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cache.AllData()
	}
}