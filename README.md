:ukraine:

# go-win32api

Win32 API bindings for Go. 

**Bindings:**
All bindings are implemented through syscall, without using cgo.

**Code Generation:**
Most of the code is generated from win32metadata, an official 
Microsoft publication that describes Win32 APIs.

**API Coverage:**
The library includes frequently used APIs. 
Less common ones are omitted. 
You can generate your own code for these using the code generator available 
[here](https://github.com/zzl/go-winapi-gen).

## Status
Currently, only the 64-bit architecture is supported and tested.

## Installation
```
go get github.com/zzl/go-win32api/v2
```
```go
import "github.com/zzl/go-win32api/v2/win32"
```

## Unions
Union fields are mapped to accessor methods in Go. 
These methods share the same name as the fields and 
return pointers to allow assigning new values. 
Additionally, a value accessor method is generated for each field, 
following the naming pattern FieldName+Val.

## Example:
Consider this C Union:
```c
union QWord {
    UINT64 WholeValue;
    struct {
        UINT32 LowPart;
        UINT32 HighPart;
    } Parts;
};
```

Here's how you would access and set field values in Go:
```go
var qw QWord
var whole uint64 = qw.WholeValueVal() // or whole = *qw.WholeValue()
var lowPart uint32 = qw.Parts().LowPart

// Setting values:
*qw.WholeValue() = 123
qw.Parts().LowPart = 4
```
