package capacity

import "reflect"

func recurse(index []int, elem reflect.Type) (out []output) {
	for i := 0; i < elem.NumField(); i += 1 {
		field := elem.Field(i)

		// If the field is a struct, resume recursion
		if field.Type.Kind() == reflect.Struct {
			out = append(out, recurse(append(index, i), field.Type)...)
			continue
		}

		// Otherwise save the field
		out = append(out, output{name: field.Name, index: append(index, i)})
	}
	return
}

type output struct {
	name  string
	index []int
}
