package config

import (
	"fmt"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	config := InitLoadConfig()
	if config.Server.Addr == "" || config.Server.Port == "" {
		t.Error("Failed to read config")
	} else {
		fmt.Println(config.Server.Addr, config.Server.Port)
	}
}
