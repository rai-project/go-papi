// Copyright (C) 2011, Los Alamos National Security, LLC.
// Use of this source code is governed by a BSD-style license.

package papi

/*
This file was generated automatically from papi.h.
*/

// #include <papi.h>
import "C"

// An EventModifier filters the set of events returned by EnumEvents().
const (
	ENUM_EVENTS           EventModifier = C.PAPI_ENUM_EVENTS           // Always enumerate all events
	ENUM_FIRST            EventModifier = C.PAPI_ENUM_FIRST            // Enumerate first event (preset or native)
	NTV_ENUM_DARR         EventModifier = C.PAPI_NTV_ENUM_DARR         // Enumerate events that support DAR (data address ranging)
	NTV_ENUM_DEAR         EventModifier = C.PAPI_NTV_ENUM_DEAR         // Enumerate DEAR (data event address register) events
	NTV_ENUM_GROUPS       EventModifier = C.PAPI_NTV_ENUM_GROUPS       // Enumerate groups an event belongs to (e.g. POWER5)
	NTV_ENUM_IARR         EventModifier = C.PAPI_NTV_ENUM_IARR         // Enumerate events that support IAR (instruction address ranging)
	NTV_ENUM_IEAR         EventModifier = C.PAPI_NTV_ENUM_IEAR         // Enumerate IEAR (instruction event address register) events
	NTV_ENUM_OPCM         EventModifier = C.PAPI_NTV_ENUM_OPCM         // Enumerate events that support OPC (opcode matching)
	NTV_ENUM_UMASKS       EventModifier = C.PAPI_NTV_ENUM_UMASKS       // all individual bits for given group
	NTV_ENUM_UMASK_COMBOS EventModifier = C.PAPI_NTV_ENUM_UMASK_COMBOS // all combinations of mask bits for given group
	NTV_GROUP_AND_MASK    EventModifier = C.PAPI_NTV_GROUP_AND_MASK    // bits occupied by group number
	NTV_GROUP_SHIFT       EventModifier = C.PAPI_NTV_GROUP_SHIFT       // bit shift to encode group number
	PRESET_BIT_BR         EventModifier = C.PAPI_PRESET_BIT_BR         // branch related preset events
	PRESET_BIT_CACH       EventModifier = C.PAPI_PRESET_BIT_CACH       // cache related preset events
	PRESET_BIT_CND        EventModifier = C.PAPI_PRESET_BIT_CND        // conditional preset events
	PRESET_BIT_FP         EventModifier = C.PAPI_PRESET_BIT_FP         // Floating Point related preset events
	PRESET_BIT_IDL        EventModifier = C.PAPI_PRESET_BIT_IDL        // Stalled or Idle preset event bit
	PRESET_BIT_INS        EventModifier = C.PAPI_PRESET_BIT_INS        // Instruction related preset event bit
	PRESET_BIT_L1         EventModifier = C.PAPI_PRESET_BIT_L1         // L1 cache related preset events
	PRESET_BIT_L2         EventModifier = C.PAPI_PRESET_BIT_L2         // L2 cache related preset events
	PRESET_BIT_L3         EventModifier = C.PAPI_PRESET_BIT_L3         // L3 cache related preset events
	PRESET_BIT_MEM        EventModifier = C.PAPI_PRESET_BIT_MEM        // memory related preset events
	PRESET_BIT_MSC        EventModifier = C.PAPI_PRESET_BIT_MSC        // Miscellaneous preset event bit
	PRESET_BIT_TLB        EventModifier = C.PAPI_PRESET_BIT_TLB        // Translation Lookaside Buffer events
	PRESET_ENUM_AVAIL     EventModifier = C.PAPI_PRESET_ENUM_AVAIL     // Enumerate events that exist here
)

// Map each of the above to a string.
var presetBitToString = map[EventModifier]string{
	PRESET_BIT_BR:   "PAPI_PRESET_BIT_BR",
	PRESET_BIT_CACH: "PAPI_PRESET_BIT_CACH",
	PRESET_BIT_CND:  "PAPI_PRESET_BIT_CND",
	PRESET_BIT_FP:   "PAPI_PRESET_BIT_FP",
	PRESET_BIT_IDL:  "PAPI_PRESET_BIT_IDL",
	PRESET_BIT_INS:  "PAPI_PRESET_BIT_INS",
	PRESET_BIT_L1:   "PAPI_PRESET_BIT_L1",
	PRESET_BIT_L2:   "PAPI_PRESET_BIT_L2",
	PRESET_BIT_L3:   "PAPI_PRESET_BIT_L3",
	PRESET_BIT_MEM:  "PAPI_PRESET_BIT_MEM",
	PRESET_BIT_MSC:  "PAPI_PRESET_BIT_MSC",
	PRESET_BIT_TLB:  "PAPI_PRESET_BIT_TLB"}
