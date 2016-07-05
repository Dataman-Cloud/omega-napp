package config

import (
	"bufio"
	"errors"
	log "github.com/Sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	DefaultTimeout     = 24 * 3600
	DefaultHost        = ""
	DefaultPort        = 6500
	DefaultDebugging   = true
	DefaultLogLevel    = "debug"
	DefaultHealthCheck = 60
)

var config Config

func Pairs() Config {
	return config
}

type Config struct {
	NumCPU       int
	Host         string
	TemplatePath string
	Port         uint16
	Debugging    bool
	Cluster      *Cluster
	Log          *LogConfig
	Db           *DbConfig
	Cache        *CacheConfig
	Mq           *MqConfig
	Drone        *DroneConfig
	Harbor       *HarborConfig
	HealthCheck  int
	HostLogDir   string
	Registry     *RegistryConfig
	Chronos      *Chronos
}

type Cluster struct {
	Host   string
	Scheme string
}

type Chronos struct {
	Host     string
	Scheme   string
	Platform string
}

type RegistryConfig struct {
	Host     string
	Port     uint16
	AuthFile string
	Domain   string
}

type LogConfig struct {
	Console   bool
	File      string
	FileNum   int
	FileSize  int
	Level     string
	Formatter string
}

type DbConfig struct {
	User         string
	Password     string
	Host         string
	Port         uint16
	Name         string
	MaxIdleConns int
	MaxOpenConns int
}

type CacheConfig struct {
	Host     string
	Port     uint16
	Password string
	DB       uint16
	PoolSize int
}

type DroneConfig struct {
	Host string
	Port uint16
}

type HarborConfig struct {
	Host string
	Port uint16
}

type MqConfig struct {
	User       string
	Password   string
	Host       string
	Port       uint16
	CbTimeout  int64 //callback timeout in seconds
	QueueTTL   int64
	MessageTTL int64
}

type EnvEntry struct {
	APP_HOST         string `required:"true"`
	APP_PORT         uint16 `required:"true"`
	APP_DEBUGGING    bool   `required:"true"`
	APP_HEALTH_CHECK int    `required:"true"`

	APP_HOST_LOG_DIR  string `required:"true"`
	APP_LOG_CONSOLE   bool   `required:"true"`
	APP_LOG_FILE      string `required:"true"`
	APP_LOG_LEVEL     string `required:"true"`
	APP_LOG_FORMATTER string `required:"true"`
	APP_LOG_FILE_SIZE int    `required:"true"`
	APP_LOG_FILE_NUM  int    `required:"true"`

	APP_DB_USER           string `required:"true"`
	APP_DB_PASSWORD       string `required:"false"`
	APP_DB_HOST           string `required:"true"`
	APP_DB_PORT           uint16 `required:"true"`
	APP_DB_NAME           string `required:"true"`
	APP_DB_MAX_IDLE_CONNS int    `required:"true"`
	APP_DB_MAX_OPEN_CONNS int    `required:"true"`

	APP_CACHE_HOST      string `required:"true"`
	APP_CACHE_PORT      uint16 `required:"true"`
	APP_CACHE_PASSWORD  string `required:"false"`
	APP_CACHE_DB        uint16 `required:"false"`
	APP_CACHE_POOL_SIZE int    `required:"true"`

	APP_MQ_HOST        string `required:"true"`
	APP_MQ_PORT        uint16 `required:"true"`
	APP_MQ_USER        string `required:"true"`
	APP_MQ_PASSWORD    string `required:"true"`
	APP_MQ_CB_TIMEOUT  int64  `required:"true"`
	APP_MQ_QUEUE_TTL   int64  `required:"true"`
	APP_MQ_MESSAGE_TTL int64  `required:"true"`

	APP_DRONE_HOST string `required:"true"`
	APP_DRONE_PORT uint16 `required:"true"`

	APP_HARBOR_HOST string `required:"true"`
	APP_HARBOR_PORT uint16 `required:"true"`

	APP_REGISTRY_HOST      string `required:"true"`
	APP_REGISTRY_PORT      uint16 `required:"true"`
	APP_REGISTRY_AUTH_FILE string `required:"true"`
	APP_REGISTRY_DOMAIN    string `required:"true"`

	APP_CLUSTER_HOST   string `required:"true"`
	APP_CLUSTER_SCHEME string `required:"true"`

	APP_CHRONOS_HOST     string `required:"true"`
	APP_CHRONOS_SCHEME   string `required:"true"`
	APP_CHRONOS_PLATFORM string `required:"true"`
}

