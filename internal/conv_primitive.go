package flaggy

import (
	"fmt"
	"reflect"
	"strconv"
)

func (f Flag) setString(args Args) error {
	if err := validate(ERROR_TYPE_STRING, args); err != nil {
		return err
	}
	f.Value.Set(reflect.ValueOf(args[0]))
	return nil
}

func (f Flag) setBool(args Args) (int, error) {
	var value reflect.Value
	if len(args) == 0 {
		b := true
		value = reflect.ValueOf(b)
		f.Value.Set(value)
		return -1, nil
	}
	isFlag := args.isFlag()
	if isFlag {
		b := true
		value = reflect.ValueOf(b)
		f.Value.Set(value)
		return 0, nil
	}
	b, err := strconv.ParseBool(args[0])
	if err != nil {
		return -1, fmt.Errorf(ERROR_INVALID_TYPE, "boolean", args[0])
	}
	value = reflect.ValueOf(b)
	f.Value.Set(value)
	return 1, nil
}

func (f Flag) setInt(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(int(i)))
	return nil
}

func (f Flag) setInt32(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(int32(i)))
	return nil
}

func (f Flag) setInt64(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(int64(i)))
	return nil
}

func (f Flag) setInt16(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseInt(args[0], 10, 16)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(int16(i)))
	return nil
}

func (f Flag) setInt8(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseInt(args[0], 10, 8)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(int8(i)))
	return nil
}

func (f Flag) setUInt(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(uint(i)))
	return nil
}

func (f Flag) setUInt32(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(uint32(i)))
	return nil
}

func (f Flag) setUInt64(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(uint64(i)))
	return nil
}

func (f Flag) setUInt16(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseUint(args[0], 10, 16)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(uint16(i)))
	return nil
}

func (f Flag) setUInt8(args Args) error {
	if err := validate(ERROR_TYPE_INTEGER, args); err != nil {
		return err
	}
	i, err := strconv.ParseInt(args[0], 10, 8)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_INTEGER, args[0])
	}
	f.Value.Set(reflect.ValueOf(uint8(i)))
	return nil
}

func (f Flag) setFloat32(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	i, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	f.Value.Set(reflect.ValueOf(float32(i)))
	return nil
}

func (f Flag) setFloat64(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	i, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	f.Value.Set(reflect.ValueOf(float64(i)))
	return nil
}

func (f Flag) setComplex64(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	i, err := strconv.ParseComplex(args[0], 64)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	f.Value.Set(reflect.ValueOf(complex64(i)))
	return nil
}

func (f Flag) setComplex128(args Args) error {
	if err := validate(ERROR_TYPE_FLOAT, args); err != nil {
		return err
	}
	i, err := strconv.ParseComplex(args[0], 128)
	if err != nil {
		return fmt.Errorf(ERROR_INVALID_TYPE, ERROR_TYPE_FLOAT, args[0])
	}
	f.Value.Set(reflect.ValueOf(complex128(i)))
	return nil
}
