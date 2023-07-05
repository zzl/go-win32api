package win32

import (
	"syscall"
	"unsafe"
)

const (
	FIND_RESOURCE_DIRECTORY_TYPES                uint32 = 0x100
	FIND_RESOURCE_DIRECTORY_NAMES                uint32 = 0x200
	FIND_RESOURCE_DIRECTORY_LANGUAGES            uint32 = 0x400
	RESOURCE_ENUM_LN                             uint32 = 0x1
	RESOURCE_ENUM_MUI                            uint32 = 0x2
	RESOURCE_ENUM_MUI_SYSTEM                     uint32 = 0x4
	RESOURCE_ENUM_VALIDATE                       uint32 = 0x8
	RESOURCE_ENUM_MODULE_EXACT                   uint32 = 0x10
	SUPPORT_LANG_NUMBER                          uint32 = 0x20
	GET_MODULE_HANDLE_EX_FLAG_PIN                uint32 = 0x1
	GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT uint32 = 0x2
	GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS       uint32 = 0x4
	CURRENT_IMPORT_REDIRECTION_VERSION           uint32 = 0x1
	LOAD_LIBRARY_OS_INTEGRITY_CONTINUITY         uint32 = 0x8000
)

// enums

// enum
// flags
type LOAD_LIBRARY_FLAGS uint32

const (
	DONT_RESOLVE_DLL_REFERENCES               LOAD_LIBRARY_FLAGS = 1
	LOAD_LIBRARY_AS_DATAFILE                  LOAD_LIBRARY_FLAGS = 2
	LOAD_WITH_ALTERED_SEARCH_PATH             LOAD_LIBRARY_FLAGS = 8
	LOAD_IGNORE_CODE_AUTHZ_LEVEL              LOAD_LIBRARY_FLAGS = 16
	LOAD_LIBRARY_AS_IMAGE_RESOURCE            LOAD_LIBRARY_FLAGS = 32
	LOAD_LIBRARY_AS_DATAFILE_EXCLUSIVE        LOAD_LIBRARY_FLAGS = 64
	LOAD_LIBRARY_REQUIRE_SIGNED_TARGET        LOAD_LIBRARY_FLAGS = 128
	LOAD_LIBRARY_SEARCH_DLL_LOAD_DIR          LOAD_LIBRARY_FLAGS = 256
	LOAD_LIBRARY_SEARCH_APPLICATION_DIR       LOAD_LIBRARY_FLAGS = 512
	LOAD_LIBRARY_SEARCH_USER_DIRS             LOAD_LIBRARY_FLAGS = 1024
	LOAD_LIBRARY_SEARCH_SYSTEM32              LOAD_LIBRARY_FLAGS = 2048
	LOAD_LIBRARY_SEARCH_DEFAULT_DIRS          LOAD_LIBRARY_FLAGS = 4096
	LOAD_LIBRARY_SAFE_CURRENT_DIRS            LOAD_LIBRARY_FLAGS = 8192
	LOAD_LIBRARY_SEARCH_SYSTEM32_NO_FORWARDER LOAD_LIBRARY_FLAGS = 16384
)

// structs

type ENUMUILANG struct {
	NumOfEnumUILang    uint32
	SizeOfEnumUIBuffer uint32
	PEnumUIBuffer      *uint16
}

type REDIRECTION_FUNCTION_DESCRIPTOR struct {
	DllName           PSTR
	FunctionName      PSTR
	RedirectionTarget unsafe.Pointer
}

type REDIRECTION_DESCRIPTOR struct {
	Version       uint32
	FunctionCount uint32
	Redirections  *REDIRECTION_FUNCTION_DESCRIPTOR
}

// func types

type ENUMRESLANGPROCA = uintptr
type ENUMRESLANGPROCA_func = func(hModule HMODULE, lpType PSTR, lpName PSTR, wLanguage uint16, lParam uintptr) BOOL

type ENUMRESLANGPROCW = uintptr
type ENUMRESLANGPROCW_func = func(hModule HMODULE, lpType PWSTR, lpName PWSTR, wLanguage uint16, lParam uintptr) BOOL

type ENUMRESNAMEPROCA = uintptr
type ENUMRESNAMEPROCA_func = func(hModule HMODULE, lpType PSTR, lpName PSTR, lParam uintptr) BOOL

type ENUMRESNAMEPROCW = uintptr
type ENUMRESNAMEPROCW_func = func(hModule HMODULE, lpType PWSTR, lpName PWSTR, lParam uintptr) BOOL

type ENUMRESTYPEPROCA = uintptr
type ENUMRESTYPEPROCA_func = func(hModule HMODULE, lpType PSTR, lParam uintptr) BOOL

