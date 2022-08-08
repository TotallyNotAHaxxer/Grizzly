package Grizzly_Encryption

import (
	"fmt"
	"log"
	"os"
)

func Does_File_Exist(filename string) bool {
	_, x := os.Stat(filename)
	if x != nil {
		return false
	} else {
		return true
	}
}

func Writer(data, file string) {
	if Does_File_Exist(file) {
		w, x := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
		if x != nil {
			fmt.Println("Got error when writing -> ", x)
		} else {
			defer w.Close()
			if _, err := w.WriteString(data); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		f, x := os.Create(file)
		if x != nil {
			log.Fatal(x)
		} else {
			defer f.Close()
			w, x := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
			if x != nil {
				fmt.Println("Got error when writing -> ", x)
			} else {
				defer w.Close()
				if _, err := w.WriteString(data); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
