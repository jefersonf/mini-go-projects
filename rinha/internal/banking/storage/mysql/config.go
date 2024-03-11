package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

type ConfigOption func(*mysql.Config) error

func NewConfig() mysql.Config {
	return mysql.Config{
		User:      "boss",
		Passwd:    "s3cr3t",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "bank",
		ParseTime: true,
		Timeout:   10 * time.Second,
	}
}

func WithUser(user string) ConfigOption {
	return func(c *mysql.Config) error {
		c.User = user
		return nil
	}
}

func WithPasswd(password string) ConfigOption {
	return func(c *mysql.Config) error {
		c.Passwd = password
		return nil
	}
}

func WithNet(net string) ConfigOption {
	return func(c *mysql.Config) error {
		c.Net = net
		return nil
	}
}

func WithAddr(addr string) ConfigOption {
	return func(c *mysql.Config) error {
		c.Addr = addr
		return nil
	}
}

func WithDBName(dbname string) ConfigOption {
	return func(c *mysql.Config) error {
		c.DBName = dbname
		return nil
	}
}