type ENUMRESTYPEPROCW = uintptr
type ENUMRESTYPEPROCW_func = func(hModule HMODULE, lpType PWSTR, lParam uintptr) BOOL

type PGET_MODULE_HANDLE_EXA = uintptr
type PGET_MODULE_HANDLE_EXA_func = func(dwFlags uint32, lpModuleName PSTR, phModule *HMODULE) BOOL

type PGET_MODULE_HANDLE_EXW = uintptr
type PGET_MODULE_HANDLE_EXW_func = func(dwFlags uint32, lpModuleName PWSTR, phModule *HMODULE) BOOL

var (
	pDisableThreadLibraryCalls uintptr
	pFindResourceExW           uintptr
	pFreeLibraryAndExitThread  uintptr
	pFreeResource              uintptr
	pGetModuleFileNameA        uintptr
	pGetModuleFileNameW        uintptr
	pGetModuleHandleA          uintptr
	pGetModuleHandleW          uintptr
	pGetModuleHandleExA        uintptr
	pGetModuleHandleExW        uintptr
	pGetProcAddress            uintptr
	pLoadLibraryExA            uintptr
	pLoadLibraryExW            uintptr
	pLoadResource              uintptr
	pLockResource              uintptr
	pSizeofResource            uintptr
	pAddDllDirectory           uintptr
	pRemoveDllDirectory        uintptr
	pSetDefaultDllDirectories  uintptr
	pEnumResourceLanguagesExA  uintptr
	pEnumResourceLanguagesExW  uintptr
	pEnumResourceNamesExA      uintptr
	pEnumResourceNamesExW      uintptr
	pEnumResourceTypesExA      uintptr
	pEnumResourceTypesExW      uintptr
	pFindResourceW             uintptr
	pLoadLibraryA              uintptr
	pLoadLibraryW              uintptr
	pEnumResourceNamesW        uintptr
	pEnumResourceNamesA        uintptr
	pLoadModule                uintptr
	pLoadPackagedLibrary       uintptr
	pFindResourceA             uintptr
	pFindResourceExA           uintptr
	pEnumResourceTypesA        uintptr
	pEnumResourceTypesW        uintptr
	pEnumResourceLanguagesA    uintptr
	pEnumResourceLanguagesW    uintptr
	pBeginUpdateResourceA      uintptr
	pBeginUpdateResourceW      uintptr
	pUpdateResourceA           uintptr
	pUpdateResourceW           uintptr
	pEndUpdateResourceA        uintptr
	pEndUpdateResourceW        uintptr
	pSetDllDirectoryA          uintptr
	pSetDllDirectoryW          uintptr
	pGetDllDirectoryA          uintptr
	pGetDllDirectoryW          uintptr
)

func DisableThreadLibraryCalls(hLibModule HMODULE) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pDisableThreadLibraryCalls, libKernel32, "DisableThreadLibraryCalls")
	ret, _, err := syscall.SyscallN(addr, hLibModule)
	return BOOL(ret), WIN32_ERROR(err)
}

var FindResourceEx = FindResourceExW

func FindResourceExW(hModule HMODULE, lpType PWSTR, lpName PWSTR, wLanguage uint16) HRSRC {
	addr := LazyAddr(&pFindResourceExW, libKernel32, "FindResourceExW")
	ret, _, _ := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), uintptr(wLanguage))
	return ret
}

func FreeLibraryAndExitThread(hLibModule HMODULE, dwExitCode uint32) {
	addr := LazyAddr(&pFreeLibraryAndExitThread, libKernel32, "FreeLibraryAndExitThread")
	syscall.SyscallN(addr, hLibModule, uintptr(dwExitCode))
}

func FreeResource(hResData HGLOBAL) BOOL {
	addr := LazyAddr(&pFreeResource, libKernel32, "FreeResource")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(hResData)))
	return BOOL(ret)
}

func GetModuleFileNameA(hModule HMODULE, lpFilename PSTR, nSize uint32) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pGetModuleFileNameA, libKernel32, "GetModuleFileNameA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpFilename)), uintptr(nSize))
	return uint32(ret), WIN32_ERROR(err)
}

var GetModuleFileName = GetModuleFileNameW

func GetModuleFileNameW(hModule HMODULE, lpFilename PWSTR, nSize uint32) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pGetModuleFileNameW, libKernel32, "GetModuleFileNameW")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpFilename)), uintptr(nSize))
	return uint32(ret), WIN32_ERROR(err)
}

