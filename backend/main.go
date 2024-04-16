package main

import (
	"fmt"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

func main() {
	//fmt.Println("hello project test")

	fmt.Println(mtg.NewQuery().Where(mtg.CardPower, "5").Random(1))

}
