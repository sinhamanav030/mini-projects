package QuizGame

import (
	"fmt"
	"math/rand"
	"time"
)

var score, total int

func init() {
	score = 0
	total = 0
}

func timer(ch chan bool, timer int) {
	if timer < 1 {
		timer = 1
	}
	time.Sleep(time.Duration(timer) * time.Second)
	ch <- true
}

func handler(ch chan string, solution string) {
	var ans string
	fmt.Scanf("%s", &ans)
	ch <- ans
}

func generateOrder(length int) []int {
	rand.Seed(time.Now().Unix())
	mp := make(map[int]bool, length)
	id := make([]int, length)
	for i := 0; i < length; i++ {
		for {
			tmp := rand.Intn(length)
			if _, ok := mp[tmp]; ok == false {
				mp[tmp] = true
				id[i] = tmp
				break
			}
		}
	}
	return id
}

func Game(quiz QuizInterface, time int) {

	//channel for timer
	tc := make(chan bool)

	//channel for answers
	ans := make(chan string)

	score, notAtTimeScore := 0, 0

	ques := quiz.GetQuestion()

	id := generateOrder(len(ques))

	fmt.Println("Quiz is Ready! Are you ready to go:\nPress enter to start")
	fmt.Scanf("%s")

	for _, i := range id {
		quiz.DisplayQuesFunc(ques[i].Question)
		go timer(tc, time)
		go handler(ans, ques[i].Solution)
		select {
		case <-tc:
			fmt.Print("\n\nTime's Up\nit wont count, but have a go:")
			sol := <-ans
			if quiz.CheckAnsFunc(sol, ques[i].Solution) {
				notAtTimeScore++
			}
			fmt.Println()
			break
		case sol := <-ans:
			fmt.Print("\nYou're Fast\nGearing up next Question\n\n")
			<-tc
			if quiz.CheckAnsFunc(sol, ques[i].Solution) {
				score++
			}
			break
		}
	}

	fmt.Printf("\nQuiz Complete\nYour Score : %d out of %d \nQuestion Correctly Answered(in time + not in time) :%d\n", score, quiz.GetTotal(), score+notAtTimeScore)
}
