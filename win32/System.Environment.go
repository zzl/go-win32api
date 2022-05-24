package win32

import "unsafe"
import "syscall"

const (
	ENCLAVE_RUNTIME_POLICY_ALLOW_FULL_DEBUG uint32 = 1
	ENCLAVE_RUNTIME_POLICY_ALLOW_DYNAMIC_DEBUG uint32 = 2
	ENCLAVE_UNSEAL_FLAG_STALE_KEY uint32 = 1
	ENCLAVE_FLAG_FULL_DEBUG_ENABLED uint32 = 1
	ENCLAVE_FLAG_DYNAMIC_DEBUG_ENABLED uint32 = 2
	ENCLAVE_FLAG_DYNAMIC_DEBUG_ACTIVE uint32 = 4
	VBS_ENCLAVE_REPORT_PKG_HEADER_VERSION_CURRENT uint32 = 1
	VBS_ENCLAVE_REPORT_SIGNATURE_SCHEME_SHA256_RSA_PSS_SHA256 uint32 = 1
	VBS_ENCLAVE_REPORT_VERSION_CURRENT uint32 = 1
	ENCLAVE_REPORT_DATA_LENGTH uint32 = 64
	VBS_ENCLAVE_VARDATA_INVALID uint32 = 0
	VBS_ENCLAVE_VARDATA_MODULE uint32 = 1
	ENCLAVE_VBS_BASIC_KEY_FLAG_MEASUREMENT uint32 = 1
	ENCLAVE_VBS_BASIC_KEY_FLAG_FAMILY_ID uint32 = 2
	ENCLAVE_VBS_BASIC_KEY_FLAG_IMAGE_ID uint32 = 4
	ENCLAVE_VBS_BASIC_KEY_FLAG_DEBUG_KEY uint32 = 8
)

// enums

// enum ENCLAVE_SEALING_IDENTITY_POLICY
type ENCLAVE_SEALING_IDENTITY_POLICY int32
const (
	ENCLAVE_IDENTITY_POLICY_SEAL_INVALID ENCLAVE_SEALING_IDENTITY_POLICY = 0
	ENCLAVE_IDENTITY_POLICY_SEAL_EXACT_CODE ENCLAVE_SEALING_IDENTITY_POLICY = 1
	ENCLAVE_IDENTITY_POLICY_SEAL_SAME_PRIMARY_CODE ENCLAVE_SEALING_IDENTITY_POLICY = 2
	ENCLAVE_IDENTITY_POLICY_SEAL_SAME_IMAGE ENCLAVE_SEALING_IDENTITY_POLICY = 3
	ENCLAVE_IDENTITY_POLICY_SEAL_SAME_FAMILY ENCLAVE_SEALING_IDENTITY_POLICY = 4
	ENCLAVE_IDENTITY_POLICY_SEAL_SAME_AUTHOR ENCLAVE_SEALING_IDENTITY_POLICY = 5
)


// structs

type ENCLAVE_IDENTITY struct {
	OwnerId [32]uint8
	UniqueId [32]uint8
	AuthorId [32]uint8
	FamilyId [16]uint8
	ImageId [16]uint8
	EnclaveSvn uint32
	SecureKernelSvn uint32
	PlatformSvn uint32
	Flags uint32
	SigningLevel uint32
	EnclaveType uint32
}

type VBS_ENCLAVE_REPORT_PKG_HEADER struct {
	PackageSize uint32
	Version uint32
	SignatureScheme uint32
	SignedStatementSize uint32
	SignatureSize uint32
	Reserved uint32
}

type VBS_ENCLAVE_REPORT struct {
	ReportSize uint32
	ReportVersion uint32
	EnclaveData [64]uint8
	EnclaveIdentity ENCLAVE_IDENTITY
}

type VBS_ENCLAVE_REPORT_VARDATA_HEADER struct {
	DataType uint32
	Size uint32
}

