package iutils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var NonDigitalRE = regexp.MustCompile(`[^0-9,\.]+`)

var rnd_letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rnd_letters[rand.Intn(len(rnd_letters))]
	}
	return string(b)
}

func AnyToStringArray(s interface{}) []string {
	out := make([]string, 0)
	switch v := s.(type) {
	case string:
		out = append(out, s.(string))
		return out
	case []string:
		return s.([]string)
	case []interface{}:
		for _, v := range s.([]interface{}) {
			out = append(out, AnyToString(v))
		}
		return out
	case []int:
		for _, v := range s.([]int) {
			out = append(out, AnyToString(v))
		}
		return out
	case nil:
		return out
	default:
		log.Printf("AnyToStringArray. unknown type %s\n", v)
		out = append(out, AnyToString(s))
		return out
	}
	return out
}

func AnyToString(s interface{}) string {
	switch v := s.(type) {
	case string:
		return s.(string)
	case []string:
		return strings.Join(s.([]string), "")
	case []uint8:
		raw := s.([]uint8)
		return string(raw)
	case float64:
		//return strconv.FormatFloat(s.(float64), 'f', 6, 64)
		return fmt.Sprintf("%f", s.(float64))
	case int:
		return strconv.Itoa(s.(int))
		//	case int32:
		//		return strconv.Itoa(s.(int32))
	case int64:
		return strconv.FormatUint(uint64(s.(int64)), 10)
		//	case *int32:
		//		return strconv.Itoa(*s.(*int32))
	case *int64:
		return strconv.FormatUint(uint64(*s.(*int64)), 10)
	case *int:
		return strconv.Itoa(*s.(*int))
	case nil:
		return ""
	default:
		log.Fatalf("AnyToString. unknown type '%t' => '%s'\v", s, v)
	}
	return ""
}

func AnyToIntArray(s interface{}) []int {
	out := []int{}
	switch v := s.(type) {
	case int:
		out = append(out, s.(int))
		return out
	case []int:
		return s.([]int)
	case []interface{}:
		for _, v := range s.([]interface{}) {
			out = append(out, AnyToInt(v))
		}
		return out
	case []string:
		for _, v := range s.([]string) {
			out = append(out, AnyToInt(v))
		}
		return out
	case nil:
		return out
	default:
		log.Printf("AnyToIntArray. unknown type %s\n", v)
		out = append(out, AnyToInt(s))
		return out
	}
	return out
}

func AnyToBoolInt(s interface{}) int {
	i := AnyToInt(s)
	if i == 0 {
		return 0
	}
	return 1
}

func AnyToInt(s interface{}, minmax ...int) int {
	i := _AnyToInt(s)
	if len(minmax) == 0 {
		return i
	}
	if len(minmax) > 0 && i < minmax[0] {
		return minmax[0]
	}
	if len(minmax) > 1 && i > minmax[1] {
		return minmax[1]
	}
	return i
}

func _AnyToInt(s interface{}) int {

	if s == nil {
		return 0
	}

	switch v := s.(type) {
	case string:
		st := NonDigitalRE.ReplaceAllString(s.(string), "")
		if st == "" {
			return 0
		}
		i, err := strconv.Atoi(st)
		if err != nil {
			log.Println(err)
			return 0
		}
		return int(i)
	case []uint8:
		raw := s.([]uint8)
		i, err := strconv.Atoi(string(raw))
		if err != nil {
			log.Println(err)
			return 0
		}
		return int(i)

	case int:
		return int(s.(int))
	case int32:
		return int(s.(int32))
	case int64:
		return int(s.(int64))
	case *int32:
		return int(*s.(*int32))
	case *int64:
		return int(*s.(*int64))
	case *int:
		return int(*s.(*int))

	case float64:
		return int(s.(float64))

	default:
		log.Fatalf("AnyToInt. unknown type %s\v", v)
	}
	return 0
}

func MakeJext(data ...interface{}) map[string]interface{} {
	out := map[string]interface{}{}
	for _, v := range data {
		line := AnyToString(v)
		var res map[string]interface{}
		raw := []byte(line)
		err := json.Unmarshal(raw, &res)
		if err != nil {
			log.Fatal("error:", err)
		}
		for k, v := range res {
			out[k] = v
		}
	}
	return out
}

func GetPath(data interface{}, path string) interface{} {
	paths := strings.Split(path, "/")
	return inGetPath(data, paths)
}

func inGetPath(data interface{}, paths []string) interface{} {

	if len(paths) == 0 {
		return data
	}

	path := paths[0]
	paths = paths[1:]
	switch t := data.(type) {
	case map[string]interface{}:
		if path == "*" {
			out := map[string]interface{}{}
			for k, v := range data.(map[string]interface{}) {
				out[k] = inGetPath(v, paths)
			}
			return out
		}
		if val, find := data.(map[string]interface{})[path]; find {
			return inGetPath(val, paths)
		}
	case []interface{}:
		d := data.([]interface{})
		if path == "*" {
			out := make([]interface{}, len(d))
			for k, v := range d {
				out[k] = inGetPath(v, paths)
			}
			return out
		}

		i := AnyToInt(path)
		if i < len(d) {
			return inGetPath(d[i], paths)
		}
	default:
		log.Printf("check type %s\n", t)
		return data

	}
	return nil
}

func GetKey(key string, p interface{}) interface{} {

	switch p.(type) {
	case map[string]interface{}:
		s := p.(map[string]interface{})
		return s[key]
	}
	return nil
}

func StringArrayToInterface(list []string) []interface{} {
	out := make([]interface{}, len(list))
	for i, _ := range list {
		out[i] = list[i]
	}
	return out
}

func GetKeyKey(key1 string, key2 string, p interface{}) interface{} {
	return GetKey(key2, GetKey(key1, p))
}
