package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type yamlConf struct {
	Port        int      `yaml:"PORT,omitempty"`
	DSN         string   `yaml:"DSN,omitempty"`
	Mode        string   `yaml:"MODE,omitempty"`
	Docs        string   `yaml:"DOCS,omitempty"`
	Secret      string   `yaml:"SECRET,omitempty"`
	ProfilePath string   `yaml:"PFP_PATH,omitempty"`
	ProfileMax  int64    `yaml:"PFP_MAX,omitempty"`
	Cors        []string `yaml:"CORS,omitempty"`
}

type devResponse struct {
	Mode   string `json:"mode"`
	Front  string `json:"front"`
	Uptime string `json:"uptime"`
}

func createConfig() {
	data, err := ioutil.ReadFile("config.example.yaml")

	if err != nil {
		fmt.Println("Could not find config.example.yaml! Did you clone the repo properly?")
		return
	}

	if err = ioutil.WriteFile("config.yaml", data, 0644); err != nil {
		panic(err)
	}

	os.Exit(0)
}

func loadConfig() yamlConf {
	conf, err := ioutil.ReadFile("config.yaml")
	c := yamlConf{}

	if err != nil {
		fmt.Println("Could not find config.yaml. Creating one for you :)")
		createConfig()
	}

	if err = yaml.Unmarshal(conf, &c); err != nil {
		fmt.Println("Invalid config!")
		panic(err)
	}

	return c
}
