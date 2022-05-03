package main

import (
	"bufio"
	"fmt"
	"os"
	"quiz/QuizGame"
	"strings"
)

type McqQuiz struct {
	Ques  []QuizGame.QuestionSet
	Total int
}

func (q *McqQuiz) DisplayQuesFunc(ques string) {
	qs := strings.Split(ques, ":")
	for _, v := range qs {
		fmt.Println(v)
	}
}

func (q *McqQuiz) CheckAnsFunc(submit string, sol string) bool {
	if strings.ToLower(submit) == strings.ToLower(sol) {
		return true
	}
	return false
}

func (q *McqQuiz) SetQuestion(ques QuizGame.QuestionSlice) {
	q.Ques = ques
	q.Total = len(ques)
}

func (q *McqQuiz) GetQuestion() QuizGame.QuestionSlice {
	return q.Ques
}

func (q *McqQuiz) GetTotal() int {
	return q.Total
}

func (q *McqQuiz) ReadFile(file *string, opts ...string) (QuizGame.QuestionSlice, error) {
	f, err := os.Open(*file)
	var sep string
	if len(opts) > 0 {
		sep = opts[0]
	} else {
		sep = ","
	}
	if err != nil {
		return nil, err
	}

	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	ques := make([]QuizGame.QuestionSet, 0)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), sep)
		// if len(s) != 2 {
		// 	return Quiz{}, errors.New("invalid format")
		// }
		if len(s) == 2 {
			temp := QuizGame.QuestionSet{}
			temp.Question = s[0]
			temp.Solution = s[1]
			ques = append(ques, temp)
			total++
		}
	}

	return ques, nil
}
