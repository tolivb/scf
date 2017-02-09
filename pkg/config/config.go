package config

import (
	"github.com/tolivb/scf/pkg/scflog"
)

//Config container
type Config struct {
	ZmqAggrAddr  string //where to send the aggregated data
	SendInterval int    //send aggregated data every sendInterval seconds
	ListenAddr   string //where to receive messages proto://addr:port to listen to
	StatusAddr   string //where to check the current status
	LogFS        string //field separator for the received logs
	RelayAddr    string //where to relay the received messages proto://addr:port
	MaxQueueLen  int    //max number of received but unread messages(separate queue for every filter)
	AppName      string
	Ver          string
	Log          scflog.Logger
}

func New() *Config {
	return &Config{}
}
