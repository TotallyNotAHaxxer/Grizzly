package Grizzly_YAML

import (
	"io/ioutil"
	"log"
)

func File_Writer(filename string, data string) {
	Pl := ioutil.WriteFile(filename, []byte(data), 0600)
	if Pl != nil {
		log.Fatal(Pl)
	}
}
