package output

import "fmt"

// control sequence introducer shorthand (aka: "\x1B["")
var CSI string = fmt.Sprintf("%s[", ESCAPE_CODE)

// control sequence used to clear the entirty of the current line
var CLEAR_LINE string = fmt.Sprintf("\r%s2K\r", CSI)

// control sequence used to clear the entire screen and move
// the cursor to (0,0).
var CLEAR_SCREEN string = fmt.Sprintf("%s2J%sH", CSI, CSI)

var RED_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_RED)
var GREEN_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_GREEN)
var YELLOW_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_YELLOW)
var BLUE_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_BLUE)

var RED_L_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_RED_L)
var GREEN_L_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_GREEN_L)
var YELLOW_L_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_YELLOW_L)
var BLUE_L_SEQ string = fmt.Sprintf("%s%d;1m", CSI, ANSI_BLUE_L)

// control sequence used to reset the output settings to the
// original values.
var ANSI_RESET string = fmt.Sprintf("%s0m", CSI)

var SUCCESS_SEQ string = fmt.Sprintf("%s[%s]%s", GREEN_L_SEQ, SUCCESS_CHR, ANSI_RESET)
var FAIL_SEQ string = fmt.Sprintf("%s[%s]%s", RED_L_SEQ, FAIL_CHR, ANSI_RESET)
var WARNING_SEQ string = fmt.Sprintf("%s[%s]%s", YELLOW_L_SEQ, WARNING_CHR, ANSI_RESET)
var INFO_SEQ string = fmt.Sprintf("%s[%s]%s", BLUE_L_SEQ, INFO_CHR, ANSI_RESET)
var SYS_SEQ string = fmt.Sprintf("%s[%s]%s", YELLOW_L_SEQ, SYS_CHR, ANSI_RESET)

// escape sequence used to indent n number of spaces "\x1b[%dC"
var INDENT_SEQ string = fmt.Sprintf("%s%%dC", CSI)
