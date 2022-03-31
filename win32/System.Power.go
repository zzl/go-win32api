package win32

import "unsafe"
import "syscall"

type HPOWERNOTIFY = uintptr

const (
	BATTERY_UNKNOWN_CAPACITY uint32 = 4294967295
	UNKNOWN_CAPACITY uint32 = 4294967295
	BATTERY_SYSTEM_BATTERY uint32 = 2147483648
	BATTERY_CAPACITY_RELATIVE uint32 = 1073741824
	BATTERY_IS_SHORT_TERM uint32 = 536870912
	BATTERY_SEALED uint32 = 268435456
	BATTERY_SET_CHARGE_SUPPORTED uint32 = 1
	BATTERY_SET_DISCHARGE_SUPPORTED uint32 = 2
	BATTERY_SET_CHARGINGSOURCE_SUPPORTED uint32 = 4
	BATTERY_SET_CHARGER_ID_SUPPORTED uint32 = 8
	BATTERY_UNKNOWN_TIME uint32 = 4294967295
	BATTERY_UNKNOWN_CURRENT uint32 = 4294967295
	UNKNOWN_CURRENT uint32 = 4294967295
	BATTERY_USB_CHARGER_STATUS_FN_DEFAULT_USB uint32 = 1
	BATTERY_USB_CHARGER_STATUS_UCM_PD uint32 = 2
	BATTERY_UNKNOWN_VOLTAGE uint32 = 4294967295
	BATTERY_UNKNOWN_RATE uint32 = 2147483648
	UNKNOWN_RATE uint32 = 2147483648
	UNKNOWN_VOLTAGE uint32 = 4294967295
	BATTERY_POWER_ON_LINE uint32 = 1
	BATTERY_DISCHARGING uint32 = 2
	BATTERY_CHARGING uint32 = 4
	BATTERY_CRITICAL uint32 = 8
	MAX_BATTERY_STRING_SIZE uint32 = 128
	IOCTL_BATTERY_QUERY_TAG uint32 = 2703424
	IOCTL_BATTERY_QUERY_INFORMATION uint32 = 2703428
	IOCTL_BATTERY_SET_INFORMATION uint32 = 2719816
	IOCTL_BATTERY_QUERY_STATUS uint32 = 2703436
	IOCTL_BATTERY_CHARGING_SOURCE_CHANGE uint32 = 2703440
	BATTERY_TAG_INVALID uint32 = 0
	MAX_ACTIVE_COOLING_LEVELS uint32 = 10
	ACTIVE_COOLING uint32 = 0
	PASSIVE_COOLING uint32 = 1
	TZ_ACTIVATION_REASON_THERMAL uint32 = 1
	TZ_ACTIVATION_REASON_CURRENT uint32 = 2
	THERMAL_POLICY_VERSION_1 uint32 = 1
	THERMAL_POLICY_VERSION_2 uint32 = 2
	IOCTL_THERMAL_QUERY_INFORMATION uint32 = 2703488
	IOCTL_THERMAL_SET_COOLING_POLICY uint32 = 2719876
	IOCTL_RUN_ACTIVE_COOLING_METHOD uint32 = 2719880
	IOCTL_THERMAL_SET_PASSIVE_LIMIT uint32 = 2719884
	IOCTL_THERMAL_READ_TEMPERATURE uint32 = 2703504
	IOCTL_THERMAL_READ_POLICY uint32 = 2703508
	IOCTL_QUERY_LID uint32 = 2703552
	IOCTL_NOTIFY_SWITCH_EVENT uint32 = 2703616
	IOCTL_GET_SYS_BUTTON_CAPS uint32 = 2703680
	IOCTL_GET_SYS_BUTTON_EVENT uint32 = 2703684
	SYS_BUTTON_POWER uint32 = 1
	SYS_BUTTON_SLEEP uint32 = 2
	SYS_BUTTON_LID uint32 = 4
	SYS_BUTTON_WAKE uint32 = 2147483648
	SYS_BUTTON_LID_STATE_MASK uint32 = 196608
	SYS_BUTTON_LID_OPEN uint32 = 65536
	SYS_BUTTON_LID_CLOSED uint32 = 131072
	SYS_BUTTON_LID_INITIAL uint32 = 262144
	SYS_BUTTON_LID_CHANGED uint32 = 524288
	IOCTL_GET_PROCESSOR_OBJ_INFO uint32 = 2703744
	THERMAL_COOLING_INTERFACE_VERSION uint32 = 1
	THERMAL_DEVICE_INTERFACE_VERSION uint32 = 1
	IOCTL_SET_SYS_MESSAGE_INDICATOR uint32 = 2720192
	IOCTL_SET_WAKE_ALARM_VALUE uint32 = 2720256
	IOCTL_SET_WAKE_ALARM_POLICY uint32 = 2720260
	IOCTL_GET_WAKE_ALARM_VALUE uint32 = 2736648
	IOCTL_GET_WAKE_ALARM_POLICY uint32 = 2736652
	ACPI_TIME_ADJUST_DAYLIGHT uint32 = 1
	ACPI_TIME_IN_DAYLIGHT uint32 = 2
	ACPI_TIME_ZONE_UNKNOWN uint32 = 2047
	IOCTL_ACPI_GET_REAL_TIME uint32 = 2703888
	IOCTL_ACPI_SET_REAL_TIME uint32 = 2720276
	IOCTL_GET_WAKE_ALARM_SYSTEM_POWERSTATE uint32 = 2703896
	BATTERY_MINIPORT_UPDATE_DATA_VER_1 uint32 = 1
	BATTERY_MINIPORT_UPDATE_DATA_VER_2 uint32 = 2
	BATTERY_CLASS_MAJOR_VERSION uint32 = 1
	BATTERY_CLASS_MINOR_VERSION uint32 = 0
	BATTERY_CLASS_MINOR_VERSION_1 uint32 = 1
	IOCTL_EMI_GET_VERSION uint32 = 2244608
	IOCTL_EMI_GET_METADATA_SIZE uint32 = 2244612
	IOCTL_EMI_GET_METADATA uint32 = 2244616
	IOCTL_EMI_GET_MEASUREMENT uint32 = 2244620
	EMI_NAME_MAX uint32 = 16
	EMI_VERSION_V1 uint32 = 1
	EMI_VERSION_V2 uint32 = 2
	EFFECTIVE_POWER_MODE_V1 uint32 = 1
	EFFECTIVE_POWER_MODE_V2 uint32 = 2
	EnableSysTrayBatteryMeter uint32 = 1
	EnableMultiBatteryDisplay uint32 = 2
	EnablePasswordLogon uint32 = 4
	EnableWakeOnRing uint32 = 8
	EnableVideoDimDisplay uint32 = 16
	POWER_ATTRIBUTE_HIDE uint32 = 1
	POWER_ATTRIBUTE_SHOW_AOAC uint32 = 2
	DEVICEPOWER_HARDWAREID uint32 = 2147483648
	DEVICEPOWER_AND_OPERATION uint32 = 1073741824
	DEVICEPOWER_FILTER_DEVICES_PRESENT uint32 = 536870912
	DEVICEPOWER_FILTER_HARDWARE uint32 = 268435456
	DEVICEPOWER_FILTER_WAKEENABLED uint32 = 134217728
	DEVICEPOWER_FILTER_WAKEPROGRAMMABLE uint32 = 67108864
	DEVICEPOWER_FILTER_ON_NAME uint32 = 33554432
	DEVICEPOWER_SET_WAKEENABLED uint32 = 1
	DEVICEPOWER_CLEAR_WAKEENABLED uint32 = 2
	PDCAP_S0_SUPPORTED uint32 = 65536
	PDCAP_S1_SUPPORTED uint32 = 131072
	PDCAP_S2_SUPPORTED uint32 = 262144
	PDCAP_S3_SUPPORTED uint32 = 524288
	PDCAP_WAKE_FROM_S0_SUPPORTED uint32 = 1048576
	PDCAP_WAKE_FROM_S1_SUPPORTED uint32 = 2097152
	PDCAP_WAKE_FROM_S2_SUPPORTED uint32 = 4194304
	PDCAP_WAKE_FROM_S3_SUPPORTED uint32 = 8388608
	PDCAP_S4_SUPPORTED uint32 = 16777216
	PDCAP_S5_SUPPORTED uint32 = 33554432
	THERMAL_EVENT_VERSION uint32 = 1
)

