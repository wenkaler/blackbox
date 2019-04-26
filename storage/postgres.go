package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const connectRetryDelay = 3 * time.Second // Initial connection retry interval.
const connectTimeout = 30                 // Initial connection timeout.

// Storage ..
type Storage struct {
	DB  *sql.DB
	cfg *Config
}

// Config ...
type Config struct {
	Logger      log.Logger
	ServiceName string

	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLmode  string
}

// New ..
func New(cfg *Config) (*Storage, error) {

	if cfg.DBName == "" {
		return nil, errors.New("empty data base name ")
	}
	if cfg.Host == "" {
		return nil, errors.New("empty host connection ")
	}
	if cfg.Port == "" {
		return nil, errors.New("empty port connection ")
	}
	if cfg.User == "" {
		return nil, errors.New("empty user connection ")
	}
	if cfg.Logger == nil {
		cfg.Logger = log.NewNopLogger()
	}

	s := &Storage{
		cfg: cfg,
	}
	go s.connectionLoop()
	return s, nil
}

func (s *Storage) connectionLoop() {
	for {
		if err := s.connect(); err != nil {
			level.Error(s.cfg.Logger).Log("msg", "failed to connect to postgres", "err", err, "retry_delay", connectRetryDelay)
			time.Sleep(connectRetryDelay)
			continue
		}
		level.Info(s.cfg.Logger).Log("msg", "established postres connection")
		return
	}
}

func (s *Storage) connect() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s fallback_application_name=%s connect_timeout=%v sslmode=%s", s.cfg.Host, s.cfg.Port, s.cfg.User, s.cfg.Password, s.cfg.DBName, s.cfg.ServiceName, connectTimeout, s.cfg.SSLmode)
	level.Info(s.cfg.Logger).Log("msg", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

// Shutdown shuts down the storage connection.
func (s *Storage) Shutdown() error {
	return s.DB.Close()
}

// FTQuery return json. F - function name, T - token.
// use for func:
// remove_project - return void
// get_project
// list_project
// confirm_setting
func (s *Storage) FTQuery(f, token string) (b []byte, err error) {
	err = s.DB.QueryRow("select "+f+"($1)", token).Scan(&b)
	if err != nil {
		return nil, fmt.Errorf("FTQuery: %v", err)
	}
	return b, nil
}

// FTJQuery return json. F - function name, T - token, J - JSON.
// use for func:
// create_project
// update_setting
func (s *Storage) FTJQuery(f, t string, j []byte) (b []byte, err error) {
	err = s.DB.QueryRow("select "+f+"($1,$2)", t, j).Scan(&b)
	if err != nil {
		return nil, fmt.Errorf("FTJQuery: %v", err)
	}
	return b, nil
}

// FTNJQuery return json. F - function name, T - token, N - name project, J - JSON.
// use for func:
// update_project
func (s *Storage) FTNJQuery(f, t, n string, j []byte) (b []byte, err error) {
	err = s.DB.QueryRow("select "+f+"($1,$2,$3)", t, n, j).Scan(&b)
	if err != nil {
		return nil, fmt.Errorf("FTJQuery: %v", err)
	}
	return b, nil
}

// FTVQuery return json. F - function name, T - token, V - Various means that can contain different kind of character values.
// use for func:
// clean_unused_settings - return void
// initial_setting
func (s *Storage) FTVQuery(f, t, v string) (b []byte, err error) {
	err = s.DB.QueryRow("select "+f+"($1,$2)", t, v).Scan(&b)
	if err != nil {
		return nil, fmt.Errorf("FTVQuery: %v", err)
	}
	return b, nil
}
