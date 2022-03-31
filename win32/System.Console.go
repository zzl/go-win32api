package win32

import "unsafe"
import "syscall"

type HPCON = uintptr

const (
	CONSOLE_TEXTMODE_BUFFER uint32 = 1
	ATTACH_PARENT_PROCESS uint32 = 4294967295
	CTRL_C_EVENT uint32 = 0
	CTRL_BREAK_EVENT uint32 = 1
	CTRL_CLOSE_EVENT uint32 = 2
	CTRL_LOGOFF_EVENT uint32 = 5
	CTRL_SHUTDOWN_EVENT uint32 = 6
	PSEUDOCONSOLE_INHERIT_CURSOR uint32 = 1
	FOREGROUND_BLUE uint32 = 1
	FOREGROUND_GREEN uint32 = 2
	FOREGROUND_RED uint32 = 4
	FOREGROUND_INTENSITY uint32 = 8
	BACKGROUND_BLUE uint32 = 16
	BACKGROUND_GREEN uint32 = 32
	BACKGROUND_RED uint32 = 64
	BACKGROUND_INTENSITY uint32 = 128
	COMMON_LVB_LEADING_BYTE uint32 = 256
	COMMON_LVB_TRAILING_BYTE uint32 = 512
	COMMON_LVB_GRID_HORIZONTAL uint32 = 1024
	COMMON_LVB_GRID_LVERTICAL uint32 = 2048
	COMMON_LVB_GRID_RVERTICAL uint32 = 4096
	COMMON_LVB_REVERSE_VIDEO uint32 = 16384
	COMMON_LVB_UNDERSCORE uint32 = 32768
	COMMON_LVB_SBCSDBCS uint32 = 768
	CONSOLE_NO_SELECTION uint32 = 0
	CONSOLE_SELECTION_IN_PROGRESS uint32 = 1
	CONSOLE_SELECTION_NOT_EMPTY uint32 = 2
	CONSOLE_MOUSE_SELECTION uint32 = 4
	CONSOLE_MOUSE_DOWN uint32 = 8
	HISTORY_NO_DUP_FLAG uint32 = 1
	CONSOLE_FULLSCREEN uint32 = 1
	CONSOLE_FULLSCREEN_HARDWARE uint32 = 2
	CONSOLE_FULLSCREEN_MODE uint32 = 1
	CONSOLE_WINDOWED_MODE uint32 = 2
	RIGHT_ALT_PRESSED uint32 = 1
	LEFT_ALT_PRESSED uint32 = 2
	RIGHT_CTRL_PRESSED uint32 = 4
	LEFT_CTRL_PRESSED uint32 = 8
	SHIFT_PRESSED uint32 = 16
	NUMLOCK_ON uint32 = 32
	SCROLLLOCK_ON uint32 = 64
	CAPSLOCK_ON uint32 = 128
	ENHANCED_KEY uint32 = 256
	NLS_DBCSCHAR uint32 = 65536
	NLS_ALPHANUMERIC uint32 = 0
	NLS_KATAKANA uint32 = 131072
	NLS_HIRAGANA uint32 = 262144
	NLS_ROMAN uint32 = 4194304
	NLS_IME_CONVERSION uint32 = 8388608
	ALTNUMPAD_BIT uint32 = 67108864
	NLS_IME_DISABLE uint32 = 536870912
	FROM_LEFT_1ST_BUTTON_PRESSED uint32 = 1
	RIGHTMOST_BUTTON_PRESSED uint32 = 2
	FROM_LEFT_2ND_BUTTON_PRESSED uint32 = 4
	FROM_LEFT_3RD_BUTTON_PRESSED uint32 = 8
	FROM_LEFT_4TH_BUTTON_PRESSED uint32 = 16
	MOUSE_MOVED uint32 = 1
	DOUBLE_CLICK uint32 = 2
	MOUSE_WHEELED uint32 = 4
	MOUSE_HWHEELED uint32 = 8
	KEY_EVENT uint32 = 1
	MOUSE_EVENT uint32 = 2
	WINDOW_BUFFER_SIZE_EVENT uint32 = 4
	MENU_EVENT uint32 = 8
	FOCUS_EVENT uint32 = 16
)

// enums

// enum CONSOLE_MODE
// flags
type CONSOLE_MODE uint32
const (
	ENABLE_PROCESSED_INPUT CONSOLE_MODE = 1
	ENABLE_LINE_INPUT CONSOLE_MODE = 2
	ENABLE_ECHO_INPUT CONSOLE_MODE = 4
	ENABLE_WINDOW_INPUT CONSOLE_MODE = 8
	ENABLE_MOUSE_INPUT CONSOLE_MODE = 16
	ENABLE_INSERT_MODE CONSOLE_MODE = 32
	ENABLE_QUICK_EDIT_MODE CONSOLE_MODE = 64
	ENABLE_EXTENDED_FLAGS CONSOLE_MODE = 128
	ENABLE_AUTO_POSITION CONSOLE_MODE = 256
	ENABLE_VIRTUAL_TERMINAL_INPUT CONSOLE_MODE = 512
	ENABLE_PROCESSED_OUTPUT CONSOLE_MODE = 1
	ENABLE_WRAP_AT_EOL_OUTPUT CONSOLE_MODE = 2
	ENABLE_VIRTUAL_TERMINAL_PROCESSING CONSOLE_MODE = 4
	DISABLE_NEWLINE_AUTO_RETURN CONSOLE_MODE = 8
	ENABLE_LVB_GRID_WORLDWIDE CONSOLE_MODE = 16
)

// enum STD_HANDLE
type STD_HANDLE uint32
const (
	STD_INPUT_HANDLE STD_HANDLE = 4294967286
	STD_OUTPUT_HANDLE STD_HANDLE = 4294967285
	STD_ERROR_HANDLE STD_HANDLE = 4294967284
)


// structs

type COORD struct {
	X int16
	Y int16
}

type SMALL_RECT struct {
	Left int16
	Top int16
	Right int16
	Bottom int16
}

type KEY_EVENT_RECORD_UChar_ struct {
	Data [1]byte
}

func (this *KEY_EVENT_RECORD_UChar_) UnicodeChar() *uint16{
	return (*uint16)(unsafe.Pointer(this))
}

func (this *KEY_EVENT_RECORD_UChar_) UnicodeCharVal() uint16{
	return *(*uint16)(unsafe.Pointer(this))
}

