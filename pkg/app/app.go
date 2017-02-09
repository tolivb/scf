package app

import (
	"github.com/tolivb/scf/pkg/config"
	"net/url"
)

type App struct {
	cfg *config.Config
}

func (a *App) Run() {
	a.cfg.Log.Debug("%s - %s(ver %s)", "App started", a.cfg.AppName, a.cfg.Ver)

	l, err := url.Parse(a.cfg.ListenAddr)

	a.cfg.Log.Debug("%v - %v - %v", l.Scheme, l.Host, err)

	//create the communication channels

	//check and start the filters

	//check and start the msg listener

	//check and start the status service

}

func (a *App) startMsgReceiver() {

}

func New(cfg *config.Config) *App {
	a := &App{
		cfg: cfg,
	}

	return a
}
