package win32

import (
	"syscall"
	"unsafe"
)

type (
	HWINSTA = uintptr
	HDESK   = uintptr
)

// enums

// enum
// flags
type BROADCAST_SYSTEM_MESSAGE_FLAGS uint32

const (
	BSF_ALLOWSFW           BROADCAST_SYSTEM_MESSAGE_FLAGS = 128
	BSF_FLUSHDISK          BROADCAST_SYSTEM_MESSAGE_FLAGS = 4
	BSF_FORCEIFHUNG        BROADCAST_SYSTEM_MESSAGE_FLAGS = 32
	BSF_IGNORECURRENTTASK  BROADCAST_SYSTEM_MESSAGE_FLAGS = 2
	BSF_NOHANG             BROADCAST_SYSTEM_MESSAGE_FLAGS = 8
	BSF_NOTIMEOUTIFNOTHUNG BROADCAST_SYSTEM_MESSAGE_FLAGS = 64
	BSF_POSTMESSAGE        BROADCAST_SYSTEM_MESSAGE_FLAGS = 16
	BSF_QUERY              BROADCAST_SYSTEM_MESSAGE_FLAGS = 1
	BSF_SENDNOTIFYMESSAGE  BROADCAST_SYSTEM_MESSAGE_FLAGS = 256
	BSF_LUID               BROADCAST_SYSTEM_MESSAGE_FLAGS = 1024
	BSF_RETURNHDESK        BROADCAST_SYSTEM_MESSAGE_FLAGS = 512
)

// enum
// flags
type BROADCAST_SYSTEM_MESSAGE_INFO uint32

const (
	BSM_ALLCOMPONENTS BROADCAST_SYSTEM_MESSAGE_INFO = 0
	BSM_ALLDESKTOPS   BROADCAST_SYSTEM_MESSAGE_INFO = 16
	BSM_APPLICATIONS  BROADCAST_SYSTEM_MESSAGE_INFO = 8
)

// enum
type USER_OBJECT_INFORMATION_INDEX int32

const (
	UOI_FLAGS    USER_OBJECT_INFORMATION_INDEX = 1
	UOI_HEAPSIZE USER_OBJECT_INFORMATION_INDEX = 5
	UOI_IO       USER_OBJECT_INFORMATION_INDEX = 6
	UOI_NAME     USER_OBJECT_INFORMATION_INDEX = 2
	UOI_TYPE     USER_OBJECT_INFORMATION_INDEX = 3
	UOI_USER_SID USER_OBJECT_INFORMATION_INDEX = 4
)

// enum
type DESKTOP_CONTROL_FLAGS uint32

const (
	DF_ALLOWOTHERACCOUNTHOOK DESKTOP_CONTROL_FLAGS = 1
)

// enum
type DESKTOP_ACCESS_FLAGS uint32

const (
	DESKTOP_DELETE          DESKTOP_ACCESS_FLAGS = 65536
	DESKTOP_READ_CONTROL    DESKTOP_ACCESS_FLAGS = 131072
	DESKTOP_WRITE_DAC       DESKTOP_ACCESS_FLAGS = 262144
	DESKTOP_WRITE_OWNER     DESKTOP_ACCESS_FLAGS = 524288
	DESKTOP_SYNCHRONIZE     DESKTOP_ACCESS_FLAGS = 1048576
	DESKTOP_READOBJECTS     DESKTOP_ACCESS_FLAGS = 1
	DESKTOP_CREATEWINDOW    DESKTOP_ACCESS_FLAGS = 2
	DESKTOP_CREATEMENU      DESKTOP_ACCESS_FLAGS = 4
	DESKTOP_HOOKCONTROL     DESKTOP_ACCESS_FLAGS = 8
	DESKTOP_JOURNALRECORD   DESKTOP_ACCESS_FLAGS = 16
	DESKTOP_JOURNALPLAYBACK DESKTOP_ACCESS_FLAGS = 32
	DESKTOP_ENUMERATE       DESKTOP_ACCESS_FLAGS = 64
	DESKTOP_WRITEOBJECTS    DESKTOP_ACCESS_FLAGS = 128
	DESKTOP_SWITCHDESKTOP   DESKTOP_ACCESS_FLAGS = 256
)