func (this *KEY_EVENT_RECORD_UChar_) AsciiChar() *CHAR{
	return (*CHAR)(unsafe.Pointer(this))
}

func (this *KEY_EVENT_RECORD_UChar_) AsciiCharVal() CHAR{
	return *(*CHAR)(unsafe.Pointer(this))
}

type KEY_EVENT_RECORD struct {
	BKeyDown BOOL
	WRepeatCount uint16
	WVirtualKeyCode uint16
	WVirtualScanCode uint16
	UChar KEY_EVENT_RECORD_UChar_
	DwControlKeyState uint32
}

type MOUSE_EVENT_RECORD struct {
	DwMousePosition COORD
	DwButtonState uint32
	DwControlKeyState uint32
	DwEventFlags uint32
}

type WINDOW_BUFFER_SIZE_RECORD struct {
	DwSize COORD
}

type MENU_EVENT_RECORD struct {
	DwCommandId uint32
}

type FOCUS_EVENT_RECORD struct {
	BSetFocus BOOL
}

type INPUT_RECORD_Event_ struct {
	Data [4]uint32
}

func (this *INPUT_RECORD_Event_) KeyEvent() *KEY_EVENT_RECORD{
	return (*KEY_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) KeyEventVal() KEY_EVENT_RECORD{
	return *(*KEY_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) MouseEvent() *MOUSE_EVENT_RECORD{
	return (*MOUSE_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) MouseEventVal() MOUSE_EVENT_RECORD{
	return *(*MOUSE_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) WindowBufferSizeEvent() *WINDOW_BUFFER_SIZE_RECORD{
	return (*WINDOW_BUFFER_SIZE_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) WindowBufferSizeEventVal() WINDOW_BUFFER_SIZE_RECORD{
	return *(*WINDOW_BUFFER_SIZE_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) MenuEvent() *MENU_EVENT_RECORD{
	return (*MENU_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) MenuEventVal() MENU_EVENT_RECORD{
	return *(*MENU_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) FocusEvent() *FOCUS_EVENT_RECORD{
	return (*FOCUS_EVENT_RECORD)(unsafe.Pointer(this))
}

func (this *INPUT_RECORD_Event_) FocusEventVal() FOCUS_EVENT_RECORD{
	return *(*FOCUS_EVENT_RECORD)(unsafe.Pointer(this))
}

type INPUT_RECORD struct {
	EventType uint16
	Event INPUT_RECORD_Event_
}

type CHAR_INFO_Char_ struct {
	Data [1]byte
}

func (this *CHAR_INFO_Char_) UnicodeChar() *uint16{
	return (*uint16)(unsafe.Pointer(this))
}

func (this *CHAR_INFO_Char_) UnicodeCharVal() uint16{
	return *(*uint16)(unsafe.Pointer(this))
}

func (this *CHAR_INFO_Char_) AsciiChar() *CHAR{
	return (*CHAR)(unsafe.Pointer(this))
}

func (this *CHAR_INFO_Char_) AsciiCharVal() CHAR{
	return *(*CHAR)(unsafe.Pointer(this))
}

type CHAR_INFO struct {
	Char CHAR_INFO_Char_
	Attributes uint16
}

type CONSOLE_FONT_INFO struct {
	NFont uint32
	DwFontSize COORD
}

type CONSOLE_READCONSOLE_CONTROL struct {
	NLength uint32
	NInitialChars uint32
	DwCtrlWakeupMask uint32
	DwControlKeyState uint32
}

type CONSOLE_CURSOR_INFO struct {
	DwSize uint32
	BVisible BOOL
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize COORD
	DwCursorPosition COORD
	WAttributes uint16
	SrWindow SMALL_RECT
	DwMaximumWindowSize COORD
}

type CONSOLE_SCREEN_BUFFER_INFOEX struct {
	CbSize uint32
	DwSize COORD
	DwCursorPosition COORD
	WAttributes uint16
	SrWindow SMALL_RECT
	DwMaximumWindowSize COORD
	WPopupAttributes uint16
	BFullscreenSupported BOOL
	ColorTable [16]uint32
}

type CONSOLE_FONT_INFOEX struct {
	CbSize uint32
	NFont uint32
	DwFontSize COORD
	FontFamily uint32
	FontWeight uint32
	FaceName [32]uint16
}

type CONSOLE_SELECTION_INFO struct {
	DwFlags uint32
	DwSelectionAnchor COORD
	SrSelection SMALL_RECT
}

type CONSOLE_HISTORY_INFO struct {
	CbSize uint32
	HistoryBufferSize uint32
	NumberOfHistoryBuffers uint32
	DwFlags uint32
}


// func types

type PHANDLER_ROUTINE func(CtrlType uint32) BOOL


var (
	pAllocConsole uintptr
	pFreeConsole uintptr
	pAttachConsole uintptr
	pGetConsoleCP uintptr
	pGetConsoleOutputCP uintptr
	pGetConsoleMode uintptr
	pSetConsoleMode uintptr
	pGetNumberOfConsoleInputEvents uintptr
	pReadConsoleInputA uintptr
	pReadConsoleInputW uintptr
	pPeekConsoleInputA uintptr
	pPeekConsoleInputW uintptr
	pReadConsoleA uintptr
	pReadConsoleW uintptr
	pWriteConsoleA uintptr
	pWriteConsoleW uintptr
	pSetConsoleCtrlHandler uintptr
	pCreatePseudoConsole uintptr
	pResizePseudoConsole uintptr
	pClosePseudoConsole uintptr
	pFillConsoleOutputCharacterA uintptr
	pFillConsoleOutputCharacterW uintptr
	pFillConsoleOutputAttribute uintptr
	pGenerateConsoleCtrlEvent uintptr
	pCreateConsoleScreenBuffer uintptr
	pSetConsoleActiveScreenBuffer uintptr
	pFlushConsoleInputBuffer uintptr
	pSetConsoleCP uintptr
	pSetConsoleOutputCP uintptr
	pGetConsoleCursorInfo uintptr
	pSetConsoleCursorInfo uintptr
	pGetConsoleScreenBufferInfo uintptr
	pGetConsoleScreenBufferInfoEx uintptr
	pSetConsoleScreenBufferInfoEx uintptr
	pSetConsoleScreenBufferSize uintptr
	pSetConsoleCursorPosition uintptr
	pGetLargestConsoleWindowSize uintptr
	pSetConsoleTextAttribute uintptr
	pSetConsoleWindowInfo uintptr
	pWriteConsoleOutputCharacterA uintptr
	pWriteConsoleOutputCharacterW uintptr
	pWriteConsoleOutputAttribute uintptr
	pReadConsoleOutputCharacterA uintptr
	pReadConsoleOutputCharacterW uintptr
	pReadConsoleOutputAttribute uintptr
	pWriteConsoleInputA uintptr
	pWriteConsoleInputW uintptr
	pScrollConsoleScreenBufferA uintptr
	pScrollConsoleScreenBufferW uintptr
	pWriteConsoleOutputA uintptr
	pWriteConsoleOutputW uintptr
	pReadConsoleOutputA uintptr
	pReadConsoleOutputW uintptr
	pGetConsoleTitleA uintptr
	pGetConsoleTitleW uintptr
	pGetConsoleOriginalTitleA uintptr
	pGetConsoleOriginalTitleW uintptr
	pSetConsoleTitleA uintptr
	pSetConsoleTitleW uintptr
	pGetNumberOfConsoleMouseButtons uintptr
	pGetConsoleFontSize uintptr
	pGetCurrentConsoleFont uintptr
	pGetCurrentConsoleFontEx uintptr
	pSetCurrentConsoleFontEx uintptr
	pGetConsoleSelectionInfo uintptr
	pGetConsoleHistoryInfo uintptr
	pSetConsoleHistoryInfo uintptr
	pGetConsoleDisplayMode uintptr
	pSetConsoleDisplayMode uintptr
	pGetConsoleWindow uintptr
	pAddConsoleAliasA uintptr
	pAddConsoleAliasW uintptr
	pGetConsoleAliasA uintptr
	pGetConsoleAliasW uintptr
	pGetConsoleAliasesLengthA uintptr
	pGetConsoleAliasesLengthW uintptr
	pGetConsoleAliasExesLengthA uintptr
	pGetConsoleAliasExesLengthW uintptr
	pGetConsoleAliasesA uintptr
	pGetConsoleAliasesW uintptr
	pGetConsoleAliasExesA uintptr
	pGetConsoleAliasExesW uintptr
	pExpungeConsoleCommandHistoryA uintptr
	pExpungeConsoleCommandHistoryW uintptr
	pSetConsoleNumberOfCommandsA uintptr
	pSetConsoleNumberOfCommandsW uintptr
	pGetConsoleCommandHistoryLengthA uintptr
	pGetConsoleCommandHistoryLengthW uintptr
	pGetConsoleCommandHistoryA uintptr
	pGetConsoleCommandHistoryW uintptr
	pGetConsoleProcessList uintptr
	pGetStdHandle uintptr
	pSetStdHandle uintptr
	pSetStdHandleEx uintptr
)

func AllocConsole() BOOL {
	addr := lazyAddr(&pAllocConsole, libKernel32, "AllocConsole")
	ret, _,  _ := syscall.SyscallN(addr)
	return BOOL(ret)
}

func FreeConsole() BOOL {
	addr := lazyAddr(&pFreeConsole, libKernel32, "FreeConsole")
	ret, _,  _ := syscall.SyscallN(addr)
	return BOOL(ret)
}

func AttachConsole(dwProcessId uint32) BOOL {
	addr := lazyAddr(&pAttachConsole, libKernel32, "AttachConsole")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(dwProcessId))
	return BOOL(ret)
}

func GetConsoleCP() uint32 {
	addr := lazyAddr(&pGetConsoleCP, libKernel32, "GetConsoleCP")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func GetConsoleOutputCP() uint32 {
	addr := lazyAddr(&pGetConsoleOutputCP, libKernel32, "GetConsoleOutputCP")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func GetConsoleMode(hConsoleHandle HANDLE, lpMode *CONSOLE_MODE) BOOL {
	addr := lazyAddr(&pGetConsoleMode, libKernel32, "GetConsoleMode")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleHandle, uintptr(unsafe.Pointer(lpMode)))
	return BOOL(ret)
}

func SetConsoleMode(hConsoleHandle HANDLE, dwMode CONSOLE_MODE) BOOL {
	addr := lazyAddr(&pSetConsoleMode, libKernel32, "SetConsoleMode")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleHandle, uintptr(dwMode))
	return BOOL(ret)
}

func GetNumberOfConsoleInputEvents(hConsoleInput HANDLE, lpNumberOfEvents *uint32) BOOL {
	addr := lazyAddr(&pGetNumberOfConsoleInputEvents, libKernel32, "GetNumberOfConsoleInputEvents")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpNumberOfEvents)))
	return BOOL(ret)
}

func ReadConsoleInputA(hConsoleInput HANDLE, lpBuffer *INPUT_RECORD, nLength uint32, lpNumberOfEventsRead *uint32) BOOL {
	addr := lazyAddr(&pReadConsoleInputA, libKernel32, "ReadConsoleInputA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpBuffer)), uintptr(nLength), uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return BOOL(ret)
}

var ReadConsoleInput = ReadConsoleInputW
func ReadConsoleInputW(hConsoleInput HANDLE, lpBuffer *INPUT_RECORD, nLength uint32, lpNumberOfEventsRead *uint32) BOOL {
	addr := lazyAddr(&pReadConsoleInputW, libKernel32, "ReadConsoleInputW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpBuffer)), uintptr(nLength), uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return BOOL(ret)
}

func PeekConsoleInputA(hConsoleInput HANDLE, lpBuffer *INPUT_RECORD, nLength uint32, lpNumberOfEventsRead *uint32) BOOL {
	addr := lazyAddr(&pPeekConsoleInputA, libKernel32, "PeekConsoleInputA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpBuffer)), uintptr(nLength), uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return BOOL(ret)
}

var PeekConsoleInput = PeekConsoleInputW
func PeekConsoleInputW(hConsoleInput HANDLE, lpBuffer *INPUT_RECORD, nLength uint32, lpNumberOfEventsRead *uint32) BOOL {
	addr := lazyAddr(&pPeekConsoleInputW, libKernel32, "PeekConsoleInputW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpBuffer)), uintptr(nLength), uintptr(unsafe.Pointer(lpNumberOfEventsRead)))
	return BOOL(ret)
}

func ReadConsoleA(hConsoleInput HANDLE, lpBuffer unsafe.Pointer, nNumberOfCharsToRead uint32, lpNumberOfCharsRead *uint32, pInputControl *CONSOLE_READCONSOLE_CONTROL) BOOL {
	addr := lazyAddr(&pReadConsoleA, libKernel32, "ReadConsoleA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(lpBuffer), uintptr(nNumberOfCharsToRead), uintptr(unsafe.Pointer(lpNumberOfCharsRead)), uintptr(unsafe.Pointer(pInputControl)))
	return BOOL(ret)
}

var ReadConsole = ReadConsoleW
func ReadConsoleW(hConsoleInput HANDLE, lpBuffer unsafe.Pointer, nNumberOfCharsToRead uint32, lpNumberOfCharsRead *uint32, pInputControl *CONSOLE_READCONSOLE_CONTROL) BOOL {
	addr := lazyAddr(&pReadConsoleW, libKernel32, "ReadConsoleW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(lpBuffer), uintptr(nNumberOfCharsToRead), uintptr(unsafe.Pointer(lpNumberOfCharsRead)), uintptr(unsafe.Pointer(pInputControl)))
	return BOOL(ret)
}

func WriteConsoleA(hConsoleOutput HANDLE, lpBuffer unsafe.Pointer, nNumberOfCharsToWrite uint32, lpNumberOfCharsWritten *uint32, lpReserved unsafe.Pointer) BOOL {
	addr := lazyAddr(&pWriteConsoleA, libKernel32, "WriteConsoleA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(lpBuffer), uintptr(nNumberOfCharsToWrite), uintptr(unsafe.Pointer(lpNumberOfCharsWritten)), uintptr(lpReserved))
	return BOOL(ret)
}

var WriteConsole = WriteConsoleW
func WriteConsoleW(hConsoleOutput HANDLE, lpBuffer unsafe.Pointer, nNumberOfCharsToWrite uint32, lpNumberOfCharsWritten *uint32, lpReserved unsafe.Pointer) BOOL {
	addr := lazyAddr(&pWriteConsoleW, libKernel32, "WriteConsoleW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(lpBuffer), uintptr(nNumberOfCharsToWrite), uintptr(unsafe.Pointer(lpNumberOfCharsWritten)), uintptr(lpReserved))
	return BOOL(ret)
}

func SetConsoleCtrlHandler(HandlerRoutine uintptr, Add BOOL) BOOL {
	addr := lazyAddr(&pSetConsoleCtrlHandler, libKernel32, "SetConsoleCtrlHandler")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(HandlerRoutine), uintptr(Add))
	return BOOL(ret)
}

func CreatePseudoConsole(size COORD, hInput HANDLE, hOutput HANDLE, dwFlags uint32, phPC *HPCON) HRESULT {
	addr := lazyAddr(&pCreatePseudoConsole, libKernel32, "CreatePseudoConsole")
	ret, _,  _ := syscall.SyscallN(addr, *(*uintptr)(unsafe.Pointer(&size)), hInput, hOutput, uintptr(dwFlags), uintptr(unsafe.Pointer(phPC)))
	return HRESULT(ret)
}

func ResizePseudoConsole(hPC HPCON, size COORD) HRESULT {
	addr := lazyAddr(&pResizePseudoConsole, libKernel32, "ResizePseudoConsole")
	ret, _,  _ := syscall.SyscallN(addr, hPC, *(*uintptr)(unsafe.Pointer(&size)))
	return HRESULT(ret)
}

func ClosePseudoConsole(hPC HPCON) {
	addr := lazyAddr(&pClosePseudoConsole, libKernel32, "ClosePseudoConsole")
	_, _,  _ = syscall.SyscallN(addr, hPC)
}

func FillConsoleOutputCharacterA(hConsoleOutput HANDLE, cCharacter CHAR, nLength uint32, dwWriteCoord COORD, lpNumberOfCharsWritten *uint32) BOOL {
	addr := lazyAddr(&pFillConsoleOutputCharacterA, libKernel32, "FillConsoleOutputCharacterA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(cCharacter), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwWriteCoord)), uintptr(unsafe.Pointer(lpNumberOfCharsWritten)))
	return BOOL(ret)
}

var FillConsoleOutputCharacter = FillConsoleOutputCharacterW
func FillConsoleOutputCharacterW(hConsoleOutput HANDLE, cCharacter uint16, nLength uint32, dwWriteCoord COORD, lpNumberOfCharsWritten *uint32) BOOL {
	addr := lazyAddr(&pFillConsoleOutputCharacterW, libKernel32, "FillConsoleOutputCharacterW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(cCharacter), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwWriteCoord)), uintptr(unsafe.Pointer(lpNumberOfCharsWritten)))
	return BOOL(ret)
}

func FillConsoleOutputAttribute(hConsoleOutput HANDLE, wAttribute uint16, nLength uint32, dwWriteCoord COORD, lpNumberOfAttrsWritten *uint32) BOOL {
	addr := lazyAddr(&pFillConsoleOutputAttribute, libKernel32, "FillConsoleOutputAttribute")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(wAttribute), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwWriteCoord)), uintptr(unsafe.Pointer(lpNumberOfAttrsWritten)))
	return BOOL(ret)
}

func GenerateConsoleCtrlEvent(dwCtrlEvent uint32, dwProcessGroupId uint32) BOOL {
	addr := lazyAddr(&pGenerateConsoleCtrlEvent, libKernel32, "GenerateConsoleCtrlEvent")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(dwCtrlEvent), uintptr(dwProcessGroupId))
	return BOOL(ret)
}

func CreateConsoleScreenBuffer(dwDesiredAccess uint32, dwShareMode uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwFlags uint32, lpScreenBufferData unsafe.Pointer) HANDLE {
	addr := lazyAddr(&pCreateConsoleScreenBuffer, libKernel32, "CreateConsoleScreenBuffer")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(dwDesiredAccess), uintptr(dwShareMode), uintptr(unsafe.Pointer(lpSecurityAttributes)), uintptr(dwFlags), uintptr(lpScreenBufferData))
	return HANDLE(ret)
}

func SetConsoleActiveScreenBuffer(hConsoleOutput HANDLE) BOOL {
	addr := lazyAddr(&pSetConsoleActiveScreenBuffer, libKernel32, "SetConsoleActiveScreenBuffer")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput)
	return BOOL(ret)
}

func FlushConsoleInputBuffer(hConsoleInput HANDLE) BOOL {
	addr := lazyAddr(&pFlushConsoleInputBuffer, libKernel32, "FlushConsoleInputBuffer")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput)
	return BOOL(ret)
}

func SetConsoleCP(wCodePageID uint32) BOOL {
	addr := lazyAddr(&pSetConsoleCP, libKernel32, "SetConsoleCP")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(wCodePageID))
	return BOOL(ret)
}

func SetConsoleOutputCP(wCodePageID uint32) BOOL {
	addr := lazyAddr(&pSetConsoleOutputCP, libKernel32, "SetConsoleOutputCP")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(wCodePageID))
	return BOOL(ret)
}

func GetConsoleCursorInfo(hConsoleOutput HANDLE, lpConsoleCursorInfo *CONSOLE_CURSOR_INFO) BOOL {
	addr := lazyAddr(&pGetConsoleCursorInfo, libKernel32, "GetConsoleCursorInfo")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpConsoleCursorInfo)))
	return BOOL(ret)
}

func SetConsoleCursorInfo(hConsoleOutput HANDLE, lpConsoleCursorInfo *CONSOLE_CURSOR_INFO) BOOL {
	addr := lazyAddr(&pSetConsoleCursorInfo, libKernel32, "SetConsoleCursorInfo")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpConsoleCursorInfo)))
	return BOOL(ret)
}

