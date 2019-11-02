package parser

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Question ,
type Question struct {
	Number        int
	Statement     string
	Code          string
	Options       map[string]string
	CorrectAnswer string
	Explanation   string
}

// GetQuestions returns a []Question with all the questions in it
func GetQuestions() []Question {
	data := getData()
	return parseGlob(data)
}

// getData will fetch data from the readme file and return it as a string
func getData() string {
	response, err := http.Get("https://rawcdn.githack.com/lydiahallie/javascript-questions/13b42090852a397f878da569cdd762bbd610f73f/en-EN/README.md")

	if err != nil {
		log.Fatal(err)
	}

	readmeContent, _ := ioutil.ReadAll(response.Body)
	return string(readmeContent)
}

// parseGlob parses data accordingly using the helper functions
func parseGlob(data string) []Question {
	dataStringArray := strings.Split(data, "---")
	dataStringArray = dataStringArray[1:]

	var questions []Question

	for questionNumber, questionBlock := range dataStringArray {
		var q Question

		q.Number = questionNumber + 1
		q.Statement = getQuestionStatement(questionBlock)
		q.Code = getQuestionCode(questionBlock)
		q.Options = getQuestionOptions(questionBlock)
		q.CorrectAnswer = getQuestionCorrectAnswer(questionBlock)
		q.Explanation = getQuestionExplanation(questionBlock)

		questions = append(questions, q)
	}

	return questions
}

// Helper Functions

func getQuestionStatement(data string) string {
	n := 0

	for i, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "######") {
			n = i
			break
		}
	}

	statement := strings.Split(data, "\n")[n]
	statement = strings.Split(statement, ".")[1]
	statement = strings.TrimSpace(statement)

	return statement
}

func getQuestionCode(data string) string {
	code := ""

	if strings.Contains(data, "```javascript") {
		code = strings.Split(data, "```javascript")[1]
		code = strings.Split(code, "```")[0]
	}

	return code
}

func getQuestionOptions(data string) map[string]string {
	finalOptionsMap := make(map[string]string)

	optionsArr := strings.Split(strings.Split(data, "<details>")[0], "\n")

	for _, o := range optionsArr {
		if len(o) != 0 && string(o[0]) == "-" {
			optionAlphabet := strings.Split(strings.Split(o, ":")[0], " ")[1]
			optionText := strings.TrimSpace(o[4:])

			finalOptionsMap[optionAlphabet] = optionText
		}
	}

	return finalOptionsMap
}

func getQuestionCorrectAnswer(data string) string {
	correctAnswer := strings.Split(data, "<p>")[1]
	correctAnswer = strings.Split(correctAnswer, "</p>")[0]
	correctAnswer = strings.Split(correctAnswer, "Answer: ")[1]
	correctAnswer = string(correctAnswer[0])

	return correctAnswer
}

func getQuestionExplanation(data string) string {
	explanation := strings.Split(data, "<p>")[1]
	explanation = strings.Split(explanation, "</p>")[0]

	return explanation
}
