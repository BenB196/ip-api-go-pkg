package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"ip-api-go-module/ip-api"
	"ip-api-go-module/util"
	"log"
	"os"
)

type Config struct {
	Fields []string `json:"fields"`
	ApiKey string `json:"apiKey"`
	Caching bool `json:"caching"`
	CacheDuration string `json:"cacheDuration"`
}

var configLocation string
var config Config

//Read config file and load config for running service
func ReadConfig()  {
	//Read flags
	flag.StringVar(&configLocation, "config","","config.json location")
	flag.Parse()

	//Check if configLocation is empty
	if configLocation == "" {
		//Assume Defaults (Default Fields, Free API, Caching On)
		config = Config{
			Fields: ip_api.API_FIELD_DEFAULTS,
			ApiKey: "",
			Caching: true,
		}
	} else {
		//Try add open the specified config file
		jsonFile, err := os.Open(configLocation)

		//If config file is specified but not found, throw error and exit
		if err != nil {
			log.Println("Specified config file not found. Config file specified: " + configLocation)
			log.Println("Error Message: " + err.Error())
			os.Exit(1)
		}
		//Read jsonFile if found
		byteValue, err := ioutil.ReadAll(jsonFile)

		//If error on reading config file, exit
		if err != nil {
			log.Println("Error reading config file.")
			log.Println("Error Message: " + err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(byteValue, &config)

		//If error on unmarshal, exit
		if err != nil {
			log.Println("Error on unmarshal of config.")
			log.Println("Error Message: " + err.Error())
			os.Exit(1)
		}

		//Validate specified fields
		if len(config.Fields) == 0 {
			//If fields is empty set default
			config.Fields = ip_api.API_FIELD_DEFAULTS
		} else {
			//If fields is not empty validate
			//Check if fields == all
			if util.ContainString(config.Fields, "all") {
				//If fields == all set to all fields
				config.Fields = ip_api.API_FIELDS
			} else {
				//If fields != all validate fields set.
			}
		}

		//Validate cacheDuration
	}
}