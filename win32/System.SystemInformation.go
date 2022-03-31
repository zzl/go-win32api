package win32

import "unsafe"
import "syscall"

type FIRMWARE_TABLE_ID = uint32

const (
	NTDDI_WIN2K uint32 = 83886080
	NTDDI_WINXP uint32 = 83951616
	NTDDI_WINXPSP2 uint32 = 83952128
	NTDDI_WS03SP1 uint32 = 84017408
	NTDDI_VISTA uint32 = 100663296
	NTDDI_VISTASP1 uint32 = 100663552
	NTDDI_WIN7 uint32 = 100728832
	NTDDI_WIN8 uint32 = 100794368
	NTDDI_WINBLUE uint32 = 100859904
	NTDDI_WINTHRESHOLD uint32 = 167772160
	SYSTEM_CPU_SET_INFORMATION_PARKED uint32 = 1
	SYSTEM_CPU_SET_INFORMATION_ALLOCATED uint32 = 2
	SYSTEM_CPU_SET_INFORMATION_ALLOCATED_TO_TARGET_PROCESS uint32 = 4
	SYSTEM_CPU_SET_INFORMATION_REALTIME uint32 = 8
	WIN32_WINNT_NT4_ uint32 = 1024
	WIN32_WINNT_WIN2K_ uint32 = 1280
	WIN32_WINNT_WINXP_ uint32 = 1281
	WIN32_WINNT_WS03_ uint32 = 1282
	WIN32_WINNT_WIN6_ uint32 = 1536
	WIN32_WINNT_VISTA_ uint32 = 1536
	WIN32_WINNT_WS08_ uint32 = 1536
	WIN32_WINNT_LONGHORN_ uint32 = 1536
	WIN32_WINNT_WIN7_ uint32 = 1537
	WIN32_WINNT_WIN8_ uint32 = 1538
	WIN32_WINNT_WINBLUE_ uint32 = 1539
	WIN32_WINNT_WINTHRESHOLD_ uint32 = 2560
	WIN32_WINNT_WIN10_ uint32 = 2560
	WIN32_IE_IE20_ uint32 = 512
	WIN32_IE_IE30_ uint32 = 768
	WIN32_IE_IE302_ uint32 = 770
	WIN32_IE_IE40_ uint32 = 1024
	WIN32_IE_IE401_ uint32 = 1025
	WIN32_IE_IE50_ uint32 = 1280
	WIN32_IE_IE501_ uint32 = 1281
	WIN32_IE_IE55_ uint32 = 1360
	WIN32_IE_IE60_ uint32 = 1536
	WIN32_IE_IE60SP1_ uint32 = 1537
	WIN32_IE_IE60SP2_ uint32 = 1539
	WIN32_IE_IE70_ uint32 = 1792
	WIN32_IE_IE80_ uint32 = 2048
	WIN32_IE_IE90_ uint32 = 2304
	WIN32_IE_IE100_ uint32 = 2560
	WIN32_IE_IE110_ uint32 = 2560
	WIN32_IE_NT4_ uint32 = 512
	WIN32_IE_NT4SP1_ uint32 = 512
	WIN32_IE_NT4SP2_ uint32 = 512
	WIN32_IE_NT4SP3_ uint32 = 770
	WIN32_IE_NT4SP4_ uint32 = 1025
	WIN32_IE_NT4SP5_ uint32 = 1025
	WIN32_IE_NT4SP6_ uint32 = 1280
	WIN32_IE_WIN98_ uint32 = 1025
	WIN32_IE_WIN98SE_ uint32 = 1280
	WIN32_IE_WINME_ uint32 = 1360
	WIN32_IE_WIN2K_ uint32 = 1281
	WIN32_IE_WIN2KSP1_ uint32 = 1281
	WIN32_IE_WIN2KSP2_ uint32 = 1281
	WIN32_IE_WIN2KSP3_ uint32 = 1281
	WIN32_IE_WIN2KSP4_ uint32 = 1281
	WIN32_IE_XP_ uint32 = 1536
	WIN32_IE_XPSP1_ uint32 = 1537
	WIN32_IE_XPSP2_ uint32 = 1539
	WIN32_IE_WS03_ uint32 = 1538
	WIN32_IE_WS03SP1_ uint32 = 1539
	WIN32_IE_WIN6_ uint32 = 1792
	WIN32_IE_LONGHORN_ uint32 = 1792
	WIN32_IE_WIN7_ uint32 = 2048
	WIN32_IE_WIN8_ uint32 = 2560
	WIN32_IE_WINBLUE_ uint32 = 2560
	WIN32_IE_WINTHRESHOLD_ uint32 = 2560
	WIN32_IE_WIN10_ uint32 = 2560
	NTDDI_WIN4 uint32 = 67108864
	NTDDI_WIN2KSP1 uint32 = 83886336
	NTDDI_WIN2KSP2 uint32 = 83886592
	NTDDI_WIN2KSP3 uint32 = 83886848
	NTDDI_WIN2KSP4 uint32 = 83887104
	NTDDI_WINXPSP1 uint32 = 83951872
	NTDDI_WINXPSP3 uint32 = 83952384
	NTDDI_WINXPSP4 uint32 = 83952640
	NTDDI_WS03 uint32 = 84017152
	NTDDI_WS03SP2 uint32 = 84017664
	NTDDI_WS03SP3 uint32 = 84017920
	NTDDI_WS03SP4 uint32 = 84018176
	NTDDI_WIN6 uint32 = 100663296
	NTDDI_WIN6SP1 uint32 = 100663552
	NTDDI_WIN6SP2 uint32 = 100663808
	NTDDI_WIN6SP3 uint32 = 100664064
	NTDDI_WIN6SP4 uint32 = 100664320
	NTDDI_VISTASP2 uint32 = 100663808
	NTDDI_VISTASP3 uint32 = 100664064
	NTDDI_VISTASP4 uint32 = 100664320
	NTDDI_LONGHORN uint32 = 100663296
	NTDDI_WS08 uint32 = 100663552
	NTDDI_WS08SP2 uint32 = 100663808
	NTDDI_WS08SP3 uint32 = 100664064
	NTDDI_WS08SP4 uint32 = 100664320
	NTDDI_WIN10 uint32 = 167772160
	NTDDI_WIN10_TH2 uint32 = 167772161
	NTDDI_WIN10_RS1 uint32 = 167772162
	NTDDI_WIN10_RS2 uint32 = 167772163
	NTDDI_WIN10_RS3 uint32 = 167772164
	NTDDI_WIN10_RS4 uint32 = 167772165
	NTDDI_WIN10_RS5 uint32 = 167772166
	NTDDI_WIN10_19H1 uint32 = 167772167
	NTDDI_WIN10_VB uint32 = 167772168
	NTDDI_WIN10_MN uint32 = 167772169
	NTDDI_WIN10_FE uint32 = 167772170
	NTDDI_WIN10_CO uint32 = 167772171
	WDK_NTDDI_VERSION uint32 = 167772171
	OSVERSION_MASK uint32 = 4294901760
	SPVERSION_MASK uint32 = 65280
	SUBVERSION_MASK uint32 = 255
	NTDDI_VERSION uint32 = 167772171
	SCEX2_ALT_NETBIOS_NAME uint32 = 1
)

// enums

