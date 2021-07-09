package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg         sync.WaitGroup
	firstWord  string = "initial"
	secondWord string = "initial"
	thirdWord  string = "initial"
	sentence   string
)

func RunAsyncWithCollectionResultExample() {

	declaredWordsWithoutWaiting()
	time.Sleep(time.Second * 2)
	sentence = fmt.Sprintf("%s %s %s", firstWord, secondWord, thirdWord)
	fmt.Println(sentence)

	resetWords()

	rightWayToDeclareWords()
	sentence = fmt.Sprintf("%s %s %s", firstWord, secondWord, thirdWord)
	fmt.Println(sentence)

	wg.Add(1)
	go declareWord(&sentence, sentence+", make it better")
	wg.Wait()
	fmt.Println(sentence)

}

func declaredWordsWithoutWaiting() {
	go delareWordWithoutDefer(&firstWord, "make")
	go delareWordWithoutDefer(&secondWord, "it")
	go delareWordWithoutDefer(&thirdWord, "happen")
}

func rightWayToDeclareWords() {
	wg.Add(1)
	go declareWord(&firstWord, "make")

	wg.Add(1)
	go declareWord(&secondWord, "it")

	wg.Add(1)
	go costlyDeclaringWord(&thirdWord, "happen")

	wg.Wait()
}

func declareWord(w *string, newWord string) {
	time.Sleep(time.Second * 2)
	*w = newWord
	defer wg.Done()
}

func costlyDeclaringWord(w *string, newWord string) {
	time.Sleep(time.Second * 5)
	*w = newWord
	defer wg.Done()
}

func delareWordWithoutDefer(w *string, newWord string) {
	time.Sleep(time.Second * 1)
	*w = newWord
}

func resetWords() {
	firstWord = "initial"
	secondWord = "initial"
	thirdWord = "initial"
}
