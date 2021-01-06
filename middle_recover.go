/*******************************************************************************
Method: 错误恢复中间件
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func Recover() HandleFunc {
	return func(ctx *Context) {
		defer func(ctx *Context) {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n", trace(msg))
				ctx.ResponseString(500, "Internal Server Error")
			}
		}(ctx)

		ctx.NextMiddlewares()
	}
}

//------------------------------------------------------------------------------

func trace(msg string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) //跳过前3层调用

	var str strings.Builder
	str.WriteString(msg + "\nTraceback: ")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}

	return str.String()
}
