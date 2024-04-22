package main

import (
	"HelloWorld/src/person"
	"HelloWorld/src/serving"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

func main() {
	homeScreen := fetchPersonInfo()
	runServer(homeScreen)
}

func fetchPersonInfo() string {
	v, err := createViper("person")
	if err != nil {
		log.Fatalf("could not create viper: %v", err)
	}

	var p person.Person
	if err := v.Unmarshal(&p); err != nil {
		log.Fatalf("could not unmarshal person: %v", err)
	}

	return fmt.Sprintf("%v's girlfriend is %v\n\nThis is the new version", p.Name, p.Family.Girlfriend.Name)
}

func runServer(homeScreen string) {
	v, err := createViper("main")
	if err != nil {
		log.Fatalf("failed initializing viper: %v", err)
	}

	var serverConfig serving.ServerConfig
	if err := v.Unmarshal(&serverConfig); err != nil {
		log.Fatalf("failed unmarshalling server config: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, homeScreen)
	})

	port := fmt.Sprintf(":%v", serverConfig.Port)
	fmt.Println("Server listening on port: ", serverConfig.Port)
	fmt.Println("Redis address is: ", serverConfig.RedisConfig.Address)

	http.ListenAndServe(port, nil)
}

func createViper(configName string) (*viper.Viper, error) {
	v := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config/")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