func GetConsoleScreenBufferInfo(hConsoleOutput HANDLE, lpConsoleScreenBufferInfo *CONSOLE_SCREEN_BUFFER_INFO) BOOL {
	addr := lazyAddr(&pGetConsoleScreenBufferInfo, libKernel32, "GetConsoleScreenBufferInfo")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpConsoleScreenBufferInfo)))
	return BOOL(ret)
}

func GetConsoleScreenBufferInfoEx(hConsoleOutput HANDLE, lpConsoleScreenBufferInfoEx *CONSOLE_SCREEN_BUFFER_INFOEX) BOOL {
	addr := lazyAddr(&pGetConsoleScreenBufferInfoEx, libKernel32, "GetConsoleScreenBufferInfoEx")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpConsoleScreenBufferInfoEx)))
	return BOOL(ret)
}

func SetConsoleScreenBufferInfoEx(hConsoleOutput HANDLE, lpConsoleScreenBufferInfoEx *CONSOLE_SCREEN_BUFFER_INFOEX) BOOL {
	addr := lazyAddr(&pSetConsoleScreenBufferInfoEx, libKernel32, "SetConsoleScreenBufferInfoEx")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpConsoleScreenBufferInfoEx)))
	return BOOL(ret)
}

func SetConsoleScreenBufferSize(hConsoleOutput HANDLE, dwSize COORD) BOOL {
	addr := lazyAddr(&pSetConsoleScreenBufferSize, libKernel32, "SetConsoleScreenBufferSize")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, *(*uintptr)(unsafe.Pointer(&dwSize)))
	return BOOL(ret)
}