type VBS_ENCLAVE_REPORT_MODULE struct {
	Header VBS_ENCLAVE_REPORT_VARDATA_HEADER
	UniqueId [32]uint8
	AuthorId [32]uint8
	FamilyId [16]uint8
	ImageId [16]uint8
	Svn uint32
	ModuleName [1]uint16
}

type ENCLAVE_INFORMATION struct {
	EnclaveType uint32
	Reserved uint32
	BaseAddress unsafe.Pointer
	Size uintptr
	Identity ENCLAVE_IDENTITY
}

type VBS_BASIC_ENCLAVE_THREAD_DESCRIPTOR32 struct {
	ThreadContext [4]uint32
	EntryPoint uint32
	StackPointer uint32
	ExceptionEntryPoint uint32
	ExceptionStack uint32
	ExceptionActive uint32
}

type VBS_BASIC_ENCLAVE_THREAD_DESCRIPTOR64 struct {
	ThreadContext [4]uint64
	EntryPoint uint64
	StackPointer uint64
	ExceptionEntryPoint uint64
	ExceptionStack uint64
	ExceptionActive uint32
}

type VBS_BASIC_ENCLAVE_EXCEPTION_AMD64 struct {
	ExceptionCode uint32
	NumberParameters uint32
	ExceptionInformation [3]uintptr
	ExceptionRAX uintptr
	ExceptionRCX uintptr
	ExceptionRIP uintptr
	ExceptionRFLAGS uintptr
	ExceptionRSP uintptr
}

type ENCLAVE_VBS_BASIC_KEY_REQUEST struct {
	RequestSize uint32
	Flags uint32
	EnclaveSVN uint32
	SystemKeyID uint32
	CurrentSystemKeyID uint32
}

type VBS_BASIC_ENCLAVE_SYSCALL_PAGE struct {
	ReturnFromEnclave uintptr
	ReturnFromException uintptr
	TerminateThread uintptr
	InterruptThread uintptr
	CommitPages uintptr
	DecommitPages uintptr
	ProtectPages uintptr
	CreateThread uintptr
	GetEnclaveInformation uintptr
	GenerateKey uintptr
	GenerateReport uintptr
	VerifyReport uintptr
	GenerateRandomData uintptr
}


// func types

type VBS_BASIC_ENCLAVE_BASIC_CALL_RETURN_FROM_ENCLAVE func(ReturnValue uintptr)

type VBS_BASIC_ENCLAVE_BASIC_CALL_RETURN_FROM_EXCEPTION func(ExceptionRecord *VBS_BASIC_ENCLAVE_EXCEPTION_AMD64) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_TERMINATE_THREAD func(ThreadDescriptor *VBS_BASIC_ENCLAVE_THREAD_DESCRIPTOR64) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_INTERRUPT_THREAD func(ThreadDescriptor *VBS_BASIC_ENCLAVE_THREAD_DESCRIPTOR64) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_COMMIT_PAGES func(EnclaveAddress unsafe.Pointer, NumberOfBytes uintptr, SourceAddress unsafe.Pointer, PageProtection uint32) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_DECOMMIT_PAGES func(EnclaveAddress unsafe.Pointer, NumberOfBytes uintptr) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_PROTECT_PAGES func(EnclaveAddress unsafe.Pointer, NumberOfytes uintptr, PageProtection uint32) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_CREATE_THREAD func(ThreadDescriptor *VBS_BASIC_ENCLAVE_THREAD_DESCRIPTOR64) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_GET_ENCLAVE_INFORMATION func(EnclaveInfo *ENCLAVE_INFORMATION) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_GENERATE_KEY func(KeyRequest *ENCLAVE_VBS_BASIC_KEY_REQUEST, RequestedKeySize uint32, ReturnedKey *uint8) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_GENERATE_REPORT func(EnclaveData *uint8, Report unsafe.Pointer, BufferSize uint32, OutputSize *uint32) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_VERIFY_REPORT func(Report unsafe.Pointer, ReportSize uint32) int32

