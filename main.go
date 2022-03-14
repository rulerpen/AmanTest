package main

import (
	"AmanTest/httpServer"
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	go httpServer.Server1()
	fmt.Println("1")
	go httpServer.Server2()
	fmt.Println("2")
	wg.Wait()
}
