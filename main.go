package main

import (
	"1-converter/3-struct/bins"
	"fmt"
)

func main() {
	bin := bins.NewBin("asd", "qwe", false)
	fmt.Println(bin)
	binList := bins.NewBinList()
	binList.AddBin(*bin)
	binList.AddBin(*bin)
	fmt.Println(binList)
}
