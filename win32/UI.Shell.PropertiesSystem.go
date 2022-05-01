package win32

import "unsafe"
import "syscall"

const (
	PKEY_PIDSTR_MAX uint32 = 10
)

// enums

// enum GETPROPERTYSTOREFLAGS
// flags
type GETPROPERTYSTOREFLAGS uint32
const (
	GPS_DEFAULT GETPROPERTYSTOREFLAGS = 0
	GPS_HANDLERPROPERTIESONLY GETPROPERTYSTOREFLAGS = 1
	GPS_READWRITE GETPROPERTYSTOREFLAGS = 2
	GPS_TEMPORARY GETPROPERTYSTOREFLAGS = 4
	GPS_FASTPROPERTIESONLY GETPROPERTYSTOREFLAGS = 8
	GPS_OPENSLOWITEM GETPROPERTYSTOREFLAGS = 16
	GPS_DELAYCREATION GETPROPERTYSTOREFLAGS = 32
	GPS_BESTEFFORT GETPROPERTYSTOREFLAGS = 64
	GPS_NO_OPLOCK GETPROPERTYSTOREFLAGS = 128
	GPS_PREFERQUERYPROPERTIES GETPROPERTYSTOREFLAGS = 256
	GPS_EXTRINSICPROPERTIES GETPROPERTYSTOREFLAGS = 512
	GPS_EXTRINSICPROPERTIESONLY GETPROPERTYSTOREFLAGS = 1024
	GPS_VOLATILEPROPERTIES GETPROPERTYSTOREFLAGS = 2048
	GPS_VOLATILEPROPERTIESONLY GETPROPERTYSTOREFLAGS = 4096
	GPS_MASK_VALID GETPROPERTYSTOREFLAGS = 8191
)

// enum PKA_FLAGS
// flags
type PKA_FLAGS uint32
const (
	PKA_SET PKA_FLAGS = 0
	PKA_APPEND PKA_FLAGS = 1
	PKA_DELETE PKA_FLAGS = 2
)

// enum PSC_STATE
type PSC_STATE int32
const (
	PSC_NORMAL PSC_STATE = 0
	PSC_NOTINSOURCE PSC_STATE = 1
	PSC_DIRTY PSC_STATE = 2
	PSC_READONLY PSC_STATE = 3
)

// enum PROPENUMTYPE
type PROPENUMTYPE int32
const (
	PET_DISCRETEVALUE PROPENUMTYPE = 0
	PET_RANGEDVALUE PROPENUMTYPE = 1
	PET_DEFAULTVALUE PROPENUMTYPE = 2
	PET_ENDRANGE PROPENUMTYPE = 3
)

// enum PROPDESC_TYPE_FLAGS
// flags
type PROPDESC_TYPE_FLAGS uint32
const (
	PDTF_DEFAULT PROPDESC_TYPE_FLAGS = 0
	PDTF_MULTIPLEVALUES PROPDESC_TYPE_FLAGS = 1
	PDTF_ISINNATE PROPDESC_TYPE_FLAGS = 2
	PDTF_ISGROUP PROPDESC_TYPE_FLAGS = 4
	PDTF_CANGROUPBY PROPDESC_TYPE_FLAGS = 8
	PDTF_CANSTACKBY PROPDESC_TYPE_FLAGS = 16
	PDTF_ISTREEPROPERTY PROPDESC_TYPE_FLAGS = 32
	PDTF_INCLUDEINFULLTEXTQUERY PROPDESC_TYPE_FLAGS = 64
	PDTF_ISVIEWABLE PROPDESC_TYPE_FLAGS = 128
	PDTF_ISQUERYABLE PROPDESC_TYPE_FLAGS = 256
	PDTF_CANBEPURGED PROPDESC_TYPE_FLAGS = 512
	PDTF_SEARCHRAWVALUE PROPDESC_TYPE_FLAGS = 1024
	PDTF_DONTCOERCEEMPTYSTRINGS PROPDESC_TYPE_FLAGS = 2048
	PDTF_ALWAYSINSUPPLEMENTALSTORE PROPDESC_TYPE_FLAGS = 4096
	PDTF_ISSYSTEMPROPERTY PROPDESC_TYPE_FLAGS = 2147483648
	PDTF_MASK_ALL PROPDESC_TYPE_FLAGS = 2147491839
)

// enum PROPDESC_VIEW_FLAGS
// flags
type PROPDESC_VIEW_FLAGS uint32
const (
	PDVF_DEFAULT PROPDESC_VIEW_FLAGS = 0
	PDVF_CENTERALIGN PROPDESC_VIEW_FLAGS = 1
	PDVF_RIGHTALIGN PROPDESC_VIEW_FLAGS = 2
	PDVF_BEGINNEWGROUP PROPDESC_VIEW_FLAGS = 4
	PDVF_FILLAREA PROPDESC_VIEW_FLAGS = 8
	PDVF_SORTDESCENDING PROPDESC_VIEW_FLAGS = 16
	PDVF_SHOWONLYIFPRESENT PROPDESC_VIEW_FLAGS = 32
	PDVF_SHOWBYDEFAULT PROPDESC_VIEW_FLAGS = 64
	PDVF_SHOWINPRIMARYLIST PROPDESC_VIEW_FLAGS = 128
	PDVF_SHOWINSECONDARYLIST PROPDESC_VIEW_FLAGS = 256
	PDVF_HIDELABEL PROPDESC_VIEW_FLAGS = 512
	PDVF_HIDDEN PROPDESC_VIEW_FLAGS = 2048
	PDVF_CANWRAP PROPDESC_VIEW_FLAGS = 4096
	PDVF_MASK_ALL PROPDESC_VIEW_FLAGS = 7167
)

// enum PROPDESC_DISPLAYTYPE
type PROPDESC_DISPLAYTYPE int32
const (
	PDDT_STRING PROPDESC_DISPLAYTYPE = 0
	PDDT_NUMBER PROPDESC_DISPLAYTYPE = 1
	PDDT_BOOLEAN PROPDESC_DISPLAYTYPE = 2
	PDDT_DATETIME PROPDESC_DISPLAYTYPE = 3
	PDDT_ENUMERATED PROPDESC_DISPLAYTYPE = 4
)

// enum PROPDESC_GROUPING_RANGE
type PROPDESC_GROUPING_RANGE int32
const (
	PDGR_DISCRETE PROPDESC_GROUPING_RANGE = 0
	PDGR_ALPHANUMERIC PROPDESC_GROUPING_RANGE = 1
	PDGR_SIZE PROPDESC_GROUPING_RANGE = 2
	PDGR_DYNAMIC PROPDESC_GROUPING_RANGE = 3
	PDGR_DATE PROPDESC_GROUPING_RANGE = 4
	PDGR_PERCENT PROPDESC_GROUPING_RANGE = 5
	PDGR_ENUMERATED PROPDESC_GROUPING_RANGE = 6
)

// enum PROPDESC_FORMAT_FLAGS
// flags
type PROPDESC_FORMAT_FLAGS uint32
const (
	PDFF_DEFAULT PROPDESC_FORMAT_FLAGS = 0
	PDFF_PREFIXNAME PROPDESC_FORMAT_FLAGS = 1
	PDFF_FILENAME PROPDESC_FORMAT_FLAGS = 2
	PDFF_ALWAYSKB PROPDESC_FORMAT_FLAGS = 4
	PDFF_RESERVED_RIGHTTOLEFT PROPDESC_FORMAT_FLAGS = 8
	PDFF_SHORTTIME PROPDESC_FORMAT_FLAGS = 16
	PDFF_LONGTIME PROPDESC_FORMAT_FLAGS = 32
	PDFF_HIDETIME PROPDESC_FORMAT_FLAGS = 64
	PDFF_SHORTDATE PROPDESC_FORMAT_FLAGS = 128
	PDFF_LONGDATE PROPDESC_FORMAT_FLAGS = 256
	PDFF_HIDEDATE PROPDESC_FORMAT_FLAGS = 512
	PDFF_RELATIVEDATE PROPDESC_FORMAT_FLAGS = 1024
	PDFF_USEEDITINVITATION PROPDESC_FORMAT_FLAGS = 2048
	PDFF_READONLY PROPDESC_FORMAT_FLAGS = 4096
	PDFF_NOAUTOREADINGORDER PROPDESC_FORMAT_FLAGS = 8192
)

// enum PROPDESC_SORTDESCRIPTION
type PROPDESC_SORTDESCRIPTION int32
const (
	PDSD_GENERAL PROPDESC_SORTDESCRIPTION = 0
	PDSD_A_Z PROPDESC_SORTDESCRIPTION = 1
	PDSD_LOWEST_HIGHEST PROPDESC_SORTDESCRIPTION = 2
	PDSD_SMALLEST_BIGGEST PROPDESC_SORTDESCRIPTION = 3
	PDSD_OLDEST_NEWEST PROPDESC_SORTDESCRIPTION = 4
)

