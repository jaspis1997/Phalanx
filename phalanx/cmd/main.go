package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	VERSION = ""
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

var rootCommand = &cobra.Command{}

func main() {
	version := rootCommand.Flags().BoolP("version", "v", false, "")
	if *version {
		log.Print(VERSION)
	}
	rootCommand.Execute()
}
