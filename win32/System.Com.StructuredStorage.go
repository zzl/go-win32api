package win32

import "unsafe"
import "syscall"

const (
	PROPSETFLAG_DEFAULT uint32 = 0
	PROPSETFLAG_NONSIMPLE uint32 = 1
	PROPSETFLAG_ANSI uint32 = 2
	PROPSETFLAG_UNBUFFERED uint32 = 4
	PROPSETFLAG_CASE_SENSITIVE uint32 = 8
	PROPSET_BEHAVIOR_CASE_SENSITIVE uint32 = 1
	PID_DICTIONARY uint32 = 0
	PID_CODEPAGE uint32 = 1
	PID_FIRST_USABLE uint32 = 2
	PID_FIRST_NAME_DEFAULT uint32 = 4095
	PID_LOCALE uint32 = 2147483648
	PID_MODIFY_TIME uint32 = 2147483649
	PID_SECURITY uint32 = 2147483650
	PID_BEHAVIOR uint32 = 2147483651
	PID_ILLEGAL uint32 = 4294967295
	PID_MIN_READONLY uint32 = 2147483648
	PID_MAX_READONLY uint32 = 3221225471
	PRSPEC_INVALID uint32 = 4294967295
	PROPSETHDR_OSVERSION_UNKNOWN uint32 = 4294967295
	PIDDI_THUMBNAIL int32 = 2
	PIDSI_TITLE int32 = 2
	PIDSI_SUBJECT int32 = 3
	PIDSI_AUTHOR int32 = 4
	PIDSI_KEYWORDS int32 = 5
	PIDSI_COMMENTS int32 = 6
	PIDSI_TEMPLATE int32 = 7
	PIDSI_LASTAUTHOR int32 = 8
	PIDSI_REVNUMBER int32 = 9
	PIDSI_EDITTIME int32 = 10
	PIDSI_LASTPRINTED int32 = 11
	PIDSI_CREATE_DTM int32 = 12
	PIDSI_LASTSAVE_DTM int32 = 13
	PIDSI_PAGECOUNT int32 = 14
	PIDSI_WORDCOUNT int32 = 15
	PIDSI_CHARCOUNT int32 = 16
	PIDSI_THUMBNAIL int32 = 17
	PIDSI_APPNAME int32 = 18
	PIDSI_DOC_SECURITY int32 = 19
	PIDDSI_CATEGORY uint32 = 2
	PIDDSI_PRESFORMAT uint32 = 3
	PIDDSI_BYTECOUNT uint32 = 4
	PIDDSI_LINECOUNT uint32 = 5
	PIDDSI_PARCOUNT uint32 = 6
	PIDDSI_SLIDECOUNT uint32 = 7
	PIDDSI_NOTECOUNT uint32 = 8
	PIDDSI_HIDDENCOUNT uint32 = 9
	PIDDSI_MMCLIPCOUNT uint32 = 10
	PIDDSI_SCALE uint32 = 11
	PIDDSI_HEADINGPAIR uint32 = 12
	PIDDSI_DOCPARTS uint32 = 13
	PIDDSI_MANAGER uint32 = 14
	PIDDSI_COMPANY uint32 = 15
	PIDDSI_LINKSDIRTY uint32 = 16
	PIDMSI_EDITOR int32 = 2
	PIDMSI_SUPPLIER int32 = 3
	PIDMSI_SOURCE int32 = 4
	PIDMSI_SEQUENCE_NO int32 = 5
	PIDMSI_PROJECT int32 = 6
	PIDMSI_STATUS int32 = 7
	PIDMSI_OWNER int32 = 8
	PIDMSI_RATING int32 = 9
	PIDMSI_PRODUCTION int32 = 10
	PIDMSI_COPYRIGHT int32 = 11
	CWCSTORAGENAME uint32 = 32
	STGM_DIRECT int32 = 0
	STGM_TRANSACTED int32 = 65536
	STGM_SIMPLE int32 = 134217728
	STGM_READ int32 = 0
	STGM_WRITE int32 = 1
	STGM_READWRITE int32 = 2
	STGM_SHARE_DENY_NONE int32 = 64
	STGM_SHARE_DENY_READ int32 = 48
	STGM_SHARE_DENY_WRITE int32 = 32
	STGM_SHARE_EXCLUSIVE int32 = 16
	STGM_PRIORITY int32 = 262144
	STGM_DELETEONRELEASE int32 = 67108864
	STGM_NOSCRATCH int32 = 1048576
	STGM_CREATE int32 = 4096
	STGM_CONVERT int32 = 131072
	STGM_FAILIFTHERE int32 = 0
	STGM_NOSNAPSHOT int32 = 2097152
	STGM_DIRECT_SWMR int32 = 4194304
	STGFMT_STORAGE uint32 = 0
	STGFMT_NATIVE uint32 = 1
	STGFMT_FILE uint32 = 3
	STGFMT_ANY uint32 = 4
	STGFMT_DOCFILE uint32 = 5
	STGFMT_DOCUMENT uint32 = 0
	STGOPTIONS_VERSION uint32 = 1
	CCH_MAX_PROPSTG_NAME uint32 = 31
)

// enums

// enum PROPSPEC_KIND
type PROPSPEC_KIND uint32
const (
	PRSPEC_LPWSTR PROPSPEC_KIND = 0
	PRSPEC_PROPID PROPSPEC_KIND = 1
)

// enum STGC
type STGC int32
const (
	STGC_DEFAULT STGC = 0
	STGC_OVERWRITE STGC = 1
	STGC_ONLYIFCURRENT STGC = 2
	STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE STGC = 4
	STGC_CONSOLIDATE STGC = 8
)

// enum STGMOVE
type STGMOVE int32
const (
	STGMOVE_MOVE STGMOVE = 0
	STGMOVE_COPY STGMOVE = 1
	STGMOVE_SHALLOWCOPY STGMOVE = 2
)

// enum STATFLAG
type STATFLAG int32
const (
	STATFLAG_DEFAULT STATFLAG = 0
	STATFLAG_NONAME STATFLAG = 1
	STATFLAG_NOOPEN STATFLAG = 2
)

// enum LOCKTYPE
type LOCKTYPE int32
const (
	LOCK_WRITE LOCKTYPE = 1
	LOCK_EXCLUSIVE LOCKTYPE = 2
	LOCK_ONLYONCE LOCKTYPE = 4
)

// enum PIDMSI_STATUS_VALUE
type PIDMSI_STATUS_VALUE int32
const (
	PIDMSI_STATUS_NORMAL PIDMSI_STATUS_VALUE = 0
	PIDMSI_STATUS_NEW PIDMSI_STATUS_VALUE = 1
	PIDMSI_STATUS_PRELIM PIDMSI_STATUS_VALUE = 2
	PIDMSI_STATUS_DRAFT PIDMSI_STATUS_VALUE = 3
	PIDMSI_STATUS_INPROGRESS PIDMSI_STATUS_VALUE = 4
	PIDMSI_STATUS_EDIT PIDMSI_STATUS_VALUE = 5
	PIDMSI_STATUS_REVIEW PIDMSI_STATUS_VALUE = 6
	PIDMSI_STATUS_PROOF PIDMSI_STATUS_VALUE = 7
	PIDMSI_STATUS_FINAL PIDMSI_STATUS_VALUE = 8
	PIDMSI_STATUS_OTHER PIDMSI_STATUS_VALUE = 32767
)


// structs

type BSTRBLOB struct {
	CbSize uint32
	PData *uint8
}

type CLIPDATA struct {
	CbSize uint32
	UlClipFmt int32
	PClipData *uint8
}

