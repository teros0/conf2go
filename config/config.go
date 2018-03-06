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
