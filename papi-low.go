// Copyright (C) 2011, Los Alamos National Security, LLC.
// Use of this source code is governed by a BSD-style license.

// This file provides an interface to PAPI's low-level functions.

package papi

/*
#include <stdio.h>
#include <papi.h>

typedef struct _papi_os_option {
	char name[PAPI_MAX_STR_LEN];
    char version[PAPI_MAX_STR_LEN];
    int os_version;
    int itimer_sig;
    int itimer_num;
    int itimer_ns;
    int itimer_res_ns;
    int clock_ticks;
    unsigned long reserved[8];       // For future expansion
} PAPI_os_info_t;
extern PAPI_os_info_t _papi_os_info;

// As of this writing, cgo doesn't seem to support bit fields.  We
// therefore have to use a wrapper function to access the bits in a
// PAPI_component_info_t.
void get_component_bits(PAPI_component_info_t *info, int *bitfields)
{
  int i = 0;
  bitfields[i++] = info->hardware_intr;
  bitfields[i++] = info->precise_intr;
  bitfields[i++] = info->posix1b_timers;
  bitfields[i++] = info->kernel_profile;
  bitfields[i++] = info->kernel_multiplex;
//	 bitfields[i++] = info->data_address_range;
i++;
//   bitfields[i++] = info->instr_address_range;
i++;
  bitfields[i++] = info->fast_counter_read;
  bitfields[i++] = info->fast_real_timer;
  bitfields[i++] = info->fast_virtual_timer;
  bitfields[i++] = info->attach;
  bitfields[i++] = info->attach_must_ptrace;
  bitfields[i++] = info->cpu;
  bitfields[i++] = info->inherit;
//   bitfields[i++] = info->edge_detect;
i++;
//   bitfields[i++] = info->invert;
i++;
//   bitfields[i++] = info->profile_ear;
i++;
//   bitfields[i++] = info->cntr_groups;
i++;
  bitfields[i++] = info->cntr_umasks;
//   bitfields[i++] = info->cntr_IEAR_events;
i++;
//   bitfields[i++] = info->cntr_DEAR_events;
i++;
//   bitfields[i++] = info->cntr_OPCM_events;
i++;
}

*/
import "C"
import "unsafe"

// Return the real-time counter's value in clock cycles.
func GetRealCyc() int64 {
	return int64(C.PAPI_get_real_cyc())
}

// Return the real-time counter's value in microseconds.
func GetRealUsec() int64 {
	return int64(C.PAPI_get_real_usec())
}

// Return the virtual-time counter's value in clock cycles.
func GetVirtCyc() int64 {
	return int64(C.PAPI_get_virt_cyc())
}

// Return the virtual-time counter's value in microseconds.
func GetVirtUsec() int64 {
	return int64(C.PAPI_get_virt_usec())
}

// ----------------------------------------------------------------------

// Return the executable's address-space information.
func GetExecutableInfo() ProgramInfo {
	cinfo := C.PAPI_get_executable_info()
	if cinfo == nil {
		// I can't imagine this ever happening, but we should
		// do something just in case.
		panic("PAPI_get_executable_info() failed unexpectedly")
	}
	addrInfo := cinfo.address_info
	return ProgramInfo{
		FullName: C.GoString(&cinfo.fullname[0]),
		AddressInfo: AddressMap{
			Name:      C.GoString(&addrInfo.name[0]),
			TextStart: uintptr(unsafe.Pointer(addrInfo.text_start)),
			TextEnd:   uintptr(unsafe.Pointer(addrInfo.text_end)),
			DataStart: uintptr(unsafe.Pointer(addrInfo.data_start)),
			DataEnd:   uintptr(unsafe.Pointer(addrInfo.data_end)),
			BssStart:  uintptr(unsafe.Pointer(addrInfo.bss_start)),
			BssEnd:    uintptr(unsafe.Pointer(addrInfo.bss_end))}}
}

