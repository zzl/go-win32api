package win32

import "unsafe"
import "syscall"

const (
	PERCEIVEDFLAG_UNDEFINED uint32 = 0
	PERCEIVEDFLAG_SOFTCODED uint32 = 1
	PERCEIVEDFLAG_HARDCODED uint32 = 2
	PERCEIVEDFLAG_NATIVESUPPORT uint32 = 4
	PERCEIVEDFLAG_GDIPLUS uint32 = 16
	PERCEIVEDFLAG_WMSDK uint32 = 32
	PERCEIVEDFLAG_ZIPFOLDER uint32 = 64
)

// enums

// enum STRRET_TYPE
type STRRET_TYPE int32
const (
	STRRET_WSTR STRRET_TYPE = 0
	STRRET_OFFSET STRRET_TYPE = 1
	STRRET_CSTR STRRET_TYPE = 2
)

// enum PERCEIVED
type PERCEIVED int32
const (
	PERCEIVED_TYPE_FIRST PERCEIVED = -3
	PERCEIVED_TYPE_CUSTOM PERCEIVED = -3
	PERCEIVED_TYPE_UNSPECIFIED PERCEIVED = -2
	PERCEIVED_TYPE_FOLDER PERCEIVED = -1
	PERCEIVED_TYPE_UNKNOWN PERCEIVED = 0
	PERCEIVED_TYPE_TEXT PERCEIVED = 1
	PERCEIVED_TYPE_IMAGE PERCEIVED = 2
	PERCEIVED_TYPE_AUDIO PERCEIVED = 3
	PERCEIVED_TYPE_VIDEO PERCEIVED = 4
	PERCEIVED_TYPE_COMPRESSED PERCEIVED = 5
	PERCEIVED_TYPE_DOCUMENT PERCEIVED = 6
	PERCEIVED_TYPE_SYSTEM PERCEIVED = 7
	PERCEIVED_TYPE_APPLICATION PERCEIVED = 8
	PERCEIVED_TYPE_GAMEMEDIA PERCEIVED = 9
	PERCEIVED_TYPE_CONTACTS PERCEIVED = 10
	PERCEIVED_TYPE_LAST PERCEIVED = 10
)

// enum SHCOLSTATE
type SHCOLSTATE int32
const (
	SHCOLSTATE_DEFAULT SHCOLSTATE = 0
	SHCOLSTATE_TYPE_STR SHCOLSTATE = 1
	SHCOLSTATE_TYPE_INT SHCOLSTATE = 2
	SHCOLSTATE_TYPE_DATE SHCOLSTATE = 3
	SHCOLSTATE_TYPEMASK SHCOLSTATE = 15
	SHCOLSTATE_ONBYDEFAULT SHCOLSTATE = 16
	SHCOLSTATE_SLOW SHCOLSTATE = 32
	SHCOLSTATE_EXTENDED SHCOLSTATE = 64
	SHCOLSTATE_SECONDARYUI SHCOLSTATE = 128
	SHCOLSTATE_HIDDEN SHCOLSTATE = 256
	SHCOLSTATE_PREFER_VARCMP SHCOLSTATE = 512
	SHCOLSTATE_PREFER_FMTCMP SHCOLSTATE = 1024
	SHCOLSTATE_NOSORTBYFOLDERNESS SHCOLSTATE = 2048
	SHCOLSTATE_VIEWONLY SHCOLSTATE = 65536
	SHCOLSTATE_BATCHREAD SHCOLSTATE = 131072
	SHCOLSTATE_NO_GROUPBY SHCOLSTATE = 262144
	SHCOLSTATE_FIXED_WIDTH SHCOLSTATE = 4096
	SHCOLSTATE_NODPISCALE SHCOLSTATE = 8192
	SHCOLSTATE_FIXED_RATIO SHCOLSTATE = 16384
	SHCOLSTATE_DISPLAYMASK SHCOLSTATE = 61440
)

