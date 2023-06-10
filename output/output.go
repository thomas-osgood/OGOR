package output

import "fmt"

func (o *Outputter) SucMsg(message string) {
	fmt.Printf("%s%s %s\n", CLEAR_LINE, SUCCESS_SEQ, message)
	return
}

func (o *Outputter) ErrMsg(message string) {
	fmt.Printf("%s%s %s\n", CLEAR_LINE, FAIL_SEQ, message)
	return
}

func (o *Outputter) InfMsg(message string) {
	fmt.Printf("%s%s %s\n", CLEAR_LINE, INFO_SEQ, message)
	return
}

func (o *Outputter) InfMsgNB(message string) {
	fmt.Printf("%s%s %s\r", CLEAR_LINE, INFO_SEQ, message)
	return
}

func (o *Outputter) SysMsg(message string) {
	fmt.Printf("%s%s %s\n", CLEAR_LINE, SYS_SEQ, message)
	return
}

func (o *Outputter) SysMsgNB(message string) {
	fmt.Printf("%s%s %s\r", CLEAR_LINE, SYS_SEQ, message)
	return
}

func (o *Outputter) WrnMsg(message string) {
	fmt.Printf("%s%s %s\n", CLEAR_LINE, WARNING_SEQ, message)
	return
}
