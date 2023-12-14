package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// One way would be to split the text based
// on the colon char ':' but if there is a link included such as http://, this would split
// a value despite not intending to split. One way to resolve this issue could be to
// url-encode each value on the client side or in our code. For now I will url/encode beforehand to get
// a running mock going
func main() {
	headers := readFile()
	createExpected(headers)

}

// readFile will open and read the text input. Each line corresponds
// to a key-value pair/singular header line.
func readFile() string {
	headers := ""
	readFile, err := os.Open("./RawTxt.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		headers += fmt.Sprintf("%s\n", convertToGo(fileScanner.Text()))
	}
	return headers
}

func convertToGo(line string) string {
	var (
		format = "\"%s\":{\"%s\"},"
	)
	splitTxt := strings.Split(line, ":")
	splitTxt[1] = strings.TrimSpace(splitTxt[1])

	return fmt.Sprintf(format, splitTxt[0], splitTxt[1])
}

func createExpected(headers string) {
	var (
		format = "req.Header = map[string][]string{%s}"
	)
	fmt.Printf(format, headers)
}
