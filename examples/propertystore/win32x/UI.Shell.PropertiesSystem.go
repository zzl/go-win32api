package win32x

import (
	. "github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var libPropsys = windows.NewLazySystemDLL("Propsys.dll")
var (
	pPropVariantToWinRTPropertyValue             uintptr
	pWinRTPropertyValueToPropVariant             uintptr
	pPSFormatForDisplay                          uintptr
	pPSFormatForDisplayAlloc                     uintptr
	pPSFormatPropertyValue                       uintptr
	pPSGetImageReferenceForValue                 uintptr
	pPSStringFromPropertyKey                     uintptr
	pPSPropertyKeyFromString                     uintptr
	pPSCreateMemoryPropertyStore                 uintptr
	pPSCreateDelayedMultiplexPropertyStore       uintptr
	pPSCreateMultiplexPropertyStore              uintptr
	pPSCreatePropertyChangeArray                 uintptr
	pPSCreateSimplePropertyChange                uintptr
	pPSGetPropertyDescription                    uintptr
	pPSGetPropertyDescriptionByName              uintptr
	pPSLookupPropertyHandlerCLSID                uintptr
	pPSGetItemPropertyHandler                    uintptr
	pPSGetItemPropertyHandlerWithCreateObject    uintptr
	pPSGetPropertyValue                          uintptr
	pPSSetPropertyValue                          uintptr
	pPSRegisterPropertySchema                    uintptr
	pPSUnregisterPropertySchema                  uintptr
	pPSRefreshPropertySchema                     uintptr
	pPSEnumeratePropertyDescriptions             uintptr
	pPSGetPropertyKeyFromName                    uintptr
	pPSGetNameFromPropertyKey                    uintptr
	pPSCoerceToCanonicalValue                    uintptr
	pPSGetPropertyDescriptionListFromString      uintptr
	pPSCreatePropertyStoreFromPropertySetStorage uintptr
	pPSCreatePropertyStoreFromObject             uintptr
	pPSCreateAdapterFromPropertyStore            uintptr
	pPSGetPropertySystem                         uintptr
	pPSGetPropertyFromPropertyStorage            uintptr
	pPSGetNamedPropertyFromPropertyStorage       uintptr
	pPSPropertyBag_ReadType                      uintptr
	pPSPropertyBag_ReadStr                       uintptr
	pPSPropertyBag_ReadStrAlloc                  uintptr
	pPSPropertyBag_ReadBSTR                      uintptr
	pPSPropertyBag_WriteStr                      uintptr
	pPSPropertyBag_WriteBSTR                     uintptr
	pPSPropertyBag_ReadInt                       uintptr
	pPSPropertyBag_WriteInt                      uintptr
	pPSPropertyBag_ReadSHORT                     uintptr
	pPSPropertyBag_WriteSHORT                    uintptr
	pPSPropertyBag_ReadLONG                      uintptr
	pPSPropertyBag_WriteLONG                     uintptr
	pPSPropertyBag_ReadDWORD                     uintptr
	pPSPropertyBag_WriteDWORD                    uintptr
	pPSPropertyBag_ReadBOOL                      uintptr
	pPSPropertyBag_WriteBOOL                     uintptr
	pPSPropertyBag_ReadPOINTL                    uintptr
	pPSPropertyBag_WritePOINTL                   uintptr
	pPSPropertyBag_ReadPOINTS                    uintptr
	pPSPropertyBag_WritePOINTS                   uintptr
	pPSPropertyBag_ReadRECTL                     uintptr
	pPSPropertyBag_WriteRECTL                    uintptr
	pPSPropertyBag_ReadStream                    uintptr
	pPSPropertyBag_WriteStream                   uintptr
	pPSPropertyBag_Delete                        uintptr
	pPSPropertyBag_ReadULONGLONG                 uintptr
	pPSPropertyBag_WriteULONGLONG                uintptr
	pPSPropertyBag_ReadUnknown                   uintptr
	pPSPropertyBag_WriteUnknown                  uintptr
	pPSPropertyBag_ReadGUID                      uintptr
	pPSPropertyBag_WriteGUID                     uintptr
	pPSPropertyBag_ReadPropertyKey               uintptr
	pPSPropertyBag_WritePropertyKey              uintptr
	pInitPropVariantFromResource                 uintptr
	pInitPropVariantFromBuffer                   uintptr
	pInitPropVariantFromCLSID                    uintptr
	pInitPropVariantFromGUIDAsString             uintptr
	pInitPropVariantFromFileTime                 uintptr
	pInitPropVariantFromPropVariantVectorElem    uintptr
	pInitPropVariantVectorFromPropVariant        uintptr
	pInitPropVariantFromStrRet                   uintptr
	pInitPropVariantFromBooleanVector            uintptr
	pInitPropVariantFromInt16Vector              uintptr
	pInitPropVariantFromUInt16Vector             uintptr
	pInitPropVariantFromInt32Vector              uintptr
	pInitPropVariantFromUInt32Vector             uintptr
	pInitPropVariantFromInt64Vector              uintptr
	pInitPropVariantFromUInt64Vector             uintptr
	pInitPropVariantFromDoubleVector             uintptr
	pInitPropVariantFromFileTimeVector           uintptr
	pInitPropVariantFromStringVector             uintptr
	pInitPropVariantFromStringAsVector           uintptr
	pPropVariantToBooleanWithDefault             uintptr
	pPropVariantToInt16WithDefault               uintptr
	pPropVariantToUInt16WithDefault              uintptr
	pPropVariantToInt32WithDefault               uintptr
	pPropVariantToUInt32WithDefault              uintptr
	pPropVariantToInt64WithDefault               uintptr
	pPropVariantToUInt64WithDefault              uintptr
	pPropVariantToDoubleWithDefault              uintptr
	pPropVariantToStringWithDefault              uintptr
	pPropVariantToBoolean                        uintptr
	pPropVariantToInt16                          uintptr
	pPropVariantToUInt16                         uintptr
	pPropVariantToInt32                          uintptr
	pPropVariantToUInt32                         uintptr
	pPropVariantToInt64                          uintptr
	pPropVariantToUInt64                         uintptr
	pPropVariantToDouble                         uintptr
	pPropVariantToBuffer                         uintptr
	pPropVariantToString                         uintptr
	pPropVariantToGUID                           uintptr
	pPropVariantToStringAlloc                    uintptr
	pPropVariantToBSTR                           uintptr
	pPropVariantToStrRet                         uintptr
	pPropVariantToFileTime                       uintptr
	pPropVariantGetElementCount                  uintptr
	pPropVariantToBooleanVector                  uintptr
	pPropVariantToInt16Vector                    uintptr
	pPropVariantToUInt16Vector                   uintptr
	pPropVariantToInt32Vector                    uintptr
	pPropVariantToUInt32Vector                   uintptr
	pPropVariantToInt64Vector                    uintptr
	pPropVariantToUInt64Vector                   uintptr
	pPropVariantToDoubleVector                   uintptr
	pPropVariantToFileTimeVector                 uintptr
	pPropVariantToStringVector                   uintptr
	pPropVariantToBooleanVectorAlloc             uintptr
	pPropVariantToInt16VectorAlloc               uintptr
	pPropVariantToUInt16VectorAlloc              uintptr
	pPropVariantToInt32VectorAlloc               uintptr
	pPropVariantToUInt32VectorAlloc              uintptr
	pPropVariantToInt64VectorAlloc               uintptr
	pPropVariantToUInt64VectorAlloc              uintptr
	pPropVariantToDoubleVectorAlloc              uintptr
	pPropVariantToFileTimeVectorAlloc            uintptr
	pPropVariantToStringVectorAlloc              uintptr
	pPropVariantGetBooleanElem                   uintptr
	pPropVariantGetInt16Elem                     uintptr
	pPropVariantGetUInt16Elem                    uintptr
	pPropVariantGetInt32Elem                     uintptr
	pPropVariantGetUInt32Elem                    uintptr
	pPropVariantGetInt64Elem                     uintptr
	pPropVariantGetUInt64Elem                    uintptr
	pPropVariantGetDoubleElem                    uintptr
	pPropVariantGetFileTimeElem                  uintptr
	pPropVariantGetStringElem                    uintptr
	pClearPropVariantArray                       uintptr
	pPropVariantCompareEx                        uintptr
	pPropVariantChangeType                       uintptr
	pPropVariantToVariant                        uintptr
	pVariantToPropVariant                        uintptr
	pInitVariantFromResource                     uintptr
	pInitVariantFromBuffer                       uintptr
	pInitVariantFromGUIDAsString                 uintptr
	pInitVariantFromFileTime                     uintptr
	pInitVariantFromFileTimeArray                uintptr
	pInitVariantFromStrRet                       uintptr
	pInitVariantFromVariantArrayElem             uintptr
	pInitVariantFromBooleanArray                 uintptr
	pInitVariantFromInt16Array                   uintptr
	pInitVariantFromUInt16Array                  uintptr
	pInitVariantFromInt32Array                   uintptr
	pInitVariantFromUInt32Array                  uintptr
	pInitVariantFromInt64Array                   uintptr
	pInitVariantFromUInt64Array                  uintptr
	pInitVariantFromDoubleArray                  uintptr
	pInitVariantFromStringArray                  uintptr
	pVariantToBooleanWithDefault                 uintptr
	pVariantToInt16WithDefault                   uintptr
	pVariantToUInt16WithDefault                  uintptr
	pVariantToInt32WithDefault                   uintptr
	pVariantToUInt32WithDefault                  uintptr
	pVariantToInt64WithDefault                   uintptr
	pVariantToUInt64WithDefault                  uintptr
	pVariantToDoubleWithDefault                  uintptr
	pVariantToStringWithDefault                  uintptr
	pVariantToBoolean                            uintptr
	pVariantToInt16                              uintptr
	pVariantToUInt16                             uintptr
	pVariantToInt32                              uintptr
	pVariantToUInt32                             uintptr
	pVariantToInt64                              uintptr
	pVariantToUInt64                             uintptr
	pVariantToDouble                             uintptr
	pVariantToBuffer                             uintptr
	pVariantToGUID                               uintptr
	pVariantToString                             uintptr
	pVariantToStringAlloc                        uintptr
	pVariantToDosDateTime                        uintptr
	pVariantToStrRet                             uintptr
	pVariantToFileTime                           uintptr
	pVariantGetElementCount                      uintptr
	pVariantToBooleanArray                       uintptr
	pVariantToInt16Array                         uintptr
	pVariantToUInt16Array                        uintptr
	pVariantToInt32Array                         uintptr
	pVariantToUInt32Array                        uintptr
	pVariantToInt64Array                         uintptr
	pVariantToUInt64Array                        uintptr
	pVariantToDoubleArray                        uintptr
	pVariantToStringArray                        uintptr
	pVariantToBooleanArrayAlloc                  uintptr
	pVariantToInt16ArrayAlloc                    uintptr
	pVariantToUInt16ArrayAlloc                   uintptr
	pVariantToInt32ArrayAlloc                    uintptr
	pVariantToUInt32ArrayAlloc                   uintptr
	pVariantToInt64ArrayAlloc                    uintptr
	pVariantToUInt64ArrayAlloc                   uintptr
	pVariantToDoubleArrayAlloc                   uintptr
	pVariantToStringArrayAlloc                   uintptr
	pVariantGetBooleanElem                       uintptr
	pVariantGetInt16Elem                         uintptr
	pVariantGetUInt16Elem                        uintptr
	pVariantGetInt32Elem                         uintptr
	pVariantGetUInt32Elem                        uintptr
	pVariantGetInt64Elem                         uintptr
	pVariantGetUInt64Elem                        uintptr
	pVariantGetDoubleElem                        uintptr
	pVariantGetStringElem                        uintptr
	pClearVariantArray                           uintptr
	pVariantCompare                              uintptr
)

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

func PSFormatForDisplay(propkey *PROPERTYKEY, propvar *PROPVARIANT, pdfFlags PROPDESC_FORMAT_FLAGS, pwszText PWSTR, cchText uint32) HRESULT {
	addr := LazyAddr(&pPSFormatForDisplay, libPropsys, "PSFormatForDisplay")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propkey)), uintptr(unsafe.Pointer(propvar)), uintptr(pdfFlags), uintptr(unsafe.Pointer(pwszText)), uintptr(cchText))
	return HRESULT(ret)
}

