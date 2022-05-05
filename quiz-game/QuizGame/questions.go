package QuizGame

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type QuizInterface interface {
	ReadFile(file *string, opts ...string) (QuestionSlice, error)
	SetQuestion(QuestionSlice)
	GetQuestion() QuestionSlice
	GetTotal() int
	DisplayQuesFunc(ques string)
	CheckAnsFunc(submit string, sol string) bool
}

type Quiz struct {
	Ques  []QuestionSet
	Total int
	// displayQuesFunc func(ques string)
	// ansCheckFunc    func(submit string, ans string) bool
}

type QuestionSet struct {
	Question string
	Solution string
}

type QuestionSlice []QuestionSet

// type QuizOption func(q *Quiz)

// func WithDisplayQuesFunc(fn func(ques string)) QuizOption {
// 	return func(q *Quiz) {
// 		q.displayQuesFunc = fn
// 	}
// }

// func WithAnswerCheckFunc(fn func(submit string, sol string) bool) QuizOption {
// 	return func(q *Quiz) {
// 		q.ansCheckFunc = fn
// 	}
// }

func (q *Quiz) DisplayQuesFunc(ques string) {
	fmt.Print(ques, ":")
}

func (q *Quiz) CheckAnsFunc(submit string, sol string) bool {
	if submit == sol {
		return true
	}
	return false
}

func (q *Quiz) SetQuestion(ques QuestionSlice) {
	q.Ques = ques
	q.Total = len(ques)
}

func (q *Quiz) GetQuestion() QuestionSlice {
	return q.Ques
}

func (q *Quiz) GetTotal() int {
	return q.Total
}

func (q *Quiz) ReadFile(file *string, opts ...string) (QuestionSlice, error) {
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
	ques := make([]QuestionSet, 0)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), sep)
		// if len(s) != 2 {
		// 	return Quiz{}, errors.New("invalid format")
		// }
		if len(s) == 2 {
			temp := QuestionSet{s[0], s[1]}
			ques = append(ques, temp)
			total++
		}
	}

	return ques, nil
}
