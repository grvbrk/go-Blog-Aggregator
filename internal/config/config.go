package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"

func Read() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home dir", err)
		return Config{}
	}

	data, err := os.ReadFile(homeDir + configFileName)
	if err != nil {
		fmt.Println("Error opening .gatorconfig.json", err)
		return Config{}
	}

	config := Config{}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return Config{}
	}

	return config
}

func (c *Config) SetUser(username string) {

	c.CurrentUserName = username

	jsondata, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home dir", err)
		return
	}

	err = os.WriteFile(homeDir+configFileName, jsondata, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