// enum VER_FLAGS
// flags
type VER_FLAGS uint32
const (
	VER_MINORVERSION VER_FLAGS = 1
	VER_MAJORVERSION VER_FLAGS = 2
	VER_BUILDNUMBER VER_FLAGS = 4
	VER_PLATFORMID VER_FLAGS = 8
	VER_SERVICEPACKMINOR VER_FLAGS = 16
	VER_SERVICEPACKMAJOR VER_FLAGS = 32
	VER_SUITENAME VER_FLAGS = 64
	VER_PRODUCT_TYPE VER_FLAGS = 128
)

// enum FIRMWARE_TABLE_PROVIDER
type FIRMWARE_TABLE_PROVIDER uint32
const (
	ACPI FIRMWARE_TABLE_PROVIDER = 1094930505
	FIRM FIRMWARE_TABLE_PROVIDER = 1179210317
	RSMB FIRMWARE_TABLE_PROVIDER = 1381190978
)

// enum USER_CET_ENVIRONMENT
type USER_CET_ENVIRONMENT uint32
const (
	USER_CET_ENVIRONMENT_WIN32_PROCESS USER_CET_ENVIRONMENT = 0
	USER_CET_ENVIRONMENT_SGX2_ENCLAVE USER_CET_ENVIRONMENT = 2
	USER_CET_ENVIRONMENT_VBS_ENCLAVE USER_CET_ENVIRONMENT = 16
	USER_CET_ENVIRONMENT_VBS_BASIC_ENCLAVE USER_CET_ENVIRONMENT = 17
)

// enum OS_PRODUCT_TYPE
type OS_PRODUCT_TYPE uint32
const (
	PRODUCT_BUSINESS OS_PRODUCT_TYPE = 6
	PRODUCT_BUSINESS_N OS_PRODUCT_TYPE = 16
	PRODUCT_CLUSTER_SERVER OS_PRODUCT_TYPE = 18
	PRODUCT_CLUSTER_SERVER_V OS_PRODUCT_TYPE = 64
	PRODUCT_CORE OS_PRODUCT_TYPE = 101
	PRODUCT_CORE_COUNTRYSPECIFIC OS_PRODUCT_TYPE = 99
	PRODUCT_CORE_N OS_PRODUCT_TYPE = 98
	PRODUCT_CORE_SINGLELANGUAGE OS_PRODUCT_TYPE = 100
	PRODUCT_DATACENTER_EVALUATION_SERVER OS_PRODUCT_TYPE = 80
	PRODUCT_DATACENTER_A_SERVER_CORE OS_PRODUCT_TYPE = 145
	PRODUCT_STANDARD_A_SERVER_CORE OS_PRODUCT_TYPE = 146
	PRODUCT_DATACENTER_SERVER OS_PRODUCT_TYPE = 8
	PRODUCT_DATACENTER_SERVER_CORE OS_PRODUCT_TYPE = 12
	PRODUCT_DATACENTER_SERVER_CORE_V OS_PRODUCT_TYPE = 39
	PRODUCT_DATACENTER_SERVER_V OS_PRODUCT_TYPE = 37
	PRODUCT_EDUCATION OS_PRODUCT_TYPE = 121
	PRODUCT_EDUCATION_N OS_PRODUCT_TYPE = 122
	PRODUCT_ENTERPRISE OS_PRODUCT_TYPE = 4
	PRODUCT_ENTERPRISE_E OS_PRODUCT_TYPE = 70
	PRODUCT_ENTERPRISE_EVALUATION OS_PRODUCT_TYPE = 72
	PRODUCT_ENTERPRISE_N OS_PRODUCT_TYPE = 27
	PRODUCT_ENTERPRISE_N_EVALUATION OS_PRODUCT_TYPE = 84
	PRODUCT_ENTERPRISE_S OS_PRODUCT_TYPE = 125
	PRODUCT_ENTERPRISE_S_EVALUATION OS_PRODUCT_TYPE = 129
	PRODUCT_ENTERPRISE_S_N OS_PRODUCT_TYPE = 126
	PRODUCT_ENTERPRISE_S_N_EVALUATION OS_PRODUCT_TYPE = 130
	PRODUCT_ENTERPRISE_SERVER OS_PRODUCT_TYPE = 10
	PRODUCT_ENTERPRISE_SERVER_CORE OS_PRODUCT_TYPE = 14
	PRODUCT_ENTERPRISE_SERVER_CORE_V OS_PRODUCT_TYPE = 41
	PRODUCT_ENTERPRISE_SERVER_IA64 OS_PRODUCT_TYPE = 15
	PRODUCT_ENTERPRISE_SERVER_V OS_PRODUCT_TYPE = 38
	PRODUCT_ESSENTIALBUSINESS_SERVER_ADDL OS_PRODUCT_TYPE = 60
	PRODUCT_ESSENTIALBUSINESS_SERVER_ADDLSVC OS_PRODUCT_TYPE = 62
	PRODUCT_ESSENTIALBUSINESS_SERVER_MGMT OS_PRODUCT_TYPE = 59
	PRODUCT_ESSENTIALBUSINESS_SERVER_MGMTSVC OS_PRODUCT_TYPE = 61
	PRODUCT_HOME_BASIC OS_PRODUCT_TYPE = 2
	PRODUCT_HOME_BASIC_E OS_PRODUCT_TYPE = 67
	PRODUCT_HOME_BASIC_N OS_PRODUCT_TYPE = 5
	PRODUCT_HOME_PREMIUM OS_PRODUCT_TYPE = 3
	PRODUCT_HOME_PREMIUM_E OS_PRODUCT_TYPE = 68
	PRODUCT_HOME_PREMIUM_N OS_PRODUCT_TYPE = 26
	PRODUCT_HOME_PREMIUM_SERVER OS_PRODUCT_TYPE = 34
	PRODUCT_HOME_SERVER OS_PRODUCT_TYPE = 19
	PRODUCT_HYPERV OS_PRODUCT_TYPE = 42
	PRODUCT_IOTUAP OS_PRODUCT_TYPE = 123
	PRODUCT_IOTUAPCOMMERCIAL OS_PRODUCT_TYPE = 131
	PRODUCT_MEDIUMBUSINESS_SERVER_MANAGEMENT OS_PRODUCT_TYPE = 30
	PRODUCT_MEDIUMBUSINESS_SERVER_MESSAGING OS_PRODUCT_TYPE = 32
	PRODUCT_MEDIUMBUSINESS_SERVER_SECURITY OS_PRODUCT_TYPE = 31
	PRODUCT_MOBILE_CORE OS_PRODUCT_TYPE = 104
	PRODUCT_MOBILE_ENTERPRISE OS_PRODUCT_TYPE = 133
	PRODUCT_MULTIPOINT_PREMIUM_SERVER OS_PRODUCT_TYPE = 77
	PRODUCT_MULTIPOINT_STANDARD_SERVER OS_PRODUCT_TYPE = 76
	PRODUCT_PRO_WORKSTATION OS_PRODUCT_TYPE = 161
	PRODUCT_PRO_WORKSTATION_N OS_PRODUCT_TYPE = 162
	PRODUCT_PROFESSIONAL OS_PRODUCT_TYPE = 48
	PRODUCT_PROFESSIONAL_E OS_PRODUCT_TYPE = 69
	PRODUCT_PROFESSIONAL_N OS_PRODUCT_TYPE = 49
	PRODUCT_PROFESSIONAL_WMC OS_PRODUCT_TYPE = 103
	PRODUCT_SB_SOLUTION_SERVER OS_PRODUCT_TYPE = 50
	PRODUCT_SB_SOLUTION_SERVER_EM OS_PRODUCT_TYPE = 54
	PRODUCT_SERVER_FOR_SB_SOLUTIONS OS_PRODUCT_TYPE = 51
	PRODUCT_SERVER_FOR_SB_SOLUTIONS_EM OS_PRODUCT_TYPE = 55
	PRODUCT_SERVER_FOR_SMALLBUSINESS OS_PRODUCT_TYPE = 24
	PRODUCT_SERVER_FOR_SMALLBUSINESS_V OS_PRODUCT_TYPE = 35
	PRODUCT_SERVER_FOUNDATION OS_PRODUCT_TYPE = 33
	PRODUCT_SMALLBUSINESS_SERVER OS_PRODUCT_TYPE = 9
	PRODUCT_SMALLBUSINESS_SERVER_PREMIUM OS_PRODUCT_TYPE = 25
	PRODUCT_SMALLBUSINESS_SERVER_PREMIUM_CORE OS_PRODUCT_TYPE = 63
	PRODUCT_SOLUTION_EMBEDDEDSERVER OS_PRODUCT_TYPE = 56
	PRODUCT_STANDARD_EVALUATION_SERVER OS_PRODUCT_TYPE = 79
	PRODUCT_STANDARD_SERVER OS_PRODUCT_TYPE = 7
	PRODUCT_STANDARD_SERVER_CORE_ OS_PRODUCT_TYPE = 13
	PRODUCT_STANDARD_SERVER_CORE_V OS_PRODUCT_TYPE = 40
	PRODUCT_STANDARD_SERVER_V OS_PRODUCT_TYPE = 36
	PRODUCT_STANDARD_SERVER_SOLUTIONS OS_PRODUCT_TYPE = 52
	PRODUCT_STANDARD_SERVER_SOLUTIONS_CORE OS_PRODUCT_TYPE = 53
	PRODUCT_STARTER OS_PRODUCT_TYPE = 11
	PRODUCT_STARTER_E OS_PRODUCT_TYPE = 66
	PRODUCT_STARTER_N OS_PRODUCT_TYPE = 47
	PRODUCT_STORAGE_ENTERPRISE_SERVER OS_PRODUCT_TYPE = 23
	PRODUCT_STORAGE_ENTERPRISE_SERVER_CORE OS_PRODUCT_TYPE = 46
	PRODUCT_STORAGE_EXPRESS_SERVER OS_PRODUCT_TYPE = 20
	PRODUCT_STORAGE_EXPRESS_SERVER_CORE OS_PRODUCT_TYPE = 43
	PRODUCT_STORAGE_STANDARD_EVALUATION_SERVER OS_PRODUCT_TYPE = 96
	PRODUCT_STORAGE_STANDARD_SERVER OS_PRODUCT_TYPE = 21
	PRODUCT_STORAGE_STANDARD_SERVER_CORE OS_PRODUCT_TYPE = 44
	PRODUCT_STORAGE_WORKGROUP_EVALUATION_SERVER OS_PRODUCT_TYPE = 95
	PRODUCT_STORAGE_WORKGROUP_SERVER OS_PRODUCT_TYPE = 22
	PRODUCT_STORAGE_WORKGROUP_SERVER_CORE OS_PRODUCT_TYPE = 45
	PRODUCT_ULTIMATE OS_PRODUCT_TYPE = 1
	PRODUCT_ULTIMATE_E OS_PRODUCT_TYPE = 71
	PRODUCT_ULTIMATE_N OS_PRODUCT_TYPE = 28
	PRODUCT_UNDEFINED OS_PRODUCT_TYPE = 0
	PRODUCT_WEB_SERVER OS_PRODUCT_TYPE = 17
	PRODUCT_WEB_SERVER_CORE OS_PRODUCT_TYPE = 29
)

