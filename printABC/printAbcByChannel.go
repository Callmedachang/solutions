package main

import (
	"log"
)

func printAA() {
	log.Println("A")
}

func printBB() {
	log.Println("B")
}

func printCC() {
	log.Println("C")
}

func main() {
	ch := make(chan int, 0)
	<-ch
}
