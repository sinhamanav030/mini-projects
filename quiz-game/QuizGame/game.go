package QuizGame

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

	fp, err := os.ReadFile(*f)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(strings.NewReader(string(fp)))
	// fmt.Println(r)
	var ans string

	ques, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	total = len(ques)
	fmt.Println("Quiz is Ready! Are you ready to go:\nPress enter to start")
	fmt.Scanf("%s")
	tc := make(chan bool)
	go Timer(tc, *t)
	go handler(tc)
	for _, v := range ques {
		fmt.Print(v[0], " :")
		fmt.Scanf("%s\n", &ans)
		if v[1] == ans {
			score++
		}
	}
	fmt.Printf("\nQuiz Complete\nYour Score : %d out of %d \n", score, total)
}
