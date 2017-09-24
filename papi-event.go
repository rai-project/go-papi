// Copyright (C) 2011, Los Alamos National Security, LLC.
// Use of this source code is governed by a BSD-style license.

package papi

/*
This file was generated automatically from papiStdEventDefs.h.
*/

// #include <papi.h>
import "C"

// The following constants represent PAPI's standard event types.
const (
	BRU_IDL Event = C.PAPI_BRU_IDL // Cycles branch units are idle
	BR_CN   Event = C.PAPI_BR_CN   // Conditional branch instructions executed
	BR_INS  Event = C.PAPI_BR_INS  // Total branch instructions executed
	BR_MSP  Event = C.PAPI_BR_MSP  // Conditional branch instructions mispred
	BR_NTK  Event = C.PAPI_BR_NTK  // Conditional branch instructions not taken
	BR_PRC  Event = C.PAPI_BR_PRC  // Conditional branch instructions corr. pred
	BR_TKN  Event = C.PAPI_BR_TKN  // Conditional branch instructions taken
	BR_UCN  Event = C.PAPI_BR_UCN  // Unconditional branch instructions executed
	BTAC_M  Event = C.PAPI_BTAC_M  // BTAC miss
	CA_CLN  Event = C.PAPI_CA_CLN  // Request for clean cache line (SMP)
	CA_INV  Event = C.PAPI_CA_INV  // Request for cache line Invalidation (SMP)
	CA_ITV  Event = C.PAPI_CA_ITV  // Request for cache line Intervention (SMP)
	CA_SHR  Event = C.PAPI_CA_SHR  // Request for shared cache line (SMP)
	CA_SNP  Event = C.PAPI_CA_SNP  // Snoops
	CSR_FAL Event = C.PAPI_CSR_FAL // Failed store conditional instructions
	CSR_SUC Event = C.PAPI_CSR_SUC // Successful store conditional instructions
	CSR_TOT Event = C.PAPI_CSR_TOT // Total store conditional instructions
	DP_OPS  Event = C.PAPI_DP_OPS  // Floating point operations executed; optimized to count scaled double precision vector operations
	FAD_INS Event = C.PAPI_FAD_INS // FA ins
	FDV_INS Event = C.PAPI_FDV_INS // FD ins
	FMA_INS Event = C.PAPI_FMA_INS // FMA instructions completed
	FML_INS Event = C.PAPI_FML_INS // FM ins
	FNV_INS Event = C.PAPI_FNV_INS // Finv ins
	FPU_IDL Event = C.PAPI_FPU_IDL // Cycles floating point units are idle
	FP_INS  Event = C.PAPI_FP_INS  // Floating point instructions executed
	FP_OPS  Event = C.PAPI_FP_OPS  // Floating point operations executed
	FP_STAL Event = C.PAPI_FP_STAL // Cycles any FP units are stalled
	FSQ_INS Event = C.PAPI_FSQ_INS // FSq ins
	FUL_CCY Event = C.PAPI_FUL_CCY // Cycles with Maximum Instruction Completion
	FUL_ICY Event = C.PAPI_FUL_ICY // Cycles with Maximum Instruction Issue
	FXU_IDL Event = C.PAPI_FXU_IDL // Cycles integer units are idle
	HW_INT  Event = C.PAPI_HW_INT  // Hardware interrupts
	INT_INS Event = C.PAPI_INT_INS // Integer instructions executed
	L1_DCA  Event = C.PAPI_L1_DCA  // L1 D Cache Access
	L1_DCH  Event = C.PAPI_L1_DCH  // L1 D Cache Hit
	L1_DCM  Event = C.PAPI_L1_DCM  // Level 1 data cache misses
	L1_DCR  Event = C.PAPI_L1_DCR  // L1 D Cache Read
	L1_DCW  Event = C.PAPI_L1_DCW  // L1 D Cache Write
	L1_ICA  Event = C.PAPI_L1_ICA  // L1 instruction cache accesses
	L1_ICH  Event = C.PAPI_L1_ICH  // L1 instruction cache hits
	L1_ICM  Event = C.PAPI_L1_ICM  // Level 1 instruction cache misses
	L1_ICR  Event = C.PAPI_L1_ICR  // L1 instruction cache reads
	L1_ICW  Event = C.PAPI_L1_ICW  // L1 instruction cache writes
	L1_LDM  Event = C.PAPI_L1_LDM  // Level 1 load misses
	L1_STM  Event = C.PAPI_L1_STM  // Level 1 store misses
	L1_TCA  Event = C.PAPI_L1_TCA  // L1 total cache accesses
	L1_TCH  Event = C.PAPI_L1_TCH  // L1 total cache hits
	L1_TCM  Event = C.PAPI_L1_TCM  // Level 1 total cache misses
	L1_TCR  Event = C.PAPI_L1_TCR  // L1 total cache reads
	L1_TCW  Event = C.PAPI_L1_TCW  // L1 total cache writes
	L2_DCA  Event = C.PAPI_L2_DCA  // L2 D Cache Access
	L2_DCH  Event = C.PAPI_L2_DCH  // L2 D Cache Hit
	L2_DCM  Event = C.PAPI_L2_DCM  // Level 2 data cache misses
	L2_DCR  Event = C.PAPI_L2_DCR  // L2 D Cache Read
	L2_DCW  Event = C.PAPI_L2_DCW  // L2 D Cache Write
	L2_ICA  Event = C.PAPI_L2_ICA  // L2 instruction cache accesses
	L2_ICH  Event = C.PAPI_L2_ICH  // L2 instruction cache hits
	L2_ICM  Event = C.PAPI_L2_ICM  // Level 2 instruction cache misses
	L2_ICR  Event = C.PAPI_L2_ICR  // L2 instruction cache reads
	L2_ICW  Event = C.PAPI_L2_ICW  // L2 instruction cache writes
	L2_LDM  Event = C.PAPI_L2_LDM  // Level 2 load misses
	L2_STM  Event = C.PAPI_L2_STM  // Level 2 store misses
	L2_TCA  Event = C.PAPI_L2_TCA  // L2 total cache accesses
	L2_TCH  Event = C.PAPI_L2_TCH  // L2 total cache hits
	L2_TCM  Event = C.PAPI_L2_TCM  // Level 2 total cache misses
	L2_TCR  Event = C.PAPI_L2_TCR  // L2 total cache reads
	L2_TCW  Event = C.PAPI_L2_TCW  // L2 total cache writes
	L3_DCA  Event = C.PAPI_L3_DCA  // L3 D Cache Access
	L3_DCH  Event = C.PAPI_L3_DCH  // Level 3 Data Cache Hit
	L3_DCM  Event = C.PAPI_L3_DCM  // Level 3 data cache misses
	L3_DCR  Event = C.PAPI_L3_DCR  // L3 D Cache Read
	L3_DCW  Event = C.PAPI_L3_DCW  // L3 D Cache Write
	L3_ICA  Event = C.PAPI_L3_ICA  // L3 instruction cache accesses
	L3_ICH  Event = C.PAPI_L3_ICH  // L3 instruction cache hits
	L3_ICM  Event = C.PAPI_L3_ICM  // Level 3 instruction cache misses
	L3_ICR  Event = C.PAPI_L3_ICR  // L3 instruction cache reads
	L3_ICW  Event = C.PAPI_L3_ICW  // L3 instruction cache writes
	L3_LDM  Event = C.PAPI_L3_LDM  // Level 3 load misses
	L3_STM  Event = C.PAPI_L3_STM  // Level 3 store misses
	L3_TCA  Event = C.PAPI_L3_TCA  // L3 total cache accesses
	L3_TCH  Event = C.PAPI_L3_TCH  // L3 total cache hits
	L3_TCM  Event = C.PAPI_L3_TCM  // Level 3 total cache misses
	L3_TCR  Event = C.PAPI_L3_TCR  // L3 total cache reads
	L3_TCW  Event = C.PAPI_L3_TCW  // L3 total cache writes
	LD_INS  Event = C.PAPI_LD_INS  // Load instructions executed
	LST_INS Event = C.PAPI_LST_INS // Total load/store inst. executed
	LSU_IDL Event = C.PAPI_LSU_IDL // Cycles load/store units are idle
	MEM_RCY Event = C.PAPI_MEM_RCY // Cycles Stalled Waiting for Memory Read
	MEM_SCY Event = C.PAPI_MEM_SCY // Cycles Stalled Waiting for Memory Access
	MEM_WCY Event = C.PAPI_MEM_WCY // Cycles Stalled Waiting for Memory Write
	PRF_DM  Event = C.PAPI_PRF_DM  // Prefetch data instruction caused a miss
	REF_CYC Event = C.PAPI_REF_CYC // Reference clock cycles
	RES_STL Event = C.PAPI_RES_STL // Cycles processor is stalled on resource
	SP_OPS  Event = C.PAPI_SP_OPS  // Floating point operations executed; optimized to count scaled single precision vector operations
	SR_INS  Event = C.PAPI_SR_INS  // Store instructions executed
	STL_CCY Event = C.PAPI_STL_CCY // Cycles with No Instruction Completion
	STL_ICY Event = C.PAPI_STL_ICY // Cycles with No Instruction Issue
	SYC_INS Event = C.PAPI_SYC_INS // Sync. inst. executed
	TLB_DM  Event = C.PAPI_TLB_DM  // Data translation lookaside buffer misses
	TLB_IM  Event = C.PAPI_TLB_IM  // Instr translation lookaside buffer misses
	TLB_SD  Event = C.PAPI_TLB_SD  // Xlation lookaside buffer shootdowns (SMP)
	TLB_TL  Event = C.PAPI_TLB_TL  // Total translation lookaside buffer misses
	TOT_CYC Event = C.PAPI_TOT_CYC // Total cycles executed
	TOT_IIS Event = C.PAPI_TOT_IIS // Total instructions issued
	TOT_INS Event = C.PAPI_TOT_INS // Total instructions executed
	VEC_DP  Event = C.PAPI_VEC_DP  // Double precision vector/SIMD instructions
	VEC_INS Event = C.PAPI_VEC_INS // Vector/SIMD instructions executed (could include integer)
	VEC_SP  Event = C.PAPI_VEC_SP  // Single precision vector/SIMD instructions
)
