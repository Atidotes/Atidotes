/**
 ***  根据日期随机生成账号
 */
package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 随机生成三位数字字符串
func randomString() (randomNumber string) {
	randomString := strconv.Itoa(rand.Intn(1001))

	if len(randomString) == 2 {
		randomNumber = "0" + randomString
	} else if len(randomString) == 1 {
		randomNumber = "00" + randomString
	} else {
		randomNumber = randomString
	}
	return randomNumber
}

// 生成随机账号
func GenerateAccount() (accountNumber string) {
	currentTime := time.Now()
	dateString := currentTime.String()
	year := dateString[2:4]
	month := dateString[5:7]
	day := dateString[8:10]
	randomString := randomString()

	// 拼接字符串
	var build strings.Builder
	build.WriteString(year)
	build.WriteString(month)
	build.WriteString(day)
	build.WriteString(randomString)
	accountNumber = build.String()

	return accountNumber
}
