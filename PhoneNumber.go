package main

import (
	"fmt"
)

var phoneNum string

func getStrPhoneNum(phoneNum string) string {
	if len(phoneNum) < 10 {
		return phoneNum
	}
	return phoneNum[0:3] + "****" + phoneNum[7:]
}
func main() {
	phoneNum = "13643816203"
	newPhoneNum := getStrPhoneNum(phoneNum)
	fmt.Println(newPhoneNum)
	temp := ""
	for _, v := range newPhoneNum {
		if v {
			temp = temp + "#"
		} else {
			temp = temp + string(v)
		}
	}
	fmt.Println(temp)
}
