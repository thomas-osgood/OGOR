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
	formatted = fmt.Sprintf("%s%s %v%s", CLEAR_LINE, FAIL_SEQ, text, NEWLINE)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's SucMsg.
func (f *Formatter) SucMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v%s", CLEAR_LINE, SUCCESS_SEQ, text, NEWLINE)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's InfMsg.
func (f *Formatter) InfMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v%s", CLEAR_LINE, INFO_SEQ, text, NEWLINE)
	return formatted
}

// function designed to take an input and return a formatted Success string.
// this string output will be formatted the same as an Outputter's SysMsg.
func (f *Formatter) SysMsg(text interface{}) (formatted string) {
	formatted = fmt.Sprintf("%s%s %v%s", CLEAR_LINE, SYS_SEQ, text, NEWLINE)
	return formatted
}

// function designed to print a string in the center of a line with
// a given length "n".
//
// currently, this includes non-printable characters in the length calculation.
func (f *Formatter) CenterString(msg interface{}, n int) (formatted string) {
	var indent_format string
	var indent int
	var msgstr string = fmt.Sprintf("%v", msg)

	fmt.Printf("%s", CLEAR_LINE)
	if len(msgstr) > n {
		return fmt.Sprintf("%s%s", msgstr, NEWLINE)
	}

	indent = (n - len(msgstr)) / 2
	indent_format = fmt.Sprintf(INDENT_SEQ, indent)

	return fmt.Sprintf("%s%s%s", indent_format, msg, NEWLINE)
}

// function designed to print a given char "c", "n" number of times.
func (f *Formatter) PrintChar(char byte, n int) (outline string) {
	if n < 1 {
		return ""
	}

	for i := 0; i < n; i++ {
		outline = fmt.Sprintf("%s%s", outline, string(char))
	}

	return outline
}
