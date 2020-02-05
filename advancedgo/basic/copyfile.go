package main

import (
	"os"
	"io"
)

// 错误处理示例  defer很重要
func copyFile(dstName,srcName string)(written int64,err error){
	src,err:=os.Open(srcName)
	if err!=nil{
		return
	}
	defer src.Close()
	dst,err:=os.Create(dstName)
	if err!=nil{
		return
	}
	defer dst.Close()
	written,err=io.Copy(dst,src)

	return
}