func PSFormatForDisplayAlloc(key *PROPERTYKEY, propvar *PROPVARIANT, pdff PROPDESC_FORMAT_FLAGS, ppszDisplay *PWSTR) HRESULT {
	addr := LazyAddr(&pPSFormatForDisplayAlloc, libPropsys, "PSFormatForDisplayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(propvar)), uintptr(pdff), uintptr(unsafe.Pointer(ppszDisplay)))
	return HRESULT(ret)
}

func PSFormatPropertyValue(pps *IPropertyStore, ppd *IPropertyDescription, pdff PROPDESC_FORMAT_FLAGS, ppszDisplay *PWSTR) HRESULT {
	addr := LazyAddr(&pPSFormatPropertyValue, libPropsys, "PSFormatPropertyValue")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pps)), uintptr(unsafe.Pointer(ppd)), uintptr(pdff), uintptr(unsafe.Pointer(ppszDisplay)))
	return HRESULT(ret)
}

func PSGetImageReferenceForValue(propkey *PROPERTYKEY, propvar *PROPVARIANT, ppszImageRes *PWSTR) HRESULT {
	addr := LazyAddr(&pPSGetImageReferenceForValue, libPropsys, "PSGetImageReferenceForValue")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propkey)), uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(ppszImageRes)))
	return HRESULT(ret)
}

