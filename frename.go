
package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"strconv"
	"io/ioutil"
	"path/filepath"
)

func main() {
	groupCommand := flag.NewFlagSet( "group", flag.ExitOnError )
	seqCommand := flag.NewFlagSet( "seq", flag.ExitOnError )

	groupDir := groupCommand.String( "folder", ".", "The folder containing the files to be renamed" )
	prefix := groupCommand.String( "prefix", "", "A prefix to be appended to the front of each file name" )
	postfix := groupCommand.String( "postfix", "", "A prefix to be appended to the back of each file name" )
	find := groupCommand.String( "find", "", "A substring to be changed" )
	replace := groupCommand.String( "replace", "", "The replacement string. Defaults to empty (deletion)" )

	seqDir := seqCommand.String( "folder", ".", "The folder containing the files to be renamed" )
	name := seqCommand.String( "base", "", "The shared base name among all files in the sequence" )

	if len( os.Args ) < 2 {
		fmt.Println( "Group or Seq subcommand required" )
		os.Exit(1)
	}

	switch os.Args[1] {
	case "group":
		groupCommand.Parse( os.Args[2:] )
	case "seq":
		seqCommand.Parse( os.Args[2:] )
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

