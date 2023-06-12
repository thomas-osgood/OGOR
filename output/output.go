package output

import "fmt"

func (o *Outputter) SucMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, SUCCESS_SEQ, message)
	return
}

func (o *Outputter) ErrMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, FAIL_SEQ, message)
	return
}

func (o *Outputter) InfMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, INFO_SEQ, message)
	return
}

func (o *Outputter) InfMsgNB(message interface{}) {
	fmt.Printf("%s%s %v\r", CLEAR_LINE, INFO_SEQ, message)
	return
}

func (o *Outputter) SysMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, SYS_SEQ, message)
	return
}

func (o *Outputter) SysMsgNB(message interface{}) {
	fmt.Printf("%s%s %v\r", CLEAR_LINE, SYS_SEQ, message)
	return
}

func (o *Outputter) WrnMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, WARNING_SEQ, message)
	return
}
