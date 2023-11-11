package utils

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 验证码存储位置 默认10240 个 过期时间10分钟
var store = base64Captcha.DefaultMemStore

// 生成验证码
func MakeCaptcha() (id, base64 string, code string, err error) {
	// 定义一个Driver
	var driver base64Captcha.Driver

	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          38,                                 // 高度
		Width:           100,                                // 宽度
		NoiseCount:      0,                                  // 干扰数
		ShowLineOptions: 2 | 4,                              //展示个数
		Length:          4,                                  // 长度
		Source:          "1234567890qwertyupkjhgfdsaxcvbnm", // 验证码随机字符来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体
	}

	driver = driverString.ConvertFonts()

	// 生成验证码
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, base64, err = captcha.Generate()
	code = store.Get(id, false)
	return id, base64, code, err
}

// 验证captcha是否正确
func VerifyCaptcha(id string, verifyValue string) bool {
	if id == "" || verifyValue == "" {
		return false
	}
	return store.Verify(id, verifyValue, true)
}
