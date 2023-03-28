package service

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/nub06/go-jsoncompare/conf"
	"github.com/nub06/go-jsoncompare/util"
)

func Run() {

	firstData := util.ReadFromFile(conf.FirstInput)
	secData := util.ReadFromFile(conf.SecondInput)

	k := util.ParseJson([]byte(firstData))
	v := util.ParseJson([]byte(secData))

	diffs := compare(k, v)

	fmt.Println(diffs)

}

func compare(obj1 interface{}, obj2 interface{}) []string {

	var diffList []string
	if reflect.DeepEqual(obj1, obj2) {
		diffList = append(diffList, "JSONs are equal")
	} else {
		diff := findDiff("", obj1, obj2)
		diffList = append(diffList, diff...)
	}

	return diffList
}

func findDiff(path string, obj1 interface{}, obj2 interface{}) []string {
	v1 := reflect.ValueOf(obj1)
	v2 := reflect.ValueOf(obj2)
	firstInp := filepath.Base(conf.FirstInput)
	secInp := filepath.Base(conf.SecondInput)

	var diffs []string

	if v1.Kind() != v2.Kind() {
		firstChar := path[0]
		if string(firstChar) == "." {
			path = strings.TrimPrefix(path, string(path[0]))
		}
		res := fmt.Sprintf("Types are different: %v=%v type is %v on %v %v=%v type is %v on %v\n", path, (v1), v1.Kind(), firstInp, path, (v2), v2.Kind(), secInp)
		diffs = append(diffs, res)

	}

	if v1.Kind() == reflect.Map {

		for _, key := range v1.MapKeys() {
			val1 := v1.MapIndex(key)
			val2 := v2.MapIndex(key)
			if !val2.IsValid() {
				keyValue := fmt.Sprintf("%v", key.Interface())
				newPath := fmt.Sprintf("%s.%s", path, keyValue)
				firstChar := newPath[0]
				if string(firstChar) == "." {
					newPath = strings.TrimPrefix(newPath, string(newPath[0]))
				}
				if !strings.Contains(path, ".") {

					res := fmt.Sprintf("Group is missing: %s is missing from %s \n", newPath, secInp)
					diffs = append(diffs, res)

				} else {
					res := fmt.Sprintf("Key is missing: %s is missing from %s \n", newPath, secInp)
					diffs = append(diffs, res)
				}

				continue
			}
			keyValue := fmt.Sprintf("%v", key.Interface())
			newPath := fmt.Sprintf("%s.%s", path, keyValue)
			diffs = append(diffs, findDiff(newPath, val1.Interface(), val2.Interface())...)
		}
		for _, key := range v2.MapKeys() {
			if !v1.MapIndex(key).IsValid() {
				keyValue := fmt.Sprintf("%v", key.Interface())
				newPath := fmt.Sprintf("%s.%s", path, keyValue)
				firstChar := newPath[0]
				if string(firstChar) == "." {
					newPath = strings.TrimPrefix(newPath, string(newPath[0]))
				}
				if !strings.Contains(path, ".") {

					res := fmt.Sprintf("Group is missing: %s is missing from %s \n", newPath, firstInp)
					diffs = append(diffs, res)
				} else {
					res := fmt.Sprintf("Key is missing: %s is missing from %s \n", newPath, firstInp)
					diffs = append(diffs, res)
				}

			}
		}
	} else {
		if !reflect.DeepEqual(obj1, obj2) {
			firstChar := path[0]
			if string(firstChar) == "." {
				path = strings.TrimPrefix(path, string(path[0]))
			}
			res := fmt.Sprintf("Values are different: %s=%v on %s %s=%v on %s\n", path, obj1, firstInp, path, obj2, secInp)
			diffs = append(diffs, res)
		}
	}

	return diffs
}
