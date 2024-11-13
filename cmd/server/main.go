package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/homearchbishop/teachermate-auto/internal/controller"
	"github.com/homearchbishop/teachermate-auto/internal/model"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}
}

var config Config

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	model.InitDB()
	defer model.CloseDB()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		controller.WsHandler(w, r)
	})
	http.HandleFunc("/api/*", func(w http.ResponseWriter, r *http.Request) {
		controller.HttpHandler(w, r)
	})

	log.Println("Server started at :" + strconv.Itoa(config.Server.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), nil)
	if err != nil {
		log.Fatalln("ListenAndServe error:", err)
	}
}
