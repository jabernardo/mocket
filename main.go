package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

func generateHandleFunc(e Endpoint) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(e.Delay) * time.Millisecond)

		for i := 0; i <= len(e.Headers)-1; i++ {
			current := e.Headers[i]
			w.Header().Add(current.Name, current.Value)
		}

		w.WriteHeader(e.Status)

		strBody, ok := e.Body.(string)

		if ok {
			fmt.Fprintf(w, "%s", strBody)
		} else {
			jsonEncoded, _ := json.Marshal(e.Body)
			fmt.Fprintf(w, "%s", string(jsonEncoded))
		}
	}
}

func serveMockAPI(data []byte) {
	mux := http.NewServeMux()

	t := Config{}

	err := yaml.Unmarshal([]byte(data), &t)

	if err != nil {
		log.Fatalf("err %v", err)
	}

	for i := 0; i <= len(t.Endpoints)-1; i++ {
		currentEndpoint := t.Endpoints[i]

		var muxPath string = fmt.Sprintf("%s %s", currentEndpoint.Method, currentEndpoint.Path)
		log.Println("Route Added: ", muxPath)

		mux.HandleFunc(muxPath, generateHandleFunc(t.Endpoints[i]))
	}

	log.Println("Running Server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", "", t.Port), mux))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Missing configuration (*.yml) file.")
	}

	configFile := os.Args[1]
	fullConfigPath, err := filepath.Abs(configFile)

	if err != nil {
		log.Fatalf("Could not get absolute path of configuration: %v", err)
	}

	log.Println("Configuration File Loaded:", fullConfigPath)

	config, err := os.ReadFile(configFile)

	if err != nil {
		log.Fatalf("Could not read configuration: %v", err)
	}

	serveMockAPI([]byte(config))
}
