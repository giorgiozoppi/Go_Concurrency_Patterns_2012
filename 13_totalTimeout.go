package main

import (
	"fmt"
	"math/rand"
	"time"
)

//////////////////////////////////////////////
/// aliases
//////////////////////////////////////////////

var println = fmt.Println
var sprintf = fmt.Sprintf
var printf = fmt.Printf

//////////////////////////////////////////////
/// init
//////////////////////////////////////////////

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//////////////////////////////////////////////
/// functions
//////////////////////////////////////////////

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			s := sprintf("%s %d", msg, i)
			c <- s
			ms := time.Duration(rand.Intn(1.1e3))
			time.Sleep(ms * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("boring!")
	totalTimeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			println(s)
		case <-totalTimeout:
			println("You talk too much")
			return
		}
	}
}
