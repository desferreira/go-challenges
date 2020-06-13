package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	file := flag.String("file", "./questions.csv", "Your file here!")
	timeLimit := flag.Int("limit", 30, "Limit time in seconds")
	flag.Parse()

	content := readCsv(*file)

	QA := getQuestionsAndAnswers(content)

	points := 0

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))

	for _, v := range QA {
		fmt.Printf("%v = ", v.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\n ========= Time is up! Your final result is: %v out of %v =========\n", points, len(QA))
			return
		case answer := <-answerCh:
			if answer != v.a {
				fmt.Println("Wrong answer!")
			} else {
				fmt.Println("Correct answer!")
				points++
			}
		}
	}
}

func getQuestionsAndAnswers(content [][]string) []problem {

	QA := make([]problem, len(content))

	for i, lines := range content {
		QA[i] = problem{
			lines[0], strings.TrimSpace(lines[1]),
		}
	}

	return QA

}

func readCsv(path string) [][]string {

	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Can't open file: %v ", path)
		os.Exit(1)
	}

	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("Can't parse the file: %v ", path)
		os.Exit(1)
	}

	return lines
}
