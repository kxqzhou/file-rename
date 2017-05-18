
package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"strings"
	"strconv"
	"io/ioutil"
	"path/filepath"
)

type replFlag []string

func (i *replFlag) String() string {
	return *i + " " + *(i + 1)
}

func(i *replFlag) Set(value string) error {
	args := strings.Split( value, " " )

	if len(args) > 2 {
		fmt.Println( "Too many arguments to replace flag, ignoring after the first 2" )
	}

	*i = args[0]
	*(i + 1) = args[1]
	return nil
}

func main() {
	groupCommand := flag.NewFlagSet( "group", flag.ExitOnError )
	seqCommand := flag.NewFlagSet( "seq", flag.ExitOnError )

	groupDir := groupCommand.String( "folder", ".", "The folder containing the files to be renamed" )
	prefix := groupCommand.String( "prefix", "", "A prefix to be appended to the front of each file name" )
	postfix := groupCommand.String( "postfix", "", "A prefix to be appended to the back of each file name" )
	var swap replFlag
	groupCommand.Var( &swap, "replace", "", "A substring and its replacement value, separated by a space" )

	seqDir := seqCommand.String( "folder", ".", "The folder containing the files to be renamed" )
	name := seqCommand.String( "base", "", "The shared base name among all files in the sequence" )

	if len( os.Args ) < 2 {
		fmt.Println( "Group or Seq subcommand required" )
		os.Exit(1)
	}

	switch os.Args[1] {
	case "group":
		groupCommand.Parse()
	case "seq":
		seqCommand.Parse()
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	/*
	fmt.Println( "Current files:" )
	for i, file := range files {
		fmt.Println( file.Name() )
	}
	*/

	if groupCommand.Parsed() {
		files, err := ioutil.ReadDir( *groupDir )
		if err != nil {
			log.Fatal( err )
		}

	} else if seqCommand.Parsed() {
		files, err := ioutil.ReadDir( *seqDir )
		if err != nil {
			log.Fatal( err )
		}

		ext := filepath.Ext( files[0].Name() )
		for i, file := range files {
			os.Rename( file.Name(), *name + strconv.Itoa(i + 1) + ext )
		}
	}
}

