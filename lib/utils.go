package lib

import (
	"io"
	"log"
)

//Close closes IO interfaces and handle error
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Printf("io close: %v", err)
	}
}
