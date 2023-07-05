package win32

import (
	"syscall"
	"unsafe"
)

const (
	WOW64_SIZE_OF_80387_REGISTERS                                      uint32 = 0x50
	WOW64_MAXIMUM_SUPPORTED_EXTENSION                                  uint32 = 0x200
	RESTORE_LAST_ERROR_NAME_A                                          string = "RestoreLastError"
	RESTORE_LAST_ERROR_NAME_W                                          string = "RestoreLastError"
	RESTORE_LAST_ERROR_NAME                                            string = "RestoreLastError"
	MAX_SYM_NAME                                                       uint32 = 0x7d0
	BIND_NO_BOUND_IMPORTS                                              uint32 = 0x1
	BIND_NO_UPDATE                                                     uint32 = 0x2
	BIND_ALL_IMAGES                                                    uint32 = 0x4
	BIND_CACHE_IMPORT_DLLS                                             uint32 = 0x8
	BIND_REPORT_64BIT_VA                                               uint32 = 0x10
	CHECKSUM_SUCCESS                                                   uint32 = 0x0
	CHECKSUM_OPEN_FAILURE                                              uint32 = 0x1
	CHECKSUM_MAP_FAILURE                                               uint32 = 0x2
	CHECKSUM_MAPVIEW_FAILURE                                           uint32 = 0x3
	CHECKSUM_UNICODE_FAILURE                                           uint32 = 0x4
	SPLITSYM_REMOVE_PRIVATE                                            uint32 = 0x1
	SPLITSYM_EXTRACT_ALL                                               uint32 = 0x2
	SPLITSYM_SYMBOLPATH_IS_SRC                                         uint32 = 0x4
	CERT_PE_IMAGE_DIGEST_DEBUG_INFO                                    uint32 = 0x1
	CERT_PE_IMAGE_DIGEST_RESOURCES                                     uint32 = 0x2
	CERT_PE_IMAGE_DIGEST_ALL_IMPORT_INFO                               uint32 = 0x4
	CERT_PE_IMAGE_DIGEST_NON_PE_INFO                                   uint32 = 0x8
	CERT_SECTION_TYPE_ANY                                              uint32 = 0xff
	ERROR_IMAGE_NOT_STRIPPED                                           uint32 = 0x8800
	ERROR_NO_DBG_POINTER                                               uint32 = 0x8801
	ERROR_NO_PDB_POINTER                                               uint32 = 0x8802
	UNDNAME_COMPLETE                                                   uint32 = 0x0
	UNDNAME_NO_LEADING_UNDERSCORES                                     uint32 = 0x1
	UNDNAME_NO_MS_KEYWORDS                                             uint32 = 0x2
	UNDNAME_NO_FUNCTION_RETURNS                                        uint32 = 0x4
	UNDNAME_NO_ALLOCATION_MODEL                                        uint32 = 0x8
	UNDNAME_NO_ALLOCATION_LANGUAGE                                     uint32 = 0x10
	UNDNAME_NO_MS_THISTYPE                                             uint32 = 0x20
	UNDNAME_NO_CV_THISTYPE                                             uint32 = 0x40
	UNDNAME_NO_THISTYPE                                                uint32 = 0x60
	UNDNAME_NO_ACCESS_SPECIFIERS                                       uint32 = 0x80
	UNDNAME_NO_THROW_SIGNATURES                                        uint32 = 0x100
	UNDNAME_NO_MEMBER_TYPE                                             uint32 = 0x200
	UNDNAME_NO_RETURN_UDT_MODEL                                        uint32 = 0x400
	UNDNAME_32_BIT_DECODE                                              uint32 = 0x800
	UNDNAME_NAME_ONLY                                                  uint32 = 0x1000
	UNDNAME_NO_ARGUMENTS                                               uint32 = 0x2000
	UNDNAME_NO_SPECIAL_SYMS                                            uint32 = 0x4000
	DBHHEADER_PDBGUID                                                  uint32 = 0x3
	INLINE_FRAME_CONTEXT_INIT                                          uint32 = 0x0
	INLINE_FRAME_CONTEXT_IGNORE                                        uint32 = 0xffffffff
	TARGET_ATTRIBUTE_PACMASK                                           uint32 = 0x1
	SYM_STKWALK_DEFAULT                                                uint32 = 0x0
	SYM_STKWALK_FORCE_FRAMEPTR                                         uint32 = 0x1
	SYM_STKWALK_ZEROEXTEND_PTRS                                        uint32 = 0x2
	API_VERSION_NUMBER                                                 uint32 = 0xc
	SYMFLAG_NULL                                                       uint32 = 0x80000
	SYMFLAG_FUNC_NO_RETURN                                             uint32 = 0x100000
	SYMFLAG_SYNTHETIC_ZEROBASE                                         uint32 = 0x200000
	SYMFLAG_PUBLIC_CODE                                                uint32 = 0x400000
	SYMFLAG_REGREL_ALIASINDIR                                          uint32 = 0x800000
	SYMFLAG_FIXUP_ARM64X                                               uint32 = 0x1000000
	SYMFLAG_GLOBAL                                                     uint32 = 0x2000000
	SYMFLAG_RESET                                                      uint32 = 0x80000000
	IMAGEHLP_MODULE_REGION_DLLBASE                                     uint32 = 0x1
	IMAGEHLP_MODULE_REGION_DLLRANGE                                    uint32 = 0x2
	IMAGEHLP_MODULE_REGION_ADDITIONAL                                  uint32 = 0x4
	IMAGEHLP_MODULE_REGION_JIT                                         uint32 = 0x8
	IMAGEHLP_MODULE_REGION_ALL                                         uint32 = 0xff
	CBA_DEFERRED_SYMBOL_LOAD_START                                     uint32 = 0x1
	CBA_DEFERRED_SYMBOL_LOAD_COMPLETE                                  uint32 = 0x2
	CBA_DEFERRED_SYMBOL_LOAD_FAILURE                                   uint32 = 0x3
	CBA_SYMBOLS_UNLOADED                                               uint32 = 0x4
	CBA_DUPLICATE_SYMBOL                                               uint32 = 0x5
	CBA_READ_MEMORY                                                    uint32 = 0x6
	CBA_DEFERRED_SYMBOL_LOAD_CANCEL                                    uint32 = 0x7
	CBA_SET_OPTIONS                                                    uint32 = 0x8
	CBA_EVENT                                                          uint32 = 0x10
	CBA_DEFERRED_SYMBOL_LOAD_PARTIAL                                   uint32 = 0x20
	CBA_DEBUG_INFO                                                     uint32 = 0x10000000
	CBA_SRCSRV_INFO                                                    uint32 = 0x20000000
	CBA_SRCSRV_EVENT                                                   uint32 = 0x40000000
	CBA_UPDATE_STATUS_BAR                                              uint32 = 0x50000000
	CBA_ENGINE_PRESENT                                                 uint32 = 0x60000000
	CBA_CHECK_ENGOPT_DISALLOW_NETWORK_PATHS                            uint32 = 0x70000000
	CBA_CHECK_ARM_MACHINE_THUMB_TYPE_OVERRIDE                          uint32 = 0x80000000
	CBA_XML_LOG                                                        uint32 = 0x90000000
	CBA_MAP_JIT_SYMBOL                                                 uint32 = 0xa0000000
	EVENT_SRCSPEW_START                                                uint32 = 0x64
	EVENT_SRCSPEW                                                      uint32 = 0x64
	EVENT_SRCSPEW_END                                                  uint32 = 0xc7
	DSLFLAG_MISMATCHED_PDB                                             uint32 = 0x1
	DSLFLAG_MISMATCHED_DBG                                             uint32 = 0x2
	FLAG_ENGINE_PRESENT                                                uint32 = 0x4
	FLAG_ENGOPT_DISALLOW_NETWORK_PATHS                                 uint32 = 0x8
	FLAG_OVERRIDE_ARM_MACHINE_TYPE                                     uint32 = 0x10
	SYMOPT_CASE_INSENSITIVE                                            uint32 = 0x1
	SYMOPT_UNDNAME                                                     uint32 = 0x2
	SYMOPT_DEFERRED_LOADS                                              uint32 = 0x4
	SYMOPT_NO_CPP                                                      uint32 = 0x8
	SYMOPT_LOAD_LINES                                                  uint32 = 0x10
	SYMOPT_OMAP_FIND_NEAREST                                           uint32 = 0x20
	SYMOPT_LOAD_ANYTHING                                               uint32 = 0x40
	SYMOPT_IGNORE_CVREC                                                uint32 = 0x80
	SYMOPT_NO_UNQUALIFIED_LOADS                                        uint32 = 0x100
	SYMOPT_FAIL_CRITICAL_ERRORS                                        uint32 = 0x200
	SYMOPT_EXACT_SYMBOLS                                               uint32 = 0x400
	SYMOPT_ALLOW_ABSOLUTE_SYMBOLS                                      uint32 = 0x800
	SYMOPT_IGNORE_NT_SYMPATH                                           uint32 = 0x1000
	SYMOPT_INCLUDE_32BIT_MODULES                                       uint32 = 0x2000
	SYMOPT_PUBLICS_ONLY                                                uint32 = 0x4000
	SYMOPT_NO_PUBLICS                                                  uint32 = 0x8000
	SYMOPT_AUTO_PUBLICS                                                uint32 = 0x10000
	SYMOPT_NO_IMAGE_SEARCH                                             uint32 = 0x20000
	SYMOPT_SECURE                                                      uint32 = 0x40000
	SYMOPT_NO_PROMPTS                                                  uint32 = 0x80000
	SYMOPT_OVERWRITE                                                   uint32 = 0x100000
	SYMOPT_IGNORE_IMAGEDIR                                             uint32 = 0x200000
	SYMOPT_FLAT_DIRECTORY                                              uint32 = 0x400000
	SYMOPT_FAVOR_COMPRESSED                                            uint32 = 0x800000
	SYMOPT_ALLOW_ZERO_ADDRESS                                          uint32 = 0x1000000
	SYMOPT_DISABLE_SYMSRV_AUTODETECT                                   uint32 = 0x2000000
	SYMOPT_READONLY_CACHE                                              uint32 = 0x4000000
	SYMOPT_SYMPATH_LAST                                                uint32 = 0x8000000
	SYMOPT_DISABLE_FAST_SYMBOLS                                        uint32 = 0x10000000
	SYMOPT_DISABLE_SYMSRV_TIMEOUT                                      uint32 = 0x20000000
	SYMOPT_DISABLE_SRVSTAR_ON_STARTUP                                  uint32 = 0x40000000
	SYMOPT_DEBUG                                                       uint32 = 0x80000000
	SYM_INLINE_COMP_ERROR                                              uint32 = 0x0
	SYM_INLINE_COMP_IDENTICAL                                          uint32 = 0x1
	SYM_INLINE_COMP_STEPIN                                             uint32 = 0x2
	SYM_INLINE_COMP_STEPOUT                                            uint32 = 0x3
	SYM_INLINE_COMP_STEPOVER                                           uint32 = 0x4
	SYM_INLINE_COMP_DIFFERENT                                          uint32 = 0x5
	ESLFLAG_FULLPATH                                                   uint32 = 0x1
	ESLFLAG_NEAREST                                                    uint32 = 0x2
	ESLFLAG_PREV                                                       uint32 = 0x4
	ESLFLAG_NEXT                                                       uint32 = 0x8
	ESLFLAG_INLINE_SITE                                                uint32 = 0x10
	SYMENUM_OPTIONS_DEFAULT                                            uint32 = 0x1
	SYMENUM_OPTIONS_INLINE                                             uint32 = 0x2
	SYMSEARCH_MASKOBJS                                                 uint32 = 0x1
	SYMSEARCH_RECURSE                                                  uint32 = 0x2
	SYMSEARCH_GLOBALSONLY                                              uint32 = 0x4
	SYMSEARCH_ALLITEMS                                                 uint32 = 0x8
	EXT_OUTPUT_VER                                                     uint32 = 0x1
	SYMSRV_VERSION                                                     uint32 = 0x2
	SSRVOPT_CALLBACK                                                   uint32 = 0x1
	SSRVOPT_OLDGUIDPTR                                                 uint32 = 0x10
	SSRVOPT_UNATTENDED                                                 uint32 = 0x20
	SSRVOPT_NOCOPY                                                     uint32 = 0x40
	SSRVOPT_GETPATH                                                    uint32 = 0x40
	SSRVOPT_PARENTWIN                                                  uint32 = 0x80
	SSRVOPT_PARAMTYPE                                                  uint32 = 0x100
	SSRVOPT_SECURE                                                     uint32 = 0x200
	SSRVOPT_TRACE                                                      uint32 = 0x400
	SSRVOPT_SETCONTEXT                                                 uint32 = 0x800
	SSRVOPT_PROXY                                                      uint32 = 0x1000
	SSRVOPT_DOWNSTREAM_STORE                                           uint32 = 0x2000
	SSRVOPT_OVERWRITE                                                  uint32 = 0x4000
	SSRVOPT_RESETTOU                                                   uint32 = 0x8000
	SSRVOPT_CALLBACKW                                                  uint32 = 0x10000
	SSRVOPT_FLAT_DEFAULT_STORE                                         uint32 = 0x20000
	SSRVOPT_PROXYW                                                     uint32 = 0x40000
	SSRVOPT_MESSAGE                                                    uint32 = 0x80000
	SSRVOPT_SERVICE                                                    uint32 = 0x100000
	SSRVOPT_FAVOR_COMPRESSED                                           uint32 = 0x200000
	SSRVOPT_STRING                                                     uint32 = 0x400000
	SSRVOPT_WINHTTP                                                    uint32 = 0x800000
	SSRVOPT_WININET                                                    uint32 = 0x1000000
	SSRVOPT_DONT_UNCOMPRESS                                            uint32 = 0x2000000
	SSRVOPT_DISABLE_PING_HOST                                          uint32 = 0x4000000
	SSRVOPT_DISABLE_TIMEOUT                                            uint32 = 0x8000000
	SSRVOPT_ENABLE_COMM_MSG                                            uint32 = 0x10000000
	SSRVOPT_URI_FILTER                                                 uint32 = 0x20000000
	SSRVOPT_URI_TIERS                                                  uint32 = 0x40000000
	SSRVOPT_RETRY_APP_HANG                                             uint32 = 0x80000000
	SSRVOPT_MAX                                                        uint32 = 0x80000000
	NUM_SSRVOPTS                                                       uint32 = 0x20
	SSRVURI_HTTP_NORMAL                                                uint32 = 0x1
	SSRVURI_HTTP_COMPRESSED                                            uint32 = 0x2
	SSRVURI_HTTP_FILEPTR                                               uint32 = 0x4
	SSRVURI_UNC_NORMAL                                                 uint32 = 0x10
	SSRVURI_UNC_COMPRESSED                                             uint32 = 0x20
	SSRVURI_UNC_FILEPTR                                                uint32 = 0x40
	SSRVURI_HTTP_MASK                                                  uint32 = 0xf
	SSRVURI_UNC_MASK                                                   uint32 = 0xf0
	SSRVURI_ALL                                                        uint32 = 0xff
	SSRVURI_NORMAL                                                     uint32 = 0x1
	SSRVURI_COMPRESSED                                                 uint32 = 0x2
	SSRVURI_FILEPTR                                                    uint32 = 0x4
	SSRVACTION_TRACE                                                   uint32 = 0x1
	SSRVACTION_QUERYCANCEL                                             uint32 = 0x2
	SSRVACTION_EVENT                                                   uint32 = 0x3
	SSRVACTION_EVENTW                                                  uint32 = 0x4
	SSRVACTION_SIZE                                                    uint32 = 0x5
	SSRVACTION_HTTPSTATUS                                              uint32 = 0x6
	SSRVACTION_XMLOUTPUT                                               uint32 = 0x7
	SSRVACTION_CHECKSUMSTATUS                                          uint32 = 0x8
	SYMSTOREOPT_ALT_INDEX                                              uint32 = 0x10
	SYMSTOREOPT_UNICODE                                                uint32 = 0x20
	SYMF_OMAP_GENERATED                                                uint32 = 0x1
	SYMF_OMAP_MODIFIED                                                 uint32 = 0x2
	SYMF_REGISTER                                                      uint32 = 0x8
	SYMF_REGREL                                                        uint32 = 0x10
	SYMF_FRAMEREL                                                      uint32 = 0x20
	SYMF_PARAMETER                                                     uint32 = 0x40
	SYMF_LOCAL                                                         uint32 = 0x80
	SYMF_CONSTANT                                                      uint32 = 0x100
	SYMF_EXPORT                                                        uint32 = 0x200
	SYMF_FORWARDER                                                     uint32 = 0x400
	SYMF_FUNCTION                                                      uint32 = 0x800
	SYMF_VIRTUAL                                                       uint32 = 0x1000
	SYMF_THUNK                                                         uint32 = 0x2000
	SYMF_TLSREL                                                        uint32 = 0x4000
	IMAGEHLP_SYMBOL_INFO_VALUEPRESENT                                  uint32 = 0x1
	IMAGEHLP_SYMBOL_INFO_REGISTER                                      uint32 = 0x8
	IMAGEHLP_SYMBOL_INFO_REGRELATIVE                                   uint32 = 0x10
	IMAGEHLP_SYMBOL_INFO_FRAMERELATIVE                                 uint32 = 0x20
	IMAGEHLP_SYMBOL_INFO_PARAMETER                                     uint32 = 0x40
	IMAGEHLP_SYMBOL_INFO_LOCAL                                         uint32 = 0x80
	IMAGEHLP_SYMBOL_INFO_CONSTANT                                      uint32 = 0x100
	IMAGEHLP_SYMBOL_FUNCTION                                           uint32 = 0x800
	IMAGEHLP_SYMBOL_VIRTUAL                                            uint32 = 0x1000
	IMAGEHLP_SYMBOL_THUNK                                              uint32 = 0x2000
	IMAGEHLP_SYMBOL_INFO_TLSRELATIVE                                   uint32 = 0x4000
	IMAGEHLP_RMAP_MAPPED_FLAT                                          uint32 = 0x1
	IMAGEHLP_RMAP_BIG_ENDIAN                                           uint32 = 0x2
	IMAGEHLP_RMAP_IGNORE_MISCOMPARE                                    uint32 = 0x4
	IMAGEHLP_RMAP_FIXUP_ARM64X                                         uint32 = 0x10000000
	IMAGEHLP_RMAP_LOAD_RW_DATA_SECTIONS                                uint32 = 0x20000000
	IMAGEHLP_RMAP_OMIT_SHARED_RW_DATA_SECTIONS                         uint32 = 0x40000000
	IMAGEHLP_RMAP_FIXUP_IMAGEBASE                                      uint32 = 0x80000000
	DMP_PHYSICAL_MEMORY_BLOCK_SIZE_32                                  uint32 = 0x2bc
	DMP_CONTEXT_RECORD_SIZE_32                                         uint32 = 0x4b0
	DMP_RESERVED_0_SIZE_32                                             uint32 = 0x6e0
	DMP_RESERVED_2_SIZE_32                                             uint32 = 0x10
	DMP_RESERVED_3_SIZE_32                                             uint32 = 0x38
	DMP_PHYSICAL_MEMORY_BLOCK_SIZE_64                                  uint32 = 0x2bc
	DMP_CONTEXT_RECORD_SIZE_64                                         uint32 = 0xbb8
	DMP_RESERVED_0_SIZE_64                                             uint32 = 0xfa8
	DMP_HEADER_COMMENT_SIZE                                            uint32 = 0x80
	DUMP_SUMMARY_VALID_KERNEL_VA                                       uint32 = 0x1
	DUMP_SUMMARY_VALID_CURRENT_USER_VA                                 uint32 = 0x2
	MINIDUMP_VERSION                                                   uint32 = 0xa793
	MINIDUMP_MISC1_PROCESSOR_POWER_INFO                                uint32 = 0x4
	MINIDUMP_MISC3_PROCESS_INTEGRITY                                   uint32 = 0x10
	MINIDUMP_MISC3_PROCESS_EXECUTE_FLAGS                               uint32 = 0x20
	MINIDUMP_MISC3_TIMEZONE                                            uint32 = 0x40
	MINIDUMP_MISC3_PROTECTED_PROCESS                                   uint32 = 0x80
	MINIDUMP_MISC4_BUILDSTRING                                         uint32 = 0x100
	MINIDUMP_MISC5_PROCESS_COOKIE                                      uint32 = 0x200
	MINIDUMP_SYSMEMINFO1_FILECACHE_TRANSITIONREPURPOSECOUNT_FLAGS      uint32 = 0x1
	MINIDUMP_SYSMEMINFO1_BASICPERF                                     uint32 = 0x2
	MINIDUMP_SYSMEMINFO1_PERF_CCTOTALDIRTYPAGES_CCDIRTYPAGETHRESHOLD   uint32 = 0x4
	MINIDUMP_SYSMEMINFO1_PERF_RESIDENTAVAILABLEPAGES_SHAREDCOMMITPAGES uint32 = 0x8
	MINIDUMP_PROCESS_VM_COUNTERS                                       uint32 = 0x1
	MINIDUMP_PROCESS_VM_COUNTERS_VIRTUALSIZE                           uint32 = 0x2
	MINIDUMP_PROCESS_VM_COUNTERS_EX                                    uint32 = 0x4
	MINIDUMP_PROCESS_VM_COUNTERS_EX2                                   uint32 = 0x8
	MINIDUMP_PROCESS_VM_COUNTERS_JOB                                   uint32 = 0x10
	INTERFACESAFE_FOR_UNTRUSTED_CALLER                                 uint32 = 0x1
	INTERFACESAFE_FOR_UNTRUSTED_DATA                                   uint32 = 0x2
	INTERFACE_USES_DISPEX                                              uint32 = 0x4
	INTERFACE_USES_SECURITY_MANAGER                                    uint32 = 0x8
	WCT_MAX_NODE_COUNT                                                 uint32 = 0x10
	WCT_OBJNAME_LENGTH                                                 uint32 = 0x80
	WCT_NETWORK_IO_FLAG                                                uint32 = 0x8
	WHEA_ERROR_SOURCE_DESCRIPTOR_VERSION_10                            uint32 = 0xa
	WHEA_ERROR_SOURCE_DESCRIPTOR_VERSION_11                            uint32 = 0xb
	WHEA_MAX_MC_BANKS                                                  uint32 = 0x20
	WHEA_ERROR_SOURCE_FLAG_FIRMWAREFIRST                               uint32 = 0x1
	WHEA_ERROR_SOURCE_FLAG_GLOBAL                                      uint32 = 0x2
	WHEA_ERROR_SOURCE_FLAG_GHES_ASSIST                                 uint32 = 0x4
	WHEA_ERROR_SOURCE_FLAG_DEFAULTSOURCE                               uint32 = 0x80000000
	WHEA_ERROR_SOURCE_INVALID_RELATED_SOURCE                           uint32 = 0xffff
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_XPFMCE                           uint32 = 0x0
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_XPFCMC                           uint32 = 0x1
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_XPFNMI                           uint32 = 0x2
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_IPFMCA                           uint32 = 0x3
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_IPFCMC                           uint32 = 0x4
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_IPFCPE                           uint32 = 0x5
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_AERROOTPORT                      uint32 = 0x6
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_AERENDPOINT                      uint32 = 0x7
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_AERBRIDGE                        uint32 = 0x8
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_GENERIC                          uint32 = 0x9
	WHEA_ERROR_SOURCE_DESCRIPTOR_TYPE_GENERIC_V2                       uint32 = 0xa
	WHEA_XPF_MC_BANK_STATUSFORMAT_IA32MCA                              uint32 = 0x0
	WHEA_XPF_MC_BANK_STATUSFORMAT_Intel64MCA                           uint32 = 0x1
	WHEA_XPF_MC_BANK_STATUSFORMAT_AMD64MCA                             uint32 = 0x2
	WHEA_NOTIFICATION_TYPE_POLLED                                      uint32 = 0x0
	WHEA_NOTIFICATION_TYPE_EXTERNALINTERRUPT                           uint32 = 0x1
	WHEA_NOTIFICATION_TYPE_LOCALINTERRUPT                              uint32 = 0x2
	WHEA_NOTIFICATION_TYPE_SCI                                         uint32 = 0x3
	WHEA_NOTIFICATION_TYPE_NMI                                         uint32 = 0x4
	WHEA_NOTIFICATION_TYPE_CMCI                                        uint32 = 0x5
	WHEA_NOTIFICATION_TYPE_MCE                                         uint32 = 0x6
	WHEA_NOTIFICATION_TYPE_GPIO_SIGNAL                                 uint32 = 0x7
	WHEA_NOTIFICATION_TYPE_ARMV8_SEA                                   uint32 = 0x8
	WHEA_NOTIFICATION_TYPE_ARMV8_SEI                                   uint32 = 0x9
	WHEA_NOTIFICATION_TYPE_EXTERNALINTERRUPT_GSIV                      uint32 = 0xa
	WHEA_NOTIFICATION_TYPE_SDEI                                        uint32 = 0xb
	WHEA_DEVICE_DRIVER_CONFIG_V1                                       uint32 = 0x1
	WHEA_DEVICE_DRIVER_CONFIG_V2                                       uint32 = 0x2
	WHEA_DEVICE_DRIVER_CONFIG_MIN                                      uint32 = 0x1
	WHEA_DEVICE_DRIVER_CONFIG_MAX                                      uint32 = 0x2
	WHEA_DEVICE_DRIVER_BUFFER_SET_V1                                   uint32 = 0x1
	WHEA_DEVICE_DRIVER_BUFFER_SET_MIN                                  uint32 = 0x1
	WHEA_DEVICE_DRIVER_BUFFER_SET_MAX                                  uint32 = 0x1
	WHEA_DISABLE_OFFLINE                                               uint32 = 0x0
	WHEA_MEM_PERSISTOFFLINE                                            uint32 = 0x1
	WHEA_MEM_PFA_DISABLE                                               uint32 = 0x2
	WHEA_MEM_PFA_PAGECOUNT                                             uint32 = 0x3
	WHEA_MEM_PFA_THRESHOLD                                             uint32 = 0x4
	WHEA_MEM_PFA_TIMEOUT                                               uint32 = 0x5
	WHEA_DISABLE_DUMMY_WRITE                                           uint32 = 0x6
	WHEA_RESTORE_CMCI_ENABLED                                          uint32 = 0x7
	WHEA_RESTORE_CMCI_ATTEMPTS                                         uint32 = 0x8
	WHEA_RESTORE_CMCI_ERR_LIMIT                                        uint32 = 0x9
	WHEA_CMCI_THRESHOLD_COUNT                                          uint32 = 0xa
	WHEA_CMCI_THRESHOLD_TIME                                           uint32 = 0xb
	WHEA_CMCI_THRESHOLD_POLL_COUNT                                     uint32 = 0xc
	WHEA_PENDING_PAGE_LIST_SZ                                          uint32 = 0xd
	WHEA_BAD_PAGE_LIST_MAX_SIZE                                        uint32 = 0xe
	WHEA_BAD_PAGE_LIST_LOCATION                                        uint32 = 0xf
	WHEA_NOTIFY_ALL_OFFLINES                                           uint32 = 0x10
	WHEA_ROW_FAIL_CHECK_EXTENT                                         uint32 = 0x11
	WHEA_ROW_FAIL_CHECK_ENABLE                                         uint32 = 0x12
	WHEA_ROW_FAIL_CHECK_THRESHOLD                                      uint32 = 0x13
	IPMI_OS_SEL_RECORD_VERSION_1                                       uint32 = 0x1
	IPMI_OS_SEL_RECORD_VERSION                                         uint32 = 0x1
	IPMI_IOCTL_INDEX                                                   uint32 = 0x400
	IOCTL_IPMI_INTERNAL_RECORD_SEL_EVENT                               uint32 = 0x221000
	IPMI_OS_SEL_RECORD_MASK                                            uint32 = 0xffff
	SevMax                                                             int32  = 4
)

// enums

// enum
// flags
type SYM_LOAD_FLAGS uint32

const (
	SLMFLAG_NONE       SYM_LOAD_FLAGS = 0
	SLMFLAG_VIRTUAL    SYM_LOAD_FLAGS = 1
	SLMFLAG_ALT_INDEX  SYM_LOAD_FLAGS = 2
	SLMFLAG_NO_SYMBOLS SYM_LOAD_FLAGS = 4
)

// enum
// flags
type IMAGE_SECTION_CHARACTERISTICS uint32

const (
	IMAGE_SCN_TYPE_NO_PAD            IMAGE_SECTION_CHARACTERISTICS = 8
	IMAGE_SCN_CNT_CODE               IMAGE_SECTION_CHARACTERISTICS = 32
	IMAGE_SCN_CNT_INITIALIZED_DATA   IMAGE_SECTION_CHARACTERISTICS = 64
	IMAGE_SCN_CNT_UNINITIALIZED_DATA IMAGE_SECTION_CHARACTERISTICS = 128
	IMAGE_SCN_LNK_OTHER              IMAGE_SECTION_CHARACTERISTICS = 256
	IMAGE_SCN_LNK_INFO               IMAGE_SECTION_CHARACTERISTICS = 512
	IMAGE_SCN_LNK_REMOVE             IMAGE_SECTION_CHARACTERISTICS = 2048
	IMAGE_SCN_LNK_COMDAT             IMAGE_SECTION_CHARACTERISTICS = 4096
	IMAGE_SCN_NO_DEFER_SPEC_EXC      IMAGE_SECTION_CHARACTERISTICS = 16384
	IMAGE_SCN_GPREL                  IMAGE_SECTION_CHARACTERISTICS = 32768
	IMAGE_SCN_MEM_FARDATA            IMAGE_SECTION_CHARACTERISTICS = 32768
	IMAGE_SCN_MEM_PURGEABLE          IMAGE_SECTION_CHARACTERISTICS = 131072
	IMAGE_SCN_MEM_16BIT              IMAGE_SECTION_CHARACTERISTICS = 131072
	IMAGE_SCN_MEM_LOCKED             IMAGE_SECTION_CHARACTERISTICS = 262144
	IMAGE_SCN_MEM_PRELOAD            IMAGE_SECTION_CHARACTERISTICS = 524288
	IMAGE_SCN_ALIGN_1BYTES           IMAGE_SECTION_CHARACTERISTICS = 1048576
	IMAGE_SCN_ALIGN_2BYTES           IMAGE_SECTION_CHARACTERISTICS = 2097152
	IMAGE_SCN_ALIGN_4BYTES           IMAGE_SECTION_CHARACTERISTICS = 3145728
	IMAGE_SCN_ALIGN_8BYTES           IMAGE_SECTION_CHARACTERISTICS = 4194304
	IMAGE_SCN_ALIGN_16BYTES          IMAGE_SECTION_CHARACTERISTICS = 5242880
	IMAGE_SCN_ALIGN_32BYTES          IMAGE_SECTION_CHARACTERISTICS = 6291456
	IMAGE_SCN_ALIGN_64BYTES          IMAGE_SECTION_CHARACTERISTICS = 7340032
	IMAGE_SCN_ALIGN_128BYTES         IMAGE_SECTION_CHARACTERISTICS = 8388608
	IMAGE_SCN_ALIGN_256BYTES         IMAGE_SECTION_CHARACTERISTICS = 9437184
	IMAGE_SCN_ALIGN_512BYTES         IMAGE_SECTION_CHARACTERISTICS = 10485760
	IMAGE_SCN_ALIGN_1024BYTES        IMAGE_SECTION_CHARACTERISTICS = 11534336
	IMAGE_SCN_ALIGN_2048BYTES        IMAGE_SECTION_CHARACTERISTICS = 12582912
	IMAGE_SCN_ALIGN_4096BYTES        IMAGE_SECTION_CHARACTERISTICS = 13631488
	IMAGE_SCN_ALIGN_8192BYTES        IMAGE_SECTION_CHARACTERISTICS = 14680064
	IMAGE_SCN_ALIGN_MASK             IMAGE_SECTION_CHARACTERISTICS = 15728640
	IMAGE_SCN_LNK_NRELOC_OVFL        IMAGE_SECTION_CHARACTERISTICS = 16777216
	IMAGE_SCN_MEM_DISCARDABLE        IMAGE_SECTION_CHARACTERISTICS = 33554432
	IMAGE_SCN_MEM_NOT_CACHED         IMAGE_SECTION_CHARACTERISTICS = 67108864
	IMAGE_SCN_MEM_NOT_PAGED          IMAGE_SECTION_CHARACTERISTICS = 134217728
	IMAGE_SCN_MEM_SHARED             IMAGE_SECTION_CHARACTERISTICS = 268435456
	IMAGE_SCN_MEM_EXECUTE            IMAGE_SECTION_CHARACTERISTICS = 536870912
	IMAGE_SCN_MEM_READ               IMAGE_SECTION_CHARACTERISTICS = 1073741824
	IMAGE_SCN_MEM_WRITE              IMAGE_SECTION_CHARACTERISTICS = 2147483648
	IMAGE_SCN_SCALE_INDEX            IMAGE_SECTION_CHARACTERISTICS = 1
)

// enum
type IMAGE_SUBSYSTEM uint16

const (
	IMAGE_SUBSYSTEM_UNKNOWN                  IMAGE_SUBSYSTEM = 0
	IMAGE_SUBSYSTEM_NATIVE                   IMAGE_SUBSYSTEM = 1
	IMAGE_SUBSYSTEM_WINDOWS_GUI              IMAGE_SUBSYSTEM = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI              IMAGE_SUBSYSTEM = 3
	IMAGE_SUBSYSTEM_OS2_CUI                  IMAGE_SUBSYSTEM = 5
	IMAGE_SUBSYSTEM_POSIX_CUI                IMAGE_SUBSYSTEM = 7
	IMAGE_SUBSYSTEM_NATIVE_WINDOWS           IMAGE_SUBSYSTEM = 8
	IMAGE_SUBSYSTEM_WINDOWS_CE_GUI           IMAGE_SUBSYSTEM = 9
	IMAGE_SUBSYSTEM_EFI_APPLICATION          IMAGE_SUBSYSTEM = 10
	IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER  IMAGE_SUBSYSTEM = 11
	IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER       IMAGE_SUBSYSTEM = 12
	IMAGE_SUBSYSTEM_EFI_ROM                  IMAGE_SUBSYSTEM = 13
	IMAGE_SUBSYSTEM_XBOX                     IMAGE_SUBSYSTEM = 14
	IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION IMAGE_SUBSYSTEM = 16
	IMAGE_SUBSYSTEM_XBOX_CODE_CATALOG        IMAGE_SUBSYSTEM = 17
)

// enum
// flags
type IMAGE_DLL_CHARACTERISTICS uint16

const (
	IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA                               IMAGE_DLL_CHARACTERISTICS = 32
	IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE                                  IMAGE_DLL_CHARACTERISTICS = 64
	IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY                               IMAGE_DLL_CHARACTERISTICS = 128
	IMAGE_DLLCHARACTERISTICS_NX_COMPAT                                     IMAGE_DLL_CHARACTERISTICS = 256
	IMAGE_DLLCHARACTERISTICS_NO_ISOLATION                                  IMAGE_DLL_CHARACTERISTICS = 512
	IMAGE_DLLCHARACTERISTICS_NO_SEH                                        IMAGE_DLL_CHARACTERISTICS = 1024
	IMAGE_DLLCHARACTERISTICS_NO_BIND                                       IMAGE_DLL_CHARACTERISTICS = 2048
	IMAGE_DLLCHARACTERISTICS_APPCONTAINER                                  IMAGE_DLL_CHARACTERISTICS = 4096
	IMAGE_DLLCHARACTERISTICS_WDM_DRIVER                                    IMAGE_DLL_CHARACTERISTICS = 8192
	IMAGE_DLLCHARACTERISTICS_GUARD_CF                                      IMAGE_DLL_CHARACTERISTICS = 16384
	IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE                         IMAGE_DLL_CHARACTERISTICS = 32768
	IMAGE_DLLCHARACTERISTICS_EX_CET_COMPAT                                 IMAGE_DLL_CHARACTERISTICS = 1
	IMAGE_DLLCHARACTERISTICS_EX_CET_COMPAT_STRICT_MODE                     IMAGE_DLL_CHARACTERISTICS = 2
	IMAGE_DLLCHARACTERISTICS_EX_CET_SET_CONTEXT_IP_VALIDATION_RELAXED_MODE IMAGE_DLL_CHARACTERISTICS = 4
	IMAGE_DLLCHARACTERISTICS_EX_CET_DYNAMIC_APIS_ALLOW_IN_PROC             IMAGE_DLL_CHARACTERISTICS = 8
	IMAGE_DLLCHARACTERISTICS_EX_CET_RESERVED_1                             IMAGE_DLL_CHARACTERISTICS = 16
	IMAGE_DLLCHARACTERISTICS_EX_CET_RESERVED_2                             IMAGE_DLL_CHARACTERISTICS = 32
)

// enum
type IMAGE_OPTIONAL_HEADER_MAGIC uint16

const (
	IMAGE_NT_OPTIONAL_HDR_MAGIC   IMAGE_OPTIONAL_HEADER_MAGIC = 523
	IMAGE_NT_OPTIONAL_HDR32_MAGIC IMAGE_OPTIONAL_HEADER_MAGIC = 267
	IMAGE_NT_OPTIONAL_HDR64_MAGIC IMAGE_OPTIONAL_HEADER_MAGIC = 523
	IMAGE_ROM_OPTIONAL_HDR_MAGIC  IMAGE_OPTIONAL_HEADER_MAGIC = 263
)

// enum
type BUGCHECK_ERROR uint32