func PSStringFromPropertyKey(pkey *PROPERTYKEY, psz PWSTR, cch uint32) HRESULT {
	addr := LazyAddr(&pPSStringFromPropertyKey, libPropsys, "PSStringFromPropertyKey")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pkey)), uintptr(unsafe.Pointer(psz)), uintptr(cch))
	return HRESULT(ret)
}

func PSPropertyKeyFromString(pszString PWSTR, pkey *PROPERTYKEY) HRESULT {
	addr := LazyAddr(&pPSPropertyKeyFromString, libPropsys, "PSPropertyKeyFromString")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszString)), uintptr(unsafe.Pointer(pkey)))
	return HRESULT(ret)
}

func PSCreateMemoryPropertyStore(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreateMemoryPropertyStore, libPropsys, "PSCreateMemoryPropertyStore")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreateDelayedMultiplexPropertyStore(flags GETPROPERTYSTOREFLAGS, pdpsf *IDelayedPropertyStoreFactory, rgStoreIds *uint32, cStores uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreateDelayedMultiplexPropertyStore, libPropsys, "PSCreateDelayedMultiplexPropertyStore")
	ret, _, _ := syscall.SyscallN(addr, uintptr(flags), uintptr(unsafe.Pointer(pdpsf)), uintptr(unsafe.Pointer(rgStoreIds)), uintptr(cStores), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreateMultiplexPropertyStore(prgpunkStores **IUnknown, cStores uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreateMultiplexPropertyStore, libPropsys, "PSCreateMultiplexPropertyStore")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgpunkStores)), uintptr(cStores), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreatePropertyChangeArray(rgpropkey *PROPERTYKEY, rgflags *PKA_FLAGS, rgpropvar *PROPVARIANT, cChanges uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreatePropertyChangeArray, libPropsys, "PSCreatePropertyChangeArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(rgpropkey)), uintptr(unsafe.Pointer(rgflags)), uintptr(unsafe.Pointer(rgpropvar)), uintptr(cChanges), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreateSimplePropertyChange(flags PKA_FLAGS, key *PROPERTYKEY, propvar *PROPVARIANT, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreateSimplePropertyChange, libPropsys, "PSCreateSimplePropertyChange")
	ret, _, _ := syscall.SyscallN(addr, uintptr(flags), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetPropertyDescription(propkey *PROPERTYKEY, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSGetPropertyDescription, libPropsys, "PSGetPropertyDescription")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propkey)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetPropertyDescriptionByName(pszCanonicalName PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSGetPropertyDescriptionByName, libPropsys, "PSGetPropertyDescriptionByName")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszCanonicalName)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSLookupPropertyHandlerCLSID(pszFilePath PWSTR, pclsid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pPSLookupPropertyHandlerCLSID, libPropsys, "PSLookupPropertyHandlerCLSID")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszFilePath)), uintptr(unsafe.Pointer(pclsid)))
	return HRESULT(ret)
}