func GetModuleHandleA(lpModuleName PSTR) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pGetModuleHandleA, libKernel32, "GetModuleHandleA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpModuleName)))
	return ret, WIN32_ERROR(err)
}

var GetModuleHandle = GetModuleHandleW

func GetModuleHandleW(lpModuleName PWSTR) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pGetModuleHandleW, libKernel32, "GetModuleHandleW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpModuleName)))
	return ret, WIN32_ERROR(err)
}

func GetModuleHandleExA(dwFlags uint32, lpModuleName PSTR, phModule *HMODULE) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetModuleHandleExA, libKernel32, "GetModuleHandleExA")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwFlags), uintptr(unsafe.Pointer(lpModuleName)), uintptr(unsafe.Pointer(phModule)))
	return BOOL(ret), WIN32_ERROR(err)
}

var GetModuleHandleEx = GetModuleHandleExW

func GetModuleHandleExW(dwFlags uint32, lpModuleName PWSTR, phModule *HMODULE) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetModuleHandleExW, libKernel32, "GetModuleHandleExW")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwFlags), uintptr(unsafe.Pointer(lpModuleName)), uintptr(unsafe.Pointer(phModule)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetProcAddress(hModule HMODULE, lpProcName PSTR) (FARPROC, WIN32_ERROR) {
	addr := LazyAddr(&pGetProcAddress, libKernel32, "GetProcAddress")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpProcName)))
	return FARPROC(ret), WIN32_ERROR(err)
}

func LoadLibraryExA(lpLibFileName PSTR, hFile HANDLE, dwFlags LOAD_LIBRARY_FLAGS) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pLoadLibraryExA, libKernel32, "LoadLibraryExA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpLibFileName)), hFile, uintptr(dwFlags))
	return ret, WIN32_ERROR(err)
}

var LoadLibraryEx = LoadLibraryExW

func LoadLibraryExW(lpLibFileName PWSTR, hFile HANDLE, dwFlags LOAD_LIBRARY_FLAGS) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pLoadLibraryExW, libKernel32, "LoadLibraryExW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpLibFileName)), hFile, uintptr(dwFlags))
	return ret, WIN32_ERROR(err)
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) (HGLOBAL, WIN32_ERROR) {
	addr := LazyAddr(&pLoadResource, libKernel32, "LoadResource")
	ret, _, err := syscall.SyscallN(addr, hModule, hResInfo)
	return (HGLOBAL)(unsafe.Pointer(ret)), WIN32_ERROR(err)
}

func LockResource(hResData HGLOBAL) unsafe.Pointer {
	addr := LazyAddr(&pLockResource, libKernel32, "LockResource")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(hResData)))
	return (unsafe.Pointer)(ret)
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pSizeofResource, libKernel32, "SizeofResource")
	ret, _, err := syscall.SyscallN(addr, hModule, hResInfo)
	return uint32(ret), WIN32_ERROR(err)
}

func AddDllDirectory(NewDirectory PWSTR) (unsafe.Pointer, WIN32_ERROR) {
	addr := LazyAddr(&pAddDllDirectory, libKernel32, "AddDllDirectory")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(NewDirectory)))
	return (unsafe.Pointer)(ret), WIN32_ERROR(err)
}

func RemoveDllDirectory(Cookie unsafe.Pointer) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pRemoveDllDirectory, libKernel32, "RemoveDllDirectory")
	ret, _, err := syscall.SyscallN(addr, uintptr(Cookie))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetDefaultDllDirectories(DirectoryFlags LOAD_LIBRARY_FLAGS) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetDefaultDllDirectories, libKernel32, "SetDefaultDllDirectories")
	ret, _, err := syscall.SyscallN(addr, uintptr(DirectoryFlags))
	return BOOL(ret), WIN32_ERROR(err)
}

func EnumResourceLanguagesExA(hModule HMODULE, lpType PSTR, lpName PSTR, lpEnumFunc ENUMRESLANGPROCA, lParam uintptr, dwFlags uint32, LangId uint16) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceLanguagesExA, libKernel32, "EnumResourceLanguagesExA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), lpEnumFunc, lParam, uintptr(dwFlags), uintptr(LangId))
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumResourceLanguagesEx = EnumResourceLanguagesExW

func EnumResourceLanguagesExW(hModule HMODULE, lpType PWSTR, lpName PWSTR, lpEnumFunc ENUMRESLANGPROCW, lParam uintptr, dwFlags uint32, LangId uint16) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceLanguagesExW, libKernel32, "EnumResourceLanguagesExW")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), lpEnumFunc, lParam, uintptr(dwFlags), uintptr(LangId))
	return BOOL(ret), WIN32_ERROR(err)
}