const (
	HARDWARE_PROFILE_UNDOCKED_STRING                         BUGCHECK_ERROR = 1073807361
	HARDWARE_PROFILE_DOCKED_STRING                           BUGCHECK_ERROR = 1073807362
	HARDWARE_PROFILE_UNKNOWN_STRING                          BUGCHECK_ERROR = 1073807363
	WINDOWS_NT_BANNER                                        BUGCHECK_ERROR = 1073741950
	WINDOWS_NT_CSD_STRING                                    BUGCHECK_ERROR = 1073741959
	WINDOWS_NT_INFO_STRING                                   BUGCHECK_ERROR = 1073741960
	WINDOWS_NT_MP_STRING                                     BUGCHECK_ERROR = 1073741961
	THREAD_TERMINATE_HELD_MUTEX                              BUGCHECK_ERROR = 1073741962
	WINDOWS_NT_INFO_STRING_PLURAL                            BUGCHECK_ERROR = 1073741981
	WINDOWS_NT_RC_STRING                                     BUGCHECK_ERROR = 1073741982
	APC_INDEX_MISMATCH                                       BUGCHECK_ERROR = 1
	DEVICE_QUEUE_NOT_BUSY                                    BUGCHECK_ERROR = 2
	INVALID_AFFINITY_SET                                     BUGCHECK_ERROR = 3
	INVALID_DATA_ACCESS_TRAP                                 BUGCHECK_ERROR = 4
	INVALID_PROCESS_ATTACH_ATTEMPT                           BUGCHECK_ERROR = 5
	INVALID_PROCESS_DETACH_ATTEMPT                           BUGCHECK_ERROR = 6
	INVALID_SOFTWARE_INTERRUPT                               BUGCHECK_ERROR = 7
	IRQL_NOT_DISPATCH_LEVEL                                  BUGCHECK_ERROR = 8
	IRQL_NOT_GREATER_OR_EQUAL                                BUGCHECK_ERROR = 9
	IRQL_NOT_LESS_OR_EQUAL                                   BUGCHECK_ERROR = 10
	NO_EXCEPTION_HANDLING_SUPPORT                            BUGCHECK_ERROR = 11
	MAXIMUM_WAIT_OBJECTS_EXCEEDED                            BUGCHECK_ERROR = 12
	MUTEX_LEVEL_NUMBER_VIOLATION                             BUGCHECK_ERROR = 13
	NO_USER_MODE_CONTEXT                                     BUGCHECK_ERROR = 14
	SPIN_LOCK_ALREADY_OWNED                                  BUGCHECK_ERROR = 15
	SPIN_LOCK_NOT_OWNED                                      BUGCHECK_ERROR = 16
	THREAD_NOT_MUTEX_OWNER                                   BUGCHECK_ERROR = 17
	TRAP_CAUSE_UNKNOWN                                       BUGCHECK_ERROR = 18
	EMPTY_THREAD_REAPER_LIST                                 BUGCHECK_ERROR = 19
	CREATE_DELETE_LOCK_NOT_LOCKED                            BUGCHECK_ERROR = 20
	LAST_CHANCE_CALLED_FROM_KMODE                            BUGCHECK_ERROR = 21
	CID_HANDLE_CREATION                                      BUGCHECK_ERROR = 22
	CID_HANDLE_DELETION                                      BUGCHECK_ERROR = 23
	REFERENCE_BY_POINTER                                     BUGCHECK_ERROR = 24
	BAD_POOL_HEADER                                          BUGCHECK_ERROR = 25
	MEMORY_MANAGEMENT                                        BUGCHECK_ERROR = 26
	PFN_SHARE_COUNT                                          BUGCHECK_ERROR = 27
	PFN_REFERENCE_COUNT                                      BUGCHECK_ERROR = 28
	NO_SPIN_LOCK_AVAILABLE                                   BUGCHECK_ERROR = 29
	KMODE_EXCEPTION_NOT_HANDLED                              BUGCHECK_ERROR = 30
	SHARED_RESOURCE_CONV_ERROR                               BUGCHECK_ERROR = 31
	KERNEL_APC_PENDING_DURING_EXIT                           BUGCHECK_ERROR = 32
	QUOTA_UNDERFLOW                                          BUGCHECK_ERROR = 33
	FILE_SYSTEM                                              BUGCHECK_ERROR = 34
	FAT_FILE_SYSTEM                                          BUGCHECK_ERROR = 35
	NTFS_FILE_SYSTEM                                         BUGCHECK_ERROR = 36
	NPFS_FILE_SYSTEM                                         BUGCHECK_ERROR = 37
	CDFS_FILE_SYSTEM                                         BUGCHECK_ERROR = 38
	RDR_FILE_SYSTEM                                          BUGCHECK_ERROR = 39
	CORRUPT_ACCESS_TOKEN                                     BUGCHECK_ERROR = 40
	SECURITY_SYSTEM                                          BUGCHECK_ERROR = 41
	INCONSISTENT_IRP                                         BUGCHECK_ERROR = 42
	PANIC_STACK_SWITCH                                       BUGCHECK_ERROR = 43
	PORT_DRIVER_INTERNAL                                     BUGCHECK_ERROR = 44
	SCSI_DISK_DRIVER_INTERNAL                                BUGCHECK_ERROR = 45
	DATA_BUS_ERROR                                           BUGCHECK_ERROR = 46
	INSTRUCTION_BUS_ERROR                                    BUGCHECK_ERROR = 47
	SET_OF_INVALID_CONTEXT                                   BUGCHECK_ERROR = 48
	PHASE0_INITIALIZATION_FAILED                             BUGCHECK_ERROR = 49
	PHASE1_INITIALIZATION_FAILED                             BUGCHECK_ERROR = 50
	UNEXPECTED_INITIALIZATION_CALL                           BUGCHECK_ERROR = 51
	CACHE_MANAGER                                            BUGCHECK_ERROR = 52
	NO_MORE_IRP_STACK_LOCATIONS                              BUGCHECK_ERROR = 53
	DEVICE_REFERENCE_COUNT_NOT_ZERO                          BUGCHECK_ERROR = 54
	FLOPPY_INTERNAL_ERROR                                    BUGCHECK_ERROR = 55
	SERIAL_DRIVER_INTERNAL                                   BUGCHECK_ERROR = 56
	SYSTEM_EXIT_OWNED_MUTEX                                  BUGCHECK_ERROR = 57
	SYSTEM_UNWIND_PREVIOUS_USER                              BUGCHECK_ERROR = 58
	SYSTEM_SERVICE_EXCEPTION                                 BUGCHECK_ERROR = 59
	INTERRUPT_UNWIND_ATTEMPTED                               BUGCHECK_ERROR = 60
	INTERRUPT_EXCEPTION_NOT_HANDLED                          BUGCHECK_ERROR = 61
	MULTIPROCESSOR_CONFIGURATION_NOT_SUPPORTED               BUGCHECK_ERROR = 62
	NO_MORE_SYSTEM_PTES                                      BUGCHECK_ERROR = 63
	TARGET_MDL_TOO_SMALL                                     BUGCHECK_ERROR = 64
	MUST_SUCCEED_POOL_EMPTY                                  BUGCHECK_ERROR = 65
	ATDISK_DRIVER_INTERNAL                                   BUGCHECK_ERROR = 66
	NO_SUCH_PARTITION                                        BUGCHECK_ERROR = 67
	MULTIPLE_IRP_COMPLETE_REQUESTS                           BUGCHECK_ERROR = 68
	INSUFFICIENT_SYSTEM_MAP_REGS                             BUGCHECK_ERROR = 69
	DEREF_UNKNOWN_LOGON_SESSION                              BUGCHECK_ERROR = 70
	REF_UNKNOWN_LOGON_SESSION                                BUGCHECK_ERROR = 71
	CANCEL_STATE_IN_COMPLETED_IRP                            BUGCHECK_ERROR = 72
	PAGE_FAULT_WITH_INTERRUPTS_OFF                           BUGCHECK_ERROR = 73
	IRQL_GT_ZERO_AT_SYSTEM_SERVICE                           BUGCHECK_ERROR = 74
	STREAMS_INTERNAL_ERROR                                   BUGCHECK_ERROR = 75
	FATAL_UNHANDLED_HARD_ERROR                               BUGCHECK_ERROR = 76
	NO_PAGES_AVAILABLE                                       BUGCHECK_ERROR = 77
	PFN_LIST_CORRUPT                                         BUGCHECK_ERROR = 78
	NDIS_INTERNAL_ERROR                                      BUGCHECK_ERROR = 79
	PAGE_FAULT_IN_NONPAGED_AREA                              BUGCHECK_ERROR = 80
	PAGE_FAULT_IN_NONPAGED_AREA_M                            BUGCHECK_ERROR = 268435536
	REGISTRY_ERROR                                           BUGCHECK_ERROR = 81
	MAILSLOT_FILE_SYSTEM                                     BUGCHECK_ERROR = 82
	NO_BOOT_DEVICE                                           BUGCHECK_ERROR = 83
	LM_SERVER_INTERNAL_ERROR                                 BUGCHECK_ERROR = 84
	DATA_COHERENCY_EXCEPTION                                 BUGCHECK_ERROR = 85
	INSTRUCTION_COHERENCY_EXCEPTION                          BUGCHECK_ERROR = 86
	XNS_INTERNAL_ERROR                                       BUGCHECK_ERROR = 87
	VOLMGRX_INTERNAL_ERROR                                   BUGCHECK_ERROR = 88
	PINBALL_FILE_SYSTEM                                      BUGCHECK_ERROR = 89
	CRITICAL_SERVICE_FAILED                                  BUGCHECK_ERROR = 90
	SET_ENV_VAR_FAILED                                       BUGCHECK_ERROR = 91
	HAL_INITIALIZATION_FAILED                                BUGCHECK_ERROR = 92
	UNSUPPORTED_PROCESSOR                                    BUGCHECK_ERROR = 93
	OBJECT_INITIALIZATION_FAILED                             BUGCHECK_ERROR = 94
	SECURITY_INITIALIZATION_FAILED                           BUGCHECK_ERROR = 95
	PROCESS_INITIALIZATION_FAILED                            BUGCHECK_ERROR = 96
	HAL1_INITIALIZATION_FAILED                               BUGCHECK_ERROR = 97
	OBJECT1_INITIALIZATION_FAILED                            BUGCHECK_ERROR = 98
	SECURITY1_INITIALIZATION_FAILED                          BUGCHECK_ERROR = 99
	SYMBOLIC_INITIALIZATION_FAILED                           BUGCHECK_ERROR = 100
	MEMORY1_INITIALIZATION_FAILED                            BUGCHECK_ERROR = 101
	CACHE_INITIALIZATION_FAILED                              BUGCHECK_ERROR = 102
	CONFIG_INITIALIZATION_FAILED                             BUGCHECK_ERROR = 103
	FILE_INITIALIZATION_FAILED                               BUGCHECK_ERROR = 104
	IO1_INITIALIZATION_FAILED                                BUGCHECK_ERROR = 105
	LPC_INITIALIZATION_FAILED                                BUGCHECK_ERROR = 106
	PROCESS1_INITIALIZATION_FAILED                           BUGCHECK_ERROR = 107
	REFMON_INITIALIZATION_FAILED                             BUGCHECK_ERROR = 108
	SESSION1_INITIALIZATION_FAILED                           BUGCHECK_ERROR = 109
	BOOTPROC_INITIALIZATION_FAILED                           BUGCHECK_ERROR = 110
	VSL_INITIALIZATION_FAILED                                BUGCHECK_ERROR = 111
	SOFT_RESTART_FATAL_ERROR                                 BUGCHECK_ERROR = 112
	ASSIGN_DRIVE_LETTERS_FAILED                              BUGCHECK_ERROR = 114
	CONFIG_LIST_FAILED                                       BUGCHECK_ERROR = 115
	BAD_SYSTEM_CONFIG_INFO                                   BUGCHECK_ERROR = 116
	CANNOT_WRITE_CONFIGURATION                               BUGCHECK_ERROR = 117
	PROCESS_HAS_LOCKED_PAGES                                 BUGCHECK_ERROR = 118
	KERNEL_STACK_INPAGE_ERROR                                BUGCHECK_ERROR = 119
	PHASE0_EXCEPTION                                         BUGCHECK_ERROR = 120
	MISMATCHED_HAL                                           BUGCHECK_ERROR = 121
	KERNEL_DATA_INPAGE_ERROR                                 BUGCHECK_ERROR = 122
	INACCESSIBLE_BOOT_DEVICE                                 BUGCHECK_ERROR = 123
	BUGCODE_NDIS_DRIVER                                      BUGCHECK_ERROR = 124
	INSTALL_MORE_MEMORY                                      BUGCHECK_ERROR = 125
	SYSTEM_THREAD_EXCEPTION_NOT_HANDLED                      BUGCHECK_ERROR = 126
	SYSTEM_THREAD_EXCEPTION_NOT_HANDLED_M                    BUGCHECK_ERROR = 268435582
	UNEXPECTED_KERNEL_MODE_TRAP                              BUGCHECK_ERROR = 127
	UNEXPECTED_KERNEL_MODE_TRAP_M                            BUGCHECK_ERROR = 268435583
	NMI_HARDWARE_FAILURE                                     BUGCHECK_ERROR = 128
	SPIN_LOCK_INIT_FAILURE                                   BUGCHECK_ERROR = 129
	DFS_FILE_SYSTEM                                          BUGCHECK_ERROR = 130
	OFS_FILE_SYSTEM                                          BUGCHECK_ERROR = 131
	RECOM_DRIVER                                             BUGCHECK_ERROR = 132
	SETUP_FAILURE                                            BUGCHECK_ERROR = 133
	AUDIT_FAILURE                                            BUGCHECK_ERROR = 134
	MBR_CHECKSUM_MISMATCH                                    BUGCHECK_ERROR = 139
	KERNEL_MODE_EXCEPTION_NOT_HANDLED                        BUGCHECK_ERROR = 142
	KERNEL_MODE_EXCEPTION_NOT_HANDLED_M                      BUGCHECK_ERROR = 268435598
	PP0_INITIALIZATION_FAILED                                BUGCHECK_ERROR = 143
	PP1_INITIALIZATION_FAILED                                BUGCHECK_ERROR = 144
	WIN32K_INIT_OR_RIT_FAILURE                               BUGCHECK_ERROR = 145
	UP_DRIVER_ON_MP_SYSTEM                                   BUGCHECK_ERROR = 146
	INVALID_KERNEL_HANDLE                                    BUGCHECK_ERROR = 147
	KERNEL_STACK_LOCKED_AT_EXIT                              BUGCHECK_ERROR = 148
	PNP_INTERNAL_ERROR                                       BUGCHECK_ERROR = 149
	INVALID_WORK_QUEUE_ITEM                                  BUGCHECK_ERROR = 150
	BOUND_IMAGE_UNSUPPORTED                                  BUGCHECK_ERROR = 151
	END_OF_NT_EVALUATION_PERIOD                              BUGCHECK_ERROR = 152
	INVALID_REGION_OR_SEGMENT                                BUGCHECK_ERROR = 153
	SYSTEM_LICENSE_VIOLATION                                 BUGCHECK_ERROR = 154
	UDFS_FILE_SYSTEM                                         BUGCHECK_ERROR = 155
	MACHINE_CHECK_EXCEPTION                                  BUGCHECK_ERROR = 156
	USER_MODE_HEALTH_MONITOR                                 BUGCHECK_ERROR = 158
	DRIVER_POWER_STATE_FAILURE                               BUGCHECK_ERROR = 159
	INTERNAL_POWER_ERROR                                     BUGCHECK_ERROR = 160
	PCI_BUS_DRIVER_INTERNAL                                  BUGCHECK_ERROR = 161
	MEMORY_IMAGE_CORRUPT                                     BUGCHECK_ERROR = 162
	ACPI_DRIVER_INTERNAL                                     BUGCHECK_ERROR = 163
	CNSS_FILE_SYSTEM_FILTER                                  BUGCHECK_ERROR = 164
	ACPI_BIOS_ERROR                                          BUGCHECK_ERROR = 165
	FP_EMULATION_ERROR                                       BUGCHECK_ERROR = 166
	BAD_EXHANDLE                                             BUGCHECK_ERROR = 167
	BOOTING_IN_SAFEMODE_MINIMAL                              BUGCHECK_ERROR = 168
	BOOTING_IN_SAFEMODE_NETWORK                              BUGCHECK_ERROR = 169
	BOOTING_IN_SAFEMODE_DSREPAIR                             BUGCHECK_ERROR = 170
	SESSION_HAS_VALID_POOL_ON_EXIT                           BUGCHECK_ERROR = 171
	HAL_MEMORY_ALLOCATION                                    BUGCHECK_ERROR = 172
	VIDEO_DRIVER_DEBUG_REPORT_REQUEST                        BUGCHECK_ERROR = 1073741997
	BGI_DETECTED_VIOLATION                                   BUGCHECK_ERROR = 177
	VIDEO_DRIVER_INIT_FAILURE                                BUGCHECK_ERROR = 180
	BOOTLOG_LOADED                                           BUGCHECK_ERROR = 181
	BOOTLOG_NOT_LOADED                                       BUGCHECK_ERROR = 182
	BOOTLOG_ENABLED                                          BUGCHECK_ERROR = 183
	ATTEMPTED_SWITCH_FROM_DPC                                BUGCHECK_ERROR = 184
	CHIPSET_DETECTED_ERROR                                   BUGCHECK_ERROR = 185
	SESSION_HAS_VALID_VIEWS_ON_EXIT                          BUGCHECK_ERROR = 186
	NETWORK_BOOT_INITIALIZATION_FAILED                       BUGCHECK_ERROR = 187
	NETWORK_BOOT_DUPLICATE_ADDRESS                           BUGCHECK_ERROR = 188
	INVALID_HIBERNATED_STATE                                 BUGCHECK_ERROR = 189
	ATTEMPTED_WRITE_TO_READONLY_MEMORY                       BUGCHECK_ERROR = 190
	MUTEX_ALREADY_OWNED                                      BUGCHECK_ERROR = 191
	PCI_CONFIG_SPACE_ACCESS_FAILURE                          BUGCHECK_ERROR = 192
	SPECIAL_POOL_DETECTED_MEMORY_CORRUPTION                  BUGCHECK_ERROR = 193
	BAD_POOL_CALLER                                          BUGCHECK_ERROR = 194
	SYSTEM_IMAGE_BAD_SIGNATURE                               BUGCHECK_ERROR = 195
	DRIVER_VERIFIER_DETECTED_VIOLATION                       BUGCHECK_ERROR = 196
	DRIVER_CORRUPTED_EXPOOL                                  BUGCHECK_ERROR = 197
	DRIVER_CAUGHT_MODIFYING_FREED_POOL                       BUGCHECK_ERROR = 198
	TIMER_OR_DPC_INVALID                                     BUGCHECK_ERROR = 199
	IRQL_UNEXPECTED_VALUE                                    BUGCHECK_ERROR = 200
	DRIVER_VERIFIER_IOMANAGER_VIOLATION                      BUGCHECK_ERROR = 201
	PNP_DETECTED_FATAL_ERROR                                 BUGCHECK_ERROR = 202
	DRIVER_LEFT_LOCKED_PAGES_IN_PROCESS                      BUGCHECK_ERROR = 203
	PAGE_FAULT_IN_FREED_SPECIAL_POOL                         BUGCHECK_ERROR = 204
	PAGE_FAULT_BEYOND_END_OF_ALLOCATION                      BUGCHECK_ERROR = 205
	DRIVER_UNLOADED_WITHOUT_CANCELLING_PENDING_OPERATIONS    BUGCHECK_ERROR = 206
	TERMINAL_SERVER_DRIVER_MADE_INCORRECT_MEMORY_REFERENCE   BUGCHECK_ERROR = 207
	DRIVER_CORRUPTED_MMPOOL                                  BUGCHECK_ERROR = 208
	DRIVER_IRQL_NOT_LESS_OR_EQUAL                            BUGCHECK_ERROR = 209
	BUGCODE_ID_DRIVER                                        BUGCHECK_ERROR = 210
	DRIVER_PORTION_MUST_BE_NONPAGED                          BUGCHECK_ERROR = 211
	SYSTEM_SCAN_AT_RAISED_IRQL_CAUGHT_IMPROPER_DRIVER_UNLOAD BUGCHECK_ERROR = 212
	DRIVER_PAGE_FAULT_IN_FREED_SPECIAL_POOL                  BUGCHECK_ERROR = 213
	DRIVER_PAGE_FAULT_BEYOND_END_OF_ALLOCATION               BUGCHECK_ERROR = 214
	DRIVER_PAGE_FAULT_BEYOND_END_OF_ALLOCATION_M             BUGCHECK_ERROR = 268435670
	DRIVER_UNMAPPING_INVALID_VIEW                            BUGCHECK_ERROR = 215
	DRIVER_USED_EXCESSIVE_PTES                               BUGCHECK_ERROR = 216
	LOCKED_PAGES_TRACKER_CORRUPTION                          BUGCHECK_ERROR = 217
	SYSTEM_PTE_MISUSE                                        BUGCHECK_ERROR = 218
	DRIVER_CORRUPTED_SYSPTES                                 BUGCHECK_ERROR = 219
	DRIVER_INVALID_STACK_ACCESS                              BUGCHECK_ERROR = 220
	POOL_CORRUPTION_IN_FILE_AREA                             BUGCHECK_ERROR = 222
	IMPERSONATING_WORKER_THREAD                              BUGCHECK_ERROR = 223
	ACPI_BIOS_FATAL_ERROR                                    BUGCHECK_ERROR = 224
	WORKER_THREAD_RETURNED_AT_BAD_IRQL                       BUGCHECK_ERROR = 225
	MANUALLY_INITIATED_CRASH                                 BUGCHECK_ERROR = 226
	RESOURCE_NOT_OWNED                                       BUGCHECK_ERROR = 227
	WORKER_INVALID                                           BUGCHECK_ERROR = 228
	POWER_FAILURE_SIMULATE                                   BUGCHECK_ERROR = 229
	DRIVER_VERIFIER_DMA_VIOLATION                            BUGCHECK_ERROR = 230
	INVALID_FLOATING_POINT_STATE                             BUGCHECK_ERROR = 231
	INVALID_CANCEL_OF_FILE_OPEN                              BUGCHECK_ERROR = 232
	ACTIVE_EX_WORKER_THREAD_TERMINATION                      BUGCHECK_ERROR = 233
	SAVER_UNSPECIFIED                                        BUGCHECK_ERROR = 61440
	SAVER_BLANKSCREEN                                        BUGCHECK_ERROR = 61442
	SAVER_INPUT                                              BUGCHECK_ERROR = 61443
	SAVER_WATCHDOG                                           BUGCHECK_ERROR = 61444
	SAVER_STARTNOTVISIBLE                                    BUGCHECK_ERROR = 61445
	SAVER_NAVIGATIONMODEL                                    BUGCHECK_ERROR = 61446
	SAVER_OUTOFMEMORY                                        BUGCHECK_ERROR = 61447
	SAVER_GRAPHICS                                           BUGCHECK_ERROR = 61448
	SAVER_NAVSERVERTIMEOUT                                   BUGCHECK_ERROR = 61449
	SAVER_CHROMEPROCESSCRASH                                 BUGCHECK_ERROR = 61450
	SAVER_NOTIFICATIONDISMISSAL                              BUGCHECK_ERROR = 61451
	SAVER_SPEECHDISMISSAL                                    BUGCHECK_ERROR = 61452
	SAVER_CALLDISMISSAL                                      BUGCHECK_ERROR = 61453
	SAVER_APPBARDISMISSAL                                    BUGCHECK_ERROR = 61454
	SAVER_RILADAPTATIONCRASH                                 BUGCHECK_ERROR = 61455
	SAVER_APPLISTUNREACHABLE                                 BUGCHECK_ERROR = 61456
	SAVER_REPORTNOTIFICATIONFAILURE                          BUGCHECK_ERROR = 61457
	SAVER_UNEXPECTEDSHUTDOWN                                 BUGCHECK_ERROR = 61458
	SAVER_RPCFAILURE                                         BUGCHECK_ERROR = 61459
	SAVER_AUXILIARYFULLDUMP                                  BUGCHECK_ERROR = 61460
	SAVER_ACCOUNTPROVSVCINITFAILURE                          BUGCHECK_ERROR = 61461
	SAVER_MTBFCOMMANDTIMEOUT                                 BUGCHECK_ERROR = 789
	SAVER_MTBFCOMMANDHANG                                    BUGCHECK_ERROR = 61697
	SAVER_MTBFPASSBUGCHECK                                   BUGCHECK_ERROR = 61698
	SAVER_MTBFIOERROR                                        BUGCHECK_ERROR = 61699
	SAVER_RENDERTHREADHANG                                   BUGCHECK_ERROR = 61952
	SAVER_RENDERMOBILEUIOOM                                  BUGCHECK_ERROR = 61953
	SAVER_DEVICEUPDATEUNSPECIFIED                            BUGCHECK_ERROR = 62208
	SAVER_AUDIODRIVERHANG                                    BUGCHECK_ERROR = 62464
	SAVER_BATTERYPULLOUT                                     BUGCHECK_ERROR = 62720
	SAVER_MEDIACORETESTHANG                                  BUGCHECK_ERROR = 62976
	SAVER_RESOURCEMANAGEMENT                                 BUGCHECK_ERROR = 63232
	SAVER_CAPTURESERVICE                                     BUGCHECK_ERROR = 63488
	SAVER_WAITFORSHELLREADY                                  BUGCHECK_ERROR = 63744
	SAVER_NONRESPONSIVEPROCESS                               BUGCHECK_ERROR = 404
	SAVER_SICKAPPLICATION                                    BUGCHECK_ERROR = 34918
	THREAD_STUCK_IN_DEVICE_DRIVER                            BUGCHECK_ERROR = 234
	THREAD_STUCK_IN_DEVICE_DRIVER_M                          BUGCHECK_ERROR = 268435690
	DIRTY_MAPPED_PAGES_CONGESTION                            BUGCHECK_ERROR = 235
	SESSION_HAS_VALID_SPECIAL_POOL_ON_EXIT                   BUGCHECK_ERROR = 236
	UNMOUNTABLE_BOOT_VOLUME                                  BUGCHECK_ERROR = 237
	CRITICAL_PROCESS_DIED                                    BUGCHECK_ERROR = 239
	STORAGE_MINIPORT_ERROR                                   BUGCHECK_ERROR = 240
	SCSI_VERIFIER_DETECTED_VIOLATION                         BUGCHECK_ERROR = 241
	HARDWARE_INTERRUPT_STORM                                 BUGCHECK_ERROR = 242
	DISORDERLY_SHUTDOWN                                      BUGCHECK_ERROR = 243
	CRITICAL_OBJECT_TERMINATION                              BUGCHECK_ERROR = 244
	FLTMGR_FILE_SYSTEM                                       BUGCHECK_ERROR = 245
	PCI_VERIFIER_DETECTED_VIOLATION                          BUGCHECK_ERROR = 246
	DRIVER_OVERRAN_STACK_BUFFER                              BUGCHECK_ERROR = 247
	RAMDISK_BOOT_INITIALIZATION_FAILED                       BUGCHECK_ERROR = 248
	DRIVER_RETURNED_STATUS_REPARSE_FOR_VOLUME_OPEN           BUGCHECK_ERROR = 249
	HTTP_DRIVER_CORRUPTED                                    BUGCHECK_ERROR = 250
	RECURSIVE_MACHINE_CHECK                                  BUGCHECK_ERROR = 251
	ATTEMPTED_EXECUTE_OF_NOEXECUTE_MEMORY                    BUGCHECK_ERROR = 252
	DIRTY_NOWRITE_PAGES_CONGESTION                           BUGCHECK_ERROR = 253
	BUGCODE_USB_DRIVER                                       BUGCHECK_ERROR = 254
	BC_BLUETOOTH_VERIFIER_FAULT                              BUGCHECK_ERROR = 3070
	BC_BTHMINI_VERIFIER_FAULT                                BUGCHECK_ERROR = 3071
	RESERVE_QUEUE_OVERFLOW                                   BUGCHECK_ERROR = 255
	LOADER_BLOCK_MISMATCH                                    BUGCHECK_ERROR = 256
	CLOCK_WATCHDOG_TIMEOUT                                   BUGCHECK_ERROR = 257
	DPC_WATCHDOG_TIMEOUT                                     BUGCHECK_ERROR = 258
	MUP_FILE_SYSTEM                                          BUGCHECK_ERROR = 259
	AGP_INVALID_ACCESS                                       BUGCHECK_ERROR = 260
	AGP_GART_CORRUPTION                                      BUGCHECK_ERROR = 261
	AGP_ILLEGALLY_REPROGRAMMED                               BUGCHECK_ERROR = 262
	KERNEL_EXPAND_STACK_ACTIVE                               BUGCHECK_ERROR = 263
	THIRD_PARTY_FILE_SYSTEM_FAILURE                          BUGCHECK_ERROR = 264
	CRITICAL_STRUCTURE_CORRUPTION                            BUGCHECK_ERROR = 265
	APP_TAGGING_INITIALIZATION_FAILED                        BUGCHECK_ERROR = 266
	DFSC_FILE_SYSTEM                                         BUGCHECK_ERROR = 267
	FSRTL_EXTRA_CREATE_PARAMETER_VIOLATION                   BUGCHECK_ERROR = 268
	WDF_VIOLATION                                            BUGCHECK_ERROR = 269
	VIDEO_MEMORY_MANAGEMENT_INTERNAL                         BUGCHECK_ERROR = 270
	DRIVER_INVALID_CRUNTIME_PARAMETER                        BUGCHECK_ERROR = 272
	RECURSIVE_NMI                                            BUGCHECK_ERROR = 273
	MSRPC_STATE_VIOLATION                                    BUGCHECK_ERROR = 274
	VIDEO_DXGKRNL_FATAL_ERROR                                BUGCHECK_ERROR = 275
	VIDEO_SHADOW_DRIVER_FATAL_ERROR                          BUGCHECK_ERROR = 276
	AGP_INTERNAL                                             BUGCHECK_ERROR = 277
	VIDEO_TDR_FAILURE                                        BUGCHECK_ERROR = 278
	VIDEO_TDR_TIMEOUT_DETECTED                               BUGCHECK_ERROR = 279
	NTHV_GUEST_ERROR                                         BUGCHECK_ERROR = 280
	VIDEO_SCHEDULER_INTERNAL_ERROR                           BUGCHECK_ERROR = 281
	EM_INITIALIZATION_ERROR                                  BUGCHECK_ERROR = 282
	DRIVER_RETURNED_HOLDING_CANCEL_LOCK                      BUGCHECK_ERROR = 283
	ATTEMPTED_WRITE_TO_CM_PROTECTED_STORAGE                  BUGCHECK_ERROR = 284
	EVENT_TRACING_FATAL_ERROR                                BUGCHECK_ERROR = 285
	TOO_MANY_RECURSIVE_FAULTS                                BUGCHECK_ERROR = 286
	INVALID_DRIVER_HANDLE                                    BUGCHECK_ERROR = 287
	BITLOCKER_FATAL_ERROR                                    BUGCHECK_ERROR = 288
	DRIVER_VIOLATION                                         BUGCHECK_ERROR = 289
	WHEA_INTERNAL_ERROR                                      BUGCHECK_ERROR = 290
	CRYPTO_SELF_TEST_FAILURE                                 BUGCHECK_ERROR = 291
	WHEA_UNCORRECTABLE_ERROR                                 BUGCHECK_ERROR = 292
	NMR_INVALID_STATE                                        BUGCHECK_ERROR = 293
	NETIO_INVALID_POOL_CALLER                                BUGCHECK_ERROR = 294
	PAGE_NOT_ZERO                                            BUGCHECK_ERROR = 295
	WORKER_THREAD_RETURNED_WITH_BAD_IO_PRIORITY              BUGCHECK_ERROR = 296
	WORKER_THREAD_RETURNED_WITH_BAD_PAGING_IO_PRIORITY       BUGCHECK_ERROR = 297
	MUI_NO_VALID_SYSTEM_LANGUAGE                             BUGCHECK_ERROR = 298
	FAULTY_HARDWARE_CORRUPTED_PAGE                           BUGCHECK_ERROR = 299
	EXFAT_FILE_SYSTEM                                        BUGCHECK_ERROR = 300
	VOLSNAP_OVERLAPPED_TABLE_ACCESS                          BUGCHECK_ERROR = 301
	INVALID_MDL_RANGE                                        BUGCHECK_ERROR = 302
	VHD_BOOT_INITIALIZATION_FAILED                           BUGCHECK_ERROR = 303
	DYNAMIC_ADD_PROCESSOR_MISMATCH                           BUGCHECK_ERROR = 304
	INVALID_EXTENDED_PROCESSOR_STATE                         BUGCHECK_ERROR = 305
	RESOURCE_OWNER_POINTER_INVALID                           BUGCHECK_ERROR = 306
	DPC_WATCHDOG_VIOLATION                                   BUGCHECK_ERROR = 307
	DRIVE_EXTENDER                                           BUGCHECK_ERROR = 308
	REGISTRY_FILTER_DRIVER_EXCEPTION                         BUGCHECK_ERROR = 309
	VHD_BOOT_HOST_VOLUME_NOT_ENOUGH_SPACE                    BUGCHECK_ERROR = 310
	WIN32K_HANDLE_MANAGER                                    BUGCHECK_ERROR = 311
	GPIO_CONTROLLER_DRIVER_ERROR                             BUGCHECK_ERROR = 312
	KERNEL_SECURITY_CHECK_FAILURE                            BUGCHECK_ERROR = 313
	KERNEL_MODE_HEAP_CORRUPTION                              BUGCHECK_ERROR = 314
	PASSIVE_INTERRUPT_ERROR                                  BUGCHECK_ERROR = 315
	INVALID_IO_BOOST_STATE                                   BUGCHECK_ERROR = 316
	CRITICAL_INITIALIZATION_FAILURE                          BUGCHECK_ERROR = 317
	ERRATA_WORKAROUND_UNSUCCESSFUL                           BUGCHECK_ERROR = 318
	REGISTRY_CALLBACK_DRIVER_EXCEPTION                       BUGCHECK_ERROR = 319
	STORAGE_DEVICE_ABNORMALITY_DETECTED                      BUGCHECK_ERROR = 320
	VIDEO_ENGINE_TIMEOUT_DETECTED                            BUGCHECK_ERROR = 321
	VIDEO_TDR_APPLICATION_BLOCKED                            BUGCHECK_ERROR = 322
	PROCESSOR_DRIVER_INTERNAL                                BUGCHECK_ERROR = 323
	BUGCODE_USB3_DRIVER                                      BUGCHECK_ERROR = 324
	SECURE_BOOT_VIOLATION                                    BUGCHECK_ERROR = 325
	NDIS_NET_BUFFER_LIST_INFO_ILLEGALLY_TRANSFERRED          BUGCHECK_ERROR = 326
	ABNORMAL_RESET_DETECTED                                  BUGCHECK_ERROR = 327
	IO_OBJECT_INVALID                                        BUGCHECK_ERROR = 328
	REFS_FILE_SYSTEM                                         BUGCHECK_ERROR = 329
	KERNEL_WMI_INTERNAL                                      BUGCHECK_ERROR = 330
	SOC_SUBSYSTEM_FAILURE                                    BUGCHECK_ERROR = 331
	FATAL_ABNORMAL_RESET_ERROR                               BUGCHECK_ERROR = 332
	EXCEPTION_SCOPE_INVALID                                  BUGCHECK_ERROR = 333
	SOC_CRITICAL_DEVICE_REMOVED                              BUGCHECK_ERROR = 334
	PDC_WATCHDOG_TIMEOUT                                     BUGCHECK_ERROR = 335
	TCPIP_AOAC_NIC_ACTIVE_REFERENCE_LEAK                     BUGCHECK_ERROR = 336
	UNSUPPORTED_INSTRUCTION_MODE                             BUGCHECK_ERROR = 337
	INVALID_PUSH_LOCK_FLAGS                                  BUGCHECK_ERROR = 338
	KERNEL_LOCK_ENTRY_LEAKED_ON_THREAD_TERMINATION           BUGCHECK_ERROR = 339
	UNEXPECTED_STORE_EXCEPTION                               BUGCHECK_ERROR = 340
	OS_DATA_TAMPERING                                        BUGCHECK_ERROR = 341
	WINSOCK_DETECTED_HUNG_CLOSESOCKET_LIVEDUMP               BUGCHECK_ERROR = 342
	KERNEL_THREAD_PRIORITY_FLOOR_VIOLATION                   BUGCHECK_ERROR = 343
	ILLEGAL_IOMMU_PAGE_FAULT                                 BUGCHECK_ERROR = 344
	HAL_ILLEGAL_IOMMU_PAGE_FAULT                             BUGCHECK_ERROR = 345
	SDBUS_INTERNAL_ERROR                                     BUGCHECK_ERROR = 346
	WORKER_THREAD_RETURNED_WITH_SYSTEM_PAGE_PRIORITY_ACTIVE  BUGCHECK_ERROR = 347
	PDC_WATCHDOG_TIMEOUT_LIVEDUMP                            BUGCHECK_ERROR = 348
	SOC_SUBSYSTEM_FAILURE_LIVEDUMP                           BUGCHECK_ERROR = 349
	BUGCODE_NDIS_DRIVER_LIVE_DUMP                            BUGCHECK_ERROR = 350
	CONNECTED_STANDBY_WATCHDOG_TIMEOUT_LIVEDUMP              BUGCHECK_ERROR = 351
	WIN32K_ATOMIC_CHECK_FAILURE                              BUGCHECK_ERROR = 352
	LIVE_SYSTEM_DUMP                                         BUGCHECK_ERROR = 353
	KERNEL_AUTO_BOOST_INVALID_LOCK_RELEASE                   BUGCHECK_ERROR = 354
	WORKER_THREAD_TEST_CONDITION                             BUGCHECK_ERROR = 355
	WIN32K_CRITICAL_FAILURE                                  BUGCHECK_ERROR = 356
	CLUSTER_CSV_STATUS_IO_TIMEOUT_LIVEDUMP                   BUGCHECK_ERROR = 357
	CLUSTER_RESOURCE_CALL_TIMEOUT_LIVEDUMP                   BUGCHECK_ERROR = 358
	CLUSTER_CSV_SNAPSHOT_DEVICE_INFO_TIMEOUT_LIVEDUMP        BUGCHECK_ERROR = 359
	CLUSTER_CSV_STATE_TRANSITION_TIMEOUT_LIVEDUMP            BUGCHECK_ERROR = 360
	CLUSTER_CSV_VOLUME_ARRIVAL_LIVEDUMP                      BUGCHECK_ERROR = 361
	CLUSTER_CSV_VOLUME_REMOVAL_LIVEDUMP                      BUGCHECK_ERROR = 362
	CLUSTER_CSV_CLUSTER_WATCHDOG_LIVEDUMP                    BUGCHECK_ERROR = 363
	INVALID_RUNDOWN_PROTECTION_FLAGS                         BUGCHECK_ERROR = 364
	INVALID_SLOT_ALLOCATOR_FLAGS                             BUGCHECK_ERROR = 365
	ERESOURCE_INVALID_RELEASE                                BUGCHECK_ERROR = 366
	CLUSTER_CSV_STATE_TRANSITION_INTERVAL_TIMEOUT_LIVEDUMP   BUGCHECK_ERROR = 367
	CLUSTER_CSV_CLUSSVC_DISCONNECT_WATCHDOG                  BUGCHECK_ERROR = 368
	CRYPTO_LIBRARY_INTERNAL_ERROR                            BUGCHECK_ERROR = 369
	COREMSGCALL_INTERNAL_ERROR                               BUGCHECK_ERROR = 371
	COREMSG_INTERNAL_ERROR                                   BUGCHECK_ERROR = 372
	PREVIOUS_FATAL_ABNORMAL_RESET_ERROR                      BUGCHECK_ERROR = 373
	ELAM_DRIVER_DETECTED_FATAL_ERROR                         BUGCHECK_ERROR = 376
	CLUSTER_CLUSPORT_STATUS_IO_TIMEOUT_LIVEDUMP              BUGCHECK_ERROR = 377
	PROFILER_CONFIGURATION_ILLEGAL                           BUGCHECK_ERROR = 379
	PDC_LOCK_WATCHDOG_LIVEDUMP                               BUGCHECK_ERROR = 380
	PDC_UNEXPECTED_REVOCATION_LIVEDUMP                       BUGCHECK_ERROR = 381
	MICROCODE_REVISION_MISMATCH                              BUGCHECK_ERROR = 382
	HYPERGUARD_INITIALIZATION_FAILURE                        BUGCHECK_ERROR = 383
	WVR_LIVEDUMP_REPLICATION_IOCONTEXT_TIMEOUT               BUGCHECK_ERROR = 384
	WVR_LIVEDUMP_STATE_TRANSITION_TIMEOUT                    BUGCHECK_ERROR = 385
	WVR_LIVEDUMP_RECOVERY_IOCONTEXT_TIMEOUT                  BUGCHECK_ERROR = 386
	WVR_LIVEDUMP_APP_IO_TIMEOUT                              BUGCHECK_ERROR = 387
	WVR_LIVEDUMP_MANUALLY_INITIATED                          BUGCHECK_ERROR = 388
	WVR_LIVEDUMP_STATE_FAILURE                               BUGCHECK_ERROR = 389
	WVR_LIVEDUMP_CRITICAL_ERROR                              BUGCHECK_ERROR = 390
	VIDEO_DWMINIT_TIMEOUT_FALLBACK_BDD                       BUGCHECK_ERROR = 391
	CLUSTER_CSVFS_LIVEDUMP                                   BUGCHECK_ERROR = 392
	BAD_OBJECT_HEADER                                        BUGCHECK_ERROR = 393
	SILO_CORRUPT                                             BUGCHECK_ERROR = 394
	SECURE_KERNEL_ERROR                                      BUGCHECK_ERROR = 395
	HYPERGUARD_VIOLATION                                     BUGCHECK_ERROR = 396
	SECURE_FAULT_UNHANDLED                                   BUGCHECK_ERROR = 397
	KERNEL_PARTITION_REFERENCE_VIOLATION                     BUGCHECK_ERROR = 398
	SYNTHETIC_EXCEPTION_UNHANDLED                            BUGCHECK_ERROR = 399
	WIN32K_CRITICAL_FAILURE_LIVEDUMP                         BUGCHECK_ERROR = 400
	PF_DETECTED_CORRUPTION                                   BUGCHECK_ERROR = 401
	KERNEL_AUTO_BOOST_LOCK_ACQUISITION_WITH_RAISED_IRQL      BUGCHECK_ERROR = 402
	VIDEO_DXGKRNL_LIVEDUMP                                   BUGCHECK_ERROR = 403
	KERNEL_STORAGE_SLOT_IN_USE                               BUGCHECK_ERROR = 409
	SMB_SERVER_LIVEDUMP                                      BUGCHECK_ERROR = 405
	LOADER_ROLLBACK_DETECTED                                 BUGCHECK_ERROR = 406
	WIN32K_SECURITY_FAILURE                                  BUGCHECK_ERROR = 407
	UFX_LIVEDUMP                                             BUGCHECK_ERROR = 408
	WORKER_THREAD_RETURNED_WHILE_ATTACHED_TO_SILO            BUGCHECK_ERROR = 410
	TTM_FATAL_ERROR                                          BUGCHECK_ERROR = 411
	WIN32K_POWER_WATCHDOG_TIMEOUT                            BUGCHECK_ERROR = 412
	CLUSTER_SVHDX_LIVEDUMP                                   BUGCHECK_ERROR = 413
	BUGCODE_NETADAPTER_DRIVER                                BUGCHECK_ERROR = 414
	PDC_PRIVILEGE_CHECK_LIVEDUMP                             BUGCHECK_ERROR = 415
	TTM_WATCHDOG_TIMEOUT                                     BUGCHECK_ERROR = 416
	WIN32K_CALLOUT_WATCHDOG_LIVEDUMP                         BUGCHECK_ERROR = 417
	WIN32K_CALLOUT_WATCHDOG_BUGCHECK                         BUGCHECK_ERROR = 418
	CALL_HAS_NOT_RETURNED_WATCHDOG_TIMEOUT_LIVEDUMP          BUGCHECK_ERROR = 419
	DRIPS_SW_HW_DIVERGENCE_LIVEDUMP                          BUGCHECK_ERROR = 420
	USB_DRIPS_BLOCKER_SURPRISE_REMOVAL_LIVEDUMP              BUGCHECK_ERROR = 421
	BLUETOOTH_ERROR_RECOVERY_LIVEDUMP                        BUGCHECK_ERROR = 422
	SMB_REDIRECTOR_LIVEDUMP                                  BUGCHECK_ERROR = 423
	VIDEO_DXGKRNL_BLACK_SCREEN_LIVEDUMP                      BUGCHECK_ERROR = 424
	DIRECTED_FX_TRANSITION_LIVEDUMP                          BUGCHECK_ERROR = 425
	EXCEPTION_ON_INVALID_STACK                               BUGCHECK_ERROR = 426
	UNWIND_ON_INVALID_STACK                                  BUGCHECK_ERROR = 427
	VIDEO_MINIPORT_FAILED_LIVEDUMP                           BUGCHECK_ERROR = 432
	VIDEO_MINIPORT_BLACK_SCREEN_LIVEDUMP                     BUGCHECK_ERROR = 440
	DRIVER_VERIFIER_DETECTED_VIOLATION_LIVEDUMP              BUGCHECK_ERROR = 452
	IO_THREADPOOL_DEADLOCK_LIVEDUMP                          BUGCHECK_ERROR = 453
	FAST_ERESOURCE_PRECONDITION_VIOLATION                    BUGCHECK_ERROR = 454
	STORE_DATA_STRUCTURE_CORRUPTION                          BUGCHECK_ERROR = 455
	MANUALLY_INITIATED_POWER_BUTTON_HOLD                     BUGCHECK_ERROR = 456
	USER_MODE_HEALTH_MONITOR_LIVEDUMP                        BUGCHECK_ERROR = 457
	SYNTHETIC_WATCHDOG_TIMEOUT                               BUGCHECK_ERROR = 458
	INVALID_SILO_DETACH                                      BUGCHECK_ERROR = 459
	EXRESOURCE_TIMEOUT_LIVEDUMP                              BUGCHECK_ERROR = 460
	INVALID_CALLBACK_STACK_ADDRESS                           BUGCHECK_ERROR = 461
	INVALID_KERNEL_STACK_ADDRESS                             BUGCHECK_ERROR = 462
	HARDWARE_WATCHDOG_TIMEOUT                                BUGCHECK_ERROR = 463
	ACPI_FIRMWARE_WATCHDOG_TIMEOUT                           BUGCHECK_ERROR = 464
	TELEMETRY_ASSERTS_LIVEDUMP                               BUGCHECK_ERROR = 465
	WORKER_THREAD_INVALID_STATE                              BUGCHECK_ERROR = 466
	WFP_INVALID_OPERATION                                    BUGCHECK_ERROR = 467
	UCMUCSI_LIVEDUMP                                         BUGCHECK_ERROR = 468
	DRIVER_PNP_WATCHDOG                                      BUGCHECK_ERROR = 469
	WORKER_THREAD_RETURNED_WITH_NON_DEFAULT_WORKLOAD_CLASS   BUGCHECK_ERROR = 470
	EFS_FATAL_ERROR                                          BUGCHECK_ERROR = 471
	UCMUCSI_FAILURE                                          BUGCHECK_ERROR = 472
	HAL_IOMMU_INTERNAL_ERROR                                 BUGCHECK_ERROR = 473
	HAL_BLOCKED_PROCESSOR_INTERNAL_ERROR                     BUGCHECK_ERROR = 474
	IPI_WATCHDOG_TIMEOUT                                     BUGCHECK_ERROR = 475
	DMA_COMMON_BUFFER_VECTOR_ERROR                           BUGCHECK_ERROR = 476
	BUGCODE_MBBADAPTER_DRIVER                                BUGCHECK_ERROR = 477
	BUGCODE_WIFIADAPTER_DRIVER                               BUGCHECK_ERROR = 478
	PROCESSOR_START_TIMEOUT                                  BUGCHECK_ERROR = 479
	INVALID_ALTERNATE_SYSTEM_CALL_HANDLER_REGISTRATION       BUGCHECK_ERROR = 480
	DEVICE_DIAGNOSTIC_LOG_LIVEDUMP                           BUGCHECK_ERROR = 481
	AZURE_DEVICE_FW_DUMP                                     BUGCHECK_ERROR = 482
	BREAKAWAY_CABLE_TRANSITION                               BUGCHECK_ERROR = 483
	VIDEO_DXGKRNL_SYSMM_FATAL_ERROR                          BUGCHECK_ERROR = 484
	DRIVER_VERIFIER_TRACKING_LIVE_DUMP                       BUGCHECK_ERROR = 485
	CRASHDUMP_WATCHDOG_TIMEOUT                               BUGCHECK_ERROR = 486
	REGISTRY_LIVE_DUMP                                       BUGCHECK_ERROR = 487
	INVALID_THREAD_AFFINITY_STATE                            BUGCHECK_ERROR = 488
	ILLEGAL_ATS_INITIALIZATION                               BUGCHECK_ERROR = 489
	SECURE_PCI_CONFIG_SPACE_ACCESS_VIOLATION                 BUGCHECK_ERROR = 490
	DAM_WATCHDOG_TIMEOUT                                     BUGCHECK_ERROR = 491
	HANDLE_LIVE_DUMP                                         BUGCHECK_ERROR = 492
	HANDLE_ERROR_ON_CRITICAL_THREAD                          BUGCHECK_ERROR = 493
	MPSDRV_QUERY_USER                                        BUGCHECK_ERROR = 1073742318
	VMBUS_LIVEDUMP                                           BUGCHECK_ERROR = 1073742319
	USB4_HARDWARE_VIOLATION                                  BUGCHECK_ERROR = 496
	KASAN_ENLIGHTENMENT_VIOLATION                            BUGCHECK_ERROR = 497
	KASAN_ILLEGAL_ACCESS                                     BUGCHECK_ERROR = 498
	IORING                                                   BUGCHECK_ERROR = 499
	MDL_CACHE                                                BUGCHECK_ERROR = 500
	MISALIGNED_POINTER_PARAMETER                             BUGCHECK_ERROR = 502
	XBOX_VMCTRL_CS_TIMEOUT                                   BUGCHECK_ERROR = 854
	XBOX_CORRUPTED_IMAGE                                     BUGCHECK_ERROR = 855
	XBOX_INVERTED_FUNCTION_TABLE_OVERFLOW                    BUGCHECK_ERROR = 856
	XBOX_CORRUPTED_IMAGE_BASE                                BUGCHECK_ERROR = 857
	XBOX_XDS_WATCHDOG_TIMEOUT                                BUGCHECK_ERROR = 858
	XBOX_SHUTDOWN_WATCHDOG_TIMEOUT                           BUGCHECK_ERROR = 859
	XBOX_360_SYSTEM_CRASH                                    BUGCHECK_ERROR = 864
	XBOX_360_SYSTEM_CRASH_RESERVED                           BUGCHECK_ERROR = 1056
	XBOX_SECURITY_FAILUE                                     BUGCHECK_ERROR = 1057
	KERNEL_CFG_INIT_FAILURE                                  BUGCHECK_ERROR = 1058
	MANUALLY_INITIATED_POWER_BUTTON_HOLD_LIVE_DUMP           BUGCHECK_ERROR = 4552
	HYPERVISOR_ERROR                                         BUGCHECK_ERROR = 131073
	XBOX_MANUALLY_INITIATED_CRASH                            BUGCHECK_ERROR = 196614
	MANUALLY_INITIATED_BLACKSCREEN_HOTKEY_LIVE_DUMP          BUGCHECK_ERROR = 8648
	WINLOGON_FATAL_ERROR                                     BUGCHECK_ERROR = 3221226010
	MANUALLY_INITIATED_CRASH1                                BUGCHECK_ERROR = 3735936685
	BUGCHECK_CONTEXT_MODIFIER                                BUGCHECK_ERROR = 2147483648
)

