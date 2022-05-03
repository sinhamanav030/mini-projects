package QuizGame

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var score, total int

func init() {
	score = 0
	total = 0
}

func timer(ch chan bool, timer int) {
	time.Sleep(time.Duration(timer) * time.Second)
	ch <- true
}

func handler(ch chan string, solution string) {
	var ans string
	fmt.Scanf("%s", &ans)
	ch <- ans
}

func QuizGame() {
	f := flag.String("fp", "./QuizGame/CSV/problems.csv", "Get problem file path")
	t := flag.Int("timer", 10, "Set timer for quiz")
	flag.Parse()

	quiz, err := ReadFile(f)
	if err != nil {
		log.Fatal(err)
		return
	}
	tc := make(chan bool)
	ans := make(chan string)
	score, notAtTimeScore := 0, 0
	fmt.Println("Quiz is Ready! Are you ready to go:\nPress enter to start")
	fmt.Scanf("%s")
	for _, ques := range quiz.ques {
		fmt.Print(ques.Question, ":")
		go timer(tc, *t)
		go handler(ans, ques.Solution)
		select {
		case <-tc:
			fmt.Print("\n\nTime's Up\nit wont count, but have a go:")
			sol := <-ans
			if sol == ques.Solution {
				notAtTimeScore++
			}
			fmt.Println()
			break
		case sol := <-ans:
			fmt.Print("\nYou're Fast\nGearing up next Question\n\n")
			<-tc
			if sol == ques.Solution {
				score++
			}
			break
		}
	}

	fmt.Printf("\nQuiz Complete\nYour Score : %d out of %d \nQuestion Correctly Answered(in time + not in time) :%d\n", score, quiz.Total, score+notAtTimeScore)
}
