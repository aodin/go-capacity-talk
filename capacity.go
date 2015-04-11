package capacity

import "reflect"

type simple struct {
	A0 int
	A1 struct {
		B0 int
		B1 int
	}
}

type nested struct {
	A0 int
	A1 struct {
		B0 struct {
			C0 struct {
				D0 int
				D1 int
			}
		}
	}
}

type output struct {
	name  string
	index []int
}

func recurse(index []int, elem reflect.Type) (outputs []output) {
	for i := 0; i < elem.NumField(); i += 1 {
		field := elem.Field(i)

		// If the field is a struct, resume recursion
		if field.Type.Kind() == reflect.Struct {
			outputs = append(
				outputs,
				recurse(append(index, i), field.Type)...,
			)
			continue
		}

		// Otherwise save the field
		outputs = append(outputs, output{
			name:  field.Name,
			index: append(index, i),
		})
	}
	return
}