// enum
type FACILITY_CODE uint32

const (
	FACILITY_NULL                                     FACILITY_CODE = 0
	FACILITY_RPC                                      FACILITY_CODE = 1
	FACILITY_DISPATCH                                 FACILITY_CODE = 2
	FACILITY_STORAGE                                  FACILITY_CODE = 3
	FACILITY_ITF                                      FACILITY_CODE = 4
	FACILITY_WIN32                                    FACILITY_CODE = 7
	FACILITY_WINDOWS                                  FACILITY_CODE = 8
	FACILITY_SSPI                                     FACILITY_CODE = 9
	FACILITY_SECURITY                                 FACILITY_CODE = 9
	FACILITY_CONTROL                                  FACILITY_CODE = 10
	FACILITY_CERT                                     FACILITY_CODE = 11
	FACILITY_INTERNET                                 FACILITY_CODE = 12
	FACILITY_MEDIASERVER                              FACILITY_CODE = 13
	FACILITY_MSMQ                                     FACILITY_CODE = 14
	FACILITY_SETUPAPI                                 FACILITY_CODE = 15
	FACILITY_SCARD                                    FACILITY_CODE = 16
	FACILITY_COMPLUS                                  FACILITY_CODE = 17
	FACILITY_AAF                                      FACILITY_CODE = 18
	FACILITY_URT                                      FACILITY_CODE = 19
	FACILITY_ACS                                      FACILITY_CODE = 20
	FACILITY_DPLAY                                    FACILITY_CODE = 21
	FACILITY_UMI                                      FACILITY_CODE = 22
	FACILITY_SXS                                      FACILITY_CODE = 23
	FACILITY_WINDOWS_CE                               FACILITY_CODE = 24
	FACILITY_HTTP                                     FACILITY_CODE = 25
	FACILITY_USERMODE_COMMONLOG                       FACILITY_CODE = 26
	FACILITY_WER                                      FACILITY_CODE = 27
	FACILITY_USERMODE_FILTER_MANAGER                  FACILITY_CODE = 31
	FACILITY_BACKGROUNDCOPY                           FACILITY_CODE = 32
	FACILITY_CONFIGURATION                            FACILITY_CODE = 33
	FACILITY_WIA                                      FACILITY_CODE = 33
	FACILITY_STATE_MANAGEMENT                         FACILITY_CODE = 34
	FACILITY_METADIRECTORY                            FACILITY_CODE = 35
	FACILITY_WINDOWSUPDATE                            FACILITY_CODE = 36
	FACILITY_DIRECTORYSERVICE                         FACILITY_CODE = 37
	FACILITY_GRAPHICS                                 FACILITY_CODE = 38
	FACILITY_SHELL                                    FACILITY_CODE = 39
	FACILITY_NAP                                      FACILITY_CODE = 39
	FACILITY_TPM_SERVICES                             FACILITY_CODE = 40
	FACILITY_TPM_SOFTWARE                             FACILITY_CODE = 41
	FACILITY_UI                                       FACILITY_CODE = 42
	FACILITY_XAML                                     FACILITY_CODE = 43
	FACILITY_ACTION_QUEUE                             FACILITY_CODE = 44
	FACILITY_PLA                                      FACILITY_CODE = 48
	FACILITY_WINDOWS_SETUP                            FACILITY_CODE = 48
	FACILITY_FVE                                      FACILITY_CODE = 49
	FACILITY_FWP                                      FACILITY_CODE = 50
	FACILITY_WINRM                                    FACILITY_CODE = 51
	FACILITY_NDIS                                     FACILITY_CODE = 52
	FACILITY_USERMODE_HYPERVISOR                      FACILITY_CODE = 53
	FACILITY_CMI                                      FACILITY_CODE = 54
	FACILITY_USERMODE_VIRTUALIZATION                  FACILITY_CODE = 55
	FACILITY_USERMODE_VOLMGR                          FACILITY_CODE = 56
	FACILITY_BCD                                      FACILITY_CODE = 57
	FACILITY_USERMODE_VHD                             FACILITY_CODE = 58
	FACILITY_USERMODE_HNS                             FACILITY_CODE = 59
	FACILITY_SDIAG                                    FACILITY_CODE = 60
	FACILITY_WEBSERVICES                              FACILITY_CODE = 61
	FACILITY_WINPE                                    FACILITY_CODE = 61
	FACILITY_WPN                                      FACILITY_CODE = 62
	FACILITY_WINDOWS_STORE                            FACILITY_CODE = 63
	FACILITY_INPUT                                    FACILITY_CODE = 64
	FACILITY_QUIC                                     FACILITY_CODE = 65
	FACILITY_EAP                                      FACILITY_CODE = 66
	FACILITY_IORING                                   FACILITY_CODE = 70
	FACILITY_WINDOWS_DEFENDER                         FACILITY_CODE = 80
	FACILITY_OPC                                      FACILITY_CODE = 81
	FACILITY_XPS                                      FACILITY_CODE = 82
	FACILITY_MBN                                      FACILITY_CODE = 84
	FACILITY_POWERSHELL                               FACILITY_CODE = 84
	FACILITY_RAS                                      FACILITY_CODE = 83
	FACILITY_P2P_INT                                  FACILITY_CODE = 98
	FACILITY_P2P                                      FACILITY_CODE = 99
	FACILITY_DAF                                      FACILITY_CODE = 100
	FACILITY_BLUETOOTH_ATT                            FACILITY_CODE = 101
	FACILITY_AUDIO                                    FACILITY_CODE = 102
	FACILITY_STATEREPOSITORY                          FACILITY_CODE = 103
	FACILITY_VISUALCPP                                FACILITY_CODE = 109
	FACILITY_SCRIPT                                   FACILITY_CODE = 112
	FACILITY_PARSE                                    FACILITY_CODE = 113
	FACILITY_BLB                                      FACILITY_CODE = 120
	FACILITY_BLB_CLI                                  FACILITY_CODE = 121
	FACILITY_WSBAPP                                   FACILITY_CODE = 122
	FACILITY_BLBUI                                    FACILITY_CODE = 128
	FACILITY_USN                                      FACILITY_CODE = 129
	FACILITY_USERMODE_VOLSNAP                         FACILITY_CODE = 130
	FACILITY_TIERING                                  FACILITY_CODE = 131
	FACILITY_WSB_ONLINE                               FACILITY_CODE = 133
	FACILITY_ONLINE_ID                                FACILITY_CODE = 134
	FACILITY_DEVICE_UPDATE_AGENT                      FACILITY_CODE = 135
	FACILITY_DRVSERVICING                             FACILITY_CODE = 136
	FACILITY_DLS                                      FACILITY_CODE = 153
	FACILITY_DELIVERY_OPTIMIZATION                    FACILITY_CODE = 208
	FACILITY_USERMODE_SPACES                          FACILITY_CODE = 231
	FACILITY_USER_MODE_SECURITY_CORE                  FACILITY_CODE = 232
	FACILITY_USERMODE_LICENSING                       FACILITY_CODE = 234
	FACILITY_SOS                                      FACILITY_CODE = 160
	FACILITY_OCP_UPDATE_AGENT                         FACILITY_CODE = 173
	FACILITY_DEBUGGERS                                FACILITY_CODE = 176
	FACILITY_SPP                                      FACILITY_CODE = 256
	FACILITY_RESTORE                                  FACILITY_CODE = 256
	FACILITY_DMSERVER                                 FACILITY_CODE = 256
	FACILITY_DEPLOYMENT_SERVICES_SERVER               FACILITY_CODE = 257
	FACILITY_DEPLOYMENT_SERVICES_IMAGING              FACILITY_CODE = 258
	FACILITY_DEPLOYMENT_SERVICES_MANAGEMENT           FACILITY_CODE = 259
	FACILITY_DEPLOYMENT_SERVICES_UTIL                 FACILITY_CODE = 260
	FACILITY_DEPLOYMENT_SERVICES_BINLSVC              FACILITY_CODE = 261
	FACILITY_DEPLOYMENT_SERVICES_PXE                  FACILITY_CODE = 263
	FACILITY_DEPLOYMENT_SERVICES_TFTP                 FACILITY_CODE = 264
	FACILITY_DEPLOYMENT_SERVICES_TRANSPORT_MANAGEMENT FACILITY_CODE = 272
	FACILITY_DEPLOYMENT_SERVICES_DRIVER_PROVISIONING  FACILITY_CODE = 278
	FACILITY_DEPLOYMENT_SERVICES_MULTICAST_SERVER     FACILITY_CODE = 289
	FACILITY_DEPLOYMENT_SERVICES_MULTICAST_CLIENT     FACILITY_CODE = 290
	FACILITY_DEPLOYMENT_SERVICES_CONTENT_PROVIDER     FACILITY_CODE = 293
	FACILITY_HSP_SERVICES                             FACILITY_CODE = 296
	FACILITY_HSP_SOFTWARE                             FACILITY_CODE = 297
	FACILITY_LINGUISTIC_SERVICES                      FACILITY_CODE = 305
	FACILITY_AUDIOSTREAMING                           FACILITY_CODE = 1094
	FACILITY_TTD                                      FACILITY_CODE = 1490
	FACILITY_ACCELERATOR                              FACILITY_CODE = 1536
	FACILITY_WMAAECMA                                 FACILITY_CODE = 1996
	FACILITY_DIRECTMUSIC                              FACILITY_CODE = 2168
	FACILITY_DIRECT3D10                               FACILITY_CODE = 2169
	FACILITY_DXGI                                     FACILITY_CODE = 2170
	FACILITY_DXGI_DDI                                 FACILITY_CODE = 2171
	FACILITY_DIRECT3D11                               FACILITY_CODE = 2172
	FACILITY_DIRECT3D11_DEBUG                         FACILITY_CODE = 2173
	FACILITY_DIRECT3D12                               FACILITY_CODE = 2174
	FACILITY_DIRECT3D12_DEBUG                         FACILITY_CODE = 2175
	FACILITY_DXCORE                                   FACILITY_CODE = 2176
	FACILITY_PRESENTATION                             FACILITY_CODE = 2177
	FACILITY_LEAP                                     FACILITY_CODE = 2184
	FACILITY_AUDCLNT                                  FACILITY_CODE = 2185
	FACILITY_WINCODEC_DWRITE_DWM                      FACILITY_CODE = 2200
	FACILITY_WINML                                    FACILITY_CODE = 2192
	FACILITY_DIRECT2D                                 FACILITY_CODE = 2201
	FACILITY_DEFRAG                                   FACILITY_CODE = 2304
	FACILITY_USERMODE_SDBUS                           FACILITY_CODE = 2305
	FACILITY_JSCRIPT                                  FACILITY_CODE = 2306
	FACILITY_PIDGENX                                  FACILITY_CODE = 2561
	FACILITY_EAS                                      FACILITY_CODE = 85
	FACILITY_WEB                                      FACILITY_CODE = 885
	FACILITY_WEB_SOCKET                               FACILITY_CODE = 886
	FACILITY_MOBILE                                   FACILITY_CODE = 1793
	FACILITY_SQLITE                                   FACILITY_CODE = 1967
	FACILITY_SERVICE_FABRIC                           FACILITY_CODE = 1968
	FACILITY_UTC                                      FACILITY_CODE = 1989
	FACILITY_WEP                                      FACILITY_CODE = 2049
	FACILITY_SYNCENGINE                               FACILITY_CODE = 2050
	FACILITY_XBOX                                     FACILITY_CODE = 2339
	FACILITY_GAME                                     FACILITY_CODE = 2340
	FACILITY_PIX                                      FACILITY_CODE = 2748
	FACILITY_NT_BIT                                   FACILITY_CODE = 268435456
)

// enum
// flags
type THREAD_ERROR_MODE uint32

const (
	SEM_ALL_ERRORS             THREAD_ERROR_MODE = 0
	SEM_FAILCRITICALERRORS     THREAD_ERROR_MODE = 1
	SEM_NOGPFAULTERRORBOX      THREAD_ERROR_MODE = 2
	SEM_NOOPENFILEERRORBOX     THREAD_ERROR_MODE = 32768
	SEM_NOALIGNMENTFAULTEXCEPT THREAD_ERROR_MODE = 4
)

// enum
// flags
type FORMAT_MESSAGE_OPTIONS uint32

const (
	FORMAT_MESSAGE_ALLOCATE_BUFFER FORMAT_MESSAGE_OPTIONS = 256
	FORMAT_MESSAGE_ARGUMENT_ARRAY  FORMAT_MESSAGE_OPTIONS = 8192
	FORMAT_MESSAGE_FROM_HMODULE    FORMAT_MESSAGE_OPTIONS = 2048
	FORMAT_MESSAGE_FROM_STRING     FORMAT_MESSAGE_OPTIONS = 1024
	FORMAT_MESSAGE_FROM_SYSTEM     FORMAT_MESSAGE_OPTIONS = 4096
	FORMAT_MESSAGE_IGNORE_INSERTS  FORMAT_MESSAGE_OPTIONS = 512
)

// enum
type RTL_VIRTUAL_UNWIND_HANDLER_TYPE uint32

const (
	UNW_FLAG_NHANDLER  RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 0
	UNW_FLAG_EHANDLER  RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 1
	UNW_FLAG_UHANDLER  RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 2
	UNW_FLAG_CHAININFO RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 4
)

// enum
type OPEN_THREAD_WAIT_CHAIN_SESSION_FLAGS uint32

const (
	WCT_ASYNC_OPEN_FLAG OPEN_THREAD_WAIT_CHAIN_SESSION_FLAGS = 1
)

// enum
type SYM_SRV_STORE_FILE_FLAGS uint32

const (
	SYMSTOREOPT_COMPRESS       SYM_SRV_STORE_FILE_FLAGS = 1
	SYMSTOREOPT_OVERWRITE      SYM_SRV_STORE_FILE_FLAGS = 2
	SYMSTOREOPT_PASS_IF_EXISTS SYM_SRV_STORE_FILE_FLAGS = 64
	SYMSTOREOPT_POINTER        SYM_SRV_STORE_FILE_FLAGS = 8
	SYMSTOREOPT_RETURNINDEX    SYM_SRV_STORE_FILE_FLAGS = 4
)

// enum
type IMAGE_DIRECTORY_ENTRY uint16

const (
	IMAGE_DIRECTORY_ENTRY_ARCHITECTURE   IMAGE_DIRECTORY_ENTRY = 7
	IMAGE_DIRECTORY_ENTRY_BASERELOC      IMAGE_DIRECTORY_ENTRY = 5
	IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT   IMAGE_DIRECTORY_ENTRY = 11
	IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR IMAGE_DIRECTORY_ENTRY = 14
	IMAGE_DIRECTORY_ENTRY_DEBUG          IMAGE_DIRECTORY_ENTRY = 6
	IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT   IMAGE_DIRECTORY_ENTRY = 13
	IMAGE_DIRECTORY_ENTRY_EXCEPTION      IMAGE_DIRECTORY_ENTRY = 3
	IMAGE_DIRECTORY_ENTRY_EXPORT         IMAGE_DIRECTORY_ENTRY = 0
	IMAGE_DIRECTORY_ENTRY_GLOBALPTR      IMAGE_DIRECTORY_ENTRY = 8
	IMAGE_DIRECTORY_ENTRY_IAT            IMAGE_DIRECTORY_ENTRY = 12
	IMAGE_DIRECTORY_ENTRY_IMPORT         IMAGE_DIRECTORY_ENTRY = 1
	IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG    IMAGE_DIRECTORY_ENTRY = 10
	IMAGE_DIRECTORY_ENTRY_RESOURCE       IMAGE_DIRECTORY_ENTRY = 2
	IMAGE_DIRECTORY_ENTRY_SECURITY       IMAGE_DIRECTORY_ENTRY = 4
	IMAGE_DIRECTORY_ENTRY_TLS            IMAGE_DIRECTORY_ENTRY = 9
)

// enum
type WAIT_CHAIN_THREAD_OPTIONS uint32

const (
	WCT_OUT_OF_PROC_COM_FLAG WAIT_CHAIN_THREAD_OPTIONS = 2
	WCT_OUT_OF_PROC_CS_FLAG  WAIT_CHAIN_THREAD_OPTIONS = 4
	WCT_OUT_OF_PROC_FLAG     WAIT_CHAIN_THREAD_OPTIONS = 1
)

// enum
type SYM_FIND_ID_OPTION uint32

const (
	SSRVOPT_DWORD    SYM_FIND_ID_OPTION = 2
	SSRVOPT_DWORDPTR SYM_FIND_ID_OPTION = 4
	SSRVOPT_GUIDPTR  SYM_FIND_ID_OPTION = 8
)

// enum
// flags
type IMAGE_FILE_CHARACTERISTICS uint16

const (
	IMAGE_FILE_RELOCS_STRIPPED         IMAGE_FILE_CHARACTERISTICS = 1
	IMAGE_FILE_EXECUTABLE_IMAGE        IMAGE_FILE_CHARACTERISTICS = 2
	IMAGE_FILE_LINE_NUMS_STRIPPED      IMAGE_FILE_CHARACTERISTICS = 4
	IMAGE_FILE_LOCAL_SYMS_STRIPPED     IMAGE_FILE_CHARACTERISTICS = 8
	IMAGE_FILE_AGGRESIVE_WS_TRIM       IMAGE_FILE_CHARACTERISTICS = 16
	IMAGE_FILE_LARGE_ADDRESS_AWARE     IMAGE_FILE_CHARACTERISTICS = 32
	IMAGE_FILE_BYTES_REVERSED_LO       IMAGE_FILE_CHARACTERISTICS = 128
	IMAGE_FILE_32BIT_MACHINE           IMAGE_FILE_CHARACTERISTICS = 256
	IMAGE_FILE_DEBUG_STRIPPED          IMAGE_FILE_CHARACTERISTICS = 512
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP IMAGE_FILE_CHARACTERISTICS = 1024
	IMAGE_FILE_NET_RUN_FROM_SWAP       IMAGE_FILE_CHARACTERISTICS = 2048
	IMAGE_FILE_SYSTEM                  IMAGE_FILE_CHARACTERISTICS = 4096
	IMAGE_FILE_DLL                     IMAGE_FILE_CHARACTERISTICS = 8192
	IMAGE_FILE_UP_SYSTEM_ONLY          IMAGE_FILE_CHARACTERISTICS = 16384
	IMAGE_FILE_BYTES_REVERSED_HI       IMAGE_FILE_CHARACTERISTICS = 32768
)

// enum
// flags
type IMAGE_FILE_CHARACTERISTICS2 uint32

const (
	IMAGE_FILE_RELOCS_STRIPPED2         IMAGE_FILE_CHARACTERISTICS2 = 1
	IMAGE_FILE_EXECUTABLE_IMAGE2        IMAGE_FILE_CHARACTERISTICS2 = 2
	IMAGE_FILE_LINE_NUMS_STRIPPED2      IMAGE_FILE_CHARACTERISTICS2 = 4
	IMAGE_FILE_LOCAL_SYMS_STRIPPED2     IMAGE_FILE_CHARACTERISTICS2 = 8
	IMAGE_FILE_AGGRESIVE_WS_TRIM2       IMAGE_FILE_CHARACTERISTICS2 = 16
	IMAGE_FILE_LARGE_ADDRESS_AWARE2     IMAGE_FILE_CHARACTERISTICS2 = 32
	IMAGE_FILE_BYTES_REVERSED_LO2       IMAGE_FILE_CHARACTERISTICS2 = 128
	IMAGE_FILE_32BIT_MACHINE2           IMAGE_FILE_CHARACTERISTICS2 = 256
	IMAGE_FILE_DEBUG_STRIPPED2          IMAGE_FILE_CHARACTERISTICS2 = 512
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP2 IMAGE_FILE_CHARACTERISTICS2 = 1024
	IMAGE_FILE_NET_RUN_FROM_SWAP2       IMAGE_FILE_CHARACTERISTICS2 = 2048
	IMAGE_FILE_SYSTEM_2                 IMAGE_FILE_CHARACTERISTICS2 = 4096
	IMAGE_FILE_DLL_2                    IMAGE_FILE_CHARACTERISTICS2 = 8192
	IMAGE_FILE_UP_SYSTEM_ONLY_2         IMAGE_FILE_CHARACTERISTICS2 = 16384
	IMAGE_FILE_BYTES_REVERSED_HI_2      IMAGE_FILE_CHARACTERISTICS2 = 32768
)

// enum
// flags
type SYMBOL_INFO_FLAGS uint32

const (
	SYMFLAG_CLR_TOKEN    SYMBOL_INFO_FLAGS = 262144
	SYMFLAG_CONSTANT     SYMBOL_INFO_FLAGS = 256
	SYMFLAG_EXPORT       SYMBOL_INFO_FLAGS = 512
	SYMFLAG_FORWARDER    SYMBOL_INFO_FLAGS = 1024
	SYMFLAG_FRAMEREL     SYMBOL_INFO_FLAGS = 32
	SYMFLAG_FUNCTION     SYMBOL_INFO_FLAGS = 2048
	SYMFLAG_ILREL        SYMBOL_INFO_FLAGS = 65536
	SYMFLAG_LOCAL        SYMBOL_INFO_FLAGS = 128
	SYMFLAG_METADATA     SYMBOL_INFO_FLAGS = 131072
	SYMFLAG_PARAMETER    SYMBOL_INFO_FLAGS = 64
	SYMFLAG_REGISTER     SYMBOL_INFO_FLAGS = 8
	SYMFLAG_REGREL       SYMBOL_INFO_FLAGS = 16
	SYMFLAG_SLOT         SYMBOL_INFO_FLAGS = 32768
	SYMFLAG_THUNK        SYMBOL_INFO_FLAGS = 8192
	SYMFLAG_TLSREL       SYMBOL_INFO_FLAGS = 16384
	SYMFLAG_VALUEPRESENT SYMBOL_INFO_FLAGS = 1
	SYMFLAG_VIRTUAL      SYMBOL_INFO_FLAGS = 4096
)

// enum
type IMAGEHLP_CBA_EVENT_SEVERITY uint32

const (
	SevInfo    IMAGEHLP_CBA_EVENT_SEVERITY = 0
	SevProblem IMAGEHLP_CBA_EVENT_SEVERITY = 1
	SevAttn    IMAGEHLP_CBA_EVENT_SEVERITY = 2
	SevFatal   IMAGEHLP_CBA_EVENT_SEVERITY = 3
)

// enum
type IMAGEHLP_GET_TYPE_INFO_FLAGS uint32

const (
	IMAGEHLP_GET_TYPE_INFO_CHILDREN IMAGEHLP_GET_TYPE_INFO_FLAGS = 2
	IMAGEHLP_GET_TYPE_INFO_UNCACHED IMAGEHLP_GET_TYPE_INFO_FLAGS = 1
)

// enum
type RIP_INFO_TYPE uint32

const (
	SLE_ERROR      RIP_INFO_TYPE = 1
	SLE_MINORERROR RIP_INFO_TYPE = 2
	SLE_WARNING    RIP_INFO_TYPE = 3
)

// enum
type VER_PLATFORM uint32

const (
	VER_PLATFORM_WIN32s        VER_PLATFORM = 0
	VER_PLATFORM_WIN32_WINDOWS VER_PLATFORM = 1
	VER_PLATFORM_WIN32_NT      VER_PLATFORM = 2
)

// enum
type IMAGE_DEBUG_TYPE uint32

const (
	IMAGE_DEBUG_TYPE_UNKNOWN   IMAGE_DEBUG_TYPE = 0
	IMAGE_DEBUG_TYPE_COFF      IMAGE_DEBUG_TYPE = 1
	IMAGE_DEBUG_TYPE_CODEVIEW  IMAGE_DEBUG_TYPE = 2
	IMAGE_DEBUG_TYPE_FPO       IMAGE_DEBUG_TYPE = 3
	IMAGE_DEBUG_TYPE_MISC      IMAGE_DEBUG_TYPE = 4
	IMAGE_DEBUG_TYPE_EXCEPTION IMAGE_DEBUG_TYPE = 5
	IMAGE_DEBUG_TYPE_FIXUP     IMAGE_DEBUG_TYPE = 6
	IMAGE_DEBUG_TYPE_BORLAND   IMAGE_DEBUG_TYPE = 9
)

// enum
type MINIDUMP_THREAD_INFO_DUMP_FLAGS uint32

const (
	MINIDUMP_THREAD_INFO_ERROR_THREAD    MINIDUMP_THREAD_INFO_DUMP_FLAGS = 1
	MINIDUMP_THREAD_INFO_EXITED_THREAD   MINIDUMP_THREAD_INFO_DUMP_FLAGS = 4
	MINIDUMP_THREAD_INFO_INVALID_CONTEXT MINIDUMP_THREAD_INFO_DUMP_FLAGS = 16
	MINIDUMP_THREAD_INFO_INVALID_INFO    MINIDUMP_THREAD_INFO_DUMP_FLAGS = 8
	MINIDUMP_THREAD_INFO_INVALID_TEB     MINIDUMP_THREAD_INFO_DUMP_FLAGS = 32
	MINIDUMP_THREAD_INFO_WRITING_THREAD  MINIDUMP_THREAD_INFO_DUMP_FLAGS = 2
)

// enum
type DEBUG_EVENT_CODE uint32

const (
	CREATE_PROCESS_DEBUG_EVENT DEBUG_EVENT_CODE = 3
	CREATE_THREAD_DEBUG_EVENT  DEBUG_EVENT_CODE = 2
	EXCEPTION_DEBUG_EVENT      DEBUG_EVENT_CODE = 1
	EXIT_PROCESS_DEBUG_EVENT   DEBUG_EVENT_CODE = 5
	EXIT_THREAD_DEBUG_EVENT    DEBUG_EVENT_CODE = 4
	LOAD_DLL_DEBUG_EVENT       DEBUG_EVENT_CODE = 6
	OUTPUT_DEBUG_STRING_EVENT  DEBUG_EVENT_CODE = 8
	RIP_EVENT                  DEBUG_EVENT_CODE = 9
	UNLOAD_DLL_DEBUG_EVENT     DEBUG_EVENT_CODE = 7
)

// enum
// flags
type MINIDUMP_MISC_INFO_FLAGS uint32

const (
	MINIDUMP_MISC1_PROCESS_ID    MINIDUMP_MISC_INFO_FLAGS = 1
	MINIDUMP_MISC1_PROCESS_TIMES MINIDUMP_MISC_INFO_FLAGS = 2
)

// enum
type MODLOAD_DATA_TYPE uint32

const (
	DBHHEADER_DEBUGDIRS MODLOAD_DATA_TYPE = 1
	DBHHEADER_CVMISC    MODLOAD_DATA_TYPE = 2
)

// enum
// flags
type CONTEXT_FLAGS uint32

