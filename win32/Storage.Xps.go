package win32

import (
	"syscall"
	"unsafe"
)

const (
	XPS_E_SIGREQUESTID_DUP                           HRESULT = -2142108795
	XPS_E_PACKAGE_NOT_OPENED                         HRESULT = -2142108794
	XPS_E_PACKAGE_ALREADY_OPENED                     HRESULT = -2142108793
	XPS_E_SIGNATUREID_DUP                            HRESULT = -2142108792
	XPS_E_MARKUP_COMPATIBILITY_ELEMENTS              HRESULT = -2142108791
	XPS_E_OBJECT_DETACHED                            HRESULT = -2142108790
	XPS_E_INVALID_SIGNATUREBLOCK_MARKUP              HRESULT = -2142108789
	XPS_E_INVALID_NUMBER_OF_POINTS_IN_CURVE_SEGMENTS HRESULT = -2142108160
	XPS_E_ABSOLUTE_REFERENCE                         HRESULT = -2142108159
	XPS_E_INVALID_NUMBER_OF_COLOR_CHANNELS           HRESULT = -2142108158
	XPS_E_INVALID_LANGUAGE                           HRESULT = -2142109696
	XPS_E_INVALID_NAME                               HRESULT = -2142109695
	XPS_E_INVALID_RESOURCE_KEY                       HRESULT = -2142109694
	XPS_E_INVALID_PAGE_SIZE                          HRESULT = -2142109693
	XPS_E_INVALID_BLEED_BOX                          HRESULT = -2142109692
	XPS_E_INVALID_THUMBNAIL_IMAGE_TYPE               HRESULT = -2142109691
	XPS_E_INVALID_LOOKUP_TYPE                        HRESULT = -2142109690
	XPS_E_INVALID_FLOAT                              HRESULT = -2142109689
	XPS_E_UNEXPECTED_CONTENT_TYPE                    HRESULT = -2142109688
	XPS_E_INVALID_FONT_URI                           HRESULT = -2142109686
	XPS_E_INVALID_CONTENT_BOX                        HRESULT = -2142109685
	XPS_E_INVALID_MARKUP                             HRESULT = -2142109684
	XPS_E_INVALID_XML_ENCODING                       HRESULT = -2142109683
	XPS_E_INVALID_CONTENT_TYPE                       HRESULT = -2142109682
	XPS_E_INVALID_OBFUSCATED_FONT_URI                HRESULT = -2142109681
	XPS_E_UNEXPECTED_RELATIONSHIP_TYPE               HRESULT = -2142109680
	XPS_E_UNEXPECTED_RESTRICTED_FONT_RELATIONSHIP    HRESULT = -2142109679
	XPS_E_MISSING_NAME                               HRESULT = -2142109440
	XPS_E_MISSING_LOOKUP                             HRESULT = -2142109439
	XPS_E_MISSING_GLYPHS                             HRESULT = -2142109438
	XPS_E_MISSING_SEGMENT_DATA                       HRESULT = -2142109437
	XPS_E_MISSING_COLORPROFILE                       HRESULT = -2142109436
	XPS_E_MISSING_RELATIONSHIP_TARGET                HRESULT = -2142109435
	XPS_E_MISSING_RESOURCE_RELATIONSHIP              HRESULT = -2142109434
	XPS_E_MISSING_FONTURI                            HRESULT = -2142109433
	XPS_E_MISSING_DOCUMENTSEQUENCE_RELATIONSHIP      HRESULT = -2142109432
	XPS_E_MISSING_DOCUMENT                           HRESULT = -2142109431
	XPS_E_MISSING_REFERRED_DOCUMENT                  HRESULT = -2142109430
	XPS_E_MISSING_REFERRED_PAGE                      HRESULT = -2142109429
	XPS_E_MISSING_PAGE_IN_DOCUMENT                   HRESULT = -2142109428
	XPS_E_MISSING_PAGE_IN_PAGEREFERENCE              HRESULT = -2142109427
	XPS_E_MISSING_IMAGE_IN_IMAGEBRUSH                HRESULT = -2142109426
	XPS_E_MISSING_RESOURCE_KEY                       HRESULT = -2142109425
	XPS_E_MISSING_PART_REFERENCE                     HRESULT = -2142109424
	XPS_E_MISSING_RESTRICTED_FONT_RELATIONSHIP       HRESULT = -2142109423
	XPS_E_MISSING_DISCARDCONTROL                     HRESULT = -2142109422
	XPS_E_MISSING_PART_STREAM                        HRESULT = -2142109421
	XPS_E_UNAVAILABLE_PACKAGE                        HRESULT = -2142109420
	XPS_E_DUPLICATE_RESOURCE_KEYS                    HRESULT = -2142109184
	XPS_E_MULTIPLE_RESOURCES                         HRESULT = -2142109183
	XPS_E_MULTIPLE_DOCUMENTSEQUENCE_RELATIONSHIPS    HRESULT = -2142109182
	XPS_E_MULTIPLE_THUMBNAILS_ON_PAGE                HRESULT = -2142109181
	XPS_E_MULTIPLE_THUMBNAILS_ON_PACKAGE             HRESULT = -2142109180
	XPS_E_MULTIPLE_PRINTTICKETS_ON_PAGE              HRESULT = -2142109179
	XPS_E_MULTIPLE_PRINTTICKETS_ON_DOCUMENT          HRESULT = -2142109178
	XPS_E_MULTIPLE_PRINTTICKETS_ON_DOCUMENTSEQUENCE  HRESULT = -2142109177
	XPS_E_MULTIPLE_REFERENCES_TO_PART                HRESULT = -2142109176
	XPS_E_DUPLICATE_NAMES                            HRESULT = -2142109175
	XPS_E_STRING_TOO_LONG                            HRESULT = -2142108928
	XPS_E_TOO_MANY_INDICES                           HRESULT = -2142108927
	XPS_E_MAPPING_OUT_OF_ORDER                       HRESULT = -2142108926
	XPS_E_MAPPING_OUTSIDE_STRING                     HRESULT = -2142108925
	XPS_E_MAPPING_OUTSIDE_INDICES                    HRESULT = -2142108924
	XPS_E_CARET_OUTSIDE_STRING                       HRESULT = -2142108923
	XPS_E_CARET_OUT_OF_ORDER                         HRESULT = -2142108922
	XPS_E_ODD_BIDILEVEL                              HRESULT = -2142108921
	XPS_E_ONE_TO_ONE_MAPPING_EXPECTED                HRESULT = -2142108920
	XPS_E_RESTRICTED_FONT_NOT_OBFUSCATED             HRESULT = -2142108919
	XPS_E_NEGATIVE_FLOAT                             HRESULT = -2142108918
	XPS_E_XKEY_ATTR_PRESENT_OUTSIDE_RES_DICT         HRESULT = -2142108672
	XPS_E_DICTIONARY_ITEM_NAMED                      HRESULT = -2142108671
	XPS_E_NESTED_REMOTE_DICTIONARY                   HRESULT = -2142108670
	XPS_E_INDEX_OUT_OF_RANGE                         HRESULT = -2142108416
	XPS_E_VISUAL_CIRCULAR_REF                        HRESULT = -2142108415
	XPS_E_NO_CUSTOM_OBJECTS                          HRESULT = -2142108414
	XPS_E_ALREADY_OWNED                              HRESULT = -2142108413
	XPS_E_RESOURCE_NOT_OWNED                         HRESULT = -2142108412
	XPS_E_UNEXPECTED_COLORPROFILE                    HRESULT = -2142108411
	XPS_E_COLOR_COMPONENT_OUT_OF_RANGE               HRESULT = -2142108410
	XPS_E_BOTH_PATHFIGURE_AND_ABBR_SYNTAX_PRESENT    HRESULT = -2142108409
	XPS_E_BOTH_RESOURCE_AND_SOURCEATTR_PRESENT       HRESULT = -2142108408
	XPS_E_BLEED_BOX_PAGE_DIMENSIONS_NOT_IN_SYNC      HRESULT = -2142108407
	XPS_E_RELATIONSHIP_EXTERNAL                      HRESULT = -2142108406
	XPS_E_NOT_ENOUGH_GRADIENT_STOPS                  HRESULT = -2142108405
	XPS_E_PACKAGE_WRITER_NOT_CLOSED                  HRESULT = -2142108404
)

// enums

// enum
type PRINT_WINDOW_FLAGS uint32

const (
	PW_CLIENTONLY PRINT_WINDOW_FLAGS = 1
)

// enum
type PRINTER_DEVICE_CAPABILITIES uint16

const (
	DC_BINNAMES         PRINTER_DEVICE_CAPABILITIES = 12
	DC_BINS             PRINTER_DEVICE_CAPABILITIES = 6
	DC_COLLATE          PRINTER_DEVICE_CAPABILITIES = 22
	DC_COLORDEVICE      PRINTER_DEVICE_CAPABILITIES = 32
	DC_COPIES           PRINTER_DEVICE_CAPABILITIES = 18
	DC_DRIVER           PRINTER_DEVICE_CAPABILITIES = 11
	DC_DUPLEX           PRINTER_DEVICE_CAPABILITIES = 7
	DC_ENUMRESOLUTIONS  PRINTER_DEVICE_CAPABILITIES = 13
	DC_EXTRA            PRINTER_DEVICE_CAPABILITIES = 9
	DC_FIELDS           PRINTER_DEVICE_CAPABILITIES = 1
	DC_FILEDEPENDENCIES PRINTER_DEVICE_CAPABILITIES = 14
	DC_MAXEXTENT        PRINTER_DEVICE_CAPABILITIES = 5
	DC_MEDIAREADY       PRINTER_DEVICE_CAPABILITIES = 29
	DC_MEDIATYPENAMES   PRINTER_DEVICE_CAPABILITIES = 34
	DC_MEDIATYPES       PRINTER_DEVICE_CAPABILITIES = 35
	DC_MINEXTENT        PRINTER_DEVICE_CAPABILITIES = 4
	DC_ORIENTATION      PRINTER_DEVICE_CAPABILITIES = 17
	DC_NUP              PRINTER_DEVICE_CAPABILITIES = 33
	DC_PAPERNAMES       PRINTER_DEVICE_CAPABILITIES = 16
	DC_PAPERS           PRINTER_DEVICE_CAPABILITIES = 2
	DC_PAPERSIZE        PRINTER_DEVICE_CAPABILITIES = 3
	DC_PERSONALITY      PRINTER_DEVICE_CAPABILITIES = 25
	DC_PRINTERMEM       PRINTER_DEVICE_CAPABILITIES = 28
	DC_PRINTRATE        PRINTER_DEVICE_CAPABILITIES = 26
	DC_PRINTRATEPPM     PRINTER_DEVICE_CAPABILITIES = 31
	DC_PRINTRATEUNIT    PRINTER_DEVICE_CAPABILITIES = 27
	DC_SIZE             PRINTER_DEVICE_CAPABILITIES = 8
	DC_STAPLE           PRINTER_DEVICE_CAPABILITIES = 30
	DC_TRUETYPE         PRINTER_DEVICE_CAPABILITIES = 15
	DC_VERSION          PRINTER_DEVICE_CAPABILITIES = 10
)

// enum
// flags
type PSINJECT_POINT uint16

const (
	PSINJECT_BEGINSTREAM                PSINJECT_POINT = 1
	PSINJECT_PSADOBE                    PSINJECT_POINT = 2
	PSINJECT_PAGESATEND                 PSINJECT_POINT = 3
	PSINJECT_PAGES                      PSINJECT_POINT = 4
	PSINJECT_DOCNEEDEDRES               PSINJECT_POINT = 5
	PSINJECT_DOCSUPPLIEDRES             PSINJECT_POINT = 6
	PSINJECT_PAGEORDER                  PSINJECT_POINT = 7
	PSINJECT_ORIENTATION                PSINJECT_POINT = 8
	PSINJECT_BOUNDINGBOX                PSINJECT_POINT = 9
	PSINJECT_DOCUMENTPROCESSCOLORS      PSINJECT_POINT = 10
	PSINJECT_COMMENTS                   PSINJECT_POINT = 11
	PSINJECT_BEGINDEFAULTS              PSINJECT_POINT = 12
	PSINJECT_ENDDEFAULTS                PSINJECT_POINT = 13
	PSINJECT_BEGINPROLOG                PSINJECT_POINT = 14
	PSINJECT_ENDPROLOG                  PSINJECT_POINT = 15
	PSINJECT_BEGINSETUP                 PSINJECT_POINT = 16
	PSINJECT_ENDSETUP                   PSINJECT_POINT = 17
	PSINJECT_TRAILER                    PSINJECT_POINT = 18
	PSINJECT_EOF                        PSINJECT_POINT = 19
	PSINJECT_ENDSTREAM                  PSINJECT_POINT = 20
	PSINJECT_DOCUMENTPROCESSCOLORSATEND PSINJECT_POINT = 21
	PSINJECT_PAGENUMBER                 PSINJECT_POINT = 100
	PSINJECT_BEGINPAGESETUP             PSINJECT_POINT = 101
	PSINJECT_ENDPAGESETUP               PSINJECT_POINT = 102
	PSINJECT_PAGETRAILER                PSINJECT_POINT = 103
	PSINJECT_PLATECOLOR                 PSINJECT_POINT = 104
	PSINJECT_SHOWPAGE                   PSINJECT_POINT = 105
	PSINJECT_PAGEBBOX                   PSINJECT_POINT = 106
	PSINJECT_ENDPAGECOMMENTS            PSINJECT_POINT = 107
	PSINJECT_VMSAVE                     PSINJECT_POINT = 200
	PSINJECT_VMRESTORE                  PSINJECT_POINT = 201
)

// enum
type XPS_TILE_MODE int32

const (
	XPS_TILE_MODE_NONE   XPS_TILE_MODE = 1
	XPS_TILE_MODE_TILE   XPS_TILE_MODE = 2
	XPS_TILE_MODE_FLIPX  XPS_TILE_MODE = 3
	XPS_TILE_MODE_FLIPY  XPS_TILE_MODE = 4
	XPS_TILE_MODE_FLIPXY XPS_TILE_MODE = 5
)

// enum
type XPS_COLOR_INTERPOLATION int32

const (
	XPS_COLOR_INTERPOLATION_SCRGBLINEAR XPS_COLOR_INTERPOLATION = 1
	XPS_COLOR_INTERPOLATION_SRGBLINEAR  XPS_COLOR_INTERPOLATION = 2
)

// enum
type XPS_SPREAD_METHOD int32

const (
	XPS_SPREAD_METHOD_PAD     XPS_SPREAD_METHOD = 1
	XPS_SPREAD_METHOD_REFLECT XPS_SPREAD_METHOD = 2
	XPS_SPREAD_METHOD_REPEAT  XPS_SPREAD_METHOD = 3
)

// enum
type XPS_STYLE_SIMULATION int32

const (
	XPS_STYLE_SIMULATION_NONE       XPS_STYLE_SIMULATION = 1
	XPS_STYLE_SIMULATION_ITALIC     XPS_STYLE_SIMULATION = 2
	XPS_STYLE_SIMULATION_BOLD       XPS_STYLE_SIMULATION = 3
	XPS_STYLE_SIMULATION_BOLDITALIC XPS_STYLE_SIMULATION = 4
)

// enum
type XPS_LINE_CAP int32

const (
	XPS_LINE_CAP_FLAT     XPS_LINE_CAP = 1
	XPS_LINE_CAP_ROUND    XPS_LINE_CAP = 2
	XPS_LINE_CAP_SQUARE   XPS_LINE_CAP = 3
	XPS_LINE_CAP_TRIANGLE XPS_LINE_CAP = 4
)

// enum
type XPS_DASH_CAP int32

const (
	XPS_DASH_CAP_FLAT     XPS_DASH_CAP = 1
	XPS_DASH_CAP_ROUND    XPS_DASH_CAP = 2
	XPS_DASH_CAP_SQUARE   XPS_DASH_CAP = 3
	XPS_DASH_CAP_TRIANGLE XPS_DASH_CAP = 4
)

// enum
type XPS_LINE_JOIN int32

const (
	XPS_LINE_JOIN_MITER XPS_LINE_JOIN = 1
	XPS_LINE_JOIN_BEVEL XPS_LINE_JOIN = 2
	XPS_LINE_JOIN_ROUND XPS_LINE_JOIN = 3
)

// enum
type XPS_IMAGE_TYPE int32

const (
	XPS_IMAGE_TYPE_JPEG XPS_IMAGE_TYPE = 1
	XPS_IMAGE_TYPE_PNG  XPS_IMAGE_TYPE = 2
	XPS_IMAGE_TYPE_TIFF XPS_IMAGE_TYPE = 3
	XPS_IMAGE_TYPE_WDP  XPS_IMAGE_TYPE = 4
	XPS_IMAGE_TYPE_JXR  XPS_IMAGE_TYPE = 5
)

// enum
type XPS_COLOR_TYPE int32

const (
	XPS_COLOR_TYPE_SRGB    XPS_COLOR_TYPE = 1
	XPS_COLOR_TYPE_SCRGB   XPS_COLOR_TYPE = 2
	XPS_COLOR_TYPE_CONTEXT XPS_COLOR_TYPE = 3
)

// enum
type XPS_FILL_RULE int32

const (
	XPS_FILL_RULE_EVENODD XPS_FILL_RULE = 1
	XPS_FILL_RULE_NONZERO XPS_FILL_RULE = 2
)

// enum
type XPS_SEGMENT_TYPE int32

const (
	XPS_SEGMENT_TYPE_ARC_LARGE_CLOCKWISE        XPS_SEGMENT_TYPE = 1
	XPS_SEGMENT_TYPE_ARC_LARGE_COUNTERCLOCKWISE XPS_SEGMENT_TYPE = 2
	XPS_SEGMENT_TYPE_ARC_SMALL_CLOCKWISE        XPS_SEGMENT_TYPE = 3
	XPS_SEGMENT_TYPE_ARC_SMALL_COUNTERCLOCKWISE XPS_SEGMENT_TYPE = 4
	XPS_SEGMENT_TYPE_BEZIER                     XPS_SEGMENT_TYPE = 5
	XPS_SEGMENT_TYPE_LINE                       XPS_SEGMENT_TYPE = 6
	XPS_SEGMENT_TYPE_QUADRATIC_BEZIER           XPS_SEGMENT_TYPE = 7
)

// enum
type XPS_SEGMENT_STROKE_PATTERN int32

const (
	XPS_SEGMENT_STROKE_PATTERN_ALL   XPS_SEGMENT_STROKE_PATTERN = 1
	XPS_SEGMENT_STROKE_PATTERN_NONE  XPS_SEGMENT_STROKE_PATTERN = 2
	XPS_SEGMENT_STROKE_PATTERN_MIXED XPS_SEGMENT_STROKE_PATTERN = 3
)

// enum
type XPS_FONT_EMBEDDING int32

const (
	XPS_FONT_EMBEDDING_NORMAL                  XPS_FONT_EMBEDDING = 1
	XPS_FONT_EMBEDDING_OBFUSCATED              XPS_FONT_EMBEDDING = 2
	XPS_FONT_EMBEDDING_RESTRICTED              XPS_FONT_EMBEDDING = 3
	XPS_FONT_EMBEDDING_RESTRICTED_UNOBFUSCATED XPS_FONT_EMBEDDING = 4
)

// enum
type XPS_OBJECT_TYPE int32

const (
	XPS_OBJECT_TYPE_CANVAS                XPS_OBJECT_TYPE = 1
	XPS_OBJECT_TYPE_GLYPHS                XPS_OBJECT_TYPE = 2
	XPS_OBJECT_TYPE_PATH                  XPS_OBJECT_TYPE = 3
	XPS_OBJECT_TYPE_MATRIX_TRANSFORM      XPS_OBJECT_TYPE = 4
	XPS_OBJECT_TYPE_GEOMETRY              XPS_OBJECT_TYPE = 5
	XPS_OBJECT_TYPE_SOLID_COLOR_BRUSH     XPS_OBJECT_TYPE = 6
	XPS_OBJECT_TYPE_IMAGE_BRUSH           XPS_OBJECT_TYPE = 7
	XPS_OBJECT_TYPE_LINEAR_GRADIENT_BRUSH XPS_OBJECT_TYPE = 8
	XPS_OBJECT_TYPE_RADIAL_GRADIENT_BRUSH XPS_OBJECT_TYPE = 9
	XPS_OBJECT_TYPE_VISUAL_BRUSH          XPS_OBJECT_TYPE = 10
)

// enum
type XPS_THUMBNAIL_SIZE int32

const (
	XPS_THUMBNAIL_SIZE_VERYSMALL XPS_THUMBNAIL_SIZE = 1
	XPS_THUMBNAIL_SIZE_SMALL     XPS_THUMBNAIL_SIZE = 2
	XPS_THUMBNAIL_SIZE_MEDIUM    XPS_THUMBNAIL_SIZE = 3
	XPS_THUMBNAIL_SIZE_LARGE     XPS_THUMBNAIL_SIZE = 4
)

// enum
type XPS_INTERLEAVING int32

const (
	XPS_INTERLEAVING_OFF XPS_INTERLEAVING = 1
	XPS_INTERLEAVING_ON  XPS_INTERLEAVING = 2
)

// enum
type XPS_DOCUMENT_TYPE int32

const (
	XPS_DOCUMENT_TYPE_UNSPECIFIED XPS_DOCUMENT_TYPE = 1
	XPS_DOCUMENT_TYPE_XPS         XPS_DOCUMENT_TYPE = 2
	XPS_DOCUMENT_TYPE_OPENXPS     XPS_DOCUMENT_TYPE = 3
)

// enum
type XPS_SIGNATURE_STATUS int32

const (
	XPS_SIGNATURE_STATUS_INCOMPLIANT  XPS_SIGNATURE_STATUS = 1
	XPS_SIGNATURE_STATUS_INCOMPLETE   XPS_SIGNATURE_STATUS = 2
	XPS_SIGNATURE_STATUS_BROKEN       XPS_SIGNATURE_STATUS = 3
	XPS_SIGNATURE_STATUS_QUESTIONABLE XPS_SIGNATURE_STATUS = 4
	XPS_SIGNATURE_STATUS_VALID        XPS_SIGNATURE_STATUS = 5
)

// enum
// flags
type XPS_SIGN_POLICY int32

const (
	XPS_SIGN_POLICY_NONE                    XPS_SIGN_POLICY = 0
	XPS_SIGN_POLICY_CORE_PROPERTIES         XPS_SIGN_POLICY = 1
	XPS_SIGN_POLICY_SIGNATURE_RELATIONSHIPS XPS_SIGN_POLICY = 2
	XPS_SIGN_POLICY_PRINT_TICKET            XPS_SIGN_POLICY = 4
	XPS_SIGN_POLICY_DISCARD_CONTROL         XPS_SIGN_POLICY = 8
	XPS_SIGN_POLICY_ALL                     XPS_SIGN_POLICY = 15
)

// enum
// flags
type XPS_SIGN_FLAGS int32

const (
	XPS_SIGN_FLAGS_NONE                        XPS_SIGN_FLAGS = 0
	XPS_SIGN_FLAGS_IGNORE_MARKUP_COMPATIBILITY XPS_SIGN_FLAGS = 1
)

// structs

type DRAWPATRECT struct {
	PtPosition POINT
	PtSize     POINT
	WStyle     uint16
	WPattern   uint16
}

type PSINJECTDATA struct {
	DataBytes      uint32
	InjectionPoint PSINJECT_POINT
	PageNumber     uint16
}

type PSFEATURE_OUTPUT struct {
	BPageIndependent BOOL
	BSetPageDevice   BOOL
}

type PSFEATURE_CUSTPAPER struct {
	LOrientation  int32
	LWidth        int32
	LHeight       int32
	LWidthOffset  int32
	LHeightOffset int32
}

type DOCINFOA struct {
	CbSize       int32
	LpszDocName  PSTR
	LpszOutput   PSTR
	LpszDatatype PSTR
	FwType       uint32
}

type DOCINFO = DOCINFOW
type DOCINFOW struct {
	CbSize       int32
	LpszDocName  PWSTR
	LpszOutput   PWSTR
	LpszDatatype PWSTR
	FwType       uint32
}

type XPS_POINT struct {
	X float32
	Y float32
}

type XPS_SIZE struct {
	Width  float32
	Height float32
}

type XPS_RECT struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type XPS_DASH struct {
	Length float32
	Gap    float32
}

type XPS_GLYPH_INDEX struct {
	Index            int32
	AdvanceWidth     float32
	HorizontalOffset float32
	VerticalOffset   float32
}

type XPS_GLYPH_MAPPING struct {
	UnicodeStringStart  uint32
	UnicodeStringLength uint16
	GlyphIndicesStart   uint32
	GlyphIndicesLength  uint16
}

type XPS_MATRIX struct {
	M11 float32
	M12 float32
	M21 float32
	M22 float32
	M31 float32
	M32 float32
}

type XPS_COLOR_XPS_COLOR_VALUE_SRGB struct {
	Alpha byte
	Red   byte
	Green byte
	Blue  byte
}

type XPS_COLOR_XPS_COLOR_VALUE_ScRGB struct {
	Alpha float32
	Red   float32
	Green float32
	Blue  float32
}

