package cmd

import "fmt"

func handler(flag string) {
	fmt.Println(fmt.Sprintf("Flag: %s is not defined!", flag))
}