// structs

type USEROBJECTFLAGS struct {
	FInherit  BOOL
	FReserved BOOL
	DwFlags   uint32
}

type BSMINFO struct {
	CbSize uint32
	Hdesk  HDESK
	Hwnd   HWND
	Luid   LUID
}

// func types

type WINSTAENUMPROCA = uintptr
type WINSTAENUMPROCA_func = func(param0 PSTR, param1 LPARAM) BOOL

type WINSTAENUMPROCW = uintptr
type WINSTAENUMPROCW_func = func(param0 PWSTR, param1 LPARAM) BOOL

type DESKTOPENUMPROCA = uintptr
type DESKTOPENUMPROCA_func = func(param0 PSTR, param1 LPARAM) BOOL

type DESKTOPENUMPROCW = uintptr
type DESKTOPENUMPROCW_func = func(param0 PWSTR, param1 LPARAM) BOOL

var (
	pCreateDesktopA            uintptr
	pCreateDesktopW            uintptr
	pCreateDesktopExA          uintptr
	pCreateDesktopExW          uintptr
	pOpenDesktopA              uintptr
	pOpenDesktopW              uintptr
	pOpenInputDesktop          uintptr
	pEnumDesktopsA             uintptr
	pEnumDesktopsW             uintptr
	pEnumDesktopWindows        uintptr
	pSwitchDesktop             uintptr
	pSetThreadDesktop          uintptr
	pCloseDesktop              uintptr
	pGetThreadDesktop          uintptr
	pCreateWindowStationA      uintptr
	pCreateWindowStationW      uintptr
	pOpenWindowStationA        uintptr
	pOpenWindowStationW        uintptr
	pEnumWindowStationsA       uintptr
	pEnumWindowStationsW       uintptr
	pCloseWindowStation        uintptr
	pSetProcessWindowStation   uintptr
	pGetProcessWindowStation   uintptr
	pGetUserObjectInformationA uintptr
	pGetUserObjectInformationW uintptr
	pSetUserObjectInformationA uintptr
	pSetUserObjectInformationW uintptr
	pBroadcastSystemMessageExA uintptr
	pBroadcastSystemMessageExW uintptr
	pBroadcastSystemMessageA   uintptr
	pBroadcastSystemMessageW   uintptr
)

func CreateDesktopA(lpszDesktop PSTR, lpszDevice PSTR, pDevmode *DEVMODEA, dwFlags DESKTOP_CONTROL_FLAGS, dwDesiredAccess uint32, lpsa *SECURITY_ATTRIBUTES) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pCreateDesktopA, libUser32, "CreateDesktopA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszDesktop)), uintptr(unsafe.Pointer(lpszDevice)), uintptr(unsafe.Pointer(pDevmode)), uintptr(dwFlags), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpsa)))
	return ret, WIN32_ERROR(err)
}

var CreateDesktop = CreateDesktopW

func CreateDesktopW(lpszDesktop PWSTR, lpszDevice PWSTR, pDevmode *DEVMODEW, dwFlags DESKTOP_CONTROL_FLAGS, dwDesiredAccess uint32, lpsa *SECURITY_ATTRIBUTES) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pCreateDesktopW, libUser32, "CreateDesktopW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszDesktop)), uintptr(unsafe.Pointer(lpszDevice)), uintptr(unsafe.Pointer(pDevmode)), uintptr(dwFlags), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpsa)))
	return ret, WIN32_ERROR(err)
}