func SetConsoleCursorPosition(hConsoleOutput HANDLE, dwCursorPosition COORD) BOOL {
	addr := lazyAddr(&pSetConsoleCursorPosition, libKernel32, "SetConsoleCursorPosition")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, *(*uintptr)(unsafe.Pointer(&dwCursorPosition)))
	return BOOL(ret)
}

func GetLargestConsoleWindowSize(hConsoleOutput HANDLE) COORD {
	addr := lazyAddr(&pGetLargestConsoleWindowSize, libKernel32, "GetLargestConsoleWindowSize")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput)
	return *(*COORD)(unsafe.Pointer(ret))
}

func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes uint16) BOOL {
	addr := lazyAddr(&pSetConsoleTextAttribute, libKernel32, "SetConsoleTextAttribute")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(wAttributes))
	return BOOL(ret)
}

func SetConsoleWindowInfo(hConsoleOutput HANDLE, bAbsolute BOOL, lpConsoleWindow *SMALL_RECT) BOOL {
	addr := lazyAddr(&pSetConsoleWindowInfo, libKernel32, "SetConsoleWindowInfo")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(bAbsolute), uintptr(unsafe.Pointer(lpConsoleWindow)))
	return BOOL(ret)
}

func WriteConsoleOutputCharacterA(hConsoleOutput HANDLE, lpCharacter *uint8, nLength uint32, dwWriteCoord COORD, lpNumberOfCharsWritten *uint32) BOOL {
	addr := lazyAddr(&pWriteConsoleOutputCharacterA, libKernel32, "WriteConsoleOutputCharacterA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpCharacter)), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwWriteCoord)), uintptr(unsafe.Pointer(lpNumberOfCharsWritten)))
	return BOOL(ret)
}

var WriteConsoleOutputCharacter = WriteConsoleOutputCharacterW
func WriteConsoleOutputCharacterW(hConsoleOutput HANDLE, lpCharacter *uint16, nLength uint32, dwWriteCoord COORD, lpNumberOfCharsWritten *uint32) BOOL {
	addr := lazyAddr(&pWriteConsoleOutputCharacterW, libKernel32, "WriteConsoleOutputCharacterW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpCharacter)), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwWriteCoord)), uintptr(unsafe.Pointer(lpNumberOfCharsWritten)))
	return BOOL(ret)
}

func WriteConsoleOutputAttribute(hConsoleOutput HANDLE, lpAttribute *uint16, nLength uint32, dwWriteCoord COORD, lpNumberOfAttrsWritten *uint32) BOOL {
	addr := lazyAddr(&pWriteConsoleOutputAttribute, libKernel32, "WriteConsoleOutputAttribute")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpAttribute)), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwWriteCoord)), uintptr(unsafe.Pointer(lpNumberOfAttrsWritten)))
	return BOOL(ret)
}

