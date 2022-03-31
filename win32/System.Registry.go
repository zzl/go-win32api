package win32

import "unsafe"
import "syscall"

type HKEY = uintptr

const (
	HKEY_CLASSES_ROOT HKEY = ^HKEY(2147483647)
	HKEY_CURRENT_USER HKEY = ^HKEY(2147483646)
	HKEY_LOCAL_MACHINE HKEY = ^HKEY(2147483645)
	HKEY_USERS HKEY = ^HKEY(2147483644)
	HKEY_PERFORMANCE_DATA HKEY = ^HKEY(2147483643)
	HKEY_PERFORMANCE_TEXT HKEY = ^HKEY(2147483567)
	HKEY_PERFORMANCE_NLSTEXT HKEY = ^HKEY(2147483551)
	HKEY_CURRENT_CONFIG HKEY = ^HKEY(2147483642)
	HKEY_DYN_DATA HKEY = ^HKEY(2147483641)
	HKEY_CURRENT_USER_LOCAL_SETTINGS HKEY = ^HKEY(2147483640)
	RRF_SUBKEY_WOW6464KEY uint32 = 65536
	RRF_SUBKEY_WOW6432KEY uint32 = 131072
	RRF_WOW64_MASK uint32 = 196608
	RRF_NOEXPAND uint32 = 268435456
	RRF_ZEROONFAILURE uint32 = 536870912
	REG_PROCESS_APPKEY uint32 = 1
	REG_USE_CURRENT_SECURITY_CONTEXT uint32 = 2
	PROVIDER_KEEPS_VALUE_LENGTH uint32 = 1
	REG_MUI_STRING_TRUNCATE uint32 = 1
	REG_SECURE_CONNECTION uint32 = 1
	REGSTR_MAX_VALUE_LENGTH uint32 = 256
	IT_COMPACT uint32 = 0
	IT_TYPICAL uint32 = 1
	IT_PORTABLE uint32 = 2
	IT_CUSTOM uint32 = 3
	DRIVERSIGN_NONE uint32 = 0
	DRIVERSIGN_WARNING uint32 = 1
	DRIVERSIGN_BLOCKING uint32 = 2
	DOSOPTGF_DEFCLEAN int32 = 1
	DOSOPTF_DEFAULT int32 = 1
	DOSOPTF_SUPPORTED int32 = 2
	DOSOPTF_ALWAYSUSE int32 = 4
	DOSOPTF_USESPMODE int32 = 8
	DOSOPTF_PROVIDESUMB int32 = 16
	DOSOPTF_NEEDSETUP int32 = 32
	DOSOPTF_INDOSSTART int32 = 64
	DOSOPTF_MULTIPLE int32 = 128
	SUF_FIRSTTIME int32 = 1
	SUF_EXPRESS int32 = 2
	SUF_BATCHINF int32 = 4
	SUF_CLEAN int32 = 8
	SUF_INSETUP int32 = 16
	SUF_NETSETUP int32 = 32
	SUF_NETHDBOOT int32 = 64
	SUF_NETRPLBOOT int32 = 128
	SUF_SBSCOPYOK int32 = 256
	VPDF_DISABLEPWRMGMT uint32 = 1
	VPDF_FORCEAPM10MODE uint32 = 2
	VPDF_SKIPINTELSLCHECK uint32 = 4
	VPDF_DISABLEPWRSTATUSPOLL uint32 = 8
	VPDF_DISABLERINGRESUME uint32 = 16
	VPDF_SHOWMULTIBATT uint32 = 32
	BIF_SHOWSIMILARDRIVERS uint32 = 1
	BIF_RAWDEVICENEEDSDRIVER uint32 = 2
	PCMCIA_OPT_HAVE_SOCKET int32 = 1
	PCMCIA_OPT_AUTOMEM int32 = 4
	PCMCIA_OPT_NO_SOUND int32 = 8
	PCMCIA_OPT_NO_AUDIO int32 = 16
	PCMCIA_OPT_NO_APMREMOVE int32 = 32
	PCMCIA_DEF_MEMBEGIN uint32 = 786432
	PCMCIA_DEF_MEMEND uint32 = 16777215
	PCMCIA_DEF_MEMLEN uint32 = 4096
	PCMCIA_DEF_MIN_REGION uint32 = 65536
	PCI_OPTIONS_USE_BIOS int32 = 1
	PCI_OPTIONS_USE_IRQ_STEERING int32 = 2
	AGP_FLAG_NO_1X_RATE int32 = 1
	AGP_FLAG_NO_2X_RATE int32 = 2
	AGP_FLAG_NO_4X_RATE int32 = 4
	AGP_FLAG_NO_8X_RATE int32 = 8
	AGP_FLAG_REVERSE_INITIALIZATION int32 = 128
	AGP_FLAG_NO_SBA_ENABLE int32 = 256
	AGP_FLAG_NO_FW_ENABLE int32 = 512
	AGP_FLAG_SPECIAL_TARGET int32 = 1048575
	AGP_FLAG_SPECIAL_RESERVE int32 = 1015808
	REGSTR_VAL_MAX_HCID_LEN uint32 = 1024
	REGDF_NOTDETIO uint32 = 1
	REGDF_NOTDETMEM uint32 = 2
	REGDF_NOTDETIRQ uint32 = 4
	REGDF_NOTDETDMA uint32 = 8
	REGDF_NEEDFULLCONFIG uint32 = 16
	REGDF_GENFORCEDCONFIG uint32 = 32
	REGDF_NODETCONFIG uint32 = 32768
	REGDF_CONFLICTIO uint32 = 65536
	REGDF_CONFLICTMEM uint32 = 131072
	REGDF_CONFLICTIRQ uint32 = 262144
	REGDF_CONFLICTDMA uint32 = 524288
	REGDF_MAPIRQ2TO9 uint32 = 1048576
	REGDF_NOTVERIFIED uint32 = 2147483648
	APMMENUSUSPEND_DISABLED uint32 = 0
	APMMENUSUSPEND_ENABLED uint32 = 1
	APMMENUSUSPEND_UNDOCKED uint32 = 2
	APMMENUSUSPEND_NOCHANGE uint32 = 128
	APMTIMEOUT_DISABLED uint32 = 0
	CONFIGFLAG_DISABLED uint32 = 1
	CONFIGFLAG_REMOVED uint32 = 2
	CONFIGFLAG_MANUAL_INSTALL uint32 = 4
	CONFIGFLAG_IGNORE_BOOT_LC uint32 = 8
	CONFIGFLAG_NET_BOOT uint32 = 16
	CONFIGFLAG_REINSTALL uint32 = 32
	CONFIGFLAG_FAILEDINSTALL uint32 = 64
	CONFIGFLAG_CANTSTOPACHILD uint32 = 128
	CONFIGFLAG_OKREMOVEROM uint32 = 256
	CONFIGFLAG_NOREMOVEEXIT uint32 = 512
	CONFIGFLAG_FINISH_INSTALL uint32 = 1024
	CONFIGFLAG_NEEDS_FORCED_CONFIG uint32 = 2048
	CONFIGFLAG_NETBOOT_CARD uint32 = 4096
	CONFIGFLAG_PARTIAL_LOG_CONF uint32 = 8192
	CONFIGFLAG_SUPPRESS_SURPRISE uint32 = 16384
	CONFIGFLAG_VERIFY_HARDWARE uint32 = 32768
	CONFIGFLAG_FINISHINSTALL_UI uint32 = 65536
	CONFIGFLAG_FINISHINSTALL_ACTION uint32 = 131072
	CONFIGFLAG_BOOT_DEVICE uint32 = 262144
	CONFIGFLAG_NEEDS_CLASS_CONFIG uint32 = 524288
	CSCONFIGFLAG_BITS uint32 = 7
	CSCONFIGFLAG_DISABLED uint32 = 1
	CSCONFIGFLAG_DO_NOT_CREATE uint32 = 2
	CSCONFIGFLAG_DO_NOT_START uint32 = 4
	DMSTATEFLAG_APPLYTOALL uint32 = 1
	NUM_RESOURCE_MAP uint32 = 256
	MF_FLAGS_EVEN_IF_NO_RESOURCE uint32 = 1
	MF_FLAGS_NO_CREATE_IF_NO_RESOURCE uint32 = 2
	MF_FLAGS_FILL_IN_UNKNOWN_RESOURCE uint32 = 4
	MF_FLAGS_CREATE_BUT_NO_SHOW_DISABLED uint32 = 8
	EISAFLAG_NO_IO_MERGE uint32 = 1
	EISAFLAG_SLOT_IO_FIRST uint32 = 2
	EISA_NO_MAX_FUNCTION uint32 = 255
	NUM_EISA_RANGES uint32 = 4
	PCIC_DEFAULT_IRQMASK uint32 = 20152
	PCIC_DEFAULT_NUMSOCKETS uint32 = 0
	DTRESULTOK uint32 = 0
	DTRESULTFIX uint32 = 1
	DTRESULTPROB uint32 = 2
	DTRESULTPART uint32 = 3
	PIR_OPTION_ENABLED uint32 = 1
	PIR_OPTION_REGISTRY uint32 = 2
	PIR_OPTION_MSSPEC uint32 = 4
	PIR_OPTION_REALMODE uint32 = 8
	PIR_OPTION_DEFAULT uint32 = 15
	PIR_STATUS_ERROR uint32 = 0
	PIR_STATUS_ENABLED uint32 = 1
	PIR_STATUS_DISABLED uint32 = 2
	PIR_STATUS_MAX uint32 = 3
	PIR_STATUS_TABLE_REGISTRY uint32 = 0
	PIR_STATUS_TABLE_MSSPEC uint32 = 1
	PIR_STATUS_TABLE_REALMODE uint32 = 2
	PIR_STATUS_TABLE_NONE uint32 = 3
	PIR_STATUS_TABLE_ERROR uint32 = 4
	PIR_STATUS_TABLE_BAD uint32 = 5
	PIR_STATUS_TABLE_SUCCESS uint32 = 6
	PIR_STATUS_TABLE_MAX uint32 = 7
	PIR_STATUS_MINIPORT_NORMAL uint32 = 0
	PIR_STATUS_MINIPORT_COMPATIBLE uint32 = 1
	PIR_STATUS_MINIPORT_OVERRIDE uint32 = 2
	PIR_STATUS_MINIPORT_NONE uint32 = 3
	PIR_STATUS_MINIPORT_ERROR uint32 = 4
	PIR_STATUS_MINIPORT_NOKEY uint32 = 5
	PIR_STATUS_MINIPORT_SUCCESS uint32 = 6
	PIR_STATUS_MINIPORT_INVALID uint32 = 7
	PIR_STATUS_MINIPORT_MAX uint32 = 8
	LASTGOOD_OPERATION uint32 = 255
	LASTGOOD_OPERATION_NOPOSTPROC uint32 = 0
	LASTGOOD_OPERATION_DELETE uint32 = 1
)