// enum PROPDESC_RELATIVEDESCRIPTION_TYPE
type PROPDESC_RELATIVEDESCRIPTION_TYPE int32
const (
	PDRDT_GENERAL PROPDESC_RELATIVEDESCRIPTION_TYPE = 0
	PDRDT_DATE PROPDESC_RELATIVEDESCRIPTION_TYPE = 1
	PDRDT_SIZE PROPDESC_RELATIVEDESCRIPTION_TYPE = 2
	PDRDT_COUNT PROPDESC_RELATIVEDESCRIPTION_TYPE = 3
	PDRDT_REVISION PROPDESC_RELATIVEDESCRIPTION_TYPE = 4
	PDRDT_LENGTH PROPDESC_RELATIVEDESCRIPTION_TYPE = 5
	PDRDT_DURATION PROPDESC_RELATIVEDESCRIPTION_TYPE = 6
	PDRDT_SPEED PROPDESC_RELATIVEDESCRIPTION_TYPE = 7
	PDRDT_RATE PROPDESC_RELATIVEDESCRIPTION_TYPE = 8
	PDRDT_RATING PROPDESC_RELATIVEDESCRIPTION_TYPE = 9
	PDRDT_PRIORITY PROPDESC_RELATIVEDESCRIPTION_TYPE = 10
)

// enum PROPDESC_AGGREGATION_TYPE
type PROPDESC_AGGREGATION_TYPE int32
const (
	PDAT_DEFAULT PROPDESC_AGGREGATION_TYPE = 0
	PDAT_FIRST PROPDESC_AGGREGATION_TYPE = 1
	PDAT_SUM PROPDESC_AGGREGATION_TYPE = 2
	PDAT_AVERAGE PROPDESC_AGGREGATION_TYPE = 3
	PDAT_DATERANGE PROPDESC_AGGREGATION_TYPE = 4
	PDAT_UNION PROPDESC_AGGREGATION_TYPE = 5
	PDAT_MAX PROPDESC_AGGREGATION_TYPE = 6
	PDAT_MIN PROPDESC_AGGREGATION_TYPE = 7
)

// enum PROPDESC_CONDITION_TYPE
type PROPDESC_CONDITION_TYPE int32
const (
	PDCOT_NONE PROPDESC_CONDITION_TYPE = 0
	PDCOT_STRING PROPDESC_CONDITION_TYPE = 1
	PDCOT_SIZE PROPDESC_CONDITION_TYPE = 2
	PDCOT_DATETIME PROPDESC_CONDITION_TYPE = 3
	PDCOT_BOOLEAN PROPDESC_CONDITION_TYPE = 4
	PDCOT_NUMBER PROPDESC_CONDITION_TYPE = 5
)

// enum PROPDESC_SEARCHINFO_FLAGS
// flags
type PROPDESC_SEARCHINFO_FLAGS uint32
const (
	PDSIF_DEFAULT PROPDESC_SEARCHINFO_FLAGS = 0
	PDSIF_ININVERTEDINDEX PROPDESC_SEARCHINFO_FLAGS = 1
	PDSIF_ISCOLUMN PROPDESC_SEARCHINFO_FLAGS = 2
	PDSIF_ISCOLUMNSPARSE PROPDESC_SEARCHINFO_FLAGS = 4
	PDSIF_ALWAYSINCLUDE PROPDESC_SEARCHINFO_FLAGS = 8
	PDSIF_USEFORTYPEAHEAD PROPDESC_SEARCHINFO_FLAGS = 16
)

// enum PROPDESC_COLUMNINDEX_TYPE
type PROPDESC_COLUMNINDEX_TYPE int32
const (
	PDCIT_NONE PROPDESC_COLUMNINDEX_TYPE = 0
	PDCIT_ONDISK PROPDESC_COLUMNINDEX_TYPE = 1
	PDCIT_INMEMORY PROPDESC_COLUMNINDEX_TYPE = 2
	PDCIT_ONDEMAND PROPDESC_COLUMNINDEX_TYPE = 3
	PDCIT_ONDISKALL PROPDESC_COLUMNINDEX_TYPE = 4
	PDCIT_ONDISKVECTOR PROPDESC_COLUMNINDEX_TYPE = 5
)

// enum PROPDESC_ENUMFILTER
type PROPDESC_ENUMFILTER int32
const (
	PDEF_ALL PROPDESC_ENUMFILTER = 0
	PDEF_SYSTEM PROPDESC_ENUMFILTER = 1
	PDEF_NONSYSTEM PROPDESC_ENUMFILTER = 2
	PDEF_VIEWABLE PROPDESC_ENUMFILTER = 3
	PDEF_QUERYABLE PROPDESC_ENUMFILTER = 4
	PDEF_INFULLTEXTQUERY PROPDESC_ENUMFILTER = 5
	PDEF_COLUMN PROPDESC_ENUMFILTER = 6
)

// enum PERSIST_SPROPSTORE_FLAGS_
type PERSIST_SPROPSTORE_FLAGS_ int32
const (
	FPSPS_DEFAULT PERSIST_SPROPSTORE_FLAGS_ = 0
	FPSPS_READONLY PERSIST_SPROPSTORE_FLAGS_ = 1
	FPSPS_TREAT_NEW_VALUES_AS_DIRTY PERSIST_SPROPSTORE_FLAGS_ = 2
)

// enum PSTIME_FLAGS
// flags
type PSTIME_FLAGS uint32
const (
	PSTF_UTC PSTIME_FLAGS = 0
	PSTF_LOCAL PSTIME_FLAGS = 1
)

// enum PROPVAR_COMPARE_UNIT
type PROPVAR_COMPARE_UNIT int32
const (
	PVCU_DEFAULT PROPVAR_COMPARE_UNIT = 0
	PVCU_SECOND PROPVAR_COMPARE_UNIT = 1
	PVCU_MINUTE PROPVAR_COMPARE_UNIT = 2
	PVCU_HOUR PROPVAR_COMPARE_UNIT = 3
	PVCU_DAY PROPVAR_COMPARE_UNIT = 4
	PVCU_MONTH PROPVAR_COMPARE_UNIT = 5
	PVCU_YEAR PROPVAR_COMPARE_UNIT = 6
)

// enum PROPVAR_COMPARE_FLAGS
// flags
type PROPVAR_COMPARE_FLAGS uint32
const (
	PVCF_DEFAULT PROPVAR_COMPARE_FLAGS = 0
	PVCF_TREATEMPTYASGREATERTHAN PROPVAR_COMPARE_FLAGS = 1
	PVCF_USESTRCMP PROPVAR_COMPARE_FLAGS = 2
	PVCF_USESTRCMPC PROPVAR_COMPARE_FLAGS = 4
	PVCF_USESTRCMPI PROPVAR_COMPARE_FLAGS = 8
	PVCF_USESTRCMPIC PROPVAR_COMPARE_FLAGS = 16
	PVCF_DIGITSASNUMBERS_CASESENSITIVE PROPVAR_COMPARE_FLAGS = 32
)

// enum PROPVAR_CHANGE_FLAGS
// flags
type PROPVAR_CHANGE_FLAGS uint32
const (
	PVCHF_DEFAULT PROPVAR_CHANGE_FLAGS = 0
	PVCHF_NOVALUEPROP PROPVAR_CHANGE_FLAGS = 1
	PVCHF_ALPHABOOL PROPVAR_CHANGE_FLAGS = 2
	PVCHF_NOUSEROVERRIDE PROPVAR_CHANGE_FLAGS = 4
	PVCHF_LOCALBOOL PROPVAR_CHANGE_FLAGS = 8
	PVCHF_NOHEXSTRING PROPVAR_CHANGE_FLAGS = 16
)

// enum DRAWPROGRESSFLAGS
// flags
type DRAWPROGRESSFLAGS uint32
const (
	DPF_NONE DRAWPROGRESSFLAGS = 0
	DPF_MARQUEE DRAWPROGRESSFLAGS = 1
	DPF_MARQUEE_COMPLETE DRAWPROGRESSFLAGS = 2
	DPF_ERROR DRAWPROGRESSFLAGS = 4
	DPF_WARNING DRAWPROGRESSFLAGS = 8
	DPF_STOPPED DRAWPROGRESSFLAGS = 16
)

// enum SYNC_TRANSFER_STATUS
// flags
type SYNC_TRANSFER_STATUS uint32
const (
	STS_NONE SYNC_TRANSFER_STATUS = 0
	STS_NEEDSUPLOAD SYNC_TRANSFER_STATUS = 1
	STS_NEEDSDOWNLOAD SYNC_TRANSFER_STATUS = 2
	STS_TRANSFERRING SYNC_TRANSFER_STATUS = 4
	STS_PAUSED SYNC_TRANSFER_STATUS = 8
	STS_HASERROR SYNC_TRANSFER_STATUS = 16
	STS_FETCHING_METADATA SYNC_TRANSFER_STATUS = 32
	STS_USER_REQUESTED_REFRESH SYNC_TRANSFER_STATUS = 64
	STS_HASWARNING SYNC_TRANSFER_STATUS = 128
	STS_EXCLUDED SYNC_TRANSFER_STATUS = 256
	STS_INCOMPLETE SYNC_TRANSFER_STATUS = 512
	STS_PLACEHOLDER_IFEMPTY SYNC_TRANSFER_STATUS = 1024
)