// Acquire and return all sorts of information about the underlying hardware.
func GetHardwareInfo() HardwareInfo {
	hw := C.PAPI_get_hardware_info()
	maxLevels := int(C.PAPI_MH_MAX_LEVELS)

	// Describe all levels of the memory hierarchy.
	mh := make([]MHLevelInfo, hw.mem_hierarchy.levels)
	for level := range mh {
		cLevel := hw.mem_hierarchy.level[level]

		// Populate the TLB information.
		tlbData := make([]TLBInfo, maxLevels)
		var validTLBLevels int
		for i := range tlbData {
			ctlb := cLevel.tlb[i]
			tlbData[i].Type = MHAttrs(ctlb._type)
			if tlbData[i].Type == MH_TYPE_EMPTY {
				break
			}
			tlbData[i].NumEntries = int32(ctlb.num_entries)
			tlbData[i].PageSize = int32(ctlb.page_size)
			tlbData[i].Associativity = int32(ctlb.associativity)
			validTLBLevels++
		}
		mh[level].TLB = tlbData[0:validTLBLevels]

		// Populate the cache information.
		cacheData := make([]CacheInfo, maxLevels)
		var validCacheLevels int
		for i := range cacheData {
			ccache := cLevel.cache[i]
			cacheData[i].Type = MHAttrs(ccache._type)
			if cacheData[i].Type == MH_TYPE_EMPTY {
				break
			}
			cacheData[i].Size = int32(ccache.size)
			cacheData[i].LineSize = int32(ccache.line_size)
			cacheData[i].NumLines = int32(ccache.num_lines)
			cacheData[i].Associativity = int32(ccache.associativity)
			validCacheLevels++
		}
		mh[level].Cache = cacheData[0:validCacheLevels]
	}

	// Populate and return the set of available hardware information.
	return HardwareInfo{
		CPUs:          int32(hw.ncpu),
		Threads:       int32(hw.threads),
		Cores:         int32(hw.cores),
		Sockets:       int32(hw.sockets),
		NUMANodes:     int32(hw.nnodes),
		TotalCPUs:     int32(hw.totalcpus),
		Vendor:        int32(hw.vendor),
		VendorName:    C.GoString(&hw.vendor_string[0]),
		Model:         int32(hw.model),
		ModelName:     C.GoString(&hw.model_string[0]),
		Revision:      float32(hw.revision),
		CPUIDFamily:   int32(hw.cpuid_family),
		CPUIDModel:    int32(hw.cpuid_model),
		CPUIDStepping: int32(hw.cpuid_stepping),
		MHz:           float32(hw.mhz),
		ClockMHz:      int32(hw.clock_mhz),
		MemHierarchy:  mh}
}

// Acquire and return all sorts of information about the current
// process's dynamic memory usage.  In addition to returning an
// overall error code, GetDynMemInfo() can also return an Errno cast
// to an int64 for any individual field.  To check for that case, note
// that all errors are represented as negative values.
func GetDynMemInfo() (dmem DynMemInfo, err error) {
	var c_dmem C.PAPI_dmem_info_t
	if errno := Errno(C.PAPI_get_dmem_info(&c_dmem)); errno != papi_ok {
		err = errno
		return
	}
	dmem = DynMemInfo{
		Peak:          int64(c_dmem.peak),
		Size:          int64(c_dmem.size),
		Resident:      int64(c_dmem.resident),
		HighWaterMark: int64(c_dmem.high_water_mark),
		Shared:        int64(c_dmem.shared),
		Text:          int64(c_dmem.text),
		Library:       int64(c_dmem.library),
		Heap:          int64(c_dmem.heap),
		Locked:        int64(c_dmem.locked),
		Stack:         int64(c_dmem.stack),
		PageSize:      int64(c_dmem.pagesize),
		PTE:           int64(c_dmem.pte)}
	return
}

// ----------------------------------------------------------------------

// Allocate a new event set and return a handler to it.
func CreateEventSet() (es EventSet, err error) {
	es = C.PAPI_NULL
	if errno := Errno(C.PAPI_create_eventset((*C.int)(&es))); errno != papi_ok {
		err = errno
	}
	return
}

// Add an event to an event set.
func (es EventSet) AddEvent(ecode Event) (err error) {
	if errno := Errno(C.PAPI_add_event(C.int(es), C.int(ecode))); errno != papi_ok {
		err = errno
	}
	return
}

// Add multiple events to an event set.
func (es EventSet) AddEvents(ecodes []Event) (err error) {
	if errno := Errno(C.PAPI_add_events(C.int(es), (*C.int)(&ecodes[0]), C.int(len(ecodes)))); errno != papi_ok {
		err = errno
	}
	return
}

// Return the number of events in an event set.
func (es EventSet) NumEvents() (numEvents int, err error) {
	if cNumEvents := C.PAPI_num_events(C.int(es)); cNumEvents >= 0 {
		numEvents = int(cNumEvents)
	} else {
		err = Errno(cNumEvents)
	}
	return
}

