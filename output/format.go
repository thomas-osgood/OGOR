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