func ReadConsoleOutputCharacterA(hConsoleOutput HANDLE, lpCharacter *uint8, nLength uint32, dwReadCoord COORD, lpNumberOfCharsRead *uint32) BOOL {
	addr := lazyAddr(&pReadConsoleOutputCharacterA, libKernel32, "ReadConsoleOutputCharacterA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpCharacter)), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwReadCoord)), uintptr(unsafe.Pointer(lpNumberOfCharsRead)))
	return BOOL(ret)
}

var ReadConsoleOutputCharacter = ReadConsoleOutputCharacterW
func ReadConsoleOutputCharacterW(hConsoleOutput HANDLE, lpCharacter *uint16, nLength uint32, dwReadCoord COORD, lpNumberOfCharsRead *uint32) BOOL {
	addr := lazyAddr(&pReadConsoleOutputCharacterW, libKernel32, "ReadConsoleOutputCharacterW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpCharacter)), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwReadCoord)), uintptr(unsafe.Pointer(lpNumberOfCharsRead)))
	return BOOL(ret)
}

func ReadConsoleOutputAttribute(hConsoleOutput HANDLE, lpAttribute *uint16, nLength uint32, dwReadCoord COORD, lpNumberOfAttrsRead *uint32) BOOL {
	addr := lazyAddr(&pReadConsoleOutputAttribute, libKernel32, "ReadConsoleOutputAttribute")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpAttribute)), uintptr(nLength), *(*uintptr)(unsafe.Pointer(&dwReadCoord)), uintptr(unsafe.Pointer(lpNumberOfAttrsRead)))
	return BOOL(ret)
}

