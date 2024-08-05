package gobotapi

import "fmt"

var noticeDisplayed = false

func showNotice() {
	if !noticeDisplayed {
		noticeDisplayed = true
		message := "GoBotAPI v1.7.7, Bot API v7.8, Copyright (C) "
		message += "2022 Laky-64 <https://github.com/Laky-64>\n"
		message += "Licensed under the terms of the GNU Lesser "
		message += "General Public License v3 or later (LGPLv3+)\n"
		fmt.Println(message)
	}
}
