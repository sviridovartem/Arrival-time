package main

import (
	"os"

	core "svyrydovartem.com/lunar-colony-arrival-time/internal/app/backend"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func start(context *cli.Context) error {
	return core.Main(context)
}

func createApp() *cli.App {
	app := cli.NewApp()
	app.Name = core.ApplicationName
	app.Usage = core.ApplicationUsage
	app.Action = start
	app.Flags = core.SetupAppFlags()

	return app
}

func main() {
	log.SetLevel(log.InfoLevel)
	if err := createApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
