// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package structreflectexample

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

// Reflect returns the fields-value pairs of a struct
// in the string representation.
func Reflect(v any) (string, error) {
	if v == nil {
		return "", nil
	}

	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		return "", errors.New("not struct")
	}

	num := val.NumField()
	names := make([]string, 0, num)
	values := make([]any, 0, num)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if !field.CanInterface() {
			num--
			continue
		}

		names = append(names, val.Type().Field(i).Name)
		values = append(values, field.Interface())
	}

	if num == 0 {
		return "", nil
	}

	var buf bytes.Buffer

	buf.WriteString(val.Type().String() + "{")

	for i := 0; i < num; i++ {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(names[i] + ":")
		buf.WriteString(fmt.Sprint(values[i]))
	}

	buf.WriteString("}")

	return buf.String(), nil
}
