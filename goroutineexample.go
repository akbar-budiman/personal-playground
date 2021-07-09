package main

import (
	"fmt"
	"net/http"
	"time"
)

func RunGoRoutineExampleController(w http.ResponseWriter, r *http.Request) {
	RunGoRoutineExample()
}

func RunGoRoutineExample() {
	fmt.Println("goroutine example started")

	go func() {
		fmt.Println("first goroutine started")
		for i := 0; i < 3; i = i + 1 {
			time.Sleep(time.Second)
			fmt.Println("first goroutine :", i)
		}
	}()

	fmt.Println("first goroutine should have been started")

	go func() {
		fmt.Println("second goroutine started")
		for i := 0; i < 3; i = i + 1 {
			time.Sleep(time.Second)
			fmt.Println("second goroutine :", i)
		}
	}()

	fmt.Println("second goroutine should have been started")

	go f("third goroutine")

	fmt.Println("third goroutine should have been started")

	f("sync")
}

func f(from string) {
	fmt.Println(from, "started")
	time.Sleep(time.Second * 3)
	fmt.Println(from, "ended")
}
