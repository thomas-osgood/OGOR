package output

import "fmt"

func SucMsg(message string) {
	fmt.Printf("%s %s\n", SUCCESS_SEQ, message)
	return
}

func ErrMsg(message string) {
	fmt.Printf("%s %s\n", FAIL_SEQ, message)
	return
}

func InfMsg(message string) {
	fmt.Printf("%s %s\n", INFO_SEQ, message)
	return
}

func InfMsgNB(message string) {
	fmt.Printf("%s %s\r", INFO_SEQ, message)
	return
}

func SysMsg(message string) {
	fmt.Printf("%s %s\n", SYS_SEQ, message)
	return
}

func SysMsgNB(message string) {
	fmt.Printf("%s %s\r", SYS_SEQ, message)
	return
}

func WrnMsg(message string) {
	fmt.Printf("%s %s\n", WARNING_SEQ, message)
	return
}

