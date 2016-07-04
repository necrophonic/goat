package main

import (
	"github.com/necrophonic/log"
	flag "github.com/ogier/pflag"
)

func main() {

	logLevel := flag.String("loglevel", "", "set the log level (trace/debug/warn/info/error - default info)")
	flag.Parse()

	if *logLevel != "" {
		if err := log.InitFromString(*logLevel); err != nil {
			log.Fatal("%v", err)
		}
	}
}
