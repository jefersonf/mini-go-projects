package postgres

import (
	"errors"
	"fmt"
)

const driverName = "postgres"

type ConfigOption func(*Config) error

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

func NewConfig() Config {
	return Config{
		host:     "localhost",
		port:     5432,
		user:     "boss",
		password: "s3cr3t",
		dbname:   "postgres",
		sslmode:  "disable",
	}
}

func (c *Config) FormatDSN() string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s"+
		"password=%s dbname=%s sslmode=%s",
		c.host, c.port, c.user, c.password, c.dbname, c.sslmode)
	return connStr
}

func WithHost(host string) ConfigOption {
	return func(c *Config) error {
		c.host = host
		return nil
	}
}

func WithPort(port int) ConfigOption {
	return func(c *Config) error {
		if port <= 0 {
			return errors.New("invalid port number")
		}
		c.port = port
		return nil
	}
}

func WithUser(user string) ConfigOption {
	return func(c *Config) error {
		c.user = user
		return nil
	}
}

func WithPassword(password string) ConfigOption {
	return func(c *Config) error {
		c.password = password
		return nil
	}
}

func WithDatabase(dbname string) ConfigOption {
	return func(c *Config) error {
		c.dbname = dbname
		return nil
	}
}

func WithSSLDisabled() ConfigOption {
	return func(c *Config) error {
		c.sslmode = "disable"
		return nil
	}
}