// enums

// enum REG_VALUE_TYPE
type REG_VALUE_TYPE uint32
const (
	REG_NONE REG_VALUE_TYPE = 0
	REG_SZ REG_VALUE_TYPE = 1
	REG_EXPAND_SZ REG_VALUE_TYPE = 2
	REG_BINARY REG_VALUE_TYPE = 3
	REG_DWORD REG_VALUE_TYPE = 4
	REG_DWORD_LITTLE_ENDIAN REG_VALUE_TYPE = 4
	REG_DWORD_BIG_ENDIAN REG_VALUE_TYPE = 5
	REG_LINK REG_VALUE_TYPE = 6
	REG_MULTI_SZ REG_VALUE_TYPE = 7
	REG_RESOURCE_LIST REG_VALUE_TYPE = 8
	REG_FULL_RESOURCE_DESCRIPTOR REG_VALUE_TYPE = 9
	REG_RESOURCE_REQUIREMENTS_LIST REG_VALUE_TYPE = 10
	REG_QWORD REG_VALUE_TYPE = 11
	REG_QWORD_LITTLE_ENDIAN REG_VALUE_TYPE = 11
)

// enum REG_SAM_FLAGS
// flags
type REG_SAM_FLAGS uint32
const (
	KEY_QUERY_VALUE REG_SAM_FLAGS = 1
	KEY_SET_VALUE REG_SAM_FLAGS = 2
	KEY_CREATE_SUB_KEY REG_SAM_FLAGS = 4
	KEY_ENUMERATE_SUB_KEYS REG_SAM_FLAGS = 8
	KEY_NOTIFY REG_SAM_FLAGS = 16
	KEY_CREATE_LINK REG_SAM_FLAGS = 32
	KEY_WOW64_32KEY REG_SAM_FLAGS = 512
	KEY_WOW64_64KEY REG_SAM_FLAGS = 256
	KEY_WOW64_RES REG_SAM_FLAGS = 768
	KEY_READ REG_SAM_FLAGS = 131097
	KEY_WRITE REG_SAM_FLAGS = 131078
	KEY_EXECUTE REG_SAM_FLAGS = 131097
	KEY_ALL_ACCESS REG_SAM_FLAGS = 983103
)

// enum REG_OPEN_CREATE_OPTIONS
// flags
type REG_OPEN_CREATE_OPTIONS uint32
const (
	REG_OPTION_RESERVED REG_OPEN_CREATE_OPTIONS = 0
	REG_OPTION_NON_VOLATILE REG_OPEN_CREATE_OPTIONS = 0
	REG_OPTION_VOLATILE REG_OPEN_CREATE_OPTIONS = 1
	REG_OPTION_CREATE_LINK REG_OPEN_CREATE_OPTIONS = 2
	REG_OPTION_BACKUP_RESTORE REG_OPEN_CREATE_OPTIONS = 4
	REG_OPTION_OPEN_LINK REG_OPEN_CREATE_OPTIONS = 8
	REG_OPTION_DONT_VIRTUALIZE REG_OPEN_CREATE_OPTIONS = 16
)

