package middleware

import (
	"fmt"
	"server/setting"
)

func MailboxLogin(params *setting.User) {
	fmt.Printf("params: %v\n", params)
}
