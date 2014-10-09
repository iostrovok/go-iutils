package iutils

import (
	"log"
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

}
