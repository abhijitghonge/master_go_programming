/////////////////////////////////
// Reading Files in Go
// Go Playground: https://play.golang.org/p/LJnTSVfaJW_R
/////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	myFile, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	bw := bufio.NewWriter(myFile)

	_, err = bw.WriteString("I love Golang!")
	if err != nil {
		log.Fatal(err)
	}

	bw.Flush()
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	//** READING INTO A BYTE SLICE USING io.ReadFull() **//

	// Opening the file in read-only mode. The file must exist (in the current working directory)
	// Use a valid path!
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// declaring a byte slice and initializing it with a length of 2
	byteSlice := make([]byte, 2)

	// io.ReadFull() returns an error if the file is smaller than the byte slice.
	// it reads the file into the byte slice up to its length
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	fmt.Println(strings.Repeat("#", 20))

	//** READING WHOLE FILE INTO A BYTESLICE USING ioutil.ReadAll() **//

	// Opening another file (from the current working directory)
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	// ioutil.ReadAll() reads every byte from the file and return a slice of unknown size
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	//** READING WHOLE FILE INTO MEMORY USING ioutil.ReadFile() **//

	// ioutil.ReadFile() reads a file into byte slice
	// this function handles opening and closing the file.
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)

	// fileInfo, err := os.Stat("info.txt")

	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal(err)
	// 	}
	// }

	// err = os.Rename(fileInfo.Name(), "new.txt")

	// if err != nil {
	// 	log.Println(err)
	// }

	file, err = os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	// ioutil.ReadAll() reads every byte from the file and return a slice of unknown size
	data, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	file, err = os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
