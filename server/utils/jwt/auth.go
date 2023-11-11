package jwt

import (
	"fmt"
	"server/config"
	"server/setting"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由中间件，验证token是否过期
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取配置文件
		v, _ := config.LoadConfigFromYaml()
		tokenSecretKey := &setting.TokenSecretKey{}
		err := v.UnmarshalKey("tokenSecretKey", tokenSecretKey)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		// 获取token信息
		authString := ctx.Request.Header.Get("Authorization")

		// 判断是否获取到用户token
		if len(authString) == 0 {
			utils.Error(ctx, "missing authorization header", "missing authorization header")
			ctx.Abort()
			return
		}

		// 切割token信息
		kv := strings.Split(authString, " ")

		// 判断是否符合规则
		if len(kv) != 2 || kv[0] != "Bearer" {
			// 返回错误代码
			utils.Error(ctx, "token无效，请重新登录", "token无效，请重新登录")
			// 终止代码执行 返回中间件
			ctx.Abort()
			return
		}

		tokenString := kv[1]

		// 解密token
		token, err := DecryptToken(tokenString, tokenSecretKey.SecretKey)

		// 判断token是否过期
		if err != nil {
			utils.Error(ctx, "token过期，请重新登录", "token过期，请重新登录")
			ctx.Abort()
			return
		}

		// 储存用户信息，方便以后获取用户信息
		ctx.Set("userInfo", token)

		// 放行
		ctx.Next()
	}
}
