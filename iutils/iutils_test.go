package iutils

import (
	"log"
	"strings"
	"testing"
)

func Test_PreInit(t *testing.T) {
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
