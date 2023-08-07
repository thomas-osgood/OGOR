# APIs

## Overview

This package is designed to aid developers in creating APIs in Golang. There are various generic functions that
can be utilized, along with a generic middleware object. The middleware allows the user to create IP blacklists, 
set a built-in authorization function that gets checked prior to sending the request to the router, log request
events and more.