func EnumResourceNamesExA(hModule HMODULE, lpType PSTR, lpEnumFunc ENUMRESNAMEPROCA, lParam uintptr, dwFlags uint32, LangId uint16) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceNamesExA, libKernel32, "EnumResourceNamesExA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), lpEnumFunc, lParam, uintptr(dwFlags), uintptr(LangId))
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumResourceNamesEx = EnumResourceNamesExW

func EnumResourceNamesExW(hModule HMODULE, lpType PWSTR, lpEnumFunc ENUMRESNAMEPROCW, lParam uintptr, dwFlags uint32, LangId uint16) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceNamesExW, libKernel32, "EnumResourceNamesExW")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), lpEnumFunc, lParam, uintptr(dwFlags), uintptr(LangId))
	return BOOL(ret), WIN32_ERROR(err)
}

func EnumResourceTypesExA(hModule HMODULE, lpEnumFunc ENUMRESTYPEPROCA, lParam uintptr, dwFlags uint32, LangId uint16) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceTypesExA, libKernel32, "EnumResourceTypesExA")
	ret, _, err := syscall.SyscallN(addr, hModule, lpEnumFunc, lParam, uintptr(dwFlags), uintptr(LangId))
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumResourceTypesEx = EnumResourceTypesExW

func EnumResourceTypesExW(hModule HMODULE, lpEnumFunc ENUMRESTYPEPROCW, lParam uintptr, dwFlags uint32, LangId uint16) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceTypesExW, libKernel32, "EnumResourceTypesExW")
	ret, _, err := syscall.SyscallN(addr, hModule, lpEnumFunc, lParam, uintptr(dwFlags), uintptr(LangId))
	return BOOL(ret), WIN32_ERROR(err)
}

var FindResource = FindResourceW

func FindResourceW(hModule HMODULE, lpName PWSTR, lpType PWSTR) HRSRC {
	addr := LazyAddr(&pFindResourceW, libKernel32, "FindResourceW")
	ret, _, _ := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpType)))
	return ret
}

func LoadLibraryA(lpLibFileName PSTR) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pLoadLibraryA, libKernel32, "LoadLibraryA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpLibFileName)))
	return ret, WIN32_ERROR(err)
}

var LoadLibrary = LoadLibraryW

func LoadLibraryW(lpLibFileName PWSTR) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pLoadLibraryW, libKernel32, "LoadLibraryW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpLibFileName)))
	return ret, WIN32_ERROR(err)
}

var EnumResourceNames = EnumResourceNamesW

func EnumResourceNamesW(hModule HMODULE, lpType PWSTR, lpEnumFunc ENUMRESNAMEPROCW, lParam uintptr) BOOL {
	addr := LazyAddr(&pEnumResourceNamesW, libKernel32, "EnumResourceNamesW")
	ret, _, _ := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), lpEnumFunc, lParam)
	return BOOL(ret)
}

func EnumResourceNamesA(hModule HMODULE, lpType PSTR, lpEnumFunc ENUMRESNAMEPROCA, lParam uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceNamesA, libKernel32, "EnumResourceNamesA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

func LoadModule(lpModuleName PSTR, lpParameterBlock unsafe.Pointer) uint32 {
	addr := LazyAddr(&pLoadModule, libKernel32, "LoadModule")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpModuleName)), uintptr(lpParameterBlock))
	return uint32(ret)
}

func LoadPackagedLibrary(lpwLibFileName PWSTR, Reserved uint32) (HMODULE, WIN32_ERROR) {
	addr := LazyAddr(&pLoadPackagedLibrary, libKernel32, "LoadPackagedLibrary")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpwLibFileName)), uintptr(Reserved))
	return ret, WIN32_ERROR(err)
}