func PSGetItemPropertyHandler(punkItem *IUnknown, fReadWrite BOOL, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSGetItemPropertyHandler, libPropsys, "PSGetItemPropertyHandler")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(punkItem)), uintptr(fReadWrite), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetItemPropertyHandlerWithCreateObject(punkItem *IUnknown, fReadWrite BOOL, punkCreateObject *IUnknown, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSGetItemPropertyHandlerWithCreateObject, libPropsys, "PSGetItemPropertyHandlerWithCreateObject")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(punkItem)), uintptr(fReadWrite), uintptr(unsafe.Pointer(punkCreateObject)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetPropertyValue(pps *IPropertyStore, ppd *IPropertyDescription, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPSGetPropertyValue, libPropsys, "PSGetPropertyValue")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pps)), uintptr(unsafe.Pointer(ppd)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func PSSetPropertyValue(pps *IPropertyStore, ppd *IPropertyDescription, propvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPSSetPropertyValue, libPropsys, "PSSetPropertyValue")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pps)), uintptr(unsafe.Pointer(ppd)), uintptr(unsafe.Pointer(propvar)))
	return HRESULT(ret)
}

func PSRegisterPropertySchema(pszPath PWSTR) HRESULT {
	addr := LazyAddr(&pPSRegisterPropertySchema, libPropsys, "PSRegisterPropertySchema")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszPath)))
	return HRESULT(ret)
}

func PSUnregisterPropertySchema(pszPath PWSTR) HRESULT {
	addr := LazyAddr(&pPSUnregisterPropertySchema, libPropsys, "PSUnregisterPropertySchema")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszPath)))
	return HRESULT(ret)
}

func PSRefreshPropertySchema() HRESULT {
	addr := LazyAddr(&pPSRefreshPropertySchema, libPropsys, "PSRefreshPropertySchema")
	ret, _, _ := syscall.SyscallN(addr)
	return HRESULT(ret)
}

func PSEnumeratePropertyDescriptions(filterOn PROPDESC_ENUMFILTER, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSEnumeratePropertyDescriptions, libPropsys, "PSEnumeratePropertyDescriptions")
	ret, _, _ := syscall.SyscallN(addr, uintptr(filterOn), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetPropertyKeyFromName(pszName PWSTR, ppropkey *PROPERTYKEY) HRESULT {
	addr := LazyAddr(&pPSGetPropertyKeyFromName, libPropsys, "PSGetPropertyKeyFromName")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszName)), uintptr(unsafe.Pointer(ppropkey)))
	return HRESULT(ret)
}

func PSGetNameFromPropertyKey(propkey *PROPERTYKEY, ppszCanonicalName *PWSTR) HRESULT {
	addr := LazyAddr(&pPSGetNameFromPropertyKey, libPropsys, "PSGetNameFromPropertyKey")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propkey)), uintptr(unsafe.Pointer(ppszCanonicalName)))
	return HRESULT(ret)
}

func PSCoerceToCanonicalValue(key *PROPERTYKEY, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPSCoerceToCanonicalValue, libPropsys, "PSCoerceToCanonicalValue")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func PSGetPropertyDescriptionListFromString(pszPropList PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSGetPropertyDescriptionListFromString, libPropsys, "PSGetPropertyDescriptionListFromString")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pszPropList)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreatePropertyStoreFromPropertySetStorage(ppss *IPropertySetStorage, grfMode uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreatePropertyStoreFromPropertySetStorage, libPropsys, "PSCreatePropertyStoreFromPropertySetStorage")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(ppss)), uintptr(grfMode), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreatePropertyStoreFromObject(punk *IUnknown, grfMode uint32, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreatePropertyStoreFromObject, libPropsys, "PSCreatePropertyStoreFromObject")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(punk)), uintptr(grfMode), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSCreateAdapterFromPropertyStore(pps *IPropertyStore, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSCreateAdapterFromPropertyStore, libPropsys, "PSCreateAdapterFromPropertyStore")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pps)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetPropertySystem(riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSGetPropertySystem, libPropsys, "PSGetPropertySystem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSGetPropertyFromPropertyStorage(psps *SERIALIZEDPROPSTORAGE, cb uint32, rpkey *PROPERTYKEY, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPSGetPropertyFromPropertyStorage, libPropsys, "PSGetPropertyFromPropertyStorage")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(psps)), uintptr(cb), uintptr(unsafe.Pointer(rpkey)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func PSGetNamedPropertyFromPropertyStorage(psps *SERIALIZEDPROPSTORAGE, cb uint32, pszName PWSTR, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pPSGetNamedPropertyFromPropertyStorage, libPropsys, "PSGetNamedPropertyFromPropertyStorage")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(psps)), uintptr(cb), uintptr(unsafe.Pointer(pszName)), uintptr(unsafe.Pointer(ppropvar)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadType(propBag *IPropertyBag, propName PWSTR, var_ *VARIANT, type_ VARENUM) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadType, libPropsys, "PSPropertyBag_ReadType")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(var_)), uintptr(type_))
	return HRESULT(ret)
}

