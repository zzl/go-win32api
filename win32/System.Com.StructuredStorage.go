package win32

import (
	"syscall"
	"unsafe"
)

const (
	PROPSETFLAG_DEFAULT             uint32 = 0x0
	PROPSETFLAG_NONSIMPLE           uint32 = 0x1
	PROPSETFLAG_ANSI                uint32 = 0x2
	PROPSETFLAG_UNBUFFERED          uint32 = 0x4
	PROPSETFLAG_CASE_SENSITIVE      uint32 = 0x8
	PROPSET_BEHAVIOR_CASE_SENSITIVE uint32 = 0x1
	PID_DICTIONARY                  uint32 = 0x0
	PID_CODEPAGE                    uint32 = 0x1
	PID_FIRST_USABLE                uint32 = 0x2
	PID_FIRST_NAME_DEFAULT          uint32 = 0xfff
	PID_LOCALE                      uint32 = 0x80000000
	PID_MODIFY_TIME                 uint32 = 0x80000001
	PID_SECURITY                    uint32 = 0x80000002
	PID_BEHAVIOR                    uint32 = 0x80000003
	PID_ILLEGAL                     uint32 = 0xffffffff
	PID_MIN_READONLY                uint32 = 0x80000000
	PID_MAX_READONLY                uint32 = 0xbfffffff
	PRSPEC_INVALID                  uint32 = 0xffffffff
	PROPSETHDR_OSVERSION_UNKNOWN    uint32 = 0xffffffff
	PIDDI_THUMBNAIL                 int32  = 2
	PIDSI_TITLE                     int32  = 2
	PIDSI_SUBJECT                   int32  = 3
	PIDSI_AUTHOR                    int32  = 4
	PIDSI_KEYWORDS                  int32  = 5
	PIDSI_COMMENTS                  int32  = 6
	PIDSI_TEMPLATE                  int32  = 7
	PIDSI_LASTAUTHOR                int32  = 8
	PIDSI_REVNUMBER                 int32  = 9
	PIDSI_EDITTIME                  int32  = 10
	PIDSI_LASTPRINTED               int32  = 11
	PIDSI_CREATE_DTM                int32  = 12
	PIDSI_LASTSAVE_DTM              int32  = 13
	PIDSI_PAGECOUNT                 int32  = 14
	PIDSI_WORDCOUNT                 int32  = 15
	PIDSI_CHARCOUNT                 int32  = 16
	PIDSI_THUMBNAIL                 int32  = 17
	PIDSI_APPNAME                   int32  = 18
	PIDSI_DOC_SECURITY              int32  = 19
	PIDDSI_CATEGORY                 uint32 = 0x2
	PIDDSI_PRESFORMAT               uint32 = 0x3
	PIDDSI_BYTECOUNT                uint32 = 0x4
	PIDDSI_LINECOUNT                uint32 = 0x5
	PIDDSI_PARCOUNT                 uint32 = 0x6
	PIDDSI_SLIDECOUNT               uint32 = 0x7
	PIDDSI_NOTECOUNT                uint32 = 0x8
	PIDDSI_HIDDENCOUNT              uint32 = 0x9
	PIDDSI_MMCLIPCOUNT              uint32 = 0xa
	PIDDSI_SCALE                    uint32 = 0xb
	PIDDSI_HEADINGPAIR              uint32 = 0xc
	PIDDSI_DOCPARTS                 uint32 = 0xd
	PIDDSI_MANAGER                  uint32 = 0xe
	PIDDSI_COMPANY                  uint32 = 0xf
	PIDDSI_LINKSDIRTY               uint32 = 0x10
	PIDMSI_EDITOR                   int32  = 2
	PIDMSI_SUPPLIER                 int32  = 3
	PIDMSI_SOURCE                   int32  = 4
	PIDMSI_SEQUENCE_NO              int32  = 5
	PIDMSI_PROJECT                  int32  = 6
	PIDMSI_STATUS                   int32  = 7
	PIDMSI_OWNER                    int32  = 8
	PIDMSI_RATING                   int32  = 9
	PIDMSI_PRODUCTION               int32  = 10
	PIDMSI_COPYRIGHT                int32  = 11
	CWCSTORAGENAME                  uint32 = 0x20
	STGOPTIONS_VERSION              uint32 = 0x1
	CCH_MAX_PROPSTG_NAME            uint32 = 0x1f
)

// enums

// enum
type PROPSPEC_KIND uint32

const (
	PRSPEC_LPWSTR PROPSPEC_KIND = 0
	PRSPEC_PROPID PROPSPEC_KIND = 1
)

// enum
type STGFMT uint32

const (
	STGFMT_STORAGE  STGFMT = 0
	STGFMT_NATIVE   STGFMT = 1
	STGFMT_FILE     STGFMT = 3
	STGFMT_ANY      STGFMT = 4
	STGFMT_DOCFILE  STGFMT = 5
	STGFMT_DOCUMENT STGFMT = 0
)

// enum
type STGMOVE int32

const (
	STGMOVE_MOVE        STGMOVE = 0
	STGMOVE_COPY        STGMOVE = 1
	STGMOVE_SHALLOWCOPY STGMOVE = 2
)

// enum
type PIDMSI_STATUS_VALUE int32

const (
	PIDMSI_STATUS_NORMAL     PIDMSI_STATUS_VALUE = 0
	PIDMSI_STATUS_NEW        PIDMSI_STATUS_VALUE = 1
	PIDMSI_STATUS_PRELIM     PIDMSI_STATUS_VALUE = 2
	PIDMSI_STATUS_DRAFT      PIDMSI_STATUS_VALUE = 3
	PIDMSI_STATUS_INPROGRESS PIDMSI_STATUS_VALUE = 4
	PIDMSI_STATUS_EDIT       PIDMSI_STATUS_VALUE = 5
	PIDMSI_STATUS_REVIEW     PIDMSI_STATUS_VALUE = 6
	PIDMSI_STATUS_PROOF      PIDMSI_STATUS_VALUE = 7
	PIDMSI_STATUS_FINAL      PIDMSI_STATUS_VALUE = 8
	PIDMSI_STATUS_OTHER      PIDMSI_STATUS_VALUE = 32767
)

// enum
type PROPVAR_COMPARE_UNIT int32

const (
	PVCU_DEFAULT PROPVAR_COMPARE_UNIT = 0
	PVCU_SECOND  PROPVAR_COMPARE_UNIT = 1
	PVCU_MINUTE  PROPVAR_COMPARE_UNIT = 2
	PVCU_HOUR    PROPVAR_COMPARE_UNIT = 3
	PVCU_DAY     PROPVAR_COMPARE_UNIT = 4
	PVCU_MONTH   PROPVAR_COMPARE_UNIT = 5
	PVCU_YEAR    PROPVAR_COMPARE_UNIT = 6
)

// enum
// flags
type PROPVAR_COMPARE_FLAGS int32

const (
	PVCF_DEFAULT                       PROPVAR_COMPARE_FLAGS = 0
	PVCF_TREATEMPTYASGREATERTHAN       PROPVAR_COMPARE_FLAGS = 1
	PVCF_USESTRCMP                     PROPVAR_COMPARE_FLAGS = 2
	PVCF_USESTRCMPC                    PROPVAR_COMPARE_FLAGS = 4
	PVCF_USESTRCMPI                    PROPVAR_COMPARE_FLAGS = 8
	PVCF_USESTRCMPIC                   PROPVAR_COMPARE_FLAGS = 16
	PVCF_DIGITSASNUMBERS_CASESENSITIVE PROPVAR_COMPARE_FLAGS = 32
)

// enum
// flags
type PROPVAR_CHANGE_FLAGS int32

const (
	PVCHF_DEFAULT        PROPVAR_CHANGE_FLAGS = 0
	PVCHF_NOVALUEPROP    PROPVAR_CHANGE_FLAGS = 1
	PVCHF_ALPHABOOL      PROPVAR_CHANGE_FLAGS = 2
	PVCHF_NOUSEROVERRIDE PROPVAR_CHANGE_FLAGS = 4
	PVCHF_LOCALBOOL      PROPVAR_CHANGE_FLAGS = 8
	PVCHF_NOHEXSTRING    PROPVAR_CHANGE_FLAGS = 16
)

// structs

type BSTRBLOB struct {
	CbSize uint32
	PData  *byte
}

type CLIPDATA struct {
	CbSize    uint32
	UlClipFmt int32
	PClipData *byte
}

type RemSNB struct {
	UlCntStr  uint32
	UlCntChar uint32
	RgString  [1]uint16
}

type VERSIONEDSTREAM struct {
	GuidVersion syscall.GUID
	PStream     *IStream
}

type CAC struct {
	CElems uint32
	PElems PSTR
}

type CAUB struct {
	CElems uint32
	PElems *byte
}

type CAI struct {
	CElems uint32
	PElems *int16
}

type CAUI struct {
	CElems uint32
	PElems *uint16
}

type CAL struct {
	CElems uint32
	PElems *int32
}

type CAUL struct {
	CElems uint32
	PElems *uint32
}

type CAFLT struct {
	CElems uint32
	PElems *float32
}

type CADBL struct {
	CElems uint32
	PElems *float64
}

type CACY struct {
	CElems uint32
	PElems *CY
}

type CADATE struct {
	CElems uint32
	PElems *float64
}

type CABSTR struct {
	CElems uint32
	PElems *BSTR
}

type CABSTRBLOB struct {
	CElems uint32
	PElems *BSTRBLOB
}

type CABOOL struct {
	CElems uint32
	PElems *VARIANT_BOOL
}

type CASCODE struct {
	CElems uint32
	PElems *int32
}

type CAPROPVARIANT struct {
	CElems uint32
	PElems *PROPVARIANT
}

type CAH struct {
	CElems uint32
	PElems *int64
}

type CAUH struct {
	CElems uint32
	PElems *uint64
}

type CALPSTR struct {
	CElems uint32
	PElems *PSTR
}

type CALPWSTR struct {
	CElems uint32
	PElems *PWSTR
}

type CAFILETIME struct {
	CElems uint32
	PElems *FILETIME
}

type CACLIPDATA struct {
	CElems uint32
	PElems *CLIPDATA
}

type CACLSID struct {
	CElems uint32
	PElems *syscall.GUID
}