func WriteConsoleInputA(hConsoleInput HANDLE, lpBuffer *INPUT_RECORD, nLength uint32, lpNumberOfEventsWritten *uint32) BOOL {
	addr := lazyAddr(&pWriteConsoleInputA, libKernel32, "WriteConsoleInputA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpBuffer)), uintptr(nLength), uintptr(unsafe.Pointer(lpNumberOfEventsWritten)))
	return BOOL(ret)
}

var WriteConsoleInput = WriteConsoleInputW
func WriteConsoleInputW(hConsoleInput HANDLE, lpBuffer *INPUT_RECORD, nLength uint32, lpNumberOfEventsWritten *uint32) BOOL {
	addr := lazyAddr(&pWriteConsoleInputW, libKernel32, "WriteConsoleInputW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleInput, uintptr(unsafe.Pointer(lpBuffer)), uintptr(nLength), uintptr(unsafe.Pointer(lpNumberOfEventsWritten)))
	return BOOL(ret)
}

func ScrollConsoleScreenBufferA(hConsoleOutput HANDLE, lpScrollRectangle *SMALL_RECT, lpClipRectangle *SMALL_RECT, dwDestinationOrigin COORD, lpFill *CHAR_INFO) BOOL {
	addr := lazyAddr(&pScrollConsoleScreenBufferA, libKernel32, "ScrollConsoleScreenBufferA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpScrollRectangle)), uintptr(unsafe.Pointer(lpClipRectangle)), *(*uintptr)(unsafe.Pointer(&dwDestinationOrigin)), uintptr(unsafe.Pointer(lpFill)))
	return BOOL(ret)
}

var ScrollConsoleScreenBuffer = ScrollConsoleScreenBufferW
func ScrollConsoleScreenBufferW(hConsoleOutput HANDLE, lpScrollRectangle *SMALL_RECT, lpClipRectangle *SMALL_RECT, dwDestinationOrigin COORD, lpFill *CHAR_INFO) BOOL {
	addr := lazyAddr(&pScrollConsoleScreenBufferW, libKernel32, "ScrollConsoleScreenBufferW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpScrollRectangle)), uintptr(unsafe.Pointer(lpClipRectangle)), *(*uintptr)(unsafe.Pointer(&dwDestinationOrigin)), uintptr(unsafe.Pointer(lpFill)))
	return BOOL(ret)
}

func WriteConsoleOutputA(hConsoleOutput HANDLE, lpBuffer *CHAR_INFO, dwBufferSize COORD, dwBufferCoord COORD, lpWriteRegion *SMALL_RECT) BOOL {
	addr := lazyAddr(&pWriteConsoleOutputA, libKernel32, "WriteConsoleOutputA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpBuffer)), *(*uintptr)(unsafe.Pointer(&dwBufferSize)), *(*uintptr)(unsafe.Pointer(&dwBufferCoord)), uintptr(unsafe.Pointer(lpWriteRegion)))
	return BOOL(ret)
}

var WriteConsoleOutput = WriteConsoleOutputW
func WriteConsoleOutputW(hConsoleOutput HANDLE, lpBuffer *CHAR_INFO, dwBufferSize COORD, dwBufferCoord COORD, lpWriteRegion *SMALL_RECT) BOOL {
	addr := lazyAddr(&pWriteConsoleOutputW, libKernel32, "WriteConsoleOutputW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpBuffer)), *(*uintptr)(unsafe.Pointer(&dwBufferSize)), *(*uintptr)(unsafe.Pointer(&dwBufferCoord)), uintptr(unsafe.Pointer(lpWriteRegion)))
	return BOOL(ret)
}

func ReadConsoleOutputA(hConsoleOutput HANDLE, lpBuffer *CHAR_INFO, dwBufferSize COORD, dwBufferCoord COORD, lpReadRegion *SMALL_RECT) BOOL {
	addr := lazyAddr(&pReadConsoleOutputA, libKernel32, "ReadConsoleOutputA")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpBuffer)), *(*uintptr)(unsafe.Pointer(&dwBufferSize)), *(*uintptr)(unsafe.Pointer(&dwBufferCoord)), uintptr(unsafe.Pointer(lpReadRegion)))
	return BOOL(ret)
}

var ReadConsoleOutput = ReadConsoleOutputW
func ReadConsoleOutputW(hConsoleOutput HANDLE, lpBuffer *CHAR_INFO, dwBufferSize COORD, dwBufferCoord COORD, lpReadRegion *SMALL_RECT) BOOL {
	addr := lazyAddr(&pReadConsoleOutputW, libKernel32, "ReadConsoleOutputW")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(unsafe.Pointer(lpBuffer)), *(*uintptr)(unsafe.Pointer(&dwBufferSize)), *(*uintptr)(unsafe.Pointer(&dwBufferCoord)), uintptr(unsafe.Pointer(lpReadRegion)))
	return BOOL(ret)
}

func GetConsoleTitleA(lpConsoleTitle *uint8, nSize uint32) uint32 {
	addr := lazyAddr(&pGetConsoleTitleA, libKernel32, "GetConsoleTitleA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleTitle)), uintptr(nSize))
	return uint32(ret)
}

var GetConsoleTitle = GetConsoleTitleW
func GetConsoleTitleW(lpConsoleTitle *uint16, nSize uint32) uint32 {
	addr := lazyAddr(&pGetConsoleTitleW, libKernel32, "GetConsoleTitleW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleTitle)), uintptr(nSize))
	return uint32(ret)
}

func GetConsoleOriginalTitleA(lpConsoleTitle *uint8, nSize uint32) uint32 {
	addr := lazyAddr(&pGetConsoleOriginalTitleA, libKernel32, "GetConsoleOriginalTitleA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleTitle)), uintptr(nSize))
	return uint32(ret)
}

var GetConsoleOriginalTitle = GetConsoleOriginalTitleW
func GetConsoleOriginalTitleW(lpConsoleTitle *uint16, nSize uint32) uint32 {
	addr := lazyAddr(&pGetConsoleOriginalTitleW, libKernel32, "GetConsoleOriginalTitleW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleTitle)), uintptr(nSize))
	return uint32(ret)
}

func SetConsoleTitleA(lpConsoleTitle PSTR) BOOL {
	addr := lazyAddr(&pSetConsoleTitleA, libKernel32, "SetConsoleTitleA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleTitle)))
	return BOOL(ret)
}

var SetConsoleTitle = SetConsoleTitleW
func SetConsoleTitleW(lpConsoleTitle PWSTR) BOOL {
	addr := lazyAddr(&pSetConsoleTitleW, libKernel32, "SetConsoleTitleW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleTitle)))
	return BOOL(ret)
}

func GetNumberOfConsoleMouseButtons(lpNumberOfMouseButtons *uint32) BOOL {
	addr := lazyAddr(&pGetNumberOfConsoleMouseButtons, libKernel32, "GetNumberOfConsoleMouseButtons")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpNumberOfMouseButtons)))
	return BOOL(ret)
}

func GetConsoleFontSize(hConsoleOutput HANDLE, nFont uint32) COORD {
	addr := lazyAddr(&pGetConsoleFontSize, libKernel32, "GetConsoleFontSize")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(nFont))
	return *(*COORD)(unsafe.Pointer(ret))
}

func GetCurrentConsoleFont(hConsoleOutput HANDLE, bMaximumWindow BOOL, lpConsoleCurrentFont *CONSOLE_FONT_INFO) BOOL {
	addr := lazyAddr(&pGetCurrentConsoleFont, libKernel32, "GetCurrentConsoleFont")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(bMaximumWindow), uintptr(unsafe.Pointer(lpConsoleCurrentFont)))
	return BOOL(ret)
}

func GetCurrentConsoleFontEx(hConsoleOutput HANDLE, bMaximumWindow BOOL, lpConsoleCurrentFontEx *CONSOLE_FONT_INFOEX) BOOL {
	addr := lazyAddr(&pGetCurrentConsoleFontEx, libKernel32, "GetCurrentConsoleFontEx")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(bMaximumWindow), uintptr(unsafe.Pointer(lpConsoleCurrentFontEx)))
	return BOOL(ret)
}

func SetCurrentConsoleFontEx(hConsoleOutput HANDLE, bMaximumWindow BOOL, lpConsoleCurrentFontEx *CONSOLE_FONT_INFOEX) BOOL {
	addr := lazyAddr(&pSetCurrentConsoleFontEx, libKernel32, "SetCurrentConsoleFontEx")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(bMaximumWindow), uintptr(unsafe.Pointer(lpConsoleCurrentFontEx)))
	return BOOL(ret)
}

func GetConsoleSelectionInfo(lpConsoleSelectionInfo *CONSOLE_SELECTION_INFO) BOOL {
	addr := lazyAddr(&pGetConsoleSelectionInfo, libKernel32, "GetConsoleSelectionInfo")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleSelectionInfo)))
	return BOOL(ret)
}

