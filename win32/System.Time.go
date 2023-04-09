package win32

import (
	"syscall"
	"unsafe"
)

const (
	WszW32TimeRegKeyTimeProviders       string = "System\\CurrentControlSet\\Services\\W32Time\\TimeProviders"
	WszW32TimeRegKeyPolicyTimeProviders string = "Software\\Policies\\Microsoft\\W32Time\\TimeProviders"
	WszW32TimeRegValueEnabled           string = "Enabled"
	WszW32TimeRegValueDllName           string = "DllName"
	WszW32TimeRegValueInputProvider     string = "InputProvider"
	WszW32TimeRegValueMetaDataProvider  string = "MetaDataProvider"
	TSF_Hardware                        uint32 = 0x1
	TSF_Authenticated                   uint32 = 0x2
	TSF_IPv6                            uint32 = 0x4
	TSF_SignatureAuthenticated          uint32 = 0x8
	TIME_ZONE_ID_INVALID                uint32 = 0xffffffff
)

// structs

type TIME_ZONE_INFORMATION struct {
	Bias         int32
	StandardName [32]uint16
	StandardDate SYSTEMTIME
	StandardBias int32
	DaylightName [32]uint16
	DaylightDate SYSTEMTIME
	DaylightBias int32
}

type DYNAMIC_TIME_ZONE_INFORMATION struct {
	Bias                        int32
	StandardName                [32]uint16
	StandardDate                SYSTEMTIME
	StandardBias                int32
	DaylightName                [32]uint16
	DaylightDate                SYSTEMTIME
	DaylightBias                int32
	TimeZoneKeyName             [128]uint16
	DynamicDaylightTimeDisabled BOOLEAN
}

var (
	pSystemTimeToTzSpecificLocalTime             uintptr
	pTzSpecificLocalTimeToSystemTime             uintptr
	pFileTimeToSystemTime                        uintptr
	pSystemTimeToFileTime                        uintptr
	pGetTimeZoneInformation                      uintptr
	pSetTimeZoneInformation                      uintptr
	pSetDynamicTimeZoneInformation               uintptr
	pGetDynamicTimeZoneInformation               uintptr
	pGetTimeZoneInformationForYear               uintptr
	pEnumDynamicTimeZoneInformation              uintptr
	pGetDynamicTimeZoneInformationEffectiveYears uintptr
	pSystemTimeToTzSpecificLocalTimeEx           uintptr
	pTzSpecificLocalTimeToSystemTimeEx           uintptr
	pLocalFileTimeToLocalSystemTime              uintptr
	pLocalSystemTimeToLocalFileTime              uintptr
)

func SystemTimeToTzSpecificLocalTime(lpTimeZoneInformation *TIME_ZONE_INFORMATION, lpUniversalTime *SYSTEMTIME, lpLocalTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSystemTimeToTzSpecificLocalTime, libKernel32, "SystemTimeToTzSpecificLocalTime")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)), uintptr(unsafe.Pointer(lpUniversalTime)), uintptr(unsafe.Pointer(lpLocalTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func TzSpecificLocalTimeToSystemTime(lpTimeZoneInformation *TIME_ZONE_INFORMATION, lpLocalTime *SYSTEMTIME, lpUniversalTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pTzSpecificLocalTimeToSystemTime, libKernel32, "TzSpecificLocalTimeToSystemTime")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)), uintptr(unsafe.Pointer(lpLocalTime)), uintptr(unsafe.Pointer(lpUniversalTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func FileTimeToSystemTime(lpFileTime *FILETIME, lpSystemTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pFileTimeToSystemTime, libKernel32, "FileTimeToSystemTime")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpFileTime)), uintptr(unsafe.Pointer(lpSystemTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SystemTimeToFileTime(lpSystemTime *SYSTEMTIME, lpFileTime *FILETIME) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSystemTimeToFileTime, libKernel32, "SystemTimeToFileTime")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTime)), uintptr(unsafe.Pointer(lpFileTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetTimeZoneInformation(lpTimeZoneInformation *TIME_ZONE_INFORMATION) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pGetTimeZoneInformation, libKernel32, "GetTimeZoneInformation")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)))
	return uint32(ret), WIN32_ERROR(err)
}

