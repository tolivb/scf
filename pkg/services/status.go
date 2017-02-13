package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/tolivb/scf/pkg/config"
)

const (
	ECLEAR = iota
	EINCIN
	EINCOUT
	EERR
)

type statusMsg struct {
	src   string
	etype int
	err   error
}

type Status struct {
	cfg         *config.Config
	conn        *http.Server
	addr        string
	in          chan statusMsg
	srvDone     chan struct{}
	inMsgCount  int               `json`
	outMsgCount int               `json`
	errors      map[string]string `json`
	mutex       *sync.Mutex
}

func (s *Status) Start() error {
	if s.conn != nil {
		return fmt.Errorf("Status service already started on %s", s.addr)
	}

	s.cfg.Log.Info("%s addr", s.cfg.StatusAddr)
	addr, err := url.Parse(s.cfg.StatusAddr)

	if err != nil {
		return err
	}

	s.addr = addr.Host
	s.conn = &http.Server{Addr: s.addr, Handler: s}
	s.cfg.Log.Info("%s addr1", s.addr)

	go s.handleEvents()
	go s.listenAndServe()

	s.cfg.Log.Info("HTTP status service started on %s", s.addr)
	return nil
}

func (s *Status) listenAndServe() error {
	defer s.cfg.Wg.Done()
	defer func() { s.srvDone <- struct{}{} }()
	err := s.conn.ListenAndServe()

	if err != nil {
		s.cfg.Log.Error("%s", err)
		return err
	}

	s.cfg.Log.Info("%s", "Status service exited")
	return nil
}

func (s *Status) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	data := struct {
		Accepts int               `json:"accepts"`
		Handled int               `json:"handled"`
		Errors  map[string]string `json:"errors"`
	}{
		s.inMsgCount,
		s.outMsgCount,
		s.errors,
	}

	b, e := json.Marshal(data)
	s.mutex.Unlock()

	s.cfg.Log.Debug("New HTTP request from %s for %s ----- %s", r.RemoteAddr, r.RequestURI, e)
	fmt.Fprintf(w, "Hi there %s - %s", r.URL.Path[1:], string(b))
}

func (s *Status) NewEvent(src string, etype int, err error) {
	s.in <- statusMsg{
		src:   src,
		etype: etype,
		err:   err,
	}

	s.cfg.Log.Debug("%s", "Registered event from %s: %s %v", src, etype, err)
}

func (s *Status) handleEvents() {
	s.cfg.Log.Debug("%s", "HTTP status events handler started")
	for m := range s.in {

		s.cfg.Log.Debug("%s", "Received event from %s: %s %v", m.src, m.etype, m.err)
		s.mutex.Lock()

		if m.etype == ECLEAR {
			delete(s.errors, m.src)
			s.mutex.Unlock()
			continue
		}

		if m.err != nil {
			s.errors[m.src] = m.err.Error()
			s.mutex.Unlock()
			continue
		}

		if m.etype == EINCIN {
			s.inMsgCount++
			s.mutex.Unlock()
			continue
		}

		if m.etype == EINCOUT {
			s.outMsgCount++
		}
		s.mutex.Unlock()
	}
}

func NewStatusHTTP(cfg *config.Config) (*Status, error) {
	in := make(chan statusMsg, 3)
	done := make(chan struct{})
	errs := make(map[string]string)

	return &Status{
		cfg:     cfg,
		in:      in,
		srvDone: done,
		errors:  errs,
		mutex:   &sync.Mutex{},
	}, nil
}
