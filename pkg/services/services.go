package services

import (
	"github.com/tolivb/scf/pkg/config"
	"net/url",
    "fmt"
)

type Services interface{
    Start() error
}

type Syslog struct {
	cfg *config.Config
    addr string
}

func (s *Syslog) Start() error {
    pc, err := net.ListenUDP("udp", "host:port")
}

func New(service string, cfg *config.Config)  &Syslog, error{
    if a.cfg.ListenAddr == "" {
        err := fmt.Errorf("%s", "Listen address is empty")
        return nil, err
    }
        
    addr, err := url.Parse(a.cfg.ListenAddr)

    if err != nil{
        return nil, err
    }

    
    switch service{
        case "syslog":

            if addr.Scheme != "udp"{
                return nil, fmt.Errorf(
                    "Unsupported scheme %s for service %s", 
                    addr.Scheme
                    service
                )
            }

            return &Syslog{
                cfg: cfg,
                addr: addr.Host
            }
        default:
            return nil, fmt.Errorf("Unsupported service `%s`", service)

    }
}
