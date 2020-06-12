package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := flag.String("file", "/home/desferreira/Documents/Estudos/go/go-projects/questions.csv", "Seu arquivo aqui!")
	flag.Parse()

	content := readCsv(*file)

	QA := getQuestionsAndAnswers(content)

	reader := bufio.NewReader(os.Stdin)

	points := 0

	for i, v := range QA {
		fmt.Printf("Qual a resposta de %v ", i)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text != v {
			fmt.Println("Resposta errada!")
		} else {

			fmt.Println("Resposta correta!")
			points++

		}
	}
	fmt.Printf("Seu resultado final é: %v acertos de %v chances", points, len(QA))
}

func getQuestionsAndAnswers(content [][]string) map[string]string {

	QA := map[string]string{}

	for i := 0; i < len(content); i++ {
		QA[strings.Split(content[i][0], " ")[0]] = strings.Split(content[i][1], " ")[0]
	}

	return QA

}

func readCsv(path string) [][]string {

	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return lines
}
