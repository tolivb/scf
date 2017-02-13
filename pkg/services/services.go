package services

import (
	"fmt"

	"net"
	"net/url"

	"github.com/tolivb/scf/pkg/config"
)

// Services interface
type Services interface {
	Start() error
}

// Syslog service responsible for receiving log messages
type Syslog struct {
	cfg    *config.Config
	conn   *net.UDPConn
	status *Status
	addr   string
	net    string
}

// Start implements Services interface
func (s *Syslog) Start() error {

	if s.conn != nil {
		return fmt.Errorf("Syslog service already started on %s", s.addr)
	}

	udpAddr, err := net.ResolveUDPAddr(s.net, s.addr)

	if err != nil {
		return fmt.Errorf("Syslog unable to resolve listen addr %s: %s", s.addr, err)
	}

	s.cfg.Log.Debug("Syslog listen addr resolved to %", udpAddr)

	conn, err := net.ListenUDP(s.net, udpAddr)

	if err != nil {
		return err
	}

	s.conn = conn
	s.cfg.Log.Info("Syslog rcv service started on %s://%s", s.net, s.addr)

	go func(s *Syslog) {
		defer s.conn.Close()
		defer s.cfg.Wg.Done()
		rcvBuff := make([]byte, 2048)
		s.cfg.Log.Debug("Syslog buffer[2048] created %s")

		for {
			n, _, err := s.conn.ReadFromUDP(rcvBuff)
			if err != nil {
				s.cfg.Log.Error("Syslog read from %s failed with: %s", s.addr, err)
			} else {
				s.cfg.Log.Info("%d - %s ", n, rcvBuff[0:n])
				s.status.NewEvent("syslog", EINCIN, nil)
			}
		}
	}(s)

	return nil
}

// NewLogReceiver used to create a new listen service that receives log messages(syslog)
func NewLogReceiver(cfg *config.Config, status *Status) (Services, error) {
	if cfg.ListenAddr == "" {
		return nil, fmt.Errorf("%s", "Listen address is empty")
	}

	addr, err := url.Parse(cfg.ListenAddr)

	if err != nil {
		return nil, err
	}

	if addr.Scheme != "udp" {
		return nil, fmt.Errorf("Unsupported schema `%s`", addr.Scheme)
	}

	return &Syslog{
		cfg:    cfg,
		net:    addr.Scheme,
		addr:   addr.Host,
		status: status,
	}, nil
}
