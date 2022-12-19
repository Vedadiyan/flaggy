package flaggy

import (
	"fmt"
)

var (
	parsedFlags int
)

// Parses command line arguments to flag structs
func Parse(inst any, args []string) error {
	parsedFlags = 0
	return parse(inst, args)
}

// Checks if any command line argument has been parsed
func Parsed() bool {
	return parsedFlags > 0
}

// Prints the current level of help
func PrintHelp() {
	flgs := make([]string, 0)
	help := make([]string, 0)
	longestFlg := 0
	maxFlagLevel := 0
	flagName := ""
	for key, value := range flagLevels {
		if value > maxFlagLevel {
			maxFlagLevel = value
			flagName = key
		}
	}

	for _, flag := range flags[flagName] {
		var flg string
		if flag.Short != nil {
			flg = fmt.Sprintf("%s|%s", flag.Long, *flag.Short)
		} else {
			flg = flag.Long
		}
		flgs = append(flgs, flg)
		help = append(help, flag.Help)
		len := len(flg)
		if len > longestFlg {
			longestFlg = len
		}
	}
	for i := 0; i < len(flgs); i++ {
		fmt.Print(flgs[i])
		for space := 0; space < longestFlg-len(flgs[i])+5; space++ {
			fmt.Print(" ")
		}
		fmt.Println(help[i])
	}
}

func parse(inst any, args []string) error {
	name := reflector(inst)
	return parseValue(name, args)
}

func parseValue(name string, args []string) error {
	if len(args) != 0 {
		cmd := args[0]
		for _, flg := range flags[name] {
			if flg.Long == cmd || (flg.Short != nil && *flg.Short == cmd) {
				parsedFlags++
				return flg.parse(args)
			}
		}
	}
	return ignore(name, args)
}

func ignore(name string, args Args) error {
	if len(args) > 2 {
		if Args(args[1:]).isFlag() {
			return parseValue(name, args[1:])
		}
		return parseValue(name, args[2:])
	}
	if runMethod, ok := runMethods[name]; ok {
		return runMethod()
	}
	return nil
}