// enum PLACEHOLDER_STATES
// flags
type PLACEHOLDER_STATES uint32
const (
	PS_NONE PLACEHOLDER_STATES = 0
	PS_MARKED_FOR_OFFLINE_AVAILABILITY PLACEHOLDER_STATES = 1
	PS_FULL_PRIMARY_STREAM_AVAILABLE PLACEHOLDER_STATES = 2
	PS_CREATE_FILE_ACCESSIBLE PLACEHOLDER_STATES = 4
	PS_CLOUDFILE_PLACEHOLDER PLACEHOLDER_STATES = 8
	PS_DEFAULT PLACEHOLDER_STATES = 7
	PS_ALL PLACEHOLDER_STATES = 15
)

// enum PROPERTYUI_NAME_FLAGS
// flags
type PROPERTYUI_NAME_FLAGS uint32
const (
	PUIFNF_DEFAULT PROPERTYUI_NAME_FLAGS = 0
	PUIFNF_MNEMONIC PROPERTYUI_NAME_FLAGS = 1
)

// enum PROPERTYUI_FLAGS
// flags
type PROPERTYUI_FLAGS uint32
const (
	PUIF_DEFAULT PROPERTYUI_FLAGS = 0
	PUIF_RIGHTALIGN PROPERTYUI_FLAGS = 1
	PUIF_NOLABELININFOTIP PROPERTYUI_FLAGS = 2
)

// enum PROPERTYUI_FORMAT_FLAGS
// flags
type PROPERTYUI_FORMAT_FLAGS uint32
const (
	PUIFFDF_DEFAULT PROPERTYUI_FORMAT_FLAGS = 0
	PUIFFDF_RIGHTTOLEFT PROPERTYUI_FORMAT_FLAGS = 1
	PUIFFDF_SHORTFORMAT PROPERTYUI_FORMAT_FLAGS = 2
	PUIFFDF_NOTIME PROPERTYUI_FORMAT_FLAGS = 4
	PUIFFDF_FRIENDLYDATE PROPERTYUI_FORMAT_FLAGS = 8
)

// enum PDOPSTATUS
type PDOPSTATUS int32
const (
	PDOPS_RUNNING PDOPSTATUS = 1
	PDOPS_PAUSED PDOPSTATUS = 2
	PDOPS_CANCELLED PDOPSTATUS = 3
	PDOPS_STOPPED PDOPSTATUS = 4
	PDOPS_ERRORS PDOPSTATUS = 5
)

// enum SYNC_ENGINE_STATE_FLAGS
// flags
type SYNC_ENGINE_STATE_FLAGS uint32
const (
	SESF_NONE SYNC_ENGINE_STATE_FLAGS = 0
	SESF_SERVICE_QUOTA_NEARING_LIMIT SYNC_ENGINE_STATE_FLAGS = 1
	SESF_SERVICE_QUOTA_EXCEEDED_LIMIT SYNC_ENGINE_STATE_FLAGS = 2
	SESF_AUTHENTICATION_ERROR SYNC_ENGINE_STATE_FLAGS = 4
	SESF_PAUSED_DUE_TO_METERED_NETWORK SYNC_ENGINE_STATE_FLAGS = 8
	SESF_PAUSED_DUE_TO_DISK_SPACE_FULL SYNC_ENGINE_STATE_FLAGS = 16
	SESF_PAUSED_DUE_TO_CLIENT_POLICY SYNC_ENGINE_STATE_FLAGS = 32
	SESF_PAUSED_DUE_TO_SERVICE_POLICY SYNC_ENGINE_STATE_FLAGS = 64
	SESF_SERVICE_UNAVAILABLE SYNC_ENGINE_STATE_FLAGS = 128
	SESF_PAUSED_DUE_TO_USER_REQUEST SYNC_ENGINE_STATE_FLAGS = 256
	SESF_ALL_FLAGS SYNC_ENGINE_STATE_FLAGS = 511
)


// structs

type PROPERTYKEY struct {
	Fmtid syscall.GUID
	Pid uint32
}

type SERIALIZEDPROPSTORAGE struct {
}

type PROPPRG struct {
	FlPrg uint16
	FlPrgInit uint16
	AchTitle [30]CHAR
	AchCmdLine [128]CHAR
	AchWorkDir [64]CHAR
	WHotKey uint16
	AchIconFile [80]CHAR
	WIconIndex uint16
	DwEnhModeFlags uint32
	DwRealModeFlags uint32
	AchOtherFile [80]CHAR
	AchPIFFile [260]CHAR
}


// coms

// b7d14566-0509-4cce-a71f-0a554233bd9b
var IID_IInitializeWithFile = syscall.GUID{0xb7d14566, 0x0509, 0x4cce, 
	[8]byte{0xa7, 0x1f, 0x0a, 0x55, 0x42, 0x33, 0xbd, 0x9b}}

type IInitializeWithFileInterface interface {
	IUnknownInterface
	Initialize(pszFilePath PWSTR, grfMode uint32) HRESULT
}

type IInitializeWithFileVtbl struct {
	IUnknownVtbl
	Initialize uintptr
}

type IInitializeWithFile struct {
	IUnknown
}

func (this *IInitializeWithFile) Vtbl() *IInitializeWithFileVtbl {
	return (*IInitializeWithFileVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInitializeWithFile) Initialize(pszFilePath PWSTR, grfMode uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Initialize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszFilePath)), uintptr(grfMode))
	return HRESULT(ret)
}

// b824b49d-22ac-4161-ac8a-9916e8fa3f7f
var IID_IInitializeWithStream = syscall.GUID{0xb824b49d, 0x22ac, 0x4161, 
	[8]byte{0xac, 0x8a, 0x99, 0x16, 0xe8, 0xfa, 0x3f, 0x7f}}

type IInitializeWithStreamInterface interface {
	IUnknownInterface
	Initialize(pstream *IStream, grfMode uint32) HRESULT
}

type IInitializeWithStreamVtbl struct {
	IUnknownVtbl
	Initialize uintptr
}

type IInitializeWithStream struct {
	IUnknown
}

func (this *IInitializeWithStream) Vtbl() *IInitializeWithStreamVtbl {
	return (*IInitializeWithStreamVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInitializeWithStream) Initialize(pstream *IStream, grfMode uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Initialize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstream)), uintptr(grfMode))
	return HRESULT(ret)
}

// 886d8eeb-8cf2-4446-8d02-cdba1dbdcf99
var IID_IPropertyStore = syscall.GUID{0x886d8eeb, 0x8cf2, 0x4446, 
	[8]byte{0x8d, 0x02, 0xcd, 0xba, 0x1d, 0xbd, 0xcf, 0x99}}

type IPropertyStoreInterface interface {
	IUnknownInterface
	GetCount(cProps *uint32) HRESULT
	GetAt(iProp uint32, pkey *PROPERTYKEY) HRESULT
	GetValue(key *PROPERTYKEY, pv *PROPVARIANT) HRESULT
	SetValue(key *PROPERTYKEY, propvar *PROPVARIANT) HRESULT
	Commit() HRESULT
}

type IPropertyStoreVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt uintptr
	GetValue uintptr
	SetValue uintptr
	Commit uintptr
}

type IPropertyStore struct {
	IUnknown
}

func (this *IPropertyStore) Vtbl() *IPropertyStoreVtbl {
	return (*IPropertyStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyStore) GetCount(cProps *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cProps)))
	return HRESULT(ret)
}

func (this *IPropertyStore) GetAt(iProp uint32, pkey *PROPERTYKEY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(iProp), uintptr(unsafe.Pointer(pkey)))
	return HRESULT(ret)
}

func (this *IPropertyStore) GetValue(key *PROPERTYKEY, pv *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(pv)))
	return HRESULT(ret)
}

func (this *IPropertyStore) SetValue(key *PROPERTYKEY, propvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(propvar)))
	return HRESULT(ret)
}

func (this *IPropertyStore) Commit() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Commit, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 71604b0f-97b0-4764-8577-2f13e98a1422
var IID_INamedPropertyStore = syscall.GUID{0x71604b0f, 0x97b0, 0x4764, 
	[8]byte{0x85, 0x77, 0x2f, 0x13, 0xe9, 0x8a, 0x14, 0x22}}

type INamedPropertyStoreInterface interface {
	IUnknownInterface
	GetNamedValue(pszName PWSTR, ppropvar *PROPVARIANT) HRESULT
	SetNamedValue(pszName PWSTR, propvar *PROPVARIANT) HRESULT
	GetNameCount(pdwCount *uint32) HRESULT
	GetNameAt(iProp uint32, pbstrName *BSTR) HRESULT
}

type INamedPropertyStoreVtbl struct {
	IUnknownVtbl
	GetNamedValue uintptr
	SetNamedValue uintptr
	GetNameCount uintptr
	GetNameAt uintptr
}

type INamedPropertyStore struct {
	IUnknown
}

