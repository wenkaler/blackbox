package server

import (
	"github.com/go-kit/kit/log"
)

type service interface {
	FT(f, t string) (b []byte, err error)
	FTJ(f, t string, j []byte) (b []byte, err error)
	FTV(f, t, v string) (b []byte, err error)
	FTNJ(f, t, n string, j []byte) (b []byte, err error)
}

// Storage is a persistent service storage.
type Storage interface {
	FTQuery(f, t string) (b []byte, err error)
	FTJQuery(f, t string, j []byte) (b []byte, err error)
	FTVQuery(f, t, v string) (b []byte, err error)
	FTNJQuery(f, t, n string, j []byte) (b []byte, err error)
}

type basicService struct {
	logger  log.Logger
	storage Storage
}

func (s *basicService) FT(f, t string) (b []byte, err error) {
	b, err = s.storage.FTQuery(f, t)
	return
}
func (s *basicService) FTJ(f, t string, j []byte) (b []byte, err error) {
	b, err = s.storage.FTJQuery(f, t, j)
	return
}
func (s *basicService) FTV(f, t, v string) (b []byte, err error) {
	b, err = s.storage.FTVQuery(f, t, v)
	return
}
func (s *basicService) FTNJ(f, t, n string, j []byte) (b []byte, err error) {
	b, err = s.storage.FTNJQuery(f, t, n, j)
	return
}
