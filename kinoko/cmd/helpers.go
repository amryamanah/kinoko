package cmd

import (
	"log"
	"os"
)

func Fatal(v ...interface{}) {
	log.Fatal("Error:", v)
	os.Exit(1)
}