// enum DEVICEFAMILYINFOENUM
type DEVICEFAMILYINFOENUM uint32
const (
	DEVICEFAMILYINFOENUM_UAP DEVICEFAMILYINFOENUM = 0
	DEVICEFAMILYINFOENUM_WINDOWS_8X DEVICEFAMILYINFOENUM = 1
	DEVICEFAMILYINFOENUM_WINDOWS_PHONE_8X DEVICEFAMILYINFOENUM = 2
	DEVICEFAMILYINFOENUM_DESKTOP DEVICEFAMILYINFOENUM = 3
	DEVICEFAMILYINFOENUM_MOBILE DEVICEFAMILYINFOENUM = 4
	DEVICEFAMILYINFOENUM_XBOX DEVICEFAMILYINFOENUM = 5
	DEVICEFAMILYINFOENUM_TEAM DEVICEFAMILYINFOENUM = 6
	DEVICEFAMILYINFOENUM_IOT DEVICEFAMILYINFOENUM = 7
	DEVICEFAMILYINFOENUM_IOT_HEADLESS DEVICEFAMILYINFOENUM = 8
	DEVICEFAMILYINFOENUM_SERVER DEVICEFAMILYINFOENUM = 9
	DEVICEFAMILYINFOENUM_HOLOGRAPHIC DEVICEFAMILYINFOENUM = 10
	DEVICEFAMILYINFOENUM_XBOXSRA DEVICEFAMILYINFOENUM = 11
	DEVICEFAMILYINFOENUM_XBOXERA DEVICEFAMILYINFOENUM = 12
	DEVICEFAMILYINFOENUM_SERVER_NANO DEVICEFAMILYINFOENUM = 13
	DEVICEFAMILYINFOENUM_8828080 DEVICEFAMILYINFOENUM = 14
	DEVICEFAMILYINFOENUM_7067329 DEVICEFAMILYINFOENUM = 15
	DEVICEFAMILYINFOENUM_WINDOWS_CORE DEVICEFAMILYINFOENUM = 16
	DEVICEFAMILYINFOENUM_WINDOWS_CORE_HEADLESS DEVICEFAMILYINFOENUM = 17
	DEVICEFAMILYINFOENUM_MAX DEVICEFAMILYINFOENUM = 17
)

