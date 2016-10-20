// Package main provides the entry point for standalone GoAT testing
package main

import (
	"os"

	"github.com/necrophonic/goat/initialise"
	"github.com/necrophonic/log"
	flag "github.com/ogier/pflag"
)

func main() {

	logLevel := flag.String("loglevel", "", "set the log level (trace/debug/warn/info/error - default info)")
	doInit := flag.BoolP("init", "i", false, "initialise current folder with a default layout. Folder must be empty")
	flag.Parse()

	if *logLevel != "" {
		if err := log.InitFromString(*logLevel); err != nil {
			log.Fatalf("%v", err)
		}
	}

	if *doInit {
		if err := initialise.CurrentFolder(); err != nil {
			log.Fatalf("unable to initialise current folder: %v", err)
		}
		// If initialising then we're not actually running tests so just exit
		os.Exit(0)
	}

	os.Exit(0)
}
