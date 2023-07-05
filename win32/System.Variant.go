package win32

import (
	"syscall"
	"unsafe"
)

// enums

// enum
// flags
type VAR_CHANGE_FLAGS uint16

const (
	VARIANT_NOVALUEPROP        VAR_CHANGE_FLAGS = 1
	VARIANT_ALPHABOOL          VAR_CHANGE_FLAGS = 2
	VARIANT_NOUSEROVERRIDE     VAR_CHANGE_FLAGS = 4
	VARIANT_CALENDAR_HIJRI     VAR_CHANGE_FLAGS = 8
	VARIANT_LOCALBOOL          VAR_CHANGE_FLAGS = 16
	VARIANT_CALENDAR_THAI      VAR_CHANGE_FLAGS = 32
	VARIANT_CALENDAR_GREGORIAN VAR_CHANGE_FLAGS = 64
	VARIANT_USE_NLS            VAR_CHANGE_FLAGS = 128
)

// enum
type VARENUM uint16

const (
	VT_EMPTY            VARENUM = 0
	VT_NULL             VARENUM = 1
	VT_I2               VARENUM = 2
	VT_I4               VARENUM = 3
	VT_R4               VARENUM = 4
	VT_R8               VARENUM = 5
	VT_CY               VARENUM = 6
	VT_DATE             VARENUM = 7
	VT_BSTR             VARENUM = 8
	VT_DISPATCH         VARENUM = 9
	VT_ERROR            VARENUM = 10
	VT_BOOL             VARENUM = 11
	VT_VARIANT          VARENUM = 12
	VT_UNKNOWN          VARENUM = 13
	VT_DECIMAL          VARENUM = 14
	VT_I1               VARENUM = 16
	VT_UI1              VARENUM = 17
	VT_UI2              VARENUM = 18
	VT_UI4              VARENUM = 19
	VT_I8               VARENUM = 20
	VT_UI8              VARENUM = 21
	VT_INT              VARENUM = 22
	VT_UINT             VARENUM = 23
	VT_VOID             VARENUM = 24
	VT_HRESULT          VARENUM = 25
	VT_PTR              VARENUM = 26
	VT_SAFEARRAY        VARENUM = 27
	VT_CARRAY           VARENUM = 28
	VT_USERDEFINED      VARENUM = 29
	VT_LPSTR            VARENUM = 30
	VT_LPWSTR           VARENUM = 31
	VT_RECORD           VARENUM = 36
	VT_INT_PTR          VARENUM = 37
	VT_UINT_PTR         VARENUM = 38
	VT_FILETIME         VARENUM = 64
	VT_BLOB             VARENUM = 65
	VT_STREAM           VARENUM = 66
	VT_STORAGE          VARENUM = 67
	VT_STREAMED_OBJECT  VARENUM = 68
	VT_STORED_OBJECT    VARENUM = 69
	VT_BLOB_OBJECT      VARENUM = 70
	VT_CF               VARENUM = 71
	VT_CLSID            VARENUM = 72
	VT_VERSIONED_STREAM VARENUM = 73
	VT_BSTR_BLOB        VARENUM = 4095
	VT_VECTOR           VARENUM = 4096
	VT_ARRAY            VARENUM = 8192
	VT_BYREF            VARENUM = 16384
	VT_RESERVED         VARENUM = 32768
	VT_ILLEGAL          VARENUM = 65535
	VT_ILLEGALMASKED    VARENUM = 4095
	VT_TYPEMASK         VARENUM = 4095
)

// enum
// flags
type PSTIME_FLAGS int32

const (
	PSTF_UTC   PSTIME_FLAGS = 0
	PSTF_LOCAL PSTIME_FLAGS = 1
)

// enum
// flags
type DRAWPROGRESSFLAGS int32

const (
	DPF_NONE             DRAWPROGRESSFLAGS = 0
	DPF_MARQUEE          DRAWPROGRESSFLAGS = 1
	DPF_MARQUEE_COMPLETE DRAWPROGRESSFLAGS = 2
	DPF_ERROR            DRAWPROGRESSFLAGS = 4
	DPF_WARNING          DRAWPROGRESSFLAGS = 8
	DPF_STOPPED          DRAWPROGRESSFLAGS = 16
)

