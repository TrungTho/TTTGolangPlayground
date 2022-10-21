package settings

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Environment string

const (
	LiveEnv          Environment = "LIVE"
	TestEnv          Environment = "TEST"
 	DevEnv           Environment = "DEV"
	connectionString             = "%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	Env      = DevEnv
	PORT=""
	AES_KEY=""
)

func IsInTest() bool {
	if viper.GetBool("LOCAL_TEST") == true {
		return true
	}

	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.v=") {
			return true
		}
	}
	return false
}
 

func init() {
	setDefaultEnvVariables()
	loadEnvVariables()
	 
 
 	initRedis()
}

func initRedis(){
	//todo: for now just return
	return
	// _, err := redis.GetRedis().Ping( ).Result()
	// if err != nil {
	// 	fmt.Printf("connect redis error %s %d %d", RedisHost, RedisPort, RedisDb)
	// 	panic(err)
	// }
}

func setDefaultEnvVariables() {}

func loadEnvVariables() {
	viper.AutomaticEnv()

	Env = Environment(viper.GetString("ENV"))
	PORT=viper.GetString("PORT")
	AES_KEY=viper.GetString("AES_KEY")
 
}