func CreateDesktopExA(lpszDesktop PSTR, lpszDevice PSTR, pDevmode *DEVMODEA, dwFlags DESKTOP_CONTROL_FLAGS, dwDesiredAccess uint32, lpsa *SECURITY_ATTRIBUTES, ulHeapSize uint32, pvoid unsafe.Pointer) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pCreateDesktopExA, libUser32, "CreateDesktopExA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszDesktop)), uintptr(unsafe.Pointer(lpszDevice)), uintptr(unsafe.Pointer(pDevmode)), uintptr(dwFlags), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpsa)), uintptr(ulHeapSize), uintptr(pvoid))
	return ret, WIN32_ERROR(err)
}

var CreateDesktopEx = CreateDesktopExW

func CreateDesktopExW(lpszDesktop PWSTR, lpszDevice PWSTR, pDevmode *DEVMODEW, dwFlags DESKTOP_CONTROL_FLAGS, dwDesiredAccess uint32, lpsa *SECURITY_ATTRIBUTES, ulHeapSize uint32, pvoid unsafe.Pointer) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pCreateDesktopExW, libUser32, "CreateDesktopExW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszDesktop)), uintptr(unsafe.Pointer(lpszDevice)), uintptr(unsafe.Pointer(pDevmode)), uintptr(dwFlags), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpsa)), uintptr(ulHeapSize), uintptr(pvoid))
	return ret, WIN32_ERROR(err)
}

func OpenDesktopA(lpszDesktop PSTR, dwFlags DESKTOP_CONTROL_FLAGS, fInherit BOOL, dwDesiredAccess uint32) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pOpenDesktopA, libUser32, "OpenDesktopA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszDesktop)), uintptr(dwFlags), uintptr(fInherit), uintptr(dwDesiredAccess))
	return ret, WIN32_ERROR(err)
}

var OpenDesktop = OpenDesktopW

func OpenDesktopW(lpszDesktop PWSTR, dwFlags DESKTOP_CONTROL_FLAGS, fInherit BOOL, dwDesiredAccess uint32) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pOpenDesktopW, libUser32, "OpenDesktopW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszDesktop)), uintptr(dwFlags), uintptr(fInherit), uintptr(dwDesiredAccess))
	return ret, WIN32_ERROR(err)
}

func OpenInputDesktop(dwFlags DESKTOP_CONTROL_FLAGS, fInherit BOOL, dwDesiredAccess DESKTOP_ACCESS_FLAGS) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pOpenInputDesktop, libUser32, "OpenInputDesktop")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwFlags), uintptr(fInherit), uintptr(dwDesiredAccess))
	return ret, WIN32_ERROR(err)
}

func EnumDesktopsA(hwinsta HWINSTA, lpEnumFunc DESKTOPENUMPROCA, lParam LPARAM) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumDesktopsA, libUser32, "EnumDesktopsA")
	ret, _, err := syscall.SyscallN(addr, hwinsta, lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumDesktops = EnumDesktopsW

func EnumDesktopsW(hwinsta HWINSTA, lpEnumFunc DESKTOPENUMPROCW, lParam LPARAM) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumDesktopsW, libUser32, "EnumDesktopsW")
	ret, _, err := syscall.SyscallN(addr, hwinsta, lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

func EnumDesktopWindows(hDesktop HDESK, lpfn WNDENUMPROC, lParam LPARAM) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumDesktopWindows, libUser32, "EnumDesktopWindows")
	ret, _, err := syscall.SyscallN(addr, hDesktop, lpfn, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

func SwitchDesktop(hDesktop HDESK) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSwitchDesktop, libUser32, "SwitchDesktop")
	ret, _, err := syscall.SyscallN(addr, hDesktop)
	return BOOL(ret), WIN32_ERROR(err)
}

func SetThreadDesktop(hDesktop HDESK) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetThreadDesktop, libUser32, "SetThreadDesktop")
	ret, _, err := syscall.SyscallN(addr, hDesktop)
	return BOOL(ret), WIN32_ERROR(err)
}

func CloseDesktop(hDesktop HDESK) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pCloseDesktop, libUser32, "CloseDesktop")
	ret, _, err := syscall.SyscallN(addr, hDesktop)
	return BOOL(ret), WIN32_ERROR(err)
}

