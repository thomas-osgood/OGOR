package output

import "fmt"

func (o *Outputter) SucMsg(message string) {
	fmt.Printf("%s %s\n", SUCCESS_SEQ, message)
	return
}

func (o *Outputter) ErrMsg(message string) {
	fmt.Printf("%s %s\n", FAIL_SEQ, message)
	return
}

func (o *Outputter) InfMsg(message string) {
	fmt.Printf("%s %s\n", INFO_SEQ, message)
	return
}

func (o *Outputter) InfMsgNB(message string) {
	fmt.Printf("%s %s\r", INFO_SEQ, message)
	return
}

func (o *Outputter) SysMsg(message string) {
	fmt.Printf("%s %s\n", SYS_SEQ, message)
	return
}

func (o *Outputter) SysMsgNB(message string) {
	fmt.Printf("%s %s\r", SYS_SEQ, message)
	return
}

func (o *Outputter) WrnMsg(message string) {
	fmt.Printf("%s %s\n", WARNING_SEQ, message)
	return
}