type VBS_BASIC_ENCLAVE_BASIC_CALL_GENERATE_RANDOM_DATA func(Buffer *uint8, NumberOfBytes uint32, Generation *uint64) int32


var (
	pSetEnvironmentStringsW uintptr
	pGetCommandLineA uintptr
	pGetCommandLineW uintptr
	pGetEnvironmentStrings uintptr
	pGetEnvironmentStringsW uintptr
	pFreeEnvironmentStringsA uintptr
	pFreeEnvironmentStringsW uintptr
	pGetEnvironmentVariableA uintptr
	pGetEnvironmentVariableW uintptr
	pSetEnvironmentVariableA uintptr
	pSetEnvironmentVariableW uintptr
	pExpandEnvironmentStringsA uintptr
	pExpandEnvironmentStringsW uintptr
	pSetCurrentDirectoryA uintptr
	pSetCurrentDirectoryW uintptr
	pGetCurrentDirectoryA uintptr
	pGetCurrentDirectoryW uintptr
	pNeedCurrentDirectoryForExePathA uintptr
	pNeedCurrentDirectoryForExePathW uintptr
	pCreateEnvironmentBlock uintptr
	pDestroyEnvironmentBlock uintptr
	pExpandEnvironmentStringsForUserA uintptr
	pExpandEnvironmentStringsForUserW uintptr
	pIsEnclaveTypeSupported uintptr
	pCreateEnclave uintptr
	pLoadEnclaveData uintptr
	pInitializeEnclave uintptr
)

func SetEnvironmentStringsW(NewEnvironment PWSTR) BOOL {
	addr := lazyAddr(&pSetEnvironmentStringsW, libKernel32, "SetEnvironmentStringsW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(NewEnvironment)))
	return BOOL(ret)
}

func GetCommandLineA() PSTR {
	addr := lazyAddr(&pGetCommandLineA, libKernel32, "GetCommandLineA")
	ret, _,  _ := syscall.SyscallN(addr)
	return (PSTR)(unsafe.Pointer(ret))
}

var GetCommandLine = GetCommandLineW
func GetCommandLineW() PWSTR {
	addr := lazyAddr(&pGetCommandLineW, libKernel32, "GetCommandLineW")
	ret, _,  _ := syscall.SyscallN(addr)
	return (PWSTR)(unsafe.Pointer(ret))
}

func GetEnvironmentStrings() PSTR {
	addr := lazyAddr(&pGetEnvironmentStrings, libKernel32, "GetEnvironmentStrings")
	ret, _,  _ := syscall.SyscallN(addr)
	return (PSTR)(unsafe.Pointer(ret))
}

func GetEnvironmentStringsW() PWSTR {
	addr := lazyAddr(&pGetEnvironmentStringsW, libKernel32, "GetEnvironmentStringsW")
	ret, _,  _ := syscall.SyscallN(addr)
	return (PWSTR)(unsafe.Pointer(ret))
}

func FreeEnvironmentStringsA(penv PSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pFreeEnvironmentStringsA, libKernel32, "FreeEnvironmentStringsA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(penv)))
	return BOOL(ret), WIN32_ERROR(err)
}

var FreeEnvironmentStrings = FreeEnvironmentStringsW
func FreeEnvironmentStringsW(penv PWSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pFreeEnvironmentStringsW, libKernel32, "FreeEnvironmentStringsW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(penv)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetEnvironmentVariableA(lpName PSTR, lpBuffer *uint8, nSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetEnvironmentVariableA, libKernel32, "GetEnvironmentVariableA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpBuffer)), uintptr(nSize))
	return uint32(ret), WIN32_ERROR(err)
}

