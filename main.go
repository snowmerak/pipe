package pipe

import (
	"reflect"
)

const errorTypeName = "error"

func Link[T any](funs ...interface{}) func(in ...interface{}) T {
	wrapped := func(in ...interface{}) (result T) {
		values := make([]reflect.Value, len(in))
		for i, v := range in {
			values[i] = reflect.ValueOf(v)
		}
		for _, fun := range funs {
			values = reflect.ValueOf(fun).Call(values)
			funcType := reflect.TypeOf(fun)
			if funcType.Out(funcType.NumOut()-1).Name() == errorTypeName {
				lastValue := values[len(values)-1]
				if !lastValue.IsNil() {
					return
				}
				values = values[:len(values)-1]
			}
		}
		out := make([]interface{}, len(values))
		for i, v := range values {
			out[i] = v.Interface()
		}
		resultType := reflect.TypeOf(result)
		if len(out) == 1 && resultType.Kind() == reflect.TypeOf(out[0]).Kind() {
			return out[0].(T)
		}
		resultValue := reflect.ValueOf(result)
		for i := 0; i < resultType.NumField(); i++ {
			field := resultType.Field(i)
			for j := 0; j < len(out); j++ {
				if field.Name == reflect.TypeOf(out[j]).Name() {
					resultValue.Field(i).Set(reflect.ValueOf(out[j]))
					out = append(out[:j], out[j+1:]...)
					break
				}
			}
		}
		return result
	}
	return wrapped
}
