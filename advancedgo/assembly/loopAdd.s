"".main STEXT size=77 args=0x0 locals=0x28
	0x0000 00000 (loopAdd.go:3)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (loopAdd.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (loopAdd.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (loopAdd.go:3)	JLS	70
	0x000f 00015 (loopAdd.go:3)	SUBQ	$40, SP
	0x0013 00019 (loopAdd.go:3)	MOVQ	BP, 32(SP)
	0x0018 00024 (loopAdd.go:3)	LEAQ	32(SP), BP
	0x001d 00029 (loopAdd.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (loopAdd.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (loopAdd.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (loopAdd.go:4)	PCDATA	$0, $0
	0x001d 00029 (loopAdd.go:4)	PCDATA	$1, $0
	0x001d 00029 (loopAdd.go:4)	MOVQ	$10, (SP)
	0x0025 00037 (loopAdd.go:4)	MOVQ	$0, 8(SP)
	0x002e 00046 (loopAdd.go:4)	MOVQ	$1, 16(SP)
	0x0037 00055 (loopAdd.go:4)	CALL	"".LoopAdd(SB)
	0x003c 00060 (loopAdd.go:5)	MOVQ	32(SP), BP
	0x0041 00065 (loopAdd.go:5)	ADDQ	$40, SP
	0x0045 00069 (loopAdd.go:5)	RET
	0x0046 00070 (loopAdd.go:5)	NOP
	0x0046 00070 (loopAdd.go:3)	PCDATA	$1, $-1
	0x0046 00070 (loopAdd.go:3)	PCDATA	$0, $-1
	0x0046 00070 (loopAdd.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (loopAdd.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 37 48  eH..%....H;a.v7H
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 04  ..(H.l$ H.l$ H..
	0x0020 24 0a 00 00 00 48 c7 44 24 08 00 00 00 00 48 c7  $....H.D$.....H.
	0x0030 44 24 10 01 00 00 00 e8 00 00 00 00 48 8b 6c 24  D$..........H.l$
	0x0040 20 48 83 c4 28 c3 e8 00 00 00 00 eb b3            H..(........
	rel 5+4 t=16 TLS+0
	rel 56+4 t=8 "".LoopAdd+0
	rel 71+4 t=8 runtime.morestack_noctxt+0
"".LoopAdd STEXT nosplit size=36 args=0x20 locals=0x0
	0x0000 00000 (loopAdd.go:7)	TEXT	"".LoopAdd(SB), NOSPLIT|ABIInternal, $0-32
	0x0000 00000 (loopAdd.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (loopAdd.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (loopAdd.go:7)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (loopAdd.go:9)	PCDATA	$0, $0
	0x0000 00000 (loopAdd.go:9)	PCDATA	$1, $0
	0x0000 00000 (loopAdd.go:9)	MOVQ	"".step+24(SP), AX
	0x0005 00005 (loopAdd.go:9)	MOVQ	"".cnt+8(SP), CX
	0x000a 00010 (loopAdd.go:9)	MOVQ	"".v0+16(SP), DX
	0x000f 00015 (loopAdd.go:9)	XORL	BX, BX
	0x0011 00017 (loopAdd.go:9)	JMP	25
	0x0013 00019 (loopAdd.go:9)	INCQ	BX
	0x0016 00022 (loopAdd.go:10)	ADDQ	AX, DX
	0x0019 00025 (loopAdd.go:9)	CMPQ	BX, CX
	0x001c 00028 (loopAdd.go:9)	JLT	19
	0x001e 00030 (loopAdd.go:12)	MOVQ	DX, "".~r3+32(SP)
	0x0023 00035 (loopAdd.go:12)	RET
	0x0000 48 8b 44 24 18 48 8b 4c 24 08 48 8b 54 24 10 31  H.D$.H.L$.H.T$.1
	0x0010 db eb 06 48 ff c3 48 01 c2 48 39 cb 7c f5 48 89  ...H..H..H9.|.H.
	0x0020 54 24 20 c3                                      T$ .
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+77
	rel 27+4 t=29 gofile../Users/sun/VSCode/Go/LearnGo/advancedgo/assembly/loopAdd.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 08 01 12 02 16 00                 ...........
go.loc."".LoopAdd SDWARFLOC size=178
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 02 00 91 08 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 01 00 9c 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 02 00 91 08 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 02 00 91 10 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x00a0 00 53 00 00 00 00 00 00 00 00 00 00 00 00 00 00  .S..............
	0x00b0 00 00                                            ..
	rel 0+8 t=50 "".LoopAdd+0
	rel 8+8 t=50 "".LoopAdd+36
	rel 36+8 t=50 "".LoopAdd+0
	rel 44+8 t=50 "".LoopAdd+36
	rel 71+8 t=50 "".LoopAdd+0
	rel 79+8 t=50 "".LoopAdd+36
	rel 107+8 t=50 "".LoopAdd+0
	rel 115+8 t=50 "".LoopAdd+36
	rel 143+8 t=50 "".LoopAdd+22
	rel 151+8 t=50 "".LoopAdd+36
go.info."".LoopAdd SDWARFINFO size=137
	0x0000 03 22 22 2e 4c 6f 6f 70 41 64 64 00 00 00 00 00  ."".LoopAdd.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0020 00 00 01 0b 72 65 73 00 08 00 00 00 00 00 00 00  ....res.........
	0x0030 00 10 63 6e 74 00 00 07 00 00 00 00 00 00 00 00  ..cnt...........
	0x0040 10 76 30 00 00 07 00 00 00 00 00 00 00 00 10 73  .v0............s
	0x0050 74 65 70 00 00 07 00 00 00 00 00 00 00 00 0f 7e  tep............~
	0x0060 72 33 00 01 07 00 00 00 00 00 15 00 00 00 00 00  r3..............
	0x0070 00 00 00 00 00 00 00 00 00 00 00 0b 69 00 09 00  ............i...
	0x0080 00 00 00 00 00 00 00 00 00                       .........
	rel 12+8 t=1 "".LoopAdd+0
	rel 20+8 t=1 "".LoopAdd+36
	rel 30+4 t=29 gofile../Users/sun/VSCode/Go/LearnGo/advancedgo/assembly/loopAdd.go+0
	rel 41+4 t=28 go.info.int+0
	rel 45+4 t=28 go.loc."".LoopAdd+0
	rel 56+4 t=28 go.info.int+0
	rel 60+4 t=28 go.loc."".LoopAdd+36
	rel 70+4 t=28 go.info.int+0
	rel 74+4 t=28 go.loc."".LoopAdd+71
	rel 86+4 t=28 go.info.int+0
	rel 90+4 t=28 go.loc."".LoopAdd+107
	rel 101+4 t=28 go.info.int+0
	rel 107+8 t=1 "".LoopAdd+0
	rel 115+8 t=1 "".LoopAdd+30
	rel 127+4 t=28 go.info.int+0
	rel 131+4 t=28 go.loc."".LoopAdd+143
go.range."".LoopAdd SDWARFRANGE size=0
go.isstmt."".LoopAdd SDWARFMISC size=0
	0x0000 04 05 01 0e 02 09 01 02 02 05 01 01 00           .............
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
