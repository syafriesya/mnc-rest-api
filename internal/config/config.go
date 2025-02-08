package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var Cfg Config
var EnvCfg EnvConfig

type ServerConfig struct {
	Name string `yaml:"SERVER_NAME"`
	Port string `yaml:"SERVER_PORT"`
}

type DbConfig struct {
	Host     string `yaml:"DB_HOST"`
	Port     string `yaml:"DB_PORT"`
	User     string `yaml:"DB_USER"`
	Password string `yaml:"DB_PASSWORD"`
	Name     string `yaml:"DB_NAME"`
	Schema   string `yaml:"DB_SCHEMA"`
}

type EnvConfig struct {
	Environment string       `yaml:"ENV"`
	Tz          string       `yaml:"TZ"`
	Server      ServerConfig `yaml:"server"`
	Db          DbConfig     `yaml:"db"`
}

type Config struct {
	Env EnvConfig `yaml:"env"`
}

var env string

var folder = "variables"

func New() (Config, error) {
	env = os.Getenv("ENV")
	if env == "" || env == "local" {
		env = "local"

		path := fmt.Sprintf("%s/%s.yaml", folder, env)
		data, err := os.ReadFile(path)
		if err != nil {
			return Config{}, err
		}

		Cfg = Config{}
		err = yaml.Unmarshal(data, &Cfg)
		if err != nil {
			return Config{}, err
		}

		loadEnv(&Cfg)
	} else {
		err := bindEnvs(viper.GetViper())
		if err != nil {
			fmt.Println(err)
		}
		flattenConfig(viper.GetViper())
		viper.AutomaticEnv()
		err = viper.Unmarshal(&EnvCfg)
		if err != nil {
			fmt.Println(err)
		}

		Cfg.Env = EnvCfg
	}

	return Cfg, nil
}

func loadEnv(cfg *Config) {
	if env == "local" {
		godotenv.Load()
	}
}

// viper have structure as KEY_CHILDKEY and we want KEY.CHILDKEY
// if KEY_CHILDKEY_CHILDKEY to KEY.CHILDKEYCHILDKEY
func flattenConfig(v *viper.Viper) {
	for _, key := range v.AllKeys() {
		value := v.Get(key)
		newKey := strings.Replace(key, "_", ".", 1)
		newString := strings.Replace(newKey, "_", "", -1)

		v.Set(newString, value)
	}
}

// get all value env in server
func bindEnvs(v *viper.Viper) error {
	var pair []string
	for _, env := range os.Environ() {
		pair = strings.SplitN(env, "=", 2)
		if len(pair) != 2 {
			continue
		}
		err := v.BindEnv(pair[0])
		if err != nil {
			return err
		}
	}
	return nil
}