// enum DEVICEFAMILYDEVICEFORM
type DEVICEFAMILYDEVICEFORM uint32
const (
	DEVICEFAMILYDEVICEFORM_UNKNOWN DEVICEFAMILYDEVICEFORM = 0
	DEVICEFAMILYDEVICEFORM_PHONE DEVICEFAMILYDEVICEFORM = 1
	DEVICEFAMILYDEVICEFORM_TABLET DEVICEFAMILYDEVICEFORM = 2
	DEVICEFAMILYDEVICEFORM_DESKTOP DEVICEFAMILYDEVICEFORM = 3
	DEVICEFAMILYDEVICEFORM_NOTEBOOK DEVICEFAMILYDEVICEFORM = 4
	DEVICEFAMILYDEVICEFORM_CONVERTIBLE DEVICEFAMILYDEVICEFORM = 5
	DEVICEFAMILYDEVICEFORM_DETACHABLE DEVICEFAMILYDEVICEFORM = 6
	DEVICEFAMILYDEVICEFORM_ALLINONE DEVICEFAMILYDEVICEFORM = 7
	DEVICEFAMILYDEVICEFORM_STICKPC DEVICEFAMILYDEVICEFORM = 8
	DEVICEFAMILYDEVICEFORM_PUCK DEVICEFAMILYDEVICEFORM = 9
	DEVICEFAMILYDEVICEFORM_LARGESCREEN DEVICEFAMILYDEVICEFORM = 10
	DEVICEFAMILYDEVICEFORM_HMD DEVICEFAMILYDEVICEFORM = 11
	DEVICEFAMILYDEVICEFORM_INDUSTRY_HANDHELD DEVICEFAMILYDEVICEFORM = 12
	DEVICEFAMILYDEVICEFORM_INDUSTRY_TABLET DEVICEFAMILYDEVICEFORM = 13
	DEVICEFAMILYDEVICEFORM_BANKING DEVICEFAMILYDEVICEFORM = 14
	DEVICEFAMILYDEVICEFORM_BUILDING_AUTOMATION DEVICEFAMILYDEVICEFORM = 15
	DEVICEFAMILYDEVICEFORM_DIGITAL_SIGNAGE DEVICEFAMILYDEVICEFORM = 16
	DEVICEFAMILYDEVICEFORM_GAMING DEVICEFAMILYDEVICEFORM = 17
	DEVICEFAMILYDEVICEFORM_HOME_AUTOMATION DEVICEFAMILYDEVICEFORM = 18
	DEVICEFAMILYDEVICEFORM_INDUSTRIAL_AUTOMATION DEVICEFAMILYDEVICEFORM = 19
	DEVICEFAMILYDEVICEFORM_KIOSK DEVICEFAMILYDEVICEFORM = 20
	DEVICEFAMILYDEVICEFORM_MAKER_BOARD DEVICEFAMILYDEVICEFORM = 21
	DEVICEFAMILYDEVICEFORM_MEDICAL DEVICEFAMILYDEVICEFORM = 22
	DEVICEFAMILYDEVICEFORM_NETWORKING DEVICEFAMILYDEVICEFORM = 23
	DEVICEFAMILYDEVICEFORM_POINT_OF_SERVICE DEVICEFAMILYDEVICEFORM = 24
	DEVICEFAMILYDEVICEFORM_PRINTING DEVICEFAMILYDEVICEFORM = 25
	DEVICEFAMILYDEVICEFORM_THIN_CLIENT DEVICEFAMILYDEVICEFORM = 26
	DEVICEFAMILYDEVICEFORM_TOY DEVICEFAMILYDEVICEFORM = 27
	DEVICEFAMILYDEVICEFORM_VENDING DEVICEFAMILYDEVICEFORM = 28
	DEVICEFAMILYDEVICEFORM_INDUSTRY_OTHER DEVICEFAMILYDEVICEFORM = 29
	DEVICEFAMILYDEVICEFORM_XBOX_ONE DEVICEFAMILYDEVICEFORM = 30
	DEVICEFAMILYDEVICEFORM_XBOX_ONE_S DEVICEFAMILYDEVICEFORM = 31
	DEVICEFAMILYDEVICEFORM_XBOX_ONE_X DEVICEFAMILYDEVICEFORM = 32
	DEVICEFAMILYDEVICEFORM_XBOX_ONE_X_DEVKIT DEVICEFAMILYDEVICEFORM = 33
	DEVICEFAMILYDEVICEFORM_XBOX_SERIES_X DEVICEFAMILYDEVICEFORM = 34
	DEVICEFAMILYDEVICEFORM_XBOX_SERIES_X_DEVKIT DEVICEFAMILYDEVICEFORM = 35
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_00 DEVICEFAMILYDEVICEFORM = 36
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_01 DEVICEFAMILYDEVICEFORM = 37
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_02 DEVICEFAMILYDEVICEFORM = 38
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_03 DEVICEFAMILYDEVICEFORM = 39
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_04 DEVICEFAMILYDEVICEFORM = 40
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_05 DEVICEFAMILYDEVICEFORM = 41
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_06 DEVICEFAMILYDEVICEFORM = 42
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_07 DEVICEFAMILYDEVICEFORM = 43
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_08 DEVICEFAMILYDEVICEFORM = 44
	DEVICEFAMILYDEVICEFORM_XBOX_RESERVED_09 DEVICEFAMILYDEVICEFORM = 45
	DEVICEFAMILYDEVICEFORM_MAX DEVICEFAMILYDEVICEFORM = 45
)

// enum COMPUTER_NAME_FORMAT
type COMPUTER_NAME_FORMAT int32
const (
	ComputerNameNetBIOS COMPUTER_NAME_FORMAT = 0
	ComputerNameDnsHostname COMPUTER_NAME_FORMAT = 1
	ComputerNameDnsDomain COMPUTER_NAME_FORMAT = 2
	ComputerNameDnsFullyQualified COMPUTER_NAME_FORMAT = 3
	ComputerNamePhysicalNetBIOS COMPUTER_NAME_FORMAT = 4
	ComputerNamePhysicalDnsHostname COMPUTER_NAME_FORMAT = 5
	ComputerNamePhysicalDnsDomain COMPUTER_NAME_FORMAT = 6
	ComputerNamePhysicalDnsFullyQualified COMPUTER_NAME_FORMAT = 7
	ComputerNameMax COMPUTER_NAME_FORMAT = 8
)

// enum FIRMWARE_TYPE
type FIRMWARE_TYPE int32
const (
	FirmwareTypeUnknown FIRMWARE_TYPE = 0
	FirmwareTypeBios FIRMWARE_TYPE = 1
	FirmwareTypeUefi FIRMWARE_TYPE = 2
	FirmwareTypeMax FIRMWARE_TYPE = 3
)

// enum LOGICAL_PROCESSOR_RELATIONSHIP
type LOGICAL_PROCESSOR_RELATIONSHIP int32
const (
	RelationProcessorCore LOGICAL_PROCESSOR_RELATIONSHIP = 0
	RelationNumaNode LOGICAL_PROCESSOR_RELATIONSHIP = 1
	RelationCache LOGICAL_PROCESSOR_RELATIONSHIP = 2
	RelationProcessorPackage LOGICAL_PROCESSOR_RELATIONSHIP = 3
	RelationGroup LOGICAL_PROCESSOR_RELATIONSHIP = 4
	RelationProcessorDie LOGICAL_PROCESSOR_RELATIONSHIP = 5
	RelationNumaNodeEx LOGICAL_PROCESSOR_RELATIONSHIP = 6
	RelationProcessorModule LOGICAL_PROCESSOR_RELATIONSHIP = 7
	RelationAll LOGICAL_PROCESSOR_RELATIONSHIP = 65535
)

// enum PROCESSOR_CACHE_TYPE
type PROCESSOR_CACHE_TYPE int32
const (
	CacheUnified PROCESSOR_CACHE_TYPE = 0
	CacheInstruction PROCESSOR_CACHE_TYPE = 1
	CacheData PROCESSOR_CACHE_TYPE = 2
	CacheTrace PROCESSOR_CACHE_TYPE = 3
)

// enum CPU_SET_INFORMATION_TYPE
type CPU_SET_INFORMATION_TYPE int32
const (
	CpuSetInformation CPU_SET_INFORMATION_TYPE = 0
)

// enum OS_DEPLOYEMENT_STATE_VALUES
type OS_DEPLOYEMENT_STATE_VALUES int32
const (
	OS_DEPLOYMENT_STANDARD OS_DEPLOYEMENT_STATE_VALUES = 1
	OS_DEPLOYMENT_COMPACT OS_DEPLOYEMENT_STATE_VALUES = 2
)

