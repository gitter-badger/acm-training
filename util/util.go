package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

// CheckArgument Checks the condition matches, otherwise the program will
// close according to the error. On the other hand, the cusotmized
// error message will be shown before exiting the program.
func CheckArgument(condition bool, messages ...interface{}) {
	if !condition {
		log.Fatal(messages...)
		os.Exit(1)
	}
}

// CheckIOError Checks the IO error thrown by IO library reading or writing.
// This function will exit the program when EOF was read from the input,
// or any unexcepted IO exception thrown.
func CheckIOError(err error) {
	if err != nil {
		if err == io.EOF {
			fmt.Println("Interrupted, canceled and existing the tool.")
			os.Exit(1)
		} else {
			log.Fatal(err)
		}
	}
}
