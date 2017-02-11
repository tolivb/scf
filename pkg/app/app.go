package app

import (
	"github.com/tolivb/scf/pkg/config"
	"github.com/tolivb/scf/pkg/services"
)

type App struct {
	cfg *config.Config
}

func (a *App) Run() int {
	a.cfg.Log.Info("%s - %s(ver %s)", "App started", a.cfg.AppName, a.cfg.Ver)

	syslog, err := services.New(a.cfg)

	if err != nil {
		a.cfg.Log.Error("%s", err)
		return 2
	}

	err = syslog.Start()

	if err != nil {
		a.cfg.Log.Error("%s", err)
	} else {
		a.cfg.Wg.Add(1)
	}

	//create the communication channels

	//check and start the filters

	//check and start the msg listener

	//check and start the status service
	a.cfg.Wg.Wait()
	return 0

}

func (a *App) startMsgReceiver() {

}

func New(cfg *config.Config) *App {
	a := &App{
		cfg: cfg,
	}

	return a
}
