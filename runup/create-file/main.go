package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("dummy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	example := "This is dummy text"

	file.WriteString(example)
	file.Seek(0, 0)

	// read whole file at once
	// fileToRead, err := os.ReadFile("dummy.txt")
	// fmt.Println(string(fileToRead))

	// Token by token or line by line read
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fmt.Println(scanner.Text())

	// //Streaming reader, good for large files and reading until specific delimiter
	// reader := bufio.NewReader(file)

	// for {
	// 	line, err := reader.ReadString('\n')

	// 	if err == io.EOF {
	// 		fmt.Print(line)
	// 		break
	// 	}

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Print(line)
	// }

	// io read
	// data, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

}
