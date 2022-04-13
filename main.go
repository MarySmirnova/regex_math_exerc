package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inFileName, outFileName, err := enterData()
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadFile(inFileName)
	if err != nil {
		panic(err)
	}

	err = writeToFile(outFileName, calculate(findMath(input)))
	if err != nil {
		panic(err)
	}
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

func findMath(input []byte) [][]string {
	mathRegex := regexp.MustCompile(`([0-9]+)([\+\-/*])([0-9]+)=`)

	exprs := mathRegex.FindAllStringSubmatch(string(input), -1)

	return exprs
}

func calculate(exprs [][]string) []string {
	var output []string

	for _, ex := range exprs {
		if len(ex) < 4 {
			continue
		}

		//не проверяем ошибки т.к. регулярка пропустит сюда только числа
		val1, _ := strconv.Atoi(ex[1])
		val2, _ := strconv.Atoi(ex[3])

		total := ex[0]

		var res int
		switch ex[2] {
		case "+":
			res = val1 + val2
		case "-":
			res = val1 - val2
		case "*":
			res = val1 * val2
		case "/":
			if val2 == 0 {
				total = total + "can't divide by zero"
				output = append(output, total)
				continue
			}
			res = val1 / val2
		default:
			continue
		}

		total = total + strconv.Itoa(res)
		output = append(output, total)
	}

	return output
}

func writeToFile(filename string, content []string) error {
	outFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	for _, c := range content {
		writer.Write([]byte(c + "\n"))
	}
	writer.Flush()

	return nil
}