// enum REG_CREATE_KEY_DISPOSITION
type REG_CREATE_KEY_DISPOSITION uint32
const (
	REG_CREATED_NEW_KEY REG_CREATE_KEY_DISPOSITION = 1
	REG_OPENED_EXISTING_KEY REG_CREATE_KEY_DISPOSITION = 2
)

// enum REG_SAVE_FORMAT
type REG_SAVE_FORMAT uint32
const (
	REG_STANDARD_FORMAT REG_SAVE_FORMAT = 1
	REG_LATEST_FORMAT REG_SAVE_FORMAT = 2
	REG_NO_COMPRESSION REG_SAVE_FORMAT = 4
)

// enum REG_RESTORE_KEY_FLAGS
type REG_RESTORE_KEY_FLAGS int32
const (
	REG_FORCE_RESTORE REG_RESTORE_KEY_FLAGS = 8
	REG_WHOLE_HIVE_VOLATILE REG_RESTORE_KEY_FLAGS = 1
)

// enum REG_NOTIFY_FILTER
// flags
type REG_NOTIFY_FILTER uint32
const (
	REG_NOTIFY_CHANGE_NAME REG_NOTIFY_FILTER = 1
	REG_NOTIFY_CHANGE_ATTRIBUTES REG_NOTIFY_FILTER = 2
	REG_NOTIFY_CHANGE_LAST_SET REG_NOTIFY_FILTER = 4
	REG_NOTIFY_CHANGE_SECURITY REG_NOTIFY_FILTER = 8
	REG_NOTIFY_THREAD_AGNOSTIC REG_NOTIFY_FILTER = 268435456
)

// enum RRF_RT
// flags
type RRF_RT uint32
const (
	RRF_RT_ANY RRF_RT = 65535
	RRF_RT_DWORD RRF_RT = 24
	RRF_RT_QWORD RRF_RT = 72
	RRF_RT_REG_BINARY RRF_RT = 8
	RRF_RT_REG_DWORD RRF_RT = 16
	RRF_RT_REG_EXPAND_SZ RRF_RT = 4
	RRF_RT_REG_MULTI_SZ RRF_RT = 32
	RRF_RT_REG_NONE RRF_RT = 1
	RRF_RT_REG_QWORD RRF_RT = 64
	RRF_RT_REG_SZ RRF_RT = 2
)


// structs

type Val_context struct {
	Valuelen int32
	Value_context unsafe.Pointer
	Val_buff_ptr unsafe.Pointer
}

type PvalueA struct {
	Pv_valuename PSTR
	Pv_valuelen int32
	Pv_value_context unsafe.Pointer
	Pv_type uint32
}

type Pvalue = PvalueW
type PvalueW struct {
	Pv_valuename PWSTR
	Pv_valuelen int32
	Pv_value_context unsafe.Pointer
	Pv_type uint32
}

type Provider_info struct {
	Pi_R0_1val uintptr
	Pi_R0_allvals uintptr
	Pi_R3_1val uintptr
	Pi_R3_allvals uintptr
	Pi_flags uint32
	Pi_key_context unsafe.Pointer
}

type VALENTA struct {
	Ve_valuename PSTR
	Ve_valuelen uint32
	Ve_valueptr uintptr
	Ve_type REG_VALUE_TYPE
}

type VALENT = VALENTW
type VALENTW struct {
	Ve_valuename PWSTR
	Ve_valuelen uint32
	Ve_valueptr uintptr
	Ve_type REG_VALUE_TYPE
}

type DSKTLSYSTEMTIME struct {
	WYear uint16
	WMonth uint16
	WDayOfWeek uint16
	WDay uint16
	WHour uint16
	WMinute uint16
	WSecond uint16
	WMilliseconds uint16
	WResult uint16
}


// func types

type PQUERYHANDLER func(keycontext unsafe.Pointer, val_list *Val_context, num_vals uint32, outputbuffer unsafe.Pointer, total_outlen *uint32, input_blen uint32) uint32


var (
	pRegCloseKey uintptr
	pRegOverridePredefKey uintptr
	pRegOpenUserClassesRoot uintptr
	pRegOpenCurrentUser uintptr
	pRegDisablePredefinedCache uintptr
	pRegDisablePredefinedCacheEx uintptr
	pRegConnectRegistryA uintptr
	pRegConnectRegistryW uintptr
	pRegConnectRegistryExA uintptr
	pRegConnectRegistryExW uintptr
	pRegCreateKeyA uintptr
	pRegCreateKeyW uintptr
	pRegCreateKeyExA uintptr
	pRegCreateKeyExW uintptr
	pRegCreateKeyTransactedA uintptr
	pRegCreateKeyTransactedW uintptr
	pRegDeleteKeyA uintptr
	pRegDeleteKeyW uintptr
	pRegDeleteKeyExA uintptr
	pRegDeleteKeyExW uintptr
	pRegDeleteKeyTransactedA uintptr
	pRegDeleteKeyTransactedW uintptr
	pRegDisableReflectionKey uintptr
	pRegEnableReflectionKey uintptr
	pRegQueryReflectionKey uintptr
	pRegDeleteValueA uintptr
	pRegDeleteValueW uintptr
	pRegEnumKeyA uintptr
	pRegEnumKeyW uintptr
	pRegEnumKeyExA uintptr
	pRegEnumKeyExW uintptr
	pRegEnumValueA uintptr
	pRegEnumValueW uintptr
	pRegFlushKey uintptr
	pRegGetKeySecurity uintptr
	pRegLoadKeyA uintptr
	pRegLoadKeyW uintptr
	pRegNotifyChangeKeyValue uintptr
	pRegOpenKeyA uintptr
	pRegOpenKeyW uintptr
	pRegOpenKeyExA uintptr
	pRegOpenKeyExW uintptr
	pRegOpenKeyTransactedA uintptr
	pRegOpenKeyTransactedW uintptr
	pRegQueryInfoKeyA uintptr
	pRegQueryInfoKeyW uintptr
	pRegQueryValueA uintptr
	pRegQueryValueW uintptr
	pRegQueryMultipleValuesA uintptr
	pRegQueryMultipleValuesW uintptr
	pRegQueryValueExA uintptr
	pRegQueryValueExW uintptr
	pRegReplaceKeyA uintptr
	pRegReplaceKeyW uintptr
	pRegRestoreKeyA uintptr
	pRegRestoreKeyW uintptr
	pRegRenameKey uintptr
	pRegSaveKeyA uintptr
	pRegSaveKeyW uintptr
	pRegSetKeySecurity uintptr
	pRegSetValueA uintptr
	pRegSetValueW uintptr
	pRegSetValueExA uintptr
	pRegSetValueExW uintptr
	pRegUnLoadKeyA uintptr
	pRegUnLoadKeyW uintptr
	pRegDeleteKeyValueA uintptr
	pRegDeleteKeyValueW uintptr
	pRegSetKeyValueA uintptr
	pRegSetKeyValueW uintptr
	pRegDeleteTreeA uintptr
	pRegDeleteTreeW uintptr
	pRegCopyTreeA uintptr
	pRegGetValueA uintptr
	pRegGetValueW uintptr
	pRegCopyTreeW uintptr
	pRegLoadMUIStringA uintptr
	pRegLoadMUIStringW uintptr
	pRegLoadAppKeyA uintptr
	pRegLoadAppKeyW uintptr
	pRegSaveKeyExA uintptr
	pRegSaveKeyExW uintptr
)

