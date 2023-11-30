package initializers

import (
	"errors"
	"fmt"
	"os"
)

var DataFile *os.File

func LoadDataFile() {
  // Create the file and link it to the variable. 
  var err error
  if _, err = os.Stat("data.csv"); errors.Is(err, os.ErrNotExist) {
    // If file doesn't exist
    DataFile, err = os.Create("data.csv")
    if err != nil {
      fmt.Println("X -> error creating data file")
      fmt.Println(err)
    }
  } else {
    // If file exist
    DataFile, err = os.OpenFile("data.csv", os.O_RDWR|os.O_APPEND, 0644)
  }
}

func RestartFile() {
  // Remove the file
  err := os.Remove("data.csv")
  if err != nil {
    fmt.Println("X -> error removing file")
  }

  // Create the file again
  LoadDataFile() 
}