var (
	GUID_DEVICE_BATTERY syscall.GUID = syscall.GUID{0x72631e54, 0x78a4, 0x11d0, 
	[8]byte{0xbc, 0xf7, 0x00, 0xaa, 0x00, 0xb7, 0xb3, 0x2a}}
	GUID_DEVICE_APPLICATIONLAUNCH_BUTTON syscall.GUID = syscall.GUID{0x629758ee, 0x986e, 0x4d9e, 
	[8]byte{0x8e, 0x47, 0xde, 0x27, 0xf8, 0xab, 0x05, 0x4d}}
	GUID_DEVICE_SYS_BUTTON syscall.GUID = syscall.GUID{0x4afa3d53, 0x74a7, 0x11d0, 
	[8]byte{0xbe, 0x5e, 0x00, 0xa0, 0xc9, 0x06, 0x28, 0x57}}
	GUID_DEVICE_LID syscall.GUID = syscall.GUID{0x4afa3d52, 0x74a7, 0x11d0, 
	[8]byte{0xbe, 0x5e, 0x00, 0xa0, 0xc9, 0x06, 0x28, 0x57}}
	GUID_DEVICE_THERMAL_ZONE syscall.GUID = syscall.GUID{0x4afa3d51, 0x74a7, 0x11d0, 
	[8]byte{0xbe, 0x5e, 0x00, 0xa0, 0xc9, 0x06, 0x28, 0x57}}
	GUID_DEVICE_FAN syscall.GUID = syscall.GUID{0x05ecd13d, 0x81da, 0x4a2a, 
	[8]byte{0x8a, 0x4c, 0x52, 0x4f, 0x23, 0xdd, 0x4d, 0xc9}}
	GUID_DEVICE_PROCESSOR syscall.GUID = syscall.GUID{0x97fadb10, 0x4e33, 0x40ae, 
	[8]byte{0x35, 0x9c, 0x8b, 0xef, 0x02, 0x9d, 0xbd, 0xd0}}
	GUID_DEVICE_MEMORY syscall.GUID = syscall.GUID{0x3fd0f03d, 0x92e0, 0x45fb, 
	[8]byte{0xb7, 0x5c, 0x5e, 0xd8, 0xff, 0xb0, 0x10, 0x21}}
	GUID_DEVICE_ACPI_TIME syscall.GUID = syscall.GUID{0x97f99bf6, 0x4497, 0x4f18, 
	[8]byte{0xbb, 0x22, 0x4b, 0x9f, 0xb2, 0xfb, 0xef, 0x9c}}
	GUID_DEVICE_MESSAGE_INDICATOR syscall.GUID = syscall.GUID{0xcd48a365, 0xfa94, 0x4ce2, 
	[8]byte{0xa2, 0x32, 0xa1, 0xb7, 0x64, 0xe5, 0xd8, 0xb4}}
	GUID_CLASS_INPUT syscall.GUID = syscall.GUID{0x4d1e55b2, 0xf16f, 0x11cf, 
	[8]byte{0x88, 0xcb, 0x00, 0x11, 0x11, 0x00, 0x00, 0x30}}
	GUID_DEVINTERFACE_THERMAL_COOLING syscall.GUID = syscall.GUID{0xdbe4373d, 0x3c81, 0x40cb, 
	[8]byte{0xac, 0xe4, 0xe0, 0xe5, 0xd0, 0x5f, 0x0c, 0x9f}}
	GUID_DEVINTERFACE_THERMAL_MANAGER syscall.GUID = syscall.GUID{0x927ec093, 0x69a4, 0x4bc0, 
	[8]byte{0xbd, 0x02, 0x71, 0x16, 0x64, 0x71, 0x44, 0x63}}
	BATTERY_STATUS_WMI_GUID syscall.GUID = syscall.GUID{0xfc4670d1, 0xebbf, 0x416e, 
	[8]byte{0x87, 0xce, 0x37, 0x4a, 0x4e, 0xbc, 0x11, 0x1a}}
	BATTERY_RUNTIME_WMI_GUID syscall.GUID = syscall.GUID{0x535a3767, 0x1ac2, 0x49bc, 
	[8]byte{0xa0, 0x77, 0x3f, 0x7a, 0x02, 0xe4, 0x0a, 0xec}}
	BATTERY_TEMPERATURE_WMI_GUID syscall.GUID = syscall.GUID{0x1a52a14d, 0xadce, 0x4a44, 
	[8]byte{0x9a, 0x3e, 0xc8, 0xd8, 0xf1, 0x5f, 0xf2, 0xc2}}
	BATTERY_FULL_CHARGED_CAPACITY_WMI_GUID syscall.GUID = syscall.GUID{0x40b40565, 0x96f7, 0x4435, 
	[8]byte{0x86, 0x94, 0x97, 0xe0, 0xe4, 0x39, 0x59, 0x05}}
	BATTERY_CYCLE_COUNT_WMI_GUID syscall.GUID = syscall.GUID{0xef98db24, 0x0014, 0x4c25, 
	[8]byte{0xa5, 0x0b, 0xc7, 0x24, 0xae, 0x5c, 0xd3, 0x71}}
	BATTERY_STATIC_DATA_WMI_GUID syscall.GUID = syscall.GUID{0x05e1e463, 0xe4e2, 0x4ea9, 
	[8]byte{0x80, 0xcb, 0x9b, 0xd4, 0xb3, 0xca, 0x06, 0x55}}
	BATTERY_STATUS_CHANGE_WMI_GUID syscall.GUID = syscall.GUID{0xcddfa0c3, 0x7c5b, 0x4e43, 
	[8]byte{0xa0, 0x34, 0x05, 0x9f, 0xa5, 0xb8, 0x43, 0x64}}
	BATTERY_TAG_CHANGE_WMI_GUID syscall.GUID = syscall.GUID{0x5e1f6e19, 0x8786, 0x4d23, 
	[8]byte{0x94, 0xfc, 0x9e, 0x74, 0x6b, 0xd5, 0xd8, 0x88}}
	GUID_DEVICE_ENERGY_METER syscall.GUID = syscall.GUID{0x45bd8344, 0x7ed6, 0x49cf, 
	[8]byte{0xa4, 0x40, 0xc2, 0x76, 0xc9, 0x33, 0xb0, 0x53}}
)

