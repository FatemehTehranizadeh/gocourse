package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buff := make([]byte, 1024)
	n, err := r.Read(buff)
	if err != nil {
		log.Fatalln("Error in reading: ", err)
	}
	fmt.Printf("%d bytes have been read successfully and the content is %v\n", n, string(buff[:n]))
}

func writeFromWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error in writing: ", err)
	}

}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error in closign the file: ", err)
	}
}

func bufferWriteExample() {
	var buff bytes.Buffer
	n, err := buff.WriteString("An example text to write into the buffer")
	if err != nil {
		log.Fatalln("Error in writing: ", err)
	}
	fmt.Printf("%d bytes have been written successfully and the content is %v\n", n, buff.String())
}

func bufferReadExample() {
	var buff bytes.Buffer
	_, err := buff.Write([]byte("An example text to write into the buffer"))
	if err != nil {
		log.Fatalln("Error in writing: ", err)
	}
	storingBuffer := make([]byte, 1024)
	n, err := buff.Read(storingBuffer)
	if err != nil {
		log.Fatalln("Error in reading: ", err)
	}
	fmt.Printf("%d bytes have been read successfully and the content is %v\n", n, string(storingBuffer))
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello!")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	// var buff bytes.Buffer
	buff := new(bytes.Buffer)
	n, err := buff.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error in reading: ", err)
	}
	fmt.Printf("%d bytes have been read successfully and the content is %v\n", n, buff.String())
}

func pipeExample() {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		pipeWriter.Write([]byte("Hello Pipe!"))
		pipeWriter.Close()
	}()
	buff := make([]byte, 1024)
	pipeReader.Read(buff)
	fmt.Println("The read content from pipe is: ", string(buff))

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(pipeReader)
	// fmt.Println("The read content from pipe is: ", buf.String())

}

func main() {

	// r := "hello world!"
	// readFromReader(strings.NewReader(r))

	exampleFile, err := os.OpenFile("intermediate/exampleFile.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("Error in opening the file: ", err)
	}
	defer closeResource(exampleFile)

	// // using a bytes.Buffer as a writer
	// buff1 := new(bytes.Buffer)
	// writer1 := io.Writer(buff1)
	// writer1.Write([]byte("Heelo"))
	// fmt.Println(buff1)

	// Using a file as a writer
	// writer2 := io.Writer(exampleFile)
	// writer2.Write([]byte("\nHeelo Woooorld!"))

	go func() {
		writer2 := io.Writer(exampleFile)
		writer2.Write([]byte("\nHeelo Woooorld!"))
	}()

	// Using a file as a Reader
	var buff2 bytes.Buffer
	reader1 := io.Reader(exampleFile)
	buff2.ReadFrom(reader1)
	fmt.Println("buff2:", buff2.String())

	// // Using a bytes.Buffer as a Reader
	// buff4 := new(bytes.Buffer)
	// buff4.WriteString("Hellooooooo")
	// buff5 := make([]byte, 1024)
	// reader2 := io.Reader(buff4)
	// reader2.Read(buff5)
	// fmt.Println("buff5:", string(buff5))

	// readFromReader(exampleFile)

	// content := "\nAnother example data for testing the write function1"
	// writeFromWriter(exampleFile, content)

	// var buff bytes.Buffer
	// writeFromWriter(&buff, "hello!!!!!")
	// fmt.Println(buff.String())
	// multiReaderExample()
	// pipeExample()
}
