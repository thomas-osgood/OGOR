//go:build windows
// +build windows

package output

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const NEWLINE string = "\r\n"

var VTP bool

func init() {
	var err error
	var hdl windows.Handle
	var mode *uint32 = new(uint32)

	hdl, err = windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		fmt.Printf("[-][OUTPUTINIT] %s\n", err.Error())
		VTP = false
		return
	}

	err = windows.GetConsoleMode(hdl, mode)
	if err != nil {
		fmt.Printf("[-][OUTPUTINIT] %s\n", err.Error())
		VTP = false
		return
	}

	err = windows.SetConsoleMode(hdl, windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	if err != nil {
		fmt.Printf("[-][OUTPUTINIT] %s\n", err.Error())
		VTP = false
		return
	}

	VTP = true

	return
}
