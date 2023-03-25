package main

import (
	"log"
	"net/http"

	_ "github.com/piliphulko/restapi-txt/pkg/datatxt"
	"github.com/piliphulko/restapi-txt/pkg/httpout"
	_ "github.com/piliphulko/restapi-txt/pkg/privacy"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
	viper.SetConfigFile("../config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if !viper.IsSet("localhost") {
		log.Fatalln("localhost does not have config")
	}
}

func main() {
	http.ListenAndServe(viper.GetString("localhost"), httpout.StartRouter())
}
