package Grizzly_YAML

import "fmt"

func E(x error, y string) {
	if x != nil {
		fmt.Println(y, x)
	}
}