// enum RTL_SYSTEM_GLOBAL_DATA_ID
type RTL_SYSTEM_GLOBAL_DATA_ID int32
const (
	GlobalDataIdUnknown RTL_SYSTEM_GLOBAL_DATA_ID = 0
	GlobalDataIdRngSeedVersion RTL_SYSTEM_GLOBAL_DATA_ID = 1
	GlobalDataIdInterruptTime RTL_SYSTEM_GLOBAL_DATA_ID = 2
	GlobalDataIdTimeZoneBias RTL_SYSTEM_GLOBAL_DATA_ID = 3
	GlobalDataIdImageNumberLow RTL_SYSTEM_GLOBAL_DATA_ID = 4
	GlobalDataIdImageNumberHigh RTL_SYSTEM_GLOBAL_DATA_ID = 5
	GlobalDataIdTimeZoneId RTL_SYSTEM_GLOBAL_DATA_ID = 6
	GlobalDataIdNtMajorVersion RTL_SYSTEM_GLOBAL_DATA_ID = 7
	GlobalDataIdNtMinorVersion RTL_SYSTEM_GLOBAL_DATA_ID = 8
	GlobalDataIdSystemExpirationDate RTL_SYSTEM_GLOBAL_DATA_ID = 9
	GlobalDataIdKdDebuggerEnabled RTL_SYSTEM_GLOBAL_DATA_ID = 10
	GlobalDataIdCyclesPerYield RTL_SYSTEM_GLOBAL_DATA_ID = 11
	GlobalDataIdSafeBootMode RTL_SYSTEM_GLOBAL_DATA_ID = 12
	GlobalDataIdLastSystemRITEventTickCount RTL_SYSTEM_GLOBAL_DATA_ID = 13
)

// enum DEP_SYSTEM_POLICY_TYPE
type DEP_SYSTEM_POLICY_TYPE int32
const (
	DEPPolicyAlwaysOff DEP_SYSTEM_POLICY_TYPE = 0
	DEPPolicyAlwaysOn DEP_SYSTEM_POLICY_TYPE = 1
	DEPPolicyOptIn DEP_SYSTEM_POLICY_TYPE = 2
	DEPPolicyOptOut DEP_SYSTEM_POLICY_TYPE = 3
	DEPTotalPolicyCount DEP_SYSTEM_POLICY_TYPE = 4
)


// structs

type GROUP_AFFINITY struct {
	Mask uintptr
	Group uint16
	Reserved [3]uint16
}

type SYSTEM_INFO_Anonymous__Anonymous_ struct {
	WProcessorArchitecture PROCESSOR_ARCHITECTURE
	WReserved uint16
}

type SYSTEM_INFO_Anonymous_ struct {
	 SYSTEM_INFO_Anonymous__Anonymous_
}

func (this *SYSTEM_INFO_Anonymous_) DwOemId() *uint32{
	return (*uint32)(unsafe.Pointer(this))
}

func (this *SYSTEM_INFO_Anonymous_) DwOemIdVal() uint32{
	return *(*uint32)(unsafe.Pointer(this))
}

type SYSTEM_INFO struct {
	SYSTEM_INFO_Anonymous_
	DwPageSize uint32
	LpMinimumApplicationAddress unsafe.Pointer
	LpMaximumApplicationAddress unsafe.Pointer
	DwActiveProcessorMask uintptr
	DwNumberOfProcessors uint32
	DwProcessorType uint32
	DwAllocationGranularity uint32
	WProcessorLevel uint16
	WProcessorRevision uint16
}

type MEMORYSTATUSEX struct {
	DwLength uint32
	DwMemoryLoad uint32
	UllTotalPhys uint64
	UllAvailPhys uint64
	UllTotalPageFile uint64
	UllAvailPageFile uint64
	UllTotalVirtual uint64
	UllAvailVirtual uint64
	UllAvailExtendedVirtual uint64
}

