package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go language"
	writer, err := os.Create("./input.txt")
	CheckError(err)
	length, err := io.WriteString(writer, content)
	CheckError(err)
	fmt.Printf("%v character writtern to file\n", length)
	defer func(writer *os.File) {
		err := writer.Close()
		if err != nil {
			panic(any(err))
		}
	}(writer)
	defer read("./input.txt")
}

func CheckError(err error) {
	if err != nil {
		panic(any(err))
	}
}

func read(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	CheckError(err)
	fmt.Println(string(data))
}
