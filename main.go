package pipe

import (
	"reflect"
)

const errorTypeName = "error"

func Link(funs ...interface{}) func(in ...interface{}) []interface{} {
	wrapped := func(in ...interface{}) []interface{} {
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
					return []interface{}{lastValue.Interface()}
				}
				values = values[:len(values)-1]
			}
		}
		out := make([]interface{}, len(values))
		for i, v := range values {
			out[i] = v.Interface()
		}
		return out
	}
	return wrapped
}
