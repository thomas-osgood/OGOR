package output

import "fmt"

// function designed to take an input and return a Red colored
// string containing the given text.
func (f *Formatter) RedText(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%v%s", RED_SEQ, text, ANSI_RESET)
	return formatted
}

// function designed to take an input and return a Green colored
// string containing the given text.
func (f *Formatter) GreenText(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%v%s", GREEN_SEQ, text, ANSI_RESET)
	return formatted
}

// function designed to take an input and return a Yellow colored
// string containing the given text.
func (f *Formatter) YellowText(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%v%s", YELLOW_SEQ, text, ANSI_RESET)
	return formatted
}

// function designed to take an input and return a Blue colored
// string containing the given text.
func (f *Formatter) BlueText(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%v%s", BLUE_SEQ, text, ANSI_RESET)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's ErrMsg.
func (f *Formatter) ErrMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v\n", CLEAR_LINE, FAIL_SEQ, text)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's SucMsg.
func (f *Formatter) SucMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v\n", CLEAR_LINE, SUCCESS_SEQ, text)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's InfMsg.
func (f *Formatter) InfMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v\n", CLEAR_LINE, INFO_SEQ, text)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's SysMsg.
func (f *Formatter) SysMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v\n", CLEAR_LINE, SYS_SEQ, text)
	return formatted
}