// enums

// enum POWER_PLATFORM_ROLE_VERSION
type POWER_PLATFORM_ROLE_VERSION uint32
const (
	POWER_PLATFORM_ROLE_V1 POWER_PLATFORM_ROLE_VERSION = 1
	POWER_PLATFORM_ROLE_V2 POWER_PLATFORM_ROLE_VERSION = 2
)

// enum POWER_SETTING_REGISTER_NOTIFICATION_FLAGS
type POWER_SETTING_REGISTER_NOTIFICATION_FLAGS uint32
const (
	DEVICE_NOTIFY_SERVICE_HANDLE POWER_SETTING_REGISTER_NOTIFICATION_FLAGS = 1
	DEVICE_NOTIFY_CALLBACK POWER_SETTING_REGISTER_NOTIFICATION_FLAGS = 2
	DEVICE_NOTIFY_WINDOW_HANDLE POWER_SETTING_REGISTER_NOTIFICATION_FLAGS = 0
)

// enum EXECUTION_STATE
// flags
type EXECUTION_STATE uint32
const (
	ES_AWAYMODE_REQUIRED EXECUTION_STATE = 64
	ES_CONTINUOUS EXECUTION_STATE = 2147483648
	ES_DISPLAY_REQUIRED EXECUTION_STATE = 2
	ES_SYSTEM_REQUIRED EXECUTION_STATE = 1
	ES_USER_PRESENT EXECUTION_STATE = 4
)

// enum POWER_ACTION_POLICY_EVENT_CODE
// flags
type POWER_ACTION_POLICY_EVENT_CODE uint32
const (
	POWER_FORCE_TRIGGER_RESET POWER_ACTION_POLICY_EVENT_CODE = 2147483648
	POWER_LEVEL_USER_NOTIFY_EXEC POWER_ACTION_POLICY_EVENT_CODE = 4
	POWER_LEVEL_USER_NOTIFY_SOUND POWER_ACTION_POLICY_EVENT_CODE = 2
	POWER_LEVEL_USER_NOTIFY_TEXT POWER_ACTION_POLICY_EVENT_CODE = 1
	POWER_USER_NOTIFY_BUTTON POWER_ACTION_POLICY_EVENT_CODE = 8
	POWER_USER_NOTIFY_SHUTDOWN POWER_ACTION_POLICY_EVENT_CODE = 16
)

// enum EFFECTIVE_POWER_MODE
type EFFECTIVE_POWER_MODE int32
const (
	EffectivePowerModeBatterySaver EFFECTIVE_POWER_MODE = 0
	EffectivePowerModeBetterBattery EFFECTIVE_POWER_MODE = 1
	EffectivePowerModeBalanced EFFECTIVE_POWER_MODE = 2
	EffectivePowerModeHighPerformance EFFECTIVE_POWER_MODE = 3
	EffectivePowerModeMaxPerformance EFFECTIVE_POWER_MODE = 4
	EffectivePowerModeGameMode EFFECTIVE_POWER_MODE = 5
	EffectivePowerModeMixedReality EFFECTIVE_POWER_MODE = 6
)

// enum POWER_DATA_ACCESSOR
type POWER_DATA_ACCESSOR int32
const (
	ACCESS_AC_POWER_SETTING_INDEX POWER_DATA_ACCESSOR = 0
	ACCESS_DC_POWER_SETTING_INDEX POWER_DATA_ACCESSOR = 1
	ACCESS_FRIENDLY_NAME POWER_DATA_ACCESSOR = 2
	ACCESS_DESCRIPTION POWER_DATA_ACCESSOR = 3
	ACCESS_POSSIBLE_POWER_SETTING POWER_DATA_ACCESSOR = 4
	ACCESS_POSSIBLE_POWER_SETTING_FRIENDLY_NAME POWER_DATA_ACCESSOR = 5
	ACCESS_POSSIBLE_POWER_SETTING_DESCRIPTION POWER_DATA_ACCESSOR = 6
	ACCESS_DEFAULT_AC_POWER_SETTING POWER_DATA_ACCESSOR = 7
	ACCESS_DEFAULT_DC_POWER_SETTING POWER_DATA_ACCESSOR = 8
	ACCESS_POSSIBLE_VALUE_MIN POWER_DATA_ACCESSOR = 9
	ACCESS_POSSIBLE_VALUE_MAX POWER_DATA_ACCESSOR = 10
	ACCESS_POSSIBLE_VALUE_INCREMENT POWER_DATA_ACCESSOR = 11
	ACCESS_POSSIBLE_VALUE_UNITS POWER_DATA_ACCESSOR = 12
	ACCESS_ICON_RESOURCE POWER_DATA_ACCESSOR = 13
	ACCESS_DEFAULT_SECURITY_DESCRIPTOR POWER_DATA_ACCESSOR = 14
	ACCESS_ATTRIBUTES POWER_DATA_ACCESSOR = 15
	ACCESS_SCHEME POWER_DATA_ACCESSOR = 16
	ACCESS_SUBGROUP POWER_DATA_ACCESSOR = 17
	ACCESS_INDIVIDUAL_SETTING POWER_DATA_ACCESSOR = 18
	ACCESS_ACTIVE_SCHEME POWER_DATA_ACCESSOR = 19
	ACCESS_CREATE_SCHEME POWER_DATA_ACCESSOR = 20
	ACCESS_AC_POWER_SETTING_MAX POWER_DATA_ACCESSOR = 21
	ACCESS_DC_POWER_SETTING_MAX POWER_DATA_ACCESSOR = 22
	ACCESS_AC_POWER_SETTING_MIN POWER_DATA_ACCESSOR = 23
	ACCESS_DC_POWER_SETTING_MIN POWER_DATA_ACCESSOR = 24
	ACCESS_PROFILE POWER_DATA_ACCESSOR = 25
	ACCESS_OVERLAY_SCHEME POWER_DATA_ACCESSOR = 26
	ACCESS_ACTIVE_OVERLAY_SCHEME POWER_DATA_ACCESSOR = 27
)

