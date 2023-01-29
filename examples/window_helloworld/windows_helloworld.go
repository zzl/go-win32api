package main

import (
	"github.com/zzl/go-win32api/v2/win32"
	"runtime"
	"syscall"
)

/*
	https://github.com/microsoft/Windows-classic-samples
		/blob/main/Samples/Win7Samples/begin/LearnWin32/HelloWorld/cpp/main.cpp
*/

func main() {

	runtime.LockOSThread()
	hInstance, _ := win32.GetModuleHandle(nil)

	// Register the window class.
	CLASS_NAME := win32.StrToPwstr("Sample Window Class")

	var wc win32.WNDCLASS
	wc.LpfnWndProc = syscall.NewCallback(WindowProc)
	wc.HCursor, _ = win32.LoadCursor(win32.NULL, win32.IDC_ARROW)
	wc.LpszClassName = CLASS_NAME

	win32.RegisterClass(&wc)

	// Create the window.
	hwnd, _ := win32.CreateWindowEx(
		0,          // Optional window styles.
		CLASS_NAME, // Window class
		win32.StrToPwstr("Learn to Program Windows"), // Window text
		win32.WS_OVERLAPPEDWINDOW,                    // Window style

		// Size and position
		win32.CW_USEDEFAULT, win32.CW_USEDEFAULT, win32.CW_USEDEFAULT, win32.CW_USEDEFAULT,

		win32.NULL, // Parent window
		win32.NULL, // Menu
		hInstance,  // Instance handle
		nil,        // Additional application data
	)

	if hwnd == win32.NULL {
		return
	}

	win32.ShowWindow(hwnd, win32.SW_SHOWDEFAULT)

	// Run the message loop.
	var msg win32.MSG
	for {
		ok, _ := win32.GetMessage(&msg, win32.NULL, 0, 0)
		if ok == win32.FALSE {
			break
		}
		win32.TranslateMessage(&msg)
		win32.DispatchMessage(&msg)
	}
}

func WindowProc(hwnd win32.HWND, uMsg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT {
	switch uMsg {
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		return 0
	case win32.WM_PAINT:
		var ps win32.PAINTSTRUCT
		hdc := win32.BeginPaint(hwnd, &ps)

		// All painting occurs here, between BeginPaint and EndPaint.
		win32.FillRect(hdc, &ps.RcPaint, win32.HBRUSH(win32.COLOR_WINDOW+1))

		wszText, _ := syscall.UTF16FromString("Hello world!")
		win32.TextOut(hdc, 10, 10, &wszText[0], int32(len(wszText)-1))

		win32.EndPaint(hwnd, &ps)
		return 0
	}
	return win32.DefWindowProc(hwnd, uMsg, wParam, lParam)
}