type XPS_COLOR_XPS_COLOR_VALUE_Context struct {
	ChannelCount byte
	Channels     [9]float32
}

type XPS_COLOR_XPS_COLOR_VALUE struct {
	Data [10]uint32
}

func (this *XPS_COLOR_XPS_COLOR_VALUE) SRGB() *XPS_COLOR_XPS_COLOR_VALUE_SRGB {
	return (*XPS_COLOR_XPS_COLOR_VALUE_SRGB)(unsafe.Pointer(this))
}

func (this *XPS_COLOR_XPS_COLOR_VALUE) SRGBVal() XPS_COLOR_XPS_COLOR_VALUE_SRGB {
	return *(*XPS_COLOR_XPS_COLOR_VALUE_SRGB)(unsafe.Pointer(this))
}

func (this *XPS_COLOR_XPS_COLOR_VALUE) ScRGB() *XPS_COLOR_XPS_COLOR_VALUE_ScRGB {
	return (*XPS_COLOR_XPS_COLOR_VALUE_ScRGB)(unsafe.Pointer(this))
}

func (this *XPS_COLOR_XPS_COLOR_VALUE) ScRGBVal() XPS_COLOR_XPS_COLOR_VALUE_ScRGB {
	return *(*XPS_COLOR_XPS_COLOR_VALUE_ScRGB)(unsafe.Pointer(this))
}

func (this *XPS_COLOR_XPS_COLOR_VALUE) Context() *XPS_COLOR_XPS_COLOR_VALUE_Context {
	return (*XPS_COLOR_XPS_COLOR_VALUE_Context)(unsafe.Pointer(this))
}

func (this *XPS_COLOR_XPS_COLOR_VALUE) ContextVal() XPS_COLOR_XPS_COLOR_VALUE_Context {
	return *(*XPS_COLOR_XPS_COLOR_VALUE_Context)(unsafe.Pointer(this))
}

type XPS_COLOR struct {
	ColorType XPS_COLOR_TYPE
	Value     XPS_COLOR_XPS_COLOR_VALUE
}

type XpsOMObjectFactory struct {
}

type XpsOMThumbnailGenerator struct {
}

type XpsSignatureManager struct {
}

// func types

type ABORTPROC = uintptr
type ABORTPROC_func = func(param0 HDC, param1 int32) BOOL

// interfaces

// 7137398F-2FC1-454D-8C6A-2C3115A16ECE
var IID_IXpsOMShareable = syscall.GUID{0x7137398F, 0x2FC1, 0x454D,
	[8]byte{0x8C, 0x6A, 0x2C, 0x31, 0x15, 0xA1, 0x6E, 0xCE}}

type IXpsOMShareableInterface interface {
	IUnknownInterface
	GetOwner(owner **IUnknown) HRESULT
	GetType(type_ *XPS_OBJECT_TYPE) HRESULT
}

type IXpsOMShareableVtbl struct {
	IUnknownVtbl
	GetOwner uintptr
	GetType  uintptr
}

type IXpsOMShareable struct {
	IUnknown
}

func (this *IXpsOMShareable) Vtbl() *IXpsOMShareableVtbl {
	return (*IXpsOMShareableVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMShareable) GetOwner(owner **IUnknown) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMShareable) GetType(type_ *XPS_OBJECT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(type_)))
	return HRESULT(ret)
}

// BC3E7333-FB0B-4AF3-A819-0B4EAAD0D2FD
var IID_IXpsOMVisual = syscall.GUID{0xBC3E7333, 0xFB0B, 0x4AF3,
	[8]byte{0xA8, 0x19, 0x0B, 0x4E, 0xAA, 0xD0, 0xD2, 0xFD}}

type IXpsOMVisualInterface interface {
	IXpsOMShareableInterface
	GetTransform(matrixTransform **IXpsOMMatrixTransform) HRESULT
	GetTransformLocal(matrixTransform **IXpsOMMatrixTransform) HRESULT
	SetTransformLocal(matrixTransform *IXpsOMMatrixTransform) HRESULT
	GetTransformLookup(key *PWSTR) HRESULT
	SetTransformLookup(key PWSTR) HRESULT
	GetClipGeometry(clipGeometry **IXpsOMGeometry) HRESULT
	GetClipGeometryLocal(clipGeometry **IXpsOMGeometry) HRESULT
	SetClipGeometryLocal(clipGeometry *IXpsOMGeometry) HRESULT
	GetClipGeometryLookup(key *PWSTR) HRESULT
	SetClipGeometryLookup(key PWSTR) HRESULT
	GetOpacity(opacity *float32) HRESULT
	SetOpacity(opacity float32) HRESULT
	GetOpacityMaskBrush(opacityMaskBrush **IXpsOMBrush) HRESULT
	GetOpacityMaskBrushLocal(opacityMaskBrush **IXpsOMBrush) HRESULT
	SetOpacityMaskBrushLocal(opacityMaskBrush *IXpsOMBrush) HRESULT
	GetOpacityMaskBrushLookup(key *PWSTR) HRESULT
	SetOpacityMaskBrushLookup(key PWSTR) HRESULT
	GetName(name *PWSTR) HRESULT
	SetName(name PWSTR) HRESULT
	GetIsHyperlinkTarget(isHyperlink *BOOL) HRESULT
	SetIsHyperlinkTarget(isHyperlink BOOL) HRESULT
	GetHyperlinkNavigateUri(hyperlinkUri **IUri) HRESULT
	SetHyperlinkNavigateUri(hyperlinkUri *IUri) HRESULT
	GetLanguage(language *PWSTR) HRESULT
	SetLanguage(language PWSTR) HRESULT
}

type IXpsOMVisualVtbl struct {
	IXpsOMShareableVtbl
	GetTransform              uintptr
	GetTransformLocal         uintptr
	SetTransformLocal         uintptr
	GetTransformLookup        uintptr
	SetTransformLookup        uintptr
	GetClipGeometry           uintptr
	GetClipGeometryLocal      uintptr
	SetClipGeometryLocal      uintptr
	GetClipGeometryLookup     uintptr
	SetClipGeometryLookup     uintptr
	GetOpacity                uintptr
	SetOpacity                uintptr
	GetOpacityMaskBrush       uintptr
	GetOpacityMaskBrushLocal  uintptr
	SetOpacityMaskBrushLocal  uintptr
	GetOpacityMaskBrushLookup uintptr
	SetOpacityMaskBrushLookup uintptr
	GetName                   uintptr
	SetName                   uintptr
	GetIsHyperlinkTarget      uintptr
	SetIsHyperlinkTarget      uintptr
	GetHyperlinkNavigateUri   uintptr
	SetHyperlinkNavigateUri   uintptr
	GetLanguage               uintptr
	SetLanguage               uintptr
}

type IXpsOMVisual struct {
	IXpsOMShareable
}

func (this *IXpsOMVisual) Vtbl() *IXpsOMVisualVtbl {
	return (*IXpsOMVisualVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMVisual) GetTransform(matrixTransform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrixTransform)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetTransformLocal(matrixTransform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrixTransform)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetTransformLocal(matrixTransform *IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrixTransform)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetTransformLookup(key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetTransformLookup(key PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetClipGeometry(clipGeometry **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetClipGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clipGeometry)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetClipGeometryLocal(clipGeometry **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetClipGeometryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clipGeometry)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetClipGeometryLocal(clipGeometry *IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetClipGeometryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clipGeometry)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetClipGeometryLookup(key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetClipGeometryLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetClipGeometryLookup(key PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetClipGeometryLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetOpacity(opacity *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOpacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(opacity)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetOpacity(opacity float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetOpacity, uintptr(unsafe.Pointer(this)), uintptr(opacity))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetOpacityMaskBrush(opacityMaskBrush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOpacityMaskBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(opacityMaskBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetOpacityMaskBrushLocal(opacityMaskBrush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOpacityMaskBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(opacityMaskBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetOpacityMaskBrushLocal(opacityMaskBrush *IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetOpacityMaskBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(opacityMaskBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetOpacityMaskBrushLookup(key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOpacityMaskBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetOpacityMaskBrushLookup(key PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetOpacityMaskBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetName(name *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetName(name PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetIsHyperlinkTarget(isHyperlink *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIsHyperlinkTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isHyperlink)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetIsHyperlinkTarget(isHyperlink BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIsHyperlinkTarget, uintptr(unsafe.Pointer(this)), uintptr(isHyperlink))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetHyperlinkNavigateUri(hyperlinkUri **IUri) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetHyperlinkNavigateUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hyperlinkUri)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetHyperlinkNavigateUri(hyperlinkUri *IUri) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHyperlinkNavigateUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hyperlinkUri)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) GetLanguage(language *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)))
	return HRESULT(ret)
}

func (this *IXpsOMVisual) SetLanguage(language PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)))
	return HRESULT(ret)
}

// 74EB2F0B-A91E-4486-AFAC-0FABECA3DFC6
var IID_IXpsOMPart = syscall.GUID{0x74EB2F0B, 0xA91E, 0x4486,
	[8]byte{0xAF, 0xAC, 0x0F, 0xAB, 0xEC, 0xA3, 0xDF, 0xC6}}

type IXpsOMPartInterface interface {
	IUnknownInterface
	GetPartName(partUri unsafe.Pointer) HRESULT
	SetPartName(partUri unsafe.Pointer) HRESULT
}

type IXpsOMPartVtbl struct {
	IUnknownVtbl
	GetPartName uintptr
	SetPartName uintptr
}

type IXpsOMPart struct {
	IUnknown
}

func (this *IXpsOMPart) Vtbl() *IXpsOMPartVtbl {
	return (*IXpsOMPartVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPart) GetPartName(partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPartName, uintptr(unsafe.Pointer(this)), uintptr(partUri))
	return HRESULT(ret)
}

func (this *IXpsOMPart) SetPartName(partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPartName, uintptr(unsafe.Pointer(this)), uintptr(partUri))
	return HRESULT(ret)
}

// A5AB8616-5B16-4B9F-9629-89B323ED7909
var IID_IXpsOMGlyphsEditor = syscall.GUID{0xA5AB8616, 0x5B16, 0x4B9F,
	[8]byte{0x96, 0x29, 0x89, 0xB3, 0x23, 0xED, 0x79, 0x09}}

type IXpsOMGlyphsEditorInterface interface {
	IUnknownInterface
	ApplyEdits() HRESULT
	GetUnicodeString(unicodeString *PWSTR) HRESULT
	SetUnicodeString(unicodeString PWSTR) HRESULT
	GetGlyphIndexCount(indexCount *uint32) HRESULT
	GetGlyphIndices(indexCount *uint32, glyphIndices *XPS_GLYPH_INDEX) HRESULT
	SetGlyphIndices(indexCount uint32, glyphIndices *XPS_GLYPH_INDEX) HRESULT
	GetGlyphMappingCount(glyphMappingCount *uint32) HRESULT
	GetGlyphMappings(glyphMappingCount *uint32, glyphMappings *XPS_GLYPH_MAPPING) HRESULT
	SetGlyphMappings(glyphMappingCount uint32, glyphMappings *XPS_GLYPH_MAPPING) HRESULT
	GetProhibitedCaretStopCount(prohibitedCaretStopCount *uint32) HRESULT
	GetProhibitedCaretStops(count *uint32, prohibitedCaretStops *uint32) HRESULT
	SetProhibitedCaretStops(count uint32, prohibitedCaretStops *uint32) HRESULT
	GetBidiLevel(bidiLevel *uint32) HRESULT
	SetBidiLevel(bidiLevel uint32) HRESULT
	GetIsSideways(isSideways *BOOL) HRESULT
	SetIsSideways(isSideways BOOL) HRESULT
	GetDeviceFontName(deviceFontName *PWSTR) HRESULT
	SetDeviceFontName(deviceFontName PWSTR) HRESULT
}

type IXpsOMGlyphsEditorVtbl struct {
	IUnknownVtbl
	ApplyEdits                  uintptr
	GetUnicodeString            uintptr
	SetUnicodeString            uintptr
	GetGlyphIndexCount          uintptr
	GetGlyphIndices             uintptr
	SetGlyphIndices             uintptr
	GetGlyphMappingCount        uintptr
	GetGlyphMappings            uintptr
	SetGlyphMappings            uintptr
	GetProhibitedCaretStopCount uintptr
	GetProhibitedCaretStops     uintptr
	SetProhibitedCaretStops     uintptr
	GetBidiLevel                uintptr
	SetBidiLevel                uintptr
	GetIsSideways               uintptr
	SetIsSideways               uintptr
	GetDeviceFontName           uintptr
	SetDeviceFontName           uintptr
}

type IXpsOMGlyphsEditor struct {
	IUnknown
}

func (this *IXpsOMGlyphsEditor) Vtbl() *IXpsOMGlyphsEditorVtbl {
	return (*IXpsOMGlyphsEditorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGlyphsEditor) ApplyEdits() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ApplyEdits, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetUnicodeString(unicodeString *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetUnicodeString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(unicodeString)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetUnicodeString(unicodeString PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetUnicodeString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(unicodeString)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetGlyphIndexCount(indexCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphIndexCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(indexCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetGlyphIndices(indexCount *uint32, glyphIndices *XPS_GLYPH_INDEX) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(indexCount)), uintptr(unsafe.Pointer(glyphIndices)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetGlyphIndices(indexCount uint32, glyphIndices *XPS_GLYPH_INDEX) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetGlyphIndices, uintptr(unsafe.Pointer(this)), uintptr(indexCount), uintptr(unsafe.Pointer(glyphIndices)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetGlyphMappingCount(glyphMappingCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphMappingCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(glyphMappingCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetGlyphMappings(glyphMappingCount *uint32, glyphMappings *XPS_GLYPH_MAPPING) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphMappings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(glyphMappingCount)), uintptr(unsafe.Pointer(glyphMappings)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetGlyphMappings(glyphMappingCount uint32, glyphMappings *XPS_GLYPH_MAPPING) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetGlyphMappings, uintptr(unsafe.Pointer(this)), uintptr(glyphMappingCount), uintptr(unsafe.Pointer(glyphMappings)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetProhibitedCaretStopCount(prohibitedCaretStopCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetProhibitedCaretStopCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(prohibitedCaretStopCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetProhibitedCaretStops(count *uint32, prohibitedCaretStops *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetProhibitedCaretStops, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)), uintptr(unsafe.Pointer(prohibitedCaretStops)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetProhibitedCaretStops(count uint32, prohibitedCaretStops *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetProhibitedCaretStops, uintptr(unsafe.Pointer(this)), uintptr(count), uintptr(unsafe.Pointer(prohibitedCaretStops)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetBidiLevel(bidiLevel *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetBidiLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bidiLevel)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetBidiLevel(bidiLevel uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetBidiLevel, uintptr(unsafe.Pointer(this)), uintptr(bidiLevel))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetIsSideways(isSideways *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIsSideways, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isSideways)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetIsSideways(isSideways BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIsSideways, uintptr(unsafe.Pointer(this)), uintptr(isSideways))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) GetDeviceFontName(deviceFontName *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceFontName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceFontName)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphsEditor) SetDeviceFontName(deviceFontName PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDeviceFontName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceFontName)))
	return HRESULT(ret)
}

// 819B3199-0A5A-4B64-BEC7-A9E17E780DE2
var IID_IXpsOMGlyphs = syscall.GUID{0x819B3199, 0x0A5A, 0x4B64,
	[8]byte{0xBE, 0xC7, 0xA9, 0xE1, 0x7E, 0x78, 0x0D, 0xE2}}

type IXpsOMGlyphsInterface interface {
	IXpsOMVisualInterface
	GetUnicodeString(unicodeString *PWSTR) HRESULT
	GetGlyphIndexCount(indexCount *uint32) HRESULT
	GetGlyphIndices(indexCount *uint32, glyphIndices *XPS_GLYPH_INDEX) HRESULT
	GetGlyphMappingCount(glyphMappingCount *uint32) HRESULT
	GetGlyphMappings(glyphMappingCount *uint32, glyphMappings *XPS_GLYPH_MAPPING) HRESULT
	GetProhibitedCaretStopCount(prohibitedCaretStopCount *uint32) HRESULT
	GetProhibitedCaretStops(prohibitedCaretStopCount *uint32, prohibitedCaretStops *uint32) HRESULT
	GetBidiLevel(bidiLevel *uint32) HRESULT
	GetIsSideways(isSideways *BOOL) HRESULT
	GetDeviceFontName(deviceFontName *PWSTR) HRESULT
	GetStyleSimulations(styleSimulations *XPS_STYLE_SIMULATION) HRESULT
	SetStyleSimulations(styleSimulations XPS_STYLE_SIMULATION) HRESULT
	GetOrigin(origin *XPS_POINT) HRESULT
	SetOrigin(origin *XPS_POINT) HRESULT
	GetFontRenderingEmSize(fontRenderingEmSize *float32) HRESULT
	SetFontRenderingEmSize(fontRenderingEmSize float32) HRESULT
	GetFontResource(fontResource **IXpsOMFontResource) HRESULT
	SetFontResource(fontResource *IXpsOMFontResource) HRESULT
	GetFontFaceIndex(fontFaceIndex *int16) HRESULT
	SetFontFaceIndex(fontFaceIndex int16) HRESULT
	GetFillBrush(fillBrush **IXpsOMBrush) HRESULT
	GetFillBrushLocal(fillBrush **IXpsOMBrush) HRESULT
	SetFillBrushLocal(fillBrush *IXpsOMBrush) HRESULT
	GetFillBrushLookup(key *PWSTR) HRESULT
	SetFillBrushLookup(key PWSTR) HRESULT
	GetGlyphsEditor(editor **IXpsOMGlyphsEditor) HRESULT
	Clone(glyphs **IXpsOMGlyphs) HRESULT
}

type IXpsOMGlyphsVtbl struct {
	IXpsOMVisualVtbl
	GetUnicodeString            uintptr
	GetGlyphIndexCount          uintptr
	GetGlyphIndices             uintptr
	GetGlyphMappingCount        uintptr
	GetGlyphMappings            uintptr
	GetProhibitedCaretStopCount uintptr
	GetProhibitedCaretStops     uintptr
	GetBidiLevel                uintptr
	GetIsSideways               uintptr
	GetDeviceFontName           uintptr
	GetStyleSimulations         uintptr
	SetStyleSimulations         uintptr
	GetOrigin                   uintptr
	SetOrigin                   uintptr
	GetFontRenderingEmSize      uintptr
	SetFontRenderingEmSize      uintptr
	GetFontResource             uintptr
	SetFontResource             uintptr
	GetFontFaceIndex            uintptr
	SetFontFaceIndex            uintptr
	GetFillBrush                uintptr
	GetFillBrushLocal           uintptr
	SetFillBrushLocal           uintptr
	GetFillBrushLookup          uintptr
	SetFillBrushLookup          uintptr
	GetGlyphsEditor             uintptr
	Clone                       uintptr
}

type IXpsOMGlyphs struct {
	IXpsOMVisual
}

func (this *IXpsOMGlyphs) Vtbl() *IXpsOMGlyphsVtbl {
	return (*IXpsOMGlyphsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGlyphs) GetUnicodeString(unicodeString *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetUnicodeString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(unicodeString)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetGlyphIndexCount(indexCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphIndexCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(indexCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetGlyphIndices(indexCount *uint32, glyphIndices *XPS_GLYPH_INDEX) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphIndices, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(indexCount)), uintptr(unsafe.Pointer(glyphIndices)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetGlyphMappingCount(glyphMappingCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphMappingCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(glyphMappingCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetGlyphMappings(glyphMappingCount *uint32, glyphMappings *XPS_GLYPH_MAPPING) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphMappings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(glyphMappingCount)), uintptr(unsafe.Pointer(glyphMappings)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetProhibitedCaretStopCount(prohibitedCaretStopCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetProhibitedCaretStopCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(prohibitedCaretStopCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetProhibitedCaretStops(prohibitedCaretStopCount *uint32, prohibitedCaretStops *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetProhibitedCaretStops, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(prohibitedCaretStopCount)), uintptr(unsafe.Pointer(prohibitedCaretStops)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetBidiLevel(bidiLevel *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetBidiLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bidiLevel)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetIsSideways(isSideways *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIsSideways, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isSideways)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetDeviceFontName(deviceFontName *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDeviceFontName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deviceFontName)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetStyleSimulations(styleSimulations *XPS_STYLE_SIMULATION) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStyleSimulations, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(styleSimulations)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetStyleSimulations(styleSimulations XPS_STYLE_SIMULATION) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStyleSimulations, uintptr(unsafe.Pointer(this)), uintptr(styleSimulations))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetOrigin(origin *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOrigin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(origin)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetOrigin(origin *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetOrigin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(origin)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetFontRenderingEmSize(fontRenderingEmSize *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFontRenderingEmSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fontRenderingEmSize)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetFontRenderingEmSize(fontRenderingEmSize float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFontRenderingEmSize, uintptr(unsafe.Pointer(this)), uintptr(fontRenderingEmSize))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetFontResource(fontResource **IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFontResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fontResource)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetFontResource(fontResource *IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFontResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fontResource)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetFontFaceIndex(fontFaceIndex *int16) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFontFaceIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fontFaceIndex)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetFontFaceIndex(fontFaceIndex int16) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFontFaceIndex, uintptr(unsafe.Pointer(this)), uintptr(fontFaceIndex))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetFillBrush(fillBrush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fillBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetFillBrushLocal(fillBrush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fillBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetFillBrushLocal(fillBrush *IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fillBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetFillBrushLookup(key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) SetFillBrushLookup(key PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) GetGlyphsEditor(editor **IXpsOMGlyphsEditor) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGlyphsEditor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(editor)))
	return HRESULT(ret)
}

func (this *IXpsOMGlyphs) Clone(glyphs **IXpsOMGlyphs) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(glyphs)))
	return HRESULT(ret)
}

// 081613F4-74EB-48F2-83B3-37A9CE2D7DC6
var IID_IXpsOMDashCollection = syscall.GUID{0x081613F4, 0x74EB, 0x48F2,
	[8]byte{0x83, 0xB3, 0x37, 0xA9, 0xCE, 0x2D, 0x7D, 0xC6}}

type IXpsOMDashCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, dash *XPS_DASH) HRESULT
	InsertAt(index uint32, dash *XPS_DASH) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, dash *XPS_DASH) HRESULT
	Append(dash *XPS_DASH) HRESULT
}

type IXpsOMDashCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMDashCollection struct {
	IUnknown
}

func (this *IXpsOMDashCollection) Vtbl() *IXpsOMDashCollectionVtbl {
	return (*IXpsOMDashCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMDashCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMDashCollection) GetAt(index uint32, dash *XPS_DASH) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(dash)))
	return HRESULT(ret)
}

func (this *IXpsOMDashCollection) InsertAt(index uint32, dash *XPS_DASH) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(dash)))
	return HRESULT(ret)
}

func (this *IXpsOMDashCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMDashCollection) SetAt(index uint32, dash *XPS_DASH) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(dash)))
	return HRESULT(ret)
}

func (this *IXpsOMDashCollection) Append(dash *XPS_DASH) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dash)))
	return HRESULT(ret)
}

// B77330FF-BB37-4501-A93E-F1B1E50BFC46
var IID_IXpsOMMatrixTransform = syscall.GUID{0xB77330FF, 0xBB37, 0x4501,
	[8]byte{0xA9, 0x3E, 0xF1, 0xB1, 0xE5, 0x0B, 0xFC, 0x46}}

type IXpsOMMatrixTransformInterface interface {
	IXpsOMShareableInterface
	GetMatrix(matrix *XPS_MATRIX) HRESULT
	SetMatrix(matrix *XPS_MATRIX) HRESULT
	Clone(matrixTransform **IXpsOMMatrixTransform) HRESULT
}

type IXpsOMMatrixTransformVtbl struct {
	IXpsOMShareableVtbl
	GetMatrix uintptr
	SetMatrix uintptr
	Clone     uintptr
}

type IXpsOMMatrixTransform struct {
	IXpsOMShareable
}

func (this *IXpsOMMatrixTransform) Vtbl() *IXpsOMMatrixTransformVtbl {
	return (*IXpsOMMatrixTransformVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMMatrixTransform) GetMatrix(matrix *XPS_MATRIX) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrix)))
	return HRESULT(ret)
}

func (this *IXpsOMMatrixTransform) SetMatrix(matrix *XPS_MATRIX) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetMatrix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrix)))
	return HRESULT(ret)
}

func (this *IXpsOMMatrixTransform) Clone(matrixTransform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrixTransform)))
	return HRESULT(ret)
}

// 64FCF3D7-4D58-44BA-AD73-A13AF6492072
var IID_IXpsOMGeometry = syscall.GUID{0x64FCF3D7, 0x4D58, 0x44BA,
	[8]byte{0xAD, 0x73, 0xA1, 0x3A, 0xF6, 0x49, 0x20, 0x72}}