func (this *INamedPropertyStore) Vtbl() *INamedPropertyStoreVtbl {
	return (*INamedPropertyStoreVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *INamedPropertyStore) GetNamedValue(pszName PWSTR, ppropvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetNamedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszName)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func (this *INamedPropertyStore) SetNamedValue(pszName PWSTR, propvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetNamedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszName)), uintptr(unsafe.Pointer(propvar)))
	return HRESULT(ret)
}

func (this *INamedPropertyStore) GetNameCount(pdwCount *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetNameCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwCount)))
	return HRESULT(ret)
}

func (this *INamedPropertyStore) GetNameAt(iProp uint32, pbstrName *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetNameAt, uintptr(unsafe.Pointer(this)), uintptr(iProp), uintptr(unsafe.Pointer(pbstrName)))
	return HRESULT(ret)
}

// fc0ca0a7-c316-4fd2-9031-3e628e6d4f23
var IID_IObjectWithPropertyKey = syscall.GUID{0xfc0ca0a7, 0xc316, 0x4fd2, 
	[8]byte{0x90, 0x31, 0x3e, 0x62, 0x8e, 0x6d, 0x4f, 0x23}}

type IObjectWithPropertyKeyInterface interface {
	IUnknownInterface
	SetPropertyKey(key *PROPERTYKEY) HRESULT
	GetPropertyKey(pkey *PROPERTYKEY) HRESULT
}

type IObjectWithPropertyKeyVtbl struct {
	IUnknownVtbl
	SetPropertyKey uintptr
	GetPropertyKey uintptr
}

type IObjectWithPropertyKey struct {
	IUnknown
}

func (this *IObjectWithPropertyKey) Vtbl() *IObjectWithPropertyKeyVtbl {
	return (*IObjectWithPropertyKeyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObjectWithPropertyKey) SetPropertyKey(key *PROPERTYKEY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPropertyKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IObjectWithPropertyKey) GetPropertyKey(pkey *PROPERTYKEY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pkey)))
	return HRESULT(ret)
}

// f917bc8a-1bba-4478-a245-1bde03eb9431
var IID_IPropertyChange = syscall.GUID{0xf917bc8a, 0x1bba, 0x4478, 
	[8]byte{0xa2, 0x45, 0x1b, 0xde, 0x03, 0xeb, 0x94, 0x31}}

type IPropertyChangeInterface interface {
	IObjectWithPropertyKeyInterface
	ApplyToPropVariant(propvarIn *PROPVARIANT, ppropvarOut *PROPVARIANT) HRESULT
}

type IPropertyChangeVtbl struct {
	IObjectWithPropertyKeyVtbl
	ApplyToPropVariant uintptr
}

type IPropertyChange struct {
	IObjectWithPropertyKey
}

func (this *IPropertyChange) Vtbl() *IPropertyChangeVtbl {
	return (*IPropertyChangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyChange) ApplyToPropVariant(propvarIn *PROPVARIANT, ppropvarOut *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ApplyToPropVariant, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(ppropvarOut)))
	return HRESULT(ret)
}

// 380f5cad-1b5e-42f2-805d-637fd392d31e
var IID_IPropertyChangeArray = syscall.GUID{0x380f5cad, 0x1b5e, 0x42f2, 
	[8]byte{0x80, 0x5d, 0x63, 0x7f, 0xd3, 0x92, 0xd3, 0x1e}}

type IPropertyChangeArrayInterface interface {
	IUnknownInterface
	GetCount(pcOperations *uint32) HRESULT
	GetAt(iIndex uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	InsertAt(iIndex uint32, ppropChange *IPropertyChange) HRESULT
	Append(ppropChange *IPropertyChange) HRESULT
	AppendOrReplace(ppropChange *IPropertyChange) HRESULT
	RemoveAt(iIndex uint32) HRESULT
	IsKeyInArray(key *PROPERTYKEY) HRESULT
}

type IPropertyChangeArrayVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt uintptr
	InsertAt uintptr
	Append uintptr
	AppendOrReplace uintptr
	RemoveAt uintptr
	IsKeyInArray uintptr
}

type IPropertyChangeArray struct {
	IUnknown
}

func (this *IPropertyChangeArray) Vtbl() *IPropertyChangeArrayVtbl {
	return (*IPropertyChangeArrayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyChangeArray) GetCount(pcOperations *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcOperations)))
	return HRESULT(ret)
}

func (this *IPropertyChangeArray) GetAt(iIndex uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(iIndex), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertyChangeArray) InsertAt(iIndex uint32, ppropChange *IPropertyChange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(iIndex), uintptr(unsafe.Pointer(ppropChange)))
	return HRESULT(ret)
}

func (this *IPropertyChangeArray) Append(ppropChange *IPropertyChange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppropChange)))
	return HRESULT(ret)
}

func (this *IPropertyChangeArray) AppendOrReplace(ppropChange *IPropertyChange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AppendOrReplace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppropChange)))
	return HRESULT(ret)
}

func (this *IPropertyChangeArray) RemoveAt(iIndex uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(iIndex))
	return HRESULT(ret)
}

func (this *IPropertyChangeArray) IsKeyInArray(key *PROPERTYKEY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsKeyInArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

// c8e2d566-186e-4d49-bf41-6909ead56acc
var IID_IPropertyStoreCapabilities = syscall.GUID{0xc8e2d566, 0x186e, 0x4d49, 
	[8]byte{0xbf, 0x41, 0x69, 0x09, 0xea, 0xd5, 0x6a, 0xcc}}

type IPropertyStoreCapabilitiesInterface interface {
	IUnknownInterface
	IsPropertyWritable(key *PROPERTYKEY) HRESULT
}

type IPropertyStoreCapabilitiesVtbl struct {
	IUnknownVtbl
	IsPropertyWritable uintptr
}

type IPropertyStoreCapabilities struct {
	IUnknown
}

func (this *IPropertyStoreCapabilities) Vtbl() *IPropertyStoreCapabilitiesVtbl {
	return (*IPropertyStoreCapabilitiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyStoreCapabilities) IsPropertyWritable(key *PROPERTYKEY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsPropertyWritable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

// 3017056d-9a91-4e90-937d-746c72abbf4f
var IID_IPropertyStoreCache = syscall.GUID{0x3017056d, 0x9a91, 0x4e90, 
	[8]byte{0x93, 0x7d, 0x74, 0x6c, 0x72, 0xab, 0xbf, 0x4f}}

type IPropertyStoreCacheInterface interface {
	IPropertyStoreInterface
	GetState(key *PROPERTYKEY, pstate *PSC_STATE) HRESULT
	GetValueAndState(key *PROPERTYKEY, ppropvar *PROPVARIANT, pstate *PSC_STATE) HRESULT
	SetState(key *PROPERTYKEY, state PSC_STATE) HRESULT
	SetValueAndState(key *PROPERTYKEY, ppropvar *PROPVARIANT, state PSC_STATE) HRESULT
}

type IPropertyStoreCacheVtbl struct {
	IPropertyStoreVtbl
	GetState uintptr
	GetValueAndState uintptr
	SetState uintptr
	SetValueAndState uintptr
}

type IPropertyStoreCache struct {
	IPropertyStore
}

func (this *IPropertyStoreCache) Vtbl() *IPropertyStoreCacheVtbl {
	return (*IPropertyStoreCacheVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyStoreCache) GetState(key *PROPERTYKEY, pstate *PSC_STATE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(pstate)))
	return HRESULT(ret)
}

func (this *IPropertyStoreCache) GetValueAndState(key *PROPERTYKEY, ppropvar *PROPVARIANT, pstate *PSC_STATE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetValueAndState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(ppropvar)), uintptr(unsafe.Pointer(pstate)))
	return HRESULT(ret)
}

func (this *IPropertyStoreCache) SetState(key *PROPERTYKEY, state PSC_STATE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(state))
	return HRESULT(ret)
}

func (this *IPropertyStoreCache) SetValueAndState(key *PROPERTYKEY, ppropvar *PROPVARIANT, state PSC_STATE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValueAndState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(ppropvar)), uintptr(state))
	return HRESULT(ret)
}

// 11e1fbf9-2d56-4a6b-8db3-7cd193a471f2
var IID_IPropertyEnumType = syscall.GUID{0x11e1fbf9, 0x2d56, 0x4a6b, 
	[8]byte{0x8d, 0xb3, 0x7c, 0xd1, 0x93, 0xa4, 0x71, 0xf2}}

type IPropertyEnumTypeInterface interface {
	IUnknownInterface
	GetEnumType(penumtype *PROPENUMTYPE) HRESULT
	GetValue(ppropvar *PROPVARIANT) HRESULT
	GetRangeMinValue(ppropvarMin *PROPVARIANT) HRESULT
	GetRangeSetValue(ppropvarSet *PROPVARIANT) HRESULT
	GetDisplayText(ppszDisplay *PWSTR) HRESULT
}

type IPropertyEnumTypeVtbl struct {
	IUnknownVtbl
	GetEnumType uintptr
	GetValue uintptr
	GetRangeMinValue uintptr
	GetRangeSetValue uintptr
	GetDisplayText uintptr
}

type IPropertyEnumType struct {
	IUnknown
}