var GetEnvironmentVariable = GetEnvironmentVariableW
func GetEnvironmentVariableW(lpName PWSTR, lpBuffer *uint16, nSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetEnvironmentVariableW, libKernel32, "GetEnvironmentVariableW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpBuffer)), uintptr(nSize))
	return uint32(ret), WIN32_ERROR(err)
}

func SetEnvironmentVariableA(lpName PSTR, lpValue PSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetEnvironmentVariableA, libKernel32, "SetEnvironmentVariableA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpValue)))
	return BOOL(ret), WIN32_ERROR(err)
}

var SetEnvironmentVariable = SetEnvironmentVariableW
func SetEnvironmentVariableW(lpName PWSTR, lpValue PWSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetEnvironmentVariableW, libKernel32, "SetEnvironmentVariableW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpValue)))
	return BOOL(ret), WIN32_ERROR(err)
}

func ExpandEnvironmentStringsA(lpSrc PSTR, lpDst *uint8, nSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pExpandEnvironmentStringsA, libKernel32, "ExpandEnvironmentStringsA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSrc)), uintptr(unsafe.Pointer(lpDst)), uintptr(nSize))
	return uint32(ret), WIN32_ERROR(err)
}

var ExpandEnvironmentStrings = ExpandEnvironmentStringsW
func ExpandEnvironmentStringsW(lpSrc PWSTR, lpDst *uint16, nSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pExpandEnvironmentStringsW, libKernel32, "ExpandEnvironmentStringsW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSrc)), uintptr(unsafe.Pointer(lpDst)), uintptr(nSize))
	return uint32(ret), WIN32_ERROR(err)
}

func SetCurrentDirectoryA(lpPathName PSTR) BOOL {
	addr := lazyAddr(&pSetCurrentDirectoryA, libKernel32, "SetCurrentDirectoryA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpPathName)))
	return BOOL(ret)
}

var SetCurrentDirectory = SetCurrentDirectoryW
func SetCurrentDirectoryW(lpPathName PWSTR) BOOL {
	addr := lazyAddr(&pSetCurrentDirectoryW, libKernel32, "SetCurrentDirectoryW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpPathName)))
	return BOOL(ret)
}

func GetCurrentDirectoryA(nBufferLength uint32, lpBuffer *uint8) uint32 {
	addr := lazyAddr(&pGetCurrentDirectoryA, libKernel32, "GetCurrentDirectoryA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(nBufferLength), uintptr(unsafe.Pointer(lpBuffer)))
	return uint32(ret)
}

var GetCurrentDirectory = GetCurrentDirectoryW
func GetCurrentDirectoryW(nBufferLength uint32, lpBuffer *uint16) uint32 {
	addr := lazyAddr(&pGetCurrentDirectoryW, libKernel32, "GetCurrentDirectoryW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(nBufferLength), uintptr(unsafe.Pointer(lpBuffer)))
	return uint32(ret)
}

func NeedCurrentDirectoryForExePathA(ExeName PSTR) BOOL {
	addr := lazyAddr(&pNeedCurrentDirectoryForExePathA, libKernel32, "NeedCurrentDirectoryForExePathA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
	return BOOL(ret)
}

var NeedCurrentDirectoryForExePath = NeedCurrentDirectoryForExePathW
func NeedCurrentDirectoryForExePathW(ExeName PWSTR) BOOL {
	addr := lazyAddr(&pNeedCurrentDirectoryForExePathW, libKernel32, "NeedCurrentDirectoryForExePathW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
	return BOOL(ret)
}

func CreateEnvironmentBlock(lpEnvironment unsafe.Pointer, hToken HANDLE, bInherit BOOL) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pCreateEnvironmentBlock, libUserenv, "CreateEnvironmentBlock")
	ret, _,  err := syscall.SyscallN(addr, uintptr(lpEnvironment), hToken, uintptr(bInherit))
	return BOOL(ret), WIN32_ERROR(err)
}

func DestroyEnvironmentBlock(lpEnvironment unsafe.Pointer) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pDestroyEnvironmentBlock, libUserenv, "DestroyEnvironmentBlock")
	ret, _,  err := syscall.SyscallN(addr, uintptr(lpEnvironment))
	return BOOL(ret), WIN32_ERROR(err)
}

func ExpandEnvironmentStringsForUserA(hToken HANDLE, lpSrc PSTR, lpDest *uint8, dwSize uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pExpandEnvironmentStringsForUserA, libUserenv, "ExpandEnvironmentStringsForUserA")
	ret, _,  err := syscall.SyscallN(addr, hToken, uintptr(unsafe.Pointer(lpSrc)), uintptr(unsafe.Pointer(lpDest)), uintptr(dwSize))
	return BOOL(ret), WIN32_ERROR(err)
}

var ExpandEnvironmentStringsForUser = ExpandEnvironmentStringsForUserW
func ExpandEnvironmentStringsForUserW(hToken HANDLE, lpSrc PWSTR, lpDest *uint16, dwSize uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pExpandEnvironmentStringsForUserW, libUserenv, "ExpandEnvironmentStringsForUserW")
	ret, _,  err := syscall.SyscallN(addr, hToken, uintptr(unsafe.Pointer(lpSrc)), uintptr(unsafe.Pointer(lpDest)), uintptr(dwSize))
	return BOOL(ret), WIN32_ERROR(err)
}

