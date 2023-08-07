# Public IP Grabber

## Overview

This package is designed to contact [WhatIsMyIP](https://www.whatismyip.com) and pull down information. 
This has the ability to grab the public IP of the machine querying the site, or information related to
a public IP specified by the user. This contacts the site's API (api.whatismyip.com) to pull down the
desired information.

The package has one main struct that all the functionality is linked to. The main struct is called `PublicIPGrabber`. To
initialize this struct, the `NewPublicIPGrabber` function should be called. This function returns a 
pointer to a `PublicIPGrabber` object.  The user can set the configuration options for the new object
by passing in `PublicIPGrabberOptFunc`s. 

## Examples

The following snippet will initialize a `PublicIPGrabber` object which uses the default http client:

```go
import "github.com/thomas-osgood/OGOR/networking/publicipgrabber"

var err error
var grabber *publicipgrabber.PublicIPGrabber

grabber, err = NewPublicIPGrabber()
if err != nil {
    // process error here
}
```

The following snippet will pull down and display the public IP info for the current machine:

```go
err = grabber.GetMyIPInformation()
if err != nil {
    // process error here
}

log.Printf("IP: %s\n", grabber.PublicIP.Ip)
log.Printf("Location: %s\n", grabber.PublicIP.Location)
log.Printf("Provider: %s\n", grabber.PublicIP.Provider)
```