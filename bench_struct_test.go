package main

import "testing"

func BenchmarkEncodeSingleStruct(b *testing.B) {
	benchmarks := []struct {
		name       string
		structData interface{}
		fn         func(interface{}) []byte
	}{
		{"algo=Gob:struct_size=small", newSmallStruct(1), encodeGob},
		{"algo=JSON:struct_size=small", newSmallStruct(1), encodeJSON},
		{"algo=Gob:struct_size=medium", newMediumStruct(1), encodeGob},
		{"algo=JSON:struct_size=medium", newMediumStruct(1), encodeJSON},
		{"algo=Gob:struct_size=big", newBigStruct(1), encodeGob},
		{"algo=JSON:struct_size=big", newBigStruct(1), encodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.fn(m)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeSingleStruct(b *testing.B) {
	type benchmark struct {
		name       string
		structData interface{}
		encode     func(interface{}) []byte
		decode     func([]byte, interface{})
	}

	benchmarks := []benchmark{
		{"algo=Gob:struct_size=small", newSmallStruct(1), encodeGob, decodeGob},
		{"algo=JSON:struct_size=small", newSmallStruct(1), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result SmallStruct
				bm.decode(byt, &result)
			}
		})
	}

	benchmarks = []benchmark{
		{"algo=Gob:struct_size=medium", newMediumStruct(1), encodeGob, decodeGob},
		{"algo=JSON:struct_size=medium", newMediumStruct(1), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result MediumStruct
				bm.decode(byt, &result)
			}
		})
	}

	benchmarks = []benchmark{
		{"algo=Gob:struct_size=big", newBigStruct(1), encodeGob, decodeGob},
		{"algo=JSON:struct_size=big", newBigStruct(1), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result BigStruct
				bm.decode(byt, &result)
			}
		})
	}
}

func BenchmarkEncodeSliceStruct(b *testing.B) {
	benchmarks := []struct {
		name       string
		structData interface{}
		fn         func(interface{}) []byte
	}{
		{"algo=Gob:struct_size=small:len=100", createSmallStructList(100), encodeGob},
		{"algo=JSON:struct_size=small:len=100", createSmallStructList(100), encodeJSON},
		{"algo=Gob:struct_size=medium:len=100", createMediumStructList(100), encodeGob},
		{"algo=JSON:struct_size=medium:len=100", createMediumStructList(100), encodeJSON},
		{"algo=Gob:struct_size=big:len=100", createBigStructList(100), encodeGob},
		{"algo=JSON:struct_size=big:len=100", createBigStructList(100), encodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.fn(m)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeSliceStruct(b *testing.B) {
	type benchmark struct {
		name       string
		structData interface{}
		encode     func(interface{}) []byte
		decode     func([]byte, interface{})
	}

	benchmarks := []benchmark{
		{"algo=Gob:struct_size=small:len=100", createSmallStructList(100), encodeGob, decodeGob},
		{"algo=JSON:struct_size=small:len=100", createSmallStructList(100), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []SmallStruct
				bm.decode(byt, &result)
			}
		})
	}

	benchmarks = []benchmark{
		{"algo=Gob:struct_size=medium:len=100", createMediumStructList(100), encodeGob, decodeGob},
		{"algo=JSON:struct_size=medium:len=100", createMediumStructList(100), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []MediumStruct
				bm.decode(byt, &result)
			}
		})
	}

	benchmarks = []benchmark{
		{"algo=Gob:struct_size=big:len=100", createBigStructList(100), encodeGob, decodeGob},
		{"algo=JSON:struct_size=big:len=100", createBigStructList(100), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []BigStruct
				bm.decode(byt, &result)
			}
		})
	}
}

func BenchmarkEncodeSliceStructPointer(b *testing.B) {
	benchmarks := []struct {
		name       string
		structData interface{}
		fn         func(interface{}) []byte
	}{
		{"algo=Gob:struct_size=small:len=100", createSmallStructListPtr(100), encodeGob},
		{"algo=JSON:struct_size=small:len=100", createSmallStructListPtr(100), encodeJSON},
		{"algo=Gob:struct_size=medium:len=100", createMediumStructListPtr(100), encodeGob},
		{"algo=JSON:struct_size=medium:len=100", createMediumStructListPtr(100), encodeJSON},
		{"algo=Gob:struct_size=big:len=100", createBigStructListPtr(100), encodeGob},
		{"algo=JSON:struct_size=big:len=100", createBigStructListPtr(100), encodeJSON},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.fn(m)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeSliceStructPointer(b *testing.B) {
	type benchmark struct {
		name       string
		structData interface{}
		encode     func(interface{}) []byte
		decode     func([]byte, interface{})
	}

	benchmarks := []benchmark{
		{"algo=Gob:struct_size=small:len=100", createSmallStructListPtr(100), encodeGob, decodeGob},
		{"algo=JSON:struct_size=small:len=100", createSmallStructListPtr(100), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*SmallStruct
				bm.decode(byt, &result)
			}
		})
	}

	benchmarks = []benchmark{
		{"algo=Gob:struct_size=medium:len=100", createMediumStructListPtr(100), encodeGob, decodeGob},
		{"algo=JSON:struct_size=medium:len=100", createMediumStructListPtr(100), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*MediumStruct
				bm.decode(byt, &result)
			}
		})
	}

	benchmarks = []benchmark{
		{"algo=Gob:struct_size=big:len=100", createBigStructListPtr(100), encodeGob, decodeGob},
		{"algo=JSON:struct_size=big:len=100", createBigStructListPtr(100), encodeJSON, decodeJSON},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			m := bm.structData
			byt := bm.encode(m)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var result []*BigStruct
				bm.decode(byt, &result)
			}
		})
	}
}
