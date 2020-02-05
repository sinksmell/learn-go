"".main STEXT size=75 args=0x0 locals=0x10
	0x0000 00000 (func.go:4)	TEXT	"".main(SB), ABIInternal, $16-0
	0x0000 00000 (func.go:4)	MOVQ	(TLS), CX
	0x0009 00009 (func.go:4)	CMPQ	SP, 16(CX)
	0x000d 00013 (func.go:4)	JLS	68
	0x000f 00015 (func.go:4)	SUBQ	$16, SP
	0x0013 00019 (func.go:4)	MOVQ	BP, 8(SP)
	0x0018 00024 (func.go:4)	LEAQ	8(SP), BP
	0x001d 00029 (func.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (func.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (func.go:4)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (func.go:5)	PCDATA	$0, $0
	0x001d 00029 (func.go:5)	PCDATA	$1, $0
	0x001d 00029 (func.go:5)	XCHGL	AX, AX
	0x001e 00030 (func.go:10)	CALL	runtime.printlock(SB)
	0x0023 00035 (func.go:10)	MOVQ	$3, (SP)
	0x002b 00043 (func.go:10)	CALL	runtime.printint(SB)
	0x0030 00048 (func.go:10)	CALL	runtime.printnl(SB)
	0x0035 00053 (func.go:10)	CALL	runtime.printunlock(SB)
	0x003a 00058 (<unknown line number>)	MOVQ	8(SP), BP
	0x003f 00063 (<unknown line number>)	ADDQ	$16, SP
	0x0043 00067 (<unknown line number>)	RET
	0x0044 00068 (<unknown line number>)	NOP
	0x0044 00068 (func.go:4)	PCDATA	$1, $-1
	0x0044 00068 (func.go:4)	PCDATA	$0, $-1
	0x0044 00068 (func.go:4)	CALL	runtime.morestack_noctxt(SB)
	0x0049 00073 (func.go:4)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 35 48  eH..%....H;a.v5H
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 90 e8 00  ...H.l$.H.l$....
	0x0020 00 00 00 48 c7 04 24 03 00 00 00 e8 00 00 00 00  ...H..$.........
	0x0030 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 08 48  ..........H.l$.H
	0x0040 83 c4 10 c3 e8 00 00 00 00 eb b5                 ...........
	rel 5+4 t=16 TLS+0
	rel 31+4 t=8 runtime.printlock+0
	rel 44+4 t=8 runtime.printint+0
	rel 49+4 t=8 runtime.printnl+0
	rel 54+4 t=8 runtime.printunlock+0
	rel 69+4 t=8 runtime.morestack_noctxt+0
"".printsum STEXT size=84 args=0x10 locals=0x10
	0x0000 00000 (func.go:8)	TEXT	"".printsum(SB), ABIInternal, $16-16
	0x0000 00000 (func.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (func.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (func.go:8)	JLS	77
	0x000f 00015 (func.go:8)	SUBQ	$16, SP
	0x0013 00019 (func.go:8)	MOVQ	BP, 8(SP)
	0x0018 00024 (func.go:8)	LEAQ	8(SP), BP
	0x001d 00029 (func.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (func.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (func.go:8)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (func.go:9)	PCDATA	$0, $0
	0x001d 00029 (func.go:9)	PCDATA	$1, $0
	0x001d 00029 (func.go:9)	XCHGL	AX, AX
	0x001e 00030 (func.go:10)	CALL	runtime.printlock(SB)
	0x0023 00035 (func.go:14)	MOVQ	"".b+32(SP), AX
	0x0028 00040 (func.go:14)	MOVQ	"".a+24(SP), CX
	0x002d 00045 (func.go:14)	ADDQ	CX, AX
	0x0030 00048 (func.go:10)	MOVQ	AX, (SP)
	0x0034 00052 (func.go:10)	CALL	runtime.printint(SB)
	0x0039 00057 (func.go:10)	CALL	runtime.printnl(SB)
	0x003e 00062 (func.go:10)	CALL	runtime.printunlock(SB)
	0x0043 00067 (func.go:11)	MOVQ	8(SP), BP
	0x0048 00072 (func.go:11)	ADDQ	$16, SP
	0x004c 00076 (func.go:11)	RET
	0x004d 00077 (func.go:11)	NOP
	0x004d 00077 (func.go:8)	PCDATA	$1, $-1
	0x004d 00077 (func.go:8)	PCDATA	$0, $-1
	0x004d 00077 (func.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x0052 00082 (func.go:8)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 3e 48  eH..%....H;a.v>H
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 90 e8 00  ...H.l$.H.l$....
	0x0020 00 00 00 48 8b 44 24 20 48 8b 4c 24 18 48 01 c8  ...H.D$ H.L$.H..
	0x0030 48 89 04 24 e8 00 00 00 00 e8 00 00 00 00 e8 00  H..$............
	0x0040 00 00 00 48 8b 6c 24 08 48 83 c4 10 c3 e8 00 00  ...H.l$.H.......
	0x0050 00 00 eb ac                                      ....
	rel 5+4 t=16 TLS+0
	rel 31+4 t=8 runtime.printlock+0
	rel 53+4 t=8 runtime.printint+0
	rel 58+4 t=8 runtime.printnl+0
	rel 63+4 t=8 runtime.printunlock+0
	rel 78+4 t=8 runtime.morestack_noctxt+0
"".sum STEXT nosplit size=19 args=0x18 locals=0x0
	0x0000 00000 (func.go:13)	TEXT	"".sum(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (func.go:13)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (func.go:13)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (func.go:13)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (func.go:14)	PCDATA	$0, $0
	0x0000 00000 (func.go:14)	PCDATA	$1, $0
	0x0000 00000 (func.go:14)	MOVQ	"".b+16(SP), AX
	0x0005 00005 (func.go:14)	MOVQ	"".a+8(SP), CX
	0x000a 00010 (func.go:14)	ADDQ	CX, AX
	0x000d 00013 (func.go:14)	MOVQ	AX, "".~r2+24(SP)
	0x0012 00018 (func.go:14)	RET
	0x0000 48 8b 44 24 10 48 8b 4c 24 08 48 01 c8 48 89 44  H.D$.H.L$.H..H.D
	0x0010 24 18 c3                                         $..
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".sum$abstract SDWARFINFO dupok size=25
	0x0000 04 2e 73 75 6d 00 01 01 11 61 00 00 00 00 00 00  ..sum....a......
	0x0010 11 62 00 00 00 00 00 00 00                       .b.......
	rel 12+4 t=28 go.info.int+0
	rel 20+4 t=28 go.info.int+0
go.info."".printsum$abstract SDWARFINFO dupok size=40
	0x0000 04 2e 70 72 69 6e 74 73 75 6d 00 01 01 11 61 00  ..printsum....a.
	0x0010 00 00 00 00 00 11 62 00 00 00 00 00 00 0c 72 65  ......b.......re
	0x0020 74 00 09 00 00 00 00 00                          t.......
	rel 17+4 t=28 go.info.int+0
	rel 25+4 t=28 go.info.int+0
	rel 35+4 t=28 go.info.int+0
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=63
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 06 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 05 00 00 00 00 00     ...............
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+75
	rel 27+4 t=29 gofile../Users/sun/VSCode/Go/LearnGo/advancedgo/assembly/func.go+0
	rel 33+4 t=28 go.info."".printsum$abstract+0
	rel 37+8 t=1 "".main+30
	rel 45+8 t=1 "".main+58
	rel 53+4 t=29 gofile../Users/sun/VSCode/Go/LearnGo/advancedgo/assembly/func.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0e 05 01 02 05 01 08 02 0f 01 0a 02 07  ................
	0x0010 00                                               .
go.loc."".printsum SDWARFLOC size=177
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 01 00 50 00 00 00 00 00 00  .........P......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 01 00 9c 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 02 00 91  ................
	0x00a0 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00b0 00                                               .
	rel 0+8 t=50 "".printsum+0
	rel 8+8 t=50 "".printsum+84
	rel 35+8 t=50 "".printsum+0
	rel 43+8 t=50 "".printsum+84
	rel 71+8 t=50 "".printsum+48
	rel 79+8 t=50 "".printsum+57
	rel 106+8 t=50 "".printsum+0
	rel 114+8 t=50 "".printsum+84
	rel 141+8 t=50 "".printsum+0
	rel 149+8 t=50 "".printsum+84
go.info."".printsum SDWARFINFO size=99
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 01 9c 13 00 00 00 00 00 00 00 00  ................
	0x0020 13 00 00 00 00 00 00 00 00 0e 00 00 00 00 00 00  ................
	0x0030 00 00 06 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 09 00 00 00 13  ................
	0x0050 00 00 00 00 00 00 00 00 13 00 00 00 00 00 00 00  ................
	0x0060 00 00 00                                         ...
	rel 1+4 t=28 go.info."".printsum$abstract+0
	rel 5+8 t=1 "".printsum+0
	rel 13+8 t=1 "".printsum+84
	rel 24+4 t=28 go.info."".printsum$abstract+13
	rel 28+4 t=28 go.loc."".printsum+0
	rel 33+4 t=28 go.info."".printsum$abstract+21
	rel 37+4 t=28 go.loc."".printsum+35
	rel 42+4 t=28 go.info."".printsum$abstract+29
	rel 46+4 t=28 go.loc."".printsum+71
	rel 51+4 t=28 go.info."".sum$abstract+0
	rel 55+8 t=1 "".printsum+35
	rel 63+8 t=1 "".printsum+48
	rel 71+4 t=29 gofile../Users/sun/VSCode/Go/LearnGo/advancedgo/assembly/func.go+0
	rel 80+4 t=28 go.info."".sum$abstract+8
	rel 84+4 t=28 go.loc."".printsum+106
	rel 89+4 t=28 go.info."".sum$abstract+16
	rel 93+4 t=28 go.loc."".printsum+141
go.range."".printsum SDWARFRANGE size=0
go.isstmt."".printsum SDWARFMISC size=0
	0x0000 04 0f 04 0e 05 01 02 0a 01 0c 02 20 00           ........... .
go.loc."".sum SDWARFLOC size=71
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00                             .......
	rel 0+8 t=50 "".sum+0
	rel 8+8 t=50 "".sum+19
	rel 35+8 t=50 "".sum+0
	rel 43+8 t=50 "".sum+19
go.info."".sum SDWARFINFO size=54
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 01 9c 13 00 00 00 00 00 00 00 00  ................
	0x0020 13 00 00 00 00 00 00 00 00 0f 7e 72 32 00 01 0d  ..........~r2...
	0x0030 00 00 00 00 00 00                                ......
	rel 1+4 t=28 go.info."".sum$abstract+0
	rel 5+8 t=1 "".sum+0
	rel 13+8 t=1 "".sum+19
	rel 24+4 t=28 go.info."".sum$abstract+8
	rel 28+4 t=28 go.loc."".sum+0
	rel 33+4 t=28 go.info."".sum$abstract+16
	rel 37+4 t=28 go.loc."".sum+35
	rel 48+4 t=28 go.info.int+0
go.range."".sum SDWARFRANGE size=0
go.isstmt."".sum SDWARFMISC size=0
	0x0000 04 05 01 0e 00                                   .....
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