func (this *IPropertyEnumType) Vtbl() *IPropertyEnumTypeVtbl {
	return (*IPropertyEnumTypeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyEnumType) GetEnumType(penumtype *PROPENUMTYPE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEnumType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penumtype)))
	return HRESULT(ret)
}

func (this *IPropertyEnumType) GetValue(ppropvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func (this *IPropertyEnumType) GetRangeMinValue(ppropvarMin *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRangeMinValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppropvarMin)))
	return HRESULT(ret)
}

func (this *IPropertyEnumType) GetRangeSetValue(ppropvarSet *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRangeSetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppropvarSet)))
	return HRESULT(ret)
}

func (this *IPropertyEnumType) GetDisplayText(ppszDisplay *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppszDisplay)))
	return HRESULT(ret)
}

// 9b6e051c-5ddd-4321-9070-fe2acb55e794
var IID_IPropertyEnumType2 = syscall.GUID{0x9b6e051c, 0x5ddd, 0x4321, 
	[8]byte{0x90, 0x70, 0xfe, 0x2a, 0xcb, 0x55, 0xe7, 0x94}}

type IPropertyEnumType2Interface interface {
	IPropertyEnumTypeInterface
	GetImageReference(ppszImageRes *PWSTR) HRESULT
}

type IPropertyEnumType2Vtbl struct {
	IPropertyEnumTypeVtbl
	GetImageReference uintptr
}

type IPropertyEnumType2 struct {
	IPropertyEnumType
}

func (this *IPropertyEnumType2) Vtbl() *IPropertyEnumType2Vtbl {
	return (*IPropertyEnumType2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyEnumType2) GetImageReference(ppszImageRes *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetImageReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppszImageRes)))
	return HRESULT(ret)
}

// a99400f4-3d84-4557-94ba-1242fb2cc9a6
var IID_IPropertyEnumTypeList = syscall.GUID{0xa99400f4, 0x3d84, 0x4557, 
	[8]byte{0x94, 0xba, 0x12, 0x42, 0xfb, 0x2c, 0xc9, 0xa6}}

type IPropertyEnumTypeListInterface interface {
	IUnknownInterface
	GetCount(pctypes *uint32) HRESULT
	GetAt(itype uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	GetConditionAt(nIndex uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	FindMatchingIndex(propvarCmp *PROPVARIANT, pnIndex *uint32) HRESULT
}

type IPropertyEnumTypeListVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt uintptr
	GetConditionAt uintptr
	FindMatchingIndex uintptr
}

type IPropertyEnumTypeList struct {
	IUnknown
}

func (this *IPropertyEnumTypeList) Vtbl() *IPropertyEnumTypeListVtbl {
	return (*IPropertyEnumTypeListVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyEnumTypeList) GetCount(pctypes *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pctypes)))
	return HRESULT(ret)
}

func (this *IPropertyEnumTypeList) GetAt(itype uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(itype), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertyEnumTypeList) GetConditionAt(nIndex uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetConditionAt, uintptr(unsafe.Pointer(this)), uintptr(nIndex), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertyEnumTypeList) FindMatchingIndex(propvarCmp *PROPVARIANT, pnIndex *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindMatchingIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propvarCmp)), uintptr(unsafe.Pointer(pnIndex)))
	return HRESULT(ret)
}

// 6f79d558-3e96-4549-a1d1-7d75d2288814
var IID_IPropertyDescription = syscall.GUID{0x6f79d558, 0x3e96, 0x4549, 
	[8]byte{0xa1, 0xd1, 0x7d, 0x75, 0xd2, 0x28, 0x88, 0x14}}

type IPropertyDescriptionInterface interface {
	IUnknownInterface
	GetPropertyKey(pkey *PROPERTYKEY) HRESULT
	GetCanonicalName(ppszName *PWSTR) HRESULT
	GetPropertyType(pvartype *uint16) HRESULT
	GetDisplayName(ppszName *PWSTR) HRESULT
	GetEditInvitation(ppszInvite *PWSTR) HRESULT
	GetTypeFlags(mask PROPDESC_TYPE_FLAGS, ppdtFlags *PROPDESC_TYPE_FLAGS) HRESULT
	GetViewFlags(ppdvFlags *PROPDESC_VIEW_FLAGS) HRESULT
	GetDefaultColumnWidth(pcxChars *uint32) HRESULT
	GetDisplayType(pdisplaytype *PROPDESC_DISPLAYTYPE) HRESULT
	GetColumnState(pcsFlags *uint32) HRESULT
	GetGroupingRange(pgr *PROPDESC_GROUPING_RANGE) HRESULT
	GetRelativeDescriptionType(prdt *PROPDESC_RELATIVEDESCRIPTION_TYPE) HRESULT
	GetRelativeDescription(propvar1 *PROPVARIANT, propvar2 *PROPVARIANT, ppszDesc1 *PWSTR, ppszDesc2 *PWSTR) HRESULT
	GetSortDescription(psd *PROPDESC_SORTDESCRIPTION) HRESULT
	GetSortDescriptionLabel(fDescending BOOL, ppszDescription *PWSTR) HRESULT
	GetAggregationType(paggtype *PROPDESC_AGGREGATION_TYPE) HRESULT
	GetConditionType(pcontype *PROPDESC_CONDITION_TYPE, popDefault *CONDITION_OPERATION) HRESULT
	GetEnumTypeList(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	CoerceToCanonicalValue(ppropvar *PROPVARIANT) HRESULT
	FormatForDisplay(propvar *PROPVARIANT, pdfFlags PROPDESC_FORMAT_FLAGS, ppszDisplay *PWSTR) HRESULT
	IsValueCanonical(propvar *PROPVARIANT) HRESULT
}

type IPropertyDescriptionVtbl struct {
	IUnknownVtbl
	GetPropertyKey uintptr
	GetCanonicalName uintptr
	GetPropertyType uintptr
	GetDisplayName uintptr
	GetEditInvitation uintptr
	GetTypeFlags uintptr
	GetViewFlags uintptr
	GetDefaultColumnWidth uintptr
	GetDisplayType uintptr
	GetColumnState uintptr
	GetGroupingRange uintptr
	GetRelativeDescriptionType uintptr
	GetRelativeDescription uintptr
	GetSortDescription uintptr
	GetSortDescriptionLabel uintptr
	GetAggregationType uintptr
	GetConditionType uintptr
	GetEnumTypeList uintptr
	CoerceToCanonicalValue uintptr
	FormatForDisplay uintptr
	IsValueCanonical uintptr
}

type IPropertyDescription struct {
	IUnknown
}

func (this *IPropertyDescription) Vtbl() *IPropertyDescriptionVtbl {
	return (*IPropertyDescriptionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyDescription) GetPropertyKey(pkey *PROPERTYKEY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pkey)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetCanonicalName(ppszName *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCanonicalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppszName)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetPropertyType(pvartype *uint16) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pvartype)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetDisplayName(ppszName *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppszName)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetEditInvitation(ppszInvite *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEditInvitation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppszInvite)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetTypeFlags(mask PROPDESC_TYPE_FLAGS, ppdtFlags *PROPDESC_TYPE_FLAGS) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTypeFlags, uintptr(unsafe.Pointer(this)), uintptr(mask), uintptr(unsafe.Pointer(ppdtFlags)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetViewFlags(ppdvFlags *PROPDESC_VIEW_FLAGS) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetViewFlags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppdvFlags)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetDefaultColumnWidth(pcxChars *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultColumnWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcxChars)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetDisplayType(pdisplaytype *PROPDESC_DISPLAYTYPE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdisplaytype)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetColumnState(pcsFlags *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColumnState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcsFlags)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetGroupingRange(pgr *PROPDESC_GROUPING_RANGE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGroupingRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pgr)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetRelativeDescriptionType(prdt *PROPDESC_RELATIVEDESCRIPTION_TYPE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRelativeDescriptionType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(prdt)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetRelativeDescription(propvar1 *PROPVARIANT, propvar2 *PROPVARIANT, ppszDesc1 *PWSTR, ppszDesc2 *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRelativeDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propvar1)), uintptr(unsafe.Pointer(propvar2)), uintptr(unsafe.Pointer(ppszDesc1)), uintptr(unsafe.Pointer(ppszDesc2)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetSortDescription(psd *PROPDESC_SORTDESCRIPTION) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSortDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(psd)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetSortDescriptionLabel(fDescending BOOL, ppszDescription *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSortDescriptionLabel, uintptr(unsafe.Pointer(this)), uintptr(fDescending), uintptr(unsafe.Pointer(ppszDescription)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetAggregationType(paggtype *PROPDESC_AGGREGATION_TYPE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAggregationType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(paggtype)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetConditionType(pcontype *PROPDESC_CONDITION_TYPE, popDefault *CONDITION_OPERATION) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetConditionType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcontype)), uintptr(unsafe.Pointer(popDefault)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) GetEnumTypeList(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEnumTypeList, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) CoerceToCanonicalValue(ppropvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CoerceToCanonicalValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) FormatForDisplay(propvar *PROPVARIANT, pdfFlags PROPDESC_FORMAT_FLAGS, ppszDisplay *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FormatForDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propvar)), uintptr(pdfFlags), uintptr(unsafe.Pointer(ppszDisplay)))
	return HRESULT(ret)
}

