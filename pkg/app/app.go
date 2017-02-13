package app

import (
	"fmt"

	"github.com/tolivb/scf/pkg/config"
	"github.com/tolivb/scf/pkg/services"
)

//App struct
type App struct {
	cfg *config.Config
}

//Run application entry point
func (a *App) Run() int {
	a.cfg.Log.Info("%s - %s(ver %s)", "App started", a.cfg.AppName, a.cfg.Ver)
	a.cfg.Log.Debug("%v", a.cfg)

	//Status service (http)
	statusHttp, err := services.NewStatusHTTP(a.cfg)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	err = statusHttp.Start()
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	a.cfg.Wg.Add(1)

	//MSG receiver
	logReceiver, err := services.NewLogReceiver(a.cfg, statusHttp)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	err = logReceiver.Start()
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	a.cfg.Wg.Add(1)

	//create the communication channels

	//check and start the filters

	//check and start the msg listener

	//check and start the status service
	a.cfg.Wg.Wait()
	return 0

}

func (a *App) startMsgReceiver() {

}

//New create new App
func New(cfg *config.Config) *App {
	a := &App{
		cfg: cfg,
	}

	return a
}
