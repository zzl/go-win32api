package win32

import (
	"syscall"
	"unsafe"
)

// enums

// enum
type PRINT_WINDOW_FLAGS uint32

const (
	PW_CLIENTONLY PRINT_WINDOW_FLAGS = 1
)

// structs

type DRAWPATRECT struct {
	PtPosition POINT
	PtSize     POINT
	WStyle     uint16
	WPattern   uint16
}

type PSFEATURE_OUTPUT struct {
	BPageIndependent BOOL
	BSetPageDevice   BOOL
}

type PSFEATURE_CUSTPAPER struct {
	LOrientation  int32
	LWidth        int32
	LHeight       int32
	LWidthOffset  int32
	LHeightOffset int32
}

type DOCINFOA struct {
	CbSize       int32
	LpszDocName  PSTR
	LpszOutput   PSTR
	LpszDatatype PSTR
	FwType       uint32
}

type DOCINFO = DOCINFOW
type DOCINFOW struct {
	CbSize       int32
	LpszDocName  PWSTR
	LpszOutput   PWSTR
	LpszDatatype PWSTR
	FwType       uint32
}

// func types

type ABORTPROC = uintptr
type ABORTPROC_func = func(param0 HDC, param1 int32) BOOL

var (
	pEscape       uintptr
	pExtEscape    uintptr
	pStartDocA    uintptr
	pStartDocW    uintptr
	pEndDoc       uintptr
	pStartPage    uintptr
	pEndPage      uintptr
	pAbortDoc     uintptr
	pSetAbortProc uintptr
	pPrintWindow  uintptr
)

func Escape(hdc HDC, iEscape int32, cjIn int32, pvIn PSTR, pvOut unsafe.Pointer) int32 {
	addr := lazyAddr(&pEscape, libGdi32, "Escape")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(iEscape), uintptr(cjIn), uintptr(unsafe.Pointer(pvIn)), uintptr(pvOut))
	return int32(ret)
}

func ExtEscape(hdc HDC, iEscape int32, cjInput int32, lpInData PSTR, cjOutput int32, lpOutData PSTR) int32 {
	addr := lazyAddr(&pExtEscape, libGdi32, "ExtEscape")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(iEscape), uintptr(cjInput), uintptr(unsafe.Pointer(lpInData)), uintptr(cjOutput), uintptr(unsafe.Pointer(lpOutData)))
	return int32(ret)
}

func StartDocA(hdc HDC, lpdi *DOCINFOA) int32 {
	addr := lazyAddr(&pStartDocA, libGdi32, "StartDocA")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(unsafe.Pointer(lpdi)))
	return int32(ret)
}

var StartDoc = StartDocW

func StartDocW(hdc HDC, lpdi *DOCINFOW) int32 {
	addr := lazyAddr(&pStartDocW, libGdi32, "StartDocW")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(unsafe.Pointer(lpdi)))
	return int32(ret)
}

func EndDoc(hdc HDC) int32 {
	addr := lazyAddr(&pEndDoc, libGdi32, "EndDoc")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func StartPage(hdc HDC) int32 {
	addr := lazyAddr(&pStartPage, libGdi32, "StartPage")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func EndPage(hdc HDC) int32 {
	addr := lazyAddr(&pEndPage, libGdi32, "EndPage")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func AbortDoc(hdc HDC) int32 {
	addr := lazyAddr(&pAbortDoc, libGdi32, "AbortDoc")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func SetAbortProc(hdc HDC, proc ABORTPROC) int32 {
	addr := lazyAddr(&pSetAbortProc, libGdi32, "SetAbortProc")
	ret, _, _ := syscall.SyscallN(addr, hdc, proc)
	return int32(ret)
}

func PrintWindow(hwnd HWND, hdcBlt HDC, nFlags PRINT_WINDOW_FLAGS) BOOL {
	addr := lazyAddr(&pPrintWindow, libUser32, "PrintWindow")
	ret, _, _ := syscall.SyscallN(addr, hwnd, hdcBlt, uintptr(nFlags))
	return BOOL(ret)
}