func (this *IPropertyDescription) IsValueCanonical(propvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsValueCanonical, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propvar)))
	return HRESULT(ret)
}

// 57d2eded-5062-400e-b107-5dae79fe57a6
var IID_IPropertyDescription2 = syscall.GUID{0x57d2eded, 0x5062, 0x400e, 
	[8]byte{0xb1, 0x07, 0x5d, 0xae, 0x79, 0xfe, 0x57, 0xa6}}

type IPropertyDescription2Interface interface {
	IPropertyDescriptionInterface
	GetImageReferenceForValue(propvar *PROPVARIANT, ppszImageRes *PWSTR) HRESULT
}

type IPropertyDescription2Vtbl struct {
	IPropertyDescriptionVtbl
	GetImageReferenceForValue uintptr
}

type IPropertyDescription2 struct {
	IPropertyDescription
}

func (this *IPropertyDescription2) Vtbl() *IPropertyDescription2Vtbl {
	return (*IPropertyDescription2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyDescription2) GetImageReferenceForValue(propvar *PROPVARIANT, ppszImageRes *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetImageReferenceForValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(ppszImageRes)))
	return HRESULT(ret)
}

// f67104fc-2af9-46fd-b32d-243c1404f3d1
var IID_IPropertyDescriptionAliasInfo = syscall.GUID{0xf67104fc, 0x2af9, 0x46fd, 
	[8]byte{0xb3, 0x2d, 0x24, 0x3c, 0x14, 0x04, 0xf3, 0xd1}}

type IPropertyDescriptionAliasInfoInterface interface {
	IPropertyDescriptionInterface
	GetSortByAlias(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	GetAdditionalSortByAliases(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type IPropertyDescriptionAliasInfoVtbl struct {
	IPropertyDescriptionVtbl
	GetSortByAlias uintptr
	GetAdditionalSortByAliases uintptr
}

type IPropertyDescriptionAliasInfo struct {
	IPropertyDescription
}

func (this *IPropertyDescriptionAliasInfo) Vtbl() *IPropertyDescriptionAliasInfoVtbl {
	return (*IPropertyDescriptionAliasInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyDescriptionAliasInfo) GetSortByAlias(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSortByAlias, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertyDescriptionAliasInfo) GetAdditionalSortByAliases(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAdditionalSortByAliases, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// 078f91bd-29a2-440f-924e-46a291524520
var IID_IPropertyDescriptionSearchInfo = syscall.GUID{0x078f91bd, 0x29a2, 0x440f, 
	[8]byte{0x92, 0x4e, 0x46, 0xa2, 0x91, 0x52, 0x45, 0x20}}

type IPropertyDescriptionSearchInfoInterface interface {
	IPropertyDescriptionInterface
	GetSearchInfoFlags(ppdsiFlags *PROPDESC_SEARCHINFO_FLAGS) HRESULT
	GetColumnIndexType(ppdciType *PROPDESC_COLUMNINDEX_TYPE) HRESULT
	GetProjectionString(ppszProjection *PWSTR) HRESULT
	GetMaxSize(pcbMaxSize *uint32) HRESULT
}

type IPropertyDescriptionSearchInfoVtbl struct {
	IPropertyDescriptionVtbl
	GetSearchInfoFlags uintptr
	GetColumnIndexType uintptr
	GetProjectionString uintptr
	GetMaxSize uintptr
}

type IPropertyDescriptionSearchInfo struct {
	IPropertyDescription
}

func (this *IPropertyDescriptionSearchInfo) Vtbl() *IPropertyDescriptionSearchInfoVtbl {
	return (*IPropertyDescriptionSearchInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyDescriptionSearchInfo) GetSearchInfoFlags(ppdsiFlags *PROPDESC_SEARCHINFO_FLAGS) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSearchInfoFlags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppdsiFlags)))
	return HRESULT(ret)
}

func (this *IPropertyDescriptionSearchInfo) GetColumnIndexType(ppdciType *PROPDESC_COLUMNINDEX_TYPE) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColumnIndexType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppdciType)))
	return HRESULT(ret)
}

func (this *IPropertyDescriptionSearchInfo) GetProjectionString(ppszProjection *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetProjectionString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppszProjection)))
	return HRESULT(ret)
}

func (this *IPropertyDescriptionSearchInfo) GetMaxSize(pcbMaxSize *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetMaxSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcbMaxSize)))
	return HRESULT(ret)
}

// 507393f4-2a3d-4a60-b59e-d9c75716c2dd
var IID_IPropertyDescriptionRelatedPropertyInfo = syscall.GUID{0x507393f4, 0x2a3d, 0x4a60, 
	[8]byte{0xb5, 0x9e, 0xd9, 0xc7, 0x57, 0x16, 0xc2, 0xdd}}

type IPropertyDescriptionRelatedPropertyInfoInterface interface {
	IPropertyDescriptionInterface
	GetRelatedProperty(pszRelationshipName PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type IPropertyDescriptionRelatedPropertyInfoVtbl struct {
	IPropertyDescriptionVtbl
	GetRelatedProperty uintptr
}

type IPropertyDescriptionRelatedPropertyInfo struct {
	IPropertyDescription
}

func (this *IPropertyDescriptionRelatedPropertyInfo) Vtbl() *IPropertyDescriptionRelatedPropertyInfoVtbl {
	return (*IPropertyDescriptionRelatedPropertyInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyDescriptionRelatedPropertyInfo) GetRelatedProperty(pszRelationshipName PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRelatedProperty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszRelationshipName)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// ca724e8a-c3e6-442b-88a4-6fb0db8035a3
var IID_IPropertySystem = syscall.GUID{0xca724e8a, 0xc3e6, 0x442b, 
	[8]byte{0x88, 0xa4, 0x6f, 0xb0, 0xdb, 0x80, 0x35, 0xa3}}

type IPropertySystemInterface interface {
	IUnknownInterface
	GetPropertyDescription(propkey *PROPERTYKEY, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	GetPropertyDescriptionByName(pszCanonicalName PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	GetPropertyDescriptionListFromString(pszPropList PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	EnumeratePropertyDescriptions(filterOn PROPDESC_ENUMFILTER, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	FormatForDisplay(key *PROPERTYKEY, propvar *PROPVARIANT, pdff PROPDESC_FORMAT_FLAGS, pszText *uint16, cchText uint32) HRESULT
	FormatForDisplayAlloc(key *PROPERTYKEY, propvar *PROPVARIANT, pdff PROPDESC_FORMAT_FLAGS, ppszDisplay *PWSTR) HRESULT
	RegisterPropertySchema(pszPath PWSTR) HRESULT
	UnregisterPropertySchema(pszPath PWSTR) HRESULT
	RefreshPropertySchema() HRESULT
}

type IPropertySystemVtbl struct {
	IUnknownVtbl
	GetPropertyDescription uintptr
	GetPropertyDescriptionByName uintptr
	GetPropertyDescriptionListFromString uintptr
	EnumeratePropertyDescriptions uintptr
	FormatForDisplay uintptr
	FormatForDisplayAlloc uintptr
	RegisterPropertySchema uintptr
	UnregisterPropertySchema uintptr
	RefreshPropertySchema uintptr
}

type IPropertySystem struct {
	IUnknown
}

func (this *IPropertySystem) Vtbl() *IPropertySystemVtbl {
	return (*IPropertySystemVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertySystem) GetPropertyDescription(propkey *PROPERTYKEY, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propkey)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertySystem) GetPropertyDescriptionByName(pszCanonicalName PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyDescriptionByName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszCanonicalName)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertySystem) GetPropertyDescriptionListFromString(pszPropList PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyDescriptionListFromString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPropList)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertySystem) EnumeratePropertyDescriptions(filterOn PROPDESC_ENUMFILTER, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().EnumeratePropertyDescriptions, uintptr(unsafe.Pointer(this)), uintptr(filterOn), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertySystem) FormatForDisplay(key *PROPERTYKEY, propvar *PROPVARIANT, pdff PROPDESC_FORMAT_FLAGS, pszText *uint16, cchText uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FormatForDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(propvar)), uintptr(pdff), uintptr(unsafe.Pointer(pszText)), uintptr(cchText))
	return HRESULT(ret)
}

func (this *IPropertySystem) FormatForDisplayAlloc(key *PROPERTYKEY, propvar *PROPVARIANT, pdff PROPDESC_FORMAT_FLAGS, ppszDisplay *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FormatForDisplayAlloc, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(propvar)), uintptr(pdff), uintptr(unsafe.Pointer(ppszDisplay)))
	return HRESULT(ret)
}

func (this *IPropertySystem) RegisterPropertySchema(pszPath PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RegisterPropertySchema, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPath)))
	return HRESULT(ret)
}

func (this *IPropertySystem) UnregisterPropertySchema(pszPath PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().UnregisterPropertySchema, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPath)))
	return HRESULT(ret)
}

