package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/spf13/viper"
)

type config struct {
	Server struct {
		Host string
		Port string
	}
	AuthClient struct {
		Host string
		Port string
	}
	Casbin struct {
		Model  string
		Policy string
	}
	JWT struct {
		SigningKey string
		Expire     int
		RefreshKey string
		RExpire    int
	}
	Environment string
}

var C config

func init() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.AutomaticEnv()
	viper.Set("environment", os.Getenv("ENVIRONMENT"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
