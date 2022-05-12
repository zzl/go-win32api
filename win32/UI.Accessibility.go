package win32

import "unsafe"
import "syscall"

type HWINEVENTHOOK = uintptr
type HUIANODE = uintptr
type HUIAPATTERNOBJECT = uintptr
type HUIATEXTRANGE = uintptr
type HUIAEVENT = uintptr

const (
	ANRUS_PRIORITY_AUDIO_DYNAMIC_DUCK uint32 = 16
	MSAA_MENU_SIG int32 = -1441927155
	DISPID_ACC_PARENT int32 = -5000
	DISPID_ACC_CHILDCOUNT int32 = -5001
	DISPID_ACC_CHILD int32 = -5002
	DISPID_ACC_NAME int32 = -5003
	DISPID_ACC_VALUE int32 = -5004
	DISPID_ACC_DESCRIPTION int32 = -5005
	DISPID_ACC_ROLE int32 = -5006
	DISPID_ACC_STATE int32 = -5007
	DISPID_ACC_HELP int32 = -5008
	DISPID_ACC_HELPTOPIC int32 = -5009
	DISPID_ACC_KEYBOARDSHORTCUT int32 = -5010
	DISPID_ACC_FOCUS int32 = -5011
	DISPID_ACC_SELECTION int32 = -5012
	DISPID_ACC_DEFAULTACTION int32 = -5013
	DISPID_ACC_SELECT int32 = -5014
	DISPID_ACC_LOCATION int32 = -5015
	DISPID_ACC_NAVIGATE int32 = -5016
	DISPID_ACC_HITTEST int32 = -5017
	DISPID_ACC_DODEFAULTACTION int32 = -5018
	NAVDIR_MIN uint32 = 0
	NAVDIR_UP uint32 = 1
	NAVDIR_DOWN uint32 = 2
	NAVDIR_LEFT uint32 = 3
	NAVDIR_RIGHT uint32 = 4
	NAVDIR_NEXT uint32 = 5
	NAVDIR_PREVIOUS uint32 = 6
	NAVDIR_FIRSTCHILD uint32 = 7
	NAVDIR_LASTCHILD uint32 = 8
	NAVDIR_MAX uint32 = 9
	SELFLAG_NONE uint32 = 0
	SELFLAG_TAKEFOCUS uint32 = 1
	SELFLAG_TAKESELECTION uint32 = 2
	SELFLAG_EXTENDSELECTION uint32 = 4
	SELFLAG_ADDSELECTION uint32 = 8
	SELFLAG_REMOVESELECTION uint32 = 16
	SELFLAG_VALID uint32 = 31
	STATE_SYSTEM_NORMAL uint32 = 0
	STATE_SYSTEM_HASPOPUP uint32 = 1073741824
	ROLE_SYSTEM_TITLEBAR uint32 = 1
	ROLE_SYSTEM_MENUBAR uint32 = 2
	ROLE_SYSTEM_SCROLLBAR uint32 = 3
	ROLE_SYSTEM_GRIP uint32 = 4
	ROLE_SYSTEM_SOUND uint32 = 5
	ROLE_SYSTEM_CURSOR uint32 = 6
	ROLE_SYSTEM_CARET uint32 = 7
	ROLE_SYSTEM_ALERT uint32 = 8
	ROLE_SYSTEM_WINDOW uint32 = 9
	ROLE_SYSTEM_CLIENT uint32 = 10
	ROLE_SYSTEM_MENUPOPUP uint32 = 11
	ROLE_SYSTEM_MENUITEM uint32 = 12
	ROLE_SYSTEM_TOOLTIP uint32 = 13
	ROLE_SYSTEM_APPLICATION uint32 = 14
	ROLE_SYSTEM_DOCUMENT uint32 = 15
	ROLE_SYSTEM_PANE uint32 = 16
	ROLE_SYSTEM_CHART uint32 = 17
	ROLE_SYSTEM_DIALOG uint32 = 18
	ROLE_SYSTEM_BORDER uint32 = 19
	ROLE_SYSTEM_GROUPING uint32 = 20
	ROLE_SYSTEM_SEPARATOR uint32 = 21
	ROLE_SYSTEM_TOOLBAR uint32 = 22
	ROLE_SYSTEM_STATUSBAR uint32 = 23
	ROLE_SYSTEM_TABLE uint32 = 24
	ROLE_SYSTEM_COLUMNHEADER uint32 = 25
	ROLE_SYSTEM_ROWHEADER uint32 = 26
	ROLE_SYSTEM_COLUMN uint32 = 27
	ROLE_SYSTEM_ROW uint32 = 28
	ROLE_SYSTEM_CELL uint32 = 29
	ROLE_SYSTEM_LINK uint32 = 30
	ROLE_SYSTEM_HELPBALLOON uint32 = 31
	ROLE_SYSTEM_CHARACTER uint32 = 32
	ROLE_SYSTEM_LIST uint32 = 33
	ROLE_SYSTEM_LISTITEM uint32 = 34
	ROLE_SYSTEM_OUTLINE uint32 = 35
	ROLE_SYSTEM_OUTLINEITEM uint32 = 36
	ROLE_SYSTEM_PAGETAB uint32 = 37
	ROLE_SYSTEM_PROPERTYPAGE uint32 = 38
	ROLE_SYSTEM_INDICATOR uint32 = 39
	ROLE_SYSTEM_GRAPHIC uint32 = 40
	ROLE_SYSTEM_STATICTEXT uint32 = 41
	ROLE_SYSTEM_TEXT uint32 = 42
	ROLE_SYSTEM_PUSHBUTTON uint32 = 43
	ROLE_SYSTEM_CHECKBUTTON uint32 = 44
	ROLE_SYSTEM_RADIOBUTTON uint32 = 45
	ROLE_SYSTEM_COMBOBOX uint32 = 46
	ROLE_SYSTEM_DROPLIST uint32 = 47
	ROLE_SYSTEM_PROGRESSBAR uint32 = 48
	ROLE_SYSTEM_DIAL uint32 = 49
	ROLE_SYSTEM_HOTKEYFIELD uint32 = 50
	ROLE_SYSTEM_SLIDER uint32 = 51
	ROLE_SYSTEM_SPINBUTTON uint32 = 52
	ROLE_SYSTEM_DIAGRAM uint32 = 53
	ROLE_SYSTEM_ANIMATION uint32 = 54
	ROLE_SYSTEM_EQUATION uint32 = 55
	ROLE_SYSTEM_BUTTONDROPDOWN uint32 = 56
	ROLE_SYSTEM_BUTTONMENU uint32 = 57
	ROLE_SYSTEM_BUTTONDROPDOWNGRID uint32 = 58
	ROLE_SYSTEM_WHITESPACE uint32 = 59
	ROLE_SYSTEM_PAGETABLIST uint32 = 60
	ROLE_SYSTEM_CLOCK uint32 = 61
	ROLE_SYSTEM_SPLITBUTTON uint32 = 62
	ROLE_SYSTEM_IPADDRESS uint32 = 63
	ROLE_SYSTEM_OUTLINEBUTTON uint32 = 64
	UIA_E_ELEMENTNOTENABLED uint32 = 2147746304
	UIA_E_ELEMENTNOTAVAILABLE uint32 = 2147746305
	UIA_E_NOCLICKABLEPOINT uint32 = 2147746306
	UIA_E_PROXYASSEMBLYNOTLOADED uint32 = 2147746307
	UIA_E_NOTSUPPORTED uint32 = 2147746308
	UIA_E_INVALIDOPERATION uint32 = 2148734217
	UIA_E_TIMEOUT uint32 = 2148734213
	UiaAppendRuntimeId uint32 = 3
	UiaRootObjectId int32 = -25
	UIA_IAFP_DEFAULT uint32 = 0
	UIA_IAFP_UNWRAP_BRIDGE uint32 = 1
	UIA_PFIA_DEFAULT uint32 = 0
	UIA_PFIA_UNWRAP_BRIDGE uint32 = 1
	UIA_ScrollPatternNoScroll float64 = -1
	UIA_InvokePatternId int32 = 10000
	UIA_SelectionPatternId int32 = 10001
	UIA_ValuePatternId int32 = 10002
	UIA_RangeValuePatternId int32 = 10003
	UIA_ScrollPatternId int32 = 10004
	UIA_ExpandCollapsePatternId int32 = 10005
	UIA_GridPatternId int32 = 10006
	UIA_GridItemPatternId int32 = 10007
	UIA_MultipleViewPatternId int32 = 10008
	UIA_WindowPatternId int32 = 10009
	UIA_SelectionItemPatternId int32 = 10010
	UIA_DockPatternId int32 = 10011
	UIA_TablePatternId int32 = 10012
	UIA_TableItemPatternId int32 = 10013
	UIA_TextPatternId int32 = 10014
	UIA_TogglePatternId int32 = 10015
	UIA_TransformPatternId int32 = 10016
	UIA_ScrollItemPatternId int32 = 10017
	UIA_LegacyIAccessiblePatternId int32 = 10018
	UIA_ItemContainerPatternId int32 = 10019
	UIA_VirtualizedItemPatternId int32 = 10020
	UIA_SynchronizedInputPatternId int32 = 10021
	UIA_ObjectModelPatternId int32 = 10022
	UIA_AnnotationPatternId int32 = 10023
	UIA_TextPattern2Id int32 = 10024
	UIA_StylesPatternId int32 = 10025
	UIA_SpreadsheetPatternId int32 = 10026
	UIA_SpreadsheetItemPatternId int32 = 10027
	UIA_TransformPattern2Id int32 = 10028
	UIA_TextChildPatternId int32 = 10029
	UIA_DragPatternId int32 = 10030
	UIA_DropTargetPatternId int32 = 10031
	UIA_TextEditPatternId int32 = 10032
	UIA_CustomNavigationPatternId int32 = 10033
	UIA_SelectionPattern2Id int32 = 10034
	UIA_ToolTipOpenedEventId int32 = 20000
	UIA_ToolTipClosedEventId int32 = 20001
	UIA_StructureChangedEventId int32 = 20002
	UIA_MenuOpenedEventId int32 = 20003
	UIA_AutomationPropertyChangedEventId int32 = 20004
	UIA_AutomationFocusChangedEventId int32 = 20005
	UIA_AsyncContentLoadedEventId int32 = 20006
	UIA_MenuClosedEventId int32 = 20007
	UIA_LayoutInvalidatedEventId int32 = 20008
	UIA_Invoke_InvokedEventId int32 = 20009
	UIA_SelectionItem_ElementAddedToSelectionEventId int32 = 20010
	UIA_SelectionItem_ElementRemovedFromSelectionEventId int32 = 20011
	UIA_SelectionItem_ElementSelectedEventId int32 = 20012
	UIA_Selection_InvalidatedEventId int32 = 20013
	UIA_Text_TextSelectionChangedEventId int32 = 20014
	UIA_Text_TextChangedEventId int32 = 20015
	UIA_Window_WindowOpenedEventId int32 = 20016
	UIA_Window_WindowClosedEventId int32 = 20017
	UIA_MenuModeStartEventId int32 = 20018
	UIA_MenuModeEndEventId int32 = 20019
	UIA_InputReachedTargetEventId int32 = 20020
	UIA_InputReachedOtherElementEventId int32 = 20021
	UIA_InputDiscardedEventId int32 = 20022
	UIA_SystemAlertEventId int32 = 20023
	UIA_LiveRegionChangedEventId int32 = 20024
	UIA_HostedFragmentRootsInvalidatedEventId int32 = 20025
	UIA_Drag_DragStartEventId int32 = 20026
	UIA_Drag_DragCancelEventId int32 = 20027
	UIA_Drag_DragCompleteEventId int32 = 20028
	UIA_DropTarget_DragEnterEventId int32 = 20029
	UIA_DropTarget_DragLeaveEventId int32 = 20030
	UIA_DropTarget_DroppedEventId int32 = 20031
	UIA_TextEdit_TextChangedEventId int32 = 20032
	UIA_TextEdit_ConversionTargetChangedEventId int32 = 20033
	UIA_ChangesEventId int32 = 20034
	UIA_NotificationEventId int32 = 20035
	UIA_ActiveTextPositionChangedEventId int32 = 20036
	UIA_RuntimeIdPropertyId int32 = 30000
	UIA_BoundingRectanglePropertyId int32 = 30001
	UIA_ProcessIdPropertyId int32 = 30002
	UIA_ControlTypePropertyId int32 = 30003
	UIA_LocalizedControlTypePropertyId int32 = 30004
	UIA_NamePropertyId int32 = 30005
	UIA_AcceleratorKeyPropertyId int32 = 30006
	UIA_AccessKeyPropertyId int32 = 30007
	UIA_HasKeyboardFocusPropertyId int32 = 30008
	UIA_IsKeyboardFocusablePropertyId int32 = 30009
	UIA_IsEnabledPropertyId int32 = 30010
	UIA_AutomationIdPropertyId int32 = 30011
	UIA_ClassNamePropertyId int32 = 30012
	UIA_HelpTextPropertyId int32 = 30013
	UIA_ClickablePointPropertyId int32 = 30014
	UIA_CulturePropertyId int32 = 30015
	UIA_IsControlElementPropertyId int32 = 30016
	UIA_IsContentElementPropertyId int32 = 30017
	UIA_LabeledByPropertyId int32 = 30018
	UIA_IsPasswordPropertyId int32 = 30019
	UIA_NativeWindowHandlePropertyId int32 = 30020
	UIA_ItemTypePropertyId int32 = 30021
	UIA_IsOffscreenPropertyId int32 = 30022
	UIA_OrientationPropertyId int32 = 30023
	UIA_FrameworkIdPropertyId int32 = 30024
	UIA_IsRequiredForFormPropertyId int32 = 30025
	UIA_ItemStatusPropertyId int32 = 30026
	UIA_IsDockPatternAvailablePropertyId int32 = 30027
	UIA_IsExpandCollapsePatternAvailablePropertyId int32 = 30028
	UIA_IsGridItemPatternAvailablePropertyId int32 = 30029
	UIA_IsGridPatternAvailablePropertyId int32 = 30030
	UIA_IsInvokePatternAvailablePropertyId int32 = 30031
	UIA_IsMultipleViewPatternAvailablePropertyId int32 = 30032
	UIA_IsRangeValuePatternAvailablePropertyId int32 = 30033
	UIA_IsScrollPatternAvailablePropertyId int32 = 30034
	UIA_IsScrollItemPatternAvailablePropertyId int32 = 30035
	UIA_IsSelectionItemPatternAvailablePropertyId int32 = 30036
	UIA_IsSelectionPatternAvailablePropertyId int32 = 30037
	UIA_IsTablePatternAvailablePropertyId int32 = 30038
	UIA_IsTableItemPatternAvailablePropertyId int32 = 30039
	UIA_IsTextPatternAvailablePropertyId int32 = 30040
	UIA_IsTogglePatternAvailablePropertyId int32 = 30041
	UIA_IsTransformPatternAvailablePropertyId int32 = 30042
	UIA_IsValuePatternAvailablePropertyId int32 = 30043
	UIA_IsWindowPatternAvailablePropertyId int32 = 30044
	UIA_ValueValuePropertyId int32 = 30045
	UIA_ValueIsReadOnlyPropertyId int32 = 30046
	UIA_RangeValueValuePropertyId int32 = 30047
	UIA_RangeValueIsReadOnlyPropertyId int32 = 30048
	UIA_RangeValueMinimumPropertyId int32 = 30049
	UIA_RangeValueMaximumPropertyId int32 = 30050
	UIA_RangeValueLargeChangePropertyId int32 = 30051
	UIA_RangeValueSmallChangePropertyId int32 = 30052
	UIA_ScrollHorizontalScrollPercentPropertyId int32 = 30053
	UIA_ScrollHorizontalViewSizePropertyId int32 = 30054
	UIA_ScrollVerticalScrollPercentPropertyId int32 = 30055
	UIA_ScrollVerticalViewSizePropertyId int32 = 30056
	UIA_ScrollHorizontallyScrollablePropertyId int32 = 30057
	UIA_ScrollVerticallyScrollablePropertyId int32 = 30058
	UIA_SelectionSelectionPropertyId int32 = 30059
	UIA_SelectionCanSelectMultiplePropertyId int32 = 30060
	UIA_SelectionIsSelectionRequiredPropertyId int32 = 30061
	UIA_GridRowCountPropertyId int32 = 30062
	UIA_GridColumnCountPropertyId int32 = 30063
	UIA_GridItemRowPropertyId int32 = 30064
	UIA_GridItemColumnPropertyId int32 = 30065
	UIA_GridItemRowSpanPropertyId int32 = 30066
	UIA_GridItemColumnSpanPropertyId int32 = 30067
	UIA_GridItemContainingGridPropertyId int32 = 30068
	UIA_DockDockPositionPropertyId int32 = 30069
	UIA_ExpandCollapseExpandCollapseStatePropertyId int32 = 30070
	UIA_MultipleViewCurrentViewPropertyId int32 = 30071
	UIA_MultipleViewSupportedViewsPropertyId int32 = 30072
	UIA_WindowCanMaximizePropertyId int32 = 30073
	UIA_WindowCanMinimizePropertyId int32 = 30074
	UIA_WindowWindowVisualStatePropertyId int32 = 30075
	UIA_WindowWindowInteractionStatePropertyId int32 = 30076
	UIA_WindowIsModalPropertyId int32 = 30077
	UIA_WindowIsTopmostPropertyId int32 = 30078
	UIA_SelectionItemIsSelectedPropertyId int32 = 30079
	UIA_SelectionItemSelectionContainerPropertyId int32 = 30080
	UIA_TableRowHeadersPropertyId int32 = 30081
	UIA_TableColumnHeadersPropertyId int32 = 30082
	UIA_TableRowOrColumnMajorPropertyId int32 = 30083
	UIA_TableItemRowHeaderItemsPropertyId int32 = 30084
	UIA_TableItemColumnHeaderItemsPropertyId int32 = 30085
	UIA_ToggleToggleStatePropertyId int32 = 30086
	UIA_TransformCanMovePropertyId int32 = 30087
	UIA_TransformCanResizePropertyId int32 = 30088
	UIA_TransformCanRotatePropertyId int32 = 30089
	UIA_IsLegacyIAccessiblePatternAvailablePropertyId int32 = 30090
	UIA_LegacyIAccessibleChildIdPropertyId int32 = 30091
	UIA_LegacyIAccessibleNamePropertyId int32 = 30092
	UIA_LegacyIAccessibleValuePropertyId int32 = 30093
	UIA_LegacyIAccessibleDescriptionPropertyId int32 = 30094
	UIA_LegacyIAccessibleRolePropertyId int32 = 30095
	UIA_LegacyIAccessibleStatePropertyId int32 = 30096
	UIA_LegacyIAccessibleHelpPropertyId int32 = 30097
	UIA_LegacyIAccessibleKeyboardShortcutPropertyId int32 = 30098
	UIA_LegacyIAccessibleSelectionPropertyId int32 = 30099
	UIA_LegacyIAccessibleDefaultActionPropertyId int32 = 30100
	UIA_AriaRolePropertyId int32 = 30101
	UIA_AriaPropertiesPropertyId int32 = 30102
	UIA_IsDataValidForFormPropertyId int32 = 30103
	UIA_ControllerForPropertyId int32 = 30104
	UIA_DescribedByPropertyId int32 = 30105
	UIA_FlowsToPropertyId int32 = 30106
	UIA_ProviderDescriptionPropertyId int32 = 30107
	UIA_IsItemContainerPatternAvailablePropertyId int32 = 30108
	UIA_IsVirtualizedItemPatternAvailablePropertyId int32 = 30109
	UIA_IsSynchronizedInputPatternAvailablePropertyId int32 = 30110
	UIA_OptimizeForVisualContentPropertyId int32 = 30111
	UIA_IsObjectModelPatternAvailablePropertyId int32 = 30112
	UIA_AnnotationAnnotationTypeIdPropertyId int32 = 30113
	UIA_AnnotationAnnotationTypeNamePropertyId int32 = 30114
	UIA_AnnotationAuthorPropertyId int32 = 30115
	UIA_AnnotationDateTimePropertyId int32 = 30116
	UIA_AnnotationTargetPropertyId int32 = 30117
	UIA_IsAnnotationPatternAvailablePropertyId int32 = 30118
	UIA_IsTextPattern2AvailablePropertyId int32 = 30119
	UIA_StylesStyleIdPropertyId int32 = 30120
	UIA_StylesStyleNamePropertyId int32 = 30121
	UIA_StylesFillColorPropertyId int32 = 30122
	UIA_StylesFillPatternStylePropertyId int32 = 30123
	UIA_StylesShapePropertyId int32 = 30124
	UIA_StylesFillPatternColorPropertyId int32 = 30125
	UIA_StylesExtendedPropertiesPropertyId int32 = 30126
	UIA_IsStylesPatternAvailablePropertyId int32 = 30127
	UIA_IsSpreadsheetPatternAvailablePropertyId int32 = 30128
	UIA_SpreadsheetItemFormulaPropertyId int32 = 30129
	UIA_SpreadsheetItemAnnotationObjectsPropertyId int32 = 30130
	UIA_SpreadsheetItemAnnotationTypesPropertyId int32 = 30131
	UIA_IsSpreadsheetItemPatternAvailablePropertyId int32 = 30132
	UIA_Transform2CanZoomPropertyId int32 = 30133
	UIA_IsTransformPattern2AvailablePropertyId int32 = 30134
	UIA_LiveSettingPropertyId int32 = 30135
	UIA_IsTextChildPatternAvailablePropertyId int32 = 30136
	UIA_IsDragPatternAvailablePropertyId int32 = 30137
	UIA_DragIsGrabbedPropertyId int32 = 30138
	UIA_DragDropEffectPropertyId int32 = 30139
	UIA_DragDropEffectsPropertyId int32 = 30140
	UIA_IsDropTargetPatternAvailablePropertyId int32 = 30141
	UIA_DropTargetDropTargetEffectPropertyId int32 = 30142
	UIA_DropTargetDropTargetEffectsPropertyId int32 = 30143
	UIA_DragGrabbedItemsPropertyId int32 = 30144
	UIA_Transform2ZoomLevelPropertyId int32 = 30145
	UIA_Transform2ZoomMinimumPropertyId int32 = 30146
	UIA_Transform2ZoomMaximumPropertyId int32 = 30147
	UIA_FlowsFromPropertyId int32 = 30148
	UIA_IsTextEditPatternAvailablePropertyId int32 = 30149
	UIA_IsPeripheralPropertyId int32 = 30150
	UIA_IsCustomNavigationPatternAvailablePropertyId int32 = 30151
	UIA_PositionInSetPropertyId int32 = 30152
	UIA_SizeOfSetPropertyId int32 = 30153
	UIA_LevelPropertyId int32 = 30154
	UIA_AnnotationTypesPropertyId int32 = 30155
	UIA_AnnotationObjectsPropertyId int32 = 30156
	UIA_LandmarkTypePropertyId int32 = 30157
	UIA_LocalizedLandmarkTypePropertyId int32 = 30158
	UIA_FullDescriptionPropertyId int32 = 30159
	UIA_FillColorPropertyId int32 = 30160
	UIA_OutlineColorPropertyId int32 = 30161
	UIA_FillTypePropertyId int32 = 30162
	UIA_VisualEffectsPropertyId int32 = 30163
	UIA_OutlineThicknessPropertyId int32 = 30164
	UIA_CenterPointPropertyId int32 = 30165
	UIA_RotationPropertyId int32 = 30166
	UIA_SizePropertyId int32 = 30167
	UIA_IsSelectionPattern2AvailablePropertyId int32 = 30168
	UIA_Selection2FirstSelectedItemPropertyId int32 = 30169
	UIA_Selection2LastSelectedItemPropertyId int32 = 30170
	UIA_Selection2CurrentSelectedItemPropertyId int32 = 30171
	UIA_Selection2ItemCountPropertyId int32 = 30172
	UIA_HeadingLevelPropertyId int32 = 30173
	UIA_IsDialogPropertyId int32 = 30174
	UIA_AnimationStyleAttributeId int32 = 40000
	UIA_BackgroundColorAttributeId int32 = 40001
	UIA_BulletStyleAttributeId int32 = 40002
	UIA_CapStyleAttributeId int32 = 40003
	UIA_CultureAttributeId int32 = 40004
	UIA_FontNameAttributeId int32 = 40005
	UIA_FontSizeAttributeId int32 = 40006
	UIA_FontWeightAttributeId int32 = 40007
	UIA_ForegroundColorAttributeId int32 = 40008
	UIA_HorizontalTextAlignmentAttributeId int32 = 40009
	UIA_IndentationFirstLineAttributeId int32 = 40010
	UIA_IndentationLeadingAttributeId int32 = 40011
	UIA_IndentationTrailingAttributeId int32 = 40012
	UIA_IsHiddenAttributeId int32 = 40013
	UIA_IsItalicAttributeId int32 = 40014
	UIA_IsReadOnlyAttributeId int32 = 40015
	UIA_IsSubscriptAttributeId int32 = 40016
	UIA_IsSuperscriptAttributeId int32 = 40017
	UIA_MarginBottomAttributeId int32 = 40018
	UIA_MarginLeadingAttributeId int32 = 40019
	UIA_MarginTopAttributeId int32 = 40020
	UIA_MarginTrailingAttributeId int32 = 40021
	UIA_OutlineStylesAttributeId int32 = 40022
	UIA_OverlineColorAttributeId int32 = 40023
	UIA_OverlineStyleAttributeId int32 = 40024
	UIA_StrikethroughColorAttributeId int32 = 40025
	UIA_StrikethroughStyleAttributeId int32 = 40026
	UIA_TabsAttributeId int32 = 40027
	UIA_TextFlowDirectionsAttributeId int32 = 40028
	UIA_UnderlineColorAttributeId int32 = 40029
	UIA_UnderlineStyleAttributeId int32 = 40030
	UIA_AnnotationTypesAttributeId int32 = 40031
	UIA_AnnotationObjectsAttributeId int32 = 40032
	UIA_StyleNameAttributeId int32 = 40033
	UIA_StyleIdAttributeId int32 = 40034
	UIA_LinkAttributeId int32 = 40035
	UIA_IsActiveAttributeId int32 = 40036
	UIA_SelectionActiveEndAttributeId int32 = 40037
	UIA_CaretPositionAttributeId int32 = 40038
	UIA_CaretBidiModeAttributeId int32 = 40039
	UIA_LineSpacingAttributeId int32 = 40040
	UIA_BeforeParagraphSpacingAttributeId int32 = 40041
	UIA_AfterParagraphSpacingAttributeId int32 = 40042
	UIA_SayAsInterpretAsAttributeId int32 = 40043
	UIA_ButtonControlTypeId int32 = 50000
	UIA_CalendarControlTypeId int32 = 50001
	UIA_CheckBoxControlTypeId int32 = 50002
	UIA_ComboBoxControlTypeId int32 = 50003
	UIA_EditControlTypeId int32 = 50004
	UIA_HyperlinkControlTypeId int32 = 50005
	UIA_ImageControlTypeId int32 = 50006
	UIA_ListItemControlTypeId int32 = 50007
	UIA_ListControlTypeId int32 = 50008
	UIA_MenuControlTypeId int32 = 50009
	UIA_MenuBarControlTypeId int32 = 50010
	UIA_MenuItemControlTypeId int32 = 50011
	UIA_ProgressBarControlTypeId int32 = 50012
	UIA_RadioButtonControlTypeId int32 = 50013
	UIA_ScrollBarControlTypeId int32 = 50014
	UIA_SliderControlTypeId int32 = 50015
	UIA_SpinnerControlTypeId int32 = 50016
	UIA_StatusBarControlTypeId int32 = 50017
	UIA_TabControlTypeId int32 = 50018
	UIA_TabItemControlTypeId int32 = 50019
	UIA_TextControlTypeId int32 = 50020
	UIA_ToolBarControlTypeId int32 = 50021
	UIA_ToolTipControlTypeId int32 = 50022
	UIA_TreeControlTypeId int32 = 50023
	UIA_TreeItemControlTypeId int32 = 50024
	UIA_CustomControlTypeId int32 = 50025
	UIA_GroupControlTypeId int32 = 50026
	UIA_ThumbControlTypeId int32 = 50027
	UIA_DataGridControlTypeId int32 = 50028
	UIA_DataItemControlTypeId int32 = 50029
	UIA_DocumentControlTypeId int32 = 50030
	UIA_SplitButtonControlTypeId int32 = 50031
	UIA_WindowControlTypeId int32 = 50032
	UIA_PaneControlTypeId int32 = 50033
	UIA_HeaderControlTypeId int32 = 50034
	UIA_HeaderItemControlTypeId int32 = 50035
	UIA_TableControlTypeId int32 = 50036
	UIA_TitleBarControlTypeId int32 = 50037
	UIA_SeparatorControlTypeId int32 = 50038
	UIA_SemanticZoomControlTypeId int32 = 50039
	UIA_AppBarControlTypeId int32 = 50040
	AnnotationType_Unknown int32 = 60000
	AnnotationType_SpellingError int32 = 60001
	AnnotationType_GrammarError int32 = 60002
	AnnotationType_Comment int32 = 60003
	AnnotationType_FormulaError int32 = 60004
	AnnotationType_TrackChanges int32 = 60005
	AnnotationType_Header int32 = 60006
	AnnotationType_Footer int32 = 60007
	AnnotationType_Highlighted int32 = 60008
	AnnotationType_Endnote int32 = 60009
	AnnotationType_Footnote int32 = 60010
	AnnotationType_InsertionChange int32 = 60011
	AnnotationType_DeletionChange int32 = 60012
	AnnotationType_MoveChange int32 = 60013
	AnnotationType_FormatChange int32 = 60014
	AnnotationType_UnsyncedChange int32 = 60015
	AnnotationType_EditingLockedChange int32 = 60016
	AnnotationType_ExternalChange int32 = 60017
	AnnotationType_ConflictingChange int32 = 60018
	AnnotationType_Author int32 = 60019
	AnnotationType_AdvancedProofingIssue int32 = 60020
	AnnotationType_DataValidationError int32 = 60021
	AnnotationType_CircularReferenceError int32 = 60022
	AnnotationType_Mathematics int32 = 60023
	AnnotationType_Sensitive int32 = 60024
	StyleId_Custom int32 = 70000
	StyleId_Heading1 int32 = 70001
	StyleId_Heading2 int32 = 70002
	StyleId_Heading3 int32 = 70003
	StyleId_Heading4 int32 = 70004
	StyleId_Heading5 int32 = 70005
	StyleId_Heading6 int32 = 70006
	StyleId_Heading7 int32 = 70007
	StyleId_Heading8 int32 = 70008
	StyleId_Heading9 int32 = 70009
	StyleId_Title int32 = 70010
	StyleId_Subtitle int32 = 70011
	StyleId_Normal int32 = 70012
	StyleId_Emphasis int32 = 70013
	StyleId_Quote int32 = 70014
	StyleId_BulletedList int32 = 70015
	StyleId_NumberedList int32 = 70016
	UIA_CustomLandmarkTypeId int32 = 80000
	UIA_FormLandmarkTypeId int32 = 80001
	UIA_MainLandmarkTypeId int32 = 80002
	UIA_NavigationLandmarkTypeId int32 = 80003
	UIA_SearchLandmarkTypeId int32 = 80004
	HeadingLevel_None int32 = 80050
	HeadingLevel1 int32 = 80051
	HeadingLevel2 int32 = 80052
	HeadingLevel3 int32 = 80053
	HeadingLevel4 int32 = 80054
	HeadingLevel5 int32 = 80055
	HeadingLevel6 int32 = 80056
	HeadingLevel7 int32 = 80057
	HeadingLevel8 int32 = 80058
	HeadingLevel9 int32 = 80059
	UIA_SummaryChangeId int32 = 90000
	UIA_SayAsInterpretAsMetadataId int32 = 100000
)

var (
	LIBID_Accessibility syscall.GUID = syscall.GUID{0x1ea4dbf0, 0x3c3b, 0x11cf, 
	[8]byte{0x81, 0x0c, 0x00, 0xaa, 0x00, 0x38, 0x9b, 0x71}}
	CLSID_AccPropServices syscall.GUID = syscall.GUID{0xb5f8350b, 0x0548, 0x48b1, 
	[8]byte{0xa6, 0xee, 0x88, 0xbd, 0x00, 0xb4, 0xa5, 0xe7}}
	IIS_IsOleaccProxy syscall.GUID = syscall.GUID{0x902697fa, 0x80e4, 0x4560, 
	[8]byte{0x80, 0x2a, 0xa1, 0x3f, 0x22, 0xa6, 0x47, 0x09}}
	IIS_ControlAccessible syscall.GUID = syscall.GUID{0x38c682a6, 0x9731, 0x43f2, 
	[8]byte{0x9f, 0xae, 0xe9, 0x01, 0xe6, 0x41, 0xb1, 0x01}}
	PROPID_ACC_NAME syscall.GUID = syscall.GUID{0x608d3df8, 0x8128, 0x4aa7, 
	[8]byte{0xa4, 0x28, 0xf5, 0x5e, 0x49, 0x26, 0x72, 0x91}}
	PROPID_ACC_VALUE syscall.GUID = syscall.GUID{0x123fe443, 0x211a, 0x4615, 
	[8]byte{0x95, 0x27, 0xc4, 0x5a, 0x7e, 0x93, 0x71, 0x7a}}
	PROPID_ACC_DESCRIPTION syscall.GUID = syscall.GUID{0x4d48dfe4, 0xbd3f, 0x491f, 
	[8]byte{0xa6, 0x48, 0x49, 0x2d, 0x6f, 0x20, 0xc5, 0x88}}
	PROPID_ACC_ROLE syscall.GUID = syscall.GUID{0xcb905ff2, 0x7bd1, 0x4c05, 
	[8]byte{0xb3, 0xc8, 0xe6, 0xc2, 0x41, 0x36, 0x4d, 0x70}}
	PROPID_ACC_STATE syscall.GUID = syscall.GUID{0xa8d4d5b0, 0x0a21, 0x42d0, 
	[8]byte{0xa5, 0xc0, 0x51, 0x4e, 0x98, 0x4f, 0x45, 0x7b}}
	PROPID_ACC_HELP syscall.GUID = syscall.GUID{0xc831e11f, 0x44db, 0x4a99, 
	[8]byte{0x97, 0x68, 0xcb, 0x8f, 0x97, 0x8b, 0x72, 0x31}}
	PROPID_ACC_KEYBOARDSHORTCUT syscall.GUID = syscall.GUID{0x7d9bceee, 0x7d1e, 0x4979, 
	[8]byte{0x93, 0x82, 0x51, 0x80, 0xf4, 0x17, 0x2c, 0x34}}
	PROPID_ACC_DEFAULTACTION syscall.GUID = syscall.GUID{0x180c072b, 0xc27f, 0x43c7, 
	[8]byte{0x99, 0x22, 0xf6, 0x35, 0x62, 0xa4, 0x63, 0x2b}}
	PROPID_ACC_HELPTOPIC syscall.GUID = syscall.GUID{0x787d1379, 0x8ede, 0x440b, 
	[8]byte{0x8a, 0xec, 0x11, 0xf7, 0xbf, 0x90, 0x30, 0xb3}}
	PROPID_ACC_FOCUS syscall.GUID = syscall.GUID{0x6eb335df, 0x1c29, 0x4127, 
	[8]byte{0xb1, 0x2c, 0xde, 0xe9, 0xfd, 0x15, 0x7f, 0x2b}}
	PROPID_ACC_SELECTION syscall.GUID = syscall.GUID{0xb99d073c, 0xd731, 0x405b, 
	[8]byte{0x90, 0x61, 0xd9, 0x5e, 0x8f, 0x84, 0x29, 0x84}}
	PROPID_ACC_PARENT syscall.GUID = syscall.GUID{0x474c22b6, 0xffc2, 0x467a, 
	[8]byte{0xb1, 0xb5, 0xe9, 0x58, 0xb4, 0x65, 0x73, 0x30}}
	PROPID_ACC_NAV_UP syscall.GUID = syscall.GUID{0x016e1a2b, 0x1a4e, 0x4767, 
	[8]byte{0x86, 0x12, 0x33, 0x86, 0xf6, 0x69, 0x35, 0xec}}
	PROPID_ACC_NAV_DOWN syscall.GUID = syscall.GUID{0x031670ed, 0x3cdf, 0x48d2, 
	[8]byte{0x96, 0x13, 0x13, 0x8f, 0x2d, 0xd8, 0xa6, 0x68}}
	PROPID_ACC_NAV_LEFT syscall.GUID = syscall.GUID{0x228086cb, 0x82f1, 0x4a39, 
	[8]byte{0x87, 0x05, 0xdc, 0xdc, 0x0f, 0xff, 0x92, 0xf5}}
	PROPID_ACC_NAV_RIGHT syscall.GUID = syscall.GUID{0xcd211d9f, 0xe1cb, 0x4fe5, 
	[8]byte{0xa7, 0x7c, 0x92, 0x0b, 0x88, 0x4d, 0x09, 0x5b}}
	PROPID_ACC_NAV_PREV syscall.GUID = syscall.GUID{0x776d3891, 0xc73b, 0x4480, 
	[8]byte{0xb3, 0xf6, 0x07, 0x6a, 0x16, 0xa1, 0x5a, 0xf6}}
	PROPID_ACC_NAV_NEXT syscall.GUID = syscall.GUID{0x1cdc5455, 0x8cd9, 0x4c92, 
	[8]byte{0xa3, 0x71, 0x39, 0x39, 0xa2, 0xfe, 0x3e, 0xee}}
	PROPID_ACC_NAV_FIRSTCHILD syscall.GUID = syscall.GUID{0xcfd02558, 0x557b, 0x4c67, 
	[8]byte{0x84, 0xf9, 0x2a, 0x09, 0xfc, 0xe4, 0x07, 0x49}}
	PROPID_ACC_NAV_LASTCHILD syscall.GUID = syscall.GUID{0x302ecaa5, 0x48d5, 0x4f8d, 
	[8]byte{0xb6, 0x71, 0x1a, 0x8d, 0x20, 0xa7, 0x78, 0x32}}
	PROPID_ACC_ROLEMAP syscall.GUID = syscall.GUID{0xf79acda2, 0x140d, 0x4fe6, 
	[8]byte{0x89, 0x14, 0x20, 0x84, 0x76, 0x32, 0x82, 0x69}}
	PROPID_ACC_VALUEMAP syscall.GUID = syscall.GUID{0xda1c3d79, 0xfc5c, 0x420e, 
	[8]byte{0xb3, 0x99, 0x9d, 0x15, 0x33, 0x54, 0x9e, 0x75}}
	PROPID_ACC_STATEMAP syscall.GUID = syscall.GUID{0x43946c5e, 0x0ac0, 0x4042, 
	[8]byte{0xb5, 0x25, 0x07, 0xbb, 0xdb, 0xe1, 0x7f, 0xa7}}
	PROPID_ACC_DESCRIPTIONMAP syscall.GUID = syscall.GUID{0x1ff1435f, 0x8a14, 0x477b, 
	[8]byte{0xb2, 0x26, 0xa0, 0xab, 0xe2, 0x79, 0x97, 0x5d}}
	PROPID_ACC_DODEFAULTACTION syscall.GUID = syscall.GUID{0x1ba09523, 0x2e3b, 0x49a6, 
	[8]byte{0xa0, 0x59, 0x59, 0x68, 0x2a, 0x3c, 0x48, 0xfd}}
	RuntimeId_Property_GUID syscall.GUID = syscall.GUID{0xa39eebfa, 0x7fba, 0x4c89, 
	[8]byte{0xb4, 0xd4, 0xb9, 0x9e, 0x2d, 0xe7, 0xd1, 0x60}}
	BoundingRectangle_Property_GUID syscall.GUID = syscall.GUID{0x7bbfe8b2, 0x3bfc, 0x48dd, 
	[8]byte{0xb7, 0x29, 0xc7, 0x94, 0xb8, 0x46, 0xe9, 0xa1}}
	ProcessId_Property_GUID syscall.GUID = syscall.GUID{0x40499998, 0x9c31, 0x4245, 
	[8]byte{0xa4, 0x03, 0x87, 0x32, 0x0e, 0x59, 0xea, 0xf6}}
	ControlType_Property_GUID syscall.GUID = syscall.GUID{0xca774fea, 0x28ac, 0x4bc2, 
	[8]byte{0x94, 0xca, 0xac, 0xec, 0x6d, 0x6c, 0x10, 0xa3}}
	LocalizedControlType_Property_GUID syscall.GUID = syscall.GUID{0x8763404f, 0xa1bd, 0x452a, 
	[8]byte{0x89, 0xc4, 0x3f, 0x01, 0xd3, 0x83, 0x38, 0x06}}
	Name_Property_GUID syscall.GUID = syscall.GUID{0xc3a6921b, 0x4a99, 0x44f1, 
	[8]byte{0xbc, 0xa6, 0x61, 0x18, 0x70, 0x52, 0xc4, 0x31}}
	AcceleratorKey_Property_GUID syscall.GUID = syscall.GUID{0x514865df, 0x2557, 0x4cb9, 
	[8]byte{0xae, 0xed, 0x6c, 0xed, 0x08, 0x4c, 0xe5, 0x2c}}
	AccessKey_Property_GUID syscall.GUID = syscall.GUID{0x06827b12, 0xa7f9, 0x4a15, 
	[8]byte{0x91, 0x7c, 0xff, 0xa5, 0xad, 0x3e, 0xb0, 0xa7}}
	HasKeyboardFocus_Property_GUID syscall.GUID = syscall.GUID{0xcf8afd39, 0x3f46, 0x4800, 
	[8]byte{0x96, 0x56, 0xb2, 0xbf, 0x12, 0x52, 0x99, 0x05}}
	IsKeyboardFocusable_Property_GUID syscall.GUID = syscall.GUID{0xf7b8552a, 0x0859, 0x4b37, 
	[8]byte{0xb9, 0xcb, 0x51, 0xe7, 0x20, 0x92, 0xf2, 0x9f}}
	IsEnabled_Property_GUID syscall.GUID = syscall.GUID{0x2109427f, 0xda60, 0x4fed, 
	[8]byte{0xbf, 0x1b, 0x26, 0x4b, 0xdc, 0xe6, 0xeb, 0x3a}}
	AutomationId_Property_GUID syscall.GUID = syscall.GUID{0xc82c0500, 0xb60e, 0x4310, 
	[8]byte{0xa2, 0x67, 0x30, 0x3c, 0x53, 0x1f, 0x8e, 0xe5}}
	ClassName_Property_GUID syscall.GUID = syscall.GUID{0x157b7215, 0x894f, 0x4b65, 
	[8]byte{0x84, 0xe2, 0xaa, 0xc0, 0xda, 0x08, 0xb1, 0x6b}}
	HelpText_Property_GUID syscall.GUID = syscall.GUID{0x08555685, 0x0977, 0x45c7, 
	[8]byte{0xa7, 0xa6, 0xab, 0xaf, 0x56, 0x84, 0x12, 0x1a}}
	ClickablePoint_Property_GUID syscall.GUID = syscall.GUID{0x0196903b, 0xb203, 0x4818, 
	[8]byte{0xa9, 0xf3, 0xf0, 0x8e, 0x67, 0x5f, 0x23, 0x41}}
	Culture_Property_GUID syscall.GUID = syscall.GUID{0xe2d74f27, 0x3d79, 0x4dc2, 
	[8]byte{0xb8, 0x8b, 0x30, 0x44, 0x96, 0x3a, 0x8a, 0xfb}}
	IsControlElement_Property_GUID syscall.GUID = syscall.GUID{0x95f35085, 0xabcc, 0x4afd, 
	[8]byte{0xa5, 0xf4, 0xdb, 0xb4, 0x6c, 0x23, 0x0f, 0xdb}}
	IsContentElement_Property_GUID syscall.GUID = syscall.GUID{0x4bda64a8, 0xf5d8, 0x480b, 
	[8]byte{0x81, 0x55, 0xef, 0x2e, 0x89, 0xad, 0xb6, 0x72}}
	LabeledBy_Property_GUID syscall.GUID = syscall.GUID{0xe5b8924b, 0xfc8a, 0x4a35, 
	[8]byte{0x80, 0x31, 0xcf, 0x78, 0xac, 0x43, 0xe5, 0x5e}}
	IsPassword_Property_GUID syscall.GUID = syscall.GUID{0xe8482eb1, 0x687c, 0x497b, 
	[8]byte{0xbe, 0xbc, 0x03, 0xbe, 0x53, 0xec, 0x14, 0x54}}
	NewNativeWindowHandle_Property_GUID syscall.GUID = syscall.GUID{0x5196b33b, 0x380a, 0x4982, 
	[8]byte{0x95, 0xe1, 0x91, 0xf3, 0xef, 0x60, 0xe0, 0x24}}
	ItemType_Property_GUID syscall.GUID = syscall.GUID{0xcdda434d, 0x6222, 0x413b, 
	[8]byte{0xa6, 0x8a, 0x32, 0x5d, 0xd1, 0xd4, 0x0f, 0x39}}
	IsOffscreen_Property_GUID syscall.GUID = syscall.GUID{0x03c3d160, 0xdb79, 0x42db, 
	[8]byte{0xa2, 0xef, 0x1c, 0x23, 0x1e, 0xed, 0xe5, 0x07}}
	Orientation_Property_GUID syscall.GUID = syscall.GUID{0xa01eee62, 0x3884, 0x4415, 
	[8]byte{0x88, 0x7e, 0x67, 0x8e, 0xc2, 0x1e, 0x39, 0xba}}
	FrameworkId_Property_GUID syscall.GUID = syscall.GUID{0xdbfd9900, 0x7e1a, 0x4f58, 
	[8]byte{0xb6, 0x1b, 0x70, 0x63, 0x12, 0x0f, 0x77, 0x3b}}
	IsRequiredForForm_Property_GUID syscall.GUID = syscall.GUID{0x4f5f43cf, 0x59fb, 0x4bde, 
	[8]byte{0xa2, 0x70, 0x60, 0x2e, 0x5e, 0x11, 0x41, 0xe9}}
	ItemStatus_Property_GUID syscall.GUID = syscall.GUID{0x51de0321, 0x3973, 0x43e7, 
	[8]byte{0x89, 0x13, 0x0b, 0x08, 0xe8, 0x13, 0xc3, 0x7f}}
	AriaRole_Property_GUID syscall.GUID = syscall.GUID{0xdd207b95, 0xbe4a, 0x4e0d, 
	[8]byte{0xb7, 0x27, 0x63, 0xac, 0xe9, 0x4b, 0x69, 0x16}}
	AriaProperties_Property_GUID syscall.GUID = syscall.GUID{0x4213678c, 0xe025, 0x4922, 
	[8]byte{0xbe, 0xb5, 0xe4, 0x3b, 0xa0, 0x8e, 0x62, 0x21}}
	IsDataValidForForm_Property_GUID syscall.GUID = syscall.GUID{0x445ac684, 0xc3fc, 0x4dd9, 
	[8]byte{0xac, 0xf8, 0x84, 0x5a, 0x57, 0x92, 0x96, 0xba}}
	ControllerFor_Property_GUID syscall.GUID = syscall.GUID{0x51124c8a, 0xa5d2, 0x4f13, 
	[8]byte{0x9b, 0xe6, 0x7f, 0xa8, 0xba, 0x9d, 0x3a, 0x90}}
	DescribedBy_Property_GUID syscall.GUID = syscall.GUID{0x7c5865b8, 0x9992, 0x40fd, 
	[8]byte{0x8d, 0xb0, 0x6b, 0xf1, 0xd3, 0x17, 0xf9, 0x98}}
	FlowsTo_Property_GUID syscall.GUID = syscall.GUID{0xe4f33d20, 0x559a, 0x47fb, 
	[8]byte{0xa8, 0x30, 0xf9, 0xcb, 0x4f, 0xf1, 0xa7, 0x0a}}
	ProviderDescription_Property_GUID syscall.GUID = syscall.GUID{0xdca5708a, 0xc16b, 0x4cd9, 
	[8]byte{0xb8, 0x89, 0xbe, 0xb1, 0x6a, 0x80, 0x49, 0x04}}
	OptimizeForVisualContent_Property_GUID syscall.GUID = syscall.GUID{0x6a852250, 0xc75a, 0x4e5d, 
	[8]byte{0xb8, 0x58, 0xe3, 0x81, 0xb0, 0xf7, 0x88, 0x61}}
	IsDockPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x2600a4c4, 0x2ff8, 0x4c96, 
	[8]byte{0xae, 0x31, 0x8f, 0xe6, 0x19, 0xa1, 0x3c, 0x6c}}
	IsExpandCollapsePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x929d3806, 0x5287, 0x4725, 
	[8]byte{0xaa, 0x16, 0x22, 0x2a, 0xfc, 0x63, 0xd5, 0x95}}
	IsGridItemPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x5a43e524, 0xf9a2, 0x4b12, 
	[8]byte{0x84, 0xc8, 0xb4, 0x8a, 0x3e, 0xfe, 0xdd, 0x34}}
	IsGridPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x5622c26c, 0xf0ef, 0x4f3b, 
	[8]byte{0x97, 0xcb, 0x71, 0x4c, 0x08, 0x68, 0x58, 0x8b}}
	IsInvokePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x4e725738, 0x8364, 0x4679, 
	[8]byte{0xaa, 0x6c, 0xf3, 0xf4, 0x19, 0x31, 0xf7, 0x50}}
	IsMultipleViewPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xff0a31eb, 0x8e25, 0x469d, 
	[8]byte{0x8d, 0x6e, 0xe7, 0x71, 0xa2, 0x7c, 0x1b, 0x90}}
	IsRangeValuePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xfda4244a, 0xeb4d, 0x43ff, 
	[8]byte{0xb5, 0xad, 0xed, 0x36, 0xd3, 0x73, 0xec, 0x4c}}
	IsScrollPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x3ebb7b4a, 0x828a, 0x4b57, 
	[8]byte{0x9d, 0x22, 0x2f, 0xea, 0x16, 0x32, 0xed, 0x0d}}
	IsScrollItemPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x1cad1a05, 0x0927, 0x4b76, 
	[8]byte{0x97, 0xe1, 0x0f, 0xcd, 0xb2, 0x09, 0xb9, 0x8a}}
	IsSelectionItemPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x8becd62d, 0x0bc3, 0x4109, 
	[8]byte{0xbe, 0xe2, 0x8e, 0x67, 0x15, 0x29, 0x0e, 0x68}}
	IsSelectionPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xf588acbe, 0xc769, 0x4838, 
	[8]byte{0x9a, 0x60, 0x26, 0x86, 0xdc, 0x11, 0x88, 0xc4}}
	IsTablePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xcb83575f, 0x45c2, 0x4048, 
	[8]byte{0x9c, 0x76, 0x15, 0x97, 0x15, 0xa1, 0x39, 0xdf}}
	IsTableItemPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xeb36b40d, 0x8ea4, 0x489b, 
	[8]byte{0xa0, 0x13, 0xe6, 0x0d, 0x59, 0x51, 0xfe, 0x34}}
	IsTextPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xfbe2d69d, 0xaff6, 0x4a45, 
	[8]byte{0x82, 0xe2, 0xfc, 0x92, 0xa8, 0x2f, 0x59, 0x17}}
	IsTogglePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x78686d53, 0xfcd0, 0x4b83, 
	[8]byte{0x9b, 0x78, 0x58, 0x32, 0xce, 0x63, 0xbb, 0x5b}}
	IsTransformPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xa7f78804, 0xd68b, 0x4077, 
	[8]byte{0xa5, 0xc6, 0x7a, 0x5e, 0xa1, 0xac, 0x31, 0xc5}}
	IsValuePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x0b5020a7, 0x2119, 0x473b, 
	[8]byte{0xbe, 0x37, 0x5c, 0xeb, 0x98, 0xbb, 0xfb, 0x22}}
	IsWindowPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xe7a57bb1, 0x5888, 0x4155, 
	[8]byte{0x98, 0xdc, 0xb4, 0x22, 0xfd, 0x57, 0xf2, 0xbc}}
	IsLegacyIAccessiblePatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xd8ebd0c7, 0x929a, 0x4ee7, 
	[8]byte{0x8d, 0x3a, 0xd3, 0xd9, 0x44, 0x13, 0x02, 0x7b}}
	IsItemContainerPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x624b5ca7, 0xfe40, 0x4957, 
	[8]byte{0xa0, 0x19, 0x20, 0xc4, 0xcf, 0x11, 0x92, 0x0f}}
	IsVirtualizedItemPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x302cb151, 0x2ac8, 0x45d6, 
	[8]byte{0x97, 0x7b, 0xd2, 0xb3, 0xa5, 0xa5, 0x3f, 0x20}}
	IsSynchronizedInputPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x75d69cc5, 0xd2bf, 0x4943, 
	[8]byte{0x87, 0x6e, 0xb4, 0x5b, 0x62, 0xa6, 0xcc, 0x66}}
	IsObjectModelPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x6b21d89b, 0x2841, 0x412f, 
	[8]byte{0x8e, 0xf2, 0x15, 0xca, 0x95, 0x23, 0x18, 0xba}}
	IsAnnotationPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x0b5b3238, 0x6d5c, 0x41b6, 
	[8]byte{0xbc, 0xc4, 0x5e, 0x80, 0x7f, 0x65, 0x51, 0xc4}}
	IsTextPattern2Available_Property_GUID syscall.GUID = syscall.GUID{0x41cf921d, 0xe3f1, 0x4b22, 
	[8]byte{0x9c, 0x81, 0xe1, 0xc3, 0xed, 0x33, 0x1c, 0x22}}
	IsTextEditPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x7843425c, 0x8b32, 0x484c, 
	[8]byte{0x9a, 0xb5, 0xe3, 0x20, 0x05, 0x71, 0xff, 0xda}}
	IsCustomNavigationPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x8f8e80d4, 0x2351, 0x48e0, 
	[8]byte{0x87, 0x4a, 0x54, 0xaa, 0x73, 0x13, 0x88, 0x9a}}
	IsStylesPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x27f353d3, 0x459c, 0x4b59, 
	[8]byte{0xa4, 0x90, 0x50, 0x61, 0x1d, 0xac, 0xaf, 0xb5}}
	IsSpreadsheetPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x6ff43732, 0xe4b4, 0x4555, 
	[8]byte{0x97, 0xbc, 0xec, 0xdb, 0xbc, 0x4d, 0x18, 0x88}}
	IsSpreadsheetItemPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x9fe79b2a, 0x2f94, 0x43fd, 
	[8]byte{0x99, 0x6b, 0x54, 0x9e, 0x31, 0x6f, 0x4a, 0xcd}}
	IsTransformPattern2Available_Property_GUID syscall.GUID = syscall.GUID{0x25980b4b, 0xbe04, 0x4710, 
	[8]byte{0xab, 0x4a, 0xfd, 0xa3, 0x1d, 0xbd, 0x28, 0x95}}
	IsTextChildPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x559e65df, 0x30ff, 0x43b5, 
	[8]byte{0xb5, 0xed, 0x5b, 0x28, 0x3b, 0x80, 0xc7, 0xe9}}
	IsDragPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xe997a7b7, 0x1d39, 0x4ca7, 
	[8]byte{0xbe, 0x0f, 0x27, 0x7f, 0xcf, 0x56, 0x05, 0xcc}}
	IsDropTargetPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0x0686b62e, 0x8e19, 0x4aaf, 
	[8]byte{0x87, 0x3d, 0x38, 0x4f, 0x6d, 0x3b, 0x92, 0xbe}}
	IsStructuredMarkupPatternAvailable_Property_GUID syscall.GUID = syscall.GUID{0xb0d4c196, 0x2c0b, 0x489c, 
	[8]byte{0xb1, 0x65, 0xa4, 0x05, 0x92, 0x8c, 0x6f, 0x3d}}
	IsPeripheral_Property_GUID syscall.GUID = syscall.GUID{0xda758276, 0x7ed5, 0x49d4, 
	[8]byte{0x8e, 0x68, 0xec, 0xc9, 0xa2, 0xd3, 0x00, 0xdd}}
	PositionInSet_Property_GUID syscall.GUID = syscall.GUID{0x33d1dc54, 0x641e, 0x4d76, 
	[8]byte{0xa6, 0xb1, 0x13, 0xf3, 0x41, 0xc1, 0xf8, 0x96}}
	SizeOfSet_Property_GUID syscall.GUID = syscall.GUID{0x1600d33c, 0x3b9f, 0x4369, 
	[8]byte{0x94, 0x31, 0xaa, 0x29, 0x3f, 0x34, 0x4c, 0xf1}}
	Level_Property_GUID syscall.GUID = syscall.GUID{0x242ac529, 0xcd36, 0x400f, 
	[8]byte{0xaa, 0xd9, 0x78, 0x76, 0xef, 0x3a, 0xf6, 0x27}}
	AnnotationTypes_Property_GUID syscall.GUID = syscall.GUID{0x64b71f76, 0x53c4, 0x4696, 
	[8]byte{0xa2, 0x19, 0x20, 0xe9, 0x40, 0xc9, 0xa1, 0x76}}
	AnnotationObjects_Property_GUID syscall.GUID = syscall.GUID{0x310910c8, 0x7c6e, 0x4f20, 
	[8]byte{0xbe, 0xcd, 0x4a, 0xaf, 0x6d, 0x19, 0x11, 0x56}}
	LandmarkType_Property_GUID syscall.GUID = syscall.GUID{0x454045f2, 0x6f61, 0x49f7, 
	[8]byte{0xa4, 0xf8, 0xb5, 0xf0, 0xcf, 0x82, 0xda, 0x1e}}
	LocalizedLandmarkType_Property_GUID syscall.GUID = syscall.GUID{0x7ac81980, 0xeafb, 0x4fb2, 
	[8]byte{0xbf, 0x91, 0xf4, 0x85, 0xbe, 0xf5, 0xe8, 0xe1}}
	FullDescription_Property_GUID syscall.GUID = syscall.GUID{0x0d4450ff, 0x6aef, 0x4f33, 
	[8]byte{0x95, 0xdd, 0x7b, 0xef, 0xa7, 0x2a, 0x43, 0x91}}
	Value_Value_Property_GUID syscall.GUID = syscall.GUID{0xe95f5e64, 0x269f, 0x4a85, 
	[8]byte{0xba, 0x99, 0x40, 0x92, 0xc3, 0xea, 0x29, 0x86}}
	Value_IsReadOnly_Property_GUID syscall.GUID = syscall.GUID{0xeb090f30, 0xe24c, 0x4799, 
	[8]byte{0xa7, 0x05, 0x0d, 0x24, 0x7b, 0xc0, 0x37, 0xf8}}
	RangeValue_Value_Property_GUID syscall.GUID = syscall.GUID{0x131f5d98, 0xc50c, 0x489d, 
	[8]byte{0xab, 0xe5, 0xae, 0x22, 0x08, 0x98, 0xc5, 0xf7}}
	RangeValue_IsReadOnly_Property_GUID syscall.GUID = syscall.GUID{0x25fa1055, 0xdebf, 0x4373, 
	[8]byte{0xa7, 0x9e, 0x1f, 0x1a, 0x19, 0x08, 0xd3, 0xc4}}
	RangeValue_Minimum_Property_GUID syscall.GUID = syscall.GUID{0x78cbd3b2, 0x684d, 0x4860, 
	[8]byte{0xaf, 0x93, 0xd1, 0xf9, 0x5c, 0xb0, 0x22, 0xfd}}
	RangeValue_Maximum_Property_GUID syscall.GUID = syscall.GUID{0x19319914, 0xf979, 0x4b35, 
	[8]byte{0xa1, 0xa6, 0xd3, 0x7e, 0x05, 0x43, 0x34, 0x73}}
	RangeValue_LargeChange_Property_GUID syscall.GUID = syscall.GUID{0xa1f96325, 0x3a3d, 0x4b44, 
	[8]byte{0x8e, 0x1f, 0x4a, 0x46, 0xd9, 0x84, 0x40, 0x19}}
	RangeValue_SmallChange_Property_GUID syscall.GUID = syscall.GUID{0x81c2c457, 0x3941, 0x4107, 
	[8]byte{0x99, 0x75, 0x13, 0x97, 0x60, 0xf7, 0xc0, 0x72}}
	Scroll_HorizontalScrollPercent_Property_GUID syscall.GUID = syscall.GUID{0xc7c13c0e, 0xeb21, 0x47ff, 
	[8]byte{0xac, 0xc4, 0xb5, 0xa3, 0x35, 0x0f, 0x51, 0x91}}
	Scroll_HorizontalViewSize_Property_GUID syscall.GUID = syscall.GUID{0x70c2e5d4, 0xfcb0, 0x4713, 
	[8]byte{0xa9, 0xaa, 0xaf, 0x92, 0xff, 0x79, 0xe4, 0xcd}}
	Scroll_VerticalScrollPercent_Property_GUID syscall.GUID = syscall.GUID{0x6c8d7099, 0xb2a8, 0x4948, 
	[8]byte{0xbf, 0xf7, 0x3c, 0xf9, 0x05, 0x8b, 0xfe, 0xfb}}
	Scroll_VerticalViewSize_Property_GUID syscall.GUID = syscall.GUID{0xde6a2e22, 0xd8c7, 0x40c5, 
	[8]byte{0x83, 0xba, 0xe5, 0xf6, 0x81, 0xd5, 0x31, 0x08}}
	Scroll_HorizontallyScrollable_Property_GUID syscall.GUID = syscall.GUID{0x8b925147, 0x28cd, 0x49ae, 
	[8]byte{0xbd, 0x63, 0xf4, 0x41, 0x18, 0xd2, 0xe7, 0x19}}
	Scroll_VerticallyScrollable_Property_GUID syscall.GUID = syscall.GUID{0x89164798, 0x0068, 0x4315, 
	[8]byte{0xb8, 0x9a, 0x1e, 0x7c, 0xfb, 0xbc, 0x3d, 0xfc}}
	Selection_Selection_Property_GUID syscall.GUID = syscall.GUID{0xaa6dc2a2, 0x0e2b, 0x4d38, 
	[8]byte{0x96, 0xd5, 0x34, 0xe4, 0x70, 0xb8, 0x18, 0x53}}
	Selection_CanSelectMultiple_Property_GUID syscall.GUID = syscall.GUID{0x49d73da5, 0xc883, 0x4500, 
	[8]byte{0x88, 0x3d, 0x8f, 0xcf, 0x8d, 0xaf, 0x6c, 0xbe}}
	Selection_IsSelectionRequired_Property_GUID syscall.GUID = syscall.GUID{0xb1ae4422, 0x63fe, 0x44e7, 
	[8]byte{0xa5, 0xa5, 0xa7, 0x38, 0xc8, 0x29, 0xb1, 0x9a}}
	Grid_RowCount_Property_GUID syscall.GUID = syscall.GUID{0x2a9505bf, 0xc2eb, 0x4fb6, 
	[8]byte{0xb3, 0x56, 0x82, 0x45, 0xae, 0x53, 0x70, 0x3e}}
	Grid_ColumnCount_Property_GUID syscall.GUID = syscall.GUID{0xfe96f375, 0x44aa, 0x4536, 
	[8]byte{0xac, 0x7a, 0x2a, 0x75, 0xd7, 0x1a, 0x3e, 0xfc}}
	GridItem_Row_Property_GUID syscall.GUID = syscall.GUID{0x6223972a, 0xc945, 0x4563, 
	[8]byte{0x93, 0x29, 0xfd, 0xc9, 0x74, 0xaf, 0x25, 0x53}}
	GridItem_Column_Property_GUID syscall.GUID = syscall.GUID{0xc774c15c, 0x62c0, 0x4519, 
	[8]byte{0x8b, 0xdc, 0x47, 0xbe, 0x57, 0x3c, 0x8a, 0xd5}}
	GridItem_RowSpan_Property_GUID syscall.GUID = syscall.GUID{0x4582291c, 0x466b, 0x4e93, 
	[8]byte{0x8e, 0x83, 0x3d, 0x17, 0x15, 0xec, 0x0c, 0x5e}}
	GridItem_ColumnSpan_Property_GUID syscall.GUID = syscall.GUID{0x583ea3f5, 0x86d0, 0x4b08, 
	[8]byte{0xa6, 0xec, 0x2c, 0x54, 0x63, 0xff, 0xc1, 0x09}}
	GridItem_Parent_Property_GUID syscall.GUID = syscall.GUID{0x9d912252, 0xb97f, 0x4ecc, 
	[8]byte{0x85, 0x10, 0xea, 0x0e, 0x33, 0x42, 0x7c, 0x72}}
	Dock_DockPosition_Property_GUID syscall.GUID = syscall.GUID{0x6d67f02e, 0xc0b0, 0x4b10, 
	[8]byte{0xb5, 0xb9, 0x18, 0xd6, 0xec, 0xf9, 0x87, 0x60}}
	ExpandCollapse_ExpandCollapseState_Property_GUID syscall.GUID = syscall.GUID{0x275a4c48, 0x85a7, 0x4f69, 
	[8]byte{0xab, 0xa0, 0xaf, 0x15, 0x76, 0x10, 0x00, 0x2b}}
	MultipleView_CurrentView_Property_GUID syscall.GUID = syscall.GUID{0x7a81a67a, 0xb94f, 0x4875, 
	[8]byte{0x91, 0x8b, 0x65, 0xc8, 0xd2, 0xf9, 0x98, 0xe5}}
	MultipleView_SupportedViews_Property_GUID syscall.GUID = syscall.GUID{0x8d5db9fd, 0xce3c, 0x4ae7, 
	[8]byte{0xb7, 0x88, 0x40, 0x0a, 0x3c, 0x64, 0x55, 0x47}}
	Window_CanMaximize_Property_GUID syscall.GUID = syscall.GUID{0x64fff53f, 0x635d, 0x41c1, 
	[8]byte{0x95, 0x0c, 0xcb, 0x5a, 0xdf, 0xbe, 0x28, 0xe3}}
	Window_CanMinimize_Property_GUID syscall.GUID = syscall.GUID{0xb73b4625, 0x5988, 0x4b97, 
	[8]byte{0xb4, 0xc2, 0xa6, 0xfe, 0x6e, 0x78, 0xc8, 0xc6}}
	Window_WindowVisualState_Property_GUID syscall.GUID = syscall.GUID{0x4ab7905f, 0xe860, 0x453e, 
	[8]byte{0xa3, 0x0a, 0xf6, 0x43, 0x1e, 0x5d, 0xaa, 0xd5}}
	Window_WindowInteractionState_Property_GUID syscall.GUID = syscall.GUID{0x4fed26a4, 0x0455, 0x4fa2, 
	[8]byte{0xb2, 0x1c, 0xc4, 0xda, 0x2d, 0xb1, 0xff, 0x9c}}
	Window_IsModal_Property_GUID syscall.GUID = syscall.GUID{0xff4e6892, 0x37b9, 0x4fca, 
	[8]byte{0x85, 0x32, 0xff, 0xe6, 0x74, 0xec, 0xfe, 0xed}}
	Window_IsTopmost_Property_GUID syscall.GUID = syscall.GUID{0xef7d85d3, 0x0937, 0x4962, 
	[8]byte{0x92, 0x41, 0xb6, 0x23, 0x45, 0xf2, 0x40, 0x41}}
	SelectionItem_IsSelected_Property_GUID syscall.GUID = syscall.GUID{0xf122835f, 0xcd5f, 0x43df, 
	[8]byte{0xb7, 0x9d, 0x4b, 0x84, 0x9e, 0x9e, 0x60, 0x20}}
	SelectionItem_SelectionContainer_Property_GUID syscall.GUID = syscall.GUID{0xa4365b6e, 0x9c1e, 0x4b63, 
	[8]byte{0x8b, 0x53, 0xc2, 0x42, 0x1d, 0xd1, 0xe8, 0xfb}}
	Table_RowHeaders_Property_GUID syscall.GUID = syscall.GUID{0xd9e35b87, 0x6eb8, 0x4562, 
	[8]byte{0xaa, 0xc6, 0xa8, 0xa9, 0x07, 0x52, 0x36, 0xa8}}
	Table_ColumnHeaders_Property_GUID syscall.GUID = syscall.GUID{0xaff1d72b, 0x968d, 0x42b1, 
	[8]byte{0xb4, 0x59, 0x15, 0x0b, 0x29, 0x9d, 0xa6, 0x64}}
	Table_RowOrColumnMajor_Property_GUID syscall.GUID = syscall.GUID{0x83be75c3, 0x29fe, 0x4a30, 
	[8]byte{0x85, 0xe1, 0x2a, 0x62, 0x77, 0xfd, 0x10, 0x6e}}
	TableItem_RowHeaderItems_Property_GUID syscall.GUID = syscall.GUID{0xb3f853a0, 0x0574, 0x4cd8, 
	[8]byte{0xbc, 0xd7, 0xed, 0x59, 0x23, 0x57, 0x2d, 0x97}}
	TableItem_ColumnHeaderItems_Property_GUID syscall.GUID = syscall.GUID{0x967a56a3, 0x74b6, 0x431e, 
	[8]byte{0x8d, 0xe6, 0x99, 0xc4, 0x11, 0x03, 0x1c, 0x58}}
	Toggle_ToggleState_Property_GUID syscall.GUID = syscall.GUID{0xb23cdc52, 0x22c2, 0x4c6c, 
	[8]byte{0x9d, 0xed, 0xf5, 0xc4, 0x22, 0x47, 0x9e, 0xde}}
	Transform_CanMove_Property_GUID syscall.GUID = syscall.GUID{0x1b75824d, 0x208b, 0x4fdf, 
	[8]byte{0xbc, 0xcd, 0xf1, 0xf4, 0xe5, 0x74, 0x1f, 0x4f}}
	Transform_CanResize_Property_GUID syscall.GUID = syscall.GUID{0xbb98dca5, 0x4c1a, 0x41d4, 
	[8]byte{0xa4, 0xf6, 0xeb, 0xc1, 0x28, 0x64, 0x41, 0x80}}
	Transform_CanRotate_Property_GUID syscall.GUID = syscall.GUID{0x10079b48, 0x3849, 0x476f, 
	[8]byte{0xac, 0x96, 0x44, 0xa9, 0x5c, 0x84, 0x40, 0xd9}}
	LegacyIAccessible_ChildId_Property_GUID syscall.GUID = syscall.GUID{0x9a191b5d, 0x9ef2, 0x4787, 
	[8]byte{0xa4, 0x59, 0xdc, 0xde, 0x88, 0x5d, 0xd4, 0xe8}}
	LegacyIAccessible_Name_Property_GUID syscall.GUID = syscall.GUID{0xcaeb063d, 0x40ae, 0x4869, 
	[8]byte{0xaa, 0x5a, 0x1b, 0x8e, 0x5d, 0x66, 0x67, 0x39}}
	LegacyIAccessible_Value_Property_GUID syscall.GUID = syscall.GUID{0xb5c5b0b6, 0x8217, 0x4a77, 
	[8]byte{0x97, 0xa5, 0x19, 0x0a, 0x85, 0xed, 0x01, 0x56}}
	LegacyIAccessible_Description_Property_GUID syscall.GUID = syscall.GUID{0x46448418, 0x7d70, 0x4ea9, 
	[8]byte{0x9d, 0x27, 0xb7, 0xe7, 0x75, 0xcf, 0x2a, 0xd7}}
	LegacyIAccessible_Role_Property_GUID syscall.GUID = syscall.GUID{0x6856e59f, 0xcbaf, 0x4e31, 
	[8]byte{0x93, 0xe8, 0xbc, 0xbf, 0x6f, 0x7e, 0x49, 0x1c}}
	LegacyIAccessible_State_Property_GUID syscall.GUID = syscall.GUID{0xdf985854, 0x2281, 0x4340, 
	[8]byte{0xab, 0x9c, 0xc6, 0x0e, 0x2c, 0x58, 0x03, 0xf6}}
	LegacyIAccessible_Help_Property_GUID syscall.GUID = syscall.GUID{0x94402352, 0x161c, 0x4b77, 
	[8]byte{0xa9, 0x8d, 0xa8, 0x72, 0xcc, 0x33, 0x94, 0x7a}}
	LegacyIAccessible_KeyboardShortcut_Property_GUID syscall.GUID = syscall.GUID{0x8f6909ac, 0x00b8, 0x4259, 
	[8]byte{0xa4, 0x1c, 0x96, 0x62, 0x66, 0xd4, 0x3a, 0x8a}}
	LegacyIAccessible_Selection_Property_GUID syscall.GUID = syscall.GUID{0x8aa8b1e0, 0x0891, 0x40cc, 
	[8]byte{0x8b, 0x06, 0x90, 0xd7, 0xd4, 0x16, 0x62, 0x19}}
	LegacyIAccessible_DefaultAction_Property_GUID syscall.GUID = syscall.GUID{0x3b331729, 0xeaad, 0x4502, 
	[8]byte{0xb8, 0x5f, 0x92, 0x61, 0x56, 0x22, 0x91, 0x3c}}
	Annotation_AnnotationTypeId_Property_GUID syscall.GUID = syscall.GUID{0x20ae484f, 0x69ef, 0x4c48, 
	[8]byte{0x8f, 0x5b, 0xc4, 0x93, 0x8b, 0x20, 0x6a, 0xc7}}
	Annotation_AnnotationTypeName_Property_GUID syscall.GUID = syscall.GUID{0x9b818892, 0x5ac9, 0x4af9, 
	[8]byte{0xaa, 0x96, 0xf5, 0x8a, 0x77, 0xb0, 0x58, 0xe3}}
	Annotation_Author_Property_GUID syscall.GUID = syscall.GUID{0x7a528462, 0x9c5c, 0x4a03, 
	[8]byte{0xa9, 0x74, 0x8b, 0x30, 0x7a, 0x99, 0x37, 0xf2}}
	Annotation_DateTime_Property_GUID syscall.GUID = syscall.GUID{0x99b5ca5d, 0x1acf, 0x414b, 
	[8]byte{0xa4, 0xd0, 0x6b, 0x35, 0x0b, 0x04, 0x75, 0x78}}
	Annotation_Target_Property_GUID syscall.GUID = syscall.GUID{0xb71b302d, 0x2104, 0x44ad, 
	[8]byte{0x9c, 0x5c, 0x09, 0x2b, 0x49, 0x07, 0xd7, 0x0f}}
	Styles_StyleId_Property_GUID syscall.GUID = syscall.GUID{0xda82852f, 0x3817, 0x4233, 
	[8]byte{0x82, 0xaf, 0x02, 0x27, 0x9e, 0x72, 0xcc, 0x77}}
	Styles_StyleName_Property_GUID syscall.GUID = syscall.GUID{0x1c12b035, 0x05d1, 0x4f55, 
	[8]byte{0x9e, 0x8e, 0x14, 0x89, 0xf3, 0xff, 0x55, 0x0d}}
	Styles_FillColor_Property_GUID syscall.GUID = syscall.GUID{0x63eff97a, 0xa1c5, 0x4b1d, 
	[8]byte{0x84, 0xeb, 0xb7, 0x65, 0xf2, 0xed, 0xd6, 0x32}}
	Styles_FillPatternStyle_Property_GUID syscall.GUID = syscall.GUID{0x81cf651f, 0x482b, 0x4451, 
	[8]byte{0xa3, 0x0a, 0xe1, 0x54, 0x5e, 0x55, 0x4f, 0xb8}}
	Styles_Shape_Property_GUID syscall.GUID = syscall.GUID{0xc71a23f8, 0x778c, 0x400d, 
	[8]byte{0x84, 0x58, 0x3b, 0x54, 0x3e, 0x52, 0x69, 0x84}}
	Styles_FillPatternColor_Property_GUID syscall.GUID = syscall.GUID{0x939a59fe, 0x8fbd, 0x4e75, 
	[8]byte{0xa2, 0x71, 0xac, 0x45, 0x95, 0x19, 0x51, 0x63}}
	Styles_ExtendedProperties_Property_GUID syscall.GUID = syscall.GUID{0xf451cda0, 0xba0a, 0x4681, 
	[8]byte{0xb0, 0xb0, 0x0d, 0xbd, 0xb5, 0x3e, 0x58, 0xf3}}
	SpreadsheetItem_Formula_Property_GUID syscall.GUID = syscall.GUID{0xe602e47d, 0x1b47, 0x4bea, 
	[8]byte{0x87, 0xcf, 0x3b, 0x0b, 0x0b, 0x5c, 0x15, 0xb6}}
	SpreadsheetItem_AnnotationObjects_Property_GUID syscall.GUID = syscall.GUID{0xa3194c38, 0xc9bc, 0x4604, 
	[8]byte{0x93, 0x96, 0xae, 0x3f, 0x9f, 0x45, 0x7f, 0x7b}}
	SpreadsheetItem_AnnotationTypes_Property_GUID syscall.GUID = syscall.GUID{0xc70c51d0, 0xd602, 0x4b45, 
	[8]byte{0xaf, 0xbc, 0xb4, 0x71, 0x2b, 0x96, 0xd7, 0x2b}}
	Transform2_CanZoom_Property_GUID syscall.GUID = syscall.GUID{0xf357e890, 0xa756, 0x4359, 
	[8]byte{0x9c, 0xa6, 0x86, 0x70, 0x2b, 0xf8, 0xf3, 0x81}}
	LiveSetting_Property_GUID syscall.GUID = syscall.GUID{0xc12bcd8e, 0x2a8e, 0x4950, 
	[8]byte{0x8a, 0xe7, 0x36, 0x25, 0x11, 0x1d, 0x58, 0xeb}}
	Drag_IsGrabbed_Property_GUID syscall.GUID = syscall.GUID{0x45f206f3, 0x75cc, 0x4cca, 
	[8]byte{0xa9, 0xb9, 0xfc, 0xdf, 0xb9, 0x82, 0xd8, 0xa2}}
	Drag_GrabbedItems_Property_GUID syscall.GUID = syscall.GUID{0x77c1562c, 0x7b86, 0x4b21, 
	[8]byte{0x9e, 0xd7, 0x3c, 0xef, 0xda, 0x6f, 0x4c, 0x43}}
	Drag_DropEffect_Property_GUID syscall.GUID = syscall.GUID{0x646f2779, 0x48d3, 0x4b23, 
	[8]byte{0x89, 0x02, 0x4b, 0xf1, 0x00, 0x00, 0x5d, 0xf3}}
	Drag_DropEffects_Property_GUID syscall.GUID = syscall.GUID{0xf5d61156, 0x7ce6, 0x49be, 
	[8]byte{0xa8, 0x36, 0x92, 0x69, 0xdc, 0xec, 0x92, 0x0f}}
	DropTarget_DropTargetEffect_Property_GUID syscall.GUID = syscall.GUID{0x8bb75975, 0xa0ca, 0x4981, 
	[8]byte{0xb8, 0x18, 0x87, 0xfc, 0x66, 0xe9, 0x50, 0x9d}}
	DropTarget_DropTargetEffects_Property_GUID syscall.GUID = syscall.GUID{0xbc1dd4ed, 0xcb89, 0x45f1, 
	[8]byte{0xa5, 0x92, 0xe0, 0x3b, 0x08, 0xae, 0x79, 0x0f}}
	Transform2_ZoomLevel_Property_GUID syscall.GUID = syscall.GUID{0xeee29f1a, 0xf4a2, 0x4b5b, 
	[8]byte{0xac, 0x65, 0x95, 0xcf, 0x93, 0x28, 0x33, 0x87}}
	Transform2_ZoomMinimum_Property_GUID syscall.GUID = syscall.GUID{0x742ccc16, 0x4ad1, 0x4e07, 
	[8]byte{0x96, 0xfe, 0xb1, 0x22, 0xc6, 0xe6, 0xb2, 0x2b}}
	Transform2_ZoomMaximum_Property_GUID syscall.GUID = syscall.GUID{0x42ab6b77, 0xceb0, 0x4eca, 
	[8]byte{0xb8, 0x2a, 0x6c, 0xfa, 0x5f, 0xa1, 0xfc, 0x08}}
	FlowsFrom_Property_GUID syscall.GUID = syscall.GUID{0x05c6844f, 0x19de, 0x48f8, 
	[8]byte{0x95, 0xfa, 0x88, 0x0d, 0x5b, 0x0f, 0xd6, 0x15}}
	FillColor_Property_GUID syscall.GUID = syscall.GUID{0x6e0ec4d0, 0xe2a8, 0x4a56, 
	[8]byte{0x9d, 0xe7, 0x95, 0x33, 0x89, 0x93, 0x3b, 0x39}}
	OutlineColor_Property_GUID syscall.GUID = syscall.GUID{0xc395d6c0, 0x4b55, 0x4762, 
	[8]byte{0xa0, 0x73, 0xfd, 0x30, 0x3a, 0x63, 0x4f, 0x52}}
	FillType_Property_GUID syscall.GUID = syscall.GUID{0xc6fc74e4, 0x8cb9, 0x429c, 
	[8]byte{0xa9, 0xe1, 0x9b, 0xc4, 0xac, 0x37, 0x2b, 0x62}}
	VisualEffects_Property_GUID syscall.GUID = syscall.GUID{0xe61a8565, 0xaad9, 0x46d7, 
	[8]byte{0x9e, 0x70, 0x4e, 0x8a, 0x84, 0x20, 0xd4, 0x20}}
	OutlineThickness_Property_GUID syscall.GUID = syscall.GUID{0x13e67cc7, 0xdac2, 0x4888, 
	[8]byte{0xbd, 0xd3, 0x37, 0x5c, 0x62, 0xfa, 0x96, 0x18}}
	CenterPoint_Property_GUID syscall.GUID = syscall.GUID{0x0cb00c08, 0x540c, 0x4edb, 
	[8]byte{0x94, 0x45, 0x26, 0x35, 0x9e, 0xa6, 0x97, 0x85}}
	Rotation_Property_GUID syscall.GUID = syscall.GUID{0x767cdc7d, 0xaec0, 0x4110, 
	[8]byte{0xad, 0x32, 0x30, 0xed, 0xd4, 0x03, 0x49, 0x2e}}
	Size_Property_GUID syscall.GUID = syscall.GUID{0x2b5f761d, 0xf885, 0x4404, 
	[8]byte{0x97, 0x3f, 0x9b, 0x1d, 0x98, 0xe3, 0x6d, 0x8f}}
	ToolTipOpened_Event_GUID syscall.GUID = syscall.GUID{0x3f4b97ff, 0x2edc, 0x451d, 
	[8]byte{0xbc, 0xa4, 0x95, 0xa3, 0x18, 0x8d, 0x5b, 0x03}}
	ToolTipClosed_Event_GUID syscall.GUID = syscall.GUID{0x276d71ef, 0x24a9, 0x49b6, 
	[8]byte{0x8e, 0x97, 0xda, 0x98, 0xb4, 0x01, 0xbb, 0xcd}}
	StructureChanged_Event_GUID syscall.GUID = syscall.GUID{0x59977961, 0x3edd, 0x4b11, 
	[8]byte{0xb1, 0x3b, 0x67, 0x6b, 0x2a, 0x2a, 0x6c, 0xa9}}
	MenuOpened_Event_GUID syscall.GUID = syscall.GUID{0xebe2e945, 0x66ca, 0x4ed1, 
	[8]byte{0x9f, 0xf8, 0x2a, 0xd7, 0xdf, 0x0a, 0x1b, 0x08}}
	AutomationPropertyChanged_Event_GUID syscall.GUID = syscall.GUID{0x2527fba1, 0x8d7a, 0x4630, 
	[8]byte{0xa4, 0xcc, 0xe6, 0x63, 0x15, 0x94, 0x2f, 0x52}}
	AutomationFocusChanged_Event_GUID syscall.GUID = syscall.GUID{0xb68a1f17, 0xf60d, 0x41a7, 
	[8]byte{0xa3, 0xcc, 0xb0, 0x52, 0x92, 0x15, 0x5f, 0xe0}}
	ActiveTextPositionChanged_Event_GUID syscall.GUID = syscall.GUID{0xa5c09e9c, 0xc77d, 0x4f25, 
	[8]byte{0xb4, 0x91, 0xe5, 0xbb, 0x70, 0x17, 0xcb, 0xd4}}
	AsyncContentLoaded_Event_GUID syscall.GUID = syscall.GUID{0x5fdee11c, 0xd2fa, 0x4fb9, 
	[8]byte{0x90, 0x4e, 0x5c, 0xbe, 0xe8, 0x94, 0xd5, 0xef}}
	MenuClosed_Event_GUID syscall.GUID = syscall.GUID{0x3cf1266e, 0x1582, 0x4041, 
	[8]byte{0xac, 0xd7, 0x88, 0xa3, 0x5a, 0x96, 0x52, 0x97}}
	LayoutInvalidated_Event_GUID syscall.GUID = syscall.GUID{0xed7d6544, 0xa6bd, 0x4595, 
	[8]byte{0x9b, 0xae, 0x3d, 0x28, 0x94, 0x6c, 0xc7, 0x15}}
	Invoke_Invoked_Event_GUID syscall.GUID = syscall.GUID{0xdfd699f0, 0xc915, 0x49dd, 
	[8]byte{0xb4, 0x22, 0xdd, 0xe7, 0x85, 0xc3, 0xd2, 0x4b}}
	SelectionItem_ElementAddedToSelectionEvent_Event_GUID syscall.GUID = syscall.GUID{0x3c822dd1, 0xc407, 0x4dba, 
	[8]byte{0x91, 0xdd, 0x79, 0xd4, 0xae, 0xd0, 0xae, 0xc6}}
	SelectionItem_ElementRemovedFromSelectionEvent_Event_GUID syscall.GUID = syscall.GUID{0x097fa8a9, 0x7079, 0x41af, 
	[8]byte{0x8b, 0x9c, 0x09, 0x34, 0xd8, 0x30, 0x5e, 0x5c}}
	SelectionItem_ElementSelectedEvent_Event_GUID syscall.GUID = syscall.GUID{0xb9c7dbfb, 0x4ebe, 0x4532, 
	[8]byte{0xaa, 0xf4, 0x00, 0x8c, 0xf6, 0x47, 0x23, 0x3c}}
	Selection_InvalidatedEvent_Event_GUID syscall.GUID = syscall.GUID{0xcac14904, 0x16b4, 0x4b53, 
	[8]byte{0x8e, 0x47, 0x4c, 0xb1, 0xdf, 0x26, 0x7b, 0xb7}}
	Text_TextSelectionChangedEvent_Event_GUID syscall.GUID = syscall.GUID{0x918edaa1, 0x71b3, 0x49ae, 
	[8]byte{0x97, 0x41, 0x79, 0xbe, 0xb8, 0xd3, 0x58, 0xf3}}
	Text_TextChangedEvent_Event_GUID syscall.GUID = syscall.GUID{0x4a342082, 0xf483, 0x48c4, 
	[8]byte{0xac, 0x11, 0xa8, 0x4b, 0x43, 0x5e, 0x2a, 0x84}}
	Window_WindowOpened_Event_GUID syscall.GUID = syscall.GUID{0xd3e81d06, 0xde45, 0x4f2f, 
	[8]byte{0x96, 0x33, 0xde, 0x9e, 0x02, 0xfb, 0x65, 0xaf}}
	Window_WindowClosed_Event_GUID syscall.GUID = syscall.GUID{0xedf141f8, 0xfa67, 0x4e22, 
	[8]byte{0xbb, 0xf7, 0x94, 0x4e, 0x05, 0x73, 0x5e, 0xe2}}
	MenuModeStart_Event_GUID syscall.GUID = syscall.GUID{0x18d7c631, 0x166a, 0x4ac9, 
	[8]byte{0xae, 0x3b, 0xef, 0x4b, 0x54, 0x20, 0xe6, 0x81}}
	MenuModeEnd_Event_GUID syscall.GUID = syscall.GUID{0x9ecd4c9f, 0x80dd, 0x47b8, 
	[8]byte{0x82, 0x67, 0x5a, 0xec, 0x06, 0xbb, 0x2c, 0xff}}
	InputReachedTarget_Event_GUID syscall.GUID = syscall.GUID{0x93ed549a, 0x0549, 0x40f0, 
	[8]byte{0xbe, 0xdb, 0x28, 0xe4, 0x4f, 0x7d, 0xe2, 0xa3}}
	InputReachedOtherElement_Event_GUID syscall.GUID = syscall.GUID{0xed201d8a, 0x4e6c, 0x415e, 
	[8]byte{0xa8, 0x74, 0x24, 0x60, 0xc9, 0xb6, 0x6b, 0xa8}}
	InputDiscarded_Event_GUID syscall.GUID = syscall.GUID{0x7f36c367, 0x7b18, 0x417c, 
	[8]byte{0x97, 0xe3, 0x9d, 0x58, 0xdd, 0xc9, 0x44, 0xab}}
	SystemAlert_Event_GUID syscall.GUID = syscall.GUID{0xd271545d, 0x7a3a, 0x47a7, 
	[8]byte{0x84, 0x74, 0x81, 0xd2, 0x9a, 0x24, 0x51, 0xc9}}
	LiveRegionChanged_Event_GUID syscall.GUID = syscall.GUID{0x102d5e90, 0xe6a9, 0x41b6, 
	[8]byte{0xb1, 0xc5, 0xa9, 0xb1, 0x92, 0x9d, 0x95, 0x10}}
	HostedFragmentRootsInvalidated_Event_GUID syscall.GUID = syscall.GUID{0xe6bdb03e, 0x0921, 0x4ec5, 
	[8]byte{0x8d, 0xcf, 0xea, 0xe8, 0x77, 0xb0, 0x42, 0x6b}}
	Drag_DragStart_Event_GUID syscall.GUID = syscall.GUID{0x883a480b, 0x3aa9, 0x429d, 
	[8]byte{0x95, 0xe4, 0xd9, 0xc8, 0xd0, 0x11, 0xf0, 0xdd}}
	Drag_DragCancel_Event_GUID syscall.GUID = syscall.GUID{0xc3ede6fa, 0x3451, 0x4e0f, 
	[8]byte{0x9e, 0x71, 0xdf, 0x9c, 0x28, 0x0a, 0x46, 0x57}}
	Drag_DragComplete_Event_GUID syscall.GUID = syscall.GUID{0x38e96188, 0xef1f, 0x463e, 
	[8]byte{0x91, 0xca, 0x3a, 0x77, 0x92, 0xc2, 0x9c, 0xaf}}
	DropTarget_DragEnter_Event_GUID syscall.GUID = syscall.GUID{0xaad9319b, 0x032c, 0x4a88, 
	[8]byte{0x96, 0x1d, 0x1c, 0xf5, 0x79, 0x58, 0x1e, 0x34}}
	DropTarget_DragLeave_Event_GUID syscall.GUID = syscall.GUID{0x0f82eb15, 0x24a2, 0x4988, 
	[8]byte{0x92, 0x17, 0xde, 0x16, 0x2a, 0xee, 0x27, 0x2b}}
	DropTarget_Dropped_Event_GUID syscall.GUID = syscall.GUID{0x622cead8, 0x1edb, 0x4a3d, 
	[8]byte{0xab, 0xbc, 0xbe, 0x22, 0x11, 0xff, 0x68, 0xb5}}
	StructuredMarkup_CompositionComplete_Event_GUID syscall.GUID = syscall.GUID{0xc48a3c17, 0x677a, 0x4047, 
	[8]byte{0xa6, 0x8d, 0xfc, 0x12, 0x57, 0x52, 0x8a, 0xef}}
	StructuredMarkup_Deleted_Event_GUID syscall.GUID = syscall.GUID{0xf9d0a020, 0xe1c1, 0x4ecf, 
	[8]byte{0xb9, 0xaa, 0x52, 0xef, 0xde, 0x7e, 0x41, 0xe1}}
	StructuredMarkup_SelectionChanged_Event_GUID syscall.GUID = syscall.GUID{0xa7c815f7, 0xff9f, 0x41c7, 
	[8]byte{0xa3, 0xa7, 0xab, 0x6c, 0xbf, 0xdb, 0x49, 0x03}}
	Invoke_Pattern_GUID syscall.GUID = syscall.GUID{0xd976c2fc, 0x66ea, 0x4a6e, 
	[8]byte{0xb2, 0x8f, 0xc2, 0x4c, 0x75, 0x46, 0xad, 0x37}}
	Selection_Pattern_GUID syscall.GUID = syscall.GUID{0x66e3b7e8, 0xd821, 0x4d25, 
	[8]byte{0x87, 0x61, 0x43, 0x5d, 0x2c, 0x8b, 0x25, 0x3f}}
	Value_Pattern_GUID syscall.GUID = syscall.GUID{0x17faad9e, 0xc877, 0x475b, 
	[8]byte{0xb9, 0x33, 0x77, 0x33, 0x27, 0x79, 0xb6, 0x37}}
	RangeValue_Pattern_GUID syscall.GUID = syscall.GUID{0x18b00d87, 0xb1c9, 0x476a, 
	[8]byte{0xbf, 0xbd, 0x5f, 0x0b, 0xdb, 0x92, 0x6f, 0x63}}
	Scroll_Pattern_GUID syscall.GUID = syscall.GUID{0x895fa4b4, 0x759d, 0x4c50, 
	[8]byte{0x8e, 0x15, 0x03, 0x46, 0x06, 0x72, 0x00, 0x3c}}
	ExpandCollapse_Pattern_GUID syscall.GUID = syscall.GUID{0xae05efa2, 0xf9d1, 0x428a, 
	[8]byte{0x83, 0x4c, 0x53, 0xa5, 0xc5, 0x2f, 0x9b, 0x8b}}
	Grid_Pattern_GUID syscall.GUID = syscall.GUID{0x260a2ccb, 0x93a8, 0x4e44, 
	[8]byte{0xa4, 0xc1, 0x3d, 0xf3, 0x97, 0xf2, 0xb0, 0x2b}}
	GridItem_Pattern_GUID syscall.GUID = syscall.GUID{0xf2d5c877, 0xa462, 0x4957, 
	[8]byte{0xa2, 0xa5, 0x2c, 0x96, 0xb3, 0x03, 0xbc, 0x63}}
	MultipleView_Pattern_GUID syscall.GUID = syscall.GUID{0x547a6ae4, 0x113f, 0x47c4, 
	[8]byte{0x85, 0x0f, 0xdb, 0x4d, 0xfa, 0x46, 0x6b, 0x1d}}
	Window_Pattern_GUID syscall.GUID = syscall.GUID{0x27901735, 0xc760, 0x4994, 
	[8]byte{0xad, 0x11, 0x59, 0x19, 0xe6, 0x06, 0xb1, 0x10}}
	SelectionItem_Pattern_GUID syscall.GUID = syscall.GUID{0x9bc64eeb, 0x87c7, 0x4b28, 
	[8]byte{0x94, 0xbb, 0x4d, 0x9f, 0xa4, 0x37, 0xb6, 0xef}}
	Dock_Pattern_GUID syscall.GUID = syscall.GUID{0x9cbaa846, 0x83c8, 0x428d, 
	[8]byte{0x82, 0x7f, 0x7e, 0x60, 0x63, 0xfe, 0x06, 0x20}}
	Table_Pattern_GUID syscall.GUID = syscall.GUID{0xc415218e, 0xa028, 0x461e, 
	[8]byte{0xaa, 0x92, 0x8f, 0x92, 0x5c, 0xf7, 0x93, 0x51}}
	TableItem_Pattern_GUID syscall.GUID = syscall.GUID{0xdf1343bd, 0x1888, 0x4a29, 
	[8]byte{0xa5, 0x0c, 0xb9, 0x2e, 0x6d, 0xe3, 0x7f, 0x6f}}
	Text_Pattern_GUID syscall.GUID = syscall.GUID{0x8615f05d, 0x7de5, 0x44fd, 
	[8]byte{0xa6, 0x79, 0x2c, 0xa4, 0xb4, 0x60, 0x33, 0xa8}}
	Toggle_Pattern_GUID syscall.GUID = syscall.GUID{0x0b419760, 0xe2f4, 0x43ff, 
	[8]byte{0x8c, 0x5f, 0x94, 0x57, 0xc8, 0x2b, 0x56, 0xe9}}
	Transform_Pattern_GUID syscall.GUID = syscall.GUID{0x24b46fdb, 0x587e, 0x49f1, 
	[8]byte{0x9c, 0x4a, 0xd8, 0xe9, 0x8b, 0x66, 0x4b, 0x7b}}
	ScrollItem_Pattern_GUID syscall.GUID = syscall.GUID{0x4591d005, 0xa803, 0x4d5c, 
	[8]byte{0xb4, 0xd5, 0x8d, 0x28, 0x00, 0xf9, 0x06, 0xa7}}
	LegacyIAccessible_Pattern_GUID syscall.GUID = syscall.GUID{0x54cc0a9f, 0x3395, 0x48af, 
	[8]byte{0xba, 0x8d, 0x73, 0xf8, 0x56, 0x90, 0xf3, 0xe0}}
	ItemContainer_Pattern_GUID syscall.GUID = syscall.GUID{0x3d13da0f, 0x8b9a, 0x4a99, 
	[8]byte{0x85, 0xfa, 0xc5, 0xc9, 0xa6, 0x9f, 0x1e, 0xd4}}
	VirtualizedItem_Pattern_GUID syscall.GUID = syscall.GUID{0xf510173e, 0x2e71, 0x45e9, 
	[8]byte{0xa6, 0xe5, 0x62, 0xf6, 0xed, 0x82, 0x89, 0xd5}}
	SynchronizedInput_Pattern_GUID syscall.GUID = syscall.GUID{0x05c288a6, 0xc47b, 0x488b, 
	[8]byte{0xb6, 0x53, 0x33, 0x97, 0x7a, 0x55, 0x1b, 0x8b}}
	ObjectModel_Pattern_GUID syscall.GUID = syscall.GUID{0x3e04acfe, 0x08fc, 0x47ec, 
	[8]byte{0x96, 0xbc, 0x35, 0x3f, 0xa3, 0xb3, 0x4a, 0xa7}}
	Annotation_Pattern_GUID syscall.GUID = syscall.GUID{0xf6c72ad7, 0x356c, 0x4850, 
	[8]byte{0x92, 0x91, 0x31, 0x6f, 0x60, 0x8a, 0x8c, 0x84}}
	Text_Pattern2_GUID syscall.GUID = syscall.GUID{0x498479a2, 0x5b22, 0x448d, 
	[8]byte{0xb6, 0xe4, 0x64, 0x74, 0x90, 0x86, 0x06, 0x98}}
	TextEdit_Pattern_GUID syscall.GUID = syscall.GUID{0x69f3ff89, 0x5af9, 0x4c75, 
	[8]byte{0x93, 0x40, 0xf2, 0xde, 0x29, 0x2e, 0x45, 0x91}}
	CustomNavigation_Pattern_GUID syscall.GUID = syscall.GUID{0xafea938a, 0x621e, 0x4054, 
	[8]byte{0xbb, 0x2c, 0x2f, 0x46, 0x11, 0x4d, 0xac, 0x3f}}
	Styles_Pattern_GUID syscall.GUID = syscall.GUID{0x1ae62655, 0xda72, 0x4d60, 
	[8]byte{0xa1, 0x53, 0xe5, 0xaa, 0x69, 0x88, 0xe3, 0xbf}}
	Spreadsheet_Pattern_GUID syscall.GUID = syscall.GUID{0x6a5b24c9, 0x9d1e, 0x4b85, 
	[8]byte{0x9e, 0x44, 0xc0, 0x2e, 0x31, 0x69, 0xb1, 0x0b}}
	SpreadsheetItem_Pattern_GUID syscall.GUID = syscall.GUID{0x32cf83ff, 0xf1a8, 0x4a8c, 
	[8]byte{0x86, 0x58, 0xd4, 0x7b, 0xa7, 0x4e, 0x20, 0xba}}
	Tranform_Pattern2_GUID syscall.GUID = syscall.GUID{0x8afcfd07, 0xa369, 0x44de, 
	[8]byte{0x98, 0x8b, 0x2f, 0x7f, 0xf4, 0x9f, 0xb8, 0xa8}}
	TextChild_Pattern_GUID syscall.GUID = syscall.GUID{0x7533cab7, 0x3bfe, 0x41ef, 
	[8]byte{0x9e, 0x85, 0xe2, 0x63, 0x8c, 0xbe, 0x16, 0x9e}}
	Drag_Pattern_GUID syscall.GUID = syscall.GUID{0xc0bee21f, 0xccb3, 0x4fed, 
	[8]byte{0x99, 0x5b, 0x11, 0x4f, 0x6e, 0x3d, 0x27, 0x28}}
	DropTarget_Pattern_GUID syscall.GUID = syscall.GUID{0x0bcbec56, 0xbd34, 0x4b7b, 
	[8]byte{0x9f, 0xd5, 0x26, 0x59, 0x90, 0x5e, 0xa3, 0xdc}}
	StructuredMarkup_Pattern_GUID syscall.GUID = syscall.GUID{0xabbd0878, 0x8665, 0x4f5c, 
	[8]byte{0x94, 0xfc, 0x36, 0xe7, 0xd8, 0xbb, 0x70, 0x6b}}
	Button_Control_GUID syscall.GUID = syscall.GUID{0x5a78e369, 0xc6a1, 0x4f33, 
	[8]byte{0xa9, 0xd7, 0x79, 0xf2, 0x0d, 0x0c, 0x78, 0x8e}}
	Calendar_Control_GUID syscall.GUID = syscall.GUID{0x8913eb88, 0x00e5, 0x46bc, 
	[8]byte{0x8e, 0x4e, 0x14, 0xa7, 0x86, 0xe1, 0x65, 0xa1}}
	CheckBox_Control_GUID syscall.GUID = syscall.GUID{0xfb50f922, 0xa3db, 0x49c0, 
	[8]byte{0x8b, 0xc3, 0x06, 0xda, 0xd5, 0x57, 0x78, 0xe2}}
	ComboBox_Control_GUID syscall.GUID = syscall.GUID{0x54cb426c, 0x2f33, 0x4fff, 
	[8]byte{0xaa, 0xa1, 0xae, 0xf6, 0x0d, 0xac, 0x5d, 0xeb}}
	Edit_Control_GUID syscall.GUID = syscall.GUID{0x6504a5c8, 0x2c86, 0x4f87, 
	[8]byte{0xae, 0x7b, 0x1a, 0xbd, 0xdc, 0x81, 0x0c, 0xf9}}
	Hyperlink_Control_GUID syscall.GUID = syscall.GUID{0x8a56022c, 0xb00d, 0x4d15, 
	[8]byte{0x8f, 0xf0, 0x5b, 0x6b, 0x26, 0x6e, 0x5e, 0x02}}
	Image_Control_GUID syscall.GUID = syscall.GUID{0x2d3736e4, 0x6b16, 0x4c57, 
	[8]byte{0xa9, 0x62, 0xf9, 0x32, 0x60, 0xa7, 0x52, 0x43}}
	ListItem_Control_GUID syscall.GUID = syscall.GUID{0x7b3717f2, 0x44d1, 0x4a58, 
	[8]byte{0x98, 0xa8, 0xf1, 0x2a, 0x9b, 0x8f, 0x78, 0xe2}}
	List_Control_GUID syscall.GUID = syscall.GUID{0x9b149ee1, 0x7cca, 0x4cfc, 
	[8]byte{0x9a, 0xf1, 0xca, 0xc7, 0xbd, 0xdd, 0x30, 0x31}}
	Menu_Control_GUID syscall.GUID = syscall.GUID{0x2e9b1440, 0x0ea8, 0x41fd, 
	[8]byte{0xb3, 0x74, 0xc1, 0xea, 0x6f, 0x50, 0x3c, 0xd1}}
	MenuBar_Control_GUID syscall.GUID = syscall.GUID{0xcc384250, 0x0e7b, 0x4ae8, 
	[8]byte{0x95, 0xae, 0xa0, 0x8f, 0x26, 0x1b, 0x52, 0xee}}
	MenuItem_Control_GUID syscall.GUID = syscall.GUID{0xf45225d3, 0xd0a0, 0x49d8, 
	[8]byte{0x98, 0x34, 0x9a, 0x00, 0x0d, 0x2a, 0xed, 0xdc}}
	ProgressBar_Control_GUID syscall.GUID = syscall.GUID{0x228c9f86, 0xc36c, 0x47bb, 
	[8]byte{0x9f, 0xb6, 0xa5, 0x83, 0x4b, 0xfc, 0x53, 0xa4}}
	RadioButton_Control_GUID syscall.GUID = syscall.GUID{0x3bdb49db, 0xfe2c, 0x4483, 
	[8]byte{0xb3, 0xe1, 0xe5, 0x7f, 0x21, 0x94, 0x40, 0xc6}}
	ScrollBar_Control_GUID syscall.GUID = syscall.GUID{0xdaf34b36, 0x5065, 0x4946, 
	[8]byte{0xb2, 0x2f, 0x92, 0x59, 0x5f, 0xc0, 0x75, 0x1a}}
	Slider_Control_GUID syscall.GUID = syscall.GUID{0xb033c24b, 0x3b35, 0x4cea, 
	[8]byte{0xb6, 0x09, 0x76, 0x36, 0x82, 0xfa, 0x66, 0x0b}}
	Spinner_Control_GUID syscall.GUID = syscall.GUID{0x60cc4b38, 0x3cb1, 0x4161, 
	[8]byte{0xb4, 0x42, 0xc6, 0xb7, 0x26, 0xc1, 0x78, 0x25}}
	StatusBar_Control_GUID syscall.GUID = syscall.GUID{0xd45e7d1b, 0x5873, 0x475f, 
	[8]byte{0x95, 0xa4, 0x04, 0x33, 0xe1, 0xf1, 0xb0, 0x0a}}
	Tab_Control_GUID syscall.GUID = syscall.GUID{0x38cd1f2d, 0x337a, 0x4bd2, 
	[8]byte{0xa5, 0xe3, 0xad, 0xb4, 0x69, 0xe3, 0x0b, 0xd3}}
	TabItem_Control_GUID syscall.GUID = syscall.GUID{0x2c6a634f, 0x921b, 0x4e6e, 
	[8]byte{0xb2, 0x6e, 0x08, 0xfc, 0xb0, 0x79, 0x8f, 0x4c}}
	Text_Control_GUID syscall.GUID = syscall.GUID{0xae9772dc, 0xd331, 0x4f09, 
	[8]byte{0xbe, 0x20, 0x7e, 0x6d, 0xfa, 0xf0, 0x7b, 0x0a}}
	ToolBar_Control_GUID syscall.GUID = syscall.GUID{0x8f06b751, 0xe182, 0x4e98, 
	[8]byte{0x88, 0x93, 0x22, 0x84, 0x54, 0x3a, 0x7d, 0xce}}
	ToolTip_Control_GUID syscall.GUID = syscall.GUID{0x05ddc6d1, 0x2137, 0x4768, 
	[8]byte{0x98, 0xea, 0x73, 0xf5, 0x2f, 0x71, 0x34, 0xf3}}
	Tree_Control_GUID syscall.GUID = syscall.GUID{0x7561349c, 0xd241, 0x43f4, 
	[8]byte{0x99, 0x08, 0xb5, 0xf0, 0x91, 0xbe, 0xe6, 0x11}}
	TreeItem_Control_GUID syscall.GUID = syscall.GUID{0x62c9feb9, 0x8ffc, 0x4878, 
	[8]byte{0xa3, 0xa4, 0x96, 0xb0, 0x30, 0x31, 0x5c, 0x18}}
	Custom_Control_GUID syscall.GUID = syscall.GUID{0xf29ea0c3, 0xadb7, 0x430a, 
	[8]byte{0xba, 0x90, 0xe5, 0x2c, 0x73, 0x13, 0xe6, 0xed}}
	Group_Control_GUID syscall.GUID = syscall.GUID{0xad50aa1c, 0xe8c8, 0x4774, 
	[8]byte{0xae, 0x1b, 0xdd, 0x86, 0xdf, 0x0b, 0x3b, 0xdc}}
	Thumb_Control_GUID syscall.GUID = syscall.GUID{0x701ca877, 0xe310, 0x4dd6, 
	[8]byte{0xb6, 0x44, 0x79, 0x7e, 0x4f, 0xae, 0xa2, 0x13}}
	DataGrid_Control_GUID syscall.GUID = syscall.GUID{0x84b783af, 0xd103, 0x4b0a, 
	[8]byte{0x84, 0x15, 0xe7, 0x39, 0x42, 0x41, 0x0f, 0x4b}}
	DataItem_Control_GUID syscall.GUID = syscall.GUID{0xa0177842, 0xd94f, 0x42a5, 
	[8]byte{0x81, 0x4b, 0x60, 0x68, 0xad, 0xdc, 0x8d, 0xa5}}
	Document_Control_GUID syscall.GUID = syscall.GUID{0x3cd6bb6f, 0x6f08, 0x4562, 
	[8]byte{0xb2, 0x29, 0xe4, 0xe2, 0xfc, 0x7a, 0x9e, 0xb4}}
	SplitButton_Control_GUID syscall.GUID = syscall.GUID{0x7011f01f, 0x4ace, 0x4901, 
	[8]byte{0xb4, 0x61, 0x92, 0x0a, 0x6f, 0x1c, 0xa6, 0x50}}
	Window_Control_GUID syscall.GUID = syscall.GUID{0xe13a7242, 0xf462, 0x4f4d, 
	[8]byte{0xae, 0xc1, 0x53, 0xb2, 0x8d, 0x6c, 0x32, 0x90}}
	Pane_Control_GUID syscall.GUID = syscall.GUID{0x5c2b3f5b, 0x9182, 0x42a3, 
	[8]byte{0x8d, 0xec, 0x8c, 0x04, 0xc1, 0xee, 0x63, 0x4d}}
	Header_Control_GUID syscall.GUID = syscall.GUID{0x5b90cbce, 0x78fb, 0x4614, 
	[8]byte{0x82, 0xb6, 0x55, 0x4d, 0x74, 0x71, 0x8e, 0x67}}
	HeaderItem_Control_GUID syscall.GUID = syscall.GUID{0xe6bc12cb, 0x7c8e, 0x49cf, 
	[8]byte{0xb1, 0x68, 0x4a, 0x93, 0xa3, 0x2b, 0xeb, 0xb0}}
	Table_Control_GUID syscall.GUID = syscall.GUID{0x773bfa0e, 0x5bc4, 0x4deb, 
	[8]byte{0x92, 0x1b, 0xde, 0x7b, 0x32, 0x06, 0x22, 0x9e}}
	TitleBar_Control_GUID syscall.GUID = syscall.GUID{0x98aa55bf, 0x3bb0, 0x4b65, 
	[8]byte{0x83, 0x6e, 0x2e, 0xa3, 0x0d, 0xbc, 0x17, 0x1f}}
	Separator_Control_GUID syscall.GUID = syscall.GUID{0x8767eba3, 0x2a63, 0x4ab0, 
	[8]byte{0xac, 0x8d, 0xaa, 0x50, 0xe2, 0x3d, 0xe9, 0x78}}
	SemanticZoom_Control_GUID syscall.GUID = syscall.GUID{0x5fd34a43, 0x061e, 0x42c8, 
	[8]byte{0xb5, 0x89, 0x9d, 0xcc, 0xf7, 0x4b, 0xc4, 0x3a}}
	AppBar_Control_GUID syscall.GUID = syscall.GUID{0x6114908d, 0xcc02, 0x4d37, 
	[8]byte{0x87, 0x5b, 0xb5, 0x30, 0xc7, 0x13, 0x95, 0x54}}
	Text_AnimationStyle_Attribute_GUID syscall.GUID = syscall.GUID{0x628209f0, 0x7c9a, 0x4d57, 
	[8]byte{0xbe, 0x64, 0x1f, 0x18, 0x36, 0x57, 0x1f, 0xf5}}
	Text_BackgroundColor_Attribute_GUID syscall.GUID = syscall.GUID{0xfdc49a07, 0x583d, 0x4f17, 
	[8]byte{0xad, 0x27, 0x77, 0xfc, 0x83, 0x2a, 0x3c, 0x0b}}
	Text_BulletStyle_Attribute_GUID syscall.GUID = syscall.GUID{0xc1097c90, 0xd5c4, 0x4237, 
	[8]byte{0x97, 0x81, 0x3b, 0xec, 0x8b, 0xa5, 0x4e, 0x48}}
	Text_CapStyle_Attribute_GUID syscall.GUID = syscall.GUID{0xfb059c50, 0x92cc, 0x49a5, 
	[8]byte{0xba, 0x8f, 0x0a, 0xa8, 0x72, 0xbb, 0xa2, 0xf3}}
	Text_Culture_Attribute_GUID syscall.GUID = syscall.GUID{0xc2025af9, 0xa42d, 0x4ced, 
	[8]byte{0xa1, 0xfb, 0xc6, 0x74, 0x63, 0x15, 0x22, 0x2e}}
	Text_FontName_Attribute_GUID syscall.GUID = syscall.GUID{0x64e63ba8, 0xf2e5, 0x476e, 
	[8]byte{0xa4, 0x77, 0x17, 0x34, 0xfe, 0xaa, 0xf7, 0x26}}
	Text_FontSize_Attribute_GUID syscall.GUID = syscall.GUID{0xdc5eeeff, 0x0506, 0x4673, 
	[8]byte{0x93, 0xf2, 0x37, 0x7e, 0x4a, 0x8e, 0x01, 0xf1}}
	Text_FontWeight_Attribute_GUID syscall.GUID = syscall.GUID{0x6fc02359, 0xb316, 0x4f5f, 
	[8]byte{0xb4, 0x01, 0xf1, 0xce, 0x55, 0x74, 0x18, 0x53}}
	Text_ForegroundColor_Attribute_GUID syscall.GUID = syscall.GUID{0x72d1c95d, 0x5e60, 0x471a, 
	[8]byte{0x96, 0xb1, 0x6c, 0x1b, 0x3b, 0x77, 0xa4, 0x36}}
	Text_HorizontalTextAlignment_Attribute_GUID syscall.GUID = syscall.GUID{0x04ea6161, 0xfba3, 0x477a, 
	[8]byte{0x95, 0x2a, 0xbb, 0x32, 0x6d, 0x02, 0x6a, 0x5b}}
	Text_IndentationFirstLine_Attribute_GUID syscall.GUID = syscall.GUID{0x206f9ad5, 0xc1d3, 0x424a, 
	[8]byte{0x81, 0x82, 0x6d, 0xa9, 0xa7, 0xf3, 0xd6, 0x32}}
	Text_IndentationLeading_Attribute_GUID syscall.GUID = syscall.GUID{0x5cf66bac, 0x2d45, 0x4a4b, 
	[8]byte{0xb6, 0xc9, 0xf7, 0x22, 0x1d, 0x28, 0x15, 0xb0}}
	Text_IndentationTrailing_Attribute_GUID syscall.GUID = syscall.GUID{0x97ff6c0f, 0x1ce4, 0x408a, 
	[8]byte{0xb6, 0x7b, 0x94, 0xd8, 0x3e, 0xb6, 0x9b, 0xf2}}
	Text_IsHidden_Attribute_GUID syscall.GUID = syscall.GUID{0x360182fb, 0xbdd7, 0x47f6, 
	[8]byte{0xab, 0x69, 0x19, 0xe3, 0x3f, 0x8a, 0x33, 0x44}}
	Text_IsItalic_Attribute_GUID syscall.GUID = syscall.GUID{0xfce12a56, 0x1336, 0x4a34, 
	[8]byte{0x96, 0x63, 0x1b, 0xab, 0x47, 0x23, 0x93, 0x20}}
	Text_IsReadOnly_Attribute_GUID syscall.GUID = syscall.GUID{0xa738156b, 0xca3e, 0x495e, 
	[8]byte{0x95, 0x14, 0x83, 0x3c, 0x44, 0x0f, 0xeb, 0x11}}
	Text_IsSubscript_Attribute_GUID syscall.GUID = syscall.GUID{0xf0ead858, 0x8f53, 0x413c, 
	[8]byte{0x87, 0x3f, 0x1a, 0x7d, 0x7f, 0x5e, 0x0d, 0xe4}}
	Text_IsSuperscript_Attribute_GUID syscall.GUID = syscall.GUID{0xda706ee4, 0xb3aa, 0x4645, 
	[8]byte{0xa4, 0x1f, 0xcd, 0x25, 0x15, 0x7d, 0xea, 0x76}}
	Text_MarginBottom_Attribute_GUID syscall.GUID = syscall.GUID{0x7ee593c4, 0x72b4, 0x4cac, 
	[8]byte{0x92, 0x71, 0x3e, 0xd2, 0x4b, 0x0e, 0x4d, 0x42}}
	Text_MarginLeading_Attribute_GUID syscall.GUID = syscall.GUID{0x9e9242d0, 0x5ed0, 0x4900, 
	[8]byte{0x8e, 0x8a, 0xee, 0xcc, 0x03, 0x83, 0x5a, 0xfc}}
	Text_MarginTop_Attribute_GUID syscall.GUID = syscall.GUID{0x683d936f, 0xc9b9, 0x4a9a, 
	[8]byte{0xb3, 0xd9, 0xd2, 0x0d, 0x33, 0x31, 0x1e, 0x2a}}
	Text_MarginTrailing_Attribute_GUID syscall.GUID = syscall.GUID{0xaf522f98, 0x999d, 0x40af, 
	[8]byte{0xa5, 0xb2, 0x01, 0x69, 0xd0, 0x34, 0x20, 0x02}}
	Text_OutlineStyles_Attribute_GUID syscall.GUID = syscall.GUID{0x5b675b27, 0xdb89, 0x46fe, 
	[8]byte{0x97, 0x0c, 0x61, 0x4d, 0x52, 0x3b, 0xb9, 0x7d}}
	Text_OverlineColor_Attribute_GUID syscall.GUID = syscall.GUID{0x83ab383a, 0xfd43, 0x40da, 
	[8]byte{0xab, 0x3e, 0xec, 0xf8, 0x16, 0x5c, 0xbb, 0x6d}}
	Text_OverlineStyle_Attribute_GUID syscall.GUID = syscall.GUID{0x0a234d66, 0x617e, 0x427f, 
	[8]byte{0x87, 0x1d, 0xe1, 0xff, 0x1e, 0x0c, 0x21, 0x3f}}
	Text_StrikethroughColor_Attribute_GUID syscall.GUID = syscall.GUID{0xbfe15a18, 0x8c41, 0x4c5a, 
	[8]byte{0x9a, 0x0b, 0x04, 0xaf, 0x0e, 0x07, 0xf4, 0x87}}
	Text_StrikethroughStyle_Attribute_GUID syscall.GUID = syscall.GUID{0x72913ef1, 0xda00, 0x4f01, 
	[8]byte{0x89, 0x9c, 0xac, 0x5a, 0x85, 0x77, 0xa3, 0x07}}
	Text_Tabs_Attribute_GUID syscall.GUID = syscall.GUID{0x2e68d00b, 0x92fe, 0x42d8, 
	[8]byte{0x89, 0x9a, 0xa7, 0x84, 0xaa, 0x44, 0x54, 0xa1}}
	Text_TextFlowDirections_Attribute_GUID syscall.GUID = syscall.GUID{0x8bdf8739, 0xf420, 0x423e, 
	[8]byte{0xaf, 0x77, 0x20, 0xa5, 0xd9, 0x73, 0xa9, 0x07}}
	Text_UnderlineColor_Attribute_GUID syscall.GUID = syscall.GUID{0xbfa12c73, 0xfde2, 0x4473, 
	[8]byte{0xbf, 0x64, 0x10, 0x36, 0xd6, 0xaa, 0x0f, 0x45}}
	Text_UnderlineStyle_Attribute_GUID syscall.GUID = syscall.GUID{0x5f3b21c0, 0xede4, 0x44bd, 
	[8]byte{0x9c, 0x36, 0x38, 0x53, 0x03, 0x8c, 0xbf, 0xeb}}
	Text_AnnotationTypes_Attribute_GUID syscall.GUID = syscall.GUID{0xad2eb431, 0xee4e, 0x4be1, 
	[8]byte{0xa7, 0xba, 0x55, 0x59, 0x15, 0x5a, 0x73, 0xef}}
	Text_AnnotationObjects_Attribute_GUID syscall.GUID = syscall.GUID{0xff41cf68, 0xe7ab, 0x40b9, 
	[8]byte{0x8c, 0x72, 0x72, 0xa8, 0xed, 0x94, 0x01, 0x7d}}
	Text_StyleName_Attribute_GUID syscall.GUID = syscall.GUID{0x22c9e091, 0x4d66, 0x45d8, 
	[8]byte{0xa8, 0x28, 0x73, 0x7b, 0xab, 0x4c, 0x98, 0xa7}}
	Text_StyleId_Attribute_GUID syscall.GUID = syscall.GUID{0x14c300de, 0xc32b, 0x449b, 
	[8]byte{0xab, 0x7c, 0xb0, 0xe0, 0x78, 0x9a, 0xea, 0x5d}}
	Text_Link_Attribute_GUID syscall.GUID = syscall.GUID{0xb38ef51d, 0x9e8d, 0x4e46, 
	[8]byte{0x91, 0x44, 0x56, 0xeb, 0xe1, 0x77, 0x32, 0x9b}}
	Text_IsActive_Attribute_GUID syscall.GUID = syscall.GUID{0xf5a4e533, 0xe1b8, 0x436b, 
	[8]byte{0x93, 0x5d, 0xb5, 0x7a, 0xa3, 0xf5, 0x58, 0xc4}}
	Text_SelectionActiveEnd_Attribute_GUID syscall.GUID = syscall.GUID{0x1f668cc3, 0x9bbf, 0x416b, 
	[8]byte{0xb0, 0xa2, 0xf8, 0x9f, 0x86, 0xf6, 0x61, 0x2c}}
	Text_CaretPosition_Attribute_GUID syscall.GUID = syscall.GUID{0xb227b131, 0x9889, 0x4752, 
	[8]byte{0xa9, 0x1b, 0x73, 0x3e, 0xfd, 0xc5, 0xc5, 0xa0}}
	Text_CaretBidiMode_Attribute_GUID syscall.GUID = syscall.GUID{0x929ee7a6, 0x51d3, 0x4715, 
	[8]byte{0x96, 0xdc, 0xb6, 0x94, 0xfa, 0x24, 0xa1, 0x68}}
	Text_BeforeParagraphSpacing_Attribute_GUID syscall.GUID = syscall.GUID{0xbe7b0ab1, 0xc822, 0x4a24, 
	[8]byte{0x85, 0xe9, 0xc8, 0xf2, 0x65, 0x0f, 0xc7, 0x9c}}
	Text_AfterParagraphSpacing_Attribute_GUID syscall.GUID = syscall.GUID{0x588cbb38, 0xe62f, 0x497c, 
	[8]byte{0xb5, 0xd1, 0xcc, 0xdf, 0x0e, 0xe8, 0x23, 0xd8}}
	Text_LineSpacing_Attribute_GUID syscall.GUID = syscall.GUID{0x63ff70ae, 0xd943, 0x4b47, 
	[8]byte{0x8a, 0xb7, 0xa7, 0xa0, 0x33, 0xd3, 0x21, 0x4b}}
	Text_BeforeSpacing_Attribute_GUID syscall.GUID = syscall.GUID{0xbe7b0ab1, 0xc822, 0x4a24, 
	[8]byte{0x85, 0xe9, 0xc8, 0xf2, 0x65, 0x0f, 0xc7, 0x9c}}
	Text_AfterSpacing_Attribute_GUID syscall.GUID = syscall.GUID{0x588cbb38, 0xe62f, 0x497c, 
	[8]byte{0xb5, 0xd1, 0xcc, 0xdf, 0x0e, 0xe8, 0x23, 0xd8}}
	Text_SayAsInterpretAs_Attribute_GUID syscall.GUID = syscall.GUID{0xb38ad6ac, 0xeee1, 0x4b6e, 
	[8]byte{0x88, 0xcc, 0x01, 0x4c, 0xef, 0xa9, 0x3f, 0xcb}}
	TextEdit_TextChanged_Event_GUID syscall.GUID = syscall.GUID{0x120b0308, 0xec22, 0x4eb8, 
	[8]byte{0x9c, 0x98, 0x98, 0x67, 0xcd, 0xa1, 0xb1, 0x65}}
	TextEdit_ConversionTargetChanged_Event_GUID syscall.GUID = syscall.GUID{0x3388c183, 0xed4f, 0x4c8b, 
	[8]byte{0x9b, 0xaa, 0x36, 0x4d, 0x51, 0xd8, 0x84, 0x7f}}
	Changes_Event_GUID syscall.GUID = syscall.GUID{0x7df26714, 0x614f, 0x4e05, 
	[8]byte{0x94, 0x88, 0x71, 0x6c, 0x5b, 0xa1, 0x94, 0x36}}
	Annotation_Custom_GUID syscall.GUID = syscall.GUID{0x9ec82750, 0x3931, 0x4952, 
	[8]byte{0x85, 0xbc, 0x1d, 0xbf, 0xf7, 0x8a, 0x43, 0xe3}}
	Annotation_SpellingError_GUID syscall.GUID = syscall.GUID{0xae85567e, 0x9ece, 0x423f, 
	[8]byte{0x81, 0xb7, 0x96, 0xc4, 0x3d, 0x53, 0xe5, 0x0e}}
	Annotation_GrammarError_GUID syscall.GUID = syscall.GUID{0x757a048d, 0x4518, 0x41c6, 
	[8]byte{0x85, 0x4c, 0xdc, 0x00, 0x9b, 0x7c, 0xfb, 0x53}}
	Annotation_Comment_GUID syscall.GUID = syscall.GUID{0xfd2fda30, 0x26b3, 0x4c06, 
	[8]byte{0x8b, 0xc7, 0x98, 0xf1, 0x53, 0x2e, 0x46, 0xfd}}
	Annotation_FormulaError_GUID syscall.GUID = syscall.GUID{0x95611982, 0x0cab, 0x46d5, 
	[8]byte{0xa2, 0xf0, 0xe3, 0x0d, 0x19, 0x05, 0xf8, 0xbf}}
	Annotation_TrackChanges_GUID syscall.GUID = syscall.GUID{0x21e6e888, 0xdc14, 0x4016, 
	[8]byte{0xac, 0x27, 0x19, 0x05, 0x53, 0xc8, 0xc4, 0x70}}
	Annotation_Header_GUID syscall.GUID = syscall.GUID{0x867b409b, 0xb216, 0x4472, 
	[8]byte{0xa2, 0x19, 0x52, 0x5e, 0x31, 0x06, 0x81, 0xf8}}
	Annotation_Footer_GUID syscall.GUID = syscall.GUID{0xcceab046, 0x1833, 0x47aa, 
	[8]byte{0x80, 0x80, 0x70, 0x1e, 0xd0, 0xb0, 0xc8, 0x32}}
	Annotation_Highlighted_GUID syscall.GUID = syscall.GUID{0x757c884e, 0x8083, 0x4081, 
	[8]byte{0x8b, 0x9c, 0xe8, 0x7f, 0x50, 0x72, 0xf0, 0xe4}}
	Annotation_Endnote_GUID syscall.GUID = syscall.GUID{0x7565725c, 0x2d99, 0x4839, 
	[8]byte{0x96, 0x0d, 0x33, 0xd3, 0xb8, 0x66, 0xab, 0xa5}}
	Annotation_Footnote_GUID syscall.GUID = syscall.GUID{0x3de10e21, 0x4125, 0x42db, 
	[8]byte{0x86, 0x20, 0xbe, 0x80, 0x83, 0x08, 0x06, 0x24}}
	Annotation_InsertionChange_GUID syscall.GUID = syscall.GUID{0x0dbeb3a6, 0xdf15, 0x4164, 
	[8]byte{0xa3, 0xc0, 0xe2, 0x1a, 0x8c, 0xe9, 0x31, 0xc4}}
	Annotation_DeletionChange_GUID syscall.GUID = syscall.GUID{0xbe3d5b05, 0x951d, 0x42e7, 
	[8]byte{0x90, 0x1d, 0xad, 0xc8, 0xc2, 0xcf, 0x34, 0xd0}}
	Annotation_MoveChange_GUID syscall.GUID = syscall.GUID{0x9da587eb, 0x23e5, 0x4490, 
	[8]byte{0xb3, 0x85, 0x1a, 0x22, 0xdd, 0xc8, 0xb1, 0x87}}
	Annotation_FormatChange_GUID syscall.GUID = syscall.GUID{0xeb247345, 0xd4f1, 0x41ce, 
	[8]byte{0x8e, 0x52, 0xf7, 0x9b, 0x69, 0x63, 0x5e, 0x48}}
	Annotation_UnsyncedChange_GUID syscall.GUID = syscall.GUID{0x1851116a, 0x0e47, 0x4b30, 
	[8]byte{0x8c, 0xb5, 0xd7, 0xda, 0xe4, 0xfb, 0xcd, 0x1b}}
	Annotation_EditingLockedChange_GUID syscall.GUID = syscall.GUID{0xc31f3e1c, 0x7423, 0x4dac, 
	[8]byte{0x83, 0x48, 0x41, 0xf0, 0x99, 0xff, 0x6f, 0x64}}
	Annotation_ExternalChange_GUID syscall.GUID = syscall.GUID{0x75a05b31, 0x5f11, 0x42fd, 
	[8]byte{0x88, 0x7d, 0xdf, 0xa0, 0x10, 0xdb, 0x23, 0x92}}
	Annotation_ConflictingChange_GUID syscall.GUID = syscall.GUID{0x98af8802, 0x517c, 0x459f, 
	[8]byte{0xaf, 0x13, 0x01, 0x6d, 0x3f, 0xab, 0x87, 0x7e}}
	Annotation_Author_GUID syscall.GUID = syscall.GUID{0xf161d3a7, 0xf81b, 0x4128, 
	[8]byte{0xb1, 0x7f, 0x71, 0xf6, 0x90, 0x91, 0x45, 0x20}}
	Annotation_AdvancedProofingIssue_GUID syscall.GUID = syscall.GUID{0xdac7b72c, 0xc0f2, 0x4b84, 
	[8]byte{0xb9, 0x0d, 0x5f, 0xaf, 0xc0, 0xf0, 0xef, 0x1c}}
	Annotation_DataValidationError_GUID syscall.GUID = syscall.GUID{0xc8649fa8, 0x9775, 0x437e, 
	[8]byte{0xad, 0x46, 0xe7, 0x09, 0xd9, 0x3c, 0x23, 0x43}}
	Annotation_CircularReferenceError_GUID syscall.GUID = syscall.GUID{0x25bd9cf4, 0x1745, 0x4659, 
	[8]byte{0xba, 0x67, 0x72, 0x7f, 0x03, 0x18, 0xc6, 0x16}}
	Annotation_Mathematics_GUID syscall.GUID = syscall.GUID{0xeaab634b, 0x26d0, 0x40c1, 
	[8]byte{0x80, 0x73, 0x57, 0xca, 0x1c, 0x63, 0x3c, 0x9b}}
	Annotation_Sensitive_GUID syscall.GUID = syscall.GUID{0x37f4c04f, 0x0f12, 0x4464, 
	[8]byte{0x92, 0x9c, 0x82, 0x8f, 0xd1, 0x52, 0x92, 0xe3}}
	Changes_Summary_GUID syscall.GUID = syscall.GUID{0x313d65a6, 0xe60f, 0x4d62, 
	[8]byte{0x98, 0x61, 0x55, 0xaf, 0xd7, 0x28, 0xd2, 0x07}}
	StyleId_Custom_GUID syscall.GUID = syscall.GUID{0xef2edd3e, 0xa999, 0x4b7c, 
	[8]byte{0xa3, 0x78, 0x09, 0xbb, 0xd5, 0x2a, 0x35, 0x16}}
	StyleId_Heading1_GUID syscall.GUID = syscall.GUID{0x7f7e8f69, 0x6866, 0x4621, 
	[8]byte{0x93, 0x0c, 0x9a, 0x5d, 0x0c, 0xa5, 0x96, 0x1c}}
	StyleId_Heading2_GUID syscall.GUID = syscall.GUID{0xbaa9b241, 0x5c69, 0x469d, 
	[8]byte{0x85, 0xad, 0x47, 0x47, 0x37, 0xb5, 0x2b, 0x14}}
	StyleId_Heading3_GUID syscall.GUID = syscall.GUID{0xbf8be9d2, 0xd8b8, 0x4ec5, 
	[8]byte{0x8c, 0x52, 0x9c, 0xfb, 0x0d, 0x03, 0x59, 0x70}}
	StyleId_Heading4_GUID syscall.GUID = syscall.GUID{0x8436ffc0, 0x9578, 0x45fc, 
	[8]byte{0x83, 0xa4, 0xff, 0x40, 0x05, 0x33, 0x15, 0xdd}}
	StyleId_Heading5_GUID syscall.GUID = syscall.GUID{0x909f424d, 0x0dbf, 0x406e, 
	[8]byte{0x97, 0xbb, 0x4e, 0x77, 0x3d, 0x97, 0x98, 0xf7}}
	StyleId_Heading6_GUID syscall.GUID = syscall.GUID{0x89d23459, 0x5d5b, 0x4824, 
	[8]byte{0xa4, 0x20, 0x11, 0xd3, 0xed, 0x82, 0xe4, 0x0f}}
	StyleId_Heading7_GUID syscall.GUID = syscall.GUID{0xa3790473, 0xe9ae, 0x422d, 
	[8]byte{0xb8, 0xe3, 0x3b, 0x67, 0x5c, 0x61, 0x81, 0xa4}}
	StyleId_Heading8_GUID syscall.GUID = syscall.GUID{0x2bc14145, 0xa40c, 0x4881, 
	[8]byte{0x84, 0xae, 0xf2, 0x23, 0x56, 0x85, 0x38, 0x0c}}
	StyleId_Heading9_GUID syscall.GUID = syscall.GUID{0xc70d9133, 0xbb2a, 0x43d3, 
	[8]byte{0x8a, 0xc6, 0x33, 0x65, 0x78, 0x84, 0xb0, 0xf0}}
	StyleId_Title_GUID syscall.GUID = syscall.GUID{0x15d8201a, 0xffcf, 0x481f, 
	[8]byte{0xb0, 0xa1, 0x30, 0xb6, 0x3b, 0xe9, 0x8f, 0x07}}
	StyleId_Subtitle_GUID syscall.GUID = syscall.GUID{0xb5d9fc17, 0x5d6f, 0x4420, 
	[8]byte{0xb4, 0x39, 0x7c, 0xb1, 0x9a, 0xd4, 0x34, 0xe2}}
	StyleId_Normal_GUID syscall.GUID = syscall.GUID{0xcd14d429, 0xe45e, 0x4475, 
	[8]byte{0xa1, 0xc5, 0x7f, 0x9e, 0x6b, 0xe9, 0x6e, 0xba}}
	StyleId_Emphasis_GUID syscall.GUID = syscall.GUID{0xca6e7dbe, 0x355e, 0x4820, 
	[8]byte{0x95, 0xa0, 0x92, 0x5f, 0x04, 0x1d, 0x34, 0x70}}
	StyleId_Quote_GUID syscall.GUID = syscall.GUID{0x5d1c21ea, 0x8195, 0x4f6c, 
	[8]byte{0x87, 0xea, 0x5d, 0xab, 0xec, 0xe6, 0x4c, 0x1d}}
	StyleId_BulletedList_GUID syscall.GUID = syscall.GUID{0x5963ed64, 0x6426, 0x4632, 
	[8]byte{0x8c, 0xaf, 0xa3, 0x2a, 0xd4, 0x02, 0xd9, 0x1a}}
	StyleId_NumberedList_GUID syscall.GUID = syscall.GUID{0x1e96dbd5, 0x64c3, 0x43d0, 
	[8]byte{0xb1, 0xee, 0xb5, 0x3b, 0x06, 0xe3, 0xed, 0xdf}}
	Notification_Event_GUID syscall.GUID = syscall.GUID{0x72c5a2f7, 0x9788, 0x480f, 
	[8]byte{0xb8, 0xeb, 0x4d, 0xee, 0x00, 0xf6, 0x18, 0x6f}}
	SID_IsUIAutomationObject syscall.GUID = syscall.GUID{0xb96fdb85, 0x7204, 0x4724, 
	[8]byte{0x84, 0x2b, 0xc7, 0x05, 0x9d, 0xed, 0xb9, 0xd0}}
	SID_ControlElementProvider syscall.GUID = syscall.GUID{0xf4791d68, 0xe254, 0x4ba3, 
	[8]byte{0x9a, 0x53, 0x26, 0xa5, 0xc5, 0x49, 0x79, 0x46}}
	IsSelectionPattern2Available_Property_GUID syscall.GUID = syscall.GUID{0x490806fb, 0x6e89, 0x4a47, 
	[8]byte{0x83, 0x19, 0xd2, 0x66, 0xe5, 0x11, 0xf0, 0x21}}
	Selection2_FirstSelectedItem_Property_GUID syscall.GUID = syscall.GUID{0xcc24ea67, 0x369c, 0x4e55, 
	[8]byte{0x9f, 0xf7, 0x38, 0xda, 0x69, 0x54, 0x0c, 0x29}}
	Selection2_LastSelectedItem_Property_GUID syscall.GUID = syscall.GUID{0xcf7bda90, 0x2d83, 0x49f8, 
	[8]byte{0x86, 0x0c, 0x9c, 0xe3, 0x94, 0xcf, 0x89, 0xb4}}
	Selection2_CurrentSelectedItem_Property_GUID syscall.GUID = syscall.GUID{0x34257c26, 0x83b5, 0x41a6, 
	[8]byte{0x93, 0x9c, 0xae, 0x84, 0x1c, 0x13, 0x62, 0x36}}
	Selection2_ItemCount_Property_GUID syscall.GUID = syscall.GUID{0xbb49eb9f, 0x456d, 0x4048, 
	[8]byte{0xb5, 0x91, 0x9c, 0x20, 0x26, 0xb8, 0x46, 0x36}}
	Selection_Pattern2_GUID syscall.GUID = syscall.GUID{0xfba25cab, 0xab98, 0x49f7, 
	[8]byte{0xa7, 0xdc, 0xfe, 0x53, 0x9d, 0xc1, 0x5b, 0xe7}}
	HeadingLevel_Property_GUID syscall.GUID = syscall.GUID{0x29084272, 0xaaaf, 0x4a30, 
	[8]byte{0x87, 0x96, 0x3c, 0x12, 0xf6, 0x2b, 0x6b, 0xbb}}
	IsDialog_Property_GUID syscall.GUID = syscall.GUID{0x9d0dfb9b, 0x8436, 0x4501, 
	[8]byte{0xbb, 0xbb, 0xe5, 0x34, 0xa4, 0xfb, 0x3b, 0x3f}}
)

// enums

// enum STICKYKEYS_FLAGS
// flags
type STICKYKEYS_FLAGS uint32
const (
	SKF_STICKYKEYSON STICKYKEYS_FLAGS = 1
	SKF_AVAILABLE STICKYKEYS_FLAGS = 2
	SKF_HOTKEYACTIVE STICKYKEYS_FLAGS = 4
	SKF_CONFIRMHOTKEY STICKYKEYS_FLAGS = 8
	SKF_HOTKEYSOUND STICKYKEYS_FLAGS = 16
	SKF_INDICATOR STICKYKEYS_FLAGS = 32
	SKF_AUDIBLEFEEDBACK STICKYKEYS_FLAGS = 64
	SKF_TRISTATE STICKYKEYS_FLAGS = 128
	SKF_TWOKEYSOFF STICKYKEYS_FLAGS = 256
	SKF_LALTLATCHED STICKYKEYS_FLAGS = 268435456
	SKF_LCTLLATCHED STICKYKEYS_FLAGS = 67108864
	SKF_LSHIFTLATCHED STICKYKEYS_FLAGS = 16777216
	SKF_RALTLATCHED STICKYKEYS_FLAGS = 536870912
	SKF_RCTLLATCHED STICKYKEYS_FLAGS = 134217728
	SKF_RSHIFTLATCHED STICKYKEYS_FLAGS = 33554432
	SKF_LWINLATCHED STICKYKEYS_FLAGS = 1073741824
	SKF_RWINLATCHED STICKYKEYS_FLAGS = 2147483648
	SKF_LALTLOCKED STICKYKEYS_FLAGS = 1048576
	SKF_LCTLLOCKED STICKYKEYS_FLAGS = 262144
	SKF_LSHIFTLOCKED STICKYKEYS_FLAGS = 65536
	SKF_RALTLOCKED STICKYKEYS_FLAGS = 2097152
	SKF_RCTLLOCKED STICKYKEYS_FLAGS = 524288
	SKF_RSHIFTLOCKED STICKYKEYS_FLAGS = 131072
	SKF_LWINLOCKED STICKYKEYS_FLAGS = 4194304
	SKF_RWINLOCKED STICKYKEYS_FLAGS = 8388608
)

// enum SOUNDSENTRY_FLAGS
// flags
type SOUNDSENTRY_FLAGS uint32
const (
	SSF_SOUNDSENTRYON SOUNDSENTRY_FLAGS = 1
	SSF_AVAILABLE SOUNDSENTRY_FLAGS = 2
	SSF_INDICATOR SOUNDSENTRY_FLAGS = 4
)

// enum ACC_UTILITY_STATE_FLAGS
// flags
type ACC_UTILITY_STATE_FLAGS uint32
const (
	ANRUS_ON_SCREEN_KEYBOARD_ACTIVE ACC_UTILITY_STATE_FLAGS = 1
	ANRUS_TOUCH_MODIFICATION_ACTIVE ACC_UTILITY_STATE_FLAGS = 2
	ANRUS_PRIORITY_AUDIO_ACTIVE ACC_UTILITY_STATE_FLAGS = 4
	ANRUS_PRIORITY_AUDIO_ACTIVE_NODUCK ACC_UTILITY_STATE_FLAGS = 8
)

// enum SOUND_SENTRY_GRAPHICS_EFFECT
type SOUND_SENTRY_GRAPHICS_EFFECT uint32
const (
	SSGF_DISPLAY SOUND_SENTRY_GRAPHICS_EFFECT = 3
	SSGF_NONE SOUND_SENTRY_GRAPHICS_EFFECT = 0
)

// enum SERIALKEYS_FLAGS
// flags
type SERIALKEYS_FLAGS uint32
const (
	SERKF_AVAILABLE SERIALKEYS_FLAGS = 2
	SERKF_INDICATOR SERIALKEYS_FLAGS = 4
	SERKF_SERIALKEYSON SERIALKEYS_FLAGS = 1
)

// enum HIGHCONTRASTW_FLAGS
// flags
type HIGHCONTRASTW_FLAGS uint32
const (
	HCF_HIGHCONTRASTON HIGHCONTRASTW_FLAGS = 1
	HCF_AVAILABLE HIGHCONTRASTW_FLAGS = 2
	HCF_HOTKEYACTIVE HIGHCONTRASTW_FLAGS = 4
	HCF_CONFIRMHOTKEY HIGHCONTRASTW_FLAGS = 8
	HCF_HOTKEYSOUND HIGHCONTRASTW_FLAGS = 16
	HCF_INDICATOR HIGHCONTRASTW_FLAGS = 32
	HCF_HOTKEYAVAILABLE HIGHCONTRASTW_FLAGS = 64
	HCF_OPTION_NOTHEMECHANGE HIGHCONTRASTW_FLAGS = 4096
)

// enum SOUNDSENTRY_TEXT_EFFECT
type SOUNDSENTRY_TEXT_EFFECT uint32
const (
	SSTF_BORDER SOUNDSENTRY_TEXT_EFFECT = 2
	SSTF_CHARS SOUNDSENTRY_TEXT_EFFECT = 1
	SSTF_DISPLAY SOUNDSENTRY_TEXT_EFFECT = 3
	SSTF_NONE SOUNDSENTRY_TEXT_EFFECT = 0
)

// enum SOUNDSENTRY_WINDOWS_EFFECT
type SOUNDSENTRY_WINDOWS_EFFECT uint32
const (
	SSWF_CUSTOM SOUNDSENTRY_WINDOWS_EFFECT = 4
	SSWF_DISPLAY SOUNDSENTRY_WINDOWS_EFFECT = 3
	SSWF_NONE SOUNDSENTRY_WINDOWS_EFFECT = 0
	SSWF_TITLE SOUNDSENTRY_WINDOWS_EFFECT = 1
	SSWF_WINDOW SOUNDSENTRY_WINDOWS_EFFECT = 2
)

// enum AnnoScope
type AnnoScope int32
const (
	ANNO_THIS AnnoScope = 0
	ANNO_CONTAINER AnnoScope = 1
)

// enum NavigateDirection
type NavigateDirection int32
const (
	NavigateDirection_Parent NavigateDirection = 0
	NavigateDirection_NextSibling NavigateDirection = 1
	NavigateDirection_PreviousSibling NavigateDirection = 2
	NavigateDirection_FirstChild NavigateDirection = 3
	NavigateDirection_LastChild NavigateDirection = 4
)

// enum ProviderOptions
type ProviderOptions int32
const (
	ProviderOptions_ClientSideProvider ProviderOptions = 1
	ProviderOptions_ServerSideProvider ProviderOptions = 2
	ProviderOptions_NonClientAreaProvider ProviderOptions = 4
	ProviderOptions_OverrideProvider ProviderOptions = 8
	ProviderOptions_ProviderOwnsSetFocus ProviderOptions = 16
	ProviderOptions_UseComThreading ProviderOptions = 32
	ProviderOptions_RefuseNonClientSupport ProviderOptions = 64
	ProviderOptions_HasNativeIAccessible ProviderOptions = 128
	ProviderOptions_UseClientCoordinates ProviderOptions = 256
)

// enum StructureChangeType
type StructureChangeType int32
const (
	StructureChangeType_ChildAdded StructureChangeType = 0
	StructureChangeType_ChildRemoved StructureChangeType = 1
	StructureChangeType_ChildrenInvalidated StructureChangeType = 2
	StructureChangeType_ChildrenBulkAdded StructureChangeType = 3
	StructureChangeType_ChildrenBulkRemoved StructureChangeType = 4
	StructureChangeType_ChildrenReordered StructureChangeType = 5
)

// enum TextEditChangeType
type TextEditChangeType int32
const (
	TextEditChangeType_None TextEditChangeType = 0
	TextEditChangeType_AutoCorrect TextEditChangeType = 1
	TextEditChangeType_Composition TextEditChangeType = 2
	TextEditChangeType_CompositionFinalized TextEditChangeType = 3
	TextEditChangeType_AutoComplete TextEditChangeType = 4
)

// enum OrientationType
type OrientationType int32
const (
	OrientationType_None OrientationType = 0
	OrientationType_Horizontal OrientationType = 1
	OrientationType_Vertical OrientationType = 2
)

// enum DockPosition
type DockPosition int32
const (
	DockPosition_Top DockPosition = 0
	DockPosition_Left DockPosition = 1
	DockPosition_Bottom DockPosition = 2
	DockPosition_Right DockPosition = 3
	DockPosition_Fill DockPosition = 4
	DockPosition_None DockPosition = 5
)

// enum ExpandCollapseState
type ExpandCollapseState int32
const (
	ExpandCollapseState_Collapsed ExpandCollapseState = 0
	ExpandCollapseState_Expanded ExpandCollapseState = 1
	ExpandCollapseState_PartiallyExpanded ExpandCollapseState = 2
	ExpandCollapseState_LeafNode ExpandCollapseState = 3
)

// enum ScrollAmount
type ScrollAmount int32
const (
	ScrollAmount_LargeDecrement ScrollAmount = 0
	ScrollAmount_SmallDecrement ScrollAmount = 1
	ScrollAmount_NoAmount ScrollAmount = 2
	ScrollAmount_LargeIncrement ScrollAmount = 3
	ScrollAmount_SmallIncrement ScrollAmount = 4
)

// enum RowOrColumnMajor
type RowOrColumnMajor int32
const (
	RowOrColumnMajor_RowMajor RowOrColumnMajor = 0
	RowOrColumnMajor_ColumnMajor RowOrColumnMajor = 1
	RowOrColumnMajor_Indeterminate RowOrColumnMajor = 2
)

// enum ToggleState
type ToggleState int32
const (
	ToggleState_Off ToggleState = 0
	ToggleState_On ToggleState = 1
	ToggleState_Indeterminate ToggleState = 2
)

// enum WindowVisualState
type WindowVisualState int32
const (
	WindowVisualState_Normal WindowVisualState = 0
	WindowVisualState_Maximized WindowVisualState = 1
	WindowVisualState_Minimized WindowVisualState = 2
)

// enum SynchronizedInputType
type SynchronizedInputType int32
const (
	SynchronizedInputType_KeyUp SynchronizedInputType = 1
	SynchronizedInputType_KeyDown SynchronizedInputType = 2
	SynchronizedInputType_LeftMouseUp SynchronizedInputType = 4
	SynchronizedInputType_LeftMouseDown SynchronizedInputType = 8
	SynchronizedInputType_RightMouseUp SynchronizedInputType = 16
	SynchronizedInputType_RightMouseDown SynchronizedInputType = 32
)

// enum WindowInteractionState
type WindowInteractionState int32
const (
	WindowInteractionState_Running WindowInteractionState = 0
	WindowInteractionState_Closing WindowInteractionState = 1
	WindowInteractionState_ReadyForUserInteraction WindowInteractionState = 2
	WindowInteractionState_BlockedByModalWindow WindowInteractionState = 3
	WindowInteractionState_NotResponding WindowInteractionState = 4
)

// enum SayAsInterpretAs
type SayAsInterpretAs int32
const (
	SayAsInterpretAs_None SayAsInterpretAs = 0
	SayAsInterpretAs_Spell SayAsInterpretAs = 1
	SayAsInterpretAs_Cardinal SayAsInterpretAs = 2
	SayAsInterpretAs_Ordinal SayAsInterpretAs = 3
	SayAsInterpretAs_Number SayAsInterpretAs = 4
	SayAsInterpretAs_Date SayAsInterpretAs = 5
	SayAsInterpretAs_Time SayAsInterpretAs = 6
	SayAsInterpretAs_Telephone SayAsInterpretAs = 7
	SayAsInterpretAs_Currency SayAsInterpretAs = 8
	SayAsInterpretAs_Net SayAsInterpretAs = 9
	SayAsInterpretAs_Url SayAsInterpretAs = 10
	SayAsInterpretAs_Address SayAsInterpretAs = 11
	SayAsInterpretAs_Alphanumeric SayAsInterpretAs = 12
	SayAsInterpretAs_Name SayAsInterpretAs = 13
	SayAsInterpretAs_Media SayAsInterpretAs = 14
	SayAsInterpretAs_Date_MonthDayYear SayAsInterpretAs = 15
	SayAsInterpretAs_Date_DayMonthYear SayAsInterpretAs = 16
	SayAsInterpretAs_Date_YearMonthDay SayAsInterpretAs = 17
	SayAsInterpretAs_Date_YearMonth SayAsInterpretAs = 18
	SayAsInterpretAs_Date_MonthYear SayAsInterpretAs = 19
	SayAsInterpretAs_Date_DayMonth SayAsInterpretAs = 20
	SayAsInterpretAs_Date_MonthDay SayAsInterpretAs = 21
	SayAsInterpretAs_Date_Year SayAsInterpretAs = 22
	SayAsInterpretAs_Time_HoursMinutesSeconds12 SayAsInterpretAs = 23
	SayAsInterpretAs_Time_HoursMinutes12 SayAsInterpretAs = 24
	SayAsInterpretAs_Time_HoursMinutesSeconds24 SayAsInterpretAs = 25
	SayAsInterpretAs_Time_HoursMinutes24 SayAsInterpretAs = 26
)

// enum TextUnit
type TextUnit int32
const (
	TextUnit_Character TextUnit = 0
	TextUnit_Format TextUnit = 1
	TextUnit_Word TextUnit = 2
	TextUnit_Line TextUnit = 3
	TextUnit_Paragraph TextUnit = 4
	TextUnit_Page TextUnit = 5
	TextUnit_Document TextUnit = 6
)

// enum TextPatternRangeEndpoint
type TextPatternRangeEndpoint int32
const (
	TextPatternRangeEndpoint_Start TextPatternRangeEndpoint = 0
	TextPatternRangeEndpoint_End TextPatternRangeEndpoint = 1
)

// enum SupportedTextSelection
type SupportedTextSelection int32
const (
	SupportedTextSelection_None SupportedTextSelection = 0
	SupportedTextSelection_Single SupportedTextSelection = 1
	SupportedTextSelection_Multiple SupportedTextSelection = 2
)

// enum LiveSetting
type LiveSetting int32
const (
	Off LiveSetting = 0
	Polite LiveSetting = 1
	Assertive LiveSetting = 2
)

// enum ActiveEnd
type ActiveEnd int32
const (
	ActiveEnd_None ActiveEnd = 0
	ActiveEnd_Start ActiveEnd = 1
	ActiveEnd_End ActiveEnd = 2
)

// enum CaretPosition
type CaretPosition int32
const (
	CaretPosition_Unknown CaretPosition = 0
	CaretPosition_EndOfLine CaretPosition = 1
	CaretPosition_BeginningOfLine CaretPosition = 2
)

// enum CaretBidiMode
type CaretBidiMode int32
const (
	CaretBidiMode_LTR CaretBidiMode = 0
	CaretBidiMode_RTL CaretBidiMode = 1
)

// enum ZoomUnit
type ZoomUnit int32
const (
	ZoomUnit_NoAmount ZoomUnit = 0
	ZoomUnit_LargeDecrement ZoomUnit = 1
	ZoomUnit_SmallDecrement ZoomUnit = 2
	ZoomUnit_LargeIncrement ZoomUnit = 3
	ZoomUnit_SmallIncrement ZoomUnit = 4
)

// enum AnimationStyle
type AnimationStyle int32
const (
	AnimationStyle_None AnimationStyle = 0
	AnimationStyle_LasVegasLights AnimationStyle = 1
	AnimationStyle_BlinkingBackground AnimationStyle = 2
	AnimationStyle_SparkleText AnimationStyle = 3
	AnimationStyle_MarchingBlackAnts AnimationStyle = 4
	AnimationStyle_MarchingRedAnts AnimationStyle = 5
	AnimationStyle_Shimmer AnimationStyle = 6
	AnimationStyle_Other AnimationStyle = -1
)

// enum BulletStyle
type BulletStyle int32
const (
	BulletStyle_None BulletStyle = 0
	BulletStyle_HollowRoundBullet BulletStyle = 1
	BulletStyle_FilledRoundBullet BulletStyle = 2
	BulletStyle_HollowSquareBullet BulletStyle = 3
	BulletStyle_FilledSquareBullet BulletStyle = 4
	BulletStyle_DashBullet BulletStyle = 5
	BulletStyle_Other BulletStyle = -1
)

// enum CapStyle
type CapStyle int32
const (
	CapStyle_None CapStyle = 0
	CapStyle_SmallCap CapStyle = 1
	CapStyle_AllCap CapStyle = 2
	CapStyle_AllPetiteCaps CapStyle = 3
	CapStyle_PetiteCaps CapStyle = 4
	CapStyle_Unicase CapStyle = 5
	CapStyle_Titling CapStyle = 6
	CapStyle_Other CapStyle = -1
)

// enum FillType
type FillType int32
const (
	FillType_None FillType = 0
	FillType_Color FillType = 1
	FillType_Gradient FillType = 2
	FillType_Picture FillType = 3
	FillType_Pattern FillType = 4
)

// enum FlowDirections
type FlowDirections int32
const (
	FlowDirections_Default FlowDirections = 0
	FlowDirections_RightToLeft FlowDirections = 1
	FlowDirections_BottomToTop FlowDirections = 2
	FlowDirections_Vertical FlowDirections = 4
)

// enum HorizontalTextAlignment
type HorizontalTextAlignment int32
const (
	HorizontalTextAlignment_Left HorizontalTextAlignment = 0
	HorizontalTextAlignment_Centered HorizontalTextAlignment = 1
	HorizontalTextAlignment_Right HorizontalTextAlignment = 2
	HorizontalTextAlignment_Justified HorizontalTextAlignment = 3
)

// enum OutlineStyles
type OutlineStyles int32
const (
	OutlineStyles_None OutlineStyles = 0
	OutlineStyles_Outline OutlineStyles = 1
	OutlineStyles_Shadow OutlineStyles = 2
	OutlineStyles_Engraved OutlineStyles = 4
	OutlineStyles_Embossed OutlineStyles = 8
)

// enum TextDecorationLineStyle
type TextDecorationLineStyle int32
const (
	TextDecorationLineStyle_None TextDecorationLineStyle = 0
	TextDecorationLineStyle_Single TextDecorationLineStyle = 1
	TextDecorationLineStyle_WordsOnly TextDecorationLineStyle = 2
	TextDecorationLineStyle_Double TextDecorationLineStyle = 3
	TextDecorationLineStyle_Dot TextDecorationLineStyle = 4
	TextDecorationLineStyle_Dash TextDecorationLineStyle = 5
	TextDecorationLineStyle_DashDot TextDecorationLineStyle = 6
	TextDecorationLineStyle_DashDotDot TextDecorationLineStyle = 7
	TextDecorationLineStyle_Wavy TextDecorationLineStyle = 8
	TextDecorationLineStyle_ThickSingle TextDecorationLineStyle = 9
	TextDecorationLineStyle_DoubleWavy TextDecorationLineStyle = 11
	TextDecorationLineStyle_ThickWavy TextDecorationLineStyle = 12
	TextDecorationLineStyle_LongDash TextDecorationLineStyle = 13
	TextDecorationLineStyle_ThickDash TextDecorationLineStyle = 14
	TextDecorationLineStyle_ThickDashDot TextDecorationLineStyle = 15
	TextDecorationLineStyle_ThickDashDotDot TextDecorationLineStyle = 16
	TextDecorationLineStyle_ThickDot TextDecorationLineStyle = 17
	TextDecorationLineStyle_ThickLongDash TextDecorationLineStyle = 18
	TextDecorationLineStyle_Other TextDecorationLineStyle = -1
)

// enum VisualEffects
type VisualEffects int32
const (
	VisualEffects_None VisualEffects = 0
	VisualEffects_Shadow VisualEffects = 1
	VisualEffects_Reflection VisualEffects = 2
	VisualEffects_Glow VisualEffects = 4
	VisualEffects_SoftEdges VisualEffects = 8
	VisualEffects_Bevel VisualEffects = 16
)

// enum NotificationProcessing
type NotificationProcessing int32
const (
	NotificationProcessing_ImportantAll NotificationProcessing = 0
	NotificationProcessing_ImportantMostRecent NotificationProcessing = 1
	NotificationProcessing_All NotificationProcessing = 2
	NotificationProcessing_MostRecent NotificationProcessing = 3
	NotificationProcessing_CurrentThenMostRecent NotificationProcessing = 4
)

// enum NotificationKind
type NotificationKind int32
const (
	NotificationKind_ItemAdded NotificationKind = 0
	NotificationKind_ItemRemoved NotificationKind = 1
	NotificationKind_ActionCompleted NotificationKind = 2
	NotificationKind_ActionAborted NotificationKind = 3
	NotificationKind_Other NotificationKind = 4
)

// enum UIAutomationType
type UIAutomationType int32
const (
	UIAutomationType_Int UIAutomationType = 1
	UIAutomationType_Bool UIAutomationType = 2
	UIAutomationType_String UIAutomationType = 3
	UIAutomationType_Double UIAutomationType = 4
	UIAutomationType_Point UIAutomationType = 5
	UIAutomationType_Rect UIAutomationType = 6
	UIAutomationType_Element UIAutomationType = 7
	UIAutomationType_Array UIAutomationType = 65536
	UIAutomationType_Out UIAutomationType = 131072
	UIAutomationType_IntArray UIAutomationType = 65537
	UIAutomationType_BoolArray UIAutomationType = 65538
	UIAutomationType_StringArray UIAutomationType = 65539
	UIAutomationType_DoubleArray UIAutomationType = 65540
	UIAutomationType_PointArray UIAutomationType = 65541
	UIAutomationType_RectArray UIAutomationType = 65542
	UIAutomationType_ElementArray UIAutomationType = 65543
	UIAutomationType_OutInt UIAutomationType = 131073
	UIAutomationType_OutBool UIAutomationType = 131074
	UIAutomationType_OutString UIAutomationType = 131075
	UIAutomationType_OutDouble UIAutomationType = 131076
	UIAutomationType_OutPoint UIAutomationType = 131077
	UIAutomationType_OutRect UIAutomationType = 131078
	UIAutomationType_OutElement UIAutomationType = 131079
	UIAutomationType_OutIntArray UIAutomationType = 196609
	UIAutomationType_OutBoolArray UIAutomationType = 196610
	UIAutomationType_OutStringArray UIAutomationType = 196611
	UIAutomationType_OutDoubleArray UIAutomationType = 196612
	UIAutomationType_OutPointArray UIAutomationType = 196613
	UIAutomationType_OutRectArray UIAutomationType = 196614
	UIAutomationType_OutElementArray UIAutomationType = 196615
)

// enum TreeScope
type TreeScope int32
const (
	TreeScope_None TreeScope = 0
	TreeScope_Element TreeScope = 1
	TreeScope_Children TreeScope = 2
	TreeScope_Descendants TreeScope = 4
	TreeScope_Parent TreeScope = 8
	TreeScope_Ancestors TreeScope = 16
	TreeScope_Subtree TreeScope = 7
)

// enum PropertyConditionFlags
type PropertyConditionFlags int32
const (
	PropertyConditionFlags_None PropertyConditionFlags = 0
	PropertyConditionFlags_IgnoreCase PropertyConditionFlags = 1
	PropertyConditionFlags_MatchSubstring PropertyConditionFlags = 2
)

// enum AutomationElementMode
type AutomationElementMode int32
const (
	AutomationElementMode_None AutomationElementMode = 0
	AutomationElementMode_Full AutomationElementMode = 1
)

// enum TreeTraversalOptions
type TreeTraversalOptions int32
const (
	TreeTraversalOptions_Default TreeTraversalOptions = 0
	TreeTraversalOptions_PostOrder TreeTraversalOptions = 1
	TreeTraversalOptions_LastToFirstOrder TreeTraversalOptions = 2
)

// enum ConnectionRecoveryBehaviorOptions
type ConnectionRecoveryBehaviorOptions int32
const (
	ConnectionRecoveryBehaviorOptions_Disabled ConnectionRecoveryBehaviorOptions = 0
	ConnectionRecoveryBehaviorOptions_Enabled ConnectionRecoveryBehaviorOptions = 1
)

// enum CoalesceEventsOptions
type CoalesceEventsOptions int32
const (
	CoalesceEventsOptions_Disabled CoalesceEventsOptions = 0
	CoalesceEventsOptions_Enabled CoalesceEventsOptions = 1
)

// enum ConditionType
type ConditionType int32
const (
	ConditionType_True ConditionType = 0
	ConditionType_False ConditionType = 1
	ConditionType_Property ConditionType = 2
	ConditionType_And ConditionType = 3
	ConditionType_Or ConditionType = 4
	ConditionType_Not ConditionType = 5
)

// enum NormalizeState
type NormalizeState int32
const (
	NormalizeState_None NormalizeState = 0
	NormalizeState_View NormalizeState = 1
	NormalizeState_Custom NormalizeState = 2
)

// enum ProviderType
type ProviderType int32
const (
	ProviderType_BaseHwnd ProviderType = 0
	ProviderType_Proxy ProviderType = 1
	ProviderType_NonClientArea ProviderType = 2
)

// enum AutomationIdentifierType
type AutomationIdentifierType int32
const (
	AutomationIdentifierType_Property AutomationIdentifierType = 0
	AutomationIdentifierType_Pattern AutomationIdentifierType = 1
	AutomationIdentifierType_Event AutomationIdentifierType = 2
	AutomationIdentifierType_ControlType AutomationIdentifierType = 3
	AutomationIdentifierType_TextAttribute AutomationIdentifierType = 4
	AutomationIdentifierType_LandmarkType AutomationIdentifierType = 5
	AutomationIdentifierType_Annotation AutomationIdentifierType = 6
	AutomationIdentifierType_Changes AutomationIdentifierType = 7
	AutomationIdentifierType_Style AutomationIdentifierType = 8
)

// enum EventArgsType
type EventArgsType int32
const (
	EventArgsType_Simple EventArgsType = 0
	EventArgsType_PropertyChanged EventArgsType = 1
	EventArgsType_StructureChanged EventArgsType = 2
	EventArgsType_AsyncContentLoaded EventArgsType = 3
	EventArgsType_WindowClosed EventArgsType = 4
	EventArgsType_TextEditTextChanged EventArgsType = 5
	EventArgsType_Changes EventArgsType = 6
	EventArgsType_Notification EventArgsType = 7
	EventArgsType_ActiveTextPositionChanged EventArgsType = 8
	EventArgsType_StructuredMarkup EventArgsType = 9
)

// enum AsyncContentLoadedState
type AsyncContentLoadedState int32
const (
	AsyncContentLoadedState_Beginning AsyncContentLoadedState = 0
	AsyncContentLoadedState_Progress AsyncContentLoadedState = 1
	AsyncContentLoadedState_Completed AsyncContentLoadedState = 2
)


// structs

type MSAAMENUINFO struct {
	DwMSAASignature uint32
	CchWText uint32
	PszWText PWSTR
}

type UiaRect struct {
	Left float64
	Top float64
	Width float64
	Height float64
}

type UiaPoint struct {
	X float64
	Y float64
}

type UiaChangeInfo struct {
	UiaId int32
	Payload VARIANT
	ExtraInfo VARIANT
}

type UIAutomationParameter struct {
	Type UIAutomationType
	PData unsafe.Pointer
}

type UIAutomationPropertyInfo struct {
	Guid syscall.GUID
	PProgrammaticName PWSTR
	Type UIAutomationType
}

type UIAutomationEventInfo struct {
	Guid syscall.GUID
	PProgrammaticName PWSTR
}

type UIAutomationMethodInfo struct {
	PProgrammaticName PWSTR
	DoSetFocus BOOL
	CInParameters uint32
	COutParameters uint32
	PParameterTypes *UIAutomationType
	PParameterNames *PWSTR
}

type UIAutomationPatternInfo struct {
	Guid syscall.GUID
	PProgrammaticName PWSTR
	ProviderInterfaceId syscall.GUID
	ClientInterfaceId syscall.GUID
	CProperties uint32
	PProperties *UIAutomationPropertyInfo
	CMethods uint32
	PMethods *UIAutomationMethodInfo
	CEvents uint32
	PEvents *UIAutomationEventInfo
	PPatternHandler *IUIAutomationPatternHandler
}

type ExtendedProperty struct {
	PropertyName BSTR
	PropertyValue BSTR
}

type UiaCondition struct {
	ConditionType ConditionType
}

type UiaPropertyCondition struct {
	ConditionType ConditionType
	PropertyId int32
	Value VARIANT
	Flags PropertyConditionFlags
}

type UiaAndOrCondition struct {
	ConditionType ConditionType
	PpConditions **UiaCondition
	CConditions int32
}

type UiaNotCondition struct {
	ConditionType ConditionType
	PCondition *UiaCondition
}

type UiaCacheRequest struct {
	PViewCondition *UiaCondition
	Scope TreeScope
	PProperties *int32
	CProperties int32
	PPatterns *int32
	CPatterns int32
	AutomationElementMode AutomationElementMode
}

type UiaFindParams struct {
	MaxDepth int32
	FindFirst BOOL
	ExcludeRoot BOOL
	PFindCondition *UiaCondition
}

type UiaEventArgs struct {
	Type EventArgsType
	EventId int32
}

type UiaPropertyChangedEventArgs struct {
	Type EventArgsType
	EventId int32
	PropertyId int32
	OldValue VARIANT
	NewValue VARIANT
}

type UiaStructureChangedEventArgs struct {
	Type EventArgsType
	EventId int32
	StructureChangeType StructureChangeType
	PRuntimeId *int32
	CRuntimeIdLen int32
}

type UiaTextEditTextChangedEventArgs struct {
	Type EventArgsType
	EventId int32
	TextEditChangeType TextEditChangeType
	PTextChange *SAFEARRAY
}

type UiaChangesEventArgs struct {
	Type EventArgsType
	EventId int32
	EventIdCount int32
	PUiaChanges *UiaChangeInfo
}

type UiaAsyncContentLoadedEventArgs struct {
	Type EventArgsType
	EventId int32
	AsyncContentLoadedState AsyncContentLoadedState
	PercentComplete float64
}

type UiaWindowClosedEventArgs struct {
	Type EventArgsType
	EventId int32
	PRuntimeId *int32
	CRuntimeIdLen int32
}

type SERIALKEYSA struct {
	CbSize uint32
	DwFlags SERIALKEYS_FLAGS
	LpszActivePort PSTR
	LpszPort PSTR
	IBaudRate uint32
	IPortState uint32
	IActive uint32
}

type SERIALKEYS = SERIALKEYSW
type SERIALKEYSW struct {
	CbSize uint32
	DwFlags SERIALKEYS_FLAGS
	LpszActivePort PWSTR
	LpszPort PWSTR
	IBaudRate uint32
	IPortState uint32
	IActive uint32
}

type HIGHCONTRASTA struct {
	CbSize uint32
	DwFlags HIGHCONTRASTW_FLAGS
	LpszDefaultScheme PSTR
}

type HIGHCONTRAST = HIGHCONTRASTW
type HIGHCONTRASTW struct {
	CbSize uint32
	DwFlags HIGHCONTRASTW_FLAGS
	LpszDefaultScheme PWSTR
}

type FILTERKEYS struct {
	CbSize uint32
	DwFlags uint32
	IWaitMSec uint32
	IDelayMSec uint32
	IRepeatMSec uint32
	IBounceMSec uint32
}

type STICKYKEYS struct {
	CbSize uint32
	DwFlags STICKYKEYS_FLAGS
}

type MOUSEKEYS struct {
	CbSize uint32
	DwFlags uint32
	IMaxSpeed uint32
	ITimeToMaxSpeed uint32
	ICtrlSpeed uint32
	DwReserved1 uint32
	DwReserved2 uint32
}

type ACCESSTIMEOUT struct {
	CbSize uint32
	DwFlags uint32
	ITimeOutMSec uint32
}

type SOUNDSENTRYA struct {
	CbSize uint32
	DwFlags SOUNDSENTRY_FLAGS
	IFSTextEffect SOUNDSENTRY_TEXT_EFFECT
	IFSTextEffectMSec uint32
	IFSTextEffectColorBits uint32
	IFSGrafEffect SOUND_SENTRY_GRAPHICS_EFFECT
	IFSGrafEffectMSec uint32
	IFSGrafEffectColor uint32
	IWindowsEffect SOUNDSENTRY_WINDOWS_EFFECT
	IWindowsEffectMSec uint32
	LpszWindowsEffectDLL PSTR
	IWindowsEffectOrdinal uint32
}

type SOUNDSENTRY = SOUNDSENTRYW
type SOUNDSENTRYW struct {
	CbSize uint32
	DwFlags SOUNDSENTRY_FLAGS
	IFSTextEffect SOUNDSENTRY_TEXT_EFFECT
	IFSTextEffectMSec uint32
	IFSTextEffectColorBits uint32
	IFSGrafEffect SOUND_SENTRY_GRAPHICS_EFFECT
	IFSGrafEffectMSec uint32
	IFSGrafEffectColor uint32
	IWindowsEffect SOUNDSENTRY_WINDOWS_EFFECT
	IWindowsEffectMSec uint32
	LpszWindowsEffectDLL PWSTR
	IWindowsEffectOrdinal uint32
}

type TOGGLEKEYS struct {
	CbSize uint32
	DwFlags uint32
}


// func types

type LPFNLRESULTFROMOBJECT func(riid *syscall.GUID, wParam WPARAM, punk *IUnknown) LRESULT

type LPFNOBJECTFROMLRESULT func(lResult LRESULT, riid *syscall.GUID, wParam WPARAM, ppvObject unsafe.Pointer) HRESULT

type LPFNACCESSIBLEOBJECTFROMWINDOW func(hwnd HWND, dwId uint32, riid *syscall.GUID, ppvObject unsafe.Pointer) HRESULT

type LPFNACCESSIBLEOBJECTFROMPOINT func(ptScreen POINT, ppacc **IAccessible, pvarChild *VARIANT) HRESULT

type LPFNCREATESTDACCESSIBLEOBJECT func(hwnd HWND, idObject int32, riid *syscall.GUID, ppvObject unsafe.Pointer) HRESULT

type LPFNACCESSIBLECHILDREN func(paccContainer *IAccessible, iChildStart int32, cChildren int32, rgvarChildren *VARIANT, pcObtained *int32) HRESULT

type UiaProviderCallback func(hwnd HWND, providerType ProviderType) *SAFEARRAY

type UiaEventCallback func(pArgs *UiaEventArgs, pRequestedData *SAFEARRAY, pTreeStructure BSTR)

type WINEVENTPROC func(hWinEventHook HWINEVENTHOOK, event uint32, hwnd HWND, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32)


// coms

type IRicheditWindowlessAccessibilityInterface interface {
	IUnknownInterface
	CreateProvider(pSite *IRawElementProviderWindowlessSite, ppProvider **IRawElementProviderSimple) HRESULT
}

type IRicheditWindowlessAccessibilityVtbl struct {
	IUnknownVtbl
	CreateProvider uintptr
}

type IRicheditWindowlessAccessibility struct {
	IUnknown
}

func (this *IRicheditWindowlessAccessibility) Vtbl() *IRicheditWindowlessAccessibilityVtbl {
	return (*IRicheditWindowlessAccessibilityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRicheditWindowlessAccessibility) CreateProvider(pSite *IRawElementProviderWindowlessSite, ppProvider **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pSite)), uintptr(unsafe.Pointer(ppProvider)))
	return HRESULT(ret)
}

type IRichEditUiaInformationInterface interface {
	IUnknownInterface
	GetBoundaryRectangle(pUiaRect *UiaRect) HRESULT
	IsVisible() HRESULT
}

type IRichEditUiaInformationVtbl struct {
	IUnknownVtbl
	GetBoundaryRectangle uintptr
	IsVisible uintptr
}

type IRichEditUiaInformation struct {
	IUnknown
}

func (this *IRichEditUiaInformation) Vtbl() *IRichEditUiaInformationVtbl {
	return (*IRichEditUiaInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRichEditUiaInformation) GetBoundaryRectangle(pUiaRect *UiaRect) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetBoundaryRectangle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pUiaRect)))
	return HRESULT(ret)
}

func (this *IRichEditUiaInformation) IsVisible() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IsVisible, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 618736e0-3c3d-11cf-810c-00aa00389b71
var IID_IAccessible = syscall.GUID{0x618736e0, 0x3c3d, 0x11cf, 
	[8]byte{0x81, 0x0c, 0x00, 0xaa, 0x00, 0x38, 0x9b, 0x71}}

type IAccessibleInterface interface {
	IDispatchInterface
	Get_accParent(ppdispParent **IDispatch) HRESULT
	Get_accChildCount(pcountChildren *int32) HRESULT
	Get_accChild(varChild VARIANT, ppdispChild **IDispatch) HRESULT
	Get_accName(varChild VARIANT, pszName *BSTR) HRESULT
	Get_accValue(varChild VARIANT, pszValue *BSTR) HRESULT
	Get_accDescription(varChild VARIANT, pszDescription *BSTR) HRESULT
	Get_accRole(varChild VARIANT, pvarRole *VARIANT) HRESULT
	Get_accState(varChild VARIANT, pvarState *VARIANT) HRESULT
	Get_accHelp(varChild VARIANT, pszHelp *BSTR) HRESULT
	Get_accHelpTopic(pszHelpFile *BSTR, varChild VARIANT, pidTopic *int32) HRESULT
	Get_accKeyboardShortcut(varChild VARIANT, pszKeyboardShortcut *BSTR) HRESULT
	Get_accFocus(pvarChild *VARIANT) HRESULT
	Get_accSelection(pvarChildren *VARIANT) HRESULT
	Get_accDefaultAction(varChild VARIANT, pszDefaultAction *BSTR) HRESULT
	AccSelect(flagsSelect int32, varChild VARIANT) HRESULT
	AccLocation(pxLeft *int32, pyTop *int32, pcxWidth *int32, pcyHeight *int32, varChild VARIANT) HRESULT
	AccNavigate(navDir int32, varStart VARIANT, pvarEndUpAt *VARIANT) HRESULT
	AccHitTest(xLeft int32, yTop int32, pvarChild *VARIANT) HRESULT
	AccDoDefaultAction(varChild VARIANT) HRESULT
	Put_accName(varChild VARIANT, szName BSTR) HRESULT
	Put_accValue(varChild VARIANT, szValue BSTR) HRESULT
}

type IAccessibleVtbl struct {
	IDispatchVtbl
	Get_accParent uintptr
	Get_accChildCount uintptr
	Get_accChild uintptr
	Get_accName uintptr
	Get_accValue uintptr
	Get_accDescription uintptr
	Get_accRole uintptr
	Get_accState uintptr
	Get_accHelp uintptr
	Get_accHelpTopic uintptr
	Get_accKeyboardShortcut uintptr
	Get_accFocus uintptr
	Get_accSelection uintptr
	Get_accDefaultAction uintptr
	AccSelect uintptr
	AccLocation uintptr
	AccNavigate uintptr
	AccHitTest uintptr
	AccDoDefaultAction uintptr
	Put_accName uintptr
	Put_accValue uintptr
}

type IAccessible struct {
	IDispatch
}

func (this *IAccessible) Vtbl() *IAccessibleVtbl {
	return (*IAccessibleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccessible) Get_accParent(ppdispParent **IDispatch) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accParent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppdispParent)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accChildCount(pcountChildren *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accChildCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pcountChildren)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accChild(varChild VARIANT, ppdispChild **IDispatch) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accChild, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(ppdispChild)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accName(varChild VARIANT, pszName *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accName, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pszName)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accValue(varChild VARIANT, pszValue *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accValue, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pszValue)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accDescription(varChild VARIANT, pszDescription *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accDescription, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pszDescription)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accRole(varChild VARIANT, pvarRole *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accRole, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pvarRole)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accState(varChild VARIANT, pvarState *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accState, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pvarState)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accHelp(varChild VARIANT, pszHelp *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accHelp, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pszHelp)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accHelpTopic(pszHelpFile *BSTR, varChild VARIANT, pidTopic *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accHelpTopic, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszHelpFile)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pidTopic)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accKeyboardShortcut(varChild VARIANT, pszKeyboardShortcut *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accKeyboardShortcut, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pszKeyboardShortcut)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accFocus(pvarChild *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accFocus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pvarChild)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accSelection(pvarChildren *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pvarChildren)))
	return HRESULT(ret)
}

func (this *IAccessible) Get_accDefaultAction(varChild VARIANT, pszDefaultAction *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_accDefaultAction, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(pszDefaultAction)))
	return HRESULT(ret)
}

func (this *IAccessible) AccSelect(flagsSelect int32, varChild VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AccSelect, uintptr(unsafe.Pointer(this)), uintptr(flagsSelect), (uintptr)(unsafe.Pointer(&varChild)))
	return HRESULT(ret)
}

func (this *IAccessible) AccLocation(pxLeft *int32, pyTop *int32, pcxWidth *int32, pcyHeight *int32, varChild VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AccLocation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pxLeft)), uintptr(unsafe.Pointer(pyTop)), uintptr(unsafe.Pointer(pcxWidth)), uintptr(unsafe.Pointer(pcyHeight)), (uintptr)(unsafe.Pointer(&varChild)))
	return HRESULT(ret)
}

func (this *IAccessible) AccNavigate(navDir int32, varStart VARIANT, pvarEndUpAt *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AccNavigate, uintptr(unsafe.Pointer(this)), uintptr(navDir), (uintptr)(unsafe.Pointer(&varStart)), uintptr(unsafe.Pointer(pvarEndUpAt)))
	return HRESULT(ret)
}

func (this *IAccessible) AccHitTest(xLeft int32, yTop int32, pvarChild *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AccHitTest, uintptr(unsafe.Pointer(this)), uintptr(xLeft), uintptr(yTop), uintptr(unsafe.Pointer(pvarChild)))
	return HRESULT(ret)
}

func (this *IAccessible) AccDoDefaultAction(varChild VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AccDoDefaultAction, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)))
	return HRESULT(ret)
}

func (this *IAccessible) Put_accName(varChild VARIANT, szName BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_accName, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(szName)))
	return HRESULT(ret)
}

func (this *IAccessible) Put_accValue(varChild VARIANT, szValue BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_accValue, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&varChild)), uintptr(unsafe.Pointer(szValue)))
	return HRESULT(ret)
}

// 03022430-abc4-11d0-bde2-00aa001a1953
var IID_IAccessibleHandler = syscall.GUID{0x03022430, 0xabc4, 0x11d0, 
	[8]byte{0xbd, 0xe2, 0x00, 0xaa, 0x00, 0x1a, 0x19, 0x53}}

type IAccessibleHandlerInterface interface {
	IUnknownInterface
	AccessibleObjectFromID(hwnd int32, lObjectID int32, pIAccessible **IAccessible) HRESULT
}

type IAccessibleHandlerVtbl struct {
	IUnknownVtbl
	AccessibleObjectFromID uintptr
}

type IAccessibleHandler struct {
	IUnknown
}

func (this *IAccessibleHandler) Vtbl() *IAccessibleHandlerVtbl {
	return (*IAccessibleHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccessibleHandler) AccessibleObjectFromID(hwnd int32, lObjectID int32, pIAccessible **IAccessible) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AccessibleObjectFromID, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(lObjectID), uintptr(unsafe.Pointer(pIAccessible)))
	return HRESULT(ret)
}

// bf3abd9c-76da-4389-9eb6-1427d25abab7
var IID_IAccessibleWindowlessSite = syscall.GUID{0xbf3abd9c, 0x76da, 0x4389, 
	[8]byte{0x9e, 0xb6, 0x14, 0x27, 0xd2, 0x5a, 0xba, 0xb7}}

type IAccessibleWindowlessSiteInterface interface {
	IUnknownInterface
	AcquireObjectIdRange(rangeSize int32, pRangeOwner *IAccessibleHandler, pRangeBase *int32) HRESULT
	ReleaseObjectIdRange(rangeBase int32, pRangeOwner *IAccessibleHandler) HRESULT
	QueryObjectIdRanges(pRangesOwner *IAccessibleHandler, psaRanges **SAFEARRAY) HRESULT
	GetParentAccessible(ppParent **IAccessible) HRESULT
}

type IAccessibleWindowlessSiteVtbl struct {
	IUnknownVtbl
	AcquireObjectIdRange uintptr
	ReleaseObjectIdRange uintptr
	QueryObjectIdRanges uintptr
	GetParentAccessible uintptr
}

type IAccessibleWindowlessSite struct {
	IUnknown
}

func (this *IAccessibleWindowlessSite) Vtbl() *IAccessibleWindowlessSiteVtbl {
	return (*IAccessibleWindowlessSiteVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccessibleWindowlessSite) AcquireObjectIdRange(rangeSize int32, pRangeOwner *IAccessibleHandler, pRangeBase *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AcquireObjectIdRange, uintptr(unsafe.Pointer(this)), uintptr(rangeSize), uintptr(unsafe.Pointer(pRangeOwner)), uintptr(unsafe.Pointer(pRangeBase)))
	return HRESULT(ret)
}

func (this *IAccessibleWindowlessSite) ReleaseObjectIdRange(rangeBase int32, pRangeOwner *IAccessibleHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ReleaseObjectIdRange, uintptr(unsafe.Pointer(this)), uintptr(rangeBase), uintptr(unsafe.Pointer(pRangeOwner)))
	return HRESULT(ret)
}

func (this *IAccessibleWindowlessSite) QueryObjectIdRanges(pRangesOwner *IAccessibleHandler, psaRanges **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().QueryObjectIdRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRangesOwner)), uintptr(unsafe.Pointer(psaRanges)))
	return HRESULT(ret)
}

func (this *IAccessibleWindowlessSite) GetParentAccessible(ppParent **IAccessible) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetParentAccessible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppParent)))
	return HRESULT(ret)
}

// 7852b78d-1cfd-41c1-a615-9c0c85960b5f
var IID_IAccIdentity = syscall.GUID{0x7852b78d, 0x1cfd, 0x41c1, 
	[8]byte{0xa6, 0x15, 0x9c, 0x0c, 0x85, 0x96, 0x0b, 0x5f}}

type IAccIdentityInterface interface {
	IUnknownInterface
	GetIdentityString(dwIDChild uint32, ppIDString **uint8, pdwIDStringLen *uint32) HRESULT
}

type IAccIdentityVtbl struct {
	IUnknownVtbl
	GetIdentityString uintptr
}

type IAccIdentity struct {
	IUnknown
}

func (this *IAccIdentity) Vtbl() *IAccIdentityVtbl {
	return (*IAccIdentityVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccIdentity) GetIdentityString(dwIDChild uint32, ppIDString **uint8, pdwIDStringLen *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIdentityString, uintptr(unsafe.Pointer(this)), uintptr(dwIDChild), uintptr(unsafe.Pointer(ppIDString)), uintptr(unsafe.Pointer(pdwIDStringLen)))
	return HRESULT(ret)
}

// 76c0dbbb-15e0-4e7b-b61b-20eeea2001e0
var IID_IAccPropServer = syscall.GUID{0x76c0dbbb, 0x15e0, 0x4e7b, 
	[8]byte{0xb6, 0x1b, 0x20, 0xee, 0xea, 0x20, 0x01, 0xe0}}

type IAccPropServerInterface interface {
	IUnknownInterface
	GetPropValue(pIDString *uint8, dwIDStringLen uint32, idProp syscall.GUID, pvarValue *VARIANT, pfHasProp *BOOL) HRESULT
}

type IAccPropServerVtbl struct {
	IUnknownVtbl
	GetPropValue uintptr
}

type IAccPropServer struct {
	IUnknown
}

func (this *IAccPropServer) Vtbl() *IAccPropServerVtbl {
	return (*IAccPropServerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccPropServer) GetPropValue(pIDString *uint8, dwIDStringLen uint32, idProp syscall.GUID, pvarValue *VARIANT, pfHasProp *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIDString)), uintptr(dwIDStringLen), (uintptr)(unsafe.Pointer(&idProp)), uintptr(unsafe.Pointer(pvarValue)), uintptr(unsafe.Pointer(pfHasProp)))
	return HRESULT(ret)
}

// 6e26e776-04f0-495d-80e4-3330352e3169
var IID_IAccPropServices = syscall.GUID{0x6e26e776, 0x04f0, 0x495d, 
	[8]byte{0x80, 0xe4, 0x33, 0x30, 0x35, 0x2e, 0x31, 0x69}}

type IAccPropServicesInterface interface {
	IUnknownInterface
	SetPropValue(pIDString *uint8, dwIDStringLen uint32, idProp syscall.GUID, var_ VARIANT) HRESULT
	SetPropServer(pIDString *uint8, dwIDStringLen uint32, paProps *syscall.GUID, cProps int32, pServer *IAccPropServer, annoScope AnnoScope) HRESULT
	ClearProps(pIDString *uint8, dwIDStringLen uint32, paProps *syscall.GUID, cProps int32) HRESULT
	SetHwndProp(hwnd HWND, idObject uint32, idChild uint32, idProp syscall.GUID, var_ VARIANT) HRESULT
	SetHwndPropStr(hwnd HWND, idObject uint32, idChild uint32, idProp syscall.GUID, str PWSTR) HRESULT
	SetHwndPropServer(hwnd HWND, idObject uint32, idChild uint32, paProps *syscall.GUID, cProps int32, pServer *IAccPropServer, annoScope AnnoScope) HRESULT
	ClearHwndProps(hwnd HWND, idObject uint32, idChild uint32, paProps *syscall.GUID, cProps int32) HRESULT
	ComposeHwndIdentityString(hwnd HWND, idObject uint32, idChild uint32, ppIDString **uint8, pdwIDStringLen *uint32) HRESULT
	DecomposeHwndIdentityString(pIDString *uint8, dwIDStringLen uint32, phwnd *HWND, pidObject *uint32, pidChild *uint32) HRESULT
	SetHmenuProp(hmenu HMENU, idChild uint32, idProp syscall.GUID, var_ VARIANT) HRESULT
	SetHmenuPropStr(hmenu HMENU, idChild uint32, idProp syscall.GUID, str PWSTR) HRESULT
	SetHmenuPropServer(hmenu HMENU, idChild uint32, paProps *syscall.GUID, cProps int32, pServer *IAccPropServer, annoScope AnnoScope) HRESULT
	ClearHmenuProps(hmenu HMENU, idChild uint32, paProps *syscall.GUID, cProps int32) HRESULT
	ComposeHmenuIdentityString(hmenu HMENU, idChild uint32, ppIDString **uint8, pdwIDStringLen *uint32) HRESULT
	DecomposeHmenuIdentityString(pIDString *uint8, dwIDStringLen uint32, phmenu *HMENU, pidChild *uint32) HRESULT
}

type IAccPropServicesVtbl struct {
	IUnknownVtbl
	SetPropValue uintptr
	SetPropServer uintptr
	ClearProps uintptr
	SetHwndProp uintptr
	SetHwndPropStr uintptr
	SetHwndPropServer uintptr
	ClearHwndProps uintptr
	ComposeHwndIdentityString uintptr
	DecomposeHwndIdentityString uintptr
	SetHmenuProp uintptr
	SetHmenuPropStr uintptr
	SetHmenuPropServer uintptr
	ClearHmenuProps uintptr
	ComposeHmenuIdentityString uintptr
	DecomposeHmenuIdentityString uintptr
}

type IAccPropServices struct {
	IUnknown
}

func (this *IAccPropServices) Vtbl() *IAccPropServicesVtbl {
	return (*IAccPropServicesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccPropServices) SetPropValue(pIDString *uint8, dwIDStringLen uint32, idProp syscall.GUID, var_ VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPropValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIDString)), uintptr(dwIDStringLen), (uintptr)(unsafe.Pointer(&idProp)), (uintptr)(unsafe.Pointer(&var_)))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetPropServer(pIDString *uint8, dwIDStringLen uint32, paProps *syscall.GUID, cProps int32, pServer *IAccPropServer, annoScope AnnoScope) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetPropServer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIDString)), uintptr(dwIDStringLen), uintptr(unsafe.Pointer(paProps)), uintptr(cProps), uintptr(unsafe.Pointer(pServer)), uintptr(annoScope))
	return HRESULT(ret)
}

func (this *IAccPropServices) ClearProps(pIDString *uint8, dwIDStringLen uint32, paProps *syscall.GUID, cProps int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ClearProps, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIDString)), uintptr(dwIDStringLen), uintptr(unsafe.Pointer(paProps)), uintptr(cProps))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetHwndProp(hwnd HWND, idObject uint32, idChild uint32, idProp syscall.GUID, var_ VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHwndProp, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(idObject), uintptr(idChild), (uintptr)(unsafe.Pointer(&idProp)), (uintptr)(unsafe.Pointer(&var_)))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetHwndPropStr(hwnd HWND, idObject uint32, idChild uint32, idProp syscall.GUID, str PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHwndPropStr, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(idObject), uintptr(idChild), (uintptr)(unsafe.Pointer(&idProp)), uintptr(unsafe.Pointer(str)))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetHwndPropServer(hwnd HWND, idObject uint32, idChild uint32, paProps *syscall.GUID, cProps int32, pServer *IAccPropServer, annoScope AnnoScope) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHwndPropServer, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(idObject), uintptr(idChild), uintptr(unsafe.Pointer(paProps)), uintptr(cProps), uintptr(unsafe.Pointer(pServer)), uintptr(annoScope))
	return HRESULT(ret)
}

func (this *IAccPropServices) ClearHwndProps(hwnd HWND, idObject uint32, idChild uint32, paProps *syscall.GUID, cProps int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ClearHwndProps, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(idObject), uintptr(idChild), uintptr(unsafe.Pointer(paProps)), uintptr(cProps))
	return HRESULT(ret)
}

func (this *IAccPropServices) ComposeHwndIdentityString(hwnd HWND, idObject uint32, idChild uint32, ppIDString **uint8, pdwIDStringLen *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ComposeHwndIdentityString, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(idObject), uintptr(idChild), uintptr(unsafe.Pointer(ppIDString)), uintptr(unsafe.Pointer(pdwIDStringLen)))
	return HRESULT(ret)
}

func (this *IAccPropServices) DecomposeHwndIdentityString(pIDString *uint8, dwIDStringLen uint32, phwnd *HWND, pidObject *uint32, pidChild *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DecomposeHwndIdentityString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIDString)), uintptr(dwIDStringLen), uintptr(unsafe.Pointer(phwnd)), uintptr(unsafe.Pointer(pidObject)), uintptr(unsafe.Pointer(pidChild)))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetHmenuProp(hmenu HMENU, idChild uint32, idProp syscall.GUID, var_ VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHmenuProp, uintptr(unsafe.Pointer(this)), uintptr(hmenu), uintptr(idChild), (uintptr)(unsafe.Pointer(&idProp)), (uintptr)(unsafe.Pointer(&var_)))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetHmenuPropStr(hmenu HMENU, idChild uint32, idProp syscall.GUID, str PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHmenuPropStr, uintptr(unsafe.Pointer(this)), uintptr(hmenu), uintptr(idChild), (uintptr)(unsafe.Pointer(&idProp)), uintptr(unsafe.Pointer(str)))
	return HRESULT(ret)
}

func (this *IAccPropServices) SetHmenuPropServer(hmenu HMENU, idChild uint32, paProps *syscall.GUID, cProps int32, pServer *IAccPropServer, annoScope AnnoScope) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetHmenuPropServer, uintptr(unsafe.Pointer(this)), uintptr(hmenu), uintptr(idChild), uintptr(unsafe.Pointer(paProps)), uintptr(cProps), uintptr(unsafe.Pointer(pServer)), uintptr(annoScope))
	return HRESULT(ret)
}

func (this *IAccPropServices) ClearHmenuProps(hmenu HMENU, idChild uint32, paProps *syscall.GUID, cProps int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ClearHmenuProps, uintptr(unsafe.Pointer(this)), uintptr(hmenu), uintptr(idChild), uintptr(unsafe.Pointer(paProps)), uintptr(cProps))
	return HRESULT(ret)
}

func (this *IAccPropServices) ComposeHmenuIdentityString(hmenu HMENU, idChild uint32, ppIDString **uint8, pdwIDStringLen *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ComposeHmenuIdentityString, uintptr(unsafe.Pointer(this)), uintptr(hmenu), uintptr(idChild), uintptr(unsafe.Pointer(ppIDString)), uintptr(unsafe.Pointer(pdwIDStringLen)))
	return HRESULT(ret)
}

func (this *IAccPropServices) DecomposeHmenuIdentityString(pIDString *uint8, dwIDStringLen uint32, phmenu *HMENU, pidChild *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DecomposeHmenuIdentityString, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIDString)), uintptr(dwIDStringLen), uintptr(unsafe.Pointer(phmenu)), uintptr(unsafe.Pointer(pidChild)))
	return HRESULT(ret)
}

// d6dd68d1-86fd-4332-8666-9abedea2d24c
var IID_IRawElementProviderSimple = syscall.GUID{0xd6dd68d1, 0x86fd, 0x4332, 
	[8]byte{0x86, 0x66, 0x9a, 0xbe, 0xde, 0xa2, 0xd2, 0x4c}}

type IRawElementProviderSimpleInterface interface {
	IUnknownInterface
	Get_ProviderOptions(pRetVal *ProviderOptions) HRESULT
	GetPatternProvider(patternId int32, pRetVal **IUnknown) HRESULT
	GetPropertyValue(propertyId int32, pRetVal *VARIANT) HRESULT
	Get_HostRawElementProvider(pRetVal **IRawElementProviderSimple) HRESULT
}

type IRawElementProviderSimpleVtbl struct {
	IUnknownVtbl
	Get_ProviderOptions uintptr
	GetPatternProvider uintptr
	GetPropertyValue uintptr
	Get_HostRawElementProvider uintptr
}

type IRawElementProviderSimple struct {
	IUnknown
}

func (this *IRawElementProviderSimple) Vtbl() *IRawElementProviderSimpleVtbl {
	return (*IRawElementProviderSimpleVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderSimple) Get_ProviderOptions(pRetVal *ProviderOptions) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ProviderOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderSimple) GetPatternProvider(patternId int32, pRetVal **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPatternProvider, uintptr(unsafe.Pointer(this)), uintptr(patternId), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderSimple) GetPropertyValue(propertyId int32, pRetVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyValue, uintptr(unsafe.Pointer(this)), uintptr(propertyId), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderSimple) Get_HostRawElementProvider(pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_HostRawElementProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// f8b80ada-2c44-48d0-89be-5ff23c9cd875
var IID_IAccessibleEx = syscall.GUID{0xf8b80ada, 0x2c44, 0x48d0, 
	[8]byte{0x89, 0xbe, 0x5f, 0xf2, 0x3c, 0x9c, 0xd8, 0x75}}

type IAccessibleExInterface interface {
	IUnknownInterface
	GetObjectForChild(idChild int32, pRetVal **IAccessibleEx) HRESULT
	GetIAccessiblePair(ppAcc **IAccessible, pidChild *int32) HRESULT
	GetRuntimeId(pRetVal **SAFEARRAY) HRESULT
	ConvertReturnedElement(pIn *IRawElementProviderSimple, ppRetValOut **IAccessibleEx) HRESULT
}

type IAccessibleExVtbl struct {
	IUnknownVtbl
	GetObjectForChild uintptr
	GetIAccessiblePair uintptr
	GetRuntimeId uintptr
	ConvertReturnedElement uintptr
}

type IAccessibleEx struct {
	IUnknown
}

func (this *IAccessibleEx) Vtbl() *IAccessibleExVtbl {
	return (*IAccessibleExVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccessibleEx) GetObjectForChild(idChild int32, pRetVal **IAccessibleEx) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetObjectForChild, uintptr(unsafe.Pointer(this)), uintptr(idChild), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IAccessibleEx) GetIAccessiblePair(ppAcc **IAccessible, pidChild *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIAccessiblePair, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppAcc)), uintptr(unsafe.Pointer(pidChild)))
	return HRESULT(ret)
}

func (this *IAccessibleEx) GetRuntimeId(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRuntimeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IAccessibleEx) ConvertReturnedElement(pIn *IRawElementProviderSimple, ppRetValOut **IAccessibleEx) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ConvertReturnedElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pIn)), uintptr(unsafe.Pointer(ppRetValOut)))
	return HRESULT(ret)
}

// a0a839a9-8da1-4a82-806a-8e0d44e79f56
var IID_IRawElementProviderSimple2 = syscall.GUID{0xa0a839a9, 0x8da1, 0x4a82, 
	[8]byte{0x80, 0x6a, 0x8e, 0x0d, 0x44, 0xe7, 0x9f, 0x56}}

type IRawElementProviderSimple2Interface interface {
	IRawElementProviderSimpleInterface
	ShowContextMenu() HRESULT
}

type IRawElementProviderSimple2Vtbl struct {
	IRawElementProviderSimpleVtbl
	ShowContextMenu uintptr
}

type IRawElementProviderSimple2 struct {
	IRawElementProviderSimple
}

func (this *IRawElementProviderSimple2) Vtbl() *IRawElementProviderSimple2Vtbl {
	return (*IRawElementProviderSimple2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderSimple2) ShowContextMenu() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ShowContextMenu, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// fcf5d820-d7ec-4613-bdf6-42a84ce7daaf
var IID_IRawElementProviderSimple3 = syscall.GUID{0xfcf5d820, 0xd7ec, 0x4613, 
	[8]byte{0xbd, 0xf6, 0x42, 0xa8, 0x4c, 0xe7, 0xda, 0xaf}}

type IRawElementProviderSimple3Interface interface {
	IRawElementProviderSimple2Interface
	GetMetadataValue(targetId int32, metadataId int32, returnVal *VARIANT) HRESULT
}

type IRawElementProviderSimple3Vtbl struct {
	IRawElementProviderSimple2Vtbl
	GetMetadataValue uintptr
}

type IRawElementProviderSimple3 struct {
	IRawElementProviderSimple2
}

func (this *IRawElementProviderSimple3) Vtbl() *IRawElementProviderSimple3Vtbl {
	return (*IRawElementProviderSimple3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderSimple3) GetMetadataValue(targetId int32, metadataId int32, returnVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetMetadataValue, uintptr(unsafe.Pointer(this)), uintptr(targetId), uintptr(metadataId), uintptr(unsafe.Pointer(returnVal)))
	return HRESULT(ret)
}

// 620ce2a5-ab8f-40a9-86cb-de3c75599b58
var IID_IRawElementProviderFragmentRoot = syscall.GUID{0x620ce2a5, 0xab8f, 0x40a9, 
	[8]byte{0x86, 0xcb, 0xde, 0x3c, 0x75, 0x59, 0x9b, 0x58}}

type IRawElementProviderFragmentRootInterface interface {
	IUnknownInterface
	ElementProviderFromPoint(x float64, y float64, pRetVal **IRawElementProviderFragment) HRESULT
	GetFocus(pRetVal **IRawElementProviderFragment) HRESULT
}

type IRawElementProviderFragmentRootVtbl struct {
	IUnknownVtbl
	ElementProviderFromPoint uintptr
	GetFocus uintptr
}

type IRawElementProviderFragmentRoot struct {
	IUnknown
}

func (this *IRawElementProviderFragmentRoot) Vtbl() *IRawElementProviderFragmentRootVtbl {
	return (*IRawElementProviderFragmentRootVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderFragmentRoot) ElementProviderFromPoint(x float64, y float64, pRetVal **IRawElementProviderFragment) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementProviderFromPoint, uintptr(unsafe.Pointer(this)), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderFragmentRoot) GetFocus(pRetVal **IRawElementProviderFragment) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFocus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// f7063da8-8359-439c-9297-bbc5299a7d87
var IID_IRawElementProviderFragment = syscall.GUID{0xf7063da8, 0x8359, 0x439c, 
	[8]byte{0x92, 0x97, 0xbb, 0xc5, 0x29, 0x9a, 0x7d, 0x87}}

type IRawElementProviderFragmentInterface interface {
	IUnknownInterface
	Navigate(direction NavigateDirection, pRetVal **IRawElementProviderFragment) HRESULT
	GetRuntimeId(pRetVal **SAFEARRAY) HRESULT
	Get_BoundingRectangle(pRetVal *UiaRect) HRESULT
	GetEmbeddedFragmentRoots(pRetVal **SAFEARRAY) HRESULT
	SetFocus() HRESULT
	Get_FragmentRoot(pRetVal **IRawElementProviderFragmentRoot) HRESULT
}

type IRawElementProviderFragmentVtbl struct {
	IUnknownVtbl
	Navigate uintptr
	GetRuntimeId uintptr
	Get_BoundingRectangle uintptr
	GetEmbeddedFragmentRoots uintptr
	SetFocus uintptr
	Get_FragmentRoot uintptr
}

type IRawElementProviderFragment struct {
	IUnknown
}

func (this *IRawElementProviderFragment) Vtbl() *IRawElementProviderFragmentVtbl {
	return (*IRawElementProviderFragmentVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderFragment) Navigate(direction NavigateDirection, pRetVal **IRawElementProviderFragment) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Navigate, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderFragment) GetRuntimeId(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRuntimeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderFragment) Get_BoundingRectangle(pRetVal *UiaRect) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_BoundingRectangle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderFragment) GetEmbeddedFragmentRoots(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEmbeddedFragmentRoots, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRawElementProviderFragment) SetFocus() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFocus, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IRawElementProviderFragment) Get_FragmentRoot(pRetVal **IRawElementProviderFragmentRoot) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_FragmentRoot, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// a407b27b-0f6d-4427-9292-473c7bf93258
var IID_IRawElementProviderAdviseEvents = syscall.GUID{0xa407b27b, 0x0f6d, 0x4427, 
	[8]byte{0x92, 0x92, 0x47, 0x3c, 0x7b, 0xf9, 0x32, 0x58}}

type IRawElementProviderAdviseEventsInterface interface {
	IUnknownInterface
	AdviseEventAdded(eventId int32, propertyIDs *SAFEARRAY) HRESULT
	AdviseEventRemoved(eventId int32, propertyIDs *SAFEARRAY) HRESULT
}

type IRawElementProviderAdviseEventsVtbl struct {
	IUnknownVtbl
	AdviseEventAdded uintptr
	AdviseEventRemoved uintptr
}

type IRawElementProviderAdviseEvents struct {
	IUnknown
}

func (this *IRawElementProviderAdviseEvents) Vtbl() *IRawElementProviderAdviseEventsVtbl {
	return (*IRawElementProviderAdviseEventsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderAdviseEvents) AdviseEventAdded(eventId int32, propertyIDs *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AdviseEventAdded, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(unsafe.Pointer(propertyIDs)))
	return HRESULT(ret)
}

func (this *IRawElementProviderAdviseEvents) AdviseEventRemoved(eventId int32, propertyIDs *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AdviseEventRemoved, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(unsafe.Pointer(propertyIDs)))
	return HRESULT(ret)
}

// 1d5df27c-8947-4425-b8d9-79787bb460b8
var IID_IRawElementProviderHwndOverride = syscall.GUID{0x1d5df27c, 0x8947, 0x4425, 
	[8]byte{0xb8, 0xd9, 0x79, 0x78, 0x7b, 0xb4, 0x60, 0xb8}}

type IRawElementProviderHwndOverrideInterface interface {
	IUnknownInterface
	GetOverrideProviderForHwnd(hwnd HWND, pRetVal **IRawElementProviderSimple) HRESULT
}

type IRawElementProviderHwndOverrideVtbl struct {
	IUnknownVtbl
	GetOverrideProviderForHwnd uintptr
}

type IRawElementProviderHwndOverride struct {
	IUnknown
}

func (this *IRawElementProviderHwndOverride) Vtbl() *IRawElementProviderHwndOverrideVtbl {
	return (*IRawElementProviderHwndOverrideVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderHwndOverride) GetOverrideProviderForHwnd(hwnd HWND, pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetOverrideProviderForHwnd, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 4fd82b78-a43e-46ac-9803-0a6969c7c183
var IID_IProxyProviderWinEventSink = syscall.GUID{0x4fd82b78, 0xa43e, 0x46ac, 
	[8]byte{0x98, 0x03, 0x0a, 0x69, 0x69, 0xc7, 0xc1, 0x83}}

type IProxyProviderWinEventSinkInterface interface {
	IUnknownInterface
	AddAutomationPropertyChangedEvent(pProvider *IRawElementProviderSimple, id int32, newValue VARIANT) HRESULT
	AddAutomationEvent(pProvider *IRawElementProviderSimple, id int32) HRESULT
	AddStructureChangedEvent(pProvider *IRawElementProviderSimple, structureChangeType StructureChangeType, runtimeId *SAFEARRAY) HRESULT
}

type IProxyProviderWinEventSinkVtbl struct {
	IUnknownVtbl
	AddAutomationPropertyChangedEvent uintptr
	AddAutomationEvent uintptr
	AddStructureChangedEvent uintptr
}

type IProxyProviderWinEventSink struct {
	IUnknown
}

func (this *IProxyProviderWinEventSink) Vtbl() *IProxyProviderWinEventSinkVtbl {
	return (*IProxyProviderWinEventSinkVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProxyProviderWinEventSink) AddAutomationPropertyChangedEvent(pProvider *IRawElementProviderSimple, id int32, newValue VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddAutomationPropertyChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pProvider)), uintptr(id), (uintptr)(unsafe.Pointer(&newValue)))
	return HRESULT(ret)
}

func (this *IProxyProviderWinEventSink) AddAutomationEvent(pProvider *IRawElementProviderSimple, id int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddAutomationEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pProvider)), uintptr(id))
	return HRESULT(ret)
}

func (this *IProxyProviderWinEventSink) AddStructureChangedEvent(pProvider *IRawElementProviderSimple, structureChangeType StructureChangeType, runtimeId *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddStructureChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pProvider)), uintptr(structureChangeType), uintptr(unsafe.Pointer(runtimeId)))
	return HRESULT(ret)
}

// 89592ad4-f4e0-43d5-a3b6-bad7e111b435
var IID_IProxyProviderWinEventHandler = syscall.GUID{0x89592ad4, 0xf4e0, 0x43d5, 
	[8]byte{0xa3, 0xb6, 0xba, 0xd7, 0xe1, 0x11, 0xb4, 0x35}}

type IProxyProviderWinEventHandlerInterface interface {
	IUnknownInterface
	RespondToWinEvent(idWinEvent uint32, hwnd HWND, idObject int32, idChild int32, pSink *IProxyProviderWinEventSink) HRESULT
}

type IProxyProviderWinEventHandlerVtbl struct {
	IUnknownVtbl
	RespondToWinEvent uintptr
}

type IProxyProviderWinEventHandler struct {
	IUnknown
}

func (this *IProxyProviderWinEventHandler) Vtbl() *IProxyProviderWinEventHandlerVtbl {
	return (*IProxyProviderWinEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IProxyProviderWinEventHandler) RespondToWinEvent(idWinEvent uint32, hwnd HWND, idObject int32, idChild int32, pSink *IProxyProviderWinEventSink) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RespondToWinEvent, uintptr(unsafe.Pointer(this)), uintptr(idWinEvent), uintptr(hwnd), uintptr(idObject), uintptr(idChild), uintptr(unsafe.Pointer(pSink)))
	return HRESULT(ret)
}

// 0a2a93cc-bfad-42ac-9b2e-0991fb0d3ea0
var IID_IRawElementProviderWindowlessSite = syscall.GUID{0x0a2a93cc, 0xbfad, 0x42ac, 
	[8]byte{0x9b, 0x2e, 0x09, 0x91, 0xfb, 0x0d, 0x3e, 0xa0}}

type IRawElementProviderWindowlessSiteInterface interface {
	IUnknownInterface
	GetAdjacentFragment(direction NavigateDirection, ppParent **IRawElementProviderFragment) HRESULT
	GetRuntimeIdPrefix(pRetVal **SAFEARRAY) HRESULT
}

type IRawElementProviderWindowlessSiteVtbl struct {
	IUnknownVtbl
	GetAdjacentFragment uintptr
	GetRuntimeIdPrefix uintptr
}

type IRawElementProviderWindowlessSite struct {
	IUnknown
}

func (this *IRawElementProviderWindowlessSite) Vtbl() *IRawElementProviderWindowlessSiteVtbl {
	return (*IRawElementProviderWindowlessSiteVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderWindowlessSite) GetAdjacentFragment(direction NavigateDirection, ppParent **IRawElementProviderFragment) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAdjacentFragment, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(unsafe.Pointer(ppParent)))
	return HRESULT(ret)
}

func (this *IRawElementProviderWindowlessSite) GetRuntimeIdPrefix(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRuntimeIdPrefix, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 33ac331b-943e-4020-b295-db37784974a3
var IID_IAccessibleHostingElementProviders = syscall.GUID{0x33ac331b, 0x943e, 0x4020, 
	[8]byte{0xb2, 0x95, 0xdb, 0x37, 0x78, 0x49, 0x74, 0xa3}}

type IAccessibleHostingElementProvidersInterface interface {
	IUnknownInterface
	GetEmbeddedFragmentRoots(pRetVal **SAFEARRAY) HRESULT
	GetObjectIdForProvider(pProvider *IRawElementProviderSimple, pidObject *int32) HRESULT
}

type IAccessibleHostingElementProvidersVtbl struct {
	IUnknownVtbl
	GetEmbeddedFragmentRoots uintptr
	GetObjectIdForProvider uintptr
}

type IAccessibleHostingElementProviders struct {
	IUnknown
}

func (this *IAccessibleHostingElementProviders) Vtbl() *IAccessibleHostingElementProvidersVtbl {
	return (*IAccessibleHostingElementProvidersVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAccessibleHostingElementProviders) GetEmbeddedFragmentRoots(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEmbeddedFragmentRoots, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IAccessibleHostingElementProviders) GetObjectIdForProvider(pProvider *IRawElementProviderSimple, pidObject *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetObjectIdForProvider, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pProvider)), uintptr(unsafe.Pointer(pidObject)))
	return HRESULT(ret)
}

// 24be0b07-d37d-487a-98cf-a13ed465e9b3
var IID_IRawElementProviderHostingAccessibles = syscall.GUID{0x24be0b07, 0xd37d, 0x487a, 
	[8]byte{0x98, 0xcf, 0xa1, 0x3e, 0xd4, 0x65, 0xe9, 0xb3}}

type IRawElementProviderHostingAccessiblesInterface interface {
	IUnknownInterface
	GetEmbeddedAccessibles(pRetVal **SAFEARRAY) HRESULT
}

type IRawElementProviderHostingAccessiblesVtbl struct {
	IUnknownVtbl
	GetEmbeddedAccessibles uintptr
}

type IRawElementProviderHostingAccessibles struct {
	IUnknown
}

func (this *IRawElementProviderHostingAccessibles) Vtbl() *IRawElementProviderHostingAccessiblesVtbl {
	return (*IRawElementProviderHostingAccessiblesVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRawElementProviderHostingAccessibles) GetEmbeddedAccessibles(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEmbeddedAccessibles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 159bc72c-4ad3-485e-9637-d7052edf0146
var IID_IDockProvider = syscall.GUID{0x159bc72c, 0x4ad3, 0x485e, 
	[8]byte{0x96, 0x37, 0xd7, 0x05, 0x2e, 0xdf, 0x01, 0x46}}

type IDockProviderInterface interface {
	IUnknownInterface
	SetDockPosition(dockPosition DockPosition) HRESULT
	Get_DockPosition(pRetVal *DockPosition) HRESULT
}

type IDockProviderVtbl struct {
	IUnknownVtbl
	SetDockPosition uintptr
	Get_DockPosition uintptr
}

type IDockProvider struct {
	IUnknown
}

func (this *IDockProvider) Vtbl() *IDockProviderVtbl {
	return (*IDockProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDockProvider) SetDockPosition(dockPosition DockPosition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDockPosition, uintptr(unsafe.Pointer(this)), uintptr(dockPosition))
	return HRESULT(ret)
}

func (this *IDockProvider) Get_DockPosition(pRetVal *DockPosition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DockPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// d847d3a5-cab0-4a98-8c32-ecb45c59ad24
var IID_IExpandCollapseProvider = syscall.GUID{0xd847d3a5, 0xcab0, 0x4a98, 
	[8]byte{0x8c, 0x32, 0xec, 0xb4, 0x5c, 0x59, 0xad, 0x24}}

type IExpandCollapseProviderInterface interface {
	IUnknownInterface
	Expand() HRESULT
	Collapse() HRESULT
	Get_ExpandCollapseState(pRetVal *ExpandCollapseState) HRESULT
}

type IExpandCollapseProviderVtbl struct {
	IUnknownVtbl
	Expand uintptr
	Collapse uintptr
	Get_ExpandCollapseState uintptr
}

type IExpandCollapseProvider struct {
	IUnknown
}

func (this *IExpandCollapseProvider) Vtbl() *IExpandCollapseProviderVtbl {
	return (*IExpandCollapseProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IExpandCollapseProvider) Expand() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Expand, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IExpandCollapseProvider) Collapse() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Collapse, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IExpandCollapseProvider) Get_ExpandCollapseState(pRetVal *ExpandCollapseState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ExpandCollapseState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// b17d6187-0907-464b-a168-0ef17a1572b1
var IID_IGridProvider = syscall.GUID{0xb17d6187, 0x0907, 0x464b, 
	[8]byte{0xa1, 0x68, 0x0e, 0xf1, 0x7a, 0x15, 0x72, 0xb1}}

type IGridProviderInterface interface {
	IUnknownInterface
	GetItem(row int32, column int32, pRetVal **IRawElementProviderSimple) HRESULT
	Get_RowCount(pRetVal *int32) HRESULT
	Get_ColumnCount(pRetVal *int32) HRESULT
}

type IGridProviderVtbl struct {
	IUnknownVtbl
	GetItem uintptr
	Get_RowCount uintptr
	Get_ColumnCount uintptr
}

type IGridProvider struct {
	IUnknown
}

func (this *IGridProvider) Vtbl() *IGridProviderVtbl {
	return (*IGridProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGridProvider) GetItem(row int32, column int32, pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetItem, uintptr(unsafe.Pointer(this)), uintptr(row), uintptr(column), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IGridProvider) Get_RowCount(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_RowCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IGridProvider) Get_ColumnCount(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ColumnCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// d02541f1-fb81-4d64-ae32-f520f8a6dbd1
var IID_IGridItemProvider = syscall.GUID{0xd02541f1, 0xfb81, 0x4d64, 
	[8]byte{0xae, 0x32, 0xf5, 0x20, 0xf8, 0xa6, 0xdb, 0xd1}}

type IGridItemProviderInterface interface {
	IUnknownInterface
	Get_Row(pRetVal *int32) HRESULT
	Get_Column(pRetVal *int32) HRESULT
	Get_RowSpan(pRetVal *int32) HRESULT
	Get_ColumnSpan(pRetVal *int32) HRESULT
	Get_ContainingGrid(pRetVal **IRawElementProviderSimple) HRESULT
}

type IGridItemProviderVtbl struct {
	IUnknownVtbl
	Get_Row uintptr
	Get_Column uintptr
	Get_RowSpan uintptr
	Get_ColumnSpan uintptr
	Get_ContainingGrid uintptr
}

type IGridItemProvider struct {
	IUnknown
}

func (this *IGridItemProvider) Vtbl() *IGridItemProviderVtbl {
	return (*IGridItemProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IGridItemProvider) Get_Row(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Row, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IGridItemProvider) Get_Column(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Column, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IGridItemProvider) Get_RowSpan(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_RowSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IGridItemProvider) Get_ColumnSpan(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ColumnSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IGridItemProvider) Get_ContainingGrid(pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ContainingGrid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 54fcb24b-e18e-47a2-b4d3-eccbe77599a2
var IID_IInvokeProvider = syscall.GUID{0x54fcb24b, 0xe18e, 0x47a2, 
	[8]byte{0xb4, 0xd3, 0xec, 0xcb, 0xe7, 0x75, 0x99, 0xa2}}

type IInvokeProviderInterface interface {
	IUnknownInterface
	Invoke() HRESULT
}

type IInvokeProviderVtbl struct {
	IUnknownVtbl
	Invoke uintptr
}

type IInvokeProvider struct {
	IUnknown
}

func (this *IInvokeProvider) Vtbl() *IInvokeProviderVtbl {
	return (*IInvokeProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInvokeProvider) Invoke() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Invoke, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 6278cab1-b556-4a1a-b4e0-418acc523201
var IID_IMultipleViewProvider = syscall.GUID{0x6278cab1, 0xb556, 0x4a1a, 
	[8]byte{0xb4, 0xe0, 0x41, 0x8a, 0xcc, 0x52, 0x32, 0x01}}

type IMultipleViewProviderInterface interface {
	IUnknownInterface
	GetViewName(viewId int32, pRetVal *BSTR) HRESULT
	SetCurrentView(viewId int32) HRESULT
	Get_CurrentView(pRetVal *int32) HRESULT
	GetSupportedViews(pRetVal **SAFEARRAY) HRESULT
}

type IMultipleViewProviderVtbl struct {
	IUnknownVtbl
	GetViewName uintptr
	SetCurrentView uintptr
	Get_CurrentView uintptr
	GetSupportedViews uintptr
}

type IMultipleViewProvider struct {
	IUnknown
}

func (this *IMultipleViewProvider) Vtbl() *IMultipleViewProviderVtbl {
	return (*IMultipleViewProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IMultipleViewProvider) GetViewName(viewId int32, pRetVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetViewName, uintptr(unsafe.Pointer(this)), uintptr(viewId), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IMultipleViewProvider) SetCurrentView(viewId int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCurrentView, uintptr(unsafe.Pointer(this)), uintptr(viewId))
	return HRESULT(ret)
}

func (this *IMultipleViewProvider) Get_CurrentView(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IMultipleViewProvider) GetSupportedViews(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSupportedViews, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 36dc7aef-33e6-4691-afe1-2be7274b3d33
var IID_IRangeValueProvider = syscall.GUID{0x36dc7aef, 0x33e6, 0x4691, 
	[8]byte{0xaf, 0xe1, 0x2b, 0xe7, 0x27, 0x4b, 0x3d, 0x33}}

type IRangeValueProviderInterface interface {
	IUnknownInterface
	SetValue(val float64) HRESULT
	Get_Value(pRetVal *float64) HRESULT
	Get_IsReadOnly(pRetVal *BOOL) HRESULT
	Get_Maximum(pRetVal *float64) HRESULT
	Get_Minimum(pRetVal *float64) HRESULT
	Get_LargeChange(pRetVal *float64) HRESULT
	Get_SmallChange(pRetVal *float64) HRESULT
}

type IRangeValueProviderVtbl struct {
	IUnknownVtbl
	SetValue uintptr
	Get_Value uintptr
	Get_IsReadOnly uintptr
	Get_Maximum uintptr
	Get_Minimum uintptr
	Get_LargeChange uintptr
	Get_SmallChange uintptr
}

type IRangeValueProvider struct {
	IUnknown
}

func (this *IRangeValueProvider) Vtbl() *IRangeValueProviderVtbl {
	return (*IRangeValueProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRangeValueProvider) SetValue(val float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(val))
	return HRESULT(ret)
}

func (this *IRangeValueProvider) Get_Value(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRangeValueProvider) Get_IsReadOnly(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRangeValueProvider) Get_Maximum(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Maximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRangeValueProvider) Get_Minimum(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Minimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRangeValueProvider) Get_LargeChange(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_LargeChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IRangeValueProvider) Get_SmallChange(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_SmallChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 2360c714-4bf1-4b26-ba65-9b21316127eb
var IID_IScrollItemProvider = syscall.GUID{0x2360c714, 0x4bf1, 0x4b26, 
	[8]byte{0xba, 0x65, 0x9b, 0x21, 0x31, 0x61, 0x27, 0xeb}}

type IScrollItemProviderInterface interface {
	IUnknownInterface
	ScrollIntoView() HRESULT
}

type IScrollItemProviderVtbl struct {
	IUnknownVtbl
	ScrollIntoView uintptr
}

type IScrollItemProvider struct {
	IUnknown
}

func (this *IScrollItemProvider) Vtbl() *IScrollItemProviderVtbl {
	return (*IScrollItemProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScrollItemProvider) ScrollIntoView() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ScrollIntoView, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// fb8b03af-3bdf-48d4-bd36-1a65793be168
var IID_ISelectionProvider = syscall.GUID{0xfb8b03af, 0x3bdf, 0x48d4, 
	[8]byte{0xbd, 0x36, 0x1a, 0x65, 0x79, 0x3b, 0xe1, 0x68}}

type ISelectionProviderInterface interface {
	IUnknownInterface
	GetSelection(pRetVal **SAFEARRAY) HRESULT
	Get_CanSelectMultiple(pRetVal *BOOL) HRESULT
	Get_IsSelectionRequired(pRetVal *BOOL) HRESULT
}

type ISelectionProviderVtbl struct {
	IUnknownVtbl
	GetSelection uintptr
	Get_CanSelectMultiple uintptr
	Get_IsSelectionRequired uintptr
}

type ISelectionProvider struct {
	IUnknown
}

func (this *ISelectionProvider) Vtbl() *ISelectionProviderVtbl {
	return (*ISelectionProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISelectionProvider) GetSelection(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ISelectionProvider) Get_CanSelectMultiple(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanSelectMultiple, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ISelectionProvider) Get_IsSelectionRequired(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelectionRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 14f68475-ee1c-44f6-a869-d239381f0fe7
var IID_ISelectionProvider2 = syscall.GUID{0x14f68475, 0xee1c, 0x44f6, 
	[8]byte{0xa8, 0x69, 0xd2, 0x39, 0x38, 0x1f, 0x0f, 0xe7}}

type ISelectionProvider2Interface interface {
	ISelectionProviderInterface
	Get_FirstSelectedItem(retVal **IRawElementProviderSimple) HRESULT
	Get_LastSelectedItem(retVal **IRawElementProviderSimple) HRESULT
	Get_CurrentSelectedItem(retVal **IRawElementProviderSimple) HRESULT
	Get_ItemCount(retVal *int32) HRESULT
}

type ISelectionProvider2Vtbl struct {
	ISelectionProviderVtbl
	Get_FirstSelectedItem uintptr
	Get_LastSelectedItem uintptr
	Get_CurrentSelectedItem uintptr
	Get_ItemCount uintptr
}

type ISelectionProvider2 struct {
	ISelectionProvider
}

func (this *ISelectionProvider2) Vtbl() *ISelectionProvider2Vtbl {
	return (*ISelectionProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISelectionProvider2) Get_FirstSelectedItem(retVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_FirstSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *ISelectionProvider2) Get_LastSelectedItem(retVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_LastSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *ISelectionProvider2) Get_CurrentSelectedItem(retVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *ISelectionProvider2) Get_ItemCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ItemCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// b38b8077-1fc3-42a5-8cae-d40c2215055a
var IID_IScrollProvider = syscall.GUID{0xb38b8077, 0x1fc3, 0x42a5, 
	[8]byte{0x8c, 0xae, 0xd4, 0x0c, 0x22, 0x15, 0x05, 0x5a}}

type IScrollProviderInterface interface {
	IUnknownInterface
	Scroll(horizontalAmount ScrollAmount, verticalAmount ScrollAmount) HRESULT
	SetScrollPercent(horizontalPercent float64, verticalPercent float64) HRESULT
	Get_HorizontalScrollPercent(pRetVal *float64) HRESULT
	Get_VerticalScrollPercent(pRetVal *float64) HRESULT
	Get_HorizontalViewSize(pRetVal *float64) HRESULT
	Get_VerticalViewSize(pRetVal *float64) HRESULT
	Get_HorizontallyScrollable(pRetVal *BOOL) HRESULT
	Get_VerticallyScrollable(pRetVal *BOOL) HRESULT
}

type IScrollProviderVtbl struct {
	IUnknownVtbl
	Scroll uintptr
	SetScrollPercent uintptr
	Get_HorizontalScrollPercent uintptr
	Get_VerticalScrollPercent uintptr
	Get_HorizontalViewSize uintptr
	Get_VerticalViewSize uintptr
	Get_HorizontallyScrollable uintptr
	Get_VerticallyScrollable uintptr
}

type IScrollProvider struct {
	IUnknown
}

func (this *IScrollProvider) Vtbl() *IScrollProviderVtbl {
	return (*IScrollProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IScrollProvider) Scroll(horizontalAmount ScrollAmount, verticalAmount ScrollAmount) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Scroll, uintptr(unsafe.Pointer(this)), uintptr(horizontalAmount), uintptr(verticalAmount))
	return HRESULT(ret)
}

func (this *IScrollProvider) SetScrollPercent(horizontalPercent float64, verticalPercent float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(horizontalPercent), uintptr(verticalPercent))
	return HRESULT(ret)
}

func (this *IScrollProvider) Get_HorizontalScrollPercent(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IScrollProvider) Get_VerticalScrollPercent(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IScrollProvider) Get_HorizontalViewSize(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontalViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IScrollProvider) Get_VerticalViewSize(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticalViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IScrollProvider) Get_HorizontallyScrollable(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_HorizontallyScrollable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IScrollProvider) Get_VerticallyScrollable(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_VerticallyScrollable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 2acad808-b2d4-452d-a407-91ff1ad167b2
var IID_ISelectionItemProvider = syscall.GUID{0x2acad808, 0xb2d4, 0x452d, 
	[8]byte{0xa4, 0x07, 0x91, 0xff, 0x1a, 0xd1, 0x67, 0xb2}}

type ISelectionItemProviderInterface interface {
	IUnknownInterface
	Select() HRESULT
	AddToSelection() HRESULT
	RemoveFromSelection() HRESULT
	Get_IsSelected(pRetVal *BOOL) HRESULT
	Get_SelectionContainer(pRetVal **IRawElementProviderSimple) HRESULT
}

type ISelectionItemProviderVtbl struct {
	IUnknownVtbl
	Select uintptr
	AddToSelection uintptr
	RemoveFromSelection uintptr
	Get_IsSelected uintptr
	Get_SelectionContainer uintptr
}

type ISelectionItemProvider struct {
	IUnknown
}

func (this *ISelectionItemProvider) Vtbl() *ISelectionItemProviderVtbl {
	return (*ISelectionItemProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISelectionItemProvider) Select() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Select, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ISelectionItemProvider) AddToSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddToSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ISelectionItemProvider) RemoveFromSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveFromSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ISelectionItemProvider) Get_IsSelected(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ISelectionItemProvider) Get_SelectionContainer(pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_SelectionContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 29db1a06-02ce-4cf7-9b42-565d4fab20ee
var IID_ISynchronizedInputProvider = syscall.GUID{0x29db1a06, 0x02ce, 0x4cf7, 
	[8]byte{0x9b, 0x42, 0x56, 0x5d, 0x4f, 0xab, 0x20, 0xee}}

type ISynchronizedInputProviderInterface interface {
	IUnknownInterface
	StartListening(inputType SynchronizedInputType) HRESULT
	Cancel() HRESULT
}

type ISynchronizedInputProviderVtbl struct {
	IUnknownVtbl
	StartListening uintptr
	Cancel uintptr
}

type ISynchronizedInputProvider struct {
	IUnknown
}

func (this *ISynchronizedInputProvider) Vtbl() *ISynchronizedInputProviderVtbl {
	return (*ISynchronizedInputProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISynchronizedInputProvider) StartListening(inputType SynchronizedInputType) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().StartListening, uintptr(unsafe.Pointer(this)), uintptr(inputType))
	return HRESULT(ret)
}

func (this *ISynchronizedInputProvider) Cancel() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Cancel, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 9c860395-97b3-490a-b52a-858cc22af166
var IID_ITableProvider = syscall.GUID{0x9c860395, 0x97b3, 0x490a, 
	[8]byte{0xb5, 0x2a, 0x85, 0x8c, 0xc2, 0x2a, 0xf1, 0x66}}

type ITableProviderInterface interface {
	IUnknownInterface
	GetRowHeaders(pRetVal **SAFEARRAY) HRESULT
	GetColumnHeaders(pRetVal **SAFEARRAY) HRESULT
	Get_RowOrColumnMajor(pRetVal *RowOrColumnMajor) HRESULT
}

type ITableProviderVtbl struct {
	IUnknownVtbl
	GetRowHeaders uintptr
	GetColumnHeaders uintptr
	Get_RowOrColumnMajor uintptr
}

type ITableProvider struct {
	IUnknown
}

func (this *ITableProvider) Vtbl() *ITableProviderVtbl {
	return (*ITableProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITableProvider) GetRowHeaders(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRowHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITableProvider) GetColumnHeaders(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColumnHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITableProvider) Get_RowOrColumnMajor(pRetVal *RowOrColumnMajor) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_RowOrColumnMajor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// b9734fa6-771f-4d78-9c90-2517999349cd
var IID_ITableItemProvider = syscall.GUID{0xb9734fa6, 0x771f, 0x4d78, 
	[8]byte{0x9c, 0x90, 0x25, 0x17, 0x99, 0x93, 0x49, 0xcd}}

type ITableItemProviderInterface interface {
	IUnknownInterface
	GetRowHeaderItems(pRetVal **SAFEARRAY) HRESULT
	GetColumnHeaderItems(pRetVal **SAFEARRAY) HRESULT
}

type ITableItemProviderVtbl struct {
	IUnknownVtbl
	GetRowHeaderItems uintptr
	GetColumnHeaderItems uintptr
}

type ITableItemProvider struct {
	IUnknown
}

func (this *ITableItemProvider) Vtbl() *ITableItemProviderVtbl {
	return (*ITableItemProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITableItemProvider) GetRowHeaderItems(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRowHeaderItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITableItemProvider) GetColumnHeaderItems(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetColumnHeaderItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 56d00bd0-c4f4-433c-a836-1a52a57e0892
var IID_IToggleProvider = syscall.GUID{0x56d00bd0, 0xc4f4, 0x433c, 
	[8]byte{0xa8, 0x36, 0x1a, 0x52, 0xa5, 0x7e, 0x08, 0x92}}

type IToggleProviderInterface interface {
	IUnknownInterface
	Toggle() HRESULT
	Get_ToggleState(pRetVal *ToggleState) HRESULT
}

type IToggleProviderVtbl struct {
	IUnknownVtbl
	Toggle uintptr
	Get_ToggleState uintptr
}

type IToggleProvider struct {
	IUnknown
}

func (this *IToggleProvider) Vtbl() *IToggleProviderVtbl {
	return (*IToggleProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IToggleProvider) Toggle() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Toggle, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IToggleProvider) Get_ToggleState(pRetVal *ToggleState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ToggleState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 6829ddc4-4f91-4ffa-b86f-bd3e2987cb4c
var IID_ITransformProvider = syscall.GUID{0x6829ddc4, 0x4f91, 0x4ffa, 
	[8]byte{0xb8, 0x6f, 0xbd, 0x3e, 0x29, 0x87, 0xcb, 0x4c}}

type ITransformProviderInterface interface {
	IUnknownInterface
	Move(x float64, y float64) HRESULT
	Resize(width float64, height float64) HRESULT
	Rotate(degrees float64) HRESULT
	Get_CanMove(pRetVal *BOOL) HRESULT
	Get_CanResize(pRetVal *BOOL) HRESULT
	Get_CanRotate(pRetVal *BOOL) HRESULT
}

type ITransformProviderVtbl struct {
	IUnknownVtbl
	Move uintptr
	Resize uintptr
	Rotate uintptr
	Get_CanMove uintptr
	Get_CanResize uintptr
	Get_CanRotate uintptr
}

type ITransformProvider struct {
	IUnknown
}

func (this *ITransformProvider) Vtbl() *ITransformProviderVtbl {
	return (*ITransformProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITransformProvider) Move(x float64, y float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Move, uintptr(unsafe.Pointer(this)), uintptr(x), uintptr(y))
	return HRESULT(ret)
}

func (this *ITransformProvider) Resize(width float64, height float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Resize, uintptr(unsafe.Pointer(this)), uintptr(width), uintptr(height))
	return HRESULT(ret)
}

func (this *ITransformProvider) Rotate(degrees float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Rotate, uintptr(unsafe.Pointer(this)), uintptr(degrees))
	return HRESULT(ret)
}

func (this *ITransformProvider) Get_CanMove(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITransformProvider) Get_CanResize(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanResize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITransformProvider) Get_CanRotate(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanRotate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// c7935180-6fb3-4201-b174-7df73adbf64a
var IID_IValueProvider = syscall.GUID{0xc7935180, 0x6fb3, 0x4201, 
	[8]byte{0xb1, 0x74, 0x7d, 0xf7, 0x3a, 0xdb, 0xf6, 0x4a}}

type IValueProviderInterface interface {
	IUnknownInterface
	SetValue(val PWSTR) HRESULT
	Get_Value(pRetVal *BSTR) HRESULT
	Get_IsReadOnly(pRetVal *BOOL) HRESULT
}

type IValueProviderVtbl struct {
	IUnknownVtbl
	SetValue uintptr
	Get_Value uintptr
	Get_IsReadOnly uintptr
}

type IValueProvider struct {
	IUnknown
}

func (this *IValueProvider) Vtbl() *IValueProviderVtbl {
	return (*IValueProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IValueProvider) SetValue(val PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(val)))
	return HRESULT(ret)
}

func (this *IValueProvider) Get_Value(pRetVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IValueProvider) Get_IsReadOnly(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 987df77b-db06-4d77-8f8a-86a9c3bb90b9
var IID_IWindowProvider = syscall.GUID{0x987df77b, 0xdb06, 0x4d77, 
	[8]byte{0x8f, 0x8a, 0x86, 0xa9, 0xc3, 0xbb, 0x90, 0xb9}}

type IWindowProviderInterface interface {
	IUnknownInterface
	SetVisualState(state WindowVisualState) HRESULT
	Close() HRESULT
	WaitForInputIdle(milliseconds int32, pRetVal *BOOL) HRESULT
	Get_CanMaximize(pRetVal *BOOL) HRESULT
	Get_CanMinimize(pRetVal *BOOL) HRESULT
	Get_IsModal(pRetVal *BOOL) HRESULT
	Get_WindowVisualState(pRetVal *WindowVisualState) HRESULT
	Get_WindowInteractionState(pRetVal *WindowInteractionState) HRESULT
	Get_IsTopmost(pRetVal *BOOL) HRESULT
}

type IWindowProviderVtbl struct {
	IUnknownVtbl
	SetVisualState uintptr
	Close uintptr
	WaitForInputIdle uintptr
	Get_CanMaximize uintptr
	Get_CanMinimize uintptr
	Get_IsModal uintptr
	Get_WindowVisualState uintptr
	Get_WindowInteractionState uintptr
	Get_IsTopmost uintptr
}

type IWindowProvider struct {
	IUnknown
}

func (this *IWindowProvider) Vtbl() *IWindowProviderVtbl {
	return (*IWindowProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IWindowProvider) SetVisualState(state WindowVisualState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetVisualState, uintptr(unsafe.Pointer(this)), uintptr(state))
	return HRESULT(ret)
}

func (this *IWindowProvider) Close() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IWindowProvider) WaitForInputIdle(milliseconds int32, pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().WaitForInputIdle, uintptr(unsafe.Pointer(this)), uintptr(milliseconds), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IWindowProvider) Get_CanMaximize(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMaximize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IWindowProvider) Get_CanMinimize(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanMinimize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IWindowProvider) Get_IsModal(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsModal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IWindowProvider) Get_WindowVisualState(pRetVal *WindowVisualState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_WindowVisualState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IWindowProvider) Get_WindowInteractionState(pRetVal *WindowInteractionState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_WindowInteractionState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IWindowProvider) Get_IsTopmost(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsTopmost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// e44c3566-915d-4070-99c6-047bff5a08f5
var IID_ILegacyIAccessibleProvider = syscall.GUID{0xe44c3566, 0x915d, 0x4070, 
	[8]byte{0x99, 0xc6, 0x04, 0x7b, 0xff, 0x5a, 0x08, 0xf5}}

type ILegacyIAccessibleProviderInterface interface {
	IUnknownInterface
	Select(flagsSelect int32) HRESULT
	DoDefaultAction() HRESULT
	SetValue(szValue PWSTR) HRESULT
	GetIAccessible(ppAccessible **IAccessible) HRESULT
	Get_ChildId(pRetVal *int32) HRESULT
	Get_Name(pszName *BSTR) HRESULT
	Get_Value(pszValue *BSTR) HRESULT
	Get_Description(pszDescription *BSTR) HRESULT
	Get_Role(pdwRole *uint32) HRESULT
	Get_State(pdwState *uint32) HRESULT
	Get_Help(pszHelp *BSTR) HRESULT
	Get_KeyboardShortcut(pszKeyboardShortcut *BSTR) HRESULT
	GetSelection(pvarSelectedChildren **SAFEARRAY) HRESULT
	Get_DefaultAction(pszDefaultAction *BSTR) HRESULT
}

type ILegacyIAccessibleProviderVtbl struct {
	IUnknownVtbl
	Select uintptr
	DoDefaultAction uintptr
	SetValue uintptr
	GetIAccessible uintptr
	Get_ChildId uintptr
	Get_Name uintptr
	Get_Value uintptr
	Get_Description uintptr
	Get_Role uintptr
	Get_State uintptr
	Get_Help uintptr
	Get_KeyboardShortcut uintptr
	GetSelection uintptr
	Get_DefaultAction uintptr
}

type ILegacyIAccessibleProvider struct {
	IUnknown
}

func (this *ILegacyIAccessibleProvider) Vtbl() *ILegacyIAccessibleProviderVtbl {
	return (*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ILegacyIAccessibleProvider) Select(flagsSelect int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Select, uintptr(unsafe.Pointer(this)), uintptr(flagsSelect))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) DoDefaultAction() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DoDefaultAction, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) SetValue(szValue PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(szValue)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) GetIAccessible(ppAccessible **IAccessible) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIAccessible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppAccessible)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_ChildId(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ChildId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_Name(pszName *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszName)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_Value(pszValue *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Value, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszValue)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_Description(pszDescription *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Description, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszDescription)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_Role(pdwRole *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Role, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwRole)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_State(pdwState *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_State, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwState)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_Help(pszHelp *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Help, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszHelp)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_KeyboardShortcut(pszKeyboardShortcut *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_KeyboardShortcut, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszKeyboardShortcut)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) GetSelection(pvarSelectedChildren **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pvarSelectedChildren)))
	return HRESULT(ret)
}

func (this *ILegacyIAccessibleProvider) Get_DefaultAction(pszDefaultAction *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DefaultAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszDefaultAction)))
	return HRESULT(ret)
}

// e747770b-39ce-4382-ab30-d8fb3f336f24
var IID_IItemContainerProvider = syscall.GUID{0xe747770b, 0x39ce, 0x4382, 
	[8]byte{0xab, 0x30, 0xd8, 0xfb, 0x3f, 0x33, 0x6f, 0x24}}

type IItemContainerProviderInterface interface {
	IUnknownInterface
	FindItemByProperty(pStartAfter *IRawElementProviderSimple, propertyId int32, value VARIANT, pFound **IRawElementProviderSimple) HRESULT
}

type IItemContainerProviderVtbl struct {
	IUnknownVtbl
	FindItemByProperty uintptr
}

type IItemContainerProvider struct {
	IUnknown
}

func (this *IItemContainerProvider) Vtbl() *IItemContainerProviderVtbl {
	return (*IItemContainerProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IItemContainerProvider) FindItemByProperty(pStartAfter *IRawElementProviderSimple, propertyId int32, value VARIANT, pFound **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindItemByProperty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStartAfter)), uintptr(propertyId), (uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(pFound)))
	return HRESULT(ret)
}

// cb98b665-2d35-4fac-ad35-f3c60d0c0b8b
var IID_IVirtualizedItemProvider = syscall.GUID{0xcb98b665, 0x2d35, 0x4fac, 
	[8]byte{0xad, 0x35, 0xf3, 0xc6, 0x0d, 0x0c, 0x0b, 0x8b}}

type IVirtualizedItemProviderInterface interface {
	IUnknownInterface
	Realize() HRESULT
}

type IVirtualizedItemProviderVtbl struct {
	IUnknownVtbl
	Realize uintptr
}

type IVirtualizedItemProvider struct {
	IUnknown
}

func (this *IVirtualizedItemProvider) Vtbl() *IVirtualizedItemProviderVtbl {
	return (*IVirtualizedItemProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IVirtualizedItemProvider) Realize() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Realize, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 3ad86ebd-f5ef-483d-bb18-b1042a475d64
var IID_IObjectModelProvider = syscall.GUID{0x3ad86ebd, 0xf5ef, 0x483d, 
	[8]byte{0xbb, 0x18, 0xb1, 0x04, 0x2a, 0x47, 0x5d, 0x64}}

type IObjectModelProviderInterface interface {
	IUnknownInterface
	GetUnderlyingObjectModel(ppUnknown **IUnknown) HRESULT
}

type IObjectModelProviderVtbl struct {
	IUnknownVtbl
	GetUnderlyingObjectModel uintptr
}

type IObjectModelProvider struct {
	IUnknown
}

func (this *IObjectModelProvider) Vtbl() *IObjectModelProviderVtbl {
	return (*IObjectModelProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IObjectModelProvider) GetUnderlyingObjectModel(ppUnknown **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetUnderlyingObjectModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppUnknown)))
	return HRESULT(ret)
}

// f95c7e80-bd63-4601-9782-445ebff011fc
var IID_IAnnotationProvider = syscall.GUID{0xf95c7e80, 0xbd63, 0x4601, 
	[8]byte{0x97, 0x82, 0x44, 0x5e, 0xbf, 0xf0, 0x11, 0xfc}}

type IAnnotationProviderInterface interface {
	IUnknownInterface
	Get_AnnotationTypeId(retVal *int32) HRESULT
	Get_AnnotationTypeName(retVal *BSTR) HRESULT
	Get_Author(retVal *BSTR) HRESULT
	Get_DateTime(retVal *BSTR) HRESULT
	Get_Target(retVal **IRawElementProviderSimple) HRESULT
}

type IAnnotationProviderVtbl struct {
	IUnknownVtbl
	Get_AnnotationTypeId uintptr
	Get_AnnotationTypeName uintptr
	Get_Author uintptr
	Get_DateTime uintptr
	Get_Target uintptr
}

type IAnnotationProvider struct {
	IUnknown
}

func (this *IAnnotationProvider) Vtbl() *IAnnotationProviderVtbl {
	return (*IAnnotationProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAnnotationProvider) Get_AnnotationTypeId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_AnnotationTypeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IAnnotationProvider) Get_AnnotationTypeName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_AnnotationTypeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IAnnotationProvider) Get_Author(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Author, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IAnnotationProvider) Get_DateTime(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IAnnotationProvider) Get_Target(retVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Target, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 19b6b649-f5d7-4a6d-bdcb-129252be588a
var IID_IStylesProvider = syscall.GUID{0x19b6b649, 0xf5d7, 0x4a6d, 
	[8]byte{0xbd, 0xcb, 0x12, 0x92, 0x52, 0xbe, 0x58, 0x8a}}

type IStylesProviderInterface interface {
	IUnknownInterface
	Get_StyleId(retVal *int32) HRESULT
	Get_StyleName(retVal *BSTR) HRESULT
	Get_FillColor(retVal *int32) HRESULT
	Get_FillPatternStyle(retVal *BSTR) HRESULT
	Get_Shape(retVal *BSTR) HRESULT
	Get_FillPatternColor(retVal *int32) HRESULT
	Get_ExtendedProperties(retVal *BSTR) HRESULT
}

type IStylesProviderVtbl struct {
	IUnknownVtbl
	Get_StyleId uintptr
	Get_StyleName uintptr
	Get_FillColor uintptr
	Get_FillPatternStyle uintptr
	Get_Shape uintptr
	Get_FillPatternColor uintptr
	Get_ExtendedProperties uintptr
}

type IStylesProvider struct {
	IUnknown
}

func (this *IStylesProvider) Vtbl() *IStylesProviderVtbl {
	return (*IStylesProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStylesProvider) Get_StyleId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_StyleId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IStylesProvider) Get_StyleName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_StyleName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IStylesProvider) Get_FillColor(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_FillColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IStylesProvider) Get_FillPatternStyle(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_FillPatternStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IStylesProvider) Get_Shape(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Shape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IStylesProvider) Get_FillPatternColor(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_FillPatternColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IStylesProvider) Get_ExtendedProperties(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 6f6b5d35-5525-4f80-b758-85473832ffc7
var IID_ISpreadsheetProvider = syscall.GUID{0x6f6b5d35, 0x5525, 0x4f80, 
	[8]byte{0xb7, 0x58, 0x85, 0x47, 0x38, 0x32, 0xff, 0xc7}}

type ISpreadsheetProviderInterface interface {
	IUnknownInterface
	GetItemByName(name PWSTR, pRetVal **IRawElementProviderSimple) HRESULT
}

type ISpreadsheetProviderVtbl struct {
	IUnknownVtbl
	GetItemByName uintptr
}

type ISpreadsheetProvider struct {
	IUnknown
}

func (this *ISpreadsheetProvider) Vtbl() *ISpreadsheetProviderVtbl {
	return (*ISpreadsheetProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpreadsheetProvider) GetItemByName(name PWSTR, pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetItemByName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// eaed4660-7b3d-4879-a2e6-365ce603f3d0
var IID_ISpreadsheetItemProvider = syscall.GUID{0xeaed4660, 0x7b3d, 0x4879, 
	[8]byte{0xa2, 0xe6, 0x36, 0x5c, 0xe6, 0x03, 0xf3, 0xd0}}

type ISpreadsheetItemProviderInterface interface {
	IUnknownInterface
	Get_Formula(pRetVal *BSTR) HRESULT
	GetAnnotationObjects(pRetVal **SAFEARRAY) HRESULT
	GetAnnotationTypes(pRetVal **SAFEARRAY) HRESULT
}

type ISpreadsheetItemProviderVtbl struct {
	IUnknownVtbl
	Get_Formula uintptr
	GetAnnotationObjects uintptr
	GetAnnotationTypes uintptr
}

type ISpreadsheetItemProvider struct {
	IUnknown
}

func (this *ISpreadsheetItemProvider) Vtbl() *ISpreadsheetItemProviderVtbl {
	return (*ISpreadsheetItemProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISpreadsheetItemProvider) Get_Formula(pRetVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Formula, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ISpreadsheetItemProvider) GetAnnotationObjects(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAnnotationObjects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ISpreadsheetItemProvider) GetAnnotationTypes(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAnnotationTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 4758742f-7ac2-460c-bc48-09fc09308a93
var IID_ITransformProvider2 = syscall.GUID{0x4758742f, 0x7ac2, 0x460c, 
	[8]byte{0xbc, 0x48, 0x09, 0xfc, 0x09, 0x30, 0x8a, 0x93}}

type ITransformProvider2Interface interface {
	ITransformProviderInterface
	Zoom(zoom float64) HRESULT
	Get_CanZoom(pRetVal *BOOL) HRESULT
	Get_ZoomLevel(pRetVal *float64) HRESULT
	Get_ZoomMinimum(pRetVal *float64) HRESULT
	Get_ZoomMaximum(pRetVal *float64) HRESULT
	ZoomByUnit(zoomUnit ZoomUnit) HRESULT
}

type ITransformProvider2Vtbl struct {
	ITransformProviderVtbl
	Zoom uintptr
	Get_CanZoom uintptr
	Get_ZoomLevel uintptr
	Get_ZoomMinimum uintptr
	Get_ZoomMaximum uintptr
	ZoomByUnit uintptr
}

type ITransformProvider2 struct {
	ITransformProvider
}

func (this *ITransformProvider2) Vtbl() *ITransformProvider2Vtbl {
	return (*ITransformProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITransformProvider2) Zoom(zoom float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Zoom, uintptr(unsafe.Pointer(this)), uintptr(zoom))
	return HRESULT(ret)
}

func (this *ITransformProvider2) Get_CanZoom(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanZoom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITransformProvider2) Get_ZoomLevel(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ZoomLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITransformProvider2) Get_ZoomMinimum(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ZoomMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITransformProvider2) Get_ZoomMaximum(pRetVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ZoomMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITransformProvider2) ZoomByUnit(zoomUnit ZoomUnit) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ZoomByUnit, uintptr(unsafe.Pointer(this)), uintptr(zoomUnit))
	return HRESULT(ret)
}

// 6aa7bbbb-7ff9-497d-904f-d20b897929d8
var IID_IDragProvider = syscall.GUID{0x6aa7bbbb, 0x7ff9, 0x497d, 
	[8]byte{0x90, 0x4f, 0xd2, 0x0b, 0x89, 0x79, 0x29, 0xd8}}

type IDragProviderInterface interface {
	IUnknownInterface
	Get_IsGrabbed(pRetVal *BOOL) HRESULT
	Get_DropEffect(pRetVal *BSTR) HRESULT
	Get_DropEffects(pRetVal **SAFEARRAY) HRESULT
	GetGrabbedItems(pRetVal **SAFEARRAY) HRESULT
}

type IDragProviderVtbl struct {
	IUnknownVtbl
	Get_IsGrabbed uintptr
	Get_DropEffect uintptr
	Get_DropEffects uintptr
	GetGrabbedItems uintptr
}

type IDragProvider struct {
	IUnknown
}

func (this *IDragProvider) Vtbl() *IDragProviderVtbl {
	return (*IDragProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDragProvider) Get_IsGrabbed(pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_IsGrabbed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IDragProvider) Get_DropEffect(pRetVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DropEffect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IDragProvider) Get_DropEffects(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DropEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IDragProvider) GetGrabbedItems(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetGrabbedItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// bae82bfd-358a-481c-85a0-d8b4d90a5d61
var IID_IDropTargetProvider = syscall.GUID{0xbae82bfd, 0x358a, 0x481c, 
	[8]byte{0x85, 0xa0, 0xd8, 0xb4, 0xd9, 0x0a, 0x5d, 0x61}}

type IDropTargetProviderInterface interface {
	IUnknownInterface
	Get_DropTargetEffect(pRetVal *BSTR) HRESULT
	Get_DropTargetEffects(pRetVal **SAFEARRAY) HRESULT
}

type IDropTargetProviderVtbl struct {
	IUnknownVtbl
	Get_DropTargetEffect uintptr
	Get_DropTargetEffects uintptr
}

type IDropTargetProvider struct {
	IUnknown
}

func (this *IDropTargetProvider) Vtbl() *IDropTargetProviderVtbl {
	return (*IDropTargetProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDropTargetProvider) Get_DropTargetEffect(pRetVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DropTargetEffect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IDropTargetProvider) Get_DropTargetEffects(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DropTargetEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 5347ad7b-c355-46f8-aff5-909033582f63
var IID_ITextRangeProvider = syscall.GUID{0x5347ad7b, 0xc355, 0x46f8, 
	[8]byte{0xaf, 0xf5, 0x90, 0x90, 0x33, 0x58, 0x2f, 0x63}}

type ITextRangeProviderInterface interface {
	IUnknownInterface
	Clone(pRetVal **ITextRangeProvider) HRESULT
	Compare(range_ *ITextRangeProvider, pRetVal *BOOL) HRESULT
	CompareEndpoints(endpoint TextPatternRangeEndpoint, targetRange *ITextRangeProvider, targetEndpoint TextPatternRangeEndpoint, pRetVal *int32) HRESULT
	ExpandToEnclosingUnit(unit TextUnit) HRESULT
	FindAttribute(attributeId int32, val VARIANT, backward BOOL, pRetVal **ITextRangeProvider) HRESULT
	FindText(text BSTR, backward BOOL, ignoreCase BOOL, pRetVal **ITextRangeProvider) HRESULT
	GetAttributeValue(attributeId int32, pRetVal *VARIANT) HRESULT
	GetBoundingRectangles(pRetVal **SAFEARRAY) HRESULT
	GetEnclosingElement(pRetVal **IRawElementProviderSimple) HRESULT
	GetText(maxLength int32, pRetVal *BSTR) HRESULT
	Move(unit TextUnit, count int32, pRetVal *int32) HRESULT
	MoveEndpointByUnit(endpoint TextPatternRangeEndpoint, unit TextUnit, count int32, pRetVal *int32) HRESULT
	MoveEndpointByRange(endpoint TextPatternRangeEndpoint, targetRange *ITextRangeProvider, targetEndpoint TextPatternRangeEndpoint) HRESULT
	Select() HRESULT
	AddToSelection() HRESULT
	RemoveFromSelection() HRESULT
	ScrollIntoView(alignToTop BOOL) HRESULT
	GetChildren(pRetVal **SAFEARRAY) HRESULT
}

type ITextRangeProviderVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Compare uintptr
	CompareEndpoints uintptr
	ExpandToEnclosingUnit uintptr
	FindAttribute uintptr
	FindText uintptr
	GetAttributeValue uintptr
	GetBoundingRectangles uintptr
	GetEnclosingElement uintptr
	GetText uintptr
	Move uintptr
	MoveEndpointByUnit uintptr
	MoveEndpointByRange uintptr
	Select uintptr
	AddToSelection uintptr
	RemoveFromSelection uintptr
	ScrollIntoView uintptr
	GetChildren uintptr
}

type ITextRangeProvider struct {
	IUnknown
}

func (this *ITextRangeProvider) Vtbl() *ITextRangeProviderVtbl {
	return (*ITextRangeProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextRangeProvider) Clone(pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) Compare(range_ *ITextRangeProvider, pRetVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Compare, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(range_)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) CompareEndpoints(endpoint TextPatternRangeEndpoint, targetRange *ITextRangeProvider, targetEndpoint TextPatternRangeEndpoint, pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CompareEndpoints, uintptr(unsafe.Pointer(this)), uintptr(endpoint), uintptr(unsafe.Pointer(targetRange)), uintptr(targetEndpoint), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) ExpandToEnclosingUnit(unit TextUnit) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ExpandToEnclosingUnit, uintptr(unsafe.Pointer(this)), uintptr(unit))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) FindAttribute(attributeId int32, val VARIANT, backward BOOL, pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindAttribute, uintptr(unsafe.Pointer(this)), uintptr(attributeId), (uintptr)(unsafe.Pointer(&val)), uintptr(backward), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) FindText(text BSTR, backward BOOL, ignoreCase BOOL, pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(text)), uintptr(backward), uintptr(ignoreCase), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) GetAttributeValue(attributeId int32, pRetVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAttributeValue, uintptr(unsafe.Pointer(this)), uintptr(attributeId), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) GetBoundingRectangles(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetBoundingRectangles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) GetEnclosingElement(pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEnclosingElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) GetText(maxLength int32, pRetVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetText, uintptr(unsafe.Pointer(this)), uintptr(maxLength), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) Move(unit TextUnit, count int32, pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Move, uintptr(unsafe.Pointer(this)), uintptr(unit), uintptr(count), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) MoveEndpointByUnit(endpoint TextPatternRangeEndpoint, unit TextUnit, count int32, pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().MoveEndpointByUnit, uintptr(unsafe.Pointer(this)), uintptr(endpoint), uintptr(unit), uintptr(count), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) MoveEndpointByRange(endpoint TextPatternRangeEndpoint, targetRange *ITextRangeProvider, targetEndpoint TextPatternRangeEndpoint) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().MoveEndpointByRange, uintptr(unsafe.Pointer(this)), uintptr(endpoint), uintptr(unsafe.Pointer(targetRange)), uintptr(targetEndpoint))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) Select() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Select, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) AddToSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddToSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) RemoveFromSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveFromSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) ScrollIntoView(alignToTop BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ScrollIntoView, uintptr(unsafe.Pointer(this)), uintptr(alignToTop))
	return HRESULT(ret)
}

func (this *ITextRangeProvider) GetChildren(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildren, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 3589c92c-63f3-4367-99bb-ada653b77cf2
var IID_ITextProvider = syscall.GUID{0x3589c92c, 0x63f3, 0x4367, 
	[8]byte{0x99, 0xbb, 0xad, 0xa6, 0x53, 0xb7, 0x7c, 0xf2}}

type ITextProviderInterface interface {
	IUnknownInterface
	GetSelection(pRetVal **SAFEARRAY) HRESULT
	GetVisibleRanges(pRetVal **SAFEARRAY) HRESULT
	RangeFromChild(childElement *IRawElementProviderSimple, pRetVal **ITextRangeProvider) HRESULT
	RangeFromPoint(point UiaPoint, pRetVal **ITextRangeProvider) HRESULT
	Get_DocumentRange(pRetVal **ITextRangeProvider) HRESULT
	Get_SupportedTextSelection(pRetVal *SupportedTextSelection) HRESULT
}

type ITextProviderVtbl struct {
	IUnknownVtbl
	GetSelection uintptr
	GetVisibleRanges uintptr
	RangeFromChild uintptr
	RangeFromPoint uintptr
	Get_DocumentRange uintptr
	Get_SupportedTextSelection uintptr
}

type ITextProvider struct {
	IUnknown
}

func (this *ITextProvider) Vtbl() *ITextProviderVtbl {
	return (*ITextProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextProvider) GetSelection(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextProvider) GetVisibleRanges(pRetVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisibleRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextProvider) RangeFromChild(childElement *IRawElementProviderSimple, pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RangeFromChild, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childElement)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextProvider) RangeFromPoint(point UiaPoint, pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RangeFromPoint, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&point)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextProvider) Get_DocumentRange(pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextProvider) Get_SupportedTextSelection(pRetVal *SupportedTextSelection) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedTextSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 0dc5e6ed-3e16-4bf1-8f9a-a979878bc195
var IID_ITextProvider2 = syscall.GUID{0x0dc5e6ed, 0x3e16, 0x4bf1, 
	[8]byte{0x8f, 0x9a, 0xa9, 0x79, 0x87, 0x8b, 0xc1, 0x95}}

type ITextProvider2Interface interface {
	ITextProviderInterface
	RangeFromAnnotation(annotationElement *IRawElementProviderSimple, pRetVal **ITextRangeProvider) HRESULT
	GetCaretRange(isActive *BOOL, pRetVal **ITextRangeProvider) HRESULT
}

type ITextProvider2Vtbl struct {
	ITextProviderVtbl
	RangeFromAnnotation uintptr
	GetCaretRange uintptr
}

type ITextProvider2 struct {
	ITextProvider
}

func (this *ITextProvider2) Vtbl() *ITextProvider2Vtbl {
	return (*ITextProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextProvider2) RangeFromAnnotation(annotationElement *IRawElementProviderSimple, pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RangeFromAnnotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(annotationElement)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextProvider2) GetCaretRange(isActive *BOOL, pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCaretRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isActive)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// ea3605b4-3a05-400e-b5f9-4e91b40f6176
var IID_ITextEditProvider = syscall.GUID{0xea3605b4, 0x3a05, 0x400e, 
	[8]byte{0xb5, 0xf9, 0x4e, 0x91, 0xb4, 0x0f, 0x61, 0x76}}

type ITextEditProviderInterface interface {
	ITextProviderInterface
	GetActiveComposition(pRetVal **ITextRangeProvider) HRESULT
	GetConversionTarget(pRetVal **ITextRangeProvider) HRESULT
}

type ITextEditProviderVtbl struct {
	ITextProviderVtbl
	GetActiveComposition uintptr
	GetConversionTarget uintptr
}

type ITextEditProvider struct {
	ITextProvider
}

func (this *ITextEditProvider) Vtbl() *ITextEditProviderVtbl {
	return (*ITextEditProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextEditProvider) GetActiveComposition(pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetActiveComposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextEditProvider) GetConversionTarget(pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetConversionTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 9bbce42c-1921-4f18-89ca-dba1910a0386
var IID_ITextRangeProvider2 = syscall.GUID{0x9bbce42c, 0x1921, 0x4f18, 
	[8]byte{0x89, 0xca, 0xdb, 0xa1, 0x91, 0x0a, 0x03, 0x86}}

type ITextRangeProvider2Interface interface {
	ITextRangeProviderInterface
	ShowContextMenu() HRESULT
}

type ITextRangeProvider2Vtbl struct {
	ITextRangeProviderVtbl
	ShowContextMenu uintptr
}

type ITextRangeProvider2 struct {
	ITextRangeProvider
}

func (this *ITextRangeProvider2) Vtbl() *ITextRangeProvider2Vtbl {
	return (*ITextRangeProvider2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextRangeProvider2) ShowContextMenu() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ShowContextMenu, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 4c2de2b9-c88f-4f88-a111-f1d336b7d1a9
var IID_ITextChildProvider = syscall.GUID{0x4c2de2b9, 0xc88f, 0x4f88, 
	[8]byte{0xa1, 0x11, 0xf1, 0xd3, 0x36, 0xb7, 0xd1, 0xa9}}

type ITextChildProviderInterface interface {
	IUnknownInterface
	Get_TextContainer(pRetVal **IRawElementProviderSimple) HRESULT
	Get_TextRange(pRetVal **ITextRangeProvider) HRESULT
}

type ITextChildProviderVtbl struct {
	IUnknownVtbl
	Get_TextContainer uintptr
	Get_TextRange uintptr
}

type ITextChildProvider struct {
	IUnknown
}

func (this *ITextChildProvider) Vtbl() *ITextChildProviderVtbl {
	return (*ITextChildProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ITextChildProvider) Get_TextContainer(pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TextContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *ITextChildProvider) Get_TextRange(pRetVal **ITextRangeProvider) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TextRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// 2062a28a-8c07-4b94-8e12-7037c622aeb8
var IID_ICustomNavigationProvider = syscall.GUID{0x2062a28a, 0x8c07, 0x4b94, 
	[8]byte{0x8e, 0x12, 0x70, 0x37, 0xc6, 0x22, 0xae, 0xb8}}

type ICustomNavigationProviderInterface interface {
	IUnknownInterface
	Navigate(direction NavigateDirection, pRetVal **IRawElementProviderSimple) HRESULT
}

type ICustomNavigationProviderVtbl struct {
	IUnknownVtbl
	Navigate uintptr
}

type ICustomNavigationProvider struct {
	IUnknown
}

func (this *ICustomNavigationProvider) Vtbl() *ICustomNavigationProviderVtbl {
	return (*ICustomNavigationProviderVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICustomNavigationProvider) Navigate(direction NavigateDirection, pRetVal **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Navigate, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// c03a7fe4-9431-409f-bed8-ae7c2299bc8d
var IID_IUIAutomationPatternInstance = syscall.GUID{0xc03a7fe4, 0x9431, 0x409f, 
	[8]byte{0xbe, 0xd8, 0xae, 0x7c, 0x22, 0x99, 0xbc, 0x8d}}

type IUIAutomationPatternInstanceInterface interface {
	IUnknownInterface
	GetProperty(index uint32, cached BOOL, type_ UIAutomationType, pPtr unsafe.Pointer) HRESULT
	CallMethod(index uint32, pParams *UIAutomationParameter, cParams uint32) HRESULT
}

type IUIAutomationPatternInstanceVtbl struct {
	IUnknownVtbl
	GetProperty uintptr
	CallMethod uintptr
}

type IUIAutomationPatternInstance struct {
	IUnknown
}

func (this *IUIAutomationPatternInstance) Vtbl() *IUIAutomationPatternInstanceVtbl {
	return (*IUIAutomationPatternInstanceVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationPatternInstance) GetProperty(index uint32, cached BOOL, type_ UIAutomationType, pPtr unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetProperty, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(cached), uintptr(type_), uintptr(unsafe.Pointer(pPtr)))
	return HRESULT(ret)
}

func (this *IUIAutomationPatternInstance) CallMethod(index uint32, pParams *UIAutomationParameter, cParams uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CallMethod, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(pParams)), uintptr(cParams))
	return HRESULT(ret)
}

// d97022f3-a947-465e-8b2a-ac4315fa54e8
var IID_IUIAutomationPatternHandler = syscall.GUID{0xd97022f3, 0xa947, 0x465e, 
	[8]byte{0x8b, 0x2a, 0xac, 0x43, 0x15, 0xfa, 0x54, 0xe8}}

type IUIAutomationPatternHandlerInterface interface {
	IUnknownInterface
	CreateClientWrapper(pPatternInstance *IUIAutomationPatternInstance, pClientWrapper **IUnknown) HRESULT
	Dispatch(pTarget *IUnknown, index uint32, pParams *UIAutomationParameter, cParams uint32) HRESULT
}

type IUIAutomationPatternHandlerVtbl struct {
	IUnknownVtbl
	CreateClientWrapper uintptr
	Dispatch uintptr
}

type IUIAutomationPatternHandler struct {
	IUnknown
}

func (this *IUIAutomationPatternHandler) Vtbl() *IUIAutomationPatternHandlerVtbl {
	return (*IUIAutomationPatternHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationPatternHandler) CreateClientWrapper(pPatternInstance *IUIAutomationPatternInstance, pClientWrapper **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateClientWrapper, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pPatternInstance)), uintptr(unsafe.Pointer(pClientWrapper)))
	return HRESULT(ret)
}

func (this *IUIAutomationPatternHandler) Dispatch(pTarget *IUnknown, index uint32, pParams *UIAutomationParameter, cParams uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Dispatch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pTarget)), uintptr(index), uintptr(unsafe.Pointer(pParams)), uintptr(cParams))
	return HRESULT(ret)
}

// 8609c4ec-4a1a-4d88-a357-5a66e060e1cf
var IID_IUIAutomationRegistrar = syscall.GUID{0x8609c4ec, 0x4a1a, 0x4d88, 
	[8]byte{0xa3, 0x57, 0x5a, 0x66, 0xe0, 0x60, 0xe1, 0xcf}}

type IUIAutomationRegistrarInterface interface {
	IUnknownInterface
	RegisterProperty(property *UIAutomationPropertyInfo, propertyId *int32) HRESULT
	RegisterEvent(event *UIAutomationEventInfo, eventId *int32) HRESULT
	RegisterPattern(pattern *UIAutomationPatternInfo, pPatternId *int32, pPatternAvailablePropertyId *int32, propertyIdCount uint32, pPropertyIds *int32, eventIdCount uint32, pEventIds *int32) HRESULT
}

type IUIAutomationRegistrarVtbl struct {
	IUnknownVtbl
	RegisterProperty uintptr
	RegisterEvent uintptr
	RegisterPattern uintptr
}

type IUIAutomationRegistrar struct {
	IUnknown
}

func (this *IUIAutomationRegistrar) Vtbl() *IUIAutomationRegistrarVtbl {
	return (*IUIAutomationRegistrarVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationRegistrar) RegisterProperty(property *UIAutomationPropertyInfo, propertyId *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RegisterProperty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(property)), uintptr(unsafe.Pointer(propertyId)))
	return HRESULT(ret)
}

func (this *IUIAutomationRegistrar) RegisterEvent(event *UIAutomationEventInfo, eventId *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RegisterEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(event)), uintptr(unsafe.Pointer(eventId)))
	return HRESULT(ret)
}

func (this *IUIAutomationRegistrar) RegisterPattern(pattern *UIAutomationPatternInfo, pPatternId *int32, pPatternAvailablePropertyId *int32, propertyIdCount uint32, pPropertyIds *int32, eventIdCount uint32, pEventIds *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RegisterPattern, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pattern)), uintptr(unsafe.Pointer(pPatternId)), uintptr(unsafe.Pointer(pPatternAvailablePropertyId)), uintptr(propertyIdCount), uintptr(unsafe.Pointer(pPropertyIds)), uintptr(eventIdCount), uintptr(unsafe.Pointer(pEventIds)))
	return HRESULT(ret)
}

// d22108aa-8ac5-49a5-837b-37bbb3d7591e
var IID_IUIAutomationElement = syscall.GUID{0xd22108aa, 0x8ac5, 0x49a5, 
	[8]byte{0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7, 0x59, 0x1e}}

type IUIAutomationElementInterface interface {
	IUnknownInterface
	SetFocus() HRESULT
	GetRuntimeId(runtimeId **SAFEARRAY) HRESULT
	FindFirst(scope TreeScope, condition *IUIAutomationCondition, found **IUIAutomationElement) HRESULT
	FindAll(scope TreeScope, condition *IUIAutomationCondition, found **IUIAutomationElementArray) HRESULT
	FindFirstBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, found **IUIAutomationElement) HRESULT
	FindAllBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, found **IUIAutomationElementArray) HRESULT
	BuildUpdatedCache(cacheRequest *IUIAutomationCacheRequest, updatedElement **IUIAutomationElement) HRESULT
	GetCurrentPropertyValue(propertyId int32, retVal *VARIANT) HRESULT
	GetCurrentPropertyValueEx(propertyId int32, ignoreDefaultValue BOOL, retVal *VARIANT) HRESULT
	GetCachedPropertyValue(propertyId int32, retVal *VARIANT) HRESULT
	GetCachedPropertyValueEx(propertyId int32, ignoreDefaultValue BOOL, retVal *VARIANT) HRESULT
	GetCurrentPatternAs(patternId int32, riid *syscall.GUID, patternObject unsafe.Pointer) HRESULT
	GetCachedPatternAs(patternId int32, riid *syscall.GUID, patternObject unsafe.Pointer) HRESULT
	GetCurrentPattern(patternId int32, patternObject **IUnknown) HRESULT
	GetCachedPattern(patternId int32, patternObject **IUnknown) HRESULT
	GetCachedParent(parent **IUIAutomationElement) HRESULT
	GetCachedChildren(children **IUIAutomationElementArray) HRESULT
	Get_CurrentProcessId(retVal *int32) HRESULT
	Get_CurrentControlType(retVal *int32) HRESULT
	Get_CurrentLocalizedControlType(retVal *BSTR) HRESULT
	Get_CurrentName(retVal *BSTR) HRESULT
	Get_CurrentAcceleratorKey(retVal *BSTR) HRESULT
	Get_CurrentAccessKey(retVal *BSTR) HRESULT
	Get_CurrentHasKeyboardFocus(retVal *BOOL) HRESULT
	Get_CurrentIsKeyboardFocusable(retVal *BOOL) HRESULT
	Get_CurrentIsEnabled(retVal *BOOL) HRESULT
	Get_CurrentAutomationId(retVal *BSTR) HRESULT
	Get_CurrentClassName(retVal *BSTR) HRESULT
	Get_CurrentHelpText(retVal *BSTR) HRESULT
	Get_CurrentCulture(retVal *int32) HRESULT
	Get_CurrentIsControlElement(retVal *BOOL) HRESULT
	Get_CurrentIsContentElement(retVal *BOOL) HRESULT
	Get_CurrentIsPassword(retVal *BOOL) HRESULT
	Get_CurrentNativeWindowHandle(retVal *HWND) HRESULT
	Get_CurrentItemType(retVal *BSTR) HRESULT
	Get_CurrentIsOffscreen(retVal *BOOL) HRESULT
	Get_CurrentOrientation(retVal *OrientationType) HRESULT
	Get_CurrentFrameworkId(retVal *BSTR) HRESULT
	Get_CurrentIsRequiredForForm(retVal *BOOL) HRESULT
	Get_CurrentItemStatus(retVal *BSTR) HRESULT
	Get_CurrentBoundingRectangle(retVal *RECT) HRESULT
	Get_CurrentLabeledBy(retVal **IUIAutomationElement) HRESULT
	Get_CurrentAriaRole(retVal *BSTR) HRESULT
	Get_CurrentAriaProperties(retVal *BSTR) HRESULT
	Get_CurrentIsDataValidForForm(retVal *BOOL) HRESULT
	Get_CurrentControllerFor(retVal **IUIAutomationElementArray) HRESULT
	Get_CurrentDescribedBy(retVal **IUIAutomationElementArray) HRESULT
	Get_CurrentFlowsTo(retVal **IUIAutomationElementArray) HRESULT
	Get_CurrentProviderDescription(retVal *BSTR) HRESULT
	Get_CachedProcessId(retVal *int32) HRESULT
	Get_CachedControlType(retVal *int32) HRESULT
	Get_CachedLocalizedControlType(retVal *BSTR) HRESULT
	Get_CachedName(retVal *BSTR) HRESULT
	Get_CachedAcceleratorKey(retVal *BSTR) HRESULT
	Get_CachedAccessKey(retVal *BSTR) HRESULT
	Get_CachedHasKeyboardFocus(retVal *BOOL) HRESULT
	Get_CachedIsKeyboardFocusable(retVal *BOOL) HRESULT
	Get_CachedIsEnabled(retVal *BOOL) HRESULT
	Get_CachedAutomationId(retVal *BSTR) HRESULT
	Get_CachedClassName(retVal *BSTR) HRESULT
	Get_CachedHelpText(retVal *BSTR) HRESULT
	Get_CachedCulture(retVal *int32) HRESULT
	Get_CachedIsControlElement(retVal *BOOL) HRESULT
	Get_CachedIsContentElement(retVal *BOOL) HRESULT
	Get_CachedIsPassword(retVal *BOOL) HRESULT
	Get_CachedNativeWindowHandle(retVal *HWND) HRESULT
	Get_CachedItemType(retVal *BSTR) HRESULT
	Get_CachedIsOffscreen(retVal *BOOL) HRESULT
	Get_CachedOrientation(retVal *OrientationType) HRESULT
	Get_CachedFrameworkId(retVal *BSTR) HRESULT
	Get_CachedIsRequiredForForm(retVal *BOOL) HRESULT
	Get_CachedItemStatus(retVal *BSTR) HRESULT
	Get_CachedBoundingRectangle(retVal *RECT) HRESULT
	Get_CachedLabeledBy(retVal **IUIAutomationElement) HRESULT
	Get_CachedAriaRole(retVal *BSTR) HRESULT
	Get_CachedAriaProperties(retVal *BSTR) HRESULT
	Get_CachedIsDataValidForForm(retVal *BOOL) HRESULT
	Get_CachedControllerFor(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedDescribedBy(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedFlowsTo(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedProviderDescription(retVal *BSTR) HRESULT
	GetClickablePoint(clickable *POINT, gotClickable *BOOL) HRESULT
}

type IUIAutomationElementVtbl struct {
	IUnknownVtbl
	SetFocus uintptr
	GetRuntimeId uintptr
	FindFirst uintptr
	FindAll uintptr
	FindFirstBuildCache uintptr
	FindAllBuildCache uintptr
	BuildUpdatedCache uintptr
	GetCurrentPropertyValue uintptr
	GetCurrentPropertyValueEx uintptr
	GetCachedPropertyValue uintptr
	GetCachedPropertyValueEx uintptr
	GetCurrentPatternAs uintptr
	GetCachedPatternAs uintptr
	GetCurrentPattern uintptr
	GetCachedPattern uintptr
	GetCachedParent uintptr
	GetCachedChildren uintptr
	Get_CurrentProcessId uintptr
	Get_CurrentControlType uintptr
	Get_CurrentLocalizedControlType uintptr
	Get_CurrentName uintptr
	Get_CurrentAcceleratorKey uintptr
	Get_CurrentAccessKey uintptr
	Get_CurrentHasKeyboardFocus uintptr
	Get_CurrentIsKeyboardFocusable uintptr
	Get_CurrentIsEnabled uintptr
	Get_CurrentAutomationId uintptr
	Get_CurrentClassName uintptr
	Get_CurrentHelpText uintptr
	Get_CurrentCulture uintptr
	Get_CurrentIsControlElement uintptr
	Get_CurrentIsContentElement uintptr
	Get_CurrentIsPassword uintptr
	Get_CurrentNativeWindowHandle uintptr
	Get_CurrentItemType uintptr
	Get_CurrentIsOffscreen uintptr
	Get_CurrentOrientation uintptr
	Get_CurrentFrameworkId uintptr
	Get_CurrentIsRequiredForForm uintptr
	Get_CurrentItemStatus uintptr
	Get_CurrentBoundingRectangle uintptr
	Get_CurrentLabeledBy uintptr
	Get_CurrentAriaRole uintptr
	Get_CurrentAriaProperties uintptr
	Get_CurrentIsDataValidForForm uintptr
	Get_CurrentControllerFor uintptr
	Get_CurrentDescribedBy uintptr
	Get_CurrentFlowsTo uintptr
	Get_CurrentProviderDescription uintptr
	Get_CachedProcessId uintptr
	Get_CachedControlType uintptr
	Get_CachedLocalizedControlType uintptr
	Get_CachedName uintptr
	Get_CachedAcceleratorKey uintptr
	Get_CachedAccessKey uintptr
	Get_CachedHasKeyboardFocus uintptr
	Get_CachedIsKeyboardFocusable uintptr
	Get_CachedIsEnabled uintptr
	Get_CachedAutomationId uintptr
	Get_CachedClassName uintptr
	Get_CachedHelpText uintptr
	Get_CachedCulture uintptr
	Get_CachedIsControlElement uintptr
	Get_CachedIsContentElement uintptr
	Get_CachedIsPassword uintptr
	Get_CachedNativeWindowHandle uintptr
	Get_CachedItemType uintptr
	Get_CachedIsOffscreen uintptr
	Get_CachedOrientation uintptr
	Get_CachedFrameworkId uintptr
	Get_CachedIsRequiredForForm uintptr
	Get_CachedItemStatus uintptr
	Get_CachedBoundingRectangle uintptr
	Get_CachedLabeledBy uintptr
	Get_CachedAriaRole uintptr
	Get_CachedAriaProperties uintptr
	Get_CachedIsDataValidForForm uintptr
	Get_CachedControllerFor uintptr
	Get_CachedDescribedBy uintptr
	Get_CachedFlowsTo uintptr
	Get_CachedProviderDescription uintptr
	GetClickablePoint uintptr
}

type IUIAutomationElement struct {
	IUnknown
}

func (this *IUIAutomationElement) Vtbl() *IUIAutomationElementVtbl {
	return (*IUIAutomationElementVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement) SetFocus() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetFocus, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetRuntimeId(runtimeId **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRuntimeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(runtimeId)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) FindFirst(scope TreeScope, condition *IUIAutomationCondition, found **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindFirst, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) FindAll(scope TreeScope, condition *IUIAutomationCondition, found **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindAll, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) FindFirstBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, found **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindFirstBuildCache, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) FindAllBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, found **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindAllBuildCache, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) BuildUpdatedCache(cacheRequest *IUIAutomationCacheRequest, updatedElement **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().BuildUpdatedCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(updatedElement)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCurrentPropertyValue(propertyId int32, retVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentPropertyValue, uintptr(unsafe.Pointer(this)), uintptr(propertyId), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCurrentPropertyValueEx(propertyId int32, ignoreDefaultValue BOOL, retVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentPropertyValueEx, uintptr(unsafe.Pointer(this)), uintptr(propertyId), uintptr(ignoreDefaultValue), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCachedPropertyValue(propertyId int32, retVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedPropertyValue, uintptr(unsafe.Pointer(this)), uintptr(propertyId), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCachedPropertyValueEx(propertyId int32, ignoreDefaultValue BOOL, retVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedPropertyValueEx, uintptr(unsafe.Pointer(this)), uintptr(propertyId), uintptr(ignoreDefaultValue), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCurrentPatternAs(patternId int32, riid *syscall.GUID, patternObject unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentPatternAs, uintptr(unsafe.Pointer(this)), uintptr(patternId), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(patternObject)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCachedPatternAs(patternId int32, riid *syscall.GUID, patternObject unsafe.Pointer) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedPatternAs, uintptr(unsafe.Pointer(this)), uintptr(patternId), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(patternObject)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCurrentPattern(patternId int32, patternObject **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentPattern, uintptr(unsafe.Pointer(this)), uintptr(patternId), uintptr(unsafe.Pointer(patternObject)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCachedPattern(patternId int32, patternObject **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedPattern, uintptr(unsafe.Pointer(this)), uintptr(patternId), uintptr(unsafe.Pointer(patternObject)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCachedParent(parent **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedParent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parent)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetCachedChildren(children **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedChildren, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(children)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentProcessId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentProcessId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentControlType(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentControlType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentLocalizedControlType(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLocalizedControlType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentAcceleratorKey(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAcceleratorKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentAccessKey(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAccessKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentHasKeyboardFocus(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHasKeyboardFocus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsKeyboardFocusable(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsKeyboardFocusable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsEnabled(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentAutomationId(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAutomationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentClassName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentClassName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentHelpText(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHelpText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentCulture(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCulture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsControlElement(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsControlElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsContentElement(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsContentElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsPassword(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsPassword, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentNativeWindowHandle(retVal *HWND) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentNativeWindowHandle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentItemType(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentItemType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsOffscreen(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsOffscreen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentOrientation(retVal *OrientationType) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentFrameworkId(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFrameworkId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsRequiredForForm(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsRequiredForForm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentItemStatus(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentItemStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentBoundingRectangle(retVal *RECT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentBoundingRectangle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentLabeledBy(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLabeledBy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentAriaRole(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAriaRole, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentAriaProperties(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAriaProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentIsDataValidForForm(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsDataValidForForm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentControllerFor(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentControllerFor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentDescribedBy(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDescribedBy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentFlowsTo(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFlowsTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CurrentProviderDescription(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentProviderDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedProcessId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedProcessId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedControlType(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedControlType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedLocalizedControlType(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLocalizedControlType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedAcceleratorKey(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAcceleratorKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedAccessKey(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAccessKey, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedHasKeyboardFocus(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHasKeyboardFocus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsKeyboardFocusable(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsKeyboardFocusable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsEnabled(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedAutomationId(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAutomationId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedClassName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedClassName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedHelpText(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHelpText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedCulture(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCulture, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsControlElement(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsControlElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsContentElement(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsContentElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsPassword(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsPassword, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedNativeWindowHandle(retVal *HWND) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedNativeWindowHandle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedItemType(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedItemType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsOffscreen(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsOffscreen, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedOrientation(retVal *OrientationType) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedOrientation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedFrameworkId(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFrameworkId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsRequiredForForm(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsRequiredForForm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedItemStatus(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedItemStatus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedBoundingRectangle(retVal *RECT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedBoundingRectangle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedLabeledBy(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLabeledBy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedAriaRole(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAriaRole, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedAriaProperties(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAriaProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedIsDataValidForForm(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsDataValidForForm, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedControllerFor(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedControllerFor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedDescribedBy(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDescribedBy, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedFlowsTo(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFlowsTo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) Get_CachedProviderDescription(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedProviderDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement) GetClickablePoint(clickable *POINT, gotClickable *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetClickablePoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clickable)), uintptr(unsafe.Pointer(gotClickable)))
	return HRESULT(ret)
}

// 14314595-b4bc-4055-95f2-58f2e42c9855
var IID_IUIAutomationElementArray = syscall.GUID{0x14314595, 0xb4bc, 0x4055, 
	[8]byte{0x95, 0xf2, 0x58, 0xf2, 0xe4, 0x2c, 0x98, 0x55}}

type IUIAutomationElementArrayInterface interface {
	IUnknownInterface
	Get_Length(length *int32) HRESULT
	GetElement(index int32, element **IUIAutomationElement) HRESULT
}

type IUIAutomationElementArrayVtbl struct {
	IUnknownVtbl
	Get_Length uintptr
	GetElement uintptr
}

type IUIAutomationElementArray struct {
	IUnknown
}

func (this *IUIAutomationElementArray) Vtbl() *IUIAutomationElementArrayVtbl {
	return (*IUIAutomationElementArrayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElementArray) Get_Length(length *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(length)))
	return HRESULT(ret)
}

func (this *IUIAutomationElementArray) GetElement(index int32, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetElement, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

// 352ffba8-0973-437c-a61f-f64cafd81df9
var IID_IUIAutomationCondition = syscall.GUID{0x352ffba8, 0x0973, 0x437c, 
	[8]byte{0xa6, 0x1f, 0xf6, 0x4c, 0xaf, 0xd8, 0x1d, 0xf9}}

type IUIAutomationConditionInterface interface {
	IUnknownInterface
}

type IUIAutomationConditionVtbl struct {
	IUnknownVtbl
}

type IUIAutomationCondition struct {
	IUnknown
}

func (this *IUIAutomationCondition) Vtbl() *IUIAutomationConditionVtbl {
	return (*IUIAutomationConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

// 1b4e1f2e-75eb-4d0b-8952-5a69988e2307
var IID_IUIAutomationBoolCondition = syscall.GUID{0x1b4e1f2e, 0x75eb, 0x4d0b, 
	[8]byte{0x89, 0x52, 0x5a, 0x69, 0x98, 0x8e, 0x23, 0x07}}

type IUIAutomationBoolConditionInterface interface {
	IUIAutomationConditionInterface
	Get_BooleanValue(boolVal *BOOL) HRESULT
}

type IUIAutomationBoolConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_BooleanValue uintptr
}

type IUIAutomationBoolCondition struct {
	IUIAutomationCondition
}

func (this *IUIAutomationBoolCondition) Vtbl() *IUIAutomationBoolConditionVtbl {
	return (*IUIAutomationBoolConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationBoolCondition) Get_BooleanValue(boolVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_BooleanValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(boolVal)))
	return HRESULT(ret)
}

// 99ebf2cb-5578-4267-9ad4-afd6ea77e94b
var IID_IUIAutomationPropertyCondition = syscall.GUID{0x99ebf2cb, 0x5578, 0x4267, 
	[8]byte{0x9a, 0xd4, 0xaf, 0xd6, 0xea, 0x77, 0xe9, 0x4b}}

type IUIAutomationPropertyConditionInterface interface {
	IUIAutomationConditionInterface
	Get_PropertyId(propertyId *int32) HRESULT
	Get_PropertyValue(propertyValue *VARIANT) HRESULT
	Get_PropertyConditionFlags(flags *PropertyConditionFlags) HRESULT
}

type IUIAutomationPropertyConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_PropertyId uintptr
	Get_PropertyValue uintptr
	Get_PropertyConditionFlags uintptr
}

type IUIAutomationPropertyCondition struct {
	IUIAutomationCondition
}

func (this *IUIAutomationPropertyCondition) Vtbl() *IUIAutomationPropertyConditionVtbl {
	return (*IUIAutomationPropertyConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationPropertyCondition) Get_PropertyId(propertyId *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_PropertyId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertyId)))
	return HRESULT(ret)
}

func (this *IUIAutomationPropertyCondition) Get_PropertyValue(propertyValue *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_PropertyValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertyValue)))
	return HRESULT(ret)
}

func (this *IUIAutomationPropertyCondition) Get_PropertyConditionFlags(flags *PropertyConditionFlags) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_PropertyConditionFlags, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(flags)))
	return HRESULT(ret)
}

// a7d0af36-b912-45fe-9855-091ddc174aec
var IID_IUIAutomationAndCondition = syscall.GUID{0xa7d0af36, 0xb912, 0x45fe, 
	[8]byte{0x98, 0x55, 0x09, 0x1d, 0xdc, 0x17, 0x4a, 0xec}}

type IUIAutomationAndConditionInterface interface {
	IUIAutomationConditionInterface
	Get_ChildCount(childCount *int32) HRESULT
	GetChildrenAsNativeArray(childArray ***IUIAutomationCondition, childArrayCount *int32) HRESULT
	GetChildren(childArray **SAFEARRAY) HRESULT
}

type IUIAutomationAndConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_ChildCount uintptr
	GetChildrenAsNativeArray uintptr
	GetChildren uintptr
}

type IUIAutomationAndCondition struct {
	IUIAutomationCondition
}

func (this *IUIAutomationAndCondition) Vtbl() *IUIAutomationAndConditionVtbl {
	return (*IUIAutomationAndConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationAndCondition) Get_ChildCount(childCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ChildCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childCount)))
	return HRESULT(ret)
}

func (this *IUIAutomationAndCondition) GetChildrenAsNativeArray(childArray ***IUIAutomationCondition, childArrayCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildrenAsNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childArray)), uintptr(unsafe.Pointer(childArrayCount)))
	return HRESULT(ret)
}

func (this *IUIAutomationAndCondition) GetChildren(childArray **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildren, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childArray)))
	return HRESULT(ret)
}

// 8753f032-3db1-47b5-a1fc-6e34a266c712
var IID_IUIAutomationOrCondition = syscall.GUID{0x8753f032, 0x3db1, 0x47b5, 
	[8]byte{0xa1, 0xfc, 0x6e, 0x34, 0xa2, 0x66, 0xc7, 0x12}}

type IUIAutomationOrConditionInterface interface {
	IUIAutomationConditionInterface
	Get_ChildCount(childCount *int32) HRESULT
	GetChildrenAsNativeArray(childArray ***IUIAutomationCondition, childArrayCount *int32) HRESULT
	GetChildren(childArray **SAFEARRAY) HRESULT
}

type IUIAutomationOrConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_ChildCount uintptr
	GetChildrenAsNativeArray uintptr
	GetChildren uintptr
}

type IUIAutomationOrCondition struct {
	IUIAutomationCondition
}

func (this *IUIAutomationOrCondition) Vtbl() *IUIAutomationOrConditionVtbl {
	return (*IUIAutomationOrConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationOrCondition) Get_ChildCount(childCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ChildCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childCount)))
	return HRESULT(ret)
}

func (this *IUIAutomationOrCondition) GetChildrenAsNativeArray(childArray ***IUIAutomationCondition, childArrayCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildrenAsNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childArray)), uintptr(unsafe.Pointer(childArrayCount)))
	return HRESULT(ret)
}

func (this *IUIAutomationOrCondition) GetChildren(childArray **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildren, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(childArray)))
	return HRESULT(ret)
}

// f528b657-847b-498c-8896-d52b565407a1
var IID_IUIAutomationNotCondition = syscall.GUID{0xf528b657, 0x847b, 0x498c, 
	[8]byte{0x88, 0x96, 0xd5, 0x2b, 0x56, 0x54, 0x07, 0xa1}}

type IUIAutomationNotConditionInterface interface {
	IUIAutomationConditionInterface
	GetChild(condition **IUIAutomationCondition) HRESULT
}

type IUIAutomationNotConditionVtbl struct {
	IUIAutomationConditionVtbl
	GetChild uintptr
}

type IUIAutomationNotCondition struct {
	IUIAutomationCondition
}

func (this *IUIAutomationNotCondition) Vtbl() *IUIAutomationNotConditionVtbl {
	return (*IUIAutomationNotConditionVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationNotCondition) GetChild(condition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChild, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)))
	return HRESULT(ret)
}

// b32a92b5-bc25-4078-9c08-d7ee95c48e03
var IID_IUIAutomationCacheRequest = syscall.GUID{0xb32a92b5, 0xbc25, 0x4078, 
	[8]byte{0x9c, 0x08, 0xd7, 0xee, 0x95, 0xc4, 0x8e, 0x03}}

type IUIAutomationCacheRequestInterface interface {
	IUnknownInterface
	AddProperty(propertyId int32) HRESULT
	AddPattern(patternId int32) HRESULT
	Clone(clonedRequest **IUIAutomationCacheRequest) HRESULT
	Get_TreeScope(scope *TreeScope) HRESULT
	Put_TreeScope(scope TreeScope) HRESULT
	Get_TreeFilter(filter **IUIAutomationCondition) HRESULT
	Put_TreeFilter(filter *IUIAutomationCondition) HRESULT
	Get_AutomationElementMode(mode *AutomationElementMode) HRESULT
	Put_AutomationElementMode(mode AutomationElementMode) HRESULT
}

type IUIAutomationCacheRequestVtbl struct {
	IUnknownVtbl
	AddProperty uintptr
	AddPattern uintptr
	Clone uintptr
	Get_TreeScope uintptr
	Put_TreeScope uintptr
	Get_TreeFilter uintptr
	Put_TreeFilter uintptr
	Get_AutomationElementMode uintptr
	Put_AutomationElementMode uintptr
}

type IUIAutomationCacheRequest struct {
	IUnknown
}

func (this *IUIAutomationCacheRequest) Vtbl() *IUIAutomationCacheRequestVtbl {
	return (*IUIAutomationCacheRequestVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationCacheRequest) AddProperty(propertyId int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddProperty, uintptr(unsafe.Pointer(this)), uintptr(propertyId))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) AddPattern(patternId int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddPattern, uintptr(unsafe.Pointer(this)), uintptr(patternId))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Clone(clonedRequest **IUIAutomationCacheRequest) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clonedRequest)))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Get_TreeScope(scope *TreeScope) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TreeScope, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scope)))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Put_TreeScope(scope TreeScope) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_TreeScope, uintptr(unsafe.Pointer(this)), uintptr(scope))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Get_TreeFilter(filter **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TreeFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filter)))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Put_TreeFilter(filter *IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_TreeFilter, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(filter)))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Get_AutomationElementMode(mode *AutomationElementMode) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_AutomationElementMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mode)))
	return HRESULT(ret)
}

func (this *IUIAutomationCacheRequest) Put_AutomationElementMode(mode AutomationElementMode) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_AutomationElementMode, uintptr(unsafe.Pointer(this)), uintptr(mode))
	return HRESULT(ret)
}

// 4042c624-389c-4afc-a630-9df854a541fc
var IID_IUIAutomationTreeWalker = syscall.GUID{0x4042c624, 0x389c, 0x4afc, 
	[8]byte{0xa6, 0x30, 0x9d, 0xf8, 0x54, 0xa5, 0x41, 0xfc}}

type IUIAutomationTreeWalkerInterface interface {
	IUnknownInterface
	GetParentElement(element *IUIAutomationElement, parent **IUIAutomationElement) HRESULT
	GetFirstChildElement(element *IUIAutomationElement, first **IUIAutomationElement) HRESULT
	GetLastChildElement(element *IUIAutomationElement, last **IUIAutomationElement) HRESULT
	GetNextSiblingElement(element *IUIAutomationElement, next **IUIAutomationElement) HRESULT
	GetPreviousSiblingElement(element *IUIAutomationElement, previous **IUIAutomationElement) HRESULT
	NormalizeElement(element *IUIAutomationElement, normalized **IUIAutomationElement) HRESULT
	GetParentElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, parent **IUIAutomationElement) HRESULT
	GetFirstChildElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, first **IUIAutomationElement) HRESULT
	GetLastChildElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, last **IUIAutomationElement) HRESULT
	GetNextSiblingElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, next **IUIAutomationElement) HRESULT
	GetPreviousSiblingElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, previous **IUIAutomationElement) HRESULT
	NormalizeElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, normalized **IUIAutomationElement) HRESULT
	Get_Condition(condition **IUIAutomationCondition) HRESULT
}

type IUIAutomationTreeWalkerVtbl struct {
	IUnknownVtbl
	GetParentElement uintptr
	GetFirstChildElement uintptr
	GetLastChildElement uintptr
	GetNextSiblingElement uintptr
	GetPreviousSiblingElement uintptr
	NormalizeElement uintptr
	GetParentElementBuildCache uintptr
	GetFirstChildElementBuildCache uintptr
	GetLastChildElementBuildCache uintptr
	GetNextSiblingElementBuildCache uintptr
	GetPreviousSiblingElementBuildCache uintptr
	NormalizeElementBuildCache uintptr
	Get_Condition uintptr
}

type IUIAutomationTreeWalker struct {
	IUnknown
}

func (this *IUIAutomationTreeWalker) Vtbl() *IUIAutomationTreeWalkerVtbl {
	return (*IUIAutomationTreeWalkerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTreeWalker) GetParentElement(element *IUIAutomationElement, parent **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetParentElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(parent)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetFirstChildElement(element *IUIAutomationElement, first **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFirstChildElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(first)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetLastChildElement(element *IUIAutomationElement, last **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLastChildElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(last)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetNextSiblingElement(element *IUIAutomationElement, next **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetNextSiblingElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(next)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetPreviousSiblingElement(element *IUIAutomationElement, previous **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPreviousSiblingElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(previous)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) NormalizeElement(element *IUIAutomationElement, normalized **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().NormalizeElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(normalized)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetParentElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, parent **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetParentElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(parent)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetFirstChildElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, first **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFirstChildElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(first)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetLastChildElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, last **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetLastChildElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(last)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetNextSiblingElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, next **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetNextSiblingElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(next)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) GetPreviousSiblingElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, previous **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPreviousSiblingElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(previous)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) NormalizeElementBuildCache(element *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest, normalized **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().NormalizeElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(normalized)))
	return HRESULT(ret)
}

func (this *IUIAutomationTreeWalker) Get_Condition(condition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Condition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)))
	return HRESULT(ret)
}

// 146c3c17-f12e-4e22-8c27-f894b9b79c69
var IID_IUIAutomationEventHandler = syscall.GUID{0x146c3c17, 0xf12e, 0x4e22, 
	[8]byte{0x8c, 0x27, 0xf8, 0x94, 0xb9, 0xb7, 0x9c, 0x69}}

type IUIAutomationEventHandlerInterface interface {
	IUnknownInterface
	HandleAutomationEvent(sender *IUIAutomationElement, eventId int32) HRESULT
}

type IUIAutomationEventHandlerVtbl struct {
	IUnknownVtbl
	HandleAutomationEvent uintptr
}

type IUIAutomationEventHandler struct {
	IUnknown
}

func (this *IUIAutomationEventHandler) Vtbl() *IUIAutomationEventHandlerVtbl {
	return (*IUIAutomationEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationEventHandler) HandleAutomationEvent(sender *IUIAutomationElement, eventId int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleAutomationEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(eventId))
	return HRESULT(ret)
}

// 40cd37d4-c756-4b0c-8c6f-bddfeeb13b50
var IID_IUIAutomationPropertyChangedEventHandler = syscall.GUID{0x40cd37d4, 0xc756, 0x4b0c, 
	[8]byte{0x8c, 0x6f, 0xbd, 0xdf, 0xee, 0xb1, 0x3b, 0x50}}

type IUIAutomationPropertyChangedEventHandlerInterface interface {
	IUnknownInterface
	HandlePropertyChangedEvent(sender *IUIAutomationElement, propertyId int32, newValue VARIANT) HRESULT
}

type IUIAutomationPropertyChangedEventHandlerVtbl struct {
	IUnknownVtbl
	HandlePropertyChangedEvent uintptr
}

type IUIAutomationPropertyChangedEventHandler struct {
	IUnknown
}

func (this *IUIAutomationPropertyChangedEventHandler) Vtbl() *IUIAutomationPropertyChangedEventHandlerVtbl {
	return (*IUIAutomationPropertyChangedEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationPropertyChangedEventHandler) HandlePropertyChangedEvent(sender *IUIAutomationElement, propertyId int32, newValue VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandlePropertyChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(propertyId), (uintptr)(unsafe.Pointer(&newValue)))
	return HRESULT(ret)
}

// e81d1b4e-11c5-42f8-9754-e7036c79f054
var IID_IUIAutomationStructureChangedEventHandler = syscall.GUID{0xe81d1b4e, 0x11c5, 0x42f8, 
	[8]byte{0x97, 0x54, 0xe7, 0x03, 0x6c, 0x79, 0xf0, 0x54}}

type IUIAutomationStructureChangedEventHandlerInterface interface {
	IUnknownInterface
	HandleStructureChangedEvent(sender *IUIAutomationElement, changeType StructureChangeType, runtimeId *SAFEARRAY) HRESULT
}

type IUIAutomationStructureChangedEventHandlerVtbl struct {
	IUnknownVtbl
	HandleStructureChangedEvent uintptr
}

type IUIAutomationStructureChangedEventHandler struct {
	IUnknown
}

func (this *IUIAutomationStructureChangedEventHandler) Vtbl() *IUIAutomationStructureChangedEventHandlerVtbl {
	return (*IUIAutomationStructureChangedEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationStructureChangedEventHandler) HandleStructureChangedEvent(sender *IUIAutomationElement, changeType StructureChangeType, runtimeId *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleStructureChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(changeType), uintptr(unsafe.Pointer(runtimeId)))
	return HRESULT(ret)
}

// c270f6b5-5c69-4290-9745-7a7f97169468
var IID_IUIAutomationFocusChangedEventHandler = syscall.GUID{0xc270f6b5, 0x5c69, 0x4290, 
	[8]byte{0x97, 0x45, 0x7a, 0x7f, 0x97, 0x16, 0x94, 0x68}}

type IUIAutomationFocusChangedEventHandlerInterface interface {
	IUnknownInterface
	HandleFocusChangedEvent(sender *IUIAutomationElement) HRESULT
}

type IUIAutomationFocusChangedEventHandlerVtbl struct {
	IUnknownVtbl
	HandleFocusChangedEvent uintptr
}

type IUIAutomationFocusChangedEventHandler struct {
	IUnknown
}

func (this *IUIAutomationFocusChangedEventHandler) Vtbl() *IUIAutomationFocusChangedEventHandlerVtbl {
	return (*IUIAutomationFocusChangedEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationFocusChangedEventHandler) HandleFocusChangedEvent(sender *IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleFocusChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)))
	return HRESULT(ret)
}

// 92faa680-e704-4156-931a-e32d5bb38f3f
var IID_IUIAutomationTextEditTextChangedEventHandler = syscall.GUID{0x92faa680, 0xe704, 0x4156, 
	[8]byte{0x93, 0x1a, 0xe3, 0x2d, 0x5b, 0xb3, 0x8f, 0x3f}}

type IUIAutomationTextEditTextChangedEventHandlerInterface interface {
	IUnknownInterface
	HandleTextEditTextChangedEvent(sender *IUIAutomationElement, textEditChangeType TextEditChangeType, eventStrings *SAFEARRAY) HRESULT
}

type IUIAutomationTextEditTextChangedEventHandlerVtbl struct {
	IUnknownVtbl
	HandleTextEditTextChangedEvent uintptr
}

type IUIAutomationTextEditTextChangedEventHandler struct {
	IUnknown
}

func (this *IUIAutomationTextEditTextChangedEventHandler) Vtbl() *IUIAutomationTextEditTextChangedEventHandlerVtbl {
	return (*IUIAutomationTextEditTextChangedEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextEditTextChangedEventHandler) HandleTextEditTextChangedEvent(sender *IUIAutomationElement, textEditChangeType TextEditChangeType, eventStrings *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleTextEditTextChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(textEditChangeType), uintptr(unsafe.Pointer(eventStrings)))
	return HRESULT(ret)
}

// 58edca55-2c3e-4980-b1b9-56c17f27a2a0
var IID_IUIAutomationChangesEventHandler = syscall.GUID{0x58edca55, 0x2c3e, 0x4980, 
	[8]byte{0xb1, 0xb9, 0x56, 0xc1, 0x7f, 0x27, 0xa2, 0xa0}}

type IUIAutomationChangesEventHandlerInterface interface {
	IUnknownInterface
	HandleChangesEvent(sender *IUIAutomationElement, uiaChanges *UiaChangeInfo, changesCount int32) HRESULT
}

type IUIAutomationChangesEventHandlerVtbl struct {
	IUnknownVtbl
	HandleChangesEvent uintptr
}

type IUIAutomationChangesEventHandler struct {
	IUnknown
}

func (this *IUIAutomationChangesEventHandler) Vtbl() *IUIAutomationChangesEventHandlerVtbl {
	return (*IUIAutomationChangesEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationChangesEventHandler) HandleChangesEvent(sender *IUIAutomationElement, uiaChanges *UiaChangeInfo, changesCount int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleChangesEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(uiaChanges)), uintptr(changesCount))
	return HRESULT(ret)
}

// c7cb2637-e6c2-4d0c-85de-4948c02175c7
var IID_IUIAutomationNotificationEventHandler = syscall.GUID{0xc7cb2637, 0xe6c2, 0x4d0c, 
	[8]byte{0x85, 0xde, 0x49, 0x48, 0xc0, 0x21, 0x75, 0xc7}}

type IUIAutomationNotificationEventHandlerInterface interface {
	IUnknownInterface
	HandleNotificationEvent(sender *IUIAutomationElement, notificationKind NotificationKind, notificationProcessing NotificationProcessing, displayString BSTR, activityId BSTR) HRESULT
}

type IUIAutomationNotificationEventHandlerVtbl struct {
	IUnknownVtbl
	HandleNotificationEvent uintptr
}

type IUIAutomationNotificationEventHandler struct {
	IUnknown
}

func (this *IUIAutomationNotificationEventHandler) Vtbl() *IUIAutomationNotificationEventHandlerVtbl {
	return (*IUIAutomationNotificationEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationNotificationEventHandler) HandleNotificationEvent(sender *IUIAutomationElement, notificationKind NotificationKind, notificationProcessing NotificationProcessing, displayString BSTR, activityId BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleNotificationEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(notificationKind), uintptr(notificationProcessing), uintptr(unsafe.Pointer(displayString)), uintptr(unsafe.Pointer(activityId)))
	return HRESULT(ret)
}

// fb377fbe-8ea6-46d5-9c73-6499642d3059
var IID_IUIAutomationInvokePattern = syscall.GUID{0xfb377fbe, 0x8ea6, 0x46d5, 
	[8]byte{0x9c, 0x73, 0x64, 0x99, 0x64, 0x2d, 0x30, 0x59}}

type IUIAutomationInvokePatternInterface interface {
	IUnknownInterface
	Invoke() HRESULT
}

type IUIAutomationInvokePatternVtbl struct {
	IUnknownVtbl
	Invoke uintptr
}

type IUIAutomationInvokePattern struct {
	IUnknown
}

func (this *IUIAutomationInvokePattern) Vtbl() *IUIAutomationInvokePatternVtbl {
	return (*IUIAutomationInvokePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationInvokePattern) Invoke() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Invoke, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// fde5ef97-1464-48f6-90bf-43d0948e86ec
var IID_IUIAutomationDockPattern = syscall.GUID{0xfde5ef97, 0x1464, 0x48f6, 
	[8]byte{0x90, 0xbf, 0x43, 0xd0, 0x94, 0x8e, 0x86, 0xec}}

type IUIAutomationDockPatternInterface interface {
	IUnknownInterface
	SetDockPosition(dockPos DockPosition) HRESULT
	Get_CurrentDockPosition(retVal *DockPosition) HRESULT
	Get_CachedDockPosition(retVal *DockPosition) HRESULT
}

type IUIAutomationDockPatternVtbl struct {
	IUnknownVtbl
	SetDockPosition uintptr
	Get_CurrentDockPosition uintptr
	Get_CachedDockPosition uintptr
}

type IUIAutomationDockPattern struct {
	IUnknown
}

func (this *IUIAutomationDockPattern) Vtbl() *IUIAutomationDockPatternVtbl {
	return (*IUIAutomationDockPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationDockPattern) SetDockPosition(dockPos DockPosition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetDockPosition, uintptr(unsafe.Pointer(this)), uintptr(dockPos))
	return HRESULT(ret)
}

func (this *IUIAutomationDockPattern) Get_CurrentDockPosition(retVal *DockPosition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDockPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDockPattern) Get_CachedDockPosition(retVal *DockPosition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDockPosition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 619be086-1f4e-4ee4-bafa-210128738730
var IID_IUIAutomationExpandCollapsePattern = syscall.GUID{0x619be086, 0x1f4e, 0x4ee4, 
	[8]byte{0xba, 0xfa, 0x21, 0x01, 0x28, 0x73, 0x87, 0x30}}

type IUIAutomationExpandCollapsePatternInterface interface {
	IUnknownInterface
	Expand() HRESULT
	Collapse() HRESULT
	Get_CurrentExpandCollapseState(retVal *ExpandCollapseState) HRESULT
	Get_CachedExpandCollapseState(retVal *ExpandCollapseState) HRESULT
}

type IUIAutomationExpandCollapsePatternVtbl struct {
	IUnknownVtbl
	Expand uintptr
	Collapse uintptr
	Get_CurrentExpandCollapseState uintptr
	Get_CachedExpandCollapseState uintptr
}

type IUIAutomationExpandCollapsePattern struct {
	IUnknown
}

func (this *IUIAutomationExpandCollapsePattern) Vtbl() *IUIAutomationExpandCollapsePatternVtbl {
	return (*IUIAutomationExpandCollapsePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationExpandCollapsePattern) Expand() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Expand, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationExpandCollapsePattern) Collapse() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Collapse, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationExpandCollapsePattern) Get_CurrentExpandCollapseState(retVal *ExpandCollapseState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentExpandCollapseState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationExpandCollapsePattern) Get_CachedExpandCollapseState(retVal *ExpandCollapseState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedExpandCollapseState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 414c3cdc-856b-4f5b-8538-3131c6302550
var IID_IUIAutomationGridPattern = syscall.GUID{0x414c3cdc, 0x856b, 0x4f5b, 
	[8]byte{0x85, 0x38, 0x31, 0x31, 0xc6, 0x30, 0x25, 0x50}}

type IUIAutomationGridPatternInterface interface {
	IUnknownInterface
	GetItem(row int32, column int32, element **IUIAutomationElement) HRESULT
	Get_CurrentRowCount(retVal *int32) HRESULT
	Get_CurrentColumnCount(retVal *int32) HRESULT
	Get_CachedRowCount(retVal *int32) HRESULT
	Get_CachedColumnCount(retVal *int32) HRESULT
}

type IUIAutomationGridPatternVtbl struct {
	IUnknownVtbl
	GetItem uintptr
	Get_CurrentRowCount uintptr
	Get_CurrentColumnCount uintptr
	Get_CachedRowCount uintptr
	Get_CachedColumnCount uintptr
}

type IUIAutomationGridPattern struct {
	IUnknown
}

func (this *IUIAutomationGridPattern) Vtbl() *IUIAutomationGridPatternVtbl {
	return (*IUIAutomationGridPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationGridPattern) GetItem(row int32, column int32, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetItem, uintptr(unsafe.Pointer(this)), uintptr(row), uintptr(column), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridPattern) Get_CurrentRowCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentRowCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridPattern) Get_CurrentColumnCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentColumnCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridPattern) Get_CachedRowCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedRowCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridPattern) Get_CachedColumnCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedColumnCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 78f8ef57-66c3-4e09-bd7c-e79b2004894d
var IID_IUIAutomationGridItemPattern = syscall.GUID{0x78f8ef57, 0x66c3, 0x4e09, 
	[8]byte{0xbd, 0x7c, 0xe7, 0x9b, 0x20, 0x04, 0x89, 0x4d}}

type IUIAutomationGridItemPatternInterface interface {
	IUnknownInterface
	Get_CurrentContainingGrid(retVal **IUIAutomationElement) HRESULT
	Get_CurrentRow(retVal *int32) HRESULT
	Get_CurrentColumn(retVal *int32) HRESULT
	Get_CurrentRowSpan(retVal *int32) HRESULT
	Get_CurrentColumnSpan(retVal *int32) HRESULT
	Get_CachedContainingGrid(retVal **IUIAutomationElement) HRESULT
	Get_CachedRow(retVal *int32) HRESULT
	Get_CachedColumn(retVal *int32) HRESULT
	Get_CachedRowSpan(retVal *int32) HRESULT
	Get_CachedColumnSpan(retVal *int32) HRESULT
}

type IUIAutomationGridItemPatternVtbl struct {
	IUnknownVtbl
	Get_CurrentContainingGrid uintptr
	Get_CurrentRow uintptr
	Get_CurrentColumn uintptr
	Get_CurrentRowSpan uintptr
	Get_CurrentColumnSpan uintptr
	Get_CachedContainingGrid uintptr
	Get_CachedRow uintptr
	Get_CachedColumn uintptr
	Get_CachedRowSpan uintptr
	Get_CachedColumnSpan uintptr
}

type IUIAutomationGridItemPattern struct {
	IUnknown
}

func (this *IUIAutomationGridItemPattern) Vtbl() *IUIAutomationGridItemPatternVtbl {
	return (*IUIAutomationGridItemPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationGridItemPattern) Get_CurrentContainingGrid(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentContainingGrid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CurrentRow(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentRow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CurrentColumn(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentColumn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CurrentRowSpan(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentRowSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CurrentColumnSpan(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentColumnSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CachedContainingGrid(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedContainingGrid, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CachedRow(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedRow, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CachedColumn(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedColumn, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CachedRowSpan(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedRowSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationGridItemPattern) Get_CachedColumnSpan(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedColumnSpan, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 8d253c91-1dc5-4bb5-b18f-ade16fa495e8
var IID_IUIAutomationMultipleViewPattern = syscall.GUID{0x8d253c91, 0x1dc5, 0x4bb5, 
	[8]byte{0xb1, 0x8f, 0xad, 0xe1, 0x6f, 0xa4, 0x95, 0xe8}}

type IUIAutomationMultipleViewPatternInterface interface {
	IUnknownInterface
	GetViewName(view int32, name *BSTR) HRESULT
	SetCurrentView(view int32) HRESULT
	Get_CurrentCurrentView(retVal *int32) HRESULT
	GetCurrentSupportedViews(retVal **SAFEARRAY) HRESULT
	Get_CachedCurrentView(retVal *int32) HRESULT
	GetCachedSupportedViews(retVal **SAFEARRAY) HRESULT
}

type IUIAutomationMultipleViewPatternVtbl struct {
	IUnknownVtbl
	GetViewName uintptr
	SetCurrentView uintptr
	Get_CurrentCurrentView uintptr
	GetCurrentSupportedViews uintptr
	Get_CachedCurrentView uintptr
	GetCachedSupportedViews uintptr
}

type IUIAutomationMultipleViewPattern struct {
	IUnknown
}

func (this *IUIAutomationMultipleViewPattern) Vtbl() *IUIAutomationMultipleViewPatternVtbl {
	return (*IUIAutomationMultipleViewPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationMultipleViewPattern) GetViewName(view int32, name *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetViewName, uintptr(unsafe.Pointer(this)), uintptr(view), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IUIAutomationMultipleViewPattern) SetCurrentView(view int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetCurrentView, uintptr(unsafe.Pointer(this)), uintptr(view))
	return HRESULT(ret)
}

func (this *IUIAutomationMultipleViewPattern) Get_CurrentCurrentView(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationMultipleViewPattern) GetCurrentSupportedViews(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSupportedViews, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationMultipleViewPattern) Get_CachedCurrentView(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCurrentView, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationMultipleViewPattern) GetCachedSupportedViews(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedSupportedViews, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 71c284b3-c14d-4d14-981e-19751b0d756d
var IID_IUIAutomationObjectModelPattern = syscall.GUID{0x71c284b3, 0xc14d, 0x4d14, 
	[8]byte{0x98, 0x1e, 0x19, 0x75, 0x1b, 0x0d, 0x75, 0x6d}}

type IUIAutomationObjectModelPatternInterface interface {
	IUnknownInterface
	GetUnderlyingObjectModel(retVal **IUnknown) HRESULT
}

type IUIAutomationObjectModelPatternVtbl struct {
	IUnknownVtbl
	GetUnderlyingObjectModel uintptr
}

type IUIAutomationObjectModelPattern struct {
	IUnknown
}

func (this *IUIAutomationObjectModelPattern) Vtbl() *IUIAutomationObjectModelPatternVtbl {
	return (*IUIAutomationObjectModelPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationObjectModelPattern) GetUnderlyingObjectModel(retVal **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetUnderlyingObjectModel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 59213f4f-7346-49e5-b120-80555987a148
var IID_IUIAutomationRangeValuePattern = syscall.GUID{0x59213f4f, 0x7346, 0x49e5, 
	[8]byte{0xb1, 0x20, 0x80, 0x55, 0x59, 0x87, 0xa1, 0x48}}

type IUIAutomationRangeValuePatternInterface interface {
	IUnknownInterface
	SetValue(val float64) HRESULT
	Get_CurrentValue(retVal *float64) HRESULT
	Get_CurrentIsReadOnly(retVal *BOOL) HRESULT
	Get_CurrentMaximum(retVal *float64) HRESULT
	Get_CurrentMinimum(retVal *float64) HRESULT
	Get_CurrentLargeChange(retVal *float64) HRESULT
	Get_CurrentSmallChange(retVal *float64) HRESULT
	Get_CachedValue(retVal *float64) HRESULT
	Get_CachedIsReadOnly(retVal *BOOL) HRESULT
	Get_CachedMaximum(retVal *float64) HRESULT
	Get_CachedMinimum(retVal *float64) HRESULT
	Get_CachedLargeChange(retVal *float64) HRESULT
	Get_CachedSmallChange(retVal *float64) HRESULT
}

type IUIAutomationRangeValuePatternVtbl struct {
	IUnknownVtbl
	SetValue uintptr
	Get_CurrentValue uintptr
	Get_CurrentIsReadOnly uintptr
	Get_CurrentMaximum uintptr
	Get_CurrentMinimum uintptr
	Get_CurrentLargeChange uintptr
	Get_CurrentSmallChange uintptr
	Get_CachedValue uintptr
	Get_CachedIsReadOnly uintptr
	Get_CachedMaximum uintptr
	Get_CachedMinimum uintptr
	Get_CachedLargeChange uintptr
	Get_CachedSmallChange uintptr
}

type IUIAutomationRangeValuePattern struct {
	IUnknown
}

func (this *IUIAutomationRangeValuePattern) Vtbl() *IUIAutomationRangeValuePatternVtbl {
	return (*IUIAutomationRangeValuePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationRangeValuePattern) SetValue(val float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(val))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CurrentValue(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CurrentIsReadOnly(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CurrentMaximum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CurrentMinimum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CurrentLargeChange(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLargeChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CurrentSmallChange(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentSmallChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CachedValue(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CachedIsReadOnly(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CachedMaximum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CachedMinimum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CachedLargeChange(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLargeChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationRangeValuePattern) Get_CachedSmallChange(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedSmallChange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 88f4d42a-e881-459d-a77c-73bbbb7e02dc
var IID_IUIAutomationScrollPattern = syscall.GUID{0x88f4d42a, 0xe881, 0x459d, 
	[8]byte{0xa7, 0x7c, 0x73, 0xbb, 0xbb, 0x7e, 0x02, 0xdc}}

type IUIAutomationScrollPatternInterface interface {
	IUnknownInterface
	Scroll(horizontalAmount ScrollAmount, verticalAmount ScrollAmount) HRESULT
	SetScrollPercent(horizontalPercent float64, verticalPercent float64) HRESULT
	Get_CurrentHorizontalScrollPercent(retVal *float64) HRESULT
	Get_CurrentVerticalScrollPercent(retVal *float64) HRESULT
	Get_CurrentHorizontalViewSize(retVal *float64) HRESULT
	Get_CurrentVerticalViewSize(retVal *float64) HRESULT
	Get_CurrentHorizontallyScrollable(retVal *BOOL) HRESULT
	Get_CurrentVerticallyScrollable(retVal *BOOL) HRESULT
	Get_CachedHorizontalScrollPercent(retVal *float64) HRESULT
	Get_CachedVerticalScrollPercent(retVal *float64) HRESULT
	Get_CachedHorizontalViewSize(retVal *float64) HRESULT
	Get_CachedVerticalViewSize(retVal *float64) HRESULT
	Get_CachedHorizontallyScrollable(retVal *BOOL) HRESULT
	Get_CachedVerticallyScrollable(retVal *BOOL) HRESULT
}

type IUIAutomationScrollPatternVtbl struct {
	IUnknownVtbl
	Scroll uintptr
	SetScrollPercent uintptr
	Get_CurrentHorizontalScrollPercent uintptr
	Get_CurrentVerticalScrollPercent uintptr
	Get_CurrentHorizontalViewSize uintptr
	Get_CurrentVerticalViewSize uintptr
	Get_CurrentHorizontallyScrollable uintptr
	Get_CurrentVerticallyScrollable uintptr
	Get_CachedHorizontalScrollPercent uintptr
	Get_CachedVerticalScrollPercent uintptr
	Get_CachedHorizontalViewSize uintptr
	Get_CachedVerticalViewSize uintptr
	Get_CachedHorizontallyScrollable uintptr
	Get_CachedVerticallyScrollable uintptr
}

type IUIAutomationScrollPattern struct {
	IUnknown
}

func (this *IUIAutomationScrollPattern) Vtbl() *IUIAutomationScrollPatternVtbl {
	return (*IUIAutomationScrollPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationScrollPattern) Scroll(horizontalAmount ScrollAmount, verticalAmount ScrollAmount) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Scroll, uintptr(unsafe.Pointer(this)), uintptr(horizontalAmount), uintptr(verticalAmount))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) SetScrollPercent(horizontalPercent float64, verticalPercent float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(horizontalPercent), uintptr(verticalPercent))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CurrentHorizontalScrollPercent(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHorizontalScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CurrentVerticalScrollPercent(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentVerticalScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CurrentHorizontalViewSize(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHorizontalViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CurrentVerticalViewSize(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentVerticalViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CurrentHorizontallyScrollable(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHorizontallyScrollable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CurrentVerticallyScrollable(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentVerticallyScrollable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CachedHorizontalScrollPercent(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHorizontalScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CachedVerticalScrollPercent(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedVerticalScrollPercent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CachedHorizontalViewSize(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHorizontalViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CachedVerticalViewSize(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedVerticalViewSize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CachedHorizontallyScrollable(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHorizontallyScrollable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationScrollPattern) Get_CachedVerticallyScrollable(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedVerticallyScrollable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// b488300f-d015-4f19-9c29-bb595e3645ef
var IID_IUIAutomationScrollItemPattern = syscall.GUID{0xb488300f, 0xd015, 0x4f19, 
	[8]byte{0x9c, 0x29, 0xbb, 0x59, 0x5e, 0x36, 0x45, 0xef}}

type IUIAutomationScrollItemPatternInterface interface {
	IUnknownInterface
	ScrollIntoView() HRESULT
}

type IUIAutomationScrollItemPatternVtbl struct {
	IUnknownVtbl
	ScrollIntoView uintptr
}

type IUIAutomationScrollItemPattern struct {
	IUnknown
}

func (this *IUIAutomationScrollItemPattern) Vtbl() *IUIAutomationScrollItemPatternVtbl {
	return (*IUIAutomationScrollItemPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationScrollItemPattern) ScrollIntoView() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ScrollIntoView, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 5ed5202e-b2ac-47a6-b638-4b0bf140d78e
var IID_IUIAutomationSelectionPattern = syscall.GUID{0x5ed5202e, 0xb2ac, 0x47a6, 
	[8]byte{0xb6, 0x38, 0x4b, 0x0b, 0xf1, 0x40, 0xd7, 0x8e}}

type IUIAutomationSelectionPatternInterface interface {
	IUnknownInterface
	GetCurrentSelection(retVal **IUIAutomationElementArray) HRESULT
	Get_CurrentCanSelectMultiple(retVal *BOOL) HRESULT
	Get_CurrentIsSelectionRequired(retVal *BOOL) HRESULT
	GetCachedSelection(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedCanSelectMultiple(retVal *BOOL) HRESULT
	Get_CachedIsSelectionRequired(retVal *BOOL) HRESULT
}

type IUIAutomationSelectionPatternVtbl struct {
	IUnknownVtbl
	GetCurrentSelection uintptr
	Get_CurrentCanSelectMultiple uintptr
	Get_CurrentIsSelectionRequired uintptr
	GetCachedSelection uintptr
	Get_CachedCanSelectMultiple uintptr
	Get_CachedIsSelectionRequired uintptr
}

type IUIAutomationSelectionPattern struct {
	IUnknown
}

func (this *IUIAutomationSelectionPattern) Vtbl() *IUIAutomationSelectionPatternVtbl {
	return (*IUIAutomationSelectionPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationSelectionPattern) GetCurrentSelection(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern) Get_CurrentCanSelectMultiple(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanSelectMultiple, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern) Get_CurrentIsSelectionRequired(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsSelectionRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern) GetCachedSelection(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern) Get_CachedCanSelectMultiple(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanSelectMultiple, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern) Get_CachedIsSelectionRequired(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsSelectionRequired, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 0532bfae-c011-4e32-a343-6d642d798555
var IID_IUIAutomationSelectionPattern2 = syscall.GUID{0x0532bfae, 0xc011, 0x4e32, 
	[8]byte{0xa3, 0x43, 0x6d, 0x64, 0x2d, 0x79, 0x85, 0x55}}

type IUIAutomationSelectionPattern2Interface interface {
	IUIAutomationSelectionPatternInterface
	Get_CurrentFirstSelectedItem(retVal **IUIAutomationElement) HRESULT
	Get_CurrentLastSelectedItem(retVal **IUIAutomationElement) HRESULT
	Get_CurrentCurrentSelectedItem(retVal **IUIAutomationElement) HRESULT
	Get_CurrentItemCount(retVal *int32) HRESULT
	Get_CachedFirstSelectedItem(retVal **IUIAutomationElement) HRESULT
	Get_CachedLastSelectedItem(retVal **IUIAutomationElement) HRESULT
	Get_CachedCurrentSelectedItem(retVal **IUIAutomationElement) HRESULT
	Get_CachedItemCount(retVal *int32) HRESULT
}

type IUIAutomationSelectionPattern2Vtbl struct {
	IUIAutomationSelectionPatternVtbl
	Get_CurrentFirstSelectedItem uintptr
	Get_CurrentLastSelectedItem uintptr
	Get_CurrentCurrentSelectedItem uintptr
	Get_CurrentItemCount uintptr
	Get_CachedFirstSelectedItem uintptr
	Get_CachedLastSelectedItem uintptr
	Get_CachedCurrentSelectedItem uintptr
	Get_CachedItemCount uintptr
}

type IUIAutomationSelectionPattern2 struct {
	IUIAutomationSelectionPattern
}

func (this *IUIAutomationSelectionPattern2) Vtbl() *IUIAutomationSelectionPattern2Vtbl {
	return (*IUIAutomationSelectionPattern2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationSelectionPattern2) Get_CurrentFirstSelectedItem(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFirstSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CurrentLastSelectedItem(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLastSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CurrentCurrentSelectedItem(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCurrentSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CurrentItemCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentItemCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CachedFirstSelectedItem(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFirstSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CachedLastSelectedItem(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLastSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CachedCurrentSelectedItem(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCurrentSelectedItem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionPattern2) Get_CachedItemCount(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedItemCount, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// a8efa66a-0fda-421a-9194-38021f3578ea
var IID_IUIAutomationSelectionItemPattern = syscall.GUID{0xa8efa66a, 0x0fda, 0x421a, 
	[8]byte{0x91, 0x94, 0x38, 0x02, 0x1f, 0x35, 0x78, 0xea}}

type IUIAutomationSelectionItemPatternInterface interface {
	IUnknownInterface
	Select() HRESULT
	AddToSelection() HRESULT
	RemoveFromSelection() HRESULT
	Get_CurrentIsSelected(retVal *BOOL) HRESULT
	Get_CurrentSelectionContainer(retVal **IUIAutomationElement) HRESULT
	Get_CachedIsSelected(retVal *BOOL) HRESULT
	Get_CachedSelectionContainer(retVal **IUIAutomationElement) HRESULT
}

type IUIAutomationSelectionItemPatternVtbl struct {
	IUnknownVtbl
	Select uintptr
	AddToSelection uintptr
	RemoveFromSelection uintptr
	Get_CurrentIsSelected uintptr
	Get_CurrentSelectionContainer uintptr
	Get_CachedIsSelected uintptr
	Get_CachedSelectionContainer uintptr
}

type IUIAutomationSelectionItemPattern struct {
	IUnknown
}

func (this *IUIAutomationSelectionItemPattern) Vtbl() *IUIAutomationSelectionItemPatternVtbl {
	return (*IUIAutomationSelectionItemPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationSelectionItemPattern) Select() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Select, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionItemPattern) AddToSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddToSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionItemPattern) RemoveFromSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveFromSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionItemPattern) Get_CurrentIsSelected(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionItemPattern) Get_CurrentSelectionContainer(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentSelectionContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionItemPattern) Get_CachedIsSelected(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsSelected, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSelectionItemPattern) Get_CachedSelectionContainer(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedSelectionContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 2233be0b-afb7-448b-9fda-3b378aa5eae1
var IID_IUIAutomationSynchronizedInputPattern = syscall.GUID{0x2233be0b, 0xafb7, 0x448b, 
	[8]byte{0x9f, 0xda, 0x3b, 0x37, 0x8a, 0xa5, 0xea, 0xe1}}

type IUIAutomationSynchronizedInputPatternInterface interface {
	IUnknownInterface
	StartListening(inputType SynchronizedInputType) HRESULT
	Cancel() HRESULT
}

type IUIAutomationSynchronizedInputPatternVtbl struct {
	IUnknownVtbl
	StartListening uintptr
	Cancel uintptr
}

type IUIAutomationSynchronizedInputPattern struct {
	IUnknown
}

func (this *IUIAutomationSynchronizedInputPattern) Vtbl() *IUIAutomationSynchronizedInputPatternVtbl {
	return (*IUIAutomationSynchronizedInputPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationSynchronizedInputPattern) StartListening(inputType SynchronizedInputType) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().StartListening, uintptr(unsafe.Pointer(this)), uintptr(inputType))
	return HRESULT(ret)
}

func (this *IUIAutomationSynchronizedInputPattern) Cancel() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Cancel, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 620e691c-ea96-4710-a850-754b24ce2417
var IID_IUIAutomationTablePattern = syscall.GUID{0x620e691c, 0xea96, 0x4710, 
	[8]byte{0xa8, 0x50, 0x75, 0x4b, 0x24, 0xce, 0x24, 0x17}}

type IUIAutomationTablePatternInterface interface {
	IUnknownInterface
	GetCurrentRowHeaders(retVal **IUIAutomationElementArray) HRESULT
	GetCurrentColumnHeaders(retVal **IUIAutomationElementArray) HRESULT
	Get_CurrentRowOrColumnMajor(retVal *RowOrColumnMajor) HRESULT
	GetCachedRowHeaders(retVal **IUIAutomationElementArray) HRESULT
	GetCachedColumnHeaders(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedRowOrColumnMajor(retVal *RowOrColumnMajor) HRESULT
}

type IUIAutomationTablePatternVtbl struct {
	IUnknownVtbl
	GetCurrentRowHeaders uintptr
	GetCurrentColumnHeaders uintptr
	Get_CurrentRowOrColumnMajor uintptr
	GetCachedRowHeaders uintptr
	GetCachedColumnHeaders uintptr
	Get_CachedRowOrColumnMajor uintptr
}

type IUIAutomationTablePattern struct {
	IUnknown
}

func (this *IUIAutomationTablePattern) Vtbl() *IUIAutomationTablePatternVtbl {
	return (*IUIAutomationTablePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTablePattern) GetCurrentRowHeaders(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentRowHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTablePattern) GetCurrentColumnHeaders(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentColumnHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTablePattern) Get_CurrentRowOrColumnMajor(retVal *RowOrColumnMajor) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentRowOrColumnMajor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTablePattern) GetCachedRowHeaders(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedRowHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTablePattern) GetCachedColumnHeaders(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedColumnHeaders, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTablePattern) Get_CachedRowOrColumnMajor(retVal *RowOrColumnMajor) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedRowOrColumnMajor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 0b964eb3-ef2e-4464-9c79-61d61737a27e
var IID_IUIAutomationTableItemPattern = syscall.GUID{0x0b964eb3, 0xef2e, 0x4464, 
	[8]byte{0x9c, 0x79, 0x61, 0xd6, 0x17, 0x37, 0xa2, 0x7e}}

type IUIAutomationTableItemPatternInterface interface {
	IUnknownInterface
	GetCurrentRowHeaderItems(retVal **IUIAutomationElementArray) HRESULT
	GetCurrentColumnHeaderItems(retVal **IUIAutomationElementArray) HRESULT
	GetCachedRowHeaderItems(retVal **IUIAutomationElementArray) HRESULT
	GetCachedColumnHeaderItems(retVal **IUIAutomationElementArray) HRESULT
}

type IUIAutomationTableItemPatternVtbl struct {
	IUnknownVtbl
	GetCurrentRowHeaderItems uintptr
	GetCurrentColumnHeaderItems uintptr
	GetCachedRowHeaderItems uintptr
	GetCachedColumnHeaderItems uintptr
}

type IUIAutomationTableItemPattern struct {
	IUnknown
}

func (this *IUIAutomationTableItemPattern) Vtbl() *IUIAutomationTableItemPatternVtbl {
	return (*IUIAutomationTableItemPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTableItemPattern) GetCurrentRowHeaderItems(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentRowHeaderItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTableItemPattern) GetCurrentColumnHeaderItems(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentColumnHeaderItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTableItemPattern) GetCachedRowHeaderItems(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedRowHeaderItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTableItemPattern) GetCachedColumnHeaderItems(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedColumnHeaderItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 94cf8058-9b8d-4ab9-8bfd-4cd0a33c8c70
var IID_IUIAutomationTogglePattern = syscall.GUID{0x94cf8058, 0x9b8d, 0x4ab9, 
	[8]byte{0x8b, 0xfd, 0x4c, 0xd0, 0xa3, 0x3c, 0x8c, 0x70}}

type IUIAutomationTogglePatternInterface interface {
	IUnknownInterface
	Toggle() HRESULT
	Get_CurrentToggleState(retVal *ToggleState) HRESULT
	Get_CachedToggleState(retVal *ToggleState) HRESULT
}

type IUIAutomationTogglePatternVtbl struct {
	IUnknownVtbl
	Toggle uintptr
	Get_CurrentToggleState uintptr
	Get_CachedToggleState uintptr
}

type IUIAutomationTogglePattern struct {
	IUnknown
}

func (this *IUIAutomationTogglePattern) Vtbl() *IUIAutomationTogglePatternVtbl {
	return (*IUIAutomationTogglePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTogglePattern) Toggle() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Toggle, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationTogglePattern) Get_CurrentToggleState(retVal *ToggleState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentToggleState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTogglePattern) Get_CachedToggleState(retVal *ToggleState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedToggleState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// a9b55844-a55d-4ef0-926d-569c16ff89bb
var IID_IUIAutomationTransformPattern = syscall.GUID{0xa9b55844, 0xa55d, 0x4ef0, 
	[8]byte{0x92, 0x6d, 0x56, 0x9c, 0x16, 0xff, 0x89, 0xbb}}

type IUIAutomationTransformPatternInterface interface {
	IUnknownInterface
	Move(x float64, y float64) HRESULT
	Resize(width float64, height float64) HRESULT
	Rotate(degrees float64) HRESULT
	Get_CurrentCanMove(retVal *BOOL) HRESULT
	Get_CurrentCanResize(retVal *BOOL) HRESULT
	Get_CurrentCanRotate(retVal *BOOL) HRESULT
	Get_CachedCanMove(retVal *BOOL) HRESULT
	Get_CachedCanResize(retVal *BOOL) HRESULT
	Get_CachedCanRotate(retVal *BOOL) HRESULT
}

type IUIAutomationTransformPatternVtbl struct {
	IUnknownVtbl
	Move uintptr
	Resize uintptr
	Rotate uintptr
	Get_CurrentCanMove uintptr
	Get_CurrentCanResize uintptr
	Get_CurrentCanRotate uintptr
	Get_CachedCanMove uintptr
	Get_CachedCanResize uintptr
	Get_CachedCanRotate uintptr
}

type IUIAutomationTransformPattern struct {
	IUnknown
}

func (this *IUIAutomationTransformPattern) Vtbl() *IUIAutomationTransformPatternVtbl {
	return (*IUIAutomationTransformPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTransformPattern) Move(x float64, y float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Move, uintptr(unsafe.Pointer(this)), uintptr(x), uintptr(y))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Resize(width float64, height float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Resize, uintptr(unsafe.Pointer(this)), uintptr(width), uintptr(height))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Rotate(degrees float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Rotate, uintptr(unsafe.Pointer(this)), uintptr(degrees))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Get_CurrentCanMove(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanMove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Get_CurrentCanResize(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanResize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Get_CurrentCanRotate(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanRotate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Get_CachedCanMove(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanMove, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Get_CachedCanResize(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanResize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern) Get_CachedCanRotate(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanRotate, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// a94cd8b1-0844-4cd6-9d2d-640537ab39e9
var IID_IUIAutomationValuePattern = syscall.GUID{0xa94cd8b1, 0x0844, 0x4cd6, 
	[8]byte{0x9d, 0x2d, 0x64, 0x05, 0x37, 0xab, 0x39, 0xe9}}

type IUIAutomationValuePatternInterface interface {
	IUnknownInterface
	SetValue(val BSTR) HRESULT
	Get_CurrentValue(retVal *BSTR) HRESULT
	Get_CurrentIsReadOnly(retVal *BOOL) HRESULT
	Get_CachedValue(retVal *BSTR) HRESULT
	Get_CachedIsReadOnly(retVal *BOOL) HRESULT
}

type IUIAutomationValuePatternVtbl struct {
	IUnknownVtbl
	SetValue uintptr
	Get_CurrentValue uintptr
	Get_CurrentIsReadOnly uintptr
	Get_CachedValue uintptr
	Get_CachedIsReadOnly uintptr
}

type IUIAutomationValuePattern struct {
	IUnknown
}

func (this *IUIAutomationValuePattern) Vtbl() *IUIAutomationValuePatternVtbl {
	return (*IUIAutomationValuePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationValuePattern) SetValue(val BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(val)))
	return HRESULT(ret)
}

func (this *IUIAutomationValuePattern) Get_CurrentValue(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationValuePattern) Get_CurrentIsReadOnly(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationValuePattern) Get_CachedValue(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationValuePattern) Get_CachedIsReadOnly(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsReadOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 0faef453-9208-43ef-bbb2-3b485177864f
var IID_IUIAutomationWindowPattern = syscall.GUID{0x0faef453, 0x9208, 0x43ef, 
	[8]byte{0xbb, 0xb2, 0x3b, 0x48, 0x51, 0x77, 0x86, 0x4f}}

type IUIAutomationWindowPatternInterface interface {
	IUnknownInterface
	Close() HRESULT
	WaitForInputIdle(milliseconds int32, success *BOOL) HRESULT
	SetWindowVisualState(state WindowVisualState) HRESULT
	Get_CurrentCanMaximize(retVal *BOOL) HRESULT
	Get_CurrentCanMinimize(retVal *BOOL) HRESULT
	Get_CurrentIsModal(retVal *BOOL) HRESULT
	Get_CurrentIsTopmost(retVal *BOOL) HRESULT
	Get_CurrentWindowVisualState(retVal *WindowVisualState) HRESULT
	Get_CurrentWindowInteractionState(retVal *WindowInteractionState) HRESULT
	Get_CachedCanMaximize(retVal *BOOL) HRESULT
	Get_CachedCanMinimize(retVal *BOOL) HRESULT
	Get_CachedIsModal(retVal *BOOL) HRESULT
	Get_CachedIsTopmost(retVal *BOOL) HRESULT
	Get_CachedWindowVisualState(retVal *WindowVisualState) HRESULT
	Get_CachedWindowInteractionState(retVal *WindowInteractionState) HRESULT
}

type IUIAutomationWindowPatternVtbl struct {
	IUnknownVtbl
	Close uintptr
	WaitForInputIdle uintptr
	SetWindowVisualState uintptr
	Get_CurrentCanMaximize uintptr
	Get_CurrentCanMinimize uintptr
	Get_CurrentIsModal uintptr
	Get_CurrentIsTopmost uintptr
	Get_CurrentWindowVisualState uintptr
	Get_CurrentWindowInteractionState uintptr
	Get_CachedCanMaximize uintptr
	Get_CachedCanMinimize uintptr
	Get_CachedIsModal uintptr
	Get_CachedIsTopmost uintptr
	Get_CachedWindowVisualState uintptr
	Get_CachedWindowInteractionState uintptr
}

type IUIAutomationWindowPattern struct {
	IUnknown
}

func (this *IUIAutomationWindowPattern) Vtbl() *IUIAutomationWindowPatternVtbl {
	return (*IUIAutomationWindowPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationWindowPattern) Close() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Close, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) WaitForInputIdle(milliseconds int32, success *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().WaitForInputIdle, uintptr(unsafe.Pointer(this)), uintptr(milliseconds), uintptr(unsafe.Pointer(success)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) SetWindowVisualState(state WindowVisualState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetWindowVisualState, uintptr(unsafe.Pointer(this)), uintptr(state))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CurrentCanMaximize(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanMaximize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CurrentCanMinimize(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanMinimize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CurrentIsModal(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsModal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CurrentIsTopmost(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsTopmost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CurrentWindowVisualState(retVal *WindowVisualState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentWindowVisualState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CurrentWindowInteractionState(retVal *WindowInteractionState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentWindowInteractionState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CachedCanMaximize(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanMaximize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CachedCanMinimize(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanMinimize, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CachedIsModal(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsModal, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CachedIsTopmost(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsTopmost, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CachedWindowVisualState(retVal *WindowVisualState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedWindowVisualState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationWindowPattern) Get_CachedWindowInteractionState(retVal *WindowInteractionState) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedWindowInteractionState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// a543cc6a-f4ae-494b-8239-c814481187a8
var IID_IUIAutomationTextRange = syscall.GUID{0xa543cc6a, 0xf4ae, 0x494b, 
	[8]byte{0x82, 0x39, 0xc8, 0x14, 0x48, 0x11, 0x87, 0xa8}}

type IUIAutomationTextRangeInterface interface {
	IUnknownInterface
	Clone(clonedRange **IUIAutomationTextRange) HRESULT
	Compare(range_ *IUIAutomationTextRange, areSame *BOOL) HRESULT
	CompareEndpoints(srcEndPoint TextPatternRangeEndpoint, range_ *IUIAutomationTextRange, targetEndPoint TextPatternRangeEndpoint, compValue *int32) HRESULT
	ExpandToEnclosingUnit(textUnit TextUnit) HRESULT
	FindAttribute(attr int32, val VARIANT, backward BOOL, found **IUIAutomationTextRange) HRESULT
	FindText(text BSTR, backward BOOL, ignoreCase BOOL, found **IUIAutomationTextRange) HRESULT
	GetAttributeValue(attr int32, value *VARIANT) HRESULT
	GetBoundingRectangles(boundingRects **SAFEARRAY) HRESULT
	GetEnclosingElement(enclosingElement **IUIAutomationElement) HRESULT
	GetText(maxLength int32, text *BSTR) HRESULT
	Move(unit TextUnit, count int32, moved *int32) HRESULT
	MoveEndpointByUnit(endpoint TextPatternRangeEndpoint, unit TextUnit, count int32, moved *int32) HRESULT
	MoveEndpointByRange(srcEndPoint TextPatternRangeEndpoint, range_ *IUIAutomationTextRange, targetEndPoint TextPatternRangeEndpoint) HRESULT
	Select() HRESULT
	AddToSelection() HRESULT
	RemoveFromSelection() HRESULT
	ScrollIntoView(alignToTop BOOL) HRESULT
	GetChildren(children **IUIAutomationElementArray) HRESULT
}

type IUIAutomationTextRangeVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Compare uintptr
	CompareEndpoints uintptr
	ExpandToEnclosingUnit uintptr
	FindAttribute uintptr
	FindText uintptr
	GetAttributeValue uintptr
	GetBoundingRectangles uintptr
	GetEnclosingElement uintptr
	GetText uintptr
	Move uintptr
	MoveEndpointByUnit uintptr
	MoveEndpointByRange uintptr
	Select uintptr
	AddToSelection uintptr
	RemoveFromSelection uintptr
	ScrollIntoView uintptr
	GetChildren uintptr
}

type IUIAutomationTextRange struct {
	IUnknown
}

func (this *IUIAutomationTextRange) Vtbl() *IUIAutomationTextRangeVtbl {
	return (*IUIAutomationTextRangeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextRange) Clone(clonedRange **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Clone, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(clonedRange)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) Compare(range_ *IUIAutomationTextRange, areSame *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Compare, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(range_)), uintptr(unsafe.Pointer(areSame)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) CompareEndpoints(srcEndPoint TextPatternRangeEndpoint, range_ *IUIAutomationTextRange, targetEndPoint TextPatternRangeEndpoint, compValue *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CompareEndpoints, uintptr(unsafe.Pointer(this)), uintptr(srcEndPoint), uintptr(unsafe.Pointer(range_)), uintptr(targetEndPoint), uintptr(unsafe.Pointer(compValue)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) ExpandToEnclosingUnit(textUnit TextUnit) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ExpandToEnclosingUnit, uintptr(unsafe.Pointer(this)), uintptr(textUnit))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) FindAttribute(attr int32, val VARIANT, backward BOOL, found **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindAttribute, uintptr(unsafe.Pointer(this)), uintptr(attr), (uintptr)(unsafe.Pointer(&val)), uintptr(backward), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) FindText(text BSTR, backward BOOL, ignoreCase BOOL, found **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(text)), uintptr(backward), uintptr(ignoreCase), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) GetAttributeValue(attr int32, value *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAttributeValue, uintptr(unsafe.Pointer(this)), uintptr(attr), uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) GetBoundingRectangles(boundingRects **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetBoundingRectangles, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(boundingRects)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) GetEnclosingElement(enclosingElement **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEnclosingElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(enclosingElement)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) GetText(maxLength int32, text *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetText, uintptr(unsafe.Pointer(this)), uintptr(maxLength), uintptr(unsafe.Pointer(text)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) Move(unit TextUnit, count int32, moved *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Move, uintptr(unsafe.Pointer(this)), uintptr(unit), uintptr(count), uintptr(unsafe.Pointer(moved)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) MoveEndpointByUnit(endpoint TextPatternRangeEndpoint, unit TextUnit, count int32, moved *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().MoveEndpointByUnit, uintptr(unsafe.Pointer(this)), uintptr(endpoint), uintptr(unit), uintptr(count), uintptr(unsafe.Pointer(moved)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) MoveEndpointByRange(srcEndPoint TextPatternRangeEndpoint, range_ *IUIAutomationTextRange, targetEndPoint TextPatternRangeEndpoint) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().MoveEndpointByRange, uintptr(unsafe.Pointer(this)), uintptr(srcEndPoint), uintptr(unsafe.Pointer(range_)), uintptr(targetEndPoint))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) Select() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Select, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) AddToSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddToSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) RemoveFromSelection() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveFromSelection, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) ScrollIntoView(alignToTop BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ScrollIntoView, uintptr(unsafe.Pointer(this)), uintptr(alignToTop))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange) GetChildren(children **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildren, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(children)))
	return HRESULT(ret)
}

// bb9b40e0-5e04-46bd-9be0-4b601b9afad4
var IID_IUIAutomationTextRange2 = syscall.GUID{0xbb9b40e0, 0x5e04, 0x46bd, 
	[8]byte{0x9b, 0xe0, 0x4b, 0x60, 0x1b, 0x9a, 0xfa, 0xd4}}

type IUIAutomationTextRange2Interface interface {
	IUIAutomationTextRangeInterface
	ShowContextMenu() HRESULT
}

type IUIAutomationTextRange2Vtbl struct {
	IUIAutomationTextRangeVtbl
	ShowContextMenu uintptr
}

type IUIAutomationTextRange2 struct {
	IUIAutomationTextRange
}

func (this *IUIAutomationTextRange2) Vtbl() *IUIAutomationTextRange2Vtbl {
	return (*IUIAutomationTextRange2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextRange2) ShowContextMenu() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ShowContextMenu, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 6a315d69-5512-4c2e-85f0-53fce6dd4bc2
var IID_IUIAutomationTextRange3 = syscall.GUID{0x6a315d69, 0x5512, 0x4c2e, 
	[8]byte{0x85, 0xf0, 0x53, 0xfc, 0xe6, 0xdd, 0x4b, 0xc2}}

type IUIAutomationTextRange3Interface interface {
	IUIAutomationTextRange2Interface
	GetEnclosingElementBuildCache(cacheRequest *IUIAutomationCacheRequest, enclosingElement **IUIAutomationElement) HRESULT
	GetChildrenBuildCache(cacheRequest *IUIAutomationCacheRequest, children **IUIAutomationElementArray) HRESULT
	GetAttributeValues(attributeIds *int32, attributeIdCount int32, attributeValues **SAFEARRAY) HRESULT
}

type IUIAutomationTextRange3Vtbl struct {
	IUIAutomationTextRange2Vtbl
	GetEnclosingElementBuildCache uintptr
	GetChildrenBuildCache uintptr
	GetAttributeValues uintptr
}

type IUIAutomationTextRange3 struct {
	IUIAutomationTextRange2
}

func (this *IUIAutomationTextRange3) Vtbl() *IUIAutomationTextRange3Vtbl {
	return (*IUIAutomationTextRange3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextRange3) GetEnclosingElementBuildCache(cacheRequest *IUIAutomationCacheRequest, enclosingElement **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEnclosingElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(enclosingElement)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange3) GetChildrenBuildCache(cacheRequest *IUIAutomationCacheRequest, children **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetChildrenBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(children)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRange3) GetAttributeValues(attributeIds *int32, attributeIdCount int32, attributeValues **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetAttributeValues, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(attributeIds)), uintptr(attributeIdCount), uintptr(unsafe.Pointer(attributeValues)))
	return HRESULT(ret)
}

// ce4ae76a-e717-4c98-81ea-47371d028eb6
var IID_IUIAutomationTextRangeArray = syscall.GUID{0xce4ae76a, 0xe717, 0x4c98, 
	[8]byte{0x81, 0xea, 0x47, 0x37, 0x1d, 0x02, 0x8e, 0xb6}}

type IUIAutomationTextRangeArrayInterface interface {
	IUnknownInterface
	Get_Length(length *int32) HRESULT
	GetElement(index int32, element **IUIAutomationTextRange) HRESULT
}

type IUIAutomationTextRangeArrayVtbl struct {
	IUnknownVtbl
	Get_Length uintptr
	GetElement uintptr
}

type IUIAutomationTextRangeArray struct {
	IUnknown
}

func (this *IUIAutomationTextRangeArray) Vtbl() *IUIAutomationTextRangeArrayVtbl {
	return (*IUIAutomationTextRangeArrayVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextRangeArray) Get_Length(length *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Length, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(length)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextRangeArray) GetElement(index int32, element **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetElement, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

// 32eba289-3583-42c9-9c59-3b6d9a1e9b6a
var IID_IUIAutomationTextPattern = syscall.GUID{0x32eba289, 0x3583, 0x42c9, 
	[8]byte{0x9c, 0x59, 0x3b, 0x6d, 0x9a, 0x1e, 0x9b, 0x6a}}

type IUIAutomationTextPatternInterface interface {
	IUnknownInterface
	RangeFromPoint(pt POINT, range_ **IUIAutomationTextRange) HRESULT
	RangeFromChild(child *IUIAutomationElement, range_ **IUIAutomationTextRange) HRESULT
	GetSelection(ranges **IUIAutomationTextRangeArray) HRESULT
	GetVisibleRanges(ranges **IUIAutomationTextRangeArray) HRESULT
	Get_DocumentRange(range_ **IUIAutomationTextRange) HRESULT
	Get_SupportedTextSelection(supportedTextSelection *SupportedTextSelection) HRESULT
}

type IUIAutomationTextPatternVtbl struct {
	IUnknownVtbl
	RangeFromPoint uintptr
	RangeFromChild uintptr
	GetSelection uintptr
	GetVisibleRanges uintptr
	Get_DocumentRange uintptr
	Get_SupportedTextSelection uintptr
}

type IUIAutomationTextPattern struct {
	IUnknown
}

func (this *IUIAutomationTextPattern) Vtbl() *IUIAutomationTextPatternVtbl {
	return (*IUIAutomationTextPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextPattern) RangeFromPoint(pt POINT, range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RangeFromPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pt)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextPattern) RangeFromChild(child *IUIAutomationElement, range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RangeFromChild, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(child)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextPattern) GetSelection(ranges **IUIAutomationTextRangeArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ranges)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextPattern) GetVisibleRanges(ranges **IUIAutomationTextRangeArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetVisibleRanges, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ranges)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextPattern) Get_DocumentRange(range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_DocumentRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextPattern) Get_SupportedTextSelection(supportedTextSelection *SupportedTextSelection) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportedTextSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(supportedTextSelection)))
	return HRESULT(ret)
}

// 506a921a-fcc9-409f-b23b-37eb74106872
var IID_IUIAutomationTextPattern2 = syscall.GUID{0x506a921a, 0xfcc9, 0x409f, 
	[8]byte{0xb2, 0x3b, 0x37, 0xeb, 0x74, 0x10, 0x68, 0x72}}

type IUIAutomationTextPattern2Interface interface {
	IUIAutomationTextPatternInterface
	RangeFromAnnotation(annotation *IUIAutomationElement, range_ **IUIAutomationTextRange) HRESULT
	GetCaretRange(isActive *BOOL, range_ **IUIAutomationTextRange) HRESULT
}

type IUIAutomationTextPattern2Vtbl struct {
	IUIAutomationTextPatternVtbl
	RangeFromAnnotation uintptr
	GetCaretRange uintptr
}

type IUIAutomationTextPattern2 struct {
	IUIAutomationTextPattern
}

func (this *IUIAutomationTextPattern2) Vtbl() *IUIAutomationTextPattern2Vtbl {
	return (*IUIAutomationTextPattern2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextPattern2) RangeFromAnnotation(annotation *IUIAutomationElement, range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RangeFromAnnotation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(annotation)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextPattern2) GetCaretRange(isActive *BOOL, range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCaretRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isActive)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

// 17e21576-996c-4870-99d9-bff323380c06
var IID_IUIAutomationTextEditPattern = syscall.GUID{0x17e21576, 0x996c, 0x4870, 
	[8]byte{0x99, 0xd9, 0xbf, 0xf3, 0x23, 0x38, 0x0c, 0x06}}

type IUIAutomationTextEditPatternInterface interface {
	IUIAutomationTextPatternInterface
	GetActiveComposition(range_ **IUIAutomationTextRange) HRESULT
	GetConversionTarget(range_ **IUIAutomationTextRange) HRESULT
}

type IUIAutomationTextEditPatternVtbl struct {
	IUIAutomationTextPatternVtbl
	GetActiveComposition uintptr
	GetConversionTarget uintptr
}

type IUIAutomationTextEditPattern struct {
	IUIAutomationTextPattern
}

func (this *IUIAutomationTextEditPattern) Vtbl() *IUIAutomationTextEditPatternVtbl {
	return (*IUIAutomationTextEditPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextEditPattern) GetActiveComposition(range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetActiveComposition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextEditPattern) GetConversionTarget(range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetConversionTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

// 01ea217a-1766-47ed-a6cc-acf492854b1f
var IID_IUIAutomationCustomNavigationPattern = syscall.GUID{0x01ea217a, 0x1766, 0x47ed, 
	[8]byte{0xa6, 0xcc, 0xac, 0xf4, 0x92, 0x85, 0x4b, 0x1f}}

type IUIAutomationCustomNavigationPatternInterface interface {
	IUnknownInterface
	Navigate(direction NavigateDirection, pRetVal **IUIAutomationElement) HRESULT
}

type IUIAutomationCustomNavigationPatternVtbl struct {
	IUnknownVtbl
	Navigate uintptr
}

type IUIAutomationCustomNavigationPattern struct {
	IUnknown
}

func (this *IUIAutomationCustomNavigationPattern) Vtbl() *IUIAutomationCustomNavigationPatternVtbl {
	return (*IUIAutomationCustomNavigationPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationCustomNavigationPattern) Navigate(direction NavigateDirection, pRetVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Navigate, uintptr(unsafe.Pointer(this)), uintptr(direction), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

// f97933b0-8dae-4496-8997-5ba015fe0d82
var IID_IUIAutomationActiveTextPositionChangedEventHandler = syscall.GUID{0xf97933b0, 0x8dae, 0x4496, 
	[8]byte{0x89, 0x97, 0x5b, 0xa0, 0x15, 0xfe, 0x0d, 0x82}}

type IUIAutomationActiveTextPositionChangedEventHandlerInterface interface {
	IUnknownInterface
	HandleActiveTextPositionChangedEvent(sender *IUIAutomationElement, range_ *IUIAutomationTextRange) HRESULT
}

type IUIAutomationActiveTextPositionChangedEventHandlerVtbl struct {
	IUnknownVtbl
	HandleActiveTextPositionChangedEvent uintptr
}

type IUIAutomationActiveTextPositionChangedEventHandler struct {
	IUnknown
}

func (this *IUIAutomationActiveTextPositionChangedEventHandler) Vtbl() *IUIAutomationActiveTextPositionChangedEventHandlerVtbl {
	return (*IUIAutomationActiveTextPositionChangedEventHandlerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationActiveTextPositionChangedEventHandler) HandleActiveTextPositionChangedEvent(sender *IUIAutomationElement, range_ *IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().HandleActiveTextPositionChangedEvent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

// 828055ad-355b-4435-86d5-3b51c14a9b1b
var IID_IUIAutomationLegacyIAccessiblePattern = syscall.GUID{0x828055ad, 0x355b, 0x4435, 
	[8]byte{0x86, 0xd5, 0x3b, 0x51, 0xc1, 0x4a, 0x9b, 0x1b}}

type IUIAutomationLegacyIAccessiblePatternInterface interface {
	IUnknownInterface
	Select(flagsSelect int32) HRESULT
	DoDefaultAction() HRESULT
	SetValue(szValue PWSTR) HRESULT
	Get_CurrentChildId(pRetVal *int32) HRESULT
	Get_CurrentName(pszName *BSTR) HRESULT
	Get_CurrentValue(pszValue *BSTR) HRESULT
	Get_CurrentDescription(pszDescription *BSTR) HRESULT
	Get_CurrentRole(pdwRole *uint32) HRESULT
	Get_CurrentState(pdwState *uint32) HRESULT
	Get_CurrentHelp(pszHelp *BSTR) HRESULT
	Get_CurrentKeyboardShortcut(pszKeyboardShortcut *BSTR) HRESULT
	GetCurrentSelection(pvarSelectedChildren **IUIAutomationElementArray) HRESULT
	Get_CurrentDefaultAction(pszDefaultAction *BSTR) HRESULT
	Get_CachedChildId(pRetVal *int32) HRESULT
	Get_CachedName(pszName *BSTR) HRESULT
	Get_CachedValue(pszValue *BSTR) HRESULT
	Get_CachedDescription(pszDescription *BSTR) HRESULT
	Get_CachedRole(pdwRole *uint32) HRESULT
	Get_CachedState(pdwState *uint32) HRESULT
	Get_CachedHelp(pszHelp *BSTR) HRESULT
	Get_CachedKeyboardShortcut(pszKeyboardShortcut *BSTR) HRESULT
	GetCachedSelection(pvarSelectedChildren **IUIAutomationElementArray) HRESULT
	Get_CachedDefaultAction(pszDefaultAction *BSTR) HRESULT
	GetIAccessible(ppAccessible **IAccessible) HRESULT
}

type IUIAutomationLegacyIAccessiblePatternVtbl struct {
	IUnknownVtbl
	Select uintptr
	DoDefaultAction uintptr
	SetValue uintptr
	Get_CurrentChildId uintptr
	Get_CurrentName uintptr
	Get_CurrentValue uintptr
	Get_CurrentDescription uintptr
	Get_CurrentRole uintptr
	Get_CurrentState uintptr
	Get_CurrentHelp uintptr
	Get_CurrentKeyboardShortcut uintptr
	GetCurrentSelection uintptr
	Get_CurrentDefaultAction uintptr
	Get_CachedChildId uintptr
	Get_CachedName uintptr
	Get_CachedValue uintptr
	Get_CachedDescription uintptr
	Get_CachedRole uintptr
	Get_CachedState uintptr
	Get_CachedHelp uintptr
	Get_CachedKeyboardShortcut uintptr
	GetCachedSelection uintptr
	Get_CachedDefaultAction uintptr
	GetIAccessible uintptr
}

type IUIAutomationLegacyIAccessiblePattern struct {
	IUnknown
}

func (this *IUIAutomationLegacyIAccessiblePattern) Vtbl() *IUIAutomationLegacyIAccessiblePatternVtbl {
	return (*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationLegacyIAccessiblePattern) Select(flagsSelect int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Select, uintptr(unsafe.Pointer(this)), uintptr(flagsSelect))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) DoDefaultAction() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().DoDefaultAction, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) SetValue(szValue PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(szValue)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentChildId(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentChildId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentName(pszName *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszName)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentValue(pszValue *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszValue)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentDescription(pszDescription *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszDescription)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentRole(pdwRole *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentRole, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwRole)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentState(pdwState *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwState)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentHelp(pszHelp *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHelp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszHelp)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentKeyboardShortcut(pszKeyboardShortcut *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentKeyboardShortcut, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszKeyboardShortcut)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) GetCurrentSelection(pvarSelectedChildren **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pvarSelectedChildren)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CurrentDefaultAction(pszDefaultAction *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDefaultAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszDefaultAction)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedChildId(pRetVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedChildId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pRetVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedName(pszName *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszName)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedValue(pszValue *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszValue)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedDescription(pszDescription *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszDescription)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedRole(pdwRole *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedRole, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwRole)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedState(pdwState *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pdwState)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedHelp(pszHelp *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHelp, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszHelp)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedKeyboardShortcut(pszKeyboardShortcut *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedKeyboardShortcut, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszKeyboardShortcut)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) GetCachedSelection(pvarSelectedChildren **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedSelection, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pvarSelectedChildren)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) Get_CachedDefaultAction(pszDefaultAction *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDefaultAction, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pszDefaultAction)))
	return HRESULT(ret)
}

func (this *IUIAutomationLegacyIAccessiblePattern) GetIAccessible(ppAccessible **IAccessible) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetIAccessible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(ppAccessible)))
	return HRESULT(ret)
}

// c690fdb2-27a8-423c-812d-429773c9084e
var IID_IUIAutomationItemContainerPattern = syscall.GUID{0xc690fdb2, 0x27a8, 0x423c, 
	[8]byte{0x81, 0x2d, 0x42, 0x97, 0x73, 0xc9, 0x08, 0x4e}}

type IUIAutomationItemContainerPatternInterface interface {
	IUnknownInterface
	FindItemByProperty(pStartAfter *IUIAutomationElement, propertyId int32, value VARIANT, pFound **IUIAutomationElement) HRESULT
}

type IUIAutomationItemContainerPatternVtbl struct {
	IUnknownVtbl
	FindItemByProperty uintptr
}

type IUIAutomationItemContainerPattern struct {
	IUnknown
}

func (this *IUIAutomationItemContainerPattern) Vtbl() *IUIAutomationItemContainerPatternVtbl {
	return (*IUIAutomationItemContainerPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationItemContainerPattern) FindItemByProperty(pStartAfter *IUIAutomationElement, propertyId int32, value VARIANT, pFound **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindItemByProperty, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pStartAfter)), uintptr(propertyId), (uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(pFound)))
	return HRESULT(ret)
}

// 6ba3d7a6-04cf-4f11-8793-a8d1cde9969f
var IID_IUIAutomationVirtualizedItemPattern = syscall.GUID{0x6ba3d7a6, 0x04cf, 0x4f11, 
	[8]byte{0x87, 0x93, 0xa8, 0xd1, 0xcd, 0xe9, 0x96, 0x9f}}

type IUIAutomationVirtualizedItemPatternInterface interface {
	IUnknownInterface
	Realize() HRESULT
}

type IUIAutomationVirtualizedItemPatternVtbl struct {
	IUnknownVtbl
	Realize uintptr
}

type IUIAutomationVirtualizedItemPattern struct {
	IUnknown
}

func (this *IUIAutomationVirtualizedItemPattern) Vtbl() *IUIAutomationVirtualizedItemPatternVtbl {
	return (*IUIAutomationVirtualizedItemPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationVirtualizedItemPattern) Realize() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Realize, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// 9a175b21-339e-41b1-8e8b-623f6b681098
var IID_IUIAutomationAnnotationPattern = syscall.GUID{0x9a175b21, 0x339e, 0x41b1, 
	[8]byte{0x8e, 0x8b, 0x62, 0x3f, 0x6b, 0x68, 0x10, 0x98}}

type IUIAutomationAnnotationPatternInterface interface {
	IUnknownInterface
	Get_CurrentAnnotationTypeId(retVal *int32) HRESULT
	Get_CurrentAnnotationTypeName(retVal *BSTR) HRESULT
	Get_CurrentAuthor(retVal *BSTR) HRESULT
	Get_CurrentDateTime(retVal *BSTR) HRESULT
	Get_CurrentTarget(retVal **IUIAutomationElement) HRESULT
	Get_CachedAnnotationTypeId(retVal *int32) HRESULT
	Get_CachedAnnotationTypeName(retVal *BSTR) HRESULT
	Get_CachedAuthor(retVal *BSTR) HRESULT
	Get_CachedDateTime(retVal *BSTR) HRESULT
	Get_CachedTarget(retVal **IUIAutomationElement) HRESULT
}

type IUIAutomationAnnotationPatternVtbl struct {
	IUnknownVtbl
	Get_CurrentAnnotationTypeId uintptr
	Get_CurrentAnnotationTypeName uintptr
	Get_CurrentAuthor uintptr
	Get_CurrentDateTime uintptr
	Get_CurrentTarget uintptr
	Get_CachedAnnotationTypeId uintptr
	Get_CachedAnnotationTypeName uintptr
	Get_CachedAuthor uintptr
	Get_CachedDateTime uintptr
	Get_CachedTarget uintptr
}

type IUIAutomationAnnotationPattern struct {
	IUnknown
}

func (this *IUIAutomationAnnotationPattern) Vtbl() *IUIAutomationAnnotationPatternVtbl {
	return (*IUIAutomationAnnotationPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationAnnotationPattern) Get_CurrentAnnotationTypeId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAnnotationTypeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CurrentAnnotationTypeName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAnnotationTypeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CurrentAuthor(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAuthor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CurrentDateTime(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CurrentTarget(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CachedAnnotationTypeId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAnnotationTypeId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CachedAnnotationTypeName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAnnotationTypeName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CachedAuthor(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAuthor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CachedDateTime(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDateTime, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationAnnotationPattern) Get_CachedTarget(retVal **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedTarget, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 85b5f0a2-bd79-484a-ad2b-388c9838d5fb
var IID_IUIAutomationStylesPattern = syscall.GUID{0x85b5f0a2, 0xbd79, 0x484a, 
	[8]byte{0xad, 0x2b, 0x38, 0x8c, 0x98, 0x38, 0xd5, 0xfb}}

type IUIAutomationStylesPatternInterface interface {
	IUnknownInterface
	Get_CurrentStyleId(retVal *int32) HRESULT
	Get_CurrentStyleName(retVal *BSTR) HRESULT
	Get_CurrentFillColor(retVal *int32) HRESULT
	Get_CurrentFillPatternStyle(retVal *BSTR) HRESULT
	Get_CurrentShape(retVal *BSTR) HRESULT
	Get_CurrentFillPatternColor(retVal *int32) HRESULT
	Get_CurrentExtendedProperties(retVal *BSTR) HRESULT
	GetCurrentExtendedPropertiesAsArray(propertyArray **ExtendedProperty, propertyCount *int32) HRESULT
	Get_CachedStyleId(retVal *int32) HRESULT
	Get_CachedStyleName(retVal *BSTR) HRESULT
	Get_CachedFillColor(retVal *int32) HRESULT
	Get_CachedFillPatternStyle(retVal *BSTR) HRESULT
	Get_CachedShape(retVal *BSTR) HRESULT
	Get_CachedFillPatternColor(retVal *int32) HRESULT
	Get_CachedExtendedProperties(retVal *BSTR) HRESULT
	GetCachedExtendedPropertiesAsArray(propertyArray **ExtendedProperty, propertyCount *int32) HRESULT
}

type IUIAutomationStylesPatternVtbl struct {
	IUnknownVtbl
	Get_CurrentStyleId uintptr
	Get_CurrentStyleName uintptr
	Get_CurrentFillColor uintptr
	Get_CurrentFillPatternStyle uintptr
	Get_CurrentShape uintptr
	Get_CurrentFillPatternColor uintptr
	Get_CurrentExtendedProperties uintptr
	GetCurrentExtendedPropertiesAsArray uintptr
	Get_CachedStyleId uintptr
	Get_CachedStyleName uintptr
	Get_CachedFillColor uintptr
	Get_CachedFillPatternStyle uintptr
	Get_CachedShape uintptr
	Get_CachedFillPatternColor uintptr
	Get_CachedExtendedProperties uintptr
	GetCachedExtendedPropertiesAsArray uintptr
}

type IUIAutomationStylesPattern struct {
	IUnknown
}

func (this *IUIAutomationStylesPattern) Vtbl() *IUIAutomationStylesPatternVtbl {
	return (*IUIAutomationStylesPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationStylesPattern) Get_CurrentStyleId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentStyleId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CurrentStyleName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentStyleName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CurrentFillColor(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFillColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CurrentFillPatternStyle(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFillPatternStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CurrentShape(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentShape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CurrentFillPatternColor(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFillPatternColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CurrentExtendedProperties(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentExtendedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) GetCurrentExtendedPropertiesAsArray(propertyArray **ExtendedProperty, propertyCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentExtendedPropertiesAsArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertyArray)), uintptr(unsafe.Pointer(propertyCount)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedStyleId(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedStyleId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedStyleName(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedStyleName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedFillColor(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFillColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedFillPatternStyle(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFillPatternStyle, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedShape(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedShape, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedFillPatternColor(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFillPatternColor, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) Get_CachedExtendedProperties(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedExtendedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationStylesPattern) GetCachedExtendedPropertiesAsArray(propertyArray **ExtendedProperty, propertyCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedExtendedPropertiesAsArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(propertyArray)), uintptr(unsafe.Pointer(propertyCount)))
	return HRESULT(ret)
}

// 7517a7c8-faae-4de9-9f08-29b91e8595c1
var IID_IUIAutomationSpreadsheetPattern = syscall.GUID{0x7517a7c8, 0xfaae, 0x4de9, 
	[8]byte{0x9f, 0x08, 0x29, 0xb9, 0x1e, 0x85, 0x95, 0xc1}}

type IUIAutomationSpreadsheetPatternInterface interface {
	IUnknownInterface
	GetItemByName(name BSTR, element **IUIAutomationElement) HRESULT
}

type IUIAutomationSpreadsheetPatternVtbl struct {
	IUnknownVtbl
	GetItemByName uintptr
}

type IUIAutomationSpreadsheetPattern struct {
	IUnknown
}

func (this *IUIAutomationSpreadsheetPattern) Vtbl() *IUIAutomationSpreadsheetPatternVtbl {
	return (*IUIAutomationSpreadsheetPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationSpreadsheetPattern) GetItemByName(name BSTR, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetItemByName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

// 7d4fb86c-8d34-40e1-8e83-62c15204e335
var IID_IUIAutomationSpreadsheetItemPattern = syscall.GUID{0x7d4fb86c, 0x8d34, 0x40e1, 
	[8]byte{0x8e, 0x83, 0x62, 0xc1, 0x52, 0x04, 0xe3, 0x35}}

type IUIAutomationSpreadsheetItemPatternInterface interface {
	IUnknownInterface
	Get_CurrentFormula(retVal *BSTR) HRESULT
	GetCurrentAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT
	GetCurrentAnnotationTypes(retVal **SAFEARRAY) HRESULT
	Get_CachedFormula(retVal *BSTR) HRESULT
	GetCachedAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT
	GetCachedAnnotationTypes(retVal **SAFEARRAY) HRESULT
}

type IUIAutomationSpreadsheetItemPatternVtbl struct {
	IUnknownVtbl
	Get_CurrentFormula uintptr
	GetCurrentAnnotationObjects uintptr
	GetCurrentAnnotationTypes uintptr
	Get_CachedFormula uintptr
	GetCachedAnnotationObjects uintptr
	GetCachedAnnotationTypes uintptr
}

type IUIAutomationSpreadsheetItemPattern struct {
	IUnknown
}

func (this *IUIAutomationSpreadsheetItemPattern) Vtbl() *IUIAutomationSpreadsheetItemPatternVtbl {
	return (*IUIAutomationSpreadsheetItemPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationSpreadsheetItemPattern) Get_CurrentFormula(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFormula, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSpreadsheetItemPattern) GetCurrentAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentAnnotationObjects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSpreadsheetItemPattern) GetCurrentAnnotationTypes(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentAnnotationTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSpreadsheetItemPattern) Get_CachedFormula(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFormula, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSpreadsheetItemPattern) GetCachedAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedAnnotationObjects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationSpreadsheetItemPattern) GetCachedAnnotationTypes(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedAnnotationTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 6d74d017-6ecb-4381-b38b-3c17a48ff1c2
var IID_IUIAutomationTransformPattern2 = syscall.GUID{0x6d74d017, 0x6ecb, 0x4381, 
	[8]byte{0xb3, 0x8b, 0x3c, 0x17, 0xa4, 0x8f, 0xf1, 0xc2}}

type IUIAutomationTransformPattern2Interface interface {
	IUIAutomationTransformPatternInterface
	Zoom(zoomValue float64) HRESULT
	ZoomByUnit(zoomUnit ZoomUnit) HRESULT
	Get_CurrentCanZoom(retVal *BOOL) HRESULT
	Get_CachedCanZoom(retVal *BOOL) HRESULT
	Get_CurrentZoomLevel(retVal *float64) HRESULT
	Get_CachedZoomLevel(retVal *float64) HRESULT
	Get_CurrentZoomMinimum(retVal *float64) HRESULT
	Get_CachedZoomMinimum(retVal *float64) HRESULT
	Get_CurrentZoomMaximum(retVal *float64) HRESULT
	Get_CachedZoomMaximum(retVal *float64) HRESULT
}

type IUIAutomationTransformPattern2Vtbl struct {
	IUIAutomationTransformPatternVtbl
	Zoom uintptr
	ZoomByUnit uintptr
	Get_CurrentCanZoom uintptr
	Get_CachedCanZoom uintptr
	Get_CurrentZoomLevel uintptr
	Get_CachedZoomLevel uintptr
	Get_CurrentZoomMinimum uintptr
	Get_CachedZoomMinimum uintptr
	Get_CurrentZoomMaximum uintptr
	Get_CachedZoomMaximum uintptr
}

type IUIAutomationTransformPattern2 struct {
	IUIAutomationTransformPattern
}

func (this *IUIAutomationTransformPattern2) Vtbl() *IUIAutomationTransformPattern2Vtbl {
	return (*IUIAutomationTransformPattern2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTransformPattern2) Zoom(zoomValue float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Zoom, uintptr(unsafe.Pointer(this)), uintptr(zoomValue))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) ZoomByUnit(zoomUnit ZoomUnit) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ZoomByUnit, uintptr(unsafe.Pointer(this)), uintptr(zoomUnit))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CurrentCanZoom(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentCanZoom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CachedCanZoom(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedCanZoom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CurrentZoomLevel(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentZoomLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CachedZoomLevel(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedZoomLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CurrentZoomMinimum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentZoomMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CachedZoomMinimum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedZoomMinimum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CurrentZoomMaximum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentZoomMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationTransformPattern2) Get_CachedZoomMaximum(retVal *float64) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedZoomMaximum, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 6552b038-ae05-40c8-abfd-aa08352aab86
var IID_IUIAutomationTextChildPattern = syscall.GUID{0x6552b038, 0xae05, 0x40c8, 
	[8]byte{0xab, 0xfd, 0xaa, 0x08, 0x35, 0x2a, 0xab, 0x86}}

type IUIAutomationTextChildPatternInterface interface {
	IUnknownInterface
	Get_TextContainer(container **IUIAutomationElement) HRESULT
	Get_TextRange(range_ **IUIAutomationTextRange) HRESULT
}

type IUIAutomationTextChildPatternVtbl struct {
	IUnknownVtbl
	Get_TextContainer uintptr
	Get_TextRange uintptr
}

type IUIAutomationTextChildPattern struct {
	IUnknown
}

func (this *IUIAutomationTextChildPattern) Vtbl() *IUIAutomationTextChildPatternVtbl {
	return (*IUIAutomationTextChildPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationTextChildPattern) Get_TextContainer(container **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TextContainer, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(container)))
	return HRESULT(ret)
}

func (this *IUIAutomationTextChildPattern) Get_TextRange(range_ **IUIAutomationTextRange) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TextRange, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(range_)))
	return HRESULT(ret)
}

// 1dc7b570-1f54-4bad-bcda-d36a722fb7bd
var IID_IUIAutomationDragPattern = syscall.GUID{0x1dc7b570, 0x1f54, 0x4bad, 
	[8]byte{0xbc, 0xda, 0xd3, 0x6a, 0x72, 0x2f, 0xb7, 0xbd}}

type IUIAutomationDragPatternInterface interface {
	IUnknownInterface
	Get_CurrentIsGrabbed(retVal *BOOL) HRESULT
	Get_CachedIsGrabbed(retVal *BOOL) HRESULT
	Get_CurrentDropEffect(retVal *BSTR) HRESULT
	Get_CachedDropEffect(retVal *BSTR) HRESULT
	Get_CurrentDropEffects(retVal **SAFEARRAY) HRESULT
	Get_CachedDropEffects(retVal **SAFEARRAY) HRESULT
	GetCurrentGrabbedItems(retVal **IUIAutomationElementArray) HRESULT
	GetCachedGrabbedItems(retVal **IUIAutomationElementArray) HRESULT
}

type IUIAutomationDragPatternVtbl struct {
	IUnknownVtbl
	Get_CurrentIsGrabbed uintptr
	Get_CachedIsGrabbed uintptr
	Get_CurrentDropEffect uintptr
	Get_CachedDropEffect uintptr
	Get_CurrentDropEffects uintptr
	Get_CachedDropEffects uintptr
	GetCurrentGrabbedItems uintptr
	GetCachedGrabbedItems uintptr
}

type IUIAutomationDragPattern struct {
	IUnknown
}

func (this *IUIAutomationDragPattern) Vtbl() *IUIAutomationDragPatternVtbl {
	return (*IUIAutomationDragPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationDragPattern) Get_CurrentIsGrabbed(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsGrabbed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) Get_CachedIsGrabbed(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsGrabbed, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) Get_CurrentDropEffect(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDropEffect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) Get_CachedDropEffect(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDropEffect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) Get_CurrentDropEffects(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDropEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) Get_CachedDropEffects(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDropEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) GetCurrentGrabbedItems(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentGrabbedItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDragPattern) GetCachedGrabbedItems(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCachedGrabbedItems, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 69a095f7-eee4-430e-a46b-fb73b1ae39a5
var IID_IUIAutomationDropTargetPattern = syscall.GUID{0x69a095f7, 0xeee4, 0x430e, 
	[8]byte{0xa4, 0x6b, 0xfb, 0x73, 0xb1, 0xae, 0x39, 0xa5}}

type IUIAutomationDropTargetPatternInterface interface {
	IUnknownInterface
	Get_CurrentDropTargetEffect(retVal *BSTR) HRESULT
	Get_CachedDropTargetEffect(retVal *BSTR) HRESULT
	Get_CurrentDropTargetEffects(retVal **SAFEARRAY) HRESULT
	Get_CachedDropTargetEffects(retVal **SAFEARRAY) HRESULT
}

type IUIAutomationDropTargetPatternVtbl struct {
	IUnknownVtbl
	Get_CurrentDropTargetEffect uintptr
	Get_CachedDropTargetEffect uintptr
	Get_CurrentDropTargetEffects uintptr
	Get_CachedDropTargetEffects uintptr
}

type IUIAutomationDropTargetPattern struct {
	IUnknown
}

func (this *IUIAutomationDropTargetPattern) Vtbl() *IUIAutomationDropTargetPatternVtbl {
	return (*IUIAutomationDropTargetPatternVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationDropTargetPattern) Get_CurrentDropTargetEffect(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDropTargetEffect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDropTargetPattern) Get_CachedDropTargetEffect(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDropTargetEffect, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDropTargetPattern) Get_CurrentDropTargetEffects(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentDropTargetEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationDropTargetPattern) Get_CachedDropTargetEffects(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedDropTargetEffects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 6749c683-f70d-4487-a698-5f79d55290d6
var IID_IUIAutomationElement2 = syscall.GUID{0x6749c683, 0xf70d, 0x4487, 
	[8]byte{0xa6, 0x98, 0x5f, 0x79, 0xd5, 0x52, 0x90, 0xd6}}

type IUIAutomationElement2Interface interface {
	IUIAutomationElementInterface
	Get_CurrentOptimizeForVisualContent(retVal *BOOL) HRESULT
	Get_CachedOptimizeForVisualContent(retVal *BOOL) HRESULT
	Get_CurrentLiveSetting(retVal *LiveSetting) HRESULT
	Get_CachedLiveSetting(retVal *LiveSetting) HRESULT
	Get_CurrentFlowsFrom(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedFlowsFrom(retVal **IUIAutomationElementArray) HRESULT
}

type IUIAutomationElement2Vtbl struct {
	IUIAutomationElementVtbl
	Get_CurrentOptimizeForVisualContent uintptr
	Get_CachedOptimizeForVisualContent uintptr
	Get_CurrentLiveSetting uintptr
	Get_CachedLiveSetting uintptr
	Get_CurrentFlowsFrom uintptr
	Get_CachedFlowsFrom uintptr
}

type IUIAutomationElement2 struct {
	IUIAutomationElement
}

func (this *IUIAutomationElement2) Vtbl() *IUIAutomationElement2Vtbl {
	return (*IUIAutomationElement2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement2) Get_CurrentOptimizeForVisualContent(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentOptimizeForVisualContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement2) Get_CachedOptimizeForVisualContent(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedOptimizeForVisualContent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement2) Get_CurrentLiveSetting(retVal *LiveSetting) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLiveSetting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement2) Get_CachedLiveSetting(retVal *LiveSetting) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLiveSetting, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement2) Get_CurrentFlowsFrom(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFlowsFrom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement2) Get_CachedFlowsFrom(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFlowsFrom, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 8471df34-aee0-4a01-a7de-7db9af12c296
var IID_IUIAutomationElement3 = syscall.GUID{0x8471df34, 0xaee0, 0x4a01, 
	[8]byte{0xa7, 0xde, 0x7d, 0xb9, 0xaf, 0x12, 0xc2, 0x96}}

type IUIAutomationElement3Interface interface {
	IUIAutomationElement2Interface
	ShowContextMenu() HRESULT
	Get_CurrentIsPeripheral(retVal *BOOL) HRESULT
	Get_CachedIsPeripheral(retVal *BOOL) HRESULT
}

type IUIAutomationElement3Vtbl struct {
	IUIAutomationElement2Vtbl
	ShowContextMenu uintptr
	Get_CurrentIsPeripheral uintptr
	Get_CachedIsPeripheral uintptr
}

type IUIAutomationElement3 struct {
	IUIAutomationElement2
}

func (this *IUIAutomationElement3) Vtbl() *IUIAutomationElement3Vtbl {
	return (*IUIAutomationElement3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement3) ShowContextMenu() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ShowContextMenu, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement3) Get_CurrentIsPeripheral(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsPeripheral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement3) Get_CachedIsPeripheral(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsPeripheral, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 3b6e233c-52fb-4063-a4c9-77c075c2a06b
var IID_IUIAutomationElement4 = syscall.GUID{0x3b6e233c, 0x52fb, 0x4063, 
	[8]byte{0xa4, 0xc9, 0x77, 0xc0, 0x75, 0xc2, 0xa0, 0x6b}}

type IUIAutomationElement4Interface interface {
	IUIAutomationElement3Interface
	Get_CurrentPositionInSet(retVal *int32) HRESULT
	Get_CurrentSizeOfSet(retVal *int32) HRESULT
	Get_CurrentLevel(retVal *int32) HRESULT
	Get_CurrentAnnotationTypes(retVal **SAFEARRAY) HRESULT
	Get_CurrentAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT
	Get_CachedPositionInSet(retVal *int32) HRESULT
	Get_CachedSizeOfSet(retVal *int32) HRESULT
	Get_CachedLevel(retVal *int32) HRESULT
	Get_CachedAnnotationTypes(retVal **SAFEARRAY) HRESULT
	Get_CachedAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT
}

type IUIAutomationElement4Vtbl struct {
	IUIAutomationElement3Vtbl
	Get_CurrentPositionInSet uintptr
	Get_CurrentSizeOfSet uintptr
	Get_CurrentLevel uintptr
	Get_CurrentAnnotationTypes uintptr
	Get_CurrentAnnotationObjects uintptr
	Get_CachedPositionInSet uintptr
	Get_CachedSizeOfSet uintptr
	Get_CachedLevel uintptr
	Get_CachedAnnotationTypes uintptr
	Get_CachedAnnotationObjects uintptr
}

type IUIAutomationElement4 struct {
	IUIAutomationElement3
}

func (this *IUIAutomationElement4) Vtbl() *IUIAutomationElement4Vtbl {
	return (*IUIAutomationElement4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement4) Get_CurrentPositionInSet(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentPositionInSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CurrentSizeOfSet(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentSizeOfSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CurrentLevel(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CurrentAnnotationTypes(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAnnotationTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CurrentAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentAnnotationObjects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CachedPositionInSet(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedPositionInSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CachedSizeOfSet(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedSizeOfSet, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CachedLevel(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CachedAnnotationTypes(retVal **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAnnotationTypes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement4) Get_CachedAnnotationObjects(retVal **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedAnnotationObjects, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 98141c1d-0d0e-4175-bbe2-6bff455842a7
var IID_IUIAutomationElement5 = syscall.GUID{0x98141c1d, 0x0d0e, 0x4175, 
	[8]byte{0xbb, 0xe2, 0x6b, 0xff, 0x45, 0x58, 0x42, 0xa7}}

type IUIAutomationElement5Interface interface {
	IUIAutomationElement4Interface
	Get_CurrentLandmarkType(retVal *int32) HRESULT
	Get_CurrentLocalizedLandmarkType(retVal *BSTR) HRESULT
	Get_CachedLandmarkType(retVal *int32) HRESULT
	Get_CachedLocalizedLandmarkType(retVal *BSTR) HRESULT
}

type IUIAutomationElement5Vtbl struct {
	IUIAutomationElement4Vtbl
	Get_CurrentLandmarkType uintptr
	Get_CurrentLocalizedLandmarkType uintptr
	Get_CachedLandmarkType uintptr
	Get_CachedLocalizedLandmarkType uintptr
}

type IUIAutomationElement5 struct {
	IUIAutomationElement4
}

func (this *IUIAutomationElement5) Vtbl() *IUIAutomationElement5Vtbl {
	return (*IUIAutomationElement5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement5) Get_CurrentLandmarkType(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLandmarkType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement5) Get_CurrentLocalizedLandmarkType(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentLocalizedLandmarkType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement5) Get_CachedLandmarkType(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLandmarkType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement5) Get_CachedLocalizedLandmarkType(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedLocalizedLandmarkType, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 4780d450-8bca-4977-afa5-a4a517f555e3
var IID_IUIAutomationElement6 = syscall.GUID{0x4780d450, 0x8bca, 0x4977, 
	[8]byte{0xaf, 0xa5, 0xa4, 0xa5, 0x17, 0xf5, 0x55, 0xe3}}

type IUIAutomationElement6Interface interface {
	IUIAutomationElement5Interface
	Get_CurrentFullDescription(retVal *BSTR) HRESULT
	Get_CachedFullDescription(retVal *BSTR) HRESULT
}

type IUIAutomationElement6Vtbl struct {
	IUIAutomationElement5Vtbl
	Get_CurrentFullDescription uintptr
	Get_CachedFullDescription uintptr
}

type IUIAutomationElement6 struct {
	IUIAutomationElement5
}

func (this *IUIAutomationElement6) Vtbl() *IUIAutomationElement6Vtbl {
	return (*IUIAutomationElement6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement6) Get_CurrentFullDescription(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentFullDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement6) Get_CachedFullDescription(retVal *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedFullDescription, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 204e8572-cfc3-4c11-b0c8-7da7420750b7
var IID_IUIAutomationElement7 = syscall.GUID{0x204e8572, 0xcfc3, 0x4c11, 
	[8]byte{0xb0, 0xc8, 0x7d, 0xa7, 0x42, 0x07, 0x50, 0xb7}}

type IUIAutomationElement7Interface interface {
	IUIAutomationElement6Interface
	FindFirstWithOptions(scope TreeScope, condition *IUIAutomationCondition, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElement) HRESULT
	FindAllWithOptions(scope TreeScope, condition *IUIAutomationCondition, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElementArray) HRESULT
	FindFirstWithOptionsBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElement) HRESULT
	FindAllWithOptionsBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElementArray) HRESULT
	GetCurrentMetadataValue(targetId int32, metadataId int32, returnVal *VARIANT) HRESULT
}

type IUIAutomationElement7Vtbl struct {
	IUIAutomationElement6Vtbl
	FindFirstWithOptions uintptr
	FindAllWithOptions uintptr
	FindFirstWithOptionsBuildCache uintptr
	FindAllWithOptionsBuildCache uintptr
	GetCurrentMetadataValue uintptr
}

type IUIAutomationElement7 struct {
	IUIAutomationElement6
}

func (this *IUIAutomationElement7) Vtbl() *IUIAutomationElement7Vtbl {
	return (*IUIAutomationElement7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement7) FindFirstWithOptions(scope TreeScope, condition *IUIAutomationCondition, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindFirstWithOptions, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(traversalOptions), uintptr(unsafe.Pointer(root)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement7) FindAllWithOptions(scope TreeScope, condition *IUIAutomationCondition, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindAllWithOptions, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(traversalOptions), uintptr(unsafe.Pointer(root)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement7) FindFirstWithOptionsBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindFirstWithOptionsBuildCache, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(traversalOptions), uintptr(unsafe.Pointer(root)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement7) FindAllWithOptionsBuildCache(scope TreeScope, condition *IUIAutomationCondition, cacheRequest *IUIAutomationCacheRequest, traversalOptions TreeTraversalOptions, root *IUIAutomationElement, found **IUIAutomationElementArray) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().FindAllWithOptionsBuildCache, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(traversalOptions), uintptr(unsafe.Pointer(root)), uintptr(unsafe.Pointer(found)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement7) GetCurrentMetadataValue(targetId int32, metadataId int32, returnVal *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetCurrentMetadataValue, uintptr(unsafe.Pointer(this)), uintptr(targetId), uintptr(metadataId), uintptr(unsafe.Pointer(returnVal)))
	return HRESULT(ret)
}

// 8c60217d-5411-4cde-bcc0-1ceda223830c
var IID_IUIAutomationElement8 = syscall.GUID{0x8c60217d, 0x5411, 0x4cde, 
	[8]byte{0xbc, 0xc0, 0x1c, 0xed, 0xa2, 0x23, 0x83, 0x0c}}

type IUIAutomationElement8Interface interface {
	IUIAutomationElement7Interface
	Get_CurrentHeadingLevel(retVal *int32) HRESULT
	Get_CachedHeadingLevel(retVal *int32) HRESULT
}

type IUIAutomationElement8Vtbl struct {
	IUIAutomationElement7Vtbl
	Get_CurrentHeadingLevel uintptr
	Get_CachedHeadingLevel uintptr
}

type IUIAutomationElement8 struct {
	IUIAutomationElement7
}

func (this *IUIAutomationElement8) Vtbl() *IUIAutomationElement8Vtbl {
	return (*IUIAutomationElement8Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement8) Get_CurrentHeadingLevel(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentHeadingLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement8) Get_CachedHeadingLevel(retVal *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedHeadingLevel, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 39325fac-039d-440e-a3a3-5eb81a5cecc3
var IID_IUIAutomationElement9 = syscall.GUID{0x39325fac, 0x039d, 0x440e, 
	[8]byte{0xa3, 0xa3, 0x5e, 0xb8, 0x1a, 0x5c, 0xec, 0xc3}}

type IUIAutomationElement9Interface interface {
	IUIAutomationElement8Interface
	Get_CurrentIsDialog(retVal *BOOL) HRESULT
	Get_CachedIsDialog(retVal *BOOL) HRESULT
}

type IUIAutomationElement9Vtbl struct {
	IUIAutomationElement8Vtbl
	Get_CurrentIsDialog uintptr
	Get_CachedIsDialog uintptr
}

type IUIAutomationElement9 struct {
	IUIAutomationElement8
}

func (this *IUIAutomationElement9) Vtbl() *IUIAutomationElement9Vtbl {
	return (*IUIAutomationElement9Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationElement9) Get_CurrentIsDialog(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CurrentIsDialog, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

func (this *IUIAutomationElement9) Get_CachedIsDialog(retVal *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CachedIsDialog, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(retVal)))
	return HRESULT(ret)
}

// 85b94ecd-849d-42b6-b94d-d6db23fdf5a4
var IID_IUIAutomationProxyFactory = syscall.GUID{0x85b94ecd, 0x849d, 0x42b6, 
	[8]byte{0xb9, 0x4d, 0xd6, 0xdb, 0x23, 0xfd, 0xf5, 0xa4}}

type IUIAutomationProxyFactoryInterface interface {
	IUnknownInterface
	CreateProvider(hwnd HWND, idObject int32, idChild int32, provider **IRawElementProviderSimple) HRESULT
	Get_ProxyFactoryId(factoryId *BSTR) HRESULT
}

type IUIAutomationProxyFactoryVtbl struct {
	IUnknownVtbl
	CreateProvider uintptr
	Get_ProxyFactoryId uintptr
}

type IUIAutomationProxyFactory struct {
	IUnknown
}

func (this *IUIAutomationProxyFactory) Vtbl() *IUIAutomationProxyFactoryVtbl {
	return (*IUIAutomationProxyFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationProxyFactory) CreateProvider(hwnd HWND, idObject int32, idChild int32, provider **IRawElementProviderSimple) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateProvider, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(idObject), uintptr(idChild), uintptr(unsafe.Pointer(provider)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactory) Get_ProxyFactoryId(factoryId *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyFactoryId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factoryId)))
	return HRESULT(ret)
}

// d50e472e-b64b-490c-bca1-d30696f9f289
var IID_IUIAutomationProxyFactoryEntry = syscall.GUID{0xd50e472e, 0xb64b, 0x490c, 
	[8]byte{0xbc, 0xa1, 0xd3, 0x06, 0x96, 0xf9, 0xf2, 0x89}}

type IUIAutomationProxyFactoryEntryInterface interface {
	IUnknownInterface
	Get_ProxyFactory(factory **IUIAutomationProxyFactory) HRESULT
	Get_ClassName(className *BSTR) HRESULT
	Get_ImageName(imageName *BSTR) HRESULT
	Get_AllowSubstringMatch(allowSubstringMatch *BOOL) HRESULT
	Get_CanCheckBaseClass(canCheckBaseClass *BOOL) HRESULT
	Get_NeedsAdviseEvents(adviseEvents *BOOL) HRESULT
	Put_ClassName(className PWSTR) HRESULT
	Put_ImageName(imageName PWSTR) HRESULT
	Put_AllowSubstringMatch(allowSubstringMatch BOOL) HRESULT
	Put_CanCheckBaseClass(canCheckBaseClass BOOL) HRESULT
	Put_NeedsAdviseEvents(adviseEvents BOOL) HRESULT
	SetWinEventsForAutomationEvent(eventId int32, propertyId int32, winEvents *SAFEARRAY) HRESULT
	GetWinEventsForAutomationEvent(eventId int32, propertyId int32, winEvents **SAFEARRAY) HRESULT
}

type IUIAutomationProxyFactoryEntryVtbl struct {
	IUnknownVtbl
	Get_ProxyFactory uintptr
	Get_ClassName uintptr
	Get_ImageName uintptr
	Get_AllowSubstringMatch uintptr
	Get_CanCheckBaseClass uintptr
	Get_NeedsAdviseEvents uintptr
	Put_ClassName uintptr
	Put_ImageName uintptr
	Put_AllowSubstringMatch uintptr
	Put_CanCheckBaseClass uintptr
	Put_NeedsAdviseEvents uintptr
	SetWinEventsForAutomationEvent uintptr
	GetWinEventsForAutomationEvent uintptr
}

type IUIAutomationProxyFactoryEntry struct {
	IUnknown
}

func (this *IUIAutomationProxyFactoryEntry) Vtbl() *IUIAutomationProxyFactoryEntryVtbl {
	return (*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationProxyFactoryEntry) Get_ProxyFactory(factory **IUIAutomationProxyFactory) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyFactory, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Get_ClassName(className *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ClassName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(className)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Get_ImageName(imageName *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ImageName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageName)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Get_AllowSubstringMatch(allowSubstringMatch *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowSubstringMatch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(allowSubstringMatch)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Get_CanCheckBaseClass(canCheckBaseClass *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CanCheckBaseClass, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(canCheckBaseClass)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Get_NeedsAdviseEvents(adviseEvents *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_NeedsAdviseEvents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(adviseEvents)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Put_ClassName(className PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_ClassName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(className)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Put_ImageName(imageName PWSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_ImageName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(imageName)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Put_AllowSubstringMatch(allowSubstringMatch BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowSubstringMatch, uintptr(unsafe.Pointer(this)), uintptr(allowSubstringMatch))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Put_CanCheckBaseClass(canCheckBaseClass BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_CanCheckBaseClass, uintptr(unsafe.Pointer(this)), uintptr(canCheckBaseClass))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) Put_NeedsAdviseEvents(adviseEvents BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_NeedsAdviseEvents, uintptr(unsafe.Pointer(this)), uintptr(adviseEvents))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) SetWinEventsForAutomationEvent(eventId int32, propertyId int32, winEvents *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetWinEventsForAutomationEvent, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(propertyId), uintptr(unsafe.Pointer(winEvents)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryEntry) GetWinEventsForAutomationEvent(eventId int32, propertyId int32, winEvents **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetWinEventsForAutomationEvent, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(propertyId), uintptr(unsafe.Pointer(winEvents)))
	return HRESULT(ret)
}

// 09e31e18-872d-4873-93d1-1e541ec133fd
var IID_IUIAutomationProxyFactoryMapping = syscall.GUID{0x09e31e18, 0x872d, 0x4873, 
	[8]byte{0x93, 0xd1, 0x1e, 0x54, 0x1e, 0xc1, 0x33, 0xfd}}

type IUIAutomationProxyFactoryMappingInterface interface {
	IUnknownInterface
	Get_Count(count *uint32) HRESULT
	GetTable(table **SAFEARRAY) HRESULT
	GetEntry(index uint32, entry **IUIAutomationProxyFactoryEntry) HRESULT
	SetTable(factoryList *SAFEARRAY) HRESULT
	InsertEntries(before uint32, factoryList *SAFEARRAY) HRESULT
	InsertEntry(before uint32, factory *IUIAutomationProxyFactoryEntry) HRESULT
	RemoveEntry(index uint32) HRESULT
	ClearTable() HRESULT
	RestoreDefaultTable() HRESULT
}

type IUIAutomationProxyFactoryMappingVtbl struct {
	IUnknownVtbl
	Get_Count uintptr
	GetTable uintptr
	GetEntry uintptr
	SetTable uintptr
	InsertEntries uintptr
	InsertEntry uintptr
	RemoveEntry uintptr
	ClearTable uintptr
	RestoreDefaultTable uintptr
}

type IUIAutomationProxyFactoryMapping struct {
	IUnknown
}

func (this *IUIAutomationProxyFactoryMapping) Vtbl() *IUIAutomationProxyFactoryMappingVtbl {
	return (*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationProxyFactoryMapping) Get_Count(count *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_Count, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) GetTable(table **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetTable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(table)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) GetEntry(index uint32, entry **IUIAutomationProxyFactoryEntry) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetEntry, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(entry)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) SetTable(factoryList *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SetTable, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factoryList)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) InsertEntries(before uint32, factoryList *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertEntries, uintptr(unsafe.Pointer(this)), uintptr(before), uintptr(unsafe.Pointer(factoryList)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) InsertEntry(before uint32, factory *IUIAutomationProxyFactoryEntry) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().InsertEntry, uintptr(unsafe.Pointer(this)), uintptr(before), uintptr(unsafe.Pointer(factory)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) RemoveEntry(index uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveEntry, uintptr(unsafe.Pointer(this)), uintptr(index))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) ClearTable() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ClearTable, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomationProxyFactoryMapping) RestoreDefaultTable() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RestoreDefaultTable, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

// c9ee12f2-c13b-4408-997c-639914377f4e
var IID_IUIAutomationEventHandlerGroup = syscall.GUID{0xc9ee12f2, 0xc13b, 0x4408, 
	[8]byte{0x99, 0x7c, 0x63, 0x99, 0x14, 0x37, 0x7f, 0x4e}}

type IUIAutomationEventHandlerGroupInterface interface {
	IUnknownInterface
	AddActiveTextPositionChangedEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationActiveTextPositionChangedEventHandler) HRESULT
	AddAutomationEventHandler(eventId int32, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationEventHandler) HRESULT
	AddChangesEventHandler(scope TreeScope, changeTypes *int32, changesCount int32, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationChangesEventHandler) HRESULT
	AddNotificationEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationNotificationEventHandler) HRESULT
	AddPropertyChangedEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationPropertyChangedEventHandler, propertyArray *int32, propertyCount int32) HRESULT
	AddStructureChangedEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationStructureChangedEventHandler) HRESULT
	AddTextEditTextChangedEventHandler(scope TreeScope, textEditChangeType TextEditChangeType, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationTextEditTextChangedEventHandler) HRESULT
}

type IUIAutomationEventHandlerGroupVtbl struct {
	IUnknownVtbl
	AddActiveTextPositionChangedEventHandler uintptr
	AddAutomationEventHandler uintptr
	AddChangesEventHandler uintptr
	AddNotificationEventHandler uintptr
	AddPropertyChangedEventHandler uintptr
	AddStructureChangedEventHandler uintptr
	AddTextEditTextChangedEventHandler uintptr
}

type IUIAutomationEventHandlerGroup struct {
	IUnknown
}

func (this *IUIAutomationEventHandlerGroup) Vtbl() *IUIAutomationEventHandlerGroupVtbl {
	return (*IUIAutomationEventHandlerGroupVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomationEventHandlerGroup) AddActiveTextPositionChangedEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationActiveTextPositionChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddActiveTextPositionChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomationEventHandlerGroup) AddAutomationEventHandler(eventId int32, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddAutomationEventHandler, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomationEventHandlerGroup) AddChangesEventHandler(scope TreeScope, changeTypes *int32, changesCount int32, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationChangesEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddChangesEventHandler, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(changeTypes)), uintptr(changesCount), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomationEventHandlerGroup) AddNotificationEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationNotificationEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddNotificationEventHandler, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomationEventHandlerGroup) AddPropertyChangedEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationPropertyChangedEventHandler, propertyArray *int32, propertyCount int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddPropertyChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(propertyArray)), uintptr(propertyCount))
	return HRESULT(ret)
}

func (this *IUIAutomationEventHandlerGroup) AddStructureChangedEventHandler(scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationStructureChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddStructureChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomationEventHandlerGroup) AddTextEditTextChangedEventHandler(scope TreeScope, textEditChangeType TextEditChangeType, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationTextEditTextChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddTextEditTextChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(scope), uintptr(textEditChangeType), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

// 30cbe57d-d9d0-452a-ab13-7ac5ac4825ee
var IID_IUIAutomation = syscall.GUID{0x30cbe57d, 0xd9d0, 0x452a, 
	[8]byte{0xab, 0x13, 0x7a, 0xc5, 0xac, 0x48, 0x25, 0xee}}

type IUIAutomationInterface interface {
	IUnknownInterface
	CompareElements(el1 *IUIAutomationElement, el2 *IUIAutomationElement, areSame *BOOL) HRESULT
	CompareRuntimeIds(runtimeId1 *SAFEARRAY, runtimeId2 *SAFEARRAY, areSame *BOOL) HRESULT
	GetRootElement(root **IUIAutomationElement) HRESULT
	ElementFromHandle(hwnd HWND, element **IUIAutomationElement) HRESULT
	ElementFromPoint(pt POINT, element **IUIAutomationElement) HRESULT
	GetFocusedElement(element **IUIAutomationElement) HRESULT
	GetRootElementBuildCache(cacheRequest *IUIAutomationCacheRequest, root **IUIAutomationElement) HRESULT
	ElementFromHandleBuildCache(hwnd HWND, cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT
	ElementFromPointBuildCache(pt POINT, cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT
	GetFocusedElementBuildCache(cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT
	CreateTreeWalker(pCondition *IUIAutomationCondition, walker **IUIAutomationTreeWalker) HRESULT
	Get_ControlViewWalker(walker **IUIAutomationTreeWalker) HRESULT
	Get_ContentViewWalker(walker **IUIAutomationTreeWalker) HRESULT
	Get_RawViewWalker(walker **IUIAutomationTreeWalker) HRESULT
	Get_RawViewCondition(condition **IUIAutomationCondition) HRESULT
	Get_ControlViewCondition(condition **IUIAutomationCondition) HRESULT
	Get_ContentViewCondition(condition **IUIAutomationCondition) HRESULT
	CreateCacheRequest(cacheRequest **IUIAutomationCacheRequest) HRESULT
	CreateTrueCondition(newCondition **IUIAutomationCondition) HRESULT
	CreateFalseCondition(newCondition **IUIAutomationCondition) HRESULT
	CreatePropertyCondition(propertyId int32, value VARIANT, newCondition **IUIAutomationCondition) HRESULT
	CreatePropertyConditionEx(propertyId int32, value VARIANT, flags PropertyConditionFlags, newCondition **IUIAutomationCondition) HRESULT
	CreateAndCondition(condition1 *IUIAutomationCondition, condition2 *IUIAutomationCondition, newCondition **IUIAutomationCondition) HRESULT
	CreateAndConditionFromArray(conditions *SAFEARRAY, newCondition **IUIAutomationCondition) HRESULT
	CreateAndConditionFromNativeArray(conditions **IUIAutomationCondition, conditionCount int32, newCondition **IUIAutomationCondition) HRESULT
	CreateOrCondition(condition1 *IUIAutomationCondition, condition2 *IUIAutomationCondition, newCondition **IUIAutomationCondition) HRESULT
	CreateOrConditionFromArray(conditions *SAFEARRAY, newCondition **IUIAutomationCondition) HRESULT
	CreateOrConditionFromNativeArray(conditions **IUIAutomationCondition, conditionCount int32, newCondition **IUIAutomationCondition) HRESULT
	CreateNotCondition(condition *IUIAutomationCondition, newCondition **IUIAutomationCondition) HRESULT
	AddAutomationEventHandler(eventId int32, element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationEventHandler) HRESULT
	RemoveAutomationEventHandler(eventId int32, element *IUIAutomationElement, handler *IUIAutomationEventHandler) HRESULT
	AddPropertyChangedEventHandlerNativeArray(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationPropertyChangedEventHandler, propertyArray *int32, propertyCount int32) HRESULT
	AddPropertyChangedEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationPropertyChangedEventHandler, propertyArray *SAFEARRAY) HRESULT
	RemovePropertyChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationPropertyChangedEventHandler) HRESULT
	AddStructureChangedEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationStructureChangedEventHandler) HRESULT
	RemoveStructureChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationStructureChangedEventHandler) HRESULT
	AddFocusChangedEventHandler(cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationFocusChangedEventHandler) HRESULT
	RemoveFocusChangedEventHandler(handler *IUIAutomationFocusChangedEventHandler) HRESULT
	RemoveAllEventHandlers() HRESULT
	IntNativeArrayToSafeArray(array *int32, arrayCount int32, safeArray **SAFEARRAY) HRESULT
	IntSafeArrayToNativeArray(intArray *SAFEARRAY, array **int32, arrayCount *int32) HRESULT
	RectToVariant(rc RECT, var_ *VARIANT) HRESULT
	VariantToRect(var_ VARIANT, rc *RECT) HRESULT
	SafeArrayToRectNativeArray(rects *SAFEARRAY, rectArray **RECT, rectArrayCount *int32) HRESULT
	CreateProxyFactoryEntry(factory *IUIAutomationProxyFactory, factoryEntry **IUIAutomationProxyFactoryEntry) HRESULT
	Get_ProxyFactoryMapping(factoryMapping **IUIAutomationProxyFactoryMapping) HRESULT
	GetPropertyProgrammaticName(property int32, name *BSTR) HRESULT
	GetPatternProgrammaticName(pattern int32, name *BSTR) HRESULT
	PollForPotentialSupportedPatterns(pElement *IUIAutomationElement, patternIds **SAFEARRAY, patternNames **SAFEARRAY) HRESULT
	PollForPotentialSupportedProperties(pElement *IUIAutomationElement, propertyIds **SAFEARRAY, propertyNames **SAFEARRAY) HRESULT
	CheckNotSupported(value VARIANT, isNotSupported *BOOL) HRESULT
	Get_ReservedNotSupportedValue(notSupportedValue **IUnknown) HRESULT
	Get_ReservedMixedAttributeValue(mixedAttributeValue **IUnknown) HRESULT
	ElementFromIAccessible(accessible *IAccessible, childId int32, element **IUIAutomationElement) HRESULT
	ElementFromIAccessibleBuildCache(accessible *IAccessible, childId int32, cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT
}

type IUIAutomationVtbl struct {
	IUnknownVtbl
	CompareElements uintptr
	CompareRuntimeIds uintptr
	GetRootElement uintptr
	ElementFromHandle uintptr
	ElementFromPoint uintptr
	GetFocusedElement uintptr
	GetRootElementBuildCache uintptr
	ElementFromHandleBuildCache uintptr
	ElementFromPointBuildCache uintptr
	GetFocusedElementBuildCache uintptr
	CreateTreeWalker uintptr
	Get_ControlViewWalker uintptr
	Get_ContentViewWalker uintptr
	Get_RawViewWalker uintptr
	Get_RawViewCondition uintptr
	Get_ControlViewCondition uintptr
	Get_ContentViewCondition uintptr
	CreateCacheRequest uintptr
	CreateTrueCondition uintptr
	CreateFalseCondition uintptr
	CreatePropertyCondition uintptr
	CreatePropertyConditionEx uintptr
	CreateAndCondition uintptr
	CreateAndConditionFromArray uintptr
	CreateAndConditionFromNativeArray uintptr
	CreateOrCondition uintptr
	CreateOrConditionFromArray uintptr
	CreateOrConditionFromNativeArray uintptr
	CreateNotCondition uintptr
	AddAutomationEventHandler uintptr
	RemoveAutomationEventHandler uintptr
	AddPropertyChangedEventHandlerNativeArray uintptr
	AddPropertyChangedEventHandler uintptr
	RemovePropertyChangedEventHandler uintptr
	AddStructureChangedEventHandler uintptr
	RemoveStructureChangedEventHandler uintptr
	AddFocusChangedEventHandler uintptr
	RemoveFocusChangedEventHandler uintptr
	RemoveAllEventHandlers uintptr
	IntNativeArrayToSafeArray uintptr
	IntSafeArrayToNativeArray uintptr
	RectToVariant uintptr
	VariantToRect uintptr
	SafeArrayToRectNativeArray uintptr
	CreateProxyFactoryEntry uintptr
	Get_ProxyFactoryMapping uintptr
	GetPropertyProgrammaticName uintptr
	GetPatternProgrammaticName uintptr
	PollForPotentialSupportedPatterns uintptr
	PollForPotentialSupportedProperties uintptr
	CheckNotSupported uintptr
	Get_ReservedNotSupportedValue uintptr
	Get_ReservedMixedAttributeValue uintptr
	ElementFromIAccessible uintptr
	ElementFromIAccessibleBuildCache uintptr
}

type IUIAutomation struct {
	IUnknown
}

func (this *IUIAutomation) Vtbl() *IUIAutomationVtbl {
	return (*IUIAutomationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomation) CompareElements(el1 *IUIAutomationElement, el2 *IUIAutomationElement, areSame *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CompareElements, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(el1)), uintptr(unsafe.Pointer(el2)), uintptr(unsafe.Pointer(areSame)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CompareRuntimeIds(runtimeId1 *SAFEARRAY, runtimeId2 *SAFEARRAY, areSame *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CompareRuntimeIds, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(runtimeId1)), uintptr(unsafe.Pointer(runtimeId2)), uintptr(unsafe.Pointer(areSame)))
	return HRESULT(ret)
}

func (this *IUIAutomation) GetRootElement(root **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRootElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(root)))
	return HRESULT(ret)
}

func (this *IUIAutomation) ElementFromHandle(hwnd HWND, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementFromHandle, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) ElementFromPoint(pt POINT, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementFromPoint, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pt)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) GetFocusedElement(element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFocusedElement, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) GetRootElementBuildCache(cacheRequest *IUIAutomationCacheRequest, root **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetRootElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(root)))
	return HRESULT(ret)
}

func (this *IUIAutomation) ElementFromHandleBuildCache(hwnd HWND, cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementFromHandleBuildCache, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) ElementFromPointBuildCache(pt POINT, cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementFromPointBuildCache, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pt)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) GetFocusedElementBuildCache(cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetFocusedElementBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateTreeWalker(pCondition *IUIAutomationCondition, walker **IUIAutomationTreeWalker) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateTreeWalker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pCondition)), uintptr(unsafe.Pointer(walker)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ControlViewWalker(walker **IUIAutomationTreeWalker) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlViewWalker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(walker)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ContentViewWalker(walker **IUIAutomationTreeWalker) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentViewWalker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(walker)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_RawViewWalker(walker **IUIAutomationTreeWalker) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_RawViewWalker, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(walker)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_RawViewCondition(condition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_RawViewCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ControlViewCondition(condition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ControlViewCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ContentViewCondition(condition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ContentViewCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateCacheRequest(cacheRequest **IUIAutomationCacheRequest) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateCacheRequest, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateTrueCondition(newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateTrueCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateFalseCondition(newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateFalseCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreatePropertyCondition(propertyId int32, value VARIANT, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePropertyCondition, uintptr(unsafe.Pointer(this)), uintptr(propertyId), (uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreatePropertyConditionEx(propertyId int32, value VARIANT, flags PropertyConditionFlags, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreatePropertyConditionEx, uintptr(unsafe.Pointer(this)), uintptr(propertyId), (uintptr)(unsafe.Pointer(&value)), uintptr(flags), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateAndCondition(condition1 *IUIAutomationCondition, condition2 *IUIAutomationCondition, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateAndCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition1)), uintptr(unsafe.Pointer(condition2)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateAndConditionFromArray(conditions *SAFEARRAY, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateAndConditionFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditions)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateAndConditionFromNativeArray(conditions **IUIAutomationCondition, conditionCount int32, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateAndConditionFromNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditions)), uintptr(conditionCount), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateOrCondition(condition1 *IUIAutomationCondition, condition2 *IUIAutomationCondition, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateOrCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition1)), uintptr(unsafe.Pointer(condition2)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateOrConditionFromArray(conditions *SAFEARRAY, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateOrConditionFromArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditions)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateOrConditionFromNativeArray(conditions **IUIAutomationCondition, conditionCount int32, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateOrConditionFromNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(conditions)), uintptr(conditionCount), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateNotCondition(condition *IUIAutomationCondition, newCondition **IUIAutomationCondition) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateNotCondition, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(condition)), uintptr(unsafe.Pointer(newCondition)))
	return HRESULT(ret)
}

func (this *IUIAutomation) AddAutomationEventHandler(eventId int32, element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddAutomationEventHandler, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) RemoveAutomationEventHandler(eventId int32, element *IUIAutomationElement, handler *IUIAutomationEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAutomationEventHandler, uintptr(unsafe.Pointer(this)), uintptr(eventId), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) AddPropertyChangedEventHandlerNativeArray(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationPropertyChangedEventHandler, propertyArray *int32, propertyCount int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddPropertyChangedEventHandlerNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(propertyArray)), uintptr(propertyCount))
	return HRESULT(ret)
}

func (this *IUIAutomation) AddPropertyChangedEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationPropertyChangedEventHandler, propertyArray *SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddPropertyChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(propertyArray)))
	return HRESULT(ret)
}

func (this *IUIAutomation) RemovePropertyChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationPropertyChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemovePropertyChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) AddStructureChangedEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationStructureChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddStructureChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) RemoveStructureChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationStructureChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveStructureChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) AddFocusChangedEventHandler(cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationFocusChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddFocusChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) RemoveFocusChangedEventHandler(handler *IUIAutomationFocusChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveFocusChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation) RemoveAllEventHandlers() HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveAllEventHandlers, uintptr(unsafe.Pointer(this)))
	return HRESULT(ret)
}

func (this *IUIAutomation) IntNativeArrayToSafeArray(array *int32, arrayCount int32, safeArray **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IntNativeArrayToSafeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(array)), uintptr(arrayCount), uintptr(unsafe.Pointer(safeArray)))
	return HRESULT(ret)
}

func (this *IUIAutomation) IntSafeArrayToNativeArray(intArray *SAFEARRAY, array **int32, arrayCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().IntSafeArrayToNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(intArray)), uintptr(unsafe.Pointer(array)), uintptr(unsafe.Pointer(arrayCount)))
	return HRESULT(ret)
}

func (this *IUIAutomation) RectToVariant(rc RECT, var_ *VARIANT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RectToVariant, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&rc)), uintptr(unsafe.Pointer(var_)))
	return HRESULT(ret)
}

func (this *IUIAutomation) VariantToRect(var_ VARIANT, rc *RECT) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().VariantToRect, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&var_)), uintptr(unsafe.Pointer(rc)))
	return HRESULT(ret)
}

func (this *IUIAutomation) SafeArrayToRectNativeArray(rects *SAFEARRAY, rectArray **RECT, rectArrayCount *int32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().SafeArrayToRectNativeArray, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(rects)), uintptr(unsafe.Pointer(rectArray)), uintptr(unsafe.Pointer(rectArrayCount)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CreateProxyFactoryEntry(factory *IUIAutomationProxyFactory, factoryEntry **IUIAutomationProxyFactoryEntry) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateProxyFactoryEntry, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factory)), uintptr(unsafe.Pointer(factoryEntry)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ProxyFactoryMapping(factoryMapping **IUIAutomationProxyFactoryMapping) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ProxyFactoryMapping, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(factoryMapping)))
	return HRESULT(ret)
}

func (this *IUIAutomation) GetPropertyProgrammaticName(property int32, name *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPropertyProgrammaticName, uintptr(unsafe.Pointer(this)), uintptr(property), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IUIAutomation) GetPatternProgrammaticName(pattern int32, name *BSTR) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().GetPatternProgrammaticName, uintptr(unsafe.Pointer(this)), uintptr(pattern), uintptr(unsafe.Pointer(name)))
	return HRESULT(ret)
}

func (this *IUIAutomation) PollForPotentialSupportedPatterns(pElement *IUIAutomationElement, patternIds **SAFEARRAY, patternNames **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().PollForPotentialSupportedPatterns, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pElement)), uintptr(unsafe.Pointer(patternIds)), uintptr(unsafe.Pointer(patternNames)))
	return HRESULT(ret)
}

func (this *IUIAutomation) PollForPotentialSupportedProperties(pElement *IUIAutomationElement, propertyIds **SAFEARRAY, propertyNames **SAFEARRAY) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().PollForPotentialSupportedProperties, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pElement)), uintptr(unsafe.Pointer(propertyIds)), uintptr(unsafe.Pointer(propertyNames)))
	return HRESULT(ret)
}

func (this *IUIAutomation) CheckNotSupported(value VARIANT, isNotSupported *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CheckNotSupported, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(isNotSupported)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ReservedNotSupportedValue(notSupportedValue **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ReservedNotSupportedValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(notSupportedValue)))
	return HRESULT(ret)
}

func (this *IUIAutomation) Get_ReservedMixedAttributeValue(mixedAttributeValue **IUnknown) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ReservedMixedAttributeValue, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mixedAttributeValue)))
	return HRESULT(ret)
}

func (this *IUIAutomation) ElementFromIAccessible(accessible *IAccessible, childId int32, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementFromIAccessible, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(accessible)), uintptr(childId), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

func (this *IUIAutomation) ElementFromIAccessibleBuildCache(accessible *IAccessible, childId int32, cacheRequest *IUIAutomationCacheRequest, element **IUIAutomationElement) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().ElementFromIAccessibleBuildCache, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(accessible)), uintptr(childId), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(element)))
	return HRESULT(ret)
}

// 34723aff-0c9d-49d0-9896-7ab52df8cd8a
var IID_IUIAutomation2 = syscall.GUID{0x34723aff, 0x0c9d, 0x49d0, 
	[8]byte{0x98, 0x96, 0x7a, 0xb5, 0x2d, 0xf8, 0xcd, 0x8a}}

type IUIAutomation2Interface interface {
	IUIAutomationInterface
	Get_AutoSetFocus(autoSetFocus *BOOL) HRESULT
	Put_AutoSetFocus(autoSetFocus BOOL) HRESULT
	Get_ConnectionTimeout(timeout *uint32) HRESULT
	Put_ConnectionTimeout(timeout uint32) HRESULT
	Get_TransactionTimeout(timeout *uint32) HRESULT
	Put_TransactionTimeout(timeout uint32) HRESULT
}

type IUIAutomation2Vtbl struct {
	IUIAutomationVtbl
	Get_AutoSetFocus uintptr
	Put_AutoSetFocus uintptr
	Get_ConnectionTimeout uintptr
	Put_ConnectionTimeout uintptr
	Get_TransactionTimeout uintptr
	Put_TransactionTimeout uintptr
}

type IUIAutomation2 struct {
	IUIAutomation
}

func (this *IUIAutomation2) Vtbl() *IUIAutomation2Vtbl {
	return (*IUIAutomation2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomation2) Get_AutoSetFocus(autoSetFocus *BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_AutoSetFocus, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(autoSetFocus)))
	return HRESULT(ret)
}

func (this *IUIAutomation2) Put_AutoSetFocus(autoSetFocus BOOL) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_AutoSetFocus, uintptr(unsafe.Pointer(this)), uintptr(autoSetFocus))
	return HRESULT(ret)
}

func (this *IUIAutomation2) Get_ConnectionTimeout(timeout *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timeout)))
	return HRESULT(ret)
}

func (this *IUIAutomation2) Put_ConnectionTimeout(timeout uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_ConnectionTimeout, uintptr(unsafe.Pointer(this)), uintptr(timeout))
	return HRESULT(ret)
}

func (this *IUIAutomation2) Get_TransactionTimeout(timeout *uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_TransactionTimeout, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(timeout)))
	return HRESULT(ret)
}

func (this *IUIAutomation2) Put_TransactionTimeout(timeout uint32) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_TransactionTimeout, uintptr(unsafe.Pointer(this)), uintptr(timeout))
	return HRESULT(ret)
}

// 73d768da-9b51-4b89-936e-c209290973e7
var IID_IUIAutomation3 = syscall.GUID{0x73d768da, 0x9b51, 0x4b89, 
	[8]byte{0x93, 0x6e, 0xc2, 0x09, 0x29, 0x09, 0x73, 0xe7}}

type IUIAutomation3Interface interface {
	IUIAutomation2Interface
	AddTextEditTextChangedEventHandler(element *IUIAutomationElement, scope TreeScope, textEditChangeType TextEditChangeType, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationTextEditTextChangedEventHandler) HRESULT
	RemoveTextEditTextChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationTextEditTextChangedEventHandler) HRESULT
}

type IUIAutomation3Vtbl struct {
	IUIAutomation2Vtbl
	AddTextEditTextChangedEventHandler uintptr
	RemoveTextEditTextChangedEventHandler uintptr
}

type IUIAutomation3 struct {
	IUIAutomation2
}

func (this *IUIAutomation3) Vtbl() *IUIAutomation3Vtbl {
	return (*IUIAutomation3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomation3) AddTextEditTextChangedEventHandler(element *IUIAutomationElement, scope TreeScope, textEditChangeType TextEditChangeType, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationTextEditTextChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddTextEditTextChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(textEditChangeType), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation3) RemoveTextEditTextChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationTextEditTextChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveTextEditTextChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

// 1189c02a-05f8-4319-8e21-e817e3db2860
var IID_IUIAutomation4 = syscall.GUID{0x1189c02a, 0x05f8, 0x4319, 
	[8]byte{0x8e, 0x21, 0xe8, 0x17, 0xe3, 0xdb, 0x28, 0x60}}

type IUIAutomation4Interface interface {
	IUIAutomation3Interface
	AddChangesEventHandler(element *IUIAutomationElement, scope TreeScope, changeTypes *int32, changesCount int32, pCacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationChangesEventHandler) HRESULT
	RemoveChangesEventHandler(element *IUIAutomationElement, handler *IUIAutomationChangesEventHandler) HRESULT
}

type IUIAutomation4Vtbl struct {
	IUIAutomation3Vtbl
	AddChangesEventHandler uintptr
	RemoveChangesEventHandler uintptr
}

type IUIAutomation4 struct {
	IUIAutomation3
}

func (this *IUIAutomation4) Vtbl() *IUIAutomation4Vtbl {
	return (*IUIAutomation4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomation4) AddChangesEventHandler(element *IUIAutomationElement, scope TreeScope, changeTypes *int32, changesCount int32, pCacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationChangesEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddChangesEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(changeTypes)), uintptr(changesCount), uintptr(unsafe.Pointer(pCacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation4) RemoveChangesEventHandler(element *IUIAutomationElement, handler *IUIAutomationChangesEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveChangesEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

// 25f700c8-d816-4057-a9dc-3cbdee77e256
var IID_IUIAutomation5 = syscall.GUID{0x25f700c8, 0xd816, 0x4057, 
	[8]byte{0xa9, 0xdc, 0x3c, 0xbd, 0xee, 0x77, 0xe2, 0x56}}

type IUIAutomation5Interface interface {
	IUIAutomation4Interface
	AddNotificationEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationNotificationEventHandler) HRESULT
	RemoveNotificationEventHandler(element *IUIAutomationElement, handler *IUIAutomationNotificationEventHandler) HRESULT
}

type IUIAutomation5Vtbl struct {
	IUIAutomation4Vtbl
	AddNotificationEventHandler uintptr
	RemoveNotificationEventHandler uintptr
}

type IUIAutomation5 struct {
	IUIAutomation4
}

func (this *IUIAutomation5) Vtbl() *IUIAutomation5Vtbl {
	return (*IUIAutomation5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomation5) AddNotificationEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationNotificationEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddNotificationEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation5) RemoveNotificationEventHandler(element *IUIAutomationElement, handler *IUIAutomationNotificationEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveNotificationEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

// aae072da-29e3-413d-87a7-192dbf81ed10
var IID_IUIAutomation6 = syscall.GUID{0xaae072da, 0x29e3, 0x413d, 
	[8]byte{0x87, 0xa7, 0x19, 0x2d, 0xbf, 0x81, 0xed, 0x10}}

type IUIAutomation6Interface interface {
	IUIAutomation5Interface
	CreateEventHandlerGroup(handlerGroup **IUIAutomationEventHandlerGroup) HRESULT
	AddEventHandlerGroup(element *IUIAutomationElement, handlerGroup *IUIAutomationEventHandlerGroup) HRESULT
	RemoveEventHandlerGroup(element *IUIAutomationElement, handlerGroup *IUIAutomationEventHandlerGroup) HRESULT
	Get_ConnectionRecoveryBehavior(connectionRecoveryBehaviorOptions *ConnectionRecoveryBehaviorOptions) HRESULT
	Put_ConnectionRecoveryBehavior(connectionRecoveryBehaviorOptions ConnectionRecoveryBehaviorOptions) HRESULT
	Get_CoalesceEvents(coalesceEventsOptions *CoalesceEventsOptions) HRESULT
	Put_CoalesceEvents(coalesceEventsOptions CoalesceEventsOptions) HRESULT
	AddActiveTextPositionChangedEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationActiveTextPositionChangedEventHandler) HRESULT
	RemoveActiveTextPositionChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationActiveTextPositionChangedEventHandler) HRESULT
}

type IUIAutomation6Vtbl struct {
	IUIAutomation5Vtbl
	CreateEventHandlerGroup uintptr
	AddEventHandlerGroup uintptr
	RemoveEventHandlerGroup uintptr
	Get_ConnectionRecoveryBehavior uintptr
	Put_ConnectionRecoveryBehavior uintptr
	Get_CoalesceEvents uintptr
	Put_CoalesceEvents uintptr
	AddActiveTextPositionChangedEventHandler uintptr
	RemoveActiveTextPositionChangedEventHandler uintptr
}

type IUIAutomation6 struct {
	IUIAutomation5
}

func (this *IUIAutomation6) Vtbl() *IUIAutomation6Vtbl {
	return (*IUIAutomation6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUIAutomation6) CreateEventHandlerGroup(handlerGroup **IUIAutomationEventHandlerGroup) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().CreateEventHandlerGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handlerGroup)))
	return HRESULT(ret)
}

func (this *IUIAutomation6) AddEventHandlerGroup(element *IUIAutomationElement, handlerGroup *IUIAutomationEventHandlerGroup) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddEventHandlerGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handlerGroup)))
	return HRESULT(ret)
}

func (this *IUIAutomation6) RemoveEventHandlerGroup(element *IUIAutomationElement, handlerGroup *IUIAutomationEventHandlerGroup) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveEventHandlerGroup, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handlerGroup)))
	return HRESULT(ret)
}

func (this *IUIAutomation6) Get_ConnectionRecoveryBehavior(connectionRecoveryBehaviorOptions *ConnectionRecoveryBehaviorOptions) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_ConnectionRecoveryBehavior, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(connectionRecoveryBehaviorOptions)))
	return HRESULT(ret)
}

func (this *IUIAutomation6) Put_ConnectionRecoveryBehavior(connectionRecoveryBehaviorOptions ConnectionRecoveryBehaviorOptions) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_ConnectionRecoveryBehavior, uintptr(unsafe.Pointer(this)), uintptr(connectionRecoveryBehaviorOptions))
	return HRESULT(ret)
}

func (this *IUIAutomation6) Get_CoalesceEvents(coalesceEventsOptions *CoalesceEventsOptions) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Get_CoalesceEvents, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coalesceEventsOptions)))
	return HRESULT(ret)
}

func (this *IUIAutomation6) Put_CoalesceEvents(coalesceEventsOptions CoalesceEventsOptions) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().Put_CoalesceEvents, uintptr(unsafe.Pointer(this)), uintptr(coalesceEventsOptions))
	return HRESULT(ret)
}

func (this *IUIAutomation6) AddActiveTextPositionChangedEventHandler(element *IUIAutomationElement, scope TreeScope, cacheRequest *IUIAutomationCacheRequest, handler *IUIAutomationActiveTextPositionChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().AddActiveTextPositionChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(scope), uintptr(unsafe.Pointer(cacheRequest)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}

func (this *IUIAutomation6) RemoveActiveTextPositionChangedEventHandler(element *IUIAutomationElement, handler *IUIAutomationActiveTextPositionChangedEventHandler) HRESULT{
	ret, _, _ := syscall.SyscallN(this.Vtbl().RemoveActiveTextPositionChangedEventHandler, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(element)), uintptr(unsafe.Pointer(handler)))
	return HRESULT(ret)
}


var (
	pRegisterPointerInputTarget uintptr
	pUnregisterPointerInputTarget uintptr
	pRegisterPointerInputTargetEx uintptr
	pUnregisterPointerInputTargetEx uintptr
	pNotifyWinEvent uintptr
	pSetWinEventHook uintptr
	pIsWinEventHookInstalled uintptr
	pUnhookWinEvent uintptr
)

func RegisterPointerInputTarget(hwnd HWND, pointerType POINTER_INPUT_TYPE) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pRegisterPointerInputTarget, libUser32, "RegisterPointerInputTarget")
	ret, _,  err := syscall.SyscallN(addr, hwnd, uintptr(pointerType))
	return BOOL(ret), WIN32_ERROR(err)
}

func UnregisterPointerInputTarget(hwnd HWND, pointerType POINTER_INPUT_TYPE) (BOOL, WIN32_ERROR) {
	addr := lazyAddr(&pUnregisterPointerInputTarget, libUser32, "UnregisterPointerInputTarget")
	ret, _,  err := syscall.SyscallN(addr, hwnd, uintptr(pointerType))
	return BOOL(ret), WIN32_ERROR(err)
}

func RegisterPointerInputTargetEx(hwnd HWND, pointerType POINTER_INPUT_TYPE, fObserve BOOL) BOOL {
	addr := lazyAddr(&pRegisterPointerInputTargetEx, libUser32, "RegisterPointerInputTargetEx")
	ret, _,  _ := syscall.SyscallN(addr, hwnd, uintptr(pointerType), uintptr(fObserve))
	return BOOL(ret)
}

func UnregisterPointerInputTargetEx(hwnd HWND, pointerType POINTER_INPUT_TYPE) BOOL {
	addr := lazyAddr(&pUnregisterPointerInputTargetEx, libUser32, "UnregisterPointerInputTargetEx")
	ret, _,  _ := syscall.SyscallN(addr, hwnd, uintptr(pointerType))
	return BOOL(ret)
}

func NotifyWinEvent(event uint32, hwnd HWND, idObject int32, idChild int32) {
	addr := lazyAddr(&pNotifyWinEvent, libUser32, "NotifyWinEvent")
	_, _,  _ = syscall.SyscallN(addr, uintptr(event), hwnd, uintptr(idObject), uintptr(idChild))
}

func SetWinEventHook(eventMin uint32, eventMax uint32, hmodWinEventProc HINSTANCE, pfnWinEventProc uintptr, idProcess uint32, idThread uint32, dwFlags uint32) HWINEVENTHOOK {
	addr := lazyAddr(&pSetWinEventHook, libUser32, "SetWinEventHook")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(eventMin), uintptr(eventMax), hmodWinEventProc, uintptr(pfnWinEventProc), uintptr(idProcess), uintptr(idThread), uintptr(dwFlags))
	return HWINEVENTHOOK(ret)
}

func IsWinEventHookInstalled(event uint32) BOOL {
	addr := lazyAddr(&pIsWinEventHookInstalled, libUser32, "IsWinEventHookInstalled")
	ret, _,  _ := syscall.SyscallN(addr, uintptr(event))
	return BOOL(ret)
}

func UnhookWinEvent(hWinEventHook HWINEVENTHOOK) BOOL {
	addr := lazyAddr(&pUnhookWinEvent, libUser32, "UnhookWinEvent")
	ret, _,  _ := syscall.SyscallN(addr, hWinEventHook)
	return BOOL(ret)
}


