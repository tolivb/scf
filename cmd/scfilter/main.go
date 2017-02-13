package main

import (
	"flag"

	"github.com/tolivb/scf/pkg/app"
	"github.com/tolivb/scf/pkg/config"
	"github.com/tolivb/scf/pkg/scflog"
)

const appVer = "1.0"
const appName = "SCFilters"

func main() {

	cfg := config.New()
	debug := flag.Bool("D", false, "Print DEBUG messages")
	flag.StringVar(&cfg.ZmqAggrAddr, "zmq", "", "Zmq broker URL")
	flag.IntVar(&cfg.SendInterval, "i", 5, "Send interval in seconds")
	flag.StringVar(&cfg.ListenAddr, "l", "", "Listen URL for messages")
	flag.StringVar(&cfg.StatusAddr, "s", "", "Status url")
	flag.StringVar(&cfg.LogFS, "fs", "``", "Field separator for incomming messages")
	flag.StringVar(&cfg.RelayAddr, "r", "", "Relay URL(where to relay the messages)")

	flag.IntVar(
		&cfg.MaxQueueLen,
		"ql",
		1000,
		"Max number of received but unread messages(separate queue for every filter)",
	)

	flag.Parse()

	if *debug {
		cfg.Log = scflog.New(scflog.DEBUG)
	} else {
		cfg.Log = scflog.New(scflog.ERROR)
	}

	cfg.AppName = appName
	cfg.Ver = appVer

	app := app.New(cfg)
	app.Run()
}
