/**
 * Copyright 2017 wfs Author. All Rights Reserved.
 * email: donnie4w@gmail.com
 */
package main

import (
	"flag"

	. "github.com/charliexp/wfs/conf"
	"github.com/charliexp/wfs/httpserver"
)

func main() {
	ParseFlag()
	flag.Parse()
	httpserver.Start()
}