func InitConfig(envFile string) *Config {
	loadEnvFile(envFile)

	envEntry := NewEnvEntry()

	config.NumCPU = 1
	config.Host = envEntry.APP_HOST
	config.Port = envEntry.APP_PORT
	config.Debugging = envEntry.APP_DEBUGGING
	config.HealthCheck = envEntry.APP_HEALTH_CHECK
	config.HostLogDir = envEntry.APP_HOST_LOG_DIR

	config.Cluster = &Cluster{}
	config.Cluster.Host = envEntry.APP_CLUSTER_HOST
	config.Cluster.Scheme = envEntry.APP_CLUSTER_SCHEME

	config.Log = &LogConfig{}
	config.Log.Console = envEntry.APP_LOG_CONSOLE
	config.Log.File = envEntry.APP_LOG_FILE
	config.Log.FileNum = envEntry.APP_LOG_FILE_NUM
	config.Log.FileSize = envEntry.APP_LOG_FILE_SIZE
	config.Log.Level = envEntry.APP_LOG_LEVEL
	config.Log.Formatter = envEntry.APP_LOG_FORMATTER

	config.Db = &DbConfig{}
	config.Db.User = envEntry.APP_DB_USER
	config.Db.Password = envEntry.APP_DB_PASSWORD
	config.Db.Host = envEntry.APP_DB_HOST
	config.Db.Port = envEntry.APP_DB_PORT
	config.Db.Name = envEntry.APP_DB_NAME
	config.Db.MaxIdleConns = envEntry.APP_DB_MAX_IDLE_CONNS
	config.Db.MaxOpenConns = envEntry.APP_DB_MAX_OPEN_CONNS

	config.Cache = &CacheConfig{}
	config.Cache.Host = envEntry.APP_CACHE_HOST
	config.Cache.Port = envEntry.APP_CACHE_PORT
	config.Cache.Password = envEntry.APP_CACHE_PASSWORD
	config.Cache.DB = envEntry.APP_CACHE_DB
	config.Cache.PoolSize = envEntry.APP_CACHE_POOL_SIZE

	config.Mq = &MqConfig{}
	config.Mq.User = envEntry.APP_MQ_USER
	config.Mq.Password = envEntry.APP_MQ_PASSWORD
	config.Mq.Host = envEntry.APP_MQ_HOST
	config.Mq.Port = envEntry.APP_MQ_PORT
	config.Mq.CbTimeout = envEntry.APP_MQ_CB_TIMEOUT
	config.Mq.QueueTTL = envEntry.APP_MQ_QUEUE_TTL
	config.Mq.MessageTTL = envEntry.APP_MQ_MESSAGE_TTL

	config.Drone = &DroneConfig{}
	config.Drone.Host = envEntry.APP_DRONE_HOST
	config.Drone.Port = envEntry.APP_DRONE_PORT

	config.Harbor = &HarborConfig{}
	config.Harbor.Host = envEntry.APP_HARBOR_HOST
	config.Harbor.Port = envEntry.APP_HARBOR_PORT

	config.Registry = &RegistryConfig{}
	config.Registry.Host = envEntry.APP_REGISTRY_HOST
	config.Registry.Port = envEntry.APP_REGISTRY_PORT
	config.Registry.AuthFile = envEntry.APP_REGISTRY_AUTH_FILE
	config.Registry.Domain = envEntry.APP_REGISTRY_DOMAIN

	config.Chronos = &Chronos{}
	config.Chronos.Host = envEntry.APP_CHRONOS_HOST
	config.Chronos.Scheme = envEntry.APP_CHRONOS_SCHEME
	config.Chronos.Platform = envEntry.APP_CHRONOS_PLATFORM

	return &config
}

func NewEnvEntry() *EnvEntry {
	envEntry := &EnvEntry{}

	val := reflect.ValueOf(envEntry).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		required := typeField.Tag.Get("required")

		env := os.Getenv(typeField.Name)

		if env == "" && required == "true" {
			exitMissingEnv(typeField.Name)
		}

		var envEntryValue interface{}
		var err error
		valueField := val.Field(i).Interface()
		value := val.Field(i)
		switch valueField.(type) {
		case int:
			envEntryValue, err = strconv.ParseInt(env, 10, 64)
			envEntryValue, _ = envEntryValue.(int)
		case int64:
			envEntryValue, err = strconv.ParseInt(env, 10, 64)
		case int16:
			envEntryValue, err = strconv.ParseInt(env, 10, 16)
			_, ok := envEntryValue.(int64)
			if !ok {
				exitCheckEnv(typeField.Name, err)
			}
			envEntryValue = int16(envEntryValue.(int64))
		case uint16:
			envEntryValue, err = strconv.ParseUint(env, 10, 16)

			_, ok := envEntryValue.(uint64)
			if !ok {
				exitCheckEnv(typeField.Name, err)
			}
			envEntryValue = uint16(envEntryValue.(uint64))
		case uint64:
			envEntryValue, err = strconv.ParseUint(env, 10, 64)
		case bool:
			envEntryValue, err = strconv.ParseBool(env)
		default:
			envEntryValue = env
		}

		if err != nil {
			exitCheckEnv(typeField.Name, err)
		}
		value.Set(reflect.ValueOf(envEntryValue))
	}

	return envEntry
}

func loadEnvFile(envfile string) {
	// load the environment file
	f, err := os.Open(envfile)
	if err == nil {
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}

			if len(line) == 0 {
				continue
			}

			key, val, err := parseln(string(line))
			if err != nil {
				continue
			}

			if len(os.Getenv(strings.ToUpper(key))) == 0 {
				log.Debug("setting env ", strings.ToUpper(key), val)
				err1 := os.Setenv(strings.ToUpper(key), val)
				if err1 != nil {
					log.Error(err1.Error(), strings.ToUpper(key), val)
				}
			}
		}
	}
}

// helper function to parse a "key=value" environment variable string.
func parseln(line string) (key string, val string, err error) {
	line = removeComments(line)
	if len(line) == 0 {
		return
	}
	splits := strings.SplitN(line, "=", 2)

	if len(splits) < 2 {
		err = errors.New("missing delimiter '='")
		return
	}

	key = strings.Trim(splits[0], " ")
	val = strings.Trim(splits[1], ` "'`)
	return

}

// helper function to trim comments and whitespace from a string.
func removeComments(s string) (_ string) {
	if len(s) == 0 || string(s[0]) == "#" {
		return
	} else {
		index := strings.Index(s, " #")
		if index > -1 {
			s = strings.TrimSpace(s[0:index])
		}
	}
	return s
}

func exitMissingEnv(env string) {
	log.Fatalf("program exit missing config for env %s", env)
}

func exitCheckEnv(env string, err error) {
	log.Errorf("Check env %s, %s", env, err.Error())
}