const (
	CONTEXT_AMD64                     CONTEXT_FLAGS = 1048576
	CONTEXT_CONTROL_AMD64             CONTEXT_FLAGS = 1048577
	CONTEXT_INTEGER_AMD64             CONTEXT_FLAGS = 1048578
	CONTEXT_SEGMENTS_AMD64            CONTEXT_FLAGS = 1048580
	CONTEXT_FLOATING_POINT_AMD64      CONTEXT_FLAGS = 1048584
	CONTEXT_DEBUG_REGISTERS_AMD64     CONTEXT_FLAGS = 1048592
	CONTEXT_FULL_AMD64                CONTEXT_FLAGS = 1048587
	CONTEXT_ALL_AMD64                 CONTEXT_FLAGS = 1048607
	CONTEXT_XSTATE_AMD64              CONTEXT_FLAGS = 1048640
	CONTEXT_KERNEL_CET_AMD64          CONTEXT_FLAGS = 1048704
	CONTEXT_KERNEL_DEBUGGER_AMD64     CONTEXT_FLAGS = 67108864
	CONTEXT_EXCEPTION_ACTIVE_AMD64    CONTEXT_FLAGS = 134217728
	CONTEXT_SERVICE_ACTIVE_AMD64      CONTEXT_FLAGS = 268435456
	CONTEXT_EXCEPTION_REQUEST_AMD64   CONTEXT_FLAGS = 1073741824
	CONTEXT_EXCEPTION_REPORTING_AMD64 CONTEXT_FLAGS = 2147483648
	CONTEXT_UNWOUND_TO_CALL_AMD64     CONTEXT_FLAGS = 536870912
	CONTEXT_X86                       CONTEXT_FLAGS = 65536
	CONTEXT_CONTROL_X86               CONTEXT_FLAGS = 65537
	CONTEXT_INTEGER_X86               CONTEXT_FLAGS = 65538
	CONTEXT_SEGMENTS_X86              CONTEXT_FLAGS = 65540
	CONTEXT_FLOATING_POINT_X86        CONTEXT_FLAGS = 65544
	CONTEXT_DEBUG_REGISTERS_X86       CONTEXT_FLAGS = 65552
	CONTEXT_EXTENDED_REGISTERS_X86    CONTEXT_FLAGS = 65568
	CONTEXT_FULL_X86                  CONTEXT_FLAGS = 65543
	CONTEXT_ALL_X86                   CONTEXT_FLAGS = 65599
	CONTEXT_XSTATE_X86                CONTEXT_FLAGS = 65600
	CONTEXT_EXCEPTION_ACTIVE_X86      CONTEXT_FLAGS = 134217728
	CONTEXT_SERVICE_ACTIVE_X86        CONTEXT_FLAGS = 268435456
	CONTEXT_EXCEPTION_REQUEST_X86     CONTEXT_FLAGS = 1073741824
	CONTEXT_EXCEPTION_REPORTING_X86   CONTEXT_FLAGS = 2147483648
	CONTEXT_ARM64                     CONTEXT_FLAGS = 4194304
	CONTEXT_CONTROL_ARM64             CONTEXT_FLAGS = 4194305
	CONTEXT_INTEGER_ARM64             CONTEXT_FLAGS = 4194306
	CONTEXT_FLOATING_POINT_ARM64      CONTEXT_FLAGS = 4194308
	CONTEXT_DEBUG_REGISTERS_ARM64     CONTEXT_FLAGS = 4194312
	CONTEXT_X18_ARM64                 CONTEXT_FLAGS = 4194320
	CONTEXT_FULL_ARM64                CONTEXT_FLAGS = 4194311
	CONTEXT_ALL_ARM64                 CONTEXT_FLAGS = 4194335
	CONTEXT_EXCEPTION_ACTIVE_ARM64    CONTEXT_FLAGS = 134217728
	CONTEXT_SERVICE_ACTIVE_ARM64      CONTEXT_FLAGS = 268435456
	CONTEXT_EXCEPTION_REQUEST_ARM64   CONTEXT_FLAGS = 1073741824
	CONTEXT_EXCEPTION_REPORTING_ARM64 CONTEXT_FLAGS = 2147483648
	CONTEXT_UNWOUND_TO_CALL_ARM64     CONTEXT_FLAGS = 536870912
	CONTEXT_RET_TO_GUEST_ARM64        CONTEXT_FLAGS = 1073741824
	CONTEXT_ARM                       CONTEXT_FLAGS = 2097152
	CONTEXT_CONTROL_ARM               CONTEXT_FLAGS = 2097153
	CONTEXT_INTEGER_ARM               CONTEXT_FLAGS = 2097154
	CONTEXT_FLOATING_POINT_ARM        CONTEXT_FLAGS = 2097156
	CONTEXT_DEBUG_REGISTERS_ARM       CONTEXT_FLAGS = 2097160
	CONTEXT_FULL_ARM                  CONTEXT_FLAGS = 2097159
	CONTEXT_ALL_ARM                   CONTEXT_FLAGS = 2097167
	CONTEXT_EXCEPTION_ACTIVE_ARM      CONTEXT_FLAGS = 134217728
	CONTEXT_SERVICE_ACTIVE_ARM        CONTEXT_FLAGS = 268435456
	CONTEXT_EXCEPTION_REQUEST_ARM     CONTEXT_FLAGS = 1073741824
	CONTEXT_EXCEPTION_REPORTING_ARM   CONTEXT_FLAGS = 2147483648
	CONTEXT_UNWOUND_TO_CALL_ARM       CONTEXT_FLAGS = 536870912
)

// enum
// flags
type WOW64_CONTEXT_FLAGS uint32

const (
	WOW64_CONTEXT_X86                 WOW64_CONTEXT_FLAGS = 65536
	WOW64_CONTEXT_CONTROL             WOW64_CONTEXT_FLAGS = 65537
	WOW64_CONTEXT_INTEGER             WOW64_CONTEXT_FLAGS = 65538
	WOW64_CONTEXT_SEGMENTS            WOW64_CONTEXT_FLAGS = 65540
	WOW64_CONTEXT_FLOATING_POINT      WOW64_CONTEXT_FLAGS = 65544
	WOW64_CONTEXT_DEBUG_REGISTERS     WOW64_CONTEXT_FLAGS = 65552
	WOW64_CONTEXT_EXTENDED_REGISTERS  WOW64_CONTEXT_FLAGS = 65568
	WOW64_CONTEXT_FULL                WOW64_CONTEXT_FLAGS = 65543
	WOW64_CONTEXT_ALL                 WOW64_CONTEXT_FLAGS = 65599
	WOW64_CONTEXT_XSTATE              WOW64_CONTEXT_FLAGS = 65600
	WOW64_CONTEXT_EXCEPTION_ACTIVE    WOW64_CONTEXT_FLAGS = 134217728
	WOW64_CONTEXT_SERVICE_ACTIVE      WOW64_CONTEXT_FLAGS = 268435456
	WOW64_CONTEXT_EXCEPTION_REQUEST   WOW64_CONTEXT_FLAGS = 1073741824
	WOW64_CONTEXT_EXCEPTION_REPORTING WOW64_CONTEXT_FLAGS = 2147483648
)

// enum
type WCT_OBJECT_TYPE int32

const (
	WctCriticalSectionType WCT_OBJECT_TYPE = 1
	WctSendMessageType     WCT_OBJECT_TYPE = 2
	WctMutexType           WCT_OBJECT_TYPE = 3
	WctAlpcType            WCT_OBJECT_TYPE = 4
	WctComType             WCT_OBJECT_TYPE = 5
	WctThreadWaitType      WCT_OBJECT_TYPE = 6
	WctProcessWaitType     WCT_OBJECT_TYPE = 7
	WctThreadType          WCT_OBJECT_TYPE = 8
	WctComActivationType   WCT_OBJECT_TYPE = 9
	WctUnknownType         WCT_OBJECT_TYPE = 10
	WctSocketIoType        WCT_OBJECT_TYPE = 11
	WctSmbIoType           WCT_OBJECT_TYPE = 12
	WctMaxType             WCT_OBJECT_TYPE = 13
)

// enum
type WCT_OBJECT_STATUS int32

const (
	WctStatusNoAccess     WCT_OBJECT_STATUS = 1
	WctStatusRunning      WCT_OBJECT_STATUS = 2
	WctStatusBlocked      WCT_OBJECT_STATUS = 3
	WctStatusPidOnly      WCT_OBJECT_STATUS = 4
	WctStatusPidOnlyRpcss WCT_OBJECT_STATUS = 5
	WctStatusOwned        WCT_OBJECT_STATUS = 6
	WctStatusNotOwned     WCT_OBJECT_STATUS = 7
	WctStatusAbandoned    WCT_OBJECT_STATUS = 8
	WctStatusUnknown      WCT_OBJECT_STATUS = 9
	WctStatusError        WCT_OBJECT_STATUS = 10
	WctStatusMax          WCT_OBJECT_STATUS = 11
)

// enum
type MINIDUMP_STREAM_TYPE int32

const (
	UnusedStream                MINIDUMP_STREAM_TYPE = 0
	ReservedStream0             MINIDUMP_STREAM_TYPE = 1
	ReservedStream1             MINIDUMP_STREAM_TYPE = 2
	ThreadListStream            MINIDUMP_STREAM_TYPE = 3
	ModuleListStream            MINIDUMP_STREAM_TYPE = 4
	MemoryListStream            MINIDUMP_STREAM_TYPE = 5
	ExceptionStream             MINIDUMP_STREAM_TYPE = 6
	SystemInfoStream            MINIDUMP_STREAM_TYPE = 7
	ThreadExListStream          MINIDUMP_STREAM_TYPE = 8
	Memory64ListStream          MINIDUMP_STREAM_TYPE = 9
	CommentStreamA              MINIDUMP_STREAM_TYPE = 10
	CommentStreamW              MINIDUMP_STREAM_TYPE = 11
	HandleDataStream            MINIDUMP_STREAM_TYPE = 12
	FunctionTableStream         MINIDUMP_STREAM_TYPE = 13
	UnloadedModuleListStream    MINIDUMP_STREAM_TYPE = 14
	MiscInfoStream              MINIDUMP_STREAM_TYPE = 15
	MemoryInfoListStream        MINIDUMP_STREAM_TYPE = 16
	ThreadInfoListStream        MINIDUMP_STREAM_TYPE = 17
	HandleOperationListStream   MINIDUMP_STREAM_TYPE = 18
	TokenStream                 MINIDUMP_STREAM_TYPE = 19
	JavaScriptDataStream        MINIDUMP_STREAM_TYPE = 20
	SystemMemoryInfoStream      MINIDUMP_STREAM_TYPE = 21
	ProcessVmCountersStream     MINIDUMP_STREAM_TYPE = 22
	IptTraceStream              MINIDUMP_STREAM_TYPE = 23
	ThreadNamesStream           MINIDUMP_STREAM_TYPE = 24
	CeStreamNull                MINIDUMP_STREAM_TYPE = 32768
	CeStreamSystemInfo          MINIDUMP_STREAM_TYPE = 32769
	CeStreamException           MINIDUMP_STREAM_TYPE = 32770
	CeStreamModuleList          MINIDUMP_STREAM_TYPE = 32771
	CeStreamProcessList         MINIDUMP_STREAM_TYPE = 32772
	CeStreamThreadList          MINIDUMP_STREAM_TYPE = 32773
	CeStreamThreadContextList   MINIDUMP_STREAM_TYPE = 32774
	CeStreamThreadCallStackList MINIDUMP_STREAM_TYPE = 32775
	CeStreamMemoryVirtualList   MINIDUMP_STREAM_TYPE = 32776
	CeStreamMemoryPhysicalList  MINIDUMP_STREAM_TYPE = 32777
	CeStreamBucketParameters    MINIDUMP_STREAM_TYPE = 32778
	CeStreamProcessModuleMap    MINIDUMP_STREAM_TYPE = 32779
	CeStreamDiagnosisList       MINIDUMP_STREAM_TYPE = 32780
	LastReservedStream          MINIDUMP_STREAM_TYPE = 65535
)

// enum
type MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE int32

const (
	MiniHandleObjectInformationNone    MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 0
	MiniThreadInformation1             MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 1
	MiniMutantInformation1             MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 2
	MiniMutantInformation2             MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 3
	MiniProcessInformation1            MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 4
	MiniProcessInformation2            MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 5
	MiniEventInformation1              MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 6
	MiniSectionInformation1            MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 7
	MiniSemaphoreInformation1          MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 8
	MiniHandleObjectInformationTypeMax MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = 9
)

// enum
type MINIDUMP_CALLBACK_TYPE int32

const (
	ModuleCallback               MINIDUMP_CALLBACK_TYPE = 0
	ThreadCallback               MINIDUMP_CALLBACK_TYPE = 1
	ThreadExCallback             MINIDUMP_CALLBACK_TYPE = 2
	IncludeThreadCallback        MINIDUMP_CALLBACK_TYPE = 3
	IncludeModuleCallback        MINIDUMP_CALLBACK_TYPE = 4
	MemoryCallback               MINIDUMP_CALLBACK_TYPE = 5
	CancelCallback               MINIDUMP_CALLBACK_TYPE = 6
	WriteKernelMinidumpCallback  MINIDUMP_CALLBACK_TYPE = 7
	KernelMinidumpStatusCallback MINIDUMP_CALLBACK_TYPE = 8
	RemoveMemoryCallback         MINIDUMP_CALLBACK_TYPE = 9
	IncludeVmRegionCallback      MINIDUMP_CALLBACK_TYPE = 10
	IoStartCallback              MINIDUMP_CALLBACK_TYPE = 11
	IoWriteAllCallback           MINIDUMP_CALLBACK_TYPE = 12
	IoFinishCallback             MINIDUMP_CALLBACK_TYPE = 13
	ReadMemoryFailureCallback    MINIDUMP_CALLBACK_TYPE = 14
	SecondaryFlagsCallback       MINIDUMP_CALLBACK_TYPE = 15
	IsProcessSnapshotCallback    MINIDUMP_CALLBACK_TYPE = 16
	VmStartCallback              MINIDUMP_CALLBACK_TYPE = 17
	VmQueryCallback              MINIDUMP_CALLBACK_TYPE = 18
	VmPreReadCallback            MINIDUMP_CALLBACK_TYPE = 19
	VmPostReadCallback           MINIDUMP_CALLBACK_TYPE = 20
)

// enum
type THREAD_WRITE_FLAGS int32

const (
	ThreadWriteThread            THREAD_WRITE_FLAGS = 1
	ThreadWriteStack             THREAD_WRITE_FLAGS = 2
	ThreadWriteContext           THREAD_WRITE_FLAGS = 4
	ThreadWriteBackingStore      THREAD_WRITE_FLAGS = 8
	ThreadWriteInstructionWindow THREAD_WRITE_FLAGS = 16
	ThreadWriteThreadData        THREAD_WRITE_FLAGS = 32
	ThreadWriteThreadInfo        THREAD_WRITE_FLAGS = 64
)

// enum
type MODULE_WRITE_FLAGS int32

const (
	ModuleWriteModule        MODULE_WRITE_FLAGS = 1
	ModuleWriteDataSeg       MODULE_WRITE_FLAGS = 2
	ModuleWriteMiscRecord    MODULE_WRITE_FLAGS = 4
	ModuleWriteCvRecord      MODULE_WRITE_FLAGS = 8
	ModuleReferencedByMemory MODULE_WRITE_FLAGS = 16
	ModuleWriteTlsData       MODULE_WRITE_FLAGS = 32
	ModuleWriteCodeSegs      MODULE_WRITE_FLAGS = 64
)

// enum
// flags
type MINIDUMP_TYPE int32

const (
	MiniDumpNormal                         MINIDUMP_TYPE = 0
	MiniDumpWithDataSegs                   MINIDUMP_TYPE = 1
	MiniDumpWithFullMemory                 MINIDUMP_TYPE = 2
	MiniDumpWithHandleData                 MINIDUMP_TYPE = 4
	MiniDumpFilterMemory                   MINIDUMP_TYPE = 8
	MiniDumpScanMemory                     MINIDUMP_TYPE = 16
	MiniDumpWithUnloadedModules            MINIDUMP_TYPE = 32
	MiniDumpWithIndirectlyReferencedMemory MINIDUMP_TYPE = 64
	MiniDumpFilterModulePaths              MINIDUMP_TYPE = 128
	MiniDumpWithProcessThreadData          MINIDUMP_TYPE = 256
	MiniDumpWithPrivateReadWriteMemory     MINIDUMP_TYPE = 512
	MiniDumpWithoutOptionalData            MINIDUMP_TYPE = 1024
	MiniDumpWithFullMemoryInfo             MINIDUMP_TYPE = 2048
	MiniDumpWithThreadInfo                 MINIDUMP_TYPE = 4096
	MiniDumpWithCodeSegs                   MINIDUMP_TYPE = 8192
	MiniDumpWithoutAuxiliaryState          MINIDUMP_TYPE = 16384
	MiniDumpWithFullAuxiliaryState         MINIDUMP_TYPE = 32768
	MiniDumpWithPrivateWriteCopyMemory     MINIDUMP_TYPE = 65536
	MiniDumpIgnoreInaccessibleMemory       MINIDUMP_TYPE = 131072
	MiniDumpWithTokenInformation           MINIDUMP_TYPE = 262144
	MiniDumpWithModuleHeaders              MINIDUMP_TYPE = 524288
	MiniDumpFilterTriage                   MINIDUMP_TYPE = 1048576
	MiniDumpWithAvxXStateContext           MINIDUMP_TYPE = 2097152
	MiniDumpWithIptTrace                   MINIDUMP_TYPE = 4194304
	MiniDumpScanInaccessiblePartialPages   MINIDUMP_TYPE = 8388608
	MiniDumpFilterWriteCombinedMemory      MINIDUMP_TYPE = 16777216
	MiniDumpValidTypeFlags                 MINIDUMP_TYPE = 33554431
)

// enum
type MINIDUMP_SECONDARY_FLAGS int32

const (
	MiniSecondaryWithoutPowerInfo MINIDUMP_SECONDARY_FLAGS = 1
	MiniSecondaryValidFlags       MINIDUMP_SECONDARY_FLAGS = 1
)

// enum
type IMAGEHLP_STATUS_REASON int32

const (
	BindOutOfMemory           IMAGEHLP_STATUS_REASON = 0
	BindRvaToVaFailed         IMAGEHLP_STATUS_REASON = 1
	BindNoRoomInImage         IMAGEHLP_STATUS_REASON = 2
	BindImportModuleFailed    IMAGEHLP_STATUS_REASON = 3
	BindImportProcedureFailed IMAGEHLP_STATUS_REASON = 4
	BindImportModule          IMAGEHLP_STATUS_REASON = 5
	BindImportProcedure       IMAGEHLP_STATUS_REASON = 6
	BindForwarder             IMAGEHLP_STATUS_REASON = 7
	BindForwarderNOT          IMAGEHLP_STATUS_REASON = 8
	BindImageModified         IMAGEHLP_STATUS_REASON = 9
	BindExpandFileHeaders     IMAGEHLP_STATUS_REASON = 10
	BindImageComplete         IMAGEHLP_STATUS_REASON = 11
	BindMismatchedSymbols     IMAGEHLP_STATUS_REASON = 12
	BindSymbolsNotUpdated     IMAGEHLP_STATUS_REASON = 13
	BindImportProcedure32     IMAGEHLP_STATUS_REASON = 14
	BindImportProcedure64     IMAGEHLP_STATUS_REASON = 15
	BindForwarder32           IMAGEHLP_STATUS_REASON = 16
	BindForwarder64           IMAGEHLP_STATUS_REASON = 17
	BindForwarderNOT32        IMAGEHLP_STATUS_REASON = 18
	BindForwarderNOT64        IMAGEHLP_STATUS_REASON = 19
)

// enum
type ADDRESS_MODE int32

const (
	AddrMode1616 ADDRESS_MODE = 0
	AddrMode1632 ADDRESS_MODE = 1
	AddrModeReal ADDRESS_MODE = 2
	AddrModeFlat ADDRESS_MODE = 3
)

// enum
type SYM_TYPE int32

const (
	SymNone     SYM_TYPE = 0
	SymCoff     SYM_TYPE = 1
	SymCv       SYM_TYPE = 2
	SymPdb      SYM_TYPE = 3
	SymExport   SYM_TYPE = 4
	SymDeferred SYM_TYPE = 5
	SymSym      SYM_TYPE = 6
	SymDia      SYM_TYPE = 7
	SymVirtual  SYM_TYPE = 8
	NumSymTypes SYM_TYPE = 9
)

// enum
type IMAGEHLP_HD_TYPE int32

const (
	HdBase IMAGEHLP_HD_TYPE = 0
	HdSym  IMAGEHLP_HD_TYPE = 1
	HdSrc  IMAGEHLP_HD_TYPE = 2
	HdMax  IMAGEHLP_HD_TYPE = 3
)

// enum
type IMAGEHLP_EXTENDED_OPTIONS int32

const (
	SYMOPT_EX_DISABLEACCESSTIMEUPDATE IMAGEHLP_EXTENDED_OPTIONS = 0
	SYMOPT_EX_LASTVALIDDEBUGDIRECTORY IMAGEHLP_EXTENDED_OPTIONS = 1
	SYMOPT_EX_NOIMPLICITPATTERNSEARCH IMAGEHLP_EXTENDED_OPTIONS = 2
	SYMOPT_EX_NEVERLOADSYMBOLS        IMAGEHLP_EXTENDED_OPTIONS = 3
	SYMOPT_EX_MAX                     IMAGEHLP_EXTENDED_OPTIONS = 4
)

// enum
type IMAGEHLP_SYMBOL_TYPE_INFO int32

const (
	TI_GET_SYMTAG                   IMAGEHLP_SYMBOL_TYPE_INFO = 0
	TI_GET_SYMNAME                  IMAGEHLP_SYMBOL_TYPE_INFO = 1
	TI_GET_LENGTH                   IMAGEHLP_SYMBOL_TYPE_INFO = 2
	TI_GET_TYPE                     IMAGEHLP_SYMBOL_TYPE_INFO = 3
	TI_GET_TYPEID                   IMAGEHLP_SYMBOL_TYPE_INFO = 4
	TI_GET_BASETYPE                 IMAGEHLP_SYMBOL_TYPE_INFO = 5
	TI_GET_ARRAYINDEXTYPEID         IMAGEHLP_SYMBOL_TYPE_INFO = 6
	TI_FINDCHILDREN                 IMAGEHLP_SYMBOL_TYPE_INFO = 7
	TI_GET_DATAKIND                 IMAGEHLP_SYMBOL_TYPE_INFO = 8
	TI_GET_ADDRESSOFFSET            IMAGEHLP_SYMBOL_TYPE_INFO = 9
	TI_GET_OFFSET                   IMAGEHLP_SYMBOL_TYPE_INFO = 10
	TI_GET_VALUE                    IMAGEHLP_SYMBOL_TYPE_INFO = 11
	TI_GET_COUNT                    IMAGEHLP_SYMBOL_TYPE_INFO = 12
	TI_GET_CHILDRENCOUNT            IMAGEHLP_SYMBOL_TYPE_INFO = 13
	TI_GET_BITPOSITION              IMAGEHLP_SYMBOL_TYPE_INFO = 14
	TI_GET_VIRTUALBASECLASS         IMAGEHLP_SYMBOL_TYPE_INFO = 15
	TI_GET_VIRTUALTABLESHAPEID      IMAGEHLP_SYMBOL_TYPE_INFO = 16
	TI_GET_VIRTUALBASEPOINTEROFFSET IMAGEHLP_SYMBOL_TYPE_INFO = 17
	TI_GET_CLASSPARENTID            IMAGEHLP_SYMBOL_TYPE_INFO = 18
	TI_GET_NESTED                   IMAGEHLP_SYMBOL_TYPE_INFO = 19
	TI_GET_SYMINDEX                 IMAGEHLP_SYMBOL_TYPE_INFO = 20
	TI_GET_LEXICALPARENT            IMAGEHLP_SYMBOL_TYPE_INFO = 21
	TI_GET_ADDRESS                  IMAGEHLP_SYMBOL_TYPE_INFO = 22
	TI_GET_THISADJUST               IMAGEHLP_SYMBOL_TYPE_INFO = 23
	TI_GET_UDTKIND                  IMAGEHLP_SYMBOL_TYPE_INFO = 24
	TI_IS_EQUIV_TO                  IMAGEHLP_SYMBOL_TYPE_INFO = 25
	TI_GET_CALLING_CONVENTION       IMAGEHLP_SYMBOL_TYPE_INFO = 26
	TI_IS_CLOSE_EQUIV_TO            IMAGEHLP_SYMBOL_TYPE_INFO = 27
	TI_GTIEX_REQS_VALID             IMAGEHLP_SYMBOL_TYPE_INFO = 28
	TI_GET_VIRTUALBASEOFFSET        IMAGEHLP_SYMBOL_TYPE_INFO = 29
	TI_GET_VIRTUALBASEDISPINDEX     IMAGEHLP_SYMBOL_TYPE_INFO = 30
	TI_GET_IS_REFERENCE             IMAGEHLP_SYMBOL_TYPE_INFO = 31
	TI_GET_INDIRECTVIRTUALBASECLASS IMAGEHLP_SYMBOL_TYPE_INFO = 32
	TI_GET_VIRTUALBASETABLETYPE     IMAGEHLP_SYMBOL_TYPE_INFO = 33
	TI_GET_OBJECTPOINTERTYPE        IMAGEHLP_SYMBOL_TYPE_INFO = 34
	IMAGEHLP_SYMBOL_TYPE_INFO_MAX   IMAGEHLP_SYMBOL_TYPE_INFO = 35
)

// enum
type IMAGEHLP_SF_TYPE int32

const (
	SfImage IMAGEHLP_SF_TYPE = 0
	SfDbg   IMAGEHLP_SF_TYPE = 1
	SfPdb   IMAGEHLP_SF_TYPE = 2
	SfMpd   IMAGEHLP_SF_TYPE = 3
	SfMax   IMAGEHLP_SF_TYPE = 4
)

// enum
type DUMP_TYPE int32

const (
	DUMP_TYPE_INVALID       DUMP_TYPE = -1
	DUMP_TYPE_UNKNOWN       DUMP_TYPE = 0
	DUMP_TYPE_FULL          DUMP_TYPE = 1
	DUMP_TYPE_SUMMARY       DUMP_TYPE = 2
	DUMP_TYPE_HEADER        DUMP_TYPE = 3
	DUMP_TYPE_TRIAGE        DUMP_TYPE = 4
	DUMP_TYPE_BITMAP_FULL   DUMP_TYPE = 5
	DUMP_TYPE_BITMAP_KERNEL DUMP_TYPE = 6
	DUMP_TYPE_AUTOMATIC     DUMP_TYPE = 7
)

// enum
type WHEA_ERROR_SOURCE_TYPE int32

const (
	WheaErrSrcTypeMCE          WHEA_ERROR_SOURCE_TYPE = 0
	WheaErrSrcTypeCMC          WHEA_ERROR_SOURCE_TYPE = 1
	WheaErrSrcTypeCPE          WHEA_ERROR_SOURCE_TYPE = 2
	WheaErrSrcTypeNMI          WHEA_ERROR_SOURCE_TYPE = 3
	WheaErrSrcTypePCIe         WHEA_ERROR_SOURCE_TYPE = 4
	WheaErrSrcTypeGeneric      WHEA_ERROR_SOURCE_TYPE = 5
	WheaErrSrcTypeINIT         WHEA_ERROR_SOURCE_TYPE = 6
	WheaErrSrcTypeBOOT         WHEA_ERROR_SOURCE_TYPE = 7
	WheaErrSrcTypeSCIGeneric   WHEA_ERROR_SOURCE_TYPE = 8
	WheaErrSrcTypeIPFMCA       WHEA_ERROR_SOURCE_TYPE = 9
	WheaErrSrcTypeIPFCMC       WHEA_ERROR_SOURCE_TYPE = 10
	WheaErrSrcTypeIPFCPE       WHEA_ERROR_SOURCE_TYPE = 11
	WheaErrSrcTypeGenericV2    WHEA_ERROR_SOURCE_TYPE = 12
	WheaErrSrcTypeSCIGenericV2 WHEA_ERROR_SOURCE_TYPE = 13
	WheaErrSrcTypeBMC          WHEA_ERROR_SOURCE_TYPE = 14
	WheaErrSrcTypePMEM         WHEA_ERROR_SOURCE_TYPE = 15
	WheaErrSrcTypeDeviceDriver WHEA_ERROR_SOURCE_TYPE = 16
	WheaErrSrcTypeSea          WHEA_ERROR_SOURCE_TYPE = 17
	WheaErrSrcTypeSei          WHEA_ERROR_SOURCE_TYPE = 18
	WheaErrSrcTypeMax          WHEA_ERROR_SOURCE_TYPE = 19
)

// enum
type WHEA_ERROR_SOURCE_STATE int32

const (
	WheaErrSrcStateStopped       WHEA_ERROR_SOURCE_STATE = 1
	WheaErrSrcStateStarted       WHEA_ERROR_SOURCE_STATE = 2
	WheaErrSrcStateRemoved       WHEA_ERROR_SOURCE_STATE = 3
	WheaErrSrcStateRemovePending WHEA_ERROR_SOURCE_STATE = 4
)

// enum
type IPMI_OS_SEL_RECORD_TYPE int32

const (
	IpmiOsSelRecordTypeWhea             IPMI_OS_SEL_RECORD_TYPE = 0
	IpmiOsSelRecordTypeOther            IPMI_OS_SEL_RECORD_TYPE = 1
	IpmiOsSelRecordTypeWheaErrorXpfMca  IPMI_OS_SEL_RECORD_TYPE = 2
	IpmiOsSelRecordTypeWheaErrorPci     IPMI_OS_SEL_RECORD_TYPE = 3
	IpmiOsSelRecordTypeWheaErrorNmi     IPMI_OS_SEL_RECORD_TYPE = 4
	IpmiOsSelRecordTypeWheaErrorOther   IPMI_OS_SEL_RECORD_TYPE = 5
	IpmiOsSelRecordTypeRaw              IPMI_OS_SEL_RECORD_TYPE = 6
	IpmiOsSelRecordTypeDriver           IPMI_OS_SEL_RECORD_TYPE = 7
	IpmiOsSelRecordTypeBugcheckRecovery IPMI_OS_SEL_RECORD_TYPE = 8
	IpmiOsSelRecordTypeBugcheckData     IPMI_OS_SEL_RECORD_TYPE = 9
	IpmiOsSelRecordTypeMax              IPMI_OS_SEL_RECORD_TYPE = 10
)

// enum
// flags
type DBGPROP_ATTRIB_FLAGS int32

const (
	DBGPROP_ATTRIB_NO_ATTRIB              DBGPROP_ATTRIB_FLAGS = 0
	DBGPROP_ATTRIB_VALUE_IS_INVALID       DBGPROP_ATTRIB_FLAGS = 8
	DBGPROP_ATTRIB_VALUE_IS_EXPANDABLE    DBGPROP_ATTRIB_FLAGS = 16
	DBGPROP_ATTRIB_VALUE_IS_FAKE          DBGPROP_ATTRIB_FLAGS = 32
	DBGPROP_ATTRIB_VALUE_IS_METHOD        DBGPROP_ATTRIB_FLAGS = 256
	DBGPROP_ATTRIB_VALUE_IS_EVENT         DBGPROP_ATTRIB_FLAGS = 512
	DBGPROP_ATTRIB_VALUE_IS_RAW_STRING    DBGPROP_ATTRIB_FLAGS = 1024
	DBGPROP_ATTRIB_VALUE_READONLY         DBGPROP_ATTRIB_FLAGS = 2048
	DBGPROP_ATTRIB_ACCESS_PUBLIC          DBGPROP_ATTRIB_FLAGS = 4096
	DBGPROP_ATTRIB_ACCESS_PRIVATE         DBGPROP_ATTRIB_FLAGS = 8192
	DBGPROP_ATTRIB_ACCESS_PROTECTED       DBGPROP_ATTRIB_FLAGS = 16384
	DBGPROP_ATTRIB_ACCESS_FINAL           DBGPROP_ATTRIB_FLAGS = 32768
	DBGPROP_ATTRIB_STORAGE_GLOBAL         DBGPROP_ATTRIB_FLAGS = 65536
	DBGPROP_ATTRIB_STORAGE_STATIC         DBGPROP_ATTRIB_FLAGS = 131072
	DBGPROP_ATTRIB_STORAGE_FIELD          DBGPROP_ATTRIB_FLAGS = 262144
	DBGPROP_ATTRIB_STORAGE_VIRTUAL        DBGPROP_ATTRIB_FLAGS = 524288
	DBGPROP_ATTRIB_TYPE_IS_CONSTANT       DBGPROP_ATTRIB_FLAGS = 1048576
	DBGPROP_ATTRIB_TYPE_IS_SYNCHRONIZED   DBGPROP_ATTRIB_FLAGS = 2097152
	DBGPROP_ATTRIB_TYPE_IS_VOLATILE       DBGPROP_ATTRIB_FLAGS = 4194304
	DBGPROP_ATTRIB_HAS_EXTENDED_ATTRIBS   DBGPROP_ATTRIB_FLAGS = 8388608
	DBGPROP_ATTRIB_FRAME_INTRYBLOCK       DBGPROP_ATTRIB_FLAGS = 16777216
	DBGPROP_ATTRIB_FRAME_INCATCHBLOCK     DBGPROP_ATTRIB_FLAGS = 33554432
	DBGPROP_ATTRIB_FRAME_INFINALLYBLOCK   DBGPROP_ATTRIB_FLAGS = 67108864
	DBGPROP_ATTRIB_VALUE_IS_RETURN_VALUE  DBGPROP_ATTRIB_FLAGS = 134217728
	DBGPROP_ATTRIB_VALUE_PENDING_MUTATION DBGPROP_ATTRIB_FLAGS = 268435456
)

// enum
// flags
type DBGPROP_INFO int32

const (
	DBGPROP_INFO_NAME         DBGPROP_INFO = 1
	DBGPROP_INFO_TYPE         DBGPROP_INFO = 2
	DBGPROP_INFO_VALUE        DBGPROP_INFO = 4
	DBGPROP_INFO_FULLNAME     DBGPROP_INFO = 32
	DBGPROP_INFO_ATTRIBUTES   DBGPROP_INFO = 8
	DBGPROP_INFO_DEBUGPROP    DBGPROP_INFO = 16
	DBGPROP_INFO_BEAUTIFY     DBGPROP_INFO = 33554432
	DBGPROP_INFO_CALLTOSTRING DBGPROP_INFO = 67108864
	DBGPROP_INFO_AUTOEXPAND   DBGPROP_INFO = 134217728
)

// enum
type OBJECT_ATTRIB_FLAGS int32

const (
	OBJECT_ATTRIB_NO_ATTRIB            OBJECT_ATTRIB_FLAGS = 0
	OBJECT_ATTRIB_NO_NAME              OBJECT_ATTRIB_FLAGS = 1
	OBJECT_ATTRIB_NO_TYPE              OBJECT_ATTRIB_FLAGS = 2
	OBJECT_ATTRIB_NO_VALUE             OBJECT_ATTRIB_FLAGS = 4
	OBJECT_ATTRIB_VALUE_IS_INVALID     OBJECT_ATTRIB_FLAGS = 8
	OBJECT_ATTRIB_VALUE_IS_OBJECT      OBJECT_ATTRIB_FLAGS = 16
	OBJECT_ATTRIB_VALUE_IS_ENUM        OBJECT_ATTRIB_FLAGS = 32
	OBJECT_ATTRIB_VALUE_IS_CUSTOM      OBJECT_ATTRIB_FLAGS = 64
	OBJECT_ATTRIB_OBJECT_IS_EXPANDABLE OBJECT_ATTRIB_FLAGS = 112
	OBJECT_ATTRIB_VALUE_HAS_CODE       OBJECT_ATTRIB_FLAGS = 128
	OBJECT_ATTRIB_TYPE_IS_OBJECT       OBJECT_ATTRIB_FLAGS = 256
	OBJECT_ATTRIB_TYPE_HAS_CODE        OBJECT_ATTRIB_FLAGS = 512
	OBJECT_ATTRIB_TYPE_IS_EXPANDABLE   OBJECT_ATTRIB_FLAGS = 256
	OBJECT_ATTRIB_SLOT_IS_CATEGORY     OBJECT_ATTRIB_FLAGS = 1024
	OBJECT_ATTRIB_VALUE_READONLY       OBJECT_ATTRIB_FLAGS = 2048
	OBJECT_ATTRIB_ACCESS_PUBLIC        OBJECT_ATTRIB_FLAGS = 4096
	OBJECT_ATTRIB_ACCESS_PRIVATE       OBJECT_ATTRIB_FLAGS = 8192
	OBJECT_ATTRIB_ACCESS_PROTECTED     OBJECT_ATTRIB_FLAGS = 16384
	OBJECT_ATTRIB_ACCESS_FINAL         OBJECT_ATTRIB_FLAGS = 32768
	OBJECT_ATTRIB_STORAGE_GLOBAL       OBJECT_ATTRIB_FLAGS = 65536
	OBJECT_ATTRIB_STORAGE_STATIC       OBJECT_ATTRIB_FLAGS = 131072
	OBJECT_ATTRIB_STORAGE_FIELD        OBJECT_ATTRIB_FLAGS = 262144
	OBJECT_ATTRIB_STORAGE_VIRTUAL      OBJECT_ATTRIB_FLAGS = 524288
	OBJECT_ATTRIB_TYPE_IS_CONSTANT     OBJECT_ATTRIB_FLAGS = 1048576
	OBJECT_ATTRIB_TYPE_IS_SYNCHRONIZED OBJECT_ATTRIB_FLAGS = 2097152
	OBJECT_ATTRIB_TYPE_IS_VOLATILE     OBJECT_ATTRIB_FLAGS = 4194304
	OBJECT_ATTRIB_HAS_EXTENDED_ATTRIBS OBJECT_ATTRIB_FLAGS = 8388608
	OBJECT_ATTRIB_IS_CLASS             OBJECT_ATTRIB_FLAGS = 16777216
	OBJECT_ATTRIB_IS_FUNCTION          OBJECT_ATTRIB_FLAGS = 33554432
	OBJECT_ATTRIB_IS_VARIABLE          OBJECT_ATTRIB_FLAGS = 67108864
	OBJECT_ATTRIB_IS_PROPERTY          OBJECT_ATTRIB_FLAGS = 134217728
	OBJECT_ATTRIB_IS_MACRO             OBJECT_ATTRIB_FLAGS = 268435456
	OBJECT_ATTRIB_IS_TYPE              OBJECT_ATTRIB_FLAGS = 536870912
	OBJECT_ATTRIB_IS_INHERITED         OBJECT_ATTRIB_FLAGS = 1073741824
	OBJECT_ATTRIB_IS_INTERFACE         OBJECT_ATTRIB_FLAGS = -2147483648
)

// enum
type PROP_INFO_FLAGS int32

const (
	PROP_INFO_NAME       PROP_INFO_FLAGS = 1
	PROP_INFO_TYPE       PROP_INFO_FLAGS = 2
	PROP_INFO_VALUE      PROP_INFO_FLAGS = 4
	PROP_INFO_FULLNAME   PROP_INFO_FLAGS = 32
	PROP_INFO_ATTRIBUTES PROP_INFO_FLAGS = 8
	PROP_INFO_DEBUGPROP  PROP_INFO_FLAGS = 16
	PROP_INFO_AUTOEXPAND PROP_INFO_FLAGS = 134217728
)

// enum
type EX_PROP_INFO_FLAGS int32

const (
	EX_PROP_INFO_ID           EX_PROP_INFO_FLAGS = 256
	EX_PROP_INFO_NTYPE        EX_PROP_INFO_FLAGS = 512
	EX_PROP_INFO_NVALUE       EX_PROP_INFO_FLAGS = 1024
	EX_PROP_INFO_LOCKBYTES    EX_PROP_INFO_FLAGS = 2048
	EX_PROP_INFO_DEBUGEXTPROP EX_PROP_INFO_FLAGS = 4096
)

// structs

type EXCEPTION_DEBUG_INFO struct {
	ExceptionRecord EXCEPTION_RECORD
	DwFirstChance   uint32
}

type CREATE_THREAD_DEBUG_INFO struct {
	HThread           HANDLE
	LpThreadLocalBase unsafe.Pointer
	LpStartAddress    LPTHREAD_START_ROUTINE
}

type CREATE_PROCESS_DEBUG_INFO struct {
	HFile                 HANDLE
	HProcess              HANDLE
	HThread               HANDLE
	LpBaseOfImage         unsafe.Pointer
	DwDebugInfoFileOffset uint32
	NDebugInfoSize        uint32
	LpThreadLocalBase     unsafe.Pointer
	LpStartAddress        LPTHREAD_START_ROUTINE
	LpImageName           unsafe.Pointer
	FUnicode              uint16
}

type EXIT_THREAD_DEBUG_INFO struct {
	DwExitCode uint32
}

type EXIT_PROCESS_DEBUG_INFO struct {
	DwExitCode uint32
}

type LOAD_DLL_DEBUG_INFO struct {
	HFile                 HANDLE
	LpBaseOfDll           unsafe.Pointer
	DwDebugInfoFileOffset uint32
	NDebugInfoSize        uint32
	LpImageName           unsafe.Pointer
	FUnicode              uint16
}

type UNLOAD_DLL_DEBUG_INFO struct {
	LpBaseOfDll unsafe.Pointer
}

type OUTPUT_DEBUG_STRING_INFO struct {
	LpDebugStringData  PSTR
	FUnicode           uint16
	NDebugStringLength uint16
}

type RIP_INFO struct {
	DwError uint32
	DwType  RIP_INFO_TYPE
}

type DEBUG_EVENT_U struct {
	Data [20]uint64
}

