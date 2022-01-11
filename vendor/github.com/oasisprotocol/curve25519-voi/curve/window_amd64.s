// Code generated by command: go run window.go. DO NOT EDIT.

// +build amd64,!purego,!forcenoasm,!force32bit

#include "textflag.h"

// func lookupAffineNiels(table *affineNielsPointLookupTable, out *affineNielsPoint, xabs uint8)
// Requires: SSE2
TEXT ·lookupAffineNiels(SB), NOSPLIT|NOFRAME, $0-24
	// This is moderately annoying due to having 3x5 64-bit elements,
	// which does not nicely fit into vector registers.  This is
	// handled by duplicating one element in 2 registers, since
	// doing so keeps the rest of the code straight forward.
	// 
	// v0 = y_plus_x[0],  y_plus_x[1]
	// v1 = y_plus_x[2],  y_plus_x[3]
	// v2 = y_plus_x[4],  y_minus_x[0]
	// v3 = y_minus_x[1], y_minus_x[2]
	// v4 = y_minus_x[3], y_minus_x[4]
	// v5 = xy2d[0],      xy2d[1]
	// v6 = xy2d[1] (*),  xy2d[2]
	// v7 = xy2d[3],      xy2d[4]
	// 
	// Note: Before I get tempted to rewrite this to use AVX2 again
	// I will to take a moment to remind myself that the AVX2 backend
	// does not use this table.

	MOVQ table+0(FP), AX

	// Build the mask, zero all the registers
	MOVBQZX xabs+16(FP), CX
	MOVD    CX, X0
	PSHUFD  $0x00, X0, X0
	PXOR    X2, X2
	PXOR    X3, X3
	PXOR    X4, X4
	PXOR    X5, X5
	PXOR    X6, X6
	PXOR    X7, X7
	PXOR    X8, X8
	PXOR    X9, X9

	// 0: Identity element (1, 1, 0)
	PXOR       X1, X1
	PCMPEQL    X0, X1
	MOVQ       $0x0000000000000001, CX
	MOVQ       CX, X10
	PXOR       X11, X11
	PUNPCKLQDQ X10, X11
	PAND       X1, X10
	PAND       X1, X11
	POR        X10, X2
	POR        X11, X4

	// 1 .. 8
	MOVQ $0x0000000000000001, CX

affine_lookup_loop:
	MOVD    CX, X1
	PSHUFD  $0x00, X1, X1
	PCMPEQL X0, X1
	MOVOU   (AX), X10
	MOVOU   16(AX), X11
	MOVOU   32(AX), X12
	MOVOU   48(AX), X13
	PAND    X1, X10
	PAND    X1, X11
	PAND    X1, X12
	PAND    X1, X13
	POR     X10, X2
	POR     X11, X3
	POR     X12, X4
	POR     X13, X5
	MOVOU   64(AX), X10
	MOVOU   80(AX), X11
	MOVOU   88(AX), X12
	MOVOU   104(AX), X13
	PAND    X1, X10
	PAND    X1, X11
	PAND    X1, X12
	PAND    X1, X13
	POR     X10, X6
	POR     X11, X7
	POR     X12, X8
	POR     X13, X9
	ADDQ    $0x78, AX
	INCQ    CX
	CMPQ    CX, $0x08
	JLE     affine_lookup_loop

	// Write out the result
	MOVQ  out+8(FP), AX
	MOVOU X2, (AX)
	MOVOU X3, 16(AX)
	MOVOU X4, 32(AX)
	MOVOU X5, 48(AX)
	MOVOU X6, 64(AX)
	MOVOU X7, 80(AX)
	MOVOU X8, 88(AX)
	MOVOU X9, 104(AX)
	RET

DATA cached_id_0<>+0(SB)/4, $0x0001db2f
DATA cached_id_0<>+4(SB)/4, $0x0001db42
DATA cached_id_0<>+8(SB)/4, $0x00000000
DATA cached_id_0<>+12(SB)/4, $0x00000000
DATA cached_id_0<>+16(SB)/4, $0x0003b684
DATA cached_id_0<>+20(SB)/4, $0x03ffffed
DATA cached_id_0<>+24(SB)/4, $0x00000000
DATA cached_id_0<>+28(SB)/4, $0x01ffffff
GLOBL cached_id_0<>(SB), RODATA|NOPTR, $32

