// # Osgood GOlang Repository (OGOR)
//
// a collection of custom packages designed to simplify various,
// ranging from basic output/logging to networking and beyond.
package main

import (
	_ "github.com/thomas-osgood/OGOR/commandline/arguments"
	_ "github.com/thomas-osgood/OGOR/crypto/decoders"
	_ "github.com/thomas-osgood/OGOR/crypto/encoders"
	_ "github.com/thomas-osgood/OGOR/misc/crosscompile"
	_ "github.com/thomas-osgood/OGOR/misc/error-handling"
	_ "github.com/thomas-osgood/OGOR/misc/generators"
	_ "github.com/thomas-osgood/OGOR/misc/generics"
	_ "github.com/thomas-osgood/OGOR/networking/apis"
	_ "github.com/thomas-osgood/OGOR/networking/ipgrabber"
	_ "github.com/thomas-osgood/OGOR/networking/ipinfo"
	_ "github.com/thomas-osgood/OGOR/networking/proxyscrape"
	_ "github.com/thomas-osgood/OGOR/networking/publicipgrabber"
	_ "github.com/thomas-osgood/OGOR/networking/scanning/dnsenum"
	_ "github.com/thomas-osgood/OGOR/networking/validations"
	_ "github.com/thomas-osgood/OGOR/output"
	_ "github.com/thomas-osgood/OGOR/protobufs/definitions/common"
	_ "github.com/thomas-osgood/OGOR/protobufs/definitions/filehandler"
	_ "github.com/thomas-osgood/OGOR/protobufs/filehandler"
	_ "github.com/thomas-osgood/OGOR/protobufs/general"
	_ "github.com/thomas-osgood/OGOR/sysinfo/firewalls"
	_ "github.com/thomas-osgood/OGOR/sysinfo/general"
	_ "github.com/thomas-osgood/OGOR/sysinfo/networking"
	_ "github.com/thomas-osgood/OGOR/vulns/lfi"
)

func main() {
}
