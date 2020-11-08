package template

import (
	"reflect"
	"text/template"
	"time"

	"goby/pkg/util"

	log "github.com/sirupsen/logrus"
)

// var (
// 	// Context :
// 	Context map[string]interface{}
// )

// Data :
type Data map[string]interface{}

func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func divide(i, j int) int {
	return i / j
}

func mod(i, j int) int {
	return i % j
}

func slicy(number int) []int {
	var s []int
	for i := 0; i < number; i++ {
		s = append(s, i)
	}
	return s
}

func hasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

func truncate(s string, i int, dots bool) string {
	return util.Truncate(s, i, dots)
}

func extractHTMLText(html string) string {
	return util.ExtractHTMLText(html, 0)
}

func ymd2USDateFormat(ymd string) string {
	t, err := time.Parse("2006-01-02", ymd)
	if err != nil {
		log.Error(err)
		return ""
	}
	return t.Format("January 2, 2006")
}

func formatNumber(number interface{}) string {
	// number.(int64)
	var numberReal int64
	switch number.(type) {
	case int:
		numberReal = int64(number.(int))
	case int8:
		numberReal = int64(number.(int8))
	case int16:
		numberReal = int64(number.(int16))
	case int32:
		numberReal = int64(number.(int32))
	case int64:
		numberReal = number.(int64)
	case uint:
		numberReal = int64(number.(uint))
	case uint8:
		numberReal = int64(number.(uint8))
	case uint16:
		numberReal = int64(number.(uint16))
	case uint32:
		numberReal = int64(number.(uint32))
	case uint64:
		numberReal = int64(number.(uint64))
	case float32:
		numberReal = int64(number.(float32))
	case float64:
		numberReal = int64(number.(float64))
	default:
		log.Error("Unsupported type => " + reflect.TypeOf(number).Kind().String())
	}

	return util.FormatNumber2String(numberReal)
}

func funcMap() map[string]interface{} {
	return template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"add":              add,
		"sub":              sub,
		"divide":           divide,
		"mul":              mul,
		"mod":              mod,
		"slicy":            slicy,
		"hasField":         hasField,
		"truncate":         truncate,
		"formatNumber":     formatNumber,
		"extractHTMLText":  extractHTMLText,
		"ymd2USDateFormat": ymd2USDateFormat,
	}
}

// New :
func New(filenames ...string) (*template.Template, error) {
	return template.New("tmpl").Funcs(funcMap()).ParseFiles(filenames...)
}
