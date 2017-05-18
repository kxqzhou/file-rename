
package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	//"flag"
	"io/ioutil"
)

func main() {

	dir := "."
	
	files, err := ioutil.ReadDir( dir )
	if err != nil {
		log.Fatal( err )
	}

	for i, file := range files {
		fmt.Println( file.Name() )
		os.Rename( file.Name(), "t" + strconv.Itoa(i + 1) )
	}

}