// structs

type VARIANT_Anonymous_Anonymous_Anonymous_Anonymous struct {
	PvRecord unsafe.Pointer
	PRecInfo *IRecordInfo
}

type VARIANT_Anonymous_Anonymous_Anonymous struct {
	VARIANT_Anonymous_Anonymous_Anonymous_Anonymous
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) LlVal() *int64 {
	return (*int64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) LlValVal() int64 {
	return *(*int64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) LVal() *int32 {
	return (*int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) LValVal() int32 {
	return *(*int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) BVal() *byte {
	return (*byte)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) BValVal() byte {
	return *(*byte)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) IVal() *int16 {
	return (*int16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) IValVal() int16 {
	return *(*int16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) FltVal() *float32 {
	return (*float32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) FltValVal() float32 {
	return *(*float32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) DblVal() *float64 {
	return (*float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) DblValVal() float64 {
	return *(*float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) BoolVal() *VARIANT_BOOL {
	return (*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) BoolValVal() VARIANT_BOOL {
	return *(*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) OBSOLETE__VARIANT_BOOL__() *VARIANT_BOOL {
	return (*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) OBSOLETE__VARIANT_BOOL__Val() VARIANT_BOOL {
	return *(*VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Scode() *int32 {
	return (*int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) ScodeVal() int32 {
	return *(*int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) CyVal() *CY {
	return (*CY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) CyValVal() CY {
	return *(*CY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Date() *float64 {
	return (*float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) DateVal() float64 {
	return *(*float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) BstrVal() *BSTR {
	return (*BSTR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) BstrValVal() BSTR {
	return *(*BSTR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PunkVal() **IUnknown {
	return (**IUnknown)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PunkValVal() *IUnknown {
	return *(**IUnknown)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdispVal() **IDispatch {
	return (**IDispatch)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdispValVal() *IDispatch {
	return *(**IDispatch)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Parray() **SAFEARRAY {
	return (**SAFEARRAY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) ParrayVal() *SAFEARRAY {
	return *(**SAFEARRAY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PbVal() **byte {
	return (**byte)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PbValVal() *byte {
	return *(**byte)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PiVal() **int16 {
	return (**int16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PiValVal() *int16 {
	return *(**int16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PlVal() **int32 {
	return (**int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PlValVal() *int32 {
	return *(**int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PllVal() **int64 {
	return (**int64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PllValVal() *int64 {
	return *(**int64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PfltVal() **float32 {
	return (**float32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PfltValVal() *float32 {
	return *(**float32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdblVal() **float64 {
	return (**float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdblValVal() *float64 {
	return *(**float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PboolVal() **VARIANT_BOOL {
	return (**VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PboolValVal() *VARIANT_BOOL {
	return *(**VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) OBSOLETE__VARIANT_PBOOL__() **VARIANT_BOOL {
	return (**VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) OBSOLETE__VARIANT_PBOOL__Val() *VARIANT_BOOL {
	return *(**VARIANT_BOOL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Pscode() **int32 {
	return (**int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PscodeVal() *int32 {
	return *(**int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PcyVal() **CY {
	return (**CY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PcyValVal() *CY {
	return *(**CY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Pdate() **float64 {
	return (**float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdateVal() *float64 {
	return *(**float64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PbstrVal() **BSTR {
	return (**BSTR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PbstrValVal() *BSTR {
	return *(**BSTR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PpunkVal() ***IUnknown {
	return (***IUnknown)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PpunkValVal() **IUnknown {
	return *(***IUnknown)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PpdispVal() ***IDispatch {
	return (***IDispatch)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PpdispValVal() **IDispatch {
	return *(***IDispatch)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Pparray() ***SAFEARRAY {
	return (***SAFEARRAY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PparrayVal() **SAFEARRAY {
	return *(***SAFEARRAY)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PvarVal() **VARIANT {
	return (**VARIANT)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PvarValVal() *VARIANT {
	return *(**VARIANT)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Byref() *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) ByrefVal() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) CVal() *CHAR {
	return (*CHAR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) CValVal() CHAR {
	return *(*CHAR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UiVal() *uint16 {
	return (*uint16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UiValVal() uint16 {
	return *(*uint16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UlVal() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UlValVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UllVal() *uint64 {
	return (*uint64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UllValVal() uint64 {
	return *(*uint64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) IntVal() *int32 {
	return (*int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) IntValVal() int32 {
	return *(*int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UintVal() *uint32 {
	return (*uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) UintValVal() uint32 {
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdecVal() **DECIMAL {
	return (**DECIMAL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PdecValVal() *DECIMAL {
	return *(**DECIMAL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PcVal() *PSTR {
	return (*PSTR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PcValVal() PSTR {
	return *(*PSTR)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PuiVal() **uint16 {
	return (**uint16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PuiValVal() *uint16 {
	return *(**uint16)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PulVal() **uint32 {
	return (**uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PulValVal() *uint32 {
	return *(**uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PullVal() **uint64 {
	return (**uint64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PullValVal() *uint64 {
	return *(**uint64)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PintVal() **int32 {
	return (**int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PintValVal() *int32 {
	return *(**int32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PuintVal() **uint32 {
	return (**uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) PuintValVal() *uint32 {
	return *(**uint32)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) Anonymous() *VARIANT_Anonymous_Anonymous_Anonymous_Anonymous {
	return (*VARIANT_Anonymous_Anonymous_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous_Anonymous_Anonymous) AnonymousVal() VARIANT_Anonymous_Anonymous_Anonymous_Anonymous {
	return *(*VARIANT_Anonymous_Anonymous_Anonymous_Anonymous)(unsafe.Pointer(this))
}

type VARIANT_Anonymous_Anonymous struct {
	Vt         VARENUM
	WReserved1 uint16
	WReserved2 uint16
	WReserved3 uint16
	VARIANT_Anonymous_Anonymous_Anonymous
}

type VARIANT_Anonymous struct {
	VARIANT_Anonymous_Anonymous
}

func (this *VARIANT_Anonymous) Anonymous() *VARIANT_Anonymous_Anonymous {
	return (*VARIANT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous) AnonymousVal() VARIANT_Anonymous_Anonymous {
	return *(*VARIANT_Anonymous_Anonymous)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous) DecVal() *DECIMAL {
	return (*DECIMAL)(unsafe.Pointer(this))
}

func (this *VARIANT_Anonymous) DecValVal() DECIMAL {
	return *(*DECIMAL)(unsafe.Pointer(this))
}

type VARIANT struct {
	VARIANT_Anonymous
}

var (
	pVARIANT_UserSize                uintptr
	pVARIANT_UserMarshal             uintptr
	pVARIANT_UserUnmarshal           uintptr
	pVARIANT_UserFree                uintptr
	pVARIANT_UserSize64              uintptr
	pVARIANT_UserMarshal64           uintptr
	pVARIANT_UserUnmarshal64         uintptr
	pVARIANT_UserFree64              uintptr
	pDosDateTimeToVariantTime        uintptr
	pVariantTimeToDosDateTime        uintptr
	pSystemTimeToVariantTime         uintptr
	pVariantTimeToSystemTime         uintptr
	pVariantInit                     uintptr
	pVariantClear                    uintptr
	pVariantCopy                     uintptr
	pVariantCopyInd                  uintptr
	pVariantChangeType               uintptr
	pVariantChangeTypeEx             uintptr
	pInitVariantFromResource         uintptr
	pInitVariantFromBuffer           uintptr
	pInitVariantFromGUIDAsString     uintptr
	pInitVariantFromFileTime         uintptr
	pInitVariantFromFileTimeArray    uintptr
	pInitVariantFromVariantArrayElem uintptr
	pInitVariantFromBooleanArray     uintptr
	pInitVariantFromInt16Array       uintptr
	pInitVariantFromUInt16Array      uintptr
	pInitVariantFromInt32Array       uintptr
	pInitVariantFromUInt32Array      uintptr
	pInitVariantFromInt64Array       uintptr
	pInitVariantFromUInt64Array      uintptr
	pInitVariantFromDoubleArray      uintptr
	pInitVariantFromStringArray      uintptr
	pVariantToBooleanWithDefault     uintptr
	pVariantToInt16WithDefault       uintptr
	pVariantToUInt16WithDefault      uintptr
	pVariantToInt32WithDefault       uintptr
	pVariantToUInt32WithDefault      uintptr
	pVariantToInt64WithDefault       uintptr
	pVariantToUInt64WithDefault      uintptr
	pVariantToDoubleWithDefault      uintptr
	pVariantToStringWithDefault      uintptr
	pVariantToBoolean                uintptr
	pVariantToInt16                  uintptr
	pVariantToUInt16                 uintptr
	pVariantToInt32                  uintptr
	pVariantToUInt32                 uintptr
	pVariantToInt64                  uintptr
	pVariantToUInt64                 uintptr
	pVariantToDouble                 uintptr
	pVariantToBuffer                 uintptr
	pVariantToGUID                   uintptr
	pVariantToString                 uintptr
	pVariantToStringAlloc            uintptr
	pVariantToDosDateTime            uintptr
	pVariantToFileTime               uintptr
	pVariantGetElementCount          uintptr
	pVariantToBooleanArray           uintptr
	pVariantToInt16Array             uintptr
	pVariantToUInt16Array            uintptr
	pVariantToInt32Array             uintptr
	pVariantToUInt32Array            uintptr
	pVariantToInt64Array             uintptr
	pVariantToUInt64Array            uintptr
	pVariantToDoubleArray            uintptr
	pVariantToStringArray            uintptr
	pVariantToBooleanArrayAlloc      uintptr
	pVariantToInt16ArrayAlloc        uintptr
	pVariantToUInt16ArrayAlloc       uintptr
	pVariantToInt32ArrayAlloc        uintptr
	pVariantToUInt32ArrayAlloc       uintptr
	pVariantToInt64ArrayAlloc        uintptr
	pVariantToUInt64ArrayAlloc       uintptr
	pVariantToDoubleArrayAlloc       uintptr
	pVariantToStringArrayAlloc       uintptr
	pVariantGetBooleanElem           uintptr
	pVariantGetInt16Elem             uintptr
	pVariantGetUInt16Elem            uintptr
	pVariantGetInt32Elem             uintptr
	pVariantGetUInt32Elem            uintptr
	pVariantGetInt64Elem             uintptr
	pVariantGetUInt64Elem            uintptr
	pVariantGetDoubleElem            uintptr
	pVariantGetStringElem            uintptr
	pClearVariantArray               uintptr
	pVariantCompare                  uintptr
)

func VARIANT_UserSize(param0 *uint32, param1 uint32, param2 *VARIANT) uint32 {
	addr := LazyAddr(&pVARIANT_UserSize, libOleaut32, "VARIANT_UserSize")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(param1), uintptr(unsafe.Pointer(param2)))
	return uint32(ret)
}

func VARIANT_UserMarshal(param0 *uint32, param1 *byte, param2 *VARIANT) *byte {
	addr := LazyAddr(&pVARIANT_UserMarshal, libOleaut32, "VARIANT_UserMarshal")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(unsafe.Pointer(param1)), uintptr(unsafe.Pointer(param2)))
	return (*byte)(unsafe.Pointer(ret))
}

func VARIANT_UserUnmarshal(param0 *uint32, param1 *byte, param2 *VARIANT) *byte {
	addr := LazyAddr(&pVARIANT_UserUnmarshal, libOleaut32, "VARIANT_UserUnmarshal")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(unsafe.Pointer(param1)), uintptr(unsafe.Pointer(param2)))
	return (*byte)(unsafe.Pointer(ret))
}

func VARIANT_UserFree(param0 *uint32, param1 *VARIANT) {
	addr := LazyAddr(&pVARIANT_UserFree, libOleaut32, "VARIANT_UserFree")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(unsafe.Pointer(param1)))
}

func VARIANT_UserSize64(param0 *uint32, param1 uint32, param2 *VARIANT) uint32 {
	addr := LazyAddr(&pVARIANT_UserSize64, libOleaut32, "VARIANT_UserSize64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(param1), uintptr(unsafe.Pointer(param2)))
	return uint32(ret)
}

func VARIANT_UserMarshal64(param0 *uint32, param1 *byte, param2 *VARIANT) *byte {
	addr := LazyAddr(&pVARIANT_UserMarshal64, libOleaut32, "VARIANT_UserMarshal64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(unsafe.Pointer(param1)), uintptr(unsafe.Pointer(param2)))
	return (*byte)(unsafe.Pointer(ret))
}

func VARIANT_UserUnmarshal64(param0 *uint32, param1 *byte, param2 *VARIANT) *byte {
	addr := LazyAddr(&pVARIANT_UserUnmarshal64, libOleaut32, "VARIANT_UserUnmarshal64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(unsafe.Pointer(param1)), uintptr(unsafe.Pointer(param2)))
	return (*byte)(unsafe.Pointer(ret))
}

func VARIANT_UserFree64(param0 *uint32, param1 *VARIANT) {
	addr := LazyAddr(&pVARIANT_UserFree64, libOleaut32, "VARIANT_UserFree64")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(param0)), uintptr(unsafe.Pointer(param1)))
}

func DosDateTimeToVariantTime(wDosDate uint16, wDosTime uint16, pvtime *float64) int32 {
	addr := LazyAddr(&pDosDateTimeToVariantTime, libOleaut32, "DosDateTimeToVariantTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(wDosDate), uintptr(wDosTime), uintptr(unsafe.Pointer(pvtime)))
	return int32(ret)
}

func VariantTimeToDosDateTime(vtime float64, pwDosDate *uint16, pwDosTime *uint16) int32 {
	addr := LazyAddr(&pVariantTimeToDosDateTime, libOleaut32, "VariantTimeToDosDateTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(vtime), uintptr(unsafe.Pointer(pwDosDate)), uintptr(unsafe.Pointer(pwDosTime)))
	return int32(ret)
}

func SystemTimeToVariantTime(lpSystemTime *SYSTEMTIME, pvtime *float64) int32 {
	addr := LazyAddr(&pSystemTimeToVariantTime, libOleaut32, "SystemTimeToVariantTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpSystemTime)), uintptr(unsafe.Pointer(pvtime)))
	return int32(ret)
}

func VariantTimeToSystemTime(vtime float64, lpSystemTime *SYSTEMTIME) int32 {
	addr := LazyAddr(&pVariantTimeToSystemTime, libOleaut32, "VariantTimeToSystemTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(vtime), uintptr(unsafe.Pointer(lpSystemTime)))
	return int32(ret)
}

func VariantInit(pvarg *VARIANT) {
	addr := LazyAddr(&pVariantInit, libOleaut32, "VariantInit")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvarg)))
}

func VariantClear(pvarg *VARIANT) HRESULT {
	addr := LazyAddr(&pVariantClear, libOleaut32, "VariantClear")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvarg)))
	return HRESULT(ret)
}

func VariantCopy(pvargDest *VARIANT, pvargSrc *VARIANT) HRESULT {
	addr := LazyAddr(&pVariantCopy, libOleaut32, "VariantCopy")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvargDest)), uintptr(unsafe.Pointer(pvargSrc)))
	return HRESULT(ret)
}

func VariantCopyInd(pvarDest *VARIANT, pvargSrc *VARIANT) HRESULT {
	addr := LazyAddr(&pVariantCopyInd, libOleaut32, "VariantCopyInd")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvarDest)), uintptr(unsafe.Pointer(pvargSrc)))
	return HRESULT(ret)
}

func VariantChangeType(pvargDest *VARIANT, pvarSrc *VARIANT, wFlags VAR_CHANGE_FLAGS, vt VARENUM) HRESULT {
	addr := LazyAddr(&pVariantChangeType, libOleaut32, "VariantChangeType")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvargDest)), uintptr(unsafe.Pointer(pvarSrc)), uintptr(wFlags), uintptr(vt))
	return HRESULT(ret)
}

func VariantChangeTypeEx(pvargDest *VARIANT, pvarSrc *VARIANT, lcid uint32, wFlags VAR_CHANGE_FLAGS, vt VARENUM) HRESULT {
	addr := LazyAddr(&pVariantChangeTypeEx, libOleaut32, "VariantChangeTypeEx")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvargDest)), uintptr(unsafe.Pointer(pvarSrc)), uintptr(lcid), uintptr(wFlags), uintptr(vt))
	return HRESULT(ret)
}