func PSPropertyBag_ReadStr(propBag *IPropertyBag, propName PWSTR, value PWSTR, characterCount int32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadStr, libPropsys, "PSPropertyBag_ReadStr")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)), uintptr(characterCount))
	return HRESULT(ret)
}

func PSPropertyBag_ReadStrAlloc(propBag *IPropertyBag, propName PWSTR, value *PWSTR) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadStrAlloc, libPropsys, "PSPropertyBag_ReadStrAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadBSTR(propBag *IPropertyBag, propName PWSTR, value *BSTR) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadBSTR, libPropsys, "PSPropertyBag_ReadBSTR")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteStr(propBag *IPropertyBag, propName PWSTR, value PWSTR) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteStr, libPropsys, "PSPropertyBag_WriteStr")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteBSTR(propBag *IPropertyBag, propName PWSTR, value BSTR) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteBSTR, libPropsys, "PSPropertyBag_WriteBSTR")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadInt(propBag *IPropertyBag, propName PWSTR, value *int32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadInt, libPropsys, "PSPropertyBag_ReadInt")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteInt(propBag *IPropertyBag, propName PWSTR, value int32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteInt, libPropsys, "PSPropertyBag_WriteInt")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(value))
	return HRESULT(ret)
}

func PSPropertyBag_ReadSHORT(propBag *IPropertyBag, propName PWSTR, value *int16) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadSHORT, libPropsys, "PSPropertyBag_ReadSHORT")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteSHORT(propBag *IPropertyBag, propName PWSTR, value int16) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteSHORT, libPropsys, "PSPropertyBag_WriteSHORT")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(value))
	return HRESULT(ret)
}

func PSPropertyBag_ReadLONG(propBag *IPropertyBag, propName PWSTR, value *int32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadLONG, libPropsys, "PSPropertyBag_ReadLONG")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteLONG(propBag *IPropertyBag, propName PWSTR, value int32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteLONG, libPropsys, "PSPropertyBag_WriteLONG")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(value))
	return HRESULT(ret)
}

func PSPropertyBag_ReadDWORD(propBag *IPropertyBag, propName PWSTR, value *uint32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadDWORD, libPropsys, "PSPropertyBag_ReadDWORD")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteDWORD(propBag *IPropertyBag, propName PWSTR, value uint32) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteDWORD, libPropsys, "PSPropertyBag_WriteDWORD")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(value))
	return HRESULT(ret)
}

func PSPropertyBag_ReadBOOL(propBag *IPropertyBag, propName PWSTR, value *BOOL) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadBOOL, libPropsys, "PSPropertyBag_ReadBOOL")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteBOOL(propBag *IPropertyBag, propName PWSTR, value BOOL) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteBOOL, libPropsys, "PSPropertyBag_WriteBOOL")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(value))
	return HRESULT(ret)
}

func PSPropertyBag_ReadPOINTL(propBag *IPropertyBag, propName PWSTR, value *POINTL) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadPOINTL, libPropsys, "PSPropertyBag_ReadPOINTL")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WritePOINTL(propBag *IPropertyBag, propName PWSTR, value *POINTL) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WritePOINTL, libPropsys, "PSPropertyBag_WritePOINTL")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadPOINTS(propBag *IPropertyBag, propName PWSTR, value *POINTS) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadPOINTS, libPropsys, "PSPropertyBag_ReadPOINTS")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WritePOINTS(propBag *IPropertyBag, propName PWSTR, value *POINTS) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WritePOINTS, libPropsys, "PSPropertyBag_WritePOINTS")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadRECTL(propBag *IPropertyBag, propName PWSTR, value *RECTL) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadRECTL, libPropsys, "PSPropertyBag_ReadRECTL")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteRECTL(propBag *IPropertyBag, propName PWSTR, value *RECTL) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteRECTL, libPropsys, "PSPropertyBag_WriteRECTL")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadStream(propBag *IPropertyBag, propName PWSTR, value **IStream) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadStream, libPropsys, "PSPropertyBag_ReadStream")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteStream(propBag *IPropertyBag, propName PWSTR, value *IStream) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteStream, libPropsys, "PSPropertyBag_WriteStream")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_Delete(propBag *IPropertyBag, propName PWSTR) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_Delete, libPropsys, "PSPropertyBag_Delete")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadULONGLONG(propBag *IPropertyBag, propName PWSTR, value *uint64) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadULONGLONG, libPropsys, "PSPropertyBag_ReadULONGLONG")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteULONGLONG(propBag *IPropertyBag, propName PWSTR, value uint64) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteULONGLONG, libPropsys, "PSPropertyBag_WriteULONGLONG")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(value))
	return HRESULT(ret)
}

func PSPropertyBag_ReadUnknown(propBag *IPropertyBag, propName PWSTR, riid *syscall.GUID, ppv unsafe.Pointer) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadUnknown, libPropsys, "PSPropertyBag_ReadUnknown")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(riid)), uintptr(ppv))
	return HRESULT(ret)
}

