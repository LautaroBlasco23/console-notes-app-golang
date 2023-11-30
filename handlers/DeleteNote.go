package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"take-notes/initializers"
)


func DeleteNote() {
  // Get file access
  bytes, err := os.ReadFile("data.csv")
  if err != nil {
    fmt.Println("X -> error reading file")
  }

  // Read index input
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("-> Insert the node index you want to delete: ")
  index, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("X -> error reading input")
  }
  
  // Parse input
  index_int, err := strconv.ParseInt(index[:len(index)-1], 10, 0)
  if err != nil {
    fmt.Println("X -> invalid input")
    return
  }

  // get list of notes and pop the last element (empty)
  list_of_notes := strings.Split(string(bytes), "\n")
  list_of_notes = list_of_notes[:len(list_of_notes)-1]

  // if index is right, remove the note
  if index_int > 0 && int(index_int) < len(list_of_notes)+1 {
    // Restart file
    initializers.RestartFile()
    // Write notes in restarted file
    for i, note := range list_of_notes {
      if i+1 != int(index_int) {
        if i < len(list_of_notes) {
          initializers.DataFile.Write([]byte(note + "\n"))
        } else {
          initializers.DataFile.Write([]byte(note))
        }
      }
    }
  } else {
    fmt.Println("X -> invalid index")
  }
}
