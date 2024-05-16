package main

import "testing"

func BenchmarkEncodeMap(b *testing.B) {
	benchmarks := []struct {
		name    string
		keySize int
		fn      func(interface{}) []byte
	}{
		{"algo=Gob:key=10", 10, encodeGob},
		{"algo=JSON:key=10", 10, encodeJSON},
		{"algo=Gob:key=50", 50, encodeGob},
		{"algo=JSON:key=50", 50, encodeJSON},
		{"algo=Gob:key=100", 100, encodeGob},
		{"algo=JSON:key=100", 100, encodeJSON},
		{"algo=Gob:key=1000", 1000, encodeGob},
		{"algo=JSON:key=1000", 1000, encodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := createMap(bm.keySize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.fn(m)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeMap(b *testing.B) {
	benchmarks := []struct {
		name    string
		keySize int
		encode  func(interface{}) []byte
		decode  func([]byte, interface{})
	}{
		{"algo=Gob:key=10", 10, encodeGob, decodeGob},
		{"algo=JSON:key=10", 10, encodeJSON, decodeJSON},
		{"algo=Gob:key=50", 50, encodeGob, decodeGob},
		{"algo=JSON:key=50", 50, encodeJSON, decodeJSON},
		{"algo=Gob:key=100", 100, encodeGob, decodeGob},
		{"algo=JSON:key=100", 100, encodeJSON, decodeJSON},
		{"algo=Gob:key=1000", 1000, encodeGob, decodeGob},
		{"algo=JSON:key=1000", 1000, encodeJSON, decodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := createMap(bm.keySize)
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result map[int64]float64
				bm.decode(byt, &result)
			}
		})
	}
}

func BenchmarkEncodeSliceMap(b *testing.B) {
	benchmarks := []struct {
		name    string
		keySize int
		lenSize int
		fn      func(interface{}) []byte
	}{
		{"algo=Gob:key=10:len=10", 10, 10, encodeGob},
		{"algo=JSON:key=10:len=10", 10, 10, encodeJSON},
		{"algo=Gob:key=10:len=1000", 10, 1000, encodeGob},
		{"algo=JSON:key=10:len=1000", 10, 1000, encodeJSON},
		{"algo=Gob:key=100:len=10", 100, 100, encodeGob},
		{"algo=JSON:key=100:len=10", 100, 100, encodeJSON},
		{"algo=Gob:key=1000:len=10", 1000, 1000, encodeGob},
		{"algo=JSON:key=1000:len=10", 1000, 1000, encodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := createSliceMap(bm.keySize, bm.lenSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.fn(m)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeSliceMap(b *testing.B) {
	benchmarks := []struct {
		name    string
		keySize int
		lenSize int
		encode  func(v interface{}) []byte
		decode  func([]byte, interface{})
	}{
		{"algo=Gob:key=10:len=10", 10, 10, encodeGob, decodeGob},
		{"algo=JSON:key=10:len=10", 10, 10, encodeJSON, decodeJSON},
		{"algo=Gob:key=10:len=1000", 10, 1000, encodeGob, decodeGob},
		{"algo=JSON:key=10:len=1000", 10, 1000, encodeJSON, decodeJSON},
		{"algo=Gob:key=100:len=10", 100, 100, encodeGob, decodeGob},
		{"algo=JSON:key=100:len=10", 100, 100, encodeJSON, decodeJSON},
		{"algo=Gob:key=1000:len=10", 1000, 1000, encodeGob, decodeGob},
		{"algo=JSON:key=1000:len=10", 1000, 1000, encodeJSON, decodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := createSliceMap(bm.keySize, bm.lenSize)
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []map[int64]float64
				bm.decode(byt, &result)
			}
		})
	}
}
