package QuizGame

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	ques            []QuestionSet
	Total           int
	displayQuesFunc func(ques string)
	ansCheckFunc    func(submit string, ans string) bool
}

type QuestionSet struct {
	Question string
	Solution string
}

type QuizOption func(q *Quiz)

func WithDisplayQuesFunc(fn func(ques string)) QuizOption {
	return func(q *Quiz) {
		q.displayQuesFunc = fn
	}
}

func WithAnswerCheckFunc(fn func(submit string, sol string) bool) QuizOption {
	return func(q *Quiz) {
		q.ansCheckFunc = fn
	}
}

func defaultDisplayQuesFunc(ques string) {
	fmt.Print(ques, ":")
}

func defaultAnswerCheckFunc(submit string, sol string) bool {
	if submit == sol {
		return true
	}
	return false
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

func NewQuiz(file *string, opts ...QuizOption) (Quiz, error) {
	f, err := os.Open(*file)
	sep := ","
	if err != nil {
		panic(err)
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

	q := Quiz{ques, total, defaultDisplayQuesFunc, defaultAnswerCheckFunc}

	for _, opt := range opts {
		opt(&q)
	}

	return q, nil
}
