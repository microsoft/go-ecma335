

// APIs for Windows.Win32.System.Diagnostics.Debug
//sys	GetEnabledXStateFeatures() (r uint64)
//sys	GetXStateFeaturesMask(Context *CONTEXT, FeatureMask *uint64) (r BOOL)
//sys	LocateXStateFeature(Context *CONTEXT, FeatureId uint32, Length *uint32) (r unsafe.Pointer)
//sys	SetXStateFeaturesMask(Context *CONTEXT, FeatureMask uint64) (r BOOL)
//sys	CheckSumMappedFile(BaseAddress unsafe.Pointer, FileLength uint32, HeaderSum *uint32, CheckSum *uint32) (r *IMAGE_NT_HEADERS32, err error) = imagehlp.CheckSumMappedFile
//sys	GetImageConfigInformation(LoadedImage *LOADED_IMAGE, ImageConfigInformation *IMAGE_LOAD_CONFIG_DIRECTORY32) (r BOOL, err error) = imagehlp.GetImageConfigInformation
//sys	SetImageConfigInformation(LoadedImage *LOADED_IMAGE, ImageConfigInformation *IMAGE_LOAD_CONFIG_DIRECTORY32) (r BOOL, err error) = imagehlp.SetImageConfigInformation
//sys	ImageNtHeader(Base unsafe.Pointer) (r *IMAGE_NT_HEADERS32, err error) = dbghelp.ImageNtHeader
//sys	ImageRvaToSection(NtHeaders *IMAGE_NT_HEADERS32, Base unsafe.Pointer, Rva uint32) (r *IMAGE_SECTION_HEADER, err error) = dbghelp.ImageRvaToSection
//sys	ImageRvaToVa(NtHeaders *IMAGE_NT_HEADERS32, Base unsafe.Pointer, Rva uint32, LastRvaSection **IMAGE_SECTION_HEADER) (r unsafe.Pointer, err error) = dbghelp.ImageRvaToVa
//sys	StackWalk(MachineType uint32, hProcess HANDLE, hThread HANDLE, StackFrame *STACKFRAME, ContextRecord unsafe.Pointer, ReadMemoryRoutine PREAD_PROCESS_MEMORY_ROUTINE, FunctionTableAccessRoutine PFUNCTION_TABLE_ACCESS_ROUTINE, GetModuleBaseRoutine PGET_MODULE_BASE_ROUTINE, TranslateAddress PTRANSLATE_ADDRESS_ROUTINE) (r BOOL) = dbghelp.StackWalk
//sys	SymEnumerateModules(hProcess HANDLE, EnumModulesCallback PSYM_ENUMMODULES_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateModules
//sys	EnumerateLoadedModules(hProcess HANDLE, EnumLoadedModulesCallback PENUMLOADED_MODULES_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumerateLoadedModules
//sys	SymFunctionTableAccess(hProcess HANDLE, AddrBase uint32) (r unsafe.Pointer, err error) = dbghelp.SymFunctionTableAccess
//sys	SymGetModuleInfo(hProcess HANDLE, dwAddr uint32, ModuleInfo *IMAGEHLP_MODULE) (r BOOL, err error) = dbghelp.SymGetModuleInfo
//sys	SymGetModuleInfoW(hProcess HANDLE, dwAddr uint32, ModuleInfo *IMAGEHLP_MODULEW) (r BOOL, err error) = dbghelp.SymGetModuleInfoW
//sys	SymGetModuleBase(hProcess HANDLE, dwAddr uint32) (r uint32, err error) = dbghelp.SymGetModuleBase
//sys	SymGetLineFromAddr(hProcess HANDLE, dwAddr uint32, pdwDisplacement *uint32, Line *IMAGEHLP_LINE) (r BOOL, err error) = dbghelp.SymGetLineFromAddr
//sys	SymGetLineFromName(hProcess HANDLE, ModuleName *PSTRElement, FileName *PSTRElement, dwLineNumber uint32, plDisplacement *int32, Line *IMAGEHLP_LINE) (r BOOL, err error) = dbghelp.SymGetLineFromName
//sys	SymGetLineNext(hProcess HANDLE, Line *IMAGEHLP_LINE) (r BOOL, err error) = dbghelp.SymGetLineNext
//sys	SymGetLinePrev(hProcess HANDLE, Line *IMAGEHLP_LINE) (r BOOL, err error) = dbghelp.SymGetLinePrev
//sys	SymUnloadModule(hProcess HANDLE, BaseOfDll uint32) (r BOOL, err error) = dbghelp.SymUnloadModule
//sys	SymUnDName(sym *IMAGEHLP_SYMBOL, UnDecName *PSTRElement, UnDecNameLength uint32) (r BOOL, err error) = dbghelp.SymUnDName
//sys	SymRegisterCallback(hProcess HANDLE, CallbackFunction PSYMBOL_REGISTERED_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymRegisterCallback
//sys	SymRegisterFunctionEntryCallback(hProcess HANDLE, CallbackFunction PSYMBOL_FUNCENTRY_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymRegisterFunctionEntryCallback
//sys	SymGetSymFromAddr(hProcess HANDLE, dwAddr uint32, pdwDisplacement *uint32, Symbol *IMAGEHLP_SYMBOL) (r BOOL, err error) = dbghelp.SymGetSymFromAddr
//sys	SymGetSymFromName(hProcess HANDLE, Name *PSTRElement, Symbol *IMAGEHLP_SYMBOL) (r BOOL, err error) = dbghelp.SymGetSymFromName
//sys	SymEnumerateSymbols(hProcess HANDLE, BaseOfDll uint32, EnumSymbolsCallback PSYM_ENUMSYMBOLS_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateSymbols
//sys	SymEnumerateSymbolsW(hProcess HANDLE, BaseOfDll uint32, EnumSymbolsCallback PSYM_ENUMSYMBOLS_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateSymbolsW
//sys	SymLoadModule(hProcess HANDLE, hFile HANDLE, ImageName *PSTRElement, ModuleName *PSTRElement, BaseOfDll uint32, SizeOfDll uint32) (r uint32, err error) = dbghelp.SymLoadModule
//sys	SymGetSymNext(hProcess HANDLE, Symbol *IMAGEHLP_SYMBOL) (r BOOL, err error) = dbghelp.SymGetSymNext
//sys	SymGetSymPrev(hProcess HANDLE, Symbol *IMAGEHLP_SYMBOL) (r BOOL, err error) = dbghelp.SymGetSymPrev

