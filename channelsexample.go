package main

import (
	"fmt"
	"time"
)

var (
	start    time.Time
	duration time.Duration
)

func RunChannelsExample() {
	greetings := make(chan string)

	go CreateQuickGreeting("World", greetings)
	go CreateLongGreeting("There", greetings)

	fmt.Println("Greeting is coming...")

	start = time.Now()
	fmt.Println(<-greetings)
	fmt.Println(<-greetings)
	duration = time.Since(start)
	fmt.Println("duration:", duration)

	// start = time.Now()
	// fullGreeting := fmt.Sprintf("%s\n%s", <-c, <-c)
	// duration = time.Since(start)
	// fmt.Println(fullGreeting)
	// fmt.Println("duration:", duration)
}

func CreateQuickGreeting(s string, c chan string) {
	fmt.Println("CreateQuickGreeting() running")
	time.Sleep(time.Second * 1)
	greeting := fmt.Sprintf("Hello %s!", s)
	c <- greeting
}

func CreateLongGreeting(s string, c chan string) {
	fmt.Println("CreateLongGreeting() running")
	time.Sleep(time.Second * 4)
	greeting := fmt.Sprintf("Hello %s! How are you? Hope you fine today", s)
	c <- greeting
}
