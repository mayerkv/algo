package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func generateInput(fileName string, n, m int) {
	src := rand.NewSource(123)
	rnd := rand.New(src)

	in := openFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	defer in.Close()

	writer := bufio.NewWriter(in)
	for i := 0; i < n; i++ {
		writer.WriteString(fmt.Sprintln(rnd.Intn(m-1) + 1))
	}
	writer.Flush()
}

func openFile(fileName string, flag int) *os.File {
	file, err := os.OpenFile(fileName, flag, 0600)
	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
