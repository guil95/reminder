package main

import (
	"flag"
	"fmt"

	"github.com/martinlindhe/notify"
	"github.com/robfig/cron/v3"
)

var timer string
var text string

func init() {
	flag.StringVar(&timer, "timer", "10m", "timer cron")
	flag.StringVar(&text, "text", "Reminder time", "text of notification")
}

func main() {
	flag.Parse()

	c := cron.New()

	c.AddFunc(fmt.Sprintf("@every %s", timer), func() {
		notify.Alert("Reminder", "", text, "glass.png")
	})

	c.Run()
}
