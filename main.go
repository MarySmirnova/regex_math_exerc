package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	inFile, _, err := enterData()
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadFile(inFile)
	if err != nil {
		panic(err)
	}

	revealMath(input)

}

func enterData() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите название файла с данными: ")
	inFile, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Println("Введите название файла для вывода данных: ")
	outFile, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(inFile), strings.TrimSpace(outFile), nil
}

func revealMath(input []byte) {
	mathRegex := regexp.MustCompile(`([0-9]+)([\+\-\\*])([0-9]+)+`)

	paths := mathRegex.FindAllStringSubmatch(string(input), -1)

	fmt.Println(paths)

}
