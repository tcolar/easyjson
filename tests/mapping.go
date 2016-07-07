package tests

import "github.com/mailru/easyjson"

//easyjson:json
type Mapping struct {
	Field string
	Str   string
	Str2  string

	Mapping *easyjson.JsonMapping `json:"-"`
}

var mappingValue = Mapping{
	Field: "test",
	Str:   "str1",
	Str2:  "str2",
	Mapping: easyjson.Mapping().
		Omit("Str").
		Rename("Str2", "string2").
		Add("foo", "52")}

var mappingString = `{"Field":"test","string2":"str2","foo":52}`