func (this *DEBUG_EVENT_U) Exception() *EXCEPTION_DEBUG_INFO {
	return (*EXCEPTION_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) ExceptionVal() EXCEPTION_DEBUG_INFO {
	return *(*EXCEPTION_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) CreateThread() *CREATE_THREAD_DEBUG_INFO {
	return (*CREATE_THREAD_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) CreateThreadVal() CREATE_THREAD_DEBUG_INFO {
	return *(*CREATE_THREAD_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) CreateProcessInfo() *CREATE_PROCESS_DEBUG_INFO {
	return (*CREATE_PROCESS_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) CreateProcessInfoVal() CREATE_PROCESS_DEBUG_INFO {
	return *(*CREATE_PROCESS_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) ExitThread() *EXIT_THREAD_DEBUG_INFO {
	return (*EXIT_THREAD_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) ExitThreadVal() EXIT_THREAD_DEBUG_INFO {
	return *(*EXIT_THREAD_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) ExitProcess() *EXIT_PROCESS_DEBUG_INFO {
	return (*EXIT_PROCESS_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) ExitProcessVal() EXIT_PROCESS_DEBUG_INFO {
	return *(*EXIT_PROCESS_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) LoadDll() *LOAD_DLL_DEBUG_INFO {
	return (*LOAD_DLL_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) LoadDllVal() LOAD_DLL_DEBUG_INFO {
	return *(*LOAD_DLL_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) UnloadDll() *UNLOAD_DLL_DEBUG_INFO {
	return (*UNLOAD_DLL_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) UnloadDllVal() UNLOAD_DLL_DEBUG_INFO {
	return *(*UNLOAD_DLL_DEBUG_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) DebugString() *OUTPUT_DEBUG_STRING_INFO {
	return (*OUTPUT_DEBUG_STRING_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) DebugStringVal() OUTPUT_DEBUG_STRING_INFO {
	return *(*OUTPUT_DEBUG_STRING_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) RipInfo() *RIP_INFO {
	return (*RIP_INFO)(unsafe.Pointer(this))
}

func (this *DEBUG_EVENT_U) RipInfoVal() RIP_INFO {
	return *(*RIP_INFO)(unsafe.Pointer(this))
}

type DEBUG_EVENT struct {
	DwDebugEventCode DEBUG_EVENT_CODE
	DwProcessId      uint32
	DwThreadId       uint32
	U                DEBUG_EVENT_U
}

type APC_CALLBACK_DATA struct {
	Parameter     uintptr
	ContextRecord *CONTEXT
	Reserved0     uintptr
	Reserved1     uintptr
}

type XSAVE_FORMAT struct {
	ControlWord    uint16
	StatusWord     uint16
	TagWord        byte
	Reserved1      byte
	ErrorOpcode    uint16
	ErrorOffset    uint32
	ErrorSelector  uint16
	Reserved2      uint16
	DataOffset     uint32
	DataSelector   uint16
	Reserved3      uint16
	MxCsr          uint32
	MxCsr_Mask     uint32
	FloatRegisters [8]M128A
	XmmRegisters   [16]M128A
	Reserved4      [96]byte
}

type XSTATE_CONTEXT struct {
	Mask      uint64
	Length    uint32
	Reserved1 uint32
	Area      *XSAVE_AREA
	Buffer    unsafe.Pointer
}

type CONTEXT_Anonymous_Anonymous struct {
	X0  uint64
	X1  uint64
	X2  uint64
	X3  uint64
	X4  uint64
	X5  uint64
	X6  uint64
	X7  uint64
	X8  uint64
	X9  uint64
	X10 uint64
	X11 uint64
	X12 uint64
	X13 uint64
	X14 uint64
	X15 uint64
	X16 uint64
	X17 uint64
	X18 uint64
	X19 uint64
	X20 uint64
	X21 uint64
	X22 uint64
	X23 uint64
	X24 uint64
	X25 uint64
	X26 uint64
	X27 uint64
	X28 uint64
	Fp  uint64
	Lr  uint64
}

type CONTEXT_Anonymous struct {
	CONTEXT_Anonymous_Anonymous
}

func (this *CONTEXT_Anonymous) Anonymous() *CONTEXT_Anonymous_Anonymous {
	return (*CONTEXT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *CONTEXT_Anonymous) AnonymousVal() CONTEXT_Anonymous_Anonymous {
	return *(*CONTEXT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *CONTEXT_Anonymous) X() *[31]uint64 {
	return (*[31]uint64)(unsafe.Pointer(this))
}

func (this *CONTEXT_Anonymous) XVal() [31]uint64 {
	return *(*[31]uint64)(unsafe.Pointer(this))
}

type CONTEXT struct {
	P1Home       uint64
	P2Home       uint64
	P3Home       uint64
	P4Home       uint64
	P5Home       uint64
	P6Home       uint64
	ContextFlags CONTEXT_FLAGS
	MxCsr        uint32
	SegCs        uint16
	SegDs        uint16
	SegEs        uint16
	SegFs        uint16
	SegGs        uint16
	SegSs        uint16
	EFlags       uint32
	Dr0          uint64
	Dr1          uint64
	Dr2          uint64
	Dr3          uint64
	Dr6          uint64
	Dr7          uint64
	Rax          uint64
	Rcx          uint64
	Rdx          uint64
	Rbx          uint64
	Rsp          uint64
	Rbp          uint64
	Rsi          uint64
	Rdi          uint64
	R8           uint64
	R9           uint64
	R10          uint64
	R11          uint64
	R12          uint64
	R13          uint64
	R14          uint64
	R15          uint64
	Rip          uint64
	CONTEXT_Anonymous
	VectorRegister       [26]M128A
	VectorControl        uint64
	DebugControl         uint64
	LastBranchToRip      uint64
	LastBranchFromRip    uint64
	LastExceptionToRip   uint64
	LastExceptionFromRip uint64
}

type DISPATCHER_CONTEXT struct {
	ControlPc        uint64
	ImageBase        uint64
	FunctionEntry    *IMAGE_RUNTIME_FUNCTION_ENTRY
	EstablisherFrame uint64
	TargetIp         uint64
	ContextRecord    *CONTEXT
	LanguageHandler  EXCEPTION_ROUTINE
	HandlerData      unsafe.Pointer
	HistoryTable     *UNWIND_HISTORY_TABLE
	ScopeIndex       uint32
	Fill0            uint32
}

type KNONVOLATILE_CONTEXT_POINTERS_Anonymous1_Anonymous struct {
	Xmm0  *M128A
	Xmm1  *M128A
	Xmm2  *M128A
	Xmm3  *M128A
	Xmm4  *M128A
	Xmm5  *M128A
	Xmm6  *M128A
	Xmm7  *M128A
	Xmm8  *M128A
	Xmm9  *M128A
	Xmm10 *M128A
	Xmm11 *M128A
	Xmm12 *M128A
	Xmm13 *M128A
	Xmm14 *M128A
	Xmm15 *M128A
}

type KNONVOLATILE_CONTEXT_POINTERS_Anonymous1 struct {
	KNONVOLATILE_CONTEXT_POINTERS_Anonymous1_Anonymous
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous1) FloatingContext() *[16]*M128A {
	return (*[16]*M128A)(unsafe.Pointer(this))
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous1) FloatingContextVal() [16]*M128A {
	return *(*[16]*M128A)(unsafe.Pointer(this))
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous1) Anonymous() *KNONVOLATILE_CONTEXT_POINTERS_Anonymous1_Anonymous {
	return (*KNONVOLATILE_CONTEXT_POINTERS_Anonymous1_Anonymous)(unsafe.Pointer(this))
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous1) AnonymousVal() KNONVOLATILE_CONTEXT_POINTERS_Anonymous1_Anonymous {
	return *(*KNONVOLATILE_CONTEXT_POINTERS_Anonymous1_Anonymous)(unsafe.Pointer(this))
}

type KNONVOLATILE_CONTEXT_POINTERS_Anonymous2_Anonymous struct {
	Rax *uint64
	Rcx *uint64
	Rdx *uint64
	Rbx *uint64
	Rsp *uint64
	Rbp *uint64
	Rsi *uint64
	Rdi *uint64
	R8  *uint64
	R9  *uint64
	R10 *uint64
	R11 *uint64
	R12 *uint64
	R13 *uint64
	R14 *uint64
	R15 *uint64
}

type KNONVOLATILE_CONTEXT_POINTERS_Anonymous2 struct {
	KNONVOLATILE_CONTEXT_POINTERS_Anonymous2_Anonymous
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous2) IntegerContext() *[16]*uint64 {
	return (*[16]*uint64)(unsafe.Pointer(this))
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous2) IntegerContextVal() [16]*uint64 {
	return *(*[16]*uint64)(unsafe.Pointer(this))
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous2) Anonymous() *KNONVOLATILE_CONTEXT_POINTERS_Anonymous2_Anonymous {
	return (*KNONVOLATILE_CONTEXT_POINTERS_Anonymous2_Anonymous)(unsafe.Pointer(this))
}

func (this *KNONVOLATILE_CONTEXT_POINTERS_Anonymous2) AnonymousVal() KNONVOLATILE_CONTEXT_POINTERS_Anonymous2_Anonymous {
	return *(*KNONVOLATILE_CONTEXT_POINTERS_Anonymous2_Anonymous)(unsafe.Pointer(this))
}

type KNONVOLATILE_CONTEXT_POINTERS struct {
	KNONVOLATILE_CONTEXT_POINTERS_Anonymous1
	KNONVOLATILE_CONTEXT_POINTERS_Anonymous2
}

type UNWIND_HISTORY_TABLE_ENTRY struct {
	ImageBase     uintptr
	FunctionEntry *IMAGE_RUNTIME_FUNCTION_ENTRY
}

type UNWIND_HISTORY_TABLE struct {
	Count       uint32
	LocalHint   byte
	GlobalHint  byte
	Search      byte
	Once        byte
	LowAddress  uintptr
	HighAddress uintptr
	Entry       [12]UNWIND_HISTORY_TABLE_ENTRY
}

type MINIDUMP_EXCEPTION_INFORMATION struct {
	ThreadId          uint32
	ExceptionPointers *EXCEPTION_POINTERS
	ClientPointers    BOOL
}

type MINIDUMP_USER_STREAM struct {
	Type       uint32
	BufferSize uint32
	Buffer     unsafe.Pointer
}

type MINIDUMP_USER_STREAM_INFORMATION struct {
	UserStreamCount uint32
	UserStreamArray *MINIDUMP_USER_STREAM
}

type MINIDUMP_CALLBACK_INFORMATION struct {
	CallbackRoutine MINIDUMP_CALLBACK_ROUTINE
	CallbackParam   unsafe.Pointer
}

type LOADED_IMAGE struct {
	ModuleName       PSTR
	HFile            HANDLE
	MappedAddress    *byte
	FileHeader       *IMAGE_NT_HEADERS64
	LastRvaSection   *IMAGE_SECTION_HEADER
	NumberOfSections uint32
	Sections         *IMAGE_SECTION_HEADER
	Characteristics  IMAGE_FILE_CHARACTERISTICS2
	FSystemImage     BOOLEAN
	FDOSImage        BOOLEAN
	FReadOnly        BOOLEAN
	Version          byte
	Links            LIST_ENTRY
	SizeOfImage      uint32
}

type M128A struct {
	Low  uint64
	High int64
}

type XSAVE_AREA_HEADER struct {
	Mask           uint64
	CompactionMask uint64
	Reserved2      [6]uint64
}

type XSAVE_AREA struct {
	LegacyState XSAVE_FORMAT
	Header      XSAVE_AREA_HEADER
}

type ARM64_NT_NEON128_Anonymous struct {
	Low  uint64
	High int64
}

type ARM64_NT_NEON128 struct {
	ARM64_NT_NEON128_Anonymous
}

func (this *ARM64_NT_NEON128) Anonymous() *ARM64_NT_NEON128_Anonymous {
	return (*ARM64_NT_NEON128_Anonymous)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) AnonymousVal() ARM64_NT_NEON128_Anonymous {
	return *(*ARM64_NT_NEON128_Anonymous)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) D() *[2]float64 {
	return (*[2]float64)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) DVal() [2]float64 {
	return *(*[2]float64)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) S() *[4]float32 {
	return (*[4]float32)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) SVal() [4]float32 {
	return *(*[4]float32)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) H() *[8]uint16 {
	return (*[8]uint16)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) HVal() [8]uint16 {
	return *(*[8]uint16)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) B() *[16]byte {
	return (*[16]byte)(unsafe.Pointer(this))
}

func (this *ARM64_NT_NEON128) BVal() [16]byte {
	return *(*[16]byte)(unsafe.Pointer(this))
}

type ARM64_NT_CONTEXT_Anonymous_Anonymous struct {
	X0  uint64
	X1  uint64
	X2  uint64
	X3  uint64
	X4  uint64
	X5  uint64
	X6  uint64
	X7  uint64
	X8  uint64
	X9  uint64
	X10 uint64
	X11 uint64
	X12 uint64
	X13 uint64
	X14 uint64
	X15 uint64
	X16 uint64
	X17 uint64
	X18 uint64
	X19 uint64
	X20 uint64
	X21 uint64
	X22 uint64
	X23 uint64
	X24 uint64
	X25 uint64
	X26 uint64
	X27 uint64
	X28 uint64
	Fp  uint64
	Lr  uint64
}

type ARM64_NT_CONTEXT_Anonymous struct {
	ARM64_NT_CONTEXT_Anonymous_Anonymous
}

func (this *ARM64_NT_CONTEXT_Anonymous) Anonymous() *ARM64_NT_CONTEXT_Anonymous_Anonymous {
	return (*ARM64_NT_CONTEXT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *ARM64_NT_CONTEXT_Anonymous) AnonymousVal() ARM64_NT_CONTEXT_Anonymous_Anonymous {
	return *(*ARM64_NT_CONTEXT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *ARM64_NT_CONTEXT_Anonymous) X() *[31]uint64 {
	return (*[31]uint64)(unsafe.Pointer(this))
}

func (this *ARM64_NT_CONTEXT_Anonymous) XVal() [31]uint64 {
	return *(*[31]uint64)(unsafe.Pointer(this))
}

type ARM64_NT_CONTEXT struct {
	ContextFlags uint32
	Cpsr         uint32
	ARM64_NT_CONTEXT_Anonymous
	Sp   uint64
	Pc   uint64
	V    [32]ARM64_NT_NEON128
	Fpcr uint32
	Fpsr uint32
	Bcr  [8]uint32
	Bvr  [8]uint64
	Wcr  [2]uint32
	Wvr  [2]uint64
}

type LDT_ENTRY_HighWord_Bytes struct {
	BaseMid byte
	Flags1  byte
	Flags2  byte
	BaseHi  byte
}

type LDT_ENTRY_HighWord_Bits struct {
	Bitfield_ uint32
}

type LDT_ENTRY_HighWord struct {
	Data [1]uint32
}

func (this *LDT_ENTRY_HighWord) Bytes() *LDT_ENTRY_HighWord_Bytes {
	return (*LDT_ENTRY_HighWord_Bytes)(unsafe.Pointer(this))
}

func (this *LDT_ENTRY_HighWord) BytesVal() LDT_ENTRY_HighWord_Bytes {
	return *(*LDT_ENTRY_HighWord_Bytes)(unsafe.Pointer(this))
}

func (this *LDT_ENTRY_HighWord) Bits() *LDT_ENTRY_HighWord_Bits {
	return (*LDT_ENTRY_HighWord_Bits)(unsafe.Pointer(this))
}

func (this *LDT_ENTRY_HighWord) BitsVal() LDT_ENTRY_HighWord_Bits {
	return *(*LDT_ENTRY_HighWord_Bits)(unsafe.Pointer(this))
}

type LDT_ENTRY struct {
	LimitLow uint16
	BaseLow  uint16
	HighWord LDT_ENTRY_HighWord
}

type WOW64_FLOATING_SAVE_AREA struct {
	ControlWord   uint32
	StatusWord    uint32
	TagWord       uint32
	ErrorOffset   uint32
	ErrorSelector uint32
	DataOffset    uint32
	DataSelector  uint32
	RegisterArea  [80]byte
	Cr0NpxState   uint32
}

type WOW64_CONTEXT struct {
	ContextFlags      WOW64_CONTEXT_FLAGS
	Dr0               uint32
	Dr1               uint32
	Dr2               uint32
	Dr3               uint32
	Dr6               uint32
	Dr7               uint32
	FloatSave         WOW64_FLOATING_SAVE_AREA
	SegGs             uint32
	SegFs             uint32
	SegEs             uint32
	SegDs             uint32
	Edi               uint32
	Esi               uint32
	Ebx               uint32
	Edx               uint32
	Ecx               uint32
	Eax               uint32
	Ebp               uint32
	Eip               uint32
	SegCs             uint32
	EFlags            uint32
	Esp               uint32
	SegSs             uint32
	ExtendedRegisters [512]byte
}

type WOW64_LDT_ENTRY_HighWord_Bytes struct {
	BaseMid byte
	Flags1  byte
	Flags2  byte
	BaseHi  byte
}

type WOW64_LDT_ENTRY_HighWord_Bits struct {
	Bitfield_ uint32
}

type WOW64_LDT_ENTRY_HighWord struct {
	Data [1]uint32
}

func (this *WOW64_LDT_ENTRY_HighWord) Bytes() *WOW64_LDT_ENTRY_HighWord_Bytes {
	return (*WOW64_LDT_ENTRY_HighWord_Bytes)(unsafe.Pointer(this))
}

func (this *WOW64_LDT_ENTRY_HighWord) BytesVal() WOW64_LDT_ENTRY_HighWord_Bytes {
	return *(*WOW64_LDT_ENTRY_HighWord_Bytes)(unsafe.Pointer(this))
}

func (this *WOW64_LDT_ENTRY_HighWord) Bits() *WOW64_LDT_ENTRY_HighWord_Bits {
	return (*WOW64_LDT_ENTRY_HighWord_Bits)(unsafe.Pointer(this))
}

func (this *WOW64_LDT_ENTRY_HighWord) BitsVal() WOW64_LDT_ENTRY_HighWord_Bits {
	return *(*WOW64_LDT_ENTRY_HighWord_Bits)(unsafe.Pointer(this))
}

type WOW64_LDT_ENTRY struct {
	LimitLow uint16
	BaseLow  uint16
	HighWord WOW64_LDT_ENTRY_HighWord
}

type WOW64_DESCRIPTOR_TABLE_ENTRY struct {
	Selector   uint32
	Descriptor WOW64_LDT_ENTRY
}

type EXCEPTION_RECORD struct {
	ExceptionCode        NTSTATUS
	ExceptionFlags       uint32
	ExceptionRecord      *EXCEPTION_RECORD
	ExceptionAddress     unsafe.Pointer
	NumberParameters     uint32
	ExceptionInformation [15]uintptr
}

type EXCEPTION_RECORD32 struct {
	ExceptionCode        NTSTATUS
	ExceptionFlags       uint32
	ExceptionRecord      uint32
	ExceptionAddress     uint32
	NumberParameters     uint32
	ExceptionInformation [15]uint32
}

type EXCEPTION_RECORD64 struct {
	ExceptionCode        NTSTATUS
	ExceptionFlags       uint32
	ExceptionRecord      uint64
	ExceptionAddress     uint64
	NumberParameters     uint32
	UnusedAlignment__    uint32
	ExceptionInformation [15]uint64
}

type EXCEPTION_POINTERS struct {
	ExceptionRecord *EXCEPTION_RECORD
	ContextRecord   *CONTEXT
}

type XSTATE_FEATURE struct {
	Offset uint32
	Size   uint32
}

type XSTATE_CONFIGURATION_Anonymous_Anonymous struct {
	Bitfield_ uint32
}

type XSTATE_CONFIGURATION_Anonymous struct {
	XSTATE_CONFIGURATION_Anonymous_Anonymous
}

func (this *XSTATE_CONFIGURATION_Anonymous) ControlFlags() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *XSTATE_CONFIGURATION_Anonymous) ControlFlagsVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *XSTATE_CONFIGURATION_Anonymous) Anonymous() *XSTATE_CONFIGURATION_Anonymous_Anonymous {
	return (*XSTATE_CONFIGURATION_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *XSTATE_CONFIGURATION_Anonymous) AnonymousVal() XSTATE_CONFIGURATION_Anonymous_Anonymous {
	return *(*XSTATE_CONFIGURATION_Anonymous_Anonymous)(unsafe.Pointer(this))
}

type XSTATE_CONFIGURATION struct {
	EnabledFeatures         uint64
	EnabledVolatileFeatures uint64
	Size                    uint32
	XSTATE_CONFIGURATION_Anonymous
	Features                             [64]XSTATE_FEATURE
	EnabledSupervisorFeatures            uint64
	AlignedFeatures                      uint64
	AllFeatureSize                       uint32
	AllFeatures                          [64]uint32
	EnabledUserVisibleSupervisorFeatures uint64
	ExtendedFeatureDisableFeatures       uint64
	AllNonLargeFeatureSize               uint32
	Spare                                uint32
}

type IMAGE_FILE_HEADER struct {
	Machine              IMAGE_FILE_MACHINE
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      IMAGE_FILE_CHARACTERISTICS
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress uint32
	Size           uint32
}

type IMAGE_OPTIONAL_HEADER32 struct {
	Magic                       IMAGE_OPTIONAL_HEADER_MAGIC
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   IMAGE_SUBSYSTEM
	DllCharacteristics          IMAGE_DLL_CHARACTERISTICS
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_ROM_OPTIONAL_HEADER struct {
	Magic                   uint16
	MajorLinkerVersion      byte
	MinorLinkerVersion      byte
	SizeOfCode              uint32
	SizeOfInitializedData   uint32
	SizeOfUninitializedData uint32
	AddressOfEntryPoint     uint32
	BaseOfCode              uint32
	BaseOfData              uint32
	BaseOfBss               uint32
	GprMask                 uint32
	CprMask                 [4]uint32
	GpValue                 uint32
}

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       IMAGE_OPTIONAL_HEADER_MAGIC
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   IMAGE_SUBSYSTEM
	DllCharacteristics          IMAGE_DLL_CHARACTERISTICS
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_NT_HEADERS64 struct {
	Signature      uint32
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}

type IMAGE_NT_HEADERS32 struct {
	Signature      uint32
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER32
}

type IMAGE_ROM_HEADERS struct {
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_ROM_OPTIONAL_HEADER
}

type IMAGE_SECTION_HEADER_Misc struct {
	Data [1]uint32
}

func (this *IMAGE_SECTION_HEADER_Misc) PhysicalAddress() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_SECTION_HEADER_Misc) PhysicalAddressVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_SECTION_HEADER_Misc) VirtualSize() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_SECTION_HEADER_Misc) VirtualSizeVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type IMAGE_SECTION_HEADER struct {
	Name                 [8]byte
	Misc                 IMAGE_SECTION_HEADER_Misc
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLinenumbers uint32
	NumberOfRelocations  uint16
	NumberOfLinenumbers  uint16
	Characteristics      IMAGE_SECTION_CHARACTERISTICS
}

type IMAGE_LOAD_CONFIG_CODE_INTEGRITY struct {
	Flags         uint16
	Catalog       uint16
	CatalogOffset uint32
	Reserved      uint32
}

type IMAGE_LOAD_CONFIG_DIRECTORY32 struct {
	Size                                     uint32
	TimeDateStamp                            uint32
	MajorVersion                             uint16
	MinorVersion                             uint16
	GlobalFlagsClear                         uint32
	GlobalFlagsSet                           uint32
	CriticalSectionDefaultTimeout            uint32
	DeCommitFreeBlockThreshold               uint32
	DeCommitTotalFreeThreshold               uint32
	LockPrefixTable                          uint32
	MaximumAllocationSize                    uint32
	VirtualMemoryThreshold                   uint32
	ProcessHeapFlags                         uint32
	ProcessAffinityMask                      uint32
	CSDVersion                               uint16
	DependentLoadFlags                       uint16
	EditList                                 uint32
	SecurityCookie                           uint32
	SEHandlerTable                           uint32
	SEHandlerCount                           uint32
	GuardCFCheckFunctionPointer              uint32
	GuardCFDispatchFunctionPointer           uint32
	GuardCFFunctionTable                     uint32
	GuardCFFunctionCount                     uint32
	GuardFlags                               uint32
	CodeIntegrity                            IMAGE_LOAD_CONFIG_CODE_INTEGRITY
	GuardAddressTakenIatEntryTable           uint32
	GuardAddressTakenIatEntryCount           uint32
	GuardLongJumpTargetTable                 uint32
	GuardLongJumpTargetCount                 uint32
	DynamicValueRelocTable                   uint32
	CHPEMetadataPointer                      uint32
	GuardRFFailureRoutine                    uint32
	GuardRFFailureRoutineFunctionPointer     uint32
	DynamicValueRelocTableOffset             uint32
	DynamicValueRelocTableSection            uint16
	Reserved2                                uint16
	GuardRFVerifyStackPointerFunctionPointer uint32
	HotPatchTableOffset                      uint32
	Reserved3                                uint32
	EnclaveConfigurationPointer              uint32
	VolatileMetadataPointer                  uint32
	GuardEHContinuationTable                 uint32
	GuardEHContinuationCount                 uint32
	GuardXFGCheckFunctionPointer             uint32
	GuardXFGDispatchFunctionPointer          uint32
	GuardXFGTableDispatchFunctionPointer     uint32
	CastGuardOsDeterminedFailureMode         uint32
	GuardMemcpyFunctionPointer               uint32
}

type IMAGE_LOAD_CONFIG_DIRECTORY64 struct {
	Size                                     uint32
	TimeDateStamp                            uint32
	MajorVersion                             uint16
	MinorVersion                             uint16
	GlobalFlagsClear                         uint32
	GlobalFlagsSet                           uint32
	CriticalSectionDefaultTimeout            uint32
	DeCommitFreeBlockThreshold               uint64
	DeCommitTotalFreeThreshold               uint64
	LockPrefixTable                          uint64
	MaximumAllocationSize                    uint64
	VirtualMemoryThreshold                   uint64
	ProcessAffinityMask                      uint64
	ProcessHeapFlags                         uint32
	CSDVersion                               uint16
	DependentLoadFlags                       uint16
	EditList                                 uint64
	SecurityCookie                           uint64
	SEHandlerTable                           uint64
	SEHandlerCount                           uint64
	GuardCFCheckFunctionPointer              uint64
	GuardCFDispatchFunctionPointer           uint64
	GuardCFFunctionTable                     uint64
	GuardCFFunctionCount                     uint64
	GuardFlags                               uint32
	CodeIntegrity                            IMAGE_LOAD_CONFIG_CODE_INTEGRITY
	GuardAddressTakenIatEntryTable           uint64
	GuardAddressTakenIatEntryCount           uint64
	GuardLongJumpTargetTable                 uint64
	GuardLongJumpTargetCount                 uint64
	DynamicValueRelocTable                   uint64
	CHPEMetadataPointer                      uint64
	GuardRFFailureRoutine                    uint64
	GuardRFFailureRoutineFunctionPointer     uint64
	DynamicValueRelocTableOffset             uint32
	DynamicValueRelocTableSection            uint16
	Reserved2                                uint16
	GuardRFVerifyStackPointerFunctionPointer uint64
	HotPatchTableOffset                      uint32
	Reserved3                                uint32
	EnclaveConfigurationPointer              uint64
	VolatileMetadataPointer                  uint64
	GuardEHContinuationTable                 uint64
	GuardEHContinuationCount                 uint64
	GuardXFGCheckFunctionPointer             uint64
	GuardXFGDispatchFunctionPointer          uint64
	GuardXFGTableDispatchFunctionPointer     uint64
	CastGuardOsDeterminedFailureMode         uint64
	GuardMemcpyFunctionPointer               uint64
}

type IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous_Anonymous struct {
	Bitfield_ uint32
}

type IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous struct {
	IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous_Anonymous
}

func (this *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous) UnwindData() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous) UnwindDataVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous) Anonymous() *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous_Anonymous {
	return (*IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous) AnonymousVal() IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous_Anonymous {
	return *(*IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous_Anonymous)(unsafe.Pointer(this))
}

type IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY struct {
	BeginAddress uint32
	IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY_Anonymous
}

type IMAGE_RUNTIME_FUNCTION_ENTRY_Anonymous struct {
	Data [1]uint32
}

func (this *IMAGE_RUNTIME_FUNCTION_ENTRY_Anonymous) UnwindInfoAddress() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_RUNTIME_FUNCTION_ENTRY_Anonymous) UnwindInfoAddressVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_RUNTIME_FUNCTION_ENTRY_Anonymous) UnwindData() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_RUNTIME_FUNCTION_ENTRY_Anonymous) UnwindDataVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type IMAGE_RUNTIME_FUNCTION_ENTRY struct {
	BeginAddress uint32
	EndAddress   uint32
	IMAGE_RUNTIME_FUNCTION_ENTRY_Anonymous
}

type IMAGE_DEBUG_DIRECTORY struct {
	Characteristics  uint32
	TimeDateStamp    uint32
	MajorVersion     uint16
	MinorVersion     uint16
	Type             IMAGE_DEBUG_TYPE
	SizeOfData       uint32
	AddressOfRawData uint32
	PointerToRawData uint32
}

type IMAGE_COFF_SYMBOLS_HEADER struct {
	NumberOfSymbols      uint32
	LvaToFirstSymbol     uint32
	NumberOfLinenumbers  uint32
	LvaToFirstLinenumber uint32
	RvaToFirstByteOfCode uint32
	RvaToLastByteOfCode  uint32
	RvaToFirstByteOfData uint32
	RvaToLastByteOfData  uint32
}

type FPO_DATA struct {
	UlOffStart uint32
	CbProcSize uint32
	CdwLocals  uint32
	CdwParams  uint16
	Bitfield_  uint16
}

type IMAGE_FUNCTION_ENTRY struct {
	StartingAddress uint32
	EndingAddress   uint32
	EndOfPrologue   uint32
}

type IMAGE_FUNCTION_ENTRY64_Anonymous struct {
	Data [1]uint64
}

func (this *IMAGE_FUNCTION_ENTRY64_Anonymous) EndOfPrologue() *uint64 {
	return (*uint64)(unsafe.Pointer(this))
}

func (this *IMAGE_FUNCTION_ENTRY64_Anonymous) EndOfPrologueVal() uint64 {
	return *(*uint64)(unsafe.Pointer(this))
}

func (this *IMAGE_FUNCTION_ENTRY64_Anonymous) UnwindInfoAddress() *uint64 {
	return (*uint64)(unsafe.Pointer(this))
}

func (this *IMAGE_FUNCTION_ENTRY64_Anonymous) UnwindInfoAddressVal() uint64 {
	return *(*uint64)(unsafe.Pointer(this))
}

type IMAGE_FUNCTION_ENTRY64 struct {
	StartingAddress uint64
	EndingAddress   uint64
	IMAGE_FUNCTION_ENTRY64_Anonymous
}

type IMAGE_COR20_HEADER_Anonymous struct {
	Data [1]uint32
}

func (this *IMAGE_COR20_HEADER_Anonymous) EntryPointToken() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_COR20_HEADER_Anonymous) EntryPointTokenVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_COR20_HEADER_Anonymous) EntryPointRVA() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *IMAGE_COR20_HEADER_Anonymous) EntryPointRVAVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type IMAGE_COR20_HEADER struct {
	Cb                  uint32
	MajorRuntimeVersion uint16
	MinorRuntimeVersion uint16
	MetaData            IMAGE_DATA_DIRECTORY
	Flags               uint32
	IMAGE_COR20_HEADER_Anonymous
	Resources               IMAGE_DATA_DIRECTORY
	StrongNameSignature     IMAGE_DATA_DIRECTORY
	CodeManagerTable        IMAGE_DATA_DIRECTORY
	VTableFixups            IMAGE_DATA_DIRECTORY
	ExportAddressTableJumps IMAGE_DATA_DIRECTORY
	ManagedNativeHeader     IMAGE_DATA_DIRECTORY
}

type WAITCHAIN_NODE_INFO_Anonymous_LockObject struct {
	ObjectName [128]uint16
	Timeout    int64
	Alertable  BOOL
}

type WAITCHAIN_NODE_INFO_Anonymous_ThreadObject struct {
	ProcessId       uint32
	ThreadId        uint32
	WaitTime        uint32
	ContextSwitches uint32
}

type WAITCHAIN_NODE_INFO_Anonymous struct {
	Data [34]uint64
}

func (this *WAITCHAIN_NODE_INFO_Anonymous) LockObject() *WAITCHAIN_NODE_INFO_Anonymous_LockObject {
	return (*WAITCHAIN_NODE_INFO_Anonymous_LockObject)(unsafe.Pointer(this))
}

func (this *WAITCHAIN_NODE_INFO_Anonymous) LockObjectVal() WAITCHAIN_NODE_INFO_Anonymous_LockObject {
	return *(*WAITCHAIN_NODE_INFO_Anonymous_LockObject)(unsafe.Pointer(this))
}

func (this *WAITCHAIN_NODE_INFO_Anonymous) ThreadObject() *WAITCHAIN_NODE_INFO_Anonymous_ThreadObject {
	return (*WAITCHAIN_NODE_INFO_Anonymous_ThreadObject)(unsafe.Pointer(this))
}

func (this *WAITCHAIN_NODE_INFO_Anonymous) ThreadObjectVal() WAITCHAIN_NODE_INFO_Anonymous_ThreadObject {
	return *(*WAITCHAIN_NODE_INFO_Anonymous_ThreadObject)(unsafe.Pointer(this))
}

type WAITCHAIN_NODE_INFO struct {
	ObjectType   WCT_OBJECT_TYPE
	ObjectStatus WCT_OBJECT_STATUS
	WAITCHAIN_NODE_INFO_Anonymous
}

type MINIDUMP_LOCATION_DESCRIPTOR struct {
	DataSize uint32
	Rva      uint32
}

type MINIDUMP_LOCATION_DESCRIPTOR64 struct {
	DataSize uint64
	Rva      uint64
}

type MINIDUMP_MEMORY_DESCRIPTOR struct {
	StartOfMemoryRange uint64
	Memory             MINIDUMP_LOCATION_DESCRIPTOR
}

type MINIDUMP_MEMORY_DESCRIPTOR64 struct {
	StartOfMemoryRange uint64
	DataSize           uint64
}

type MINIDUMP_HEADER_Anonymous struct {
	Data [1]uint32
}

func (this *MINIDUMP_HEADER_Anonymous) Reserved() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_HEADER_Anonymous) ReservedVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_HEADER_Anonymous) TimeDateStamp() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_HEADER_Anonymous) TimeDateStampVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type MINIDUMP_HEADER struct {
	Signature          uint32
	Version            uint32
	NumberOfStreams    uint32
	StreamDirectoryRva uint32
	CheckSum           uint32
	MINIDUMP_HEADER_Anonymous
	Flags uint64
}

type MINIDUMP_DIRECTORY struct {
	StreamType uint32
	Location   MINIDUMP_LOCATION_DESCRIPTOR
}

type MINIDUMP_STRING struct {
	Length uint32
	Buffer [1]uint16
}

type CPU_INFORMATION_X86CpuInfo struct {
	VendorId               [3]uint32
	VersionInformation     uint32
	FeatureInformation     uint32
	AMDExtendedCpuFeatures uint32
}

type CPU_INFORMATION_OtherCpuInfo struct {
	ProcessorFeatures [2]uint64
}

type CPU_INFORMATION struct {
	Data [3]uint64
}

func (this *CPU_INFORMATION) X86CpuInfo() *CPU_INFORMATION_X86CpuInfo {
	return (*CPU_INFORMATION_X86CpuInfo)(unsafe.Pointer(this))
}

func (this *CPU_INFORMATION) X86CpuInfoVal() CPU_INFORMATION_X86CpuInfo {
	return *(*CPU_INFORMATION_X86CpuInfo)(unsafe.Pointer(this))
}

func (this *CPU_INFORMATION) OtherCpuInfo() *CPU_INFORMATION_OtherCpuInfo {
	return (*CPU_INFORMATION_OtherCpuInfo)(unsafe.Pointer(this))
}

func (this *CPU_INFORMATION) OtherCpuInfoVal() CPU_INFORMATION_OtherCpuInfo {
	return *(*CPU_INFORMATION_OtherCpuInfo)(unsafe.Pointer(this))
}

type MINIDUMP_SYSTEM_INFO_Anonymous1_Anonymous struct {
	NumberOfProcessors byte
	ProductType        byte
}

type MINIDUMP_SYSTEM_INFO_Anonymous1 struct {
	MINIDUMP_SYSTEM_INFO_Anonymous1_Anonymous
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous1) Reserved0() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous1) Reserved0Val() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous1) Anonymous() *MINIDUMP_SYSTEM_INFO_Anonymous1_Anonymous {
	return (*MINIDUMP_SYSTEM_INFO_Anonymous1_Anonymous)(unsafe.Pointer(this))
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous1) AnonymousVal() MINIDUMP_SYSTEM_INFO_Anonymous1_Anonymous {
	return *(*MINIDUMP_SYSTEM_INFO_Anonymous1_Anonymous)(unsafe.Pointer(this))
}

type MINIDUMP_SYSTEM_INFO_Anonymous2_Anonymous struct {
	SuiteMask uint16
	Reserved2 uint16
}

type MINIDUMP_SYSTEM_INFO_Anonymous2 struct {
	MINIDUMP_SYSTEM_INFO_Anonymous2_Anonymous
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous2) Reserved1() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous2) Reserved1Val() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous2) Anonymous() *MINIDUMP_SYSTEM_INFO_Anonymous2_Anonymous {
	return (*MINIDUMP_SYSTEM_INFO_Anonymous2_Anonymous)(unsafe.Pointer(this))
}

func (this *MINIDUMP_SYSTEM_INFO_Anonymous2) AnonymousVal() MINIDUMP_SYSTEM_INFO_Anonymous2_Anonymous {
	return *(*MINIDUMP_SYSTEM_INFO_Anonymous2_Anonymous)(unsafe.Pointer(this))
}

type MINIDUMP_SYSTEM_INFO struct {
	ProcessorArchitecture PROCESSOR_ARCHITECTURE
	ProcessorLevel        uint16
	ProcessorRevision     uint16
	MINIDUMP_SYSTEM_INFO_Anonymous1
	MajorVersion  uint32
	MinorVersion  uint32
	BuildNumber   uint32
	PlatformId    VER_PLATFORM
	CSDVersionRva uint32
	MINIDUMP_SYSTEM_INFO_Anonymous2
	Cpu CPU_INFORMATION
}

type MINIDUMP_THREAD struct {
	ThreadId      uint32
	SuspendCount  uint32
	PriorityClass uint32
	Priority      uint32
	Teb           uint64
	Stack         MINIDUMP_MEMORY_DESCRIPTOR
	ThreadContext MINIDUMP_LOCATION_DESCRIPTOR
}

type MINIDUMP_THREAD_LIST struct {
	NumberOfThreads uint32
	Threads         [1]MINIDUMP_THREAD
}

type MINIDUMP_THREAD_EX struct {
	ThreadId      uint32
	SuspendCount  uint32
	PriorityClass uint32
	Priority      uint32
	Teb           uint64
	Stack         MINIDUMP_MEMORY_DESCRIPTOR
	ThreadContext MINIDUMP_LOCATION_DESCRIPTOR
	BackingStore  MINIDUMP_MEMORY_DESCRIPTOR
}

type MINIDUMP_THREAD_EX_LIST struct {
	NumberOfThreads uint32
	Threads         [1]MINIDUMP_THREAD_EX
}

type MINIDUMP_EXCEPTION struct {
	ExceptionCode        uint32
	ExceptionFlags       uint32
	ExceptionRecord      uint64
	ExceptionAddress     uint64
	NumberParameters     uint32
	UnusedAlignment__    uint32
	ExceptionInformation [15]uint64
}

type MINIDUMP_EXCEPTION_STREAM struct {
	ThreadId        uint32
	Alignment__     uint32
	ExceptionRecord MINIDUMP_EXCEPTION
	ThreadContext   MINIDUMP_LOCATION_DESCRIPTOR
}

type MINIDUMP_MODULE struct {
	BaseOfImage   uint64
	SizeOfImage   uint32
	CheckSum      uint32
	TimeDateStamp uint32
	ModuleNameRva uint32
	VersionInfo   VS_FIXEDFILEINFO
	CvRecord      MINIDUMP_LOCATION_DESCRIPTOR
	MiscRecord    MINIDUMP_LOCATION_DESCRIPTOR
	Reserved0     uint64
	Reserved1     uint64
}

type MINIDUMP_MODULE_LIST struct {
	NumberOfModules uint32
	Modules         [1]MINIDUMP_MODULE
}

type MINIDUMP_MEMORY_LIST struct {
	NumberOfMemoryRanges uint32
	MemoryRanges         [1]MINIDUMP_MEMORY_DESCRIPTOR
}

type MINIDUMP_MEMORY64_LIST struct {
	NumberOfMemoryRanges uint64
	BaseRva              uint64
	MemoryRanges         [1]MINIDUMP_MEMORY_DESCRIPTOR64
}

type MINIDUMP_EXCEPTION_INFORMATION64 struct {
	ThreadId        uint32
	ExceptionRecord uint64
	ContextRecord   uint64
	ClientPointers  BOOL
}

type MINIDUMP_HANDLE_OBJECT_INFORMATION struct {
	NextInfoRva uint32
	InfoType    uint32
	SizeOfInfo  uint32
}

type MINIDUMP_HANDLE_DESCRIPTOR struct {
	Handle        uint64
	TypeNameRva   uint32
	ObjectNameRva uint32
	Attributes    uint32
	GrantedAccess uint32
	HandleCount   uint32
	PointerCount  uint32
}

type MINIDUMP_HANDLE_DESCRIPTOR_2 struct {
	Handle        uint64
	TypeNameRva   uint32
	ObjectNameRva uint32
	Attributes    uint32
	GrantedAccess uint32
	HandleCount   uint32
	PointerCount  uint32
	ObjectInfoRva uint32
	Reserved0     uint32
}

type MINIDUMP_HANDLE_DATA_STREAM struct {
	SizeOfHeader        uint32
	SizeOfDescriptor    uint32
	NumberOfDescriptors uint32
	Reserved            uint32
}

type MINIDUMP_HANDLE_OPERATION_LIST struct {
	SizeOfHeader    uint32
	SizeOfEntry     uint32
	NumberOfEntries uint32
	Reserved        uint32
}

type MINIDUMP_FUNCTION_TABLE_DESCRIPTOR struct {
	MinimumAddress uint64
	MaximumAddress uint64
	BaseAddress    uint64
	EntryCount     uint32
	SizeOfAlignPad uint32
}

type MINIDUMP_FUNCTION_TABLE_STREAM struct {
	SizeOfHeader           uint32
	SizeOfDescriptor       uint32
	SizeOfNativeDescriptor uint32
	SizeOfFunctionEntry    uint32
	NumberOfDescriptors    uint32
	SizeOfAlignPad         uint32
}

type MINIDUMP_UNLOADED_MODULE struct {
	BaseOfImage   uint64
	SizeOfImage   uint32
	CheckSum      uint32
	TimeDateStamp uint32
	ModuleNameRva uint32
}

type MINIDUMP_UNLOADED_MODULE_LIST struct {
	SizeOfHeader    uint32
	SizeOfEntry     uint32
	NumberOfEntries uint32
}

type XSTATE_CONFIG_FEATURE_MSC_INFO struct {
	SizeOfInfo      uint32
	ContextSize     uint32
	EnabledFeatures uint64
	Features        [64]XSTATE_FEATURE
}

type MINIDUMP_MISC_INFO struct {
	SizeOfInfo        uint32
	Flags1            MINIDUMP_MISC_INFO_FLAGS
	ProcessId         uint32
	ProcessCreateTime uint32
	ProcessUserTime   uint32
	ProcessKernelTime uint32
}

type MINIDUMP_MISC_INFO_2 struct {
	SizeOfInfo                uint32
	Flags1                    uint32
	ProcessId                 uint32
	ProcessCreateTime         uint32
	ProcessUserTime           uint32
	ProcessKernelTime         uint32
	ProcessorMaxMhz           uint32
	ProcessorCurrentMhz       uint32
	ProcessorMhzLimit         uint32
	ProcessorMaxIdleState     uint32
	ProcessorCurrentIdleState uint32
}

type MINIDUMP_MISC_INFO_3 struct {
	SizeOfInfo                uint32
	Flags1                    uint32
	ProcessId                 uint32
	ProcessCreateTime         uint32
	ProcessUserTime           uint32
	ProcessKernelTime         uint32
	ProcessorMaxMhz           uint32
	ProcessorCurrentMhz       uint32
	ProcessorMhzLimit         uint32
	ProcessorMaxIdleState     uint32
	ProcessorCurrentIdleState uint32
	ProcessIntegrityLevel     uint32
	ProcessExecuteFlags       uint32
	ProtectedProcess          uint32
	TimeZoneId                uint32
	TimeZone                  TIME_ZONE_INFORMATION
}

type MINIDUMP_MISC_INFO_4 struct {
	SizeOfInfo                uint32
	Flags1                    uint32
	ProcessId                 uint32
	ProcessCreateTime         uint32
	ProcessUserTime           uint32
	ProcessKernelTime         uint32
	ProcessorMaxMhz           uint32
	ProcessorCurrentMhz       uint32
	ProcessorMhzLimit         uint32
	ProcessorMaxIdleState     uint32
	ProcessorCurrentIdleState uint32
	ProcessIntegrityLevel     uint32
	ProcessExecuteFlags       uint32
	ProtectedProcess          uint32
	TimeZoneId                uint32
	TimeZone                  TIME_ZONE_INFORMATION
	BuildString               [260]uint16
	DbgBldStr                 [40]uint16
}

type MINIDUMP_MISC_INFO_5 struct {
	SizeOfInfo                uint32
	Flags1                    uint32
	ProcessId                 uint32
	ProcessCreateTime         uint32
	ProcessUserTime           uint32
	ProcessKernelTime         uint32
	ProcessorMaxMhz           uint32
	ProcessorCurrentMhz       uint32
	ProcessorMhzLimit         uint32
	ProcessorMaxIdleState     uint32
	ProcessorCurrentIdleState uint32
	ProcessIntegrityLevel     uint32
	ProcessExecuteFlags       uint32
	ProtectedProcess          uint32
	TimeZoneId                uint32
	TimeZone                  TIME_ZONE_INFORMATION
	BuildString               [260]uint16
	DbgBldStr                 [40]uint16
	XStateData                XSTATE_CONFIG_FEATURE_MSC_INFO
	ProcessCookie             uint32
}

type MINIDUMP_MEMORY_INFO struct {
	BaseAddress       uint64
	AllocationBase    uint64
	AllocationProtect uint32
	Alignment1__      uint32
	RegionSize        uint64
	State             VIRTUAL_ALLOCATION_TYPE
	Protect           uint32
	Type              uint32
	Alignment2__      uint32
}

type MINIDUMP_MEMORY_INFO_LIST struct {
	SizeOfHeader    uint32
	SizeOfEntry     uint32
	NumberOfEntries uint64
}

type MINIDUMP_THREAD_NAME struct {
	ThreadId        uint32
	RvaOfThreadName uint64
}

type MINIDUMP_THREAD_NAME_LIST struct {
	NumberOfThreadNames uint32
	ThreadNames         [1]MINIDUMP_THREAD_NAME
}

type MINIDUMP_THREAD_INFO struct {
	ThreadId     uint32
	DumpFlags    MINIDUMP_THREAD_INFO_DUMP_FLAGS
	DumpError    uint32
	ExitStatus   uint32
	CreateTime   uint64
	ExitTime     uint64
	KernelTime   uint64
	UserTime     uint64
	StartAddress uint64
	Affinity     uint64
}

type MINIDUMP_THREAD_INFO_LIST struct {
	SizeOfHeader    uint32
	SizeOfEntry     uint32
	NumberOfEntries uint32
}

type MINIDUMP_TOKEN_INFO_HEADER struct {
	TokenSize   uint32
	TokenId     uint32
	TokenHandle uint64
}

type MINIDUMP_TOKEN_INFO_LIST struct {
	TokenListSize     uint32
	TokenListEntries  uint32
	ListHeaderSize    uint32
	ElementHeaderSize uint32
}

type MINIDUMP_SYSTEM_BASIC_INFORMATION struct {
	TimerResolution              uint32
	PageSize                     uint32
	NumberOfPhysicalPages        uint32
	LowestPhysicalPageNumber     uint32
	HighestPhysicalPageNumber    uint32
	AllocationGranularity        uint32
	MinimumUserModeAddress       uint64
	MaximumUserModeAddress       uint64
	ActiveProcessorsAffinityMask uint64
	NumberOfProcessors           uint32
}

type MINIDUMP_SYSTEM_FILECACHE_INFORMATION struct {
	CurrentSize                           uint64
	PeakSize                              uint64
	PageFaultCount                        uint32
	MinimumWorkingSet                     uint64
	MaximumWorkingSet                     uint64
	CurrentSizeIncludingTransitionInPages uint64
	PeakSizeIncludingTransitionInPages    uint64
	TransitionRePurposeCount              uint32
	Flags                                 uint32
}

type MINIDUMP_SYSTEM_BASIC_PERFORMANCE_INFORMATION struct {
	AvailablePages uint64
	CommittedPages uint64
	CommitLimit    uint64
	PeakCommitment uint64
}

type MINIDUMP_SYSTEM_PERFORMANCE_INFORMATION struct {
	IdleProcessTime           uint64
	IoReadTransferCount       uint64
	IoWriteTransferCount      uint64
	IoOtherTransferCount      uint64
	IoReadOperationCount      uint32
	IoWriteOperationCount     uint32
	IoOtherOperationCount     uint32
	AvailablePages            uint32
	CommittedPages            uint32
	CommitLimit               uint32
	PeakCommitment            uint32
	PageFaultCount            uint32
	CopyOnWriteCount          uint32
	TransitionCount           uint32
	CacheTransitionCount      uint32
	DemandZeroCount           uint32
	PageReadCount             uint32
	PageReadIoCount           uint32
	CacheReadCount            uint32
	CacheIoCount              uint32
	DirtyPagesWriteCount      uint32
	DirtyWriteIoCount         uint32
	MappedPagesWriteCount     uint32
	MappedWriteIoCount        uint32
	PagedPoolPages            uint32
	NonPagedPoolPages         uint32
	PagedPoolAllocs           uint32
	PagedPoolFrees            uint32
	NonPagedPoolAllocs        uint32
	NonPagedPoolFrees         uint32
	FreeSystemPtes            uint32
	ResidentSystemCodePage    uint32
	TotalSystemDriverPages    uint32
	TotalSystemCodePages      uint32
	NonPagedPoolLookasideHits uint32
	PagedPoolLookasideHits    uint32
	AvailablePagedPoolPages   uint32
	ResidentSystemCachePage   uint32
	ResidentPagedPoolPage     uint32
	ResidentSystemDriverPage  uint32
	CcFastReadNoWait          uint32
	CcFastReadWait            uint32
	CcFastReadResourceMiss    uint32
	CcFastReadNotPossible     uint32
	CcFastMdlReadNoWait       uint32
	CcFastMdlReadWait         uint32
	CcFastMdlReadResourceMiss uint32
	CcFastMdlReadNotPossible  uint32
	CcMapDataNoWait           uint32
	CcMapDataWait             uint32
	CcMapDataNoWaitMiss       uint32
	CcMapDataWaitMiss         uint32
	CcPinMappedDataCount      uint32
	CcPinReadNoWait           uint32
	CcPinReadWait             uint32
	CcPinReadNoWaitMiss       uint32
	CcPinReadWaitMiss         uint32
	CcCopyReadNoWait          uint32
	CcCopyReadWait            uint32
	CcCopyReadNoWaitMiss      uint32
	CcCopyReadWaitMiss        uint32
	CcMdlReadNoWait           uint32
	CcMdlReadWait             uint32
	CcMdlReadNoWaitMiss       uint32
	CcMdlReadWaitMiss         uint32
	CcReadAheadIos            uint32
	CcLazyWriteIos            uint32
	CcLazyWritePages          uint32
	CcDataFlushes             uint32
	CcDataPages               uint32
	ContextSwitches           uint32
	FirstLevelTbFills         uint32
	SecondLevelTbFills        uint32
	SystemCalls               uint32
	CcTotalDirtyPages         uint64
	CcDirtyPageThreshold      uint64
	ResidentAvailablePages    int64
	SharedCommittedPages      uint64
}

type MINIDUMP_SYSTEM_MEMORY_INFO_1 struct {
	Revision      uint16
	Flags         uint16
	BasicInfo     MINIDUMP_SYSTEM_BASIC_INFORMATION
	FileCacheInfo MINIDUMP_SYSTEM_FILECACHE_INFORMATION
	BasicPerfInfo MINIDUMP_SYSTEM_BASIC_PERFORMANCE_INFORMATION
	PerfInfo      MINIDUMP_SYSTEM_PERFORMANCE_INFORMATION
}

type MINIDUMP_PROCESS_VM_COUNTERS_1 struct {
	Revision                   uint16
	PageFaultCount             uint32
	PeakWorkingSetSize         uint64
	WorkingSetSize             uint64
	QuotaPeakPagedPoolUsage    uint64
	QuotaPagedPoolUsage        uint64
	QuotaPeakNonPagedPoolUsage uint64
	QuotaNonPagedPoolUsage     uint64
	PagefileUsage              uint64
	PeakPagefileUsage          uint64
	PrivateUsage               uint64
}

type MINIDUMP_PROCESS_VM_COUNTERS_2 struct {
	Revision                   uint16
	Flags                      uint16
	PageFaultCount             uint32
	PeakWorkingSetSize         uint64
	WorkingSetSize             uint64
	QuotaPeakPagedPoolUsage    uint64
	QuotaPagedPoolUsage        uint64
	QuotaPeakNonPagedPoolUsage uint64
	QuotaNonPagedPoolUsage     uint64
	PagefileUsage              uint64
	PeakPagefileUsage          uint64
	PeakVirtualSize            uint64
	VirtualSize                uint64
	PrivateUsage               uint64
	PrivateWorkingSetSize      uint64
	SharedCommitUsage          uint64
	JobSharedCommitUsage       uint64
	JobPrivateCommitUsage      uint64
	JobPeakPrivateCommitUsage  uint64
	JobPrivateCommitLimit      uint64
	JobTotalCommitLimit        uint64
}

type MINIDUMP_USER_RECORD struct {
	Type   uint32
	Memory MINIDUMP_LOCATION_DESCRIPTOR
}

type MINIDUMP_THREAD_CALLBACK struct {
	ThreadId      uint32
	ThreadHandle  HANDLE
	Context       CONTEXT
	SizeOfContext uint32
	StackBase     uint64
	StackEnd      uint64
}

type MINIDUMP_THREAD_EX_CALLBACK struct {
	ThreadId         uint32
	ThreadHandle     HANDLE
	Context          CONTEXT
	SizeOfContext    uint32
	StackBase        uint64
	StackEnd         uint64
	BackingStoreBase uint64
	BackingStoreEnd  uint64
}

type MINIDUMP_INCLUDE_THREAD_CALLBACK struct {
	ThreadId uint32
}

type MINIDUMP_MODULE_CALLBACK struct {
	FullPath         PWSTR
	BaseOfImage      uint64
	SizeOfImage      uint32
	CheckSum         uint32
	TimeDateStamp    uint32
	VersionInfo      VS_FIXEDFILEINFO
	CvRecord         unsafe.Pointer
	SizeOfCvRecord   uint32
	MiscRecord       unsafe.Pointer
	SizeOfMiscRecord uint32
}

type MINIDUMP_INCLUDE_MODULE_CALLBACK struct {
	BaseOfImage uint64
}

type MINIDUMP_IO_CALLBACK struct {
	Handle      HANDLE
	Offset      uint64
	Buffer      unsafe.Pointer
	BufferBytes uint32
}

type MINIDUMP_READ_MEMORY_FAILURE_CALLBACK struct {
	Offset        uint64
	Bytes         uint32
	FailureStatus HRESULT
}

type MINIDUMP_VM_QUERY_CALLBACK struct {
	Offset uint64
}

type MINIDUMP_VM_PRE_READ_CALLBACK struct {
	Offset uint64
	Buffer unsafe.Pointer
	Size   uint32
}

type MINIDUMP_VM_POST_READ_CALLBACK struct {
	Offset    uint64
	Buffer    unsafe.Pointer
	Size      uint32
	Completed uint32
	Status    HRESULT
}

type MINIDUMP_CALLBACK_INPUT_Anonymous struct {
	Data [128]uint64
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) Status() *HRESULT {
	return (*HRESULT)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) StatusVal() HRESULT {
	return *(*HRESULT)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) Thread() *MINIDUMP_THREAD_CALLBACK {
	return (*MINIDUMP_THREAD_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) ThreadVal() MINIDUMP_THREAD_CALLBACK {
	return *(*MINIDUMP_THREAD_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) ThreadEx() *MINIDUMP_THREAD_EX_CALLBACK {
	return (*MINIDUMP_THREAD_EX_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) ThreadExVal() MINIDUMP_THREAD_EX_CALLBACK {
	return *(*MINIDUMP_THREAD_EX_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) Module() *MINIDUMP_MODULE_CALLBACK {
	return (*MINIDUMP_MODULE_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) ModuleVal() MINIDUMP_MODULE_CALLBACK {
	return *(*MINIDUMP_MODULE_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) IncludeThread() *MINIDUMP_INCLUDE_THREAD_CALLBACK {
	return (*MINIDUMP_INCLUDE_THREAD_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) IncludeThreadVal() MINIDUMP_INCLUDE_THREAD_CALLBACK {
	return *(*MINIDUMP_INCLUDE_THREAD_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) IncludeModule() *MINIDUMP_INCLUDE_MODULE_CALLBACK {
	return (*MINIDUMP_INCLUDE_MODULE_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) IncludeModuleVal() MINIDUMP_INCLUDE_MODULE_CALLBACK {
	return *(*MINIDUMP_INCLUDE_MODULE_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) Io() *MINIDUMP_IO_CALLBACK {
	return (*MINIDUMP_IO_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) IoVal() MINIDUMP_IO_CALLBACK {
	return *(*MINIDUMP_IO_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) ReadMemoryFailure() *MINIDUMP_READ_MEMORY_FAILURE_CALLBACK {
	return (*MINIDUMP_READ_MEMORY_FAILURE_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) ReadMemoryFailureVal() MINIDUMP_READ_MEMORY_FAILURE_CALLBACK {
	return *(*MINIDUMP_READ_MEMORY_FAILURE_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) SecondaryFlags() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) SecondaryFlagsVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) VmQuery() *MINIDUMP_VM_QUERY_CALLBACK {
	return (*MINIDUMP_VM_QUERY_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) VmQueryVal() MINIDUMP_VM_QUERY_CALLBACK {
	return *(*MINIDUMP_VM_QUERY_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) VmPreRead() *MINIDUMP_VM_PRE_READ_CALLBACK {
	return (*MINIDUMP_VM_PRE_READ_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) VmPreReadVal() MINIDUMP_VM_PRE_READ_CALLBACK {
	return *(*MINIDUMP_VM_PRE_READ_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) VmPostRead() *MINIDUMP_VM_POST_READ_CALLBACK {
	return (*MINIDUMP_VM_POST_READ_CALLBACK)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_INPUT_Anonymous) VmPostReadVal() MINIDUMP_VM_POST_READ_CALLBACK {
	return *(*MINIDUMP_VM_POST_READ_CALLBACK)(unsafe.Pointer(this))
}

type MINIDUMP_CALLBACK_INPUT struct {
	ProcessId     uint32
	ProcessHandle HANDLE
	CallbackType  uint32
	MINIDUMP_CALLBACK_INPUT_Anonymous
}

type MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous1 struct {
	MemoryBase uint64
	MemorySize uint32
}

type MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous2 struct {
	CheckCancel BOOL
	Cancel      BOOL
}

type MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous3 struct {
	VmRegion MINIDUMP_MEMORY_INFO
	Continue BOOL
}

type MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous4 struct {
	VmQueryStatus HRESULT
	VmQueryResult MINIDUMP_MEMORY_INFO
}

type MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous5 struct {
	VmReadStatus         HRESULT
	VmReadBytesCompleted uint32
}

type MINIDUMP_CALLBACK_OUTPUT_Anonymous struct {
	Data [7]uint64
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) ModuleWriteFlags() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) ModuleWriteFlagsVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) ThreadWriteFlags() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) ThreadWriteFlagsVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) SecondaryFlags() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) SecondaryFlagsVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous1() *MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous1 {
	return (*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous1)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous1Val() MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous1 {
	return *(*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous1)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous2() *MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous2 {
	return (*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous2)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous2Val() MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous2 {
	return *(*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous2)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Handle() *HANDLE {
	return (*HANDLE)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) HandleVal() HANDLE {
	return *(*HANDLE)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous3() *MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous3 {
	return (*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous3)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous3Val() MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous3 {
	return *(*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous3)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous4() *MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous4 {
	return (*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous4)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous4Val() MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous4 {
	return *(*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous4)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous5() *MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous5 {
	return (*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous5)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Anonymous5Val() MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous5 {
	return *(*MINIDUMP_CALLBACK_OUTPUT_Anonymous_Anonymous5)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) Status() *HRESULT {
	return (*HRESULT)(unsafe.Pointer(this))
}

func (this *MINIDUMP_CALLBACK_OUTPUT_Anonymous) StatusVal() HRESULT {
	return *(*HRESULT)(unsafe.Pointer(this))
}

type MINIDUMP_CALLBACK_OUTPUT struct {
	MINIDUMP_CALLBACK_OUTPUT_Anonymous
}

type MODLOAD_DATA struct {
	Ssize uint32
	Ssig  MODLOAD_DATA_TYPE
	Data  unsafe.Pointer
	Size  uint32
	Flags uint32
}

type MODLOAD_CVMISC struct {
	OCV     uint32
	CCV     uintptr
	OMisc   uint32
	CMisc   uintptr
	DtImage uint32
	CImage  uint32
}

type MODLOAD_PDBGUID_PDBAGE struct {
	PdbGuid syscall.GUID
	PdbAge  uint32
}

type ADDRESS64 struct {
	Offset  uint64
	Segment uint16
	Mode    ADDRESS_MODE
}

type KDHELP64 struct {
	Thread                         uint64
	ThCallbackStack                uint32
	ThCallbackBStore               uint32
	NextCallback                   uint32
	FramePointer                   uint32
	KiCallUserMode                 uint64
	KeUserCallbackDispatcher       uint64
	SystemRangeStart               uint64
	KiUserExceptionDispatcher      uint64
	StackBase                      uint64
	StackLimit                     uint64
	BuildVersion                   uint32
	RetpolineStubFunctionTableSize uint32
	RetpolineStubFunctionTable     uint64
	RetpolineStubOffset            uint32
	RetpolineStubSize              uint32
	Reserved0                      [2]uint64
}

type STACKFRAME64 struct {
	AddrPC         ADDRESS64
	AddrReturn     ADDRESS64
	AddrFrame      ADDRESS64
	AddrStack      ADDRESS64
	AddrBStore     ADDRESS64
	FuncTableEntry unsafe.Pointer
	Params         [4]uint64
	Far            BOOL
	Virtual        BOOL
	Reserved       [3]uint64
	KdHelp         KDHELP64
}

type STACKFRAME_EX struct {
	AddrPC             ADDRESS64
	AddrReturn         ADDRESS64
	AddrFrame          ADDRESS64
	AddrStack          ADDRESS64
	AddrBStore         ADDRESS64
	FuncTableEntry     unsafe.Pointer
	Params             [4]uint64
	Far                BOOL
	Virtual            BOOL
	Reserved           [3]uint64
	KdHelp             KDHELP64
	StackFrameSize     uint32
	InlineFrameContext uint32
}

type API_VERSION struct {
	MajorVersion uint16
	MinorVersion uint16
	Revision     uint16
	Reserved     uint16
}

type IMAGEHLP_SYMBOL64 struct {
	SizeOfStruct  uint32
	Address       uint64
	Size          uint32
	Flags         uint32
	MaxNameLength uint32
	Name          [1]CHAR
}

type IMAGEHLP_SYMBOL64_PACKAGE struct {
	Sym  IMAGEHLP_SYMBOL64
	Name [2001]CHAR
}

type IMAGEHLP_SYMBOLW64 struct {
	SizeOfStruct  uint32
	Address       uint64
	Size          uint32
	Flags         uint32
	MaxNameLength uint32
	Name          [1]uint16
}

type IMAGEHLP_SYMBOLW64_PACKAGE struct {
	Sym  IMAGEHLP_SYMBOLW64
	Name [2001]uint16
}

type IMAGEHLP_MODULE64 struct {
	SizeOfStruct    uint32
	BaseOfImage     uint64
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         SYM_TYPE
	ModuleName      [32]CHAR
	ImageName       [256]CHAR
	LoadedImageName [256]CHAR
	LoadedPdbName   [256]CHAR
	CVSig           uint32
	CVData          [780]CHAR
	PdbSig          uint32
	PdbSig70        syscall.GUID
	PdbAge          uint32
	PdbUnmatched    BOOL
	DbgUnmatched    BOOL
	LineNumbers     BOOL
	GlobalSymbols   BOOL
	TypeInfo        BOOL
	SourceIndexed   BOOL
	Publics         BOOL
	MachineType     uint32
	Reserved        uint32
}

type IMAGEHLP_MODULE64_EX struct {
	Module      IMAGEHLP_MODULE64
	RegionFlags uint32
}

type IMAGEHLP_MODULEW64 struct {
	SizeOfStruct    uint32
	BaseOfImage     uint64
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         SYM_TYPE
	ModuleName      [32]uint16
	ImageName       [256]uint16
	LoadedImageName [256]uint16
	LoadedPdbName   [256]uint16
	CVSig           uint32
	CVData          [780]uint16
	PdbSig          uint32
	PdbSig70        syscall.GUID
	PdbAge          uint32
	PdbUnmatched    BOOL
	DbgUnmatched    BOOL
	LineNumbers     BOOL
	GlobalSymbols   BOOL
	TypeInfo        BOOL
	SourceIndexed   BOOL
	Publics         BOOL
	MachineType     uint32
	Reserved        uint32
}

type IMAGEHLP_MODULEW64_EX struct {
	Module      IMAGEHLP_MODULEW64
	RegionFlags uint32
}

type IMAGEHLP_LINE64 struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	LineNumber   uint32
	FileName     PSTR
	Address      uint64
}

type IMAGEHLP_LINEW64 struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	LineNumber   uint32
	FileName     PWSTR
	Address      uint64
}

type SOURCEFILE struct {
	ModBase  uint64
	FileName PSTR
}

type SOURCEFILEW struct {
	ModBase  uint64
	FileName PWSTR
}

type IMAGEHLP_CBA_READ_MEMORY struct {
	Addr      uint64
	Buf       unsafe.Pointer
	Bytes     uint32
	Bytesread *uint32
}

type IMAGEHLP_CBA_EVENT struct {
	Severity IMAGEHLP_CBA_EVENT_SEVERITY
	Code     uint32
	Desc     PSTR
	Object   unsafe.Pointer
}

type IMAGEHLP_CBA_EVENTW struct {
	Severity IMAGEHLP_CBA_EVENT_SEVERITY
	Code     uint32
	Desc     PWSTR
	Object   unsafe.Pointer
}

type IMAGEHLP_DEFERRED_SYMBOL_LOAD64 struct {
	SizeOfStruct  uint32
	BaseOfImage   uint64
	CheckSum      uint32
	TimeDateStamp uint32
	FileName      [260]CHAR
	Reparse       BOOLEAN
	HFile         HANDLE
	Flags         uint32
}

type IMAGEHLP_DEFERRED_SYMBOL_LOADW64 struct {
	SizeOfStruct  uint32
	BaseOfImage   uint64
	CheckSum      uint32
	TimeDateStamp uint32
	FileName      [261]uint16
	Reparse       BOOLEAN
	HFile         HANDLE
	Flags         uint32
}

type IMAGEHLP_DUPLICATE_SYMBOL64 struct {
	SizeOfStruct   uint32
	NumberOfDups   uint32
	Symbol         *IMAGEHLP_SYMBOL64
	SelectedSymbol uint32
}

type IMAGEHLP_JIT_SYMBOLMAP struct {
	SizeOfStruct uint32
	Address      uint64
	BaseOfImage  uint64
}

type OMAP struct {
	Rva   uint32
	RvaTo uint32
}

type SRCCODEINFO struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	ModBase      uint64
	Obj          [261]CHAR
	FileName     [261]CHAR
	LineNumber   uint32
	Address      uint64
}

type SRCCODEINFOW struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	ModBase      uint64
	Obj          [261]uint16
	FileName     [261]uint16
	LineNumber   uint32
	Address      uint64
}

type IMAGEHLP_SYMBOL_SRC struct {
	Sizeofstruct uint32
	Type_        uint32
	File         [260]CHAR
}

type MODULE_TYPE_INFO struct {
	DataLength uint16
	Leaf       uint16
	Data       [1]byte
}

type SYMBOL_INFO struct {
	SizeOfStruct uint32
	TypeIndex    uint32
	Reserved     [2]uint64
	Index        uint32
	Size         uint32
	ModBase      uint64
	Flags        SYMBOL_INFO_FLAGS
	Value        uint64
	Address      uint64
	Register     uint32
	Scope        uint32
	Tag          uint32
	NameLen      uint32
	MaxNameLen   uint32
	Name         [1]CHAR
}

type SYMBOL_INFO_PACKAGE struct {
	Si   SYMBOL_INFO
	Name [2001]CHAR
}

type SYMBOL_INFOW struct {
	SizeOfStruct uint32
	TypeIndex    uint32
	Reserved     [2]uint64
	Index        uint32
	Size         uint32
	ModBase      uint64
	Flags        SYMBOL_INFO_FLAGS
	Value        uint64
	Address      uint64
	Register     uint32
	Scope        uint32
	Tag          uint32
	NameLen      uint32
	MaxNameLen   uint32
	Name         [1]uint16
}

type SYMBOL_INFO_PACKAGEW struct {
	Si   SYMBOL_INFOW
	Name [2001]uint16
}

type IMAGEHLP_STACK_FRAME struct {
	InstructionOffset  uint64
	ReturnOffset       uint64
	FrameOffset        uint64
	StackOffset        uint64
	BackingStoreOffset uint64
	FuncTableEntry     uint64
	Params             [4]uint64
	Reserved           [5]uint64
	Virtual            BOOL
	Reserved2          uint32
}

type TI_FINDCHILDREN_PARAMS struct {
	Count   uint32
	Start   uint32
	ChildId [1]uint32
}

type IMAGEHLP_GET_TYPE_INFO_PARAMS struct {
	SizeOfStruct   uint32
	Flags          IMAGEHLP_GET_TYPE_INFO_FLAGS
	NumIds         uint32
	TypeIds        *uint32
	TagFilter      uint64
	NumReqs        uint32
	ReqKinds       *IMAGEHLP_SYMBOL_TYPE_INFO
	ReqOffsets     *uintptr
	ReqSizes       *uint32
	ReqStride      uintptr
	BufferSize     uintptr
	Buffer         unsafe.Pointer
	EntriesMatched uint32
	EntriesFilled  uint32
	TagsFound      uint64
	AllReqsValid   uint64
	NumReqsValid   uint32
	ReqsValid      *uint64
}

type SYMSRV_INDEX_INFO struct {
	Sizeofstruct uint32
	File         [261]CHAR
	Stripped     BOOL
	Timestamp    uint32
	Size         uint32
	Dbgfile      [261]CHAR
	Pdbfile      [261]CHAR
	Guid         syscall.GUID
	Sig          uint32
	Age          uint32
}

type SYMSRV_INDEX_INFOW struct {
	Sizeofstruct uint32
	File         [261]uint16
	Stripped     BOOL
	Timestamp    uint32
	Size         uint32
	Dbgfile      [261]uint16
	Pdbfile      [261]uint16
	Guid         syscall.GUID
	Sig          uint32
	Age          uint32
}

type SYMSRV_EXTENDED_OUTPUT_DATA struct {
	SizeOfStruct uint32
	Version      uint32
	FilePtrMsg   [261]uint16
}

type DBGHELP_DATA_REPORT_STRUCT struct {
	PBinPathNonExist    PWSTR
	PSymbolPathNonExist PWSTR
}

type PHYSICAL_MEMORY_RUN32 struct {
	BasePage  uint32
	PageCount uint32
}

type PHYSICAL_MEMORY_DESCRIPTOR32 struct {
	NumberOfRuns  uint32
	NumberOfPages uint32
	Run           [1]PHYSICAL_MEMORY_RUN32
}

type PHYSICAL_MEMORY_RUN64 struct {
	BasePage  uint64
	PageCount uint64
}

type PHYSICAL_MEMORY_DESCRIPTOR64 struct {
	NumberOfRuns  uint32
	NumberOfPages uint64
	Run           [1]PHYSICAL_MEMORY_RUN64
}

type DUMP_FILE_ATTRIBUTES_Anonymous struct {
	Bitfield_ uint32
}

type DUMP_FILE_ATTRIBUTES struct {
	DUMP_FILE_ATTRIBUTES_Anonymous
}

func (this *DUMP_FILE_ATTRIBUTES) Anonymous() *DUMP_FILE_ATTRIBUTES_Anonymous {
	return (*DUMP_FILE_ATTRIBUTES_Anonymous)(unsafe.Pointer(this))
}

func (this *DUMP_FILE_ATTRIBUTES) AnonymousVal() DUMP_FILE_ATTRIBUTES_Anonymous {
	return *(*DUMP_FILE_ATTRIBUTES_Anonymous)(unsafe.Pointer(this))
}

func (this *DUMP_FILE_ATTRIBUTES) Attributes() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *DUMP_FILE_ATTRIBUTES) AttributesVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type DUMP_HEADER32_Anonymous struct {
	Data [175]uint32
}

func (this *DUMP_HEADER32_Anonymous) PhysicalMemoryBlock() *PHYSICAL_MEMORY_DESCRIPTOR32 {
	return (*PHYSICAL_MEMORY_DESCRIPTOR32)(unsafe.Pointer(this))
}

func (this *DUMP_HEADER32_Anonymous) PhysicalMemoryBlockVal() PHYSICAL_MEMORY_DESCRIPTOR32 {
	return *(*PHYSICAL_MEMORY_DESCRIPTOR32)(unsafe.Pointer(this))
}

func (this *DUMP_HEADER32_Anonymous) PhysicalMemoryBlockBuffer() *[700]byte {
	return (*[700]byte)(unsafe.Pointer(this))
}

func (this *DUMP_HEADER32_Anonymous) PhysicalMemoryBlockBufferVal() [700]byte {
	return *(*[700]byte)(unsafe.Pointer(this))
}

type DUMP_HEADER32 struct {
	Signature           uint32
	ValidDump           uint32
	MajorVersion        uint32
	MinorVersion        uint32
	DirectoryTableBase  uint32
	PfnDataBase         uint32
	PsLoadedModuleList  uint32
	PsActiveProcessHead uint32
	MachineImageType    uint32
	NumberProcessors    uint32
	BugCheckCode        uint32
	BugCheckParameter1  uint32
	BugCheckParameter2  uint32
	BugCheckParameter3  uint32
	BugCheckParameter4  uint32
	VersionUser         [32]CHAR
	PaeEnabled          byte
	KdSecondaryVersion  byte
	Spare3              [2]byte
	KdDebuggerDataBlock uint32
	DUMP_HEADER32_Anonymous
	ContextRecord      [1200]byte
	Exception          EXCEPTION_RECORD32
	Comment            [128]CHAR
	Attributes         DUMP_FILE_ATTRIBUTES
	BootId             uint32
	Reserved0_         [1760]byte
	DumpType           uint32
	MiniDumpFields     uint32
	SecondaryDataState uint32
	ProductType        uint32
	SuiteMask          uint32
	WriterStatus       uint32
	RequiredDumpSpace  int64
	Reserved2_         [16]byte
	SystemUpTime       int64
	SystemTime         int64
	Reserved3_         [56]byte
}

type DUMP_HEADER64_Anonymous struct {
	Data [87]uint64
}

func (this *DUMP_HEADER64_Anonymous) PhysicalMemoryBlock() *PHYSICAL_MEMORY_DESCRIPTOR64 {
	return (*PHYSICAL_MEMORY_DESCRIPTOR64)(unsafe.Pointer(this))
}

func (this *DUMP_HEADER64_Anonymous) PhysicalMemoryBlockVal() PHYSICAL_MEMORY_DESCRIPTOR64 {
	return *(*PHYSICAL_MEMORY_DESCRIPTOR64)(unsafe.Pointer(this))
}

func (this *DUMP_HEADER64_Anonymous) PhysicalMemoryBlockBuffer() *[700]byte {
	return (*[700]byte)(unsafe.Pointer(this))
}

func (this *DUMP_HEADER64_Anonymous) PhysicalMemoryBlockBufferVal() [700]byte {
	return *(*[700]byte)(unsafe.Pointer(this))
}

type DUMP_HEADER64 struct {
	Signature           uint32
	ValidDump           uint32
	MajorVersion        uint32
	MinorVersion        uint32
	DirectoryTableBase  uint64
	PfnDataBase         uint64
	PsLoadedModuleList  uint64
	PsActiveProcessHead uint64
	MachineImageType    uint32
	NumberProcessors    uint32
	BugCheckCode        uint32
	BugCheckParameter1  uint64
	BugCheckParameter2  uint64
	BugCheckParameter3  uint64
	BugCheckParameter4  uint64
	VersionUser         [32]CHAR
	KdDebuggerDataBlock uint64
	DUMP_HEADER64_Anonymous
	ContextRecord      [3000]byte
	Exception          EXCEPTION_RECORD64
	DumpType           uint32
	RequiredDumpSpace  int64
	SystemTime         int64
	Comment            [128]CHAR
	SystemUpTime       int64
	MiniDumpFields     uint32
	SecondaryDataState uint32
	ProductType        uint32
	SuiteMask          uint32
	WriterStatus       uint32
	Unused1            byte
	KdSecondaryVersion byte
	Unused             [2]byte
	Attributes         DUMP_FILE_ATTRIBUTES
	BootId             uint32
	Reserved0_         [4008]byte
}

type WHEA_ERROR_SOURCE_CONFIGURATION_DD struct {
	Initialize   WHEA_ERROR_SOURCE_INITIALIZE_DEVICE_DRIVER
	Uninitialize WHEA_ERROR_SOURCE_UNINITIALIZE_DEVICE_DRIVER
	Correct      WHEA_ERROR_SOURCE_CORRECT_DEVICE_DRIVER
}

type WHEA_ERROR_SOURCE_CONFIGURATION_DEVICE_DRIVER_V1 struct {
	Version      uint32
	SourceGuid   syscall.GUID
	LogTag       uint16
	Reserved     [6]byte
	Initialize   WHEA_ERROR_SOURCE_INITIALIZE_DEVICE_DRIVER
	Uninitialize WHEA_ERROR_SOURCE_UNINITIALIZE_DEVICE_DRIVER
}

type WHEA_ERROR_SOURCE_CONFIGURATION_DEVICE_DRIVER struct {
	Version              uint32
	SourceGuid           syscall.GUID
	LogTag               uint16
	Reserved             [6]byte
	Initialize           WHEA_ERROR_SOURCE_INITIALIZE_DEVICE_DRIVER
	Uninitialize         WHEA_ERROR_SOURCE_UNINITIALIZE_DEVICE_DRIVER
	MaxSectionDataLength uint32
	MaxSectionsPerReport uint32
	CreatorId            syscall.GUID
	PartitionId          syscall.GUID
}

type WHEA_DRIVER_BUFFER_SET struct {
	Version             uint32
	Data                *byte
	DataSize            uint32
	SectionTypeGuid     *syscall.GUID
	SectionFriendlyName *byte
	Flags               *byte
}

type WHEA_NOTIFICATION_FLAGS_Anonymous struct {
	Bitfield_ uint16
}

type WHEA_NOTIFICATION_FLAGS struct {
	WHEA_NOTIFICATION_FLAGS_Anonymous
}

func (this *WHEA_NOTIFICATION_FLAGS) Anonymous() *WHEA_NOTIFICATION_FLAGS_Anonymous {
	return (*WHEA_NOTIFICATION_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_FLAGS) AnonymousVal() WHEA_NOTIFICATION_FLAGS_Anonymous {
	return *(*WHEA_NOTIFICATION_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_FLAGS) AsUSHORT() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_FLAGS) AsUSHORTVal() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

type XPF_MC_BANK_FLAGS_Anonymous struct {
	Bitfield_ byte
}

type XPF_MC_BANK_FLAGS struct {
	XPF_MC_BANK_FLAGS_Anonymous
}

func (this *XPF_MC_BANK_FLAGS) Anonymous() *XPF_MC_BANK_FLAGS_Anonymous {
	return (*XPF_MC_BANK_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *XPF_MC_BANK_FLAGS) AnonymousVal() XPF_MC_BANK_FLAGS_Anonymous {
	return *(*XPF_MC_BANK_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *XPF_MC_BANK_FLAGS) AsUCHAR() *byte {
	return (*byte)(unsafe.Pointer(this))
}

func (this *XPF_MC_BANK_FLAGS) AsUCHARVal() byte {
	return *(*byte)(unsafe.Pointer(this))
}

type XPF_MCE_FLAGS_Anonymous struct {
	Bitfield_ uint32
}

type XPF_MCE_FLAGS struct {
	XPF_MCE_FLAGS_Anonymous
}

func (this *XPF_MCE_FLAGS) Anonymous() *XPF_MCE_FLAGS_Anonymous {
	return (*XPF_MCE_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *XPF_MCE_FLAGS) AnonymousVal() XPF_MCE_FLAGS_Anonymous {
	return *(*XPF_MCE_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *XPF_MCE_FLAGS) AsULONG() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *XPF_MCE_FLAGS) AsULONGVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type AER_ROOTPORT_DESCRIPTOR_FLAGS_Anonymous struct {
	Bitfield_ uint16
}

type AER_ROOTPORT_DESCRIPTOR_FLAGS struct {
	AER_ROOTPORT_DESCRIPTOR_FLAGS_Anonymous
}

func (this *AER_ROOTPORT_DESCRIPTOR_FLAGS) Anonymous() *AER_ROOTPORT_DESCRIPTOR_FLAGS_Anonymous {
	return (*AER_ROOTPORT_DESCRIPTOR_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *AER_ROOTPORT_DESCRIPTOR_FLAGS) AnonymousVal() AER_ROOTPORT_DESCRIPTOR_FLAGS_Anonymous {
	return *(*AER_ROOTPORT_DESCRIPTOR_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *AER_ROOTPORT_DESCRIPTOR_FLAGS) AsUSHORT() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *AER_ROOTPORT_DESCRIPTOR_FLAGS) AsUSHORTVal() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

type AER_ENDPOINT_DESCRIPTOR_FLAGS_Anonymous struct {
	Bitfield_ uint16
}

type AER_ENDPOINT_DESCRIPTOR_FLAGS struct {
	AER_ENDPOINT_DESCRIPTOR_FLAGS_Anonymous
}

func (this *AER_ENDPOINT_DESCRIPTOR_FLAGS) Anonymous() *AER_ENDPOINT_DESCRIPTOR_FLAGS_Anonymous {
	return (*AER_ENDPOINT_DESCRIPTOR_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *AER_ENDPOINT_DESCRIPTOR_FLAGS) AnonymousVal() AER_ENDPOINT_DESCRIPTOR_FLAGS_Anonymous {
	return *(*AER_ENDPOINT_DESCRIPTOR_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *AER_ENDPOINT_DESCRIPTOR_FLAGS) AsUSHORT() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *AER_ENDPOINT_DESCRIPTOR_FLAGS) AsUSHORTVal() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

type AER_BRIDGE_DESCRIPTOR_FLAGS_Anonymous struct {
	Bitfield_ uint16
}

type AER_BRIDGE_DESCRIPTOR_FLAGS struct {
	AER_BRIDGE_DESCRIPTOR_FLAGS_Anonymous
}

func (this *AER_BRIDGE_DESCRIPTOR_FLAGS) Anonymous() *AER_BRIDGE_DESCRIPTOR_FLAGS_Anonymous {
	return (*AER_BRIDGE_DESCRIPTOR_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *AER_BRIDGE_DESCRIPTOR_FLAGS) AnonymousVal() AER_BRIDGE_DESCRIPTOR_FLAGS_Anonymous {
	return *(*AER_BRIDGE_DESCRIPTOR_FLAGS_Anonymous)(unsafe.Pointer(this))
}

func (this *AER_BRIDGE_DESCRIPTOR_FLAGS) AsUSHORT() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *AER_BRIDGE_DESCRIPTOR_FLAGS) AsUSHORTVal() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Polled struct {
	PollInterval uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Interrupt struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_LocalInterrupt struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Sci struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Nmi struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Sea struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Sei struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U_Gsiv struct {
	PollInterval             uint32
	Vector                   uint32
	SwitchToPollingThreshold uint32
	SwitchToPollingWindow    uint32
	ErrorThreshold           uint32
	ErrorThresholdWindow     uint32
}

type WHEA_NOTIFICATION_DESCRIPTOR_U struct {
	Data [6]uint32
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Polled() *WHEA_NOTIFICATION_DESCRIPTOR_U_Polled {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Polled)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) PolledVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Polled {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Polled)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Interrupt() *WHEA_NOTIFICATION_DESCRIPTOR_U_Interrupt {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Interrupt)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) InterruptVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Interrupt {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Interrupt)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) LocalInterrupt() *WHEA_NOTIFICATION_DESCRIPTOR_U_LocalInterrupt {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_LocalInterrupt)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) LocalInterruptVal() WHEA_NOTIFICATION_DESCRIPTOR_U_LocalInterrupt {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_LocalInterrupt)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Sci() *WHEA_NOTIFICATION_DESCRIPTOR_U_Sci {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Sci)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) SciVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Sci {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Sci)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Nmi() *WHEA_NOTIFICATION_DESCRIPTOR_U_Nmi {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Nmi)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) NmiVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Nmi {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Nmi)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Sea() *WHEA_NOTIFICATION_DESCRIPTOR_U_Sea {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Sea)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) SeaVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Sea {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Sea)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Sei() *WHEA_NOTIFICATION_DESCRIPTOR_U_Sei {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Sei)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) SeiVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Sei {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Sei)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) Gsiv() *WHEA_NOTIFICATION_DESCRIPTOR_U_Gsiv {
	return (*WHEA_NOTIFICATION_DESCRIPTOR_U_Gsiv)(unsafe.Pointer(this))
}

func (this *WHEA_NOTIFICATION_DESCRIPTOR_U) GsivVal() WHEA_NOTIFICATION_DESCRIPTOR_U_Gsiv {
	return *(*WHEA_NOTIFICATION_DESCRIPTOR_U_Gsiv)(unsafe.Pointer(this))
}

type WHEA_NOTIFICATION_DESCRIPTOR struct {
	Type   byte
	Length byte
	Flags  WHEA_NOTIFICATION_FLAGS
	U      WHEA_NOTIFICATION_DESCRIPTOR_U
}

type WHEA_XPF_MC_BANK_DESCRIPTOR struct {
	BankNumber            byte
	ClearOnInitialization BOOLEAN
	StatusDataFormat      byte
	Flags                 XPF_MC_BANK_FLAGS
	ControlMsr            uint32
	StatusMsr             uint32
	AddressMsr            uint32
	MiscMsr               uint32
	ControlData           uint64
}

type WHEA_XPF_MCE_DESCRIPTOR struct {
	Type              uint16
	Enabled           byte
	NumberOfBanks     byte
	Flags             XPF_MCE_FLAGS
	MCG_Capability    uint64
	MCG_GlobalControl uint64
	Banks             [32]WHEA_XPF_MC_BANK_DESCRIPTOR
}

type WHEA_XPF_CMC_DESCRIPTOR struct {
	Type          uint16
	Enabled       BOOLEAN
	NumberOfBanks byte
	Reserved      uint32
	Notify        WHEA_NOTIFICATION_DESCRIPTOR
	Banks         [32]WHEA_XPF_MC_BANK_DESCRIPTOR
}

type WHEA_PCI_SLOT_NUMBER_U_Bits struct {
	Bitfield_ uint32
}

type WHEA_PCI_SLOT_NUMBER_U struct {
	Data [1]uint32
}

func (this *WHEA_PCI_SLOT_NUMBER_U) Bits() *WHEA_PCI_SLOT_NUMBER_U_Bits {
	return (*WHEA_PCI_SLOT_NUMBER_U_Bits)(unsafe.Pointer(this))
}

func (this *WHEA_PCI_SLOT_NUMBER_U) BitsVal() WHEA_PCI_SLOT_NUMBER_U_Bits {
	return *(*WHEA_PCI_SLOT_NUMBER_U_Bits)(unsafe.Pointer(this))
}

func (this *WHEA_PCI_SLOT_NUMBER_U) AsULONG() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *WHEA_PCI_SLOT_NUMBER_U) AsULONGVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

type WHEA_PCI_SLOT_NUMBER struct {
	U WHEA_PCI_SLOT_NUMBER_U
}

type WHEA_XPF_NMI_DESCRIPTOR struct {
	Type    uint16
	Enabled BOOLEAN
}

type WHEA_AER_ROOTPORT_DESCRIPTOR struct {
	Type                       uint16
	Enabled                    BOOLEAN
	Reserved                   byte
	BusNumber                  uint32
	Slot                       WHEA_PCI_SLOT_NUMBER
	DeviceControl              uint16
	Flags                      AER_ROOTPORT_DESCRIPTOR_FLAGS
	UncorrectableErrorMask     uint32
	UncorrectableErrorSeverity uint32
	CorrectableErrorMask       uint32
	AdvancedCapsAndControl     uint32
	RootErrorCommand           uint32
}

type WHEA_AER_ENDPOINT_DESCRIPTOR struct {
	Type                       uint16
	Enabled                    BOOLEAN
	Reserved                   byte
	BusNumber                  uint32
	Slot                       WHEA_PCI_SLOT_NUMBER
	DeviceControl              uint16
	Flags                      AER_ENDPOINT_DESCRIPTOR_FLAGS
	UncorrectableErrorMask     uint32
	UncorrectableErrorSeverity uint32
	CorrectableErrorMask       uint32
	AdvancedCapsAndControl     uint32
}

type WHEA_AER_BRIDGE_DESCRIPTOR struct {
	Type                            uint16
	Enabled                         BOOLEAN
	Reserved                        byte
	BusNumber                       uint32
	Slot                            WHEA_PCI_SLOT_NUMBER
	DeviceControl                   uint16
	Flags                           AER_BRIDGE_DESCRIPTOR_FLAGS
	UncorrectableErrorMask          uint32
	UncorrectableErrorSeverity      uint32
	CorrectableErrorMask            uint32
	AdvancedCapsAndControl          uint32
	SecondaryUncorrectableErrorMask uint32
	SecondaryUncorrectableErrorSev  uint32
	SecondaryCapsAndControl         uint32
}

type WHEA_GENERIC_ERROR_DESCRIPTOR struct {
	Type                       uint16
	Reserved                   byte
	Enabled                    byte
	ErrStatusBlockLength       uint32
	RelatedErrorSourceId       uint32
	ErrStatusAddressSpaceID    byte
	ErrStatusAddressBitWidth   byte
	ErrStatusAddressBitOffset  byte
	ErrStatusAddressAccessSize byte
	ErrStatusAddress           int64
	Notify                     WHEA_NOTIFICATION_DESCRIPTOR
}

type WHEA_GENERIC_ERROR_DESCRIPTOR_V2 struct {
	Type                       uint16
	Reserved                   byte
	Enabled                    byte
	ErrStatusBlockLength       uint32
	RelatedErrorSourceId       uint32
	ErrStatusAddressSpaceID    byte
	ErrStatusAddressBitWidth   byte
	ErrStatusAddressBitOffset  byte
	ErrStatusAddressAccessSize byte
	ErrStatusAddress           int64
	Notify                     WHEA_NOTIFICATION_DESCRIPTOR
	ReadAckAddressSpaceID      byte
	ReadAckAddressBitWidth     byte
	ReadAckAddressBitOffset    byte
	ReadAckAddressAccessSize   byte
	ReadAckAddress             int64
	ReadAckPreserveMask        uint64
	ReadAckWriteMask           uint64
}

type WHEA_DEVICE_DRIVER_DESCRIPTOR struct {
	Type                 uint16
	Enabled              BOOLEAN
	Reserved             byte
	SourceGuid           syscall.GUID
	LogTag               uint16
	Reserved2            uint16
	PacketLength         uint32
	PacketCount          uint32
	PacketBuffer         *byte
	Config               WHEA_ERROR_SOURCE_CONFIGURATION_DD
	CreatorId            syscall.GUID
	PartitionId          syscall.GUID
	MaxSectionDataLength uint32
	MaxSectionsPerRecord uint32
	PacketStateBuffer    *byte
	OpenHandles          int32
}

type WHEA_IPF_MCA_DESCRIPTOR struct {
	Type     uint16
	Enabled  byte
	Reserved byte
}

type WHEA_IPF_CMC_DESCRIPTOR struct {
	Type     uint16
	Enabled  byte
	Reserved byte
}

type WHEA_IPF_CPE_DESCRIPTOR struct {
	Type     uint16
	Enabled  byte
	Reserved byte
}

type WHEA_ERROR_SOURCE_DESCRIPTOR_Info struct {
	Data [133]uint64
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) XpfMceDescriptor() *WHEA_XPF_MCE_DESCRIPTOR {
	return (*WHEA_XPF_MCE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) XpfMceDescriptorVal() WHEA_XPF_MCE_DESCRIPTOR {
	return *(*WHEA_XPF_MCE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) XpfCmcDescriptor() *WHEA_XPF_CMC_DESCRIPTOR {
	return (*WHEA_XPF_CMC_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) XpfCmcDescriptorVal() WHEA_XPF_CMC_DESCRIPTOR {
	return *(*WHEA_XPF_CMC_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) XpfNmiDescriptor() *WHEA_XPF_NMI_DESCRIPTOR {
	return (*WHEA_XPF_NMI_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) XpfNmiDescriptorVal() WHEA_XPF_NMI_DESCRIPTOR {
	return *(*WHEA_XPF_NMI_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) IpfMcaDescriptor() *WHEA_IPF_MCA_DESCRIPTOR {
	return (*WHEA_IPF_MCA_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) IpfMcaDescriptorVal() WHEA_IPF_MCA_DESCRIPTOR {
	return *(*WHEA_IPF_MCA_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) IpfCmcDescriptor() *WHEA_IPF_CMC_DESCRIPTOR {
	return (*WHEA_IPF_CMC_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) IpfCmcDescriptorVal() WHEA_IPF_CMC_DESCRIPTOR {
	return *(*WHEA_IPF_CMC_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) IpfCpeDescriptor() *WHEA_IPF_CPE_DESCRIPTOR {
	return (*WHEA_IPF_CPE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) IpfCpeDescriptorVal() WHEA_IPF_CPE_DESCRIPTOR {
	return *(*WHEA_IPF_CPE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) AerRootportDescriptor() *WHEA_AER_ROOTPORT_DESCRIPTOR {
	return (*WHEA_AER_ROOTPORT_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) AerRootportDescriptorVal() WHEA_AER_ROOTPORT_DESCRIPTOR {
	return *(*WHEA_AER_ROOTPORT_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) AerEndpointDescriptor() *WHEA_AER_ENDPOINT_DESCRIPTOR {
	return (*WHEA_AER_ENDPOINT_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) AerEndpointDescriptorVal() WHEA_AER_ENDPOINT_DESCRIPTOR {
	return *(*WHEA_AER_ENDPOINT_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) AerBridgeDescriptor() *WHEA_AER_BRIDGE_DESCRIPTOR {
	return (*WHEA_AER_BRIDGE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) AerBridgeDescriptorVal() WHEA_AER_BRIDGE_DESCRIPTOR {
	return *(*WHEA_AER_BRIDGE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) GenErrDescriptor() *WHEA_GENERIC_ERROR_DESCRIPTOR {
	return (*WHEA_GENERIC_ERROR_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) GenErrDescriptorVal() WHEA_GENERIC_ERROR_DESCRIPTOR {
	return *(*WHEA_GENERIC_ERROR_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) GenErrDescriptorV2() *WHEA_GENERIC_ERROR_DESCRIPTOR_V2 {
	return (*WHEA_GENERIC_ERROR_DESCRIPTOR_V2)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) GenErrDescriptorV2Val() WHEA_GENERIC_ERROR_DESCRIPTOR_V2 {
	return *(*WHEA_GENERIC_ERROR_DESCRIPTOR_V2)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) DeviceDriverDescriptor() *WHEA_DEVICE_DRIVER_DESCRIPTOR {
	return (*WHEA_DEVICE_DRIVER_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *WHEA_ERROR_SOURCE_DESCRIPTOR_Info) DeviceDriverDescriptorVal() WHEA_DEVICE_DRIVER_DESCRIPTOR {
	return *(*WHEA_DEVICE_DRIVER_DESCRIPTOR)(unsafe.Pointer(this))
}

type WHEA_ERROR_SOURCE_DESCRIPTOR struct {
	Length                  uint32
	Version                 uint32
	Type                    WHEA_ERROR_SOURCE_TYPE
	State                   WHEA_ERROR_SOURCE_STATE
	MaxRawDataLength        uint32
	NumRecordsToPreallocate uint32
	MaxSectionsPerRecord    uint32
	ErrorSourceId           uint32
	PlatformErrorSourceId   uint32
	Flags                   uint32
	Info                    WHEA_ERROR_SOURCE_DESCRIPTOR_Info
}

type IPMI_OS_SEL_RECORD struct {
	Signature  uint32
	Version    uint32
	Length     uint32
	RecordType IPMI_OS_SEL_RECORD_TYPE
	DataLength uint32
	Data       [1]byte
}

type DebugPropertyInfo struct {
	M_dwValidFields uint32
	M_bstrName      BSTR
	M_bstrType      BSTR
	M_bstrValue     BSTR
	M_bstrFullName  BSTR
	M_dwAttrib      uint32
	M_pDebugProp    *IDebugProperty
}

type ExtendedDebugPropertyInfo struct {
	DwValidFields uint32
	PszName       PWSTR
	PszType       PWSTR
	PszValue      PWSTR
	PszFullName   PWSTR
	DwAttrib      uint32
	PDebugProp    *IDebugProperty
	NDISPID       uint32
	NType         uint32
	VarValue      VARIANT
	PlbValue      *ILockBytes
	PDebugExtProp *IDebugExtendedProperty
}

// func types

type PGET_RUNTIME_FUNCTION_CALLBACK = uintptr
type PGET_RUNTIME_FUNCTION_CALLBACK_func = func(ControlPc uint64, Context unsafe.Pointer) *IMAGE_RUNTIME_FUNCTION_ENTRY

type PVECTORED_EXCEPTION_HANDLER = uintptr
type PVECTORED_EXCEPTION_HANDLER_func = func(ExceptionInfo *EXCEPTION_POINTERS) int32

type LPTOP_LEVEL_EXCEPTION_FILTER = uintptr
type LPTOP_LEVEL_EXCEPTION_FILTER_func = func(ExceptionInfo *EXCEPTION_POINTERS) int32

type PWAITCHAINCALLBACK = uintptr
type PWAITCHAINCALLBACK_func = func(WctHandle unsafe.Pointer, Context uintptr, CallbackStatus uint32, NodeCount *uint32, NodeInfoArray *WAITCHAIN_NODE_INFO, IsCycle *int32)

type PCOGETCALLSTATE = uintptr
type PCOGETCALLSTATE_func = func(param0 int32, param1 *uint32) HRESULT

type PCOGETACTIVATIONSTATE = uintptr
type PCOGETACTIVATIONSTATE_func = func(param0 syscall.GUID, param1 uint32, param2 *uint32) HRESULT

type MINIDUMP_CALLBACK_ROUTINE = uintptr
type MINIDUMP_CALLBACK_ROUTINE_func = func(CallbackParam unsafe.Pointer, CallbackInput *MINIDUMP_CALLBACK_INPUT, CallbackOutput *MINIDUMP_CALLBACK_OUTPUT) BOOL

type PIMAGEHLP_STATUS_ROUTINE = uintptr
type PIMAGEHLP_STATUS_ROUTINE_func = func(Reason IMAGEHLP_STATUS_REASON, ImageName PSTR, DllName PSTR, Va uintptr, Parameter uintptr) BOOL

type PIMAGEHLP_STATUS_ROUTINE32 = uintptr
type PIMAGEHLP_STATUS_ROUTINE32_func = func(Reason IMAGEHLP_STATUS_REASON, ImageName PSTR, DllName PSTR, Va uint32, Parameter uintptr) BOOL

type PIMAGEHLP_STATUS_ROUTINE64 = uintptr
type PIMAGEHLP_STATUS_ROUTINE64_func = func(Reason IMAGEHLP_STATUS_REASON, ImageName PSTR, DllName PSTR, Va uint64, Parameter uintptr) BOOL

type DIGEST_FUNCTION = uintptr
type DIGEST_FUNCTION_func = func(refdata unsafe.Pointer, pData *byte, dwLength uint32) BOOL

type PFIND_DEBUG_FILE_CALLBACK = uintptr
type PFIND_DEBUG_FILE_CALLBACK_func = func(FileHandle HANDLE, FileName PSTR, CallerData unsafe.Pointer) BOOL

type PFIND_DEBUG_FILE_CALLBACKW = uintptr
type PFIND_DEBUG_FILE_CALLBACKW_func = func(FileHandle HANDLE, FileName PWSTR, CallerData unsafe.Pointer) BOOL

type PFINDFILEINPATHCALLBACK = uintptr
type PFINDFILEINPATHCALLBACK_func = func(filename PSTR, context unsafe.Pointer) BOOL

type PFINDFILEINPATHCALLBACKW = uintptr
type PFINDFILEINPATHCALLBACKW_func = func(filename PWSTR, context unsafe.Pointer) BOOL

type PFIND_EXE_FILE_CALLBACK = uintptr
type PFIND_EXE_FILE_CALLBACK_func = func(FileHandle HANDLE, FileName PSTR, CallerData unsafe.Pointer) BOOL

type PFIND_EXE_FILE_CALLBACKW = uintptr
type PFIND_EXE_FILE_CALLBACKW_func = func(FileHandle HANDLE, FileName PWSTR, CallerData unsafe.Pointer) BOOL

type PENUMDIRTREE_CALLBACK = uintptr
type PENUMDIRTREE_CALLBACK_func = func(FilePath PSTR, CallerData unsafe.Pointer) BOOL

type PENUMDIRTREE_CALLBACKW = uintptr
type PENUMDIRTREE_CALLBACKW_func = func(FilePath PWSTR, CallerData unsafe.Pointer) BOOL

type PREAD_PROCESS_MEMORY_ROUTINE64 = uintptr
type PREAD_PROCESS_MEMORY_ROUTINE64_func = func(hProcess HANDLE, qwBaseAddress uint64, lpBuffer unsafe.Pointer, nSize uint32, lpNumberOfBytesRead *uint32) BOOL

type PFUNCTION_TABLE_ACCESS_ROUTINE64 = uintptr
type PFUNCTION_TABLE_ACCESS_ROUTINE64_func = func(ahProcess HANDLE, AddrBase uint64) unsafe.Pointer

type PGET_MODULE_BASE_ROUTINE64 = uintptr
type PGET_MODULE_BASE_ROUTINE64_func = func(hProcess HANDLE, Address uint64) uint64

type PTRANSLATE_ADDRESS_ROUTINE64 = uintptr
type PTRANSLATE_ADDRESS_ROUTINE64_func = func(hProcess HANDLE, hThread HANDLE, lpaddr *ADDRESS64) uint64

type PGET_TARGET_ATTRIBUTE_VALUE64 = uintptr
type PGET_TARGET_ATTRIBUTE_VALUE64_func = func(hProcess HANDLE, Attribute uint32, AttributeData uint64, AttributeValue *uint64) BOOL

type PSYM_ENUMMODULES_CALLBACK64 = uintptr
type PSYM_ENUMMODULES_CALLBACK64_func = func(ModuleName PSTR, BaseOfDll uint64, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMMODULES_CALLBACKW64 = uintptr
type PSYM_ENUMMODULES_CALLBACKW64_func = func(ModuleName PWSTR, BaseOfDll uint64, UserContext unsafe.Pointer) BOOL

type PENUMLOADED_MODULES_CALLBACK64 = uintptr
type PENUMLOADED_MODULES_CALLBACK64_func = func(ModuleName PSTR, ModuleBase uint64, ModuleSize uint32, UserContext unsafe.Pointer) BOOL

type PENUMLOADED_MODULES_CALLBACKW64 = uintptr
type PENUMLOADED_MODULES_CALLBACKW64_func = func(ModuleName PWSTR, ModuleBase uint64, ModuleSize uint32, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMSYMBOLS_CALLBACK64 = uintptr
type PSYM_ENUMSYMBOLS_CALLBACK64_func = func(SymbolName PSTR, SymbolAddress uint64, SymbolSize uint32, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMSYMBOLS_CALLBACK64W = uintptr
type PSYM_ENUMSYMBOLS_CALLBACK64W_func = func(SymbolName PWSTR, SymbolAddress uint64, SymbolSize uint32, UserContext unsafe.Pointer) BOOL

type PSYMBOL_REGISTERED_CALLBACK64 = uintptr
type PSYMBOL_REGISTERED_CALLBACK64_func = func(hProcess HANDLE, ActionCode uint32, CallbackData uint64, UserContext uint64) BOOL

type PSYMBOL_FUNCENTRY_CALLBACK = uintptr
type PSYMBOL_FUNCENTRY_CALLBACK_func = func(hProcess HANDLE, AddrBase uint32, UserContext unsafe.Pointer) unsafe.Pointer

type PSYMBOL_FUNCENTRY_CALLBACK64 = uintptr
type PSYMBOL_FUNCENTRY_CALLBACK64_func = func(hProcess HANDLE, AddrBase uint64, UserContext uint64) unsafe.Pointer

type PSYM_ENUMSOURCEFILES_CALLBACK = uintptr
type PSYM_ENUMSOURCEFILES_CALLBACK_func = func(pSourceFile *SOURCEFILE, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMSOURCEFILES_CALLBACKW = uintptr
type PSYM_ENUMSOURCEFILES_CALLBACKW_func = func(pSourceFile *SOURCEFILEW, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMLINES_CALLBACK = uintptr
type PSYM_ENUMLINES_CALLBACK_func = func(LineInfo *SRCCODEINFO, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMLINES_CALLBACKW = uintptr
type PSYM_ENUMLINES_CALLBACKW_func = func(LineInfo *SRCCODEINFOW, UserContext unsafe.Pointer) BOOL

type PENUMSOURCEFILETOKENSCALLBACK = uintptr
type PENUMSOURCEFILETOKENSCALLBACK_func = func(token unsafe.Pointer, size uintptr) BOOL

type PSYM_ENUMPROCESSES_CALLBACK = uintptr
type PSYM_ENUMPROCESSES_CALLBACK_func = func(hProcess HANDLE, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMERATESYMBOLS_CALLBACK = uintptr
type PSYM_ENUMERATESYMBOLS_CALLBACK_func = func(pSymInfo *SYMBOL_INFO, SymbolSize uint32, UserContext unsafe.Pointer) BOOL

type PSYM_ENUMERATESYMBOLS_CALLBACKW = uintptr
type PSYM_ENUMERATESYMBOLS_CALLBACKW_func = func(pSymInfo *SYMBOL_INFOW, SymbolSize uint32, UserContext unsafe.Pointer) BOOL

type SYMADDSOURCESTREAM = uintptr
type SYMADDSOURCESTREAM_func = func(param0 HANDLE, param1 uint64, param2 PSTR, param3 *byte, param4 uintptr) BOOL

type SYMADDSOURCESTREAMA = uintptr
type SYMADDSOURCESTREAMA_func = func(param0 HANDLE, param1 uint64, param2 PSTR, param3 *byte, param4 uintptr) BOOL

type PDBGHELP_CREATE_USER_DUMP_CALLBACK = uintptr
type PDBGHELP_CREATE_USER_DUMP_CALLBACK_func = func(DataType uint32, Data unsafe.Pointer, DataLength *uint32, UserData unsafe.Pointer) BOOL

type PSYMBOLSERVERPROC = uintptr
type PSYMBOLSERVERPROC_func = func(param0 PSTR, param1 PSTR, param2 unsafe.Pointer, param3 uint32, param4 uint32, param5 PSTR) BOOL

type PSYMBOLSERVERPROCA = uintptr
type PSYMBOLSERVERPROCA_func = func(param0 PSTR, param1 PSTR, param2 unsafe.Pointer, param3 uint32, param4 uint32, param5 PSTR) BOOL

type PSYMBOLSERVERPROCW = uintptr
type PSYMBOLSERVERPROCW_func = func(param0 PWSTR, param1 PWSTR, param2 unsafe.Pointer, param3 uint32, param4 uint32, param5 PWSTR) BOOL

type PSYMBOLSERVERBYINDEXPROC = uintptr
type PSYMBOLSERVERBYINDEXPROC_func = func(param0 PSTR, param1 PSTR, param2 PSTR, param3 PSTR) BOOL

type PSYMBOLSERVERBYINDEXPROCA = uintptr
type PSYMBOLSERVERBYINDEXPROCA_func = func(param0 PSTR, param1 PSTR, param2 PSTR, param3 PSTR) BOOL

type PSYMBOLSERVERBYINDEXPROCW = uintptr
type PSYMBOLSERVERBYINDEXPROCW_func = func(param0 PWSTR, param1 PWSTR, param2 PWSTR, param3 PWSTR) BOOL

type PSYMBOLSERVEROPENPROC = uintptr
type PSYMBOLSERVEROPENPROC_func = func() BOOL

type PSYMBOLSERVERCLOSEPROC = uintptr
type PSYMBOLSERVERCLOSEPROC_func = func() BOOL

type PSYMBOLSERVERSETOPTIONSPROC = uintptr
type PSYMBOLSERVERSETOPTIONSPROC_func = func(param0 uintptr, param1 uint64) BOOL

type PSYMBOLSERVERSETOPTIONSWPROC = uintptr
type PSYMBOLSERVERSETOPTIONSWPROC_func = func(param0 uintptr, param1 uint64) BOOL

type PSYMBOLSERVERCALLBACKPROC = uintptr
type PSYMBOLSERVERCALLBACKPROC_func = func(action uintptr, data uint64, context uint64) BOOL

type PSYMBOLSERVERGETOPTIONSPROC = uintptr
type PSYMBOLSERVERGETOPTIONSPROC_func = func() uintptr

type PSYMBOLSERVERPINGPROC = uintptr
type PSYMBOLSERVERPINGPROC_func = func(param0 PSTR) BOOL

type PSYMBOLSERVERPINGPROCA = uintptr
type PSYMBOLSERVERPINGPROCA_func = func(param0 PSTR) BOOL

type PSYMBOLSERVERPINGPROCW = uintptr
type PSYMBOLSERVERPINGPROCW_func = func(param0 PWSTR) BOOL

type PSYMBOLSERVERGETVERSION = uintptr
type PSYMBOLSERVERGETVERSION_func = func(param0 *API_VERSION) BOOL

type PSYMBOLSERVERDELTANAME = uintptr
type PSYMBOLSERVERDELTANAME_func = func(param0 PSTR, param1 unsafe.Pointer, param2 uint32, param3 uint32, param4 unsafe.Pointer, param5 uint32, param6 uint32, param7 PSTR, param8 uintptr) BOOL

type PSYMBOLSERVERDELTANAMEW = uintptr
type PSYMBOLSERVERDELTANAMEW_func = func(param0 PWSTR, param1 unsafe.Pointer, param2 uint32, param3 uint32, param4 unsafe.Pointer, param5 uint32, param6 uint32, param7 PWSTR, param8 uintptr) BOOL

type PSYMBOLSERVERGETSUPPLEMENT = uintptr
type PSYMBOLSERVERGETSUPPLEMENT_func = func(param0 PSTR, param1 PSTR, param2 PSTR, param3 PSTR, param4 uintptr) BOOL

type PSYMBOLSERVERGETSUPPLEMENTW = uintptr
type PSYMBOLSERVERGETSUPPLEMENTW_func = func(param0 PWSTR, param1 PWSTR, param2 PWSTR, param3 PWSTR, param4 uintptr) BOOL

type PSYMBOLSERVERSTORESUPPLEMENT = uintptr
type PSYMBOLSERVERSTORESUPPLEMENT_func = func(param0 PSTR, param1 PSTR, param2 PSTR, param3 PSTR, param4 uintptr, param5 uint32) BOOL

type PSYMBOLSERVERSTORESUPPLEMENTW = uintptr
type PSYMBOLSERVERSTORESUPPLEMENTW_func = func(param0 PWSTR, param1 PWSTR, param2 PWSTR, param3 PWSTR, param4 uintptr, param5 uint32) BOOL

type PSYMBOLSERVERGETINDEXSTRING = uintptr
type PSYMBOLSERVERGETINDEXSTRING_func = func(param0 unsafe.Pointer, param1 uint32, param2 uint32, param3 PSTR, param4 uintptr) BOOL

type PSYMBOLSERVERGETINDEXSTRINGW = uintptr
type PSYMBOLSERVERGETINDEXSTRINGW_func = func(param0 unsafe.Pointer, param1 uint32, param2 uint32, param3 PWSTR, param4 uintptr) BOOL

type PSYMBOLSERVERSTOREFILE = uintptr
type PSYMBOLSERVERSTOREFILE_func = func(param0 PSTR, param1 PSTR, param2 unsafe.Pointer, param3 uint32, param4 uint32, param5 PSTR, param6 uintptr, param7 uint32) BOOL

type PSYMBOLSERVERSTOREFILEW = uintptr
type PSYMBOLSERVERSTOREFILEW_func = func(param0 PWSTR, param1 PWSTR, param2 unsafe.Pointer, param3 uint32, param4 uint32, param5 PWSTR, param6 uintptr, param7 uint32) BOOL

type PSYMBOLSERVERISSTORE = uintptr
type PSYMBOLSERVERISSTORE_func = func(param0 PSTR) BOOL

type PSYMBOLSERVERISSTOREW = uintptr
type PSYMBOLSERVERISSTOREW_func = func(param0 PWSTR) BOOL

type PSYMBOLSERVERVERSION = uintptr
type PSYMBOLSERVERVERSION_func = func() uint32

type PSYMBOLSERVERMESSAGEPROC = uintptr
type PSYMBOLSERVERMESSAGEPROC_func = func(action uintptr, data uint64, context uint64) BOOL

type PSYMBOLSERVERWEXPROC = uintptr
type PSYMBOLSERVERWEXPROC_func = func(param0 PWSTR, param1 PWSTR, param2 unsafe.Pointer, param3 uint32, param4 uint32, param5 PWSTR, param6 *SYMSRV_EXTENDED_OUTPUT_DATA) BOOL

type PSYMBOLSERVERPINGPROCWEX = uintptr
type PSYMBOLSERVERPINGPROCWEX_func = func(param0 PWSTR) BOOL

type PSYMBOLSERVERGETOPTIONDATAPROC = uintptr
type PSYMBOLSERVERGETOPTIONDATAPROC_func = func(param0 uintptr, param1 *uint64) BOOL

type PSYMBOLSERVERSETHTTPAUTHHEADER = uintptr
type PSYMBOLSERVERSETHTTPAUTHHEADER_func = func(pszAuthHeader PWSTR) BOOL

type LPCALL_BACK_USER_INTERRUPT_ROUTINE = uintptr
type LPCALL_BACK_USER_INTERRUPT_ROUTINE_func = func() uint32

type WHEA_ERROR_SOURCE_INITIALIZE_DEVICE_DRIVER = uintptr
type WHEA_ERROR_SOURCE_INITIALIZE_DEVICE_DRIVER_func = func(Context unsafe.Pointer, ErrorSourceId uint32) NTSTATUS

type WHEA_ERROR_SOURCE_UNINITIALIZE_DEVICE_DRIVER = uintptr
type WHEA_ERROR_SOURCE_UNINITIALIZE_DEVICE_DRIVER_func = func(Context unsafe.Pointer)

type WHEA_ERROR_SOURCE_CORRECT_DEVICE_DRIVER = uintptr
type WHEA_ERROR_SOURCE_CORRECT_DEVICE_DRIVER_func = func(ErrorSourceDesc unsafe.Pointer, MaximumSectionLength *uint32) NTSTATUS

// interfaces

// CB5BDC81-93C1-11CF-8F20-00805F2CD064
var IID_IObjectSafety = syscall.GUID{0xCB5BDC81, 0x93C1, 0x11CF,
	[8]byte{0x8F, 0x20, 0x00, 0x80, 0x5F, 0x2C, 0xD0, 0x64}}

type IObjectSafetyInterface interface {
	IUnknownInterface
	GetInterfaceSafetyOptions(riid *syscall.GUID, pdwSupportedOptions *uint32, pdwEnabledOptions *uint32) HRESULT
	SetInterfaceSafetyOptions(riid *syscall.GUID, dwOptionSetMask uint32, dwEnabledOptions uint32) HRESULT
}

type IObjectSafetyVtbl struct {
	IUnknownVtbl
	GetInterfaceSafetyOptions uintptr
	SetInterfaceSafetyOptions uintptr
}

type IObjectSafety struct {
	IUnknown
}

func (this *IObjectSafety) Vtbl() *IObjectSafetyVtbl {
	return (*IObjectSafetyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObjectSafety) GetInterfaceSafetyOptions(riid *syscall.GUID, pdwSupportedOptions *uint32, pdwEnabledOptions *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetInterfaceSafetyOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(pdwSupportedOptions)), uintptr(unsafe.Pointer(pdwEnabledOptions)))
	return HRESULT(ret)
}

func (this *IObjectSafety) SetInterfaceSafetyOptions(riid *syscall.GUID, dwOptionSetMask uint32, dwEnabledOptions uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetInterfaceSafetyOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(riid)), uintptr(dwOptionSetMask), uintptr(dwEnabledOptions))
	return HRESULT(ret)
}

// 51973C50-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugProperty = syscall.GUID{0x51973C50, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugPropertyInterface interface {
	IUnknownInterface
	GetPropertyInfo(dwFieldSpec uint32, nRadix uint32, pPropertyInfo *DebugPropertyInfo) HRESULT
	GetExtendedInfo(cInfos uint32, rgguidExtendedInfo *syscall.GUID, rgvar *VARIANT) HRESULT
	SetValueAsString(pszValue PWSTR, nRadix uint32) HRESULT
	EnumMembers(dwFieldSpec uint32, nRadix uint32, refiid *syscall.GUID, ppepi **IEnumDebugPropertyInfo) HRESULT
	GetParent(ppDebugProp **IDebugProperty) HRESULT
}

type IDebugPropertyVtbl struct {
	IUnknownVtbl
	GetPropertyInfo  uintptr
	GetExtendedInfo  uintptr
	SetValueAsString uintptr
	EnumMembers      uintptr
	GetParent        uintptr
}

type IDebugProperty struct {
	IUnknown
}

func (this *IDebugProperty) Vtbl() *IDebugPropertyVtbl {
	return (*IDebugPropertyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDebugProperty) GetPropertyInfo(dwFieldSpec uint32, nRadix uint32, pPropertyInfo *DebugPropertyInfo) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyInfo, uintptr(unsafe.Pointer(this)), uintptr(dwFieldSpec), uintptr(nRadix), uintptr(unsafe.Pointer(pPropertyInfo)))
	return HRESULT(ret)
}

func (this *IDebugProperty) GetExtendedInfo(cInfos uint32, rgguidExtendedInfo *syscall.GUID, rgvar *VARIANT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetExtendedInfo, uintptr(unsafe.Pointer(this)), uintptr(cInfos), uintptr(unsafe.Pointer(rgguidExtendedInfo)), uintptr(unsafe.Pointer(rgvar)))
	return HRESULT(ret)
}

func (this *IDebugProperty) SetValueAsString(pszValue PWSTR, nRadix uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValueAsString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszValue)), uintptr(nRadix))
	return HRESULT(ret)
}

func (this *IDebugProperty) EnumMembers(dwFieldSpec uint32, nRadix uint32, refiid *syscall.GUID, ppepi **IEnumDebugPropertyInfo) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().EnumMembers, uintptr(unsafe.Pointer(this)), uintptr(dwFieldSpec), uintptr(nRadix), uintptr(unsafe.Pointer(refiid)), uintptr(unsafe.Pointer(ppepi)))
	return HRESULT(ret)
}

func (this *IDebugProperty) GetParent(ppDebugProp **IDebugProperty) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetParent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppDebugProp)))
	return HRESULT(ret)
}

// 51973C51-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IEnumDebugPropertyInfo = syscall.GUID{0x51973C51, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IEnumDebugPropertyInfoInterface interface {
	IUnknownInterface
	Next(celt uint32, pi *DebugPropertyInfo, pcEltsfetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppepi **IEnumDebugPropertyInfo) HRESULT
	GetCount(pcelt *uint32) HRESULT
}

type IEnumDebugPropertyInfoVtbl struct {
	IUnknownVtbl
	Next     uintptr
	Skip     uintptr
	Reset    uintptr
	Clone    uintptr
	GetCount uintptr
}

type IEnumDebugPropertyInfo struct {
	IUnknown
}

func (this *IEnumDebugPropertyInfo) Vtbl() *IEnumDebugPropertyInfoVtbl {
	return (*IEnumDebugPropertyInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumDebugPropertyInfo) Next(celt uint32, pi *DebugPropertyInfo, pcEltsfetched *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(pi)), uintptr(unsafe.Pointer(pcEltsfetched)))
	return HRESULT(ret)
}

func (this *IEnumDebugPropertyInfo) Skip(celt uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumDebugPropertyInfo) Reset() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumDebugPropertyInfo) Clone(ppepi **IEnumDebugPropertyInfo) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppepi)))
	return HRESULT(ret)
}

func (this *IEnumDebugPropertyInfo) GetCount(pcelt *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcelt)))
	return HRESULT(ret)
}

// 51973C52-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugExtendedProperty = syscall.GUID{0x51973C52, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugExtendedPropertyInterface interface {
	IDebugPropertyInterface
	GetExtendedPropertyInfo(dwFieldSpec uint32, nRadix uint32, pExtendedPropertyInfo *ExtendedDebugPropertyInfo) HRESULT
	EnumExtendedMembers(dwFieldSpec uint32, nRadix uint32, ppeepi **IEnumDebugExtendedPropertyInfo) HRESULT
}

type IDebugExtendedPropertyVtbl struct {
	IDebugPropertyVtbl
	GetExtendedPropertyInfo uintptr
	EnumExtendedMembers     uintptr
}

type IDebugExtendedProperty struct {
	IDebugProperty
}

func (this *IDebugExtendedProperty) Vtbl() *IDebugExtendedPropertyVtbl {
	return (*IDebugExtendedPropertyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDebugExtendedProperty) GetExtendedPropertyInfo(dwFieldSpec uint32, nRadix uint32, pExtendedPropertyInfo *ExtendedDebugPropertyInfo) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetExtendedPropertyInfo, uintptr(unsafe.Pointer(this)), uintptr(dwFieldSpec), uintptr(nRadix), uintptr(unsafe.Pointer(pExtendedPropertyInfo)))
	return HRESULT(ret)
}

func (this *IDebugExtendedProperty) EnumExtendedMembers(dwFieldSpec uint32, nRadix uint32, ppeepi **IEnumDebugExtendedPropertyInfo) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().EnumExtendedMembers, uintptr(unsafe.Pointer(this)), uintptr(dwFieldSpec), uintptr(nRadix), uintptr(unsafe.Pointer(ppeepi)))
	return HRESULT(ret)
}

// 51973C53-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IEnumDebugExtendedPropertyInfo = syscall.GUID{0x51973C53, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IEnumDebugExtendedPropertyInfoInterface interface {
	IUnknownInterface
	Next(celt uint32, rgExtendedPropertyInfo *ExtendedDebugPropertyInfo, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(pedpe **IEnumDebugExtendedPropertyInfo) HRESULT
	GetCount(pcelt *uint32) HRESULT
}

type IEnumDebugExtendedPropertyInfoVtbl struct {
	IUnknownVtbl
	Next     uintptr
	Skip     uintptr
	Reset    uintptr
	Clone    uintptr
	GetCount uintptr
}

type IEnumDebugExtendedPropertyInfo struct {
	IUnknown
}

func (this *IEnumDebugExtendedPropertyInfo) Vtbl() *IEnumDebugExtendedPropertyInfoVtbl {
	return (*IEnumDebugExtendedPropertyInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumDebugExtendedPropertyInfo) Next(celt uint32, rgExtendedPropertyInfo *ExtendedDebugPropertyInfo, pceltFetched *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgExtendedPropertyInfo)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumDebugExtendedPropertyInfo) Skip(celt uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumDebugExtendedPropertyInfo) Reset() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumDebugExtendedPropertyInfo) Clone(pedpe **IEnumDebugExtendedPropertyInfo) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pedpe)))
	return HRESULT(ret)
}

func (this *IEnumDebugExtendedPropertyInfo) GetCount(pcelt *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcelt)))
	return HRESULT(ret)
}

// 51973C54-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IPerPropertyBrowsing2 = syscall.GUID{0x51973C54, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IPerPropertyBrowsing2Interface interface {
	IUnknownInterface
	GetDisplayString(dispid int32, pBstr *BSTR) HRESULT
	MapPropertyToPage(dispid int32, pClsidPropPage *syscall.GUID) HRESULT
	GetPredefinedStrings(dispid int32, pCaStrings *CALPOLESTR, pCaCookies *CADWORD) HRESULT
	SetPredefinedValue(dispid int32, dwCookie uint32) HRESULT
}

type IPerPropertyBrowsing2Vtbl struct {
	IUnknownVtbl
	GetDisplayString     uintptr
	MapPropertyToPage    uintptr
	GetPredefinedStrings uintptr
	SetPredefinedValue   uintptr
}

type IPerPropertyBrowsing2 struct {
	IUnknown
}

func (this *IPerPropertyBrowsing2) Vtbl() *IPerPropertyBrowsing2Vtbl {
	return (*IPerPropertyBrowsing2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPerPropertyBrowsing2) GetDisplayString(dispid int32, pBstr *BSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayString, uintptr(unsafe.Pointer(this)), uintptr(dispid), uintptr(unsafe.Pointer(pBstr)))
	return HRESULT(ret)
}

func (this *IPerPropertyBrowsing2) MapPropertyToPage(dispid int32, pClsidPropPage *syscall.GUID) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().MapPropertyToPage, uintptr(unsafe.Pointer(this)), uintptr(dispid), uintptr(unsafe.Pointer(pClsidPropPage)))
	return HRESULT(ret)
}

func (this *IPerPropertyBrowsing2) GetPredefinedStrings(dispid int32, pCaStrings *CALPOLESTR, pCaCookies *CADWORD) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPredefinedStrings, uintptr(unsafe.Pointer(this)), uintptr(dispid), uintptr(unsafe.Pointer(pCaStrings)), uintptr(unsafe.Pointer(pCaCookies)))
	return HRESULT(ret)
}

func (this *IPerPropertyBrowsing2) SetPredefinedValue(dispid int32, dwCookie uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPredefinedValue, uintptr(unsafe.Pointer(this)), uintptr(dispid), uintptr(dwCookie))
	return HRESULT(ret)
}

// 51973C55-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugPropertyEnumType_All = syscall.GUID{0x51973C55, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugPropertyEnumType_AllInterface interface {
	IUnknownInterface
	GetName(__MIDL__IDebugPropertyEnumType_All0000 *BSTR) HRESULT
}

type IDebugPropertyEnumType_AllVtbl struct {
	IUnknownVtbl
	GetName uintptr
}

type IDebugPropertyEnumType_All struct {
	IUnknown
}

func (this *IDebugPropertyEnumType_All) Vtbl() *IDebugPropertyEnumType_AllVtbl {
	return (*IDebugPropertyEnumType_AllVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDebugPropertyEnumType_All) GetName(__MIDL__IDebugPropertyEnumType_All0000 *BSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(__MIDL__IDebugPropertyEnumType_All0000)))
	return HRESULT(ret)
}

