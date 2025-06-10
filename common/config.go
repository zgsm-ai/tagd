package common

import (
	"os"

	"gopkg.in/yaml.v2"
)

/*
 * Database connection configuration
 * @param Type Database type (mysql/sqlite)
 * @param DatabaseName Database name
 * @param Host Database host
 * @param Port Database port
 * @param Password Database password
 * @param User Database username
 */
type DbConfig struct {
	Type         string `yaml:"type"` // Database type, mysql or sqlite
	DatabaseName string `yaml:"databaseName"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Password     string `yaml:"password"`
	User         string `yaml:"user"`
}

type ServerConfig struct {
	ListenAddr string `yaml:"listenAddr"`
}

/*
 * Main system configuration
 * @param Env Environment identifier
 * @param Db Database configuration
 * @param Redis Redis configuration
 * @param Timeout Timeout settings
 * @param WeChat WeChat notification configuration
 * @param LokiURL Loki logging service URL
 * @param Priority Task priority configuration
 */
type Config struct {
	Env    string       `yaml:"env"`
	Db     DbConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
}

/*
 * Initialize configuration
 * @param filePath Configuration file path
 * @return error Error object
 */
func (config *Config) Init(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	return nil
}