// enum DEVICE_SCALE_FACTOR
type DEVICE_SCALE_FACTOR int32
const (
	DEVICE_SCALE_FACTOR_INVALID DEVICE_SCALE_FACTOR = 0
	SCALE_100_PERCENT DEVICE_SCALE_FACTOR = 100
	SCALE_120_PERCENT DEVICE_SCALE_FACTOR = 120
	SCALE_125_PERCENT DEVICE_SCALE_FACTOR = 125
	SCALE_140_PERCENT DEVICE_SCALE_FACTOR = 140
	SCALE_150_PERCENT DEVICE_SCALE_FACTOR = 150
	SCALE_160_PERCENT DEVICE_SCALE_FACTOR = 160
	SCALE_175_PERCENT DEVICE_SCALE_FACTOR = 175
	SCALE_180_PERCENT DEVICE_SCALE_FACTOR = 180
	SCALE_200_PERCENT DEVICE_SCALE_FACTOR = 200
	SCALE_225_PERCENT DEVICE_SCALE_FACTOR = 225
	SCALE_250_PERCENT DEVICE_SCALE_FACTOR = 250
	SCALE_300_PERCENT DEVICE_SCALE_FACTOR = 300
	SCALE_350_PERCENT DEVICE_SCALE_FACTOR = 350
	SCALE_400_PERCENT DEVICE_SCALE_FACTOR = 400
	SCALE_450_PERCENT DEVICE_SCALE_FACTOR = 450
	SCALE_500_PERCENT DEVICE_SCALE_FACTOR = 500
)


// structs

type SHITEMID struct {
	Cb uint16
	AbID [1]uint8
}

type ITEMIDLIST struct {
	Mkid SHITEMID
}

type STRRET_Anonymous_ struct {
	Data [32]uint64
}

func (this *STRRET_Anonymous_) POleStr() *PWSTR{
	return (*PWSTR)(unsafe.Pointer(this))
}

func (this *STRRET_Anonymous_) POleStrVal() PWSTR{
	return *(*PWSTR)(unsafe.Pointer(this))
}

func (this *STRRET_Anonymous_) UOffset() *uint32{
	return (*uint32)(unsafe.Pointer(this))
}

func (this *STRRET_Anonymous_) UOffsetVal() uint32{
	return *(*uint32)(unsafe.Pointer(this))
}

func (this *STRRET_Anonymous_) CStr() *[260]uint8{
	return (*[260]uint8)(unsafe.Pointer(this))
}

func (this *STRRET_Anonymous_) CStrVal() [260]uint8{
	return *(*[260]uint8)(unsafe.Pointer(this))
}

type STRRET struct {
	UType uint32
	STRRET_Anonymous_
}

type SHELLDETAILS struct {
	Fmt int32
	CxChar int32
	Str STRRET
}

type COMDLG_FILTERSPEC struct {
	PszName PWSTR
	PszSpec PWSTR
}


// coms

// 92ca9dcd-5622-4bba-a805-5e9f541bd8c9
var IID_IObjectArray = syscall.GUID{0x92ca9dcd, 0x5622, 0x4bba, 
	[8]byte{0xa8, 0x05, 0x5e, 0x9f, 0x54, 0x1b, 0xd8, 0xc9}}

type IObjectArrayInterface interface {
	IUnknownInterface
	GetCount(pcObjects *uint32) HRESULT
	GetAt(uiIndex uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT
}

type IObjectArrayVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt uintptr
}

type IObjectArray struct {
	IUnknown
}

func (this *IObjectArray) Vtbl() *IObjectArrayVtbl {
	return (*IObjectArrayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObjectArray) GetCount(pcObjects *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcObjects)))
	return HRESULT(ret)
}

func (this *IObjectArray) GetAt(uiIndex uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(uiIndex), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppv)))
	return HRESULT(ret)
}

// 5632b1a4-e38a-400a-928a-d4cd63230295
var IID_IObjectCollection = syscall.GUID{0x5632b1a4, 0xe38a, 0x400a, 
	[8]byte{0x92, 0x8a, 0xd4, 0xcd, 0x63, 0x23, 0x02, 0x95}}

type IObjectCollectionInterface interface {
	IObjectArrayInterface
	AddObject(punk *IUnknown) HRESULT
	AddFromArray(poaSource *IObjectArray) HRESULT
	RemoveObjectAt(uiIndex uint32) HRESULT
	Clear() HRESULT
}

type IObjectCollectionVtbl struct {
	IObjectArrayVtbl
	AddObject uintptr
	AddFromArray uintptr
	RemoveObjectAt uintptr
	Clear uintptr
}

type IObjectCollection struct {
	IObjectArray
}

func (this *IObjectCollection) Vtbl() *IObjectCollectionVtbl {
	return (*IObjectCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObjectCollection) AddObject(punk *IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddObject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(punk)))
	return HRESULT(ret)
}

func (this *IObjectCollection) AddFromArray(poaSource *IObjectArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(poaSource)))
	return HRESULT(ret)
}

func (this *IObjectCollection) RemoveObjectAt(uiIndex uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveObjectAt, uintptr(unsafe.Pointer(this)), uintptr(uiIndex))
	return HRESULT(ret)
}

func (this *IObjectCollection) Clear() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clear, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}


var (
)


