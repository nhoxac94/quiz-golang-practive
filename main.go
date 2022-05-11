package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "file csv content question and answer")
	timeLimit := flag.Int("limit", 30, "limit for resolve problems")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("%s file have problem", *fileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse file")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0
	for i, p := range problems {
		fmt.Printf("#%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYour score %d out of %d.\n ", score, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				score++
			}
		}
	}
	fmt.Printf("\nYour score %d out of %d.\n ", score, len(problems))

}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
