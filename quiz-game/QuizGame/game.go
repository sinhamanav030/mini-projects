package QuizGame

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var score, total int

func init() {
	score = 0
	total = 0
}

func Timer(ch chan bool, timer int) {
	select {
	case <-time.After(time.Duration(timer) * time.Second):
		ch <- true
	}

}

func handler(ch chan bool) {
	select {
	case <-ch:
		fmt.Printf("\nTime Out\nYour Score : %d out of %d \n", score, total)
		os.Exit(0)
	}
}

func ReadCSV() {
	f := flag.String("fp", "./QuizGame/CSV/problems.csv", "Get problem file path")
	t := flag.Int("timer", 10, "Set timer for quiz")
	flag.Parse()

	quiz, err := ReadFile(f)
	if err != nil {
		panic(err)
	}
	var ans string
	tc := make(chan bool)
	go Timer(tc, *t)
	go handler(tc)
	total = len(quiz.ques)
	fmt.Println("Quiz is Ready! Are you ready to go:\nPress enter to start")
	fmt.Scanf("%s")
	for _, ques := range quiz.ques {
		fmt.Print(ques.Question, " :")
		fmt.Scanf("%s\n", &ans)
		if ques.Solution == ans {
			score++
		}
	}
	fmt.Printf("\nQuiz Complete\nYour Score : %d out of %d \n", score, len(quiz.ques))
}
