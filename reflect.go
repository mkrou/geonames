package geonames

import (
	"errors"
	"reflect"
)

func exploreFunc(fn interface{}) (funcVal reflect.Value, argVal reflect.Value, err error) {
	funcVal = reflect.ValueOf(fn)
	if funcVal.Kind() != reflect.Func {
		return reflect.Value{}, reflect.Value{}, errors.New("Arg must be a func")
	}

	funcType := funcVal.Type()
	if funcType.NumIn() != 1 {
		return reflect.Value{}, reflect.Value{}, errors.New("Func must input only one argument")
	}
	if funcType.NumOut() != 1 {
		return reflect.Value{}, reflect.Value{}, errors.New("Func must output only one argument")
	}

	argPtr := funcType.In(0)
	if argPtr.Kind() != reflect.Ptr {
		return reflect.Value{}, reflect.Value{}, errors.New("The input argument must be a pointer")
	}

	argType := argPtr.Elem()
	argVal = reflect.New(argType)
	outType := funcType.Out(0)

	if outType.Kind() == reflect.TypeOf((*error)(nil)).Kind() {
		return reflect.Value{}, reflect.Value{}, errors.New("The output must be an error")
	}

	return funcVal, argVal, nil
}
func getArgument(f interface{}) (interface{}, error) {
	_, arg, err := exploreFunc(f)
	if err != nil {
		return nil, err
	}
	return arg.Interface(), nil
}
func fillArgument(f interface{}, parse func(v interface{}) error) error {
	fn, arg, err := exploreFunc(f)
	if err != nil {
		return err
	}
	if err := parse(arg.Interface()); err != nil {
		return err
	}

	errVal := fn.Call([]reflect.Value{
		arg,
	})[0]

	if err, ok := errVal.Interface().(error); ok {
		return err
	}
	return nil
}