type CACHE_DESCRIPTOR struct {
	Level uint8
	Associativity uint8
	LineSize uint16
	Size uint32
	Type PROCESSOR_CACHE_TYPE
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__ProcessorCore_ struct {
	Flags uint8
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__NumaNode_ struct {
	NodeNumber uint32
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_ struct {
	Data [2]uint64
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) ProcessorCore() *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__ProcessorCore_{
	return (*SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__ProcessorCore_)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) ProcessorCoreVal() SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__ProcessorCore_{
	return *(*SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__ProcessorCore_)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) NumaNode() *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__NumaNode_{
	return (*SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__NumaNode_)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) NumaNodeVal() SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__NumaNode_{
	return *(*SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous__NumaNode_)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) Cache() *CACHE_DESCRIPTOR{
	return (*CACHE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) CacheVal() CACHE_DESCRIPTOR{
	return *(*CACHE_DESCRIPTOR)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) Reserved() *[2]uint64{
	return (*[2]uint64)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_) ReservedVal() [2]uint64{
	return *(*[2]uint64)(unsafe.Pointer(this))
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION struct {
	ProcessorMask uintptr
	Relationship LOGICAL_PROCESSOR_RELATIONSHIP
	SYSTEM_LOGICAL_PROCESSOR_INFORMATION_Anonymous_
}

type PROCESSOR_RELATIONSHIP struct {
	Flags uint8
	EfficiencyClass uint8
	Reserved [20]uint8
	GroupCount uint16
	GroupMask [1]GROUP_AFFINITY
}

type NUMA_NODE_RELATIONSHIP_Anonymous_ struct {
	Data [2]uint64
}

func (this *NUMA_NODE_RELATIONSHIP_Anonymous_) GroupMask() *GROUP_AFFINITY{
	return (*GROUP_AFFINITY)(unsafe.Pointer(this))
}

func (this *NUMA_NODE_RELATIONSHIP_Anonymous_) GroupMaskVal() GROUP_AFFINITY{
	return *(*GROUP_AFFINITY)(unsafe.Pointer(this))
}

func (this *NUMA_NODE_RELATIONSHIP_Anonymous_) GroupMasks() *[1]GROUP_AFFINITY{
	return (*[1]GROUP_AFFINITY)(unsafe.Pointer(this))
}

func (this *NUMA_NODE_RELATIONSHIP_Anonymous_) GroupMasksVal() [1]GROUP_AFFINITY{
	return *(*[1]GROUP_AFFINITY)(unsafe.Pointer(this))
}

type NUMA_NODE_RELATIONSHIP struct {
	NodeNumber uint32
	Reserved [18]uint8
	GroupCount uint16
	NUMA_NODE_RELATIONSHIP_Anonymous_
}

type CACHE_RELATIONSHIP_Anonymous_ struct {
	Data [2]uint64
}

func (this *CACHE_RELATIONSHIP_Anonymous_) GroupMask() *GROUP_AFFINITY{
	return (*GROUP_AFFINITY)(unsafe.Pointer(this))
}

func (this *CACHE_RELATIONSHIP_Anonymous_) GroupMaskVal() GROUP_AFFINITY{
	return *(*GROUP_AFFINITY)(unsafe.Pointer(this))
}

func (this *CACHE_RELATIONSHIP_Anonymous_) GroupMasks() *[1]GROUP_AFFINITY{
	return (*[1]GROUP_AFFINITY)(unsafe.Pointer(this))
}

func (this *CACHE_RELATIONSHIP_Anonymous_) GroupMasksVal() [1]GROUP_AFFINITY{
	return *(*[1]GROUP_AFFINITY)(unsafe.Pointer(this))
}

type CACHE_RELATIONSHIP struct {
	Level uint8
	Associativity uint8
	LineSize uint16
	CacheSize uint32
	Type PROCESSOR_CACHE_TYPE
	Reserved [18]uint8
	GroupCount uint16
	CACHE_RELATIONSHIP_Anonymous_
}

type PROCESSOR_GROUP_INFO struct {
	MaximumProcessorCount uint8
	ActiveProcessorCount uint8
	Reserved [38]uint8
	ActiveProcessorMask uintptr
}

type GROUP_RELATIONSHIP struct {
	MaximumGroupCount uint16
	ActiveGroupCount uint16
	Reserved [20]uint8
	GroupInfo [1]PROCESSOR_GROUP_INFO
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_ struct {
	Data [9]uint64
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) Processor() *PROCESSOR_RELATIONSHIP{
	return (*PROCESSOR_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) ProcessorVal() PROCESSOR_RELATIONSHIP{
	return *(*PROCESSOR_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) NumaNode() *NUMA_NODE_RELATIONSHIP{
	return (*NUMA_NODE_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) NumaNodeVal() NUMA_NODE_RELATIONSHIP{
	return *(*NUMA_NODE_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) Cache() *CACHE_RELATIONSHIP{
	return (*CACHE_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) CacheVal() CACHE_RELATIONSHIP{
	return *(*CACHE_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) Group() *GROUP_RELATIONSHIP{
	return (*GROUP_RELATIONSHIP)(unsafe.Pointer(this))
}

func (this *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_) GroupVal() GROUP_RELATIONSHIP{
	return *(*GROUP_RELATIONSHIP)(unsafe.Pointer(this))
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX struct {
	Relationship LOGICAL_PROCESSOR_RELATIONSHIP
	Size uint32
	SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX_Anonymous_
}

type SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous1__Anonymous_ struct {
	Bitfield_ uint8
}

type SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous1_ struct {
	 SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous1__Anonymous_
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous1_) AllFlags() *uint8{
	return (*uint8)(unsafe.Pointer(this))
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous1_) AllFlagsVal() uint8{
	return *(*uint8)(unsafe.Pointer(this))
}

type SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous2_ struct {
	Data [1]uint32
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous2_) Reserved() *uint32{
	return (*uint32)(unsafe.Pointer(this))
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous2_) ReservedVal() uint32{
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous2_) SchedulingClass() *uint8{
	return (*uint8)(unsafe.Pointer(this))
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous2_) SchedulingClassVal() uint8{
	return *(*uint8)(unsafe.Pointer(this))
}

type SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet_ struct {
	Id uint32
	Group uint16
	LogicalProcessorIndex uint8
	CoreIndex uint8
	LastLevelCacheIndex uint8
	NumaNodeIndex uint8
	EfficiencyClass uint8
	SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous1_
	SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet__Anonymous2_
	AllocationTag uint64
}

type SYSTEM_CPU_SET_INFORMATION_Anonymous_ struct {
	Data [3]uint64
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous_) CpuSet() *SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet_{
	return (*SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet_)(unsafe.Pointer(this))
}

func (this *SYSTEM_CPU_SET_INFORMATION_Anonymous_) CpuSetVal() SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet_{
	return *(*SYSTEM_CPU_SET_INFORMATION_Anonymous__CpuSet_)(unsafe.Pointer(this))
}

type SYSTEM_CPU_SET_INFORMATION struct {
	Size uint32
	Type CPU_SET_INFORMATION_TYPE
	SYSTEM_CPU_SET_INFORMATION_Anonymous_
}

type SYSTEM_POOL_ZEROING_INFORMATION struct {
	PoolZeroingSupportPresent BOOLEAN
}

type SYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION struct {
	CycleTime uint64
}

type SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION struct {
	Bitfield_ uint32
}

type OSVERSIONINFOA struct {
	DwOSVersionInfoSize uint32
	DwMajorVersion uint32
	DwMinorVersion uint32
	DwBuildNumber uint32
	DwPlatformId uint32
	SzCSDVersion [128]CHAR
}

type OSVERSIONINFO = OSVERSIONINFOW
type OSVERSIONINFOW struct {
	DwOSVersionInfoSize uint32
	DwMajorVersion uint32
	DwMinorVersion uint32
	DwBuildNumber uint32
	DwPlatformId uint32
	SzCSDVersion [128]uint16
}

type OSVERSIONINFOEXA struct {
	DwOSVersionInfoSize uint32
	DwMajorVersion uint32
	DwMinorVersion uint32
	DwBuildNumber uint32
	DwPlatformId uint32
	SzCSDVersion [128]CHAR
	WServicePackMajor uint16
	WServicePackMinor uint16
	WSuiteMask uint16
	WProductType uint8
	WReserved uint8
}

type OSVERSIONINFOEX = OSVERSIONINFOEXW
type OSVERSIONINFOEXW struct {
	DwOSVersionInfoSize uint32
	DwMajorVersion uint32
	DwMinorVersion uint32
	DwBuildNumber uint32
	DwPlatformId uint32
	SzCSDVersion [128]uint16
	WServicePackMajor uint16
	WServicePackMinor uint16
	WSuiteMask uint16
	WProductType uint8
	WReserved uint8
}

type MEMORYSTATUS struct {
	DwLength uint32
	DwMemoryLoad uint32
	DwTotalPhys uintptr
	DwAvailPhys uintptr
	DwTotalPageFile uintptr
	DwAvailPageFile uintptr
	DwTotalVirtual uintptr
	DwAvailVirtual uintptr
}


// func types

type PGET_SYSTEM_WOW64_DIRECTORY_A func(lpBuffer *uint8, uSize uint32) uint32

type PGET_SYSTEM_WOW64_DIRECTORY_W func(lpBuffer *uint16, uSize uint32) uint32


var (
	pGlobalMemoryStatusEx uintptr
	pGetSystemInfo uintptr
	pGetSystemTime uintptr
	pGetSystemTimeAsFileTime uintptr
	pGetLocalTime uintptr
	pIsUserCetAvailableInEnvironment uintptr
	pGetSystemLeapSecondInformation uintptr
	pGetVersion uintptr
	pSetLocalTime uintptr
	pGetTickCount uintptr
	pGetTickCount64 uintptr
	pGetSystemTimeAdjustment uintptr
	pGetSystemDirectoryA uintptr
	pGetSystemDirectoryW uintptr
	pGetWindowsDirectoryA uintptr
	pGetWindowsDirectoryW uintptr
	pGetSystemWindowsDirectoryA uintptr
	pGetSystemWindowsDirectoryW uintptr
	pGetComputerNameExA uintptr
	pGetComputerNameExW uintptr
	pSetComputerNameExW uintptr
	pSetSystemTime uintptr
	pGetVersionExA uintptr
	pGetVersionExW uintptr
	pGetLogicalProcessorInformation uintptr
	pGetLogicalProcessorInformationEx uintptr
	pGetNativeSystemInfo uintptr
	pGetSystemTimePreciseAsFileTime uintptr
	pGetProductInfo uintptr
	pVerSetConditionMask uintptr
	pEnumSystemFirmwareTables uintptr
	pGetSystemFirmwareTable uintptr
	pDnsHostnameToComputerNameExW uintptr
	pGetPhysicallyInstalledSystemMemory uintptr
	pSetComputerNameEx2W uintptr
	pSetSystemTimeAdjustment uintptr
	pGetProcessorSystemCycleTime uintptr
	pSetComputerNameA uintptr
	pSetComputerNameW uintptr
	pSetComputerNameExA uintptr
	pGetSystemCpuSetInformation uintptr
	pGetSystemWow64DirectoryA uintptr
	pGetSystemWow64DirectoryW uintptr
	pIsWow64GuestMachineSupported uintptr
	pGlobalMemoryStatus uintptr
	pGetSystemDEPPolicy uintptr
	pGetFirmwareType uintptr
	pVerifyVersionInfoA uintptr
	pVerifyVersionInfoW uintptr
)

func GlobalMemoryStatusEx(lpBuffer *MEMORYSTATUSEX) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGlobalMemoryStatusEx, libKernel32, "GlobalMemoryStatusEx")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetSystemInfo(lpSystemInfo *SYSTEM_INFO) {
	addr := lazyAddr(&pGetSystemInfo, libKernel32, "GetSystemInfo")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemInfo)))
}

func GetSystemTime(lpSystemTime *SYSTEMTIME) {
	addr := lazyAddr(&pGetSystemTime, libKernel32, "GetSystemTime")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTime)))
}

func GetSystemTimeAsFileTime(lpSystemTimeAsFileTime *FILETIME) {
	addr := lazyAddr(&pGetSystemTimeAsFileTime, libKernel32, "GetSystemTimeAsFileTime")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTimeAsFileTime)))
}

func GetLocalTime(lpSystemTime *SYSTEMTIME) {
	addr := lazyAddr(&pGetLocalTime, libKernel32, "GetLocalTime")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTime)))
}

