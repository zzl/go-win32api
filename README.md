# go-win32api

Win32 API bindings for Golang. 

The bindings are all implemented through syscall, no cgo used.
Most of the code are generated from win32metadata, which is officially published by microsoft to describe Win32 APIs.

The APIs included in this lib are not exhaustive, those that are less frequently used are not included.
The code generator is avaiable [here](https://github.com/zzl/go-win32api-gen), you can use it to generate your own code.

