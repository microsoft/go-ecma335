

// APIs for Windows.Win32.System.Diagnostics.Debug
//sys	RtlCaptureContext2(ContextRecord *CONTEXT)
//sys	RtlAddFunctionTable(FunctionTable *IMAGE_RUNTIME_FUNCTION_ENTRY, EntryCount uint32, BaseAddress uint64) (r BOOLEAN)
//sys	RtlDeleteFunctionTable(FunctionTable *IMAGE_RUNTIME_FUNCTION_ENTRY) (r BOOLEAN)
//sys	RtlInstallFunctionTableCallback(TableIdentifier uint64, BaseAddress uint64, Length uint32, Callback PGET_RUNTIME_FUNCTION_CALLBACK, Context unsafe.Pointer, OutOfProcessCallbackDll *PWSTRElement) (r BOOLEAN)
//sys	RtlAddGrowableFunctionTable(DynamicTable *unsafe.Pointer, FunctionTable *IMAGE_RUNTIME_FUNCTION_ENTRY, EntryCount uint32, MaximumEntryCount uint32, RangeBase uintptr, RangeEnd uintptr) (r uint32) = ntdll.RtlAddGrowableFunctionTable
//sys	RtlGrowFunctionTable(DynamicTable unsafe.Pointer, NewEntryCount uint32) = ntdll.RtlGrowFunctionTable
//sys	RtlDeleteGrowableFunctionTable(DynamicTable unsafe.Pointer) = ntdll.RtlDeleteGrowableFunctionTable
//sys	RtlLookupFunctionEntry(ControlPc uint64, ImageBase *uint64, HistoryTable *UNWIND_HISTORY_TABLE) (r *IMAGE_RUNTIME_FUNCTION_ENTRY)
//sys	RtlUnwindEx(TargetFrame unsafe.Pointer, TargetIp unsafe.Pointer, ExceptionRecord *EXCEPTION_RECORD, ReturnValue unsafe.Pointer, ContextRecord *CONTEXT, HistoryTable *UNWIND_HISTORY_TABLE)
//sys	RtlVirtualUnwind(HandlerType RTL_VIRTUAL_UNWIND_HANDLER_TYPE, ImageBase uint64, ControlPc uint64, FunctionEntry *IMAGE_RUNTIME_FUNCTION_ENTRY, ContextRecord *CONTEXT, HandlerData *unsafe.Pointer, EstablisherFrame *uint64, ContextPointers *KNONVOLATILE_CONTEXT_POINTERS) (r EXCEPTION_ROUTINE)
//sys	CheckSumMappedFile(BaseAddress unsafe.Pointer, FileLength uint32, HeaderSum *uint32, CheckSum *uint32) (r *IMAGE_NT_HEADERS64, err error) = imagehlp.CheckSumMappedFile
//sys	GetImageConfigInformation(LoadedImage *LOADED_IMAGE, ImageConfigInformation *IMAGE_LOAD_CONFIG_DIRECTORY64) (r BOOL, err error) = imagehlp.GetImageConfigInformation
//sys	SetImageConfigInformation(LoadedImage *LOADED_IMAGE, ImageConfigInformation *IMAGE_LOAD_CONFIG_DIRECTORY64) (r BOOL, err error) = imagehlp.SetImageConfigInformation
//sys	ImageNtHeader(Base unsafe.Pointer) (r *IMAGE_NT_HEADERS64, err error) = dbghelp.ImageNtHeader
//sys	ImageRvaToSection(NtHeaders *IMAGE_NT_HEADERS64, Base unsafe.Pointer, Rva uint32) (r *IMAGE_SECTION_HEADER, err error) = dbghelp.ImageRvaToSection
//sys	ImageRvaToVa(NtHeaders *IMAGE_NT_HEADERS64, Base unsafe.Pointer, Rva uint32, LastRvaSection **IMAGE_SECTION_HEADER) (r unsafe.Pointer, err error) = dbghelp.ImageRvaToVa
//sys	GetEnabledXStateFeatures() (r uint64)
//sys	GetXStateFeaturesMask(Context *CONTEXT, FeatureMask *uint64) (r BOOL)
//sys	LocateXStateFeature(Context *CONTEXT, FeatureId uint32, Length *uint32) (r unsafe.Pointer)
//sys	SetXStateFeaturesMask(Context *CONTEXT, FeatureMask uint64) (r BOOL)

// Types used in generated APIs for

type CONTEXT struct {
	P1Home               uint64
	P2Home               uint64
	P3Home               uint64
	P4Home               uint64
	P5Home               uint64
	P6Home               uint64
	ContextFlags         uint32
	MxCsr                uint32
	SegCs                uint16
	SegDs                uint16
	SegEs                uint16
	SegFs                uint16
	SegGs                uint16
	SegSs                uint16
	EFlags               uint32
	Dr0                  uint64
	Dr1                  uint64
	Dr2                  uint64
	Dr3                  uint64
	Dr6                  uint64
	Dr7                  uint64
	Rax                  uint64
	Rcx                  uint64
	Rdx                  uint64
	Rbx                  uint64
	Rsp                  uint64
	Rbp                  uint64
	Rsi                  uint64
	Rdi                  uint64
	R8                   uint64
	R9                   uint64
	R10                  uint64
	R11                  uint64
	R12                  uint64
	R13                  uint64
	R14                  uint64
	R15                  uint64
	Rip                  uint64
	FltSave              XSAVE_FORMAT
	VectorRegister       [26]M128A
	VectorControl        uint64
	DebugControl         uint64
	LastBranchToRip      uint64
	LastBranchFromRip    uint64
	LastExceptionToRip   uint64
	LastExceptionFromRip uint64
}

type PGET_RUNTIME_FUNCTION_CALLBACK uintptr

type KNONVOLATILE_CONTEXT_POINTERS struct {
	FloatingContext [16]*M128A
	IntegerContext  [16]*uint64
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

type UNWIND_HISTORY_TABLE_ENTRY struct {
	ImageBase     uintptr
	FunctionEntry *IMAGE_RUNTIME_FUNCTION_ENTRY
}

