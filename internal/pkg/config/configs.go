package config

import (
	"crypto/tls"
	"school_project/internal/pkg/util"
)

type ConfigDatabase struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	TLSConfig *tls.Config
}

// ToDo: Đang phải set db host là ip của máy
func NewConfigDatabase() *ConfigDatabase {
	return &ConfigDatabase{
		Host:      util.GetEnv("DB_HOST", "0.0.0.0"),
		Port:      util.GetEnv("DB_PORT", "5432"),
		User:      util.GetEnv("DB_USER", "enouvo"),
		Password:  util.GetEnv("DB_PASSWORD", "123qwe"),
		Database:  util.GetEnv("DB_DATABASE", "school"),
		TLSConfig: nil,
	}
}

type PathConfig struct {
	Path       string `json:"path"`
	RouterFile string `json:"routerFile"`
}

func PathSaveFile() *PathConfig {
	return &PathConfig{
		RouterFile: util.GetEnv("ROOT_DIR", ""),
	}
}