func FindResourceA(hModule HMODULE, lpName PSTR, lpType PSTR) (HRSRC, WIN32_ERROR) {
	addr := LazyAddr(&pFindResourceA, libKernel32, "FindResourceA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpType)))
	return ret, WIN32_ERROR(err)
}

func FindResourceExA(hModule HMODULE, lpType PSTR, lpName PSTR, wLanguage uint16) (HRSRC, WIN32_ERROR) {
	addr := LazyAddr(&pFindResourceExA, libKernel32, "FindResourceExA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), uintptr(wLanguage))
	return ret, WIN32_ERROR(err)
}

func EnumResourceTypesA(hModule HMODULE, lpEnumFunc ENUMRESTYPEPROCA, lParam uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceTypesA, libKernel32, "EnumResourceTypesA")
	ret, _, err := syscall.SyscallN(addr, hModule, lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumResourceTypes = EnumResourceTypesW

func EnumResourceTypesW(hModule HMODULE, lpEnumFunc ENUMRESTYPEPROCW, lParam uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceTypesW, libKernel32, "EnumResourceTypesW")
	ret, _, err := syscall.SyscallN(addr, hModule, lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

func EnumResourceLanguagesA(hModule HMODULE, lpType PSTR, lpName PSTR, lpEnumFunc ENUMRESLANGPROCA, lParam uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceLanguagesA, libKernel32, "EnumResourceLanguagesA")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

var EnumResourceLanguages = EnumResourceLanguagesW

func EnumResourceLanguagesW(hModule HMODULE, lpType PWSTR, lpName PWSTR, lpEnumFunc ENUMRESLANGPROCW, lParam uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEnumResourceLanguagesW, libKernel32, "EnumResourceLanguagesW")
	ret, _, err := syscall.SyscallN(addr, hModule, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), lpEnumFunc, lParam)
	return BOOL(ret), WIN32_ERROR(err)
}

func BeginUpdateResourceA(pFileName PSTR, bDeleteExistingResources BOOL) (HANDLE, WIN32_ERROR) {
	addr := LazyAddr(&pBeginUpdateResourceA, libKernel32, "BeginUpdateResourceA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pFileName)), uintptr(bDeleteExistingResources))
	return ret, WIN32_ERROR(err)
}

var BeginUpdateResource = BeginUpdateResourceW

func BeginUpdateResourceW(pFileName PWSTR, bDeleteExistingResources BOOL) (HANDLE, WIN32_ERROR) {
	addr := LazyAddr(&pBeginUpdateResourceW, libKernel32, "BeginUpdateResourceW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pFileName)), uintptr(bDeleteExistingResources))
	return ret, WIN32_ERROR(err)
}

func UpdateResourceA(hUpdate HANDLE, lpType PSTR, lpName PSTR, wLanguage uint16, lpData unsafe.Pointer, cb uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pUpdateResourceA, libKernel32, "UpdateResourceA")
	ret, _, err := syscall.SyscallN(addr, hUpdate, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), uintptr(wLanguage), uintptr(lpData), uintptr(cb))
	return BOOL(ret), WIN32_ERROR(err)
}

var UpdateResource = UpdateResourceW

func UpdateResourceW(hUpdate HANDLE, lpType PWSTR, lpName PWSTR, wLanguage uint16, lpData unsafe.Pointer, cb uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pUpdateResourceW, libKernel32, "UpdateResourceW")
	ret, _, err := syscall.SyscallN(addr, hUpdate, uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpName)), uintptr(wLanguage), uintptr(lpData), uintptr(cb))
	return BOOL(ret), WIN32_ERROR(err)
}

func EndUpdateResourceA(hUpdate HANDLE, fDiscard BOOL) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEndUpdateResourceA, libKernel32, "EndUpdateResourceA")
	ret, _, err := syscall.SyscallN(addr, hUpdate, uintptr(fDiscard))
	return BOOL(ret), WIN32_ERROR(err)
}

var EndUpdateResource = EndUpdateResourceW

func EndUpdateResourceW(hUpdate HANDLE, fDiscard BOOL) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pEndUpdateResourceW, libKernel32, "EndUpdateResourceW")
	ret, _, err := syscall.SyscallN(addr, hUpdate, uintptr(fDiscard))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetDllDirectoryA(lpPathName PSTR) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetDllDirectoryA, libKernel32, "SetDllDirectoryA")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpPathName)))
	return BOOL(ret), WIN32_ERROR(err)
}

var SetDllDirectory = SetDllDirectoryW

func SetDllDirectoryW(lpPathName PWSTR) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetDllDirectoryW, libKernel32, "SetDllDirectoryW")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpPathName)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetDllDirectoryA(nBufferLength uint32, lpBuffer PSTR) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pGetDllDirectoryA, libKernel32, "GetDllDirectoryA")
	ret, _, err := syscall.SyscallN(addr, uintptr(nBufferLength), uintptr(unsafe.Pointer(lpBuffer)))
	return uint32(ret), WIN32_ERROR(err)
}

var GetDllDirectory = GetDllDirectoryW

func GetDllDirectoryW(nBufferLength uint32, lpBuffer PWSTR) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pGetDllDirectoryW, libKernel32, "GetDllDirectoryW")
	ret, _, err := syscall.SyscallN(addr, uintptr(nBufferLength), uintptr(unsafe.Pointer(lpBuffer)))
	return uint32(ret), WIN32_ERROR(err)
}
