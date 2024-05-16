package main

import (
	"encoding/gob"
	"time"
)

func init() {
	gob.Register(SmallStruct{})
	gob.Register(MediumStruct{})
	gob.Register(BigStruct{})
	gob.Register([]SmallStruct{})
	gob.Register([]MediumStruct{})
	gob.Register([]BigStruct{})
	gob.Register([]*SmallStruct{})
	gob.Register([]*MediumStruct{})
	gob.Register([]*BigStruct{})
}

type SmallStruct struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func newSmallStruct(id int64) SmallStruct {
	return SmallStruct{
		ID:   id,
		Name: "name",
	}
}

type MediumStruct struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Field1    string    `json:"field1"`
	Field2    *string   `json:"field2"`
	Field3    *string   `json:"field3"`
	Field4    float64   `json:"field4"`

	SmallStruct SmallStruct `json:"small_struct"`
}

func newMediumStruct(id int64) MediumStruct {
	v := "field2"
	return MediumStruct{
		ID:          id,
		Name:        "name",
		Field2:      &v,
		SmallStruct: newSmallStruct(id),
	}
}

type BigStruct struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Field1    string    `json:"field1"`
	Field2    *string   `json:"field2"`
	Field3    *string   `json:"field3"`
	Field4    float64   `json:"field4"`

	MediumStruct1 MediumStruct  `json:"medium_struct_1"`
	MediumStruct2 MediumStruct  `json:"medium_struct_2"`
	MediumStruct3 *MediumStruct `json:"medium_struct_3"`
}

func newBigStruct(id int64) BigStruct {
	v := "field2"
	m := newMediumStruct(3)
	return BigStruct{
		ID:            id,
		Name:          "name",
		Field2:        &v,
		MediumStruct1: newMediumStruct(id),
		MediumStruct3: &m,
	}
}

func createSmallStructList(max int) []SmallStruct {
	list := make([]SmallStruct, max)
	for i := range list {
		list[i] = newSmallStruct(int64(i))
	}
	return list
}

func createSmallStructListPtr(max int) []*SmallStruct {
	list := make([]*SmallStruct, max)
	for i := range list {
		v := newSmallStruct(int64(i))
		list[i] = &v
	}
	return list
}

func createMediumStructList(max int) []MediumStruct {
	list := make([]MediumStruct, max)
	for i := range list {
		list[i] = newMediumStruct(int64(i))
	}
	return list
}

func createMediumStructListPtr(max int) []*MediumStruct {
	list := make([]*MediumStruct, max)
	for i := range list {
		v := newMediumStruct(int64(i))
		list[i] = &v
	}
	return list
}

func createBigStructList(max int) []BigStruct {
	list := make([]BigStruct, max)
	for i := range list {
		list[i] = newBigStruct(int64(i))
	}
	return list
}

func createBigStructListPtr(max int) []*BigStruct {
	list := make([]*BigStruct, max)
	for i := range list {
		v := newBigStruct(int64(i))
		list[i] = &v
	}
	return list
}
