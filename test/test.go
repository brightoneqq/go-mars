package test

import (
	"github.com/robfig/cron"
	"log"
)

func Test() {
	c := cron.New()
	schedule:= "* * * * * ?"
	c.AddFunc(schedule, func() {
		log.Print("This is Cron Demo Test")
	})
}