type IXpsOMGeometryInterface interface {
	IXpsOMShareableInterface
	GetFigures(figures **IXpsOMGeometryFigureCollection) HRESULT
	GetFillRule(fillRule *XPS_FILL_RULE) HRESULT
	SetFillRule(fillRule XPS_FILL_RULE) HRESULT
	GetTransform(transform **IXpsOMMatrixTransform) HRESULT
	GetTransformLocal(transform **IXpsOMMatrixTransform) HRESULT
	SetTransformLocal(transform *IXpsOMMatrixTransform) HRESULT
	GetTransformLookup(lookup *PWSTR) HRESULT
	SetTransformLookup(lookup PWSTR) HRESULT
	Clone(geometry **IXpsOMGeometry) HRESULT
}

type IXpsOMGeometryVtbl struct {
	IXpsOMShareableVtbl
	GetFigures         uintptr
	GetFillRule        uintptr
	SetFillRule        uintptr
	GetTransform       uintptr
	GetTransformLocal  uintptr
	SetTransformLocal  uintptr
	GetTransformLookup uintptr
	SetTransformLookup uintptr
	Clone              uintptr
}

type IXpsOMGeometry struct {
	IXpsOMShareable
}

func (this *IXpsOMGeometry) Vtbl() *IXpsOMGeometryVtbl {
	return (*IXpsOMGeometryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGeometry) GetFigures(figures **IXpsOMGeometryFigureCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFigures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(figures)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) GetFillRule(fillRule *XPS_FILL_RULE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillRule, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fillRule)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) SetFillRule(fillRule XPS_FILL_RULE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillRule, uintptr(unsafe.Pointer(this)), uintptr(fillRule))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) GetTransform(transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) GetTransformLocal(transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) SetTransformLocal(transform *IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) GetTransformLookup(lookup *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) SetTransformLookup(lookup PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometry) Clone(geometry **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)))
	return HRESULT(ret)
}

// D410DC83-908C-443E-8947-B1795D3C165A
var IID_IXpsOMGeometryFigure = syscall.GUID{0xD410DC83, 0x908C, 0x443E,
	[8]byte{0x89, 0x47, 0xB1, 0x79, 0x5D, 0x3C, 0x16, 0x5A}}

type IXpsOMGeometryFigureInterface interface {
	IUnknownInterface
	GetOwner(owner **IXpsOMGeometry) HRESULT
	GetSegmentData(dataCount *uint32, segmentData *float32) HRESULT
	GetSegmentTypes(segmentCount *uint32, segmentTypes *XPS_SEGMENT_TYPE) HRESULT
	GetSegmentStrokes(segmentCount *uint32, segmentStrokes *BOOL) HRESULT
	SetSegments(segmentCount uint32, segmentDataCount uint32, segmentTypes *XPS_SEGMENT_TYPE, segmentData *float32, segmentStrokes *BOOL) HRESULT
	GetStartPoint(startPoint *XPS_POINT) HRESULT
	SetStartPoint(startPoint *XPS_POINT) HRESULT
	GetIsClosed(isClosed *BOOL) HRESULT
	SetIsClosed(isClosed BOOL) HRESULT
	GetIsFilled(isFilled *BOOL) HRESULT
	SetIsFilled(isFilled BOOL) HRESULT
	GetSegmentCount(segmentCount *uint32) HRESULT
	GetSegmentDataCount(segmentDataCount *uint32) HRESULT
	GetSegmentStrokePattern(segmentStrokePattern *XPS_SEGMENT_STROKE_PATTERN) HRESULT
	Clone(geometryFigure **IXpsOMGeometryFigure) HRESULT
}

type IXpsOMGeometryFigureVtbl struct {
	IUnknownVtbl
	GetOwner                uintptr
	GetSegmentData          uintptr
	GetSegmentTypes         uintptr
	GetSegmentStrokes       uintptr
	SetSegments             uintptr
	GetStartPoint           uintptr
	SetStartPoint           uintptr
	GetIsClosed             uintptr
	SetIsClosed             uintptr
	GetIsFilled             uintptr
	SetIsFilled             uintptr
	GetSegmentCount         uintptr
	GetSegmentDataCount     uintptr
	GetSegmentStrokePattern uintptr
	Clone                   uintptr
}

type IXpsOMGeometryFigure struct {
	IUnknown
}

func (this *IXpsOMGeometryFigure) Vtbl() *IXpsOMGeometryFigureVtbl {
	return (*IXpsOMGeometryFigureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGeometryFigure) GetOwner(owner **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetSegmentData(dataCount *uint32, segmentData *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSegmentData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dataCount)), uintptr(unsafe.Pointer(segmentData)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetSegmentTypes(segmentCount *uint32, segmentTypes *XPS_SEGMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSegmentTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(segmentCount)), uintptr(unsafe.Pointer(segmentTypes)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetSegmentStrokes(segmentCount *uint32, segmentStrokes *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSegmentStrokes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(segmentCount)), uintptr(unsafe.Pointer(segmentStrokes)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) SetSegments(segmentCount uint32, segmentDataCount uint32, segmentTypes *XPS_SEGMENT_TYPE, segmentData *float32, segmentStrokes *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSegments, uintptr(unsafe.Pointer(this)), uintptr(segmentCount), uintptr(segmentDataCount), uintptr(unsafe.Pointer(segmentTypes)), uintptr(unsafe.Pointer(segmentData)), uintptr(unsafe.Pointer(segmentStrokes)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetStartPoint(startPoint *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStartPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startPoint)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) SetStartPoint(startPoint *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStartPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startPoint)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetIsClosed(isClosed *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIsClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isClosed)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) SetIsClosed(isClosed BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIsClosed, uintptr(unsafe.Pointer(this)), uintptr(isClosed))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetIsFilled(isFilled *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIsFilled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isFilled)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) SetIsFilled(isFilled BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIsFilled, uintptr(unsafe.Pointer(this)), uintptr(isFilled))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetSegmentCount(segmentCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSegmentCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(segmentCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetSegmentDataCount(segmentDataCount *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSegmentDataCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(segmentDataCount)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) GetSegmentStrokePattern(segmentStrokePattern *XPS_SEGMENT_STROKE_PATTERN) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSegmentStrokePattern, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(segmentStrokePattern)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigure) Clone(geometryFigure **IXpsOMGeometryFigure) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometryFigure)))
	return HRESULT(ret)
}

// FD48C3F3-A58E-4B5A-8826-1DE54ABE72B2
var IID_IXpsOMGeometryFigureCollection = syscall.GUID{0xFD48C3F3, 0xA58E, 0x4B5A,
	[8]byte{0x88, 0x26, 0x1D, 0xE5, 0x4A, 0xBE, 0x72, 0xB2}}

type IXpsOMGeometryFigureCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, geometryFigure **IXpsOMGeometryFigure) HRESULT
	InsertAt(index uint32, geometryFigure *IXpsOMGeometryFigure) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, geometryFigure *IXpsOMGeometryFigure) HRESULT
	Append(geometryFigure *IXpsOMGeometryFigure) HRESULT
}

type IXpsOMGeometryFigureCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMGeometryFigureCollection struct {
	IUnknown
}

func (this *IXpsOMGeometryFigureCollection) Vtbl() *IXpsOMGeometryFigureCollectionVtbl {
	return (*IXpsOMGeometryFigureCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGeometryFigureCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigureCollection) GetAt(index uint32, geometryFigure **IXpsOMGeometryFigure) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(geometryFigure)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigureCollection) InsertAt(index uint32, geometryFigure *IXpsOMGeometryFigure) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(geometryFigure)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigureCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigureCollection) SetAt(index uint32, geometryFigure *IXpsOMGeometryFigure) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(geometryFigure)))
	return HRESULT(ret)
}

func (this *IXpsOMGeometryFigureCollection) Append(geometryFigure *IXpsOMGeometryFigure) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometryFigure)))
	return HRESULT(ret)
}

// 37D38BB6-3EE9-4110-9312-14B194163337
var IID_IXpsOMPath = syscall.GUID{0x37D38BB6, 0x3EE9, 0x4110,
	[8]byte{0x93, 0x12, 0x14, 0xB1, 0x94, 0x16, 0x33, 0x37}}

type IXpsOMPathInterface interface {
	IXpsOMVisualInterface
	GetGeometry(geometry **IXpsOMGeometry) HRESULT
	GetGeometryLocal(geometry **IXpsOMGeometry) HRESULT
	SetGeometryLocal(geometry *IXpsOMGeometry) HRESULT
	GetGeometryLookup(lookup *PWSTR) HRESULT
	SetGeometryLookup(lookup PWSTR) HRESULT
	GetAccessibilityShortDescription(shortDescription *PWSTR) HRESULT
	SetAccessibilityShortDescription(shortDescription PWSTR) HRESULT
	GetAccessibilityLongDescription(longDescription *PWSTR) HRESULT
	SetAccessibilityLongDescription(longDescription PWSTR) HRESULT
	GetSnapsToPixels(snapsToPixels *BOOL) HRESULT
	SetSnapsToPixels(snapsToPixels BOOL) HRESULT
	GetStrokeBrush(brush **IXpsOMBrush) HRESULT
	GetStrokeBrushLocal(brush **IXpsOMBrush) HRESULT
	SetStrokeBrushLocal(brush *IXpsOMBrush) HRESULT
	GetStrokeBrushLookup(lookup *PWSTR) HRESULT
	SetStrokeBrushLookup(lookup PWSTR) HRESULT
	GetStrokeDashes(strokeDashes **IXpsOMDashCollection) HRESULT
	GetStrokeDashCap(strokeDashCap *XPS_DASH_CAP) HRESULT
	SetStrokeDashCap(strokeDashCap XPS_DASH_CAP) HRESULT
	GetStrokeDashOffset(strokeDashOffset *float32) HRESULT
	SetStrokeDashOffset(strokeDashOffset float32) HRESULT
	GetStrokeStartLineCap(strokeStartLineCap *XPS_LINE_CAP) HRESULT
	SetStrokeStartLineCap(strokeStartLineCap XPS_LINE_CAP) HRESULT
	GetStrokeEndLineCap(strokeEndLineCap *XPS_LINE_CAP) HRESULT
	SetStrokeEndLineCap(strokeEndLineCap XPS_LINE_CAP) HRESULT
	GetStrokeLineJoin(strokeLineJoin *XPS_LINE_JOIN) HRESULT
	SetStrokeLineJoin(strokeLineJoin XPS_LINE_JOIN) HRESULT
	GetStrokeMiterLimit(strokeMiterLimit *float32) HRESULT
	SetStrokeMiterLimit(strokeMiterLimit float32) HRESULT
	GetStrokeThickness(strokeThickness *float32) HRESULT
	SetStrokeThickness(strokeThickness float32) HRESULT
	GetFillBrush(brush **IXpsOMBrush) HRESULT
	GetFillBrushLocal(brush **IXpsOMBrush) HRESULT
	SetFillBrushLocal(brush *IXpsOMBrush) HRESULT
	GetFillBrushLookup(lookup *PWSTR) HRESULT
	SetFillBrushLookup(lookup PWSTR) HRESULT
	Clone(path **IXpsOMPath) HRESULT
}

type IXpsOMPathVtbl struct {
	IXpsOMVisualVtbl
	GetGeometry                      uintptr
	GetGeometryLocal                 uintptr
	SetGeometryLocal                 uintptr
	GetGeometryLookup                uintptr
	SetGeometryLookup                uintptr
	GetAccessibilityShortDescription uintptr
	SetAccessibilityShortDescription uintptr
	GetAccessibilityLongDescription  uintptr
	SetAccessibilityLongDescription  uintptr
	GetSnapsToPixels                 uintptr
	SetSnapsToPixels                 uintptr
	GetStrokeBrush                   uintptr
	GetStrokeBrushLocal              uintptr
	SetStrokeBrushLocal              uintptr
	GetStrokeBrushLookup             uintptr
	SetStrokeBrushLookup             uintptr
	GetStrokeDashes                  uintptr
	GetStrokeDashCap                 uintptr
	SetStrokeDashCap                 uintptr
	GetStrokeDashOffset              uintptr
	SetStrokeDashOffset              uintptr
	GetStrokeStartLineCap            uintptr
	SetStrokeStartLineCap            uintptr
	GetStrokeEndLineCap              uintptr
	SetStrokeEndLineCap              uintptr
	GetStrokeLineJoin                uintptr
	SetStrokeLineJoin                uintptr
	GetStrokeMiterLimit              uintptr
	SetStrokeMiterLimit              uintptr
	GetStrokeThickness               uintptr
	SetStrokeThickness               uintptr
	GetFillBrush                     uintptr
	GetFillBrushLocal                uintptr
	SetFillBrushLocal                uintptr
	GetFillBrushLookup               uintptr
	SetFillBrushLookup               uintptr
	Clone                            uintptr
}

type IXpsOMPath struct {
	IXpsOMVisual
}

func (this *IXpsOMPath) Vtbl() *IXpsOMPathVtbl {
	return (*IXpsOMPathVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPath) GetGeometry(geometry **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetGeometryLocal(geometry **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGeometryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetGeometryLocal(geometry *IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetGeometryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetGeometryLookup(lookup *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGeometryLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetGeometryLookup(lookup PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetGeometryLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetAccessibilityShortDescription(shortDescription *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAccessibilityShortDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shortDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetAccessibilityShortDescription(shortDescription PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAccessibilityShortDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shortDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetAccessibilityLongDescription(longDescription *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAccessibilityLongDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(longDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetAccessibilityLongDescription(longDescription PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAccessibilityLongDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(longDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetSnapsToPixels(snapsToPixels *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSnapsToPixels, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(snapsToPixels)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetSnapsToPixels(snapsToPixels BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSnapsToPixels, uintptr(unsafe.Pointer(this)), uintptr(snapsToPixels))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeBrush(brush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(brush)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeBrushLocal(brush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(brush)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeBrushLocal(brush *IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(brush)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeBrushLookup(lookup *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeBrushLookup(lookup PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeDashes(strokeDashes **IXpsOMDashCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeDashes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeDashes)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeDashCap(strokeDashCap *XPS_DASH_CAP) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeDashCap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeDashCap)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeDashCap(strokeDashCap XPS_DASH_CAP) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeDashCap, uintptr(unsafe.Pointer(this)), uintptr(strokeDashCap))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeDashOffset(strokeDashOffset *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeDashOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeDashOffset)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeDashOffset(strokeDashOffset float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeDashOffset, uintptr(unsafe.Pointer(this)), uintptr(strokeDashOffset))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeStartLineCap(strokeStartLineCap *XPS_LINE_CAP) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeStartLineCap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeStartLineCap)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeStartLineCap(strokeStartLineCap XPS_LINE_CAP) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeStartLineCap, uintptr(unsafe.Pointer(this)), uintptr(strokeStartLineCap))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeEndLineCap(strokeEndLineCap *XPS_LINE_CAP) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeEndLineCap, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeEndLineCap)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeEndLineCap(strokeEndLineCap XPS_LINE_CAP) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeEndLineCap, uintptr(unsafe.Pointer(this)), uintptr(strokeEndLineCap))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeLineJoin(strokeLineJoin *XPS_LINE_JOIN) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeLineJoin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeLineJoin)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeLineJoin(strokeLineJoin XPS_LINE_JOIN) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeLineJoin, uintptr(unsafe.Pointer(this)), uintptr(strokeLineJoin))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeMiterLimit(strokeMiterLimit *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeMiterLimit, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeMiterLimit)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeMiterLimit(strokeMiterLimit float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeMiterLimit, uintptr(unsafe.Pointer(this)), uintptr(strokeMiterLimit))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetStrokeThickness(strokeThickness *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStrokeThickness, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(strokeThickness)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetStrokeThickness(strokeThickness float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStrokeThickness, uintptr(unsafe.Pointer(this)), uintptr(strokeThickness))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetFillBrush(brush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(brush)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetFillBrushLocal(brush **IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(brush)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetFillBrushLocal(brush *IXpsOMBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillBrushLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(brush)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) GetFillBrushLookup(lookup *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFillBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) SetFillBrushLookup(lookup PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFillBrushLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMPath) Clone(path **IXpsOMPath) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(path)))
	return HRESULT(ret)
}

// 56A3F80C-EA4C-4187-A57B-A2A473B2B42B
var IID_IXpsOMBrush = syscall.GUID{0x56A3F80C, 0xEA4C, 0x4187,
	[8]byte{0xA5, 0x7B, 0xA2, 0xA4, 0x73, 0xB2, 0xB4, 0x2B}}

type IXpsOMBrushInterface interface {
	IXpsOMShareableInterface
	GetOpacity(opacity *float32) HRESULT
	SetOpacity(opacity float32) HRESULT
}

type IXpsOMBrushVtbl struct {
	IXpsOMShareableVtbl
	GetOpacity uintptr
	SetOpacity uintptr
}

type IXpsOMBrush struct {
	IXpsOMShareable
}

func (this *IXpsOMBrush) Vtbl() *IXpsOMBrushVtbl {
	return (*IXpsOMBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMBrush) GetOpacity(opacity *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOpacity, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(opacity)))
	return HRESULT(ret)
}

func (this *IXpsOMBrush) SetOpacity(opacity float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetOpacity, uintptr(unsafe.Pointer(this)), uintptr(opacity))
	return HRESULT(ret)
}

// C9174C3A-3CD3-4319-BDA4-11A39392CEEF
var IID_IXpsOMGradientStopCollection = syscall.GUID{0xC9174C3A, 0x3CD3, 0x4319,
	[8]byte{0xBD, 0xA4, 0x11, 0xA3, 0x93, 0x92, 0xCE, 0xEF}}

type IXpsOMGradientStopCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, stop **IXpsOMGradientStop) HRESULT
	InsertAt(index uint32, stop *IXpsOMGradientStop) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, stop *IXpsOMGradientStop) HRESULT
	Append(stop *IXpsOMGradientStop) HRESULT
}

type IXpsOMGradientStopCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMGradientStopCollection struct {
	IUnknown
}

func (this *IXpsOMGradientStopCollection) Vtbl() *IXpsOMGradientStopCollectionVtbl {
	return (*IXpsOMGradientStopCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGradientStopCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStopCollection) GetAt(index uint32, stop **IXpsOMGradientStop) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(stop)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStopCollection) InsertAt(index uint32, stop *IXpsOMGradientStop) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(stop)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStopCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStopCollection) SetAt(index uint32, stop *IXpsOMGradientStop) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(stop)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStopCollection) Append(stop *IXpsOMGradientStop) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stop)))
	return HRESULT(ret)
}

// A06F9F05-3BE9-4763-98A8-094FC672E488
var IID_IXpsOMSolidColorBrush = syscall.GUID{0xA06F9F05, 0x3BE9, 0x4763,
	[8]byte{0x98, 0xA8, 0x09, 0x4F, 0xC6, 0x72, 0xE4, 0x88}}

type IXpsOMSolidColorBrushInterface interface {
	IXpsOMBrushInterface
	GetColor(color *XPS_COLOR, colorProfile **IXpsOMColorProfileResource) HRESULT
	SetColor(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource) HRESULT
	Clone(solidColorBrush **IXpsOMSolidColorBrush) HRESULT
}

type IXpsOMSolidColorBrushVtbl struct {
	IXpsOMBrushVtbl
	GetColor uintptr
	SetColor uintptr
	Clone    uintptr
}

type IXpsOMSolidColorBrush struct {
	IXpsOMBrush
}

func (this *IXpsOMSolidColorBrush) Vtbl() *IXpsOMSolidColorBrushVtbl {
	return (*IXpsOMSolidColorBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMSolidColorBrush) GetColor(color *XPS_COLOR, colorProfile **IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(colorProfile)))
	return HRESULT(ret)
}

func (this *IXpsOMSolidColorBrush) SetColor(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(colorProfile)))
	return HRESULT(ret)
}

func (this *IXpsOMSolidColorBrush) Clone(solidColorBrush **IXpsOMSolidColorBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(solidColorBrush)))
	return HRESULT(ret)
}

// 0FC2328D-D722-4A54-B2EC-BE90218A789E
var IID_IXpsOMTileBrush = syscall.GUID{0x0FC2328D, 0xD722, 0x4A54,
	[8]byte{0xB2, 0xEC, 0xBE, 0x90, 0x21, 0x8A, 0x78, 0x9E}}

type IXpsOMTileBrushInterface interface {
	IXpsOMBrushInterface
	GetTransform(transform **IXpsOMMatrixTransform) HRESULT
	GetTransformLocal(transform **IXpsOMMatrixTransform) HRESULT
	SetTransformLocal(transform *IXpsOMMatrixTransform) HRESULT
	GetTransformLookup(key *PWSTR) HRESULT
	SetTransformLookup(key PWSTR) HRESULT
	GetViewbox(viewbox *XPS_RECT) HRESULT
	SetViewbox(viewbox *XPS_RECT) HRESULT
	GetViewport(viewport *XPS_RECT) HRESULT
	SetViewport(viewport *XPS_RECT) HRESULT
	GetTileMode(tileMode *XPS_TILE_MODE) HRESULT
	SetTileMode(tileMode XPS_TILE_MODE) HRESULT
}

type IXpsOMTileBrushVtbl struct {
	IXpsOMBrushVtbl
	GetTransform       uintptr
	GetTransformLocal  uintptr
	SetTransformLocal  uintptr
	GetTransformLookup uintptr
	SetTransformLookup uintptr
	GetViewbox         uintptr
	SetViewbox         uintptr
	GetViewport        uintptr
	SetViewport        uintptr
	GetTileMode        uintptr
	SetTileMode        uintptr
}

type IXpsOMTileBrush struct {
	IXpsOMBrush
}

func (this *IXpsOMTileBrush) Vtbl() *IXpsOMTileBrushVtbl {
	return (*IXpsOMTileBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMTileBrush) GetTransform(transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) GetTransformLocal(transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) SetTransformLocal(transform *IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) GetTransformLookup(key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) SetTransformLookup(key PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) GetViewbox(viewbox *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetViewbox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewbox)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) SetViewbox(viewbox *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetViewbox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewbox)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) GetViewport(viewport *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetViewport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewport)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) SetViewport(viewport *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetViewport, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewport)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) GetTileMode(tileMode *XPS_TILE_MODE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTileMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(tileMode)))
	return HRESULT(ret)
}

func (this *IXpsOMTileBrush) SetTileMode(tileMode XPS_TILE_MODE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTileMode, uintptr(unsafe.Pointer(this)), uintptr(tileMode))
	return HRESULT(ret)
}

// 97E294AF-5B37-46B4-8057-874D2F64119B
var IID_IXpsOMVisualBrush = syscall.GUID{0x97E294AF, 0x5B37, 0x46B4,
	[8]byte{0x80, 0x57, 0x87, 0x4D, 0x2F, 0x64, 0x11, 0x9B}}

type IXpsOMVisualBrushInterface interface {
	IXpsOMTileBrushInterface
	GetVisual(visual **IXpsOMVisual) HRESULT
	GetVisualLocal(visual **IXpsOMVisual) HRESULT
	SetVisualLocal(visual *IXpsOMVisual) HRESULT
	GetVisualLookup(lookup *PWSTR) HRESULT
	SetVisualLookup(lookup PWSTR) HRESULT
	Clone(visualBrush **IXpsOMVisualBrush) HRESULT
}

type IXpsOMVisualBrushVtbl struct {
	IXpsOMTileBrushVtbl
	GetVisual       uintptr
	GetVisualLocal  uintptr
	SetVisualLocal  uintptr
	GetVisualLookup uintptr
	SetVisualLookup uintptr
	Clone           uintptr
}

type IXpsOMVisualBrush struct {
	IXpsOMTileBrush
}

func (this *IXpsOMVisualBrush) Vtbl() *IXpsOMVisualBrushVtbl {
	return (*IXpsOMVisualBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMVisualBrush) GetVisual(visual **IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisual, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visual)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualBrush) GetVisualLocal(visual **IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisualLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visual)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualBrush) SetVisualLocal(visual *IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetVisualLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visual)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualBrush) GetVisualLookup(lookup *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisualLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualBrush) SetVisualLookup(lookup PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetVisualLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lookup)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualBrush) Clone(visualBrush **IXpsOMVisualBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visualBrush)))
	return HRESULT(ret)
}

// 3DF0B466-D382-49EF-8550-DD94C80242E4
var IID_IXpsOMImageBrush = syscall.GUID{0x3DF0B466, 0xD382, 0x49EF,
	[8]byte{0x85, 0x50, 0xDD, 0x94, 0xC8, 0x02, 0x42, 0xE4}}

type IXpsOMImageBrushInterface interface {
	IXpsOMTileBrushInterface
	GetImageResource(imageResource **IXpsOMImageResource) HRESULT
	SetImageResource(imageResource *IXpsOMImageResource) HRESULT
	GetColorProfileResource(colorProfileResource **IXpsOMColorProfileResource) HRESULT
	SetColorProfileResource(colorProfileResource *IXpsOMColorProfileResource) HRESULT
	Clone(imageBrush **IXpsOMImageBrush) HRESULT
}

type IXpsOMImageBrushVtbl struct {
	IXpsOMTileBrushVtbl
	GetImageResource        uintptr
	SetImageResource        uintptr
	GetColorProfileResource uintptr
	SetColorProfileResource uintptr
	Clone                   uintptr
}

type IXpsOMImageBrush struct {
	IXpsOMTileBrush
}