// Types used in generated APIs for

type CONTEXT struct {
	ContextFlags      uint32
	Dr0               uint32
	Dr1               uint32
	Dr2               uint32
	Dr3               uint32
	Dr6               uint32
	Dr7               uint32
	FloatSave         FLOATING_SAVE_AREA
	SegGs             uint32
	SegFs             uint32
	SegEs             uint32
	SegDs             uint32
	Edi               uint32
	Esi               uint32
	Ebx               uint32
	Edx               uint32
	Ecx               uint32
	Eax               uint32
	Ebp               uint32
	Eip               uint32
	SegCs             uint32
	EFlags            uint32
	Esp               uint32
	SegSs             uint32
	ExtendedRegisters []uint8
}

type LOADED_IMAGE struct {
	ModuleName       *PSTRElement
	HFile            HANDLE
	MappedAddress    *uint8
	FileHeader       *IMAGE_NT_HEADERS32
	LastRvaSection   *IMAGE_SECTION_HEADER
	NumberOfSections uint32
	Sections         *IMAGE_SECTION_HEADER
	Characteristics  IMAGE_FILE_CHARACTERISTICS2
	FSystemImage     BOOLEAN
	FDOSImage        BOOLEAN
	FReadOnly        BOOLEAN
	Version          uint8
	Links            LIST_ENTRY
	SizeOfImage      uint32
}

type STACKFRAME struct {
	AddrPC         ADDRESS
	AddrReturn     ADDRESS
	AddrFrame      ADDRESS
	AddrStack      ADDRESS
	FuncTableEntry unsafe.Pointer
	Params         []uint32
	Far            BOOL
	Virtual        BOOL
	Reserved       []uint32
	KdHelp         KDHELP
	AddrBStore     ADDRESS
}

type PREAD_PROCESS_MEMORY_ROUTINE uintptr

type PFUNCTION_TABLE_ACCESS_ROUTINE uintptr

type PGET_MODULE_BASE_ROUTINE uintptr

type PTRANSLATE_ADDRESS_ROUTINE uintptr

type PSYM_ENUMMODULES_CALLBACK uintptr

type PSYM_ENUMSYMBOLS_CALLBACK uintptr

type PSYM_ENUMSYMBOLS_CALLBACKW uintptr

type PENUMLOADED_MODULES_CALLBACK uintptr

type PSYMBOL_REGISTERED_CALLBACK uintptr

type IMAGEHLP_SYMBOL struct {
	SizeOfStruct  uint32
	Address       uint32
	Size          uint32
	Flags         uint32
	MaxNameLength uint32
	Name          []CHAR
}

type IMAGEHLP_MODULE struct {
	SizeOfStruct    uint32
	BaseOfImage     uint32
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         SYM_TYPE
	ModuleName      []CHAR
	ImageName       []CHAR
	LoadedImageName []CHAR
}

type IMAGEHLP_MODULEW struct {
	SizeOfStruct    uint32
	BaseOfImage     uint32
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         SYM_TYPE
	ModuleName      []uint16
	ImageName       []uint16
	LoadedImageName []uint16
}

type IMAGEHLP_LINE struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	LineNumber   uint32
	FileName     *PSTRElement
	Address      uint32
}

type FLOATING_SAVE_AREA struct {
	ControlWord   uint32
	StatusWord    uint32
	TagWord       uint32
	ErrorOffset   uint32
	ErrorSelector uint32
	DataOffset    uint32
	DataSelector  uint32
	RegisterArea  []uint8
	Spare0        uint32
}

type ADDRESS struct {
	Offset  uint32
	Segment uint16
	Mode    ADDRESS_MODE
}

type KDHELP struct {
	Thread                    uint32
	ThCallbackStack           uint32
	NextCallback              uint32
	FramePointer              uint32
	KiCallUserMode            uint32
	KeUserCallbackDispatcher  uint32
	SystemRangeStart          uint32
	ThCallbackBStore          uint32
	KiUserExceptionDispatcher uint32
	StackBase                 uint32
	StackLimit                uint32
	Reserved                  []uint32
}

