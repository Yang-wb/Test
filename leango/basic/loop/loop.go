package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)

	//for true  {
	//	fmt.Println(1)
	//}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func printFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func main() {
//	fmt.Println(convertToBin(5))
//}