// enum BATTERY_QUERY_INFORMATION_LEVEL
type BATTERY_QUERY_INFORMATION_LEVEL int32
const (
	BatteryInformation BATTERY_QUERY_INFORMATION_LEVEL = 0
	BatteryGranularityInformation BATTERY_QUERY_INFORMATION_LEVEL = 1
	BatteryTemperature BATTERY_QUERY_INFORMATION_LEVEL = 2
	BatteryEstimatedTime BATTERY_QUERY_INFORMATION_LEVEL = 3
	BatteryDeviceName BATTERY_QUERY_INFORMATION_LEVEL = 4
	BatteryManufactureDate BATTERY_QUERY_INFORMATION_LEVEL = 5
	BatteryManufactureName BATTERY_QUERY_INFORMATION_LEVEL = 6
	BatteryUniqueID BATTERY_QUERY_INFORMATION_LEVEL = 7
	BatterySerialNumber BATTERY_QUERY_INFORMATION_LEVEL = 8
)

// enum BATTERY_CHARGING_SOURCE_TYPE
type BATTERY_CHARGING_SOURCE_TYPE int32
const (
	BatteryChargingSourceType_AC BATTERY_CHARGING_SOURCE_TYPE = 1
	BatteryChargingSourceType_USB BATTERY_CHARGING_SOURCE_TYPE = 2
	BatteryChargingSourceType_Wireless BATTERY_CHARGING_SOURCE_TYPE = 3
	BatteryChargingSourceType_Max BATTERY_CHARGING_SOURCE_TYPE = 4
)

// enum USB_CHARGER_PORT
type USB_CHARGER_PORT int32
const (
	UsbChargerPort_Legacy USB_CHARGER_PORT = 0
	UsbChargerPort_TypeC USB_CHARGER_PORT = 1
	UsbChargerPort_Max USB_CHARGER_PORT = 2
)

// enum BATTERY_SET_INFORMATION_LEVEL
type BATTERY_SET_INFORMATION_LEVEL int32
const (
	BatteryCriticalBias BATTERY_SET_INFORMATION_LEVEL = 0
	BatteryCharge BATTERY_SET_INFORMATION_LEVEL = 1
	BatteryDischarge BATTERY_SET_INFORMATION_LEVEL = 2
	BatteryChargingSource BATTERY_SET_INFORMATION_LEVEL = 3
	BatteryChargerId BATTERY_SET_INFORMATION_LEVEL = 4
	BatteryChargerStatus BATTERY_SET_INFORMATION_LEVEL = 5
)

// enum EMI_MEASUREMENT_UNIT
type EMI_MEASUREMENT_UNIT int32
const (
	EmiMeasurementUnitPicowattHours EMI_MEASUREMENT_UNIT = 0
)

// enum SYSTEM_POWER_STATE
type SYSTEM_POWER_STATE int32
const (
	PowerSystemUnspecified SYSTEM_POWER_STATE = 0
	PowerSystemWorking SYSTEM_POWER_STATE = 1
	PowerSystemSleeping1 SYSTEM_POWER_STATE = 2
	PowerSystemSleeping2 SYSTEM_POWER_STATE = 3
	PowerSystemSleeping3 SYSTEM_POWER_STATE = 4
	PowerSystemHibernate SYSTEM_POWER_STATE = 5
	PowerSystemShutdown SYSTEM_POWER_STATE = 6
	PowerSystemMaximum SYSTEM_POWER_STATE = 7
)

// enum POWER_ACTION
type POWER_ACTION int32
const (
	PowerActionNone POWER_ACTION = 0
	PowerActionReserved POWER_ACTION = 1
	PowerActionSleep POWER_ACTION = 2
	PowerActionHibernate POWER_ACTION = 3
	PowerActionShutdown POWER_ACTION = 4
	PowerActionShutdownReset POWER_ACTION = 5
	PowerActionShutdownOff POWER_ACTION = 6
	PowerActionWarmEject POWER_ACTION = 7
	PowerActionDisplayOff POWER_ACTION = 8
)

// enum DEVICE_POWER_STATE
type DEVICE_POWER_STATE int32
const (
	PowerDeviceUnspecified DEVICE_POWER_STATE = 0
	PowerDeviceD0 DEVICE_POWER_STATE = 1
	PowerDeviceD1 DEVICE_POWER_STATE = 2
	PowerDeviceD2 DEVICE_POWER_STATE = 3
	PowerDeviceD3 DEVICE_POWER_STATE = 4
	PowerDeviceMaximum DEVICE_POWER_STATE = 5
)

// enum LATENCY_TIME
type LATENCY_TIME int32
const (
	LT_DONT_CARE LATENCY_TIME = 0
	LT_LOWEST_LATENCY LATENCY_TIME = 1
)

// enum POWER_REQUEST_TYPE
type POWER_REQUEST_TYPE int32
const (
	PowerRequestDisplayRequired POWER_REQUEST_TYPE = 0
	PowerRequestSystemRequired POWER_REQUEST_TYPE = 1
	PowerRequestAwayModeRequired POWER_REQUEST_TYPE = 2
	PowerRequestExecutionRequired POWER_REQUEST_TYPE = 3
)

