package iutils

import (
	"log"
	"strings"
	"testing"
)

func Test_(t *testing.T) {
	_o1_float64(t)
	_02_PreInit(t)
}

func _o1_float64(t *testing.T) {

	var s0 int
	s0 = 1234

	var s32 float32
	s32 = 1234.1256

	var s64 float64
	s64 = 1234.1256

	if s64 != AnyToFloat64("1234.1256") {
		t.Fatal("error AnyToFloat64 from str")
	}

	if 1234.0 != AnyToFloat64(s0) {
		t.Fatal("error AnyToFloat64 from int")
	}

	if 1234.0 != AnyToFloat64(&s0) {
		t.Fatal("error AnyToFloat64 from *int")
	}

	if !EqFloat64(s64, AnyToFloat64(s32), 5) {
		log.Println(AnyToFloat64(s32))
		t.Fatal("error AnyToFloat64 from float32")
	}
	if !EqFloat64(s64, AnyToFloat64(&s32), 5) {
		t.Fatal("error AnyToFloat64 from *float32")
	}

	if !EqFloat64(s64, AnyToFloat64(s64), 5) {
		t.Fatal("error AnyToFloat64 from float64")
	}
	if !EqFloat64(s64, AnyToFloat64(&s64), 5) {
		t.Fatal("error AnyToFloat64 from *float64")
	}
}

func EqFloat64(a float64, b float64, mants ...int) bool {
	mant := 5.0
	if len(mants) == 0 {
		mant = float64(mants[0])
	}

	r := mant*a - mant*b
	if -1 < r && r < 1 {
		return true
	}
	return false
}

func _02_PreInit(t *testing.T) {
	str := "123"
	if 123 != AnyToInt(str) {
		t.Fatal("error AnyToInt")
	}

	num := 123
	if "123" != AnyToString(num) {
		t.Fatal("error AnyToString")
	}

	list_interface := []interface{}{
		"123", 123, 1, "TOP",
	}

	res_s := AnyToStringArray(list_interface)
	log.Printf("AnyToStringArray: %v\n", res_s)
	if res_s[0] != "123" || res_s[1] != "123" || res_s[2] != "1" || res_s[3] != "TOP" {
		t.Fatal("error AnyToStringArray")
	}

	res_n := AnyToIntArray(list_interface)
	log.Printf("AnyToIntArray: %v\n", res_n)
	if res_n[0] != 123 || res_n[1] != 123 || res_n[2] != 1 || res_n[3] != 0 {
		t.Fatal("error AnyToStringArray")
	}

	list_interface = []interface{}{
		"123", 123, 1, "TOP",
	}

	res_a := AppendAny(list_interface, "MAIL")
	log.Printf("AppendAny: %v\n", res_a)
	if res_a[0] != "123" || res_a[1] != 123 || res_a[2] != 1 || res_a[3] != "TOP" || res_a[4] != "MAIL" {
		t.Fatal("error AppendAny")
	}

	res_str := strings.Split("A,A,B,C,D,E", ",")
	res_str_res := GrepString(res_str, "A")
	log.Printf("GrepString: %v\n", res_str_res)
	if len(res_str_res) != 2 {
		t.Fatal("error GrepString([]string, string)")
	}

	var ff = func(a string) bool {
		if a != "A" {
			return true
		}
		return false
	}

	res_str_res = GrepString(res_str, ff)
	log.Printf("GrepString: %v\n", res_str_res)
	if len(res_str_res) != 4 {
		t.Fatal("error GrepString([]string, func)")
	}

	res_str = append(res_str, "", "")
	res_str_res = GrepString(res_str)
	log.Printf("GrepString: %v\n", res_str_res)
	if len(res_str_res) != 6 {
		t.Fatal("error GrepString([]string)")
	}

	//t.Fatal("error AppendAny")

}
