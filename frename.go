
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

func main() {
	groupCommand := flag.NewFlagSet( "group", flag.ExitOnError )
	seqCommand := flag.NewFlagSet( "seq", flag.ExitOnError )

	groupDir := groupCommand.String( "folder", ".", "The folder containing the files to be renamed" )
	prefix := groupCommand.String( "prefix", "", "A prefix to be appended to the front of each file name" )
	postfix := groupCommand.String( "postfix", "", "A prefix to be appended to the back of each file name" )
	replace := groupCommand.String( "replace", "", "A substring to be changed" )
	with := groupCommand.String( "with", "", "The replacement string. Defaults to empty (deletion)" )

	seqDir := seqCommand.String( "folder", ".", "The folder containing the files to be renamed" )
	base := seqCommand.String( "base", "", "The shared base name among all files in the sequence" )

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
		fmt.Println( "No subcommand specified" )
		flag.PrintDefaults()
		os.Exit(1)
	}

	/*
	fmt.Println( "Current files:" )
	for i, file := range files {
		fmt.Println( file.Name() )
	}
	*/

	/* TODO
	need to resolve replacing names leading to two files with the same name
	commands don't seem to work in tandem (ie group -postifx -prefix will not apply one of them)
	*/

	if groupCommand.Parsed() {
		files, err := ioutil.ReadDir( *groupDir )
		if err != nil {
			log.Fatal( err )
		}

		// what if this causes 2 files to have the same name?
		if *replace != "" {
			for _, file := range files {
				if *replace != file.Name() {
					os.Rename( file.Name(), strings.Replace( file.Name(), *replace, *with, -1 ) )
				}
			}
		}

		if *prefix != "" {
			for _, file := range files {
				os.Rename( file.Name(), *prefix + file.Name() )
			}
		}

		if *postfix != "" {
			ext := filepath.Ext( files[0].Name() )

			for _, file := range files {
				dotIndex := strings.Index( file.Name(), "." )
				os.Rename( file.Name(), file.Name()[:dotIndex] + *postfix + ext )
			}
		}

	} else if seqCommand.Parsed() {
		files, err := ioutil.ReadDir( *seqDir )
		if err != nil {
			log.Fatal( err )
		}

		if *base != "" {
			ext := filepath.Ext( files[0].Name() )
			for i, file := range files {
				os.Rename( file.Name(), *base + strconv.Itoa(i + 1) + ext )
			}
		} else {
			fmt.Println( "Did not specify sequence base name" )
		}
	}
}