// enum POWER_INFORMATION_LEVEL
type POWER_INFORMATION_LEVEL int32
const (
	SystemPowerPolicyAc POWER_INFORMATION_LEVEL = 0
	SystemPowerPolicyDc POWER_INFORMATION_LEVEL = 1
	VerifySystemPolicyAc POWER_INFORMATION_LEVEL = 2
	VerifySystemPolicyDc POWER_INFORMATION_LEVEL = 3
	SystemPowerCapabilities POWER_INFORMATION_LEVEL = 4
	SystemBatteryState POWER_INFORMATION_LEVEL = 5
	SystemPowerStateHandler POWER_INFORMATION_LEVEL = 6
	ProcessorStateHandler POWER_INFORMATION_LEVEL = 7
	SystemPowerPolicyCurrent POWER_INFORMATION_LEVEL = 8
	AdministratorPowerPolicy POWER_INFORMATION_LEVEL = 9
	SystemReserveHiberFile POWER_INFORMATION_LEVEL = 10
	ProcessorInformation POWER_INFORMATION_LEVEL = 11
	SystemPowerInformation POWER_INFORMATION_LEVEL = 12
	ProcessorStateHandler2 POWER_INFORMATION_LEVEL = 13
	LastWakeTime POWER_INFORMATION_LEVEL = 14
	LastSleepTime POWER_INFORMATION_LEVEL = 15
	SystemExecutionState POWER_INFORMATION_LEVEL = 16
	SystemPowerStateNotifyHandler POWER_INFORMATION_LEVEL = 17
	ProcessorPowerPolicyAc POWER_INFORMATION_LEVEL = 18
	ProcessorPowerPolicyDc POWER_INFORMATION_LEVEL = 19
	VerifyProcessorPowerPolicyAc POWER_INFORMATION_LEVEL = 20
	VerifyProcessorPowerPolicyDc POWER_INFORMATION_LEVEL = 21
	ProcessorPowerPolicyCurrent POWER_INFORMATION_LEVEL = 22
	SystemPowerStateLogging POWER_INFORMATION_LEVEL = 23
	SystemPowerLoggingEntry POWER_INFORMATION_LEVEL = 24
	SetPowerSettingValue POWER_INFORMATION_LEVEL = 25
	NotifyUserPowerSetting POWER_INFORMATION_LEVEL = 26
	PowerInformationLevelUnused0 POWER_INFORMATION_LEVEL = 27
	SystemMonitorHiberBootPowerOff POWER_INFORMATION_LEVEL = 28
	SystemVideoState POWER_INFORMATION_LEVEL = 29
	TraceApplicationPowerMessage POWER_INFORMATION_LEVEL = 30
	TraceApplicationPowerMessageEnd POWER_INFORMATION_LEVEL = 31
	ProcessorPerfStates POWER_INFORMATION_LEVEL = 32
	ProcessorIdleStates POWER_INFORMATION_LEVEL = 33
	ProcessorCap POWER_INFORMATION_LEVEL = 34
	SystemWakeSource POWER_INFORMATION_LEVEL = 35
	SystemHiberFileInformation POWER_INFORMATION_LEVEL = 36
	TraceServicePowerMessage POWER_INFORMATION_LEVEL = 37
	ProcessorLoad POWER_INFORMATION_LEVEL = 38
	PowerShutdownNotification POWER_INFORMATION_LEVEL = 39
	MonitorCapabilities POWER_INFORMATION_LEVEL = 40
	SessionPowerInit POWER_INFORMATION_LEVEL = 41
	SessionDisplayState POWER_INFORMATION_LEVEL = 42
	PowerRequestCreate POWER_INFORMATION_LEVEL = 43
	PowerRequestAction POWER_INFORMATION_LEVEL = 44
	GetPowerRequestList POWER_INFORMATION_LEVEL = 45
	ProcessorInformationEx POWER_INFORMATION_LEVEL = 46
	NotifyUserModeLegacyPowerEvent POWER_INFORMATION_LEVEL = 47
	GroupPark POWER_INFORMATION_LEVEL = 48
	ProcessorIdleDomains POWER_INFORMATION_LEVEL = 49
	WakeTimerList POWER_INFORMATION_LEVEL = 50
	SystemHiberFileSize POWER_INFORMATION_LEVEL = 51
	ProcessorIdleStatesHv POWER_INFORMATION_LEVEL = 52
	ProcessorPerfStatesHv POWER_INFORMATION_LEVEL = 53
	ProcessorPerfCapHv POWER_INFORMATION_LEVEL = 54
	ProcessorSetIdle POWER_INFORMATION_LEVEL = 55
	LogicalProcessorIdling POWER_INFORMATION_LEVEL = 56
	UserPresence POWER_INFORMATION_LEVEL = 57
	PowerSettingNotificationName POWER_INFORMATION_LEVEL = 58
	GetPowerSettingValue POWER_INFORMATION_LEVEL = 59
	IdleResiliency POWER_INFORMATION_LEVEL = 60
	SessionRITState POWER_INFORMATION_LEVEL = 61
	SessionConnectNotification POWER_INFORMATION_LEVEL = 62
	SessionPowerCleanup POWER_INFORMATION_LEVEL = 63
	SessionLockState POWER_INFORMATION_LEVEL = 64
	SystemHiberbootState POWER_INFORMATION_LEVEL = 65
	PlatformInformation POWER_INFORMATION_LEVEL = 66
	PdcInvocation POWER_INFORMATION_LEVEL = 67
	MonitorInvocation POWER_INFORMATION_LEVEL = 68
	FirmwareTableInformationRegistered POWER_INFORMATION_LEVEL = 69
	SetShutdownSelectedTime POWER_INFORMATION_LEVEL = 70
	SuspendResumeInvocation POWER_INFORMATION_LEVEL = 71
	PlmPowerRequestCreate POWER_INFORMATION_LEVEL = 72
	ScreenOff POWER_INFORMATION_LEVEL = 73
	CsDeviceNotification POWER_INFORMATION_LEVEL = 74
	PlatformRole POWER_INFORMATION_LEVEL = 75
	LastResumePerformance POWER_INFORMATION_LEVEL = 76
	DisplayBurst POWER_INFORMATION_LEVEL = 77
	ExitLatencySamplingPercentage POWER_INFORMATION_LEVEL = 78
	RegisterSpmPowerSettings POWER_INFORMATION_LEVEL = 79
	PlatformIdleStates POWER_INFORMATION_LEVEL = 80
	ProcessorIdleVeto POWER_INFORMATION_LEVEL = 81
	PlatformIdleVeto POWER_INFORMATION_LEVEL = 82
	SystemBatteryStatePrecise POWER_INFORMATION_LEVEL = 83
	ThermalEvent POWER_INFORMATION_LEVEL = 84
	PowerRequestActionInternal POWER_INFORMATION_LEVEL = 85
	BatteryDeviceState POWER_INFORMATION_LEVEL = 86
	PowerInformationInternal POWER_INFORMATION_LEVEL = 87
	ThermalStandby POWER_INFORMATION_LEVEL = 88
	SystemHiberFileType POWER_INFORMATION_LEVEL = 89
	PhysicalPowerButtonPress POWER_INFORMATION_LEVEL = 90
	QueryPotentialDripsConstraint POWER_INFORMATION_LEVEL = 91
	EnergyTrackerCreate POWER_INFORMATION_LEVEL = 92
	EnergyTrackerQuery POWER_INFORMATION_LEVEL = 93
	UpdateBlackBoxRecorder POWER_INFORMATION_LEVEL = 94
	SessionAllowExternalDmaDevices POWER_INFORMATION_LEVEL = 95
	SendSuspendResumeNotification POWER_INFORMATION_LEVEL = 96
	PowerInformationLevelMaximum POWER_INFORMATION_LEVEL = 97
)

