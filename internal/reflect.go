package flaggy

import (
	"reflect"
	"strings"
)

type Args []string
type Tag string
type Never bool
type Flag struct {
	Short    *string
	Long     string
	Help     string
	Type     reflect.StructField
	Value    reflect.Value
	Instance any
	Index    int
	Name     string
}

func (args Args) isFlag() bool {
	if len(args) == 0 {
		return false
	}
	return strings.HasPrefix(args[0], "-")
}

func (f Flag) set(cmd string, args Args) error {
	if f.Type.Type.Kind() == reflect.Pointer {
		return f.setPtr(args)
	}
	return f.setValue(args)
}
func (f Flag) parse(args []string) error {
	cmd := args[0]
	if f.Long == cmd || (f.Short != nil && *f.Short == cmd) {
		return f.set(cmd, args[1:])
	}
	return nil
}

type Flags map[string][]Flag

const (
	TAG_SHORT          Tag = "short"
	TAG_LONG           Tag = "long"
	TAG_HELP           Tag = "help"
	ERROR_INVALID_TYPE     = "expected %s recieved %s"
)

var (
	flags      Flags
	runMethods map[string]func() error
)

func init() {
	flags = make(Flags)

}

func reflector(inst any) string {
	reflectedType := reflect.TypeOf(inst)
	reflectedValue := reflect.ValueOf(inst)
	name := reflectedType.Elem().Name()
	flags[name] = make([]Flag, 0)
	runMethods = make(map[string]func() error)
	nOfF := reflectedType.Elem().NumField()
	for fi := 0; fi < nOfF; fi++ {
		field := reflectedType.Elem().Field(fi)
		value := reflectedValue.Elem().Field(fi)
		flag := toFlag(inst, name, fi, value, field)
		flags[name] = append(flags[name], *flag)
	}
	runMethod := reflectedValue.MethodByName("Run")
	runMethods[name] = func() error {
		out := runMethod.Call(nil)
		if len(out) > 0 {
			switch t := out[0].Interface().(type) {
			case nil:
				{
					return nil
				}
			case error:
				{
					return t
				}
			}
		}
		return nil
	}
	return name
}

func toFlag(instance any, name string, index int, value reflect.Value, field reflect.StructField) *Flag {
	flag := Flag{}
	short := field.Tag.Get(string(TAG_SHORT))
	if len(short) > 0 {
		flag.Short = &short
	}
	flag.Long = field.Tag.Get(string(TAG_LONG))
	flag.Help = field.Tag.Get(string(TAG_HELP))
	flag.Type = field
	flag.Index = index
	flag.Value = value
	flag.Name = name
	return &flag
}
