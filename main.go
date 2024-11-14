package main

import (
	"fmt"
	"motd/icons"
	"motd/utils"
)

func main() {
	fmt.Println(utils.GetOS())
	fmt.Printf(icons.IconApple())
}