// Start counting every event in an event set.
func (es EventSet) Start() (err error) {
	if errno := Errno(C.PAPI_start(C.int(es))); errno != papi_ok {
		err = errno
	}
	return
}

// Stop counting events and return the final counter values.
func (es EventSet) Stop(values []int64) error {
	numEvents, err := es.NumEvents()
	if err != nil {
		return err
	}
	if len(values) < numEvents {
		return EBUF
	}
	if errno := Errno(C.PAPI_stop(C.int(es), (*C.longlong)(&values[0]))); errno != papi_ok {
		return errno
	}
	return nil
}

// Remove an event from an event set.
func (es EventSet) RemoveEvent(ecode Event) (err error) {
	if errno := Errno(C.PAPI_remove_event(C.int(es), C.int(ecode))); errno != papi_ok {
		err = errno
	}
	return
}

// Remove multiple events from an event set.
func (es EventSet) RemoveEvents(ecodes []Event) (err error) {
	if errno := Errno(C.PAPI_remove_events(C.int(es), (*C.int)(&ecodes[0]), C.int(len(ecodes)))); errno != papi_ok {
		err = errno
	}
	return
}

// Remove all events from an event set and stop counting events in the
// event set.  CleanupEventSet() can not be called if the event set
// has not been stopped.
func (es EventSet) CleanupEventSet() (err error) {
	if errno := Errno(C.PAPI_cleanup_eventset(C.int(es))); errno != papi_ok {
		err = errno
	}
	return
}

// Deallocate the memory associated with an empty event set.
func (es *EventSet) DestroyEventSet() (err error) {
	if errno := Errno(C.PAPI_destroy_eventset((*C.int)(es))); errno != papi_ok {
		err = errno
	}
	return
}

// Return a slice of all of the events in an event set.
func (es EventSet) ListEvents() (ecodes []Event, err error) {
	var numEvents int
	if numEvents, err = es.NumEvents(); err != nil {
		return
	}
	c_ecodes := make([]C.int, numEvents)
	c_num_events := C.int(numEvents)
	if errno := Errno(C.PAPI_list_events(C.int(es), &c_ecodes[0], &c_num_events)); errno != papi_ok {
		err = errno
		return
	}
	ecodes = make([]Event, c_num_events)
	for i, ev := range c_ecodes {
		ecodes[i] = Event(ev)
	}
	return
}

// Say whether an event set is multiplexed (allows more counters than
// what the underlying hardware supports).
func (es EventSet) GetMultiplex() (isMplexed bool, err error) {
	if retval := C.PAPI_get_multiplex(C.int(es)); Errno(retval) != papi_ok {
		err = Errno(retval)
	} else {
		isMplexed = (retval != 0)
	}
	return
}

// Convert an ordinary event set into a multiplexed event set,
// enabling it to handle more counters than what the underlying
// hardware supports by timesharing counters.  SetMultiplex() must be
// called after MultiplexInit() but before Start().
func (es EventSet) SetMultiplex() (err error) {
	if errno := Errno(C.PAPI_set_multiplex(C.int(es))); errno != papi_ok {
		err = errno
	}
	return
}

// Assign a component index to an event set.  Event sets are
// ordinarily automatically bound to components when the first event
// is added.  This function is useful to explicitly bind an event set
// to a component before setting component related options (e.g., via
// SetMultiplex()).
func (es EventSet) AssignComponent(idx int) (err error) {
	if errno := Errno(C.PAPI_assign_eventset_component(C.int(es), C.int(idx))); errno != papi_ok {
		err = errno
	}
	return
}

// ----------------------------------------------------------------------

// Enumerate PAPI preset or native events.  The corresponding C
// interface, PAPI_enum_event(), returns a single event at a time.
// For convenience, we return a slice of all events.
func EnumEvents(emask EventMask, modifier EventModifier) (matches []Event, err error) {
	c_event := C.int(emask)
	c_mod := C.int(modifier)
	matches = make([]Event, 0)
	var errno Errno

	// Store the complete list of events in a Vector.
	for errno = Errno(C.PAPI_enum_event(&c_event, C.int(ENUM_FIRST))); errno == papi_ok; errno = Errno(C.PAPI_enum_event(&c_event, c_mod)) {
		matches = append(matches, Event(c_event))
	}
	if errno != ENOEVNT && errno != ESBSTR {
		matches = nil
		err = errno
		return
	}
	return
}

