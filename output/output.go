package output

import "fmt"

// function designed to print a success message with the
// format: [+] <message>
func (o *Outputter) SucMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, SUCCESS_SEQ, message)
	return
}

// function designed to print an error message with the
// format: [-] <message>
func (o *Outputter) ErrMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, FAIL_SEQ, message)
	return
}

// function designed to print an info message with the
// format: [i] <message>
func (o *Outputter) InfMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, INFO_SEQ, message)
	return
}

// function designed to print an info message with the
// format: [i] <message>
//
// note: this does not add a newline character to the end.
func (o *Outputter) InfMsgNB(message interface{}) {
	fmt.Printf("%s%s %v\r", CLEAR_LINE, INFO_SEQ, message)
	return
}

// function designed to print a system message with the
// format: [*] <message>
func (o *Outputter) SysMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, SYS_SEQ, message)
	return
}

// function designed to print a system message with the
// format: [*] <message>
//
// note: this does not add a newline character to the end.
func (o *Outputter) SysMsgNB(message interface{}) {
	fmt.Printf("%s%s %v\r", CLEAR_LINE, SYS_SEQ, message)
	return
}

// function designed to print a warning message with the
// format: [!] <message>
func (o *Outputter) WrnMsg(message interface{}) {
	fmt.Printf("%s%s %v\n", CLEAR_LINE, WARNING_SEQ, message)
	return
}

// function designed to print the given message in Red. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintRed(message interface{}) {
	fmt.Printf("%s%v%s\n", RED_SEQ, message, ANSI_RESET)
}

// function designed to print the given message in Green. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintGreen(message interface{}) {
	fmt.Printf("%s%v%s\n", GREEN_SEQ, message, ANSI_RESET)
}

// function designed to print the given message in Yellow. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintYellow(message interface{}) {
	fmt.Printf("%s%v%s\n", YELLOW_SEQ, message, ANSI_RESET)
}

// function designed to print the given message in Blue. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintBlue(message interface{}) {
	fmt.Printf("%s%v%s\n", BLUE_SEQ, message, ANSI_RESET)
}
