package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'qustion, answer'")
	timeLimit := flag.Int("timer", 30, "Timer for quiz to run in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file : %s", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to read CSV file : %s", *csvFileName))
	}
	problems := parseLines(lines)
	correctAnswer := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d : %s=", i+1, p.ques)
		answerCh := make(chan string)
		//GO routine
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		//GO Case
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correctAnswer, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.ans {
				correctAnswer++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correctAnswer, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	ques string
	ans  string
}

func exit(msg string) {
	log.Fatalln(msg)
	os.Exit(1)
}