func IsUserCetAvailableInEnvironment(UserCetEnvironment USER_CET_ENVIRONMENT) BOOL {
	addr := lazyAddr(&pIsUserCetAvailableInEnvironment, libKernel32, "IsUserCetAvailableInEnvironment")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(UserCetEnvironment))
	return BOOL(ret)
}

func GetSystemLeapSecondInformation(Enabled *BOOL, Flags *uint32) BOOL {
	addr := lazyAddr(&pGetSystemLeapSecondInformation, libKernel32, "GetSystemLeapSecondInformation")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Enabled)), uintptr(unsafe.Pointer(Flags)))
	return BOOL(ret)
}

func GetVersion() uint32 {
	addr := lazyAddr(&pGetVersion, libKernel32, "GetVersion")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func SetLocalTime(lpSystemTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetLocalTime, libKernel32, "SetLocalTime")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetTickCount() uint32 {
	addr := lazyAddr(&pGetTickCount, libKernel32, "GetTickCount")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func GetTickCount64() uint64 {
	addr := lazyAddr(&pGetTickCount64, libKernel32, "GetTickCount64")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint64(ret)
}

func GetSystemTimeAdjustment(lpTimeAdjustment *uint32, lpTimeIncrement *uint32, lpTimeAdjustmentDisabled *BOOL) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemTimeAdjustment, libKernel32, "GetSystemTimeAdjustment")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpTimeAdjustment)), uintptr(unsafe.Pointer(lpTimeIncrement)), uintptr(unsafe.Pointer(lpTimeAdjustmentDisabled)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetSystemDirectoryA(lpBuffer *uint8, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemDirectoryA, libKernel32, "GetSystemDirectoryA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

var GetSystemDirectory = GetSystemDirectoryW
func GetSystemDirectoryW(lpBuffer *uint16, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemDirectoryW, libKernel32, "GetSystemDirectoryW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

func GetWindowsDirectoryA(lpBuffer *uint8, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetWindowsDirectoryA, libKernel32, "GetWindowsDirectoryA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

var GetWindowsDirectory = GetWindowsDirectoryW
func GetWindowsDirectoryW(lpBuffer *uint16, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetWindowsDirectoryW, libKernel32, "GetWindowsDirectoryW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

func GetSystemWindowsDirectoryA(lpBuffer *uint8, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemWindowsDirectoryA, libKernel32, "GetSystemWindowsDirectoryA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

var GetSystemWindowsDirectory = GetSystemWindowsDirectoryW
func GetSystemWindowsDirectoryW(lpBuffer *uint16, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemWindowsDirectoryW, libKernel32, "GetSystemWindowsDirectoryW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

func GetComputerNameExA(NameType COMPUTER_NAME_FORMAT, lpBuffer *uint8, nSize *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetComputerNameExA, libKernel32, "GetComputerNameExA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(NameType), uintptr(unsafe.Pointer(lpBuffer)), uintptr(unsafe.Pointer(nSize)))
	return BOOL(ret), WIN32_ERROR(err)
}

var GetComputerNameEx = GetComputerNameExW
func GetComputerNameExW(NameType COMPUTER_NAME_FORMAT, lpBuffer *uint16, nSize *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetComputerNameExW, libKernel32, "GetComputerNameExW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(NameType), uintptr(unsafe.Pointer(lpBuffer)), uintptr(unsafe.Pointer(nSize)))
	return BOOL(ret), WIN32_ERROR(err)
}

var SetComputerNameEx = SetComputerNameExW
func SetComputerNameExW(NameType COMPUTER_NAME_FORMAT, lpBuffer PWSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetComputerNameExW, libKernel32, "SetComputerNameExW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(NameType), uintptr(unsafe.Pointer(lpBuffer)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetSystemTime(lpSystemTime *SYSTEMTIME) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetSystemTime, libKernel32, "SetSystemTime")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTime)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetVersionExA(lpVersionInformation *OSVERSIONINFOA) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetVersionExA, libKernel32, "GetVersionExA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpVersionInformation)))
	return BOOL(ret), WIN32_ERROR(err)
}

var GetVersionEx = GetVersionExW
func GetVersionExW(lpVersionInformation *OSVERSIONINFOW) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetVersionExW, libKernel32, "GetVersionExW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpVersionInformation)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetLogicalProcessorInformation(Buffer *SYSTEM_LOGICAL_PROCESSOR_INFORMATION, ReturnedLength *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetLogicalProcessorInformation, libKernel32, "GetLogicalProcessorInformation")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(ReturnedLength)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetLogicalProcessorInformationEx(RelationshipType LOGICAL_PROCESSOR_RELATIONSHIP, Buffer *SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX, ReturnedLength *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetLogicalProcessorInformationEx, libKernel32, "GetLogicalProcessorInformationEx")
	ret, _,  err := syscall.SyscallN(addr, uintptr(RelationshipType), uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(ReturnedLength)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetNativeSystemInfo(lpSystemInfo *SYSTEM_INFO) {
	addr := lazyAddr(&pGetNativeSystemInfo, libKernel32, "GetNativeSystemInfo")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemInfo)))
}

func GetSystemTimePreciseAsFileTime(lpSystemTimeAsFileTime *FILETIME) {
	addr := lazyAddr(&pGetSystemTimePreciseAsFileTime, libKernel32, "GetSystemTimePreciseAsFileTime")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTimeAsFileTime)))
}

