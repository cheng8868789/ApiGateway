package apiConfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
	"strconv"
)

type ApiConfig struct {
	//server
	Host string
	Port string
	Path string

	//超时
	TimeoutSeconds int

	//熔断

	//限流

	//redis

	//kafka

	//url网关
	UrlData UrlData
}

type UrlData struct {
	Urls []Url `json:"urls"`
}

type Url struct {
	Url         string `json:"url"`
	ServiceName string `json:"serviceName"`
}

var config *ApiConfig
var configPath string = "./apiConfig/appConfig.ini"
var urlJsonPath string = "./apiConfig/url.json"

func GetConfig() *ApiConfig {
	return config
}

func GetAddr() string {
	addr := config.Host + ":" + config.Port
	return addr
}

func Init() error {
	config = new(ApiConfig)
	/*if  !isExist() {
		return errors.New("the config of application is not exist")
	}*/
	cfg, err := ini.Load(configPath)
	if err != nil {
		return err
	}
	err = config.getServer(cfg)
	if err != nil {
		return err
	}

	err = config.getUrls()
	if err != nil {
		return err
	}

	fmt.Println("the config of application load success ......")
	return nil

}

func isExist() bool {
	_, err := os.Stat(configPath)
	if err != nil {
		return false
	}
	return true
}

func (cfg *ApiConfig) getServer(configFile *ini.File) error {
	server, err := configFile.GetSection("server")
	if err != nil {
		return err
	}
	cfg.Host = setDefaultString(server, "host", "127.0.0.1")
	cfg.Port = setDefaultString(server, "port", "23008")
	cfg.Path = setDefaultString(server, "path", "apiTest")
	return nil
}

func (cfg *ApiConfig) getTimeout(configFile *ini.File) error {
	server, err := configFile.GetSection("timeout")
	if err != nil {
		return err
	}
	cfg.TimeoutSeconds = setDefaultInt(server, "timeoutSeconds", 30)

	return nil

}

func setDefaultString(section *ini.Section, key string, defaultValue string) string {
	isExist := section.HasKey(key)
	if !isExist {
		return defaultValue
	}
	return section.Key(key).Value()
}

func setDefaultInt(section *ini.Section, key string, defaultValue int) int {
	isExist := section.HasKey(key)
	if !isExist {
		return defaultValue
	}
	intValue, err := strconv.Atoi(section.Key(key).Value())
	if err != nil {
		return defaultValue
	}
	return intValue
}

func (cfg *ApiConfig) getUrls() error {
	urlDataJson, err := ioutil.ReadFile(urlJsonPath)
	if err != nil {
		errorMsg := "load urlJson error :" + err.Error()
		return errors.New(errorMsg)
	}
	urlData := new(UrlData)
	err = json.Unmarshal(urlDataJson, urlData)
	if err != nil {
		return errors.New("json unmarshal error")
	}
	cfg.UrlData = *urlData
	return nil
}
