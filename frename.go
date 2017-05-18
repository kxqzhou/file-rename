
package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"string"
	"strconv"
	"io/ioutil"
)

type replFlag []string

func (i *replFlag) String() string {
	return i[0] + " " + i[1]
}

func(i *replFlag) Set(value string) error {
	args := string.Split( value, " " )
	i[0] = args[0]
	i[1] = args[1]
	return nil
}

func main() {
	groupCommand := flag.NewFlagSet( "group", flag.ExitOnError )
	seqCommand := flag.NewFlagSet( "seq", flag.ExitOnError )

	groupDir := group.String( "folder", ".", "The folder containing the files to be renamed" )
	prefix := group.String( "prefix", "", "A prefix to be appended to the front of each file name" )
	postfix := group.String( "postfix", "", "A prefix to be appended to the back of each file name" )
	var swap replFlag
	group.Var( &swap "", "A substring and its replacement value, separated by a space" )

	seqDir := seq.String( "folder", ".", "The folder containing the files to be renamed" )
	name := seq.String( "base", "", "The shared base name among all files in the sequence" )

	flag.Parse()

	files, err := ioutil.ReadDir( dir )
	if err != nil {
		log.Fatal( err )
	}

	for i, file := range files {
		fmt.Println( file.Name() )
		os.Rename( file.Name(), "t" + strconv.Itoa(i + 1) )
	}

}