func (this *IXpsOMImageBrush) Vtbl() *IXpsOMImageBrushVtbl {
	return (*IXpsOMImageBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMImageBrush) GetImageResource(imageResource **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetImageResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMImageBrush) SetImageResource(imageResource *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetImageResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMImageBrush) GetColorProfileResource(colorProfileResource **IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColorProfileResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(colorProfileResource)))
	return HRESULT(ret)
}

func (this *IXpsOMImageBrush) SetColorProfileResource(colorProfileResource *IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetColorProfileResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(colorProfileResource)))
	return HRESULT(ret)
}

func (this *IXpsOMImageBrush) Clone(imageBrush **IXpsOMImageBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageBrush)))
	return HRESULT(ret)
}

// 5CF4F5CC-3969-49B5-A70A-5550B618FE49
var IID_IXpsOMGradientStop = syscall.GUID{0x5CF4F5CC, 0x3969, 0x49B5,
	[8]byte{0xA7, 0x0A, 0x55, 0x50, 0xB6, 0x18, 0xFE, 0x49}}

type IXpsOMGradientStopInterface interface {
	IUnknownInterface
	GetOwner(owner **IXpsOMGradientBrush) HRESULT
	GetOffset(offset *float32) HRESULT
	SetOffset(offset float32) HRESULT
	GetColor(color *XPS_COLOR, colorProfile **IXpsOMColorProfileResource) HRESULT
	SetColor(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource) HRESULT
	Clone(gradientStop **IXpsOMGradientStop) HRESULT
}

type IXpsOMGradientStopVtbl struct {
	IUnknownVtbl
	GetOwner  uintptr
	GetOffset uintptr
	SetOffset uintptr
	GetColor  uintptr
	SetColor  uintptr
	Clone     uintptr
}

type IXpsOMGradientStop struct {
	IUnknown
}

func (this *IXpsOMGradientStop) Vtbl() *IXpsOMGradientStopVtbl {
	return (*IXpsOMGradientStopVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGradientStop) GetOwner(owner **IXpsOMGradientBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStop) GetOffset(offset *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOffset, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(offset)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStop) SetOffset(offset float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetOffset, uintptr(unsafe.Pointer(this)), uintptr(offset))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStop) GetColor(color *XPS_COLOR, colorProfile **IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(colorProfile)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStop) SetColor(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(colorProfile)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientStop) Clone(gradientStop **IXpsOMGradientStop) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gradientStop)))
	return HRESULT(ret)
}

// EDB59622-61A2-42C3-BACE-ACF2286C06BF
var IID_IXpsOMGradientBrush = syscall.GUID{0xEDB59622, 0x61A2, 0x42C3,
	[8]byte{0xBA, 0xCE, 0xAC, 0xF2, 0x28, 0x6C, 0x06, 0xBF}}

type IXpsOMGradientBrushInterface interface {
	IXpsOMBrushInterface
	GetGradientStops(gradientStops **IXpsOMGradientStopCollection) HRESULT
	GetTransform(transform **IXpsOMMatrixTransform) HRESULT
	GetTransformLocal(transform **IXpsOMMatrixTransform) HRESULT
	SetTransformLocal(transform *IXpsOMMatrixTransform) HRESULT
	GetTransformLookup(key *PWSTR) HRESULT
	SetTransformLookup(key PWSTR) HRESULT
	GetSpreadMethod(spreadMethod *XPS_SPREAD_METHOD) HRESULT
	SetSpreadMethod(spreadMethod XPS_SPREAD_METHOD) HRESULT
	GetColorInterpolationMode(colorInterpolationMode *XPS_COLOR_INTERPOLATION) HRESULT
	SetColorInterpolationMode(colorInterpolationMode XPS_COLOR_INTERPOLATION) HRESULT
}

type IXpsOMGradientBrushVtbl struct {
	IXpsOMBrushVtbl
	GetGradientStops          uintptr
	GetTransform              uintptr
	GetTransformLocal         uintptr
	SetTransformLocal         uintptr
	GetTransformLookup        uintptr
	SetTransformLookup        uintptr
	GetSpreadMethod           uintptr
	SetSpreadMethod           uintptr
	GetColorInterpolationMode uintptr
	SetColorInterpolationMode uintptr
}

type IXpsOMGradientBrush struct {
	IXpsOMBrush
}

func (this *IXpsOMGradientBrush) Vtbl() *IXpsOMGradientBrushVtbl {
	return (*IXpsOMGradientBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMGradientBrush) GetGradientStops(gradientStops **IXpsOMGradientStopCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGradientStops, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gradientStops)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) GetTransform(transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) GetTransformLocal(transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) SetTransformLocal(transform *IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) GetTransformLookup(key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) SetTransformLookup(key PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTransformLookup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) GetSpreadMethod(spreadMethod *XPS_SPREAD_METHOD) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSpreadMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(spreadMethod)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) SetSpreadMethod(spreadMethod XPS_SPREAD_METHOD) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSpreadMethod, uintptr(unsafe.Pointer(this)), uintptr(spreadMethod))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) GetColorInterpolationMode(colorInterpolationMode *XPS_COLOR_INTERPOLATION) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColorInterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(colorInterpolationMode)))
	return HRESULT(ret)
}

func (this *IXpsOMGradientBrush) SetColorInterpolationMode(colorInterpolationMode XPS_COLOR_INTERPOLATION) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetColorInterpolationMode, uintptr(unsafe.Pointer(this)), uintptr(colorInterpolationMode))
	return HRESULT(ret)
}

// 005E279F-C30D-40FF-93EC-1950D3C528DB
var IID_IXpsOMLinearGradientBrush = syscall.GUID{0x005E279F, 0xC30D, 0x40FF,
	[8]byte{0x93, 0xEC, 0x19, 0x50, 0xD3, 0xC5, 0x28, 0xDB}}

type IXpsOMLinearGradientBrushInterface interface {
	IXpsOMGradientBrushInterface
	GetStartPoint(startPoint *XPS_POINT) HRESULT
	SetStartPoint(startPoint *XPS_POINT) HRESULT
	GetEndPoint(endPoint *XPS_POINT) HRESULT
	SetEndPoint(endPoint *XPS_POINT) HRESULT
	Clone(linearGradientBrush **IXpsOMLinearGradientBrush) HRESULT
}

type IXpsOMLinearGradientBrushVtbl struct {
	IXpsOMGradientBrushVtbl
	GetStartPoint uintptr
	SetStartPoint uintptr
	GetEndPoint   uintptr
	SetEndPoint   uintptr
	Clone         uintptr
}

type IXpsOMLinearGradientBrush struct {
	IXpsOMGradientBrush
}

func (this *IXpsOMLinearGradientBrush) Vtbl() *IXpsOMLinearGradientBrushVtbl {
	return (*IXpsOMLinearGradientBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMLinearGradientBrush) GetStartPoint(startPoint *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStartPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startPoint)))
	return HRESULT(ret)
}

func (this *IXpsOMLinearGradientBrush) SetStartPoint(startPoint *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStartPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startPoint)))
	return HRESULT(ret)
}

func (this *IXpsOMLinearGradientBrush) GetEndPoint(endPoint *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEndPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endPoint)))
	return HRESULT(ret)
}

func (this *IXpsOMLinearGradientBrush) SetEndPoint(endPoint *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetEndPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(endPoint)))
	return HRESULT(ret)
}

func (this *IXpsOMLinearGradientBrush) Clone(linearGradientBrush **IXpsOMLinearGradientBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(linearGradientBrush)))
	return HRESULT(ret)
}

// 75F207E5-08BF-413C-96B1-B82B4064176B
var IID_IXpsOMRadialGradientBrush = syscall.GUID{0x75F207E5, 0x08BF, 0x413C,
	[8]byte{0x96, 0xB1, 0xB8, 0x2B, 0x40, 0x64, 0x17, 0x6B}}

type IXpsOMRadialGradientBrushInterface interface {
	IXpsOMGradientBrushInterface
	GetCenter(center *XPS_POINT) HRESULT
	SetCenter(center *XPS_POINT) HRESULT
	GetRadiiSizes(radiiSizes *XPS_SIZE) HRESULT
	SetRadiiSizes(radiiSizes *XPS_SIZE) HRESULT
	GetGradientOrigin(origin *XPS_POINT) HRESULT
	SetGradientOrigin(origin *XPS_POINT) HRESULT
	Clone(radialGradientBrush **IXpsOMRadialGradientBrush) HRESULT
}

type IXpsOMRadialGradientBrushVtbl struct {
	IXpsOMGradientBrushVtbl
	GetCenter         uintptr
	SetCenter         uintptr
	GetRadiiSizes     uintptr
	SetRadiiSizes     uintptr
	GetGradientOrigin uintptr
	SetGradientOrigin uintptr
	Clone             uintptr
}

type IXpsOMRadialGradientBrush struct {
	IXpsOMGradientBrush
}

func (this *IXpsOMRadialGradientBrush) Vtbl() *IXpsOMRadialGradientBrushVtbl {
	return (*IXpsOMRadialGradientBrushVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMRadialGradientBrush) GetCenter(center *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(center)))
	return HRESULT(ret)
}

func (this *IXpsOMRadialGradientBrush) SetCenter(center *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCenter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(center)))
	return HRESULT(ret)
}

func (this *IXpsOMRadialGradientBrush) GetRadiiSizes(radiiSizes *XPS_SIZE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRadiiSizes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(radiiSizes)))
	return HRESULT(ret)
}

func (this *IXpsOMRadialGradientBrush) SetRadiiSizes(radiiSizes *XPS_SIZE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetRadiiSizes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(radiiSizes)))
	return HRESULT(ret)
}

func (this *IXpsOMRadialGradientBrush) GetGradientOrigin(origin *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGradientOrigin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(origin)))
	return HRESULT(ret)
}

func (this *IXpsOMRadialGradientBrush) SetGradientOrigin(origin *XPS_POINT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetGradientOrigin, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(origin)))
	return HRESULT(ret)
}

func (this *IXpsOMRadialGradientBrush) Clone(radialGradientBrush **IXpsOMRadialGradientBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(radialGradientBrush)))
	return HRESULT(ret)
}

// DA2AC0A2-73A2-4975-AD14-74097C3FF3A5
var IID_IXpsOMResource = syscall.GUID{0xDA2AC0A2, 0x73A2, 0x4975,
	[8]byte{0xAD, 0x14, 0x74, 0x09, 0x7C, 0x3F, 0xF3, 0xA5}}

type IXpsOMResourceInterface interface {
	IXpsOMPartInterface
}

type IXpsOMResourceVtbl struct {
	IXpsOMPartVtbl
}

type IXpsOMResource struct {
	IXpsOMPart
}

func (this *IXpsOMResource) Vtbl() *IXpsOMResourceVtbl {
	return (*IXpsOMResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// F4CF7729-4864-4275-99B3-A8717163ECAF
var IID_IXpsOMPartResources = syscall.GUID{0xF4CF7729, 0x4864, 0x4275,
	[8]byte{0x99, 0xB3, 0xA8, 0x71, 0x71, 0x63, 0xEC, 0xAF}}

type IXpsOMPartResourcesInterface interface {
	IUnknownInterface
	GetFontResources(fontResources **IXpsOMFontResourceCollection) HRESULT
	GetImageResources(imageResources **IXpsOMImageResourceCollection) HRESULT
	GetColorProfileResources(colorProfileResources **IXpsOMColorProfileResourceCollection) HRESULT
	GetRemoteDictionaryResources(dictionaryResources **IXpsOMRemoteDictionaryResourceCollection) HRESULT
}

type IXpsOMPartResourcesVtbl struct {
	IUnknownVtbl
	GetFontResources             uintptr
	GetImageResources            uintptr
	GetColorProfileResources     uintptr
	GetRemoteDictionaryResources uintptr
}

type IXpsOMPartResources struct {
	IUnknown
}

func (this *IXpsOMPartResources) Vtbl() *IXpsOMPartResourcesVtbl {
	return (*IXpsOMPartResourcesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPartResources) GetFontResources(fontResources **IXpsOMFontResourceCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFontResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fontResources)))
	return HRESULT(ret)
}

func (this *IXpsOMPartResources) GetImageResources(imageResources **IXpsOMImageResourceCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetImageResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResources)))
	return HRESULT(ret)
}

func (this *IXpsOMPartResources) GetColorProfileResources(colorProfileResources **IXpsOMColorProfileResourceCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColorProfileResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(colorProfileResources)))
	return HRESULT(ret)
}

func (this *IXpsOMPartResources) GetRemoteDictionaryResources(dictionaryResources **IXpsOMRemoteDictionaryResourceCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRemoteDictionaryResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionaryResources)))
	return HRESULT(ret)
}

// 897C86B8-8EAF-4AE3-BDDE-56419FCF4236
var IID_IXpsOMDictionary = syscall.GUID{0x897C86B8, 0x8EAF, 0x4AE3,
	[8]byte{0xBD, 0xDE, 0x56, 0x41, 0x9F, 0xCF, 0x42, 0x36}}

type IXpsOMDictionaryInterface interface {
	IUnknownInterface
	GetOwner(owner **IUnknown) HRESULT
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, key *PWSTR, entry **IXpsOMShareable) HRESULT
	GetByKey(key PWSTR, beforeEntry *IXpsOMShareable, entry **IXpsOMShareable) HRESULT
	GetIndex(entry *IXpsOMShareable, index *uint32) HRESULT
	Append(key PWSTR, entry *IXpsOMShareable) HRESULT
	InsertAt(index uint32, key PWSTR, entry *IXpsOMShareable) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, key PWSTR, entry *IXpsOMShareable) HRESULT
	Clone(dictionary **IXpsOMDictionary) HRESULT
}

type IXpsOMDictionaryVtbl struct {
	IUnknownVtbl
	GetOwner uintptr
	GetCount uintptr
	GetAt    uintptr
	GetByKey uintptr
	GetIndex uintptr
	Append   uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Clone    uintptr
}

type IXpsOMDictionary struct {
	IUnknown
}

func (this *IXpsOMDictionary) Vtbl() *IXpsOMDictionaryVtbl {
	return (*IXpsOMDictionaryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMDictionary) GetOwner(owner **IUnknown) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) GetAt(index uint32, key *PWSTR, entry **IXpsOMShareable) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(entry)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) GetByKey(key PWSTR, beforeEntry *IXpsOMShareable, entry **IXpsOMShareable) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetByKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(beforeEntry)), uintptr(unsafe.Pointer(entry)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) GetIndex(entry *IXpsOMShareable, index *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(entry)), uintptr(unsafe.Pointer(index)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) Append(key PWSTR, entry *IXpsOMShareable) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(entry)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) InsertAt(index uint32, key PWSTR, entry *IXpsOMShareable) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(entry)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) SetAt(index uint32, key PWSTR, entry *IXpsOMShareable) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(entry)))
	return HRESULT(ret)
}

func (this *IXpsOMDictionary) Clone(dictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionary)))
	return HRESULT(ret)
}

// A8C45708-47D9-4AF4-8D20-33B48C9B8485
var IID_IXpsOMFontResource = syscall.GUID{0xA8C45708, 0x47D9, 0x4AF4,
	[8]byte{0x8D, 0x20, 0x33, 0xB4, 0x8C, 0x9B, 0x84, 0x85}}

type IXpsOMFontResourceInterface interface {
	IXpsOMResourceInterface
	GetStream(readerStream **IStream) HRESULT
	SetContent(sourceStream *IStream, embeddingOption XPS_FONT_EMBEDDING, partName unsafe.Pointer) HRESULT
	GetEmbeddingOption(embeddingOption *XPS_FONT_EMBEDDING) HRESULT
}

type IXpsOMFontResourceVtbl struct {
	IXpsOMResourceVtbl
	GetStream          uintptr
	SetContent         uintptr
	GetEmbeddingOption uintptr
}

type IXpsOMFontResource struct {
	IXpsOMResource
}

func (this *IXpsOMFontResource) Vtbl() *IXpsOMFontResourceVtbl {
	return (*IXpsOMFontResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMFontResource) GetStream(readerStream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(readerStream)))
	return HRESULT(ret)
}

func (this *IXpsOMFontResource) SetContent(sourceStream *IStream, embeddingOption XPS_FONT_EMBEDDING, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(embeddingOption), uintptr(partName))
	return HRESULT(ret)
}

func (this *IXpsOMFontResource) GetEmbeddingOption(embeddingOption *XPS_FONT_EMBEDDING) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEmbeddingOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(embeddingOption)))
	return HRESULT(ret)
}

// 70B4A6BB-88D4-4FA8-AAF9-6D9C596FDBAD
var IID_IXpsOMFontResourceCollection = syscall.GUID{0x70B4A6BB, 0x88D4, 0x4FA8,
	[8]byte{0xAA, 0xF9, 0x6D, 0x9C, 0x59, 0x6F, 0xDB, 0xAD}}

type IXpsOMFontResourceCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, value **IXpsOMFontResource) HRESULT
	SetAt(index uint32, value *IXpsOMFontResource) HRESULT
	InsertAt(index uint32, value *IXpsOMFontResource) HRESULT
	Append(value *IXpsOMFontResource) HRESULT
	RemoveAt(index uint32) HRESULT
	GetByPartName(partName unsafe.Pointer, part **IXpsOMFontResource) HRESULT
}

type IXpsOMFontResourceCollectionVtbl struct {
	IUnknownVtbl
	GetCount      uintptr
	GetAt         uintptr
	SetAt         uintptr
	InsertAt      uintptr
	Append        uintptr
	RemoveAt      uintptr
	GetByPartName uintptr
}

type IXpsOMFontResourceCollection struct {
	IUnknown
}

func (this *IXpsOMFontResourceCollection) Vtbl() *IXpsOMFontResourceCollectionVtbl {
	return (*IXpsOMFontResourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMFontResourceCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMFontResourceCollection) GetAt(index uint32, value **IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func (this *IXpsOMFontResourceCollection) SetAt(index uint32, value *IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func (this *IXpsOMFontResourceCollection) InsertAt(index uint32, value *IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func (this *IXpsOMFontResourceCollection) Append(value *IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func (this *IXpsOMFontResourceCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMFontResourceCollection) GetByPartName(partName unsafe.Pointer, part **IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetByPartName, uintptr(unsafe.Pointer(this)), uintptr(partName), uintptr(unsafe.Pointer(part)))
	return HRESULT(ret)
}

// 3DB8417D-AE50-485E-9A44-D7758F78A23F
var IID_IXpsOMImageResource = syscall.GUID{0x3DB8417D, 0xAE50, 0x485E,
	[8]byte{0x9A, 0x44, 0xD7, 0x75, 0x8F, 0x78, 0xA2, 0x3F}}

type IXpsOMImageResourceInterface interface {
	IXpsOMResourceInterface
	GetStream(readerStream **IStream) HRESULT
	SetContent(sourceStream *IStream, imageType XPS_IMAGE_TYPE, partName unsafe.Pointer) HRESULT
	GetImageType(imageType *XPS_IMAGE_TYPE) HRESULT
}

type IXpsOMImageResourceVtbl struct {
	IXpsOMResourceVtbl
	GetStream    uintptr
	SetContent   uintptr
	GetImageType uintptr
}

type IXpsOMImageResource struct {
	IXpsOMResource
}

func (this *IXpsOMImageResource) Vtbl() *IXpsOMImageResourceVtbl {
	return (*IXpsOMImageResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMImageResource) GetStream(readerStream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(readerStream)))
	return HRESULT(ret)
}

func (this *IXpsOMImageResource) SetContent(sourceStream *IStream, imageType XPS_IMAGE_TYPE, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(imageType), uintptr(partName))
	return HRESULT(ret)
}

func (this *IXpsOMImageResource) GetImageType(imageType *XPS_IMAGE_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetImageType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageType)))
	return HRESULT(ret)
}

// 7A4A1A71-9CDE-4B71-B33F-62DE843EABFE
var IID_IXpsOMImageResourceCollection = syscall.GUID{0x7A4A1A71, 0x9CDE, 0x4B71,
	[8]byte{0xB3, 0x3F, 0x62, 0xDE, 0x84, 0x3E, 0xAB, 0xFE}}

type IXpsOMImageResourceCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, object **IXpsOMImageResource) HRESULT
	InsertAt(index uint32, object *IXpsOMImageResource) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, object *IXpsOMImageResource) HRESULT
	Append(object *IXpsOMImageResource) HRESULT
	GetByPartName(partName unsafe.Pointer, part **IXpsOMImageResource) HRESULT
}

type IXpsOMImageResourceCollectionVtbl struct {
	IUnknownVtbl
	GetCount      uintptr
	GetAt         uintptr
	InsertAt      uintptr
	RemoveAt      uintptr
	SetAt         uintptr
	Append        uintptr
	GetByPartName uintptr
}

type IXpsOMImageResourceCollection struct {
	IUnknown
}

func (this *IXpsOMImageResourceCollection) Vtbl() *IXpsOMImageResourceCollectionVtbl {
	return (*IXpsOMImageResourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMImageResourceCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMImageResourceCollection) GetAt(index uint32, object **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMImageResourceCollection) InsertAt(index uint32, object *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMImageResourceCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMImageResourceCollection) SetAt(index uint32, object *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMImageResourceCollection) Append(object *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMImageResourceCollection) GetByPartName(partName unsafe.Pointer, part **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetByPartName, uintptr(unsafe.Pointer(this)), uintptr(partName), uintptr(unsafe.Pointer(part)))
	return HRESULT(ret)
}

// 67BD7D69-1EEF-4BB1-B5E7-6F4F87BE8ABE
var IID_IXpsOMColorProfileResource = syscall.GUID{0x67BD7D69, 0x1EEF, 0x4BB1,
	[8]byte{0xB5, 0xE7, 0x6F, 0x4F, 0x87, 0xBE, 0x8A, 0xBE}}

type IXpsOMColorProfileResourceInterface interface {
	IXpsOMResourceInterface
	GetStream(stream **IStream) HRESULT
	SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT
}

type IXpsOMColorProfileResourceVtbl struct {
	IXpsOMResourceVtbl
	GetStream  uintptr
	SetContent uintptr
}

type IXpsOMColorProfileResource struct {
	IXpsOMResource
}

func (this *IXpsOMColorProfileResource) Vtbl() *IXpsOMColorProfileResourceVtbl {
	return (*IXpsOMColorProfileResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMColorProfileResource) GetStream(stream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResource) SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(partName))
	return HRESULT(ret)
}

// 12759630-5FBA-4283-8F7D-CCA849809EDB
var IID_IXpsOMColorProfileResourceCollection = syscall.GUID{0x12759630, 0x5FBA, 0x4283,
	[8]byte{0x8F, 0x7D, 0xCC, 0xA8, 0x49, 0x80, 0x9E, 0xDB}}

type IXpsOMColorProfileResourceCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, object **IXpsOMColorProfileResource) HRESULT
	InsertAt(index uint32, object *IXpsOMColorProfileResource) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, object *IXpsOMColorProfileResource) HRESULT
	Append(object *IXpsOMColorProfileResource) HRESULT
	GetByPartName(partName unsafe.Pointer, part **IXpsOMColorProfileResource) HRESULT
}

type IXpsOMColorProfileResourceCollectionVtbl struct {
	IUnknownVtbl
	GetCount      uintptr
	GetAt         uintptr
	InsertAt      uintptr
	RemoveAt      uintptr
	SetAt         uintptr
	Append        uintptr
	GetByPartName uintptr
}

type IXpsOMColorProfileResourceCollection struct {
	IUnknown
}

func (this *IXpsOMColorProfileResourceCollection) Vtbl() *IXpsOMColorProfileResourceCollectionVtbl {
	return (*IXpsOMColorProfileResourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMColorProfileResourceCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResourceCollection) GetAt(index uint32, object **IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResourceCollection) InsertAt(index uint32, object *IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResourceCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResourceCollection) SetAt(index uint32, object *IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResourceCollection) Append(object *IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMColorProfileResourceCollection) GetByPartName(partName unsafe.Pointer, part **IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetByPartName, uintptr(unsafe.Pointer(this)), uintptr(partName), uintptr(unsafe.Pointer(part)))
	return HRESULT(ret)
}

// E7FF32D2-34AA-499B-BBE9-9CD4EE6C59F7
var IID_IXpsOMPrintTicketResource = syscall.GUID{0xE7FF32D2, 0x34AA, 0x499B,
	[8]byte{0xBB, 0xE9, 0x9C, 0xD4, 0xEE, 0x6C, 0x59, 0xF7}}

type IXpsOMPrintTicketResourceInterface interface {
	IXpsOMResourceInterface
	GetStream(stream **IStream) HRESULT
	SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT
}

type IXpsOMPrintTicketResourceVtbl struct {
	IXpsOMResourceVtbl
	GetStream  uintptr
	SetContent uintptr
}

type IXpsOMPrintTicketResource struct {
	IXpsOMResource
}

func (this *IXpsOMPrintTicketResource) Vtbl() *IXpsOMPrintTicketResourceVtbl {
	return (*IXpsOMPrintTicketResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPrintTicketResource) GetStream(stream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

func (this *IXpsOMPrintTicketResource) SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(partName))
	return HRESULT(ret)
}

