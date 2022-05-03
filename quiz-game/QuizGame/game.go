package QuizGame

import (
	"flag"
	"fmt"
	"log"
	"strings"
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

func display(ques string) {
	q := strings.Split(ques, ":")
	for _, v := range q {
		fmt.Println(v)
	}
}

func checker(submit string, sol string) bool {
	if strings.ToLower(submit) == strings.ToLower(sol) {
		return true
	}
	return false
}

func QuizGame() {
	f := flag.String("fp", "./QuizGame/CSV/problems.csv", "Get problem file path")
	t := flag.Int("timer", 5, "Set timer for quiz")
	flag.Parse()

	// displayMcq := WithDisplayQuesFunc(display)
	// checkAns := WithAnswerCheckFunc(checker)
	// quiz, err := NewQuiz(f, displayMcq, checkAns)
	quiz, err := NewQuiz(f)
	if err != nil {
		log.Fatal(err)
		return
	}
	tc := make(chan bool)

	ans := make(chan string)
	score, notAtTimeScore := 0, 0

	id := generateOrder(len(quiz.ques))
	// fmt.Println(id)

	fmt.Println("Quiz is Ready! Are you ready to go:\nPress enter to start")
	fmt.Scanf("%s")

	for _, i := range id {
		quiz.displayQuesFunc(quiz.ques[i].Question)
		go timer(tc, *t)
		go handler(ans, quiz.ques[i].Solution)
		select {
		case <-tc:
			fmt.Print("\n\nTime's Up\nit wont count, but have a go:")
			sol := <-ans
			if quiz.ansCheckFunc(sol, quiz.ques[i].Solution) {
				notAtTimeScore++
			}
			fmt.Println()
			break
		case sol := <-ans:
			fmt.Print("\nYou're Fast\nGearing up next Question\n\n")
			<-tc
			if quiz.ansCheckFunc(sol, quiz.ques[i].Solution) {
				score++
			}
			break
		}
	}

	fmt.Printf("\nQuiz Complete\nYour Score : %d out of %d \nQuestion Correctly Answered(in time + not in time) :%d\n", score, quiz.Total, score+notAtTimeScore)
}
