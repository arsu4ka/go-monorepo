package app2

import (
	"io"
	"log"
)

func safeCloseBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		log.Println("Error closing request body: ", err)
	}
}