// enum SYSTEM_POWER_CONDITION
type SYSTEM_POWER_CONDITION int32
const (
	PoAc SYSTEM_POWER_CONDITION = 0
	PoDc SYSTEM_POWER_CONDITION = 1
	PoHot SYSTEM_POWER_CONDITION = 2
	PoConditionMaximum SYSTEM_POWER_CONDITION = 3
)

// enum POWER_PLATFORM_ROLE
type POWER_PLATFORM_ROLE int32
const (
	PlatformRoleUnspecified POWER_PLATFORM_ROLE = 0
	PlatformRoleDesktop POWER_PLATFORM_ROLE = 1
	PlatformRoleMobile POWER_PLATFORM_ROLE = 2
	PlatformRoleWorkstation POWER_PLATFORM_ROLE = 3
	PlatformRoleEnterpriseServer POWER_PLATFORM_ROLE = 4
	PlatformRoleSOHOServer POWER_PLATFORM_ROLE = 5
	PlatformRoleAppliancePC POWER_PLATFORM_ROLE = 6
	PlatformRolePerformanceServer POWER_PLATFORM_ROLE = 7
	PlatformRoleSlate POWER_PLATFORM_ROLE = 8
	PlatformRoleMaximum POWER_PLATFORM_ROLE = 9
)


// structs

type GLOBAL_MACHINE_POWER_POLICY struct {
	Revision uint32
	LidOpenWakeAc SYSTEM_POWER_STATE
	LidOpenWakeDc SYSTEM_POWER_STATE
	BroadcastCapacityResolution uint32
}

type GLOBAL_USER_POWER_POLICY struct {
	Revision uint32
	PowerButtonAc POWER_ACTION_POLICY
	PowerButtonDc POWER_ACTION_POLICY
	SleepButtonAc POWER_ACTION_POLICY
	SleepButtonDc POWER_ACTION_POLICY
	LidCloseAc POWER_ACTION_POLICY
	LidCloseDc POWER_ACTION_POLICY
	DischargePolicy [4]SYSTEM_POWER_LEVEL
	GlobalFlags uint32
}

type GLOBAL_POWER_POLICY struct {
	User GLOBAL_USER_POWER_POLICY
	Mach GLOBAL_MACHINE_POWER_POLICY
}

type MACHINE_POWER_POLICY struct {
	Revision uint32
	MinSleepAc SYSTEM_POWER_STATE
	MinSleepDc SYSTEM_POWER_STATE
	ReducedLatencySleepAc SYSTEM_POWER_STATE
	ReducedLatencySleepDc SYSTEM_POWER_STATE
	DozeTimeoutAc uint32
	DozeTimeoutDc uint32
	DozeS4TimeoutAc uint32
	DozeS4TimeoutDc uint32
	MinThrottleAc uint8
	MinThrottleDc uint8
	Pad1 [2]uint8
	OverThrottledAc POWER_ACTION_POLICY
	OverThrottledDc POWER_ACTION_POLICY
}

type MACHINE_PROCESSOR_POWER_POLICY struct {
	Revision uint32
	ProcessorPolicyAc PROCESSOR_POWER_POLICY
	ProcessorPolicyDc PROCESSOR_POWER_POLICY
}

type USER_POWER_POLICY struct {
	Revision uint32
	IdleAc POWER_ACTION_POLICY
	IdleDc POWER_ACTION_POLICY
	IdleTimeoutAc uint32
	IdleTimeoutDc uint32
	IdleSensitivityAc uint8
	IdleSensitivityDc uint8
	ThrottlePolicyAc uint8
	ThrottlePolicyDc uint8
	MaxSleepAc SYSTEM_POWER_STATE
	MaxSleepDc SYSTEM_POWER_STATE
	Reserved [2]uint32
	VideoTimeoutAc uint32
	VideoTimeoutDc uint32
	SpindownTimeoutAc uint32
	SpindownTimeoutDc uint32
	OptimizeForPowerAc BOOLEAN
	OptimizeForPowerDc BOOLEAN
	FanThrottleToleranceAc uint8
	FanThrottleToleranceDc uint8
	ForcedThrottleAc uint8
	ForcedThrottleDc uint8
}

type POWER_POLICY struct {
	User USER_POWER_POLICY
	Mach MACHINE_POWER_POLICY
}

type DEVICE_NOTIFY_SUBSCRIBE_PARAMETERS struct {
	Callback uintptr
	Context unsafe.Pointer
}

type THERMAL_EVENT struct {
	Version uint32
	Size uint32
	Type uint32
	Temperature uint32
	TripPointTemperature uint32
	Initiator PWSTR
}

type BATTERY_QUERY_INFORMATION struct {
	BatteryTag uint32
	InformationLevel BATTERY_QUERY_INFORMATION_LEVEL
	AtRate uint32
}

type BATTERY_INFORMATION struct {
	Capabilities uint32
	Technology uint8
	Reserved [3]uint8
	Chemistry [4]uint8
	DesignedCapacity uint32
	FullChargedCapacity uint32
	DefaultAlert1 uint32
	DefaultAlert2 uint32
	CriticalBias uint32
	CycleCount uint32
}

type BATTERY_CHARGING_SOURCE struct {
	Type BATTERY_CHARGING_SOURCE_TYPE
	MaxCurrent uint32
}

type BATTERY_CHARGING_SOURCE_INFORMATION struct {
	Type BATTERY_CHARGING_SOURCE_TYPE
	SourceOnline BOOLEAN
}

type BATTERY_SET_INFORMATION struct {
	BatteryTag uint32
	InformationLevel BATTERY_SET_INFORMATION_LEVEL
	Buffer [1]uint8
}

type BATTERY_CHARGER_STATUS struct {
	Type BATTERY_CHARGING_SOURCE_TYPE
	VaData [1]uint32
}

type BATTERY_USB_CHARGER_STATUS struct {
	Type BATTERY_CHARGING_SOURCE_TYPE
	Reserved uint32
	Flags uint32
	MaxCurrent uint32
	Voltage uint32
	PortType USB_CHARGER_PORT
	PortId uint64
	PowerSourceInformation unsafe.Pointer
	OemCharger syscall.GUID
}

type BATTERY_WAIT_STATUS struct {
	BatteryTag uint32
	Timeout uint32
	PowerState uint32
	LowCapacity uint32
	HighCapacity uint32
}

type BATTERY_STATUS struct {
	PowerState uint32
	Capacity uint32
	Voltage uint32
	Rate int32
}

type BATTERY_MANUFACTURE_DATE struct {
	Day uint8
	Month uint8
	Year uint16
}

type THERMAL_INFORMATION struct {
	ThermalStamp uint32
	ThermalConstant1 uint32
	ThermalConstant2 uint32
	Processors uintptr
	SamplingPeriod uint32
	CurrentTemperature uint32
	PassiveTripPoint uint32
	CriticalTripPoint uint32
	ActiveTripPointCount uint8
	ActiveTripPoint [10]uint32
}

