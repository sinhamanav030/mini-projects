package QuizGame

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
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
	rand.Seed(time.Now().Unix())
	mp := make(map[int]bool, len(quiz.ques))
	id := make([]int, len(quiz.ques))
	for i := 0; i < len(quiz.ques); i++ {
		for {
			tmp := rand.Intn(len(quiz.ques))
			if _, ok := mp[tmp]; ok == false {
				mp[tmp] = true
				id[i] = tmp
				break
			}
		}
	}
	// fmt.Println(id)
	fmt.Println("Quiz is Ready! Are you ready to go:\nPress enter to start")
	fmt.Scanf("%s")
	for _, i := range id {
		fmt.Print(quiz.ques[i].Question, ":")
		go timer(tc, *t)
		go handler(ans, quiz.ques[i].Solution)
		select {
		case <-tc:
			fmt.Print("\n\nTime's Up\nit wont count, but have a go:")
			sol := <-ans
			if sol == quiz.ques[i].Solution {
				notAtTimeScore++
			}
			fmt.Println()
			break
		case sol := <-ans:
			fmt.Print("\nYou're Fast\nGearing up next Question\n\n")
			<-tc
			if sol == quiz.ques[i].Solution {
				score++
			}
			break
		}
	}

	fmt.Printf("\nQuiz Complete\nYour Score : %d out of %d \nQuestion Correctly Answered(in time + not in time) :%d\n", score, quiz.Total, score+notAtTimeScore)
}
