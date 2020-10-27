package logger

import "log"

func Handle(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
