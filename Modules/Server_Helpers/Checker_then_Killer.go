package Killers_Modders_Writers_Readers_Runners

import "os"

func Check_Kill(filename string) string {
	_, x := os.Stat(filename)
	if x != nil {
		return "file not found"
	} else {
		x := os.Remove(filename)
		if x != nil {
			return "could not remove file"
		} else {
			return "removed file correctly"
		}
	}
}
