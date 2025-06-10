// Code scaffolded by goctl. Safe to edit.
// goctl <no value>

package config

import "fmt"

type Config struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`

	Jwt struct {
		Secret      string `yaml:"Secret"`
		ExpireHours int64  `yaml:"ExpireHours"`
	} `yaml:"Jwt"`

	MySQL struct {
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Database string `yaml:"Database"`
	} `yaml:"MySQL"`

	Redis struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Password string `yaml:"Password"`
	} `yaml:"Redis"`
}

func (c *Config) MySQLDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQL.User, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.Database)
}