func PSPropertyBag_WriteUnknown(propBag *IPropertyBag, propName PWSTR, punk *IUnknown) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteUnknown, libPropsys, "PSPropertyBag_WriteUnknown")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(punk)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadGUID(propBag *IPropertyBag, propName PWSTR, value *syscall.GUID) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadGUID, libPropsys, "PSPropertyBag_ReadGUID")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WriteGUID(propBag *IPropertyBag, propName PWSTR, value *syscall.GUID) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WriteGUID, libPropsys, "PSPropertyBag_WriteGUID")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_ReadPropertyKey(propBag *IPropertyBag, propName PWSTR, value *PROPERTYKEY) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_ReadPropertyKey, libPropsys, "PSPropertyBag_ReadPropertyKey")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func PSPropertyBag_WritePropertyKey(propBag *IPropertyBag, propName PWSTR, value *PROPERTYKEY) HRESULT {
	addr := LazyAddr(&pPSPropertyBag_WritePropertyKey, libPropsys, "PSPropertyBag_WritePropertyKey")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propBag)), uintptr(unsafe.Pointer(propName)), uintptr(unsafe.Pointer(value)))
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

func InitPropVariantFromStrRet(pstrret *STRRET, pidl *ITEMIDLIST, ppropvar *PROPVARIANT) HRESULT {
	addr := LazyAddr(&pInitPropVariantFromStrRet, libPropsys, "InitPropVariantFromStrRet")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstrret)), uintptr(unsafe.Pointer(pidl)), uintptr(unsafe.Pointer(ppropvar)))
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

func PropVariantToStrRet(propvar *PROPVARIANT, pstrret *STRRET) HRESULT {
	addr := LazyAddr(&pPropVariantToStrRet, libPropsys, "PropVariantToStrRet")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(propvar)), uintptr(unsafe.Pointer(pstrret)))
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

func InitVariantFromResource(hinst HINSTANCE, id uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromResource, libPropsys, "InitVariantFromResource")
	ret, _, _ := syscall.SyscallN(addr, hinst, uintptr(id), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromBuffer(pv unsafe.Pointer, cb uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromBuffer, libPropsys, "InitVariantFromBuffer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(pv), uintptr(cb), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromGUIDAsString(guid *syscall.GUID, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromGUIDAsString, libPropsys, "InitVariantFromGUIDAsString")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromFileTime(pft *FILETIME, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromFileTime, libPropsys, "InitVariantFromFileTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pft)), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromFileTimeArray(prgft *FILETIME, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromFileTimeArray, libPropsys, "InitVariantFromFileTimeArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgft)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromStrRet(pstrret *STRRET, pidl *ITEMIDLIST, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromStrRet, libPropsys, "InitVariantFromStrRet")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(pstrret)), uintptr(unsafe.Pointer(pidl)), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromVariantArrayElem(varIn *VARIANT, iElem uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromVariantArrayElem, libPropsys, "InitVariantFromVariantArrayElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(iElem), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromBooleanArray(prgf *BOOL, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromBooleanArray, libPropsys, "InitVariantFromBooleanArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgf)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromInt16Array(prgn *int16, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromInt16Array, libPropsys, "InitVariantFromInt16Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromUInt16Array(prgn *uint16, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromUInt16Array, libPropsys, "InitVariantFromUInt16Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromInt32Array(prgn *int32, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromInt32Array, libPropsys, "InitVariantFromInt32Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromUInt32Array(prgn *uint32, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromUInt32Array, libPropsys, "InitVariantFromUInt32Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromInt64Array(prgn *int64, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromInt64Array, libPropsys, "InitVariantFromInt64Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromUInt64Array(prgn *uint64, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromUInt64Array, libPropsys, "InitVariantFromUInt64Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromDoubleArray(prgn *float64, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromDoubleArray, libPropsys, "InitVariantFromDoubleArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgn)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func InitVariantFromStringArray(prgsz *PWSTR, cElems uint32, pvar *VARIANT) HRESULT {
	addr := LazyAddr(&pInitVariantFromStringArray, libPropsys, "InitVariantFromStringArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(prgsz)), uintptr(cElems), uintptr(unsafe.Pointer(pvar)))
	return HRESULT(ret)
}

func VariantToBooleanWithDefault(varIn *VARIANT, fDefault BOOL) BOOL {
	addr := LazyAddr(&pVariantToBooleanWithDefault, libPropsys, "VariantToBooleanWithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(fDefault))
	return BOOL(ret)
}

func VariantToInt16WithDefault(varIn *VARIANT, iDefault int16) int16 {
	addr := LazyAddr(&pVariantToInt16WithDefault, libPropsys, "VariantToInt16WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(iDefault))
	return int16(ret)
}

func VariantToUInt16WithDefault(varIn *VARIANT, uiDefault uint16) uint16 {
	addr := LazyAddr(&pVariantToUInt16WithDefault, libPropsys, "VariantToUInt16WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(uiDefault))
	return uint16(ret)
}

func VariantToInt32WithDefault(varIn *VARIANT, lDefault int32) int32 {
	addr := LazyAddr(&pVariantToInt32WithDefault, libPropsys, "VariantToInt32WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(lDefault))
	return int32(ret)
}

func VariantToUInt32WithDefault(varIn *VARIANT, ulDefault uint32) uint32 {
	addr := LazyAddr(&pVariantToUInt32WithDefault, libPropsys, "VariantToUInt32WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(ulDefault))
	return uint32(ret)
}

func VariantToInt64WithDefault(varIn *VARIANT, llDefault int64) int64 {
	addr := LazyAddr(&pVariantToInt64WithDefault, libPropsys, "VariantToInt64WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(llDefault))
	return int64(ret)
}

func VariantToUInt64WithDefault(varIn *VARIANT, ullDefault uint64) uint64 {
	addr := LazyAddr(&pVariantToUInt64WithDefault, libPropsys, "VariantToUInt64WithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(ullDefault))
	return uint64(ret)
}

func VariantToDoubleWithDefault(varIn *VARIANT, dblDefault float64) float64 {
	addr := LazyAddr(&pVariantToDoubleWithDefault, libPropsys, "VariantToDoubleWithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(dblDefault))
	return float64(ret)
}

func VariantToStringWithDefault(varIn *VARIANT, pszDefault PWSTR) PWSTR {
	addr := LazyAddr(&pVariantToStringWithDefault, libPropsys, "VariantToStringWithDefault")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pszDefault)))
	return (PWSTR)(unsafe.Pointer(ret))
}

func VariantToBoolean(varIn *VARIANT, pfRet *BOOL) HRESULT {
	addr := LazyAddr(&pVariantToBoolean, libPropsys, "VariantToBoolean")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pfRet)))
	return HRESULT(ret)
}

func VariantToInt16(varIn *VARIANT, piRet *int16) HRESULT {
	addr := LazyAddr(&pVariantToInt16, libPropsys, "VariantToInt16")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(piRet)))
	return HRESULT(ret)
}

func VariantToUInt16(varIn *VARIANT, puiRet *uint16) HRESULT {
	addr := LazyAddr(&pVariantToUInt16, libPropsys, "VariantToUInt16")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(puiRet)))
	return HRESULT(ret)
}

func VariantToInt32(varIn *VARIANT, plRet *int32) HRESULT {
	addr := LazyAddr(&pVariantToInt32, libPropsys, "VariantToInt32")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(plRet)))
	return HRESULT(ret)
}

func VariantToUInt32(varIn *VARIANT, pulRet *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt32, libPropsys, "VariantToUInt32")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pulRet)))
	return HRESULT(ret)
}

