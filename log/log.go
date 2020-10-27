package log

import "log"

func Fatal(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