DATA cached_id_1<>+0(SB)/4, $0x04000000
DATA cached_id_1<>+4(SB)/4, $0x00000000
DATA cached_id_1<>+8(SB)/4, $0x01ffffff
DATA cached_id_1<>+12(SB)/4, $0x00000000
DATA cached_id_1<>+16(SB)/4, $0x00000000
DATA cached_id_1<>+20(SB)/4, $0x03ffffff
DATA cached_id_1<>+24(SB)/4, $0x00000000
DATA cached_id_1<>+28(SB)/4, $0x01ffffff
GLOBL cached_id_1<>(SB), RODATA|NOPTR, $32

DATA cached_id_2_4<>+0(SB)/4, $0x03ffffff
DATA cached_id_2_4<>+4(SB)/4, $0x00000000
DATA cached_id_2_4<>+8(SB)/4, $0x01ffffff
DATA cached_id_2_4<>+12(SB)/4, $0x00000000
DATA cached_id_2_4<>+16(SB)/4, $0x00000000
DATA cached_id_2_4<>+20(SB)/4, $0x03ffffff
DATA cached_id_2_4<>+24(SB)/4, $0x00000000
DATA cached_id_2_4<>+28(SB)/4, $0x01ffffff
GLOBL cached_id_2_4<>(SB), RODATA|NOPTR, $32

// func lookupCached(table *cachedPointLookupTable, out *cachedPoint, xabs uint8)
// Requires: AVX, AVX2
TEXT ·lookupCached(SB), NOSPLIT|NOFRAME, $0-24
	MOVQ table+0(FP), AX

	// Build the mask, zero all the registers
	MOVBQZX      xabs+16(FP), CX
	VMOVD        CX, X0
	VPBROADCASTD X0, Y0
	VPXOR        Y2, Y2, Y2
	VPXOR        Y3, Y3, Y3
	VPXOR        Y4, Y4, Y4
	VPXOR        Y5, Y5, Y5
	VPXOR        Y6, Y6, Y6

	// 0: Identity element
	VPXOR    Y1, Y1, Y1
	VPCMPEQD Y0, Y1, Y1
	VMOVDQA  cached_id_0<>+0(SB), Y7
	VMOVDQA  cached_id_1<>+0(SB), Y8
	VMOVDQA  cached_id_2_4<>+0(SB), Y9
	VMOVDQA  Y9, Y10
	VMOVDQA  Y9, Y11
	VPAND    Y7, Y1, Y2
	VPAND    Y8, Y1, Y3
	VPAND    Y9, Y1, Y4
	VPAND    Y10, Y1, Y5
	VPAND    Y11, Y1, Y6

	// 1 .. 8
	MOVQ $0x0000000000000001, CX

cached_lookup_loop:
	VMOVQ        CX, X1
	VPBROADCASTD X1, Y1
	VPCMPEQD     Y0, Y1, Y1
	VPAND        (AX), Y1, Y7
	VPAND        32(AX), Y1, Y8
	VPAND        64(AX), Y1, Y9
	VPAND        96(AX), Y1, Y10
	VPAND        128(AX), Y1, Y11
	VPOR         Y2, Y7, Y2
	VPOR         Y3, Y8, Y3
	VPOR         Y4, Y9, Y4
	VPOR         Y5, Y10, Y5
	VPOR         Y6, Y11, Y6
	ADDQ         $0xa0, AX
	INCQ         CX
	CMPQ         CX, $0x08
	JLE          cached_lookup_loop

	// Write out the result
	MOVQ    out+8(FP), AX
	VMOVDQU Y2, (AX)
	VMOVDQU Y3, 32(AX)
	VMOVDQU Y4, 64(AX)
	VMOVDQU Y5, 96(AX)
	VMOVDQU Y6, 128(AX)
	VZEROUPPER
	RET