type THERMAL_WAIT_READ struct {
	Timeout uint32
	LowTemperature uint32
	HighTemperature uint32
}

type THERMAL_POLICY struct {
	Version uint32
	WaitForUpdate BOOLEAN
	Hibernate BOOLEAN
	Critical BOOLEAN
	ThermalStandby BOOLEAN
	ActivationReasons uint32
	PassiveLimit uint32
	ActiveLevel uint32
	OverThrottled BOOLEAN
}

type PROCESSOR_OBJECT_INFO struct {
	PhysicalID uint32
	PBlkAddress uint32
	PBlkLength uint8
}

type PROCESSOR_OBJECT_INFO_EX struct {
	PhysicalID uint32
	PBlkAddress uint32
	PBlkLength uint8
	InitialApicId uint32
}

type WAKE_ALARM_INFORMATION struct {
	TimerIdentifier uint32
	Timeout uint32
}

type ACPI_REAL_TIME struct {
	Year uint16
	Month uint8
	Day uint8
	Hour uint8
	Minute uint8
	Second uint8
	Valid uint8
	Milliseconds uint16
	TimeZone int16
	DayLight uint8
	Reserved1 [3]uint8
}

type EMI_VERSION struct {
	EmiVersion uint16
}

type EMI_METADATA_SIZE struct {
	MetadataSize uint32
}

type EMI_CHANNEL_MEASUREMENT_DATA struct {
	AbsoluteEnergy uint64
	AbsoluteTime uint64
}

type EMI_METADATA_V1 struct {
	MeasurementUnit EMI_MEASUREMENT_UNIT
	HardwareOEM [16]uint16
	HardwareModel [16]uint16
	HardwareRevision uint16
	MeteredHardwareNameSize uint16
	MeteredHardwareName [1]uint16
}

type EMI_CHANNEL_V2 struct {
	MeasurementUnit EMI_MEASUREMENT_UNIT
	ChannelNameSize uint16
	ChannelName [1]uint16
}

type EMI_METADATA_V2 struct {
	HardwareOEM [16]uint16
	HardwareModel [16]uint16
	HardwareRevision uint16
	ChannelCount uint16
	Channels [1]EMI_CHANNEL_V2
}

type EMI_MEASUREMENT_DATA_V2 struct {
	ChannelData [1]EMI_CHANNEL_MEASUREMENT_DATA
}

type CM_POWER_DATA struct {
	PD_Size uint32
	PD_MostRecentPowerState DEVICE_POWER_STATE
	PD_Capabilities uint32
	PD_D1Latency uint32
	PD_D2Latency uint32
	PD_D3Latency uint32
	PD_PowerStateMapping [7]DEVICE_POWER_STATE
	PD_DeepestSystemWake SYSTEM_POWER_STATE
}

type SET_POWER_SETTING_VALUE struct {
	Version uint32
	Guid syscall.GUID
	PowerCondition SYSTEM_POWER_CONDITION
	DataLength uint32
	Data [1]uint8
}

type BATTERY_REPORTING_SCALE struct {
	Granularity uint32
	Capacity uint32
}

type POWER_ACTION_POLICY struct {
	Action POWER_ACTION
	Flags uint32
	EventCode POWER_ACTION_POLICY_EVENT_CODE
}

type SYSTEM_POWER_LEVEL struct {
	Enable BOOLEAN
	Spare [3]uint8
	BatteryLevel uint32
	PowerPolicy POWER_ACTION_POLICY
	MinSystemState SYSTEM_POWER_STATE
}

type SYSTEM_POWER_POLICY struct {
	Revision uint32
	PowerButton POWER_ACTION_POLICY
	SleepButton POWER_ACTION_POLICY
	LidClose POWER_ACTION_POLICY
	LidOpenWake SYSTEM_POWER_STATE
	Reserved uint32
	Idle POWER_ACTION_POLICY
	IdleTimeout uint32
	IdleSensitivity uint8
	DynamicThrottle uint8
	Spare2 [2]uint8
	MinSleep SYSTEM_POWER_STATE
	MaxSleep SYSTEM_POWER_STATE
	ReducedLatencySleep SYSTEM_POWER_STATE
	WinLogonFlags uint32
	Spare3 uint32
	DozeS4Timeout uint32
	BroadcastCapacityResolution uint32
	DischargePolicy [4]SYSTEM_POWER_LEVEL
	VideoTimeout uint32
	VideoDimDisplay BOOLEAN
	VideoReserved [3]uint32
	SpindownTimeout uint32
	OptimizeForPower BOOLEAN
	FanThrottleTolerance uint8
	ForcedThrottle uint8
	MinThrottle uint8
	OverThrottled POWER_ACTION_POLICY
}

type PROCESSOR_POWER_POLICY_INFO struct {
	TimeCheck uint32
	DemoteLimit uint32
	PromoteLimit uint32
	DemotePercent uint8
	PromotePercent uint8
	Spare [2]uint8
	Bitfield_ uint32
}

type PROCESSOR_POWER_POLICY struct {
	Revision uint32
	DynamicThrottle uint8
	Spare [3]uint8
	Bitfield_ uint32
	PolicyCount uint32
	Policy [3]PROCESSOR_POWER_POLICY_INFO
}

type ADMINISTRATOR_POWER_POLICY struct {
	MinSleep SYSTEM_POWER_STATE
	MaxSleep SYSTEM_POWER_STATE
	MinVideoTimeout uint32
	MaxVideoTimeout uint32
	MinSpindownTimeout uint32
	MaxSpindownTimeout uint32
}

type SYSTEM_POWER_CAPABILITIES struct {
	PowerButtonPresent BOOLEAN
	SleepButtonPresent BOOLEAN
	LidPresent BOOLEAN
	SystemS1 BOOLEAN
	SystemS2 BOOLEAN
	SystemS3 BOOLEAN
	SystemS4 BOOLEAN
	SystemS5 BOOLEAN
	HiberFilePresent BOOLEAN
	FullWake BOOLEAN
	VideoDimPresent BOOLEAN
	ApmPresent BOOLEAN
	UpsPresent BOOLEAN
	ThermalControl BOOLEAN
	ProcessorThrottle BOOLEAN
	ProcessorMinThrottle uint8
	ProcessorMaxThrottle uint8
	FastSystemS4 BOOLEAN
	Hiberboot BOOLEAN
	WakeAlarmPresent BOOLEAN
	AoAc BOOLEAN
	DiskSpinDown BOOLEAN
	HiberFileType uint8
	AoAcConnectivitySupported BOOLEAN
	Spare3 [6]uint8
	SystemBatteriesPresent BOOLEAN
	BatteriesAreShortTerm BOOLEAN
	BatteryScale [3]BATTERY_REPORTING_SCALE
	AcOnLineWake SYSTEM_POWER_STATE
	SoftLidWake SYSTEM_POWER_STATE
	RtcWake SYSTEM_POWER_STATE
	MinDeviceWakeState SYSTEM_POWER_STATE
	DefaultLowLatencyWake SYSTEM_POWER_STATE
}