// C9BD7CD4-E16A-4BF8-8C84-C950AF7A3061
var IID_IXpsOMRemoteDictionaryResource = syscall.GUID{0xC9BD7CD4, 0xE16A, 0x4BF8,
	[8]byte{0x8C, 0x84, 0xC9, 0x50, 0xAF, 0x7A, 0x30, 0x61}}

type IXpsOMRemoteDictionaryResourceInterface interface {
	IXpsOMResourceInterface
	GetDictionary(dictionary **IXpsOMDictionary) HRESULT
	SetDictionary(dictionary *IXpsOMDictionary) HRESULT
}

type IXpsOMRemoteDictionaryResourceVtbl struct {
	IXpsOMResourceVtbl
	GetDictionary uintptr
	SetDictionary uintptr
}

type IXpsOMRemoteDictionaryResource struct {
	IXpsOMResource
}

func (this *IXpsOMRemoteDictionaryResource) Vtbl() *IXpsOMRemoteDictionaryResourceVtbl {
	return (*IXpsOMRemoteDictionaryResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMRemoteDictionaryResource) GetDictionary(dictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResource) SetDictionary(dictionary *IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDictionary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionary)))
	return HRESULT(ret)
}

// 5C38DB61-7FEC-464A-87BD-41E3BEF018BE
var IID_IXpsOMRemoteDictionaryResourceCollection = syscall.GUID{0x5C38DB61, 0x7FEC, 0x464A,
	[8]byte{0x87, 0xBD, 0x41, 0xE3, 0xBE, 0xF0, 0x18, 0xBE}}

type IXpsOMRemoteDictionaryResourceCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, object **IXpsOMRemoteDictionaryResource) HRESULT
	InsertAt(index uint32, object *IXpsOMRemoteDictionaryResource) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, object *IXpsOMRemoteDictionaryResource) HRESULT
	Append(object *IXpsOMRemoteDictionaryResource) HRESULT
	GetByPartName(partName unsafe.Pointer, remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT
}

type IXpsOMRemoteDictionaryResourceCollectionVtbl struct {
	IUnknownVtbl
	GetCount      uintptr
	GetAt         uintptr
	InsertAt      uintptr
	RemoveAt      uintptr
	SetAt         uintptr
	Append        uintptr
	GetByPartName uintptr
}

type IXpsOMRemoteDictionaryResourceCollection struct {
	IUnknown
}

func (this *IXpsOMRemoteDictionaryResourceCollection) Vtbl() *IXpsOMRemoteDictionaryResourceCollectionVtbl {
	return (*IXpsOMRemoteDictionaryResourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMRemoteDictionaryResourceCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResourceCollection) GetAt(index uint32, object **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResourceCollection) InsertAt(index uint32, object *IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResourceCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResourceCollection) SetAt(index uint32, object *IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResourceCollection) Append(object *IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResourceCollection) GetByPartName(partName unsafe.Pointer, remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetByPartName, uintptr(unsafe.Pointer(this)), uintptr(partName), uintptr(unsafe.Pointer(remoteDictionaryResource)))
	return HRESULT(ret)
}

// AB8F5D8E-351B-4D33-AAED-FA56F0022931
var IID_IXpsOMSignatureBlockResourceCollection = syscall.GUID{0xAB8F5D8E, 0x351B, 0x4D33,
	[8]byte{0xAA, 0xED, 0xFA, 0x56, 0xF0, 0x02, 0x29, 0x31}}

type IXpsOMSignatureBlockResourceCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, signatureBlockResource **IXpsOMSignatureBlockResource) HRESULT
	InsertAt(index uint32, signatureBlockResource *IXpsOMSignatureBlockResource) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, signatureBlockResource *IXpsOMSignatureBlockResource) HRESULT
	Append(signatureBlockResource *IXpsOMSignatureBlockResource) HRESULT
	GetByPartName(partName unsafe.Pointer, signatureBlockResource **IXpsOMSignatureBlockResource) HRESULT
}

type IXpsOMSignatureBlockResourceCollectionVtbl struct {
	IUnknownVtbl
	GetCount      uintptr
	GetAt         uintptr
	InsertAt      uintptr
	RemoveAt      uintptr
	SetAt         uintptr
	Append        uintptr
	GetByPartName uintptr
}

type IXpsOMSignatureBlockResourceCollection struct {
	IUnknown
}

func (this *IXpsOMSignatureBlockResourceCollection) Vtbl() *IXpsOMSignatureBlockResourceCollectionVtbl {
	return (*IXpsOMSignatureBlockResourceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMSignatureBlockResourceCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResourceCollection) GetAt(index uint32, signatureBlockResource **IXpsOMSignatureBlockResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(signatureBlockResource)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResourceCollection) InsertAt(index uint32, signatureBlockResource *IXpsOMSignatureBlockResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(signatureBlockResource)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResourceCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResourceCollection) SetAt(index uint32, signatureBlockResource *IXpsOMSignatureBlockResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(signatureBlockResource)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResourceCollection) Append(signatureBlockResource *IXpsOMSignatureBlockResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureBlockResource)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResourceCollection) GetByPartName(partName unsafe.Pointer, signatureBlockResource **IXpsOMSignatureBlockResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetByPartName, uintptr(unsafe.Pointer(this)), uintptr(partName), uintptr(unsafe.Pointer(signatureBlockResource)))
	return HRESULT(ret)
}

// 85FEBC8A-6B63-48A9-AF07-7064E4ECFF30
var IID_IXpsOMDocumentStructureResource = syscall.GUID{0x85FEBC8A, 0x6B63, 0x48A9,
	[8]byte{0xAF, 0x07, 0x70, 0x64, 0xE4, 0xEC, 0xFF, 0x30}}

type IXpsOMDocumentStructureResourceInterface interface {
	IXpsOMResourceInterface
	GetOwner(owner **IXpsOMDocument) HRESULT
	GetStream(stream **IStream) HRESULT
	SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT
}

type IXpsOMDocumentStructureResourceVtbl struct {
	IXpsOMResourceVtbl
	GetOwner   uintptr
	GetStream  uintptr
	SetContent uintptr
}

type IXpsOMDocumentStructureResource struct {
	IXpsOMResource
}

func (this *IXpsOMDocumentStructureResource) Vtbl() *IXpsOMDocumentStructureResourceVtbl {
	return (*IXpsOMDocumentStructureResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMDocumentStructureResource) GetOwner(owner **IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentStructureResource) GetStream(stream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentStructureResource) SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(partName))
	return HRESULT(ret)
}

// C2B3CA09-0473-4282-87AE-1780863223F0
var IID_IXpsOMStoryFragmentsResource = syscall.GUID{0xC2B3CA09, 0x0473, 0x4282,
	[8]byte{0x87, 0xAE, 0x17, 0x80, 0x86, 0x32, 0x23, 0xF0}}

type IXpsOMStoryFragmentsResourceInterface interface {
	IXpsOMResourceInterface
	GetOwner(owner **IXpsOMPageReference) HRESULT
	GetStream(stream **IStream) HRESULT
	SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT
}

type IXpsOMStoryFragmentsResourceVtbl struct {
	IXpsOMResourceVtbl
	GetOwner   uintptr
	GetStream  uintptr
	SetContent uintptr
}

type IXpsOMStoryFragmentsResource struct {
	IXpsOMResource
}

func (this *IXpsOMStoryFragmentsResource) Vtbl() *IXpsOMStoryFragmentsResourceVtbl {
	return (*IXpsOMStoryFragmentsResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMStoryFragmentsResource) GetOwner(owner **IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMStoryFragmentsResource) GetStream(stream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

func (this *IXpsOMStoryFragmentsResource) SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(partName))
	return HRESULT(ret)
}

// 4776AD35-2E04-4357-8743-EBF6C171A905
var IID_IXpsOMSignatureBlockResource = syscall.GUID{0x4776AD35, 0x2E04, 0x4357,
	[8]byte{0x87, 0x43, 0xEB, 0xF6, 0xC1, 0x71, 0xA9, 0x05}}

type IXpsOMSignatureBlockResourceInterface interface {
	IXpsOMResourceInterface
	GetOwner(owner **IXpsOMDocument) HRESULT
	GetStream(stream **IStream) HRESULT
	SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT
}

type IXpsOMSignatureBlockResourceVtbl struct {
	IXpsOMResourceVtbl
	GetOwner   uintptr
	GetStream  uintptr
	SetContent uintptr
}

type IXpsOMSignatureBlockResource struct {
	IXpsOMResource
}

func (this *IXpsOMSignatureBlockResource) Vtbl() *IXpsOMSignatureBlockResourceVtbl {
	return (*IXpsOMSignatureBlockResourceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMSignatureBlockResource) GetOwner(owner **IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(owner)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResource) GetStream(stream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

func (this *IXpsOMSignatureBlockResource) SetContent(sourceStream *IStream, partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sourceStream)), uintptr(partName))
	return HRESULT(ret)
}

// 94D8ABDE-AB91-46A8-82B7-F5B05EF01A96
var IID_IXpsOMVisualCollection = syscall.GUID{0x94D8ABDE, 0xAB91, 0x46A8,
	[8]byte{0x82, 0xB7, 0xF5, 0xB0, 0x5E, 0xF0, 0x1A, 0x96}}

type IXpsOMVisualCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, object **IXpsOMVisual) HRESULT
	InsertAt(index uint32, object *IXpsOMVisual) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, object *IXpsOMVisual) HRESULT
	Append(object *IXpsOMVisual) HRESULT
}

type IXpsOMVisualCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMVisualCollection struct {
	IUnknown
}

func (this *IXpsOMVisualCollection) Vtbl() *IXpsOMVisualCollectionVtbl {
	return (*IXpsOMVisualCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMVisualCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualCollection) GetAt(index uint32, object **IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualCollection) InsertAt(index uint32, object *IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMVisualCollection) SetAt(index uint32, object *IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

func (this *IXpsOMVisualCollection) Append(object *IXpsOMVisual) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(object)))
	return HRESULT(ret)
}

// 221D1452-331E-47C6-87E9-6CCEFB9B5BA3
var IID_IXpsOMCanvas = syscall.GUID{0x221D1452, 0x331E, 0x47C6,
	[8]byte{0x87, 0xE9, 0x6C, 0xCE, 0xFB, 0x9B, 0x5B, 0xA3}}

type IXpsOMCanvasInterface interface {
	IXpsOMVisualInterface
	GetVisuals(visuals **IXpsOMVisualCollection) HRESULT
	GetUseAliasedEdgeMode(useAliasedEdgeMode *BOOL) HRESULT
	SetUseAliasedEdgeMode(useAliasedEdgeMode BOOL) HRESULT
	GetAccessibilityShortDescription(shortDescription *PWSTR) HRESULT
	SetAccessibilityShortDescription(shortDescription PWSTR) HRESULT
	GetAccessibilityLongDescription(longDescription *PWSTR) HRESULT
	SetAccessibilityLongDescription(longDescription PWSTR) HRESULT
	GetDictionary(resourceDictionary **IXpsOMDictionary) HRESULT
	GetDictionaryLocal(resourceDictionary **IXpsOMDictionary) HRESULT
	SetDictionaryLocal(resourceDictionary *IXpsOMDictionary) HRESULT
	GetDictionaryResource(remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT
	SetDictionaryResource(remoteDictionaryResource *IXpsOMRemoteDictionaryResource) HRESULT
	Clone(canvas **IXpsOMCanvas) HRESULT
}

type IXpsOMCanvasVtbl struct {
	IXpsOMVisualVtbl
	GetVisuals                       uintptr
	GetUseAliasedEdgeMode            uintptr
	SetUseAliasedEdgeMode            uintptr
	GetAccessibilityShortDescription uintptr
	SetAccessibilityShortDescription uintptr
	GetAccessibilityLongDescription  uintptr
	SetAccessibilityLongDescription  uintptr
	GetDictionary                    uintptr
	GetDictionaryLocal               uintptr
	SetDictionaryLocal               uintptr
	GetDictionaryResource            uintptr
	SetDictionaryResource            uintptr
	Clone                            uintptr
}

type IXpsOMCanvas struct {
	IXpsOMVisual
}

func (this *IXpsOMCanvas) Vtbl() *IXpsOMCanvasVtbl {
	return (*IXpsOMCanvasVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMCanvas) GetVisuals(visuals **IXpsOMVisualCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisuals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visuals)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) GetUseAliasedEdgeMode(useAliasedEdgeMode *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetUseAliasedEdgeMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(useAliasedEdgeMode)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) SetUseAliasedEdgeMode(useAliasedEdgeMode BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetUseAliasedEdgeMode, uintptr(unsafe.Pointer(this)), uintptr(useAliasedEdgeMode))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) GetAccessibilityShortDescription(shortDescription *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAccessibilityShortDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shortDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) SetAccessibilityShortDescription(shortDescription PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAccessibilityShortDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shortDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) GetAccessibilityLongDescription(longDescription *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAccessibilityLongDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(longDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) SetAccessibilityLongDescription(longDescription PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAccessibilityLongDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(longDescription)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) GetDictionary(resourceDictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourceDictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) GetDictionaryLocal(resourceDictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionaryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourceDictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) SetDictionaryLocal(resourceDictionary *IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDictionaryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourceDictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) GetDictionaryResource(remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionaryResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteDictionaryResource)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) SetDictionaryResource(remoteDictionaryResource *IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDictionaryResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteDictionaryResource)))
	return HRESULT(ret)
}

func (this *IXpsOMCanvas) Clone(canvas **IXpsOMCanvas) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(canvas)))
	return HRESULT(ret)
}

// D3E18888-F120-4FEE-8C68-35296EAE91D4
var IID_IXpsOMPage = syscall.GUID{0xD3E18888, 0xF120, 0x4FEE,
	[8]byte{0x8C, 0x68, 0x35, 0x29, 0x6E, 0xAE, 0x91, 0xD4}}

type IXpsOMPageInterface interface {
	IXpsOMPartInterface
	GetOwner(pageReference **IXpsOMPageReference) HRESULT
	GetVisuals(visuals **IXpsOMVisualCollection) HRESULT
	GetPageDimensions(pageDimensions *XPS_SIZE) HRESULT
	SetPageDimensions(pageDimensions *XPS_SIZE) HRESULT
	GetContentBox(contentBox *XPS_RECT) HRESULT
	SetContentBox(contentBox *XPS_RECT) HRESULT
	GetBleedBox(bleedBox *XPS_RECT) HRESULT
	SetBleedBox(bleedBox *XPS_RECT) HRESULT
	GetLanguage(language *PWSTR) HRESULT
	SetLanguage(language PWSTR) HRESULT
	GetName(name *PWSTR) HRESULT
	SetName(name PWSTR) HRESULT
	GetIsHyperlinkTarget(isHyperlinkTarget *BOOL) HRESULT
	SetIsHyperlinkTarget(isHyperlinkTarget BOOL) HRESULT
	GetDictionary(resourceDictionary **IXpsOMDictionary) HRESULT
	GetDictionaryLocal(resourceDictionary **IXpsOMDictionary) HRESULT
	SetDictionaryLocal(resourceDictionary *IXpsOMDictionary) HRESULT
	GetDictionaryResource(remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT
	SetDictionaryResource(remoteDictionaryResource *IXpsOMRemoteDictionaryResource) HRESULT
	Write(stream *ISequentialStream, optimizeMarkupSize BOOL) HRESULT
	GenerateUnusedLookupKey(type_ XPS_OBJECT_TYPE, key *PWSTR) HRESULT
	Clone(page **IXpsOMPage) HRESULT
}

type IXpsOMPageVtbl struct {
	IXpsOMPartVtbl
	GetOwner                uintptr
	GetVisuals              uintptr
	GetPageDimensions       uintptr
	SetPageDimensions       uintptr
	GetContentBox           uintptr
	SetContentBox           uintptr
	GetBleedBox             uintptr
	SetBleedBox             uintptr
	GetLanguage             uintptr
	SetLanguage             uintptr
	GetName                 uintptr
	SetName                 uintptr
	GetIsHyperlinkTarget    uintptr
	SetIsHyperlinkTarget    uintptr
	GetDictionary           uintptr
	GetDictionaryLocal      uintptr
	SetDictionaryLocal      uintptr
	GetDictionaryResource   uintptr
	SetDictionaryResource   uintptr
	Write                   uintptr
	GenerateUnusedLookupKey uintptr
	Clone                   uintptr
}

type IXpsOMPage struct {
	IXpsOMPart
}

func (this *IXpsOMPage) Vtbl() *IXpsOMPageVtbl {
	return (*IXpsOMPageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPage) GetOwner(pageReference **IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetVisuals(visuals **IXpsOMVisualCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisuals, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(visuals)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetPageDimensions(pageDimensions *XPS_SIZE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPageDimensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageDimensions)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetPageDimensions(pageDimensions *XPS_SIZE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPageDimensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageDimensions)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetContentBox(contentBox *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetContentBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentBox)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetContentBox(contentBox *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContentBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentBox)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetBleedBox(bleedBox *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetBleedBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bleedBox)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetBleedBox(bleedBox *XPS_RECT) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetBleedBox, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bleedBox)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetLanguage(language *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetLanguage(language PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetName(name *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetName(name PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetIsHyperlinkTarget(isHyperlinkTarget *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIsHyperlinkTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isHyperlinkTarget)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetIsHyperlinkTarget(isHyperlinkTarget BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIsHyperlinkTarget, uintptr(unsafe.Pointer(this)), uintptr(isHyperlinkTarget))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetDictionary(resourceDictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourceDictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetDictionaryLocal(resourceDictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionaryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourceDictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetDictionaryLocal(resourceDictionary *IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDictionaryLocal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resourceDictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GetDictionaryResource(remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDictionaryResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteDictionaryResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) SetDictionaryResource(remoteDictionaryResource *IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDictionaryResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(remoteDictionaryResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) Write(stream *ISequentialStream, optimizeMarkupSize BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(optimizeMarkupSize))
	return HRESULT(ret)
}

func (this *IXpsOMPage) GenerateUnusedLookupKey(type_ XPS_OBJECT_TYPE, key *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GenerateUnusedLookupKey, uintptr(unsafe.Pointer(this)), uintptr(type_), uintptr(unsafe.Pointer(key)))
	return HRESULT(ret)
}

func (this *IXpsOMPage) Clone(page **IXpsOMPage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

// ED360180-6F92-4998-890D-2F208531A0A0
var IID_IXpsOMPageReference = syscall.GUID{0xED360180, 0x6F92, 0x4998,
	[8]byte{0x89, 0x0D, 0x2F, 0x20, 0x85, 0x31, 0xA0, 0xA0}}

type IXpsOMPageReferenceInterface interface {
	IUnknownInterface
	GetOwner(document **IXpsOMDocument) HRESULT
	GetPage(page **IXpsOMPage) HRESULT
	SetPage(page *IXpsOMPage) HRESULT
	DiscardPage() HRESULT
	IsPageLoaded(isPageLoaded *BOOL) HRESULT
	GetAdvisoryPageDimensions(pageDimensions *XPS_SIZE) HRESULT
	SetAdvisoryPageDimensions(pageDimensions *XPS_SIZE) HRESULT
	GetStoryFragmentsResource(storyFragmentsResource **IXpsOMStoryFragmentsResource) HRESULT
	SetStoryFragmentsResource(storyFragmentsResource *IXpsOMStoryFragmentsResource) HRESULT
	GetPrintTicketResource(printTicketResource **IXpsOMPrintTicketResource) HRESULT
	SetPrintTicketResource(printTicketResource *IXpsOMPrintTicketResource) HRESULT
	GetThumbnailResource(imageResource **IXpsOMImageResource) HRESULT
	SetThumbnailResource(imageResource *IXpsOMImageResource) HRESULT
	CollectLinkTargets(linkTargets **IXpsOMNameCollection) HRESULT
	CollectPartResources(partResources **IXpsOMPartResources) HRESULT
	HasRestrictedFonts(restrictedFonts *BOOL) HRESULT
	Clone(pageReference **IXpsOMPageReference) HRESULT
}

type IXpsOMPageReferenceVtbl struct {
	IUnknownVtbl
	GetOwner                  uintptr
	GetPage                   uintptr
	SetPage                   uintptr
	DiscardPage               uintptr
	IsPageLoaded              uintptr
	GetAdvisoryPageDimensions uintptr
	SetAdvisoryPageDimensions uintptr
	GetStoryFragmentsResource uintptr
	SetStoryFragmentsResource uintptr
	GetPrintTicketResource    uintptr
	SetPrintTicketResource    uintptr
	GetThumbnailResource      uintptr
	SetThumbnailResource      uintptr
	CollectLinkTargets        uintptr
	CollectPartResources      uintptr
	HasRestrictedFonts        uintptr
	Clone                     uintptr
}

type IXpsOMPageReference struct {
	IUnknown
}

func (this *IXpsOMPageReference) Vtbl() *IXpsOMPageReferenceVtbl {
	return (*IXpsOMPageReferenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPageReference) GetOwner(document **IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) GetPage(page **IXpsOMPage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) SetPage(page *IXpsOMPage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) DiscardPage() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().DiscardPage, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) IsPageLoaded(isPageLoaded *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsPageLoaded, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isPageLoaded)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) GetAdvisoryPageDimensions(pageDimensions *XPS_SIZE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAdvisoryPageDimensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageDimensions)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) SetAdvisoryPageDimensions(pageDimensions *XPS_SIZE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAdvisoryPageDimensions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageDimensions)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) GetStoryFragmentsResource(storyFragmentsResource **IXpsOMStoryFragmentsResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetStoryFragmentsResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(storyFragmentsResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) SetStoryFragmentsResource(storyFragmentsResource *IXpsOMStoryFragmentsResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetStoryFragmentsResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(storyFragmentsResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) GetPrintTicketResource(printTicketResource **IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) SetPrintTicketResource(printTicketResource *IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) GetThumbnailResource(imageResource **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) SetThumbnailResource(imageResource *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetThumbnailResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) CollectLinkTargets(linkTargets **IXpsOMNameCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CollectLinkTargets, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(linkTargets)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) CollectPartResources(partResources **IXpsOMPartResources) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CollectPartResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(partResources)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) HasRestrictedFonts(restrictedFonts *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().HasRestrictedFonts, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(restrictedFonts)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReference) Clone(pageReference **IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

// CA16BA4D-E7B9-45C5-958B-F98022473745
var IID_IXpsOMPageReferenceCollection = syscall.GUID{0xCA16BA4D, 0xE7B9, 0x45C5,
	[8]byte{0x95, 0x8B, 0xF9, 0x80, 0x22, 0x47, 0x37, 0x45}}

type IXpsOMPageReferenceCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, pageReference **IXpsOMPageReference) HRESULT
	InsertAt(index uint32, pageReference *IXpsOMPageReference) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, pageReference *IXpsOMPageReference) HRESULT
	Append(pageReference *IXpsOMPageReference) HRESULT
}

type IXpsOMPageReferenceCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMPageReferenceCollection struct {
	IUnknown
}

func (this *IXpsOMPageReferenceCollection) Vtbl() *IXpsOMPageReferenceCollectionVtbl {
	return (*IXpsOMPageReferenceCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPageReferenceCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReferenceCollection) GetAt(index uint32, pageReference **IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReferenceCollection) InsertAt(index uint32, pageReference *IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReferenceCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMPageReferenceCollection) SetAt(index uint32, pageReference *IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

func (this *IXpsOMPageReferenceCollection) Append(pageReference *IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

// 2C2C94CB-AC5F-4254-8EE9-23948309D9F0
var IID_IXpsOMDocument = syscall.GUID{0x2C2C94CB, 0xAC5F, 0x4254,
	[8]byte{0x8E, 0xE9, 0x23, 0x94, 0x83, 0x09, 0xD9, 0xF0}}

type IXpsOMDocumentInterface interface {
	IXpsOMPartInterface
	GetOwner(documentSequence **IXpsOMDocumentSequence) HRESULT
	GetPageReferences(pageReferences **IXpsOMPageReferenceCollection) HRESULT
	GetPrintTicketResource(printTicketResource **IXpsOMPrintTicketResource) HRESULT
	SetPrintTicketResource(printTicketResource *IXpsOMPrintTicketResource) HRESULT
	GetDocumentStructureResource(documentStructureResource **IXpsOMDocumentStructureResource) HRESULT
	SetDocumentStructureResource(documentStructureResource *IXpsOMDocumentStructureResource) HRESULT
	GetSignatureBlockResources(signatureBlockResources **IXpsOMSignatureBlockResourceCollection) HRESULT
	Clone(document **IXpsOMDocument) HRESULT
}

type IXpsOMDocumentVtbl struct {
	IXpsOMPartVtbl
	GetOwner                     uintptr
	GetPageReferences            uintptr
	GetPrintTicketResource       uintptr
	SetPrintTicketResource       uintptr
	GetDocumentStructureResource uintptr
	SetDocumentStructureResource uintptr
	GetSignatureBlockResources   uintptr
	Clone                        uintptr
}

type IXpsOMDocument struct {
	IXpsOMPart
}

func (this *IXpsOMDocument) Vtbl() *IXpsOMDocumentVtbl {
	return (*IXpsOMDocumentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMDocument) GetOwner(documentSequence **IXpsOMDocumentSequence) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentSequence)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) GetPageReferences(pageReferences **IXpsOMPageReferenceCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPageReferences, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageReferences)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) GetPrintTicketResource(printTicketResource **IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) SetPrintTicketResource(printTicketResource *IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) GetDocumentStructureResource(documentStructureResource **IXpsOMDocumentStructureResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentStructureResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentStructureResource)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) SetDocumentStructureResource(documentStructureResource *IXpsOMDocumentStructureResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDocumentStructureResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentStructureResource)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) GetSignatureBlockResources(signatureBlockResources **IXpsOMSignatureBlockResourceCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureBlockResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureBlockResources)))
	return HRESULT(ret)
}

func (this *IXpsOMDocument) Clone(document **IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

// D1C87F0D-E947-4754-8A25-971478F7E83E
var IID_IXpsOMDocumentCollection = syscall.GUID{0xD1C87F0D, 0xE947, 0x4754,
	[8]byte{0x8A, 0x25, 0x97, 0x14, 0x78, 0xF7, 0xE8, 0x3E}}

type IXpsOMDocumentCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, document **IXpsOMDocument) HRESULT
	InsertAt(index uint32, document *IXpsOMDocument) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, document *IXpsOMDocument) HRESULT
	Append(document *IXpsOMDocument) HRESULT
}

type IXpsOMDocumentCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMDocumentCollection struct {
	IUnknown
}

func (this *IXpsOMDocumentCollection) Vtbl() *IXpsOMDocumentCollectionVtbl {
	return (*IXpsOMDocumentCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMDocumentCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentCollection) GetAt(index uint32, document **IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentCollection) InsertAt(index uint32, document *IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentCollection) SetAt(index uint32, document *IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentCollection) Append(document *IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

// 56492EB4-D8D5-425E-8256-4C2B64AD0264
var IID_IXpsOMDocumentSequence = syscall.GUID{0x56492EB4, 0xD8D5, 0x425E,
	[8]byte{0x82, 0x56, 0x4C, 0x2B, 0x64, 0xAD, 0x02, 0x64}}

type IXpsOMDocumentSequenceInterface interface {
	IXpsOMPartInterface
	GetOwner(package_ **IXpsOMPackage) HRESULT
	GetDocuments(documents **IXpsOMDocumentCollection) HRESULT
	GetPrintTicketResource(printTicketResource **IXpsOMPrintTicketResource) HRESULT
	SetPrintTicketResource(printTicketResource *IXpsOMPrintTicketResource) HRESULT
}

type IXpsOMDocumentSequenceVtbl struct {
	IXpsOMPartVtbl
	GetOwner               uintptr
	GetDocuments           uintptr
	GetPrintTicketResource uintptr
	SetPrintTicketResource uintptr
}

type IXpsOMDocumentSequence struct {
	IXpsOMPart
}

func (this *IXpsOMDocumentSequence) Vtbl() *IXpsOMDocumentSequenceVtbl {
	return (*IXpsOMDocumentSequenceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMDocumentSequence) GetOwner(package_ **IXpsOMPackage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentSequence) GetDocuments(documents **IXpsOMDocumentCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocuments, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documents)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentSequence) GetPrintTicketResource(printTicketResource **IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

func (this *IXpsOMDocumentSequence) SetPrintTicketResource(printTicketResource *IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

// 3340FE8F-4027-4AA1-8F5F-D35AE45FE597
var IID_IXpsOMCoreProperties = syscall.GUID{0x3340FE8F, 0x4027, 0x4AA1,
	[8]byte{0x8F, 0x5F, 0xD3, 0x5A, 0xE4, 0x5F, 0xE5, 0x97}}

type IXpsOMCorePropertiesInterface interface {
	IXpsOMPartInterface
	GetOwner(package_ **IXpsOMPackage) HRESULT
	GetCategory(category *PWSTR) HRESULT
	SetCategory(category PWSTR) HRESULT
	GetContentStatus(contentStatus *PWSTR) HRESULT
	SetContentStatus(contentStatus PWSTR) HRESULT
	GetContentType(contentType *PWSTR) HRESULT
	SetContentType(contentType PWSTR) HRESULT
	GetCreated(created *SYSTEMTIME) HRESULT
	SetCreated(created *SYSTEMTIME) HRESULT
	GetCreator(creator *PWSTR) HRESULT
	SetCreator(creator PWSTR) HRESULT
	GetDescription(description *PWSTR) HRESULT
	SetDescription(description PWSTR) HRESULT
	GetIdentifier(identifier *PWSTR) HRESULT
	SetIdentifier(identifier PWSTR) HRESULT
	GetKeywords(keywords *PWSTR) HRESULT
	SetKeywords(keywords PWSTR) HRESULT
	GetLanguage(language *PWSTR) HRESULT
	SetLanguage(language PWSTR) HRESULT
	GetLastModifiedBy(lastModifiedBy *PWSTR) HRESULT
	SetLastModifiedBy(lastModifiedBy PWSTR) HRESULT
	GetLastPrinted(lastPrinted *SYSTEMTIME) HRESULT
	SetLastPrinted(lastPrinted *SYSTEMTIME) HRESULT
	GetModified(modified *SYSTEMTIME) HRESULT
	SetModified(modified *SYSTEMTIME) HRESULT
	GetRevision(revision *PWSTR) HRESULT
	SetRevision(revision PWSTR) HRESULT
	GetSubject(subject *PWSTR) HRESULT
	SetSubject(subject PWSTR) HRESULT
	GetTitle(title *PWSTR) HRESULT
	SetTitle(title PWSTR) HRESULT
	GetVersion(version *PWSTR) HRESULT
	SetVersion(version PWSTR) HRESULT
	Clone(coreProperties **IXpsOMCoreProperties) HRESULT
}

type IXpsOMCorePropertiesVtbl struct {
	IXpsOMPartVtbl
	GetOwner          uintptr
	GetCategory       uintptr
	SetCategory       uintptr
	GetContentStatus  uintptr
	SetContentStatus  uintptr
	GetContentType    uintptr
	SetContentType    uintptr
	GetCreated        uintptr
	SetCreated        uintptr
	GetCreator        uintptr
	SetCreator        uintptr
	GetDescription    uintptr
	SetDescription    uintptr
	GetIdentifier     uintptr
	SetIdentifier     uintptr
	GetKeywords       uintptr
	SetKeywords       uintptr
	GetLanguage       uintptr
	SetLanguage       uintptr
	GetLastModifiedBy uintptr
	SetLastModifiedBy uintptr
	GetLastPrinted    uintptr
	SetLastPrinted    uintptr
	GetModified       uintptr
	SetModified       uintptr
	GetRevision       uintptr
	SetRevision       uintptr
	GetSubject        uintptr
	SetSubject        uintptr
	GetTitle          uintptr
	SetTitle          uintptr
	GetVersion        uintptr
	SetVersion        uintptr
	Clone             uintptr
}

type IXpsOMCoreProperties struct {
	IXpsOMPart
}

func (this *IXpsOMCoreProperties) Vtbl() *IXpsOMCorePropertiesVtbl {
	return (*IXpsOMCorePropertiesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMCoreProperties) GetOwner(package_ **IXpsOMPackage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOwner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetCategory(category *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(category)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetCategory(category PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCategory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(category)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetContentStatus(contentStatus *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetContentStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentStatus)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetContentStatus(contentStatus PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContentStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentStatus)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetContentType(contentType *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentType)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetContentType(contentType PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetContentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentType)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetCreated(created *SYSTEMTIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(created)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetCreated(created *SYSTEMTIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCreated, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(created)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetCreator(creator *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCreator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(creator)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetCreator(creator PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCreator, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(creator)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetDescription(description *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetDescription(description PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(description)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetIdentifier(identifier *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIdentifier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(identifier)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetIdentifier(identifier PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIdentifier, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(identifier)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetKeywords(keywords *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetKeywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keywords)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetKeywords(keywords PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetKeywords, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keywords)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetLanguage(language *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetLanguage(language PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetLanguage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(language)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetLastModifiedBy(lastModifiedBy *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLastModifiedBy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lastModifiedBy)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetLastModifiedBy(lastModifiedBy PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetLastModifiedBy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lastModifiedBy)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetLastPrinted(lastPrinted *SYSTEMTIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLastPrinted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lastPrinted)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetLastPrinted(lastPrinted *SYSTEMTIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetLastPrinted, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lastPrinted)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetModified(modified *SYSTEMTIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetModified, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modified)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetModified(modified *SYSTEMTIME) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetModified, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(modified)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetRevision(revision *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRevision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(revision)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetRevision(revision PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetRevision, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(revision)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetSubject(subject *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSubject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(subject)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetSubject(subject PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSubject, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(subject)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetTitle(title *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(title)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetTitle(title PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTitle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(title)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) GetVersion(version *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(version)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) SetVersion(version PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(version)))
	return HRESULT(ret)
}

func (this *IXpsOMCoreProperties) Clone(coreProperties **IXpsOMCoreProperties) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coreProperties)))
	return HRESULT(ret)
}

// 18C3DF65-81E1-4674-91DC-FC452F5A416F
var IID_IXpsOMPackage = syscall.GUID{0x18C3DF65, 0x81E1, 0x4674,
	[8]byte{0x91, 0xDC, 0xFC, 0x45, 0x2F, 0x5A, 0x41, 0x6F}}

type IXpsOMPackageInterface interface {
	IUnknownInterface
	GetDocumentSequence(documentSequence **IXpsOMDocumentSequence) HRESULT
	SetDocumentSequence(documentSequence *IXpsOMDocumentSequence) HRESULT
	GetCoreProperties(coreProperties **IXpsOMCoreProperties) HRESULT
	SetCoreProperties(coreProperties *IXpsOMCoreProperties) HRESULT
	GetDiscardControlPartName(discardControlPartUri unsafe.Pointer) HRESULT
	SetDiscardControlPartName(discardControlPartUri unsafe.Pointer) HRESULT
	GetThumbnailResource(imageResource **IXpsOMImageResource) HRESULT
	SetThumbnailResource(imageResource *IXpsOMImageResource) HRESULT
	WriteToFile(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL) HRESULT
	WriteToStream(stream *ISequentialStream, optimizeMarkupSize BOOL) HRESULT
}

type IXpsOMPackageVtbl struct {
	IUnknownVtbl
	GetDocumentSequence       uintptr
	SetDocumentSequence       uintptr
	GetCoreProperties         uintptr
	SetCoreProperties         uintptr
	GetDiscardControlPartName uintptr
	SetDiscardControlPartName uintptr
	GetThumbnailResource      uintptr
	SetThumbnailResource      uintptr
	WriteToFile               uintptr
	WriteToStream             uintptr
}

type IXpsOMPackage struct {
	IUnknown
}

func (this *IXpsOMPackage) Vtbl() *IXpsOMPackageVtbl {
	return (*IXpsOMPackageVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPackage) GetDocumentSequence(documentSequence **IXpsOMDocumentSequence) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentSequence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentSequence)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) SetDocumentSequence(documentSequence *IXpsOMDocumentSequence) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDocumentSequence, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentSequence)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) GetCoreProperties(coreProperties **IXpsOMCoreProperties) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCoreProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coreProperties)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) SetCoreProperties(coreProperties *IXpsOMCoreProperties) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCoreProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coreProperties)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) GetDiscardControlPartName(discardControlPartUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDiscardControlPartName, uintptr(unsafe.Pointer(this)), uintptr(discardControlPartUri))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) SetDiscardControlPartName(discardControlPartUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDiscardControlPartName, uintptr(unsafe.Pointer(this)), uintptr(discardControlPartUri))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) GetThumbnailResource(imageResource **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetThumbnailResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) SetThumbnailResource(imageResource *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetThumbnailResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) WriteToFile(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteToFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(securityAttributes)), uintptr(flagsAndAttributes), uintptr(optimizeMarkupSize))
	return HRESULT(ret)
}

func (this *IXpsOMPackage) WriteToStream(stream *ISequentialStream, optimizeMarkupSize BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteToStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(optimizeMarkupSize))
	return HRESULT(ret)
}

// F9B2A685-A50D-4FC2-B764-B56E093EA0CA
var IID_IXpsOMObjectFactory = syscall.GUID{0xF9B2A685, 0xA50D, 0x4FC2,
	[8]byte{0xB7, 0x64, 0xB5, 0x6E, 0x09, 0x3E, 0xA0, 0xCA}}

type IXpsOMObjectFactoryInterface interface {
	IUnknownInterface
	CreatePackage(package_ **IXpsOMPackage) HRESULT
	CreatePackageFromFile(filename PWSTR, reuseObjects BOOL, package_ **IXpsOMPackage) HRESULT
	CreatePackageFromStream(stream *IStream, reuseObjects BOOL, package_ **IXpsOMPackage) HRESULT
	CreateStoryFragmentsResource(acquiredStream *IStream, partUri unsafe.Pointer, storyFragmentsResource **IXpsOMStoryFragmentsResource) HRESULT
	CreateDocumentStructureResource(acquiredStream *IStream, partUri unsafe.Pointer, documentStructureResource **IXpsOMDocumentStructureResource) HRESULT
	CreateSignatureBlockResource(acquiredStream *IStream, partUri unsafe.Pointer, signatureBlockResource **IXpsOMSignatureBlockResource) HRESULT
	CreateRemoteDictionaryResource(dictionary *IXpsOMDictionary, partUri unsafe.Pointer, remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT
	CreateRemoteDictionaryResourceFromStream(dictionaryMarkupStream *IStream, dictionaryPartUri unsafe.Pointer, resources *IXpsOMPartResources, dictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT
	CreatePartResources(partResources **IXpsOMPartResources) HRESULT
	CreateDocumentSequence(partUri unsafe.Pointer, documentSequence **IXpsOMDocumentSequence) HRESULT
	CreateDocument(partUri unsafe.Pointer, document **IXpsOMDocument) HRESULT
	CreatePageReference(advisoryPageDimensions *XPS_SIZE, pageReference **IXpsOMPageReference) HRESULT
	CreatePage(pageDimensions *XPS_SIZE, language PWSTR, partUri unsafe.Pointer, page **IXpsOMPage) HRESULT
	CreatePageFromStream(pageMarkupStream *IStream, partUri unsafe.Pointer, resources *IXpsOMPartResources, reuseObjects BOOL, page **IXpsOMPage) HRESULT
	CreateCanvas(canvas **IXpsOMCanvas) HRESULT
	CreateGlyphs(fontResource *IXpsOMFontResource, glyphs **IXpsOMGlyphs) HRESULT
	CreatePath(path **IXpsOMPath) HRESULT
	CreateGeometry(geometry **IXpsOMGeometry) HRESULT
	CreateGeometryFigure(startPoint *XPS_POINT, figure **IXpsOMGeometryFigure) HRESULT
	CreateMatrixTransform(matrix *XPS_MATRIX, transform **IXpsOMMatrixTransform) HRESULT
	CreateSolidColorBrush(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource, solidColorBrush **IXpsOMSolidColorBrush) HRESULT
	CreateColorProfileResource(acquiredStream *IStream, partUri unsafe.Pointer, colorProfileResource **IXpsOMColorProfileResource) HRESULT
	CreateImageBrush(image *IXpsOMImageResource, viewBox *XPS_RECT, viewPort *XPS_RECT, imageBrush **IXpsOMImageBrush) HRESULT
	CreateVisualBrush(viewBox *XPS_RECT, viewPort *XPS_RECT, visualBrush **IXpsOMVisualBrush) HRESULT
	CreateImageResource(acquiredStream *IStream, contentType XPS_IMAGE_TYPE, partUri unsafe.Pointer, imageResource **IXpsOMImageResource) HRESULT
	CreatePrintTicketResource(acquiredStream *IStream, partUri unsafe.Pointer, printTicketResource **IXpsOMPrintTicketResource) HRESULT
	CreateFontResource(acquiredStream *IStream, fontEmbedding XPS_FONT_EMBEDDING, partUri unsafe.Pointer, isObfSourceStream BOOL, fontResource **IXpsOMFontResource) HRESULT
	CreateGradientStop(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource, offset float32, gradientStop **IXpsOMGradientStop) HRESULT
	CreateLinearGradientBrush(gradStop1 *IXpsOMGradientStop, gradStop2 *IXpsOMGradientStop, startPoint *XPS_POINT, endPoint *XPS_POINT, linearGradientBrush **IXpsOMLinearGradientBrush) HRESULT
	CreateRadialGradientBrush(gradStop1 *IXpsOMGradientStop, gradStop2 *IXpsOMGradientStop, centerPoint *XPS_POINT, gradientOrigin *XPS_POINT, radiiSizes *XPS_SIZE, radialGradientBrush **IXpsOMRadialGradientBrush) HRESULT
	CreateCoreProperties(partUri unsafe.Pointer, coreProperties **IXpsOMCoreProperties) HRESULT
	CreateDictionary(dictionary **IXpsOMDictionary) HRESULT
	CreatePartUriCollection(partUriCollection **IXpsOMPartUriCollection) HRESULT
	CreatePackageWriterOnFile(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT
	CreatePackageWriterOnStream(outputStream *ISequentialStream, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT
	CreatePartUri(uri PWSTR, partUri unsafe.Pointer) HRESULT
	CreateReadOnlyStreamOnFile(filename PWSTR, stream **IStream) HRESULT
}

type IXpsOMObjectFactoryVtbl struct {
	IUnknownVtbl
	CreatePackage                            uintptr
	CreatePackageFromFile                    uintptr
	CreatePackageFromStream                  uintptr
	CreateStoryFragmentsResource             uintptr
	CreateDocumentStructureResource          uintptr
	CreateSignatureBlockResource             uintptr
	CreateRemoteDictionaryResource           uintptr
	CreateRemoteDictionaryResourceFromStream uintptr
	CreatePartResources                      uintptr
	CreateDocumentSequence                   uintptr
	CreateDocument                           uintptr
	CreatePageReference                      uintptr
	CreatePage                               uintptr
	CreatePageFromStream                     uintptr
	CreateCanvas                             uintptr
	CreateGlyphs                             uintptr
	CreatePath                               uintptr
	CreateGeometry                           uintptr
	CreateGeometryFigure                     uintptr
	CreateMatrixTransform                    uintptr
	CreateSolidColorBrush                    uintptr
	CreateColorProfileResource               uintptr
	CreateImageBrush                         uintptr
	CreateVisualBrush                        uintptr
	CreateImageResource                      uintptr
	CreatePrintTicketResource                uintptr
	CreateFontResource                       uintptr
	CreateGradientStop                       uintptr
	CreateLinearGradientBrush                uintptr
	CreateRadialGradientBrush                uintptr
	CreateCoreProperties                     uintptr
	CreateDictionary                         uintptr
	CreatePartUriCollection                  uintptr
	CreatePackageWriterOnFile                uintptr
	CreatePackageWriterOnStream              uintptr
	CreatePartUri                            uintptr
	CreateReadOnlyStreamOnFile               uintptr
}

type IXpsOMObjectFactory struct {
	IUnknown
}

func (this *IXpsOMObjectFactory) Vtbl() *IXpsOMObjectFactoryVtbl {
	return (*IXpsOMObjectFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMObjectFactory) CreatePackage(package_ **IXpsOMPackage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePackageFromFile(filename PWSTR, reuseObjects BOOL, package_ **IXpsOMPackage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageFromFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filename)), uintptr(reuseObjects), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePackageFromStream(stream *IStream, reuseObjects BOOL, package_ **IXpsOMPackage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(reuseObjects), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateStoryFragmentsResource(acquiredStream *IStream, partUri unsafe.Pointer, storyFragmentsResource **IXpsOMStoryFragmentsResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateStoryFragmentsResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(partUri), uintptr(unsafe.Pointer(storyFragmentsResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateDocumentStructureResource(acquiredStream *IStream, partUri unsafe.Pointer, documentStructureResource **IXpsOMDocumentStructureResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateDocumentStructureResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(partUri), uintptr(unsafe.Pointer(documentStructureResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateSignatureBlockResource(acquiredStream *IStream, partUri unsafe.Pointer, signatureBlockResource **IXpsOMSignatureBlockResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateSignatureBlockResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(partUri), uintptr(unsafe.Pointer(signatureBlockResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateRemoteDictionaryResource(dictionary *IXpsOMDictionary, partUri unsafe.Pointer, remoteDictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateRemoteDictionaryResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionary)), uintptr(partUri), uintptr(unsafe.Pointer(remoteDictionaryResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateRemoteDictionaryResourceFromStream(dictionaryMarkupStream *IStream, dictionaryPartUri unsafe.Pointer, resources *IXpsOMPartResources, dictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateRemoteDictionaryResourceFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionaryMarkupStream)), uintptr(dictionaryPartUri), uintptr(unsafe.Pointer(resources)), uintptr(unsafe.Pointer(dictionaryResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePartResources(partResources **IXpsOMPartResources) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePartResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(partResources)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateDocumentSequence(partUri unsafe.Pointer, documentSequence **IXpsOMDocumentSequence) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateDocumentSequence, uintptr(unsafe.Pointer(this)), uintptr(partUri), uintptr(unsafe.Pointer(documentSequence)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateDocument(partUri unsafe.Pointer, document **IXpsOMDocument) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateDocument, uintptr(unsafe.Pointer(this)), uintptr(partUri), uintptr(unsafe.Pointer(document)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePageReference(advisoryPageDimensions *XPS_SIZE, pageReference **IXpsOMPageReference) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePageReference, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(advisoryPageDimensions)), uintptr(unsafe.Pointer(pageReference)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePage(pageDimensions *XPS_SIZE, language PWSTR, partUri unsafe.Pointer, page **IXpsOMPage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageDimensions)), uintptr(unsafe.Pointer(language)), uintptr(partUri), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePageFromStream(pageMarkupStream *IStream, partUri unsafe.Pointer, resources *IXpsOMPartResources, reuseObjects BOOL, page **IXpsOMPage) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePageFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageMarkupStream)), uintptr(partUri), uintptr(unsafe.Pointer(resources)), uintptr(reuseObjects), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateCanvas(canvas **IXpsOMCanvas) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateCanvas, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(canvas)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateGlyphs(fontResource *IXpsOMFontResource, glyphs **IXpsOMGlyphs) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateGlyphs, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fontResource)), uintptr(unsafe.Pointer(glyphs)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePath(path **IXpsOMPath) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(path)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateGeometry(geometry **IXpsOMGeometry) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateGeometry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(geometry)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateGeometryFigure(startPoint *XPS_POINT, figure **IXpsOMGeometryFigure) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateGeometryFigure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(startPoint)), uintptr(unsafe.Pointer(figure)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateMatrixTransform(matrix *XPS_MATRIX, transform **IXpsOMMatrixTransform) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateMatrixTransform, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(matrix)), uintptr(unsafe.Pointer(transform)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateSolidColorBrush(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource, solidColorBrush **IXpsOMSolidColorBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateSolidColorBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(colorProfile)), uintptr(unsafe.Pointer(solidColorBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateColorProfileResource(acquiredStream *IStream, partUri unsafe.Pointer, colorProfileResource **IXpsOMColorProfileResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateColorProfileResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(partUri), uintptr(unsafe.Pointer(colorProfileResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateImageBrush(image *IXpsOMImageResource, viewBox *XPS_RECT, viewPort *XPS_RECT, imageBrush **IXpsOMImageBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateImageBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(image)), uintptr(unsafe.Pointer(viewBox)), uintptr(unsafe.Pointer(viewPort)), uintptr(unsafe.Pointer(imageBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateVisualBrush(viewBox *XPS_RECT, viewPort *XPS_RECT, visualBrush **IXpsOMVisualBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateVisualBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(viewBox)), uintptr(unsafe.Pointer(viewPort)), uintptr(unsafe.Pointer(visualBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateImageResource(acquiredStream *IStream, contentType XPS_IMAGE_TYPE, partUri unsafe.Pointer, imageResource **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateImageResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(contentType), uintptr(partUri), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePrintTicketResource(acquiredStream *IStream, partUri unsafe.Pointer, printTicketResource **IXpsOMPrintTicketResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePrintTicketResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(partUri), uintptr(unsafe.Pointer(printTicketResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateFontResource(acquiredStream *IStream, fontEmbedding XPS_FONT_EMBEDDING, partUri unsafe.Pointer, isObfSourceStream BOOL, fontResource **IXpsOMFontResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateFontResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(acquiredStream)), uintptr(fontEmbedding), uintptr(partUri), uintptr(isObfSourceStream), uintptr(unsafe.Pointer(fontResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateGradientStop(color *XPS_COLOR, colorProfile *IXpsOMColorProfileResource, offset float32, gradientStop **IXpsOMGradientStop) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateGradientStop, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(colorProfile)), uintptr(offset), uintptr(unsafe.Pointer(gradientStop)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateLinearGradientBrush(gradStop1 *IXpsOMGradientStop, gradStop2 *IXpsOMGradientStop, startPoint *XPS_POINT, endPoint *XPS_POINT, linearGradientBrush **IXpsOMLinearGradientBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateLinearGradientBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gradStop1)), uintptr(unsafe.Pointer(gradStop2)), uintptr(unsafe.Pointer(startPoint)), uintptr(unsafe.Pointer(endPoint)), uintptr(unsafe.Pointer(linearGradientBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateRadialGradientBrush(gradStop1 *IXpsOMGradientStop, gradStop2 *IXpsOMGradientStop, centerPoint *XPS_POINT, gradientOrigin *XPS_POINT, radiiSizes *XPS_SIZE, radialGradientBrush **IXpsOMRadialGradientBrush) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateRadialGradientBrush, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(gradStop1)), uintptr(unsafe.Pointer(gradStop2)), uintptr(unsafe.Pointer(centerPoint)), uintptr(unsafe.Pointer(gradientOrigin)), uintptr(unsafe.Pointer(radiiSizes)), uintptr(unsafe.Pointer(radialGradientBrush)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateCoreProperties(partUri unsafe.Pointer, coreProperties **IXpsOMCoreProperties) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateCoreProperties, uintptr(unsafe.Pointer(this)), uintptr(partUri), uintptr(unsafe.Pointer(coreProperties)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateDictionary(dictionary **IXpsOMDictionary) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateDictionary, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionary)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePartUriCollection(partUriCollection **IXpsOMPartUriCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePartUriCollection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(partUriCollection)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePackageWriterOnFile(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageWriterOnFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(securityAttributes)), uintptr(flagsAndAttributes), uintptr(optimizeMarkupSize), uintptr(interleaving), uintptr(documentSequencePartName), uintptr(unsafe.Pointer(coreProperties)), uintptr(unsafe.Pointer(packageThumbnail)), uintptr(unsafe.Pointer(documentSequencePrintTicket)), uintptr(discardControlPartName), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePackageWriterOnStream(outputStream *ISequentialStream, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageWriterOnStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(optimizeMarkupSize), uintptr(interleaving), uintptr(documentSequencePartName), uintptr(unsafe.Pointer(coreProperties)), uintptr(unsafe.Pointer(packageThumbnail)), uintptr(unsafe.Pointer(documentSequencePrintTicket)), uintptr(discardControlPartName), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreatePartUri(uri PWSTR, partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePartUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)), uintptr(partUri))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory) CreateReadOnlyStreamOnFile(filename PWSTR, stream **IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateReadOnlyStreamOnFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

// 4BDDF8EC-C915-421B-A166-D173D25653D2
var IID_IXpsOMNameCollection = syscall.GUID{0x4BDDF8EC, 0xC915, 0x421B,
	[8]byte{0xA1, 0x66, 0xD1, 0x73, 0xD2, 0x56, 0x53, 0xD2}}

type IXpsOMNameCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, name *PWSTR) HRESULT
}

type IXpsOMNameCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
}

type IXpsOMNameCollection struct {
	IUnknown
}

func (this *IXpsOMNameCollection) Vtbl() *IXpsOMNameCollectionVtbl {
	return (*IXpsOMNameCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMNameCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMNameCollection) GetAt(index uint32, name *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

// 57C650D4-067C-4893-8C33-F62A0633730F
var IID_IXpsOMPartUriCollection = syscall.GUID{0x57C650D4, 0x067C, 0x4893,
	[8]byte{0x8C, 0x33, 0xF6, 0x2A, 0x06, 0x33, 0x73, 0x0F}}

type IXpsOMPartUriCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, partUri unsafe.Pointer) HRESULT
	InsertAt(index uint32, partUri unsafe.Pointer) HRESULT
	RemoveAt(index uint32) HRESULT
	SetAt(index uint32, partUri unsafe.Pointer) HRESULT
	Append(partUri unsafe.Pointer) HRESULT
}

type IXpsOMPartUriCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	InsertAt uintptr
	RemoveAt uintptr
	SetAt    uintptr
	Append   uintptr
}

type IXpsOMPartUriCollection struct {
	IUnknown
}

func (this *IXpsOMPartUriCollection) Vtbl() *IXpsOMPartUriCollectionVtbl {
	return (*IXpsOMPartUriCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPartUriCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsOMPartUriCollection) GetAt(index uint32, partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(partUri))
	return HRESULT(ret)
}

func (this *IXpsOMPartUriCollection) InsertAt(index uint32, partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(partUri))
	return HRESULT(ret)
}

func (this *IXpsOMPartUriCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IXpsOMPartUriCollection) SetAt(index uint32, partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(partUri))
	return HRESULT(ret)
}

func (this *IXpsOMPartUriCollection) Append(partUri unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Append, uintptr(unsafe.Pointer(this)), uintptr(partUri))
	return HRESULT(ret)
}

// 4E2AA182-A443-42C6-B41B-4F8E9DE73FF9
var IID_IXpsOMPackageWriter = syscall.GUID{0x4E2AA182, 0xA443, 0x42C6,
	[8]byte{0xB4, 0x1B, 0x4F, 0x8E, 0x9D, 0xE7, 0x3F, 0xF9}}

type IXpsOMPackageWriterInterface interface {
	IUnknownInterface
	StartNewDocument(documentPartName unsafe.Pointer, documentPrintTicket *IXpsOMPrintTicketResource, documentStructure *IXpsOMDocumentStructureResource, signatureBlockResources *IXpsOMSignatureBlockResourceCollection, restrictedFonts *IXpsOMPartUriCollection) HRESULT
	AddPage(page *IXpsOMPage, advisoryPageDimensions *XPS_SIZE, discardableResourceParts *IXpsOMPartUriCollection, storyFragments *IXpsOMStoryFragmentsResource, pagePrintTicket *IXpsOMPrintTicketResource, pageThumbnail *IXpsOMImageResource) HRESULT
	AddResource(resource *IXpsOMResource) HRESULT
	Close() HRESULT
	IsClosed(isClosed *BOOL) HRESULT
}

type IXpsOMPackageWriterVtbl struct {
	IUnknownVtbl
	StartNewDocument uintptr
	AddPage          uintptr
	AddResource      uintptr
	Close            uintptr
	IsClosed         uintptr
}

type IXpsOMPackageWriter struct {
	IUnknown
}

func (this *IXpsOMPackageWriter) Vtbl() *IXpsOMPackageWriterVtbl {
	return (*IXpsOMPackageWriterVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPackageWriter) StartNewDocument(documentPartName unsafe.Pointer, documentPrintTicket *IXpsOMPrintTicketResource, documentStructure *IXpsOMDocumentStructureResource, signatureBlockResources *IXpsOMSignatureBlockResourceCollection, restrictedFonts *IXpsOMPartUriCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().StartNewDocument, uintptr(unsafe.Pointer(this)), uintptr(documentPartName), uintptr(unsafe.Pointer(documentPrintTicket)), uintptr(unsafe.Pointer(documentStructure)), uintptr(unsafe.Pointer(signatureBlockResources)), uintptr(unsafe.Pointer(restrictedFonts)))
	return HRESULT(ret)
}

func (this *IXpsOMPackageWriter) AddPage(page *IXpsOMPage, advisoryPageDimensions *XPS_SIZE, discardableResourceParts *IXpsOMPartUriCollection, storyFragments *IXpsOMStoryFragmentsResource, pagePrintTicket *IXpsOMPrintTicketResource, pageThumbnail *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddPage, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(page)), uintptr(unsafe.Pointer(advisoryPageDimensions)), uintptr(unsafe.Pointer(discardableResourceParts)), uintptr(unsafe.Pointer(storyFragments)), uintptr(unsafe.Pointer(pagePrintTicket)), uintptr(unsafe.Pointer(pageThumbnail)))
	return HRESULT(ret)
}

func (this *IXpsOMPackageWriter) AddResource(resource *IXpsOMResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddResource, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resource)))
	return HRESULT(ret)
}

func (this *IXpsOMPackageWriter) Close() HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IXpsOMPackageWriter) IsClosed(isClosed *BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsClosed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isClosed)))
	return HRESULT(ret)
}

// 219A9DB0-4959-47D0-8034-B1CE84F41A4D
var IID_IXpsOMPackageTarget = syscall.GUID{0x219A9DB0, 0x4959, 0x47D0,
	[8]byte{0x80, 0x34, 0xB1, 0xCE, 0x84, 0xF4, 0x1A, 0x4D}}

type IXpsOMPackageTargetInterface interface {
	IUnknownInterface
	CreateXpsOMPackageWriter(documentSequencePartName unsafe.Pointer, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT
}

type IXpsOMPackageTargetVtbl struct {
	IUnknownVtbl
	CreateXpsOMPackageWriter uintptr
}

type IXpsOMPackageTarget struct {
	IUnknown
}

func (this *IXpsOMPackageTarget) Vtbl() *IXpsOMPackageTargetVtbl {
	return (*IXpsOMPackageTargetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPackageTarget) CreateXpsOMPackageWriter(documentSequencePartName unsafe.Pointer, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateXpsOMPackageWriter, uintptr(unsafe.Pointer(this)), uintptr(documentSequencePartName), uintptr(unsafe.Pointer(documentSequencePrintTicket)), uintptr(discardControlPartName), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

// 15B873D5-1971-41E8-83A3-6578403064C7
var IID_IXpsOMThumbnailGenerator = syscall.GUID{0x15B873D5, 0x1971, 0x41E8,
	[8]byte{0x83, 0xA3, 0x65, 0x78, 0x40, 0x30, 0x64, 0xC7}}

type IXpsOMThumbnailGeneratorInterface interface {
	IUnknownInterface
	GenerateThumbnail(page *IXpsOMPage, thumbnailType XPS_IMAGE_TYPE, thumbnailSize XPS_THUMBNAIL_SIZE, imageResourcePartName unsafe.Pointer, imageResource **IXpsOMImageResource) HRESULT
}

type IXpsOMThumbnailGeneratorVtbl struct {
	IUnknownVtbl
	GenerateThumbnail uintptr
}

type IXpsOMThumbnailGenerator struct {
	IUnknown
}

func (this *IXpsOMThumbnailGenerator) Vtbl() *IXpsOMThumbnailGeneratorVtbl {
	return (*IXpsOMThumbnailGeneratorVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMThumbnailGenerator) GenerateThumbnail(page *IXpsOMPage, thumbnailType XPS_IMAGE_TYPE, thumbnailSize XPS_THUMBNAIL_SIZE, imageResourcePartName unsafe.Pointer, imageResource **IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GenerateThumbnail, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(page)), uintptr(thumbnailType), uintptr(thumbnailSize), uintptr(imageResourcePartName), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

// 0A91B617-D612-4181-BF7C-BE5824E9CC8F
var IID_IXpsOMObjectFactory1 = syscall.GUID{0x0A91B617, 0xD612, 0x4181,
	[8]byte{0xBF, 0x7C, 0xBE, 0x58, 0x24, 0xE9, 0xCC, 0x8F}}

type IXpsOMObjectFactory1Interface interface {
	IXpsOMObjectFactoryInterface
	GetDocumentTypeFromFile(filename PWSTR, documentType *XPS_DOCUMENT_TYPE) HRESULT
	GetDocumentTypeFromStream(xpsDocumentStream *IStream, documentType *XPS_DOCUMENT_TYPE) HRESULT
	ConvertHDPhotoToJpegXR(imageResource *IXpsOMImageResource) HRESULT
	ConvertJpegXRToHDPhoto(imageResource *IXpsOMImageResource) HRESULT
	CreatePackageWriterOnFile1(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, documentType XPS_DOCUMENT_TYPE, packageWriter **IXpsOMPackageWriter) HRESULT
	CreatePackageWriterOnStream1(outputStream *ISequentialStream, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, documentType XPS_DOCUMENT_TYPE, packageWriter **IXpsOMPackageWriter) HRESULT
	CreatePackage1(package_ **IXpsOMPackage1) HRESULT
	CreatePackageFromStream1(stream *IStream, reuseObjects BOOL, package_ **IXpsOMPackage1) HRESULT
	CreatePackageFromFile1(filename PWSTR, reuseObjects BOOL, package_ **IXpsOMPackage1) HRESULT
	CreatePage1(pageDimensions *XPS_SIZE, language PWSTR, partUri unsafe.Pointer, page **IXpsOMPage1) HRESULT
	CreatePageFromStream1(pageMarkupStream *IStream, partUri unsafe.Pointer, resources *IXpsOMPartResources, reuseObjects BOOL, page **IXpsOMPage1) HRESULT
	CreateRemoteDictionaryResourceFromStream1(dictionaryMarkupStream *IStream, partUri unsafe.Pointer, resources *IXpsOMPartResources, dictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT
}

type IXpsOMObjectFactory1Vtbl struct {
	IXpsOMObjectFactoryVtbl
	GetDocumentTypeFromFile                   uintptr
	GetDocumentTypeFromStream                 uintptr
	ConvertHDPhotoToJpegXR                    uintptr
	ConvertJpegXRToHDPhoto                    uintptr
	CreatePackageWriterOnFile1                uintptr
	CreatePackageWriterOnStream1              uintptr
	CreatePackage1                            uintptr
	CreatePackageFromStream1                  uintptr
	CreatePackageFromFile1                    uintptr
	CreatePage1                               uintptr
	CreatePageFromStream1                     uintptr
	CreateRemoteDictionaryResourceFromStream1 uintptr
}

type IXpsOMObjectFactory1 struct {
	IXpsOMObjectFactory
}

func (this *IXpsOMObjectFactory1) Vtbl() *IXpsOMObjectFactory1Vtbl {
	return (*IXpsOMObjectFactory1Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMObjectFactory1) GetDocumentTypeFromFile(filename PWSTR, documentType *XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentTypeFromFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(documentType)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) GetDocumentTypeFromStream(xpsDocumentStream *IStream, documentType *XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentTypeFromStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(xpsDocumentStream)), uintptr(unsafe.Pointer(documentType)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) ConvertHDPhotoToJpegXR(imageResource *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ConvertHDPhotoToJpegXR, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) ConvertJpegXRToHDPhoto(imageResource *IXpsOMImageResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().ConvertJpegXRToHDPhoto, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageResource)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePackageWriterOnFile1(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, documentType XPS_DOCUMENT_TYPE, packageWriter **IXpsOMPackageWriter) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageWriterOnFile1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(securityAttributes)), uintptr(flagsAndAttributes), uintptr(optimizeMarkupSize), uintptr(interleaving), uintptr(documentSequencePartName), uintptr(unsafe.Pointer(coreProperties)), uintptr(unsafe.Pointer(packageThumbnail)), uintptr(unsafe.Pointer(documentSequencePrintTicket)), uintptr(discardControlPartName), uintptr(documentType), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePackageWriterOnStream1(outputStream *ISequentialStream, optimizeMarkupSize BOOL, interleaving XPS_INTERLEAVING, documentSequencePartName unsafe.Pointer, coreProperties *IXpsOMCoreProperties, packageThumbnail *IXpsOMImageResource, documentSequencePrintTicket *IXpsOMPrintTicketResource, discardControlPartName unsafe.Pointer, documentType XPS_DOCUMENT_TYPE, packageWriter **IXpsOMPackageWriter) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageWriterOnStream1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(optimizeMarkupSize), uintptr(interleaving), uintptr(documentSequencePartName), uintptr(unsafe.Pointer(coreProperties)), uintptr(unsafe.Pointer(packageThumbnail)), uintptr(unsafe.Pointer(documentSequencePrintTicket)), uintptr(discardControlPartName), uintptr(documentType), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePackage1(package_ **IXpsOMPackage1) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackage1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePackageFromStream1(stream *IStream, reuseObjects BOOL, package_ **IXpsOMPackage1) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageFromStream1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(reuseObjects), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePackageFromFile1(filename PWSTR, reuseObjects BOOL, package_ **IXpsOMPackage1) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePackageFromFile1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filename)), uintptr(reuseObjects), uintptr(unsafe.Pointer(package_)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePage1(pageDimensions *XPS_SIZE, language PWSTR, partUri unsafe.Pointer, page **IXpsOMPage1) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePage1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageDimensions)), uintptr(unsafe.Pointer(language)), uintptr(partUri), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreatePageFromStream1(pageMarkupStream *IStream, partUri unsafe.Pointer, resources *IXpsOMPartResources, reuseObjects BOOL, page **IXpsOMPage1) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePageFromStream1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageMarkupStream)), uintptr(partUri), uintptr(unsafe.Pointer(resources)), uintptr(reuseObjects), uintptr(unsafe.Pointer(page)))
	return HRESULT(ret)
}

func (this *IXpsOMObjectFactory1) CreateRemoteDictionaryResourceFromStream1(dictionaryMarkupStream *IStream, partUri unsafe.Pointer, resources *IXpsOMPartResources, dictionaryResource **IXpsOMRemoteDictionaryResource) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateRemoteDictionaryResourceFromStream1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dictionaryMarkupStream)), uintptr(partUri), uintptr(unsafe.Pointer(resources)), uintptr(unsafe.Pointer(dictionaryResource)))
	return HRESULT(ret)
}

// 95A9435E-12BB-461B-8E7F-C6ADB04CD96A
var IID_IXpsOMPackage1 = syscall.GUID{0x95A9435E, 0x12BB, 0x461B,
	[8]byte{0x8E, 0x7F, 0xC6, 0xAD, 0xB0, 0x4C, 0xD9, 0x6A}}

type IXpsOMPackage1Interface interface {
	IXpsOMPackageInterface
	GetDocumentType(documentType *XPS_DOCUMENT_TYPE) HRESULT
	WriteToFile1(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL, documentType XPS_DOCUMENT_TYPE) HRESULT
	WriteToStream1(outputStream *ISequentialStream, optimizeMarkupSize BOOL, documentType XPS_DOCUMENT_TYPE) HRESULT
}

type IXpsOMPackage1Vtbl struct {
	IXpsOMPackageVtbl
	GetDocumentType uintptr
	WriteToFile1    uintptr
	WriteToStream1  uintptr
}

type IXpsOMPackage1 struct {
	IXpsOMPackage
}

func (this *IXpsOMPackage1) Vtbl() *IXpsOMPackage1Vtbl {
	return (*IXpsOMPackage1Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPackage1) GetDocumentType(documentType *XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentType)))
	return HRESULT(ret)
}

func (this *IXpsOMPackage1) WriteToFile1(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32, optimizeMarkupSize BOOL, documentType XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteToFile1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(securityAttributes)), uintptr(flagsAndAttributes), uintptr(optimizeMarkupSize), uintptr(documentType))
	return HRESULT(ret)
}

func (this *IXpsOMPackage1) WriteToStream1(outputStream *ISequentialStream, optimizeMarkupSize BOOL, documentType XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().WriteToStream1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(outputStream)), uintptr(optimizeMarkupSize), uintptr(documentType))
	return HRESULT(ret)
}

// 305B60EF-6892-4DDA-9CBB-3AA65974508A
var IID_IXpsOMPage1 = syscall.GUID{0x305B60EF, 0x6892, 0x4DDA,
	[8]byte{0x9C, 0xBB, 0x3A, 0xA6, 0x59, 0x74, 0x50, 0x8A}}

type IXpsOMPage1Interface interface {
	IXpsOMPageInterface
	GetDocumentType(documentType *XPS_DOCUMENT_TYPE) HRESULT
	Write1(stream *ISequentialStream, optimizeMarkupSize BOOL, documentType XPS_DOCUMENT_TYPE) HRESULT
}

type IXpsOMPage1Vtbl struct {
	IXpsOMPageVtbl
	GetDocumentType uintptr
	Write1          uintptr
}

type IXpsOMPage1 struct {
	IXpsOMPage
}

func (this *IXpsOMPage1) Vtbl() *IXpsOMPage1Vtbl {
	return (*IXpsOMPage1Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPage1) GetDocumentType(documentType *XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentType)))
	return HRESULT(ret)
}

func (this *IXpsOMPage1) Write1(stream *ISequentialStream, optimizeMarkupSize BOOL, documentType XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(optimizeMarkupSize), uintptr(documentType))
	return HRESULT(ret)
}

// 3B0B6D38-53AD-41DA-B212-D37637A6714E
var IID_IXpsDocumentPackageTarget = syscall.GUID{0x3B0B6D38, 0x53AD, 0x41DA,
	[8]byte{0xB2, 0x12, 0xD3, 0x76, 0x37, 0xA6, 0x71, 0x4E}}

type IXpsDocumentPackageTargetInterface interface {
	IUnknownInterface
	GetXpsOMPackageWriter(documentSequencePartName unsafe.Pointer, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT
	GetXpsOMFactory(xpsFactory **IXpsOMObjectFactory) HRESULT
	GetXpsType(documentType *XPS_DOCUMENT_TYPE) HRESULT
}

type IXpsDocumentPackageTargetVtbl struct {
	IUnknownVtbl
	GetXpsOMPackageWriter uintptr
	GetXpsOMFactory       uintptr
	GetXpsType            uintptr
}

type IXpsDocumentPackageTarget struct {
	IUnknown
}

func (this *IXpsDocumentPackageTarget) Vtbl() *IXpsDocumentPackageTargetVtbl {
	return (*IXpsDocumentPackageTargetVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsDocumentPackageTarget) GetXpsOMPackageWriter(documentSequencePartName unsafe.Pointer, discardControlPartName unsafe.Pointer, packageWriter **IXpsOMPackageWriter) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetXpsOMPackageWriter, uintptr(unsafe.Pointer(this)), uintptr(documentSequencePartName), uintptr(discardControlPartName), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

func (this *IXpsDocumentPackageTarget) GetXpsOMFactory(xpsFactory **IXpsOMObjectFactory) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetXpsOMFactory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(xpsFactory)))
	return HRESULT(ret)
}

func (this *IXpsDocumentPackageTarget) GetXpsType(documentType *XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetXpsType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentType)))
	return HRESULT(ret)
}

// BF8FC1D4-9D46-4141-BA5F-94BB9250D041
var IID_IXpsOMRemoteDictionaryResource1 = syscall.GUID{0xBF8FC1D4, 0x9D46, 0x4141,
	[8]byte{0xBA, 0x5F, 0x94, 0xBB, 0x92, 0x50, 0xD0, 0x41}}

type IXpsOMRemoteDictionaryResource1Interface interface {
	IXpsOMRemoteDictionaryResourceInterface
	GetDocumentType(documentType *XPS_DOCUMENT_TYPE) HRESULT
	Write1(stream *ISequentialStream, documentType XPS_DOCUMENT_TYPE) HRESULT
}

type IXpsOMRemoteDictionaryResource1Vtbl struct {
	IXpsOMRemoteDictionaryResourceVtbl
	GetDocumentType uintptr
	Write1          uintptr
}

type IXpsOMRemoteDictionaryResource1 struct {
	IXpsOMRemoteDictionaryResource
}

func (this *IXpsOMRemoteDictionaryResource1) Vtbl() *IXpsOMRemoteDictionaryResource1Vtbl {
	return (*IXpsOMRemoteDictionaryResource1Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMRemoteDictionaryResource1) GetDocumentType(documentType *XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(documentType)))
	return HRESULT(ret)
}

