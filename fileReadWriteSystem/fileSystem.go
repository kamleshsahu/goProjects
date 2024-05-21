package main

import (
	"fmt"
)

func main() {
	bs := NewBlockService("vfs")

	// Create a new file
	fd, err := bs.FOpen("test.txt", "w")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Write data to the file
	data := []byte("Hello, Virtual File System!")
	err = bs.FWrite(fd, data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = bs.FClose(fd)
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}

	// Open the file for reading
	fd, err = bs.FOpen("test.txt", "r")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Read data from the file
	readData, err := bs.FRead(fd, len(data)+64)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Read data:", string(readData))

	// Close the file
	err = bs.FClose(fd)
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}
}