func SetTimeZoneInformation(lpTimeZoneInformation *TIME_ZONE_INFORMATION) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetTimeZoneInformation, libKernel32, "SetTimeZoneInformation")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetDynamicTimeZoneInformation(lpTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetDynamicTimeZoneInformation, libKernel32, "SetDynamicTimeZoneInformation")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetDynamicTimeZoneInformation(pTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pGetDynamicTimeZoneInformation, libKernel32, "GetDynamicTimeZoneInformation")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pTimeZoneInformation)))
	return uint32(ret), WIN32_ERROR(err)
}

func GetTimeZoneInformationForYear(wYear uint16, pdtzi *DYNAMIC_TIME_ZONE_INFORMATION, ptzi *TIME_ZONE_INFORMATION) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetTimeZoneInformationForYear, libKernel32, "GetTimeZoneInformationForYear")
	ret, _, err := syscall.SyscallN(addr, uintptr(wYear), uintptr(unsafe.Pointer(pdtzi)), uintptr(unsafe.Pointer(ptzi)))
	return BOOL(ret), WIN32_ERROR(err)
}

func EnumDynamicTimeZoneInformation(dwIndex uint32, lpTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION) uint32 {
	addr := LazyAddr(&pEnumDynamicTimeZoneInformation, libAdvapi32, "EnumDynamicTimeZoneInformation")
	ret, _, _ := syscall.SyscallN(addr, uintptr(dwIndex), uintptr(unsafe.Pointer(lpTimeZoneInformation)))
	return uint32(ret)
}

func GetDynamicTimeZoneInformationEffectiveYears(lpTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION, FirstYear *uint32, LastYear *uint32) uint32 {
	addr := LazyAddr(&pGetDynamicTimeZoneInformationEffectiveYears, libAdvapi32, "GetDynamicTimeZoneInformationEffectiveYears")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)), uintptr(unsafe.Pointer(FirstYear)), uintptr(unsafe.Pointer(LastYear)))
	return uint32(ret)
}

func SystemTimeToTzSpecificLocalTimeEx(lpTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION, lpUniversalTime *SYSTEMTIME, lpLocalTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSystemTimeToTzSpecificLocalTimeEx, libKernel32, "SystemTimeToTzSpecificLocalTimeEx")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)), uintptr(unsafe.Pointer(lpUniversalTime)), uintptr(unsafe.Pointer(lpLocalTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func TzSpecificLocalTimeToSystemTimeEx(lpTimeZoneInformation *DYNAMIC_TIME_ZONE_INFORMATION, lpLocalTime *SYSTEMTIME, lpUniversalTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pTzSpecificLocalTimeToSystemTimeEx, libKernel32, "TzSpecificLocalTimeToSystemTimeEx")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeZoneInformation)), uintptr(unsafe.Pointer(lpLocalTime)), uintptr(unsafe.Pointer(lpUniversalTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func LocalFileTimeToLocalSystemTime(timeZoneInformation *TIME_ZONE_INFORMATION, localFileTime *FILETIME, localSystemTime *SYSTEMTIME) BOOL {
	addr := LazyAddr(&pLocalFileTimeToLocalSystemTime, libKernel32, "LocalFileTimeToLocalSystemTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(timeZoneInformation)), uintptr(unsafe.Pointer(localFileTime)), uintptr(unsafe.Pointer(localSystemTime)))
	return BOOL(ret)
}

func LocalSystemTimeToLocalFileTime(timeZoneInformation *TIME_ZONE_INFORMATION, localSystemTime *SYSTEMTIME, localFileTime *FILETIME) BOOL {
	addr := LazyAddr(&pLocalSystemTimeToLocalFileTime, libKernel32, "LocalSystemTimeToLocalFileTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(timeZoneInformation)), uintptr(unsafe.Pointer(localSystemTime)), uintptr(unsafe.Pointer(localFileTime)))
	return BOOL(ret)
}
