package QuizGame

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	ques  []QuestionSet
	Total int
}

type QuestionSet struct {
	Question string
	Solution string
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

func ReadFile(file *string) (Quiz, error) {
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

	return Quiz{ques, total}, nil

}