func (this *IXpsOMRemoteDictionaryResource1) Write1(stream *ISequentialStream, documentType XPS_DOCUMENT_TYPE) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Write1, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)), uintptr(documentType))
	return HRESULT(ret)
}

// E8A45033-640E-43FA-9BDF-FDDEAA31C6A0
var IID_IXpsOMPackageWriter3D = syscall.GUID{0xE8A45033, 0x640E, 0x43FA,
	[8]byte{0x9B, 0xDF, 0xFD, 0xDE, 0xAA, 0x31, 0xC6, 0xA0}}

type IXpsOMPackageWriter3DInterface interface {
	IXpsOMPackageWriterInterface
	AddModelTexture(texturePartName unsafe.Pointer, textureData *IStream) HRESULT
	SetModelPrintTicket(printTicketPartName unsafe.Pointer, printTicketData *IStream) HRESULT
}

type IXpsOMPackageWriter3DVtbl struct {
	IXpsOMPackageWriterVtbl
	AddModelTexture     uintptr
	SetModelPrintTicket uintptr
}

type IXpsOMPackageWriter3D struct {
	IXpsOMPackageWriter
}

func (this *IXpsOMPackageWriter3D) Vtbl() *IXpsOMPackageWriter3DVtbl {
	return (*IXpsOMPackageWriter3DVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsOMPackageWriter3D) AddModelTexture(texturePartName unsafe.Pointer, textureData *IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddModelTexture, uintptr(unsafe.Pointer(this)), uintptr(texturePartName), uintptr(unsafe.Pointer(textureData)))
	return HRESULT(ret)
}

func (this *IXpsOMPackageWriter3D) SetModelPrintTicket(printTicketPartName unsafe.Pointer, printTicketData *IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetModelPrintTicket, uintptr(unsafe.Pointer(this)), uintptr(printTicketPartName), uintptr(unsafe.Pointer(printTicketData)))
	return HRESULT(ret)
}

