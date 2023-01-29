package main

import (
	"github.com/zzl/go-win32api/v2/win32"
	"os"
	"unsafe"
)

//https://learn.microsoft.com/en-us/windows/console/using-the-high-level-input-and-output-functions

var hStdout, hStdin win32.HANDLE
var csbiInfo win32.CONSOLE_SCREEN_BUFFER_INFO

func main() {

	lpszPrompt1 := win32.StrToPstr("Type a line and press Enter, or q to quit: ")
	lpszPrompt2 := win32.StrToPstr("Type any key, or q to quit: ")
	var chBuffer [256]win32.CHAR
	var cRead, cWritten win32.DWORD
	var fdwMode, fdwOldMode win32.CONSOLE_MODE
	var wOldColorAttrs win32.CONSOLE_CHARACTER_ATTRIBUTES

	// Get handles to STDIN and STDOUT.

	hStdin, _ = win32.GetStdHandle(win32.STD_INPUT_HANDLE)
	hStdout, _ = win32.GetStdHandle(win32.STD_OUTPUT_HANDLE)

	if hStdin == win32.INVALID_HANDLE_VALUE ||
		hStdout == win32.INVALID_HANDLE_VALUE {
		win32.MessageBox(win32.NULL, win32.StrToPwstr("GetStdHandle"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		os.Exit(1)
	}

	// Save the current text colors.

	if ok, _ := win32.GetConsoleScreenBufferInfo(hStdout, &csbiInfo); ok == win32.FALSE {
		win32.MessageBox(win32.NULL, win32.StrToPwstr("GetConsoleScreenBufferInfo"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		os.Exit(1)
	}

	wOldColorAttrs = csbiInfo.WAttributes

	// Set the text attributes to draw red text on black background.

	if ok, _ := win32.SetConsoleTextAttribute(hStdout,
		win32.FOREGROUND_RED|win32.FOREGROUND_INTENSITY); ok == win32.FALSE {
		win32.MessageBox(win32.NULL, win32.StrToPwstr("GetConsoleScreenBufferInfo"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		os.Exit(1)
	}

	// Write to STDOUT and read from STDIN by using the default
	// modes. Input is echoed automatically, and ReadFile
	// does not return until a carriage return is typed.
	//
	// The default input modes are line, processed, and echo.
	// The default output modes are processed and wrap at EOL.

	for {
		if ok, _ := win32.WriteFile(
			hStdout,                             // output handle
			unsafe.Pointer(lpszPrompt1),         // prompt string
			uint32(win32.LstrlenA(lpszPrompt1)), // string length
			&cWritten,                           // bytes written
			nil,                                 // not overlapped
		); ok == win32.FALSE {
			os.Exit(1)
		}

		if ok, _ := win32.ReadFile(
			hStdin,                       // input handle
			unsafe.Pointer(&chBuffer[0]), // buffer to read into
			255,                          // size of buffer
			&cRead,                       // actual bytes read
			nil,                          // not overlapped
		); ok == win32.FALSE {
			break
		}

		if chBuffer[0] == 'q' {
			break
		}
	}

	// Turn off the line input and echo input modes
	if ok, _ := win32.GetConsoleMode(hStdin, &fdwOldMode); ok == win32.FALSE {
		win32.MessageBox(win32.NULL, win32.StrToPwstr("GetConsoleMode"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		os.Exit(1)
	}

	fdwMode = fdwOldMode & ^(win32.ENABLE_LINE_INPUT | win32.ENABLE_ECHO_INPUT)
	if ok, _ := win32.SetConsoleMode(hStdin, fdwMode); ok == win32.FALSE {
		win32.MessageBox(win32.NULL, win32.StrToPwstr("SetConsoleMode"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		os.Exit(1)
	}

	// ReadFile returns when any input is available.
	// WriteFile is used to echo input.

	NewLine()

	for {
		if ok, _ := win32.WriteFile(
			hStdout,                             // output handle
			unsafe.Pointer(lpszPrompt2),         // prompt string
			uint32(win32.LstrlenA(lpszPrompt2)), // string length
			&cWritten,                           // bytes written
			nil,                                 // not overlapped
		); ok == win32.FALSE {
			win32.MessageBox(win32.NULL, win32.StrToPwstr("WriteFile"),
				win32.StrToPwstr("Console Error"), win32.MB_OK)
			os.Exit(1)
		}

		if ok, _ := win32.ReadFile(hStdin,
			unsafe.Pointer(&chBuffer[0]), 1, &cRead, nil); ok == win32.FALSE {
			break
		}
		if chBuffer[0] == '\r' {
			NewLine()
		} else if ok, _ := win32.WriteFile(hStdout,
			unsafe.Pointer(&chBuffer[0]), cRead, &cWritten, nil); ok == win32.FALSE {
			break
		} else {
			NewLine()
		}
		if chBuffer[0] == 'q' {
			break
		}
	}

	// Restore the original console mode.

	win32.SetConsoleMode(hStdin, fdwOldMode)

	// Restore the original text colors.

	win32.SetConsoleTextAttribute(hStdout, wOldColorAttrs)

}

// The NewLine function handles carriage returns when the processed
// input mode is disabled. It gets the current cursor position
// and resets it to the first cell of the next row.
func NewLine() {
	if ok, _ := win32.GetConsoleScreenBufferInfo(hStdout, &csbiInfo); ok == win32.FALSE {
		win32.MessageBox(win32.NULL, win32.StrToPwstr("GetConsoleScreenBufferInfo"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		return
	}
	csbiInfo.DwCursorPosition.X = 0

	// If it is the last line in the screen buffer, scroll
	// the buffer up.
	if csbiInfo.DwSize.Y-1 == csbiInfo.DwCursorPosition.Y {
		ScrollScreenBuffer(hStdout, 1)
	} else { // Otherwise, advance the cursor to the next line.
		csbiInfo.DwCursorPosition.Y += 1
	}

	if ok, _ := win32.SetConsoleCursorPosition(hStdout,
		csbiInfo.DwCursorPosition); ok == win32.FALSE {

		win32.MessageBox(win32.NULL, win32.StrToPwstr("SetConsoleCursorPosition"),
			win32.StrToPwstr("Console Error"), win32.MB_OK)
		return
	}
}

func ScrollScreenBuffer(h win32.HANDLE, x win32.INT) {
	var srctScrollRect, srctClipRect win32.SMALL_RECT
	var chiFill win32.CHAR_INFO
	var coordDest win32.COORD

	srctScrollRect.Left = 0
	srctScrollRect.Top = 1
	srctScrollRect.Right = csbiInfo.DwSize.X - win32.SHORT(x)
	srctScrollRect.Bottom = csbiInfo.DwSize.Y - win32.SHORT(x)

	// The destination for the scroll rectangle is one row up.

	coordDest.X = 0
	coordDest.Y = 0

	// The clipping rectangle is the same as the scrolling rectangle.
	// The destination row is left unchanged.

	srctClipRect = srctScrollRect

	// Set the fill character and attributes.

	chiFill.Attributes = uint16(win32.FOREGROUND_RED | win32.FOREGROUND_INTENSITY)
	*chiFill.Char.AsciiChar() = ' '

	// Scroll up one line.

	win32.ScrollConsoleScreenBuffer(
		h,               // screen buffer handle
		&srctScrollRect, // scrolling rectangle
		&srctClipRect,   // clipping rectangle
		coordDest,       // top left destination cell
		&chiFill)        // fill character and color
}
