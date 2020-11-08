package util

import (
	"strconv"
	"strings"

	"goby/pkg/dict"
)

// IntSliceToString :
func IntSliceToString(values []int, delimeters ...string) string {
	var valuesStrList []string
	for _, v := range values {
		valuesStrList = append(valuesStrList, strconv.FormatInt(int64(v), 10))
	}

	return StringSliceToStringWithDelimeters(valuesStrList, delimeters...)
}

// Int32SliceToStringSlice :
func Int32SliceToStringSlice(values []int32) []string {
	var valuesStrList []string
	for _, v := range values {
		valuesStrList = append(valuesStrList, strconv.FormatInt(int64(v), 10))
	}

	return valuesStrList
}

// Int32SliceToString :
func Int32SliceToString(values []int32, delimeters ...string) string {
	var valuesStrList []string
	for _, v := range values {
		valuesStrList = append(valuesStrList, strconv.FormatInt(int64(v), 10))
	}

	return StringSliceToStringWithDelimeters(valuesStrList, delimeters...)
}

// Int64SliceToStringSlice :
func Int64SliceToStringSlice(values []int64) []string {
	var valuesStrList []string
	for _, v := range values {
		valuesStrList = append(valuesStrList, strconv.FormatInt(v, 10))
	}

	return valuesStrList
}

// Int64SliceToString :
func Int64SliceToString(values []int64, delimeters ...string) string {
	var valuesStrList []string
	for _, v := range values {
		valuesStrList = append(valuesStrList, strconv.FormatInt(v, 10))
	}

	return StringSliceToStringWithDelimeters(valuesStrList, delimeters...)
}

// StringSliceToInt64 :
func StringSliceToInt64(values []string) ([]int64, error) {
	var vals []int64
	for _, v := range values {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

// StringSliceToStringWithDelimeters :
func StringSliceToStringWithDelimeters(vals []string, delimeters ...string) string {
	var delimetersCom string
	for _, v := range delimeters {
		delimetersCom += v
	}
	return strings.Join(vals, delimetersCom)
}

// StringSliceToStringWithMonoQuotesAndDelimeters :
func StringSliceToStringWithMonoQuotesAndDelimeters(vals []string, delimeters ...string) string {
	var newVals []string
	for _, v := range vals {
		v = strings.TrimSpace(v)
		if !strings.EqualFold(v, dict.Blank) {
			newVals = append(newVals, dict.MonoQuote+v+dict.MonoQuote)
		}
	}
	var delimetersCom string
	for _, v := range delimeters {
		delimetersCom += v
	}
	return strings.Join(newVals, delimetersCom)
}

// TrimAndEqual :
func TrimAndEqual(toBeTrimmed, expected string) bool {
	return strings.EqualFold(strings.TrimSpace(toBeTrimmed), expected)
}
