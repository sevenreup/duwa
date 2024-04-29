package object

import "testing"

func TestStringMapKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}
	if hello1.MapKey() != hello2.MapKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.MapKey() != diff2.MapKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.MapKey() == diff1.MapKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