type SYSTEM_BATTERY_STATE struct {
	AcOnLine BOOLEAN
	BatteryPresent BOOLEAN
	Charging BOOLEAN
	Discharging BOOLEAN
	Spare1 [3]BOOLEAN
	Tag uint8
	MaxCapacity uint32
	RemainingCapacity uint32
	Rate uint32
	EstimatedTime uint32
	DefaultAlert1 uint32
	DefaultAlert2 uint32
}

type POWERBROADCAST_SETTING struct {
	PowerSetting syscall.GUID
	DataLength uint32
	Data [1]uint8
}

type SYSTEM_POWER_STATUS struct {
	ACLineStatus uint8
	BatteryFlag uint8
	BatteryLifePercent uint8
	SystemStatusFlag uint8
	BatteryLifeTime uint32
	BatteryFullLifeTime uint32
}


// func types

type EFFECTIVE_POWER_MODE_CALLBACK func(Mode EFFECTIVE_POWER_MODE, Context unsafe.Pointer)

type PWRSCHEMESENUMPROC_V1 func(Index uint32, NameSize uint32, Name *int8, DescriptionSize uint32, Description *int8, Policy *POWER_POLICY, Context LPARAM) BOOLEAN

type PWRSCHEMESENUMPROC func(Index uint32, NameSize uint32, Name PWSTR, DescriptionSize uint32, Description PWSTR, Policy *POWER_POLICY, Context LPARAM) BOOLEAN

type PDEVICE_NOTIFY_CALLBACK_ROUTINE func(Context unsafe.Pointer, Type uint32, Setting unsafe.Pointer) uint32


var (
	pRegisterPowerSettingNotification uintptr
	pUnregisterPowerSettingNotification uintptr
	pRegisterSuspendResumeNotification uintptr
	pUnregisterSuspendResumeNotification uintptr
	pRequestWakeupLatency uintptr
	pIsSystemResumeAutomatic uintptr
	pSetThreadExecutionState uintptr
	pPowerCreateRequest uintptr
	pPowerSetRequest uintptr
	pPowerClearRequest uintptr
	pGetDevicePowerState uintptr
	pSetSystemPowerState uintptr
	pGetSystemPowerStatus uintptr
)

func RegisterPowerSettingNotification(hRecipient HANDLE, PowerSettingGuid *syscall.GUID, Flags uint32) (HPOWERNOTIFY, WIN32_ERROR) {
	addr := lazyAddr(&pRegisterPowerSettingNotification, libUser32, "RegisterPowerSettingNotification")
	ret, _,  err := syscall.SyscallN(addr, hRecipient, uintptr(unsafe.Pointer(PowerSettingGuid)), uintptr(Flags))
	return HPOWERNOTIFY(ret), WIN32_ERROR(err)
}

func UnregisterPowerSettingNotification(Handle HPOWERNOTIFY) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pUnregisterPowerSettingNotification, libUser32, "UnregisterPowerSettingNotification")
	ret, _,  err := syscall.SyscallN(addr, Handle)
	return BOOL(ret), WIN32_ERROR(err)
}

func RegisterSuspendResumeNotification(hRecipient HANDLE, Flags uint32) (HPOWERNOTIFY, WIN32_ERROR) {
	addr := lazyAddr(&pRegisterSuspendResumeNotification, libUser32, "RegisterSuspendResumeNotification")
	ret, _,  err := syscall.SyscallN(addr, hRecipient, uintptr(Flags))
	return HPOWERNOTIFY(ret), WIN32_ERROR(err)
}

func UnregisterSuspendResumeNotification(Handle HPOWERNOTIFY) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pUnregisterSuspendResumeNotification, libUser32, "UnregisterSuspendResumeNotification")
	ret, _,  err := syscall.SyscallN(addr, Handle)
	return BOOL(ret), WIN32_ERROR(err)
}

func RequestWakeupLatency(latency LATENCY_TIME) BOOL {
	addr := lazyAddr(&pRequestWakeupLatency, libKernel32, "RequestWakeupLatency")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(latency))
	return BOOL(ret)
}

func IsSystemResumeAutomatic() BOOL {
	addr := lazyAddr(&pIsSystemResumeAutomatic, libKernel32, "IsSystemResumeAutomatic")
	ret, _,  _ := syscall.SyscallN(addr)
	return BOOL(ret)
}

func SetThreadExecutionState(esFlags EXECUTION_STATE) EXECUTION_STATE {
	addr := lazyAddr(&pSetThreadExecutionState, libKernel32, "SetThreadExecutionState")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(esFlags))
	return EXECUTION_STATE(ret)
}

func PowerCreateRequest(Context *REASON_CONTEXT) (HANDLE, WIN32_ERROR) {
	addr := lazyAddr(&pPowerCreateRequest, libKernel32, "PowerCreateRequest")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Context)))
	return HANDLE(ret), WIN32_ERROR(err)
}

func PowerSetRequest(PowerRequest HANDLE, RequestType POWER_REQUEST_TYPE) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pPowerSetRequest, libKernel32, "PowerSetRequest")
	ret, _,  err := syscall.SyscallN(addr, PowerRequest, uintptr(RequestType))
	return BOOL(ret), WIN32_ERROR(err)
}

func PowerClearRequest(PowerRequest HANDLE, RequestType POWER_REQUEST_TYPE) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pPowerClearRequest, libKernel32, "PowerClearRequest")
	ret, _,  err := syscall.SyscallN(addr, PowerRequest, uintptr(RequestType))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetDevicePowerState(hDevice HANDLE, pfOn *BOOL) BOOL {
	addr := lazyAddr(&pGetDevicePowerState, libKernel32, "GetDevicePowerState")
	ret, _,  _ := syscall.SyscallN(addr, hDevice, uintptr(unsafe.Pointer(pfOn)))
	return BOOL(ret)
}

func SetSystemPowerState(fSuspend BOOL, fForce BOOL) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pSetSystemPowerState, libKernel32, "SetSystemPowerState")
	ret, _,  err := syscall.SyscallN(addr, uintptr(fSuspend), uintptr(fForce))
	return BOOL(ret), WIN32_ERROR(err)
}

func GetSystemPowerStatus(lpSystemPowerStatus *SYSTEM_POWER_STATUS) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pGetSystemPowerStatus, libKernel32, "GetSystemPowerStatus")
	ret, _,  err := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemPowerStatus)))
	return BOOL(ret), WIN32_ERROR(err)
}


