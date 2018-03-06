Automate copy pasting configuration package by code generation.

Example
JSON:
{
    "bd": {
        "b": 1,
        "d":"test"
    },
    "cf": {
        "c": 1,
        "f":"test"
    },
    "server_address": "localhost:7473",
    "log_max_size": 5,
    "log_max_age": 5,
    "log_max_backups": 5,
    "db_user": "nomane",
    "db_password": "seecrets",
    "db_name": "sdpb",
    "db_host": "db-host",
    "db_port": "5432",
    "auth_url": "http://holdplacer.nya"
}
Package:
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	DbUser string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName string `json:"db_name"`
	LogMaxSize float64 `json:"log_max_size"`
	LogMaxAge float64 `json:"log_max_age"`
	LogMaxBackups float64 `json:"log_max_backups"`
	DbHost string `json:"db_host"`
	DbPort string `json:"db_port"`
	AuthUrl string `json:"auth_url"`
	Cf CfStruct `json:"cf"`
	Bd BdStruct `json:"bd"`
}

type CfStruct struct {
	C float64 `json:"c"`
	F string `json:"f"`
}
type BdStruct struct {
	B float64 `json:"b"`
	D string `json:"d"`
}

var C Config

func Init() error {
	data, err := ioutil.ReadFile("/home/teros0/go/src/conf2go/config.json")
	if err != nil {
		return fmt.Errorf("InitConfig -> couldn't read config file /home/teros0/go/src/conf2go/config.json -> %s", err)
	}
	if err = json.Unmarshal(data, &C); err != nil {
		return fmt.Errorf("InitConfig -> couldn't unmarshal json -> %s", err)
	}
	return nil
}