// 60BA71B8-3101-4984-9199-F4EA775FF01D
var IID_IXpsDocumentPackageTarget3D = syscall.GUID{0x60BA71B8, 0x3101, 0x4984,
	[8]byte{0x91, 0x99, 0xF4, 0xEA, 0x77, 0x5F, 0xF0, 0x1D}}

type IXpsDocumentPackageTarget3DInterface interface {
	IUnknownInterface
	GetXpsOMPackageWriter3D(documentSequencePartName unsafe.Pointer, discardControlPartName unsafe.Pointer, modelPartName unsafe.Pointer, modelData *IStream, packageWriter **IXpsOMPackageWriter3D) HRESULT
	GetXpsOMFactory(xpsFactory **IXpsOMObjectFactory) HRESULT
}

type IXpsDocumentPackageTarget3DVtbl struct {
	IUnknownVtbl
	GetXpsOMPackageWriter3D uintptr
	GetXpsOMFactory         uintptr
}

type IXpsDocumentPackageTarget3D struct {
	IUnknown
}

func (this *IXpsDocumentPackageTarget3D) Vtbl() *IXpsDocumentPackageTarget3DVtbl {
	return (*IXpsDocumentPackageTarget3DVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsDocumentPackageTarget3D) GetXpsOMPackageWriter3D(documentSequencePartName unsafe.Pointer, discardControlPartName unsafe.Pointer, modelPartName unsafe.Pointer, modelData *IStream, packageWriter **IXpsOMPackageWriter3D) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetXpsOMPackageWriter3D, uintptr(unsafe.Pointer(this)), uintptr(documentSequencePartName), uintptr(discardControlPartName), uintptr(modelPartName), uintptr(unsafe.Pointer(modelData)), uintptr(unsafe.Pointer(packageWriter)))
	return HRESULT(ret)
}

func (this *IXpsDocumentPackageTarget3D) GetXpsOMFactory(xpsFactory **IXpsOMObjectFactory) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetXpsOMFactory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(xpsFactory)))
	return HRESULT(ret)
}

// 7718EAE4-3215-49BE-AF5B-594FEF7FCFA6
var IID_IXpsSigningOptions = syscall.GUID{0x7718EAE4, 0x3215, 0x49BE,
	[8]byte{0xAF, 0x5B, 0x59, 0x4F, 0xEF, 0x7F, 0xCF, 0xA6}}

type IXpsSigningOptionsInterface interface {
	IUnknownInterface
	GetSignatureId(signatureId *PWSTR) HRESULT
	SetSignatureId(signatureId PWSTR) HRESULT
	GetSignatureMethod(signatureMethod *PWSTR) HRESULT
	SetSignatureMethod(signatureMethod PWSTR) HRESULT
	GetDigestMethod(digestMethod *PWSTR) HRESULT
	SetDigestMethod(digestMethod PWSTR) HRESULT
	GetSignaturePartName(signaturePartName unsafe.Pointer) HRESULT
	SetSignaturePartName(signaturePartName unsafe.Pointer) HRESULT
	GetPolicy(policy *XPS_SIGN_POLICY) HRESULT
	SetPolicy(policy XPS_SIGN_POLICY) HRESULT
	GetSigningTimeFormat(timeFormat unsafe.Pointer) HRESULT
	SetSigningTimeFormat(timeFormat unsafe.Pointer) HRESULT
	GetCustomObjects(customObjectSet unsafe.Pointer) HRESULT
	GetCustomReferences(customReferenceSet unsafe.Pointer) HRESULT
	GetCertificateSet(certificateSet unsafe.Pointer) HRESULT
	GetFlags(flags *XPS_SIGN_FLAGS) HRESULT
	SetFlags(flags XPS_SIGN_FLAGS) HRESULT
}

type IXpsSigningOptionsVtbl struct {
	IUnknownVtbl
	GetSignatureId       uintptr
	SetSignatureId       uintptr
	GetSignatureMethod   uintptr
	SetSignatureMethod   uintptr
	GetDigestMethod      uintptr
	SetDigestMethod      uintptr
	GetSignaturePartName uintptr
	SetSignaturePartName uintptr
	GetPolicy            uintptr
	SetPolicy            uintptr
	GetSigningTimeFormat uintptr
	SetSigningTimeFormat uintptr
	GetCustomObjects     uintptr
	GetCustomReferences  uintptr
	GetCertificateSet    uintptr
	GetFlags             uintptr
	SetFlags             uintptr
}

type IXpsSigningOptions struct {
	IUnknown
}

func (this *IXpsSigningOptions) Vtbl() *IXpsSigningOptionsVtbl {
	return (*IXpsSigningOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSigningOptions) GetSignatureId(signatureId *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureId)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetSignatureId(signatureId PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSignatureId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureId)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetSignatureMethod(signatureMethod *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureMethod)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetSignatureMethod(signatureMethod PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSignatureMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureMethod)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetDigestMethod(digestMethod *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDigestMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(digestMethod)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetDigestMethod(digestMethod PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDigestMethod, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(digestMethod)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetSignaturePartName(signaturePartName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignaturePartName, uintptr(unsafe.Pointer(this)), uintptr(signaturePartName))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetSignaturePartName(signaturePartName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSignaturePartName, uintptr(unsafe.Pointer(this)), uintptr(signaturePartName))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetPolicy(policy *XPS_SIGN_POLICY) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(policy)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetPolicy(policy XPS_SIGN_POLICY) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPolicy, uintptr(unsafe.Pointer(this)), uintptr(policy))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetSigningTimeFormat(timeFormat unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSigningTimeFormat, uintptr(unsafe.Pointer(this)), uintptr(timeFormat))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetSigningTimeFormat(timeFormat unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSigningTimeFormat, uintptr(unsafe.Pointer(this)), uintptr(timeFormat))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetCustomObjects(customObjectSet unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCustomObjects, uintptr(unsafe.Pointer(this)), uintptr(customObjectSet))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetCustomReferences(customReferenceSet unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCustomReferences, uintptr(unsafe.Pointer(this)), uintptr(customReferenceSet))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetCertificateSet(certificateSet unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCertificateSet, uintptr(unsafe.Pointer(this)), uintptr(certificateSet))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) GetFlags(flags *XPS_SIGN_FLAGS) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFlags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(flags)))
	return HRESULT(ret)
}

func (this *IXpsSigningOptions) SetFlags(flags XPS_SIGN_FLAGS) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFlags, uintptr(unsafe.Pointer(this)), uintptr(flags))
	return HRESULT(ret)
}

// A2D1D95D-ADD2-4DFF-AB27-6B9C645FF322
var IID_IXpsSignatureCollection = syscall.GUID{0xA2D1D95D, 0xADD2, 0x4DFF,
	[8]byte{0xAB, 0x27, 0x6B, 0x9C, 0x64, 0x5F, 0xF3, 0x22}}

type IXpsSignatureCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, signature **IXpsSignature) HRESULT
	RemoveAt(index uint32) HRESULT
}

type IXpsSignatureCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	RemoveAt uintptr
}

type IXpsSignatureCollection struct {
	IUnknown
}

func (this *IXpsSignatureCollection) Vtbl() *IXpsSignatureCollectionVtbl {
	return (*IXpsSignatureCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignatureCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsSignatureCollection) GetAt(index uint32, signature **IXpsSignature) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(signature)))
	return HRESULT(ret)
}

func (this *IXpsSignatureCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

// 6AE4C93E-1ADE-42FB-898B-3A5658284857
var IID_IXpsSignature = syscall.GUID{0x6AE4C93E, 0x1ADE, 0x42FB,
	[8]byte{0x89, 0x8B, 0x3A, 0x56, 0x58, 0x28, 0x48, 0x57}}

type IXpsSignatureInterface interface {
	IUnknownInterface
	GetSignatureId(sigId *PWSTR) HRESULT
	GetSignatureValue(signatureHashValue **byte, count *uint32) HRESULT
	GetCertificateEnumerator(certificateEnumerator unsafe.Pointer) HRESULT
	GetSigningTime(sigDateTimeString *PWSTR) HRESULT
	GetSigningTimeFormat(timeFormat unsafe.Pointer) HRESULT
	GetSignaturePartName(signaturePartName unsafe.Pointer) HRESULT
	Verify(x509Certificate unsafe.Pointer, sigStatus *XPS_SIGNATURE_STATUS) HRESULT
	GetPolicy(policy *XPS_SIGN_POLICY) HRESULT
	GetCustomObjectEnumerator(customObjectEnumerator unsafe.Pointer) HRESULT
	GetCustomReferenceEnumerator(customReferenceEnumerator unsafe.Pointer) HRESULT
	GetSignatureXml(signatureXml **byte, count *uint32) HRESULT
	SetSignatureXml(signatureXml *byte, count uint32) HRESULT
}

type IXpsSignatureVtbl struct {
	IUnknownVtbl
	GetSignatureId               uintptr
	GetSignatureValue            uintptr
	GetCertificateEnumerator     uintptr
	GetSigningTime               uintptr
	GetSigningTimeFormat         uintptr
	GetSignaturePartName         uintptr
	Verify                       uintptr
	GetPolicy                    uintptr
	GetCustomObjectEnumerator    uintptr
	GetCustomReferenceEnumerator uintptr
	GetSignatureXml              uintptr
	SetSignatureXml              uintptr
}

type IXpsSignature struct {
	IUnknown
}

func (this *IXpsSignature) Vtbl() *IXpsSignatureVtbl {
	return (*IXpsSignatureVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignature) GetSignatureId(sigId *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sigId)))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetSignatureValue(signatureHashValue **byte, count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureHashValue)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetCertificateEnumerator(certificateEnumerator unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCertificateEnumerator, uintptr(unsafe.Pointer(this)), uintptr(certificateEnumerator))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetSigningTime(sigDateTimeString *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSigningTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sigDateTimeString)))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetSigningTimeFormat(timeFormat unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSigningTimeFormat, uintptr(unsafe.Pointer(this)), uintptr(timeFormat))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetSignaturePartName(signaturePartName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignaturePartName, uintptr(unsafe.Pointer(this)), uintptr(signaturePartName))
	return HRESULT(ret)
}

func (this *IXpsSignature) Verify(x509Certificate unsafe.Pointer, sigStatus *XPS_SIGNATURE_STATUS) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Verify, uintptr(unsafe.Pointer(this)), uintptr(x509Certificate), uintptr(unsafe.Pointer(sigStatus)))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetPolicy(policy *XPS_SIGN_POLICY) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPolicy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(policy)))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetCustomObjectEnumerator(customObjectEnumerator unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCustomObjectEnumerator, uintptr(unsafe.Pointer(this)), uintptr(customObjectEnumerator))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetCustomReferenceEnumerator(customReferenceEnumerator unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCustomReferenceEnumerator, uintptr(unsafe.Pointer(this)), uintptr(customReferenceEnumerator))
	return HRESULT(ret)
}

func (this *IXpsSignature) GetSignatureXml(signatureXml **byte, count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureXml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureXml)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsSignature) SetSignatureXml(signatureXml *byte, count uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSignatureXml, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureXml)), uintptr(count))
	return HRESULT(ret)
}

// 23397050-FE99-467A-8DCE-9237F074FFE4
var IID_IXpsSignatureBlockCollection = syscall.GUID{0x23397050, 0xFE99, 0x467A,
	[8]byte{0x8D, 0xCE, 0x92, 0x37, 0xF0, 0x74, 0xFF, 0xE4}}

type IXpsSignatureBlockCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, signatureBlock **IXpsSignatureBlock) HRESULT
	RemoveAt(index uint32) HRESULT
}

type IXpsSignatureBlockCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	RemoveAt uintptr
}

type IXpsSignatureBlockCollection struct {
	IUnknown
}

func (this *IXpsSignatureBlockCollection) Vtbl() *IXpsSignatureBlockCollectionVtbl {
	return (*IXpsSignatureBlockCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignatureBlockCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsSignatureBlockCollection) GetAt(index uint32, signatureBlock **IXpsSignatureBlock) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(signatureBlock)))
	return HRESULT(ret)
}

func (this *IXpsSignatureBlockCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

// 151FAC09-0B97-4AC6-A323-5E4297D4322B
var IID_IXpsSignatureBlock = syscall.GUID{0x151FAC09, 0x0B97, 0x4AC6,
	[8]byte{0xA3, 0x23, 0x5E, 0x42, 0x97, 0xD4, 0x32, 0x2B}}

type IXpsSignatureBlockInterface interface {
	IUnknownInterface
	GetRequests(requests **IXpsSignatureRequestCollection) HRESULT
	GetPartName(partName unsafe.Pointer) HRESULT
	GetDocumentIndex(fixedDocumentIndex *uint32) HRESULT
	GetDocumentName(fixedDocumentName unsafe.Pointer) HRESULT
	CreateRequest(requestId PWSTR, signatureRequest **IXpsSignatureRequest) HRESULT
}

type IXpsSignatureBlockVtbl struct {
	IUnknownVtbl
	GetRequests      uintptr
	GetPartName      uintptr
	GetDocumentIndex uintptr
	GetDocumentName  uintptr
	CreateRequest    uintptr
}

type IXpsSignatureBlock struct {
	IUnknown
}

func (this *IXpsSignatureBlock) Vtbl() *IXpsSignatureBlockVtbl {
	return (*IXpsSignatureBlockVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignatureBlock) GetRequests(requests **IXpsSignatureRequestCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRequests, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requests)))
	return HRESULT(ret)
}

func (this *IXpsSignatureBlock) GetPartName(partName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPartName, uintptr(unsafe.Pointer(this)), uintptr(partName))
	return HRESULT(ret)
}

func (this *IXpsSignatureBlock) GetDocumentIndex(fixedDocumentIndex *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentIndex, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fixedDocumentIndex)))
	return HRESULT(ret)
}

func (this *IXpsSignatureBlock) GetDocumentName(fixedDocumentName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetDocumentName, uintptr(unsafe.Pointer(this)), uintptr(fixedDocumentName))
	return HRESULT(ret)
}

func (this *IXpsSignatureBlock) CreateRequest(requestId PWSTR, signatureRequest **IXpsSignatureRequest) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestId)), uintptr(unsafe.Pointer(signatureRequest)))
	return HRESULT(ret)
}

// F0253E68-9F19-412E-9B4F-54D3B0AC6CD9
var IID_IXpsSignatureRequestCollection = syscall.GUID{0xF0253E68, 0x9F19, 0x412E,
	[8]byte{0x9B, 0x4F, 0x54, 0xD3, 0xB0, 0xAC, 0x6C, 0xD9}}

type IXpsSignatureRequestCollectionInterface interface {
	IUnknownInterface
	GetCount(count *uint32) HRESULT
	GetAt(index uint32, signatureRequest **IXpsSignatureRequest) HRESULT
	RemoveAt(index uint32) HRESULT
}

type IXpsSignatureRequestCollectionVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	RemoveAt uintptr
}

type IXpsSignatureRequestCollection struct {
	IUnknown
}

func (this *IXpsSignatureRequestCollection) Vtbl() *IXpsSignatureRequestCollectionVtbl {
	return (*IXpsSignatureRequestCollectionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignatureRequestCollection) GetCount(count *uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequestCollection) GetAt(index uint32, signatureRequest **IXpsSignatureRequest) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAt, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(signatureRequest)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequestCollection) RemoveAt(index uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAt, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

// AC58950B-7208-4B2D-B2C4-951083D3B8EB
var IID_IXpsSignatureRequest = syscall.GUID{0xAC58950B, 0x7208, 0x4B2D,
	[8]byte{0xB2, 0xC4, 0x95, 0x10, 0x83, 0xD3, 0xB8, 0xEB}}

type IXpsSignatureRequestInterface interface {
	IUnknownInterface
	GetIntent(intent *PWSTR) HRESULT
	SetIntent(intent PWSTR) HRESULT
	GetRequestedSigner(signerName *PWSTR) HRESULT
	SetRequestedSigner(signerName PWSTR) HRESULT
	GetRequestSignByDate(dateString *PWSTR) HRESULT
	SetRequestSignByDate(dateString PWSTR) HRESULT
	GetSigningLocale(place *PWSTR) HRESULT
	SetSigningLocale(place PWSTR) HRESULT
	GetSpotLocation(pageIndex *int32, pagePartName unsafe.Pointer, x *float32, y *float32) HRESULT
	SetSpotLocation(pageIndex int32, x float32, y float32) HRESULT
	GetRequestId(requestId *PWSTR) HRESULT
	GetSignature(signature **IXpsSignature) HRESULT
}

type IXpsSignatureRequestVtbl struct {
	IUnknownVtbl
	GetIntent            uintptr
	SetIntent            uintptr
	GetRequestedSigner   uintptr
	SetRequestedSigner   uintptr
	GetRequestSignByDate uintptr
	SetRequestSignByDate uintptr
	GetSigningLocale     uintptr
	SetSigningLocale     uintptr
	GetSpotLocation      uintptr
	SetSpotLocation      uintptr
	GetRequestId         uintptr
	GetSignature         uintptr
}

type IXpsSignatureRequest struct {
	IUnknown
}

func (this *IXpsSignatureRequest) Vtbl() *IXpsSignatureRequestVtbl {
	return (*IXpsSignatureRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignatureRequest) GetIntent(intent *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIntent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(intent)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) SetIntent(intent PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetIntent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(intent)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) GetRequestedSigner(signerName *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRequestedSigner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signerName)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) SetRequestedSigner(signerName PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetRequestedSigner, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signerName)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) GetRequestSignByDate(dateString *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRequestSignByDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dateString)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) SetRequestSignByDate(dateString PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetRequestSignByDate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(dateString)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) GetSigningLocale(place *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSigningLocale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(place)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) SetSigningLocale(place PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSigningLocale, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(place)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) GetSpotLocation(pageIndex *int32, pagePartName unsafe.Pointer, x *float32, y *float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSpotLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageIndex)), uintptr(pagePartName), uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) SetSpotLocation(pageIndex int32, x float32, y float32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSpotLocation, uintptr(unsafe.Pointer(this)), uintptr(pageIndex), uintptr(x), uintptr(y))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) GetRequestId(requestId *PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRequestId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestId)))
	return HRESULT(ret)
}

func (this *IXpsSignatureRequest) GetSignature(signature **IXpsSignature) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignature, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signature)))
	return HRESULT(ret)
}

// D3E8D338-FDC4-4AFC-80B5-D532A1782EE1
var IID_IXpsSignatureManager = syscall.GUID{0xD3E8D338, 0xFDC4, 0x4AFC,
	[8]byte{0x80, 0xB5, 0xD5, 0x32, 0xA1, 0x78, 0x2E, 0xE1}}

type IXpsSignatureManagerInterface interface {
	IUnknownInterface
	LoadPackageFile(fileName PWSTR) HRESULT
	LoadPackageStream(stream *IStream) HRESULT
	Sign(signOptions *IXpsSigningOptions, x509Certificate unsafe.Pointer, signature **IXpsSignature) HRESULT
	GetSignatureOriginPartName(signatureOriginPartName unsafe.Pointer) HRESULT
	SetSignatureOriginPartName(signatureOriginPartName unsafe.Pointer) HRESULT
	GetSignatures(signatures **IXpsSignatureCollection) HRESULT
	AddSignatureBlock(partName unsafe.Pointer, fixedDocumentIndex uint32, signatureBlock **IXpsSignatureBlock) HRESULT
	GetSignatureBlocks(signatureBlocks **IXpsSignatureBlockCollection) HRESULT
	CreateSigningOptions(signingOptions **IXpsSigningOptions) HRESULT
	SavePackageToFile(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32) HRESULT
	SavePackageToStream(stream *IStream) HRESULT
}

type IXpsSignatureManagerVtbl struct {
	IUnknownVtbl
	LoadPackageFile            uintptr
	LoadPackageStream          uintptr
	Sign                       uintptr
	GetSignatureOriginPartName uintptr
	SetSignatureOriginPartName uintptr
	GetSignatures              uintptr
	AddSignatureBlock          uintptr
	GetSignatureBlocks         uintptr
	CreateSigningOptions       uintptr
	SavePackageToFile          uintptr
	SavePackageToStream        uintptr
}

type IXpsSignatureManager struct {
	IUnknown
}

func (this *IXpsSignatureManager) Vtbl() *IXpsSignatureManagerVtbl {
	return (*IXpsSignatureManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IXpsSignatureManager) LoadPackageFile(fileName PWSTR) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().LoadPackageFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileName)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) LoadPackageStream(stream *IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().LoadPackageStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) Sign(signOptions *IXpsSigningOptions, x509Certificate unsafe.Pointer, signature **IXpsSignature) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().Sign, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signOptions)), uintptr(x509Certificate), uintptr(unsafe.Pointer(signature)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) GetSignatureOriginPartName(signatureOriginPartName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureOriginPartName, uintptr(unsafe.Pointer(this)), uintptr(signatureOriginPartName))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) SetSignatureOriginPartName(signatureOriginPartName unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetSignatureOriginPartName, uintptr(unsafe.Pointer(this)), uintptr(signatureOriginPartName))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) GetSignatures(signatures **IXpsSignatureCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatures, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatures)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) AddSignatureBlock(partName unsafe.Pointer, fixedDocumentIndex uint32, signatureBlock **IXpsSignatureBlock) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddSignatureBlock, uintptr(unsafe.Pointer(this)), uintptr(partName), uintptr(fixedDocumentIndex), uintptr(unsafe.Pointer(signatureBlock)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) GetSignatureBlocks(signatureBlocks **IXpsSignatureBlockCollection) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSignatureBlocks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signatureBlocks)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) CreateSigningOptions(signingOptions **IXpsSigningOptions) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateSigningOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(signingOptions)))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) SavePackageToFile(fileName PWSTR, securityAttributes *SECURITY_ATTRIBUTES, flagsAndAttributes uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SavePackageToFile, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(securityAttributes)), uintptr(flagsAndAttributes))
	return HRESULT(ret)
}

func (this *IXpsSignatureManager) SavePackageToStream(stream *IStream) HRESULT {
	ret, _, _ := syscall.SyscallN(this.Vtbl().SavePackageToStream, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(stream)))
	return HRESULT(ret)
}

var (
	pEscape       uintptr
	pExtEscape    uintptr
	pStartDocA    uintptr
	pStartDocW    uintptr
	pEndDoc       uintptr
	pStartPage    uintptr
	pEndPage      uintptr
	pAbortDoc     uintptr
	pSetAbortProc uintptr
	pPrintWindow  uintptr
)

func Escape(hdc HDC, iEscape int32, cjIn int32, pvIn PSTR, pvOut unsafe.Pointer) int32 {
	addr := LazyAddr(&pEscape, libGdi32, "Escape")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(iEscape), uintptr(cjIn), uintptr(unsafe.Pointer(pvIn)), uintptr(pvOut))
	return int32(ret)
}

func ExtEscape(hdc HDC, iEscape int32, cjInput int32, lpInData PSTR, cjOutput int32, lpOutData PSTR) int32 {
	addr := LazyAddr(&pExtEscape, libGdi32, "ExtEscape")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(iEscape), uintptr(cjInput), uintptr(unsafe.Pointer(lpInData)), uintptr(cjOutput), uintptr(unsafe.Pointer(lpOutData)))
	return int32(ret)
}

func StartDocA(hdc HDC, lpdi *DOCINFOA) int32 {
	addr := LazyAddr(&pStartDocA, libGdi32, "StartDocA")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(unsafe.Pointer(lpdi)))
	return int32(ret)
}

var StartDoc = StartDocW

func StartDocW(hdc HDC, lpdi *DOCINFOW) int32 {
	addr := LazyAddr(&pStartDocW, libGdi32, "StartDocW")
	ret, _, _ := syscall.SyscallN(addr, hdc, uintptr(unsafe.Pointer(lpdi)))
	return int32(ret)
}

func EndDoc(hdc HDC) int32 {
	addr := LazyAddr(&pEndDoc, libGdi32, "EndDoc")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func StartPage(hdc HDC) int32 {
	addr := LazyAddr(&pStartPage, libGdi32, "StartPage")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func EndPage(hdc HDC) int32 {
	addr := LazyAddr(&pEndPage, libGdi32, "EndPage")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func AbortDoc(hdc HDC) int32 {
	addr := LazyAddr(&pAbortDoc, libGdi32, "AbortDoc")
	ret, _, _ := syscall.SyscallN(addr, hdc)
	return int32(ret)
}

func SetAbortProc(hdc HDC, proc ABORTPROC) int32 {
	addr := LazyAddr(&pSetAbortProc, libGdi32, "SetAbortProc")
	ret, _, _ := syscall.SyscallN(addr, hdc, proc)
	return int32(ret)
}

func PrintWindow(hwnd HWND, hdcBlt HDC, nFlags PRINT_WINDOW_FLAGS) BOOL {
	addr := LazyAddr(&pPrintWindow, libUser32, "PrintWindow")
	ret, _, _ := syscall.SyscallN(addr, hwnd, hdcBlt, uintptr(nFlags))
	return BOOL(ret)
}
