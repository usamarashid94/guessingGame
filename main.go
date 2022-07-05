package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Mcqs struct {
	Mcqs []Question `json:"mcqs"`
}

type Choice struct {
	ChoiceOne   string `json:"a"`
	ChoiceTwo   string `json:"b"`
	ChoiceThree string `json:"c"`
	ChoiceFour  string `json:"d"`
}

type Question struct {
	Question      string `json:"question"`
	CorrectChoice string `json:"correctChoice"`
	Choices       Choice `json:"choices"`
}

func inputUserName() {

	fmt.Println("Enter Your Name: ")
	//storing name of user
	var name string
	// Taking input from user
	fmt.Scanln(&name)
	fmt.Println("Your Name is: " + name)
}

func main() {

	inputUserName()
	fileContent, err := os.Open("mcqs.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var mcqs Mcqs

	json.Unmarshal(byteResult, &mcqs)

	var answer string
	var earned int = 0
	var status bool = true

	for i := 0; i < len(mcqs.Mcqs) && status; i++ {
		fmt.Println("THIS IS QUESTION #: ", i+1)
		fmt.Println("Question: " + mcqs.Mcqs[i].Question)
		//fmt.Println("Correct Choice: " + mcqs.Mcqs[i].CorrectChoice)
		fmt.Println("a) " + mcqs.Mcqs[i].Choices.ChoiceOne)
		fmt.Println("b) " + mcqs.Mcqs[i].Choices.ChoiceTwo)
		fmt.Println("c) " + mcqs.Mcqs[i].Choices.ChoiceThree)
		fmt.Println("d) " + mcqs.Mcqs[i].Choices.ChoiceFour)

		fmt.Println("Enter your Answer: ")
		fmt.Scanln(&answer)

		if answer != mcqs.Mcqs[i].CorrectChoice {
			println("Oops! Your answer was incorrect :(")
			fmt.Println("Correct Answer was: " + mcqs.Mcqs[i].CorrectChoice)
			status = false
		} else {
			earned += 100000
		}
	}

	if status {
		fmt.Println("Congratulations! You have won!")
	}

	fmt.Println("Final Points Earned: $", earned)

}