func GetThreadDesktop(dwThreadId uint32) (HDESK, WIN32_ERROR) {
	addr := LazyAddr(&pGetThreadDesktop, libUser32, "GetThreadDesktop")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwThreadId))
	return ret, WIN32_ERROR(err)
}

func CreateWindowStationA(lpwinsta PSTR, dwFlags uint32, dwDesiredAccess uint32, lpsa *SECURITY_ATTRIBUTES) (HWINSTA, WIN32_ERROR) {
	addr := LazyAddr(&pCreateWindowStationA, libUser32, "CreateWindowStationA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpwinsta)), uintptr(dwFlags), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpsa)))
	return ret, WIN32_ERROR(err)
}

var CreateWindowStation = CreateWindowStationW

func CreateWindowStationW(lpwinsta PWSTR, dwFlags uint32, dwDesiredAccess uint32, lpsa *SECURITY_ATTRIBUTES) (HWINSTA, WIN32_ERROR) {
	addr := LazyAddr(&pCreateWindowStationW, libUser32, "CreateWindowStationW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpwinsta)), uintptr(dwFlags), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpsa)))
	return ret, WIN32_ERROR(err)
}

func OpenWindowStationA(lpszWinSta PSTR, fInherit BOOL, dwDesiredAccess uint32) (HWINSTA, WIN32_ERROR) {
	addr := LazyAddr(&pOpenWindowStationA, libUser32, "OpenWindowStationA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszWinSta)), uintptr(fInherit), uintptr(dwDesiredAccess))
	return ret, WIN32_ERROR(err)
}

var OpenWindowStation = OpenWindowStationW

func OpenWindowStationW(lpszWinSta PWSTR, fInherit BOOL, dwDesiredAccess uint32) (HWINSTA, WIN32_ERROR) {
	addr := LazyAddr(&pOpenWindowStationW, libUser32, "OpenWindowStationW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszWinSta)), uintptr(fInherit), uintptr(dwDesiredAccess))
	return ret, WIN32_ERROR(err)
}

func EnumWindowStationsA(lpEnumFunc WINSTAENUMPROCA, lParam LPARAM) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumWindowStationsA, libUser32, "EnumWindowStationsA")
	ret, _, err := syscall.SyscallN(addr, lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumWindowStations = EnumWindowStationsW

func EnumWindowStationsW(lpEnumFunc WINSTAENUMPROCW, lParam LPARAM) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumWindowStationsW, libUser32, "EnumWindowStationsW")
	ret, _, err := syscall.SyscallN(addr, lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

func CloseWindowStation(hWinSta HWINSTA) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pCloseWindowStation, libUser32, "CloseWindowStation")
	ret, _, err := syscall.SyscallN(addr, hWinSta)
	return BOOL(ret), WIN32_ERROR(err)
}

func SetProcessWindowStation(hWinSta HWINSTA) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetProcessWindowStation, libUser32, "SetProcessWindowStation")
	ret, _, err := syscall.SyscallN(addr, hWinSta)
	return BOOL(ret), WIN32_ERROR(err)
}

func GetProcessWindowStation() (HWINSTA, WIN32_ERROR) {
	addr := LazyAddr(&pGetProcessWindowStation, libUser32, "GetProcessWindowStation")
	ret, _, err := syscall.SyscallN(addr)
	return ret, WIN32_ERROR(err)
}

func GetUserObjectInformationA(hObj HANDLE, nIndex USER_OBJECT_INFORMATION_INDEX, pvInfo unsafe.Pointer, nLength uint32, lpnLengthNeeded *uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetUserObjectInformationA, libUser32, "GetUserObjectInformationA")
	ret, _, err := syscall.SyscallN(addr, hObj, uintptr(nIndex), uintptr(pvInfo), uintptr(nLength), uintptr(unsafe.Pointer(lpnLengthNeeded)))
	return BOOL(ret), WIN32_ERROR(err)
}

var GetUserObjectInformation = GetUserObjectInformationW

