package config

import "time"

type ServerConfig struct {
	Host string
}

type AppConfig struct {
	Name     string
	LogLevel string
}

type DatabaseConfig struct {
	Hosts                    []string
	Port                     int
	ConnectTimeout           time.Duration
	MaxConnections           int
	ReadTimeout              time.Duration
	KeySpace                 string
	Retries                  int
	DisableInitialHostLookup bool
	User                     string
	Password                 string
}