func GetConsoleHistoryInfo(lpConsoleHistoryInfo *CONSOLE_HISTORY_INFO) BOOL {
	addr := lazyAddr(&pGetConsoleHistoryInfo, libKernel32, "GetConsoleHistoryInfo")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleHistoryInfo)))
	return BOOL(ret)
}

func SetConsoleHistoryInfo(lpConsoleHistoryInfo *CONSOLE_HISTORY_INFO) BOOL {
	addr := lazyAddr(&pSetConsoleHistoryInfo, libKernel32, "SetConsoleHistoryInfo")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpConsoleHistoryInfo)))
	return BOOL(ret)
}

func GetConsoleDisplayMode(lpModeFlags *uint32) BOOL {
	addr := lazyAddr(&pGetConsoleDisplayMode, libKernel32, "GetConsoleDisplayMode")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpModeFlags)))
	return BOOL(ret)
}

func SetConsoleDisplayMode(hConsoleOutput HANDLE, dwFlags uint32, lpNewScreenBufferDimensions *COORD) BOOL {
	addr := lazyAddr(&pSetConsoleDisplayMode, libKernel32, "SetConsoleDisplayMode")
	ret, _,  _ := syscall.SyscallN(addr, hConsoleOutput, uintptr(dwFlags), uintptr(unsafe.Pointer(lpNewScreenBufferDimensions)))
	return BOOL(ret)
}

func GetConsoleWindow() HWND {
	addr := lazyAddr(&pGetConsoleWindow, libKernel32, "GetConsoleWindow")
	ret, _,  _ := syscall.SyscallN(addr)
	return HWND(ret)
}

