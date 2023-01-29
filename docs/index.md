Win32 API bindings for Golang. 

The bindings are all implemented through syscall, no cgo used.
Most of the code are generated from win32metadata, which is officially published by microsoft to describe Win32 APIs.

The APIs in this lib are not exhaustive, those that are less frequently used are not included.
The code generator is avaiable [here](https://github.com/zzl/go-win32api-gen), you can use it to generate your own code.


## Project Repo
[github.com/zzl/go-win32api](https://github.com/zzl/go-win32api)

## Status
Currently only 64-bit architecture is targeted and tested.

## Installation
```
go get github.com/zzl/go-win32api/v2
```
```go
import "github.com/zzl/go-win32api/v2/win32"
```

## About unions
Union fields are mapped to accessor methods in Go. These methods has the same name as the fields, and return pointer values, so that they can be assigned with new values. 
An additional value accessor method is also generated for each field, whose name is in the pattern of FieldName+Val.

For example, with this C Union:
```
union QWord {
    UINT64 WholeValue;
    struct {
        UINT32 LowPart;
        UINT32 HighPart;
    } Parts;
};
```
The code in Go the get the field values would look like this:
```
var qw QWord
var whole uint64 = qw.WholeValueVal() //or whole = *qw.WholeValue()
var lowPart uint32 = qw.Parts().LowPart
```
To set the field values, we use:
```
*qw.WholeValue() = 123
qw.Parts().LowPart = 4
```