func GetProductInfo(dwOSMajorVersion uint32, dwOSMinorVersion uint32, dwSpMajorVersion uint32, dwSpMinorVersion uint32, pdwReturnedProductType *OS_PRODUCT_TYPE) BOOL {
	addr := lazyAddr(&pGetProductInfo, libKernel32, "GetProductInfo")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(dwOSMajorVersion), uintptr(dwOSMinorVersion), uintptr(dwSpMajorVersion), uintptr(dwSpMinorVersion), uintptr(unsafe.Pointer(pdwReturnedProductType)))
	return BOOL(ret)
}

func VerSetConditionMask(ConditionMask uint64, TypeMask VER_FLAGS, Condition uint8) uint64 {
	addr := lazyAddr(&pVerSetConditionMask, libKernel32, "VerSetConditionMask")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(ConditionMask), uintptr(TypeMask), uintptr(Condition))
	return uint64(ret)
}

func EnumSystemFirmwareTables(FirmwareTableProviderSignature FIRMWARE_TABLE_PROVIDER, pFirmwareTableEnumBuffer *FIRMWARE_TABLE_ID, BufferSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pEnumSystemFirmwareTables, libKernel32, "EnumSystemFirmwareTables")
	ret, _,  err := syscall.SyscallN(addr, uintptr(FirmwareTableProviderSignature), uintptr(unsafe.Pointer(pFirmwareTableEnumBuffer)), uintptr(BufferSize))
	return uint32(ret), WIN32_ERROR(err)
}

func GetSystemFirmwareTable(FirmwareTableProviderSignature FIRMWARE_TABLE_PROVIDER, FirmwareTableID FIRMWARE_TABLE_ID, pFirmwareTableBuffer unsafe.Pointer, BufferSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemFirmwareTable, libKernel32, "GetSystemFirmwareTable")
	ret, _,  err := syscall.SyscallN(addr, uintptr(FirmwareTableProviderSignature), uintptr(FirmwareTableID), uintptr(pFirmwareTableBuffer), uintptr(BufferSize))
	return uint32(ret), WIN32_ERROR(err)
}

func DnsHostnameToComputerNameExW(Hostname PWSTR, ComputerName *uint16, nSize *uint32) BOOL {
	addr := lazyAddr(&pDnsHostnameToComputerNameExW, libKernel32, "DnsHostnameToComputerNameExW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Hostname)), uintptr(unsafe.Pointer(ComputerName)), uintptr(unsafe.Pointer(nSize)))
	return BOOL(ret)
}

func GetPhysicallyInstalledSystemMemory(TotalMemoryInKilobytes *uint64) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetPhysicallyInstalledSystemMemory, libKernel32, "GetPhysicallyInstalledSystemMemory")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(TotalMemoryInKilobytes)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetComputerNameEx2W(NameType COMPUTER_NAME_FORMAT, Flags uint32, lpBuffer PWSTR) BOOL {
	addr := lazyAddr(&pSetComputerNameEx2W, libKernel32, "SetComputerNameEx2W")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(NameType), uintptr(Flags), uintptr(unsafe.Pointer(lpBuffer)))
	return BOOL(ret)
}

func SetSystemTimeAdjustment(dwTimeAdjustment uint32, bTimeAdjustmentDisabled BOOL) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetSystemTimeAdjustment, libKernel32, "SetSystemTimeAdjustment")
	ret, _,  err := syscall.SyscallN(addr, uintptr(dwTimeAdjustment), uintptr(bTimeAdjustmentDisabled))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetProcessorSystemCycleTime(Group uint16, Buffer *SYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION, ReturnedLength *uint32) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetProcessorSystemCycleTime, libKernel32, "GetProcessorSystemCycleTime")
	ret, _,  err := syscall.SyscallN(addr, uintptr(Group), uintptr(unsafe.Pointer(Buffer)), uintptr(unsafe.Pointer(ReturnedLength)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetComputerNameA(lpComputerName PSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetComputerNameA, libKernel32, "SetComputerNameA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpComputerName)))
	return BOOL(ret), WIN32_ERROR(err)
}

var SetComputerName = SetComputerNameW
func SetComputerNameW(lpComputerName PWSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetComputerNameW, libKernel32, "SetComputerNameW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpComputerName)))
	return BOOL(ret), WIN32_ERROR(err)
}

func SetComputerNameExA(NameType COMPUTER_NAME_FORMAT, lpBuffer PSTR) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetComputerNameExA, libKernel32, "SetComputerNameExA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(NameType), uintptr(unsafe.Pointer(lpBuffer)))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetSystemCpuSetInformation(Information *SYSTEM_CPU_SET_INFORMATION, BufferLength uint32, ReturnedLength *uint32, Process HANDLE, Flags uint32) BOOL {
	addr := lazyAddr(&pGetSystemCpuSetInformation, libKernel32, "GetSystemCpuSetInformation")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Information)), uintptr(BufferLength), uintptr(unsafe.Pointer(ReturnedLength)), Process, uintptr(Flags))
	return BOOL(ret)
}

func GetSystemWow64DirectoryA(lpBuffer *uint8, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemWow64DirectoryA, libKernel32, "GetSystemWow64DirectoryA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

var GetSystemWow64Directory = GetSystemWow64DirectoryW
func GetSystemWow64DirectoryW(lpBuffer *uint16, uSize uint32) (uint32, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemWow64DirectoryW, libKernel32, "GetSystemWow64DirectoryW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)), uintptr(uSize))
	return uint32(ret), WIN32_ERROR(err)
}

func IsWow64GuestMachineSupported(WowGuestMachine uint16, MachineIsSupported *BOOL) (HRESULT, WIN32_ERROR) {
	addr := lazyAddr(&pIsWow64GuestMachineSupported, libKernel32, "IsWow64GuestMachineSupported")
	ret, _,  err := syscall.SyscallN(addr, uintptr(WowGuestMachine), uintptr(unsafe.Pointer(MachineIsSupported)))
	return HRESULT(ret), WIN32_ERROR(err)
}

func GlobalMemoryStatus(lpBuffer *MEMORYSTATUS) {
	addr := lazyAddr(&pGlobalMemoryStatus, libKernel32, "GlobalMemoryStatus")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpBuffer)))
}

func GetSystemDEPPolicy() DEP_SYSTEM_POLICY_TYPE {
	addr := lazyAddr(&pGetSystemDEPPolicy, libKernel32, "GetSystemDEPPolicy")
	ret, _,  _ := syscall.SyscallN(addr)
	return DEP_SYSTEM_POLICY_TYPE(ret)
}

func GetFirmwareType(FirmwareType *FIRMWARE_TYPE) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetFirmwareType, libKernel32, "GetFirmwareType")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(FirmwareType)))
	return BOOL(ret), WIN32_ERROR(err)
}

func VerifyVersionInfoA(lpVersionInformation *OSVERSIONINFOEXA, dwTypeMask VER_FLAGS, dwlConditionMask uint64) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pVerifyVersionInfoA, libKernel32, "VerifyVersionInfoA")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpVersionInformation)), uintptr(dwTypeMask), uintptr(dwlConditionMask))
	return BOOL(ret), WIN32_ERROR(err)
}

var VerifyVersionInfo = VerifyVersionInfoW
func VerifyVersionInfoW(lpVersionInformation *OSVERSIONINFOEXW, dwTypeMask VER_FLAGS, dwlConditionMask uint64) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pVerifyVersionInfoW, libKernel32, "VerifyVersionInfoW")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpVersionInformation)), uintptr(dwTypeMask), uintptr(dwlConditionMask))
	return BOOL(ret), WIN32_ERROR(err)
}