func GetUserObjectInformationW(hObj HANDLE, nIndex USER_OBJECT_INFORMATION_INDEX, pvInfo unsafe.Pointer, nLength uint32, lpnLengthNeeded *uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetUserObjectInformationW, libUser32, "GetUserObjectInformationW")
	ret, _, err := syscall.SyscallN(addr, hObj, uintptr(nIndex), uintptr(pvInfo), uintptr(nLength), uintptr(unsafe.Pointer(lpnLengthNeeded)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetUserObjectInformationA(hObj HANDLE, nIndex int32, pvInfo unsafe.Pointer, nLength uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetUserObjectInformationA, libUser32, "SetUserObjectInformationA")
	ret, _, err := syscall.SyscallN(addr, hObj, uintptr(nIndex), uintptr(pvInfo), uintptr(nLength))
	return BOOL(ret), WIN32_ERROR(err)
}

var SetUserObjectInformation = SetUserObjectInformationW

func SetUserObjectInformationW(hObj HANDLE, nIndex int32, pvInfo unsafe.Pointer, nLength uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetUserObjectInformationW, libUser32, "SetUserObjectInformationW")
	ret, _, err := syscall.SyscallN(addr, hObj, uintptr(nIndex), uintptr(pvInfo), uintptr(nLength))
	return BOOL(ret), WIN32_ERROR(err)
}

func BroadcastSystemMessageExA(flags BROADCAST_SYSTEM_MESSAGE_FLAGS, lpInfo *BROADCAST_SYSTEM_MESSAGE_INFO, Msg uint32, wParam WPARAM, lParam LPARAM, pbsmInfo *BSMINFO) (int32, WIN32_ERROR) {
	addr := LazyAddr(&pBroadcastSystemMessageExA, libUser32, "BroadcastSystemMessageExA")
	ret, _, err := syscall.SyscallN(addr, uintptr(flags), uintptr(unsafe.Pointer(lpInfo)), uintptr(Msg), wParam, lParam, uintptr(unsafe.Pointer(pbsmInfo)))
	return int32(ret), WIN32_ERROR(err)
}

var BroadcastSystemMessageEx = BroadcastSystemMessageExW

func BroadcastSystemMessageExW(flags BROADCAST_SYSTEM_MESSAGE_FLAGS, lpInfo *BROADCAST_SYSTEM_MESSAGE_INFO, Msg uint32, wParam WPARAM, lParam LPARAM, pbsmInfo *BSMINFO) (int32, WIN32_ERROR) {
	addr := LazyAddr(&pBroadcastSystemMessageExW, libUser32, "BroadcastSystemMessageExW")
	ret, _, err := syscall.SyscallN(addr, uintptr(flags), uintptr(unsafe.Pointer(lpInfo)), uintptr(Msg), wParam, lParam, uintptr(unsafe.Pointer(pbsmInfo)))
	return int32(ret), WIN32_ERROR(err)
}

func BroadcastSystemMessageA(flags uint32, lpInfo *uint32, Msg uint32, wParam WPARAM, lParam LPARAM) int32 {
	addr := LazyAddr(&pBroadcastSystemMessageA, libUser32, "BroadcastSystemMessageA")
	ret, _, _ := syscall.SyscallN(addr, uintptr(flags), uintptr(unsafe.Pointer(lpInfo)), uintptr(Msg), wParam, lParam)
	return int32(ret)
}

var BroadcastSystemMessage = BroadcastSystemMessageW

func BroadcastSystemMessageW(flags BROADCAST_SYSTEM_MESSAGE_FLAGS, lpInfo *BROADCAST_SYSTEM_MESSAGE_INFO, Msg uint32, wParam WPARAM, lParam LPARAM) (int32, WIN32_ERROR) {
	addr := LazyAddr(&pBroadcastSystemMessageW, libUser32, "BroadcastSystemMessageW")
	ret, _, err := syscall.SyscallN(addr, uintptr(flags), uintptr(unsafe.Pointer(lpInfo)), uintptr(Msg), wParam, lParam)
	return int32(ret), WIN32_ERROR(err)
}
