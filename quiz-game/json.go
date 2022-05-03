package main

import (
	"encoding/json"
	"fmt"
	"os"
	"quiz/QuizGame"
	"strings"
)

type JsonHanlder struct {
	Ques  []QuizGame.QuestionSet
	Total int
}

type Ques struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func (js *JsonHanlder) ReadFile(file *string, opts ...string) (QuizGame.QuestionSlice, error) {
	f, err := os.Open(*file)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	var ques []Ques
	err = dec.Decode(&ques)
	if err != nil {
		return nil, err
	}
	// fmt.Println(ques)
	questions := make([]QuizGame.QuestionSet, 0)

	for _, q := range ques {
		question := q.Question
		for _, op := range q.Options {
			question = question + ":" + op
		}
		temp := QuizGame.QuestionSet{}
		temp.Question = question
		temp.Solution = q.Answer
		questions = append(questions, temp)

	}
	// fmt.Println(questions)

	return questions, nil

}

func (js *JsonHanlder) SetQuestion(ques QuizGame.QuestionSlice) {
	js.Ques = ques
	js.Total = len(ques)
}

func (js *JsonHanlder) GetQuestion() QuizGame.QuestionSlice {
	return js.Ques
}
func (js *JsonHanlder) GetTotal() int {
	return js.Total
}
func (js *JsonHanlder) DisplayQuesFunc(ques string) {
	qs := strings.Split(ques, ":")
	for _, v := range qs {
		fmt.Println(v)
	}

}
func (js *JsonHanlder) CheckAnsFunc(submit string, sol string) bool {
	if strings.ToLower(submit) == strings.ToLower(sol) {
		return true
	}
	return false
}

// func main() {
// 	s := "./QuizGame/json/proble.json"
// 	ReadFile(&s)
// }
