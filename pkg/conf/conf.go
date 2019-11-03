package conf

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	env := initEnv()
	configFile := fmt.Sprintf("config/config.%v.yml", env)
	viper.SetConfigFile(configFile)

	// Allow env vars to use underscores for periods
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
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

	fmt.Printf("Environment: %v\n", env)

	return env
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
