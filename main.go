package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"take-notes/handlers"
	"take-notes/initializers"
)

func main() {
  // Create data.txt file and let handlers use it.
  initializers.LoadDataFile()

  list_of_options := `---- OPTIONS ----
0) Exit program
1) Add new note
2) Modify note
3) Delete note
4) Show notes
9) Restart notes
-> Enter your choice: `

  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(list_of_options)
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("X -> error reading input")
    }
    text = text[0:len(text) - 1]

    option, err := strconv.ParseInt(text, 10, 0)
    if err != nil {
      fmt.Println("X -> Invalid option")
    } else {
      switch option {
        case 0:
          os.Exit(1)
        case 1:
          handlers.AddNote()
        case 2:
          handlers.EditNote()
        case 3:
          handlers.DeleteNote()
        case 4:
          handlers.ShowNotes()
        case 9:
          initializers.RestartFile()
        default:
          fmt.Println("X -> Invalid option, please insert a valid number")
      }
      fmt.Println()
    }
  }
}
