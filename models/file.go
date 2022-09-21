package models

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type InchSources struct {
	data string
}

func (fs *InchSources) check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func (fs *InchSources) parseFile() {

	dat, err := os.ReadFile("./files/inch.json")
	fs.check(err)
	fmt.Print(string(dat))

	f, err := os.Open("./files/inch.json")
	fs.check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	fs.check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	fs.check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fs.check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	fs.check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fs.check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	fs.check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	fs.check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
