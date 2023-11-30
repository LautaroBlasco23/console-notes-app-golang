package handlers

import (
	"bufio"
	"fmt"
	"os"
	"take-notes/initializers"
)

func AddNote() {
  reader := bufio.NewReader(os.Stdin)
  // Reading Title
  fmt.Print("-> Add title: ")
  title, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("X -> error reading input")
  }
  // remove \n from the string
  title = title[0:len(title) - 1]

  // Reading Body
  fmt.Print("-> Add body: ")
  // Read user input 
  body, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("X -> error reading input")
  }
  // remove \n from the string
  body = body[0:len(body) - 1]

  // Write body and title, with , to do a csv
  initializers.DataFile.Write([]byte(title))
  initializers.DataFile.Write([]byte(","))
  initializers.DataFile.Write([]byte(body))
  initializers.DataFile.Write([]byte("\n"))
}
