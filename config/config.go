package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

var c Config

type Config struct {
	Enums []Enum `mapstructure:"enums"`
}

func NewConfig(path string) Config {
	viper.SetConfigFile(path)
	viper.AllowEmptyEnv(true)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		errorLoadConfig(err)
	}

	// Find a *.local.yml file and merge
	localPathSlice := strings.Split(path, ".")
	idx := len(localPathSlice) - 1
	localPathSlice = append(localPathSlice[:idx], append([]string{"local"}, localPathSlice[idx:]...)...)
	path = strings.Join(localPathSlice, ".")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		viper.SetConfigFile(path)
		if err := viper.MergeInConfig(); err != nil {
			errorLoadConfig(err)
		}
	}

	if err := viper.Unmarshal(&c, viper.DecodeHook(configDecodeFunc())); err != nil {
		errorLoadConfig(err)
	}

	return c
}

func configDecodeFunc() mapstructure.DecodeHookFuncValue {
	// Wrapped in a function call to add optional input parameters (eg. separator)
	return func(
		f reflect.Value, // data type
		t reflect.Value, // target data type
	) (interface{}, error) {
		data := f.Interface()
		if data == nil {
			return data, nil
		}

		dataStr, isDataStr := data.(string)
		if !isDataStr {
			dataStr = fmt.Sprintf(`%v`, data)
		}

		tI := t.Addr().Interface()
		var err error

		if v, ok := tI.(*decimal.Decimal); ok {
			*v, err = decimal.NewFromString(dataStr)
			return *v, err
		}

		if v, ok := tI.(*time.Duration); ok {
			*v, err = time.ParseDuration(dataStr)
			return *v, err
		}

		return data, nil
	}
}

func GetConfig() Config {
	return c
}

func errorLoadConfig(err error) {
	log.Fatalf("[Config][Error] cannot read config file: %+v\n", err)
}
