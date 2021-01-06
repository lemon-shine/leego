/*******************************************************************************
Method: 认证中间件
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package leego

type Auth struct {
	Username string
	Password string
}

// //验证POST或GET的用户名和密码
// func Authenticate(auth *Auth) HandleFunc {
// 	return func(ctx *Context) {
// 		ctx.Request.BasicAuth()
// 		//检验PSOT
// 		name, err0 := ctx.GetPostForm("username")
// 		pwd, err1 := ctx.GetPostForm("password")
// 		if err0 != nil || err1 != nil {
// 			//校验GET
// 			if ctx.Query("username") != auth.Username || ctx.Query("password") != auth.Password {
// 				ctx.ResponseString(405, "User name or password authentication failed")
// 				return
// 			}
// 		}
// 		if name != auth.Username || pwd != auth.Password {
// 			ctx.ResponseString(405, "User name or password authentication failed")
// 			return
// 		}

// 		ctx.NextMiddlewares()
// 	}
// }
