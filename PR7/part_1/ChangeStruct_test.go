package part_1_test

import (
	"github.com/tevkr/gbGoLangLevel2/PR7/part_1"
	"reflect"
	"testing"
)
func TestChangeStruct(t *testing.T) {
	s := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
		Slice       []int
		Object      struct {
			NestedField int
		}
	}{
		FieldString: "stroka",
		FieldInt:    107,
		Slice:       []int{112, 107, 207},
		Object:      struct{ NestedField int }{NestedField: 302},
	}
	// Test 1
	m := map[string]interface{}{
		"FieldString": "asd",
	}
	err := part_1.ChangeStruct(&s, m)
	if err != nil || s.FieldString != m["FieldString"] {
		t.Errorf("ChangeStruct(%v, %v) failed.", s, m)
	}
	// Test 2
	m = map[string]interface{}{
		"FieldString": "asd",
		"Object": map[string]interface{} {
			"NestedField": 2,
		},
		"steak": true,
		"FieldInt": 222,
		"Slice" : []int{1212, 107, 20117},
	}
	err = part_1.ChangeStruct(&s, m)
	if err != nil ||
		s.FieldString != m["FieldString"] ||
		s.Object.NestedField != m["Object"].(map[string]interface{})["NestedField"] ||
		s.FieldInt != m["FieldInt"] ||
		!reflect.DeepEqual(s.Slice, m["Slice"]){
		t.Errorf("ChangeStruct(%v, %v) failed.", s, m)
	}
}