// Return descriptive information about an event.
func GetEventInfo(ev Event) (info EventInfo, err error) {
	var c_info C.PAPI_event_info_t
	if errno := Errno(C.PAPI_get_event_info(C.int(ev), &c_info)); errno != papi_ok {
		err = errno
		return
	}
	code := make([]uint32, c_info.count)
	name := make([]string, c_info.count)
	for i := 0; i < int(c_info.count); i++ {
		code[i] = uint32(c_info.code[i])
		name[i] = C.GoString(&c_info.name[i][0])
	}
	info = EventInfo{
		EventCode:  Event(c_info.event_code),
		EventType:  EventModifier(c_info.event_type),
		Symbol:     C.GoString(&c_info.symbol[0]),
		ShortDescr: C.GoString(&c_info.short_descr[0]),
		LongDescr:  C.GoString(&c_info.long_descr[0]),
		Derived:    C.GoString(&c_info.derived[0]),
		Postfix:    C.GoString(&c_info.postfix[0]),
		Code:       code,
		Name:       name,
		Note:       C.GoString(&c_info.note[0])}
	return
}

// ----------------------------------------------------------------------

// Return the number of counting components included in the PAPI
// library.
func GetNumComponents() int {
	return int(C.PAPI_num_components())
}

// Return the number of counters present in the specified component.
// By convention, component 0 is the CPU.
func GetNumCounters(idx int) int {
	return int(C.PAPI_num_cmp_hwctrs(C.int(idx)))
}

// Return information about the nth PAPI component.  By convention,
// component 0 is the CPU.
func GetComponentInfo(idx int) (info ComponentInfo, err error) {
	c_info := C.PAPI_get_component_info(C.int(idx))
	if c_info == nil {
		err = ENOCMP
		return
	}
	_papi_os_info := C._papi_os_info
	bitfields := make([]C.int, 22)
	C.get_component_bits(c_info, &bitfields[0])
	info = ComponentInfo{
		Name:                   C.GoString(&c_info.name[0]),
		Version:                C.GoString(&c_info.version[0]),
		SupportVersion:         C.GoString(&c_info.support_version[0]),
		KernelVersion:          C.GoString(&c_info.kernel_version[0]),
		CmpIdx:                 int(c_info.CmpIdx),
		NumCntrs:               int(c_info.num_cntrs),
		NumMpxCntrs:            int(c_info.num_mpx_cntrs),
		NumPresetEvents:        int(c_info.num_preset_events),
		NumNativeEvents:        int(c_info.num_native_events),
		DefaultDomain:          int(c_info.default_domain),
		AvailableDomains:       int(c_info.available_domains),
		DefaultGranularity:     int(c_info.default_granularity),
		AvailableGranularities: int(c_info.available_granularities),
		ItimerSig:              int(_papi_os_info.itimer_sig),
		ItimerNum:              int(_papi_os_info.itimer_num),
		ItimerNs:               int(_papi_os_info.itimer_ns),
		ItimerResNs:            int(_papi_os_info.itimer_res_ns),
		HardwareIntrSig:        int(c_info.hardware_intr_sig),
		ClockTicks:             int(_papi_os_info.clock_ticks),
		// OpcodeMatchWidth:       int(c_info.opcode_match_width),
		OSVersion:       int(_papi_os_info.os_version),
		HardwareIntr:    bitfields[0] != 0,
		PreciseIntr:     bitfields[1] != 0,
		POSIX1bTimers:   bitfields[2] != 0,
		KernelProfile:   bitfields[3] != 0,
		KernelMultiplex: bitfields[4] != 0,
		// DataAddressRange:       bitfields[5] != 0,
		// InstrAddressRange:      bitfields[6] != 0,
		FastCounterRead:  bitfields[7] != 0,
		FastRealTimer:    bitfields[8] != 0,
		FastVirtualTimer: bitfields[9] != 0,
		Attach:           bitfields[10] != 0,
		AttachMustPtrace: bitfields[11] != 0,
		CPU:              bitfields[12] != 0,
		Inherit:          bitfields[13] != 0,
		// EdgeDetect:             bitfields[14] != 0,
		// Invert:         bitfields[15] != 0,
		// ProfileEAR:     bitfields[16] != 0,
		// CntrGroups:     bitfields[17] != 0,
		CntrUmasks: bitfields[18] != 0,
		// CntrIEAREvents: bitfields[19] != 0,
		// CntrDEAREvents: bitfields[20] != 0,
		// CntrOPCMEvents: bitfields[21] != 0,
	}
	return
}
