package flaggy

import (
	"fmt"
	"reflect"
	"strconv"
)

func (f Flag) setStringArray(args Args) error {
	if err := validate(ERROR_TYPE_STRING, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]string)
	if array == nil {
		array = make([]string, 0)
	}
	array = append(array, args[0])
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setIntArray(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]int)
	if array == nil {
		array = make([]int, 0)
	}
	i, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, int(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setInt32Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]int32)
	if array == nil {
		array = make([]int32, 0)
	}
	i, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, int32(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setInt64Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]int64)
	if array == nil {
		array = make([]int64, 0)
	}
	i, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, int64(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setInt16Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]int16)
	if array == nil {
		array = make([]int16, 0)
	}
	i, err := strconv.ParseInt(args[0], 10, 16)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, int16(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setInt8Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]int8)
	if array == nil {
		array = make([]int8, 0)
	}
	i, err := strconv.ParseInt(args[0], 10, 8)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, int8(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setUIntArray(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]uint)
	if array == nil {
		array = make([]uint, 0)
	}
	i, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, uint(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setUInt32Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]uint32)
	if array == nil {
		array = make([]uint32, 0)
	}
	i, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, uint32(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setUInt64Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]uint64)
	if array == nil {
		array = make([]uint64, 0)
	}
	i, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, uint64(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setUInt16Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]uint16)
	if array == nil {
		array = make([]uint16, 0)
	}
	i, err := strconv.ParseUint(args[0], 10, 16)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, uint16(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setUInt8Array(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]uint8)
	if array == nil {
		array = make([]uint8, 0)
	}
	i, err := strconv.ParseUint(args[0], 10, 8)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	array = append(array, uint8(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setFloat32Array(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]float32)
	if array == nil {
		array = make([]float32, 0)
	}
	i, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	array = append(array, float32(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setFloat64Array(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]float64)
	if array == nil {
		array = make([]float64, 0)
	}
	i, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	array = append(array, float64(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setComplex64Array(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]complex64)
	if array == nil {
		array = make([]complex64, 0)
	}
	i, err := strconv.ParseComplex(args[0], 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	array = append(array, complex64(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}

func (f Flag) setComplex128Array(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	array := f.Value.Interface().([]complex128)
	if array == nil {
		array = make([]complex128, 0)
	}
	i, err := strconv.ParseComplex(args[0], 128)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	array = append(array, complex128(i))
	f.Value.Set(reflect.ValueOf(array))
	return nil
}
