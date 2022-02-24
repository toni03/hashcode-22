package main

import (
	"bufio"
	"os"
)

func writeFullFile(filename string, data []byte) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	_, err = f.Write(data)
	check(err)
	check(f.Sync())
}

type writeChunk = func([]byte)
type closeFile = func()

func writeFileInChunks(filename string) (wc writeChunk, cf closeFile) {
	f, err := os.Create(filename)
	check(err)

	w := bufio.NewWriter(f)
	wc = func(data []byte) {
		_, err := w.Write(data)
		check(err)
	}
	cf = func() {
		check(w.Flush())
		check(f.Sync())
		f.Close()
	}
	return
}
