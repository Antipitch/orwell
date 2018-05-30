package orwell

import (
	"fmt"
	"reflect"
)

// NOE func
func NOE(value interface{}) bool {
	if _, isNil := IsNil(value); isNil {
		return true
	}

	if IsEmpty(value) {
		return true
	}

	return false
}

// IsNil func
func IsNil(value interface{}) (interface{}, bool) {
	if value == nil {
		return nil, true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Invalid:
		return nil, true
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return nil, true
		}
		return IsNil(v.Elem().Interface())
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Func, reflect.Chan:
		if v.IsNil() {
			return nil, true
		}
	}

	return value, false
}

// IsEmpty func
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {

	case reflect.Invalid:
		return true
	case reflect.Bool:
		return !v.Bool()
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return IsEmpty(v.Elem().Interface())
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

// ToString func
func ToString(value interface{}) (string, error) {
	if str, ok := value.(string); ok {
		return str, nil
	}

	v := reflect.ValueOf(value)
	if v.Type() == reflect.TypeOf([]byte(nil)) {
		return string(v.Interface().([]byte)), nil
	}

	return "", fmt.Errorf("Argument could not be converted to string")
}

// ToInt64 func
func ToInt64(v interface{}) (int64, error) {
	switch v.(type) {
	case int:
		return int64(v.(int)), nil
	case int8:
		return int64(v.(int8)), nil
	case int16:
		return int64(v.(int16)), nil
	case int32:
		return int64(v.(int32)), nil
	case int64:
		return v.(int64), nil
	}

	return 0, fmt.Errorf("Argument could not be converted to int64")
}

// ToUInt64 func
func ToUInt64(v interface{}) (uint64, error) {
	switch v.(type) {
	case uint:
		return uint64(v.(uint)), nil
	case uint8:
		return uint64(v.(uint8)), nil
	case uint16:
		return uint64(v.(uint16)), nil
	case uint32:
		return uint64(v.(uint32)), nil
	case uint64:
		return v.(uint64), nil
	}

	return 0, fmt.Errorf("Argument could not be converted to uint64")
}

// ToFloat64 func
func ToFloat64(v interface{}) (float64, error) {
	switch v.(type) {
	case float32:
		return float64(v.(float32)), nil
	case float64:
		return v.(float64), nil
	}

	return 0, fmt.Errorf("Argument could not be converted to float64")
}

// LengthOf func
func LengthOf(value interface{}) (int, error) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return v.Len(), nil
	}
	return 0, fmt.Errorf("Could not determine length of argument type %v", v.Kind())
}

func DeepEqual(s interface{}, c interface{}) bool {
	return reflect.DeepEqual(s, c)
}
