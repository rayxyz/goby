package conf

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Basic :
type Basic struct {
	MapperTag string                 `yaml:"mapper_tag"`
	Extra     map[string]interface{} `yaml:"extra"`
}

// Redis :
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// DSN :
type DSN struct {
	DBDriver string `yaml:"db_driver"`
	DB       string `yaml:"db"`
}

// HTTPEndpoint :
type HTTPEndpoint struct {
	Host          string   `yaml:"host"`
	Port          int      `yaml:"port"`
	RoutePrefixes []string `yaml:"route_prefixes"`
}

// Service :
type Service struct {
	Name         string                 `yaml:"name"`
	DSN          *DSN                   `yaml:"dsn"`
	HTTPEndpoint *HTTPEndpoint          `yaml:"http_endpoint"`
	Extra        map[string]interface{} `yaml:"extra"`
}

// Conf :
type Conf struct {
	Basic    *Basic     `yaml:"basic"`
	Redis    *Redis     `yaml:"redis"`
	Services []*Service `yaml:"services"`
}

var conf Conf
var serviceMap = make(map[string]*Service)

func init() {
	data, err := ioutil.ReadFile("/goby/conf/default.yaml")
	// fmt.Print("The content of the service config file => \n", string(data))
	if err != nil {
		log.Error("Read configuration file error => ", err)
		log.Panic("Read configuration file error => ", err)
	}
	if err = yaml.Unmarshal(data, &conf); err != nil {
		log.Error("Unmarshal configuration file error => ", err)
		log.Panic("Unmarshal configuration file error => ", err)
	}

	for _, v := range conf.Services {
		serviceMap[v.Name] = v
	}
}

// GetConf :
func GetConf() *Conf {
	return &conf
}

// GetBasic :
func GetBasic() *Basic {
	return conf.Basic
}

// GetRedis :
func GetRedis() *Redis {
	return conf.Redis
}

// GetService :
func GetService(service string) *Service {
	svcconf, ok := serviceMap[service]
	if !ok {
		log.Fatal("service conf " + service + " not found")
	}
	return svcconf
}

// GetAllServiceList :
func GetAllServiceList() []*Service {
	if len(conf.Services) > 0 {
		return conf.Services
	}

	return nil
}

// GetAllServiceMap :
func GetAllServiceMap() map[string]*Service {
	if len(serviceMap) > 0 {
		return serviceMap
	}

	return nil
}
