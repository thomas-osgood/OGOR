package main

// main package in OGOR. this will import all subpackages.

import (
	_ "github.com/thomas-osgood/OGOR/networking/ipgrabber"
	_ "github.com/thomas-osgood/OGOR/networking/ipinfo"
	_ "github.com/thomas-osgood/OGOR/output"
)

func main() {
}