func VariantToInt64(varIn *VARIANT, pllRet *int64) HRESULT {
	addr := LazyAddr(&pVariantToInt64, libPropsys, "VariantToInt64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pllRet)))
	return HRESULT(ret)
}

func VariantToUInt64(varIn *VARIANT, pullRet *uint64) HRESULT {
	addr := LazyAddr(&pVariantToUInt64, libPropsys, "VariantToUInt64")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pullRet)))
	return HRESULT(ret)
}

func VariantToDouble(varIn *VARIANT, pdblRet *float64) HRESULT {
	addr := LazyAddr(&pVariantToDouble, libPropsys, "VariantToDouble")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pdblRet)))
	return HRESULT(ret)
}

func VariantToBuffer(varIn *VARIANT, pv unsafe.Pointer, cb uint32) HRESULT {
	addr := LazyAddr(&pVariantToBuffer, libPropsys, "VariantToBuffer")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(pv), uintptr(cb))
	return HRESULT(ret)
}

func VariantToGUID(varIn *VARIANT, pguid *syscall.GUID) HRESULT {
	addr := LazyAddr(&pVariantToGUID, libPropsys, "VariantToGUID")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pguid)))
	return HRESULT(ret)
}

func VariantToString(varIn *VARIANT, pszBuf PWSTR, cchBuf uint32) HRESULT {
	addr := LazyAddr(&pVariantToString, libPropsys, "VariantToString")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pszBuf)), uintptr(cchBuf))
	return HRESULT(ret)
}

func VariantToStringAlloc(varIn *VARIANT, ppszBuf *PWSTR) HRESULT {
	addr := LazyAddr(&pVariantToStringAlloc, libPropsys, "VariantToStringAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(ppszBuf)))
	return HRESULT(ret)
}

func VariantToDosDateTime(varIn *VARIANT, pwDate *uint16, pwTime *uint16) HRESULT {
	addr := LazyAddr(&pVariantToDosDateTime, libPropsys, "VariantToDosDateTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pwDate)), uintptr(unsafe.Pointer(pwTime)))
	return HRESULT(ret)
}

func VariantToStrRet(varIn *VARIANT, pstrret *STRRET) HRESULT {
	addr := LazyAddr(&pVariantToStrRet, libPropsys, "VariantToStrRet")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(unsafe.Pointer(pstrret)))
	return HRESULT(ret)
}

func VariantToFileTime(varIn *VARIANT, stfOut PSTIME_FLAGS, pftOut *FILETIME) HRESULT {
	addr := LazyAddr(&pVariantToFileTime, libPropsys, "VariantToFileTime")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)), uintptr(stfOut), uintptr(unsafe.Pointer(pftOut)))
	return HRESULT(ret)
}

func VariantGetElementCount(varIn *VARIANT) uint32 {
	addr := LazyAddr(&pVariantGetElementCount, libPropsys, "VariantGetElementCount")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(varIn)))
	return uint32(ret)
}

