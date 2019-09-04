package go_utils

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"unsafe"
)

func StringRemoveAllNum(string2 string) string {
	re, _ := regexp.Compile("[0-9]")
	return re.ReplaceAllString(string2, "")
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func InterFace2Int64(input interface{}) (int64, error) {
	switch input.(type) {
	case int64:
		return input.(int64), nil
	default:
		return 0, fmt.Errorf("input: %v not int", input)
	}
}

func InterFace2String(input interface{}) (string, error) {
	switch input.(type) {
	case string:
		return input.(string), nil
	default:
		return "", fmt.Errorf("input: %v not string", input)
	}
}

func RemoveDuplicateElement(addrs []int64) []int64 {
	result := make([]int64, 0, len(addrs))
	temp := map[int64]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
