package main

import "github.com/tolivb/scf/pkg/scflog"

// Config container
type Config struct {
	zqmAggrAddr  string //where to send the aggregated data
	sendInterval int    //send aggregated data every sendInterval seconds
	listenAddr   string //where to receive messages proto://addr:port to listen to
	logFS        string //field separator for the received logs
	relayAddr    string //where to relay the received messages proto://addr:port
	maxQueueLen  int    //max number of received but unread messages(separate queue for every filter)
	log          *scflog.Logger
}

func main() {
	l := scflog.New(scflog.DEBUG)

	l.Err("%s", "alabala")
	l.Dbg("%s", "alabala1")
	l.Msg("%s", "alabala1")

	l.Level(scflog.ERROR)

	l.Dbg("%s", "alabala22")
	l.Err("%s", "alabala23", "aalall")
}
