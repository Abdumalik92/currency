package main

import (
	"curr/db"
	"curr/jobs"
	"curr/routes"
	"curr/utils"
	"github.com/jasonlvhit/gocron"
)

func runScheduler() {
	scheduler := gocron.NewScheduler()

	scheduler.Every(5).Minutes().Do(jobs.NPCRBank)
	scheduler.Every(5).Minutes().Do(jobs.C2CBank)

	<-scheduler.Start()
}

func main() {
	utils.ReadSettings()
	db.StartDbConnection()
	go runScheduler()
	routes.RunRoutes()

}
