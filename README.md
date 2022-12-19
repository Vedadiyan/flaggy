# flaggy
Flaggy is a command line parser that has been inspired by go-flags (https://github.com/jessevdk/go-flags). Flaggy has been designed with the purpose of binding behavior to data when parsing command line arguments in order to make the user interaction logic cleaner and simpler to maintain. Furthermore, Flaggy also allows for creating nested commands and multiple levels of flags when required. 

## Option (Flag) Definition
Defining options (flags) in Flaggy is identical to go-flags. In order to define options, a `struct` has to be created whose fields represent the data required to be collected from the user and parsed from command line arguments. These fields are decorated with `tags` in order to define to which command they belong and what their purpose are. For exampe 


    type Options struct {
        Help bool `long:"help" short:"-h" help:"Shows help"`    
    }

### Parsing 
Parsing options (flags) is as easy as calling the `Parse` method. For example:

    type Options struct {
        Help bool `long:"help" short:"-h" help:"Shows help"`    
    }
    options := Options {}
    err := flaggy.Parse(&options, os.Args[1:])

You can also verify if any argument has been parsed successfully using the `Parsed` method.     

### Supported Data Types 
Flaggy has limited support for Go's data types. The following table lists out all data types that are currently supported by Flaggy: 

- [x] int         
- [x] *int         
- [x] int32       
- [x] *int32       
- [x] int64       
- [x] *int64       
- [x] int16       
- [x] *int16       
- [x] int8        
- [x] *int8        
- [x] uint        
- [x] *uint        
- [x] uint32      
- [x] *uint32      
- [x] uint64      
- [x] *uint64      
- [x] uint16      
- [x] *uint16      
- [x] uint8       
- [x] *uint8       
- [x] float32     
- [x] *float32     
- [x] float64     
- [x] *float64     
- [x] comples64   
- [x] *comples64   
- [x] complex128  
- [x] *complex128  
- [x] string
- [x] *string      
- [x] bool   
- [x] *bool        
- [x] struct 
- [x] *struct      
- [x] slice/array
- [ ] *slice/*array 

*Slices or arrays of structs are not supported*

### Nested Options 
Similar to any other command line interfaces (i.e. kubectl, docker, etc.), Flaggy allows for grouping commands. For instance, the following command `read -f abc.txt --skip-white-space` can be created as: 

    type Options struct {
        Read Read `long:"read" short:"" help:"reads a file"`
    }
    type Read struct {
        FileName string `long:"--filename" short:"-f" help:"the name of the file to read"`
        SkipWhiteSpace bool `long:"--skip-white-space" short:"" help:"skips white spaces"`
    }

### Binding Behavior 
When parsing a flag, Flaggy looks for an optional method called `Run` and invokes that method if it is available on the struct when the parsing process is finished. The `Run` method may also return `error` in case it encounters an error. 

    type Options struct {
        Read Read `long:"read" short:"" help:"reads a file"`
    }
    type Read struct {
        FileName string `long:"--filename" short:"-f" help:"the name of the file to read"`
        SkipWhiteSpace bool `long:"--skip-white-space" short:"" help:"skips white spaces"`
    }
    func (r Read) Run() error {
        failing := true
        if len(r.FileName) == 0 {
            failing = true
        }
        if failing {
            flaggy.PrintHelp()
            return fmt.Errorf("file name is not provided")
        }
        ...
        return nil
    }

If the `Run` method returns an error, the `Parse` method will return it back. 