func IsEnclaveTypeSupported(flEnclaveType uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pIsEnclaveTypeSupported, libKernel32, "IsEnclaveTypeSupported")
	ret, _,  err := syscall.SyscallN(addr, uintptr(flEnclaveType))
	return BOOL(ret), WIN32_ERROR(err)
}

func CreateEnclave(hProcess HANDLE, lpAddress unsafe.Pointer, dwSize uintptr, dwInitialCommitment uintptr, flEnclaveType uint32, lpEnclaveInformation unsafe.Pointer, dwInfoLength uint32, lpEnclaveError *uint32) (unsafe.Pointer, WIN32_ERROR) {
	addr := lazyAddr(&pCreateEnclave, libKernel32, "CreateEnclave")
	ret, _,  err := syscall.SyscallN(addr, hProcess, uintptr(lpAddress), uintptr(dwSize), uintptr(dwInitialCommitment), uintptr(flEnclaveType), uintptr(lpEnclaveInformation), uintptr(dwInfoLength), uintptr(unsafe.Pointer(lpEnclaveError)))
	return (unsafe.Pointer)(ret), WIN32_ERROR(err)
}

func LoadEnclaveData(hProcess HANDLE, lpAddress unsafe.Pointer, lpBuffer unsafe.Pointer, nSize uintptr, flProtect uint32, lpPageInformation unsafe.Pointer, dwInfoLength uint32, lpNumberOfBytesWritten *uintptr, lpEnclaveError *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pLoadEnclaveData, libKernel32, "LoadEnclaveData")
	ret, _,  err := syscall.SyscallN(addr, hProcess, uintptr(lpAddress), uintptr(lpBuffer), uintptr(nSize), uintptr(flProtect), uintptr(lpPageInformation), uintptr(dwInfoLength), uintptr(unsafe.Pointer(lpNumberOfBytesWritten)), uintptr(unsafe.Pointer(lpEnclaveError)))
	return BOOL(ret), WIN32_ERROR(err)
}

func InitializeEnclave(hProcess HANDLE, lpAddress unsafe.Pointer, lpEnclaveInformation unsafe.Pointer, dwInfoLength uint32, lpEnclaveError *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pInitializeEnclave, libKernel32, "InitializeEnclave")
	ret, _,  err := syscall.SyscallN(addr, hProcess, uintptr(lpAddress), uintptr(lpEnclaveInformation), uintptr(dwInfoLength), uintptr(unsafe.Pointer(lpEnclaveError)))
	return BOOL(ret), WIN32_ERROR(err)
}


