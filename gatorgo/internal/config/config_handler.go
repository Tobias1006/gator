package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	dir, errDir := os.UserHomeDir()
	if errDir != nil {
		fmt.Print("No Home-Directory found.")
	}
	path := dir + "/" + configFileName
	return path, errDir
}

func write(cfg Config) error {
	path, errPath := getConfigFilePath()
	if errPath != nil {
		return errPath
	}
	data, errMarshal := json.Marshal(cfg)
	if errMarshal != nil {
		return errMarshal
	}
	errWriting := os.WriteFile(path, data, 0666)
	if errWriting != nil {
		fmt.Print("Could not marshal config.")
	}
	return nil
}

func Read() *Config {
	var newConfig Config
	path, err := getConfigFilePath()
	if err != nil {
		fmt.Print(err)
		return &newConfig
	}
	reader, errReading := os.ReadFile(path)
	if errReading != nil {
		fmt.Print("Could not read file.")
		return &newConfig
	}
	errUnmarshal := json.Unmarshal(reader, &newConfig)
	if errUnmarshal != nil {
		fmt.Print("Could not unmarshal file.")
		return &newConfig
	}
	return &newConfig
}

func (c Config) SetUser(username string) error {
	c.Current_user_name = username
	errWriting := write(c)
	if errWriting != nil {
		return errWriting
	}
	return nil
}