// 51973C56-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugPropertyEnumType_Locals = syscall.GUID{0x51973C56, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugPropertyEnumType_LocalsInterface interface {
	IDebugPropertyEnumType_AllInterface
}

type IDebugPropertyEnumType_LocalsVtbl struct {
	IDebugPropertyEnumType_AllVtbl
}

type IDebugPropertyEnumType_Locals struct {
	IDebugPropertyEnumType_All
}

func (this *IDebugPropertyEnumType_Locals) Vtbl() *IDebugPropertyEnumType_LocalsVtbl {
	return (*IDebugPropertyEnumType_LocalsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 51973C57-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugPropertyEnumType_Arguments = syscall.GUID{0x51973C57, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugPropertyEnumType_ArgumentsInterface interface {
	IDebugPropertyEnumType_AllInterface
}

type IDebugPropertyEnumType_ArgumentsVtbl struct {
	IDebugPropertyEnumType_AllVtbl
}

type IDebugPropertyEnumType_Arguments struct {
	IDebugPropertyEnumType_All
}

func (this *IDebugPropertyEnumType_Arguments) Vtbl() *IDebugPropertyEnumType_ArgumentsVtbl {
	return (*IDebugPropertyEnumType_ArgumentsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 51973C58-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugPropertyEnumType_LocalsPlusArgs = syscall.GUID{0x51973C58, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugPropertyEnumType_LocalsPlusArgsInterface interface {
	IDebugPropertyEnumType_AllInterface
}

type IDebugPropertyEnumType_LocalsPlusArgsVtbl struct {
	IDebugPropertyEnumType_AllVtbl
}

type IDebugPropertyEnumType_LocalsPlusArgs struct {
	IDebugPropertyEnumType_All
}

func (this *IDebugPropertyEnumType_LocalsPlusArgs) Vtbl() *IDebugPropertyEnumType_LocalsPlusArgsVtbl {
	return (*IDebugPropertyEnumType_LocalsPlusArgsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 51973C59-CB0C-11D0-B5C9-00A0244A0E7A
var IID_IDebugPropertyEnumType_Registers = syscall.GUID{0x51973C59, 0xCB0C, 0x11D0,
	[8]byte{0xB5, 0xC9, 0x00, 0xA0, 0x24, 0x4A, 0x0E, 0x7A}}

type IDebugPropertyEnumType_RegistersInterface interface {
	IDebugPropertyEnumType_AllInterface
}

type IDebugPropertyEnumType_RegistersVtbl struct {
	IDebugPropertyEnumType_AllVtbl
}

type IDebugPropertyEnumType_Registers struct {
	IDebugPropertyEnumType_All
}

func (this *IDebugPropertyEnumType_Registers) Vtbl() *IDebugPropertyEnumType_RegistersVtbl {
	return (*IDebugPropertyEnumType_RegistersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

var (
	pReadProcessMemory               uintptr
	pWriteProcessMemory              uintptr
	pGetThreadContext                uintptr
	pSetThreadContext                uintptr
	pFlushInstructionCache           uintptr
	pWow64GetThreadContext           uintptr
	pWow64SetThreadContext           uintptr
	pRtlCaptureContext2              uintptr
	pRtlAddFunctionTable             uintptr
	pRtlDeleteFunctionTable          uintptr
	pRtlInstallFunctionTableCallback uintptr
	pRtlLookupFunctionEntry          uintptr
	pRtlUnwindEx                     uintptr
	pRtlVirtualUnwind                uintptr
	pRtlCaptureStackBackTrace        uintptr
	pRtlCaptureContext               uintptr
	pRtlUnwind                       uintptr
	pRtlRestoreContext               uintptr
	pRtlRaiseException               uintptr
	pRtlPcToFileHeader               uintptr
	pIsDebuggerPresent               uintptr
	pDebugBreak                      uintptr
	pOutputDebugStringA              uintptr
	pOutputDebugStringW              uintptr
	pContinueDebugEvent              uintptr
	pWaitForDebugEvent               uintptr
	pDebugActiveProcess              uintptr
	pDebugActiveProcessStop          uintptr
	pCheckRemoteDebuggerPresent      uintptr
	pWaitForDebugEventEx             uintptr
	pEncodePointer                   uintptr
	pDecodePointer                   uintptr
	pEncodeSystemPointer             uintptr
	pDecodeSystemPointer             uintptr
	pBeep                            uintptr
	pRaiseException                  uintptr
	pUnhandledExceptionFilter        uintptr
	pSetUnhandledExceptionFilter     uintptr
	pGetErrorMode                    uintptr
	pSetErrorMode                    uintptr
	pAddVectoredExceptionHandler     uintptr
	pRemoveVectoredExceptionHandler  uintptr
	pAddVectoredContinueHandler      uintptr
	pRemoveVectoredContinueHandler   uintptr
	pRaiseFailFastException          uintptr
	pFatalAppExitA                   uintptr
	pFatalAppExitW                   uintptr
	pGetThreadErrorMode              uintptr
	pSetThreadErrorMode              uintptr
	pOpenThreadWaitChainSession      uintptr
	pCloseThreadWaitChainSession     uintptr
	pGetThreadWaitChain              uintptr
	pRegisterWaitChainCOMCallback    uintptr
	pMessageBeep                     uintptr
	pFatalExit                       uintptr
	pGetThreadSelectorEntry          uintptr
	pWow64GetThreadSelectorEntry     uintptr
	pDebugSetProcessKillOnExit       uintptr
	pDebugBreakProcess               uintptr
	pFormatMessageA                  uintptr
	pFormatMessageW                  uintptr
	pCopyContext                     uintptr
	pInitializeContext               uintptr
	pInitializeContext2              uintptr
	pGetEnabledXStateFeatures        uintptr
	pGetXStateFeaturesMask           uintptr
	pLocateXStateFeature             uintptr
	pSetXStateFeaturesMask           uintptr
)

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress unsafe.Pointer, lpBuffer unsafe.Pointer, nSize uintptr, lpNumberOfBytesRead *uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pReadProcessMemory, libKernel32, "ReadProcessMemory")
	ret, _, err := syscall.SyscallN(addr, hProcess, uintptr(lpBaseAddress), uintptr(lpBuffer), nSize, uintptr(unsafe.Pointer(lpNumberOfBytesRead)))
	return BOOL(ret), WIN32_ERROR(err)
}

func WriteProcessMemory(hProcess HANDLE, lpBaseAddress unsafe.Pointer, lpBuffer unsafe.Pointer, nSize uintptr, lpNumberOfBytesWritten *uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pWriteProcessMemory, libKernel32, "WriteProcessMemory")
	ret, _, err := syscall.SyscallN(addr, hProcess, uintptr(lpBaseAddress), uintptr(lpBuffer), nSize, uintptr(unsafe.Pointer(lpNumberOfBytesWritten)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetThreadContext(hThread HANDLE, lpContext *CONTEXT) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetThreadContext, libKernel32, "GetThreadContext")
	ret, _, err := syscall.SyscallN(addr, hThread, uintptr(unsafe.Pointer(lpContext)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetThreadContext(hThread HANDLE, lpContext *CONTEXT) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetThreadContext, libKernel32, "SetThreadContext")
	ret, _, err := syscall.SyscallN(addr, hThread, uintptr(unsafe.Pointer(lpContext)))
	return BOOL(ret), WIN32_ERROR(err)
}

func FlushInstructionCache(hProcess HANDLE, lpBaseAddress unsafe.Pointer, dwSize uintptr) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pFlushInstructionCache, libKernel32, "FlushInstructionCache")
	ret, _, err := syscall.SyscallN(addr, hProcess, uintptr(lpBaseAddress), dwSize)
	return BOOL(ret), WIN32_ERROR(err)
}

func Wow64GetThreadContext(hThread HANDLE, lpContext *WOW64_CONTEXT) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pWow64GetThreadContext, libKernel32, "Wow64GetThreadContext")
	ret, _, err := syscall.SyscallN(addr, hThread, uintptr(unsafe.Pointer(lpContext)))
	return BOOL(ret), WIN32_ERROR(err)
}

func Wow64SetThreadContext(hThread HANDLE, lpContext *WOW64_CONTEXT) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pWow64SetThreadContext, libKernel32, "Wow64SetThreadContext")
	ret, _, err := syscall.SyscallN(addr, hThread, uintptr(unsafe.Pointer(lpContext)))
	return BOOL(ret), WIN32_ERROR(err)
}

func RtlCaptureContext2(ContextRecord *CONTEXT) {
	addr := LazyAddr(&pRtlCaptureContext2, libKernel32, "RtlCaptureContext2")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(ContextRecord)))
}

func RtlAddFunctionTable(FunctionTable *IMAGE_RUNTIME_FUNCTION_ENTRY, EntryCount uint32, BaseAddress uint64) BOOLEAN {
	addr := LazyAddr(&pRtlAddFunctionTable, libKernel32, "RtlAddFunctionTable")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(FunctionTable)), uintptr(EntryCount), uintptr(BaseAddress))
	return BOOLEAN(ret)
}