func VariantToBooleanArray(var_ *VARIANT, prgf *BOOL, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToBooleanArray, libPropsys, "VariantToBooleanArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgf)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToInt16Array(var_ *VARIANT, prgn *int16, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToInt16Array, libPropsys, "VariantToInt16Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToUInt16Array(var_ *VARIANT, prgn *uint16, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt16Array, libPropsys, "VariantToUInt16Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToInt32Array(var_ *VARIANT, prgn *int32, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToInt32Array, libPropsys, "VariantToInt32Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToUInt32Array(var_ *VARIANT, prgn *uint32, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt32Array, libPropsys, "VariantToUInt32Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToInt64Array(var_ *VARIANT, prgn *int64, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToInt64Array, libPropsys, "VariantToInt64Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToUInt64Array(var_ *VARIANT, prgn *uint64, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt64Array, libPropsys, "VariantToUInt64Array")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToDoubleArray(var_ *VARIANT, prgn *float64, crgn uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToDoubleArray, libPropsys, "VariantToDoubleArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgn)), uintptr(crgn), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToStringArray(var_ *VARIANT, prgsz *PWSTR, crgsz uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToStringArray, libPropsys, "VariantToStringArray")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(prgsz)), uintptr(crgsz), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToBooleanArrayAlloc(var_ *VARIANT, pprgf **BOOL, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToBooleanArrayAlloc, libPropsys, "VariantToBooleanArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgf)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToInt16ArrayAlloc(var_ *VARIANT, pprgn **int16, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToInt16ArrayAlloc, libPropsys, "VariantToInt16ArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToUInt16ArrayAlloc(var_ *VARIANT, pprgn **uint16, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt16ArrayAlloc, libPropsys, "VariantToUInt16ArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToInt32ArrayAlloc(var_ *VARIANT, pprgn **int32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToInt32ArrayAlloc, libPropsys, "VariantToInt32ArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToUInt32ArrayAlloc(var_ *VARIANT, pprgn **uint32, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt32ArrayAlloc, libPropsys, "VariantToUInt32ArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToInt64ArrayAlloc(var_ *VARIANT, pprgn **int64, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToInt64ArrayAlloc, libPropsys, "VariantToInt64ArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToUInt64ArrayAlloc(var_ *VARIANT, pprgn **uint64, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToUInt64ArrayAlloc, libPropsys, "VariantToUInt64ArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToDoubleArrayAlloc(var_ *VARIANT, pprgn **float64, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToDoubleArrayAlloc, libPropsys, "VariantToDoubleArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgn)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantToStringArrayAlloc(var_ *VARIANT, pprgsz **PWSTR, pcElem *uint32) HRESULT {
	addr := LazyAddr(&pVariantToStringArrayAlloc, libPropsys, "VariantToStringArrayAlloc")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(unsafe.Pointer(pprgsz)), uintptr(unsafe.Pointer(pcElem)))
	return HRESULT(ret)
}

func VariantGetBooleanElem(var_ *VARIANT, iElem uint32, pfVal *BOOL) HRESULT {
	addr := LazyAddr(&pVariantGetBooleanElem, libPropsys, "VariantGetBooleanElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pfVal)))
	return HRESULT(ret)
}

func VariantGetInt16Elem(var_ *VARIANT, iElem uint32, pnVal *int16) HRESULT {
	addr := LazyAddr(&pVariantGetInt16Elem, libPropsys, "VariantGetInt16Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetUInt16Elem(var_ *VARIANT, iElem uint32, pnVal *uint16) HRESULT {
	addr := LazyAddr(&pVariantGetUInt16Elem, libPropsys, "VariantGetUInt16Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetInt32Elem(var_ *VARIANT, iElem uint32, pnVal *int32) HRESULT {
	addr := LazyAddr(&pVariantGetInt32Elem, libPropsys, "VariantGetInt32Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetUInt32Elem(var_ *VARIANT, iElem uint32, pnVal *uint32) HRESULT {
	addr := LazyAddr(&pVariantGetUInt32Elem, libPropsys, "VariantGetUInt32Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetInt64Elem(var_ *VARIANT, iElem uint32, pnVal *int64) HRESULT {
	addr := LazyAddr(&pVariantGetInt64Elem, libPropsys, "VariantGetInt64Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetUInt64Elem(var_ *VARIANT, iElem uint32, pnVal *uint64) HRESULT {
	addr := LazyAddr(&pVariantGetUInt64Elem, libPropsys, "VariantGetUInt64Elem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetDoubleElem(var_ *VARIANT, iElem uint32, pnVal *float64) HRESULT {
	addr := LazyAddr(&pVariantGetDoubleElem, libPropsys, "VariantGetDoubleElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(pnVal)))
	return HRESULT(ret)
}

func VariantGetStringElem(var_ *VARIANT, iElem uint32, ppszVal *PWSTR) HRESULT {
	addr := LazyAddr(&pVariantGetStringElem, libPropsys, "VariantGetStringElem")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var_)), uintptr(iElem), uintptr(unsafe.Pointer(ppszVal)))
	return HRESULT(ret)
}

func ClearVariantArray(pvars *VARIANT, cvars uint32) {
	addr := LazyAddr(&pClearVariantArray, libPropsys, "ClearVariantArray")
	syscall.SyscallN(addr, uintptr(unsafe.Pointer(pvars)), uintptr(cvars))
}

func VariantCompare(var1 *VARIANT, var2 *VARIANT) int32 {
	addr := LazyAddr(&pVariantCompare, libPropsys, "VariantCompare")
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(var1)), uintptr(unsafe.Pointer(var2)))
	return int32(ret)
}
