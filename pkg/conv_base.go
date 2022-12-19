package flaggy

import (
	"fmt"
	"reflect"
)

type ErrorType string

const (
	ERROR_TYPE_STRING  ErrorType = "string"
	ERROR_TYPE_INTEGER ErrorType = "integer"
	ERROR_TYPE_FLOAT   ErrorType = "float"
)

func (f Flag) setPtr(args Args) error {
	switch f.Type.Type.Elem().Kind() {
	case reflect.Struct:
		{
			inst := reflect.New(f.Type.Type.Elem()).Interface()
			parse(inst, args)
			f.Value.Set(reflect.ValueOf(inst))
			return nil
		}
	case reflect.String:
		{
			if err := f.setStringPtr(args); err != nil {
				return err
			}
		}
	case reflect.Bool:
		{
			var offset int
			var err error
			if offset, err = f.setBoolPtr(args); err != nil {
				return err
			}
			if offset >= 0 {
				return parseValue(f.Name, args[offset:])
			}
			return parseValue(f.Name, []string{})
		}
	case reflect.Int:
		{
			if err := f.setIntPtr(args); err != nil {
				return err
			}
		}
	case reflect.Int32:
		{
			if err := f.setInt32Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Int64:
		{
			if err := f.setInt64Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Int16:
		{
			if err := f.setInt16Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Int8:
		{
			if err := f.setInt8Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Uint:
		{
			if err := f.setUIntPtr(args); err != nil {
				return err
			}
		}
	case reflect.Uint32:
		{
			if err := f.setUInt32Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Uint64:
		{
			if err := f.setUInt64Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Uint16:
		{
			if err := f.setUInt16Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Uint8:
		{
			if err := f.setUInt8Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Float32:
		{
			if err := f.setFloat32Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Float64:
		{
			if err := f.setFloat64Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Complex64:
		{
			if err := f.setComplex64Ptr(args); err != nil {
				return err
			}
		}
	case reflect.Complex128:
		{
			if err := f.setComplex128Ptr(args); err != nil {
				return err
			}
		}
	}
	parseValue(f.Name, args[1:])
	return nil
}

func (f Flag) setValue(args Args) error {
	switch f.Type.Type.Kind() {
	case reflect.Struct:
		{
			inst := reflect.New(f.Type.Type).Interface()
			parse(inst, args)
			f.Value.Set(reflect.ValueOf(inst).Elem())
			return nil
		}
	case reflect.String:
		{
			if err := f.setString(args); err != nil {
				return err
			}
		}
	case reflect.Bool:
		{
			var offset int
			var err error
			if offset, err = f.setBool(args); err != nil {
				return err
			}
			if offset >= 0 {
				return parseValue(f.Name, args[offset:])
			}
			return parseValue(f.Name, []string{})
		}
	case reflect.Int:
		{
			if err := f.setInt(args); err != nil {
				return err
			}
		}
	case reflect.Int32:
		{
			if err := f.setInt32(args); err != nil {
				return err
			}
		}
	case reflect.Int64:
		{
			if err := f.setInt64(args); err != nil {
				return err
			}
		}
	case reflect.Int16:
		{
			if err := f.setInt16(args); err != nil {
				return err
			}
		}
	case reflect.Int8:
		{
			if err := f.setInt8(args); err != nil {
				return err
			}
		}
	case reflect.Uint:
		{
			if err := f.setUInt(args); err != nil {
				return err
			}
		}
	case reflect.Uint32:
		{
			if err := f.setUInt32(args); err != nil {
				return err
			}
		}
	case reflect.Uint64:
		{
			if err := f.setUInt64(args); err != nil {
				return err
			}
		}
	case reflect.Uint16:
		{
			if err := f.setUInt16(args); err != nil {
				return err
			}
		}
	case reflect.Uint8:
		{
			if err := f.setUInt8(args); err != nil {
				return err
			}
		}
	case reflect.Float32:
		{
			if err := f.setFloat32(args); err != nil {
				return err
			}
		}
	case reflect.Float64:
		{
			if err := f.setFloat64(args); err != nil {
				return err
			}
		}
	case reflect.Complex64:
		{
			if err := f.setComplex64(args); err != nil {
				return err
			}
		}
	case reflect.Complex128:
		{
			if err := f.setComplex128(args); err != nil {
				return err
			}
		}
	case reflect.Slice:
		{
			return f.setArray(args)
		}
	}
	parseValue(f.Name, args[1:])
	return nil
}

func (f Flag) setArray(args Args) error {
	switch f.Type.Type.Elem().Kind() {
	case reflect.String:
		{
			if err := f.setStringArray(args); err != nil {
				return err
			}
		}
	case reflect.Int:
		{
			if err := f.setIntArray(args); err != nil {
				return err
			}
		}
	case reflect.Int16:
		{
			if err := f.setInt16Array(args); err != nil {
				return err
			}
		}
	case reflect.Int32:
		{
			if err := f.setInt32Array(args); err != nil {
				return err
			}
		}
	case reflect.Int64:
		{
			if err := f.setInt64Array(args); err != nil {
				return err
			}
		}
	case reflect.Int8:
		{
			if err := f.setInt8Array(args); err != nil {
				return err
			}
		}
	case reflect.Uint:
		{
			if err := f.setUIntArray(args); err != nil {
				return err
			}
		}
	case reflect.Uint16:
		{
			if err := f.setUInt16Array(args); err != nil {
				return err
			}
		}
	case reflect.Uint32:
		{
			if err := f.setUInt32Array(args); err != nil {
				return err
			}
		}
	case reflect.Uint64:
		{
			if err := f.setUInt64Array(args); err != nil {
				return err
			}
		}
	case reflect.Uint8:
		{
			if err := f.setUInt8Array(args); err != nil {
				return err
			}
		}
	case reflect.Float32:
		{
			if err := f.setFloat32Array(args); err != nil {
				return err
			}
		}
	case reflect.Float64:
		{
			if err := f.setFloat64Array(args); err != nil {
				return err
			}
		}
	case reflect.Complex64:
		{
			if err := f.setComplex64Array(args); err != nil {
				return err
			}
		}
	case reflect.Complex128:
		{
			if err := f.setComplex128Array(args); err != nil {
				return err
			}
		}
	}
	parseValue(f.Name, args[1:])
	return nil
}

func validate(errorType ErrorType, args Args) error {
	if len(args) == 0 {
		return fmt.Errorf(ERROR_INVALID_TYPE, errorType, "nothing")
	}
	isFlag := args.isFlag()
	if isFlag {
		return fmt.Errorf(ERROR_INVALID_TYPE, errorType, args[0])
	}
	return nil
}

func ptr[T any](value T) *T {
	return &value
}