func RtlDeleteFunctionTable(FunctionTable *IMAGE_RUNTIME_FUNCTION_ENTRY) BOOLEAN {
	addr := LazyAddr(&pRtlDeleteFunctionTable, libKernel32, "RtlDeleteFunctionTable")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(FunctionTable)))
	return BOOLEAN(ret)
}

func RtlInstallFunctionTableCallback(TableIdentifier uint64, BaseAddress uint64, Length uint32, Callback PGET_RUNTIME_FUNCTION_CALLBACK, Context unsafe.Pointer, OutOfProcessCallbackDll PWSTR) BOOLEAN {
	addr := LazyAddr(&pRtlInstallFunctionTableCallback, libKernel32, "RtlInstallFunctionTableCallback")
	ret, _, _ := syscall.SyscallN(addr, uintptr(TableIdentifier), uintptr(BaseAddress), uintptr(Length), Callback, uintptr(Context), uintptr(unsafe.Pointer(OutOfProcessCallbackDll)))
	return BOOLEAN(ret)
}

func RtlLookupFunctionEntry(ControlPc uint64, ImageBase *uint64, HistoryTable *UNWIND_HISTORY_TABLE) *IMAGE_RUNTIME_FUNCTION_ENTRY {
	addr := LazyAddr(&pRtlLookupFunctionEntry, libKernel32, "RtlLookupFunctionEntry")
	ret, _, _ := syscall.SyscallN(addr, uintptr(ControlPc), uintptr(unsafe.Pointer(ImageBase)), uintptr(unsafe.Pointer(HistoryTable)))
	return (*IMAGE_RUNTIME_FUNCTION_ENTRY)(unsafe.Pointer(ret))
}

