package main

import "facade/pkg/facade"
import "fmt"

func main() {
	server := facade.NewServer()
	result := server.Work()
	fmt.Println(result)
}
