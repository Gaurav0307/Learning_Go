package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "test.txt"
	var file *os.File
	var err error

	// Check if the file already exists
	if _, err := os.Stat(fileName); err == nil {
		// Open the existing file
		// Available flags: os.O_RDONLY = Read only, os.O_WRONLY = Write only, 
		// os.O_RDWR = Read and write, os.O_APPEND = Append to the file, 
		// os.O_CREATE = Create the file if it doesn't exist, 
		// os.O_TRUNC = Truncate the file, os.O_EXCL = Fail if the file exists
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error while opening file:", err)
			return
		}

		fmt.Println("File opened:", fileName)
	} else {
		// Create a new file
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("Error while creating file:", err)
			return
		}

		fmt.Println("File created:", fileName)
	}

	// Close the file
	defer file.Close()

	// Write the content to the file
	content := "Hi, this is a test file.\n"

	fmt.Println("Enter content to the file:")
	reader := bufio.NewReader(os.Stdin)
	content, _ = reader.ReadString('\n')

	if false {
		// Way 1
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error while writing file:", err)
			return
		}
	} else {
		// Way 2
		bytesWritten, err := io.WriteString(file, content)
		if err != nil {
			fmt.Println("Error while writing file:", err)
			return
		}

		fmt.Println("Bytes written:", bytesWritten)
	}

	fmt.Println("Content written to the file:", content)


	// Read the contents of the file
	if false {
		// Way 1
		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error while reading file:", err)
			return
		}

		fmt.Println("File content:", string(fileContent))
	} else {
		// Way 2

		// Seek to the beginning of the file
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create a buffer to store the file content
		buffer := make([]byte, 1024)
		// The bytesRead variable to store the number of bytes read
		var bytesRead int

		// Read the file content
		for {
			bytesRead, err = file.Read(buffer)	

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("Error while reading file:", err)
				return
			}

			fmt.Println("Bytes read:", bytesRead)
		}

		fmt.Println("File content:", string(buffer))
	}
}