func RtlUnwindEx(TargetFrame unsafe.Pointer, TargetIp unsafe.Pointer, ExceptionRecord *EXCEPTION_RECORD, ReturnValue unsafe.Pointer, ContextRecord *CONTEXT, HistoryTable *UNWIND_HISTORY_TABLE) {
	addr := LazyAddr(&pRtlUnwindEx, libKernel32, "RtlUnwindEx")
	syscall.SyscallN(addr, uintptr(TargetFrame), uintptr(TargetIp), uintptr(unsafe.Pointer(ExceptionRecord)), uintptr(ReturnValue), uintptr(unsafe.Pointer(ContextRecord)), uintptr(unsafe.Pointer(HistoryTable)))
}

func RtlVirtualUnwind(HandlerType RTL_VIRTUAL_UNWIND_HANDLER_TYPE, ImageBase uint64, ControlPc uint64, FunctionEntry *IMAGE_RUNTIME_FUNCTION_ENTRY, ContextRecord *CONTEXT, HandlerData unsafe.Pointer, EstablisherFrame *uint64, ContextPointers *KNONVOLATILE_CONTEXT_POINTERS) EXCEPTION_ROUTINE {
	addr := LazyAddr(&pRtlVirtualUnwind, libKernel32, "RtlVirtualUnwind")
	ret, _, _ := syscall.SyscallN(addr, uintptr(HandlerType), uintptr(ImageBase), uintptr(ControlPc), uintptr(unsafe.Pointer(FunctionEntry)), uintptr(unsafe.Pointer(ContextRecord)), uintptr(HandlerData), uintptr(unsafe.Pointer(EstablisherFrame)), uintptr(unsafe.Pointer(ContextPointers)))
	return EXCEPTION_ROUTINE(ret)
}

func RtlCaptureStackBackTrace(FramesToSkip uint32, FramesToCapture uint32, BackTrace unsafe.Pointer, BackTraceHash *uint32) uint16 {
	addr := LazyAddr(&pRtlCaptureStackBackTrace, libKernel32, "RtlCaptureStackBackTrace")
	ret, _, _ := syscall.SyscallN(addr, uintptr(FramesToSkip), uintptr(FramesToCapture), uintptr(BackTrace), uintptr(unsafe.Pointer(BackTraceHash)))
	return uint16(ret)
}

func RtlCaptureContext(ContextRecord *CONTEXT) {
	addr := LazyAddr(&pRtlCaptureContext, libKernel32, "RtlCaptureContext")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(ContextRecord)))
}

func RtlUnwind(TargetFrame unsafe.Pointer, TargetIp unsafe.Pointer, ExceptionRecord *EXCEPTION_RECORD, ReturnValue unsafe.Pointer) {
	addr := LazyAddr(&pRtlUnwind, libKernel32, "RtlUnwind")
	syscall.SyscallN(addr, uintptr(TargetFrame), uintptr(TargetIp), uintptr(unsafe.Pointer(ExceptionRecord)), uintptr(ReturnValue))
}

func RtlRestoreContext(ContextRecord *CONTEXT, ExceptionRecord *EXCEPTION_RECORD) {
	addr := LazyAddr(&pRtlRestoreContext, libKernel32, "RtlRestoreContext")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(ContextRecord)), uintptr(unsafe.Pointer(ExceptionRecord)))
}

func RtlRaiseException(ExceptionRecord *EXCEPTION_RECORD) {
	addr := LazyAddr(&pRtlRaiseException, libKernel32, "RtlRaiseException")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExceptionRecord)))
}

func RtlPcToFileHeader(PcValue unsafe.Pointer, BaseOfImage unsafe.Pointer) unsafe.Pointer {
	addr := LazyAddr(&pRtlPcToFileHeader, libKernel32, "RtlPcToFileHeader")
	ret, _, _ := syscall.SyscallN(addr, uintptr(PcValue), uintptr(BaseOfImage))
	return (unsafe.Pointer)(ret)
}

func IsDebuggerPresent() BOOL {
	addr := LazyAddr(&pIsDebuggerPresent, libKernel32, "IsDebuggerPresent")
	ret, _, _ := syscall.SyscallN(addr)
	return BOOL(ret)
}

func DebugBreak() {
	addr := LazyAddr(&pDebugBreak, libKernel32, "DebugBreak")
	syscall.SyscallN(addr)
}

func OutputDebugStringA(lpOutputString PSTR) {
	addr := LazyAddr(&pOutputDebugStringA, libKernel32, "OutputDebugStringA")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpOutputString)))
}

var OutputDebugString = OutputDebugStringW

func OutputDebugStringW(lpOutputString PWSTR) {
	addr := LazyAddr(&pOutputDebugStringW, libKernel32, "OutputDebugStringW")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpOutputString)))
}

func ContinueDebugEvent(dwProcessId uint32, dwThreadId uint32, dwContinueStatus NTSTATUS) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pContinueDebugEvent, libKernel32, "ContinueDebugEvent")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwProcessId), uintptr(dwThreadId), uintptr(dwContinueStatus))
	return BOOL(ret), WIN32_ERROR(err)
}

func WaitForDebugEvent(lpDebugEvent *DEBUG_EVENT, dwMilliseconds uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pWaitForDebugEvent, libKernel32, "WaitForDebugEvent")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpDebugEvent)), uintptr(dwMilliseconds))
	return BOOL(ret), WIN32_ERROR(err)
}

func DebugActiveProcess(dwProcessId uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pDebugActiveProcess, libKernel32, "DebugActiveProcess")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwProcessId))
	return BOOL(ret), WIN32_ERROR(err)
}

func DebugActiveProcessStop(dwProcessId uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pDebugActiveProcessStop, libKernel32, "DebugActiveProcessStop")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwProcessId))
	return BOOL(ret), WIN32_ERROR(err)
}

func CheckRemoteDebuggerPresent(hProcess HANDLE, pbDebuggerPresent *BOOL) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pCheckRemoteDebuggerPresent, libKernel32, "CheckRemoteDebuggerPresent")
	ret, _, err := syscall.SyscallN(addr, hProcess, uintptr(unsafe.Pointer(pbDebuggerPresent)))
	return BOOL(ret), WIN32_ERROR(err)
}

func WaitForDebugEventEx(lpDebugEvent *DEBUG_EVENT, dwMilliseconds uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pWaitForDebugEventEx, libKernel32, "WaitForDebugEventEx")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpDebugEvent)), uintptr(dwMilliseconds))
	return BOOL(ret), WIN32_ERROR(err)
}

func EncodePointer(Ptr unsafe.Pointer) unsafe.Pointer {
	addr := LazyAddr(&pEncodePointer, libKernel32, "EncodePointer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(Ptr))
	return (unsafe.Pointer)(ret)
}

func DecodePointer(Ptr unsafe.Pointer) unsafe.Pointer {
	addr := LazyAddr(&pDecodePointer, libKernel32, "DecodePointer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(Ptr))
	return (unsafe.Pointer)(ret)
}

func EncodeSystemPointer(Ptr unsafe.Pointer) unsafe.Pointer {
	addr := LazyAddr(&pEncodeSystemPointer, libKernel32, "EncodeSystemPointer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(Ptr))
	return (unsafe.Pointer)(ret)
}

func DecodeSystemPointer(Ptr unsafe.Pointer) unsafe.Pointer {
	addr := LazyAddr(&pDecodeSystemPointer, libKernel32, "DecodeSystemPointer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(Ptr))
	return (unsafe.Pointer)(ret)
}

func Beep(dwFreq uint32, dwDuration uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pBeep, libKernel32, "Beep")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwFreq), uintptr(dwDuration))
	return BOOL(ret), WIN32_ERROR(err)
}

func RaiseException(dwExceptionCode uint32, dwExceptionFlags uint32, nNumberOfArguments uint32, lpArguments *uintptr) {
	addr := LazyAddr(&pRaiseException, libKernel32, "RaiseException")
	syscall.SyscallN(addr, uintptr(dwExceptionCode), uintptr(dwExceptionFlags), uintptr(nNumberOfArguments), uintptr(unsafe.Pointer(lpArguments)))
}

func UnhandledExceptionFilter(ExceptionInfo *EXCEPTION_POINTERS) int32 {
	addr := LazyAddr(&pUnhandledExceptionFilter, libKernel32, "UnhandledExceptionFilter")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExceptionInfo)))
	return int32(ret)
}

func SetUnhandledExceptionFilter(lpTopLevelExceptionFilter LPTOP_LEVEL_EXCEPTION_FILTER) LPTOP_LEVEL_EXCEPTION_FILTER {
	addr := LazyAddr(&pSetUnhandledExceptionFilter, libKernel32, "SetUnhandledExceptionFilter")
	ret, _, _ := syscall.SyscallN(addr, lpTopLevelExceptionFilter)
	return LPTOP_LEVEL_EXCEPTION_FILTER(ret)
}

func GetErrorMode() uint32 {
	addr := LazyAddr(&pGetErrorMode, libKernel32, "GetErrorMode")
	ret, _, _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func SetErrorMode(uMode THREAD_ERROR_MODE) uint32 {
	addr := LazyAddr(&pSetErrorMode, libKernel32, "SetErrorMode")
	ret, _, _ := syscall.SyscallN(addr, uintptr(uMode))
	return uint32(ret)
}

func AddVectoredExceptionHandler(First uint32, Handler PVECTORED_EXCEPTION_HANDLER) unsafe.Pointer {
	addr := LazyAddr(&pAddVectoredExceptionHandler, libKernel32, "AddVectoredExceptionHandler")
	ret, _, _ := syscall.SyscallN(addr, uintptr(First), Handler)
	return (unsafe.Pointer)(ret)
}

func RemoveVectoredExceptionHandler(Handle unsafe.Pointer) uint32 {
	addr := LazyAddr(&pRemoveVectoredExceptionHandler, libKernel32, "RemoveVectoredExceptionHandler")
	ret, _, _ := syscall.SyscallN(addr, uintptr(Handle))
	return uint32(ret)
}

func AddVectoredContinueHandler(First uint32, Handler PVECTORED_EXCEPTION_HANDLER) unsafe.Pointer {
	addr := LazyAddr(&pAddVectoredContinueHandler, libKernel32, "AddVectoredContinueHandler")
	ret, _, _ := syscall.SyscallN(addr, uintptr(First), Handler)
	return (unsafe.Pointer)(ret)
}

func RemoveVectoredContinueHandler(Handle unsafe.Pointer) uint32 {
	addr := LazyAddr(&pRemoveVectoredContinueHandler, libKernel32, "RemoveVectoredContinueHandler")
	ret, _, _ := syscall.SyscallN(addr, uintptr(Handle))
	return uint32(ret)
}

func RaiseFailFastException(pExceptionRecord *EXCEPTION_RECORD, pContextRecord *CONTEXT, dwFlags uint32) {
	addr := LazyAddr(&pRaiseFailFastException, libKernel32, "RaiseFailFastException")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(pExceptionRecord)), uintptr(unsafe.Pointer(pContextRecord)), uintptr(dwFlags))
}

func FatalAppExitA(uAction uint32, lpMessageText PSTR) {
	addr := LazyAddr(&pFatalAppExitA, libKernel32, "FatalAppExitA")
	syscall.SyscallN(addr, uintptr(uAction), uintptr(unsafe.Pointer(lpMessageText)))
}

var FatalAppExit = FatalAppExitW

func FatalAppExitW(uAction uint32, lpMessageText PWSTR) {
	addr := LazyAddr(&pFatalAppExitW, libKernel32, "FatalAppExitW")
	syscall.SyscallN(addr, uintptr(uAction), uintptr(unsafe.Pointer(lpMessageText)))
}

func GetThreadErrorMode() uint32 {
	addr := LazyAddr(&pGetThreadErrorMode, libKernel32, "GetThreadErrorMode")
	ret, _, _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func SetThreadErrorMode(dwNewMode THREAD_ERROR_MODE, lpOldMode *THREAD_ERROR_MODE) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pSetThreadErrorMode, libKernel32, "SetThreadErrorMode")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwNewMode), uintptr(unsafe.Pointer(lpOldMode)))
	return BOOL(ret), WIN32_ERROR(err)
}

func OpenThreadWaitChainSession(Flags OPEN_THREAD_WAIT_CHAIN_SESSION_FLAGS, callback PWAITCHAINCALLBACK) (unsafe.Pointer, WIN32_ERROR) {
	addr := LazyAddr(&pOpenThreadWaitChainSession, libAdvapi32, "OpenThreadWaitChainSession")
	ret, _, err := syscall.SyscallN(addr, uintptr(Flags), callback)
	return (unsafe.Pointer)(ret), WIN32_ERROR(err)
}

func CloseThreadWaitChainSession(WctHandle unsafe.Pointer) {
	addr := LazyAddr(&pCloseThreadWaitChainSession, libAdvapi32, "CloseThreadWaitChainSession")
	syscall.SyscallN(addr, uintptr(WctHandle))
}

func GetThreadWaitChain(WctHandle unsafe.Pointer, Context uintptr, Flags WAIT_CHAIN_THREAD_OPTIONS, ThreadId uint32, NodeCount *uint32, NodeInfoArray *WAITCHAIN_NODE_INFO, IsCycle *int32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetThreadWaitChain, libAdvapi32, "GetThreadWaitChain")
	ret, _, err := syscall.SyscallN(addr, uintptr(WctHandle), Context, uintptr(Flags), uintptr(ThreadId), uintptr(unsafe.Pointer(NodeCount)), uintptr(unsafe.Pointer(NodeInfoArray)), uintptr(unsafe.Pointer(IsCycle)))
	return BOOL(ret), WIN32_ERROR(err)
}

func RegisterWaitChainCOMCallback(CallStateCallback PCOGETCALLSTATE, ActivationStateCallback PCOGETACTIVATIONSTATE) {
	addr := LazyAddr(&pRegisterWaitChainCOMCallback, libAdvapi32, "RegisterWaitChainCOMCallback")
	syscall.SyscallN(addr, CallStateCallback, ActivationStateCallback)
}

func MessageBeep(uType MESSAGEBOX_STYLE) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pMessageBeep, libUser32, "MessageBeep")
	ret, _, err := syscall.SyscallN(addr, uintptr(uType))
	return BOOL(ret), WIN32_ERROR(err)
}

func FatalExit(ExitCode int32) {
	addr := LazyAddr(&pFatalExit, libKernel32, "FatalExit")
	syscall.SyscallN(addr, uintptr(ExitCode))
}

func GetThreadSelectorEntry(hThread HANDLE, dwSelector uint32, lpSelectorEntry *LDT_ENTRY) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pGetThreadSelectorEntry, libKernel32, "GetThreadSelectorEntry")
	ret, _, err := syscall.SyscallN(addr, hThread, uintptr(dwSelector), uintptr(unsafe.Pointer(lpSelectorEntry)))
	return BOOL(ret), WIN32_ERROR(err)
}

func Wow64GetThreadSelectorEntry(hThread HANDLE, dwSelector uint32, lpSelectorEntry *WOW64_LDT_ENTRY) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pWow64GetThreadSelectorEntry, libKernel32, "Wow64GetThreadSelectorEntry")
	ret, _, err := syscall.SyscallN(addr, hThread, uintptr(dwSelector), uintptr(unsafe.Pointer(lpSelectorEntry)))
	return BOOL(ret), WIN32_ERROR(err)
}

func DebugSetProcessKillOnExit(KillOnExit BOOL) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pDebugSetProcessKillOnExit, libKernel32, "DebugSetProcessKillOnExit")
	ret, _, err := syscall.SyscallN(addr, uintptr(KillOnExit))
	return BOOL(ret), WIN32_ERROR(err)
}

func DebugBreakProcess(Process HANDLE) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pDebugBreakProcess, libKernel32, "DebugBreakProcess")
	ret, _, err := syscall.SyscallN(addr, Process)
	return BOOL(ret), WIN32_ERROR(err)
}

func FormatMessageA(dwFlags FORMAT_MESSAGE_OPTIONS, lpSource unsafe.Pointer, dwMessageId uint32, dwLanguageId uint32, lpBuffer PSTR, nSize uint32, Arguments **int8) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pFormatMessageA, libKernel32, "FormatMessageA")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwFlags), uintptr(lpSource), uintptr(dwMessageId), uintptr(dwLanguageId), uintptr(unsafe.Pointer(lpBuffer)), uintptr(nSize), uintptr(unsafe.Pointer(Arguments)))
	return uint32(ret), WIN32_ERROR(err)
}

var FormatMessage = FormatMessageW

func FormatMessageW(dwFlags FORMAT_MESSAGE_OPTIONS, lpSource unsafe.Pointer, dwMessageId uint32, dwLanguageId uint32, lpBuffer PWSTR, nSize uint32, Arguments **int8) (uint32, WIN32_ERROR) {
	addr := LazyAddr(&pFormatMessageW, libKernel32, "FormatMessageW")
	ret, _, err := syscall.SyscallN(addr, uintptr(dwFlags), uintptr(lpSource), uintptr(dwMessageId), uintptr(dwLanguageId), uintptr(unsafe.Pointer(lpBuffer)), uintptr(nSize), uintptr(unsafe.Pointer(Arguments)))
	return uint32(ret), WIN32_ERROR(err)
}

func CopyContext(Destination *CONTEXT, ContextFlags CONTEXT_FLAGS, Source *CONTEXT) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pCopyContext, libKernel32, "CopyContext")
	ret, _, err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Destination)), uintptr(ContextFlags), uintptr(unsafe.Pointer(Source)))
	return BOOL(ret), WIN32_ERROR(err)
}

func InitializeContext(Buffer unsafe.Pointer, ContextFlags CONTEXT_FLAGS, Context **CONTEXT, ContextLength *uint32) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pInitializeContext, libKernel32, "InitializeContext")
	ret, _, err := syscall.SyscallN(addr, uintptr(Buffer), uintptr(ContextFlags), uintptr(unsafe.Pointer(Context)), uintptr(unsafe.Pointer(ContextLength)))
	return BOOL(ret), WIN32_ERROR(err)
}

func InitializeContext2(Buffer unsafe.Pointer, ContextFlags CONTEXT_FLAGS, Context **CONTEXT, ContextLength *uint32, XStateCompactionMask uint64) (BOOL, WIN32_ERROR) {
	addr := LazyAddr(&pInitializeContext2, libKernel32, "InitializeContext2")
	ret, _, err := syscall.SyscallN(addr, uintptr(Buffer), uintptr(ContextFlags), uintptr(unsafe.Pointer(Context)), uintptr(unsafe.Pointer(ContextLength)), uintptr(XStateCompactionMask))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetEnabledXStateFeatures() uint64 {
	addr := LazyAddr(&pGetEnabledXStateFeatures, libKernel32, "GetEnabledXStateFeatures")
	ret, _, _ := syscall.SyscallN(addr)
	return uint64(ret)
}

func GetXStateFeaturesMask(Context *CONTEXT, FeatureMask *uint64) BOOL {
	addr := LazyAddr(&pGetXStateFeaturesMask, libKernel32, "GetXStateFeaturesMask")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Context)), uintptr(unsafe.Pointer(FeatureMask)))
	return BOOL(ret)
}

func LocateXStateFeature(Context *CONTEXT, FeatureId uint32, Length *uint32) unsafe.Pointer {
	addr := LazyAddr(&pLocateXStateFeature, libKernel32, "LocateXStateFeature")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Context)), uintptr(FeatureId), uintptr(unsafe.Pointer(Length)))
	return (unsafe.Pointer)(ret)
}

func SetXStateFeaturesMask(Context *CONTEXT, FeatureMask uint64) BOOL {
	addr := LazyAddr(&pSetXStateFeaturesMask, libKernel32, "SetXStateFeaturesMask")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Context)), uintptr(FeatureMask))
	return BOOL(ret)
}