type PROPVARIANT_Anonymous_Anonymous_Anonymous struct {
	Data [2]uint64
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CVal() *CHAR {
	return (*CHAR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CValVal() CHAR {
	return *(*CHAR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BVal() *byte {
	return (*byte)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BValVal() byte {
	return *(*byte)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) IVal() *int16 {
	return (*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) IValVal() int16 {
	return *(*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UiVal() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UiValVal() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) LVal() *int32 {
	return (*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) LValVal() int32 {
	return *(*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UlVal() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UlValVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) IntVal() *int32 {
	return (*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) IntValVal() int32 {
	return *(*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UintVal() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UintValVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) HVal() *int64 {
	return (*int64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) HValVal() int64 {
	return *(*int64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UhVal() *uint64 {
	return (*uint64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) UhValVal() uint64 {
	return *(*uint64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) FltVal() *float32 {
	return (*float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) FltValVal() float32 {
	return *(*float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) DblVal() *float64 {
	return (*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) DblValVal() float64 {
	return *(*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BoolVal() *VARIANT_BOOL {
	return (*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BoolValVal() VARIANT_BOOL {
	return *(*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) OBSOLETE__VARIANT_BOOL__() *VARIANT_BOOL {
	return (*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) OBSOLETE__VARIANT_BOOL__Val() VARIANT_BOOL {
	return *(*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Scode() *int32 {
	return (*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) ScodeVal() int32 {
	return *(*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CyVal() *CY {
	return (*CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CyValVal() CY {
	return *(*CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Date() *float64 {
	return (*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) DateVal() float64 {
	return *(*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Filetime() *FILETIME {
	return (*FILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) FiletimeVal() FILETIME {
	return *(*FILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Puuid() **syscall.GUID {
	return (**syscall.GUID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PuuidVal() *syscall.GUID {
	return *(**syscall.GUID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Pclipdata() **CLIPDATA {
	return (**CLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PclipdataVal() *CLIPDATA {
	return *(**CLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BstrVal() *BSTR {
	return (*BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BstrValVal() BSTR {
	return *(*BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BstrblobVal() *BSTRBLOB {
	return (*BSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BstrblobValVal() BSTRBLOB {
	return *(*BSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Blob() *BLOB {
	return (*BLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) BlobVal() BLOB {
	return *(*BLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PszVal() *PSTR {
	return (*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PszValVal() PSTR {
	return *(*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PwszVal() *PWSTR {
	return (*PWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PwszValVal() PWSTR {
	return *(*PWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PunkVal() **IUnknown {
	return (**IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PunkValVal() *IUnknown {
	return *(**IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdispVal() **IDispatch {
	return (**IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdispValVal() *IDispatch {
	return *(**IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PStream() **IStream {
	return (**IStream)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PStreamVal() *IStream {
	return *(**IStream)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PStorage() **IStorage {
	return (**IStorage)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PStorageVal() *IStorage {
	return *(**IStorage)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PVersionedStream() **VERSIONEDSTREAM {
	return (**VERSIONEDSTREAM)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PVersionedStreamVal() *VERSIONEDSTREAM {
	return *(**VERSIONEDSTREAM)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Parray() **SAFEARRAY {
	return (**SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) ParrayVal() *SAFEARRAY {
	return *(**SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cac() *CAC {
	return (*CAC)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CacVal() CAC {
	return *(*CAC)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Caub() *CAUB {
	return (*CAUB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CaubVal() CAUB {
	return *(*CAUB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cai() *CAI {
	return (*CAI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CaiVal() CAI {
	return *(*CAI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Caui() *CAUI {
	return (*CAUI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CauiVal() CAUI {
	return *(*CAUI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cal() *CAL {
	return (*CAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CalVal() CAL {
	return *(*CAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Caul() *CAUL {
	return (*CAUL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CaulVal() CAUL {
	return *(*CAUL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cah() *CAH {
	return (*CAH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CahVal() CAH {
	return *(*CAH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cauh() *CAUH {
	return (*CAUH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CauhVal() CAUH {
	return *(*CAUH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Caflt() *CAFLT {
	return (*CAFLT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CafltVal() CAFLT {
	return *(*CAFLT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cadbl() *CADBL {
	return (*CADBL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CadblVal() CADBL {
	return *(*CADBL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cabool() *CABOOL {
	return (*CABOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CaboolVal() CABOOL {
	return *(*CABOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cascode() *CASCODE {
	return (*CASCODE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CascodeVal() CASCODE {
	return *(*CASCODE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cacy() *CACY {
	return (*CACY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CacyVal() CACY {
	return *(*CACY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cadate() *CADATE {
	return (*CADATE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CadateVal() CADATE {
	return *(*CADATE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cafiletime() *CAFILETIME {
	return (*CAFILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CafiletimeVal() CAFILETIME {
	return *(*CAFILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cauuid() *CACLSID {
	return (*CACLSID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CauuidVal() CACLSID {
	return *(*CACLSID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Caclipdata() *CACLIPDATA {
	return (*CACLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CaclipdataVal() CACLIPDATA {
	return *(*CACLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cabstr() *CABSTR {
	return (*CABSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CabstrVal() CABSTR {
	return *(*CABSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Cabstrblob() *CABSTRBLOB {
	return (*CABSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CabstrblobVal() CABSTRBLOB {
	return *(*CABSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Calpstr() *CALPSTR {
	return (*CALPSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CalpstrVal() CALPSTR {
	return *(*CALPSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Calpwstr() *CALPWSTR {
	return (*CALPWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CalpwstrVal() CALPWSTR {
	return *(*CALPWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Capropvar() *CAPROPVARIANT {
	return (*CAPROPVARIANT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) CapropvarVal() CAPROPVARIANT {
	return *(*CAPROPVARIANT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PcVal() *PSTR {
	return (*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PcValVal() PSTR {
	return *(*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PbVal() **byte {
	return (**byte)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PbValVal() *byte {
	return *(**byte)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PiVal() **int16 {
	return (**int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PiValVal() *int16 {
	return *(**int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PuiVal() **uint16 {
	return (**uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PuiValVal() *uint16 {
	return *(**uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PlVal() **int32 {
	return (**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PlValVal() *int32 {
	return *(**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PulVal() **uint32 {
	return (**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PulValVal() *uint32 {
	return *(**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PintVal() **int32 {
	return (**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PintValVal() *int32 {
	return *(**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PuintVal() **uint32 {
	return (**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PuintValVal() *uint32 {
	return *(**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PfltVal() **float32 {
	return (**float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PfltValVal() *float32 {
	return *(**float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdblVal() **float64 {
	return (**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdblValVal() *float64 {
	return *(**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PboolVal() **VARIANT_BOOL {
	return (**VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PboolValVal() *VARIANT_BOOL {
	return *(**VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdecVal() **DECIMAL {
	return (**DECIMAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdecValVal() *DECIMAL {
	return *(**DECIMAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Pscode() **int32 {
	return (**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PscodeVal() *int32 {
	return *(**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PcyVal() **CY {
	return (**CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PcyValVal() *CY {
	return *(**CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Pdate() **float64 {
	return (**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PdateVal() *float64 {
	return *(**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PbstrVal() **BSTR {
	return (**BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PbstrValVal() *BSTR {
	return *(**BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PpunkVal() ***IUnknown {
	return (***IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PpunkValVal() **IUnknown {
	return *(***IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PpdispVal() ***IDispatch {
	return (***IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PpdispValVal() **IDispatch {
	return *(***IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) Pparray() ***SAFEARRAY {
	return (***SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PparrayVal() **SAFEARRAY {
	return *(***SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PvarVal() **PROPVARIANT {
	return (**PROPVARIANT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_Anonymous_Anonymous) PvarValVal() *PROPVARIANT {
	return *(**PROPVARIANT)(unsafe.Pointer(this))
}

type PROPVARIANT_Anonymous_Anonymous struct {
	Vt         VARENUM
	WReserved1 uint16
	WReserved2 uint16
	WReserved3 uint16
	PROPVARIANT_Anonymous_Anonymous_Anonymous
}

type PROPVARIANT_Anonymous struct {
	PROPVARIANT_Anonymous_Anonymous
}

func (this *PROPVARIANT_Anonymous) Anonymous() *PROPVARIANT_Anonymous_Anonymous {
	return (*PROPVARIANT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous) AnonymousVal() PROPVARIANT_Anonymous_Anonymous {
	return *(*PROPVARIANT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous) DecVal() *DECIMAL {
	return (*DECIMAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous) DecValVal() DECIMAL {
	return *(*DECIMAL)(unsafe.Pointer(this))
}

type PROPVARIANT struct {
	PROPVARIANT_Anonymous
}

type PROPSPEC_Anonymous struct {
	Data [1]uint64
}

func (this *PROPSPEC_Anonymous) Propid() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *PROPSPEC_Anonymous) PropidVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *PROPSPEC_Anonymous) Lpwstr() *PWSTR {
	return (*PWSTR)(unsafe.Pointer(this))
}

func (this *PROPSPEC_Anonymous) LpwstrVal() PWSTR {
	return *(*PWSTR)(unsafe.Pointer(this))
}

type PROPSPEC struct {
	UlKind PROPSPEC_KIND
	PROPSPEC_Anonymous
}

type STATPROPSTG struct {
	LpwstrName PWSTR
	Propid     uint32
	Vt         VARENUM
}

type STATPROPSETSTG struct {
	Fmtid       syscall.GUID
	Clsid       syscall.GUID
	GrfFlags    uint32
	Mtime       FILETIME
	Ctime       FILETIME
	Atime       FILETIME
	DwOSVersion uint32
}

type STGOPTIONS struct {
	UsVersion        uint16
	Reserved         uint16
	UlSectorSize     uint32
	PwcsTemplateFile PWSTR
}

type SERIALIZEDPROPERTYVALUE struct {
	DwType uint32
	Rgb    [1]byte
}

type OLESTREAMVTBL struct {
	Get uintptr
	Put uintptr
}

type OLESTREAM struct {
	Lpstbl *OLESTREAMVTBL
}

type PROPBAG2 struct {
	DwType   uint32
	Vt       VARENUM
	CfType   uint16
	DwHint   uint32
	PstrName PWSTR
	Clsid    syscall.GUID
}

// interfaces

// 0000000D-0000-0000-C000-000000000046
var IID_IEnumSTATSTG = syscall.GUID{0x0000000D, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumSTATSTGInterface interface {
	IUnknownInterface
	Next(celt uint32, rgelt *STATSTG, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppenum **IEnumSTATSTG) HRESULT
}

type IEnumSTATSTGVtbl struct {
	IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumSTATSTG struct {
	IUnknown
}

func (this *IEnumSTATSTG) Vtbl() *IEnumSTATSTGVtbl {
	return (*IEnumSTATSTGVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumSTATSTG) Next(celt uint32, rgelt *STATSTG, pceltFetched *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgelt)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumSTATSTG) Skip(celt uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumSTATSTG) Reset() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumSTATSTG) Clone(ppenum **IEnumSTATSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 0000000B-0000-0000-C000-000000000046
var IID_IStorage = syscall.GUID{0x0000000B, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IStorageInterface interface {
	IUnknownInterface
	CreateStream(pwcsName PWSTR, grfMode STGM, reserved1 uint32, reserved2 uint32, ppstm **IStream) HRESULT
	OpenStream(pwcsName PWSTR, reserved1 unsafe.Pointer, grfMode STGM, reserved2 uint32, ppstm **IStream) HRESULT
	CreateStorage(pwcsName PWSTR, grfMode STGM, reserved1 uint32, reserved2 uint32, ppstg **IStorage) HRESULT
	OpenStorage(pwcsName PWSTR, pstgPriority *IStorage, grfMode STGM, snbExclude **uint16, reserved uint32, ppstg **IStorage) HRESULT
	CopyTo(ciidExclude uint32, rgiidExclude *syscall.GUID, snbExclude **uint16, pstgDest *IStorage) HRESULT
	MoveElementTo(pwcsName PWSTR, pstgDest *IStorage, pwcsNewName PWSTR, grfFlags STGMOVE) HRESULT
	Commit(grfCommitFlags uint32) HRESULT
	Revert() HRESULT
	EnumElements(reserved1 uint32, reserved2 unsafe.Pointer, reserved3 uint32, ppenum **IEnumSTATSTG) HRESULT
	DestroyElement(pwcsName PWSTR) HRESULT
	RenameElement(pwcsOldName PWSTR, pwcsNewName PWSTR) HRESULT
	SetElementTimes(pwcsName PWSTR, pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT
	SetClass(clsid *syscall.GUID) HRESULT
	SetStateBits(grfStateBits uint32, grfMask uint32) HRESULT
	Stat(pstatstg *STATSTG, grfStatFlag uint32) HRESULT
}

type IStorageVtbl struct {
	IUnknownVtbl
	CreateStream    uintptr
	OpenStream      uintptr
	CreateStorage   uintptr
	OpenStorage     uintptr
	CopyTo          uintptr
	MoveElementTo   uintptr
	Commit          uintptr
	Revert          uintptr
	EnumElements    uintptr
	DestroyElement  uintptr
	RenameElement   uintptr
	SetElementTimes uintptr
	SetClass        uintptr
	SetStateBits    uintptr
	Stat            uintptr
}

type IStorage struct {
	IUnknown
}

func (this *IStorage) Vtbl() *IStorageVtbl {
	return (*IStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorage) CreateStream(pwcsName PWSTR, grfMode STGM, reserved1 uint32, reserved2 uint32, ppstm **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(reserved1), uintptr(reserved2), uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret)
}

func (this *IStorage) OpenStream(pwcsName PWSTR, reserved1 unsafe.Pointer, grfMode STGM, reserved2 uint32, ppstm **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().OpenStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(reserved1), uintptr(grfMode), uintptr(reserved2), uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret)
}

func (this *IStorage) CreateStorage(pwcsName PWSTR, grfMode STGM, reserved1 uint32, reserved2 uint32, ppstg **IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(reserved1), uintptr(reserved2), uintptr(unsafe.Pointer(ppstg)))
	return HRESULT(ret)
}

func (this *IStorage) OpenStorage(pwcsName PWSTR, pstgPriority *IStorage, grfMode STGM, snbExclude **uint16, reserved uint32, ppstg **IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().OpenStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pstgPriority)), uintptr(grfMode), uintptr(unsafe.Pointer(snbExclude)), uintptr(reserved), uintptr(unsafe.Pointer(ppstg)))
	return HRESULT(ret)
}

func (this *IStorage) CopyTo(ciidExclude uint32, rgiidExclude *syscall.GUID, snbExclude **uint16, pstgDest *IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CopyTo, uintptr(unsafe.Pointer(this)), uintptr(ciidExclude), uintptr(unsafe.Pointer(rgiidExclude)), uintptr(unsafe.Pointer(snbExclude)), uintptr(unsafe.Pointer(pstgDest)))
	return HRESULT(ret)
}

func (this *IStorage) MoveElementTo(pwcsName PWSTR, pstgDest *IStorage, pwcsNewName PWSTR, grfFlags STGMOVE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().MoveElementTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pstgDest)), uintptr(unsafe.Pointer(pwcsNewName)), uintptr(grfFlags))
	return HRESULT(ret)
}

func (this *IStorage) Commit(grfCommitFlags uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Commit, uintptr(unsafe.Pointer(this)), uintptr(grfCommitFlags))
	return HRESULT(ret)
}

func (this *IStorage) Revert() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Revert, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IStorage) EnumElements(reserved1 uint32, reserved2 unsafe.Pointer, reserved3 uint32, ppenum **IEnumSTATSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().EnumElements, uintptr(unsafe.Pointer(this)), uintptr(reserved1), uintptr(reserved2), uintptr(reserved3), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

func (this *IStorage) DestroyElement(pwcsName PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().DestroyElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)))
	return HRESULT(ret)
}

func (this *IStorage) RenameElement(pwcsOldName PWSTR, pwcsNewName PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RenameElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsOldName)), uintptr(unsafe.Pointer(pwcsNewName)))
	return HRESULT(ret)
}

func (this *IStorage) SetElementTimes(pwcsName PWSTR, pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetElementTimes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pctime)), uintptr(unsafe.Pointer(patime)), uintptr(unsafe.Pointer(pmtime)))
	return HRESULT(ret)
}

func (this *IStorage) SetClass(clsid *syscall.GUID) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clsid)))
	return HRESULT(ret)
}

func (this *IStorage) SetStateBits(grfStateBits uint32, grfMask uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStateBits, uintptr(unsafe.Pointer(this)), uintptr(grfStateBits), uintptr(grfMask))
	return HRESULT(ret)
}

func (this *IStorage) Stat(pstatstg *STATSTG, grfStatFlag uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Stat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstatstg)), uintptr(grfStatFlag))
	return HRESULT(ret)
}

// 0000010A-0000-0000-C000-000000000046
var IID_IPersistStorage = syscall.GUID{0x0000010A, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IPersistStorageInterface interface {
	IPersistInterface
	IsDirty() HRESULT
	InitNew(pStg *IStorage) HRESULT
	Load(pStg *IStorage) HRESULT
	Save(pStgSave *IStorage, fSameAsLoad BOOL) HRESULT
	SaveCompleted(pStgNew *IStorage) HRESULT
	HandsOffStorage() HRESULT
}

type IPersistStorageVtbl struct {
	IPersistVtbl
	IsDirty         uintptr
	InitNew         uintptr
	Load            uintptr
	Save            uintptr
	SaveCompleted   uintptr
	HandsOffStorage uintptr
}

type IPersistStorage struct {
	IPersist
}

func (this *IPersistStorage) Vtbl() *IPersistStorageVtbl {
	return (*IPersistStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPersistStorage) IsDirty() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsDirty, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IPersistStorage) InitNew(pStg *IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InitNew, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStg)))
	return HRESULT(ret)
}

func (this *IPersistStorage) Load(pStg *IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Load, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStg)))
	return HRESULT(ret)
}

func (this *IPersistStorage) Save(pStgSave *IStorage, fSameAsLoad BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Save, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStgSave)), uintptr(fSameAsLoad))
	return HRESULT(ret)
}

func (this *IPersistStorage) SaveCompleted(pStgNew *IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SaveCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStgNew)))
	return HRESULT(ret)
}

func (this *IPersistStorage) HandsOffStorage() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandsOffStorage, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 0000000A-0000-0000-C000-000000000046
var IID_ILockBytes = syscall.GUID{0x0000000A, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type ILockBytesInterface interface {
	IUnknownInterface
	ReadAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbRead *uint32) HRESULT
	WriteAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT
	Flush() HRESULT
	SetSize(cb uint64) HRESULT
	LockRegion(libOffset uint64, cb uint64, dwLockType uint32) HRESULT
	UnlockRegion(libOffset uint64, cb uint64, dwLockType uint32) HRESULT
	Stat(pstatstg *STATSTG, grfStatFlag uint32) HRESULT
}

type ILockBytesVtbl struct {
	IUnknownVtbl
	ReadAt       uintptr
	WriteAt      uintptr
	Flush        uintptr
	SetSize      uintptr
	LockRegion   uintptr
	UnlockRegion uintptr
	Stat         uintptr
}

type ILockBytes struct {
	IUnknown
}

func (this *ILockBytes) Vtbl() *ILockBytesVtbl {
	return (*ILockBytesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockBytes) ReadAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbRead *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReadAt, uintptr(unsafe.Pointer(this)), uintptr(ulOffset), uintptr(pv), uintptr(cb), uintptr(unsafe.Pointer(pcbRead)))
	return HRESULT(ret)
}

func (this *ILockBytes) WriteAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteAt, uintptr(unsafe.Pointer(this)), uintptr(ulOffset), uintptr(pv), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

func (this *ILockBytes) Flush() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Flush, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILockBytes) SetSize(cb uint64) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSize, uintptr(unsafe.Pointer(this)), uintptr(cb))
	return HRESULT(ret)
}

func (this *ILockBytes) LockRegion(libOffset uint64, cb uint64, dwLockType uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().LockRegion, uintptr(unsafe.Pointer(this)), uintptr(libOffset), uintptr(cb), uintptr(dwLockType))
	return HRESULT(ret)
}

func (this *ILockBytes) UnlockRegion(libOffset uint64, cb uint64, dwLockType uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().UnlockRegion, uintptr(unsafe.Pointer(this)), uintptr(libOffset), uintptr(cb), uintptr(dwLockType))
	return HRESULT(ret)
}

func (this *ILockBytes) Stat(pstatstg *STATSTG, grfStatFlag uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Stat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstatstg)), uintptr(grfStatFlag))
	return HRESULT(ret)
}

// 00000012-0000-0000-C000-000000000046
var IID_IRootStorage = syscall.GUID{0x00000012, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IRootStorageInterface interface {
	IUnknownInterface
	SwitchToFile(pszFile PWSTR) HRESULT
}

type IRootStorageVtbl struct {
	IUnknownVtbl
	SwitchToFile uintptr
}

type IRootStorage struct {
	IUnknown
}

func (this *IRootStorage) Vtbl() *IRootStorageVtbl {
	return (*IRootStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRootStorage) SwitchToFile(pszFile PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SwitchToFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszFile)))
	return HRESULT(ret)
}

// 99CAF010-415E-11CF-8814-00AA00B569F5
var IID_IFillLockBytes = syscall.GUID{0x99CAF010, 0x415E, 0x11CF,
	[8]byte{0x88, 0x14, 0x00, 0xAA, 0x00, 0xB5, 0x69, 0xF5}}

type IFillLockBytesInterface interface {
	IUnknownInterface
	FillAppend(pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT
	FillAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT
	SetFillSize(ulSize uint64) HRESULT
	Terminate(bCanceled BOOL) HRESULT
}

type IFillLockBytesVtbl struct {
	IUnknownVtbl
	FillAppend  uintptr
	FillAt      uintptr
	SetFillSize uintptr
	Terminate   uintptr
}

type IFillLockBytes struct {
	IUnknown
}

func (this *IFillLockBytes) Vtbl() *IFillLockBytesVtbl {
	return (*IFillLockBytesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFillLockBytes) FillAppend(pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().FillAppend, uintptr(unsafe.Pointer(this)), uintptr(pv), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

func (this *IFillLockBytes) FillAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().FillAt, uintptr(unsafe.Pointer(this)), uintptr(ulOffset), uintptr(pv), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

func (this *IFillLockBytes) SetFillSize(ulSize uint64) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillSize, uintptr(unsafe.Pointer(this)), uintptr(ulSize))
	return HRESULT(ret)
}

func (this *IFillLockBytes) Terminate(bCanceled BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Terminate, uintptr(unsafe.Pointer(this)), uintptr(bCanceled))
	return HRESULT(ret)
}

// 0E6D4D90-6738-11CF-9608-00AA00680DB4
var IID_ILayoutStorage = syscall.GUID{0x0E6D4D90, 0x6738, 0x11CF,
	[8]byte{0x96, 0x08, 0x00, 0xAA, 0x00, 0x68, 0x0D, 0xB4}}

type ILayoutStorageInterface interface {
	IUnknownInterface
	LayoutScript(pStorageLayout *StorageLayout, nEntries uint32, glfInterleavedFlag uint32) HRESULT
	BeginMonitor() HRESULT
	EndMonitor() HRESULT
	ReLayoutDocfile(pwcsNewDfName PWSTR) HRESULT
	ReLayoutDocfileOnILockBytes(pILockBytes *ILockBytes) HRESULT
}

type ILayoutStorageVtbl struct {
	IUnknownVtbl
	LayoutScript                uintptr
	BeginMonitor                uintptr
	EndMonitor                  uintptr
	ReLayoutDocfile             uintptr
	ReLayoutDocfileOnILockBytes uintptr
}

type ILayoutStorage struct {
	IUnknown
}

func (this *ILayoutStorage) Vtbl() *ILayoutStorageVtbl {
	return (*ILayoutStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILayoutStorage) LayoutScript(pStorageLayout *StorageLayout, nEntries uint32, glfInterleavedFlag uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().LayoutScript, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStorageLayout)), uintptr(nEntries), uintptr(glfInterleavedFlag))
	return HRESULT(ret)
}

func (this *ILayoutStorage) BeginMonitor() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().BeginMonitor, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILayoutStorage) EndMonitor() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().EndMonitor, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILayoutStorage) ReLayoutDocfile(pwcsNewDfName PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReLayoutDocfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsNewDfName)))
	return HRESULT(ret)
}

func (this *ILayoutStorage) ReLayoutDocfileOnILockBytes(pILockBytes *ILockBytes) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReLayoutDocfileOnILockBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pILockBytes)))
	return HRESULT(ret)
}

// 0E6D4D92-6738-11CF-9608-00AA00680DB4
var IID_IDirectWriterLock = syscall.GUID{0x0E6D4D92, 0x6738, 0x11CF,
	[8]byte{0x96, 0x08, 0x00, 0xAA, 0x00, 0x68, 0x0D, 0xB4}}

type IDirectWriterLockInterface interface {
	IUnknownInterface
	WaitForWriteAccess(dwTimeout uint32) HRESULT
	ReleaseWriteAccess() HRESULT
	HaveWriteAccess() HRESULT
}

type IDirectWriterLockVtbl struct {
	IUnknownVtbl
	WaitForWriteAccess uintptr
	ReleaseWriteAccess uintptr
	HaveWriteAccess    uintptr
}

type IDirectWriterLock struct {
	IUnknown
}

func (this *IDirectWriterLock) Vtbl() *IDirectWriterLockVtbl {
	return (*IDirectWriterLockVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDirectWriterLock) WaitForWriteAccess(dwTimeout uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WaitForWriteAccess, uintptr(unsafe.Pointer(this)), uintptr(dwTimeout))
	return HRESULT(ret)
}

func (this *IDirectWriterLock) ReleaseWriteAccess() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReleaseWriteAccess, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IDirectWriterLock) HaveWriteAccess() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().HaveWriteAccess, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 00000138-0000-0000-C000-000000000046
var IID_IPropertyStorage = syscall.GUID{0x00000138, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IPropertyStorageInterface interface {
	IUnknownInterface
	ReadMultiple(cpspec uint32, rgpspec *PROPSPEC, rgpropvar *PROPVARIANT) HRESULT
	WriteMultiple(cpspec uint32, rgpspec *PROPSPEC, rgpropvar *PROPVARIANT, propidNameFirst uint32) HRESULT
	DeleteMultiple(cpspec uint32, rgpspec *PROPSPEC) HRESULT
	ReadPropertyNames(cpropid uint32, rgpropid *uint32, rglpwstrName *PWSTR) HRESULT
	WritePropertyNames(cpropid uint32, rgpropid *uint32, rglpwstrName *PWSTR) HRESULT
	DeletePropertyNames(cpropid uint32, rgpropid *uint32) HRESULT
	Commit(grfCommitFlags uint32) HRESULT
	Revert() HRESULT
	Enum(ppenum **IEnumSTATPROPSTG) HRESULT
	SetTimes(pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT
	SetClass(clsid *syscall.GUID) HRESULT
	Stat(pstatpsstg *STATPROPSETSTG) HRESULT
}

type IPropertyStorageVtbl struct {
	IUnknownVtbl
	ReadMultiple        uintptr
	WriteMultiple       uintptr
	DeleteMultiple      uintptr
	ReadPropertyNames   uintptr
	WritePropertyNames  uintptr
	DeletePropertyNames uintptr
	Commit              uintptr
	Revert              uintptr
	Enum                uintptr
	SetTimes            uintptr
	SetClass            uintptr
	Stat                uintptr
}

type IPropertyStorage struct {
	IUnknown
}

func (this *IPropertyStorage) Vtbl() *IPropertyStorageVtbl {
	return (*IPropertyStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyStorage) ReadMultiple(cpspec uint32, rgpspec *PROPSPEC, rgpropvar *PROPVARIANT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReadMultiple, uintptr(unsafe.Pointer(this)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)), uintptr(unsafe.Pointer(rgpropvar)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) WriteMultiple(cpspec uint32, rgpspec *PROPSPEC, rgpropvar *PROPVARIANT, propidNameFirst uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteMultiple, uintptr(unsafe.Pointer(this)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)), uintptr(unsafe.Pointer(rgpropvar)), uintptr(propidNameFirst))
	return HRESULT(ret)
}

func (this *IPropertyStorage) DeleteMultiple(cpspec uint32, rgpspec *PROPSPEC) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().DeleteMultiple, uintptr(unsafe.Pointer(this)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) ReadPropertyNames(cpropid uint32, rgpropid *uint32, rglpwstrName *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReadPropertyNames, uintptr(unsafe.Pointer(this)), uintptr(cpropid), uintptr(unsafe.Pointer(rgpropid)), uintptr(unsafe.Pointer(rglpwstrName)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) WritePropertyNames(cpropid uint32, rgpropid *uint32, rglpwstrName *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WritePropertyNames, uintptr(unsafe.Pointer(this)), uintptr(cpropid), uintptr(unsafe.Pointer(rgpropid)), uintptr(unsafe.Pointer(rglpwstrName)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) DeletePropertyNames(cpropid uint32, rgpropid *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().DeletePropertyNames, uintptr(unsafe.Pointer(this)), uintptr(cpropid), uintptr(unsafe.Pointer(rgpropid)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Commit(grfCommitFlags uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Commit, uintptr(unsafe.Pointer(this)), uintptr(grfCommitFlags))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Revert() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Revert, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Enum(ppenum **IEnumSTATPROPSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Enum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) SetTimes(pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTimes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pctime)), uintptr(unsafe.Pointer(patime)), uintptr(unsafe.Pointer(pmtime)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) SetClass(clsid *syscall.GUID) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clsid)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Stat(pstatpsstg *STATPROPSETSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Stat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstatpsstg)))
	return HRESULT(ret)
}

// 0000013A-0000-0000-C000-000000000046
var IID_IPropertySetStorage = syscall.GUID{0x0000013A, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IPropertySetStorageInterface interface {
	IUnknownInterface
	Create(rfmtid *syscall.GUID, pclsid *syscall.GUID, grfFlags uint32, grfMode uint32, ppprstg **IPropertyStorage) HRESULT
	Open(rfmtid *syscall.GUID, grfMode uint32, ppprstg **IPropertyStorage) HRESULT
	Delete(rfmtid *syscall.GUID) HRESULT
	Enum(ppenum **IEnumSTATPROPSETSTG) HRESULT
}

type IPropertySetStorageVtbl struct {
	IUnknownVtbl
	Create uintptr
	Open   uintptr
	Delete uintptr
	Enum   uintptr
}

type IPropertySetStorage struct {
	IUnknown
}

func (this *IPropertySetStorage) Vtbl() *IPropertySetStorageVtbl {
	return (*IPropertySetStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertySetStorage) Create(rfmtid *syscall.GUID, pclsid *syscall.GUID, grfFlags uint32, grfMode uint32, ppprstg **IPropertyStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rfmtid)), uintptr(unsafe.Pointer(pclsid)), uintptr(grfFlags), uintptr(grfMode), uintptr(unsafe.Pointer(ppprstg)))
	return HRESULT(ret)
}

func (this *IPropertySetStorage) Open(rfmtid *syscall.GUID, grfMode uint32, ppprstg **IPropertyStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Open, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rfmtid)), uintptr(grfMode), uintptr(unsafe.Pointer(ppprstg)))
	return HRESULT(ret)
}

func (this *IPropertySetStorage) Delete(rfmtid *syscall.GUID) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Delete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rfmtid)))
	return HRESULT(ret)
}

func (this *IPropertySetStorage) Enum(ppenum **IEnumSTATPROPSETSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Enum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 00000139-0000-0000-C000-000000000046
var IID_IEnumSTATPROPSTG = syscall.GUID{0x00000139, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumSTATPROPSTGInterface interface {
	IUnknownInterface
	Next(celt uint32, rgelt *STATPROPSTG, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppenum **IEnumSTATPROPSTG) HRESULT
}

type IEnumSTATPROPSTGVtbl struct {
	IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumSTATPROPSTG struct {
	IUnknown
}

func (this *IEnumSTATPROPSTG) Vtbl() *IEnumSTATPROPSTGVtbl {
	return (*IEnumSTATPROPSTGVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumSTATPROPSTG) Next(celt uint32, rgelt *STATPROPSTG, pceltFetched *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgelt)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSTG) Skip(celt uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSTG) Reset() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSTG) Clone(ppenum **IEnumSTATPROPSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 0000013B-0000-0000-C000-000000000046
var IID_IEnumSTATPROPSETSTG = syscall.GUID{0x0000013B, 0x0000, 0x0000,
	[8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumSTATPROPSETSTGInterface interface {
	IUnknownInterface
	Next(celt uint32, rgelt *STATPROPSETSTG, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppenum **IEnumSTATPROPSETSTG) HRESULT
}

type IEnumSTATPROPSETSTGVtbl struct {
	IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumSTATPROPSETSTG struct {
	IUnknown
}

func (this *IEnumSTATPROPSETSTG) Vtbl() *IEnumSTATPROPSETSTGVtbl {
	return (*IEnumSTATPROPSETSTGVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumSTATPROPSETSTG) Next(celt uint32, rgelt *STATPROPSETSTG, pceltFetched *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgelt)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSETSTG) Skip(celt uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSETSTG) Reset() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSETSTG) Clone(ppenum **IEnumSTATPROPSETSTG) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 00000000-0000-0000-0000-000000000000
var IID_IMemoryAllocator = syscall.GUID{0x00000000, 0x0000, 0x0000,
	[8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}

type IMemoryAllocatorInterface interface {
	Allocate(cbSize uint32) unsafe.Pointer
	Free(pv unsafe.Pointer)
}

type IMemoryAllocatorVtbl struct {
	Allocate uintptr
	Free     uintptr
}

type IMemoryAllocator struct {
	LpVtbl *[1024]uintptr
}

func (this *IMemoryAllocator) Vtbl() *IMemoryAllocatorVtbl {
	return (*IMemoryAllocatorVtbl)(unsafe.Pointer(this.LpVtbl))
}

func (this *IMemoryAllocator) Allocate(cbSize uint32) unsafe.Pointer {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Allocate, uintptr(unsafe.Pointer(this)), uintptr(cbSize))
	return (unsafe.Pointer)(ret)
}

func (this *IMemoryAllocator) Free(pv unsafe.Pointer) {
	_, _, _ = syscall.SyscallN(this.Vtbl().Free, uintptr(unsafe.Pointer(this)), uintptr(pv))
}

// 55272A00-42CB-11CE-8135-00AA004BB851
var IID_IPropertyBag = syscall.GUID{0x55272A00, 0x42CB, 0x11CE,
	[8]byte{0x81, 0x35, 0x00, 0xAA, 0x00, 0x4B, 0xB8, 0x51}}

type IPropertyBagInterface interface {
	IUnknownInterface
	Read(pszPropName PWSTR, pVar *VARIANT, pErrorLog *IErrorLog) HRESULT
	Write(pszPropName PWSTR, pVar *VARIANT) HRESULT
}

type IPropertyBagVtbl struct {
	IUnknownVtbl
	Read  uintptr
	Write uintptr
}

type IPropertyBag struct {
	IUnknown
}

func (this *IPropertyBag) Vtbl() *IPropertyBagVtbl {
	return (*IPropertyBagVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyBag) Read(pszPropName PWSTR, pVar *VARIANT, pErrorLog *IErrorLog) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPropName)), uintptr(unsafe.Pointer(pVar)), uintptr(unsafe.Pointer(pErrorLog)))
	return HRESULT(ret)
}

func (this *IPropertyBag) Write(pszPropName PWSTR, pVar *VARIANT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPropName)), uintptr(unsafe.Pointer(pVar)))
	return HRESULT(ret)
}

// 22F55882-280B-11D0-A8A9-00A0C90C2004
var IID_IPropertyBag2 = syscall.GUID{0x22F55882, 0x280B, 0x11D0,
	[8]byte{0xA8, 0xA9, 0x00, 0xA0, 0xC9, 0x0C, 0x20, 0x04}}

type IPropertyBag2Interface interface {
	IUnknownInterface
	Read(cProperties uint32, pPropBag *PROPBAG2, pErrLog *IErrorLog, pvarValue *VARIANT, phrError *HRESULT) HRESULT
	Write(cProperties uint32, pPropBag *PROPBAG2, pvarValue *VARIANT) HRESULT
	CountProperties(pcProperties *uint32) HRESULT
	GetPropertyInfo(iProperty uint32, cProperties uint32, pPropBag *PROPBAG2, pcProperties *uint32) HRESULT
	LoadObject(pstrName PWSTR, dwHint uint32, pUnkObject *IUnknown, pErrLog *IErrorLog) HRESULT
}

type IPropertyBag2Vtbl struct {
	IUnknownVtbl
	Read            uintptr
	Write           uintptr
	CountProperties uintptr
	GetPropertyInfo uintptr
	LoadObject      uintptr
}

type IPropertyBag2 struct {
	IUnknown
}

func (this *IPropertyBag2) Vtbl() *IPropertyBag2Vtbl {
	return (*IPropertyBag2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyBag2) Read(cProperties uint32, pPropBag *PROPBAG2, pErrLog *IErrorLog, pvarValue *VARIANT, phrError *HRESULT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(cProperties), uintptr(unsafe.Pointer(pPropBag)), uintptr(unsafe.Pointer(pErrLog)), uintptr(unsafe.Pointer(pvarValue)), uintptr(unsafe.Pointer(phrError)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) Write(cProperties uint32, pPropBag *PROPBAG2, pvarValue *VARIANT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(cProperties), uintptr(unsafe.Pointer(pPropBag)), uintptr(unsafe.Pointer(pvarValue)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) CountProperties(pcProperties *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CountProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcProperties)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) GetPropertyInfo(iProperty uint32, cProperties uint32, pPropBag *PROPBAG2, pcProperties *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyInfo, uintptr(unsafe.Pointer(this)), uintptr(iProperty), uintptr(cProperties), uintptr(unsafe.Pointer(pPropBag)), uintptr(unsafe.Pointer(pcProperties)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) LoadObject(pstrName PWSTR, dwHint uint32, pUnkObject *IUnknown, pErrLog *IErrorLog) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().LoadObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstrName)), uintptr(dwHint), uintptr(unsafe.Pointer(pUnkObject)), uintptr(unsafe.Pointer(pErrLog)))
	return HRESULT(ret)
}

var (
	pCoGetInstanceFromFile                    uintptr
	pCoGetInstanceFromIStorage                uintptr
	pStgOpenAsyncDocfileOnIFillLockBytes      uintptr
	pStgGetIFillLockBytesOnILockBytes         uintptr
	pStgGetIFillLockBytesOnFile               uintptr
	pCreateStreamOnHGlobal                    uintptr
	pGetHGlobalFromStream                     uintptr
	pCoGetInterfaceAndReleaseStream           uintptr
	pPropVariantCopy                          uintptr
	pPropVariantClear                         uintptr
	pFreePropVariantArray                     uintptr
	pStgCreateDocfile                         uintptr
	pStgCreateDocfileOnILockBytes             uintptr
	pStgOpenStorage                           uintptr
	pStgOpenStorageOnILockBytes               uintptr
	pStgIsStorageFile                         uintptr
	pStgIsStorageILockBytes                   uintptr
	pStgSetTimes                              uintptr
	pStgCreateStorageEx                       uintptr
	pStgOpenStorageEx                         uintptr
	pStgCreatePropStg                         uintptr
	pStgOpenPropStg                           uintptr
	pStgCreatePropSetStg                      uintptr
	pFmtIdToPropStgName                       uintptr
	pPropStgNameToFmtId                       uintptr
	pReadClassStg                             uintptr
	pWriteClassStg                            uintptr
	pReadClassStm                             uintptr
	pWriteClassStm                            uintptr
	pGetHGlobalFromILockBytes                 uintptr
	pCreateILockBytesOnHGlobal                uintptr
	pGetConvertStg                            uintptr
	pStgConvertVariantToProperty              uintptr
	pStgConvertPropertyToVariant              uintptr
	pStgPropertyLengthAsVariant               uintptr
	pWriteFmtUserTypeStg                      uintptr
	pReadFmtUserTypeStg                       uintptr
	pOleConvertOLESTREAMToIStorage            uintptr
	pOleConvertIStorageToOLESTREAM            uintptr
	pSetConvertStg                            uintptr
	pOleConvertIStorageToOLESTREAMEx          uintptr
	pOleConvertOLESTREAMToIStorageEx          uintptr
	pPropVariantToWinRTPropertyValue          uintptr
	pWinRTPropertyValueToPropVariant          uintptr
	pInitPropVariantFromResource              uintptr
	pInitPropVariantFromBuffer                uintptr
	pInitPropVariantFromCLSID                 uintptr
	pInitPropVariantFromGUIDAsString          uintptr
	pInitPropVariantFromFileTime              uintptr
	pInitPropVariantFromPropVariantVectorElem uintptr
	pInitPropVariantVectorFromPropVariant     uintptr
	pInitPropVariantFromBooleanVector         uintptr
	pInitPropVariantFromInt16Vector           uintptr
	pInitPropVariantFromUInt16Vector          uintptr
	pInitPropVariantFromInt32Vector           uintptr
	pInitPropVariantFromUInt32Vector          uintptr
	pInitPropVariantFromInt64Vector           uintptr
	pInitPropVariantFromUInt64Vector          uintptr
	pInitPropVariantFromDoubleVector          uintptr
	pInitPropVariantFromFileTimeVector        uintptr
	pInitPropVariantFromStringVector          uintptr
	pInitPropVariantFromStringAsVector        uintptr
	pPropVariantToBooleanWithDefault          uintptr
	pPropVariantToInt16WithDefault            uintptr
	pPropVariantToUInt16WithDefault           uintptr
	pPropVariantToInt32WithDefault            uintptr
	pPropVariantToUInt32WithDefault           uintptr
	pPropVariantToInt64WithDefault            uintptr
	pPropVariantToUInt64WithDefault           uintptr
	pPropVariantToDoubleWithDefault           uintptr
	pPropVariantToStringWithDefault           uintptr
	pPropVariantToBoolean                     uintptr
	pPropVariantToInt16                       uintptr
	pPropVariantToUInt16                      uintptr
	pPropVariantToInt32                       uintptr
	pPropVariantToUInt32                      uintptr
	pPropVariantToInt64                       uintptr
	pPropVariantToUInt64                      uintptr
	pPropVariantToDouble                      uintptr
	pPropVariantToBuffer                      uintptr
	pPropVariantToString                      uintptr
	pPropVariantToGUID                        uintptr
	pPropVariantToStringAlloc                 uintptr
	pPropVariantToBSTR                        uintptr
	pPropVariantToFileTime                    uintptr
	pPropVariantGetElementCount               uintptr
	pPropVariantToBooleanVector               uintptr
	pPropVariantToInt16Vector                 uintptr
	pPropVariantToUInt16Vector                uintptr
	pPropVariantToInt32Vector                 uintptr
	pPropVariantToUInt32Vector                uintptr
	pPropVariantToInt64Vector                 uintptr
	pPropVariantToUInt64Vector                uintptr
	pPropVariantToDoubleVector                uintptr
	pPropVariantToFileTimeVector              uintptr
	pPropVariantToStringVector                uintptr
	pPropVariantToBooleanVectorAlloc          uintptr
	pPropVariantToInt16VectorAlloc            uintptr
	pPropVariantToUInt16VectorAlloc           uintptr
	pPropVariantToInt32VectorAlloc            uintptr
	pPropVariantToUInt32VectorAlloc           uintptr
	pPropVariantToInt64VectorAlloc            uintptr
	pPropVariantToUInt64VectorAlloc           uintptr
	pPropVariantToDoubleVectorAlloc           uintptr
	pPropVariantToFileTimeVectorAlloc         uintptr
	pPropVariantToStringVectorAlloc           uintptr
	pPropVariantGetBooleanElem                uintptr
	pPropVariantGetInt16Elem                  uintptr
	pPropVariantGetUInt16Elem                 uintptr
	pPropVariantGetInt32Elem                  uintptr
	pPropVariantGetUInt32Elem                 uintptr
	pPropVariantGetInt64Elem                  uintptr
	pPropVariantGetUInt64Elem                 uintptr
	pPropVariantGetDoubleElem                 uintptr
	pPropVariantGetFileTimeElem               uintptr
	pPropVariantGetStringElem                 uintptr
	pClearPropVariantArray                    uintptr
	pPropVariantCompareEx                     uintptr
	pPropVariantChangeType                    uintptr
	pPropVariantToVariant                     uintptr
	pVariantToPropVariant                     uintptr
	pStgSerializePropVariant                  uintptr
	pStgDeserializePropVariant                uintptr
)

func CoGetInstanceFromFile(pServerInfo *COSERVERINFO, pClsid *syscall.GUID, punkOuter *IUnknown, dwClsCtx CLSCTX, grfMode uint32, pwszName PWSTR, dwCount uint32, pResults *MULTI_QI) HRESULT {
	addr := LazyAddr(&pCoGetInstanceFromFile, libOle32, "CoGetInstanceFromFile")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pServerInfo)), uintptr(unsafe.Pointer(pClsid)), uintptr(unsafe.Pointer(punkOuter)), uintptr(dwClsCtx), uintptr(grfMode), uintptr(unsafe.Pointer(pwszName)), uintptr(dwCount), uintptr(unsafe.Pointer(pResults)))
	return HRESULT(ret)
}

func CoGetInstanceFromIStorage(pServerInfo *COSERVERINFO, pClsid *syscall.GUID, punkOuter *IUnknown, dwClsCtx CLSCTX, pstg *IStorage, dwCount uint32, pResults *MULTI_QI) HRESULT {
	addr := LazyAddr(&pCoGetInstanceFromIStorage, libOle32, "CoGetInstanceFromIStorage")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pServerInfo)), uintptr(unsafe.Pointer(pClsid)), uintptr(unsafe.Pointer(punkOuter)), uintptr(dwClsCtx), uintptr(unsafe.Pointer(pstg)), uintptr(dwCount), uintptr(unsafe.Pointer(pResults)))
	return HRESULT(ret)
}

func StgOpenAsyncDocfileOnIFillLockBytes(pflb *IFillLockBytes, grfMode uint32, asyncFlags uint32, ppstgOpen **IStorage) HRESULT {
	addr := LazyAddr(&pStgOpenAsyncDocfileOnIFillLockBytes, libOle32, "StgOpenAsyncDocfileOnIFillLockBytes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pflb)), uintptr(grfMode), uintptr(asyncFlags), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgGetIFillLockBytesOnILockBytes(pilb *ILockBytes, ppflb **IFillLockBytes) HRESULT {
	addr := LazyAddr(&pStgGetIFillLockBytesOnILockBytes, libOle32, "StgGetIFillLockBytesOnILockBytes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pilb)), uintptr(unsafe.Pointer(ppflb)))
	return HRESULT(ret)
}

func StgGetIFillLockBytesOnFile(pwcsName PWSTR, ppflb **IFillLockBytes) HRESULT {
	addr := LazyAddr(&pStgGetIFillLockBytesOnFile, libOle32, "StgGetIFillLockBytesOnFile")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(ppflb)))
	return HRESULT(ret)
}

func CreateStreamOnHGlobal(hGlobal HGLOBAL, fDeleteOnRelease BOOL, ppstm **IStream) HRESULT {
	addr := LazyAddr(&pCreateStreamOnHGlobal, libOle32, "CreateStreamOnHGlobal")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(hGlobal)), uintptr(fDeleteOnRelease), uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret)
}

func GetHGlobalFromStream(pstm *IStream, phglobal *HGLOBAL) HRESULT {
	addr := LazyAddr(&pGetHGlobalFromStream, libOle32, "GetHGlobalFromStream")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstm)), uintptr(unsafe.Pointer(phglobal)))
	return HRESULT(ret)
}

func CoGetInterfaceAndReleaseStream(pStm *IStream, iid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pCoGetInterfaceAndReleaseStream, libOle32, "CoGetInterfaceAndReleaseStream")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStm)), uintptr(unsafe.Pointer(iid)), uintptr(ppv))
	return HRESULT(ret)
}

func PropVariantCopy(pvarDest *PROPVARIANT, pvarSrc *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPropVariantCopy, libOle32, "PropVariantCopy")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvarDest)), uintptr(unsafe.Pointer(pvarSrc)))
	return HRESULT(ret)
}

func PropVariantClear(pvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPropVariantClear, libOle32, "PropVariantClear")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func FreePropVariantArray(cVariants uint32, rgvars *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pFreePropVariantArray, libOle32, "FreePropVariantArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(cVariants), uintptr(unsafe.Pointer(rgvars)))
	return HRESULT(ret)
}

func StgCreateDocfile(pwcsName PWSTR, grfMode STGM, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := LazyAddr(&pStgCreateDocfile, libOle32, "StgCreateDocfile")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgCreateDocfileOnILockBytes(plkbyt *ILockBytes, grfMode STGM, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := LazyAddr(&pStgCreateDocfileOnILockBytes, libOle32, "StgCreateDocfileOnILockBytes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)), uintptr(grfMode), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgOpenStorage(pwcsName PWSTR, pstgPriority *IStorage, grfMode STGM, snbExclude **uint16, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := LazyAddr(&pStgOpenStorage, libOle32, "StgOpenStorage")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pstgPriority)), uintptr(grfMode), uintptr(unsafe.Pointer(snbExclude)), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgOpenStorageOnILockBytes(plkbyt *ILockBytes, pstgPriority *IStorage, grfMode STGM, snbExclude **uint16, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := LazyAddr(&pStgOpenStorageOnILockBytes, libOle32, "StgOpenStorageOnILockBytes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)), uintptr(unsafe.Pointer(pstgPriority)), uintptr(grfMode), uintptr(unsafe.Pointer(snbExclude)), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgIsStorageFile(pwcsName PWSTR) HRESULT {
	addr := LazyAddr(&pStgIsStorageFile, libOle32, "StgIsStorageFile")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)))
	return HRESULT(ret)
}

func StgIsStorageILockBytes(plkbyt *ILockBytes) HRESULT {
	addr := LazyAddr(&pStgIsStorageILockBytes, libOle32, "StgIsStorageILockBytes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)))
	return HRESULT(ret)
}

func StgSetTimes(lpszName PWSTR, pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT {
	addr := LazyAddr(&pStgSetTimes, libOle32, "StgSetTimes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszName)), uintptr(unsafe.Pointer(pctime)), uintptr(unsafe.Pointer(patime)), uintptr(unsafe.Pointer(pmtime)))
	return HRESULT(ret)
}

func StgCreateStorageEx(pwcsName PWSTR, grfMode STGM, stgfmt STGFMT, grfAttrs uint32, pStgOptions *STGOPTIONS, pSecurityDescriptor PSECURITY_DESCRIPTOR, riid *syscall.GUID, ppObjectOpen unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pStgCreateStorageEx, libOle32, "StgCreateStorageEx")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(stgfmt), uintptr(grfAttrs), uintptr(unsafe.Pointer(pStgOptions)), uintptr(unsafe.Pointer(pSecurityDescriptor)), uintptr(unsafe.Pointer(riid)), uintptr(ppObjectOpen))
	return HRESULT(ret)
}

func StgOpenStorageEx(pwcsName PWSTR, grfMode STGM, stgfmt STGFMT, grfAttrs uint32, pStgOptions *STGOPTIONS, pSecurityDescriptor PSECURITY_DESCRIPTOR, riid *syscall.GUID, ppObjectOpen unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pStgOpenStorageEx, libOle32, "StgOpenStorageEx")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(stgfmt), uintptr(grfAttrs), uintptr(unsafe.Pointer(pStgOptions)), uintptr(unsafe.Pointer(pSecurityDescriptor)), uintptr(unsafe.Pointer(riid)), uintptr(ppObjectOpen))
	return HRESULT(ret)
}

func StgCreatePropStg(pUnk *IUnknown, fmtid *syscall.GUID, pclsid *syscall.GUID, grfFlags uint32, dwReserved uint32, ppPropStg **IPropertyStorage) HRESULT {
	addr := LazyAddr(&pStgCreatePropStg, libOle32, "StgCreatePropStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pUnk)), uintptr(unsafe.Pointer(fmtid)), uintptr(unsafe.Pointer(pclsid)), uintptr(grfFlags), uintptr(dwReserved), uintptr(unsafe.Pointer(ppPropStg)))
	return HRESULT(ret)
}

func StgOpenPropStg(pUnk *IUnknown, fmtid *syscall.GUID, grfFlags uint32, dwReserved uint32, ppPropStg **IPropertyStorage) HRESULT {
	addr := LazyAddr(&pStgOpenPropStg, libOle32, "StgOpenPropStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pUnk)), uintptr(unsafe.Pointer(fmtid)), uintptr(grfFlags), uintptr(dwReserved), uintptr(unsafe.Pointer(ppPropStg)))
	return HRESULT(ret)
}

func StgCreatePropSetStg(pStorage *IStorage, dwReserved uint32, ppPropSetStg **IPropertySetStorage) HRESULT {
	addr := LazyAddr(&pStgCreatePropSetStg, libOle32, "StgCreatePropSetStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStorage)), uintptr(dwReserved), uintptr(unsafe.Pointer(ppPropSetStg)))
	return HRESULT(ret)
}

func FmtIdToPropStgName(pfmtid *syscall.GUID, oszName PWSTR) HRESULT {
	addr := LazyAddr(&pFmtIdToPropStgName, libOle32, "FmtIdToPropStgName")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pfmtid)), uintptr(unsafe.Pointer(oszName)))
	return HRESULT(ret)
}

func PropStgNameToFmtId(oszName PWSTR, pfmtid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pPropStgNameToFmtId, libOle32, "PropStgNameToFmtId")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(oszName)), uintptr(unsafe.Pointer(pfmtid)))
	return HRESULT(ret)
}

func ReadClassStg(pStg *IStorage, pclsid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pReadClassStg, libOle32, "ReadClassStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)), uintptr(unsafe.Pointer(pclsid)))
	return HRESULT(ret)
}

func WriteClassStg(pStg *IStorage, rclsid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pWriteClassStg, libOle32, "WriteClassStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)), uintptr(unsafe.Pointer(rclsid)))
	return HRESULT(ret)
}

func ReadClassStm(pStm *IStream, pclsid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pReadClassStm, libOle32, "ReadClassStm")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStm)), uintptr(unsafe.Pointer(pclsid)))
	return HRESULT(ret)
}

func WriteClassStm(pStm *IStream, rclsid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pWriteClassStm, libOle32, "WriteClassStm")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStm)), uintptr(unsafe.Pointer(rclsid)))
	return HRESULT(ret)
}

func GetHGlobalFromILockBytes(plkbyt *ILockBytes, phglobal *HGLOBAL) HRESULT {
	addr := LazyAddr(&pGetHGlobalFromILockBytes, libOle32, "GetHGlobalFromILockBytes")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)), uintptr(unsafe.Pointer(phglobal)))
	return HRESULT(ret)
}

func CreateILockBytesOnHGlobal(hGlobal HGLOBAL, fDeleteOnRelease BOOL, pplkbyt **ILockBytes) HRESULT {
	addr := LazyAddr(&pCreateILockBytesOnHGlobal, libOle32, "CreateILockBytesOnHGlobal")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(hGlobal)), uintptr(fDeleteOnRelease), uintptr(unsafe.Pointer(pplkbyt)))
	return HRESULT(ret)
}

func GetConvertStg(pStg *IStorage) HRESULT {
	addr := LazyAddr(&pGetConvertStg, libOle32, "GetConvertStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)))
	return HRESULT(ret)
}

func StgConvertVariantToProperty(pvar *PROPVARIANT, CodePage uint16, pprop *SERIALIZEDPROPERTYVALUE, pcb *uint32, pid uint32, fReserved BOOLEAN, pcIndirect *uint32) *SERIALIZEDPROPERTYVALUE {
	addr := LazyAddr(&pStgConvertVariantToProperty, libOle32, "StgConvertVariantToProperty")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvar)), uintptr(CodePage), uintptr(unsafe.Pointer(pprop)), uintptr(unsafe.Pointer(pcb)), uintptr(pid), uintptr(fReserved), uintptr(unsafe.Pointer(pcIndirect)))
	return (*SERIALIZEDPROPERTYVALUE)(unsafe.Pointer(ret))
}

func StgConvertPropertyToVariant(pprop *SERIALIZEDPROPERTYVALUE, CodePage uint16, pvar *PROPVARIANT, pma *IMemoryAllocator) BOOLEAN {
	addr := LazyAddr(&pStgConvertPropertyToVariant, libOle32, "StgConvertPropertyToVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pprop)), uintptr(CodePage), uintptr(unsafe.Pointer(pvar)), uintptr(unsafe.Pointer(pma)))
	return BOOLEAN(ret)
}

func StgPropertyLengthAsVariant(pProp *SERIALIZEDPROPERTYVALUE, cbProp uint32, CodePage uint16, bReserved byte) uint32 {
	addr := LazyAddr(&pStgPropertyLengthAsVariant, libOle32, "StgPropertyLengthAsVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pProp)), uintptr(cbProp), uintptr(CodePage), uintptr(bReserved))
	return uint32(ret)
}

func WriteFmtUserTypeStg(pstg *IStorage, cf uint16, lpszUserType PWSTR) HRESULT {
	addr := LazyAddr(&pWriteFmtUserTypeStg, libOle32, "WriteFmtUserTypeStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(cf), uintptr(unsafe.Pointer(lpszUserType)))
	return HRESULT(ret)
}

func ReadFmtUserTypeStg(pstg *IStorage, pcf *uint16, lplpszUserType *PWSTR) HRESULT {
	addr := LazyAddr(&pReadFmtUserTypeStg, libOle32, "ReadFmtUserTypeStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(pcf)), uintptr(unsafe.Pointer(lplpszUserType)))
	return HRESULT(ret)
}

func OleConvertOLESTREAMToIStorage(lpolestream *OLESTREAM, pstg *IStorage, ptd *DVTARGETDEVICE) HRESULT {
	addr := LazyAddr(&pOleConvertOLESTREAMToIStorage, libOle32, "OleConvertOLESTREAMToIStorage")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpolestream)), uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(ptd)))
	return HRESULT(ret)
}

func OleConvertIStorageToOLESTREAM(pstg *IStorage, lpolestream *OLESTREAM) HRESULT {
	addr := LazyAddr(&pOleConvertIStorageToOLESTREAM, libOle32, "OleConvertIStorageToOLESTREAM")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(lpolestream)))
	return HRESULT(ret)
}

func SetConvertStg(pStg *IStorage, fConvert BOOL) HRESULT {
	addr := LazyAddr(&pSetConvertStg, libOle32, "SetConvertStg")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)), uintptr(fConvert))
	return HRESULT(ret)
}

func OleConvertIStorageToOLESTREAMEx(pstg *IStorage, cfFormat uint16, lWidth int32, lHeight int32, dwSize uint32, pmedium *STGMEDIUM, polestm *OLESTREAM) HRESULT {
	addr := LazyAddr(&pOleConvertIStorageToOLESTREAMEx, libOle32, "OleConvertIStorageToOLESTREAMEx")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(cfFormat), uintptr(lWidth), uintptr(lHeight), uintptr(dwSize), uintptr(unsafe.Pointer(pmedium)), uintptr(unsafe.Pointer(polestm)))
	return HRESULT(ret)
}

func OleConvertOLESTREAMToIStorageEx(polestm *OLESTREAM, pstg *IStorage, pcfFormat *uint16, plwWidth *int32, plHeight *int32, pdwSize *uint32, pmedium *STGMEDIUM) HRESULT {
	addr := LazyAddr(&pOleConvertOLESTREAMToIStorageEx, libOle32, "OleConvertOLESTREAMToIStorageEx")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(polestm)), uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(pcfFormat)), uintptr(unsafe.Pointer(plwWidth)), uintptr(unsafe.Pointer(plHeight)), uintptr(unsafe.Pointer(pdwSize)), uintptr(unsafe.Pointer(pmedium)))
	return HRESULT(ret)
}

func PropVariantToWinRTPropertyValue(propvar *PROPVARIANT, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPropVariantToWinRTPropertyValue, libPropsys, "PropVariantToWinRTPropertyValue")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func WinRTPropertyValueToPropVariant(punkPropertyValue *IUnknown, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pWinRTPropertyValueToPropVariant, libPropsys, "WinRTPropertyValueToPropVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(punkPropertyValue)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromResource(hinst HINSTANCE, id uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromResource, libPropsys, "InitPropVariantFromResource")
	ret, _, _ := syscall.SyscallN(addr, hinst, uintptr(id), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromBuffer(pv unsafe.Pointer, cb uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromBuffer, libPropsys, "InitPropVariantFromBuffer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(pv), uintptr(cb), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromCLSID(clsid *syscall.GUID, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromCLSID, libPropsys, "InitPropVariantFromCLSID")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(clsid)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromGUIDAsString(guid *syscall.GUID, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromGUIDAsString, libPropsys, "InitPropVariantFromGUIDAsString")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromFileTime(pftIn *FILETIME, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromFileTime, libPropsys, "InitPropVariantFromFileTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pftIn)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromPropVariantVectorElem(propvarIn *PROPVARIANT, iElem uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromPropVariantVectorElem, libPropsys, "InitPropVariantFromPropVariantVectorElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(iElem), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantVectorFromPropVariant(propvarSingle *PROPVARIANT, ppropvarVector *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantVectorFromPropVariant, libPropsys, "InitPropVariantVectorFromPropVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarSingle)), uintptr(unsafe.Pointer(ppropvarVector)))
	return HRESULT(ret)
}

func InitPropVariantFromBooleanVector(prgf *BOOL, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromBooleanVector, libPropsys, "InitPropVariantFromBooleanVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgf)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromInt16Vector(prgn *int16, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromInt16Vector, libPropsys, "InitPropVariantFromInt16Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromUInt16Vector(prgn *uint16, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromUInt16Vector, libPropsys, "InitPropVariantFromUInt16Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromInt32Vector(prgn *int32, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromInt32Vector, libPropsys, "InitPropVariantFromInt32Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromUInt32Vector(prgn *uint32, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromUInt32Vector, libPropsys, "InitPropVariantFromUInt32Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromInt64Vector(prgn *int64, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromInt64Vector, libPropsys, "InitPropVariantFromInt64Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromUInt64Vector(prgn *uint64, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromUInt64Vector, libPropsys, "InitPropVariantFromUInt64Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromDoubleVector(prgn *float64, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromDoubleVector, libPropsys, "InitPropVariantFromDoubleVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromFileTimeVector(prgft *FILETIME, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromFileTimeVector, libPropsys, "InitPropVariantFromFileTimeVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgft)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromStringVector(prgsz *PWSTR, cElems uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromStringVector, libPropsys, "InitPropVariantFromStringVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgsz)), uintptr(cElems), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func InitPropVariantFromStringAsVector(psz PWSTR, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromStringAsVector, libPropsys, "InitPropVariantFromStringAsVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(psz)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func PropVariantToBooleanWithDefault(propvarIn *PROPVARIANT, fDefault BOOL) BOOL {
	addr := LazyAddr(&pPropVariantToBooleanWithDefault, libPropsys, "PropVariantToBooleanWithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(fDefault))
	return BOOL(ret)
}

func PropVariantToInt16WithDefault(propvarIn *PROPVARIANT, iDefault int16) int16 {
	addr := LazyAddr(&pPropVariantToInt16WithDefault, libPropsys, "PropVariantToInt16WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(iDefault))
	return int16(ret)
}

func PropVariantToUInt16WithDefault(propvarIn *PROPVARIANT, uiDefault uint16) uint16 {
	addr := LazyAddr(&pPropVariantToUInt16WithDefault, libPropsys, "PropVariantToUInt16WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(uiDefault))
	return uint16(ret)
}

func PropVariantToInt32WithDefault(propvarIn *PROPVARIANT, lDefault int32) int32 {
	addr := LazyAddr(&pPropVariantToInt32WithDefault, libPropsys, "PropVariantToInt32WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(lDefault))
	return int32(ret)
}

func PropVariantToUInt32WithDefault(propvarIn *PROPVARIANT, ulDefault uint32) uint32 {
	addr := LazyAddr(&pPropVariantToUInt32WithDefault, libPropsys, "PropVariantToUInt32WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(ulDefault))
	return uint32(ret)
}

func PropVariantToInt64WithDefault(propvarIn *PROPVARIANT, llDefault int64) int64 {
	addr := LazyAddr(&pPropVariantToInt64WithDefault, libPropsys, "PropVariantToInt64WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(llDefault))
	return int64(ret)
}

func PropVariantToUInt64WithDefault(propvarIn *PROPVARIANT, ullDefault uint64) uint64 {
	addr := LazyAddr(&pPropVariantToUInt64WithDefault, libPropsys, "PropVariantToUInt64WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(ullDefault))
	return uint64(ret)
}

func PropVariantToDoubleWithDefault(propvarIn *PROPVARIANT, dblDefault float64) float64 {
	addr := LazyAddr(&pPropVariantToDoubleWithDefault, libPropsys, "PropVariantToDoubleWithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(dblDefault))
	return float64(ret)
}

func PropVariantToStringWithDefault(propvarIn *PROPVARIANT, pszDefault PWSTR) PWSTR {
	addr := LazyAddr(&pPropVariantToStringWithDefault, libPropsys, "PropVariantToStringWithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(pszDefault)))
	return (PWSTR)(unsafe.Pointer(ret))
}

func PropVariantToBoolean(propvarIn *PROPVARIANT, pfRet *BOOL) HRESULT {
	addr := LazyAddr(&pPropVariantToBoolean, libPropsys, "PropVariantToBoolean")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(pfRet)))
	return HRESULT(ret)
}

func PropVariantToInt16(propvarIn *PROPVARIANT, piRet *int16) HRESULT {
	addr := LazyAddr(&pPropVariantToInt16, libPropsys, "PropVariantToInt16")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(piRet)))
	return HRESULT(ret)
}

func PropVariantToUInt16(propvarIn *PROPVARIANT, puiRet *uint16) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt16, libPropsys, "PropVariantToUInt16")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(puiRet)))
	return HRESULT(ret)
}

func PropVariantToInt32(propvarIn *PROPVARIANT, plRet *int32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt32, libPropsys, "PropVariantToInt32")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(plRet)))
	return HRESULT(ret)
}

func PropVariantToUInt32(propvarIn *PROPVARIANT, pulRet *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt32, libPropsys, "PropVariantToUInt32")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(pulRet)))
	return HRESULT(ret)
}

func PropVariantToInt64(propvarIn *PROPVARIANT, pllRet *int64) HRESULT {
	addr := LazyAddr(&pPropVariantToInt64, libPropsys, "PropVariantToInt64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(pllRet)))
	return HRESULT(ret)
}

func PropVariantToUInt64(propvarIn *PROPVARIANT, pullRet *uint64) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt64, libPropsys, "PropVariantToUInt64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(pullRet)))
	return HRESULT(ret)
}

func PropVariantToDouble(propvarIn *PROPVARIANT, pdblRet *float64) HRESULT {
	addr := LazyAddr(&pPropVariantToDouble, libPropsys, "PropVariantToDouble")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvarIn)), uintptr(unsafe.Pointer(pdblRet)))
	return HRESULT(ret)
}

func PropVariantToBuffer(propvar *PROPVARIANT, pv unsafe.Pointer, cb uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToBuffer, libPropsys, "PropVariantToBuffer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(pv), uintptr(cb))
	return HRESULT(ret)
}

func PropVariantToString(propvar *PROPVARIANT, psz PWSTR, cch uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToString, libPropsys, "PropVariantToString")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(psz)), uintptr(cch))
	return HRESULT(ret)
}

func PropVariantToGUID(propvar *PROPVARIANT, pguid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pPropVariantToGUID, libPropsys, "PropVariantToGUID")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pguid)))
	return HRESULT(ret)
}

func PropVariantToStringAlloc(propvar *PROPVARIANT, ppszOut *PWSTR) HRESULT {
	addr := LazyAddr(&pPropVariantToStringAlloc, libPropsys, "PropVariantToStringAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(ppszOut)))
	return HRESULT(ret)
}

func PropVariantToBSTR(propvar *PROPVARIANT, pbstrOut *BSTR) HRESULT {
	addr := LazyAddr(&pPropVariantToBSTR, libPropsys, "PropVariantToBSTR")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pbstrOut)))
	return HRESULT(ret)
}

func PropVariantToFileTime(propvar *PROPVARIANT, pstfOut PSTIME_FLAGS, pftOut *FILETIME) HRESULT {
	addr := LazyAddr(&pPropVariantToFileTime, libPropsys, "PropVariantToFileTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(pstfOut), uintptr(unsafe.Pointer(pftOut)))
	return HRESULT(ret)
}

func PropVariantGetElementCount(propvar *PROPVARIANT) uint32 {
	addr := LazyAddr(&pPropVariantGetElementCount, libPropsys, "PropVariantGetElementCount")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)))
	return uint32(ret)
}

func PropVariantToBooleanVector(propvar *PROPVARIANT, prgf *BOOL, crgf uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToBooleanVector, libPropsys, "PropVariantToBooleanVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgf)), uintptr(crgf), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToInt16Vector(propvar *PROPVARIANT, prgn *int16, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt16Vector, libPropsys, "PropVariantToInt16Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToUInt16Vector(propvar *PROPVARIANT, prgn *uint16, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt16Vector, libPropsys, "PropVariantToUInt16Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToInt32Vector(propvar *PROPVARIANT, prgn *int32, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt32Vector, libPropsys, "PropVariantToInt32Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToUInt32Vector(propvar *PROPVARIANT, prgn *uint32, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt32Vector, libPropsys, "PropVariantToUInt32Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToInt64Vector(propvar *PROPVARIANT, prgn *int64, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt64Vector, libPropsys, "PropVariantToInt64Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToUInt64Vector(propvar *PROPVARIANT, prgn *uint64, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt64Vector, libPropsys, "PropVariantToUInt64Vector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToDoubleVector(propvar *PROPVARIANT, prgn *float64, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToDoubleVector, libPropsys, "PropVariantToDoubleVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToFileTimeVector(propvar *PROPVARIANT, prgft *FILETIME, crgft uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToFileTimeVector, libPropsys, "PropVariantToFileTimeVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgft)), uintptr(crgft), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToStringVector(propvar *PROPVARIANT, prgsz *PWSTR, crgsz uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToStringVector, libPropsys, "PropVariantToStringVector")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(prgsz)), uintptr(crgsz), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToBooleanVectorAlloc(propvar *PROPVARIANT, pprgf **BOOL, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToBooleanVectorAlloc, libPropsys, "PropVariantToBooleanVectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgf)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToInt16VectorAlloc(propvar *PROPVARIANT, pprgn **int16, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt16VectorAlloc, libPropsys, "PropVariantToInt16VectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToUInt16VectorAlloc(propvar *PROPVARIANT, pprgn **uint16, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt16VectorAlloc, libPropsys, "PropVariantToUInt16VectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToInt32VectorAlloc(propvar *PROPVARIANT, pprgn **int32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt32VectorAlloc, libPropsys, "PropVariantToInt32VectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToUInt32VectorAlloc(propvar *PROPVARIANT, pprgn **uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt32VectorAlloc, libPropsys, "PropVariantToUInt32VectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToInt64VectorAlloc(propvar *PROPVARIANT, pprgn **int64, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToInt64VectorAlloc, libPropsys, "PropVariantToInt64VectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToUInt64VectorAlloc(propvar *PROPVARIANT, pprgn **uint64, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToUInt64VectorAlloc, libPropsys, "PropVariantToUInt64VectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToDoubleVectorAlloc(propvar *PROPVARIANT, pprgn **float64, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToDoubleVectorAlloc, libPropsys, "PropVariantToDoubleVectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToFileTimeVectorAlloc(propvar *PROPVARIANT, pprgft **FILETIME, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToFileTimeVectorAlloc, libPropsys, "PropVariantToFileTimeVectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgft)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantToStringVectorAlloc(propvar *PROPVARIANT, pprgsz **PWSTR, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantToStringVectorAlloc, libPropsys, "PropVariantToStringVectorAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pprgsz)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func PropVariantGetBooleanElem(propvar *PROPVARIANT, iElem uint32, pfVal *BOOL) HRESULT {
	addr := LazyAddr(&pPropVariantGetBooleanElem, libPropsys, "PropVariantGetBooleanElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pfVal)))
	return HRESULT(ret)
}

func PropVariantGetInt16Elem(propvar *PROPVARIANT, iElem uint32, pnVal *int16) HRESULT {
	addr := LazyAddr(&pPropVariantGetInt16Elem, libPropsys, "PropVariantGetInt16Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetUInt16Elem(propvar *PROPVARIANT, iElem uint32, pnVal *uint16) HRESULT {
	addr := LazyAddr(&pPropVariantGetUInt16Elem, libPropsys, "PropVariantGetUInt16Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetInt32Elem(propvar *PROPVARIANT, iElem uint32, pnVal *int32) HRESULT {
	addr := LazyAddr(&pPropVariantGetInt32Elem, libPropsys, "PropVariantGetInt32Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetUInt32Elem(propvar *PROPVARIANT, iElem uint32, pnVal *uint32) HRESULT {
	addr := LazyAddr(&pPropVariantGetUInt32Elem, libPropsys, "PropVariantGetUInt32Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetInt64Elem(propvar *PROPVARIANT, iElem uint32, pnVal *int64) HRESULT {
	addr := LazyAddr(&pPropVariantGetInt64Elem, libPropsys, "PropVariantGetInt64Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetUInt64Elem(propvar *PROPVARIANT, iElem uint32, pnVal *uint64) HRESULT {
	addr := LazyAddr(&pPropVariantGetUInt64Elem, libPropsys, "PropVariantGetUInt64Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetDoubleElem(propvar *PROPVARIANT, iElem uint32, pnVal *float64) HRESULT {
	addr := LazyAddr(&pPropVariantGetDoubleElem, libPropsys, "PropVariantGetDoubleElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func PropVariantGetFileTimeElem(propvar *PROPVARIANT, iElem uint32, pftVal *FILETIME) HRESULT {
	addr := LazyAddr(&pPropVariantGetFileTimeElem, libPropsys, "PropVariantGetFileTimeElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(pftVal)))
	return HRESULT(ret)
}

func PropVariantGetStringElem(propvar *PROPVARIANT, iElem uint32, ppszVal *PWSTR) HRESULT {
	addr := LazyAddr(&pPropVariantGetStringElem, libPropsys, "PropVariantGetStringElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(iElem), uintptr(unsafe.Pointer(ppszVal)))
	return HRESULT(ret)
}

func ClearPropVariantArray(rgPropVar *PROPVARIANT, cVars uint32) {
	addr := LazyAddr(&pClearPropVariantArray, libPropsys, "ClearPropVariantArray")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(rgPropVar)), uintptr(cVars))
}

func PropVariantCompareEx(propvar1 *PROPVARIANT, propvar2 *PROPVARIANT, unit PROPVAR_COMPARE_UNIT, flags PROPVAR_COMPARE_FLAGS) int32 {
	addr := LazyAddr(&pPropVariantCompareEx, libPropsys, "PropVariantCompareEx")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar1)), uintptr(unsafe.Pointer(propvar2)), uintptr(unit), uintptr(flags))
	return int32(ret)
}

func PropVariantChangeType(ppropvarDest *PROPVARIANT, propvarSrc *PROPVARIANT, flags PROPVAR_CHANGE_FLAGS, vt VARENUM) HRESULT {
	addr := LazyAddr(&pPropVariantChangeType, libPropsys, "PropVariantChangeType")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ppropvarDest)), uintptr(unsafe.Pointer(propvarSrc)), uintptr(flags), uintptr(vt))
	return HRESULT(ret)
}

func PropVariantToVariant(pPropVar *PROPVARIANT, pVar *VARIANT) HRESULT {
	addr := LazyAddr(&pPropVariantToVariant, libPropsys, "PropVariantToVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pPropVar)), uintptr(unsafe.Pointer(pVar)))
	return HRESULT(ret)
}

func VariantToPropVariant(pVar *VARIANT, pPropVar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pVariantToPropVariant, libPropsys, "VariantToPropVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pVar)), uintptr(unsafe.Pointer(pPropVar)))
	return HRESULT(ret)
}

func StgSerializePropVariant(ppropvar *PROPVARIANT, ppProp **SERIALIZEDPROPERTYVALUE, pcb *uint32) HRESULT {
	addr := LazyAddr(&pStgSerializePropVariant, libPropsys, "StgSerializePropVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ppropvar)), uintptr(unsafe.Pointer(ppProp)), uintptr(unsafe.Pointer(pcb)))
	return HRESULT(ret)
}

func StgDeserializePropVariant(pprop *SERIALIZEDPROPERTYVALUE, cbMax uint32, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pStgDeserializePropVariant, libPropsys, "StgDeserializePropVariant")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pprop)), uintptr(cbMax), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}
