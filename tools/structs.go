package tools

import (
	"reflect"
	"strings"
)

func StructToMap(source interface{}, onlyFields ...string) map[string]interface{} {
	var dist = map[string]interface{}{}
	t := reflect.TypeOf(source)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// silence failed
	if t.Kind() != reflect.Struct {
		return dist
	}
	v := reflect.ValueOf(source)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Name
		tag := parseJsonTag(f.Tag.Get("json"))
		if tag != "" {
			name = tag
		}
		if len(onlyFields) == 0 || ExistInArray(onlyFields, name) || ExistInArray(onlyFields, f.Name) {
			dist[name] = v.FieldByName(f.Name).Interface()
		}
	}
	return dist
}
func parseJsonTag(tag string) string {
	if tag == "-" || tag == "" {
		return ""
	}
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx]
	}
	return tag
}
