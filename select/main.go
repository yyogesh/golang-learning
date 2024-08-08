package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	paniPuriChan := make(chan string)
	vadaPavChan := make(chan string)
	dosaChan := make(chan string)

	go servePaniPuri(paniPuriChan)
	go serveVadaPav(vadaPavChan)
	go serveDosa(dosaChan)

	select {
	case food := <-paniPuriChan:
		fmt.Println("Yay! Got", food)
	case food := <-vadaPavChan:
		fmt.Println("Yay! Got", food)
	case food := <-dosaChan:
		fmt.Println("Yay! Got", food)
	}

}

func servePaniPuri(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- "pani puri"
}

func serveVadaPav(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- "vada pav"
}

func serveDosa(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- "dosa"
}
