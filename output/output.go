package output

import "fmt"

// function designed to print a string in the center of a line with
// a given length "n".
//
// currently, this includes non-printable characters in the length calculation.
func (o *Outputter) CenterString(msg interface{}, n int) {
	var indent_format string
	var indent int
	var msgstr string = fmt.Sprintf("%v", msg)

	fmt.Printf("%s", CLEAR_LINE)
	if len(msgstr) > n {
		o.Mutex.Lock()
		fmt.Printf("%s%s", msgstr, NEWLINE)
		o.Mutex.Unlock()
	} else {
		indent = (n - len(msgstr)) / 2
		indent_format = fmt.Sprintf(INDENT_SEQ, indent)

		o.Mutex.Lock()
		fmt.Printf("%s%s%s", indent_format, msg, NEWLINE)
		o.Mutex.Unlock()
	}

	return
}

// function designed to print a given char "c", "n" number of times.
// this appends a newline character to the end of the sequence.
func (o *Outputter) PrintChar(char byte, n int) {
	var outline string

	if n < 1 {
		outline = ""
	} else {
		for i := 0; i < n; i++ {
			outline = fmt.Sprintf("%s%s", outline, string(char))
		}
	}
	o.Mutex.Lock()
	fmt.Printf("%s%s", outline, NEWLINE)
	o.Mutex.Unlock()

	return
}

// function designed to print a success message with the
// format: [+] <message>
func (o *Outputter) SucMsg(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v%s", CLEAR_LINE, SUCCESS_SEQ, message, NEWLINE)
	o.Mutex.Unlock()
	return
}

// function designed to print an error message with the
// format: [-] <message>
func (o *Outputter) ErrMsg(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v%s", CLEAR_LINE, FAIL_SEQ, message, NEWLINE)
	o.Mutex.Unlock()
	return
}

// function designed to print an info message with the
// format: [i] <message>
func (o *Outputter) InfMsg(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v%s", CLEAR_LINE, INFO_SEQ, message, NEWLINE)
	o.Mutex.Unlock()
	return
}

// function designed to print an info message with the
// format: [i] <message>
//
// note: this does not add a newline character to the end.
func (o *Outputter) InfMsgNB(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v\r", CLEAR_LINE, INFO_SEQ, message)
	o.Mutex.Unlock()
	return
}

// function designed to print a system message with the
// format: [*] <message>
func (o *Outputter) SysMsg(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v%s", CLEAR_LINE, SYS_SEQ, message, NEWLINE)
	o.Mutex.Unlock()
	return
}

// function designed to print a system message with the
// format: [*] <message>
//
// note: this does not add a newline character to the end.
func (o *Outputter) SysMsgNB(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v\r", CLEAR_LINE, SYS_SEQ, message)
	o.Mutex.Unlock()
	return
}

// function designed to print a warning message with the
// format: [!] <message>
func (o *Outputter) WrnMsg(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%s %v%s", CLEAR_LINE, WARNING_SEQ, message, NEWLINE)
	o.Mutex.Unlock()
	return
}

// function designed to print the given message in Red. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintRed(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%v%s%s", RED_SEQ, message, ANSI_RESET, NEWLINE)
	o.Mutex.Unlock()
}

// function designed to print the given message in Green. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintGreen(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%v%s%s", GREEN_SEQ, message, ANSI_RESET, NEWLINE)
	o.Mutex.Unlock()
}

// function designed to print the given message in Yellow. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintYellow(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%v%s%s", YELLOW_SEQ, message, ANSI_RESET, NEWLINE)
	o.Mutex.Unlock()
}

// function designed to print the given message in Blue. a newline
// will be appended to the message upon printing.
func (o *Outputter) PrintBlue(message interface{}) {
	o.Mutex.Lock()
	fmt.Printf("%s%v%s%s", BLUE_SEQ, message, ANSI_RESET, NEWLINE)
	o.Mutex.Unlock()
}
