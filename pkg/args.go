package godcmtk

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// MarshalArgs converts an args object with arg tag to a string slice.
func MarshalArgs(args interface{}) ([]string, error) {
	if args == nil {
		return []string{}, nil
	}

	t := reflect.TypeOf(args)
	v := reflect.ValueOf(args)

	// handle if args is pointer.
	if v.Kind() == reflect.Ptr {
		// return empty slice if args is nil.
		// if v.IsNil() {
		// 	return []string{}, nil
		// }

		v = reflect.Indirect(v)
		t = v.Type()
	}

	results := []string{}

	// iterate through the fields of the struct
	for i := 0; i < t.NumField(); i++ {

		// get the filed tag for arg.
		field := t.Field(i)
		tagValue := field.Tag.Get("arg")

		if tagValue == "" {
			continue
		}

		// get the field value.
		fieldValue := v.Field(i)

		switch fieldValue.Kind() {
		case reflect.Bool:
			if fieldValue.Bool() {
				results = append(results, tagValue)
			}
		case reflect.String:
			if s := fieldValue.String(); s != "" {
				results = append(results, tagValue, s)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			if fieldValue.CanInt() {
				if i := fieldValue.Int(); i != 0 {
					results = append(results, tagValue, fmt.Sprint(i))
				}
			} else if fieldValue.CanFloat() {
				if f := fieldValue.Float(); f != 0.0 {
					results = append(results, tagValue, fmt.Sprint(f))
				}
			}
		default:
			return nil, fmt.Errorf("invalid kind: %v", fieldValue.Kind())
		}

	}

	return results, nil
}

// StringArray converts []string in go to char*[] in c.
func StringArray(goStrings ...string) (**C.char, func()) {
	// allocate memory for an array of pointers to C-style strings
	cStrings := make([]*C.char, len(goStrings))

	// convert Go strings to C-style strings and store their pointers
	for i, str := range goStrings {
		cStrings[i] = C.CString(str)
	}

	// convert the array of pointers to C strings to a pointer to a pointer
	cStringArray := (**C.char)(unsafe.Pointer(&cStrings[0]))

	return cStringArray, func() {
		for i := range cStrings {
			// ensure memory is deallocated
			C.free(unsafe.Pointer(cStrings[i]))
		}
	}
}

// EmptyString creates a char *array with given length.
func EmptyString() (*C.char, func()) {
	emptyString := ""
	cEmptyString := C.CString(emptyString)
	return cEmptyString, func() {
		C.free(unsafe.Pointer(cEmptyString))
	}
}
