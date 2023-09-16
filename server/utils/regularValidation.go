/**
 ***  正则验证规则
 */

package utils

import "regexp"

// 手机号正则验证
func PhoneVerification(phone string) bool {
	Reg := `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`

	// 编译正则表达式
	reg := regexp.MustCompile(Reg)

	// 匹配正则
	return reg.MatchString(phone)
}