func (this *IPropertySystem) RefreshPropertySchema() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RefreshPropertySchema, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 1f9fc1d0-c39b-4b26-817f-011967d3440e
var IID_IPropertyDescriptionList = syscall.GUID{0x1f9fc1d0, 0xc39b, 0x4b26, 
	[8]byte{0x81, 0x7f, 0x01, 0x19, 0x67, 0xd3, 0x44, 0x0e}}

type IPropertyDescriptionListInterface interface {
	IUnknownInterface
	GetCount(pcElem *uint32) HRESULT
	GetAt(iElem uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type IPropertyDescriptionListVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt uintptr
}

type IPropertyDescriptionList struct {
	IUnknown
}

func (this *IPropertyDescriptionList) Vtbl() *IPropertyDescriptionListVtbl {
	return (*IPropertyDescriptionListVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyDescriptionList) GetCount(pcElem *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func (this *IPropertyDescriptionList) GetAt(iElem uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(iElem), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// bc110b6d-57e8-4148-a9c6-91015ab2f3a5
var IID_IPropertyStoreFactory = syscall.GUID{0xbc110b6d, 0x57e8, 0x4148, 
	[8]byte{0xa9, 0xc6, 0x91, 0x01, 0x5a, 0xb2, 0xf3, 0xa5}}

type IPropertyStoreFactoryInterface interface {
	IUnknownInterface
	GetPropertyStore(flags GETPROPERTYSTOREFLAGS, pUnkFactory *IUnknown, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
	GetPropertyStoreForKeys(rgKeys *PROPERTYKEY, cKeys uint32, flags GETPROPERTYSTOREFLAGS, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type IPropertyStoreFactoryVtbl struct {
	IUnknownVtbl
	GetPropertyStore uintptr
	GetPropertyStoreForKeys uintptr
}

type IPropertyStoreFactory struct {
	IUnknown
}

func (this *IPropertyStoreFactory) Vtbl() *IPropertyStoreFactoryVtbl {
	return (*IPropertyStoreFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyStoreFactory) GetPropertyStore(flags GETPROPERTYSTOREFLAGS, pUnkFactory *IUnknown, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyStore, uintptr(unsafe.Pointer(this)), uintptr(flags), uintptr(unsafe.Pointer(pUnkFactory)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

func (this *IPropertyStoreFactory) GetPropertyStoreForKeys(rgKeys *PROPERTYKEY, cKeys uint32, flags GETPROPERTYSTOREFLAGS, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyStoreForKeys, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rgKeys)), uintptr(cKeys), uintptr(flags), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// 40d4577f-e237-4bdb-bd69-58f089431b6a
var IID_IDelayedPropertyStoreFactory = syscall.GUID{0x40d4577f, 0xe237, 0x4bdb, 
	[8]byte{0xbd, 0x69, 0x58, 0xf0, 0x89, 0x43, 0x1b, 0x6a}}

type IDelayedPropertyStoreFactoryInterface interface {
	IPropertyStoreFactoryInterface
	GetDelayedPropertyStore(flags GETPROPERTYSTOREFLAGS, dwStoreId uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type IDelayedPropertyStoreFactoryVtbl struct {
	IPropertyStoreFactoryVtbl
	GetDelayedPropertyStore uintptr
}

type IDelayedPropertyStoreFactory struct {
	IPropertyStoreFactory
}

func (this *IDelayedPropertyStoreFactory) Vtbl() *IDelayedPropertyStoreFactoryVtbl {
	return (*IDelayedPropertyStoreFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDelayedPropertyStoreFactory) GetDelayedPropertyStore(flags GETPROPERTYSTOREFLAGS, dwStoreId uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDelayedPropertyStore, uintptr(unsafe.Pointer(this)), uintptr(flags), uintptr(dwStoreId), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// e318ad57-0aa0-450f-aca5-6fab7103d917
var IID_IPersistSerializedPropStorage = syscall.GUID{0xe318ad57, 0x0aa0, 0x450f, 
	[8]byte{0xac, 0xa5, 0x6f, 0xab, 0x71, 0x03, 0xd9, 0x17}}

type IPersistSerializedPropStorageInterface interface {
	IUnknownInterface
	SetFlags(flags int32) HRESULT
	SetPropertyStorage(psps *SERIALIZEDPROPSTORAGE, cb uint32) HRESULT
	GetPropertyStorage(ppsps **SERIALIZEDPROPSTORAGE, pcb *uint32) HRESULT
}

type IPersistSerializedPropStorageVtbl struct {
	IUnknownVtbl
	SetFlags uintptr
	SetPropertyStorage uintptr
	GetPropertyStorage uintptr
}

type IPersistSerializedPropStorage struct {
	IUnknown
}

func (this *IPersistSerializedPropStorage) Vtbl() *IPersistSerializedPropStorageVtbl {
	return (*IPersistSerializedPropStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPersistSerializedPropStorage) SetFlags(flags int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFlags, uintptr(unsafe.Pointer(this)), uintptr(flags))
	return HRESULT(ret)
}

func (this *IPersistSerializedPropStorage) SetPropertyStorage(psps *SERIALIZEDPROPSTORAGE, cb uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPropertyStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(psps)), uintptr(cb))
	return HRESULT(ret)
}

func (this *IPersistSerializedPropStorage) GetPropertyStorage(ppsps **SERIALIZEDPROPSTORAGE, pcb *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppsps)), uintptr(unsafe.Pointer(pcb)))
	return HRESULT(ret)
}

// 77effa68-4f98-4366-ba72-573b3d880571
var IID_IPersistSerializedPropStorage2 = syscall.GUID{0x77effa68, 0x4f98, 0x4366, 
	[8]byte{0xba, 0x72, 0x57, 0x3b, 0x3d, 0x88, 0x05, 0x71}}

type IPersistSerializedPropStorage2Interface interface {
	IPersistSerializedPropStorageInterface
	GetPropertyStorageSize(pcb *uint32) HRESULT
	GetPropertyStorageBuffer(psps *SERIALIZEDPROPSTORAGE, cb uint32, pcbWritten *uint32) HRESULT
}

type IPersistSerializedPropStorage2Vtbl struct {
	IPersistSerializedPropStorageVtbl
	GetPropertyStorageSize uintptr
	GetPropertyStorageBuffer uintptr
}

type IPersistSerializedPropStorage2 struct {
	IPersistSerializedPropStorage
}

func (this *IPersistSerializedPropStorage2) Vtbl() *IPersistSerializedPropStorage2Vtbl {
	return (*IPersistSerializedPropStorage2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPersistSerializedPropStorage2) GetPropertyStorageSize(pcb *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyStorageSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcb)))
	return HRESULT(ret)
}

func (this *IPersistSerializedPropStorage2) GetPropertyStorageBuffer(psps *SERIALIZEDPROPSTORAGE, cb uint32, pcbWritten *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyStorageBuffer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(psps)), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

// fa955fd9-38be-4879-a6ce-824cf52d609f
var IID_IPropertySystemChangeNotify = syscall.GUID{0xfa955fd9, 0x38be, 0x4879, 
	[8]byte{0xa6, 0xce, 0x82, 0x4c, 0xf5, 0x2d, 0x60, 0x9f}}

type IPropertySystemChangeNotifyInterface interface {
	IUnknownInterface
	SchemaRefreshed() HRESULT
}

type IPropertySystemChangeNotifyVtbl struct {
	IUnknownVtbl
	SchemaRefreshed uintptr
}

type IPropertySystemChangeNotify struct {
	IUnknown
}

func (this *IPropertySystemChangeNotify) Vtbl() *IPropertySystemChangeNotifyVtbl {
	return (*IPropertySystemChangeNotifyVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertySystemChangeNotify) SchemaRefreshed() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SchemaRefreshed, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 75121952-e0d0-43e5-9380-1d80483acf72
var IID_ICreateObject = syscall.GUID{0x75121952, 0xe0d0, 0x43e5, 
	[8]byte{0x93, 0x80, 0x1d, 0x80, 0x48, 0x3a, 0xcf, 0x72}}

type ICreateObjectInterface interface {
	IUnknownInterface
	CreateObject(clsid *syscall.GUID, pUnkOuter *IUnknown, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type ICreateObjectVtbl struct {
	IUnknownVtbl
	CreateObject uintptr
}

type ICreateObject struct {
	IUnknown
}

func (this *ICreateObject) Vtbl() *ICreateObjectVtbl {
	return (*ICreateObjectVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateObject) CreateObject(clsid *syscall.GUID, pUnkOuter *IUnknown, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clsid)), uintptr(unsafe.Pointer(pUnkOuter)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// 757a7d9f-919a-4118-99d7-dbb208c8cc66
var IID_IPropertyUI = syscall.GUID{0x757a7d9f, 0x919a, 0x4118, 
	[8]byte{0x99, 0xd7, 0xdb, 0xb2, 0x08, 0xc8, 0xcc, 0x66}}

type IPropertyUIInterface interface {
	IUnknownInterface
	ParsePropertyName(pszName PWSTR, pfmtid *syscall.GUID, ppid *uint32, pchEaten *uint32) HRESULT
	GetCannonicalName(fmtid *syscall.GUID, pid uint32, pwszText *uint16, cchText uint32) HRESULT
	GetDisplayName(fmtid *syscall.GUID, pid uint32, flags PROPERTYUI_NAME_FLAGS, pwszText *uint16, cchText uint32) HRESULT
	GetPropertyDescription(fmtid *syscall.GUID, pid uint32, pwszText *uint16, cchText uint32) HRESULT
	GetDefaultWidth(fmtid *syscall.GUID, pid uint32, pcxChars *uint32) HRESULT
	GetFlags(fmtid *syscall.GUID, pid uint32, pflags *PROPERTYUI_FLAGS) HRESULT
	FormatForDisplay(fmtid *syscall.GUID, pid uint32, ppropvar *PROPVARIANT, puiff PROPERTYUI_FORMAT_FLAGS, pwszText *uint16, cchText uint32) HRESULT
	GetHelpInfo(fmtid *syscall.GUID, pid uint32, pwszHelpFile *uint16, cch uint32, puHelpID *uint32) HRESULT
}

type IPropertyUIVtbl struct {
	IUnknownVtbl
	ParsePropertyName uintptr
	GetCannonicalName uintptr
	GetDisplayName uintptr
	GetPropertyDescription uintptr
	GetDefaultWidth uintptr
	GetFlags uintptr
	FormatForDisplay uintptr
	GetHelpInfo uintptr
}

type IPropertyUI struct {
	IUnknown
}

func (this *IPropertyUI) Vtbl() *IPropertyUIVtbl {
	return (*IPropertyUIVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyUI) ParsePropertyName(pszName PWSTR, pfmtid *syscall.GUID, ppid *uint32, pchEaten *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ParsePropertyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszName)), uintptr(unsafe.Pointer(pfmtid)), uintptr(unsafe.Pointer(ppid)), uintptr(unsafe.Pointer(pchEaten)))
	return HRESULT(ret)
}

func (this *IPropertyUI) GetCannonicalName(fmtid *syscall.GUID, pid uint32, pwszText *uint16, cchText uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCannonicalName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(unsafe.Pointer(pwszText)), uintptr(cchText))
	return HRESULT(ret)
}

func (this *IPropertyUI) GetDisplayName(fmtid *syscall.GUID, pid uint32, flags PROPERTYUI_NAME_FLAGS, pwszText *uint16, cchText uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(flags), uintptr(unsafe.Pointer(pwszText)), uintptr(cchText))
	return HRESULT(ret)
}

func (this *IPropertyUI) GetPropertyDescription(fmtid *syscall.GUID, pid uint32, pwszText *uint16, cchText uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(unsafe.Pointer(pwszText)), uintptr(cchText))
	return HRESULT(ret)
}

func (this *IPropertyUI) GetDefaultWidth(fmtid *syscall.GUID, pid uint32, pcxChars *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultWidth, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(unsafe.Pointer(pcxChars)))
	return HRESULT(ret)
}

func (this *IPropertyUI) GetFlags(fmtid *syscall.GUID, pid uint32, pflags *PROPERTYUI_FLAGS) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFlags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(unsafe.Pointer(pflags)))
	return HRESULT(ret)
}

func (this *IPropertyUI) FormatForDisplay(fmtid *syscall.GUID, pid uint32, ppropvar *PROPVARIANT, puiff PROPERTYUI_FORMAT_FLAGS, pwszText *uint16, cchText uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FormatForDisplay, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(unsafe.Pointer(ppropvar)), uintptr(puiff), uintptr(unsafe.Pointer(pwszText)), uintptr(cchText))
	return HRESULT(ret)
}

func (this *IPropertyUI) GetHelpInfo(fmtid *syscall.GUID, pid uint32, pwszHelpFile *uint16, cch uint32, puHelpID *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetHelpInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fmtid)), uintptr(pid), uintptr(unsafe.Pointer(pwszHelpFile)), uintptr(cch), uintptr(unsafe.Pointer(puHelpID)))
	return HRESULT(ret)
}


