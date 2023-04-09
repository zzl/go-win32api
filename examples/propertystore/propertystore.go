package main

import (
	"github.com/zzl/go-win32api/v2/examples/propertystore/win32x"
	"github.com/zzl/go-win32api/v2/win32"
	"runtime"
	"unsafe"
)

func main() {

	runtime.LockOSThread()
	win32.CoInitialize(nil)
	defer win32.CoUninitialize()

	//
	var pps *win32.IPropertyStore
	hr := win32.SHGetPropertyStoreFromParsingName(
		win32.StrToPwstr(`D:\Test\arial.ttf`),
		nil, win32.GPS_DEFAULT,
		&win32.IID_IPropertyStore,
		unsafe.Pointer(&pps),
	)
	win32.ASSERT_SUCCEEDED(hr)
	defer pps.Release()

	//
	var pv win32.PROPVARIANT
	hr = pps.GetValue(&win32x.PKEY_Trademarks, &pv)
	win32.ASSERT_SUCCEEDED(hr)
	defer win32.PropVariantClear(&pv)

	//
	var pwszTrademarks win32.PWSTR
	hr = win32x.PSFormatForDisplayAlloc(&win32x.PKEY_Trademarks, &pv, win32.PDFF_DEFAULT, &pwszTrademarks)
	win32.ASSERT_SUCCEEDED(hr)
	defer win32.CoTaskMemFree(unsafe.Pointer(pwszTrademarks))

	//
	trademarks := win32.PwstrToStr(pwszTrademarks)
	println(trademarks)

}
