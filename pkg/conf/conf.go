package conf

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
	log "github.com/sirupsen/logrus"
)

var env string

func init() {
	env = initEnv()
	viper.SetConfigFile(configFile())

	// Allow env vars to use underscores for periods
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Error reading config file. Check to make sure the config/%s.yml file exists.", env)
		panic(err)
	}

}

func MysqlConfig() mysql.Config {

	params := map[string]string{
		"parseTime": "true",
	}

	return mysql.Config{
		User:   viper.GetString("db.user"),
		Passwd: viper.GetString("db.password"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:%v", viper.GetString("db.host"), viper.GetString("db.port")),
		DBName: viper.GetString("db.database"), Params: params,
		AllowNativePasswords: true,
	}
}

func MysqlRootConfig() mysql.Config {

	params := map[string]string{
		"parseTime": "true",
	}

	return mysql.Config{
		User:   viper.GetString("db.rootUser"),
		Passwd: viper.GetString("db.rootPassword"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:%v", viper.GetString("db.host"), viper.GetString("db.port")),
		DBName: viper.GetString("db.rootDatabase"), Params: params,
		AllowNativePasswords: true,
	}

}

func GetString(path string) string {
	return viper.GetString(path)
}

func GetStrings(path string) []string {
	return viper.GetStringSlice(path)
}

// Private

func initEnv() string {
	viper.SetDefault("env", "development")
	_ = viper.BindEnv("env")

	env := viper.GetString("env")
	if !stringInSlice(env, []string{"development", "production", "prod"}) {
		panic("Invalid environment: " + env)
	}

	if env == "prod" {
		env = "production"
	}

	log.Infof("Environment: %v\n", env)

	return env
}

func configFile() string {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), fmt.Sprintf("../../config/%v.yml", env))
	return filepath
}

// UTIL

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