var (
	pSHGetPropertyStoreFromIDList uintptr
	pSHGetPropertyStoreFromParsingName uintptr
	pSHAddDefaultPropertiesByExt uintptr
	pPifMgr_OpenProperties uintptr
	pPifMgr_GetProperties uintptr
	pPifMgr_SetProperties uintptr
	pPifMgr_CloseProperties uintptr
	pSHPropStgCreate uintptr
	pSHPropStgReadMultiple uintptr
	pSHPropStgWriteMultiple uintptr
	pSHGetPropertyStoreForWindow uintptr
)

func SHGetPropertyStoreFromIDList(pidl *ITEMIDLIST, flags GETPROPERTYSTOREFLAGS, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := lazyAddr(&pSHGetPropertyStoreFromIDList, libShell32, "SHGetPropertyStoreFromIDList")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pidl)), uintptr(flags), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func SHGetPropertyStoreFromParsingName(pszPath PWSTR, pbc *IBindCtx, flags GETPROPERTYSTOREFLAGS, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := lazyAddr(&pSHGetPropertyStoreFromParsingName, libShell32, "SHGetPropertyStoreFromParsingName")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszPath)), uintptr(unsafe.Pointer(pbc)), uintptr(flags), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func SHAddDefaultPropertiesByExt(pszExt PWSTR, pPropStore *IPropertyStore) HRESULT {
	addr := lazyAddr(&pSHAddDefaultPropertiesByExt, libShell32, "SHAddDefaultPropertiesByExt")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszExt)), uintptr(unsafe.Pointer(pPropStore)))
	return HRESULT(ret)
}

func PifMgr_OpenProperties(pszApp PWSTR, pszPIF PWSTR, hInf uint32, flOpt uint32) HANDLE {
	addr := lazyAddr(&pPifMgr_OpenProperties, libShell32, "PifMgr_OpenProperties")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszApp)), uintptr(unsafe.Pointer(pszPIF)), uintptr(hInf), uintptr(flOpt))
	return HANDLE(ret)
}

func PifMgr_GetProperties(hProps HANDLE, pszGroup PSTR, lpProps unsafe.Pointer, cbProps int32, flOpt uint32) int32 {
	addr := lazyAddr(&pPifMgr_GetProperties, libShell32, "PifMgr_GetProperties")
	ret, _,  _ := syscall.SyscallN(addr, hProps, uintptr(unsafe.Pointer(pszGroup)), uintptr(lpProps), uintptr(cbProps), uintptr(flOpt))
	return int32(ret)
}

func PifMgr_SetProperties(hProps HANDLE, pszGroup PSTR, lpProps unsafe.Pointer, cbProps int32, flOpt uint32) int32 {
	addr := lazyAddr(&pPifMgr_SetProperties, libShell32, "PifMgr_SetProperties")
	ret, _,  _ := syscall.SyscallN(addr, hProps, uintptr(unsafe.Pointer(pszGroup)), uintptr(lpProps), uintptr(cbProps), uintptr(flOpt))
	return int32(ret)
}

func PifMgr_CloseProperties(hProps HANDLE, flOpt uint32) HANDLE {
	addr := lazyAddr(&pPifMgr_CloseProperties, libShell32, "PifMgr_CloseProperties")
	ret, _,  _ := syscall.SyscallN(addr, hProps, uintptr(flOpt))
	return HANDLE(ret)
}

func SHPropStgCreate(psstg *IPropertySetStorage, fmtid *syscall.GUID, pclsid *syscall.GUID, grfFlags uint32, grfMode uint32, dwDisposition uint32, ppstg **IPropertyStorage, puCodePage *uint32) HRESULT {
	addr := lazyAddr(&pSHPropStgCreate, libShell32, "SHPropStgCreate")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(psstg)), uintptr(unsafe.Pointer(fmtid)), uintptr(unsafe.Pointer(pclsid)), uintptr(grfFlags), uintptr(grfMode), uintptr(dwDisposition), uintptr(unsafe.Pointer(ppstg)), uintptr(unsafe.Pointer(puCodePage)))
	return HRESULT(ret)
}

func SHPropStgReadMultiple(pps *IPropertyStorage, uCodePage uint32, cpspec uint32, rgpspec *PROPSPEC, rgvar *PROPVARIANT) HRESULT {
	addr := lazyAddr(&pSHPropStgReadMultiple, libShell32, "SHPropStgReadMultiple")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pps)), uintptr(uCodePage), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)), uintptr(unsafe.Pointer(rgvar)))
	return HRESULT(ret)
}

func SHPropStgWriteMultiple(pps *IPropertyStorage, puCodePage *uint32, cpspec uint32, rgpspec *PROPSPEC, rgvar *PROPVARIANT, propidNameFirst uint32) HRESULT {
	addr := lazyAddr(&pSHPropStgWriteMultiple, libShell32, "SHPropStgWriteMultiple")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pps)), uintptr(unsafe.Pointer(puCodePage)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)), uintptr(unsafe.Pointer(rgvar)), uintptr(propidNameFirst))
	return HRESULT(ret)
}

func SHGetPropertyStoreForWindow(hwnd HWND, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := lazyAddr(&pSHGetPropertyStoreForWindow, libShell32, "SHGetPropertyStoreForWindow")
	ret, _,  _ := syscall.SyscallN(addr, hwnd, uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}


