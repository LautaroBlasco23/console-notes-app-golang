package handlers

import (
	"fmt"
	"os"
	"strings"
)

func ShowNotes() {
  notes, err := os.ReadFile("data.csv")
  if err != nil {
    fmt.Println("X -> error reding files")
  }
  
  notes_arr := strings.Split(string(notes), "\n")
  fmt.Println()
  fmt.Println("---- LIST OF NOTES ----")

  for i, note := range notes_arr {
    if i < len(notes_arr) - 1 {
      fmt.Printf("* note(%v): %s", i+1, note)
      fmt.Println()
    }
  }
  
}