type RemSNB struct {
	UlCntStr uint32
	UlCntChar uint32
	RgString [1]uint16
}

type VERSIONEDSTREAM struct {
	GuidVersion syscall.GUID
	PStream *IStream
}

type CAC struct {
	CElems uint32
	PElems PSTR
}

type CAUB struct {
	CElems uint32
	PElems *uint8
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
	PElems *int16
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

type PROPVARIANT_Anonymous__Anonymous__Anonymous_ struct {
	Data [2]uint64
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CVal() *CHAR{
	return (*CHAR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CValVal() CHAR{
	return *(*CHAR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BVal() *uint8{
	return (*uint8)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BValVal() uint8{
	return *(*uint8)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) IVal() *int16{
	return (*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) IValVal() int16{
	return *(*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UiVal() *uint16{
	return (*uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UiValVal() uint16{
	return *(*uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) LVal() *int32{
	return (*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) LValVal() int32{
	return *(*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UlVal() *uint32{
	return (*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UlValVal() uint32{
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) IntVal() *int32{
	return (*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) IntValVal() int32{
	return *(*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UintVal() *uint32{
	return (*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UintValVal() uint32{
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) HVal() *int64{
	return (*int64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) HValVal() int64{
	return *(*int64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UhVal() *uint64{
	return (*uint64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) UhValVal() uint64{
	return *(*uint64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) FltVal() *float32{
	return (*float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) FltValVal() float32{
	return *(*float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) DblVal() *float64{
	return (*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) DblValVal() float64{
	return *(*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BoolVal() *int16{
	return (*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BoolValVal() int16{
	return *(*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) OBSOLETE__VARIANT_BOOL__() *int16{
	return (*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) OBSOLETE__VARIANT_BOOL__Val() int16{
	return *(*int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Scode() *int32{
	return (*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) ScodeVal() int32{
	return *(*int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CyVal() *CY{
	return (*CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CyValVal() CY{
	return *(*CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Date() *float64{
	return (*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) DateVal() float64{
	return *(*float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Filetime() *FILETIME{
	return (*FILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) FiletimeVal() FILETIME{
	return *(*FILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Puuid() **syscall.GUID{
	return (**syscall.GUID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PuuidVal() *syscall.GUID{
	return *(**syscall.GUID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Pclipdata() **CLIPDATA{
	return (**CLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PclipdataVal() *CLIPDATA{
	return *(**CLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BstrVal() *BSTR{
	return (*BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BstrValVal() BSTR{
	return *(*BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BstrblobVal() *BSTRBLOB{
	return (*BSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BstrblobValVal() BSTRBLOB{
	return *(*BSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Blob() *BLOB{
	return (*BLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) BlobVal() BLOB{
	return *(*BLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PszVal() *PSTR{
	return (*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PszValVal() PSTR{
	return *(*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PwszVal() *PWSTR{
	return (*PWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PwszValVal() PWSTR{
	return *(*PWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PunkVal() **IUnknown{
	return (**IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PunkValVal() *IUnknown{
	return *(**IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdispVal() **IDispatch{
	return (**IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdispValVal() *IDispatch{
	return *(**IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PStream() **IStream{
	return (**IStream)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PStreamVal() *IStream{
	return *(**IStream)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PStorage() **IStorage{
	return (**IStorage)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PStorageVal() *IStorage{
	return *(**IStorage)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PVersionedStream() **VERSIONEDSTREAM{
	return (**VERSIONEDSTREAM)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PVersionedStreamVal() *VERSIONEDSTREAM{
	return *(**VERSIONEDSTREAM)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Parray() **SAFEARRAY{
	return (**SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) ParrayVal() *SAFEARRAY{
	return *(**SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cac() *CAC{
	return (*CAC)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CacVal() CAC{
	return *(*CAC)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Caub() *CAUB{
	return (*CAUB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CaubVal() CAUB{
	return *(*CAUB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cai() *CAI{
	return (*CAI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CaiVal() CAI{
	return *(*CAI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Caui() *CAUI{
	return (*CAUI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CauiVal() CAUI{
	return *(*CAUI)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cal() *CAL{
	return (*CAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CalVal() CAL{
	return *(*CAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Caul() *CAUL{
	return (*CAUL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CaulVal() CAUL{
	return *(*CAUL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cah() *CAH{
	return (*CAH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CahVal() CAH{
	return *(*CAH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cauh() *CAUH{
	return (*CAUH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CauhVal() CAUH{
	return *(*CAUH)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Caflt() *CAFLT{
	return (*CAFLT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CafltVal() CAFLT{
	return *(*CAFLT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cadbl() *CADBL{
	return (*CADBL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CadblVal() CADBL{
	return *(*CADBL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cabool() *CABOOL{
	return (*CABOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CaboolVal() CABOOL{
	return *(*CABOOL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cascode() *CASCODE{
	return (*CASCODE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CascodeVal() CASCODE{
	return *(*CASCODE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cacy() *CACY{
	return (*CACY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CacyVal() CACY{
	return *(*CACY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cadate() *CADATE{
	return (*CADATE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CadateVal() CADATE{
	return *(*CADATE)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cafiletime() *CAFILETIME{
	return (*CAFILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CafiletimeVal() CAFILETIME{
	return *(*CAFILETIME)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cauuid() *CACLSID{
	return (*CACLSID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CauuidVal() CACLSID{
	return *(*CACLSID)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Caclipdata() *CACLIPDATA{
	return (*CACLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CaclipdataVal() CACLIPDATA{
	return *(*CACLIPDATA)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cabstr() *CABSTR{
	return (*CABSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CabstrVal() CABSTR{
	return *(*CABSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Cabstrblob() *CABSTRBLOB{
	return (*CABSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CabstrblobVal() CABSTRBLOB{
	return *(*CABSTRBLOB)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Calpstr() *CALPSTR{
	return (*CALPSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CalpstrVal() CALPSTR{
	return *(*CALPSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Calpwstr() *CALPWSTR{
	return (*CALPWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CalpwstrVal() CALPWSTR{
	return *(*CALPWSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Capropvar() *CAPROPVARIANT{
	return (*CAPROPVARIANT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) CapropvarVal() CAPROPVARIANT{
	return *(*CAPROPVARIANT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PcVal() *PSTR{
	return (*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PcValVal() PSTR{
	return *(*PSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PbVal() **uint8{
	return (**uint8)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PbValVal() *uint8{
	return *(**uint8)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PiVal() **int16{
	return (**int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PiValVal() *int16{
	return *(**int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PuiVal() **uint16{
	return (**uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PuiValVal() *uint16{
	return *(**uint16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PlVal() **int32{
	return (**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PlValVal() *int32{
	return *(**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PulVal() **uint32{
	return (**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PulValVal() *uint32{
	return *(**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PintVal() **int32{
	return (**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PintValVal() *int32{
	return *(**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PuintVal() **uint32{
	return (**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PuintValVal() *uint32{
	return *(**uint32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PfltVal() **float32{
	return (**float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PfltValVal() *float32{
	return *(**float32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdblVal() **float64{
	return (**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdblValVal() *float64{
	return *(**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PboolVal() **int16{
	return (**int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PboolValVal() *int16{
	return *(**int16)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdecVal() **DECIMAL{
	return (**DECIMAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdecValVal() *DECIMAL{
	return *(**DECIMAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Pscode() **int32{
	return (**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PscodeVal() *int32{
	return *(**int32)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PcyVal() **CY{
	return (**CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PcyValVal() *CY{
	return *(**CY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Pdate() **float64{
	return (**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PdateVal() *float64{
	return *(**float64)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PbstrVal() **BSTR{
	return (**BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PbstrValVal() *BSTR{
	return *(**BSTR)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PpunkVal() ***IUnknown{
	return (***IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PpunkValVal() **IUnknown{
	return *(***IUnknown)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PpdispVal() ***IDispatch{
	return (***IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PpdispValVal() **IDispatch{
	return *(***IDispatch)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) Pparray() ***SAFEARRAY{
	return (***SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PparrayVal() **SAFEARRAY{
	return *(***SAFEARRAY)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PvarVal() **PROPVARIANT{
	return (**PROPVARIANT)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous__Anonymous__Anonymous_) PvarValVal() *PROPVARIANT{
	return *(**PROPVARIANT)(unsafe.Pointer(this))
}

type PROPVARIANT_Anonymous__Anonymous_ struct {
	Vt uint16
	WReserved1 uint16
	WReserved2 uint16
	WReserved3 uint16
	PROPVARIANT_Anonymous__Anonymous__Anonymous_
}

type PROPVARIANT_Anonymous_ struct {
	 PROPVARIANT_Anonymous__Anonymous_
}

func (this *PROPVARIANT_Anonymous_) DecVal() *DECIMAL{
	return (*DECIMAL)(unsafe.Pointer(this))
}

func (this *PROPVARIANT_Anonymous_) DecValVal() DECIMAL{
	return *(*DECIMAL)(unsafe.Pointer(this))
}

type PROPVARIANT struct {
	PROPVARIANT_Anonymous_
}

type PROPSPEC_Anonymous_ struct {
	Data [1]uint64
}

func (this *PROPSPEC_Anonymous_) Propid() *uint32{
	return (*uint32)(unsafe.Pointer(this))
}

func (this *PROPSPEC_Anonymous_) PropidVal() uint32{
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *PROPSPEC_Anonymous_) Lpwstr() *PWSTR{
	return (*PWSTR)(unsafe.Pointer(this))
}

func (this *PROPSPEC_Anonymous_) LpwstrVal() PWSTR{
	return *(*PWSTR)(unsafe.Pointer(this))
}

type PROPSPEC struct {
	UlKind PROPSPEC_KIND
	PROPSPEC_Anonymous_
}

type STATPROPSTG struct {
	LpwstrName PWSTR
	Propid uint32
	Vt uint16
}

type STATPROPSETSTG struct {
	Fmtid syscall.GUID
	Clsid syscall.GUID
	GrfFlags uint32
	Mtime FILETIME
	Ctime FILETIME
	Atime FILETIME
	DwOSVersion uint32
}

type STGOPTIONS struct {
	UsVersion uint16
	Reserved uint16
	UlSectorSize uint32
	PwcsTemplateFile PWSTR
}

type SERIALIZEDPROPERTYVALUE struct {
	DwType uint32
	Rgb [1]uint8
}

type PMemoryAllocator struct {
}

type OLESTREAMVTBL struct {
	Get uintptr
	Put uintptr
}

type OLESTREAM struct {
	Lpstbl *OLESTREAMVTBL
}

type PROPBAG2 struct {
	DwType uint32
	Vt uint16
	CfType uint16
	DwHint uint32
	PstrName PWSTR
	Clsid syscall.GUID
}


// coms

// 0000000d-0000-0000-c000-000000000046
var IID_IEnumSTATSTG = syscall.GUID{0x0000000d, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumSTATSTGInterface interface {
	IUnknownInterface
	Next(celt uint32, rgelt *STATSTG, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppenum **IEnumSTATSTG) HRESULT
}

type IEnumSTATSTGVtbl struct {
	IUnknownVtbl
	Next uintptr
	Skip uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumSTATSTG struct {
	IUnknown
}

func (this *IEnumSTATSTG) Vtbl() *IEnumSTATSTGVtbl {
	return (*IEnumSTATSTGVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumSTATSTG) Next(celt uint32, rgelt *STATSTG, pceltFetched *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgelt)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumSTATSTG) Skip(celt uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumSTATSTG) Reset() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumSTATSTG) Clone(ppenum **IEnumSTATSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 0000000b-0000-0000-c000-000000000046
var IID_IStorage = syscall.GUID{0x0000000b, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IStorageInterface interface {
	IUnknownInterface
	CreateStream(pwcsName PWSTR, grfMode uint32, reserved1 uint32, reserved2 uint32, ppstm **IStream) HRESULT
	OpenStream(pwcsName PWSTR, reserved1 unsafe.Pointer, grfMode uint32, reserved2 uint32, ppstm **IStream) HRESULT
	CreateStorage(pwcsName PWSTR, grfMode uint32, reserved1 uint32, reserved2 uint32, ppstg **IStorage) HRESULT
	OpenStorage(pwcsName PWSTR, pstgPriority *IStorage, grfMode uint32, snbExclude **uint16, reserved uint32, ppstg **IStorage) HRESULT
	CopyTo(ciidExclude uint32, rgiidExclude *syscall.GUID, snbExclude **uint16, pstgDest *IStorage) HRESULT
	MoveElementTo(pwcsName PWSTR, pstgDest *IStorage, pwcsNewName PWSTR, grfFlags uint32) HRESULT
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
	CreateStream uintptr
	OpenStream uintptr
	CreateStorage uintptr
	OpenStorage uintptr
	CopyTo uintptr
	MoveElementTo uintptr
	Commit uintptr
	Revert uintptr
	EnumElements uintptr
	DestroyElement uintptr
	RenameElement uintptr
	SetElementTimes uintptr
	SetClass uintptr
	SetStateBits uintptr
	Stat uintptr
}

type IStorage struct {
	IUnknown
}

func (this *IStorage) Vtbl() *IStorageVtbl {
	return (*IStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStorage) CreateStream(pwcsName PWSTR, grfMode uint32, reserved1 uint32, reserved2 uint32, ppstm **IStream) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(reserved1), uintptr(reserved2), uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret)
}

func (this *IStorage) OpenStream(pwcsName PWSTR, reserved1 unsafe.Pointer, grfMode uint32, reserved2 uint32, ppstm **IStream) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().OpenStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(reserved1)), uintptr(grfMode), uintptr(reserved2), uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret)
}

func (this *IStorage) CreateStorage(pwcsName PWSTR, grfMode uint32, reserved1 uint32, reserved2 uint32, ppstg **IStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(reserved1), uintptr(reserved2), uintptr(unsafe.Pointer(ppstg)))
	return HRESULT(ret)
}

func (this *IStorage) OpenStorage(pwcsName PWSTR, pstgPriority *IStorage, grfMode uint32, snbExclude **uint16, reserved uint32, ppstg **IStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().OpenStorage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pstgPriority)), uintptr(grfMode), uintptr(unsafe.Pointer(snbExclude)), uintptr(reserved), uintptr(unsafe.Pointer(ppstg)))
	return HRESULT(ret)
}

func (this *IStorage) CopyTo(ciidExclude uint32, rgiidExclude *syscall.GUID, snbExclude **uint16, pstgDest *IStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CopyTo, uintptr(unsafe.Pointer(this)), uintptr(ciidExclude), uintptr(unsafe.Pointer(rgiidExclude)), uintptr(unsafe.Pointer(snbExclude)), uintptr(unsafe.Pointer(pstgDest)))
	return HRESULT(ret)
}

func (this *IStorage) MoveElementTo(pwcsName PWSTR, pstgDest *IStorage, pwcsNewName PWSTR, grfFlags uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().MoveElementTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pstgDest)), uintptr(unsafe.Pointer(pwcsNewName)), uintptr(grfFlags))
	return HRESULT(ret)
}

func (this *IStorage) Commit(grfCommitFlags uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Commit, uintptr(unsafe.Pointer(this)), uintptr(grfCommitFlags))
	return HRESULT(ret)
}

func (this *IStorage) Revert() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Revert, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IStorage) EnumElements(reserved1 uint32, reserved2 unsafe.Pointer, reserved3 uint32, ppenum **IEnumSTATSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().EnumElements, uintptr(unsafe.Pointer(this)), uintptr(reserved1), uintptr(unsafe.Pointer(reserved2)), uintptr(reserved3), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

func (this *IStorage) DestroyElement(pwcsName PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DestroyElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)))
	return HRESULT(ret)
}

func (this *IStorage) RenameElement(pwcsOldName PWSTR, pwcsNewName PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RenameElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsOldName)), uintptr(unsafe.Pointer(pwcsNewName)))
	return HRESULT(ret)
}

func (this *IStorage) SetElementTimes(pwcsName PWSTR, pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetElementTimes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pctime)), uintptr(unsafe.Pointer(patime)), uintptr(unsafe.Pointer(pmtime)))
	return HRESULT(ret)
}

func (this *IStorage) SetClass(clsid *syscall.GUID) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clsid)))
	return HRESULT(ret)
}

func (this *IStorage) SetStateBits(grfStateBits uint32, grfMask uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStateBits, uintptr(unsafe.Pointer(this)), uintptr(grfStateBits), uintptr(grfMask))
	return HRESULT(ret)
}

func (this *IStorage) Stat(pstatstg *STATSTG, grfStatFlag uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Stat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstatstg)), uintptr(grfStatFlag))
	return HRESULT(ret)
}

// 0000010a-0000-0000-c000-000000000046
var IID_IPersistStorage = syscall.GUID{0x0000010a, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

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
	IsDirty uintptr
	InitNew uintptr
	Load uintptr
	Save uintptr
	SaveCompleted uintptr
	HandsOffStorage uintptr
}

type IPersistStorage struct {
	IPersist
}

func (this *IPersistStorage) Vtbl() *IPersistStorageVtbl {
	return (*IPersistStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPersistStorage) IsDirty() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsDirty, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IPersistStorage) InitNew(pStg *IStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().InitNew, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStg)))
	return HRESULT(ret)
}

func (this *IPersistStorage) Load(pStg *IStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Load, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStg)))
	return HRESULT(ret)
}

func (this *IPersistStorage) Save(pStgSave *IStorage, fSameAsLoad BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Save, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStgSave)), uintptr(fSameAsLoad))
	return HRESULT(ret)
}

func (this *IPersistStorage) SaveCompleted(pStgNew *IStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SaveCompleted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStgNew)))
	return HRESULT(ret)
}

func (this *IPersistStorage) HandsOffStorage() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandsOffStorage, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 0000000a-0000-0000-c000-000000000046
var IID_ILockBytes = syscall.GUID{0x0000000a, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

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
	ReadAt uintptr
	WriteAt uintptr
	Flush uintptr
	SetSize uintptr
	LockRegion uintptr
	UnlockRegion uintptr
	Stat uintptr
}

type ILockBytes struct {
	IUnknown
}

func (this *ILockBytes) Vtbl() *ILockBytesVtbl {
	return (*ILockBytesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILockBytes) ReadAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbRead *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReadAt, uintptr(unsafe.Pointer(this)), uintptr(ulOffset), uintptr(unsafe.Pointer(pv)), uintptr(cb), uintptr(unsafe.Pointer(pcbRead)))
	return HRESULT(ret)
}

func (this *ILockBytes) WriteAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteAt, uintptr(unsafe.Pointer(this)), uintptr(ulOffset), uintptr(unsafe.Pointer(pv)), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

func (this *ILockBytes) Flush() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Flush, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILockBytes) SetSize(cb uint64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSize, uintptr(unsafe.Pointer(this)), uintptr(cb))
	return HRESULT(ret)
}

func (this *ILockBytes) LockRegion(libOffset uint64, cb uint64, dwLockType uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().LockRegion, uintptr(unsafe.Pointer(this)), uintptr(libOffset), uintptr(cb), uintptr(dwLockType))
	return HRESULT(ret)
}

func (this *ILockBytes) UnlockRegion(libOffset uint64, cb uint64, dwLockType uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().UnlockRegion, uintptr(unsafe.Pointer(this)), uintptr(libOffset), uintptr(cb), uintptr(dwLockType))
	return HRESULT(ret)
}

func (this *ILockBytes) Stat(pstatstg *STATSTG, grfStatFlag uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Stat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstatstg)), uintptr(grfStatFlag))
	return HRESULT(ret)
}

// 00000012-0000-0000-c000-000000000046
var IID_IRootStorage = syscall.GUID{0x00000012, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

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

func (this *IRootStorage) SwitchToFile(pszFile PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SwitchToFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszFile)))
	return HRESULT(ret)
}

// 99caf010-415e-11cf-8814-00aa00b569f5
var IID_IFillLockBytes = syscall.GUID{0x99caf010, 0x415e, 0x11cf, 
	[8]byte{0x88, 0x14, 0x00, 0xaa, 0x00, 0xb5, 0x69, 0xf5}}

type IFillLockBytesInterface interface {
	IUnknownInterface
	FillAppend(pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT
	FillAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT
	SetFillSize(ulSize uint64) HRESULT
	Terminate(bCanceled BOOL) HRESULT
}

type IFillLockBytesVtbl struct {
	IUnknownVtbl
	FillAppend uintptr
	FillAt uintptr
	SetFillSize uintptr
	Terminate uintptr
}

type IFillLockBytes struct {
	IUnknown
}

func (this *IFillLockBytes) Vtbl() *IFillLockBytesVtbl {
	return (*IFillLockBytesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFillLockBytes) FillAppend(pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FillAppend, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pv)), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

func (this *IFillLockBytes) FillAt(ulOffset uint64, pv unsafe.Pointer, cb uint32, pcbWritten *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FillAt, uintptr(unsafe.Pointer(this)), uintptr(ulOffset), uintptr(unsafe.Pointer(pv)), uintptr(cb), uintptr(unsafe.Pointer(pcbWritten)))
	return HRESULT(ret)
}

func (this *IFillLockBytes) SetFillSize(ulSize uint64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillSize, uintptr(unsafe.Pointer(this)), uintptr(ulSize))
	return HRESULT(ret)
}

func (this *IFillLockBytes) Terminate(bCanceled BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Terminate, uintptr(unsafe.Pointer(this)), uintptr(bCanceled))
	return HRESULT(ret)
}

// 0e6d4d90-6738-11cf-9608-00aa00680db4
var IID_ILayoutStorage = syscall.GUID{0x0e6d4d90, 0x6738, 0x11cf, 
	[8]byte{0x96, 0x08, 0x00, 0xaa, 0x00, 0x68, 0x0d, 0xb4}}

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
	LayoutScript uintptr
	BeginMonitor uintptr
	EndMonitor uintptr
	ReLayoutDocfile uintptr
	ReLayoutDocfileOnILockBytes uintptr
}

type ILayoutStorage struct {
	IUnknown
}

func (this *ILayoutStorage) Vtbl() *ILayoutStorageVtbl {
	return (*ILayoutStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILayoutStorage) LayoutScript(pStorageLayout *StorageLayout, nEntries uint32, glfInterleavedFlag uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().LayoutScript, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStorageLayout)), uintptr(nEntries), uintptr(glfInterleavedFlag))
	return HRESULT(ret)
}

func (this *ILayoutStorage) BeginMonitor() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().BeginMonitor, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILayoutStorage) EndMonitor() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().EndMonitor, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILayoutStorage) ReLayoutDocfile(pwcsNewDfName PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReLayoutDocfile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pwcsNewDfName)))
	return HRESULT(ret)
}

func (this *ILayoutStorage) ReLayoutDocfileOnILockBytes(pILockBytes *ILockBytes) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReLayoutDocfileOnILockBytes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pILockBytes)))
	return HRESULT(ret)
}

// 0e6d4d92-6738-11cf-9608-00aa00680db4
var IID_IDirectWriterLock = syscall.GUID{0x0e6d4d92, 0x6738, 0x11cf, 
	[8]byte{0x96, 0x08, 0x00, 0xaa, 0x00, 0x68, 0x0d, 0xb4}}

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
	HaveWriteAccess uintptr
}

type IDirectWriterLock struct {
	IUnknown
}

func (this *IDirectWriterLock) Vtbl() *IDirectWriterLockVtbl {
	return (*IDirectWriterLockVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDirectWriterLock) WaitForWriteAccess(dwTimeout uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().WaitForWriteAccess, uintptr(unsafe.Pointer(this)), uintptr(dwTimeout))
	return HRESULT(ret)
}

func (this *IDirectWriterLock) ReleaseWriteAccess() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReleaseWriteAccess, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IDirectWriterLock) HaveWriteAccess() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HaveWriteAccess, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 00000138-0000-0000-c000-000000000046
var IID_IPropertyStorage = syscall.GUID{0x00000138, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

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
	ReadMultiple uintptr
	WriteMultiple uintptr
	DeleteMultiple uintptr
	ReadPropertyNames uintptr
	WritePropertyNames uintptr
	DeletePropertyNames uintptr
	Commit uintptr
	Revert uintptr
	Enum uintptr
	SetTimes uintptr
	SetClass uintptr
	Stat uintptr
}

type IPropertyStorage struct {
	IUnknown
}

func (this *IPropertyStorage) Vtbl() *IPropertyStorageVtbl {
	return (*IPropertyStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyStorage) ReadMultiple(cpspec uint32, rgpspec *PROPSPEC, rgpropvar *PROPVARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReadMultiple, uintptr(unsafe.Pointer(this)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)), uintptr(unsafe.Pointer(rgpropvar)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) WriteMultiple(cpspec uint32, rgpspec *PROPSPEC, rgpropvar *PROPVARIANT, propidNameFirst uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteMultiple, uintptr(unsafe.Pointer(this)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)), uintptr(unsafe.Pointer(rgpropvar)), uintptr(propidNameFirst))
	return HRESULT(ret)
}

func (this *IPropertyStorage) DeleteMultiple(cpspec uint32, rgpspec *PROPSPEC) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DeleteMultiple, uintptr(unsafe.Pointer(this)), uintptr(cpspec), uintptr(unsafe.Pointer(rgpspec)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) ReadPropertyNames(cpropid uint32, rgpropid *uint32, rglpwstrName *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReadPropertyNames, uintptr(unsafe.Pointer(this)), uintptr(cpropid), uintptr(unsafe.Pointer(rgpropid)), uintptr(unsafe.Pointer(rglpwstrName)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) WritePropertyNames(cpropid uint32, rgpropid *uint32, rglpwstrName *PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().WritePropertyNames, uintptr(unsafe.Pointer(this)), uintptr(cpropid), uintptr(unsafe.Pointer(rgpropid)), uintptr(unsafe.Pointer(rglpwstrName)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) DeletePropertyNames(cpropid uint32, rgpropid *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DeletePropertyNames, uintptr(unsafe.Pointer(this)), uintptr(cpropid), uintptr(unsafe.Pointer(rgpropid)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Commit(grfCommitFlags uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Commit, uintptr(unsafe.Pointer(this)), uintptr(grfCommitFlags))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Revert() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Revert, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Enum(ppenum **IEnumSTATPROPSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Enum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) SetTimes(pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTimes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pctime)), uintptr(unsafe.Pointer(patime)), uintptr(unsafe.Pointer(pmtime)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) SetClass(clsid *syscall.GUID) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clsid)))
	return HRESULT(ret)
}

func (this *IPropertyStorage) Stat(pstatpsstg *STATPROPSETSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Stat, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstatpsstg)))
	return HRESULT(ret)
}

// 0000013a-0000-0000-c000-000000000046
var IID_IPropertySetStorage = syscall.GUID{0x0000013a, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

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
	Open uintptr
	Delete uintptr
	Enum uintptr
}

type IPropertySetStorage struct {
	IUnknown
}

func (this *IPropertySetStorage) Vtbl() *IPropertySetStorageVtbl {
	return (*IPropertySetStorageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertySetStorage) Create(rfmtid *syscall.GUID, pclsid *syscall.GUID, grfFlags uint32, grfMode uint32, ppprstg **IPropertyStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Create, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rfmtid)), uintptr(unsafe.Pointer(pclsid)), uintptr(grfFlags), uintptr(grfMode), uintptr(unsafe.Pointer(ppprstg)))
	return HRESULT(ret)
}

func (this *IPropertySetStorage) Open(rfmtid *syscall.GUID, grfMode uint32, ppprstg **IPropertyStorage) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Open, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rfmtid)), uintptr(grfMode), uintptr(unsafe.Pointer(ppprstg)))
	return HRESULT(ret)
}

func (this *IPropertySetStorage) Delete(rfmtid *syscall.GUID) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Delete, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rfmtid)))
	return HRESULT(ret)
}

func (this *IPropertySetStorage) Enum(ppenum **IEnumSTATPROPSETSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Enum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 00000139-0000-0000-c000-000000000046
var IID_IEnumSTATPROPSTG = syscall.GUID{0x00000139, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumSTATPROPSTGInterface interface {
	IUnknownInterface
	Next(celt uint32, rgelt *STATPROPSTG, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppenum **IEnumSTATPROPSTG) HRESULT
}

type IEnumSTATPROPSTGVtbl struct {
	IUnknownVtbl
	Next uintptr
	Skip uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumSTATPROPSTG struct {
	IUnknown
}

func (this *IEnumSTATPROPSTG) Vtbl() *IEnumSTATPROPSTGVtbl {
	return (*IEnumSTATPROPSTGVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumSTATPROPSTG) Next(celt uint32, rgelt *STATPROPSTG, pceltFetched *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgelt)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSTG) Skip(celt uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSTG) Reset() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSTG) Clone(ppenum **IEnumSTATPROPSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 0000013b-0000-0000-c000-000000000046
var IID_IEnumSTATPROPSETSTG = syscall.GUID{0x0000013b, 0x0000, 0x0000, 
	[8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IEnumSTATPROPSETSTGInterface interface {
	IUnknownInterface
	Next(celt uint32, rgelt *STATPROPSETSTG, pceltFetched *uint32) HRESULT
	Skip(celt uint32) HRESULT
	Reset() HRESULT
	Clone(ppenum **IEnumSTATPROPSETSTG) HRESULT
}

type IEnumSTATPROPSETSTGVtbl struct {
	IUnknownVtbl
	Next uintptr
	Skip uintptr
	Reset uintptr
	Clone uintptr
}

type IEnumSTATPROPSETSTG struct {
	IUnknown
}

func (this *IEnumSTATPROPSETSTG) Vtbl() *IEnumSTATPROPSETSTGVtbl {
	return (*IEnumSTATPROPSETSTGVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IEnumSTATPROPSETSTG) Next(celt uint32, rgelt *STATPROPSETSTG, pceltFetched *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Next, uintptr(unsafe.Pointer(this)), uintptr(celt), uintptr(unsafe.Pointer(rgelt)), uintptr(unsafe.Pointer(pceltFetched)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSETSTG) Skip(celt uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Skip, uintptr(unsafe.Pointer(this)), uintptr(celt))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSETSTG) Reset() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Reset, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IEnumSTATPROPSETSTG) Clone(ppenum **IEnumSTATPROPSETSTG) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppenum)))
	return HRESULT(ret)
}

// 55272a00-42cb-11ce-8135-00aa004bb851
var IID_IPropertyBag = syscall.GUID{0x55272a00, 0x42cb, 0x11ce, 
	[8]byte{0x81, 0x35, 0x00, 0xaa, 0x00, 0x4b, 0xb8, 0x51}}

type IPropertyBagInterface interface {
	IUnknownInterface
	Read(pszPropName PWSTR, pVar *VARIANT, pErrorLog *IErrorLog) HRESULT
	Write(pszPropName PWSTR, pVar *VARIANT) HRESULT
}

type IPropertyBagVtbl struct {
	IUnknownVtbl
	Read uintptr
	Write uintptr
}

type IPropertyBag struct {
	IUnknown
}

func (this *IPropertyBag) Vtbl() *IPropertyBagVtbl {
	return (*IPropertyBagVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyBag) Read(pszPropName PWSTR, pVar *VARIANT, pErrorLog *IErrorLog) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPropName)), uintptr(unsafe.Pointer(pVar)), uintptr(unsafe.Pointer(pErrorLog)))
	return HRESULT(ret)
}

func (this *IPropertyBag) Write(pszPropName PWSTR, pVar *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszPropName)), uintptr(unsafe.Pointer(pVar)))
	return HRESULT(ret)
}

// 22f55882-280b-11d0-a8a9-00a0c90c2004
var IID_IPropertyBag2 = syscall.GUID{0x22f55882, 0x280b, 0x11d0, 
	[8]byte{0xa8, 0xa9, 0x00, 0xa0, 0xc9, 0x0c, 0x20, 0x04}}

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
	Read uintptr
	Write uintptr
	CountProperties uintptr
	GetPropertyInfo uintptr
	LoadObject uintptr
}

type IPropertyBag2 struct {
	IUnknown
}

func (this *IPropertyBag2) Vtbl() *IPropertyBag2Vtbl {
	return (*IPropertyBag2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPropertyBag2) Read(cProperties uint32, pPropBag *PROPBAG2, pErrLog *IErrorLog, pvarValue *VARIANT, phrError *HRESULT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Read, uintptr(unsafe.Pointer(this)), uintptr(cProperties), uintptr(unsafe.Pointer(pPropBag)), uintptr(unsafe.Pointer(pErrLog)), uintptr(unsafe.Pointer(pvarValue)), uintptr(unsafe.Pointer(phrError)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) Write(cProperties uint32, pPropBag *PROPBAG2, pvarValue *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(cProperties), uintptr(unsafe.Pointer(pPropBag)), uintptr(unsafe.Pointer(pvarValue)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) CountProperties(pcProperties *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CountProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcProperties)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) GetPropertyInfo(iProperty uint32, cProperties uint32, pPropBag *PROPBAG2, pcProperties *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyInfo, uintptr(unsafe.Pointer(this)), uintptr(iProperty), uintptr(cProperties), uintptr(unsafe.Pointer(pPropBag)), uintptr(unsafe.Pointer(pcProperties)))
	return HRESULT(ret)
}

func (this *IPropertyBag2) LoadObject(pstrName PWSTR, dwHint uint32, pUnkObject *IUnknown, pErrLog *IErrorLog) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().LoadObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pstrName)), uintptr(dwHint), uintptr(unsafe.Pointer(pUnkObject)), uintptr(unsafe.Pointer(pErrLog)))
	return HRESULT(ret)
}


var (
	pCoGetInstanceFromFile uintptr
	pCoGetInstanceFromIStorage uintptr
	pStgOpenAsyncDocfileOnIFillLockBytes uintptr
	pStgGetIFillLockBytesOnILockBytes uintptr
	pStgGetIFillLockBytesOnFile uintptr
	pCreateStreamOnHGlobal uintptr
	pGetHGlobalFromStream uintptr
	pCoGetInterfaceAndReleaseStream uintptr
	pPropVariantCopy uintptr
	pPropVariantClear uintptr
	pFreePropVariantArray uintptr
	pStgCreateDocfile uintptr
	pStgCreateDocfileOnILockBytes uintptr
	pStgOpenStorage uintptr
	pStgOpenStorageOnILockBytes uintptr
	pStgIsStorageFile uintptr
	pStgIsStorageILockBytes uintptr
	pStgSetTimes uintptr
	pStgCreateStorageEx uintptr
	pStgOpenStorageEx uintptr
	pStgCreatePropStg uintptr
	pStgOpenPropStg uintptr
	pStgCreatePropSetStg uintptr
	pFmtIdToPropStgName uintptr
	pPropStgNameToFmtId uintptr
	pReadClassStg uintptr
	pWriteClassStg uintptr
	pReadClassStm uintptr
	pWriteClassStm uintptr
	pGetHGlobalFromILockBytes uintptr
	pCreateILockBytesOnHGlobal uintptr
	pGetConvertStg uintptr
	pStgConvertVariantToProperty uintptr
	pStgConvertPropertyToVariant uintptr
	pStgPropertyLengthAsVariant uintptr
	pWriteFmtUserTypeStg uintptr
	pReadFmtUserTypeStg uintptr
	pOleConvertOLESTREAMToIStorage uintptr
	pOleConvertIStorageToOLESTREAM uintptr
	pSetConvertStg uintptr
	pOleConvertIStorageToOLESTREAMEx uintptr
	pOleConvertOLESTREAMToIStorageEx uintptr
)

func CoGetInstanceFromFile(pServerInfo *COSERVERINFO, pClsid *syscall.GUID, punkOuter *IUnknown, dwClsCtx CLSCTX, grfMode uint32, pwszName PWSTR, dwCount uint32, pResults *MULTI_QI) HRESULT {
	addr := lazyAddr(&pCoGetInstanceFromFile, libOle32, "CoGetInstanceFromFile")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pServerInfo)), uintptr(unsafe.Pointer(pClsid)), uintptr(unsafe.Pointer(punkOuter)), uintptr(dwClsCtx), uintptr(grfMode), uintptr(unsafe.Pointer(pwszName)), uintptr(dwCount), uintptr(unsafe.Pointer(pResults)))
	return HRESULT(ret)
}

func CoGetInstanceFromIStorage(pServerInfo *COSERVERINFO, pClsid *syscall.GUID, punkOuter *IUnknown, dwClsCtx CLSCTX, pstg *IStorage, dwCount uint32, pResults *MULTI_QI) HRESULT {
	addr := lazyAddr(&pCoGetInstanceFromIStorage, libOle32, "CoGetInstanceFromIStorage")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pServerInfo)), uintptr(unsafe.Pointer(pClsid)), uintptr(unsafe.Pointer(punkOuter)), uintptr(dwClsCtx), uintptr(unsafe.Pointer(pstg)), uintptr(dwCount), uintptr(unsafe.Pointer(pResults)))
	return HRESULT(ret)
}

func StgOpenAsyncDocfileOnIFillLockBytes(pflb *IFillLockBytes, grfMode uint32, asyncFlags uint32, ppstgOpen **IStorage) HRESULT {
	addr := lazyAddr(&pStgOpenAsyncDocfileOnIFillLockBytes, libOle32, "StgOpenAsyncDocfileOnIFillLockBytes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pflb)), uintptr(grfMode), uintptr(asyncFlags), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgGetIFillLockBytesOnILockBytes(pilb *ILockBytes, ppflb **IFillLockBytes) HRESULT {
	addr := lazyAddr(&pStgGetIFillLockBytesOnILockBytes, libOle32, "StgGetIFillLockBytesOnILockBytes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pilb)), uintptr(unsafe.Pointer(ppflb)))
	return HRESULT(ret)
}

func StgGetIFillLockBytesOnFile(pwcsName PWSTR, ppflb **IFillLockBytes) HRESULT {
	addr := lazyAddr(&pStgGetIFillLockBytesOnFile, libOle32, "StgGetIFillLockBytesOnFile")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(ppflb)))
	return HRESULT(ret)
}

func CreateStreamOnHGlobal(hGlobal uintptr, fDeleteOnRelease BOOL, ppstm **IStream) HRESULT {
	addr := lazyAddr(&pCreateStreamOnHGlobal, libOle32, "CreateStreamOnHGlobal")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(hGlobal), uintptr(fDeleteOnRelease), uintptr(unsafe.Pointer(ppstm)))
	return HRESULT(ret)
}

func GetHGlobalFromStream(pstm *IStream, phglobal *uintptr) HRESULT {
	addr := lazyAddr(&pGetHGlobalFromStream, libOle32, "GetHGlobalFromStream")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstm)), uintptr(unsafe.Pointer(phglobal)))
	return HRESULT(ret)
}

func CoGetInterfaceAndReleaseStream(pStm *IStream, iid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := lazyAddr(&pCoGetInterfaceAndReleaseStream, libOle32, "CoGetInterfaceAndReleaseStream")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStm)), uintptr(unsafe.Pointer(iid)), uintptr(ppv))
	return HRESULT(ret)
}

func PropVariantCopy(pvarDest *PROPVARIANT, pvarSrc *PROPVARIANT) HRESULT {
	addr := lazyAddr(&pPropVariantCopy, libOle32, "PropVariantCopy")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvarDest)), uintptr(unsafe.Pointer(pvarSrc)))
	return HRESULT(ret)
}

func PropVariantClear(pvar *PROPVARIANT) HRESULT {
	addr := lazyAddr(&pPropVariantClear, libOle32, "PropVariantClear")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func FreePropVariantArray(cVariants uint32, rgvars *PROPVARIANT) HRESULT {
	addr := lazyAddr(&pFreePropVariantArray, libOle32, "FreePropVariantArray")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(cVariants), uintptr(unsafe.Pointer(rgvars)))
	return HRESULT(ret)
}

func StgCreateDocfile(pwcsName PWSTR, grfMode uint32, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := lazyAddr(&pStgCreateDocfile, libOle32, "StgCreateDocfile")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgCreateDocfileOnILockBytes(plkbyt *ILockBytes, grfMode uint32, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := lazyAddr(&pStgCreateDocfileOnILockBytes, libOle32, "StgCreateDocfileOnILockBytes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)), uintptr(grfMode), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgOpenStorage(pwcsName PWSTR, pstgPriority *IStorage, grfMode uint32, snbExclude **uint16, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := lazyAddr(&pStgOpenStorage, libOle32, "StgOpenStorage")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(unsafe.Pointer(pstgPriority)), uintptr(grfMode), uintptr(unsafe.Pointer(snbExclude)), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgOpenStorageOnILockBytes(plkbyt *ILockBytes, pstgPriority *IStorage, grfMode uint32, snbExclude **uint16, reserved uint32, ppstgOpen **IStorage) HRESULT {
	addr := lazyAddr(&pStgOpenStorageOnILockBytes, libOle32, "StgOpenStorageOnILockBytes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)), uintptr(unsafe.Pointer(pstgPriority)), uintptr(grfMode), uintptr(unsafe.Pointer(snbExclude)), uintptr(reserved), uintptr(unsafe.Pointer(ppstgOpen)))
	return HRESULT(ret)
}

func StgIsStorageFile(pwcsName PWSTR) HRESULT {
	addr := lazyAddr(&pStgIsStorageFile, libOle32, "StgIsStorageFile")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)))
	return HRESULT(ret)
}

func StgIsStorageILockBytes(plkbyt *ILockBytes) HRESULT {
	addr := lazyAddr(&pStgIsStorageILockBytes, libOle32, "StgIsStorageILockBytes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)))
	return HRESULT(ret)
}

func StgSetTimes(lpszName PWSTR, pctime *FILETIME, patime *FILETIME, pmtime *FILETIME) HRESULT {
	addr := lazyAddr(&pStgSetTimes, libOle32, "StgSetTimes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpszName)), uintptr(unsafe.Pointer(pctime)), uintptr(unsafe.Pointer(patime)), uintptr(unsafe.Pointer(pmtime)))
	return HRESULT(ret)
}

func StgCreateStorageEx(pwcsName PWSTR, grfMode uint32, stgfmt uint32, grfAttrs uint32, pStgOptions *STGOPTIONS, pSecurityDescriptor *SECURITY_DESCRIPTOR, riid *syscall.GUID, ppObjectOpen unsafe.Pointer) HRESULT {
	addr := lazyAddr(&pStgCreateStorageEx, libOle32, "StgCreateStorageEx")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(stgfmt), uintptr(grfAttrs), uintptr(unsafe.Pointer(pStgOptions)), uintptr(unsafe.Pointer(pSecurityDescriptor)), uintptr(unsafe.Pointer(riid)), uintptr(ppObjectOpen))
	return HRESULT(ret)
}

func StgOpenStorageEx(pwcsName PWSTR, grfMode uint32, stgfmt uint32, grfAttrs uint32, pStgOptions *STGOPTIONS, pSecurityDescriptor *SECURITY_DESCRIPTOR, riid *syscall.GUID, ppObjectOpen unsafe.Pointer) HRESULT {
	addr := lazyAddr(&pStgOpenStorageEx, libOle32, "StgOpenStorageEx")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pwcsName)), uintptr(grfMode), uintptr(stgfmt), uintptr(grfAttrs), uintptr(unsafe.Pointer(pStgOptions)), uintptr(unsafe.Pointer(pSecurityDescriptor)), uintptr(unsafe.Pointer(riid)), uintptr(ppObjectOpen))
	return HRESULT(ret)
}

func StgCreatePropStg(pUnk *IUnknown, fmtid *syscall.GUID, pclsid *syscall.GUID, grfFlags uint32, dwReserved uint32, ppPropStg **IPropertyStorage) HRESULT {
	addr := lazyAddr(&pStgCreatePropStg, libOle32, "StgCreatePropStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pUnk)), uintptr(unsafe.Pointer(fmtid)), uintptr(unsafe.Pointer(pclsid)), uintptr(grfFlags), uintptr(dwReserved), uintptr(unsafe.Pointer(ppPropStg)))
	return HRESULT(ret)
}

func StgOpenPropStg(pUnk *IUnknown, fmtid *syscall.GUID, grfFlags uint32, dwReserved uint32, ppPropStg **IPropertyStorage) HRESULT {
	addr := lazyAddr(&pStgOpenPropStg, libOle32, "StgOpenPropStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pUnk)), uintptr(unsafe.Pointer(fmtid)), uintptr(grfFlags), uintptr(dwReserved), uintptr(unsafe.Pointer(ppPropStg)))
	return HRESULT(ret)
}

func StgCreatePropSetStg(pStorage *IStorage, dwReserved uint32, ppPropSetStg **IPropertySetStorage) HRESULT {
	addr := lazyAddr(&pStgCreatePropSetStg, libOle32, "StgCreatePropSetStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStorage)), uintptr(dwReserved), uintptr(unsafe.Pointer(ppPropSetStg)))
	return HRESULT(ret)
}

func FmtIdToPropStgName(pfmtid *syscall.GUID, oszName PWSTR) HRESULT {
	addr := lazyAddr(&pFmtIdToPropStgName, libOle32, "FmtIdToPropStgName")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pfmtid)), uintptr(unsafe.Pointer(oszName)))
	return HRESULT(ret)
}

func PropStgNameToFmtId(oszName PWSTR, pfmtid *syscall.GUID) HRESULT {
	addr := lazyAddr(&pPropStgNameToFmtId, libOle32, "PropStgNameToFmtId")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(oszName)), uintptr(unsafe.Pointer(pfmtid)))
	return HRESULT(ret)
}

func ReadClassStg(pStg *IStorage, pclsid *syscall.GUID) HRESULT {
	addr := lazyAddr(&pReadClassStg, libOle32, "ReadClassStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)), uintptr(unsafe.Pointer(pclsid)))
	return HRESULT(ret)
}

func WriteClassStg(pStg *IStorage, rclsid *syscall.GUID) HRESULT {
	addr := lazyAddr(&pWriteClassStg, libOle32, "WriteClassStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)), uintptr(unsafe.Pointer(rclsid)))
	return HRESULT(ret)
}

func ReadClassStm(pStm *IStream, pclsid *syscall.GUID) HRESULT {
	addr := lazyAddr(&pReadClassStm, libOle32, "ReadClassStm")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStm)), uintptr(unsafe.Pointer(pclsid)))
	return HRESULT(ret)
}

func WriteClassStm(pStm *IStream, rclsid *syscall.GUID) HRESULT {
	addr := lazyAddr(&pWriteClassStm, libOle32, "WriteClassStm")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStm)), uintptr(unsafe.Pointer(rclsid)))
	return HRESULT(ret)
}

func GetHGlobalFromILockBytes(plkbyt *ILockBytes, phglobal *uintptr) HRESULT {
	addr := lazyAddr(&pGetHGlobalFromILockBytes, libOle32, "GetHGlobalFromILockBytes")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(plkbyt)), uintptr(unsafe.Pointer(phglobal)))
	return HRESULT(ret)
}

func CreateILockBytesOnHGlobal(hGlobal uintptr, fDeleteOnRelease BOOL, pplkbyt **ILockBytes) HRESULT {
	addr := lazyAddr(&pCreateILockBytesOnHGlobal, libOle32, "CreateILockBytesOnHGlobal")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(hGlobal), uintptr(fDeleteOnRelease), uintptr(unsafe.Pointer(pplkbyt)))
	return HRESULT(ret)
}

func GetConvertStg(pStg *IStorage) HRESULT {
	addr := lazyAddr(&pGetConvertStg, libOle32, "GetConvertStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)))
	return HRESULT(ret)
}

func StgConvertVariantToProperty(pvar *PROPVARIANT, CodePage uint16, pprop *SERIALIZEDPROPERTYVALUE, pcb *uint32, pid uint32, fReserved BOOLEAN, pcIndirect *uint32) *SERIALIZEDPROPERTYVALUE {
	addr := lazyAddr(&pStgConvertVariantToProperty, libOle32, "StgConvertVariantToProperty")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvar)), uintptr(CodePage), uintptr(unsafe.Pointer(pprop)), uintptr(unsafe.Pointer(pcb)), uintptr(pid), uintptr(fReserved), uintptr(unsafe.Pointer(pcIndirect)))
	return (*SERIALIZEDPROPERTYVALUE)(unsafe.Pointer(ret))
}

func StgConvertPropertyToVariant(pprop *SERIALIZEDPROPERTYVALUE, CodePage uint16, pvar *PROPVARIANT, pma *PMemoryAllocator) BOOLEAN {
	addr := lazyAddr(&pStgConvertPropertyToVariant, libOle32, "StgConvertPropertyToVariant")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pprop)), uintptr(CodePage), uintptr(unsafe.Pointer(pvar)), uintptr(unsafe.Pointer(pma)))
	return BOOLEAN(ret)
}

func StgPropertyLengthAsVariant(pProp *SERIALIZEDPROPERTYVALUE, cbProp uint32, CodePage uint16, bReserved uint8) uint32 {
	addr := lazyAddr(&pStgPropertyLengthAsVariant, libOle32, "StgPropertyLengthAsVariant")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pProp)), uintptr(cbProp), uintptr(CodePage), uintptr(bReserved))
	return uint32(ret)
}

func WriteFmtUserTypeStg(pstg *IStorage, cf uint16, lpszUserType PWSTR) HRESULT {
	addr := lazyAddr(&pWriteFmtUserTypeStg, libOle32, "WriteFmtUserTypeStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(cf), uintptr(unsafe.Pointer(lpszUserType)))
	return HRESULT(ret)
}

func ReadFmtUserTypeStg(pstg *IStorage, pcf *uint16, lplpszUserType *PWSTR) HRESULT {
	addr := lazyAddr(&pReadFmtUserTypeStg, libOle32, "ReadFmtUserTypeStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(pcf)), uintptr(unsafe.Pointer(lplpszUserType)))
	return HRESULT(ret)
}

func OleConvertOLESTREAMToIStorage(lpolestream *OLESTREAM, pstg *IStorage, ptd *DVTARGETDEVICE) HRESULT {
	addr := lazyAddr(&pOleConvertOLESTREAMToIStorage, libOle32, "OleConvertOLESTREAMToIStorage")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpolestream)), uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(ptd)))
	return HRESULT(ret)
}

func OleConvertIStorageToOLESTREAM(pstg *IStorage, lpolestream *OLESTREAM) HRESULT {
	addr := lazyAddr(&pOleConvertIStorageToOLESTREAM, libOle32, "OleConvertIStorageToOLESTREAM")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(lpolestream)))
	return HRESULT(ret)
}

func SetConvertStg(pStg *IStorage, fConvert BOOL) HRESULT {
	addr := lazyAddr(&pSetConvertStg, libOle32, "SetConvertStg")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pStg)), uintptr(fConvert))
	return HRESULT(ret)
}

func OleConvertIStorageToOLESTREAMEx(pstg *IStorage, cfFormat uint16, lWidth int32, lHeight int32, dwSize uint32, pmedium *STGMEDIUM, polestm *OLESTREAM) HRESULT {
	addr := lazyAddr(&pOleConvertIStorageToOLESTREAMEx, libOle32, "OleConvertIStorageToOLESTREAMEx")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstg)), uintptr(cfFormat), uintptr(lWidth), uintptr(lHeight), uintptr(dwSize), uintptr(unsafe.Pointer(pmedium)), uintptr(unsafe.Pointer(polestm)))
	return HRESULT(ret)
}

func OleConvertOLESTREAMToIStorageEx(polestm *OLESTREAM, pstg *IStorage, pcfFormat *uint16, plwWidth *int32, plHeight *int32, pdwSize *uint32, pmedium *STGMEDIUM) HRESULT {
	addr := lazyAddr(&pOleConvertOLESTREAMToIStorageEx, libOle32, "OleConvertOLESTREAMToIStorageEx")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(polestm)), uintptr(unsafe.Pointer(pstg)), uintptr(unsafe.Pointer(pcfFormat)), uintptr(unsafe.Pointer(plwWidth)), uintptr(unsafe.Pointer(plHeight)), uintptr(unsafe.Pointer(pdwSize)), uintptr(unsafe.Pointer(pmedium)))
	return HRESULT(ret)
}