func AddConsoleAliasA(Source PSTR, Target PSTR, ExeName PSTR) BOOL {
	addr := lazyAddr(&pAddConsoleAliasA, libKernel32, "AddConsoleAliasA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Source)), uintptr(unsafe.Pointer(Target)), uintptr(unsafe.Pointer(ExeName)))
	return BOOL(ret)
}

var AddConsoleAlias = AddConsoleAliasW
func AddConsoleAliasW(Source PWSTR, Target PWSTR, ExeName PWSTR) BOOL {
	addr := lazyAddr(&pAddConsoleAliasW, libKernel32, "AddConsoleAliasW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Source)), uintptr(unsafe.Pointer(Target)), uintptr(unsafe.Pointer(ExeName)))
	return BOOL(ret)
}

func GetConsoleAliasA(Source PSTR, TargetBuffer *uint8, TargetBufferLength uint32, ExeName PSTR) uint32 {
	addr := lazyAddr(&pGetConsoleAliasA, libKernel32, "GetConsoleAliasA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Source)), uintptr(unsafe.Pointer(TargetBuffer)), uintptr(TargetBufferLength), uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

var GetConsoleAlias = GetConsoleAliasW
func GetConsoleAliasW(Source PWSTR, TargetBuffer *uint16, TargetBufferLength uint32, ExeName PWSTR) uint32 {
	addr := lazyAddr(&pGetConsoleAliasW, libKernel32, "GetConsoleAliasW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Source)), uintptr(unsafe.Pointer(TargetBuffer)), uintptr(TargetBufferLength), uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

func GetConsoleAliasesLengthA(ExeName PSTR) uint32 {
	addr := lazyAddr(&pGetConsoleAliasesLengthA, libKernel32, "GetConsoleAliasesLengthA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

var GetConsoleAliasesLength = GetConsoleAliasesLengthW
func GetConsoleAliasesLengthW(ExeName PWSTR) uint32 {
	addr := lazyAddr(&pGetConsoleAliasesLengthW, libKernel32, "GetConsoleAliasesLengthW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

func GetConsoleAliasExesLengthA() uint32 {
	addr := lazyAddr(&pGetConsoleAliasExesLengthA, libKernel32, "GetConsoleAliasExesLengthA")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint32(ret)
}

var GetConsoleAliasExesLength = GetConsoleAliasExesLengthW
func GetConsoleAliasExesLengthW() uint32 {
	addr := lazyAddr(&pGetConsoleAliasExesLengthW, libKernel32, "GetConsoleAliasExesLengthW")
	ret, _,  _ := syscall.SyscallN(addr)
	return uint32(ret)
}

func GetConsoleAliasesA(AliasBuffer *uint8, AliasBufferLength uint32, ExeName PSTR) uint32 {
	addr := lazyAddr(&pGetConsoleAliasesA, libKernel32, "GetConsoleAliasesA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(AliasBuffer)), uintptr(AliasBufferLength), uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

var GetConsoleAliases = GetConsoleAliasesW
func GetConsoleAliasesW(AliasBuffer *uint16, AliasBufferLength uint32, ExeName PWSTR) uint32 {
	addr := lazyAddr(&pGetConsoleAliasesW, libKernel32, "GetConsoleAliasesW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(AliasBuffer)), uintptr(AliasBufferLength), uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

func GetConsoleAliasExesA(ExeNameBuffer *uint8, ExeNameBufferLength uint32) uint32 {
	addr := lazyAddr(&pGetConsoleAliasExesA, libKernel32, "GetConsoleAliasExesA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeNameBuffer)), uintptr(ExeNameBufferLength))
	return uint32(ret)
}

var GetConsoleAliasExes = GetConsoleAliasExesW
func GetConsoleAliasExesW(ExeNameBuffer *uint16, ExeNameBufferLength uint32) uint32 {
	addr := lazyAddr(&pGetConsoleAliasExesW, libKernel32, "GetConsoleAliasExesW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeNameBuffer)), uintptr(ExeNameBufferLength))
	return uint32(ret)
}

func ExpungeConsoleCommandHistoryA(ExeName PSTR) {
	addr := lazyAddr(&pExpungeConsoleCommandHistoryA, libKernel32, "ExpungeConsoleCommandHistoryA")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
}

var ExpungeConsoleCommandHistory = ExpungeConsoleCommandHistoryW
func ExpungeConsoleCommandHistoryW(ExeName PWSTR) {
	addr := lazyAddr(&pExpungeConsoleCommandHistoryW, libKernel32, "ExpungeConsoleCommandHistoryW")
	_, _,  _ = syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
}

func SetConsoleNumberOfCommandsA(Number uint32, ExeName PSTR) BOOL {
	addr := lazyAddr(&pSetConsoleNumberOfCommandsA, libKernel32, "SetConsoleNumberOfCommandsA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(Number), uintptr(unsafe.Pointer(ExeName)))
	return BOOL(ret)
}

var SetConsoleNumberOfCommands = SetConsoleNumberOfCommandsW
func SetConsoleNumberOfCommandsW(Number uint32, ExeName PWSTR) BOOL {
	addr := lazyAddr(&pSetConsoleNumberOfCommandsW, libKernel32, "SetConsoleNumberOfCommandsW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(Number), uintptr(unsafe.Pointer(ExeName)))
	return BOOL(ret)
}

func GetConsoleCommandHistoryLengthA(ExeName PSTR) uint32 {
	addr := lazyAddr(&pGetConsoleCommandHistoryLengthA, libKernel32, "GetConsoleCommandHistoryLengthA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

var GetConsoleCommandHistoryLength = GetConsoleCommandHistoryLengthW
func GetConsoleCommandHistoryLengthW(ExeName PWSTR) uint32 {
	addr := lazyAddr(&pGetConsoleCommandHistoryLengthW, libKernel32, "GetConsoleCommandHistoryLengthW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

func GetConsoleCommandHistoryA(Commands PSTR, CommandBufferLength uint32, ExeName PSTR) uint32 {
	addr := lazyAddr(&pGetConsoleCommandHistoryA, libKernel32, "GetConsoleCommandHistoryA")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Commands)), uintptr(CommandBufferLength), uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

var GetConsoleCommandHistory = GetConsoleCommandHistoryW
func GetConsoleCommandHistoryW(Commands PWSTR, CommandBufferLength uint32, ExeName PWSTR) uint32 {
	addr := lazyAddr(&pGetConsoleCommandHistoryW, libKernel32, "GetConsoleCommandHistoryW")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(Commands)), uintptr(CommandBufferLength), uintptr(unsafe.Pointer(ExeName)))
	return uint32(ret)
}

func GetConsoleProcessList(lpdwProcessList *uint32, dwProcessCount uint32) uint32 {
	addr := lazyAddr(&pGetConsoleProcessList, libKernel32, "GetConsoleProcessList")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(lpdwProcessList)), uintptr(dwProcessCount))
	return uint32(ret)
}

func GetStdHandle(nStdHandle STD_HANDLE) HANDLE {
	addr := lazyAddr(&pGetStdHandle, libKernel32, "GetStdHandle")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(nStdHandle))
	return HANDLE(ret)
}

func SetStdHandle(nStdHandle STD_HANDLE, hHandle HANDLE) BOOL {
	addr := lazyAddr(&pSetStdHandle, libKernel32, "SetStdHandle")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(nStdHandle), hHandle)
	return BOOL(ret)
}

func SetStdHandleEx(nStdHandle STD_HANDLE, hHandle HANDLE, phPrevValue *HANDLE) BOOL {
	addr := lazyAddr(&pSetStdHandleEx, libKernel32, "SetStdHandleEx")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(nStdHandle), hHandle, uintptr(unsafe.Pointer(phPrevValue)))
	return BOOL(ret)
}