func RegCloseKey(hKey HKEY) LSTATUS {
	addr := lazyAddr(&pRegCloseKey, libAdvapi32, "RegCloseKey")
	ret, _,  _ := syscall.SyscallN(addr, hKey)
	return LSTATUS(ret)
}

func RegOverridePredefKey(hKey HKEY, hNewHKey HKEY) LSTATUS {
	addr := lazyAddr(&pRegOverridePredefKey, libAdvapi32, "RegOverridePredefKey")
	ret, _,  _ := syscall.SyscallN(addr, hKey, hNewHKey)
	return LSTATUS(ret)
}

func RegOpenUserClassesRoot(hToken HANDLE, dwOptions uint32, samDesired uint32, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegOpenUserClassesRoot, libAdvapi32, "RegOpenUserClassesRoot")
	ret, _,  _ := syscall.SyscallN(addr, hToken, uintptr(dwOptions), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegOpenCurrentUser(samDesired uint32, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegOpenCurrentUser, libAdvapi32, "RegOpenCurrentUser")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegDisablePredefinedCache() LSTATUS {
	addr := lazyAddr(&pRegDisablePredefinedCache, libAdvapi32, "RegDisablePredefinedCache")
	ret, _,  _ := syscall.SyscallN(addr)
	return LSTATUS(ret)
}

func RegDisablePredefinedCacheEx() LSTATUS {
	addr := lazyAddr(&pRegDisablePredefinedCacheEx, libAdvapi32, "RegDisablePredefinedCacheEx")
	ret, _,  _ := syscall.SyscallN(addr)
	return LSTATUS(ret)
}

func RegConnectRegistryA(lpMachineName PSTR, hKey HKEY, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegConnectRegistryA, libAdvapi32, "RegConnectRegistryA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpMachineName)), hKey, uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

var RegConnectRegistry = RegConnectRegistryW
func RegConnectRegistryW(lpMachineName PWSTR, hKey HKEY, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegConnectRegistryW, libAdvapi32, "RegConnectRegistryW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpMachineName)), hKey, uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegConnectRegistryExA(lpMachineName PSTR, hKey HKEY, Flags uint32, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegConnectRegistryExA, libAdvapi32, "RegConnectRegistryExA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpMachineName)), hKey, uintptr(Flags), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

var RegConnectRegistryEx = RegConnectRegistryExW
func RegConnectRegistryExW(lpMachineName PWSTR, hKey HKEY, Flags uint32, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegConnectRegistryExW, libAdvapi32, "RegConnectRegistryExW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpMachineName)), hKey, uintptr(Flags), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegCreateKeyA(hKey HKEY, lpSubKey PSTR, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegCreateKeyA, libAdvapi32, "RegCreateKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

var RegCreateKey = RegCreateKeyW
func RegCreateKeyW(hKey HKEY, lpSubKey PWSTR, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegCreateKeyW, libAdvapi32, "RegCreateKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegCreateKeyExA(hKey HKEY, lpSubKey PSTR, Reserved uint32, lpClass PSTR, dwOptions REG_OPEN_CREATE_OPTIONS, samDesired REG_SAM_FLAGS, lpSecurityAttributes *SECURITY_ATTRIBUTES, phkResult *HKEY, lpdwDisposition *REG_CREATE_KEY_DISPOSITION) LSTATUS {
	addr := lazyAddr(&pRegCreateKeyExA, libAdvapi32, "RegCreateKeyExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(Reserved), uintptr(unsafe.Pointer(lpClass)), uintptr(dwOptions), uintptr(samDesired), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(unsafe.Pointer(phkResult)), uintptr(unsafe.Pointer(lpdwDisposition)))
	return LSTATUS(ret)
}

var RegCreateKeyEx = RegCreateKeyExW
func RegCreateKeyExW(hKey HKEY, lpSubKey PWSTR, Reserved uint32, lpClass PWSTR, dwOptions REG_OPEN_CREATE_OPTIONS, samDesired REG_SAM_FLAGS, lpSecurityAttributes *SECURITY_ATTRIBUTES, phkResult *HKEY, lpdwDisposition *REG_CREATE_KEY_DISPOSITION) LSTATUS {
	addr := lazyAddr(&pRegCreateKeyExW, libAdvapi32, "RegCreateKeyExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(Reserved), uintptr(unsafe.Pointer(lpClass)), uintptr(dwOptions), uintptr(samDesired), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(unsafe.Pointer(phkResult)), uintptr(unsafe.Pointer(lpdwDisposition)))
	return LSTATUS(ret)
}

func RegCreateKeyTransactedA(hKey HKEY, lpSubKey PSTR, Reserved uint32, lpClass PSTR, dwOptions REG_OPEN_CREATE_OPTIONS, samDesired REG_SAM_FLAGS, lpSecurityAttributes *SECURITY_ATTRIBUTES, phkResult *HKEY, lpdwDisposition *REG_CREATE_KEY_DISPOSITION, hTransaction HANDLE, pExtendedParemeter unsafe.Pointer) LSTATUS {
	addr := lazyAddr(&pRegCreateKeyTransactedA, libAdvapi32, "RegCreateKeyTransactedA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(Reserved), uintptr(unsafe.Pointer(lpClass)), uintptr(dwOptions), uintptr(samDesired), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(unsafe.Pointer(phkResult)), uintptr(unsafe.Pointer(lpdwDisposition)), hTransaction, uintptr(pExtendedParemeter))
	return LSTATUS(ret)
}

var RegCreateKeyTransacted = RegCreateKeyTransactedW
func RegCreateKeyTransactedW(hKey HKEY, lpSubKey PWSTR, Reserved uint32, lpClass PWSTR, dwOptions REG_OPEN_CREATE_OPTIONS, samDesired REG_SAM_FLAGS, lpSecurityAttributes *SECURITY_ATTRIBUTES, phkResult *HKEY, lpdwDisposition *REG_CREATE_KEY_DISPOSITION, hTransaction HANDLE, pExtendedParemeter unsafe.Pointer) LSTATUS {
	addr := lazyAddr(&pRegCreateKeyTransactedW, libAdvapi32, "RegCreateKeyTransactedW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(Reserved), uintptr(unsafe.Pointer(lpClass)), uintptr(dwOptions), uintptr(samDesired), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(unsafe.Pointer(phkResult)), uintptr(unsafe.Pointer(lpdwDisposition)), hTransaction, uintptr(pExtendedParemeter))
	return LSTATUS(ret)
}

func RegDeleteKeyA(hKey HKEY, lpSubKey PSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyA, libAdvapi32, "RegDeleteKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)))
	return LSTATUS(ret)
}

var RegDeleteKey = RegDeleteKeyW
func RegDeleteKeyW(hKey HKEY, lpSubKey PWSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyW, libAdvapi32, "RegDeleteKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)))
	return LSTATUS(ret)
}

func RegDeleteKeyExA(hKey HKEY, lpSubKey PSTR, samDesired uint32, Reserved uint32) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyExA, libAdvapi32, "RegDeleteKeyExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(samDesired), uintptr(Reserved))
	return LSTATUS(ret)
}

var RegDeleteKeyEx = RegDeleteKeyExW
func RegDeleteKeyExW(hKey HKEY, lpSubKey PWSTR, samDesired uint32, Reserved uint32) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyExW, libAdvapi32, "RegDeleteKeyExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(samDesired), uintptr(Reserved))
	return LSTATUS(ret)
}

func RegDeleteKeyTransactedA(hKey HKEY, lpSubKey PSTR, samDesired uint32, Reserved uint32, hTransaction HANDLE, pExtendedParameter unsafe.Pointer) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyTransactedA, libAdvapi32, "RegDeleteKeyTransactedA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(samDesired), uintptr(Reserved), hTransaction, uintptr(pExtendedParameter))
	return LSTATUS(ret)
}

var RegDeleteKeyTransacted = RegDeleteKeyTransactedW
func RegDeleteKeyTransactedW(hKey HKEY, lpSubKey PWSTR, samDesired uint32, Reserved uint32, hTransaction HANDLE, pExtendedParameter unsafe.Pointer) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyTransactedW, libAdvapi32, "RegDeleteKeyTransactedW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(samDesired), uintptr(Reserved), hTransaction, uintptr(pExtendedParameter))
	return LSTATUS(ret)
}

func RegDisableReflectionKey(hBase HKEY) int32 {
	addr := lazyAddr(&pRegDisableReflectionKey, libAdvapi32, "RegDisableReflectionKey")
	ret, _,  _ := syscall.SyscallN(addr, hBase)
	return int32(ret)
}

func RegEnableReflectionKey(hBase HKEY) int32 {
	addr := lazyAddr(&pRegEnableReflectionKey, libAdvapi32, "RegEnableReflectionKey")
	ret, _,  _ := syscall.SyscallN(addr, hBase)
	return int32(ret)
}

func RegQueryReflectionKey(hBase HKEY, bIsReflectionDisabled *BOOL) int32 {
	addr := lazyAddr(&pRegQueryReflectionKey, libAdvapi32, "RegQueryReflectionKey")
	ret, _,  _ := syscall.SyscallN(addr, hBase, uintptr(unsafe.Pointer(bIsReflectionDisabled)))
	return int32(ret)
}

func RegDeleteValueA(hKey HKEY, lpValueName PSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteValueA, libAdvapi32, "RegDeleteValueA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpValueName)))
	return LSTATUS(ret)
}

var RegDeleteValue = RegDeleteValueW
func RegDeleteValueW(hKey HKEY, lpValueName PWSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteValueW, libAdvapi32, "RegDeleteValueW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpValueName)))
	return LSTATUS(ret)
}

func RegEnumKeyA(hKey HKEY, dwIndex uint32, lpName *uint8, cchName uint32) LSTATUS {
	addr := lazyAddr(&pRegEnumKeyA, libAdvapi32, "RegEnumKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(dwIndex), uintptr(unsafe.Pointer(lpName)), uintptr(cchName))
	return LSTATUS(ret)
}

var RegEnumKey = RegEnumKeyW
func RegEnumKeyW(hKey HKEY, dwIndex uint32, lpName *uint16, cchName uint32) LSTATUS {
	addr := lazyAddr(&pRegEnumKeyW, libAdvapi32, "RegEnumKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(dwIndex), uintptr(unsafe.Pointer(lpName)), uintptr(cchName))
	return LSTATUS(ret)
}

func RegEnumKeyExA(hKey HKEY, dwIndex uint32, lpName *uint8, lpcchName *uint32, lpReserved *uint32, lpClass *uint8, lpcchClass *uint32, lpftLastWriteTime *FILETIME) LSTATUS {
	addr := lazyAddr(&pRegEnumKeyExA, libAdvapi32, "RegEnumKeyExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(dwIndex), uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpcchName)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpClass)), uintptr(unsafe.Pointer(lpcchClass)), uintptr(unsafe.Pointer(lpftLastWriteTime)))
	return LSTATUS(ret)
}

var RegEnumKeyEx = RegEnumKeyExW
func RegEnumKeyExW(hKey HKEY, dwIndex uint32, lpName *uint16, lpcchName *uint32, lpReserved *uint32, lpClass *uint16, lpcchClass *uint32, lpftLastWriteTime *FILETIME) LSTATUS {
	addr := lazyAddr(&pRegEnumKeyExW, libAdvapi32, "RegEnumKeyExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(dwIndex), uintptr(unsafe.Pointer(lpName)), uintptr(unsafe.Pointer(lpcchName)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpClass)), uintptr(unsafe.Pointer(lpcchClass)), uintptr(unsafe.Pointer(lpftLastWriteTime)))
	return LSTATUS(ret)
}

func RegEnumValueA(hKey HKEY, dwIndex uint32, lpValueName *uint8, lpcchValueName *uint32, lpReserved *uint32, lpType *uint32, lpData *uint8, lpcbData *uint32) LSTATUS {
	addr := lazyAddr(&pRegEnumValueA, libAdvapi32, "RegEnumValueA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(dwIndex), uintptr(unsafe.Pointer(lpValueName)), uintptr(unsafe.Pointer(lpcchValueName)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpData)), uintptr(unsafe.Pointer(lpcbData)))
	return LSTATUS(ret)
}

var RegEnumValue = RegEnumValueW
func RegEnumValueW(hKey HKEY, dwIndex uint32, lpValueName *uint16, lpcchValueName *uint32, lpReserved *uint32, lpType *uint32, lpData *uint8, lpcbData *uint32) LSTATUS {
	addr := lazyAddr(&pRegEnumValueW, libAdvapi32, "RegEnumValueW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(dwIndex), uintptr(unsafe.Pointer(lpValueName)), uintptr(unsafe.Pointer(lpcchValueName)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpData)), uintptr(unsafe.Pointer(lpcbData)))
	return LSTATUS(ret)
}

func RegFlushKey(hKey HKEY) LSTATUS {
	addr := lazyAddr(&pRegFlushKey, libAdvapi32, "RegFlushKey")
	ret, _,  _ := syscall.SyscallN(addr, hKey)
	return LSTATUS(ret)
}

func RegGetKeySecurity(hKey HKEY, SecurityInformation uint32, pSecurityDescriptor *SECURITY_DESCRIPTOR, lpcbSecurityDescriptor *uint32) LSTATUS {
	addr := lazyAddr(&pRegGetKeySecurity, libAdvapi32, "RegGetKeySecurity")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(SecurityInformation), uintptr(unsafe.Pointer(pSecurityDescriptor)), uintptr(unsafe.Pointer(lpcbSecurityDescriptor)))
	return LSTATUS(ret)
}

func RegLoadKeyA(hKey HKEY, lpSubKey PSTR, lpFile PSTR) LSTATUS {
	addr := lazyAddr(&pRegLoadKeyA, libAdvapi32, "RegLoadKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpFile)))
	return LSTATUS(ret)
}

var RegLoadKey = RegLoadKeyW
func RegLoadKeyW(hKey HKEY, lpSubKey PWSTR, lpFile PWSTR) LSTATUS {
	addr := lazyAddr(&pRegLoadKeyW, libAdvapi32, "RegLoadKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpFile)))
	return LSTATUS(ret)
}

func RegNotifyChangeKeyValue(hKey HKEY, bWatchSubtree BOOL, dwNotifyFilter REG_NOTIFY_FILTER, hEvent HANDLE, fAsynchronous BOOL) LSTATUS {
	addr := lazyAddr(&pRegNotifyChangeKeyValue, libAdvapi32, "RegNotifyChangeKeyValue")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(bWatchSubtree), uintptr(dwNotifyFilter), hEvent, uintptr(fAsynchronous))
	return LSTATUS(ret)
}

func RegOpenKeyA(hKey HKEY, lpSubKey PSTR, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegOpenKeyA, libAdvapi32, "RegOpenKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

var RegOpenKey = RegOpenKeyW
func RegOpenKeyW(hKey HKEY, lpSubKey PWSTR, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegOpenKeyW, libAdvapi32, "RegOpenKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegOpenKeyExA(hKey HKEY, lpSubKey PSTR, ulOptions uint32, samDesired REG_SAM_FLAGS, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegOpenKeyExA, libAdvapi32, "RegOpenKeyExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

var RegOpenKeyEx = RegOpenKeyExW
func RegOpenKeyExW(hKey HKEY, lpSubKey PWSTR, ulOptions uint32, samDesired REG_SAM_FLAGS, phkResult *HKEY) LSTATUS {
	addr := lazyAddr(&pRegOpenKeyExW, libAdvapi32, "RegOpenKeyExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)))
	return LSTATUS(ret)
}

func RegOpenKeyTransactedA(hKey HKEY, lpSubKey PSTR, ulOptions uint32, samDesired REG_SAM_FLAGS, phkResult *HKEY, hTransaction HANDLE, pExtendedParemeter unsafe.Pointer) LSTATUS {
	addr := lazyAddr(&pRegOpenKeyTransactedA, libAdvapi32, "RegOpenKeyTransactedA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)), hTransaction, uintptr(pExtendedParemeter))
	return LSTATUS(ret)
}

var RegOpenKeyTransacted = RegOpenKeyTransactedW
func RegOpenKeyTransactedW(hKey HKEY, lpSubKey PWSTR, ulOptions uint32, samDesired REG_SAM_FLAGS, phkResult *HKEY, hTransaction HANDLE, pExtendedParemeter unsafe.Pointer) LSTATUS {
	addr := lazyAddr(&pRegOpenKeyTransactedW, libAdvapi32, "RegOpenKeyTransactedW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)), hTransaction, uintptr(pExtendedParemeter))
	return LSTATUS(ret)
}

func RegQueryInfoKeyA(hKey HKEY, lpClass *uint8, lpcchClass *uint32, lpReserved *uint32, lpcSubKeys *uint32, lpcbMaxSubKeyLen *uint32, lpcbMaxClassLen *uint32, lpcValues *uint32, lpcbMaxValueNameLen *uint32, lpcbMaxValueLen *uint32, lpcbSecurityDescriptor *uint32, lpftLastWriteTime *FILETIME) LSTATUS {
	addr := lazyAddr(&pRegQueryInfoKeyA, libAdvapi32, "RegQueryInfoKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpClass)), uintptr(unsafe.Pointer(lpcchClass)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpcSubKeys)), uintptr(unsafe.Pointer(lpcbMaxSubKeyLen)), uintptr(unsafe.Pointer(lpcbMaxClassLen)), uintptr(unsafe.Pointer(lpcValues)), uintptr(unsafe.Pointer(lpcbMaxValueNameLen)), uintptr(unsafe.Pointer(lpcbMaxValueLen)), uintptr(unsafe.Pointer(lpcbSecurityDescriptor)), uintptr(unsafe.Pointer(lpftLastWriteTime)))
	return LSTATUS(ret)
}

var RegQueryInfoKey = RegQueryInfoKeyW
func RegQueryInfoKeyW(hKey HKEY, lpClass *uint16, lpcchClass *uint32, lpReserved *uint32, lpcSubKeys *uint32, lpcbMaxSubKeyLen *uint32, lpcbMaxClassLen *uint32, lpcValues *uint32, lpcbMaxValueNameLen *uint32, lpcbMaxValueLen *uint32, lpcbSecurityDescriptor *uint32, lpftLastWriteTime *FILETIME) LSTATUS {
	addr := lazyAddr(&pRegQueryInfoKeyW, libAdvapi32, "RegQueryInfoKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpClass)), uintptr(unsafe.Pointer(lpcchClass)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpcSubKeys)), uintptr(unsafe.Pointer(lpcbMaxSubKeyLen)), uintptr(unsafe.Pointer(lpcbMaxClassLen)), uintptr(unsafe.Pointer(lpcValues)), uintptr(unsafe.Pointer(lpcbMaxValueNameLen)), uintptr(unsafe.Pointer(lpcbMaxValueLen)), uintptr(unsafe.Pointer(lpcbSecurityDescriptor)), uintptr(unsafe.Pointer(lpftLastWriteTime)))
	return LSTATUS(ret)
}

func RegQueryValueA(hKey HKEY, lpSubKey PSTR, lpData PSTR, lpcbData *int32) LSTATUS {
	addr := lazyAddr(&pRegQueryValueA, libAdvapi32, "RegQueryValueA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpData)), uintptr(unsafe.Pointer(lpcbData)))
	return LSTATUS(ret)
}

var RegQueryValue = RegQueryValueW
func RegQueryValueW(hKey HKEY, lpSubKey PWSTR, lpData PWSTR, lpcbData *int32) LSTATUS {
	addr := lazyAddr(&pRegQueryValueW, libAdvapi32, "RegQueryValueW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpData)), uintptr(unsafe.Pointer(lpcbData)))
	return LSTATUS(ret)
}

func RegQueryMultipleValuesA(hKey HKEY, val_list *VALENTA, num_vals uint32, lpValueBuf PSTR, ldwTotsize *uint32) LSTATUS {
	addr := lazyAddr(&pRegQueryMultipleValuesA, libAdvapi32, "RegQueryMultipleValuesA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(val_list)), uintptr(num_vals), uintptr(unsafe.Pointer(lpValueBuf)), uintptr(unsafe.Pointer(ldwTotsize)))
	return LSTATUS(ret)
}

var RegQueryMultipleValues = RegQueryMultipleValuesW
func RegQueryMultipleValuesW(hKey HKEY, val_list *VALENTW, num_vals uint32, lpValueBuf PWSTR, ldwTotsize *uint32) LSTATUS {
	addr := lazyAddr(&pRegQueryMultipleValuesW, libAdvapi32, "RegQueryMultipleValuesW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(val_list)), uintptr(num_vals), uintptr(unsafe.Pointer(lpValueBuf)), uintptr(unsafe.Pointer(ldwTotsize)))
	return LSTATUS(ret)
}

func RegQueryValueExA(hKey HKEY, lpValueName PSTR, lpReserved *uint32, lpType *REG_VALUE_TYPE, lpData *uint8, lpcbData *uint32) LSTATUS {
	addr := lazyAddr(&pRegQueryValueExA, libAdvapi32, "RegQueryValueExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpValueName)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpData)), uintptr(unsafe.Pointer(lpcbData)))
	return LSTATUS(ret)
}

var RegQueryValueEx = RegQueryValueExW
func RegQueryValueExW(hKey HKEY, lpValueName PWSTR, lpReserved *uint32, lpType *REG_VALUE_TYPE, lpData *uint8, lpcbData *uint32) LSTATUS {
	addr := lazyAddr(&pRegQueryValueExW, libAdvapi32, "RegQueryValueExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpValueName)), uintptr(unsafe.Pointer(lpReserved)), uintptr(unsafe.Pointer(lpType)), uintptr(unsafe.Pointer(lpData)), uintptr(unsafe.Pointer(lpcbData)))
	return LSTATUS(ret)
}

func RegReplaceKeyA(hKey HKEY, lpSubKey PSTR, lpNewFile PSTR, lpOldFile PSTR) LSTATUS {
	addr := lazyAddr(&pRegReplaceKeyA, libAdvapi32, "RegReplaceKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpNewFile)), uintptr(unsafe.Pointer(lpOldFile)))
	return LSTATUS(ret)
}

var RegReplaceKey = RegReplaceKeyW
func RegReplaceKeyW(hKey HKEY, lpSubKey PWSTR, lpNewFile PWSTR, lpOldFile PWSTR) LSTATUS {
	addr := lazyAddr(&pRegReplaceKeyW, libAdvapi32, "RegReplaceKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpNewFile)), uintptr(unsafe.Pointer(lpOldFile)))
	return LSTATUS(ret)
}

func RegRestoreKeyA(hKey HKEY, lpFile PSTR, dwFlags REG_RESTORE_KEY_FLAGS) LSTATUS {
	addr := lazyAddr(&pRegRestoreKeyA, libAdvapi32, "RegRestoreKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpFile)), uintptr(dwFlags))
	return LSTATUS(ret)
}

var RegRestoreKey = RegRestoreKeyW
func RegRestoreKeyW(hKey HKEY, lpFile PWSTR, dwFlags REG_RESTORE_KEY_FLAGS) LSTATUS {
	addr := lazyAddr(&pRegRestoreKeyW, libAdvapi32, "RegRestoreKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpFile)), uintptr(dwFlags))
	return LSTATUS(ret)
}

func RegRenameKey(hKey HKEY, lpSubKeyName PWSTR, lpNewKeyName PWSTR) LSTATUS {
	addr := lazyAddr(&pRegRenameKey, libAdvapi32, "RegRenameKey")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKeyName)), uintptr(unsafe.Pointer(lpNewKeyName)))
	return LSTATUS(ret)
}

func RegSaveKeyA(hKey HKEY, lpFile PSTR, lpSecurityAttributes *SECURITY_ATTRIBUTES) LSTATUS {
	addr := lazyAddr(&pRegSaveKeyA, libAdvapi32, "RegSaveKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(lpSecurityAttributes)))
	return LSTATUS(ret)
}

var RegSaveKey = RegSaveKeyW
func RegSaveKeyW(hKey HKEY, lpFile PWSTR, lpSecurityAttributes *SECURITY_ATTRIBUTES) LSTATUS {
	addr := lazyAddr(&pRegSaveKeyW, libAdvapi32, "RegSaveKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(lpSecurityAttributes)))
	return LSTATUS(ret)
}

func RegSetKeySecurity(hKey HKEY, SecurityInformation uint32, pSecurityDescriptor *SECURITY_DESCRIPTOR) LSTATUS {
	addr := lazyAddr(&pRegSetKeySecurity, libAdvapi32, "RegSetKeySecurity")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(SecurityInformation), uintptr(unsafe.Pointer(pSecurityDescriptor)))
	return LSTATUS(ret)
}

func RegSetValueA(hKey HKEY, lpSubKey PSTR, dwType REG_VALUE_TYPE, lpData PSTR, cbData uint32) LSTATUS {
	addr := lazyAddr(&pRegSetValueA, libAdvapi32, "RegSetValueA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(dwType), uintptr(unsafe.Pointer(lpData)), uintptr(cbData))
	return LSTATUS(ret)
}

var RegSetValue = RegSetValueW
func RegSetValueW(hKey HKEY, lpSubKey PWSTR, dwType REG_VALUE_TYPE, lpData PWSTR, cbData uint32) LSTATUS {
	addr := lazyAddr(&pRegSetValueW, libAdvapi32, "RegSetValueW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(dwType), uintptr(unsafe.Pointer(lpData)), uintptr(cbData))
	return LSTATUS(ret)
}

func RegSetValueExA(hKey HKEY, lpValueName PSTR, Reserved uint32, dwType REG_VALUE_TYPE, lpData *uint8, cbData uint32) LSTATUS {
	addr := lazyAddr(&pRegSetValueExA, libAdvapi32, "RegSetValueExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpValueName)), uintptr(Reserved), uintptr(dwType), uintptr(unsafe.Pointer(lpData)), uintptr(cbData))
	return LSTATUS(ret)
}

var RegSetValueEx = RegSetValueExW
func RegSetValueExW(hKey HKEY, lpValueName PWSTR, Reserved uint32, dwType REG_VALUE_TYPE, lpData *uint8, cbData uint32) LSTATUS {
	addr := lazyAddr(&pRegSetValueExW, libAdvapi32, "RegSetValueExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpValueName)), uintptr(Reserved), uintptr(dwType), uintptr(unsafe.Pointer(lpData)), uintptr(cbData))
	return LSTATUS(ret)
}

func RegUnLoadKeyA(hKey HKEY, lpSubKey PSTR) LSTATUS {
	addr := lazyAddr(&pRegUnLoadKeyA, libAdvapi32, "RegUnLoadKeyA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)))
	return LSTATUS(ret)
}

var RegUnLoadKey = RegUnLoadKeyW
func RegUnLoadKeyW(hKey HKEY, lpSubKey PWSTR) LSTATUS {
	addr := lazyAddr(&pRegUnLoadKeyW, libAdvapi32, "RegUnLoadKeyW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)))
	return LSTATUS(ret)
}

func RegDeleteKeyValueA(hKey HKEY, lpSubKey PSTR, lpValueName PSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyValueA, libAdvapi32, "RegDeleteKeyValueA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpValueName)))
	return LSTATUS(ret)
}

var RegDeleteKeyValue = RegDeleteKeyValueW
func RegDeleteKeyValueW(hKey HKEY, lpSubKey PWSTR, lpValueName PWSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteKeyValueW, libAdvapi32, "RegDeleteKeyValueW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpValueName)))
	return LSTATUS(ret)
}

func RegSetKeyValueA(hKey HKEY, lpSubKey PSTR, lpValueName PSTR, dwType uint32, lpData unsafe.Pointer, cbData uint32) LSTATUS {
	addr := lazyAddr(&pRegSetKeyValueA, libAdvapi32, "RegSetKeyValueA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpValueName)), uintptr(dwType), uintptr(lpData), uintptr(cbData))
	return LSTATUS(ret)
}

var RegSetKeyValue = RegSetKeyValueW
func RegSetKeyValueW(hKey HKEY, lpSubKey PWSTR, lpValueName PWSTR, dwType uint32, lpData unsafe.Pointer, cbData uint32) LSTATUS {
	addr := lazyAddr(&pRegSetKeyValueW, libAdvapi32, "RegSetKeyValueW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpValueName)), uintptr(dwType), uintptr(lpData), uintptr(cbData))
	return LSTATUS(ret)
}

func RegDeleteTreeA(hKey HKEY, lpSubKey PSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteTreeA, libAdvapi32, "RegDeleteTreeA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)))
	return LSTATUS(ret)
}

var RegDeleteTree = RegDeleteTreeW
func RegDeleteTreeW(hKey HKEY, lpSubKey PWSTR) LSTATUS {
	addr := lazyAddr(&pRegDeleteTreeW, libAdvapi32, "RegDeleteTreeW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpSubKey)))
	return LSTATUS(ret)
}

func RegCopyTreeA(hKeySrc HKEY, lpSubKey PSTR, hKeyDest HKEY) LSTATUS {
	addr := lazyAddr(&pRegCopyTreeA, libAdvapi32, "RegCopyTreeA")
	ret, _,  _ := syscall.SyscallN(addr, hKeySrc, uintptr(unsafe.Pointer(lpSubKey)), hKeyDest)
	return LSTATUS(ret)
}

func RegGetValueA(hkey HKEY, lpSubKey PSTR, lpValue PSTR, dwFlags RRF_RT, pdwType *uint32, pvData unsafe.Pointer, pcbData *uint32) LSTATUS {
	addr := lazyAddr(&pRegGetValueA, libAdvapi32, "RegGetValueA")
	ret, _,  _ := syscall.SyscallN(addr, hkey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpValue)), uintptr(dwFlags), uintptr(unsafe.Pointer(pdwType)), uintptr(pvData), uintptr(unsafe.Pointer(pcbData)))
	return LSTATUS(ret)
}

var RegGetValue = RegGetValueW
func RegGetValueW(hkey HKEY, lpSubKey PWSTR, lpValue PWSTR, dwFlags RRF_RT, pdwType *uint32, pvData unsafe.Pointer, pcbData *uint32) LSTATUS {
	addr := lazyAddr(&pRegGetValueW, libAdvapi32, "RegGetValueW")
	ret, _,  _ := syscall.SyscallN(addr, hkey, uintptr(unsafe.Pointer(lpSubKey)), uintptr(unsafe.Pointer(lpValue)), uintptr(dwFlags), uintptr(unsafe.Pointer(pdwType)), uintptr(pvData), uintptr(unsafe.Pointer(pcbData)))
	return LSTATUS(ret)
}

var RegCopyTree = RegCopyTreeW
func RegCopyTreeW(hKeySrc HKEY, lpSubKey PWSTR, hKeyDest HKEY) LSTATUS {
	addr := lazyAddr(&pRegCopyTreeW, libAdvapi32, "RegCopyTreeW")
	ret, _,  _ := syscall.SyscallN(addr, hKeySrc, uintptr(unsafe.Pointer(lpSubKey)), hKeyDest)
	return LSTATUS(ret)
}

func RegLoadMUIStringA(hKey HKEY, pszValue PSTR, pszOutBuf PSTR, cbOutBuf uint32, pcbData *uint32, Flags uint32, pszDirectory PSTR) LSTATUS {
	addr := lazyAddr(&pRegLoadMUIStringA, libAdvapi32, "RegLoadMUIStringA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(pszValue)), uintptr(unsafe.Pointer(pszOutBuf)), uintptr(cbOutBuf), uintptr(unsafe.Pointer(pcbData)), uintptr(Flags), uintptr(unsafe.Pointer(pszDirectory)))
	return LSTATUS(ret)
}

var RegLoadMUIString = RegLoadMUIStringW
func RegLoadMUIStringW(hKey HKEY, pszValue PWSTR, pszOutBuf PWSTR, cbOutBuf uint32, pcbData *uint32, Flags uint32, pszDirectory PWSTR) LSTATUS {
	addr := lazyAddr(&pRegLoadMUIStringW, libAdvapi32, "RegLoadMUIStringW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(pszValue)), uintptr(unsafe.Pointer(pszOutBuf)), uintptr(cbOutBuf), uintptr(unsafe.Pointer(pcbData)), uintptr(Flags), uintptr(unsafe.Pointer(pszDirectory)))
	return LSTATUS(ret)
}

func RegLoadAppKeyA(lpFile PSTR, phkResult *HKEY, samDesired uint32, dwOptions uint32, Reserved uint32) LSTATUS {
	addr := lazyAddr(&pRegLoadAppKeyA, libAdvapi32, "RegLoadAppKeyA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(phkResult)), uintptr(samDesired), uintptr(dwOptions), uintptr(Reserved))
	return LSTATUS(ret)
}

var RegLoadAppKey = RegLoadAppKeyW
func RegLoadAppKeyW(lpFile PWSTR, phkResult *HKEY, samDesired uint32, dwOptions uint32, Reserved uint32) LSTATUS {
	addr := lazyAddr(&pRegLoadAppKeyW, libAdvapi32, "RegLoadAppKeyW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(phkResult)), uintptr(samDesired), uintptr(dwOptions), uintptr(Reserved))
	return LSTATUS(ret)
}

func RegSaveKeyExA(hKey HKEY, lpFile PSTR, lpSecurityAttributes *SECURITY_ATTRIBUTES, Flags REG_SAVE_FORMAT) LSTATUS {
	addr := lazyAddr(&pRegSaveKeyExA, libAdvapi32, "RegSaveKeyExA")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(Flags))
	return LSTATUS(ret)
}

var RegSaveKeyEx = RegSaveKeyExW
func RegSaveKeyExW(hKey HKEY, lpFile PWSTR, lpSecurityAttributes *SECURITY_ATTRIBUTES, Flags REG_SAVE_FORMAT) LSTATUS {
	addr := lazyAddr(&pRegSaveKeyExW, libAdvapi32, "RegSaveKeyExW")
	ret, _,  _ := syscall.SyscallN(addr, hKey, uintptr(unsafe.Pointer(lpFile)), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(Flags))
	return LSTATUS(ret)
}


