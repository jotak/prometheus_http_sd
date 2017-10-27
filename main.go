package main

import (
	"encoding/json"
	"io/ioutil"
  "log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Targets struct {
	Targets   []string   `json:"targets"`
}

var (
	targetsMap = make(map[string]bool)
	fileSdPath = "/prometheus/targets.json"
)

func ReadTargetsFile() {
	file, err := ioutil.ReadFile(fileSdPath)
	if err != nil {
		return
	}
	var targets [1]Targets
	err = json.Unmarshal(file, &targets)
	if err != nil {
		panic(err)
	}
	for _, host := range targets[0].Targets {
		targetsMap[host] = true
	}
}

func WriteTargetsFile() error {
	targets := []string{}
	for host, _ := range targetsMap {
		targets = append(targets, host)
	}
	fileSD := [1]Targets{Targets{targets}}
	b, err := json.Marshal(fileSD)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileSdPath, b, 0644)
}

func CreateTarget(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	targetsMap[params["host"]] = true
	err := WriteTargetsFile()
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
	}
}

func DeleteTarget(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	delete(targetsMap, params["host"])
	err := WriteTargetsFile()
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
	}
}

func main() {
	if len(os.Args) > 1 {
		fileSdPath = os.Args[1]
	}
	log.Print("Using file_sd path: " + fileSdPath)
	ReadTargetsFile()
	router := mux.NewRouter()
	router.HandleFunc("/target", CreateTarget).Queries("host", "{host}").Methods("PUT")
	router.HandleFunc("/target", DeleteTarget).Queries("host", "{host}").Methods("DELETE")
  log.Fatal(http.ListenAndServe(":9091", router))
}
