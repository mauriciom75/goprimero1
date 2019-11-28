package main

import (
	"fmt"
	"io/ioutil"
	"os"

	//"log"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	//"github.com/snowzach/rotatefilehook"
	"github.com/spf13/viper"
)

func homeLink2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getApi(w http.ResponseWriter, r *http.Request) {
	//eventID := mux.Vars(r)["id"]
	log.Info("Invocando getApi")

	response, err := http.Get(C.Api1)

	if err != nil {
		fmt.Printf("Error en invocación: %s", err)
		return
	}

	data, _ := ioutil.ReadAll(response.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	log.Info("fin Invocando getApi")

}

type config struct {
	Api1 string
	Name string
}

var C config

func main() {

	log.Info("Iniciando..")

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	viper.SetConfigName("config")
	//viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/tmp/src")
	viper.AddConfigPath("~")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file. %s", err)
		return
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("No puedo leer la configuración. %s", err)
		return
	}
	log.Infof("Configuracion: %+v\n", C)

	//        initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink2)

	router.HandleFunc("/api/{id}", getApi).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
