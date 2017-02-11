package services

import (
	"fmt"
	"github.com/tolivb/scf/pkg/config"
	"net"
	"net/url"
)

type Services interface {
	Start() error
}

type Syslog struct {
	cfg  *config.Config
	conn *net.UDPConn
	addr string
	net  string
}

func (s *Syslog) Start() error {

	if s.conn != nil {
		return fmt.Errorf("Service already started on addr  %s", s.addr)
	}

	udpAddr, err := net.ResolveUDPAddr(s.net, s.addr)
	s.cfg.Log.Debug("Resolved syslog addr %", udpAddr)

	if err != nil {
		return err
	}

	conn, err := net.ListenUDP(s.net, udpAddr)
	s.conn = conn
	s.cfg.Log.Debug("Syslog listen started %v", conn)
	if err != nil {
		return err
	}

	s.cfg.Log.Debug("Syslog add WG %v", conn)

	go func(s *Syslog) {
		defer s.conn.Close()
		defer s.cfg.Wg.Done()
		rcvBuff := make([]byte, 1400)
		s.cfg.Log.Debug("Syslog buffer created %s", "")

		for {
			n, _, err := s.conn.ReadFromUDP(rcvBuff)

			if err != nil {
				s.cfg.Log.Error("Read from %s failed: %s", err, s.addr)
			} else {
				s.cfg.Log.Info("%d - %s ", n, rcvBuff[0:n])
			}
		}
	}(s)

	return nil
}

func New(cfg *config.Config) (Services, error) {
	if cfg.ListenAddr == "" {
		err := fmt.Errorf("%s", "Listen address is empty")
		return nil, err
	}

	addr, err := url.Parse(cfg.ListenAddr)

	if err != nil {
		return nil, err
	}

	switch addr.Scheme {
	case "udp":
		return &Syslog{
			cfg:  cfg,
			net:  addr.Scheme,
			addr: addr.Host,
		}, nil
	}

	return nil, fmt.Errorf("Unsupported schema `%s`", addr.Scheme)
}
