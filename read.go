package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func readFullFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	check(err)
	return data
}

func readFileInChunks(filePath string, from int64, bytesToRead int) []byte {
	f, err := os.Open(filePath)
	defer f.Close()
	check(err)

	f.Seek(from, 0)
	chunk := make([]byte, bytesToRead)
	_, err = f.Read(chunk)
	check(err)

	return chunk
}

func readFileByLine(filePath string) []byte {
	f, err := os.Open(filePath)
	defer f.Close()
	check(err)

	reader := bufio.NewReader(f)

	line := make([]byte, 0)
	for {
		segment, isPrefix, err := reader.ReadLine()
		check(err)

		line = append(line, segment...)

		if !isPrefix {
			break
		}
	}

	return line
}
