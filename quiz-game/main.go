package main

import (
	"flag"
	"quiz/QuizGame"
)

func main() {
	f := flag.String("fp", "./QuizGame/CSV/problems.csv", "Get problem file path")
	t := flag.Int("timer", 5, "Set timer for quiz")
	flag.Parse()

	// var quiz QuizGame.Quiz
	// ques, err := quiz.ReadFile(f, ",")
	// if err != nil {
	// 	panic(err)
	// }

	// quiz.SetQuestion(ques)

	// QuizGame.Game(&quiz, *t)

	var mcq McqQuiz

	ques, err := mcq.ReadFile(f, ",")
	if err != nil {
		panic(err)
	}

	mcq.SetQuestion(ques)

	QuizGame.Game(&mcq, *t)

}
