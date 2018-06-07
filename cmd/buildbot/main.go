package main

import (
	"github.com/tstangenberg/buildbot/buildbot"
	"os"
)

var app Application

func init() {
	app = buildbot.NewCli()
}

func main() {
	app.Run(os.Args)
}
