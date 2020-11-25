package main

import (
	"fmt"

	"github.com/SuvorovNikita/golang/02Var/nested/hello"
	"github.com/SuvorovNikita/golang/02Var/nested/say"
)

func main() {
	fmt.Println("Start propgrram....")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}
