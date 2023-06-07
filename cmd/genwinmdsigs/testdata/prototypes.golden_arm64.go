

// APIs for Windows.Win32.System.Diagnostics.Debug
//sys	RtlAddFunctionTable(FunctionTable *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY, EntryCount uint32, BaseAddress uintptr) (r BOOLEAN)
//sys	RtlDeleteFunctionTable(FunctionTable *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY) (r BOOLEAN)
//sys	RtlAddGrowableFunctionTable(DynamicTable *unsafe.Pointer, FunctionTable *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY, EntryCount uint32, MaximumEntryCount uint32, RangeBase uintptr, RangeEnd uintptr) (r uint32) = ntdll.RtlAddGrowableFunctionTable
//sys	RtlLookupFunctionEntry(ControlPc uintptr, ImageBase *uintptr, HistoryTable *UNWIND_HISTORY_TABLE) (r *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY)
//sys	RtlVirtualUnwind(HandlerType RTL_VIRTUAL_UNWIND_HANDLER_TYPE, ImageBase uintptr, ControlPc uintptr, FunctionEntry *IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY, ContextRecord *CONTEXT, HandlerData *unsafe.Pointer, EstablisherFrame *uintptr, ContextPointers *KNONVOLATILE_CONTEXT_POINTERS_ARM64) (r EXCEPTION_ROUTINE)
//sys	RtlInstallFunctionTableCallback(TableIdentifier uint64, BaseAddress uint64, Length uint32, Callback PGET_RUNTIME_FUNCTION_CALLBACK, Context unsafe.Pointer, OutOfProcessCallbackDll *PWSTRElement) (r BOOLEAN)
//sys	RtlGrowFunctionTable(DynamicTable unsafe.Pointer, NewEntryCount uint32) = ntdll.RtlGrowFunctionTable
//sys	RtlDeleteGrowableFunctionTable(DynamicTable unsafe.Pointer) = ntdll.RtlDeleteGrowableFunctionTable
//sys	RtlUnwindEx(TargetFrame unsafe.Pointer, TargetIp unsafe.Pointer, ExceptionRecord *EXCEPTION_RECORD, ReturnValue unsafe.Pointer, ContextRecord *CONTEXT, HistoryTable *UNWIND_HISTORY_TABLE)
//sys	CheckSumMappedFile(BaseAddress unsafe.Pointer, FileLength uint32, HeaderSum *uint32, CheckSum *uint32) (r *IMAGE_NT_HEADERS64, err error) = imagehlp.CheckSumMappedFile
//sys	GetImageConfigInformation(LoadedImage *LOADED_IMAGE, ImageConfigInformation *IMAGE_LOAD_CONFIG_DIRECTORY64) (r BOOL, err error) = imagehlp.GetImageConfigInformation
//sys	SetImageConfigInformation(LoadedImage *LOADED_IMAGE, ImageConfigInformation *IMAGE_LOAD_CONFIG_DIRECTORY64) (r BOOL, err error) = imagehlp.SetImageConfigInformation
//sys	ImageNtHeader(Base unsafe.Pointer) (r *IMAGE_NT_HEADERS64, err error) = dbghelp.ImageNtHeader
//sys	ImageRvaToSection(NtHeaders *IMAGE_NT_HEADERS64, Base unsafe.Pointer, Rva uint32) (r *IMAGE_SECTION_HEADER, err error) = dbghelp.ImageRvaToSection
//sys	ImageRvaToVa(NtHeaders *IMAGE_NT_HEADERS64, Base unsafe.Pointer, Rva uint32, LastRvaSection **IMAGE_SECTION_HEADER) (r unsafe.Pointer, err error) = dbghelp.ImageRvaToVa

// Types used in generated APIs for

type CONTEXT struct {
	ContextFlags uint32
	Cpsr         uint32
	X0           uint64
	X1           uint64
	X2           uint64
	X3           uint64
	X4           uint64
	X5           uint64
	X6           uint64
	X7           uint64
	X8           uint64
	X9           uint64
	X10          uint64
	X11          uint64
	X12          uint64
	X13          uint64
	X14          uint64
	X15          uint64
	X16          uint64
	X17          uint64
	X18          uint64
	X19          uint64
	X20          uint64
	X21          uint64
	X22          uint64
	X23          uint64
	X24          uint64
	X25          uint64
	X26          uint64
	X27          uint64
	X28          uint64
	Fp           uint64
	Lr           uint64
	Sp           uint64
	Pc           uint64
	V            [32]ARM64_NT_NEON128
	Fpcr         uint32
	Fpsr         uint32
	Bcr          [8]uint32
	Bvr          [8]uint64
	Wcr          [2]uint32
	Wvr          [2]uint64
}

type PGET_RUNTIME_FUNCTION_CALLBACK uintptr

type KNONVOLATILE_CONTEXT_POINTERS_ARM64 struct {
	X19 *uint64
	X20 *uint64
	X21 *uint64
	X22 *uint64
	X23 *uint64
	X24 *uint64
	X25 *uint64
	X26 *uint64
	X27 *uint64
	X28 *uint64
	Fp  *uint64
	Lr  *uint64
	D8  *uint64
	D9  *uint64
	D10 *uint64
	D11 *uint64
	D12 *uint64
	D13 *uint64
	D14 *uint64
	D15 *uint64
}

type UNWIND_HISTORY_TABLE struct {
	Count       uint32
	LocalHint   uint8
	GlobalHint  uint8
	Search      uint8
	Once        uint8
	LowAddress  uintptr
	HighAddress uintptr
	Entry       [12]UNWIND_HISTORY_TABLE_ENTRY
}

type LOADED_IMAGE struct {
	ModuleName       *PSTRElement
	HFile            HANDLE
	MappedAddress    *uint8
	FileHeader       *IMAGE_NT_HEADERS64
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

type XSAVE_FORMAT struct {
	ControlWord    uint16
	StatusWord     uint16
	TagWord        uint8
	Reserved1      uint8
	ErrorOpcode    uint16
	ErrorOffset    uint32
	ErrorSelector  uint16
	Reserved2      uint16
	DataOffset     uint32
	DataSelector   uint16
	Reserved3      uint16
	MxCsr          uint32
	MxCsr_Mask     uint32
	FloatRegisters [8]M128A
	XmmRegisters   [16]M128A
	Reserved4      [96]uint8
}

