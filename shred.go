package main

import (
  "fmt"
  "os"
  "time"
  "math/rand"
 )

func main() {
  // Seed the random number generator
  rand.Seed(time.Now().UnixNano())

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a command-line argument.")
		return
	}

  fileName := args[1]
  
  shredFile(fileName) // why do this three times?

  err := os.Remove(fileName)
  if err != nil {
    fmt.Println("Error deleting file:", err)
    return
  }
}


func shredFile(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Error getting file information:", err)
		return
	}

	fileSize := fileInfo.Size()
	
  file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
  if err != nil {
      fmt.Println("Error opening file:", err)
      return
  }
  defer file.Close()

	fmt.Printf("Shredding: %s, %d bytes\n", fileName, fileSize)
	
  // Write to the file
  data := []byte("X")
  var i int64
  for i = 1; i <= fileSize; i++ {
    data[0] = byte(32 + rand.Intn(96)); // keep the random characters printable for the moment
    _, err := file.Write(data)
    if err != nil {
      panic("Error writing to file")
    }
  }

  // Flush the buffer to ensure all data is written to the file
  err = file.Sync()
  if err != nil {
    panic("Error flushing buffer")
  }
}
