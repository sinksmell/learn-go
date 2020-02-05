package main

func main() {
	LoopAdd(10, 0, 1)
}

func LoopAdd(cnt, v0, step int) int {
	res := v0
	for i := 0; i < cnt; i++ {
		res += step
	}
	return res
}

/*

res=0
i=0

LOOP_BEGIN:
res=v0

LOOP_IF
if i< cnt {
	goto LOOP_BODY
}
goto LOOP_END

LOOP_BODY:
i+=1
res += step
goto LOOP_IF

LOOP_END
return res

*/

/*
plan9 assembly

TEXT LoopAdd(SB),NOSPLIT,$0-32
	MOVQ cnt+0(FP),AX //cnt
	MOVQ v0+8(FP),BX //v0
	MOVQ step+16(FP),CX // step

LOOP_BEGIN
	MOVQ $0,DX // i=0

LOOP_IF
	CMPQ DX,AX // i<cnt
	JL LOOP_BODY // if i<cnt goto LOOP_BODy
	goto LOOP_END

LOOP_BODY
	ADDQ $1,DX // i+1=1
	ADDQ CX,BX // res +=i
	goto LOOP_IF

LOOP_END
	MOVQ BX,ret+24(FP) // return res
	RET
*/
