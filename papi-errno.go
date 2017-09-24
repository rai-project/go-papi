// Copyright (C) 2011, Los Alamos National Security, LLC.
// Use of this source code is governed by a BSD-style license.

package papi

/*
This file was generated automatically from papi.h.
*/

// #include <papi.h>
import "C"

// The following constants can be returned as Errno values from PAPI functions.
var (
	EATTR      error = Errno(C.PAPI_EATTR)      // Invalid or missing event attributes
	EBUF       error = Errno(C.PAPI_EBUF)       // Buffer size exceeded
	EBUG       error = Errno(C.PAPI_EBUG)       // Internal error, please send mail to the developers
	ECLOST     error = Errno(C.PAPI_ECLOST)     // Access to the counters was lost or interrupted
	ECMP       error = Errno(C.PAPI_ECMP)       // Not supported by component
	ECNFLCT    error = Errno(C.PAPI_ECNFLCT)    // Event exists, but cannot be counted due to counter resource limitations
	ECOMBO     error = Errno(C.PAPI_ECOMBO)     // Bad combination of features
	ECOUNT     error = Errno(C.PAPI_ECOUNT)     // Too many events or attributes
	EINVAL     error = Errno(C.PAPI_EINVAL)     // Invalid argument
	EINVAL_DOM error = Errno(C.PAPI_EINVAL_DOM) // EventSet domain is not supported for the operation
	EISRUN     error = Errno(C.PAPI_EISRUN)     // EventSet is currently counting
	EMISC      error = Errno(C.PAPI_EMISC)      // Unknown error code
	ENOCMP     error = Errno(C.PAPI_ENOCMP)     // Component Index isn't set
	ENOCNTR    error = Errno(C.PAPI_ENOCNTR)    // Hardware does not support performance counters
	ENOEVNT    error = Errno(C.PAPI_ENOEVNT)    // Event does not exist
	ENOEVST    error = Errno(C.PAPI_ENOEVST)    // No such EventSet Available
	ENOIMPL    error = Errno(C.PAPI_ENOIMPL)    // Not implemented
	ENOINIT    error = Errno(C.PAPI_ENOINIT)    // PAPI hasn't been initialized yet
	ENOMEM     error = Errno(C.PAPI_ENOMEM)     // Insufficient memory
	ENOSUPP    error = Errno(C.PAPI_ENOSUPP)    // Not supported
	ENOTPRESET error = Errno(C.PAPI_ENOTPRESET) // Event in argument is not a valid preset
	ENOTRUN    error = Errno(C.PAPI_ENOTRUN)    // EventSet is currently not running
	EPERM      error = Errno(C.PAPI_EPERM)      // Permission level does not permit operation
	ESBSTR     error = Errno(C.PAPI_ESBSTR)     // Backwards compatibility
	ESYS       error = Errno(C.PAPI_ESYS)       // A System/C library call failed
)
