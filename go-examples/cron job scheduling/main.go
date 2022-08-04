package main

import (
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

// func init() {
// 	log.SetLevel(log.InfoLevel)
// 	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
// }

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { log.Info("Every Second\n") })

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())
	time.Sleep(10 * time.Second)

	// Funcs may also be added to a running Cron
	log.Info("Add new job to a running cron")
	_ = c.AddFunc("*/2 * * * *", func() { log.Info("Every Two Second\n") })
	printCronEntries(c.Entries())

	time.Sleep(time.Second * 20)
}

func printCronEntries(cronEntries []*cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}
