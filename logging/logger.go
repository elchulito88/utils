package logging

import "log"

// Log error tracking
func Log(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
