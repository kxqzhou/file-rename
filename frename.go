
package main

import (
	"fmt"
	//"os"
	"log"
	//"flag"
	"io/ioutil"
)

func main() {

	dir := "."
	
	files, err := ioutil.ReadDir( dir )
	if err != nil {
		log.Fatal( err )
	}

	for _, file := range files {
		fmt.Println( file.Name() )
	}

}

